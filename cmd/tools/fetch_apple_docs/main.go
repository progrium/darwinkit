package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// AppleDocumentationResponse represents the root documentation response
type AppleDocumentationResponse struct {
	Abstract []struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"abstract"`
	Identifier struct {
		URL               string `json:"url"`
		InterfaceLanguage string `json:"interfaceLanguage"`
	} `json:"identifier"`
	Metadata struct {
		Title       string `json:"title"`
		Role        string `json:"role"`
		RoleHeading string `json:"roleHeading"`
		Platforms   []struct {
			Name         string `json:"name"`
			IntroducedAt string `json:"introducedAt"`
			Current      string `json:"current"`
		} `json:"platforms"`
		ExternalID string `json:"externalID"`
		Modules    []struct {
			Name string `json:"name"`
		} `json:"modules"`
	} `json:"metadata"`
	References map[string]struct {
		Title      string `json:"title"`
		Type       string `json:"type"`
		Identifier string `json:"identifier"`
		Kind       string `json:"kind"`
		Role       string `json:"role"`
		URL        string `json:"url"`
		Abstract   []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"abstract,omitempty"`
	} `json:"references"`
	TopicSections []struct {
		Kind        string   `json:"kind"`
		Title       string   `json:"title"`
		Identifiers []string `json:"identifiers"`
	} `json:"topicSections"`
	PrimaryContentSections []struct {
		Kind         string `json:"kind"`
		Declarations []struct {
			Languages []string `json:"languages"`
			Tokens    []struct {
				Kind string `json:"kind"`
				Text string `json:"text"`
			} `json:"tokens"`
		} `json:"declarations"`
	} `json:"primaryContentSections"`
}

// SymbolSummary represents our simplified symbol format
type SymbolSummary struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Kind        string `json:"kind"`
	Description string `json:"description"`
	Declaration string `json:"declaration,omitempty"`
	ExternalID  string `json:"externalID,omitempty"`
	Platforms   []struct {
		Name         string `json:"name"`
		IntroducedAt string `json:"introducedAt"`
		Current      string `json:"current"`
	} `json:"platforms"`
	Functions []struct {
		Name        string `json:"name"`
		Declaration string `json:"declaration"`
		Description string `json:"description"`
	} `json:"functions,omitempty"`
}

func fetchDocument(url string) (*AppleDocumentationResponse, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7)")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: received status code %d\nResponse body: %s", resp.StatusCode, string(body))
	}

	var docResponse AppleDocumentationResponse
	if err := json.Unmarshal(body, &docResponse); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v\nResponse body: %s", err, string(body))
	}

	return &docResponse, nil
}

func getObjCVariant(doc *AppleDocumentationResponse) string {
	for _, variant := range doc.TopicSections {
		if variant.Kind == "taskGroup" && variant.Title == "Objective-C" {
			for _, id := range variant.Identifiers {
				if ref, ok := doc.References[id]; ok && ref.Kind == "symbol" {
					return ref.URL
				}
			}
		}
	}
	return ""
}

func extractDeclaration(doc *AppleDocumentationResponse) string {
	for _, section := range doc.PrimaryContentSections {
		if section.Kind == "declarations" {
			for _, decl := range section.Declarations {
				if contains(decl.Languages, "occ") {
					var declaration strings.Builder
					for _, token := range decl.Tokens {
						declaration.WriteString(token.Text)
					}
					return declaration.String()
				}
			}
		}
	}
	return ""
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func processSymbol(baseURL string, ref struct {
	Title      string `json:"title"`
	Type       string `json:"type"`
	Identifier string `json:"identifier"`
	Kind       string `json:"kind"`
	Role       string `json:"role"`
	URL        string `json:"url"`
	Abstract   []struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"abstract,omitempty"`
}, visited map[string]bool) (*SymbolSummary, error) {
	if visited[ref.URL] {
		return nil, nil
	}
	visited[ref.URL] = true

	url := baseURL + ref.URL + ".json"
	doc, err := fetchDocument(url)
	if err != nil {
		return nil, err
	}

	// Try to get Objective-C variant
	objcURL := getObjCVariant(doc)
	if objcURL != "" {
		doc, err = fetchDocument(baseURL + objcURL + ".json")
		if err != nil {
			return nil, err
		}
	}

	symbol := &SymbolSummary{
		Name:       doc.Metadata.Title,
		Path:       strings.TrimPrefix(doc.Identifier.URL, "doc://com.apple.documentation/documentation/"),
		Kind:       doc.Metadata.RoleHeading,
		ExternalID: doc.Metadata.ExternalID,
		Platforms:  doc.Metadata.Platforms,
	}

	// Extract description
	var description strings.Builder
	for _, ab := range doc.Abstract {
		if ab.Type == "text" {
			description.WriteString(ab.Text)
		} else if ab.Type == "codeVoice" {
			description.WriteString("`" + ab.Text + "`")
		}
	}
	symbol.Description = description.String()

	// Extract declaration
	symbol.Declaration = extractDeclaration(doc)

	return symbol, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <framework>\n", os.Args[0])
		os.Exit(1)
	}

	framework := os.Args[1]
	baseURL := "https://developer.apple.com/tutorials/data/documentation"
	url := fmt.Sprintf("%s/%s.json", baseURL, framework)

	doc, err := fetchDocument(url)
	if err != nil {
		fmt.Printf("Error fetching framework documentation: %v\n", err)
		os.Exit(1)
	}

	outDir := filepath.Join("generate", "apidocs", framework)
	if err := os.MkdirAll(outDir, 0755); err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	var symbols []SymbolSummary
	visited := make(map[string]bool)

	// Process each reference
	for _, ref := range doc.References {
		if ref.Kind == "symbol" {
			fmt.Printf("Processing %s...\n", ref.Title)
			symbol, err := processSymbol(baseURL, ref, visited)
			if err != nil {
				fmt.Printf("Error processing symbol %s: %v\n", ref.Title, err)
				continue
			}
			if symbol != nil {
				symbols = append(symbols, *symbol)
			}
			time.Sleep(100 * time.Millisecond) // Be nice to Apple's servers
		}
	}

	// Save all symbols
	outPath := filepath.Join(outDir, "symbols.json")
	prettyJSON, err := json.MarshalIndent(symbols, "", "  ")
	if err != nil {
		fmt.Printf("Error formatting JSON: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile(outPath, prettyJSON, 0644); err != nil {
		fmt.Printf("Error writing symbols file: %v\n", err)
		os.Exit(1)
	}

	// Also save the original overview
	overviewPath := filepath.Join(outDir, "overview.json")
	overviewJSON, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		fmt.Printf("Error formatting overview JSON: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile(overviewPath, overviewJSON, 0644); err != nil {
		fmt.Printf("Error writing overview file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully processed %d symbols for %s\n", len(symbols), framework)
	fmt.Printf("Saved to %s and %s\n", outPath, overviewPath)
}