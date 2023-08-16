// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MenuItemCell] class.
var MenuItemCellClass = _MenuItemCellClass{objc.GetClass("NSMenuItemCell")}

type _MenuItemCellClass struct {
	objc.Class
}

// An interface definition for the [MenuItemCell] class.
type IMenuItemCell interface {
	IButtonCell
	KeyEquivalentRectForBounds(cellFrame foundation.Rect) foundation.Rect
	StateImageRectForBounds(cellFrame foundation.Rect) foundation.Rect
	DrawSeparatorItemWithFrameInView(cellFrame foundation.Rect, controlView IView)
	CalcSize()
	DrawBorderAndBackgroundWithFrameInView(cellFrame foundation.Rect, controlView IView)
	DrawStateImageWithFrameInView(cellFrame foundation.Rect, controlView IView)
	DrawKeyEquivalentWithFrameInView(cellFrame foundation.Rect, controlView IView)
	KeyEquivalentWidth() float64
	TitleWidth() float64
	StateImageWidth() float64
	MenuItem() MenuItem
	SetMenuItem(value IMenuItem)
	NeedsSizing() bool
	SetNeedsSizing(value bool)
	ImageWidth() float64
	NeedsDisplay() bool
	SetNeedsDisplay(value bool)
}

// An object that handles the measurement and display of a single menu item in its encompassing frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell?language=objc
type MenuItemCell struct {
	ButtonCell
}

func MenuItemCellFrom(ptr unsafe.Pointer) MenuItemCell {
	return MenuItemCell{
		ButtonCell: ButtonCellFrom(ptr),
	}
}

