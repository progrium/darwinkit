// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var PrinterClass = _PrinterClass{objc.GetClass("NSPrinter")}

type _PrinterClass struct {
	objc.Class
}

type IPrinter interface {
	objc.IObject
	PageSizeForPaper(paperName PrinterPaperName) foundation.Size
	Name() string
	Type() PrinterTypeName
	LanguageLevel() int
	DeviceDescription() map[DeviceDescriptionKey]objc.Object
}

type Printer struct {
	objc.Object
}

func MakePrinter(ptr unsafe.Pointer) Printer {
	return Printer{
		Object: objc.MakeObject(ptr),
	}
}

func (pc _PrinterClass) Alloc() Printer {
	rv := objc.CallMethod[Printer](pc, objc.GetSelector("alloc"))
	return rv
}

func Printer_Alloc() Printer {
	return PrinterClass.Alloc()
}

func (pc _PrinterClass) New() Printer {
	rv := objc.CallMethod[Printer](pc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewPrinter() Printer {
	return PrinterClass.New()
}

func Printer_New() Printer {
	return PrinterClass.New()
}

func (p_ Printer) Init() Printer {
	rv := objc.CallMethod[Printer](p_, objc.GetSelector("init"))
	return rv
}

func Printer_Init() Printer {
	return PrinterClass.Alloc().Init()
}

func (pc _PrinterClass) PrinterWithName(name string) Printer {
	rv := objc.CallMethod[Printer](pc, objc.GetSelector("printerWithName:"), name)
	return rv
}

func Printer_PrinterWithName(name string) Printer {
	return PrinterClass.PrinterWithName(name)
}

func (pc _PrinterClass) PrinterWithType(type_ PrinterTypeName) Printer {
	rv := objc.CallMethod[Printer](pc, objc.GetSelector("printerWithType:"), type_)
	return rv
}

func Printer_PrinterWithType(type_ PrinterTypeName) Printer {
	return PrinterClass.PrinterWithType(type_)
}

func (p_ Printer) PageSizeForPaper(paperName PrinterPaperName) foundation.Size {
	rv := objc.CallMethod[foundation.Size](p_, objc.GetSelector("pageSizeForPaper:"), paperName)
	return rv
}

func (pc _PrinterClass) PrinterNames() []string {
	rv := objc.CallMethod[[]string](pc, objc.GetSelector("printerNames"))
	return rv
}

func Printer_PrinterNames() []string {
	return PrinterClass.PrinterNames()
}

func (pc _PrinterClass) PrinterTypes() []PrinterTypeName {
	rv := objc.CallMethod[[]PrinterTypeName](pc, objc.GetSelector("printerTypes"))
	return rv
}

func Printer_PrinterTypes() []PrinterTypeName {
	return PrinterClass.PrinterTypes()
}

func (p_ Printer) Name() string {
	rv := objc.CallMethod[string](p_, objc.GetSelector("name"))
	return rv
}

func (p_ Printer) Type() PrinterTypeName {
	rv := objc.CallMethod[PrinterTypeName](p_, objc.GetSelector("type"))
	return rv
}

func (p_ Printer) LanguageLevel() int {
	rv := objc.CallMethod[int](p_, objc.GetSelector("languageLevel"))
	return rv
}

func (p_ Printer) DeviceDescription() map[DeviceDescriptionKey]objc.Object {
	rv := objc.CallMethod[map[DeviceDescriptionKey]objc.Object](p_, objc.GetSelector("deviceDescription"))
	return rv
}
