// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/coreimage"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Color] class.
var ColorClass = _ColorClass{objc.GetClass("NSColor")}

type _ColorClass struct {
	objc.Class
}

// An interface definition for the [Color] class.
type IColor interface {
	objc.IObject
	ColorWithSystemEffect(systemEffect ColorSystemEffect) Color
	Set()
	WriteToPasteboard(pasteBoard IPasteboard)
	ColorUsingType(type_ ColorType) Color
	ShadowWithLevel(val float64) Color
	BlendedColorWithFractionOfColor(fraction float64, color IColor) Color
	GetRedGreenBlueAlpha(red *float64, green *float64, blue *float64, alpha *float64)
	GetCyanMagentaYellowBlackAlpha(cyan *float64, magenta *float64, yellow *float64, black *float64, alpha *float64)
	GetComponents(components *float64)
	GetWhiteAlpha(white *float64, alpha *float64)
	SetStroke()
	DrawSwatchInRect(rect foundation.Rect)
	HighlightWithLevel(val float64) Color
	ColorWithAlphaComponent(alpha float64) Color
	SetFill()
	GetHueSaturationBrightnessAlpha(hue *float64, saturation *float64, brightness *float64, alpha *float64)
	ColorUsingColorSpace(space IColorSpace) Color
	GreenComponent() float64
	BlackComponent() float64
	LocalizedColorNameComponent() string
	YellowComponent() float64
	PatternImage() Image
	HueComponent() float64
	ColorNameComponent() ColorName
	RedComponent() float64
	ColorSpace() ColorSpace
	WhiteComponent() float64
	BrightnessComponent() float64
	SaturationComponent() float64
	MagentaComponent() float64
	Type() ColorType
	BlueComponent() float64
	LocalizedCatalogNameComponent() string
	NumberOfComponents() int
	AlphaComponent() float64
	CyanComponent() float64
	CatalogNameComponent() ColorListName
	CGColor() coregraphics.ColorRef
}

// An object that stores color data and sometimes opacity (alpha value). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor?language=objc
type Color struct {
	objc.Object
}

func ColorFrom(ptr unsafe.Pointer) Color {
	return Color{
		Object: objc.ObjectFrom(ptr),
	}
}

func (c_ Color) Init() Color {
	rv := objc.Call[Color](c_, objc.Sel("init"))
	return rv
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

// Creates a color object using the specified Core Graphics color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1524265-colorwithcgcolor?language=objc
func (cc _ColorClass) ColorWithCGColor(cgColor coregraphics.ColorRef) Color {
	rv := objc.Call[Color](cc, objc.Sel("colorWithCGColor:"), cgColor)
	return rv
}

// Creates a color object using the specified Core Graphics color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1524265-colorwithcgcolor?language=objc
func Color_ColorWithCGColor(cgColor coregraphics.ColorRef) Color {
	return ColorClass.ColorWithCGColor(cgColor)
}

// Returns a new color object that represents the current color modified to include the specified visual effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2998826-colorwithsystemeffect?language=objc
func (c_ Color) ColorWithSystemEffect(systemEffect ColorSystemEffect) Color {
	rv := objc.Call[Color](c_, objc.Sel("colorWithSystemEffect:"), systemEffect)
	return rv
}

// Sets the color of subsequent drawing to the color that the color object represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1527089-set?language=objc
func (c_ Color) Set() {
	objc.Call[objc.Void](c_, objc.Sel("set"))
}

// Writes the color object’s data to the specified pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1532199-writetopasteboard?language=objc
func (c_ Color) WriteToPasteboard(pasteBoard IPasteboard) {
	objc.Call[objc.Void](c_, objc.Sel("writeToPasteboard:"), objc.Ptr(pasteBoard))
}

// Creates a color object from the provided name, which corresponds to a color in the default asset catalog of the app's main bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2876295-colornamed?language=objc
func (cc _ColorClass) ColorNamed(name ColorName) Color {
	rv := objc.Call[Color](cc, objc.Sel("colorNamed:"), name)
	return rv
}

// Creates a color object from the provided name, which corresponds to a color in the default asset catalog of the app's main bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2876295-colornamed?language=objc
func Color_ColorNamed(name ColorName) Color {
	return ColorClass.ColorNamed(name)
}

// Returns a version of the color object that is compatible with the specified color type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2880320-colorusingtype?language=objc
func (c_ Color) ColorUsingType(type_ ColorType) Color {
	rv := objc.Call[Color](c_, objc.Sel("colorUsingType:"), type_)
	return rv
}

// Creates a new color object that represents a blend between the current color and the shadow color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1528523-shadowwithlevel?language=objc
func (c_ Color) ShadowWithLevel(val float64) Color {
	rv := objc.Call[Color](c_, objc.Sel("shadowWithLevel:"), val)
	return rv
}

// Creates a new color object whose component values are a weighted sum of the current color object and the specified color object's. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1524689-blendedcolorwithfraction?language=objc
func (c_ Color) BlendedColorWithFractionOfColor(fraction float64, color IColor) Color {
	rv := objc.Call[Color](c_, objc.Sel("blendedColorWithFraction:ofColor:"), fraction, objc.Ptr(color))
	return rv
}

// Creates a color object with the specified hue, saturation, brightness, and alpha channel values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1530601-colorwithhue?language=objc
func (cc _ColorClass) ColorWithHueSaturationBrightnessAlpha(hue float64, saturation float64, brightness float64, alpha float64) Color {
	rv := objc.Call[Color](cc, objc.Sel("colorWithHue:saturation:brightness:alpha:"), hue, saturation, brightness, alpha)
	return rv
}

// Creates a color object with the specified hue, saturation, brightness, and alpha channel values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1530601-colorwithhue?language=objc
func Color_ColorWithHueSaturationBrightnessAlpha(hue float64, saturation float64, brightness float64, alpha float64) Color {
	return ColorClass.ColorWithHueSaturationBrightnessAlpha(hue, saturation, brightness, alpha)
}

// Creates a color object using the given opacity and RGB components. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1526323-colorwithcalibratedred?language=objc
func (cc _ColorClass) ColorWithCalibratedRedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) Color {
	rv := objc.Call[Color](cc, objc.Sel("colorWithCalibratedRed:green:blue:alpha:"), red, green, blue, alpha)
	return rv
}

