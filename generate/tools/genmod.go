package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/progrium/macdriver/generate"
	"github.com/progrium/macdriver/generate/oldgen"
	"github.com/progrium/macdriver/internal/set"
)

// TODO: replace with env var
const TargetPlatform = "macos"

// TODO: replace with autodetect+env var
const TargetVersion = 12

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	db, err := generate.OpenSymbols("../../generate/symbols.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// first generate using oldgen (for now)
	oldgen.GenerateModules(filepath.Dir(cwd), []string{os.Getenv("GOPACKAGE")})

	// then use new generate for enums and others as they come online
	gen := generate.Generator{SymbolCache: db}
	gen.Generate(TargetPlatform, TargetVersion, filepath.Dir(cwd), os.Getenv("GOPACKAGE"), set.New[string]())

}
