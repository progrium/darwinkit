// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var PrintOperationClass = _PrintOperationClass{objc.GetClass("NSPrintOperation")}

type _PrintOperationClass struct {
	objc.Class
}

type IPrintOperation interface {
	objc.IObject
	RunOperation() bool
	RunOperationModalForWindowDelegateDidRunSelectorContextInfo(docWindow IWindow, delegate objc.IObject, didRunSelector objc.Selector, contextInfo unsafe.Pointer)
	CleanUpOperation()
	DeliverResult() bool
	CreateContext() GraphicsContext
	DestroyContext()
	IsCopyingOperation() bool
	PrintInfo() PrintInfo
	SetPrintInfo(value IPrintInfo)
	View() View
	PreferredRenderingQuality() PrintRenderingQuality
	ShowsPrintPanel() bool
	SetShowsPrintPanel(value bool)
	ShowsProgressPanel() bool
	SetShowsProgressPanel(value bool)
	JobTitle() string
	SetJobTitle(value string)
	PrintPanel() PrintPanel
	SetPrintPanel(value IPrintPanel)
	PDFPanel() PDFPanel
	SetPDFPanel(value IPDFPanel)
	Context() GraphicsContext
	CurrentPage() int
	PageRange() foundation.Range
	PageOrder() PrintingPageOrder
	SetPageOrder(value PrintingPageOrder)
	CanSpawnSeparateThread() bool
	SetCanSpawnSeparateThread(value bool)
}

type PrintOperation struct {
	objc.Object
}

func MakePrintOperation(ptr unsafe.Pointer) PrintOperation {
	return PrintOperation{
		Object: objc.MakeObject(ptr),
	}
}

func (pc _PrintOperationClass) Alloc() PrintOperation {
	rv := objc.CallMethod[PrintOperation](pc, objc.GetSelector("alloc"))
	return rv
}

func PrintOperation_Alloc() PrintOperation {
	return PrintOperationClass.Alloc()
}

