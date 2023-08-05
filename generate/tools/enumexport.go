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
			if m.Name == "UIKit" {
				continue
			}
			dump := exportConstants(r, &m, TargetPlatform, TargetVersion)
			filename := "./generate/modules/enums/" + TargetPlatform + "/" + m.Package
			os.MkdirAll(filepath.Dir(filename), 0755)
			if err := ioutil.WriteFile(filename, dump, 0644); err != nil {
				log.Fatal(err)
			}
		}
	}

}

func exportConstants(db *zip.ReadCloser, framework *modules.Module, platform string, version int) []byte {
	// first get enums for int constants
	var enums []string
	for _, file := range db.File {
		if filepath.Dir(file.Name) == fmt.Sprintf("symbols/%s", framework.Package) {
			s, err := generate.LoadSymbolFrom(file)
			if err != nil {
				continue
			}
			if !s.HasPlatform(platform, version) {
				continue
			}
			if s.Kind != "Enum" {
				continue
			}
			enums = append(enums, s.Name)
		}
	}
	inEnums := func(s string) bool {
		for _, enum := range enums {
			if enum == s {
				return true
			}
		}
		return false
	}

	// now collect types of values
	var constInts []string
	var constStrs []string
	var constFloats []string
	var constFloatPtrs []string
	for _, file := range db.File {
		if filepath.Dir(file.Name) == fmt.Sprintf("symbols/%s", framework.Package) {
			s, err := generate.LoadSymbolFrom(file)
			if err != nil {
				continue
			}
			if !s.HasPlatform(platform, version) {
				continue
			}
			if s.Kind == "Constant" && s.Type == "Enumeration Case" && inEnums(s.Parent) {
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
				if stmt.Variable.Type.Name == "BOOL" ||
					stmt.Variable.Type.Name == "Boolean" {
					// ignore bool globals, prob not constants
					// ex: NSDeallocateZombies, globalUpdateOK
					continue
				}
				if stmt.Variable.Type.Name == "NSFileProviderPage" {
					// other manual ignores from:
					// NSFileProviderInitialPageSortedByName
					continue
				}
				if stmt.Variable.Type.Name == "int" ||
					stmt.Variable.Type.Name == "long long" ||
					stmt.Variable.Type.Name == "long" {
					constInts = append(constInts, s.Name)
					continue
				}
				if stmt.Variable.Type.Name == "NSString" {
					constStrs = append(constStrs, s.Name)
					continue
				}
				if stmt.Variable.Type.Name == "double" ||
					stmt.Variable.Type.Name == "float" ||
					stmt.Variable.Type.Name == "CGFloat" {
					if stmt.Variable.Type.IsPtr {
						constFloatPtrs = append(constFloatPtrs, s.Name)
					} else {
						constFloats = append(constFloats, s.Name)
					}
					continue
				}
				typ := generate.FindSymbolByName(db, stmt.Variable.Type.Name)
				if typ == nil {
					log.Fatalf("unable to find symbol: %s\n", stmt.Variable.Type.Name)
				}
				if typ.Declaration == "" {
					continue
				}
				typdef, err := typ.Parse()
				if err != nil {
					log.Fatalf("%s: %s in '%s'", typ.Name, err, typ.Declaration)
				}
				if typdef.Typedef == "" {
					if typdef.Interface != nil {
						// ignore global object pointers here
						// ex: NSApp
						continue
					}
					if typdef.Struct != nil {
						// ignore weird struct refs here
						// ex: kCFURLFileResourceTypeKey
						continue
					}
					log.Fatalf("%s: declaration statement not a typedef: %s [%s]", typ.Name, typ.Declaration, s.Name)
				}
				if typdef.Enum != nil {
					switch typdef.Enum.Type.Name {
					case "NSUInteger", "NSInteger", "long long", "short":
						constInts = append(constInts, s.Name)
						continue
					default:
						log.Fatalf("%s: unsupported enum type: %s in '%s' [%s]", typ.Name, typdef.Enum.Type.Name, typ.Declaration, s.Name)
					}
				}
				if typdef.TypeAlias == nil {
					if typdef.Struct != nil {
						// ignore global object pointers here for now
						// ex: NSMultipleValuesMarker
						continue
					}
					log.Fatalf("%s: unsupported typedef: %s [%s]", typ.Name, typ.Declaration, s.Name)
				}
				switch typdef.TypeAlias.Name {
				case "NSString", "CFStringRef":
					constStrs = append(constStrs, s.Name)
					continue
				case "float", "double", "CGFloat":
					constFloats = append(constFloats, s.Name)
					continue
				case "NSInteger", "NSUInteger", "long", "long long", "short":
					constInts = append(constInts, s.Name)
					continue
				case "CGSize", "CGPoint", "NSRect", "CGRect":
					// manual ignores for obvious reasons
					continue
				default:
					log.Fatalf("%s: unsupported typedef type: %s in '%s' [%s]", typ.Name, typdef.TypeAlias.Name, typ.Declaration, s.Name)
				}
			}

		}
	}
	sort.Strings(constInts)
	sort.Strings(constStrs)
	sort.Strings(constFloats)
	sort.Strings(constFloatPtrs)

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
	// for _, c := range constFloatPtrs {
	// 	constPrintfs = append(constPrintfs, fmt.Sprintf("	printf(\"%s %%f\\n\", (float)*%s);\n", c, c))
	// }
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