// Creates a color object using the given opacity and RGB components. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1526323-colorwithcalibratedred?language=objc
func Color_ColorWithCalibratedRedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) Color {
	return ColorClass.ColorWithCalibratedRedGreenBlueAlpha(red, green, blue, alpha)
}

// Returns the color object’s RGB component and opacity values in the respective arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1527848-getred?language=objc
func (c_ Color) GetRedGreenBlueAlpha(red *float64, green *float64, blue *float64, alpha *float64) {
	objc.Call[objc.Void](c_, objc.Sel("getRed:green:blue:alpha:"), red, green, blue, alpha)
}

// Returns the color object’s CMYK and opacity values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1531348-getcyan?language=objc
func (c_ Color) GetCyanMagentaYellowBlackAlpha(cyan *float64, magenta *float64, yellow *float64, black *float64, alpha *float64) {
	objc.Call[objc.Void](c_, objc.Sel("getCyan:magenta:yellow:black:alpha:"), cyan, magenta, yellow, black, alpha)
}

// Creates a dynamic catalog color with a provider that’s used to resolve the exact color value, calculated on first use. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/3294481-colorwithname?language=objc
func (cc _ColorClass) ColorWithNameDynamicProvider(colorName ColorName, dynamicProvider func(arg0 Appearance) Color) Color {
	rv := objc.Call[Color](cc, objc.Sel("colorWithName:dynamicProvider:"), colorName, dynamicProvider)
	return rv
}

// Creates a dynamic catalog color with a provider that’s used to resolve the exact color value, calculated on first use. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/3294481-colorwithname?language=objc
func Color_ColorWithNameDynamicProvider(colorName ColorName, dynamicProvider func(arg0 Appearance) Color) Color {
	return ColorClass.ColorWithNameDynamicProvider(colorName, dynamicProvider)
}

// Creates a color object using the given opacity value and HSB color space components. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1533653-colorwithdevicehue?language=objc
func (cc _ColorClass) ColorWithDeviceHueSaturationBrightnessAlpha(hue float64, saturation float64, brightness float64, alpha float64) Color {
	rv := objc.Call[Color](cc, objc.Sel("colorWithDeviceHue:saturation:brightness:alpha:"), hue, saturation, brightness, alpha)
	return rv
}

// Creates a color object using the given opacity value and HSB color space components. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1533653-colorwithdevicehue?language=objc
func Color_ColorWithDeviceHueSaturationBrightnessAlpha(hue float64, saturation float64, brightness float64, alpha float64) Color {
	return ColorClass.ColorWithDeviceHueSaturationBrightnessAlpha(hue, saturation, brightness, alpha)
}

// Returns the components of the color as an array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1524600-getcomponents?language=objc
func (c_ Color) GetComponents(components *float64) {
	objc.Call[objc.Void](c_, objc.Sel("getComponents:"), components)
}

// Returns the grayscale and alpha values of the color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1532613-getwhite?language=objc
func (c_ Color) GetWhiteAlpha(white *float64, alpha *float64) {
	objc.Call[objc.Void](c_, objc.Sel("getWhite:alpha:"), white, alpha)
}

// Sets the stroke color of subsequent drawing to the color object’s color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1531019-setstroke?language=objc
func (c_ Color) SetStroke() {
	objc.Call[objc.Void](c_, objc.Sel("setStroke"))
}

// Draws the current color in the specified rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1531770-drawswatchinrect?language=objc
func (c_ Color) DrawSwatchInRect(rect foundation.Rect) {
	objc.Call[objc.Void](c_, objc.Sel("drawSwatchInRect:"), rect)
}

// Creates a color object with the specified brightness and alpha channel values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1525501-colorwithwhite?language=objc
func (cc _ColorClass) ColorWithWhiteAlpha(white float64, alpha float64) Color {
	rv := objc.Call[Color](cc, objc.Sel("colorWithWhite:alpha:"), white, alpha)
	return rv
}

// Creates a color object with the specified brightness and alpha channel values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1525501-colorwithwhite?language=objc
func Color_ColorWithWhiteAlpha(white float64, alpha float64) Color {
	return ColorClass.ColorWithWhiteAlpha(white, alpha)
}

// Creates a color object from color data currently on the pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1535057-colorfrompasteboard?language=objc
func (cc _ColorClass) ColorFromPasteboard(pasteBoard IPasteboard) Color {
	rv := objc.Call[Color](cc, objc.Sel("colorFromPasteboard:"), objc.Ptr(pasteBoard))
	return rv
}

// Creates a color object from color data currently on the pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1535057-colorfrompasteboard?language=objc
func Color_ColorFromPasteboard(pasteBoard IPasteboard) Color {
	return ColorClass.ColorFromPasteboard(pasteBoard)
}

// Creates a color object from the specified components in the sRGB colorspace. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1532666-colorwithsrgbred?language=objc
func (cc _ColorClass) ColorWithSRGBRedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) Color {
	rv := objc.Call[Color](cc, objc.Sel("colorWithSRGBRed:green:blue:alpha:"), red, green, blue, alpha)
	return rv
}

// Creates a color object from the specified components in the sRGB colorspace. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1532666-colorwithsrgbred?language=objc
func Color_ColorWithSRGBRedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) Color {
	return ColorClass.ColorWithSRGBRedGreenBlueAlpha(red, green, blue, alpha)
}

// Creates a color object with the specified color space, hue, saturation, brightness, and alpha channel values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1644595-colorwithcolorspace?language=objc
func (cc _ColorClass) ColorWithColorSpaceHueSaturationBrightnessAlpha(space IColorSpace, hue float64, saturation float64, brightness float64, alpha float64) Color {
	rv := objc.Call[Color](cc, objc.Sel("colorWithColorSpace:hue:saturation:brightness:alpha:"), objc.Ptr(space), hue, saturation, brightness, alpha)
	return rv
}

// Creates a color object with the specified color space, hue, saturation, brightness, and alpha channel values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1644595-colorwithcolorspace?language=objc
func Color_ColorWithColorSpaceHueSaturationBrightnessAlpha(space IColorSpace, hue float64, saturation float64, brightness float64, alpha float64) Color {
	return ColorClass.ColorWithColorSpaceHueSaturationBrightnessAlpha(space, hue, saturation, brightness, alpha)
}

// Creates a color object using the specified asset catalog and color names. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1530559-colorwithcatalogname?language=objc
func (cc _ColorClass) ColorWithCatalogNameColorName(listName ColorListName, colorName ColorName) Color {
	rv := objc.Call[Color](cc, objc.Sel("colorWithCatalogName:colorName:"), listName, colorName)
	return rv
}