func (pc _PrintOperationClass) New() PrintOperation {
	rv := objc.CallMethod[PrintOperation](pc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewPrintOperation() PrintOperation {
	return PrintOperationClass.New()
}

func PrintOperation_New() PrintOperation {
	return PrintOperationClass.New()
}

func (p_ PrintOperation) Init() PrintOperation {
	rv := objc.CallMethod[PrintOperation](p_, objc.GetSelector("init"))
	return rv
}

func PrintOperation_Init() PrintOperation {
	return PrintOperationClass.Alloc().Init()
}

func (pc _PrintOperationClass) EPSOperationWithViewInsideRectToData(view IView, rect foundation.Rect, data foundation.IMutableData) PrintOperation {
	rv := objc.CallMethod[PrintOperation](pc, objc.GetSelector("EPSOperationWithView:insideRect:toData:"), objc.ExtractPtr(view), rect, objc.ExtractPtr(data))
	return rv
}

func PrintOperation_EPSOperationWithViewInsideRectToData(view IView, rect foundation.Rect, data foundation.IMutableData) PrintOperation {
	return PrintOperationClass.EPSOperationWithViewInsideRectToData(view, rect, data)
}

func (pc _PrintOperationClass) EPSOperationWithViewInsideRectToDataPrintInfo(view IView, rect foundation.Rect, data foundation.IMutableData, printInfo IPrintInfo) PrintOperation {
	rv := objc.CallMethod[PrintOperation](pc, objc.GetSelector("EPSOperationWithView:insideRect:toData:printInfo:"), objc.ExtractPtr(view), rect, objc.ExtractPtr(data), objc.ExtractPtr(printInfo))
	return rv
}

func PrintOperation_EPSOperationWithViewInsideRectToDataPrintInfo(view IView, rect foundation.Rect, data foundation.IMutableData, printInfo IPrintInfo) PrintOperation {
	return PrintOperationClass.EPSOperationWithViewInsideRectToDataPrintInfo(view, rect, data, printInfo)
}

func (pc _PrintOperationClass) EPSOperationWithViewInsideRectToPathPrintInfo(view IView, rect foundation.Rect, path string, printInfo IPrintInfo) PrintOperation {
	rv := objc.CallMethod[PrintOperation](pc, objc.GetSelector("EPSOperationWithView:insideRect:toPath:printInfo:"), objc.ExtractPtr(view), rect, path, objc.ExtractPtr(printInfo))
	return rv
}

func PrintOperation_EPSOperationWithViewInsideRectToPathPrintInfo(view IView, rect foundation.Rect, path string, printInfo IPrintInfo) PrintOperation {
	return PrintOperationClass.EPSOperationWithViewInsideRectToPathPrintInfo(view, rect, path, printInfo)
}

func (pc _PrintOperationClass) PDFOperationWithViewInsideRectToData(view IView, rect foundation.Rect, data foundation.IMutableData) PrintOperation {
	rv := objc.CallMethod[PrintOperation](pc, objc.GetSelector("PDFOperationWithView:insideRect:toData:"), objc.ExtractPtr(view), rect, objc.ExtractPtr(data))
	return rv
}

func PrintOperation_PDFOperationWithViewInsideRectToData(view IView, rect foundation.Rect, data foundation.IMutableData) PrintOperation {
	return PrintOperationClass.PDFOperationWithViewInsideRectToData(view, rect, data)
}

func (pc _PrintOperationClass) PDFOperationWithViewInsideRectToDataPrintInfo(view IView, rect foundation.Rect, data foundation.IMutableData, printInfo IPrintInfo) PrintOperation {
	rv := objc.CallMethod[PrintOperation](pc, objc.GetSelector("PDFOperationWithView:insideRect:toData:printInfo:"), objc.ExtractPtr(view), rect, objc.ExtractPtr(data), objc.ExtractPtr(printInfo))
	return rv
}

func PrintOperation_PDFOperationWithViewInsideRectToDataPrintInfo(view IView, rect foundation.Rect, data foundation.IMutableData, printInfo IPrintInfo) PrintOperation {
	return PrintOperationClass.PDFOperationWithViewInsideRectToDataPrintInfo(view, rect, data, printInfo)
}

func (pc _PrintOperationClass) PDFOperationWithViewInsideRectToPathPrintInfo(view IView, rect foundation.Rect, path string, printInfo IPrintInfo) PrintOperation {
	rv := objc.CallMethod[PrintOperation](pc, objc.GetSelector("PDFOperationWithView:insideRect:toPath:printInfo:"), objc.ExtractPtr(view), rect, path, objc.ExtractPtr(printInfo))
	return rv
}

func PrintOperation_PDFOperationWithViewInsideRectToPathPrintInfo(view IView, rect foundation.Rect, path string, printInfo IPrintInfo) PrintOperation {
	return PrintOperationClass.PDFOperationWithViewInsideRectToPathPrintInfo(view, rect, path, printInfo)
}

func (pc _PrintOperationClass) PrintOperationWithView(view IView) PrintOperation {
	rv := objc.CallMethod[PrintOperation](pc, objc.GetSelector("printOperationWithView:"), objc.ExtractPtr(view))
	return rv
}

func PrintOperation_PrintOperationWithView(view IView) PrintOperation {
	return PrintOperationClass.PrintOperationWithView(view)
}

func (pc _PrintOperationClass) PrintOperationWithViewPrintInfo(view IView, printInfo IPrintInfo) PrintOperation {
	rv := objc.CallMethod[PrintOperation](pc, objc.GetSelector("printOperationWithView:printInfo:"), objc.ExtractPtr(view), objc.ExtractPtr(printInfo))
	return rv
}

func PrintOperation_PrintOperationWithViewPrintInfo(view IView, printInfo IPrintInfo) PrintOperation {
	return PrintOperationClass.PrintOperationWithViewPrintInfo(view, printInfo)
}

func (p_ PrintOperation) RunOperation() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("runOperation"))
	return rv
}

func (p_ PrintOperation) RunOperationModalForWindowDelegateDidRunSelectorContextInfo(docWindow IWindow, delegate objc.IObject, didRunSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("runOperationModalForWindow:delegate:didRunSelector:contextInfo:"), objc.ExtractPtr(docWindow), objc.ExtractPtr(delegate), didRunSelector, contextInfo)
}

func (p_ PrintOperation) CleanUpOperation() {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("cleanUpOperation"))
}

func (p_ PrintOperation) DeliverResult() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("deliverResult"))
	return rv
}

