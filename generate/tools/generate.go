//go:build ignore

package main

import (
	"os"

	"github.com/progrium/macdriver/generate/oldgen"
)

// go run ./generate/tools/generate.go [dir, ex: ./macos]
func main() {
	oldgen.GenerateAll(os.Args[1])
}