// Creates a color object using the specified asset catalog and color names. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1530559-colorwithcatalogname?language=objc
func Color_ColorWithCatalogNameColorName(listName ColorListName, colorName ColorName) Color {
	return ColorClass.ColorWithCatalogNameColorName(listName, colorName)
}

// Creates a new color object that represents a blend between the current color and the highlight color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1533061-highlightwithlevel?language=objc
func (c_ Color) HighlightWithLevel(val float64) Color {
	rv := objc.Call[Color](c_, objc.Sel("highlightWithLevel:"), val)
	return rv
}

// Creates a new color object that has the same color space and component values as the current color object, but the specified alpha component. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1526906-colorwithalphacomponent?language=objc
func (c_ Color) ColorWithAlphaComponent(alpha float64) Color {
	rv := objc.Call[Color](c_, objc.Sel("colorWithAlphaComponent:"), alpha)
	return rv
}

// Creates a color object with the specified red, green, blue, and alpha channel values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1535804-colorwithred?language=objc
func (cc _ColorClass) ColorWithRedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) Color {
	rv := objc.Call[Color](cc, objc.Sel("colorWithRed:green:blue:alpha:"), red, green, blue, alpha)
	return rv
}

// Creates a color object with the specified red, green, blue, and alpha channel values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1535804-colorwithred?language=objc
func Color_ColorWithRedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) Color {
	return ColorClass.ColorWithRedGreenBlueAlpha(red, green, blue, alpha)
}

// Creates a color object that uses the specified image pattern to paint the target area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1527422-colorwithpatternimage?language=objc
func (cc _ColorClass) ColorWithPatternImage(image IImage) Color {
	rv := objc.Call[Color](cc, objc.Sel("colorWithPatternImage:"), objc.Ptr(image))
	return rv
}

// Creates a color object that uses the specified image pattern to paint the target area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1527422-colorwithpatternimage?language=objc
func Color_ColorWithPatternImage(image IImage) Color {
	return ColorClass.ColorWithPatternImage(image)
}

// Sets the fill color of subsequent drawing to the color object’s color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1524755-setfill?language=objc
func (c_ Color) SetFill() {
	objc.Call[objc.Void](c_, objc.Sel("setFill"))
}

// Creates a color object from the specified Core Image color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1525446-colorwithcicolor?language=objc
func (cc _ColorClass) ColorWithCIColor(color coreimage.IColor) Color {
	rv := objc.Call[Color](cc, objc.Sel("colorWithCIColor:"), objc.Ptr(color))
	return rv
}

// Creates a color object from the specified Core Image color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1525446-colorwithcicolor?language=objc
func Color_ColorWithCIColor(color coreimage.IColor) Color {
	return ColorClass.ColorWithCIColor(color)
}

// Creates a color object using the given opacity and HSB color space components. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1535827-colorwithcalibratedhue?language=objc
func (cc _ColorClass) ColorWithCalibratedHueSaturationBrightnessAlpha(hue float64, saturation float64, brightness float64, alpha float64) Color {
	rv := objc.Call[Color](cc, objc.Sel("colorWithCalibratedHue:saturation:brightness:alpha:"), hue, saturation, brightness, alpha)
	return rv
}

// Creates a color object using the given opacity and HSB color space components. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1535827-colorwithcalibratedhue?language=objc
func Color_ColorWithCalibratedHueSaturationBrightnessAlpha(hue float64, saturation float64, brightness float64, alpha float64) Color {
	return ColorClass.ColorWithCalibratedHueSaturationBrightnessAlpha(hue, saturation, brightness, alpha)
}

// Returns the color object’s HSB component and opacity values in the respective arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1534060-gethue?language=objc
func (c_ Color) GetHueSaturationBrightnessAlpha(hue *float64, saturation *float64, brightness *float64, alpha *float64) {
	objc.Call[objc.Void](c_, objc.Sel("getHue:saturation:brightness:alpha:"), hue, saturation, brightness, alpha)
}

// Creates a color object using the given opacity and grayscale values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1535884-colorwithcalibratedwhite?language=objc
func (cc _ColorClass) ColorWithCalibratedWhiteAlpha(white float64, alpha float64) Color {
	rv := objc.Call[Color](cc, objc.Sel("colorWithCalibratedWhite:alpha:"), white, alpha)
	return rv
}

// Creates a color object using the given opacity and grayscale values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1535884-colorwithcalibratedwhite?language=objc
func Color_ColorWithCalibratedWhiteAlpha(white float64, alpha float64) Color {
	return ColorClass.ColorWithCalibratedWhiteAlpha(white, alpha)
}

// Returns a color object with the specified white and alpha values in the GenericGamma22 colorspace. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1525670-colorwithgenericgamma22white?language=objc
func (cc _ColorClass) ColorWithGenericGamma22WhiteAlpha(white float64, alpha float64) Color {
	rv := objc.Call[Color](cc, objc.Sel("colorWithGenericGamma22White:alpha:"), white, alpha)
	return rv
}

// Returns a color object with the specified white and alpha values in the GenericGamma22 colorspace. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1525670-colorwithgenericgamma22white?language=objc
func Color_ColorWithGenericGamma22WhiteAlpha(white float64, alpha float64) Color {
	return ColorClass.ColorWithGenericGamma22WhiteAlpha(white, alpha)
}

// Creates a color object from the specified components in the Display P3 color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1644633-colorwithdisplayp3red?language=objc
func (cc _ColorClass) ColorWithDisplayP3RedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) Color {
	rv := objc.Call[Color](cc, objc.Sel("colorWithDisplayP3Red:green:blue:alpha:"), red, green, blue, alpha)
	return rv
}

// Creates a color object from the specified components in the Display P3 color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1644633-colorwithdisplayp3red?language=objc
func Color_ColorWithDisplayP3RedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) Color {
	return ColorClass.ColorWithDisplayP3RedGreenBlueAlpha(red, green, blue, alpha)
}

// Creates a color object using the given opacity value and RGB components. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1524567-colorwithdevicered?language=objc
func (cc _ColorClass) ColorWithDeviceRedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) Color {
	rv := objc.Call[Color](cc, objc.Sel("colorWithDeviceRed:green:blue:alpha:"), red, green, blue, alpha)
	return rv
}

