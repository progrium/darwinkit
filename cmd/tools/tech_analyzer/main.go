package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const (
	// URL for Apple's technologies.json file
	technologiesURL = "https://developer.apple.com/tutorials/data/documentation/technologies.json"
)

// Technology represents an Apple technology/framework
type Technology struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	Path         string   `json:"path"`
	Type         string   `json:"type"`
	Role         string   `json:"role,omitempty"`
	Paths        []string `json:"paths,omitempty"`
	Children     []string `json:"children,omitempty"`
	SymbolCount  int      `json:"symbolCount,omitempty"`
	TopicCount   int      `json:"topicCount,omitempty"`
	Platforms    []string `json:"platforms,omitempty"`
	FromOSVersion string  `json:"fromOSVersion,omitempty"`
	ToOSVersion  string   `json:"toOSVersion,omitempty"`
}

// TechnologiesResponse is the root response structure
type TechnologiesResponse struct {
	Technologies []Technology `json:"technologies"`
}

func main() {
	// Define flags
	outputDir := flag.String("outdir", ".", "Directory to write output files")
	downloadOnly := flag.Bool("download-only", false, "Only download the technologies.json file without analysis")
	listAll := flag.Bool("list-all", false, "List all technologies")
	listMacOS := flag.Bool("list-macos", false, "List macOS technologies")
	countSymbols := flag.Bool("count-symbols", false, "Show symbol counts for technologies")
	genModules := flag.Bool("gen-modules", false, "Generate modules.go entries for selected frameworks")
	frameworks := flag.String("frameworks", "", "Comma-separated list of frameworks to analyze")
	outputFile := flag.String("output", "", "Output file path for analysis results")
	
	flag.Parse()

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create output directory: %v\n", err)
		os.Exit(1)
	}

	// Download the technologies.json file
	outputPath := filepath.Join(*outputDir, "technologies.json")
	if err := downloadFile(technologiesURL, outputPath); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to download technologies.json: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Technologies data downloaded to %s\n", outputPath)

	// If download-only flag is set, exit after downloading
	if *downloadOnly {
		return
	}

	// Read the downloaded file
	data, err := os.ReadFile(outputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read technologies.json: %v\n", err)
		os.Exit(1)
	}

	// Parse the JSON data
	var response TechnologiesResponse
	if err := json.Unmarshal(data, &response); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse technologies.json: %v\n", err)
		os.Exit(1)
	}

	// Create a map for easier access to technologies by ID
	techMap := make(map[string]Technology)
	for _, tech := range response.Technologies {
		techMap[tech.ID] = tech
	}

	// Filter frameworks by name if provided
	var selectedFrameworks []Technology
	if *frameworks != "" {
		frameworkList := strings.Split(*frameworks, ",")
		for _, frameworkName := range frameworkList {
			frameworkName = strings.TrimSpace(frameworkName)
			for _, tech := range response.Technologies {
				if strings.EqualFold(tech.Title, frameworkName) || strings.EqualFold(tech.ID, frameworkName) {
					selectedFrameworks = append(selectedFrameworks, tech)
					break
				}
			}
		}
	} else {
		selectedFrameworks = response.Technologies
	}

	// List all technologies if requested
	if *listAll {
		fmt.Println("\nAll Technologies:")
		for _, tech := range selectedFrameworks {
			fmt.Printf("- %s (ID: %s, Type: %s)\n", tech.Title, tech.ID, tech.Type)
		}
	}

	// List macOS technologies if requested
	if *listMacOS {
		fmt.Println("\nmacOS Technologies:")
		for _, tech := range selectedFrameworks {
			isMacOS := false
			for _, platform := range tech.Platforms {
				if strings.Contains(strings.ToLower(platform), "macos") {
					isMacOS = true
					break
				}
			}
			if isMacOS {
				fmt.Printf("- %s (ID: %s, Type: %s)\n", tech.Title, tech.ID, tech.Type)
			}
		}
	}

	// Show symbol counts if requested
	if *countSymbols {
		fmt.Println("\nSymbol Counts:")
		sort.Slice(selectedFrameworks, func(i, j int) bool {
			return selectedFrameworks[i].SymbolCount > selectedFrameworks[j].SymbolCount
		})
		for _, tech := range selectedFrameworks {
			if tech.SymbolCount > 0 {
				fmt.Printf("- %s: %d symbols\n", tech.Title, tech.SymbolCount)
			}
		}
	}

	// Generate modules.go entries if requested
	if *genModules {
		fmt.Println("\nModules.go Entries:")
		for _, tech := range selectedFrameworks {
			if tech.Type == "framework" {
				// Format the module entry
				packageName := strings.ToLower(tech.Title)
				header := tech.Title + "/" + tech.Title + ".h"
				prefixes := getFrameworkPrefixes(tech.Title)
				entry := fmt.Sprintf("{%q, %q, %q, %q, []string{%s}},", 
					tech.Title, tech.Title, packageName, header, formatPrefixes(prefixes))
				fmt.Println(entry)
			}
		}
	}

	// Write analysis to output file if requested
	if *outputFile != "" {
		var output strings.Builder

		// Title and overview
		output.WriteString("# Apple Framework Analysis\n\n")
		output.WriteString(fmt.Sprintf("Total frameworks: %d\n\n", len(response.Technologies)))
		
		// Framework tables
		output.WriteString("## Framework Symbol Counts\n\n")
		output.WriteString("| Framework | Symbol Count | Platforms |\n")
		output.WriteString("|-----------|--------------|----------|\n")
		
		sort.Slice(response.Technologies, func(i, j int) bool {
			return response.Technologies[i].SymbolCount > response.Technologies[j].SymbolCount
		})
		
		for _, tech := range response.Technologies {
			if tech.Type == "framework" && tech.SymbolCount > 0 {
				platforms := strings.Join(tech.Platforms, ", ")
				output.WriteString(fmt.Sprintf("| %s | %d | %s |\n", tech.Title, tech.SymbolCount, platforms))
			}
		}
		
		// macOS-specific frameworks
		output.WriteString("\n## macOS-Specific Frameworks\n\n")
		var macOSFrameworks []Technology
		for _, tech := range response.Technologies {
			if tech.Type == "framework" {
				isMacOS := false
				isIOS := false
				for _, platform := range tech.Platforms {
					if strings.Contains(strings.ToLower(platform), "macos") {
						isMacOS = true
					}
					if strings.Contains(strings.ToLower(platform), "ios") {
						isIOS = true
					}
				}
				if isMacOS && !isIOS {
					macOSFrameworks = append(macOSFrameworks, tech)
				}
			}
		}
		
		sort.Slice(macOSFrameworks, func(i, j int) bool {
			return macOSFrameworks[i].SymbolCount > macOSFrameworks[j].SymbolCount
		})
		
		output.WriteString("| Framework | Symbol Count |\n")
		output.WriteString("|-----------|--------------|----------|\n")
		for _, tech := range macOSFrameworks {
			output.WriteString(fmt.Sprintf("| %s | %d |\n", tech.Title, tech.SymbolCount))
		}
		
		// Write to file
		if err := os.WriteFile(*outputFile, []byte(output.String()), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to write output file: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Analysis written to %s\n", *outputFile)
	}
}

