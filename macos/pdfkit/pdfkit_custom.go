package pdfkit

import (
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// PDFDocument represents a PDF document.
type PDFDocument struct {
	objc.Object
}

// PDFPage represents a single page in a PDF document.
type PDFPage struct {
	objc.Object
}

// PDFAnnotation represents an annotation on a PDF page.
type PDFAnnotation struct {
	objc.Object
}

// PDFOutline represents an item in the PDF document's outline.
type PDFOutline struct {
	objc.Object
}

// Document creation and loading

// NewPDFDocument creates a new PDFDocument.
func NewPDFDocument() PDFDocument {
	obj := objc.Call[objc.Object](objc.GetClass("PDFDocument"), objc.Sel("alloc"))
	return PDFDocument{objc.Call[objc.Object](obj, objc.Sel("init"))}
}

// DocumentWithURL creates a PDFDocument from a URL.
func DocumentWithURL(url foundation.URL) PDFDocument {
	return PDFDocument{objc.Call[objc.Object](objc.GetClass("PDFDocument"), objc.Sel("documentWithURL:"), url)}
}

// DocumentWithData creates a PDFDocument from data.
func DocumentWithData(data foundation.Data) PDFDocument {
	return PDFDocument{objc.Call[objc.Object](objc.GetClass("PDFDocument"), objc.Sel("documentWithData:"), data)}
}

// Document properties and operations

// PageCount returns the number of pages in the document.
func (d PDFDocument) PageCount() int {
	return objc.Call[int](d, objc.Sel("pageCount"))
}

// PageAtIndex returns the page at the specified index.
func (d PDFDocument) PageAtIndex(index int) PDFPage {
	return PDFPage{objc.Call[objc.Object](d, objc.Sel("pageAtIndex:"), index)}
}

// Outlines returns the document's outline items.
func (d PDFDocument) Outlines() []PDFOutline {
	outlineRoot := objc.Call[objc.Object](d, objc.Sel("outlineRoot"))
	outlineArray := objc.Call[objc.Object](outlineRoot, objc.Sel("children"))
	count := objc.Call[int](outlineArray, objc.Sel("count"))
	result := make([]PDFOutline, count)
	
	for i := 0; i < count; i++ {
		result[i] = PDFOutline{objc.Call[objc.Object](outlineArray, objc.Sel("objectAtIndex:"), i)}
	}
	
	return result
}

// WriteToURL writes the document to the specified URL.
func (d PDFDocument) WriteToURL(url foundation.URL) bool {
	return objc.Call[bool](d, objc.Sel("writeToURL:withOptions:"), url, nil)
}

// Page operations

// BoundsForBox returns the bounds for the specified box type.
func (p PDFPage) BoundsForBox(box string) foundation.Rect {
	return objc.Call[foundation.Rect](p, objc.Sel("boundsForBox:"), box)
}

// RotationAngle returns the rotation angle of the page.
func (p PDFPage) RotationAngle() int {
	return objc.Call[int](p, objc.Sel("rotation"))
}

// SetRotationAngle sets the rotation angle of the page.
func (p PDFPage) SetRotationAngle(angle int) {
	objc.Call[objc.Void](p, objc.Sel("setRotation:"), angle)
}

// Annotations returns all annotations on the page.
func (p PDFPage) Annotations() []PDFAnnotation {
	annotationArray := objc.Call[objc.Object](p, objc.Sel("annotations"))
	count := objc.Call[int](annotationArray, objc.Sel("count"))
	result := make([]PDFAnnotation, count)
	
	for i := 0; i < count; i++ {
		result[i] = PDFAnnotation{objc.Call[objc.Object](annotationArray, objc.Sel("objectAtIndex:"), i)}
	}
	
	return result
}

// Annotation operations

// Type returns the type of the annotation.
func (a PDFAnnotation) Type() string {
	return objc.Call[string](a, objc.Sel("type"))
}

// Contents returns the contents of the annotation.
func (a PDFAnnotation) Contents() string {
	return objc.Call[string](a, objc.Sel("contents"))
}

// SetContents sets the contents of the annotation.
func (a PDFAnnotation) SetContents(contents string) {
	objc.Call[objc.Void](a, objc.Sel("setContents:"), contents)
}

// Outline operations

// Label returns the label of the outline item.
func (o PDFOutline) Label() string {
	return objc.Call[string](o, objc.Sel("label"))
}

// Children returns the child outline items.
func (o PDFOutline) Children() []PDFOutline {
	childrenArray := objc.Call[objc.Object](o, objc.Sel("children"))
	count := objc.Call[int](childrenArray, objc.Sel("count"))
	result := make([]PDFOutline, count)
	
	for i := 0; i < count; i++ {
		result[i] = PDFOutline{objc.Call[objc.Object](childrenArray, objc.Sel("objectAtIndex:"), i)}
	}
	
	return result
}

// PDFDocument box constants
const (
	MediaBox = "kPDFDisplayBoxMediaBox"
	CropBox = "kPDFDisplayBoxCropBox"
	BleedBox = "kPDFDisplayBoxBleedBox"
	TrimBox = "kPDFDisplayBoxTrimBox"
	ArtBox = "kPDFDisplayBoxArtBox"
)

// PDFAnnotation types
const (
	TextAnnotation = "Text"
	LinkAnnotation = "Link"
	FreeTextAnnotation = "FreeText"
	LineAnnotation = "Line"
	SquareAnnotation = "Square"
	CircleAnnotation = "Circle"
	HighlightAnnotation = "Highlight"
	UnderlineAnnotation = "Underline"
	StrikeOutAnnotation = "StrikeOut"
	StampAnnotation = "Stamp"
	InkAnnotation = "Ink"
	PopupAnnotation = "Popup"
	FileAttachmentAnnotation = "FileAttachment"
	SoundAnnotation = "Sound"
	MovieAnnotation = "Movie"
	WidgetAnnotation = "Widget"
	ScreenAnnotation = "Screen"
	PrinterMarkAnnotation = "PrinterMark"
	TrapNetAnnotation = "TrapNet"
	WatermarkAnnotation = "Watermark"
	ThreeDAnnotation = "3D"
	RichMediaAnnotation = "RichMedia"
)