// Creates a color object using the given opacity value and RGB components. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1524567-colorwithdevicered?language=objc
func Color_ColorWithDeviceRedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) Color {
	return ColorClass.ColorWithDeviceRedGreenBlueAlpha(red, green, blue, alpha)
}

// Creates a color object using the given opacity value and CMYK components. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1525661-colorwithdevicecyan?language=objc
func (cc _ColorClass) ColorWithDeviceCyanMagentaYellowBlackAlpha(cyan float64, magenta float64, yellow float64, black float64, alpha float64) Color {
	rv := objc.Call[Color](cc, objc.Sel("colorWithDeviceCyan:magenta:yellow:black:alpha:"), cyan, magenta, yellow, black, alpha)
	return rv
}

// Creates a color object using the given opacity value and CMYK components. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1525661-colorwithdevicecyan?language=objc
func Color_ColorWithDeviceCyanMagentaYellowBlackAlpha(cyan float64, magenta float64, yellow float64, black float64, alpha float64) Color {
	return ColorClass.ColorWithDeviceCyanMagentaYellowBlackAlpha(cyan, magenta, yellow, black, alpha)
}

// Creates a new color object representing the color of the current color object in the specified color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1527379-colorusingcolorspace?language=objc
func (c_ Color) ColorUsingColorSpace(space IColorSpace) Color {
	rv := objc.Call[Color](c_, objc.Sel("colorUsingColorSpace:"), objc.Ptr(space))
	return rv
}

// Creates a color object using the given opacity and grayscale values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1525958-colorwithdevicewhite?language=objc
func (cc _ColorClass) ColorWithDeviceWhiteAlpha(white float64, alpha float64) Color {
	rv := objc.Call[Color](cc, objc.Sel("colorWithDeviceWhite:alpha:"), white, alpha)
	return rv
}

// Creates a color object using the given opacity and grayscale values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1525958-colorwithdevicewhite?language=objc
func Color_ColorWithDeviceWhiteAlpha(white float64, alpha float64) Color {
	return ColorClass.ColorWithDeviceWhiteAlpha(white, alpha)
}

// Returns a color object whose grayscale value is 2/3 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1535184-lightgraycolor?language=objc
func (cc _ColorClass) LightGrayColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("lightGrayColor"))
	return rv
}

// Returns a color object whose grayscale value is 2/3 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1535184-lightgraycolor?language=objc
func Color_LightGrayColor() Color {
	return ColorClass.LightGrayColor()
}

// The color to use for the background of large controls, such as scroll views or table views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1531948-controlbackgroundcolor?language=objc
func (cc _ColorClass) ControlBackgroundColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("controlBackgroundColor"))
	return rv
}

// The color to use for the background of large controls, such as scroll views or table views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1531948-controlbackgroundcolor?language=objc
func Color_ControlBackgroundColor() Color {
	return ColorClass.ControlBackgroundColor()
}

// The green component value of the color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1525935-greencomponent?language=objc
func (c_ Color) GreenComponent() float64 {
	rv := objc.Call[float64](c_, objc.Sel("greenComponent"))
	return rv
}

// Returns a color object whose grayscale and alpha values are both 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1526370-whitecolor?language=objc
func (cc _ColorClass) WhiteColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("whiteColor"))
	return rv
}

// Returns a color object whose grayscale and alpha values are both 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1526370-whitecolor?language=objc
func Color_WhiteColor() Color {
	return ColorClass.WhiteColor()
}

// Returns a color object whose RGB value is 1.0, 0.0, 1.0 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1529883-magentacolor?language=objc
func (cc _ColorClass) MagentaColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("magentaColor"))
	return rv
}

// Returns a color object whose RGB value is 1.0, 0.0, 1.0 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1529883-magentacolor?language=objc
func Color_MagentaColor() Color {
	return ColorClass.MagentaColor()
}

// The color to use for the face of a selected control—that is, a control that has been clicked or is being dragged. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1526796-selectedcontrolcolor?language=objc
func (cc _ColorClass) SelectedControlColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("selectedControlColor"))
	return rv
}

// The color to use for the face of a selected control—that is, a control that has been clicked or is being dragged. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1526796-selectedcontrolcolor?language=objc
func Color_SelectedControlColor() Color {
	return ColorClass.SelectedControlColor()
}

// The black component value of the color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1526883-blackcomponent?language=objc
func (c_ Color) BlackComponent() float64 {
	rv := objc.Call[float64](c_, objc.Sel("blackComponent"))
	return rv
}

// The localized version of the color name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1527286-localizedcolornamecomponent?language=objc
func (c_ Color) LocalizedColorNameComponent() string {
	rv := objc.Call[string](c_, objc.Sel("localizedColorNameComponent"))
	return rv
}

// The yellow component value of the color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1531965-yellowcomponent?language=objc
func (c_ Color) YellowComponent() float64 {
	rv := objc.Call[float64](c_, objc.Sel("yellowComponent"))
	return rv
}

// The current system control tint color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1526377-currentcontroltint?language=objc
func (cc _ColorClass) CurrentControlTint() ControlTint {
	rv := objc.Call[ControlTint](cc, objc.Sel("currentControlTint"))
	return rv
}

// The current system control tint color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1526377-currentcontroltint?language=objc
func Color_CurrentControlTint() ControlTint {
	return ColorClass.CurrentControlTint()
}

// The color to use for the window background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1528630-windowbackgroundcolor?language=objc
func (cc _ColorClass) WindowBackgroundColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("windowBackgroundColor"))
	return rv
}

// The color to use for the window background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1528630-windowbackgroundcolor?language=objc
func Color_WindowBackgroundColor() Color {
	return ColorClass.WindowBackgroundColor()
}

// The pattern image used to paint the target area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1531572-patternimage?language=objc
func (c_ Color) PatternImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("patternImage"))
	return rv
}

// The color to use as a virtual light source on the screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1527697-highlightcolor?language=objc
func (cc _ColorClass) HighlightColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("highlightColor"))
	return rv
}

// The color to use as a virtual light source on the screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1527697-highlightcolor?language=objc
func Color_HighlightColor() Color {
	return ColorClass.HighlightColor()
}

// The secondary color to use for text labels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1533254-secondarylabelcolor?language=objc
func (cc _ColorClass) SecondaryLabelColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("secondaryLabelColor"))
	return rv
}

