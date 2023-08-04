// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var PageLayoutClass = _PageLayoutClass{objc.GetClass("NSPageLayout")}

type _PageLayoutClass struct {
	objc.Class
}

type IPageLayout interface {
	objc.IObject
	BeginSheetWithPrintInfoModalForWindowDelegateDidEndSelectorContextInfo(printInfo IPrintInfo, docWindow IWindow, delegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer)
	RunModal() int
	RunModalWithPrintInfo(printInfo IPrintInfo) int
	AddAccessoryController(accessoryController IViewController)
	RemoveAccessoryController(accessoryController IViewController)
	AccessoryControllers() []ViewController
	PrintInfo() PrintInfo
}

type PageLayout struct {
	objc.Object
}

func MakePageLayout(ptr unsafe.Pointer) PageLayout {
	return PageLayout{
		Object: objc.MakeObject(ptr),
	}
}

func (pc _PageLayoutClass) Alloc() PageLayout {
	rv := objc.CallMethod[PageLayout](pc, objc.GetSelector("alloc"))
	return rv
}

func PageLayout_Alloc() PageLayout {
	return PageLayoutClass.Alloc()
}

func (pc _PageLayoutClass) New() PageLayout {
	rv := objc.CallMethod[PageLayout](pc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewPageLayout() PageLayout {
	return PageLayoutClass.New()
}

func PageLayout_New() PageLayout {
	return PageLayoutClass.New()
}

func (p_ PageLayout) Init() PageLayout {
	rv := objc.CallMethod[PageLayout](p_, objc.GetSelector("init"))
	return rv
}

func PageLayout_Init() PageLayout {
	return PageLayoutClass.Alloc().Init()
}

func (pc _PageLayoutClass) PageLayout() PageLayout {
	rv := objc.CallMethod[PageLayout](pc, objc.GetSelector("pageLayout"))
	return rv
}

func PageLayout_PageLayout() PageLayout {
	return PageLayoutClass.PageLayout()
}

func (p_ PageLayout) BeginSheetWithPrintInfoModalForWindowDelegateDidEndSelectorContextInfo(printInfo IPrintInfo, docWindow IWindow, delegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("beginSheetWithPrintInfo:modalForWindow:delegate:didEndSelector:contextInfo:"), objc.ExtractPtr(printInfo), objc.ExtractPtr(docWindow), objc.ExtractPtr(delegate), didEndSelector, contextInfo)
}

func (p_ PageLayout) RunModal() int {
	rv := objc.CallMethod[int](p_, objc.GetSelector("runModal"))
	return rv
}

func (p_ PageLayout) RunModalWithPrintInfo(printInfo IPrintInfo) int {
	rv := objc.CallMethod[int](p_, objc.GetSelector("runModalWithPrintInfo:"), objc.ExtractPtr(printInfo))
	return rv
}

func (p_ PageLayout) AddAccessoryController(accessoryController IViewController) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("addAccessoryController:"), objc.ExtractPtr(accessoryController))
}

func (p_ PageLayout) RemoveAccessoryController(accessoryController IViewController) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("removeAccessoryController:"), objc.ExtractPtr(accessoryController))
}

func (p_ PageLayout) AccessoryControllers() []ViewController {
	rv := objc.CallMethod[[]ViewController](p_, objc.GetSelector("accessoryControllers"))
	return rv
}

func (p_ PageLayout) PrintInfo() PrintInfo {
	rv := objc.CallMethod[PrintInfo](p_, objc.GetSelector("printInfo"))
	return rv
}
