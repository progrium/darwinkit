// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ColorPickerTouchBarItem] class.
var ColorPickerTouchBarItemClass = _ColorPickerTouchBarItemClass{objc.GetClass("NSColorPickerTouchBarItem")}

type _ColorPickerTouchBarItemClass struct {
	objc.Class
}

// An interface definition for the [ColorPickerTouchBarItem] class.
type IColorPickerTouchBarItem interface {
	ITouchBarItem
	Color() Color
	SetColor(value IColor)
	Target() objc.Object
	SetTarget(value objc.IObject)
	SetCustomizationLabel(value string)
	Action() objc.Selector
	SetAction(value objc.Selector)
	AllowedColorSpaces() []ColorSpace
	SetAllowedColorSpaces(value []IColorSpace)
	ColorList() ColorList
	SetColorList(value IColorList)
	ShowsAlpha() bool
	SetShowsAlpha(value bool)
	IsEnabled() bool
	SetEnabled(value bool)
}

// A bar item that provides a system-defined color picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickertouchbaritem?language=objc
type ColorPickerTouchBarItem struct {
	TouchBarItem
}

func ColorPickerTouchBarItemFrom(ptr unsafe.Pointer) ColorPickerTouchBarItem {
	return ColorPickerTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}

func (cc _ColorPickerTouchBarItemClass) StrokeColorPickerWithIdentifier(identifier TouchBarItemIdentifier) ColorPickerTouchBarItem {
	rv := objc.Call[ColorPickerTouchBarItem](cc, objc.Sel("strokeColorPickerWithIdentifier:"), identifier)
	return rv
}

// Creates a bar item with the standard stroke color picker icon. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickertouchbaritem/2560993-strokecolorpickerwithidentifier?language=objc
func ColorPickerTouchBarItem_StrokeColorPickerWithIdentifier(identifier TouchBarItemIdentifier) ColorPickerTouchBarItem {
	return ColorPickerTouchBarItemClass.StrokeColorPickerWithIdentifier(identifier)
}

func (cc _ColorPickerTouchBarItemClass) ColorPickerWithIdentifier(identifier TouchBarItemIdentifier) ColorPickerTouchBarItem {
	rv := objc.Call[ColorPickerTouchBarItem](cc, objc.Sel("colorPickerWithIdentifier:"), identifier)
	return rv
}

// Creates a bar item with the standard color picker icon. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickertouchbaritem/2544807-colorpickerwithidentifier?language=objc
func ColorPickerTouchBarItem_ColorPickerWithIdentifier(identifier TouchBarItemIdentifier) ColorPickerTouchBarItem {
	return ColorPickerTouchBarItemClass.ColorPickerWithIdentifier(identifier)
}

func (cc _ColorPickerTouchBarItemClass) TextColorPickerWithIdentifier(identifier TouchBarItemIdentifier) ColorPickerTouchBarItem {
	rv := objc.Call[ColorPickerTouchBarItem](cc, objc.Sel("textColorPickerWithIdentifier:"), identifier)
	return rv
}

// Creates a bar item with the standard text color picker icon. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickertouchbaritem/2560995-textcolorpickerwithidentifier?language=objc
func ColorPickerTouchBarItem_TextColorPickerWithIdentifier(identifier TouchBarItemIdentifier) ColorPickerTouchBarItem {
	return ColorPickerTouchBarItemClass.TextColorPickerWithIdentifier(identifier)
}

func (cc _ColorPickerTouchBarItemClass) Alloc() ColorPickerTouchBarItem {
	rv := objc.Call[ColorPickerTouchBarItem](cc, objc.Sel("alloc"))
	return rv
}

func ColorPickerTouchBarItem_Alloc() ColorPickerTouchBarItem {
	return ColorPickerTouchBarItemClass.Alloc()
}

