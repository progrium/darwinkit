// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var WindowTabGroupClass = _WindowTabGroupClass{objc.GetClass("NSWindowTabGroup")}

type _WindowTabGroupClass struct {
	objc.Class
}

type IWindowTabGroup interface {
	objc.IObject
	AddWindow(window IWindow)
	InsertWindowAtIndex(window IWindow, index int)
	RemoveWindow(window IWindow)
	Identifier() WindowTabbingIdentifier
	IsOverviewVisible() bool
	SetOverviewVisible(value bool)
	IsTabBarVisible() bool
	Windows() []Window
	SelectedWindow() Window
	SetSelectedWindow(value IWindow)
}

type WindowTabGroup struct {
	objc.Object
}

func MakeWindowTabGroup(ptr unsafe.Pointer) WindowTabGroup {
	return WindowTabGroup{
		Object: objc.MakeObject(ptr),
	}
}

func (wc _WindowTabGroupClass) Alloc() WindowTabGroup {
	rv := objc.CallMethod[WindowTabGroup](wc, objc.GetSelector("alloc"))
	return rv
}

func WindowTabGroup_Alloc() WindowTabGroup {
	return WindowTabGroupClass.Alloc()
}

func (wc _WindowTabGroupClass) New() WindowTabGroup {
	rv := objc.CallMethod[WindowTabGroup](wc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewWindowTabGroup() WindowTabGroup {
	return WindowTabGroupClass.New()
}

func WindowTabGroup_New() WindowTabGroup {
	return WindowTabGroupClass.New()
}

func (w_ WindowTabGroup) Init() WindowTabGroup {
	rv := objc.CallMethod[WindowTabGroup](w_, objc.GetSelector("init"))
	return rv
}

func WindowTabGroup_Init() WindowTabGroup {
	return WindowTabGroupClass.Alloc().Init()
}

func (w_ WindowTabGroup) AddWindow(window IWindow) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("addWindow:"), objc.ExtractPtr(window))
}

func (w_ WindowTabGroup) InsertWindowAtIndex(window IWindow, index int) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("insertWindow:atIndex:"), objc.ExtractPtr(window), index)
}

func (w_ WindowTabGroup) RemoveWindow(window IWindow) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("removeWindow:"), objc.ExtractPtr(window))
}

func (w_ WindowTabGroup) Identifier() WindowTabbingIdentifier {
	rv := objc.CallMethod[WindowTabbingIdentifier](w_, objc.GetSelector("identifier"))
	return rv
}

func (w_ WindowTabGroup) IsOverviewVisible() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("isOverviewVisible"))
	return rv
}

func (w_ WindowTabGroup) SetOverviewVisible(value bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setOverviewVisible:"), value)
}

func (w_ WindowTabGroup) IsTabBarVisible() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("isTabBarVisible"))
	return rv
}

func (w_ WindowTabGroup) Windows() []Window {
	rv := objc.CallMethod[[]Window](w_, objc.GetSelector("windows"))
	return rv
}

func (w_ WindowTabGroup) SelectedWindow() Window {
	rv := objc.CallMethod[Window](w_, objc.GetSelector("selectedWindow"))
	return rv
}

func (w_ WindowTabGroup) SetSelectedWindow(value IWindow) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setSelectedWindow:"), objc.ExtractPtr(value))
}
