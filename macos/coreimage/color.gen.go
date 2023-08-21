// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Color] class.
var ColorClass = _ColorClass{objc.GetClass("CIColor")}

type _ColorClass struct {
	objc.Class
}

// An interface definition for the [Color] class.
type IColor interface {
	objc.IObject
	Red() float64
	Green() float64
	ColorSpace() coregraphics.ColorSpaceRef
	Blue() float64
	Alpha() float64
	Components() *float64
	NumberOfComponents() uint
	StringRepresentation() string
}

// The component values defining a color in a specific color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor?language=objc
type Color struct {
	objc.Object
}

func ColorFrom(ptr unsafe.Pointer) Color {
	return Color{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _ColorClass) ColorWithCGColor(c coregraphics.ColorRef) Color {
	rv := objc.Call[Color](cc, objc.Sel("colorWithCGColor:"), c)
	return rv
}

// Creates a color object from a Quartz color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor/1502106-colorwithcgcolor?language=objc
func Color_ColorWithCGColor(c coregraphics.ColorRef) Color {
	return ColorClass.ColorWithCGColor(c)
}

func (c_ Color) InitWithColor(color objc.IObject) Color {
	rv := objc.Call[Color](c_, objc.Sel("initWithColor:"), objc.Ptr(color))
	return rv
}

// Initializes a Core Image color object using a UIKit (or AppKit) color object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor/1528762-initwithcolor?language=objc
func NewColorWithColor(color objc.IObject) Color {
	instance := ColorClass.Alloc().InitWithColor(color)
	instance.Autorelease()
	return instance
}

func (c_ Color) InitWithCGColor(c coregraphics.ColorRef) Color {
	rv := objc.Call[Color](c_, objc.Sel("initWithCGColor:"), c)
	return rv
}

// Initializes a Core Image color object with a Core Graphics color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor/1437821-initwithcgcolor?language=objc
func NewColorWithCGColor(c coregraphics.ColorRef) Color {
	instance := ColorClass.Alloc().InitWithCGColor(c)
	instance.Autorelease()
	return instance
}

func (cc _ColorClass) ColorWithString(representation string) Color {
	rv := objc.Call[Color](cc, objc.Sel("colorWithString:"), representation)
	return rv
}

// Creates a color object using the RGBA color component values specified by a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor/1438059-colorwithstring?language=objc
func Color_ColorWithString(representation string) Color {
	return ColorClass.ColorWithString(representation)
}

func (cc _ColorClass) ColorWithRedGreenBlue(r float64, g float64, b float64) Color {
	rv := objc.Call[Color](cc, objc.Sel("colorWithRed:green:blue:"), r, g, b)
	return rv
}

// Creates a color object using the specified RGB color component values [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor/1437941-colorwithred?language=objc
func Color_ColorWithRedGreenBlue(r float64, g float64, b float64) Color {
	return ColorClass.ColorWithRedGreenBlue(r, g, b)
}

func (c_ Color) InitWithRedGreenBlue(r float64, g float64, b float64) Color {
	rv := objc.Call[Color](c_, objc.Sel("initWithRed:green:blue:"), r, g, b)
	return rv
}

// Initializes a Core Image color object with the specified red, green, and blue component values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor/1502102-initwithred?language=objc
func NewColorWithRedGreenBlue(r float64, g float64, b float64) Color {
	instance := ColorClass.Alloc().InitWithRedGreenBlue(r, g, b)
	instance.Autorelease()
	return instance
}

func (cc _ColorClass) Alloc() Color {
	rv := objc.Call[Color](cc, objc.Sel("alloc"))
	return rv
}

func Color_Alloc() Color {
	return ColorClass.Alloc()
}

func (cc _ColorClass) New() Color {
	rv := objc.Call[Color](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewColor() Color {
	return ColorClass.New()
}

func (c_ Color) Init() Color {
	rv := objc.Call[Color](c_, objc.Sel("init"))
	return rv
}

// Returns a color object whose RGB values are all 1.0 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor/1643571-whitecolor?language=objc
func (cc _ColorClass) WhiteColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("whiteColor"))
	return rv
}

// Returns a color object whose RGB values are all 1.0 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor/1643571-whitecolor?language=objc
func Color_WhiteColor() Color {
	return ColorClass.WhiteColor()
}

// Returns a color object whose RGB values are 1.0, 0.0, and 1.0 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor/1643574-magentacolor?language=objc
func (cc _ColorClass) MagentaColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("magentaColor"))
	return rv
}

// Returns a color object whose RGB values are 1.0, 0.0, and 1.0 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor/1643574-magentacolor?language=objc
func Color_MagentaColor() Color {
	return ColorClass.MagentaColor()
}