// The secondary color to use for text labels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1533254-secondarylabelcolor?language=objc
func Color_SecondaryLabelColor() Color {
	return ColorClass.SecondaryLabelColor()
}

// The hue component value of the color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1531780-huecomponent?language=objc
func (c_ Color) HueComponent() float64 {
	rv := objc.Call[float64](c_, objc.Sel("hueComponent"))
	return rv
}

// The color to use for text in a selected control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1527413-alternateselectedcontroltextcolo?language=objc
func (cc _ColorClass) AlternateSelectedControlTextColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("alternateSelectedControlTextColor"))
	return rv
}

// The color to use for text in a selected control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1527413-alternateselectedcontroltextcolo?language=objc
func Color_AlternateSelectedControlTextColor() Color {
	return ColorClass.AlternateSelectedControlTextColor()
}

// The patterned color to use for the background of a scrubber control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2646931-scrubbertexturedbackgroundcolor?language=objc
func (cc _ColorClass) ScrubberTexturedBackgroundColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("scrubberTexturedBackgroundColor"))
	return rv
}

// The patterned color to use for the background of a scrubber control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2646931-scrubbertexturedbackgroundcolor?language=objc
func Color_ScrubberTexturedBackgroundColor() Color {
	return ColorClass.ScrubberTexturedBackgroundColor()
}

// Returns a color object whose grayscale value is 0.0 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1527194-blackcolor?language=objc
func (cc _ColorClass) BlackColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("blackColor"))
	return rv
}

// Returns a color object whose grayscale value is 0.0 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1527194-blackcolor?language=objc
func Color_BlackColor() Color {
	return ColorClass.BlackColor()
}

// The name of the color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1528278-colornamecomponent?language=objc
func (c_ Color) ColorNameComponent() ColorName {
	rv := objc.Call[ColorName](c_, objc.Sel("colorNameComponent"))
	return rv
}

// The highlight color to use for the bubble that shows inline search result values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2998827-findhighlightcolor?language=objc
func (cc _ColorClass) FindHighlightColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("findHighlightColor"))
	return rv
}

// The highlight color to use for the bubble that shows inline search result values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2998827-findhighlightcolor?language=objc
func Color_FindHighlightColor() Color {
	return ColorClass.FindHighlightColor()
}

// The color to use for selected text in an unemphasized context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2998834-unemphasizedselectedtextcolor?language=objc
func (cc _ColorClass) UnemphasizedSelectedTextColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("unemphasizedSelectedTextColor"))
	return rv
}

// The color to use for selected text in an unemphasized context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2998834-unemphasizedselectedtextcolor?language=objc
func Color_UnemphasizedSelectedTextColor() Color {
	return ColorClass.UnemphasizedSelectedTextColor()
}

// Returns a color object whose grayscale value is 0.5 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1535307-graycolor?language=objc
func (cc _ColorClass) GrayColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("grayColor"))
	return rv
}

// Returns a color object whose grayscale value is 0.5 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1535307-graycolor?language=objc
func Color_GrayColor() Color {
	return ColorClass.GrayColor()
}

// Returns a color object for pink that automatically adapts to vibrancy and accessibility settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2879261-systempinkcolor?language=objc
func (cc _ColorClass) SystemPinkColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("systemPinkColor"))
	return rv
}

// Returns a color object for pink that automatically adapts to vibrancy and accessibility settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2879261-systempinkcolor?language=objc
func Color_SystemPinkColor() Color {
	return ColorClass.SystemPinkColor()
}

// Returns a color object for purple that automatically adapts to vibrancy and accessibility settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2879259-systempurplecolor?language=objc
func (cc _ColorClass) SystemPurpleColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("systemPurpleColor"))
	return rv
}

// Returns a color object for purple that automatically adapts to vibrancy and accessibility settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2879259-systempurplecolor?language=objc
func Color_SystemPurpleColor() Color {
	return ColorClass.SystemPurpleColor()
}

// Returns a color object for green that automatically adapts to vibrancy and accessibility settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2879264-systemgreencolor?language=objc
func (cc _ColorClass) SystemGreenColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("systemGreenColor"))
	return rv
}

// Returns a color object for green that automatically adapts to vibrancy and accessibility settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2879264-systemgreencolor?language=objc
func Color_SystemGreenColor() Color {
	return ColorClass.SystemGreenColor()
}

// The color to use for placeholder text in controls or text views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2998829-placeholdertextcolor?language=objc
func (cc _ColorClass) PlaceholderTextColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("placeholderTextColor"))
	return rv
}

// The color to use for placeholder text in controls or text views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2998829-placeholdertextcolor?language=objc
func Color_PlaceholderTextColor() Color {
	return ColorClass.PlaceholderTextColor()
}

// The user's current accent color preference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/3000782-controlaccentcolor?language=objc
func (cc _ColorClass) ControlAccentColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("controlAccentColor"))
	return rv
}

// The user's current accent color preference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/3000782-controlaccentcolor?language=objc
func Color_ControlAccentColor() Color {
	return ColorClass.ControlAccentColor()
}

// Returns a color object for yellow that automatically adapts to vibrancy and accessibility settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2879267-systemyellowcolor?language=objc
func (cc _ColorClass) SystemYellowColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("systemYellowColor"))
	return rv
}

// Returns a color object for yellow that automatically adapts to vibrancy and accessibility settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2879267-systemyellowcolor?language=objc
func Color_SystemYellowColor() Color {
	return ColorClass.SystemYellowColor()
}

// Returns a color object whose grayscale and alpha values are both 0.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1527217-clearcolor?language=objc
func (cc _ColorClass) ClearColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("clearColor"))
	return rv
}

// Returns a color object whose grayscale and alpha values are both 0.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1527217-clearcolor?language=objc
func Color_ClearColor() Color {
	return ColorClass.ClearColor()
}

// Returns a color object whose RGB value is 0.0, 1.0, 0.0 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1527433-greencolor?language=objc
func (cc _ColorClass) GreenColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("greenColor"))
	return rv
}

// Returns a color object whose RGB value is 0.0, 1.0, 0.0 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1527433-greencolor?language=objc
func Color_GreenColor() Color {
	return ColorClass.GreenColor()
}

// The color to use for separators between different sections of content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2998831-separatorcolor?language=objc
func (cc _ColorClass) SeparatorColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("separatorColor"))
	return rv
}

