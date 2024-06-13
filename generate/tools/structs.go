//go:build ignore

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/progrium/macdriver/generate"
	"github.com/progrium/macdriver/generate/modules"
)

const TargetPlatform = "macos"
const TargetVersion = 14

// go run ./generate/tools/structs.go [framework]
func main() {

	db, err := generate.OpenSymbols("./generate/symbols.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if len(os.Args) > 1 {
		mod := modules.Get(strings.ToLower(os.Args[1]))

		var structDefs []string
		var badStructs []string
		docComments := make(map[string]string)

		for _, s := range db.ModuleSymbols(mod.Package) {
			if TargetPlatform == "macos" && mod.Package == "uikit" {
				// we're going to pretend to be appkit later
				// to handle the uikit symbols existing in appkit
				if !s.HasFramework("appkit") {
					continue
				}
			}
			if !s.HasPlatform(TargetPlatform, TargetVersion, true) {
				continue
			}

			if s.Kind == "Struct" {

				structDef := fmt.Sprintf("type %s C.%s\n", modules.TrimPrefix(s.Name), s.Name)

				// generate for just this struct to see that it works
				// before adding so the full list generates without error.
				// however we can't just use this because godefs will only
				// use the right sub struct type if they're generated together.
				_, err := goDefsFor(mod, []string{structDef})
				if err != nil {
					badStructs = append(badStructs, s.Name)
					continue
				}

				if s.DocURL() != "" {
					// godefs strips comments so we'll have to reassemble with
					// comments using this
					docComments[modules.TrimPrefix(s.Name)] = fmt.Sprintf("// %s [Full Topic]\n//\n// [Full Topic]: %s\n", s.Description, s.DocURL())
				}

				structDefs = append(structDefs, structDef)
			}
		}

		if len(structDefs) == 0 {
			fmt.Fprintln(os.Stderr, "No structs for this framework")
			os.Exit(1)
		}

		out, err := goDefsFor(mod, structDefs)
		if err != nil {
			fmt.Println(out)
			log.Fatal(err)
		}

		// *byte and *[0]byte fields should just be uintptr
		out = strings.ReplaceAll(out, "*byte", "uintptr")
		out = strings.ReplaceAll(out, "*[0]byte", "uintptr")

		// reassemble with comments
		fmt.Println("package", mod.Package, "\n")
		parts := strings.Split(out, "type ")
		for _, p := range parts[1:] {
			words := strings.Fields(p)
			if s, ok := docComments[words[0]]; ok {
				fmt.Print(s)
			}
			fmt.Println("type", p)
		}

		if len(badStructs) > 0 {
			fmt.Println("// TODO (unable to generate):")
			fmt.Println("//", strings.Join(badStructs, " "))
		}
	}

}

func goDefsFor(mod *modules.Module, structDefs []string) (string, error) {
	extraLoad := ""
	extraInclude := ""
	source := fmt.Sprintf(`package main

/*
#cgo CFLAGS: -w -x objective-c
#cgo LDFLAGS: -framework %s %s
#include <%s>
%s
*/
import "C"

%s
`, mod.Name, extraLoad, mod.Header, extraInclude, strings.Join(structDefs, "\n"))
	b, err := evalSource(source)
	return string(b), err
}

func evalSource(source string) ([]byte, error) {
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

	cmd := exec.Command("go", "tool", "cgo", "-godefs", "--", "-x", "objective-c", tempFile.Name())
	return cmd.CombinedOutput()
}
