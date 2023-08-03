// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var PDFInfoClass = _PDFInfoClass{objc.GetClass("NSPDFInfo")}

type _PDFInfoClass struct {
	objc.Class
}

type IPDFInfo interface {
	objc.IObject
	URL() foundation.URL
	SetURL(value foundation.IURL)
	IsFileExtensionHidden() bool
	SetFileExtensionHidden(value bool)
	TagNames() []string
	SetTagNames(value []string)
	Orientation() PaperOrientation
	SetOrientation(value PaperOrientation)
	PaperSize() foundation.Size
	SetPaperSize(value foundation.Size)
	Attributes() foundation.MutableDictionary
}

type PDFInfo struct {
	objc.Object
}

func MakePDFInfo(ptr unsafe.Pointer) PDFInfo {
	return PDFInfo{
		Object: objc.MakeObject(ptr),
	}
}

func (pc _PDFInfoClass) Alloc() PDFInfo {
	rv := objc.CallMethod[PDFInfo](pc, objc.GetSelector("alloc"))
	return rv
}

func PDFInfo_Alloc() PDFInfo {
	return PDFInfoClass.Alloc()
}

func (pc _PDFInfoClass) New() PDFInfo {
	rv := objc.CallMethod[PDFInfo](pc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewPDFInfo() PDFInfo {
	return PDFInfoClass.New()
}

func PDFInfo_New() PDFInfo {
	return PDFInfoClass.New()
}

func (p_ PDFInfo) Init() PDFInfo {
	rv := objc.CallMethod[PDFInfo](p_, objc.GetSelector("init"))
	return rv
}

func PDFInfo_Init() PDFInfo {
	return PDFInfoClass.Alloc().Init()
}

func (p_ PDFInfo) URL() foundation.URL {
	rv := objc.CallMethod[foundation.URL](p_, objc.GetSelector("URL"))
	return rv
}

func (p_ PDFInfo) SetURL(value foundation.IURL) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setURL:"), objc.ExtractPtr(value))
}

func (p_ PDFInfo) IsFileExtensionHidden() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("isFileExtensionHidden"))
	return rv
}

func (p_ PDFInfo) SetFileExtensionHidden(value bool) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setFileExtensionHidden:"), value)
}

func (p_ PDFInfo) TagNames() []string {
	rv := objc.CallMethod[[]string](p_, objc.GetSelector("tagNames"))
	return rv
}

func (p_ PDFInfo) SetTagNames(value []string) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setTagNames:"), value)
}

func (p_ PDFInfo) Orientation() PaperOrientation {
	rv := objc.CallMethod[PaperOrientation](p_, objc.GetSelector("orientation"))
	return rv
}

func (p_ PDFInfo) SetOrientation(value PaperOrientation) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setOrientation:"), value)
}

func (p_ PDFInfo) PaperSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](p_, objc.GetSelector("paperSize"))
	return rv
}

func (p_ PDFInfo) SetPaperSize(value foundation.Size) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setPaperSize:"), value)
}

func (p_ PDFInfo) Attributes() foundation.MutableDictionary {
	rv := objc.CallMethod[foundation.MutableDictionary](p_, objc.GetSelector("attributes"))
	return rv
}
