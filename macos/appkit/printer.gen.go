// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Printer] class.
var PrinterClass = _PrinterClass{objc.GetClass("NSPrinter")}

type _PrinterClass struct {
	objc.Class
}

// An interface definition for the [Printer] class.
type IPrinter interface {
	objc.IObject
	PageSizeForPaper(paperName PrinterPaperName) foundation.Size
	Name() string
	DeviceDescription() map[DeviceDescriptionKey]objc.Object
	LanguageLevel() int
	Type() PrinterTypeName
}

// An object that describes a printer’s capabilities. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprinter?language=objc
type Printer struct {
	objc.Object
}

func PrinterFrom(ptr unsafe.Pointer) Printer {
	return Printer{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PrinterClass) Alloc() Printer {
	rv := objc.Call[Printer](pc, objc.Sel("alloc"))
	return rv
}

func Printer_Alloc() Printer {
	return PrinterClass.Alloc()
}

func (pc _PrinterClass) New() Printer {
	rv := objc.Call[Printer](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPrinter() Printer {
	return PrinterClass.New()
}

func (p_ Printer) Init() Printer {
	rv := objc.Call[Printer](p_, objc.Sel("init"))
	return rv
}

// Creates and returns a printer object initialized to the first available printer with the specified make and model information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprinter/1525196-printerwithtype?language=objc
func (pc _PrinterClass) PrinterWithType(type_ PrinterTypeName) Printer {
	rv := objc.Call[Printer](pc, objc.Sel("printerWithType:"), type_)
	return rv
}

// Creates and returns a printer object initialized to the first available printer with the specified make and model information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprinter/1525196-printerwithtype?language=objc
func Printer_PrinterWithType(type_ PrinterTypeName) Printer {
	return PrinterClass.PrinterWithType(type_)
}

// Returns the size of the page for the specified paper type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprinter/1525222-pagesizeforpaper?language=objc
func (p_ Printer) PageSizeForPaper(paperName PrinterPaperName) foundation.Size {
	rv := objc.Call[foundation.Size](p_, objc.Sel("pageSizeForPaper:"), paperName)
	return rv
}

// Returns the names of all available printers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprinter/1525209-printernames?language=objc
func (pc _PrinterClass) PrinterNames() []string {
	rv := objc.Call[[]string](pc, objc.Sel("printerNames"))
	return rv
}

// Returns the names of all available printers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprinter/1525209-printernames?language=objc
func Printer_PrinterNames() []string {
	return PrinterClass.PrinterNames()
}

// The printer’s name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprinter/1525189-name?language=objc
func (p_ Printer) Name() string {
	rv := objc.Call[string](p_, objc.Sel("name"))
	return rv
}

// A dictionary of keys and values that describe the device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprinter/1525201-devicedescription?language=objc
func (p_ Printer) DeviceDescription() map[DeviceDescriptionKey]objc.Object {
	rv := objc.Call[map[DeviceDescriptionKey]objc.Object](p_, objc.Sel("deviceDescription"))
	return rv
}

// The PostScript language level recognized by the printer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprinter/1525215-languagelevel?language=objc
func (p_ Printer) LanguageLevel() int {
	rv := objc.Call[int](p_, objc.Sel("languageLevel"))
	return rv
}

// A description of the printer’s make and model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprinter/1524450-type?language=objc
func (p_ Printer) Type() PrinterTypeName {
	rv := objc.Call[PrinterTypeName](p_, objc.Sel("type"))
	return rv
}

// Returns descriptions of the makes and models of all available printers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprinter/1525216-printertypes?language=objc
func (pc _PrinterClass) PrinterTypes() []PrinterTypeName {
	rv := objc.Call[[]PrinterTypeName](pc, objc.Sel("printerTypes"))
	return rv
}

// Returns descriptions of the makes and models of all available printers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprinter/1525216-printertypes?language=objc
func Printer_PrinterTypes() []PrinterTypeName {
	return PrinterClass.PrinterTypes()
}
