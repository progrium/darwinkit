//go:build ignore

package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/progrium/macdriver/generate"
	"github.com/progrium/macdriver/generate/declparse"
)

// go run ./generate/tools/declcheck.go [optional symbol prefix, ex: appkit]
func main() {
	filter := ""
	if len(os.Args) > 1 {
		filter = os.Args[1]
	}

	db, err := generate.OpenSymbols("./generate/symbols.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	countDecls := map[string]int{}
	countPass := map[string]int{}
	for _, s := range db.AllSymbols() {
		if filter != "" {
			if !strings.HasPrefix(s.Path, filter) {
				continue
			}
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
