// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var TabViewItemClass = _TabViewItemClass{objc.GetClass("NSTabViewItem")}

type _TabViewItemClass struct {
	objc.Class
}

type ITabViewItem interface {
	objc.IObject
	DrawLabelInRect(shouldTruncateLabel bool, labelRect foundation.Rect)
	SizeOfLabel(computeMin bool) foundation.Size
	Label() string
	SetLabel(value string)
	TabState() TabState
	Identifier() objc.Object
	SetIdentifier(value objc.IObject)
	Color() Color
	SetColor(value IColor)
	View() View
	SetView(value IView)
	InitialFirstResponder() View
	SetInitialFirstResponder(value IView)
	TabView() TabView
	ToolTip() string
	SetToolTip(value string)
	Image() Image
	SetImage(value IImage)
	ViewController() ViewController
	SetViewController(value IViewController)
}

type TabViewItem struct {
	objc.Object
}

func MakeTabViewItem(ptr unsafe.Pointer) TabViewItem {
	return TabViewItem{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TabViewItem) InitWithIdentifier(identifier objc.IObject) TabViewItem {
	rv := objc.CallMethod[TabViewItem](t_, objc.GetSelector("initWithIdentifier:"), objc.ExtractPtr(identifier))
	return rv
}

func TabViewItem_InitWithIdentifier(identifier objc.IObject) TabViewItem {
	return TabViewItemClass.Alloc().InitWithIdentifier(identifier)
}

func (tc _TabViewItemClass) TabViewItemWithViewController(viewController IViewController) TabViewItem {
	rv := objc.CallMethod[TabViewItem](tc, objc.GetSelector("tabViewItemWithViewController:"), objc.ExtractPtr(viewController))
	return rv
}

func TabViewItem_TabViewItemWithViewController(viewController IViewController) TabViewItem {
	return TabViewItemClass.TabViewItemWithViewController(viewController)
}

func (tc _TabViewItemClass) Alloc() TabViewItem {
	rv := objc.CallMethod[TabViewItem](tc, objc.GetSelector("alloc"))
	return rv
}

func TabViewItem_Alloc() TabViewItem {
	return TabViewItemClass.Alloc()
}

func (tc _TabViewItemClass) New() TabViewItem {
	rv := objc.CallMethod[TabViewItem](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTabViewItem() TabViewItem {
	return TabViewItemClass.New()
}

func TabViewItem_New() TabViewItem {
	return TabViewItemClass.New()
}

func (t_ TabViewItem) Init() TabViewItem {
	rv := objc.CallMethod[TabViewItem](t_, objc.GetSelector("init"))
	return rv
}

func TabViewItem_Init() TabViewItem {
	return TabViewItemClass.Alloc().Init()
}

func (t_ TabViewItem) DrawLabelInRect(shouldTruncateLabel bool, labelRect foundation.Rect) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("drawLabel:inRect:"), shouldTruncateLabel, labelRect)
}

func (t_ TabViewItem) SizeOfLabel(computeMin bool) foundation.Size {
	rv := objc.CallMethod[foundation.Size](t_, objc.GetSelector("sizeOfLabel:"), computeMin)
	return rv
}

func (t_ TabViewItem) Label() string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("label"))
	return rv
}

func (t_ TabViewItem) SetLabel(value string) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setLabel:"), value)
}

func (t_ TabViewItem) TabState() TabState {
	rv := objc.CallMethod[TabState](t_, objc.GetSelector("tabState"))
	return rv
}

func (t_ TabViewItem) Identifier() objc.Object {
	rv := objc.CallMethod[objc.Object](t_, objc.GetSelector("identifier"))
	return rv
}

func (t_ TabViewItem) SetIdentifier(value objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setIdentifier:"), objc.ExtractPtr(value))
}

func (t_ TabViewItem) Color() Color {
	rv := objc.CallMethod[Color](t_, objc.GetSelector("color"))
	return rv
}

func (t_ TabViewItem) SetColor(value IColor) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setColor:"), objc.ExtractPtr(value))
}

func (t_ TabViewItem) View() View {
	rv := objc.CallMethod[View](t_, objc.GetSelector("view"))
	return rv
}

func (t_ TabViewItem) SetView(value IView) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setView:"), objc.ExtractPtr(value))
}

func (t_ TabViewItem) InitialFirstResponder() View {
	rv := objc.CallMethod[View](t_, objc.GetSelector("initialFirstResponder"))
	return rv
}

func (t_ TabViewItem) SetInitialFirstResponder(value IView) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setInitialFirstResponder:"), objc.ExtractPtr(value))
}

func (t_ TabViewItem) TabView() TabView {
	rv := objc.CallMethod[TabView](t_, objc.GetSelector("tabView"))
	return rv
}

func (t_ TabViewItem) ToolTip() string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("toolTip"))
	return rv
}

func (t_ TabViewItem) SetToolTip(value string) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setToolTip:"), value)
}

func (t_ TabViewItem) Image() Image {
	rv := objc.CallMethod[Image](t_, objc.GetSelector("image"))
	return rv
}

func (t_ TabViewItem) SetImage(value IImage) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setImage:"), objc.ExtractPtr(value))
}

func (t_ TabViewItem) ViewController() ViewController {
	rv := objc.CallMethod[ViewController](t_, objc.GetSelector("viewController"))
	return rv
}

func (t_ TabViewItem) SetViewController(value IViewController) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setViewController:"), objc.ExtractPtr(value))
}
