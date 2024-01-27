package generate

import (
	"fmt"
	"log"
	"strings"
	"io"
	"os"

	"github.com/progrium/macdriver/generate/codegen"
	"github.com/progrium/macdriver/generate/declparse"
	"github.com/progrium/macdriver/generate/modules"
	"github.com/progrium/macdriver/generate/typing"
	"github.com/progrium/macdriver/internal/set"
)

type Generator struct {
	*SymbolCache
	Platform      string
	Version       int
	Framework     string
	customTypes   set.Set[string]
	customMethods set.Set[string]
	enumCache     map[string][]Symbol
	genCache      map[string]codegen.CodeGen
}

func (db *Generator) Generate(platform string, version int, rootDir string, framework string, ignoreTypes set.Set[string]) {
	db.Platform = platform
	db.Version = version
	module := modules.Get(framework)
	db.Framework = module.Package
	modsym := db.ModuleSymbol(*module)

	types, methods := exportedSourceSymbols(fmt.Sprintf("%s/%s/%s_custom.go", rootDir, db.Framework, db.Framework))
	db.customTypes = set.New(types...)
	db.customMethods = set.New(methods...)

	RemoveGeneratedCode(fmt.Sprintf("%s/%s", rootDir, db.Framework))

	mw := &codegen.ModuleWriter{
		Description: modsym.Description,
		DocURL:      modsym.DocURL(),
		Module:      *module,
		PlatformDir: rootDir,
	}
	copySpecialFile(framework)

	for _, s := range db.ModuleSymbols(framework) {
		if ignoreTypes.Contains(s.Name) {
			continue
		}
		if !s.HasPlatform(platform, version, true) {
			continue
		}
		switch s.Kind {
		case "Class":
			if db.customTypes.Contains(modules.TrimPrefix(s.Name)) {
				log.Println("skipping class with custom definition", s.Name)
				continue
			}
			classGen := db.ToClassGen(s)
			if classGen == nil {
				log.Println("skipping class", s.Name)
				continue
			}
			classGen.Init()
			fw := &codegen.FileWriter{
				Name:        s.Name,
				Module:      *classGen.Type.Module,
				PlatformDir: rootDir,
			}
			fw.Add(classGen)
			removeMethods(classGen)
			fw.WriteCode()
		case "Protocol":
			protocolGen := db.ToProtocolGen(framework, s)
			if protocolGen == nil {
				log.Println("skipping protocol", s.Name)
				continue
			}
			if db.customTypes.Contains("P" + modules.TrimPrefix(s.Name)) {
				log.Println("skipping protocol interface with custom definition", s.Name)
				protocolGen.SkipInterface = true
			}
			if db.customTypes.Contains(modules.TrimPrefix(s.Name) + "Wrapper") {
				log.Println("skipping protocol wrapper with custom definition", s.Name)
				protocolGen.SkipWrapper = true
			}
			if db.customTypes.Contains(modules.TrimPrefix(s.Name)) {
				log.Println("skipping protocol delegate with custom definition", s.Name)
				protocolGen.SkipDelegate = true
			}
			protocolGen.Init()
			fw := &codegen.FileWriter{
				Name:        s.Name + "Protocol",
				Module:      *protocolGen.Type.Module,
				PlatformDir: rootDir,
			}
			fw.Add(protocolGen)
			fw.WriteCode()
			mw.Protocols = append(mw.Protocols, protocolGen.Type)
		case "Enum":
			if s.Name == "MTLArgumentAccess" {
				// todo: regen metal without this skip to see why
				continue
			}
			if db.customTypes.Contains(modules.TrimPrefix(s.Name)) {
				log.Println("skipping enum with custom definition", s.Name)
				continue
			}
			mw.EnumAliases = append(mw.EnumAliases, db.ToEnumInfo(framework, s))
		case "Type":
			if db.customTypes.Contains(modules.TrimPrefix(s.Name)) {
				log.Println("skipping type with custom definition", s.Name)
				continue
			}
			if db.AllowedEnumAlias(s) {
				mw.EnumAliases = append(mw.EnumAliases, db.ToEnumInfo(framework, s))
				continue
			}
			st, err := s.Parse(db.Platform)
			if err != nil || st.TypeAlias == nil {
				log.Printf("skipping '%s', bad decl: %s", s.Name, s.Declaration)
				continue
			}
			if st.TypeAlias.Func != nil {
				ftyp := db.ParseType(*st.TypeAlias)
				if ftyp == nil {
					fmt.Printf("skipping: name=%s declaration=%s\n", s.Name, s.Declaration)
					continue
				}
				mw.FuncAliases = append(mw.FuncAliases, &codegen.AliasInfo{
					AliasType: typing.AliasType{
						Name:  s.Name,
						GName: modules.TrimPrefix(s.Name),
						Type:  db.ParseType(*st.TypeAlias),
					},
					Description: s.Description,
					DocURL:      s.DocURL(),
				})
				continue
			}
			// any other type aliases can be added manually
			// since they're just a go type alias or an
			// unsafe.Pointer type
		}
	}
	mw.WriteCode()
	FormatCode(rootDir + "/" + module.Package)
}

func (db *Generator) ResolveTypeAlias(typeName string) (declparse.TypeInfo, bool) {
	s := db.FindTypeSymbol(typeName)
	if s == nil {
		return declparse.TypeInfo{}, false
	}
	st, err := s.Parse(db.Platform)
	if err != nil || st.TypeAlias == nil {
		return declparse.TypeInfo{}, false
	}
	return *st.TypeAlias, true
}