func (cc _ColorPickerTouchBarItemClass) New() ColorPickerTouchBarItem {
	rv := objc.Call[ColorPickerTouchBarItem](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewColorPickerTouchBarItem() ColorPickerTouchBarItem {
	return ColorPickerTouchBarItemClass.New()
}

func (c_ ColorPickerTouchBarItem) Init() ColorPickerTouchBarItem {
	rv := objc.Call[ColorPickerTouchBarItem](c_, objc.Sel("init"))
	return rv
}

func (c_ ColorPickerTouchBarItem) InitWithIdentifier(identifier TouchBarItemIdentifier) ColorPickerTouchBarItem {
	rv := objc.Call[ColorPickerTouchBarItem](c_, objc.Sel("initWithIdentifier:"), identifier)
	return rv
}

// Creates a new item with the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/2544755-initwithidentifier?language=objc
func ColorPickerTouchBarItem_InitWithIdentifier(identifier TouchBarItemIdentifier) ColorPickerTouchBarItem {
	return ColorPickerTouchBarItemClass.Alloc().InitWithIdentifier(identifier)
}

// The picker's currently selected color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickertouchbaritem/2544768-color?language=objc
func (c_ ColorPickerTouchBarItem) Color() Color {
	rv := objc.Call[Color](c_, objc.Sel("color"))
	return rv
}

// The picker's currently selected color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickertouchbaritem/2544768-color?language=objc
func (c_ ColorPickerTouchBarItem) SetColor(value IColor) {
	objc.Call[objc.Void](c_, objc.Sel("setColor:"), objc.Ptr(value))
}

// An object that is notified when a user interacts with the color picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickertouchbaritem/2544857-target?language=objc
func (c_ ColorPickerTouchBarItem) Target() objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("target"))
	return rv
}

// An object that is notified when a user interacts with the color picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickertouchbaritem/2544857-target?language=objc
func (c_ ColorPickerTouchBarItem) SetTarget(value objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setTarget:"), value)
}

// The user-visible string identifying this item during touch bar customization. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickertouchbaritem/2544803-customizationlabel?language=objc
func (c_ ColorPickerTouchBarItem) SetCustomizationLabel(value string) {
	objc.Call[objc.Void](c_, objc.Sel("setCustomizationLabel:"), value)
}

// The selector on the target object that is invoked when a user interacts with the color picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickertouchbaritem/2544673-action?language=objc
func (c_ ColorPickerTouchBarItem) Action() objc.Selector {
	rv := objc.Call[objc.Selector](c_, objc.Sel("action"))
	return rv
}

// The selector on the target object that is invoked when a user interacts with the color picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickertouchbaritem/2544673-action?language=objc
func (c_ ColorPickerTouchBarItem) SetAction(value objc.Selector) {
	objc.Call[objc.Void](c_, objc.Sel("setAction:"), value)
}

// Controls the color spaces that the color picker can produce. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickertouchbaritem/2881218-allowedcolorspaces?language=objc
func (c_ ColorPickerTouchBarItem) AllowedColorSpaces() []ColorSpace {
	rv := objc.Call[[]ColorSpace](c_, objc.Sel("allowedColorSpaces"))
	return rv
}

// Controls the color spaces that the color picker can produce. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickertouchbaritem/2881218-allowedcolorspaces?language=objc
func (c_ ColorPickerTouchBarItem) SetAllowedColorSpaces(value []IColorSpace) {
	objc.Call[objc.Void](c_, objc.Sel("setAllowedColorSpaces:"), value)
}

// The list of colors displayed in the color picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickertouchbaritem/2641681-colorlist?language=objc
func (c_ ColorPickerTouchBarItem) ColorList() ColorList {
	rv := objc.Call[ColorList](c_, objc.Sel("colorList"))
	return rv
}

// The list of colors displayed in the color picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickertouchbaritem/2641681-colorlist?language=objc
func (c_ ColorPickerTouchBarItem) SetColorList(value IColorList) {
	objc.Call[objc.Void](c_, objc.Sel("setColorList:"), objc.Ptr(value))
}

// A Boolean value that controls whether the color picker allows picking of colors with alpha values other than 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickertouchbaritem/2544828-showsalpha?language=objc
func (c_ ColorPickerTouchBarItem) ShowsAlpha() bool {
	rv := objc.Call[bool](c_, objc.Sel("showsAlpha"))
	return rv
}

// A Boolean value that controls whether the color picker allows picking of colors with alpha values other than 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickertouchbaritem/2544828-showsalpha?language=objc
func (c_ ColorPickerTouchBarItem) SetShowsAlpha(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setShowsAlpha:"), value)
}

// A Boolean value that determines whether the color picker is enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickertouchbaritem/2641680-enabled?language=objc
func (c_ ColorPickerTouchBarItem) IsEnabled() bool {
	rv := objc.Call[bool](c_, objc.Sel("isEnabled"))
	return rv
}

// A Boolean value that determines whether the color picker is enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickertouchbaritem/2641680-enabled?language=objc
func (c_ ColorPickerTouchBarItem) SetEnabled(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setEnabled:"), value)
}
