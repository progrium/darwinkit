//go:build ignore

package main

import (
	"archive/zip"
	"fmt"
	"log"
	"os"

	"github.com/progrium/macdriver/generate"
)

const TargetPlatform = "macos"
const TargetVersion = 12

// go run ./generate/tools/aliases.go <framework>
func main() {
	db, err := zip.OpenReader("./generate/symbols.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for _, s := range generate.FrameworkSymbols(db, os.Args[1]) {
		if !s.HasPlatform(TargetPlatform, TargetVersion, true) {
			continue
		}
		if s.Kind == "Enum" {
			fmt.Println(s.Name)
			continue
		}
		if s.Type == "Type Alias" {
			fmt.Println(s.Name)
			continue
		}
	}

}
