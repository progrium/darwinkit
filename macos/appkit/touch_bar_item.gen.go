// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var TouchBarItemClass = _TouchBarItemClass{objc.GetClass("NSTouchBarItem")}

type _TouchBarItemClass struct {
	objc.Class
}

type ITouchBarItem interface {
	objc.IObject
	Identifier() TouchBarItemIdentifier
	VisibilityPriority() TouchBarItemPriority
	SetVisibilityPriority(value TouchBarItemPriority)
	IsVisible() bool
	CustomizationLabel() string
	ViewController() ViewController
	View() View
}

type TouchBarItem struct {
	objc.Object
}

func MakeTouchBarItem(ptr unsafe.Pointer) TouchBarItem {
	return TouchBarItem{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TouchBarItem) InitWithIdentifier(identifier TouchBarItemIdentifier) TouchBarItem {
	rv := objc.CallMethod[TouchBarItem](t_, objc.GetSelector("initWithIdentifier:"), identifier)
	return rv
}

func TouchBarItem_InitWithIdentifier(identifier TouchBarItemIdentifier) TouchBarItem {
	return TouchBarItemClass.Alloc().InitWithIdentifier(identifier)
}

func (tc _TouchBarItemClass) Alloc() TouchBarItem {
	rv := objc.CallMethod[TouchBarItem](tc, objc.GetSelector("alloc"))
	return rv
}

func TouchBarItem_Alloc() TouchBarItem {
	return TouchBarItemClass.Alloc()
}

func (tc _TouchBarItemClass) New() TouchBarItem {
	rv := objc.CallMethod[TouchBarItem](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTouchBarItem() TouchBarItem {
	return TouchBarItemClass.New()
}

func TouchBarItem_New() TouchBarItem {
	return TouchBarItemClass.New()
}

func (t_ TouchBarItem) Init() TouchBarItem {
	rv := objc.CallMethod[TouchBarItem](t_, objc.GetSelector("init"))
	return rv
}

func TouchBarItem_Init() TouchBarItem {
	return TouchBarItemClass.Alloc().Init()
}

func (t_ TouchBarItem) Identifier() TouchBarItemIdentifier {
	rv := objc.CallMethod[TouchBarItemIdentifier](t_, objc.GetSelector("identifier"))
	return rv
}

func (t_ TouchBarItem) VisibilityPriority() TouchBarItemPriority {
	rv := objc.CallMethod[TouchBarItemPriority](t_, objc.GetSelector("visibilityPriority"))
	return rv
}

func (t_ TouchBarItem) SetVisibilityPriority(value TouchBarItemPriority) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setVisibilityPriority:"), value)
}

func (t_ TouchBarItem) IsVisible() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isVisible"))
	return rv
}

func (t_ TouchBarItem) CustomizationLabel() string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("customizationLabel"))
	return rv
}

func (t_ TouchBarItem) ViewController() ViewController {
	rv := objc.CallMethod[ViewController](t_, objc.GetSelector("viewController"))
	return rv
}

func (t_ TouchBarItem) View() View {
	rv := objc.CallMethod[View](t_, objc.GetSelector("view"))
	return rv
}
