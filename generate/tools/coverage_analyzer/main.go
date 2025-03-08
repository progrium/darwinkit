package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// AnalysisFramework represents an Apple framework and its coverage in darwinkit
type AnalysisFramework struct {
	Name              string  `json:"name"`
	TotalSymbols      int     `json:"totalSymbols"`
	CoveredSymbols    int     `json:"coveredSymbols"`
	CoveragePercent   float64 `json:"coveragePercent"`
	Status            string  `json:"status"` // "complete", "partial", "missing"
	MacOSOnly         bool    `json:"macOSOnly"`
	ImportantSymbols  int     `json:"importantSymbols"` 
	ImportantCovered  int     `json:"importantCovered"`
	ImportantPercent  float64 `json:"importantPercent"`
}

func main() {
	appleDocsDir := flag.String("apple-docs", "generate/apidocs", "Path to Apple documentation directory")
	darwinkitDir := flag.String("darwinkit", ".", "Path to darwinkit repository")
	outFile := flag.String("out", "coverage_report.json", "Output file for coverage report")
	flag.Parse()

	// Load Apple documentation stats
	appleStats, err := loadAppleStats(*appleDocsDir)
	if err != nil {
		log.Fatalf("Error loading Apple documentation stats: %v", err)
	}

	// Analyze darwinkit coverage
	coverage := analyzeFrameworkCoverage(appleStats, *darwinkitDir)

	// Save results
	if err := saveCoverageReport(coverage, *outFile); err != nil {
		log.Fatalf("Error saving coverage report: %v", err)
	}

	// Print summary
	printCoverageSummary(coverage, *outFile)
}

func loadAppleStats(docsDir string) (map[string]map[string]int, error) {
	// This would normally load from the apidocs directory
	// For demonstration, we'll use a placeholder
	
	stats := make(map[string]map[string]int)
	
	// Check if we have an analysis file from the docs_analyzer
	analysisFile := filepath.Join(docsDir, "analysis", "framework_stats.json")
	if _, err := os.Stat(analysisFile); err == nil {
		data, err := os.ReadFile(analysisFile)
		if err != nil {
			return nil, fmt.Errorf("error reading analysis file: %v", err)
		}
		
		var frameworkStats []*struct {
			Name         string
			SymbolCount  int
			ClassCount   int
			FunctionCount int
		}
		
		if err := json.Unmarshal(data, &frameworkStats); err != nil {
			return nil, fmt.Errorf("error parsing analysis file: %v", err)
		}
		
		for _, fs := range frameworkStats {
			stats[fs.Name] = map[string]int{
				"symbols":   fs.SymbolCount,
				"classes":   fs.ClassCount,
				"functions": fs.FunctionCount,
			}
		}
		
		return stats, nil
	}
	
	// If we don't have real stats, use placeholders
	frameworks := []string{
		"appkit", "foundation", "corefoundation", "coregraphics", "coredata",
		"coreaudio", "coremedia", "avfoundation", "webkit", "metal",
	}
	
	for _, fw := range frameworks {
		stats[fw] = map[string]int{
			"symbols":   500, // placeholder
			"classes":   50,  // placeholder
			"functions": 100, // placeholder
		}
	}
	
	return stats, nil
}

func analyzeFrameworkCoverage(appleStats map[string]map[string]int, darwinkitDir string) []*AnalysisFramework {
	// This would normally scan through the darwinkit code to see what's implemented
	// For demonstration, we'll simulate coverage analysis
	
	var coverage []*AnalysisFramework
	
	for fw, stats := range appleStats {
		totalSymbols := stats["symbols"]
		
		// Count symbols in darwinkit (simulated)
		coveredSymbols := estimateCoveredSymbols(fw, darwinkitDir)
		coveragePercent := float64(coveredSymbols) / float64(totalSymbols) * 100.0
		
		status := "missing"
		if coveragePercent >= 90 {
			status = "complete"
		} else if coveragePercent > 0 {
			status = "partial"
		}
		
		// Estimate important symbols (usually classes and core functions)
		importantSymbols := stats["classes"] + (stats["functions"] / 5)
		importantCovered := min(coveredSymbols, importantSymbols)
		importantPercent := float64(importantCovered) / float64(importantSymbols) * 100.0
		
		coverage = append(coverage, &AnalysisFramework{
			Name:             fw,
			TotalSymbols:     totalSymbols,
			CoveredSymbols:   coveredSymbols,
			CoveragePercent:  coveragePercent,
			Status:           status,
			MacOSOnly:        isMacOSOnly(fw),
			ImportantSymbols: importantSymbols,
			ImportantCovered: importantCovered,
			ImportantPercent: importantPercent,
		})
	}
	
	return coverage
}

func estimateCoveredSymbols(framework string, darwinkitDir string) int {
	// In a real implementation, we would scan darwinkit source files
	// to find actual symbols from this framework
	
	// For demonstration, simulate varying levels of coverage
	switch framework {
	case "foundation", "corefoundation":
		return 400 // Simulated high coverage
	case "appkit", "coregraphics":
		return 250 // Simulated medium coverage
	case "webkit", "metal":
		return 50 // Simulated low coverage
	default:
		return 20 // Simulated minimal coverage
	}
}

func isMacOSOnly(framework string) bool {
	macOSOnlyFrameworks := []string{"appkit"}
	for _, fw := range macOSOnlyFrameworks {
		if fw == framework {
			return true
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func saveCoverageReport(coverage []*AnalysisFramework, outFile string) error {
	data, err := json.MarshalIndent(coverage, "", "  ")
	if err != nil {
		return fmt.Errorf("error serializing coverage report: %v", err)
	}
	
	return os.WriteFile(outFile, data, 0644)
}

func printCoverageSummary(coverage []*AnalysisFramework, outFile string) {
	fmt.Println("\nDarwinkit Framework Coverage Summary:")
	fmt.Println("====================================")
	fmt.Printf("%-15s %-10s %-10s %-10s %-10s\n", "Framework", "Total", "Covered", "Coverage", "Status")
	fmt.Println("------------------------------------------------------------")
	
	var totalSymbols, totalCovered int
	
	for _, fw := range coverage {
		fmt.Printf("%-15s %-10d %-10d %-10.1f%% %-10s\n", 
			fw.Name, fw.TotalSymbols, fw.CoveredSymbols, fw.CoveragePercent, fw.Status)
		
		totalSymbols += fw.TotalSymbols
		totalCovered += fw.CoveredSymbols
	}
	
	fmt.Println("------------------------------------------------------------")
	totalPercent := float64(totalCovered) / float64(totalSymbols) * 100.0
	fmt.Printf("%-15s %-10d %-10d %-10.1f%%\n", "TOTAL", totalSymbols, totalCovered, totalPercent)
	
	fmt.Printf("\nCoverage report saved to %s\n", outFile)
}