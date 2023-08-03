//go:build ignore

package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/progrium/macdriver/generate/declparse"
)

// go run ./generate/tools/declcheck.go [optional symbol prefix, ex: appkit]
func main() {
	filter := ""
	if len(os.Args) > 1 {
		filter = os.Args[1]
	}

	r, err := zip.OpenReader("./symbols.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	countDecls := map[string]int{}
	countPass := map[string]int{}
	for _, file := range r.File {
		if !file.FileInfo().IsDir() {
			if filter != "" {
				if !strings.HasPrefix(strings.TrimPrefix(file.Name, "symbols/"), filter) {
					continue
				}
			}

			reader, err := file.Open()
			if err != nil {
				log.Fatal(err)
			}
			defer reader.Close()

			b, err := io.ReadAll(reader)
			if err != nil {
				log.Fatal(err)
			}
			var s Symbol
			if err := json.Unmarshal(b, &s); err != nil {
				log.Fatal(err)
			}
			if s.Declaration != "" && !strings.HasPrefix(s.Declaration, "#define") {
				countDecls[s.Kind]++
				//fmt.Println(s.Declaration)
				p := declparse.NewStringParser(normalizeStmntString(s.Declaration))
				switch s.Kind {
				case "Constant", "Property":
					p.Hint = declparse.HintVariable
				case "Function":
					p.Hint = declparse.HintFunction
				default:
				}
				if s.Type == "Enumeration Case" {
					p.Hint = declparse.HintEnumCase
				}
				_, err := p.Parse()
				if err != nil {
					log.Println(err, "::", s.Type, "::", s.Declaration)
				} else {
					countPass[s.Kind]++
				}
			}
		}
	}
	for k, _ := range countDecls {
		fmt.Printf("%s: %d/%d %d%% \n", k, countPass[k], countDecls[k], int((float64(countPass[k])/float64(countDecls[k]))*100))
	}

}

type Symbol struct {
	Name        string
	Path        string
	Kind        string
	Declaration string
	Type        string
}

func normalizeStmntString(s string) string {
	return strings.TrimRight(s, ";") + ";"
}
