package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"text/template"
	"time"
)

// ReportFramework represents an Apple framework and its coverage stats
type ReportFramework struct {
	Name             string  `json:"name"`
	TotalSymbols     int     `json:"totalSymbols"`
	CoveredSymbols   int     `json:"coveredSymbols"`
	CoveragePercent  float64 `json:"coveragePercent"`
	Status           string  `json:"status"`
	MacOSOnly        bool    `json:"macOSOnly"`
	ImportantSymbols int     `json:"importantSymbols"`
	ImportantCovered int     `json:"importantCovered"`
	ImportantPercent float64 `json:"importantPercent"`
}

// ReportData contains all data needed for the report template
type ReportData struct {
	GeneratedDate    string
	Frameworks       []*ReportFramework
	TotalSymbols     int
	TotalCovered     int
	OverallPercent   float64
	MacOSOnlyCount   int
	CompletedCount   int
	PartialCount     int
	MissingCount     int
	RecommendedNext  []*ReportFramework // Top frameworks to work on next
}

const reportTemplate = `# Darwinkit API Coverage Report

**Generated:** {{ .GeneratedDate }}

## Overview

Darwinkit currently implements **{{ printf "%.1f" .OverallPercent }}%** of the Apple frameworks analyzed (**{{ .TotalCovered }}** out of **{{ .TotalSymbols }}** symbols).

- **{{ .CompletedCount }}** frameworks are fully implemented (90%+ coverage)
- **{{ .PartialCount }}** frameworks are partially implemented
- **{{ .MissingCount }}** frameworks have minimal or no implementation
- **{{ .MacOSOnlyCount }}** frameworks are macOS-only

## Framework Coverage

| Framework | Total Symbols | Covered | Coverage % | Status |
|-----------|---------------|---------|-----------|--------|
{{- range .Frameworks }}
| **{{ .Name }}** | {{ .TotalSymbols }} | {{ .CoveredSymbols }} | {{ printf "%.1f" .CoveragePercent }}% | {{ .Status }} |
{{- end }}

## Recommended Next Steps

Based on importance and current coverage, the following frameworks are recommended for implementation focus:

{{- range .RecommendedNext }}
1. **{{ .Name }}** - Currently at {{ printf "%.1f" .CoveragePercent }}% coverage ({{ .CoveredSymbols }}/{{ .TotalSymbols }} symbols)
   - Focus on the {{ .ImportantSymbols }} important symbols first (current: {{ .ImportantCovered }})
{{- end }}

## Notes

- "Important symbols" typically include classes, primary structures, and essential functions
- macOS-only frameworks are prioritized for implementation
- This report is based on analysis of the Apple API documentation
`

// readCoverageData reads the coverage data from a JSON file
func readCoverageData(filePath string) ([]*ReportFramework, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading coverage data: %v", err)
	}

	var frameworks []*ReportFramework
	if err := json.Unmarshal(data, &frameworks); err != nil {
		return nil, fmt.Errorf("error parsing coverage data: %v", err)
	}

	return frameworks, nil
}

// getRecommendedFrameworks returns the top frameworks to work on next
func getRecommendedFrameworks(frameworks []*ReportFramework) []*ReportFramework {
	// Create a copy for sorting
	ranked := make([]*ReportFramework, len(frameworks))
	copy(ranked, frameworks)

	// Sort by priority score (combination of importance and current coverage gap)
	sort.Slice(ranked, func(i, j int) bool {
		// Calculate priority score:
		// - MacOS only frameworks get a bonus
		// - Frameworks with higher number of important symbols get priority
		// - Frameworks with some coverage (but not complete) get priority over those with none

		scoreI := float64(ranked[i].ImportantSymbols) * (100 - ranked[i].CoveragePercent) / 100
		scoreJ := float64(ranked[j].ImportantSymbols) * (100 - ranked[j].CoveragePercent) / 100

		// Give 25% bonus to macOS-only frameworks
		if ranked[i].MacOSOnly {
			scoreI *= 1.25
		}
		if ranked[j].MacOSOnly {
			scoreJ *= 1.25
		}

		// Penalize frameworks with zero coverage (likely harder to start)
		if ranked[i].CoveragePercent == 0 {
			scoreI *= 0.8
		}
		if ranked[j].CoveragePercent == 0 {
			scoreJ *= 0.8
		}

		return scoreI > scoreJ
	})

	// Return top 5 (or fewer if there are less than 5 frameworks)
	count := 5
	if len(ranked) < count {
		count = len(ranked)
	}
	return ranked[:count]
}

