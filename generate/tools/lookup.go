//go:build ignore

package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// go run ./generate/tools/lookup.go [framework/symbol, ex: appkit/nswindow]
func main() {
	dir := fmt.Sprintf("symbols/%s", strings.ToLower(os.Args[1]))
	filename := fmt.Sprintf("symbols/%s.json", strings.ToLower(os.Args[1]))

	r, err := zip.OpenReader("./generate/symbols.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	for _, file := range r.File {
		if file.Name == filename || filepath.Dir(file.Name) == dir {
			reader, err := file.Open()
			if err != nil {
				log.Fatal(err)
			}
			defer reader.Close()

			b, err := io.ReadAll(reader)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(string(b))
		}

	}

}