// The color to use for separators between different sections of content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2998831-separatorcolor?language=objc
func Color_SeparatorColor() Color {
	return ColorClass.SeparatorColor()
}

// Returns a color object whose RGB value is 1.0, 1.0, 0.0 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1524796-yellowcolor?language=objc
func (cc _ColorClass) YellowColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("yellowColor"))
	return rv
}

// Returns a color object whose RGB value is 1.0, 1.0, 0.0 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1524796-yellowcolor?language=objc
func Color_YellowColor() Color {
	return ColorClass.YellowColor()
}

// The tertiary color to use for text labels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1532376-tertiarylabelcolor?language=objc
func (cc _ColorClass) TertiaryLabelColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("tertiaryLabelColor"))
	return rv
}

// The tertiary color to use for text labels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1532376-tertiarylabelcolor?language=objc
func Color_TertiaryLabelColor() Color {
	return ColorClass.TertiaryLabelColor()
}

// The color to use for text on disabled controls. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1530573-disabledcontroltextcolor?language=objc
func (cc _ColorClass) DisabledControlTextColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("disabledControlTextColor"))
	return rv
}

// The color to use for text on disabled controls. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1530573-disabledcontroltextcolor?language=objc
func Color_DisabledControlTextColor() Color {
	return ColorClass.DisabledControlTextColor()
}

// Returns a color object whose RGB value is 0.5, 0.0, 0.5 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1526713-purplecolor?language=objc
func (cc _ColorClass) PurpleColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("purpleColor"))
	return rv
}

// Returns a color object whose RGB value is 0.5, 0.0, 0.5 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1526713-purplecolor?language=objc
func Color_PurpleColor() Color {
	return ColorClass.PurpleColor()
}

// The red component value of the color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1530483-redcomponent?language=objc
func (c_ Color) RedComponent() float64 {
	rv := objc.Call[float64](c_, objc.Sel("redComponent"))
	return rv
}

// Returns a color object whose RGB value is 0.6, 0.4, 0.2 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1534013-browncolor?language=objc
func (cc _ColorClass) BrownColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("brownColor"))
	return rv
}

// Returns a color object whose RGB value is 0.6, 0.4, 0.2 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1534013-browncolor?language=objc
func Color_BrownColor() Color {
	return ColorClass.BrownColor()
}

// The color to use for text in header cells in table views and outline views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1531237-headertextcolor?language=objc
func (cc _ColorClass) HeaderTextColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("headerTextColor"))
	return rv
}

// The color to use for text in header cells in table views and outline views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1531237-headertextcolor?language=objc
func Color_HeaderTextColor() Color {
	return ColorClass.HeaderTextColor()
}

// The color to use for selected and unemphasized content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2998832-unemphasizedselectedcontentbackg?language=objc
func (cc _ColorClass) UnemphasizedSelectedContentBackgroundColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("unemphasizedSelectedContentBackgroundColor"))
	return rv
}

// The color to use for selected and unemphasized content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2998832-unemphasizedselectedcontentbackg?language=objc
func Color_UnemphasizedSelectedContentBackgroundColor() Color {
	return ColorClass.UnemphasizedSelectedContentBackgroundColor()
}

// Returns a color object whose grayscale value is 1/3 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1533739-darkgraycolor?language=objc
func (cc _ColorClass) DarkGrayColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("darkGrayColor"))
	return rv
}

// Returns a color object whose grayscale value is 1/3 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1533739-darkgraycolor?language=objc
func Color_DarkGrayColor() Color {
	return ColorClass.DarkGrayColor()
}

// Returns a color object for gray that automatically adapts to vibrancy and accessibility settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2879265-systemgraycolor?language=objc
func (cc _ColorClass) SystemGrayColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("systemGrayColor"))
	return rv
}

// Returns a color object for gray that automatically adapts to vibrancy and accessibility settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2879265-systemgraycolor?language=objc
func Color_SystemGrayColor() Color {
	return ColorClass.SystemGrayColor()
}

// The color space associated with the color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1526733-colorspace?language=objc
func (c_ Color) ColorSpace() ColorSpace {
	rv := objc.Call[ColorSpace](c_, objc.Sel("colorSpace"))
	return rv
}

// Returns a color object for brown that automatically adapts to vibrancy and accessibility settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2879266-systembrowncolor?language=objc
func (cc _ColorClass) SystemBrownColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("systemBrownColor"))
	return rv
}

// Returns a color object for brown that automatically adapts to vibrancy and accessibility settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2879266-systembrowncolor?language=objc
func Color_SystemBrownColor() Color {
	return ColorClass.SystemBrownColor()
}

// The white component value of the color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1534051-whitecomponent?language=objc
func (c_ Color) WhiteComponent() float64 {
	rv := objc.Call[float64](c_, objc.Sel("whiteComponent"))
	return rv
}

// The color to use for the text background in an unemphasized context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2998833-unemphasizedselectedtextbackgrou?language=objc
func (cc _ColorClass) UnemphasizedSelectedTextBackgroundColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("unemphasizedSelectedTextBackgroundColor"))
	return rv
}

// The color to use for the text background in an unemphasized context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2998833-unemphasizedselectedtextbackgrou?language=objc
func Color_UnemphasizedSelectedTextBackgroundColor() Color {
	return ColorClass.UnemphasizedSelectedTextBackgroundColor()
}

// Returns a color object for red that automatically adapts to vibrancy and accessibility settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2879262-systemredcolor?language=objc
func (cc _ColorClass) SystemRedColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("systemRedColor"))
	return rv
}

// Returns a color object for red that automatically adapts to vibrancy and accessibility settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2879262-systemredcolor?language=objc
func Color_SystemRedColor() Color {
	return ColorClass.SystemRedColor()
}

// Returns a color object whose RGB value is 1.0, 0.0, 0.0 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1525063-redcolor?language=objc
func (cc _ColorClass) RedColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("redColor"))
	return rv
}

// Returns a color object whose RGB value is 1.0, 0.0, 0.0 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1525063-redcolor?language=objc
func Color_RedColor() Color {
	return ColorClass.RedColor()
}

// The color to use for text in a window's frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1524257-windowframetextcolor?language=objc
func (cc _ColorClass) WindowFrameTextColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("windowFrameTextColor"))
	return rv
}

// The color to use for text in a window's frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1524257-windowframetextcolor?language=objc
func Color_WindowFrameTextColor() Color {
	return ColorClass.WindowFrameTextColor()
}