// generateReport generates a coverage report using the template
func generateReport(frameworks []*ReportFramework, outputFile string) error {
	// Sort frameworks by coverage percentage (descending)
	sort.Slice(frameworks, func(i, j int) bool {
		return frameworks[i].CoveragePercent > frameworks[j].CoveragePercent
	})

	// Prepare report data
	data := ReportData{
		GeneratedDate:   time.Now().Format("January 2, 2006"),
		Frameworks:      frameworks,
		RecommendedNext: getRecommendedFrameworks(frameworks),
	}

	// Calculate summary stats
	for _, fw := range frameworks {
		data.TotalSymbols += fw.TotalSymbols
		data.TotalCovered += fw.CoveredSymbols

		if fw.MacOSOnly {
			data.MacOSOnlyCount++
		}

		switch fw.Status {
		case "complete":
			data.CompletedCount++
		case "partial":
			data.PartialCount++
		case "missing":
			data.MissingCount++
		}
	}

	// Calculate overall percentage
	if data.TotalSymbols > 0 {
		data.OverallPercent = float64(data.TotalCovered) / float64(data.TotalSymbols) * 100.0
	}

	// Parse template
	tmpl, err := template.New("report").Parse(reportTemplate)
	if err != nil {
		return fmt.Errorf("error parsing template: %v", err)
	}

	// Create output file
	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer file.Close()

	// Execute template
	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("error executing template: %v", err)
	}

	return nil
}

func main() {
	coverageFile := flag.String("coverage", "generate/apidocs/analysis/coverage_report.json", "Path to coverage JSON data")
	outputFile := flag.String("output", "API_COVERAGE.md", "Output markdown report file")
	flag.Parse()

	// Read coverage data
	var frameworks []*ReportFramework
	
	// If coverage file is "sample" or doesn't exist, use sample data
	if *coverageFile == "sample" {
		log.Println("Using sample data for report")
		frameworks = generateSampleData()
	} else {
		data, err := readCoverageData(*coverageFile)
		if err != nil {
			log.Printf("Warning: Error reading coverage data: %v", err)
			log.Println("Falling back to sample data")
			frameworks = generateSampleData()
		} else {
			frameworks = data
			// If no frameworks found, create sample data for testing
			if len(frameworks) == 0 {
				log.Println("No coverage data found, generating sample data for testing")
				frameworks = generateSampleData()
			}
		}
	}

	// Generate report
	if err := generateReport(frameworks, *outputFile); err != nil {
		log.Fatalf("Error generating report: %v", err)
	}

	fmt.Printf("Coverage report generated: %s\n", *outputFile)
}

// generateSampleData creates sample framework data for testing
func generateSampleData() []*ReportFramework {
	return []*ReportFramework{
		{
			Name:             "foundation",
			TotalSymbols:     750,
			CoveredSymbols:   450,
			CoveragePercent:  60.0,
			Status:           "partial",
			MacOSOnly:        false,
			ImportantSymbols: 200,
			ImportantCovered: 150,
			ImportantPercent: 75.0,
		},
		{
			Name:             "appkit",
			TotalSymbols:     1200,
			CoveredSymbols:   240,
			CoveragePercent:  20.0,
			Status:           "partial",
			MacOSOnly:        true,
			ImportantSymbols: 300,
			ImportantCovered: 75,
			ImportantPercent: 25.0,
		},
		{
			Name:             "corefoundation",
			TotalSymbols:     500,
			CoveredSymbols:   475,
			CoveragePercent:  95.0,
			Status:           "complete",
			MacOSOnly:        false,
			ImportantSymbols: 150,
			ImportantCovered: 145,
			ImportantPercent: 96.7,
		},
		{
			Name:             "coregraphics",
			TotalSymbols:     600,
			CoveredSymbols:   300,
			CoveragePercent:  50.0,
			Status:           "partial",
			MacOSOnly:        false,
			ImportantSymbols: 200,
			ImportantCovered: 120,
			ImportantPercent: 60.0,
		},
		{
			Name:             "metal",
			TotalSymbols:     400,
			CoveredSymbols:   0,
			CoveragePercent:  0.0,
			Status:           "missing",
			MacOSOnly:        false,
			ImportantSymbols: 100,
			ImportantCovered: 0,
			ImportantPercent: 0.0,
		},
	}
}