func (db *Generator) TypeFromSymbolName(name string) typing.Type {
	// hardcoded for now
	if db.Platform == "macos" && strings.HasPrefix(name, "UI") && !strings.HasPrefix(name, "UInt") {
		name = "NS" + strings.TrimPrefix(name, "UI")
	}
	s := db.FindTypeSymbol(name)
	if s == nil {
		log.Printf("using NSObject for unknown symbol: %s\n", name)
		return typing.Object
	}
	if m := s.MainModule(); m != nil && modules.CanAbstractModuleCoupling(db.Framework, m.Package) {
		log.Printf("using NSObject for '%s' to decouple '%s'\n", name, m.Package)
		return typing.Object
	}

	return db.TypeFromSymbol(*s)
}

func strIn(slice []string, str string) bool {
	for _, s := range slice {
		if strings.EqualFold(str, s) {
			return true
		}
	}
	return false
}

// used to store information from special.go files
type SkipMethod struct {
	ClassName string
	Selector string
	Note string
}

// the database of methods that should not be generated
var methodSkipList []SkipMethod
var attemptedRead = false // prevents multiple unnecessary read attempts

// reads a file called special.go and creates a database of methods to not generate "skip"
func getFilteredSelectors() []SkipMethod {
	workingPath, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	specialFilePath := workingPath + "/special.go"
	specialFile, err := os.ReadFile(specialFilePath)
	if err != nil {
		return nil
	}
	
	// The special.go file will contain a section that begins with "begin-skip"
	// and end with "end-skip". The information between these two indicators
	// is used to create a database of methods to not generate.
	// The format is "Objective-c class name", "selector", "note to display while generating".
	specialFileContents := string(specialFile)
	start := strings.Index(specialFileContents, "begin-skip") + 10
	end := strings.Index(specialFileContents, "end-skip")
	data := specialFileContents[start:end]
	rows := strings.Split(data, "\n")
	buffer := make([]SkipMethod, 0)
	for _,row := range(rows) {
		if row == "" {
			continue
		}
		row = strings.ReplaceAll(row, "\"", "")
		columns := strings.Split(row, ",")
		newSkipMethod := SkipMethod{}
		newSkipMethod.ClassName = strings.TrimSpace(columns[0])
		newSkipMethod.Selector = strings.TrimSpace(columns[1])
		newSkipMethod.Note = strings.TrimSpace(columns[2])
		buffer = append(buffer, newSkipMethod)
	}
	return buffer
}

// removes methods that are on the list of methods to skip
func removeMethods(theClass *codegen.Class) {

	if methodSkipList == nil && attemptedRead == true {
		return // no need to continue if there is no list
	}
	
	// cache the list
	if methodSkipList == nil && attemptedRead == false {
		methodSkipList = getFilteredSelectors()
		attemptedRead = true
	}

	// find and remove the skippable instance methods
	methodBuffer := make([]*codegen.Method, len(theClass.InstanceTypeMethods))
	copy(methodBuffer, theClass.InstanceTypeMethods)
	for index, method := range(theClass.InstanceTypeMethods) {
		for _, skipMethod := range(methodSkipList) {
			if skipMethod.ClassName == theClass.String() {
				if skipMethod.Selector == method.Selector() {
					fmt.Printf("Removing [%s %s] - %s\n", theClass, skipMethod.Selector, skipMethod.Note)
					//remove the element from the slice
					methodBuffer[index] = nil
				}
			}
		}
	}

	// create a new slice with all the skipped instance methods gone
	theClass.InstanceTypeMethods = nil
	for _, method := range(methodBuffer) {
		if method != nil {
			theClass.InstanceTypeMethods = append(theClass.InstanceTypeMethods, method)
		}
	}

	// find and remove the skippable methods
	methodBuffer = make([]*codegen.Method, len(theClass.Methods))
	copy(methodBuffer, theClass.Methods)
	for index, method := range(theClass.Methods) {
		for _, skipMethod := range(methodSkipList) {
			if skipMethod.ClassName == theClass.String() {
				if skipMethod.Selector == method.Selector() {
					fmt.Printf("Removing [%s %s] - %s\n", theClass, skipMethod.Selector, skipMethod.Note)
					//remove the element from the slice
					methodBuffer[index] = nil
				}
			}
		}
	}

	// create a new slice with all the skipped methods gone
	theClass.Methods = nil
	for _, method := range(methodBuffer) {
		if method != nil {
			theClass.Methods = append(theClass.Methods, method)
		}
	}
}

// Copies the special.go file to the current framework's folder
func copySpecialFile(framework string) {

	// the folder all the frameworks are in is named one of these
	platforms := []string {
		"macos",
		"ios",
		"ipados",
		"tvos",
		"watchos",
		"visionos",
	}

	workingPath, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	sourcePath := workingPath + "/../../generate/special/"

	// Go thru every platform
	for _, platform := range(platforms) {
		sourcePath = sourcePath + platform
		sourcePath = sourcePath + "/" + framework
		sourcePath = sourcePath + "/special.go"
		source, err := os.Open(sourcePath)
		if err != nil {
			continue // platform not available, move on to the next platform
		}

		destPath := workingPath + "/special.go"
		destination, err := os.Create(destPath)

		// copy the file to the framework's folder
		_, err = io.Copy(destination, source)
		if err != nil {
			fmt.Println("Error copying file:", err)
			return
		}
	}
}
