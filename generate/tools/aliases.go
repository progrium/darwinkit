//go:build ignore

package main

import (
	"log"
	"os"

	"github.com/progrium/macdriver/generate"
	"github.com/progrium/macdriver/internal/set"
)

const TargetPlatform = "macos"
const TargetVersion = 12

// go run ./generate/tools/aliases.go <framework>
func main() {
	db, err := generate.OpenSymbols("./generate/symbols.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	gen := generate.Generator{SymbolCache: db}
	gen.Generate(TargetPlatform, TargetVersion, "./macos", os.Args[1], set.New[string]())

}
