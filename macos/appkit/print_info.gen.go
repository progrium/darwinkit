// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var PrintInfoClass = _PrintInfoClass{objc.GetClass("NSPrintInfo")}

type _PrintInfoClass struct {
	objc.Class
}

type IPrintInfo interface {
	objc.IObject
	SetUpPrintOperationDefaultValues()
	Dictionary() foundation.MutableDictionary
	PMPrintSession() unsafe.Pointer
	PMPageFormat() unsafe.Pointer
	PMPrintSettings() unsafe.Pointer
	UpdateFromPMPageFormat()
	UpdateFromPMPrintSettings()
	TakeSettingsFromPDFInfo(inPDFInfo IPDFInfo)
	PaperSize() foundation.Size
	SetPaperSize(value foundation.Size)
	TopMargin() float64
	SetTopMargin(value float64)
	BottomMargin() float64
	SetBottomMargin(value float64)
	LeftMargin() float64
	SetLeftMargin(value float64)
	RightMargin() float64
	SetRightMargin(value float64)
	ImageablePageBounds() foundation.Rect
	Orientation() PaperOrientation
	SetOrientation(value PaperOrientation)
	PaperName() PrinterPaperName
	SetPaperName(value PrinterPaperName)
	LocalizedPaperName() string
	HorizontalPagination() PrintingPaginationMode
	SetHorizontalPagination(value PrintingPaginationMode)
	VerticalPagination() PrintingPaginationMode
	SetVerticalPagination(value PrintingPaginationMode)
	IsHorizontallyCentered() bool
	SetHorizontallyCentered(value bool)
	IsVerticallyCentered() bool
	SetVerticallyCentered(value bool)
	Printer() Printer
	SetPrinter(value IPrinter)
	JobDisposition() PrintJobDispositionValue
	SetJobDisposition(value PrintJobDispositionValue)
	IsSelectionOnly() bool
	SetSelectionOnly(value bool)
	ScalingFactor() float64
	SetScalingFactor(value float64)
	PrintSettings() foundation.MutableDictionary
}

type PrintInfo struct {
	objc.Object
}

func MakePrintInfo(ptr unsafe.Pointer) PrintInfo {
	return PrintInfo{
		Object: objc.MakeObject(ptr),
	}
}

func (p_ PrintInfo) InitWithDictionary(attributes map[PrintInfoAttributeKey]objc.IObject) PrintInfo {
	rv := objc.CallMethod[PrintInfo](p_, objc.GetSelector("initWithDictionary:"), attributes)
	return rv
}

func PrintInfo_InitWithDictionary(attributes map[PrintInfoAttributeKey]objc.IObject) PrintInfo {
	return PrintInfoClass.Alloc().InitWithDictionary(attributes)
}

func (p_ PrintInfo) Init() PrintInfo {
	rv := objc.CallMethod[PrintInfo](p_, objc.GetSelector("init"))
	return rv
}

func PrintInfo_Init() PrintInfo {
	return PrintInfoClass.Alloc().Init()
}

func (pc _PrintInfoClass) Alloc() PrintInfo {
	rv := objc.CallMethod[PrintInfo](pc, objc.GetSelector("alloc"))
	return rv
}

func PrintInfo_Alloc() PrintInfo {
	return PrintInfoClass.Alloc()
}

func (pc _PrintInfoClass) New() PrintInfo {
	rv := objc.CallMethod[PrintInfo](pc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewPrintInfo() PrintInfo {
	return PrintInfoClass.New()
}

func PrintInfo_New() PrintInfo {
	return PrintInfoClass.New()
}

func (p_ PrintInfo) SetUpPrintOperationDefaultValues() {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setUpPrintOperationDefaultValues"))
}

func (p_ PrintInfo) Dictionary() foundation.MutableDictionary {
	rv := objc.CallMethod[foundation.MutableDictionary](p_, objc.GetSelector("dictionary"))
	return rv
}

func (p_ PrintInfo) PMPrintSession() unsafe.Pointer {
	rv := objc.CallMethod[unsafe.Pointer](p_, objc.GetSelector("PMPrintSession"))
	return rv
}

func (p_ PrintInfo) PMPageFormat() unsafe.Pointer {
	rv := objc.CallMethod[unsafe.Pointer](p_, objc.GetSelector("PMPageFormat"))
	return rv
}

func (p_ PrintInfo) PMPrintSettings() unsafe.Pointer {
	rv := objc.CallMethod[unsafe.Pointer](p_, objc.GetSelector("PMPrintSettings"))
	return rv
}

func (p_ PrintInfo) UpdateFromPMPageFormat() {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("updateFromPMPageFormat"))
}

func (p_ PrintInfo) UpdateFromPMPrintSettings() {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("updateFromPMPrintSettings"))
}

func (p_ PrintInfo) TakeSettingsFromPDFInfo(inPDFInfo IPDFInfo) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("takeSettingsFromPDFInfo:"), objc.ExtractPtr(inPDFInfo))
}

func (pc _PrintInfoClass) SharedPrintInfo() PrintInfo {
	rv := objc.CallMethod[PrintInfo](pc, objc.GetSelector("sharedPrintInfo"))
	return rv
}

func PrintInfo_SharedPrintInfo() PrintInfo {
	return PrintInfoClass.SharedPrintInfo()
}

