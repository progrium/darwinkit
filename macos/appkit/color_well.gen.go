// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ColorWell] class.
var ColorWellClass = _ColorWellClass{objc.GetClass("NSColorWell")}

type _ColorWellClass struct {
	objc.Class
}

// An interface definition for the [ColorWell] class.
type IColorWell interface {
	IControl
	Activate(exclusive bool)
	DrawWellInside(insideRect foundation.Rect)
	TakeColorFrom(sender objc.IObject)
	Deactivate()
	Color() Color
	SetColor(value IColor)
	IsActive() bool
}

// A control that displays a color value and lets the user change that color value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorwell?language=objc
type ColorWell struct {
	Control
}

func ColorWellFrom(ptr unsafe.Pointer) ColorWell {
	return ColorWell{
		Control: ControlFrom(ptr),
	}
}

func (cc _ColorWellClass) Alloc() ColorWell {
	rv := objc.Call[ColorWell](cc, objc.Sel("alloc"))
	return rv
}

func ColorWell_Alloc() ColorWell {
	return ColorWellClass.Alloc()
}

func (cc _ColorWellClass) New() ColorWell {
	rv := objc.Call[ColorWell](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewColorWell() ColorWell {
	return ColorWellClass.New()
}

func (c_ ColorWell) Init() ColorWell {
	rv := objc.Call[ColorWell](c_, objc.Sel("init"))
	return rv
}

func (c_ ColorWell) InitWithFrame(frameRect foundation.Rect) ColorWell {
	rv := objc.Call[ColorWell](c_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes a control with the specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428900-initwithframe?language=objc
func ColorWell_InitWithFrame(frameRect foundation.Rect) ColorWell {
	return ColorWellClass.Alloc().InitWithFrame(frameRect)
}

// Activates the color well, displays the color panel, and synchronizes the two UI elements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorwell/1527152-activate?language=objc
func (c_ ColorWell) Activate(exclusive bool) {
	objc.Call[objc.Void](c_, objc.Sel("activate:"), exclusive)
}

// Draws the area inside the color well at the specified location without drawing borders. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorwell/1530268-drawwellinside?language=objc
func (c_ ColorWell) DrawWellInside(insideRect foundation.Rect) {
	objc.Call[objc.Void](c_, objc.Sel("drawWellInside:"), insideRect)
}

// Changes the currently selected color to the color of the specified object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorwell/1524407-takecolorfrom?language=objc
func (c_ ColorWell) TakeColorFrom(sender objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("takeColorFrom:"), sender)
}

// Deactivates the color well. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorwell/1524493-deactivate?language=objc
func (c_ ColorWell) Deactivate() {
	objc.Call[objc.Void](c_, objc.Sel("deactivate"))
}

// The currently selected color for the color well. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorwell/1534546-color?language=objc
func (c_ ColorWell) Color() Color {
	rv := objc.Call[Color](c_, objc.Sel("color"))
	return rv
}

// The currently selected color for the color well. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorwell/1534546-color?language=objc
func (c_ ColorWell) SetColor(value IColor) {
	objc.Call[objc.Void](c_, objc.Sel("setColor:"), objc.Ptr(value))
}

// A Boolean value that indicates whether the color well is currently active. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorwell/1528698-active?language=objc
func (c_ ColorWell) IsActive() bool {
	rv := objc.Call[bool](c_, objc.Sel("isActive"))
	return rv
}
