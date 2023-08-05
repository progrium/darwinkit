// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var WindowTabClass = _WindowTabClass{objc.GetClass("NSWindowTab")}

type _WindowTabClass struct {
	objc.Class
}

type IWindowTab interface {
	objc.IObject
	Title() string
	SetTitle(value string)
	AttributedTitle() foundation.AttributedString
	SetAttributedTitle(value foundation.IAttributedString)
	ToolTip() string
	SetToolTip(value string)
	AccessoryView() View
	SetAccessoryView(value IView)
}

type WindowTab struct {
	objc.Object
}

func MakeWindowTab(ptr unsafe.Pointer) WindowTab {
	return WindowTab{
		Object: objc.MakeObject(ptr),
	}
}

func (wc _WindowTabClass) Alloc() WindowTab {
	rv := objc.CallMethod[WindowTab](wc, objc.GetSelector("alloc"))
	return rv
}

func WindowTab_Alloc() WindowTab {
	return WindowTabClass.Alloc()
}

func (wc _WindowTabClass) New() WindowTab {
	rv := objc.CallMethod[WindowTab](wc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewWindowTab() WindowTab {
	return WindowTabClass.New()
}

func WindowTab_New() WindowTab {
	return WindowTabClass.New()
}

func (w_ WindowTab) Init() WindowTab {
	rv := objc.CallMethod[WindowTab](w_, objc.GetSelector("init"))
	return rv
}

func WindowTab_Init() WindowTab {
	return WindowTabClass.Alloc().Init()
}

func (w_ WindowTab) Title() string {
	rv := objc.CallMethod[string](w_, objc.GetSelector("title"))
	return rv
}

func (w_ WindowTab) SetTitle(value string) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setTitle:"), value)
}

func (w_ WindowTab) AttributedTitle() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](w_, objc.GetSelector("attributedTitle"))
	return rv
}

func (w_ WindowTab) SetAttributedTitle(value foundation.IAttributedString) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setAttributedTitle:"), objc.ExtractPtr(value))
}

func (w_ WindowTab) ToolTip() string {
	rv := objc.CallMethod[string](w_, objc.GetSelector("toolTip"))
	return rv
}

func (w_ WindowTab) SetToolTip(value string) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setToolTip:"), value)
}

func (w_ WindowTab) AccessoryView() View {
	rv := objc.CallMethod[View](w_, objc.GetSelector("accessoryView"))
	return rv
}

func (w_ WindowTab) SetAccessoryView(value IView) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setAccessoryView:"), objc.ExtractPtr(value))
}