// The unpremultiplied red component of the color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor/1437969-red?language=objc
func (c_ Color) Red() float64 {
	rv := objc.Call[float64](c_, objc.Sel("red"))
	return rv
}

// Returns a color object whose RGB values are all 0.0 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor/1643578-blackcolor?language=objc
func (cc _ColorClass) BlackColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("blackColor"))
	return rv
}

// Returns a color object whose RGB values are all 0.0 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor/1643578-blackcolor?language=objc
func Color_BlackColor() Color {
	return ColorClass.BlackColor()
}

// Returns a color object whose RGB values are all 0.5 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor/1643573-graycolor?language=objc
func (cc _ColorClass) GrayColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("grayColor"))
	return rv
}

// Returns a color object whose RGB values are all 0.5 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor/1643573-graycolor?language=objc
func Color_GrayColor() Color {
	return ColorClass.GrayColor()
}

// Returns a color object whose RGB and alpha values are all 0.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor/1643577-clearcolor?language=objc
func (cc _ColorClass) ClearColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("clearColor"))
	return rv
}

// Returns a color object whose RGB and alpha values are all 0.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor/1643577-clearcolor?language=objc
func Color_ClearColor() Color {
	return ColorClass.ClearColor()
}

// Returns a color object whose RGB values are 0.0, 1.0, and 0.0 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor/1643580-greencolor?language=objc
func (cc _ColorClass) GreenColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("greenColor"))
	return rv
}

// Returns a color object whose RGB values are 0.0, 1.0, and 0.0 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor/1643580-greencolor?language=objc
func Color_GreenColor() Color {
	return ColorClass.GreenColor()
}

// Returns a color object whose RGB values are 1.0, 1.0, and 0.0 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor/1643582-yellowcolor?language=objc
func (cc _ColorClass) YellowColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("yellowColor"))
	return rv
}

// Returns a color object whose RGB values are 1.0, 1.0, and 0.0 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor/1643582-yellowcolor?language=objc
func Color_YellowColor() Color {
	return ColorClass.YellowColor()
}

// The unpremultiplied green component of the color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor/1437607-green?language=objc
func (c_ Color) Green() float64 {
	rv := objc.Call[float64](c_, objc.Sel("green"))
	return rv
}

// The Quartz 2D color space associated with the color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor/1437917-colorspace?language=objc
func (c_ Color) ColorSpace() coregraphics.ColorSpaceRef {
	rv := objc.Call[coregraphics.ColorSpaceRef](c_, objc.Sel("colorSpace"))
	return rv
}

// The unpremultiplied blue component of the color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor/1438033-blue?language=objc
func (c_ Color) Blue() float64 {
	rv := objc.Call[float64](c_, objc.Sel("blue"))
	return rv
}

// Returns a color object whose RGB values are 1.0, 0.0, and 0.0 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor/1643570-redcolor?language=objc
func (cc _ColorClass) RedColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("redColor"))
	return rv
}

// Returns a color object whose RGB values are 1.0, 0.0, and 0.0 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor/1643570-redcolor?language=objc
func Color_RedColor() Color {
	return ColorClass.RedColor()
}

// The alpha value of the color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor/1437981-alpha?language=objc
func (c_ Color) Alpha() float64 {
	rv := objc.Call[float64](c_, objc.Sel("alpha"))
	return rv
}

// Returns a color object whose RGB values are 0.0, 1.0, and 1.0 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor/1643581-cyancolor?language=objc
func (cc _ColorClass) CyanColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("cyanColor"))
	return rv
}

// Returns a color object whose RGB values are 0.0, 1.0, and 1.0 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor/1643581-cyancolor?language=objc
func Color_CyanColor() Color {
	return ColorClass.CyanColor()
}

// Returns a color object whose RGB values are 0.0, 0.0, and 1.0 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor/1643569-bluecolor?language=objc
func (cc _ColorClass) BlueColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("blueColor"))
	return rv
}

// Returns a color object whose RGB values are 0.0, 0.0, and 1.0 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor/1643569-bluecolor?language=objc
func Color_BlueColor() Color {
	return ColorClass.BlueColor()
}

// The color components of the color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor/1437862-components?language=objc
func (c_ Color) Components() *float64 {
	rv := objc.Call[*float64](c_, objc.Sel("components"))
	return rv
}

// Returns the number of color components in the color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor/1438151-numberofcomponents?language=objc
func (c_ Color) NumberOfComponents() uint {
	rv := objc.Call[uint](c_, objc.Sel("numberOfComponents"))
	return rv
}

// A formatted string that specifies the components of the color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolor/1437910-stringrepresentation?language=objc
func (c_ Color) StringRepresentation() string {
	rv := objc.Call[string](c_, objc.Sel("stringRepresentation"))
	return rv
}
