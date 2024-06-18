//go:build ignore

package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/progrium/darwinkit/generate/modules"
)

// go run ./generate/tools/initmod.go [platform] [framework]
func main() {
	platform := strings.ToLower(os.Args[1])
	m := modules.Get(os.Args[2])
	if m == nil {
		log.Fatal("framework not a known module")
	}

	dir := fmt.Sprintf("./%s/%s", platform, m.Package)
	os.MkdirAll(dir, 0755)

	modsrc := fmt.Sprintf(`//go:generate go run ../../generate/tools/genmod.go
package %s

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework %s
import "C"
`, m.Package, m.Name)
	if err := os.WriteFile(fmt.Sprintf("%s/%s.go", dir, m.Package), []byte(modsrc), 0644); err != nil {
		log.Fatal(err)
	}

	testsrc := fmt.Sprintf(`package %s

import "testing"
	
func Test%sValid(t *testing.T) {}
`, m.Package, m.Name)
	if err := os.WriteFile(fmt.Sprintf("%s/%s_test.go", dir, m.Package), []byte(testsrc), 0644); err != nil {
		log.Fatal(err)
	}

	customsrc := fmt.Sprintf("package %s\n", m.Package)
	if err := os.WriteFile(fmt.Sprintf("%s/%s_custom.go", dir, m.Package), []byte(customsrc), 0644); err != nil {
		log.Fatal(err)
	}

}
