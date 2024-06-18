//go:build ignore

package main

import (
	"os"

	"github.com/progrium/darwinkit/generate"
)

// go run ./generate/tools/clobbergen.go [dir, ex: ./macos/appkit]
func main() {
	generate.RemoveGeneratedCode(os.Args[1])
}