// downloadFile downloads a file from a URL to a local path
func downloadFile(url, filepath string) error {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// getFrameworkPrefixes returns common prefixes for the given framework
func getFrameworkPrefixes(framework string) []string {
	// Common prefix patterns
	prefixMap := map[string][]string{
		"AppKit":          {"NS"},
		"Foundation":      {"NS"},
		"CoreGraphics":    {"CG", "kCG"},
		"CoreFoundation":  {"CF", "kCF"},
		"AVFoundation":    {"AV"},
		"CoreData":        {"NS"},
		"CoreImage":       {"CI"},
		"CoreAudio":       {"Audio", "kAudio"},
		"CoreLocation":    {"CL"},
		"EventKit":        {"EK"},
		"UIKit":           {"UI"},
		"WebKit":          {"WK"},
		"CloudKit":        {"CK"},
		"MapKit":          {"MK"},
		"HealthKit":       {"HK"},
		"HomeKit":         {"HM"},
		"GameKit":         {"GK"},
		"Security":        {"SEC", "kSEC"},
		"JavaScriptCore":  {"JS"},
		"UserNotifications": {"UN"},
		"Metal":           {"MTL"},
		"QuartzCore":      {"CA", "kCA"},
		"CoreML":          {"ML"},
		"Vision":          {"VN"},
		"SpriteKit":       {"SK"},
		"SceneKit":        {"SCN"},
	}
	
	// Check for known prefixes
	if prefixes, ok := prefixMap[framework]; ok {
		return prefixes
	}
	
	// Generate a standard prefix if not found
	// Extract prefix from CamelCase (e.g., CoreMedia -> CM, NetworkExtension -> NE)
	var prefix string
	parts := splitCamelCase(framework)
	for _, part := range parts {
		if len(part) > 0 {
			prefix += string(part[0])
		}
	}
	
	return []string{prefix}
}

// splitCamelCase splits a CamelCase string into its component words
func splitCamelCase(s string) []string {
	var words []string
	var currentWord strings.Builder
	
	for i, r := range s {
		if i > 0 && isUpperCase(r) && !isUpperCase(rune(s[i-1])) {
			words = append(words, currentWord.String())
			currentWord.Reset()
		}
		currentWord.WriteRune(r)
	}
	
	if currentWord.Len() > 0 {
		words = append(words, currentWord.String())
	}
	
	return words
}

// isUpperCase checks if a rune is uppercase
func isUpperCase(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

// formatPrefixes formats a slice of prefixes for use in a module definition
func formatPrefixes(prefixes []string) string {
	var quoted []string
	for _, prefix := range prefixes {
		quoted = append(quoted, fmt.Sprintf("%q", prefix))
	}
	return strings.Join(quoted, ", ")
}