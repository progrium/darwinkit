// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ToolbarItem] class.
var ToolbarItemClass = _ToolbarItemClass{objc.GetClass("NSToolbarItem")}

type _ToolbarItemClass struct {
	objc.Class
}

// An interface definition for the [ToolbarItem] class.
type IToolbarItem interface {
	objc.IObject
	Validate()
	MenuFormRepresentation() MenuItem
	SetMenuFormRepresentation(value IMenuItem)
	Target() objc.Object
	SetTarget(value objc.IObject)
	IsBordered() bool
	SetBordered(value bool)
	VisibilityPriority() ToolbarItemVisibilityPriority
	SetVisibilityPriority(value ToolbarItemVisibilityPriority)
	Action() objc.Selector
	SetAction(value objc.Selector)
	Autovalidates() bool
	SetAutovalidates(value bool)
	ToolTip() string
	SetToolTip(value string)
	View() View
	SetView(value IView)
	AllowsDuplicatesInToolbar() bool
	Toolbar() Toolbar
	ItemIdentifier() ToolbarItemIdentifier
	IsNavigational() bool
	SetNavigational(value bool)
	Tag() int
	SetTag(value int)
	Title() string
	SetTitle(value string)
	IsEnabled() bool
	SetEnabled(value bool)
	Label() string
	SetLabel(value string)
	Image() Image
	SetImage(value IImage)
	PaletteLabel() string
	SetPaletteLabel(value string)
}

// A single item that appears in a window’s toolbar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem?language=objc
type ToolbarItem struct {
	objc.Object
}

func ToolbarItemFrom(ptr unsafe.Pointer) ToolbarItem {
	return ToolbarItem{
		Object: objc.ObjectFrom(ptr),
	}
}

func (t_ ToolbarItem) InitWithItemIdentifier(itemIdentifier ToolbarItemIdentifier) ToolbarItem {
	rv := objc.Call[ToolbarItem](t_, objc.Sel("initWithItemIdentifier:"), itemIdentifier)
	return rv
}

// Creates a toolbar item with the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/1534084-initwithitemidentifier?language=objc
func ToolbarItem_InitWithItemIdentifier(itemIdentifier ToolbarItemIdentifier) ToolbarItem {
	return ToolbarItemClass.Alloc().InitWithItemIdentifier(itemIdentifier)
}

func (tc _ToolbarItemClass) Alloc() ToolbarItem {
	rv := objc.Call[ToolbarItem](tc, objc.Sel("alloc"))
	return rv
}

func ToolbarItem_Alloc() ToolbarItem {
	return ToolbarItemClass.Alloc()
}

