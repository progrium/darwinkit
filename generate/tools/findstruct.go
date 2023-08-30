//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: findstruct <struct_name>")
		return
	}

	cmd := exec.Command("xcrun", "--show-sdk-path")
	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	dir := fmt.Sprintf("%s/System/Library/Frameworks", strings.TrimSpace(string(b)))
	structName := os.Args[1]

	// Naive regex to find C struct definitions
	// It won't capture complex scenarios or structs split across lines.
	re := regexp.MustCompile(`(?s)struct\s+` + structName + `\s+\{.*?\};`)

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && (filepath.Ext(path) == ".h" || filepath.Ext(path) == ".c") {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			var content string
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				content += scanner.Text() + "\n"
			}

			matches := re.FindAllString(content, -1)
			for _, match := range matches {
				fmt.Println(match)
			}
		}
		return nil
	})
}
