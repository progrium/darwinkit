// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var PrintPanelClass = _PrintPanelClass{objc.GetClass("NSPrintPanel")}

type _PrintPanelClass struct {
	objc.Class
}

type IPrintPanel interface {
	objc.IObject
	DefaultButtonTitle() string
	SetDefaultButtonTitle(defaultButtonTitle string)
	AddAccessoryController(accessoryController IViewController)
	RemoveAccessoryController(accessoryController IViewController)
	BeginSheetWithPrintInfoModalForWindowDelegateDidEndSelectorContextInfo(printInfo IPrintInfo, docWindow IWindow, delegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer)
	RunModal() int
	RunModalWithPrintInfo(printInfo IPrintInfo) int
	JobStyleHint() PrintPanelJobStyleHint
	SetJobStyleHint(value PrintPanelJobStyleHint)
	Options() PrintPanelOptions
	SetOptions(value PrintPanelOptions)
	HelpAnchor() HelpAnchorName
	SetHelpAnchor(value HelpAnchorName)
	AccessoryControllers() []ViewController
	PrintInfo() PrintInfo
}

type PrintPanel struct {
	objc.Object
}

func MakePrintPanel(ptr unsafe.Pointer) PrintPanel {
	return PrintPanel{
		Object: objc.MakeObject(ptr),
	}
}

func (pc _PrintPanelClass) Alloc() PrintPanel {
	rv := objc.CallMethod[PrintPanel](pc, objc.GetSelector("alloc"))
	return rv
}

func PrintPanel_Alloc() PrintPanel {
	return PrintPanelClass.Alloc()
}

func (pc _PrintPanelClass) New() PrintPanel {
	rv := objc.CallMethod[PrintPanel](pc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewPrintPanel() PrintPanel {
	return PrintPanelClass.New()
}

func PrintPanel_New() PrintPanel {
	return PrintPanelClass.New()
}

func (p_ PrintPanel) Init() PrintPanel {
	rv := objc.CallMethod[PrintPanel](p_, objc.GetSelector("init"))
	return rv
}

func PrintPanel_Init() PrintPanel {
	return PrintPanelClass.Alloc().Init()
}

func (pc _PrintPanelClass) PrintPanel() PrintPanel {
	rv := objc.CallMethod[PrintPanel](pc, objc.GetSelector("printPanel"))
	return rv
}

func PrintPanel_PrintPanel() PrintPanel {
	return PrintPanelClass.PrintPanel()
}

func (p_ PrintPanel) DefaultButtonTitle() string {
	rv := objc.CallMethod[string](p_, objc.GetSelector("defaultButtonTitle"))
	return rv
}

func (p_ PrintPanel) SetDefaultButtonTitle(defaultButtonTitle string) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setDefaultButtonTitle:"), defaultButtonTitle)
}

func (p_ PrintPanel) AddAccessoryController(accessoryController IViewController) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("addAccessoryController:"), objc.ExtractPtr(accessoryController))
}

func (p_ PrintPanel) RemoveAccessoryController(accessoryController IViewController) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("removeAccessoryController:"), objc.ExtractPtr(accessoryController))
}

func (p_ PrintPanel) BeginSheetWithPrintInfoModalForWindowDelegateDidEndSelectorContextInfo(printInfo IPrintInfo, docWindow IWindow, delegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("beginSheetWithPrintInfo:modalForWindow:delegate:didEndSelector:contextInfo:"), objc.ExtractPtr(printInfo), objc.ExtractPtr(docWindow), objc.ExtractPtr(delegate), didEndSelector, contextInfo)
}

func (p_ PrintPanel) RunModal() int {
	rv := objc.CallMethod[int](p_, objc.GetSelector("runModal"))
	return rv
}

func (p_ PrintPanel) RunModalWithPrintInfo(printInfo IPrintInfo) int {
	rv := objc.CallMethod[int](p_, objc.GetSelector("runModalWithPrintInfo:"), objc.ExtractPtr(printInfo))
	return rv
}

func (p_ PrintPanel) JobStyleHint() PrintPanelJobStyleHint {
	rv := objc.CallMethod[PrintPanelJobStyleHint](p_, objc.GetSelector("jobStyleHint"))
	return rv
}

func (p_ PrintPanel) SetJobStyleHint(value PrintPanelJobStyleHint) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setJobStyleHint:"), value)
}

func (p_ PrintPanel) Options() PrintPanelOptions {
	rv := objc.CallMethod[PrintPanelOptions](p_, objc.GetSelector("options"))
	return rv
}

func (p_ PrintPanel) SetOptions(value PrintPanelOptions) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setOptions:"), value)
}

func (p_ PrintPanel) HelpAnchor() HelpAnchorName {
	rv := objc.CallMethod[HelpAnchorName](p_, objc.GetSelector("helpAnchor"))
	return rv
}

func (p_ PrintPanel) SetHelpAnchor(value HelpAnchorName) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setHelpAnchor:"), value)
}

func (p_ PrintPanel) AccessoryControllers() []ViewController {
	rv := objc.CallMethod[[]ViewController](p_, objc.GetSelector("accessoryControllers"))
	// TODO: convert slice items...
	return rv
}

func (p_ PrintPanel) PrintInfo() PrintInfo {
	rv := objc.CallMethod[PrintInfo](p_, objc.GetSelector("printInfo"))
	return rv
}