func (pc _PrintInfoClass) SetSharedPrintInfo(value IPrintInfo) {
	objc.CallMethod[objc.Void](pc, objc.GetSelector("setSharedPrintInfo:"), objc.ExtractPtr(value))
}

func PrintInfo_SetSharedPrintInfo(value IPrintInfo) {
	PrintInfoClass.SetSharedPrintInfo(value)
}

func (p_ PrintInfo) PaperSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](p_, objc.GetSelector("paperSize"))
	return rv
}

func (p_ PrintInfo) SetPaperSize(value foundation.Size) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setPaperSize:"), value)
}

func (p_ PrintInfo) TopMargin() float64 {
	rv := objc.CallMethod[float64](p_, objc.GetSelector("topMargin"))
	return rv
}

func (p_ PrintInfo) SetTopMargin(value float64) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setTopMargin:"), value)
}

func (p_ PrintInfo) BottomMargin() float64 {
	rv := objc.CallMethod[float64](p_, objc.GetSelector("bottomMargin"))
	return rv
}

func (p_ PrintInfo) SetBottomMargin(value float64) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setBottomMargin:"), value)
}

func (p_ PrintInfo) LeftMargin() float64 {
	rv := objc.CallMethod[float64](p_, objc.GetSelector("leftMargin"))
	return rv
}

func (p_ PrintInfo) SetLeftMargin(value float64) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setLeftMargin:"), value)
}

func (p_ PrintInfo) RightMargin() float64 {
	rv := objc.CallMethod[float64](p_, objc.GetSelector("rightMargin"))
	return rv
}

func (p_ PrintInfo) SetRightMargin(value float64) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setRightMargin:"), value)
}

func (p_ PrintInfo) ImageablePageBounds() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](p_, objc.GetSelector("imageablePageBounds"))
	return rv
}

func (p_ PrintInfo) Orientation() PaperOrientation {
	rv := objc.CallMethod[PaperOrientation](p_, objc.GetSelector("orientation"))
	return rv
}

func (p_ PrintInfo) SetOrientation(value PaperOrientation) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setOrientation:"), value)
}

func (p_ PrintInfo) PaperName() PrinterPaperName {
	rv := objc.CallMethod[PrinterPaperName](p_, objc.GetSelector("paperName"))
	return rv
}

func (p_ PrintInfo) SetPaperName(value PrinterPaperName) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setPaperName:"), value)
}

func (p_ PrintInfo) LocalizedPaperName() string {
	rv := objc.CallMethod[string](p_, objc.GetSelector("localizedPaperName"))
	return rv
}

func (p_ PrintInfo) HorizontalPagination() PrintingPaginationMode {
	rv := objc.CallMethod[PrintingPaginationMode](p_, objc.GetSelector("horizontalPagination"))
	return rv
}

func (p_ PrintInfo) SetHorizontalPagination(value PrintingPaginationMode) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setHorizontalPagination:"), value)
}

func (p_ PrintInfo) VerticalPagination() PrintingPaginationMode {
	rv := objc.CallMethod[PrintingPaginationMode](p_, objc.GetSelector("verticalPagination"))
	return rv
}

func (p_ PrintInfo) SetVerticalPagination(value PrintingPaginationMode) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setVerticalPagination:"), value)
}

func (p_ PrintInfo) IsHorizontallyCentered() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("isHorizontallyCentered"))
	return rv
}

func (p_ PrintInfo) SetHorizontallyCentered(value bool) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setHorizontallyCentered:"), value)
}

func (p_ PrintInfo) IsVerticallyCentered() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("isVerticallyCentered"))
	return rv
}

func (p_ PrintInfo) SetVerticallyCentered(value bool) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setVerticallyCentered:"), value)
}

func (p_ PrintInfo) Printer() Printer {
	rv := objc.CallMethod[Printer](p_, objc.GetSelector("printer"))
	return rv
}

func (p_ PrintInfo) SetPrinter(value IPrinter) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setPrinter:"), objc.ExtractPtr(value))
}

func (p_ PrintInfo) JobDisposition() PrintJobDispositionValue {
	rv := objc.CallMethod[PrintJobDispositionValue](p_, objc.GetSelector("jobDisposition"))
	return rv
}

func (p_ PrintInfo) SetJobDisposition(value PrintJobDispositionValue) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setJobDisposition:"), value)
}

func (p_ PrintInfo) IsSelectionOnly() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("isSelectionOnly"))
	return rv
}

func (p_ PrintInfo) SetSelectionOnly(value bool) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setSelectionOnly:"), value)
}

func (p_ PrintInfo) ScalingFactor() float64 {
	rv := objc.CallMethod[float64](p_, objc.GetSelector("scalingFactor"))
	return rv
}

func (p_ PrintInfo) SetScalingFactor(value float64) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setScalingFactor:"), value)
}

func (p_ PrintInfo) PrintSettings() foundation.MutableDictionary {
	rv := objc.CallMethod[foundation.MutableDictionary](p_, objc.GetSelector("printSettings"))
	return rv
}

func (pc _PrintInfoClass) DefaultPrinter() Printer {
	rv := objc.CallMethod[Printer](pc, objc.GetSelector("defaultPrinter"))
	return rv
}

func PrintInfo_DefaultPrinter() Printer {
	return PrintInfoClass.DefaultPrinter()
}
