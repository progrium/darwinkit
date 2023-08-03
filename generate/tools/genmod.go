package main

import (
	"os"
	"path/filepath"

	"github.com/progrium/macdriver/generate/oldgen"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	oldgen.GenerateModules(filepath.Dir(cwd), []string{os.Getenv("GOPACKAGE")})

}