func (m_ MenuItemCell) InitTextCell(string_ string) MenuItemCell {
	rv := objc.Call[MenuItemCell](m_, objc.Sel("initTextCell:"), string_)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/1641970-inittextcell?language=objc
func MenuItemCell_InitTextCell(string_ string) MenuItemCell {
	return MenuItemCellClass.Alloc().InitTextCell(string_)
}

func (mc _MenuItemCellClass) Alloc() MenuItemCell {
	rv := objc.Call[MenuItemCell](mc, objc.Sel("alloc"))
	return rv
}

func MenuItemCell_Alloc() MenuItemCell {
	return MenuItemCellClass.Alloc()
}

func (mc _MenuItemCellClass) New() MenuItemCell {
	rv := objc.Call[MenuItemCell](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMenuItemCell() MenuItemCell {
	return MenuItemCellClass.New()
}

func (m_ MenuItemCell) Init() MenuItemCell {
	rv := objc.Call[MenuItemCell](m_, objc.Sel("init"))
	return rv
}

func (m_ MenuItemCell) InitImageCell(image IImage) MenuItemCell {
	rv := objc.Call[MenuItemCell](m_, objc.Sel("initImageCell:"), objc.Ptr(image))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/1639152-initimagecell?language=objc
func MenuItemCell_InitImageCell(image IImage) MenuItemCell {
	return MenuItemCellClass.Alloc().InitImageCell(image)
}

// Returns the rectangle into which the menu item’s key equivalent should be drawn. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/1498859-keyequivalentrectforbounds?language=objc
func (m_ MenuItemCell) KeyEquivalentRectForBounds(cellFrame foundation.Rect) foundation.Rect {
	rv := objc.Call[foundation.Rect](m_, objc.Sel("keyEquivalentRectForBounds:"), cellFrame)
	return rv
}

// Returns the rectangle into which the menu item’s state image should be drawn. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/1498855-stateimagerectforbounds?language=objc
func (m_ MenuItemCell) StateImageRectForBounds(cellFrame foundation.Rect) foundation.Rect {
	rv := objc.Call[foundation.Rect](m_, objc.Sel("stateImageRectForBounds:"), cellFrame)
	return rv
}

// Draws a menu item separator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/1498876-drawseparatoritemwithframe?language=objc
func (m_ MenuItemCell) DrawSeparatorItemWithFrameInView(cellFrame foundation.Rect, controlView IView) {
	objc.Call[objc.Void](m_, objc.Sel("drawSeparatorItemWithFrame:inView:"), cellFrame, objc.Ptr(controlView))
}

// Calculates the minimum required width and height of the receiver’s menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/1498862-calcsize?language=objc
func (m_ MenuItemCell) CalcSize() {
	objc.Call[objc.Void](m_, objc.Sel("calcSize"))
}

// Draws the borders and background associated with the receiver’s menu item (if any). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/1498863-drawborderandbackgroundwithframe?language=objc
func (m_ MenuItemCell) DrawBorderAndBackgroundWithFrameInView(cellFrame foundation.Rect, controlView IView) {
	objc.Call[objc.Void](m_, objc.Sel("drawBorderAndBackgroundWithFrame:inView:"), cellFrame, objc.Ptr(controlView))
}

// Draws the state image associated with the menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/1498854-drawstateimagewithframe?language=objc
func (m_ MenuItemCell) DrawStateImageWithFrameInView(cellFrame foundation.Rect, controlView IView) {
	objc.Call[objc.Void](m_, objc.Sel("drawStateImageWithFrame:inView:"), cellFrame, objc.Ptr(controlView))
}

// Draws the key equivalent associated with the menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/1498851-drawkeyequivalentwithframe?language=objc
func (m_ MenuItemCell) DrawKeyEquivalentWithFrameInView(cellFrame foundation.Rect, controlView IView) {
	objc.Call[objc.Void](m_, objc.Sel("drawKeyEquivalentWithFrame:inView:"), cellFrame, objc.Ptr(controlView))
}

// The width of the menu item’s key equivalent string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/1498874-keyequivalentwidth?language=objc
func (m_ MenuItemCell) KeyEquivalentWidth() float64 {
	rv := objc.Call[float64](m_, objc.Sel("keyEquivalentWidth"))
	return rv
}

// The width of the menu item’s text, measured in points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/1498865-titlewidth?language=objc
func (m_ MenuItemCell) TitleWidth() float64 {
	rv := objc.Call[float64](m_, objc.Sel("titleWidth"))
	return rv
}

// The width of the image used to indicate the state of the menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/1498879-stateimagewidth?language=objc
func (m_ MenuItemCell) StateImageWidth() float64 {
	rv := objc.Call[float64](m_, objc.Sel("stateImageWidth"))
	return rv
}

// The menu item object associated with the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/1498871-menuitem?language=objc
func (m_ MenuItemCell) MenuItem() MenuItem {
	rv := objc.Call[MenuItem](m_, objc.Sel("menuItem"))
	return rv
}

// The menu item object associated with the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/1498871-menuitem?language=objc
func (m_ MenuItemCell) SetMenuItem(value IMenuItem) {
	objc.Call[objc.Void](m_, objc.Sel("setMenuItem:"), objc.Ptr(value))
}

// A Boolean value indicating whether the size of the menu needs to be calculated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/1498869-needssizing?language=objc
func (m_ MenuItemCell) NeedsSizing() bool {
	rv := objc.Call[bool](m_, objc.Sel("needsSizing"))
	return rv
}

// A Boolean value indicating whether the size of the menu needs to be calculated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/1498869-needssizing?language=objc
func (m_ MenuItemCell) SetNeedsSizing(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setNeedsSizing:"), value)
}

// The width of the image associated with the menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/1498877-imagewidth?language=objc
func (m_ MenuItemCell) ImageWidth() float64 {
	rv := objc.Call[float64](m_, objc.Sel("imageWidth"))
	return rv
}

// A Boolean value indicating whether the menu item needs to be displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/1498861-needsdisplay?language=objc
func (m_ MenuItemCell) NeedsDisplay() bool {
	rv := objc.Call[bool](m_, objc.Sel("needsDisplay"))
	return rv
}

// A Boolean value indicating whether the menu item needs to be displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/1498861-needsdisplay?language=objc
func (m_ MenuItemCell) SetNeedsDisplay(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setNeedsDisplay:"), value)
}
