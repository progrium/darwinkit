package pdfkit

import (
	"testing"
)

func TestPDFKitValid(t *testing.T) {
	// This test just verifies that the package is valid.
}

// TestConstants ensures our custom constants are defined
func TestConstants(t *testing.T) {
	// These constants are defined in the package
	_ = MediaBox
	_ = CropBox
	_ = BleedBox
	_ = TrimBox
	_ = ArtBox
	
	_ = TextAnnotation
	_ = LinkAnnotation
	_ = FreeTextAnnotation
}