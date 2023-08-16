// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PrintOperation] class.
var PrintOperationClass = _PrintOperationClass{objc.GetClass("NSPrintOperation")}

type _PrintOperationClass struct {
	objc.Class
}

// An interface definition for the [PrintOperation] class.
type IPrintOperation interface {
	objc.IObject
	DestroyContext()
	DeliverResult() bool
	RunOperation() bool
	CleanUpOperation()
	CreateContext() GraphicsContext
	RunOperationModalForWindowDelegateDidRunSelectorContextInfo(docWindow IWindow, delegate objc.IObject, didRunSelector objc.Selector, contextInfo unsafe.Pointer)
	PrintInfo() PrintInfo
	SetPrintInfo(value IPrintInfo)
	PreferredRenderingQuality() PrintRenderingQuality
	PageOrder() PrintingPageOrder
	SetPageOrder(value PrintingPageOrder)
	PDFPanel() PDFPanel
	SetPDFPanel(value IPDFPanel)
	ShowsPrintPanel() bool
	SetShowsPrintPanel(value bool)
	View() View
	CurrentPage() int
	JobTitle() string
	SetJobTitle(value string)
	ShowsProgressPanel() bool
	SetShowsProgressPanel(value bool)
	Context() GraphicsContext
	PrintPanel() PrintPanel
	SetPrintPanel(value IPrintPanel)
	PageRange() foundation.Range
	IsCopyingOperation() bool
	CanSpawnSeparateThread() bool
	SetCanSpawnSeparateThread(value bool)
}

// An object that controls operations that generate Encapsulated PostScript (EPS) code, Portable Document Format (PDF) code, or print jobs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation?language=objc
type PrintOperation struct {
	objc.Object
}

