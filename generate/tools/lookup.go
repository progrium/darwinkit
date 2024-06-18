//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/progrium/darwinkit/generate"
)

// go run ./generate/tools/lookup.go [framework/symbol, ex: appkit/nswindow]
func main() {
	db, err := generate.OpenSymbols("./generate/symbols.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for _, s := range db.AllSymbols() {
		if strings.HasPrefix(s.Path, strings.ToLower(os.Args[1])) {
			b, err := json.MarshalIndent(s, "", "  ")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(b))
		}

	}

}
