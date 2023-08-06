//go:build ignore

package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/progrium/macdriver/generate"
	"github.com/progrium/macdriver/generate/declparse"
	"github.com/progrium/macdriver/generate/modules"
)

const TargetPlatform = "macos"
const TargetVersion = 12

// go run ./generate/tools/enumexport.go [framework]
func main() {

	r, err := zip.OpenReader("./generate/symbols.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	if len(os.Args) > 1 {
		mod := modules.Get(strings.ToLower(os.Args[1]))
		fmt.Print(string(exportConstants(r, mod, TargetPlatform, TargetVersion)))
	} else {
		for _, m := range modules.All {
			if m.Package == "objc" {
				continue
			}
			dump := exportConstants(r, &m, TargetPlatform, TargetVersion)
			if m.Package == "uikit" && TargetPlatform == "macos" {
				// append to appkit instead
				filename := "./generate/modules/enums/" + TargetPlatform + "/appkit"
				log.Println(filename, "[uikit]")
				f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
				if err != nil {
					log.Fatal(err)
				}
				defer f.Close()
				_, err = f.Write(dump)
				if err != nil {
					log.Fatal(err)
				}
			} else {
				filename := "./generate/modules/enums/" + TargetPlatform + "/" + m.Package
				log.Println(filename)
				os.MkdirAll(filepath.Dir(filename), 0755)
				if err := os.WriteFile(filename, dump, 0644); err != nil {
					log.Fatal(err)
				}
			}

		}
	}

}