func (p_ PrintOperation) CreateContext() GraphicsContext {
	rv := objc.CallMethod[GraphicsContext](p_, objc.GetSelector("createContext"))
	return rv
}

func (p_ PrintOperation) DestroyContext() {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("destroyContext"))
}

func (pc _PrintOperationClass) CurrentOperation() PrintOperation {
	rv := objc.CallMethod[PrintOperation](pc, objc.GetSelector("currentOperation"))
	return rv
}

func PrintOperation_CurrentOperation() PrintOperation {
	return PrintOperationClass.CurrentOperation()
}

func (pc _PrintOperationClass) SetCurrentOperation(value IPrintOperation) {
	objc.CallMethod[objc.Void](pc, objc.GetSelector("setCurrentOperation:"), objc.ExtractPtr(value))
}

func PrintOperation_SetCurrentOperation(value IPrintOperation) {
	PrintOperationClass.SetCurrentOperation(value)
}

func (p_ PrintOperation) IsCopyingOperation() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("isCopyingOperation"))
	return rv
}

func (p_ PrintOperation) PrintInfo() PrintInfo {
	rv := objc.CallMethod[PrintInfo](p_, objc.GetSelector("printInfo"))
	return rv
}

func (p_ PrintOperation) SetPrintInfo(value IPrintInfo) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setPrintInfo:"), objc.ExtractPtr(value))
}

func (p_ PrintOperation) View() View {
	rv := objc.CallMethod[View](p_, objc.GetSelector("view"))
	return rv
}

func (p_ PrintOperation) PreferredRenderingQuality() PrintRenderingQuality {
	rv := objc.CallMethod[PrintRenderingQuality](p_, objc.GetSelector("preferredRenderingQuality"))
	return rv
}

func (p_ PrintOperation) ShowsPrintPanel() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("showsPrintPanel"))
	return rv
}

func (p_ PrintOperation) SetShowsPrintPanel(value bool) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setShowsPrintPanel:"), value)
}

func (p_ PrintOperation) ShowsProgressPanel() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("showsProgressPanel"))
	return rv
}

func (p_ PrintOperation) SetShowsProgressPanel(value bool) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setShowsProgressPanel:"), value)
}

func (p_ PrintOperation) JobTitle() string {
	rv := objc.CallMethod[string](p_, objc.GetSelector("jobTitle"))
	return rv
}

func (p_ PrintOperation) SetJobTitle(value string) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setJobTitle:"), value)
}

func (p_ PrintOperation) PrintPanel() PrintPanel {
	rv := objc.CallMethod[PrintPanel](p_, objc.GetSelector("printPanel"))
	return rv
}

func (p_ PrintOperation) SetPrintPanel(value IPrintPanel) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setPrintPanel:"), objc.ExtractPtr(value))
}

func (p_ PrintOperation) PDFPanel() PDFPanel {
	rv := objc.CallMethod[PDFPanel](p_, objc.GetSelector("PDFPanel"))
	return rv
}

func (p_ PrintOperation) SetPDFPanel(value IPDFPanel) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setPDFPanel:"), objc.ExtractPtr(value))
}

func (p_ PrintOperation) Context() GraphicsContext {
	rv := objc.CallMethod[GraphicsContext](p_, objc.GetSelector("context"))
	return rv
}

func (p_ PrintOperation) CurrentPage() int {
	rv := objc.CallMethod[int](p_, objc.GetSelector("currentPage"))
	return rv
}

func (p_ PrintOperation) PageRange() foundation.Range {
	rv := objc.CallMethod[foundation.Range](p_, objc.GetSelector("pageRange"))
	return rv
}

func (p_ PrintOperation) PageOrder() PrintingPageOrder {
	rv := objc.CallMethod[PrintingPageOrder](p_, objc.GetSelector("pageOrder"))
	return rv
}

func (p_ PrintOperation) SetPageOrder(value PrintingPageOrder) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setPageOrder:"), value)
}

func (p_ PrintOperation) CanSpawnSeparateThread() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("canSpawnSeparateThread"))
	return rv
}

func (p_ PrintOperation) SetCanSpawnSeparateThread(value bool) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setCanSpawnSeparateThread:"), value)
}