func (tc _ToolbarItemClass) New() ToolbarItem {
	rv := objc.Call[ToolbarItem](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewToolbarItem() ToolbarItem {
	return ToolbarItemClass.New()
}

func (t_ ToolbarItem) Init() ToolbarItem {
	rv := objc.Call[ToolbarItem](t_, objc.Sel("init"))
	return rv
}

// Validates the toolbar item’s menu and its ability to perfrom its action. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/1525295-validate?language=objc
func (t_ ToolbarItem) Validate() {
	objc.Call[objc.Void](t_, objc.Sel("validate"))
}

// The menu item to use when the toolbar item is in the overflow menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/1532562-menuformrepresentation?language=objc
func (t_ ToolbarItem) MenuFormRepresentation() MenuItem {
	rv := objc.Call[MenuItem](t_, objc.Sel("menuFormRepresentation"))
	return rv
}

// The menu item to use when the toolbar item is in the overflow menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/1532562-menuformrepresentation?language=objc
func (t_ ToolbarItem) SetMenuFormRepresentation(value IMenuItem) {
	objc.Call[objc.Void](t_, objc.Sel("setMenuFormRepresentation:"), objc.Ptr(value))
}

// The object that defines the action method the toolbar item calls when clicked. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/1525982-target?language=objc
func (t_ ToolbarItem) Target() objc.Object {
	rv := objc.Call[objc.Object](t_, objc.Sel("target"))
	return rv
}

// The object that defines the action method the toolbar item calls when clicked. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/1525982-target?language=objc
func (t_ ToolbarItem) SetTarget(value objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("setTarget:"), value)
}

// A Boolean value that indicates whether the toolbar item has a bordered style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/3237224-bordered?language=objc
func (t_ ToolbarItem) IsBordered() bool {
	rv := objc.Call[bool](t_, objc.Sel("isBordered"))
	return rv
}

// A Boolean value that indicates whether the toolbar item has a bordered style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/3237224-bordered?language=objc
func (t_ ToolbarItem) SetBordered(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setBordered:"), value)
}

// The display priority associated with the toolbar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/1527947-visibilitypriority?language=objc
func (t_ ToolbarItem) VisibilityPriority() ToolbarItemVisibilityPriority {
	rv := objc.Call[ToolbarItemVisibilityPriority](t_, objc.Sel("visibilityPriority"))
	return rv
}

// The display priority associated with the toolbar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/1527947-visibilitypriority?language=objc
func (t_ ToolbarItem) SetVisibilityPriority(value ToolbarItemVisibilityPriority) {
	objc.Call[objc.Void](t_, objc.Sel("setVisibilityPriority:"), value)
}

// The action method to call when someone clicks on the toolbar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/1525723-action?language=objc
func (t_ ToolbarItem) Action() objc.Selector {
	rv := objc.Call[objc.Selector](t_, objc.Sel("action"))
	return rv
}

// The action method to call when someone clicks on the toolbar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/1525723-action?language=objc
func (t_ ToolbarItem) SetAction(value objc.Selector) {
	objc.Call[objc.Void](t_, objc.Sel("setAction:"), value)
}

// A Boolean value that indicates whether the toolbar automatically validates the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/1524463-autovalidates?language=objc
func (t_ ToolbarItem) Autovalidates() bool {
	rv := objc.Call[bool](t_, objc.Sel("autovalidates"))
	return rv
}

// A Boolean value that indicates whether the toolbar automatically validates the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/1524463-autovalidates?language=objc
func (t_ ToolbarItem) SetAutovalidates(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAutovalidates:"), value)
}

// The tooltip to display when someone hovers over the item in the toolbar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/1524627-tooltip?language=objc
func (t_ ToolbarItem) ToolTip() string {
	rv := objc.Call[string](t_, objc.Sel("toolTip"))
	return rv
}

// The tooltip to display when someone hovers over the item in the toolbar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/1524627-tooltip?language=objc
func (t_ ToolbarItem) SetToolTip(value string) {
	objc.Call[objc.Void](t_, objc.Sel("setToolTip:"), value)
}

// The custom view you use to draw the toolbar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/1534039-view?language=objc
func (t_ ToolbarItem) View() View {
	rv := objc.Call[View](t_, objc.Sel("view"))
	return rv
}

// The custom view you use to draw the toolbar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/1534039-view?language=objc
func (t_ ToolbarItem) SetView(value IView) {
	objc.Call[objc.Void](t_, objc.Sel("setView:"), objc.Ptr(value))
}

// A Boolean value that indicates whether the toolbar item can appear more than once in a toolbar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/1530116-allowsduplicatesintoolbar?language=objc
func (t_ ToolbarItem) AllowsDuplicatesInToolbar() bool {
	rv := objc.Call[bool](t_, objc.Sel("allowsDuplicatesInToolbar"))
	return rv
}

// The toolbar that currently includes the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/1532976-toolbar?language=objc
func (t_ ToolbarItem) Toolbar() Toolbar {
	rv := objc.Call[Toolbar](t_, objc.Sel("toolbar"))
	return rv
}

// The value you use to identify the toolbar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/1524394-itemidentifier?language=objc
func (t_ ToolbarItem) ItemIdentifier() ToolbarItemIdentifier {
	rv := objc.Call[ToolbarItemIdentifier](t_, objc.Sel("itemIdentifier"))
	return rv
}

// A Boolean value that indicates whether the item behaves as a navigation item in the toolbar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/3622481-navigational?language=objc
func (t_ ToolbarItem) IsNavigational() bool {
	rv := objc.Call[bool](t_, objc.Sel("isNavigational"))
	return rv
}

// A Boolean value that indicates whether the item behaves as a navigation item in the toolbar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/3622481-navigational?language=objc
func (t_ ToolbarItem) SetNavigational(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setNavigational:"), value)
}

// An integer tag you can use to identify the toolbar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/1524618-tag?language=objc
func (t_ ToolbarItem) Tag() int {
	rv := objc.Call[int](t_, objc.Sel("tag"))
	return rv
}

// An integer tag you can use to identify the toolbar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/1524618-tag?language=objc
func (t_ ToolbarItem) SetTag(value int) {
	objc.Call[objc.Void](t_, objc.Sel("setTag:"), value)
}

// The title of the toolbar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/3237225-title?language=objc
func (t_ ToolbarItem) Title() string {
	rv := objc.Call[string](t_, objc.Sel("title"))
	return rv
}

// The title of the toolbar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/3237225-title?language=objc
func (t_ ToolbarItem) SetTitle(value string) {
	objc.Call[objc.Void](t_, objc.Sel("setTitle:"), value)
}

// A Boolean value that indicates whether the item is enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/1524277-enabled?language=objc
func (t_ ToolbarItem) IsEnabled() bool {
	rv := objc.Call[bool](t_, objc.Sel("isEnabled"))
	return rv
}

// A Boolean value that indicates whether the item is enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/1524277-enabled?language=objc
func (t_ ToolbarItem) SetEnabled(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setEnabled:"), value)
}

// The label that appears for this item in the toolbar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/1535127-label?language=objc
func (t_ ToolbarItem) Label() string {
	rv := objc.Call[string](t_, objc.Sel("label"))
	return rv
}

// The label that appears for this item in the toolbar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/1535127-label?language=objc
func (t_ ToolbarItem) SetLabel(value string) {
	objc.Call[objc.Void](t_, objc.Sel("setLabel:"), value)
}

// The image to display for the toolbar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/1527749-image?language=objc
func (t_ ToolbarItem) Image() Image {
	rv := objc.Call[Image](t_, objc.Sel("image"))
	return rv
}

// The image to display for the toolbar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/1527749-image?language=objc
func (t_ ToolbarItem) SetImage(value IImage) {
	objc.Call[objc.Void](t_, objc.Sel("setImage:"), objc.Ptr(value))
}

// The label that appears when the toolbar item is in the customization palette. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/1525421-palettelabel?language=objc
func (t_ ToolbarItem) PaletteLabel() string {
	rv := objc.Call[string](t_, objc.Sel("paletteLabel"))
	return rv
}

// The label that appears when the toolbar item is in the customization palette. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/1525421-palettelabel?language=objc
func (t_ ToolbarItem) SetPaletteLabel(value string) {
	objc.Call[objc.Void](t_, objc.Sel("setPaletteLabel:"), value)
}
