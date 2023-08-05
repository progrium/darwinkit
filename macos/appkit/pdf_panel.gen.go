// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var PDFPanelClass = _PDFPanelClass{objc.GetClass("NSPDFPanel")}

type _PDFPanelClass struct {
	objc.Class
}

type IPDFPanel interface {
	objc.IObject
	BeginSheetWithPDFInfoModalForWindowCompletionHandler(pdfInfo IPDFInfo, docWindow IWindow, completionHandler func(param1 int))
	AccessoryController() ViewController
	SetAccessoryController(value IViewController)
	Options() PDFPanelOptions
	SetOptions(value PDFPanelOptions)
	DefaultFileName() string
	SetDefaultFileName(value string)
}

type PDFPanel struct {
	objc.Object
}

func MakePDFPanel(ptr unsafe.Pointer) PDFPanel {
	return PDFPanel{
		Object: objc.MakeObject(ptr),
	}
}

func (pc _PDFPanelClass) Alloc() PDFPanel {
	rv := objc.CallMethod[PDFPanel](pc, objc.GetSelector("alloc"))
	return rv
}

func PDFPanel_Alloc() PDFPanel {
	return PDFPanelClass.Alloc()
}

func (pc _PDFPanelClass) New() PDFPanel {
	rv := objc.CallMethod[PDFPanel](pc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewPDFPanel() PDFPanel {
	return PDFPanelClass.New()
}

func PDFPanel_New() PDFPanel {
	return PDFPanelClass.New()
}

func (p_ PDFPanel) Init() PDFPanel {
	rv := objc.CallMethod[PDFPanel](p_, objc.GetSelector("init"))
	return rv
}

func PDFPanel_Init() PDFPanel {
	return PDFPanelClass.Alloc().Init()
}

func (pc _PDFPanelClass) Panel() PDFPanel {
	rv := objc.CallMethod[PDFPanel](pc, objc.GetSelector("panel"))
	return rv
}

func PDFPanel_Panel() PDFPanel {
	return PDFPanelClass.Panel()
}

func (p_ PDFPanel) BeginSheetWithPDFInfoModalForWindowCompletionHandler(pdfInfo IPDFInfo, docWindow IWindow, completionHandler func(param1 int)) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("beginSheetWithPDFInfo:modalForWindow:completionHandler:"), objc.ExtractPtr(pdfInfo), objc.ExtractPtr(docWindow), completionHandler)
}

func (p_ PDFPanel) AccessoryController() ViewController {
	rv := objc.CallMethod[ViewController](p_, objc.GetSelector("accessoryController"))
	return rv
}

func (p_ PDFPanel) SetAccessoryController(value IViewController) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setAccessoryController:"), objc.ExtractPtr(value))
}

func (p_ PDFPanel) Options() PDFPanelOptions {
	rv := objc.CallMethod[PDFPanelOptions](p_, objc.GetSelector("options"))
	return rv
}

func (p_ PDFPanel) SetOptions(value PDFPanelOptions) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setOptions:"), value)
}

func (p_ PDFPanel) DefaultFileName() string {
	rv := objc.CallMethod[string](p_, objc.GetSelector("defaultFileName"))
	return rv
}

func (p_ PDFPanel) SetDefaultFileName(value string) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setDefaultFileName:"), value)
}