func exportConstants(db *zip.ReadCloser, framework *modules.Module, platform string, version int) []byte {

	var constInts []string
	var constStrs []string
	var constFloats []string
	//var constFloatPtrs []string
	addConst := func(typeInfo declparse.TypeInfo, name string) bool {
		switch typeInfo.Name {
		case "int", "long", "long long", "short", "NSUInteger", "NSInteger":
			constInts = append(constInts, name)
			return true
		case "NSString", "CFStringRef":
			constStrs = append(constStrs, name)
			return true
		case "float", "double", "CGFloat":
			if typeInfo.IsPtr {
				// ignore for now
				//constFloatPtrs = append(constFloats, name)
				return true
			}
			constFloats = append(constFloats, name)
			return true
		case "BOOL", "Boolean":
			// ignore bool globals, prob not constants
			// ex: NSDeallocateZombies, globalUpdateOK
			return true
		case "CGSize", "CGPoint", "NSRect", "CGRect":
			// obvious ignores
			return true
		case "NSFileProviderPage":
			// other manual ignores from:
			// NSFileProviderInitialPageSortedByName
			return true
		default:
			return false
		}
	}

	for _, file := range db.File {
		if filepath.Dir(file.Name) == fmt.Sprintf("symbols/%s", framework.Package) {
			s, err := generate.LoadSymbolFrom(file)
			if err != nil {
				continue
			}
			if platform == "macos" && framework.Package == "uikit" {
				// we're going to pretend to be appkit later
				// to handle the uikit symbols existing in appkit
				if !s.HasFramework("appkit") {
					continue
				}
			}
			if !s.HasPlatform(platform, version, true) {
				continue
			}
			// ignore list. these deprecated ones aren't defined for me
			if strIn([]string{
				"QCCompositionInputRSSArticleDurationKey",
				"QCCompositionInputRSSFeedURLKey",
				"QCCompositionProtocolRSSVisualizer",
			}, s.Name) {
				continue
			}
			if s.Kind == "Constant" && s.Type == "Enumeration Case" {
				constInts = append(constInts, s.Name)
				continue
			}
			if s.Kind == "Constant" && s.Type == "Global Variable" {
				stmt, err := s.Parse()
				if err != nil {
					log.Fatalf("%s: %s in '%s'", s.Name, err, s.Declaration)
				}
				if stmt.Variable == nil {
					log.Fatalf("%s: declaration statement not a variable: %s", s.Name, s.Declaration)
				}
				if ok := addConst(stmt.Variable.Type, s.Name); ok {
					continue
				}

				lookupType := func(name string) (ti *declparse.TypeInfo, ok bool) {
					typ := generate.FindSymbolByName(db, name)
					if typ == nil {
						log.Fatalf("unable to find symbol: %s\n", name)
					}
					if typ.Declaration == "" {
						return nil, false
					}
					typdef, err := typ.Parse()
					if err != nil {
						log.Fatalf("%s: %s in '%s'", typ.Name, err, typ.Declaration)
					}
					if typdef.Interface != nil {
						// ignore global object pointers here
						// ex: NSApp
						return nil, false
					}
					if typdef.Struct != nil {
						// ignore weird struct refs here
						// ex: kCFURLFileResourceTypeKey
						// ex: NSMultipleValuesMarker
						return nil, false
					}
					if typdef.Typedef == "" {
						log.Fatalf("%s: declaration statement not a typedef: %s [%s]", typ.Name, typ.Declaration, s.Name)
					}
					if typdef.Enum != nil {
						return &typdef.Enum.Type, true
					}
					if typdef.TypeAlias != nil {
						return typdef.TypeAlias, true
					}
					log.Fatalf("%s: unsupported typedef: %s [%s]", typ.Name, typ.Declaration, s.Name)
					return nil, false
				}

				ti, ok := lookupType(stmt.Variable.Type.Name)
				if !ok {
					continue
				}
				if ok := addConst(*ti, s.Name); !ok {
					// try to resolve one more level
					tti, ok := lookupType(ti.Name)
					if !ok {
						continue
					}
					if ok := addConst(*tti, s.Name); !ok {
						log.Fatalf("%s: unsupported type: %s [%s]", stmt.Variable.Type.Name, ti.Name, s.Name)
					}
				}
			}

		}
	}
	sort.Strings(constInts)
	sort.Strings(constStrs)
	sort.Strings(constFloats)
	// sort.Strings(constFloatPtrs)

	// generate source code to export
	var constPrintfs []string
	for _, c := range constInts {
		constPrintfs = append(constPrintfs, fmt.Sprintf("	printf(\"%s %%ld\\n\", (long)%s);\n", c, c))
	}
	for _, c := range constStrs {
		constPrintfs = append(constPrintfs, fmt.Sprintf("	printf(\"%s %%s\\n\", [%s UTF8String]);\n", c, c))
	}
	for _, c := range constFloats {
		constPrintfs = append(constPrintfs, fmt.Sprintf("	printf(\"%s %%f\\n\", (float)%s);\n", c, c))
	}
	// these can segfault, so leaving them out for now
	// for _, c := range constFloatPtrs {
	// 	constPrintfs = append(constPrintfs, fmt.Sprintf("	printf(\"%s %%f\\n\", (float)*%s);\n", c, c))
	// }

	if platform == "macos" && framework.Package == "uikit" {
		// in this case we will export these out of appkit
		// since uikit does not exist on macos
		framework = modules.Get("appkit")
	}
	source := fmt.Sprintf(`package main

/*
#cgo CFLAGS: -w -x objective-c
#cgo LDFLAGS: -lobjc -framework %s
#include <%s>

void Main() {
%s
}

*/
import "C"

func main() {
	C.Main()
}
`, framework.Name, framework.Header, strings.Join(constPrintfs, ""))

	return evalSource(source)
}

func evalSource(source string) []byte {
	tempFile, err := ioutil.TempFile("", "temp-code-*.go")
	if err != nil {
		log.Fatal("Failed to create temporary file:", err)
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.WriteString(source); err != nil {
		log.Fatal("Failed to write to temporary file:", err)
	}

	if err := tempFile.Close(); err != nil {
		log.Fatal("Failed to close temporary file:", err)
	}

	cmd := exec.Command("go", "run", tempFile.Name())
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(output))
		log.Fatal("Failed to execute command:", err)
	}

	return output
}

func strIn(slice []string, str string) bool {
	for _, s := range slice {
		if strings.EqualFold(str, s) {
			return true
		}
	}
	return false
}
