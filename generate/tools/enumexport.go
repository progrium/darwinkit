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

// go run ./generate/tools/enumexport.go [framework]
func main() {
	dir := fmt.Sprintf("symbols/%s", strings.ToLower(os.Args[1]))
	mod := modules.Get(strings.ToLower(os.Args[1]))

	r, err := zip.OpenReader("./generate/symbols.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	var enums []string
	inEnums := func(s string) bool {
		for _, enum := range enums {
			if enum == s {
				return true
			}
		}
		return false
	}
	for _, file := range r.File {
		if filepath.Dir(file.Name) == dir {
			s, err := generate.LoadSymbolFrom(file)
			if err != nil {
				continue
			}
			if s.Kind != "Enum" {
				continue
			}
			enums = append(enums, s.Name)
		}
	}
	var enumConsts []string
	for _, file := range r.File {
		if filepath.Dir(file.Name) == dir {
			s, err := generate.LoadSymbolFrom(file)
			if err != nil {
				continue
			}
			if s.Kind != "Constant" {
				continue
			}
			if s.Type != "Enumeration Case" {
				continue
			}
			if !s.HasPlatform("macos") {
				continue
			}
			if !inEnums(s.Parent) {
				continue
			}
			enumConsts = append(enumConsts, s.Name)
		}
	}

	sort.Strings(enumConsts)

	for idx, c := range enumConsts {
		enumConsts[idx] = fmt.Sprintf("	printf(\"%s %%ld\\n\", (long)%s);\n", c, c)
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
`, mod.Name, mod.Header, strings.Join(enumConsts, ""))

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

	fmt.Println(string(output))

}
