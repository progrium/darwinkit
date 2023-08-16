// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PopoverTouchBarItem] class.
var PopoverTouchBarItemClass = _PopoverTouchBarItemClass{objc.GetClass("NSPopoverTouchBarItem")}

type _PopoverTouchBarItemClass struct {
	objc.Class
}

// An interface definition for the [PopoverTouchBarItem] class.
type IPopoverTouchBarItem interface {
	ITouchBarItem
	ShowPopover(sender objc.IObject)
	DismissPopover(sender objc.IObject)
	MakeStandardActivatePopoverGestureRecognizer() GestureRecognizer
	CollapsedRepresentationImage() Image
	SetCollapsedRepresentationImage(value IImage)
	PressAndHoldTouchBar() TouchBar
	SetPressAndHoldTouchBar(value ITouchBar)
	SetCustomizationLabel(value string)
	CollapsedRepresentationLabel() string
	SetCollapsedRepresentationLabel(value string)
	ShowsCloseButton() bool
	SetShowsCloseButton(value bool)
	PopoverTouchBar() TouchBar
	SetPopoverTouchBar(value ITouchBar)
	CollapsedRepresentation() View
	SetCollapsedRepresentation(value IView)
}

// A bar item that provides a two-state control that can expand into its second state, showing the contents of a bar that it owns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem?language=objc
type PopoverTouchBarItem struct {
	TouchBarItem
}

func PopoverTouchBarItemFrom(ptr unsafe.Pointer) PopoverTouchBarItem {
	return PopoverTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}

func (pc _PopoverTouchBarItemClass) Alloc() PopoverTouchBarItem {
	rv := objc.Call[PopoverTouchBarItem](pc, objc.Sel("alloc"))
	return rv
}

func PopoverTouchBarItem_Alloc() PopoverTouchBarItem {
	return PopoverTouchBarItemClass.Alloc()
}

func (pc _PopoverTouchBarItemClass) New() PopoverTouchBarItem {
	rv := objc.Call[PopoverTouchBarItem](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPopoverTouchBarItem() PopoverTouchBarItem {
	return PopoverTouchBarItemClass.New()
}

func (p_ PopoverTouchBarItem) Init() PopoverTouchBarItem {
	rv := objc.Call[PopoverTouchBarItem](p_, objc.Sel("init"))
	return rv
}

func (p_ PopoverTouchBarItem) InitWithIdentifier(identifier TouchBarItemIdentifier) PopoverTouchBarItem {
	rv := objc.Call[PopoverTouchBarItem](p_, objc.Sel("initWithIdentifier:"), identifier)
	return rv
}

// Creates a new item with the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/2544755-initwithidentifier?language=objc
func PopoverTouchBarItem_InitWithIdentifier(identifier TouchBarItemIdentifier) PopoverTouchBarItem {
	return PopoverTouchBarItemClass.Alloc().InitWithIdentifier(identifier)
}

// Replaces the main bar with this item's popover bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/2544797-showpopover?language=objc
func (p_ PopoverTouchBarItem) ShowPopover(sender objc.IObject) {
	objc.Call[objc.Void](p_, objc.Sel("showPopover:"), sender)
}

// Restores the previously visible main bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/2544697-dismisspopover?language=objc
func (p_ PopoverTouchBarItem) DismissPopover(sender objc.IObject) {
	objc.Call[objc.Void](p_, objc.Sel("dismissPopover:"), sender)
}

// Returns a gesture recognizer, configured to invoke the showPopover: method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/2544842-makestandardactivatepopovergestu?language=objc
func (p_ PopoverTouchBarItem) MakeStandardActivatePopoverGestureRecognizer() GestureRecognizer {
	rv := objc.Call[GestureRecognizer](p_, objc.Sel("makeStandardActivatePopoverGestureRecognizer"))
	return rv
}

// The image displayed by the button for the default collapsed representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/2544725-collapsedrepresentationimage?language=objc
func (p_ PopoverTouchBarItem) CollapsedRepresentationImage() Image {
	rv := objc.Call[Image](p_, objc.Sel("collapsedRepresentationImage"))
	return rv
}

// The image displayed by the button for the default collapsed representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/2544725-collapsedrepresentationimage?language=objc
func (p_ PopoverTouchBarItem) SetCollapsedRepresentationImage(value IImage) {
	objc.Call[objc.Void](p_, objc.Sel("setCollapsedRepresentationImage:"), objc.Ptr(value))
}

// The bar that is displayed when a user press-and-holds on the popover item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/2646955-pressandholdtouchbar?language=objc
func (p_ PopoverTouchBarItem) PressAndHoldTouchBar() TouchBar {
	rv := objc.Call[TouchBar](p_, objc.Sel("pressAndHoldTouchBar"))
	return rv
}

// The bar that is displayed when a user press-and-holds on the popover item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/2646955-pressandholdtouchbar?language=objc
func (p_ PopoverTouchBarItem) SetPressAndHoldTouchBar(value ITouchBar) {
	objc.Call[objc.Void](p_, objc.Sel("setPressAndHoldTouchBar:"), objc.Ptr(value))
}

// The user-visible string identifying this item during bar customization. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/2544684-customizationlabel?language=objc
func (p_ PopoverTouchBarItem) SetCustomizationLabel(value string) {
	objc.Call[objc.Void](p_, objc.Sel("setCustomizationLabel:"), value)
}

// The localized string displayed by the button for the default collapsed representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/2544772-collapsedrepresentationlabel?language=objc
func (p_ PopoverTouchBarItem) CollapsedRepresentationLabel() string {
	rv := objc.Call[string](p_, objc.Sel("collapsedRepresentationLabel"))
	return rv
}

// The localized string displayed by the button for the default collapsed representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/2544772-collapsedrepresentationlabel?language=objc
func (p_ PopoverTouchBarItem) SetCollapsedRepresentationLabel(value string) {
	objc.Call[objc.Void](p_, objc.Sel("setCollapsedRepresentationLabel:"), value)
}

// A Boolean value that determines whether a close button should be shown on the popover bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/2544711-showsclosebutton?language=objc
func (p_ PopoverTouchBarItem) ShowsCloseButton() bool {
	rv := objc.Call[bool](p_, objc.Sel("showsCloseButton"))
	return rv
}

// A Boolean value that determines whether a close button should be shown on the popover bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/2544711-showsclosebutton?language=objc
func (p_ PopoverTouchBarItem) SetShowsCloseButton(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setShowsCloseButton:"), value)
}

// The bar displayed when this item is "popped." [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/2544856-popovertouchbar?language=objc
func (p_ PopoverTouchBarItem) PopoverTouchBar() TouchBar {
	rv := objc.Call[TouchBar](p_, objc.Sel("popoverTouchBar"))
	return rv
}

// The bar displayed when this item is "popped." [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/2544856-popovertouchbar?language=objc
func (p_ PopoverTouchBarItem) SetPopoverTouchBar(value ITouchBar) {
	objc.Call[objc.Void](p_, objc.Sel("setPopoverTouchBar:"), objc.Ptr(value))
}

// The view displayed when this item is displayed in its parent bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/2544699-collapsedrepresentation?language=objc
func (p_ PopoverTouchBarItem) CollapsedRepresentation() View {
	rv := objc.Call[View](p_, objc.Sel("collapsedRepresentation"))
	return rv
}

// The view displayed when this item is displayed in its parent bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/2544699-collapsedrepresentation?language=objc
func (p_ PopoverTouchBarItem) SetCollapsedRepresentation(value IView) {
	objc.Call[objc.Void](p_, objc.Sel("setCollapsedRepresentation:"), objc.Ptr(value))
}
