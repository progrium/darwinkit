package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Known Apple frameworks that we want to analyze
var knownFrameworks = []string{
	"appkit",
	"foundation",
	"corefoundation",
	"coredata",
	"coregraphics",
	"coreaudio",
	"coremedia",
	"avfoundation",
	"gamekit",
	"metal",
	"uikit",
	"webkit",
	"cloudkit",
	"mapkit",
	"scenekit",
	"spritekit",
	"arkit",
	"healthkit",
	"homekit",
	"security",
	"networkextension",
}

// AppleDocResponse represents the root documentation response
type AppleDocResponse struct {
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

// FrameworkStats holds statistics about a framework's documentation coverage
type FrameworkStats struct {
	Name          string
	SymbolCount   int
	ClassCount    int
	FunctionCount int
	EnumCount     int
	StructCount   int
	ProtocolCount int
	OtherCount    int
	Platforms     map[string]bool
	MacOSOnly     bool
}

// fetchDoc fetches a document from Apple's documentation API
func fetchDoc(url string) (*AppleDocResponse, error) {
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
		return nil, fmt.Errorf("error: received status code %d: %s", resp.StatusCode, string(body))
	}

	var docResponse AppleDocResponse
	if err := json.Unmarshal(body, &docResponse); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	return &docResponse, nil
}

// analyzeFramework analyzes a framework and returns statistics
func analyzeFramework(framework string, baseURL string) (*FrameworkStats, error) {
	url := fmt.Sprintf("%s/%s.json", baseURL, framework)
	
	doc, err := fetchDoc(url)
	if err != nil {
		return nil, err
	}
	
	stats := &FrameworkStats{
		Name:       framework,
		Platforms:  make(map[string]bool),
	}
	
	// Count symbol categories
	for _, ref := range doc.References {
		if ref.Kind == "symbol" {
			stats.SymbolCount++
			
			switch ref.Role {
			case "class":
				stats.ClassCount++
			case "function":
				stats.FunctionCount++
			case "enumeration", "enumeration case":
				stats.EnumCount++
			case "struct":
				stats.StructCount++
			case "protocol":
				stats.ProtocolCount++
			default:
				stats.OtherCount++
			}
		}
	}
	
	// Record platform support
	for _, platform := range doc.Metadata.Platforms {
		stats.Platforms[platform.Name] = true
	}
	
	// Check if macOS only
	stats.MacOSOnly = stats.Platforms["macOS"] && !stats.Platforms["iOS"] && !stats.Platforms["tvOS"] && !stats.Platforms["watchOS"]
	
	return stats, nil
}

// contains checks if a slice contains a specific item
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func saveStatsToJSON(stats []*FrameworkStats, outPath string) error {
	jsonData, err := json.MarshalIndent(stats, "", "  ")
	if err != nil {
		return fmt.Errorf("error serializing stats: %v", err)
	}

	return os.WriteFile(outPath, jsonData, 0644)
}

func main() {
	outDir := flag.String("outdir", "analysis", "Directory to store analysis results")
	filterMacOnly := flag.Bool("macos-only", false, "Only analyze frameworks that are macOS-only")
	flag.Parse()

	baseURL := "https://developer.apple.com/tutorials/data/documentation"
	
	// Create output directory
	if err := os.MkdirAll(*outDir, 0755); err != nil {
		log.Fatalf("Error creating output directory: %v", err)
	}
	
	var allStats []*FrameworkStats
	var mutex sync.Mutex
	var wg sync.WaitGroup
	
	// Process each framework
	for _, framework := range knownFrameworks {
		wg.Add(1)
		go func(fw string) {
			defer wg.Done()
			
			fmt.Printf("Analyzing %s...\n", fw)
			stats, err := analyzeFramework(fw, baseURL)
			if err != nil {
				fmt.Printf("Error analyzing %s: %v\n", fw, err)
				return
			}
			
			// Skip if not macOS-only and filter is enabled
			if *filterMacOnly && !stats.MacOSOnly {
				fmt.Printf("Skipping %s (not macOS-only)\n", fw)
				return
			}
			
			mutex.Lock()
			allStats = append(allStats, stats)
			mutex.Unlock()
			
			fmt.Printf("Completed %s: %d symbols found\n", fw, stats.SymbolCount)
			time.Sleep(100 * time.Millisecond) // Rate limit
		}(framework)
	}
	
	wg.Wait()
	
	// Save the analysis
	outPath := filepath.Join(*outDir, "framework_stats.json")
	if err := saveStatsToJSON(allStats, outPath); err != nil {
		log.Fatalf("Error saving analysis: %v", err)
	}
	
	// Generate a simple report
	fmt.Println("\nFramework Analysis Report:")
	fmt.Println("==========================")
	
	// Sort by symbol count (we would import sort package for this in real code)
	for _, stats := range allStats {
		fmt.Printf("%-20s: %5d symbols (%d classes, %d functions, %d enums, %d structs, %d protocols)\n",
			stats.Name, stats.SymbolCount, stats.ClassCount, stats.FunctionCount, 
			stats.EnumCount, stats.StructCount, stats.ProtocolCount)
	}
	
	fmt.Printf("\nAnalysis saved to %s\n", outPath)
}