func PrintOperationFrom(ptr unsafe.Pointer) PrintOperation {
	return PrintOperation{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PrintOperationClass) Alloc() PrintOperation {
	rv := objc.Call[PrintOperation](pc, objc.Sel("alloc"))
	return rv
}

func PrintOperation_Alloc() PrintOperation {
	return PrintOperationClass.Alloc()
}

func (pc _PrintOperationClass) New() PrintOperation {
	rv := objc.Call[PrintOperation](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPrintOperation() PrintOperation {
	return PrintOperationClass.New()
}

func (p_ PrintOperation) Init() PrintOperation {
	rv := objc.Call[PrintOperation](p_, objc.Sel("init"))
	return rv
}

// Destroys the print operation’s graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1529896-destroycontext?language=objc
func (p_ PrintOperation) DestroyContext() {
	objc.Call[objc.Void](p_, objc.Sel("destroyContext"))
}

// Creates and returns an print operation object ready to control the printing of the specified view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1535893-printoperationwithview?language=objc
func (pc _PrintOperationClass) PrintOperationWithView(view IView) PrintOperation {
	rv := objc.Call[PrintOperation](pc, objc.Sel("printOperationWithView:"), objc.Ptr(view))
	return rv
}

// Creates and returns an print operation object ready to control the printing of the specified view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1535893-printoperationwithview?language=objc
func PrintOperation_PrintOperationWithView(view IView) PrintOperation {
	return PrintOperationClass.PrintOperationWithView(view)
}

// Delivers the results of the print operation to the intended destination. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1530693-deliverresult?language=objc
func (p_ PrintOperation) DeliverResult() bool {
	rv := objc.Call[bool](p_, objc.Sel("deliverResult"))
	return rv
}

// Creates and returns a new print operation object ready to control the copying of EPS graphics from the specified view and write the resulting data to the specified file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1530037-epsoperationwithview?language=objc
func (pc _PrintOperationClass) EPSOperationWithViewInsideRectToPathPrintInfo(view IView, rect foundation.Rect, path string, printInfo IPrintInfo) PrintOperation {
	rv := objc.Call[PrintOperation](pc, objc.Sel("EPSOperationWithView:insideRect:toPath:printInfo:"), objc.Ptr(view), rect, path, objc.Ptr(printInfo))
	return rv
}

// Creates and returns a new print operation object ready to control the copying of EPS graphics from the specified view and write the resulting data to the specified file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1530037-epsoperationwithview?language=objc
func PrintOperation_EPSOperationWithViewInsideRectToPathPrintInfo(view IView, rect foundation.Rect, path string, printInfo IPrintInfo) PrintOperation {
	return PrintOperationClass.EPSOperationWithViewInsideRectToPathPrintInfo(view, rect, path, printInfo)
}

// Runs the print operation on the current thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1532039-runoperation?language=objc
func (p_ PrintOperation) RunOperation() bool {
	rv := objc.Call[bool](p_, objc.Sel("runOperation"))
	return rv
}

// Creates and returns a new print operation object ready to control the copying of PDF graphics from the specified view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1529269-pdfoperationwithview?language=objc
func (pc _PrintOperationClass) PDFOperationWithViewInsideRectToData(view IView, rect foundation.Rect, data foundation.IMutableData) PrintOperation {
	rv := objc.Call[PrintOperation](pc, objc.Sel("PDFOperationWithView:insideRect:toData:"), objc.Ptr(view), rect, objc.Ptr(data))
	return rv
}

// Creates and returns a new print operation object ready to control the copying of PDF graphics from the specified view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1529269-pdfoperationwithview?language=objc
func PrintOperation_PDFOperationWithViewInsideRectToData(view IView, rect foundation.Rect, data foundation.IMutableData) PrintOperation {
	return PrintOperationClass.PDFOperationWithViewInsideRectToData(view, rect, data)
}

// Called at the end of a print operation to remove the print operation as the current operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1534126-cleanupoperation?language=objc
func (p_ PrintOperation) CleanUpOperation() {
	objc.Call[objc.Void](p_, objc.Sel("cleanUpOperation"))
}

// Creates the graphics context object used for drawing during the operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1526832-createcontext?language=objc
func (p_ PrintOperation) CreateContext() GraphicsContext {
	rv := objc.Call[GraphicsContext](p_, objc.Sel("createContext"))
	return rv
}

// Runs the print operation, calling your custom delegate method upon completion. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1532065-runoperationmodalforwindow?language=objc
func (p_ PrintOperation) RunOperationModalForWindowDelegateDidRunSelectorContextInfo(docWindow IWindow, delegate objc.IObject, didRunSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.Call[objc.Void](p_, objc.Sel("runOperationModalForWindow:delegate:didRunSelector:contextInfo:"), objc.Ptr(docWindow), delegate, didRunSelector, contextInfo)
}

// The printing information associated with the print operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1535187-printinfo?language=objc
func (p_ PrintOperation) PrintInfo() PrintInfo {
	rv := objc.Call[PrintInfo](p_, objc.Sel("printInfo"))
	return rv
}

// The printing information associated with the print operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1535187-printinfo?language=objc
func (p_ PrintOperation) SetPrintInfo(value IPrintInfo) {
	objc.Call[objc.Void](p_, objc.Sel("setPrintInfo:"), objc.Ptr(value))
}

// The printing quality. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1529716-preferredrenderingquality?language=objc
func (p_ PrintOperation) PreferredRenderingQuality() PrintRenderingQuality {
	rv := objc.Call[PrintRenderingQuality](p_, objc.Sel("preferredRenderingQuality"))
	return rv
}

// The print order for the pages of the operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1532990-pageorder?language=objc
func (p_ PrintOperation) PageOrder() PrintingPageOrder {
	rv := objc.Call[PrintingPageOrder](p_, objc.Sel("pageOrder"))
	return rv
}

// The print order for the pages of the operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1532990-pageorder?language=objc
func (p_ PrintOperation) SetPageOrder(value PrintingPageOrder) {
	objc.Call[objc.Void](p_, objc.Sel("setPageOrder:"), value)
}

// The PDF panel object to use during the operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1526838-pdfpanel?language=objc
func (p_ PrintOperation) PDFPanel() PDFPanel {
	rv := objc.Call[PDFPanel](p_, objc.Sel("PDFPanel"))
	return rv
}

// The PDF panel object to use during the operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1526838-pdfpanel?language=objc
func (p_ PrintOperation) SetPDFPanel(value IPDFPanel) {
	objc.Call[objc.Void](p_, objc.Sel("setPDFPanel:"), objc.Ptr(value))
}

// A Boolean value that determines whether the print operation displays a print panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1530666-showsprintpanel?language=objc
func (p_ PrintOperation) ShowsPrintPanel() bool {
	rv := objc.Call[bool](p_, objc.Sel("showsPrintPanel"))
	return rv
}

// A Boolean value that determines whether the print operation displays a print panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1530666-showsprintpanel?language=objc
func (p_ PrintOperation) SetShowsPrintPanel(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setShowsPrintPanel:"), value)
}

// The view object that generates the actual data for the print operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1530311-view?language=objc
func (p_ PrintOperation) View() View {
	rv := objc.Call[View](p_, objc.Sel("view"))
	return rv
}

// The current page number being printed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1534881-currentpage?language=objc
func (p_ PrintOperation) CurrentPage() int {
	rv := objc.Call[int](p_, objc.Sel("currentPage"))
	return rv
}

// The custom title of the print job. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1535322-jobtitle?language=objc
func (p_ PrintOperation) JobTitle() string {
	rv := objc.Call[string](p_, objc.Sel("jobTitle"))
	return rv
}

// The custom title of the print job. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1535322-jobtitle?language=objc
func (p_ PrintOperation) SetJobTitle(value string) {
	objc.Call[objc.Void](p_, objc.Sel("setJobTitle:"), value)
}

// A Boolean value that determines whether the print operation displays a progress panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1535665-showsprogresspanel?language=objc
func (p_ PrintOperation) ShowsProgressPanel() bool {
	rv := objc.Call[bool](p_, objc.Sel("showsProgressPanel"))
	return rv
}

// A Boolean value that determines whether the print operation displays a progress panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1535665-showsprogresspanel?language=objc
func (p_ PrintOperation) SetShowsProgressPanel(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setShowsProgressPanel:"), value)
}

// The graphics context object used for generating output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1534162-context?language=objc
func (p_ PrintOperation) Context() GraphicsContext {
	rv := objc.Call[GraphicsContext](p_, objc.Sel("context"))
	return rv
}

// The print panel object to use during the operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1529626-printpanel?language=objc
func (p_ PrintOperation) PrintPanel() PrintPanel {
	rv := objc.Call[PrintPanel](p_, objc.Sel("printPanel"))
	return rv
}

// The print panel object to use during the operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1529626-printpanel?language=objc
func (p_ PrintOperation) SetPrintPanel(value IPrintPanel) {
	objc.Call[objc.Void](p_, objc.Sel("setPrintPanel:"), objc.Ptr(value))
}

// The range of pages associated with the print operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1524601-pagerange?language=objc
func (p_ PrintOperation) PageRange() foundation.Range {
	rv := objc.Call[foundation.Range](p_, objc.Sel("pageRange"))
	return rv
}

// A Boolean value that indicates whether the print operation is an EPS or PDF copy operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1534206-copyingoperation?language=objc
func (p_ PrintOperation) IsCopyingOperation() bool {
	rv := objc.Call[bool](p_, objc.Sel("isCopyingOperation"))
	return rv
}

// The current print operation for this thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1524255-currentoperation?language=objc
func (pc _PrintOperationClass) CurrentOperation() PrintOperation {
	rv := objc.Call[PrintOperation](pc, objc.Sel("currentOperation"))
	return rv
}

// The current print operation for this thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1524255-currentoperation?language=objc
func PrintOperation_CurrentOperation() PrintOperation {
	return PrintOperationClass.CurrentOperation()
}

// The current print operation for this thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1524255-currentoperation?language=objc
func (pc _PrintOperationClass) SetCurrentOperation(value IPrintOperation) {
	objc.Call[objc.Void](pc, objc.Sel("setCurrentOperation:"), objc.Ptr(value))
}

// The current print operation for this thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1524255-currentoperation?language=objc
func PrintOperation_SetCurrentOperation(value IPrintOperation) {
	PrintOperationClass.SetCurrentOperation(value)
}

// A Boolean value that determines whether the print operation is allowed to spawn a separate printing thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1532487-canspawnseparatethread?language=objc
func (p_ PrintOperation) CanSpawnSeparateThread() bool {
	rv := objc.Call[bool](p_, objc.Sel("canSpawnSeparateThread"))
	return rv
}

// A Boolean value that determines whether the print operation is allowed to spawn a separate printing thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/1532487-canspawnseparatethread?language=objc
func (p_ PrintOperation) SetCanSpawnSeparateThread(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setCanSpawnSeparateThread:"), value)
}