// The brightness component value of the color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1529355-brightnesscomponent?language=objc
func (c_ Color) BrightnessComponent() float64 {
	rv := objc.Call[float64](c_, objc.Sel("brightnessComponent"))
	return rv
}

// The color to use for the background of selected text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1528136-selectedtextbackgroundcolor?language=objc
func (cc _ColorClass) SelectedTextBackgroundColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("selectedTextBackgroundColor"))
	return rv
}

// The color to use for the background of selected text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1528136-selectedtextbackgroundcolor?language=objc
func Color_SelectedTextBackgroundColor() Color {
	return ColorClass.SelectedTextBackgroundColor()
}

// Returns a color object for teal that automatically adapts to vibrancy and accessibility settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/3174921-systemtealcolor?language=objc
func (cc _ColorClass) SystemTealColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("systemTealColor"))
	return rv
}

// Returns a color object for teal that automatically adapts to vibrancy and accessibility settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/3174921-systemtealcolor?language=objc
func Color_SystemTealColor() Color {
	return ColorClass.SystemTealColor()
}

// Returns a color object whose RGB value is 0.0, 1.0, 1.0 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1529217-cyancolor?language=objc
func (cc _ColorClass) CyanColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("cyanColor"))
	return rv
}

// Returns a color object whose RGB value is 0.0, 1.0, 1.0 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1529217-cyancolor?language=objc
func Color_CyanColor() Color {
	return ColorClass.CyanColor()
}

// The quaternary color to use for text labels and separators. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1534635-quaternarylabelcolor?language=objc
func (cc _ColorClass) QuaternaryLabelColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("quaternaryLabelColor"))
	return rv
}

// The quaternary color to use for text labels and separators. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1534635-quaternarylabelcolor?language=objc
func Color_QuaternaryLabelColor() Color {
	return ColorClass.QuaternaryLabelColor()
}

// The saturation component value of the color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1526326-saturationcomponent?language=objc
func (c_ Color) SaturationComponent() float64 {
	rv := objc.Call[float64](c_, objc.Sel("saturationComponent"))
	return rv
}

// The color to use for selected text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1528605-selectedtextcolor?language=objc
func (cc _ColorClass) SelectedTextColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("selectedTextColor"))
	return rv
}

// The color to use for selected text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1528605-selectedtextcolor?language=objc
func Color_SelectedTextColor() Color {
	return ColorClass.SelectedTextColor()
}

// The magenta component value of the color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1535560-magentacomponent?language=objc
func (c_ Color) MagentaComponent() float64 {
	rv := objc.Call[float64](c_, objc.Sel("magentaComponent"))
	return rv
}

// Returns a color object whose RGB value is 0.0, 0.0, 1.0 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1535421-bluecolor?language=objc
func (cc _ColorClass) BlueColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("blueColor"))
	return rv
}

// Returns a color object whose RGB value is 0.0, 0.0, 1.0 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1535421-bluecolor?language=objc
func Color_BlueColor() Color {
	return ColorClass.BlueColor()
}

// The color to use for text in a selected control—that is, a control being clicked or dragged. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1535591-selectedcontroltextcolor?language=objc
func (cc _ColorClass) SelectedControlTextColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("selectedControlTextColor"))
	return rv
}

// The color to use for text in a selected control—that is, a control being clicked or dragged. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1535591-selectedcontroltextcolor?language=objc
func Color_SelectedControlTextColor() Color {
	return ColorClass.SelectedControlTextColor()
}

// The color to use for the background area behind text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1535446-textbackgroundcolor?language=objc
func (cc _ColorClass) TextBackgroundColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("textBackgroundColor"))
	return rv
}

// The color to use for the background area behind text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1535446-textbackgroundcolor?language=objc
func Color_TextBackgroundColor() Color {
	return ColorClass.TextBackgroundColor()
}

// The color to use for text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1527025-textcolor?language=objc
func (cc _ColorClass) TextColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("textColor"))
	return rv
}

// The color to use for text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1527025-textcolor?language=objc
func Color_TextColor() Color {
	return ColorClass.TextColor()
}

// The color to use for virtual shadows cast by raised objects on the screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1525121-shadowcolor?language=objc
func (cc _ColorClass) ShadowColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("shadowColor"))
	return rv
}

// The color to use for virtual shadows cast by raised objects on the screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1525121-shadowcolor?language=objc
func Color_ShadowColor() Color {
	return ColorClass.ShadowColor()
}

// The color to use for text on enabled controls. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1535230-controltextcolor?language=objc
func (cc _ColorClass) ControlTextColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("controlTextColor"))
	return rv
}

// The color to use for text on enabled controls. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1535230-controltextcolor?language=objc
func Color_ControlTextColor() Color {
	return ColorClass.ControlTextColor()
}

// Returns a color object for mint that automatically adapts to vibrancy and accessibility settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/3816006-systemmintcolor?language=objc
func (cc _ColorClass) SystemMintColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("systemMintColor"))
	return rv
}

// Returns a color object for mint that automatically adapts to vibrancy and accessibility settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/3816006-systemmintcolor?language=objc
func Color_SystemMintColor() Color {
	return ColorClass.SystemMintColor()
}

// The color to use for links. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2998828-linkcolor?language=objc
func (cc _ColorClass) LinkColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("linkColor"))
	return rv
}

// The color to use for links. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2998828-linkcolor?language=objc
func Color_LinkColor() Color {
	return ColorClass.LinkColor()
}

// The type of the color object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2880315-type?language=objc
func (c_ Color) Type() ColorType {
	rv := objc.Call[ColorType](c_, objc.Sel("type"))
	return rv
}

// The blue component value of the color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1534229-bluecomponent?language=objc
func (c_ Color) BlueComponent() float64 {
	rv := objc.Call[float64](c_, objc.Sel("blueComponent"))
	return rv
}

// The localized version of the catalog name containing the color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1535351-localizedcatalognamecomponent?language=objc
func (c_ Color) LocalizedCatalogNameComponent() string {
	rv := objc.Call[string](c_, objc.Sel("localizedCatalogNameComponent"))
	return rv
}

// The number of components in the color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1531308-numberofcomponents?language=objc
func (c_ Color) NumberOfComponents() int {
	rv := objc.Call[int](c_, objc.Sel("numberOfComponents"))
	return rv
}

