package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var ImplementedTechnologies = map[string]bool{
	// Can be populated with manual entries, we autodiscover by default.
}

type Technology struct {
	Title       string `json:"title"`
	Destination struct {
		Identifier string `json:"identifier"`
		IsActive   bool   `json:"isActive"`
	} `json:"destination"`
	Languages []string `json:"languages"`
}

type Group struct {
	Name         string       `json:"name"`
	Technologies []Technology `json:"technologies"`
}

type Section struct {
	Kind   string  `json:"kind"`
	Groups []Group `json:"groups"`
}

type Data struct {
	Sections []Section `json:"sections"`
}

func main() {
	jsonFile, err := os.Open("technologies.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var data Data
	json.Unmarshal(byteValue, &data)

	fmt.Println("| Framework | Implemented |")
	fmt.Println("|-----------|-------------|")
	for _, section := range data.Sections {
		if section.Kind != "technologies" {
			continue
		}
		for _, group := range section.Groups {
			if group.Name != "App Services" {
				continue
			}
			for _, tech := range group.Technologies {
				if !tech.Destination.IsActive {
					continue
				}
				_, implemented := ImplementedTechnologies[tech.Title]
				implemented = implemented || dirExists(tech.Title)
				langObjC := contains(tech.Languages, "occ")
				if !langObjC {
					continue
				}
				fmt.Printf("| [%s](%s) | %v |\n",
					tech.Title,
					getURL(tech.Destination.Identifier),
					emojiMap[implemented])
			}
		}
	}
}

var emojiMap = map[bool]string{
	true:  "✅",
	false: "❌",
}

// getURL returns the URL for the given identifier. it replaces the doc:// scheme with https://developer.apple.com/documentation/
// doc://com.apple.documentation/documentation/accounts	 -> https://developer.apple.com/documentation/accounts
func getURL(identifier string) string {
	return strings.Replace(identifier, "doc://com.apple.documentation/documentation/", "https://developer.apple.com/documentation/", 1)
}

// dirExists returns true if the given directory exists
func dirExists(path string) bool {
	dir := strings.ToLower(path)
	dir = strings.Replace(dir, " ", "", -1)
	if _, err := os.Stat(filepath.Join("..", dir)); os.IsNotExist(err) {
		return false
	}
	return true
}

func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}
