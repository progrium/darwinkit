// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PDFImageRep] class.
var PDFImageRepClass = _PDFImageRepClass{objc.GetClass("NSPDFImageRep")}

type _PDFImageRepClass struct {
	objc.Class
}

// An interface definition for the [PDFImageRep] class.
type IPDFImageRep interface {
	IImageRep
	PDFRepresentation() []byte
	PageCount() int
	Bounds() foundation.Rect
	CurrentPage() int
	SetCurrentPage(value int)
}

// An object that can render an image from a PDF format data stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfimagerep?language=objc
type PDFImageRep struct {
	ImageRep
}

func PDFImageRepFrom(ptr unsafe.Pointer) PDFImageRep {
	return PDFImageRep{
		ImageRep: ImageRepFrom(ptr),
	}
}

func (p_ PDFImageRep) InitWithData(pdfData []byte) PDFImageRep {
	rv := objc.Call[PDFImageRep](p_, objc.Sel("initWithData:"), pdfData)
	return rv
}

// Returns a representation of an image initialized with the specified PDF data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfimagerep/1535547-initwithdata?language=objc
func PDFImageRep_InitWithData(pdfData []byte) PDFImageRep {
	return PDFImageRepClass.Alloc().InitWithData(pdfData)
}

func (pc _PDFImageRepClass) ImageRepWithData(pdfData []byte) PDFImageRep {
	rv := objc.Call[PDFImageRep](pc, objc.Sel("imageRepWithData:"), pdfData)
	return rv
}

// Creates and returns a representation of an image initialized with the specified PDF data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfimagerep/1586713-imagerepwithdata?language=objc
func PDFImageRep_ImageRepWithData(pdfData []byte) PDFImageRep {
	return PDFImageRepClass.ImageRepWithData(pdfData)
}

func (pc _PDFImageRepClass) Alloc() PDFImageRep {
	rv := objc.Call[PDFImageRep](pc, objc.Sel("alloc"))
	return rv
}

func PDFImageRep_Alloc() PDFImageRep {
	return PDFImageRepClass.Alloc()
}

func (pc _PDFImageRepClass) New() PDFImageRep {
	rv := objc.Call[PDFImageRep](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPDFImageRep() PDFImageRep {
	return PDFImageRepClass.New()
}

func (p_ PDFImageRep) Init() PDFImageRep {
	rv := objc.Call[PDFImageRep](p_, objc.Sel("init"))
	return rv
}

// The PDF representation of the representation’s image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfimagerep/1530162-pdfrepresentation?language=objc
func (p_ PDFImageRep) PDFRepresentation() []byte {
	rv := objc.Call[[]byte](p_, objc.Sel("PDFRepresentation"))
	return rv
}

// The number of pages in the image representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfimagerep/1533063-pagecount?language=objc
func (p_ PDFImageRep) PageCount() int {
	rv := objc.Call[int](p_, objc.Sel("pageCount"))
	return rv
}

// The image representation’s bounding rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfimagerep/1533966-bounds?language=objc
func (p_ PDFImageRep) Bounds() foundation.Rect {
	rv := objc.Call[foundation.Rect](p_, objc.Sel("bounds"))
	return rv
}

// The page currently displayed by the image representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfimagerep/1528846-currentpage?language=objc
func (p_ PDFImageRep) CurrentPage() int {
	rv := objc.Call[int](p_, objc.Sel("currentPage"))
	return rv
}

// The page currently displayed by the image representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfimagerep/1528846-currentpage?language=objc
func (p_ PDFImageRep) SetCurrentPage(value int) {
	objc.Call[objc.Void](p_, objc.Sel("setCurrentPage:"), value)
}
