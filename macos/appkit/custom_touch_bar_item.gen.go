// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CustomTouchBarItem] class.
var CustomTouchBarItemClass = _CustomTouchBarItemClass{objc.GetClass("NSCustomTouchBarItem")}

type _CustomTouchBarItemClass struct {
	objc.Class
}

// An interface definition for the [CustomTouchBarItem] class.
type ICustomTouchBarItem interface {
	ITouchBarItem
	SetCustomizationLabel(value string)
	SetView(value IView)
	SetViewController(value IViewController)
}

// A bar item that contains a responder of your choice, such as a view, a button, or a scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscustomtouchbaritem?language=objc
type CustomTouchBarItem struct {
	TouchBarItem
}

func CustomTouchBarItemFrom(ptr unsafe.Pointer) CustomTouchBarItem {
	return CustomTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}

func (cc _CustomTouchBarItemClass) Alloc() CustomTouchBarItem {
	rv := objc.Call[CustomTouchBarItem](cc, objc.Sel("alloc"))
	return rv
}

func CustomTouchBarItem_Alloc() CustomTouchBarItem {
	return CustomTouchBarItemClass.Alloc()
}

func (cc _CustomTouchBarItemClass) New() CustomTouchBarItem {
	rv := objc.Call[CustomTouchBarItem](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCustomTouchBarItem() CustomTouchBarItem {
	return CustomTouchBarItemClass.New()
}

func (c_ CustomTouchBarItem) Init() CustomTouchBarItem {
	rv := objc.Call[CustomTouchBarItem](c_, objc.Sel("init"))
	return rv
}

func (c_ CustomTouchBarItem) InitWithIdentifier(identifier TouchBarItemIdentifier) CustomTouchBarItem {
	rv := objc.Call[CustomTouchBarItem](c_, objc.Sel("initWithIdentifier:"), identifier)
	return rv
}

// Creates a new item with the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/2544755-initwithidentifier?language=objc
func NewCustomTouchBarItemWithIdentifier(identifier TouchBarItemIdentifier) CustomTouchBarItem {
	instance := CustomTouchBarItemClass.Alloc().InitWithIdentifier(identifier)
	instance.Autorelease()
	return instance
}

// The user-visible string identifying this item during bar customization. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscustomtouchbaritem/2544770-customizationlabel?language=objc
func (c_ CustomTouchBarItem) SetCustomizationLabel(value string) {
	objc.Call[objc.Void](c_, objc.Sel("setCustomizationLabel:"), value)
}

// The view displayed in the bar to represent this item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscustomtouchbaritem/2544813-view?language=objc
func (c_ CustomTouchBarItem) SetView(value IView) {
	objc.Call[objc.Void](c_, objc.Sel("setView:"), objc.Ptr(value))
}

// A view controller whose view is displayed in the bar to represent this item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscustomtouchbaritem/2544745-viewcontroller?language=objc
func (c_ CustomTouchBarItem) SetViewController(value IViewController) {
	objc.Call[objc.Void](c_, objc.Sel("setViewController:"), objc.Ptr(value))
}
