// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PDFConfiguration] class.
var PDFConfigurationClass = _PDFConfigurationClass{objc.GetClass("WKPDFConfiguration")}

type _PDFConfigurationClass struct {
	objc.Class
}

// An interface definition for the [PDFConfiguration] class.
type IPDFConfiguration interface {
	objc.IObject
	Rect() coregraphics.Rect
	SetRect(value coregraphics.Rect)
}

// The configuration data to use when generating a PDF representation of a web view’s contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpdfconfiguration?language=objc
type PDFConfiguration struct {
	objc.Object
}

func PDFConfigurationFrom(ptr unsafe.Pointer) PDFConfiguration {
	return PDFConfiguration{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PDFConfigurationClass) Alloc() PDFConfiguration {
	rv := objc.Call[PDFConfiguration](pc, objc.Sel("alloc"))
	return rv
}

func PDFConfiguration_Alloc() PDFConfiguration {
	return PDFConfigurationClass.Alloc()
}

func (pc _PDFConfigurationClass) New() PDFConfiguration {
	rv := objc.Call[PDFConfiguration](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPDFConfiguration() PDFConfiguration {
	return PDFConfigurationClass.New()
}

func (p_ PDFConfiguration) Init() PDFConfiguration {
	rv := objc.Call[PDFConfiguration](p_, objc.Sel("init"))
	return rv
}

// The portion of your web view to capture, specified as a rectangle in the view’s coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpdfconfiguration/3516860-rect?language=objc
func (p_ PDFConfiguration) Rect() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](p_, objc.Sel("rect"))
	return rv
}

// The portion of your web view to capture, specified as a rectangle in the view’s coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpdfconfiguration/3516860-rect?language=objc
func (p_ PDFConfiguration) SetRect(value coregraphics.Rect) {
	objc.Call[objc.Void](p_, objc.Sel("setRect:"), value)
}