// The alpha (opacity) component value of the color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1532504-alphacomponent?language=objc
func (c_ Color) AlphaComponent() float64 {
	rv := objc.Call[float64](c_, objc.Sel("alphaComponent"))
	return rv
}

// The color to use for the optional gridlines, such as those in a table view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1527240-gridcolor?language=objc
func (cc _ColorClass) GridColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("gridColor"))
	return rv
}

// The color to use for the optional gridlines, such as those in a table view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1527240-gridcolor?language=objc
func Color_GridColor() Color {
	return ColorClass.GridColor()
}

// The color to use in the area beneath your window's views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1534707-underpagebackgroundcolor?language=objc
func (cc _ColorClass) UnderPageBackgroundColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("underPageBackgroundColor"))
	return rv
}

// The color to use in the area beneath your window's views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1534707-underpagebackgroundcolor?language=objc
func Color_UnderPageBackgroundColor() Color {
	return ColorClass.UnderPageBackgroundColor()
}

// Returns a color object for indigo that automatically adapts to vibrancy and accessibility settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/3174920-systemindigocolor?language=objc
func (cc _ColorClass) SystemIndigoColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("systemIndigoColor"))
	return rv
}

// Returns a color object for indigo that automatically adapts to vibrancy and accessibility settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/3174920-systemindigocolor?language=objc
func Color_SystemIndigoColor() Color {
	return ColorClass.SystemIndigoColor()
}

// Returns a color object for cyan that automatically adapts to vibrancy and accessibility settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/3816005-systemcyancolor?language=objc
func (cc _ColorClass) SystemCyanColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("systemCyanColor"))
	return rv
}

// Returns a color object for cyan that automatically adapts to vibrancy and accessibility settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/3816005-systemcyancolor?language=objc
func Color_SystemCyanColor() Color {
	return ColorClass.SystemCyanColor()
}

// The color to use for the background of selected and emphasized content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2998830-selectedcontentbackgroundcolor?language=objc
func (cc _ColorClass) SelectedContentBackgroundColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("selectedContentBackgroundColor"))
	return rv
}

// The color to use for the background of selected and emphasized content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2998830-selectedcontentbackgroundcolor?language=objc
func Color_SelectedContentBackgroundColor() Color {
	return ColorClass.SelectedContentBackgroundColor()
}

// The color to use for the text in menu items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1526658-selectedmenuitemtextcolor?language=objc
func (cc _ColorClass) SelectedMenuItemTextColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("selectedMenuItemTextColor"))
	return rv
}

// The color to use for the text in menu items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1526658-selectedmenuitemtextcolor?language=objc
func Color_SelectedMenuItemTextColor() Color {
	return ColorClass.SelectedMenuItemTextColor()
}

// Returns a color object whose RGB value is 1.0, 0.5, 0.0 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1526691-orangecolor?language=objc
func (cc _ColorClass) OrangeColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("orangeColor"))
	return rv
}

// Returns a color object whose RGB value is 1.0, 0.5, 0.0 and whose alpha value is 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1526691-orangecolor?language=objc
func Color_OrangeColor() Color {
	return ColorClass.OrangeColor()
}

// The cyan component value of the color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1528234-cyancomponent?language=objc
func (c_ Color) CyanComponent() float64 {
	rv := objc.Call[float64](c_, objc.Sel("cyanComponent"))
	return rv
}

// The colors to use for alternating content, typically found in table views and collection views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2998825-alternatingcontentbackgroundcolo?language=objc
func (cc _ColorClass) AlternatingContentBackgroundColors() []Color {
	rv := objc.Call[[]Color](cc, objc.Sel("alternatingContentBackgroundColors"))
	return rv
}

// The colors to use for alternating content, typically found in table views and collection views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2998825-alternatingcontentbackgroundcolo?language=objc
func Color_AlternatingContentBackgroundColors() []Color {
	return ColorClass.AlternatingContentBackgroundColors()
}

// Returns a color object for orange that automatically adapts to vibrancy and accessibility settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2879263-systemorangecolor?language=objc
func (cc _ColorClass) SystemOrangeColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("systemOrangeColor"))
	return rv
}

// Returns a color object for orange that automatically adapts to vibrancy and accessibility settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2879263-systemorangecolor?language=objc
func Color_SystemOrangeColor() Color {
	return ColorClass.SystemOrangeColor()
}

// The catalog containing the color’s name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1535443-catalognamecomponent?language=objc
func (c_ Color) CatalogNameComponent() ColorListName {
	rv := objc.Call[ColorListName](c_, objc.Sel("catalogNameComponent"))
	return rv
}

// Returns a color object for blue that automatically adapts to vibrancy and accessibility settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2879260-systembluecolor?language=objc
func (cc _ColorClass) SystemBlueColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("systemBlueColor"))
	return rv
}

// Returns a color object for blue that automatically adapts to vibrancy and accessibility settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/2879260-systembluecolor?language=objc
func Color_SystemBlueColor() Color {
	return ColorClass.SystemBlueColor()
}

// The primary color to use for text labels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1534657-labelcolor?language=objc
func (cc _ColorClass) LabelColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("labelColor"))
	return rv
}

// The primary color to use for text labels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1534657-labelcolor?language=objc
func Color_LabelColor() Color {
	return ColorClass.LabelColor()
}

// The color to use for the keyboard focus ring around controls. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1532031-keyboardfocusindicatorcolor?language=objc
func (cc _ColorClass) KeyboardFocusIndicatorColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("keyboardFocusIndicatorColor"))
	return rv
}

// The color to use for the keyboard focus ring around controls. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1532031-keyboardfocusindicatorcolor?language=objc
func Color_KeyboardFocusIndicatorColor() Color {
	return ColorClass.KeyboardFocusIndicatorColor()
}

// The color to use for the flat surfaces of a control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1524856-controlcolor?language=objc
func (cc _ColorClass) ControlColor() Color {
	rv := objc.Call[Color](cc, objc.Sel("controlColor"))
	return rv
}

// The color to use for the flat surfaces of a control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1524856-controlcolor?language=objc
func Color_ControlColor() Color {
	return ColorClass.ControlColor()
}

// The Core Graphics color object corresponding to the color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/1527738-cgcolor?language=objc
func (c_ Color) CGColor() coregraphics.ColorRef {
	rv := objc.Call[coregraphics.ColorRef](c_, objc.Sel("CGColor"))
	return rv
}
