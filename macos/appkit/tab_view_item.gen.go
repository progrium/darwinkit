// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TabViewItem] class.
var TabViewItemClass = _TabViewItemClass{objc.GetClass("NSTabViewItem")}

type _TabViewItemClass struct {
	objc.Class
}

// An interface definition for the [TabViewItem] class.
type ITabViewItem interface {
	objc.IObject
	SizeOfLabel(computeMin bool) foundation.Size
	DrawLabelInRect(shouldTruncateLabel bool, labelRect foundation.Rect)
	Color() Color
	SetColor(value IColor)
	TabView() TabView
	ToolTip() string
	SetToolTip(value string)
	View() View
	SetView(value IView)
	TabState() TabState
	ViewController() ViewController
	SetViewController(value IViewController)
	InitialFirstResponder() View
	SetInitialFirstResponder(value IView)
	Label() string
	SetLabel(value string)
	Image() Image
	SetImage(value IImage)
	Identifier() objc.Object
	SetIdentifier(value objc.IObject)
}

// An item in a tab view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewitem?language=objc
type TabViewItem struct {
	objc.Object
}

func TabViewItemFrom(ptr unsafe.Pointer) TabViewItem {
	return TabViewItem{
		Object: objc.ObjectFrom(ptr),
	}
}

func (t_ TabViewItem) InitWithIdentifier(identifier objc.IObject) TabViewItem {
	rv := objc.Call[TabViewItem](t_, objc.Sel("initWithIdentifier:"), identifier)
	return rv
}

// Performs default initialization for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewitem/1477533-initwithidentifier?language=objc
func TabViewItem_InitWithIdentifier(identifier objc.IObject) TabViewItem {
	return TabViewItemClass.Alloc().InitWithIdentifier(identifier)
}

func (tc _TabViewItemClass) TabViewItemWithViewController(viewController IViewController) TabViewItem {
	rv := objc.Call[TabViewItem](tc, objc.Sel("tabViewItemWithViewController:"), objc.Ptr(viewController))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewitem/1477509-tabviewitemwithviewcontroller?language=objc
func TabViewItem_TabViewItemWithViewController(viewController IViewController) TabViewItem {
	return TabViewItemClass.TabViewItemWithViewController(viewController)
}

func (tc _TabViewItemClass) Alloc() TabViewItem {
	rv := objc.Call[TabViewItem](tc, objc.Sel("alloc"))
	return rv
}

func TabViewItem_Alloc() TabViewItem {
	return TabViewItemClass.Alloc()
}

func (tc _TabViewItemClass) New() TabViewItem {
	rv := objc.Call[TabViewItem](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTabViewItem() TabViewItem {
	return TabViewItemClass.New()
}

func (t_ TabViewItem) Init() TabViewItem {
	rv := objc.Call[TabViewItem](t_, objc.Sel("init"))
	return rv
}

// Calculates the size of the receiver’s label. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewitem/1477513-sizeoflabel?language=objc
func (t_ TabViewItem) SizeOfLabel(computeMin bool) foundation.Size {
	rv := objc.Call[foundation.Size](t_, objc.Sel("sizeOfLabel:"), computeMin)
	return rv
}

// Draws the receiver’s label in tabRect, which is the area between the curved end caps. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewitem/1477523-drawlabel?language=objc
func (t_ TabViewItem) DrawLabelInRect(shouldTruncateLabel bool, labelRect foundation.Rect) {
	objc.Call[objc.Void](t_, objc.Sel("drawLabel:inRect:"), shouldTruncateLabel, labelRect)
}

// Sets the background color for content in the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewitem/1477525-color?language=objc
func (t_ TabViewItem) Color() Color {
	rv := objc.Call[Color](t_, objc.Sel("color"))
	return rv
}

// Sets the background color for content in the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewitem/1477525-color?language=objc
func (t_ TabViewItem) SetColor(value IColor) {
	objc.Call[objc.Void](t_, objc.Sel("setColor:"), objc.Ptr(value))
}

// Returns the parent tab view for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewitem/1477535-tabview?language=objc
func (t_ TabViewItem) TabView() TabView {
	rv := objc.Call[TabView](t_, objc.Sel("tabView"))
	return rv
}

// Sets the tooltip displayed for the tab view item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewitem/1477515-tooltip?language=objc
func (t_ TabViewItem) ToolTip() string {
	rv := objc.Call[string](t_, objc.Sel("toolTip"))
	return rv
}

// Sets the tooltip displayed for the tab view item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewitem/1477515-tooltip?language=objc
func (t_ TabViewItem) SetToolTip(value string) {
	objc.Call[objc.Void](t_, objc.Sel("setToolTip:"), value)
}

// Sets the view associated with the receiver to view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewitem/1477537-view?language=objc
func (t_ TabViewItem) View() View {
	rv := objc.Call[View](t_, objc.Sel("view"))
	return rv
}

// Sets the view associated with the receiver to view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewitem/1477537-view?language=objc
func (t_ TabViewItem) SetView(value IView) {
	objc.Call[objc.Void](t_, objc.Sel("setView:"), objc.Ptr(value))
}

// Returns the current display state of the tab associated with the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewitem/1477539-tabstate?language=objc
func (t_ TabViewItem) TabState() TabState {
	rv := objc.Call[TabState](t_, objc.Sel("tabState"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewitem/1477521-viewcontroller?language=objc
func (t_ TabViewItem) ViewController() ViewController {
	rv := objc.Call[ViewController](t_, objc.Sel("viewController"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewitem/1477521-viewcontroller?language=objc
func (t_ TabViewItem) SetViewController(value IViewController) {
	objc.Call[objc.Void](t_, objc.Sel("setViewController:"), objc.Ptr(value))
}

// Sets the initial first responder for the view associated with the receiver (the view that is displayed when a user clicks on the tab) to view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewitem/1477541-initialfirstresponder?language=objc
func (t_ TabViewItem) InitialFirstResponder() View {
	rv := objc.Call[View](t_, objc.Sel("initialFirstResponder"))
	return rv
}

// Sets the initial first responder for the view associated with the receiver (the view that is displayed when a user clicks on the tab) to view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewitem/1477541-initialfirstresponder?language=objc
func (t_ TabViewItem) SetInitialFirstResponder(value IView) {
	objc.Call[objc.Void](t_, objc.Sel("setInitialFirstResponder:"), objc.Ptr(value))
}

// Sets the label text for the receiver to label. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewitem/1477542-label?language=objc
func (t_ TabViewItem) Label() string {
	rv := objc.Call[string](t_, objc.Sel("label"))
	return rv
}

// Sets the label text for the receiver to label. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewitem/1477542-label?language=objc
func (t_ TabViewItem) SetLabel(value string) {
	objc.Call[objc.Void](t_, objc.Sel("setLabel:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewitem/1477527-image?language=objc
func (t_ TabViewItem) Image() Image {
	rv := objc.Call[Image](t_, objc.Sel("image"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewitem/1477527-image?language=objc
func (t_ TabViewItem) SetImage(value IImage) {
	objc.Call[objc.Void](t_, objc.Sel("setImage:"), objc.Ptr(value))
}

// Sets the receiver’s optional identifier object to identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewitem/1477529-identifier?language=objc
func (t_ TabViewItem) Identifier() objc.Object {
	rv := objc.Call[objc.Object](t_, objc.Sel("identifier"))
	return rv
}

// Sets the receiver’s optional identifier object to identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewitem/1477529-identifier?language=objc
func (t_ TabViewItem) SetIdentifier(value objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("setIdentifier:"), value)
}
