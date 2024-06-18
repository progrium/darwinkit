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

// go run ./generate/tools/type.go [symbol, ex: nswindow]
func main() {
	db, err := generate.OpenSymbols("./generate/symbols.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s := db.FindTypeSymbol(os.Args[1])
	if s != nil {
		b, err := json.MarshalIndent(s, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
	}

	for _, s := range db.AllSymbols() {
		if strings.EqualFold(os.Args[1], s.Name) {
			fmt.Println(s.Path)
		}

	}

}
