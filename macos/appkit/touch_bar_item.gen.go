// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TouchBarItem] class.
var TouchBarItemClass = _TouchBarItemClass{objc.GetClass("NSTouchBarItem")}

type _TouchBarItemClass struct {
	objc.Class
}

// An interface definition for the [TouchBarItem] class.
type ITouchBarItem interface {
	objc.IObject
	IsVisible() bool
	CustomizationLabel() string
	VisibilityPriority() TouchBarItemPriority
	SetVisibilityPriority(value TouchBarItemPriority)
	View() View
	ViewController() ViewController
	Identifier() TouchBarItemIdentifier
}

// A UI control shown in the Touch Bar on supported models of MacBook Pro. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem?language=objc
type TouchBarItem struct {
	objc.Object
}

func TouchBarItemFrom(ptr unsafe.Pointer) TouchBarItem {
	return TouchBarItem{
		Object: objc.ObjectFrom(ptr),
	}
}

func (t_ TouchBarItem) InitWithIdentifier(identifier TouchBarItemIdentifier) TouchBarItem {
	rv := objc.Call[TouchBarItem](t_, objc.Sel("initWithIdentifier:"), identifier)
	return rv
}

// Creates a new item with the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/2544755-initwithidentifier?language=objc
func NewTouchBarItemWithIdentifier(identifier TouchBarItemIdentifier) TouchBarItem {
	instance := TouchBarItemClass.Alloc().InitWithIdentifier(identifier)
	instance.Autorelease()
	return instance
}

func (tc _TouchBarItemClass) Alloc() TouchBarItem {
	rv := objc.Call[TouchBarItem](tc, objc.Sel("alloc"))
	return rv
}

func TouchBarItem_Alloc() TouchBarItem {
	return TouchBarItemClass.Alloc()
}

func (tc _TouchBarItemClass) New() TouchBarItem {
	rv := objc.Call[TouchBarItem](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTouchBarItem() TouchBarItem {
	return TouchBarItemClass.New()
}

func (t_ TouchBarItem) Init() TouchBarItem {
	rv := objc.Call[TouchBarItem](t_, objc.Sel("init"))
	return rv
}

// A Boolean value that reflects whether or not the item is visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/2544669-visible?language=objc
func (t_ TouchBarItem) IsVisible() bool {
	rv := objc.Call[bool](t_, objc.Sel("isVisible"))
	return rv
}

// The user-visible string identifying this item during bar customization. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/2544693-customizationlabel?language=objc
func (t_ TouchBarItem) CustomizationLabel() string {
	rv := objc.Call[string](t_, objc.Sel("customizationLabel"))
	return rv
}

// Determines which items are shown in a bar when space is limited. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/2544689-visibilitypriority?language=objc
func (t_ TouchBarItem) VisibilityPriority() TouchBarItemPriority {
	rv := objc.Call[TouchBarItemPriority](t_, objc.Sel("visibilityPriority"))
	return rv
}

// Determines which items are shown in a bar when space is limited. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/2544689-visibilitypriority?language=objc
func (t_ TouchBarItem) SetVisibilityPriority(value TouchBarItemPriority) {
	objc.Call[objc.Void](t_, objc.Sel("setVisibilityPriority:"), value)
}

// The view associated with this item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/2544719-view?language=objc
func (t_ TouchBarItem) View() View {
	rv := objc.Call[View](t_, objc.Sel("view"))
	return rv
}

// The view controller associated with this item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/2544843-viewcontroller?language=objc
func (t_ TouchBarItem) ViewController() ViewController {
	rv := objc.Call[ViewController](t_, objc.Sel("viewController"))
	return rv
}

// The identifier for this item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/2544812-identifier?language=objc
func (t_ TouchBarItem) Identifier() TouchBarItemIdentifier {
	rv := objc.Call[TouchBarItemIdentifier](t_, objc.Sel("identifier"))
	return rv
}
