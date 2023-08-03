// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var ColorClass = _ColorClass{objc.GetClass("NSColor")}

type _ColorClass struct {
	objc.Class
}

type IColor interface {
	objc.IObject
	ColorWithSystemEffect(systemEffect ColorSystemEffect) Color
	ColorUsingColorSpace(space IColorSpace) Color
	BlendedColorWithFractionOfColor(fraction float64, color IColor) Color
	ColorWithAlphaComponent(alpha float64) Color
	HighlightWithLevel(val float64) Color
	ShadowWithLevel(val float64) Color
	WriteToPasteboard(pasteBoard IPasteboard)
	GetCyanMagentaYellowBlackAlpha(cyan *float64, magenta *float64, yellow *float64, black *float64, alpha *float64)
	GetHueSaturationBrightnessAlpha(hue *float64, saturation *float64, brightness *float64, alpha *float64)
	GetRedGreenBlueAlpha(red *float64, green *float64, blue *float64, alpha *float64)
	GetWhiteAlpha(white *float64, alpha *float64)
	GetComponents(components *float64)
	ColorUsingType(type_ ColorType) Color
	DrawSwatchInRect(rect foundation.Rect)
	Set()
	SetFill()
	SetStroke()
	NumberOfComponents() int
	AlphaComponent() float64
	WhiteComponent() float64
	RedComponent() float64
	GreenComponent() float64
	BlueComponent() float64
	CyanComponent() float64
	MagentaComponent() float64
	YellowComponent() float64
	BlackComponent() float64
	HueComponent() float64
	SaturationComponent() float64
	BrightnessComponent() float64
	CatalogNameComponent() ColorListName
	LocalizedCatalogNameComponent() string
	ColorNameComponent() ColorName
	LocalizedColorNameComponent() string
	Type() ColorType
	ColorSpace() ColorSpace
	CGColor() coregraphics.ColorRef
	PatternImage() Image
}

type Color struct {
	objc.Object
}

func MakeColor(ptr unsafe.Pointer) Color {
	return Color{
		Object: objc.MakeObject(ptr),
	}
}

func (c_ Color) Init() Color {
	rv := objc.CallMethod[Color](c_, objc.GetSelector("init"))
	return rv
}

func Color_Init() Color {
	return ColorClass.Alloc().Init()
}

func (cc _ColorClass) Alloc() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("alloc"))
	return rv
}

func Color_Alloc() Color {
	return ColorClass.Alloc()
}

func (cc _ColorClass) New() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewColor() Color {
	return ColorClass.New()
}

func Color_New() Color {
	return ColorClass.New()
}

func (c_ Color) ColorWithSystemEffect(systemEffect ColorSystemEffect) Color {
	rv := objc.CallMethod[Color](c_, objc.GetSelector("colorWithSystemEffect:"), systemEffect)
	return rv
}

func (c_ Color) ColorUsingColorSpace(space IColorSpace) Color {
	rv := objc.CallMethod[Color](c_, objc.GetSelector("colorUsingColorSpace:"), objc.ExtractPtr(space))
	return rv
}

func (c_ Color) BlendedColorWithFractionOfColor(fraction float64, color IColor) Color {
	rv := objc.CallMethod[Color](c_, objc.GetSelector("blendedColorWithFraction:ofColor:"), fraction, objc.ExtractPtr(color))
	return rv
}

func (c_ Color) ColorWithAlphaComponent(alpha float64) Color {
	rv := objc.CallMethod[Color](c_, objc.GetSelector("colorWithAlphaComponent:"), alpha)
	return rv
}

func (c_ Color) HighlightWithLevel(val float64) Color {
	rv := objc.CallMethod[Color](c_, objc.GetSelector("highlightWithLevel:"), val)
	return rv
}

func (c_ Color) ShadowWithLevel(val float64) Color {
	rv := objc.CallMethod[Color](c_, objc.GetSelector("shadowWithLevel:"), val)
	return rv
}

func (cc _ColorClass) ColorFromPasteboard(pasteBoard IPasteboard) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorFromPasteboard:"), objc.ExtractPtr(pasteBoard))
	return rv
}

func Color_ColorFromPasteboard(pasteBoard IPasteboard) Color {
	return ColorClass.ColorFromPasteboard(pasteBoard)
}

func (c_ Color) WriteToPasteboard(pasteBoard IPasteboard) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("writeToPasteboard:"), objc.ExtractPtr(pasteBoard))
}

func (c_ Color) GetCyanMagentaYellowBlackAlpha(cyan *float64, magenta *float64, yellow *float64, black *float64, alpha *float64) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("getCyan:magenta:yellow:black:alpha:"), cyan, magenta, yellow, black, alpha)
}

func (c_ Color) GetHueSaturationBrightnessAlpha(hue *float64, saturation *float64, brightness *float64, alpha *float64) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("getHue:saturation:brightness:alpha:"), hue, saturation, brightness, alpha)
}

func (c_ Color) GetRedGreenBlueAlpha(red *float64, green *float64, blue *float64, alpha *float64) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("getRed:green:blue:alpha:"), red, green, blue, alpha)
}

func (c_ Color) GetWhiteAlpha(white *float64, alpha *float64) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("getWhite:alpha:"), white, alpha)
}

func (c_ Color) GetComponents(components *float64) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("getComponents:"), components)
}

func (c_ Color) ColorUsingType(type_ ColorType) Color {
	rv := objc.CallMethod[Color](c_, objc.GetSelector("colorUsingType:"), type_)
	return rv
}

func (c_ Color) DrawSwatchInRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("drawSwatchInRect:"), rect)
}

func (c_ Color) Set() {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("set"))
}

func (c_ Color) SetFill() {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setFill"))
}

func (c_ Color) SetStroke() {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setStroke"))
}

func (cc _ColorClass) ColorNamed(name ColorName) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorNamed:"), name)
	return rv
}

func Color_ColorNamed(name ColorName) Color {
	return ColorClass.ColorNamed(name)
}

func (cc _ColorClass) ColorNamedBundle(name ColorName, bundle foundation.IBundle) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorNamed:bundle:"), name, objc.ExtractPtr(bundle))
	return rv
}

func Color_ColorNamedBundle(name ColorName, bundle foundation.IBundle) Color {
	return ColorClass.ColorNamedBundle(name, bundle)
}

func (cc _ColorClass) ColorWithCatalogNameColorName(listName ColorListName, colorName ColorName) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithCatalogName:colorName:"), listName, colorName)
	return rv
}

func Color_ColorWithCatalogNameColorName(listName ColorListName, colorName ColorName) Color {
	return ColorClass.ColorWithCatalogNameColorName(listName, colorName)
}

func (cc _ColorClass) ColorWithSRGBRedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithSRGBRed:green:blue:alpha:"), red, green, blue, alpha)
	return rv
}

func Color_ColorWithSRGBRedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) Color {
	return ColorClass.ColorWithSRGBRedGreenBlueAlpha(red, green, blue, alpha)
}

func (cc _ColorClass) ColorWithDisplayP3RedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithDisplayP3Red:green:blue:alpha:"), red, green, blue, alpha)
	return rv
}

func Color_ColorWithDisplayP3RedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) Color {
	return ColorClass.ColorWithDisplayP3RedGreenBlueAlpha(red, green, blue, alpha)
}

func (cc _ColorClass) ColorWithRedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithRed:green:blue:alpha:"), red, green, blue, alpha)
	return rv
}

func Color_ColorWithRedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) Color {
	return ColorClass.ColorWithRedGreenBlueAlpha(red, green, blue, alpha)
}

func (cc _ColorClass) ColorWithCalibratedRedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithCalibratedRed:green:blue:alpha:"), red, green, blue, alpha)
	return rv
}

func Color_ColorWithCalibratedRedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) Color {
	return ColorClass.ColorWithCalibratedRedGreenBlueAlpha(red, green, blue, alpha)
}

func (cc _ColorClass) ColorWithDeviceRedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithDeviceRed:green:blue:alpha:"), red, green, blue, alpha)
	return rv
}

func Color_ColorWithDeviceRedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) Color {
	return ColorClass.ColorWithDeviceRedGreenBlueAlpha(red, green, blue, alpha)
}

func (cc _ColorClass) ColorWithCalibratedHueSaturationBrightnessAlpha(hue float64, saturation float64, brightness float64, alpha float64) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithCalibratedHue:saturation:brightness:alpha:"), hue, saturation, brightness, alpha)
	return rv
}

func Color_ColorWithCalibratedHueSaturationBrightnessAlpha(hue float64, saturation float64, brightness float64, alpha float64) Color {
	return ColorClass.ColorWithCalibratedHueSaturationBrightnessAlpha(hue, saturation, brightness, alpha)
}

func (cc _ColorClass) ColorWithDeviceHueSaturationBrightnessAlpha(hue float64, saturation float64, brightness float64, alpha float64) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithDeviceHue:saturation:brightness:alpha:"), hue, saturation, brightness, alpha)
	return rv
}

func Color_ColorWithDeviceHueSaturationBrightnessAlpha(hue float64, saturation float64, brightness float64, alpha float64) Color {
	return ColorClass.ColorWithDeviceHueSaturationBrightnessAlpha(hue, saturation, brightness, alpha)
}

func (cc _ColorClass) ColorWithHueSaturationBrightnessAlpha(hue float64, saturation float64, brightness float64, alpha float64) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithHue:saturation:brightness:alpha:"), hue, saturation, brightness, alpha)
	return rv
}

func Color_ColorWithHueSaturationBrightnessAlpha(hue float64, saturation float64, brightness float64, alpha float64) Color {
	return ColorClass.ColorWithHueSaturationBrightnessAlpha(hue, saturation, brightness, alpha)
}

func (cc _ColorClass) ColorWithColorSpaceHueSaturationBrightnessAlpha(space IColorSpace, hue float64, saturation float64, brightness float64, alpha float64) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithColorSpace:hue:saturation:brightness:alpha:"), objc.ExtractPtr(space), hue, saturation, brightness, alpha)
	return rv
}

func Color_ColorWithColorSpaceHueSaturationBrightnessAlpha(space IColorSpace, hue float64, saturation float64, brightness float64, alpha float64) Color {
	return ColorClass.ColorWithColorSpaceHueSaturationBrightnessAlpha(space, hue, saturation, brightness, alpha)
}

func (cc _ColorClass) ColorWithDeviceCyanMagentaYellowBlackAlpha(cyan float64, magenta float64, yellow float64, black float64, alpha float64) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithDeviceCyan:magenta:yellow:black:alpha:"), cyan, magenta, yellow, black, alpha)
	return rv
}

func Color_ColorWithDeviceCyanMagentaYellowBlackAlpha(cyan float64, magenta float64, yellow float64, black float64, alpha float64) Color {
	return ColorClass.ColorWithDeviceCyanMagentaYellowBlackAlpha(cyan, magenta, yellow, black, alpha)
}

func (cc _ColorClass) ColorWithWhiteAlpha(white float64, alpha float64) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithWhite:alpha:"), white, alpha)
	return rv
}

func Color_ColorWithWhiteAlpha(white float64, alpha float64) Color {
	return ColorClass.ColorWithWhiteAlpha(white, alpha)
}

func (cc _ColorClass) ColorWithCalibratedWhiteAlpha(white float64, alpha float64) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithCalibratedWhite:alpha:"), white, alpha)
	return rv
}

func Color_ColorWithCalibratedWhiteAlpha(white float64, alpha float64) Color {
	return ColorClass.ColorWithCalibratedWhiteAlpha(white, alpha)
}

func (cc _ColorClass) ColorWithDeviceWhiteAlpha(white float64, alpha float64) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithDeviceWhite:alpha:"), white, alpha)
	return rv
}

func Color_ColorWithDeviceWhiteAlpha(white float64, alpha float64) Color {
	return ColorClass.ColorWithDeviceWhiteAlpha(white, alpha)
}

func (cc _ColorClass) ColorWithGenericGamma22WhiteAlpha(white float64, alpha float64) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithGenericGamma22White:alpha:"), white, alpha)
	return rv
}

func Color_ColorWithGenericGamma22WhiteAlpha(white float64, alpha float64) Color {
	return ColorClass.ColorWithGenericGamma22WhiteAlpha(white, alpha)
}

func (cc _ColorClass) ColorWithPatternImage(image IImage) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithPatternImage:"), objc.ExtractPtr(image))
	return rv
}

func Color_ColorWithPatternImage(image IImage) Color {
	return ColorClass.ColorWithPatternImage(image)
}

func (cc _ColorClass) ColorWithNameDynamicProvider(colorName ColorName, dynamicProvider func(param1 Appearance) Color) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithName:dynamicProvider:"), colorName, dynamicProvider)
	return rv
}

func Color_ColorWithNameDynamicProvider(colorName ColorName, dynamicProvider func(param1 Appearance) Color) Color {
	return ColorClass.ColorWithNameDynamicProvider(colorName, dynamicProvider)
}

func (cc _ColorClass) ColorWithColorSpaceComponentsCount(space IColorSpace, components *float64, numberOfComponents int) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithColorSpace:components:count:"), objc.ExtractPtr(space), components, numberOfComponents)
	return rv
}

func Color_ColorWithColorSpaceComponentsCount(space IColorSpace, components *float64, numberOfComponents int) Color {
	return ColorClass.ColorWithColorSpaceComponentsCount(space, components, numberOfComponents)
}

func (cc _ColorClass) ColorWithCGColor(cgColor coregraphics.ColorRef) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithCGColor:"), cgColor)
	return rv
}

func Color_ColorWithCGColor(cgColor coregraphics.ColorRef) Color {
	return ColorClass.ColorWithCGColor(cgColor)
}

func (cc _ColorClass) IgnoresAlpha() bool {
	rv := objc.CallMethod[bool](cc, objc.GetSelector("ignoresAlpha"))
	return rv
}

func Color_IgnoresAlpha() bool {
	return ColorClass.IgnoresAlpha()
}

func (cc _ColorClass) SetIgnoresAlpha(value bool) {
	objc.CallMethod[objc.Void](cc, objc.GetSelector("setIgnoresAlpha:"), value)
}

func Color_SetIgnoresAlpha(value bool) {
	ColorClass.SetIgnoresAlpha(value)
}

func (c_ Color) NumberOfComponents() int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("numberOfComponents"))
	return rv
}

func (c_ Color) AlphaComponent() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("alphaComponent"))
	return rv
}

func (c_ Color) WhiteComponent() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("whiteComponent"))
	return rv
}

func (c_ Color) RedComponent() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("redComponent"))
	return rv
}

func (c_ Color) GreenComponent() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("greenComponent"))
	return rv
}

func (c_ Color) BlueComponent() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("blueComponent"))
	return rv
}

func (c_ Color) CyanComponent() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("cyanComponent"))
	return rv
}

func (c_ Color) MagentaComponent() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("magentaComponent"))
	return rv
}

func (c_ Color) YellowComponent() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("yellowComponent"))
	return rv
}

func (c_ Color) BlackComponent() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("blackComponent"))
	return rv
}

func (c_ Color) HueComponent() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("hueComponent"))
	return rv
}

func (c_ Color) SaturationComponent() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("saturationComponent"))
	return rv
}

func (c_ Color) BrightnessComponent() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("brightnessComponent"))
	return rv
}

func (c_ Color) CatalogNameComponent() ColorListName {
	rv := objc.CallMethod[ColorListName](c_, objc.GetSelector("catalogNameComponent"))
	return rv
}

func (c_ Color) LocalizedCatalogNameComponent() string {
	rv := objc.CallMethod[string](c_, objc.GetSelector("localizedCatalogNameComponent"))
	return rv
}

func (c_ Color) ColorNameComponent() ColorName {
	rv := objc.CallMethod[ColorName](c_, objc.GetSelector("colorNameComponent"))
	return rv
}

func (c_ Color) LocalizedColorNameComponent() string {
	rv := objc.CallMethod[string](c_, objc.GetSelector("localizedColorNameComponent"))
	return rv
}

func (c_ Color) Type() ColorType {
	rv := objc.CallMethod[ColorType](c_, objc.GetSelector("type"))
	return rv
}

func (c_ Color) ColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](c_, objc.GetSelector("colorSpace"))
	return rv
}

func (c_ Color) CGColor() coregraphics.ColorRef {
	rv := objc.CallMethod[coregraphics.ColorRef](c_, objc.GetSelector("CGColor"))
	return rv
}

func (cc _ColorClass) SystemCyanColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("systemCyanColor"))
	return rv
}

func Color_SystemCyanColor() Color {
	return ColorClass.SystemCyanColor()
}

func (cc _ColorClass) SystemMintColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("systemMintColor"))
	return rv
}

func Color_SystemMintColor() Color {
	return ColorClass.SystemMintColor()
}

func (cc _ColorClass) LabelColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("labelColor"))
	return rv
}

func Color_LabelColor() Color {
	return ColorClass.LabelColor()
}

func (cc _ColorClass) SecondaryLabelColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("secondaryLabelColor"))
	return rv
}

func Color_SecondaryLabelColor() Color {
	return ColorClass.SecondaryLabelColor()
}

func (cc _ColorClass) TertiaryLabelColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("tertiaryLabelColor"))
	return rv
}

func Color_TertiaryLabelColor() Color {
	return ColorClass.TertiaryLabelColor()
}

func (cc _ColorClass) QuaternaryLabelColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("quaternaryLabelColor"))
	return rv
}

func Color_QuaternaryLabelColor() Color {
	return ColorClass.QuaternaryLabelColor()
}

func (cc _ColorClass) TextColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("textColor"))
	return rv
}

func Color_TextColor() Color {
	return ColorClass.TextColor()
}

func (cc _ColorClass) PlaceholderTextColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("placeholderTextColor"))
	return rv
}

func Color_PlaceholderTextColor() Color {
	return ColorClass.PlaceholderTextColor()
}

func (cc _ColorClass) SelectedTextColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("selectedTextColor"))
	return rv
}

func Color_SelectedTextColor() Color {
	return ColorClass.SelectedTextColor()
}

func (cc _ColorClass) TextBackgroundColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("textBackgroundColor"))
	return rv
}

func Color_TextBackgroundColor() Color {
	return ColorClass.TextBackgroundColor()
}

func (cc _ColorClass) SelectedTextBackgroundColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("selectedTextBackgroundColor"))
	return rv
}

func Color_SelectedTextBackgroundColor() Color {
	return ColorClass.SelectedTextBackgroundColor()
}

func (cc _ColorClass) KeyboardFocusIndicatorColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("keyboardFocusIndicatorColor"))
	return rv
}

func Color_KeyboardFocusIndicatorColor() Color {
	return ColorClass.KeyboardFocusIndicatorColor()
}

func (cc _ColorClass) UnemphasizedSelectedTextColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("unemphasizedSelectedTextColor"))
	return rv
}

func Color_UnemphasizedSelectedTextColor() Color {
	return ColorClass.UnemphasizedSelectedTextColor()
}

func (cc _ColorClass) UnemphasizedSelectedTextBackgroundColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("unemphasizedSelectedTextBackgroundColor"))
	return rv
}

func Color_UnemphasizedSelectedTextBackgroundColor() Color {
	return ColorClass.UnemphasizedSelectedTextBackgroundColor()
}

func (cc _ColorClass) LinkColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("linkColor"))
	return rv
}

func Color_LinkColor() Color {
	return ColorClass.LinkColor()
}

func (cc _ColorClass) SeparatorColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("separatorColor"))
	return rv
}

func Color_SeparatorColor() Color {
	return ColorClass.SeparatorColor()
}

func (cc _ColorClass) SelectedContentBackgroundColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("selectedContentBackgroundColor"))
	return rv
}

func Color_SelectedContentBackgroundColor() Color {
	return ColorClass.SelectedContentBackgroundColor()
}

func (cc _ColorClass) UnemphasizedSelectedContentBackgroundColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("unemphasizedSelectedContentBackgroundColor"))
	return rv
}

func Color_UnemphasizedSelectedContentBackgroundColor() Color {
	return ColorClass.UnemphasizedSelectedContentBackgroundColor()
}

func (cc _ColorClass) SelectedMenuItemTextColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("selectedMenuItemTextColor"))
	return rv
}

func Color_SelectedMenuItemTextColor() Color {
	return ColorClass.SelectedMenuItemTextColor()
}

func (cc _ColorClass) GridColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("gridColor"))
	return rv
}

func Color_GridColor() Color {
	return ColorClass.GridColor()
}

func (cc _ColorClass) HeaderTextColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("headerTextColor"))
	return rv
}

func Color_HeaderTextColor() Color {
	return ColorClass.HeaderTextColor()
}

func (cc _ColorClass) AlternatingContentBackgroundColors() []Color {
	rv := objc.CallMethod[[]Color](cc, objc.GetSelector("alternatingContentBackgroundColors"))
	// TODO: convert slice items...
	return rv
}

func Color_AlternatingContentBackgroundColors() []Color {
	return ColorClass.AlternatingContentBackgroundColors()
}

func (cc _ColorClass) ControlAccentColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("controlAccentColor"))
	return rv
}

func Color_ControlAccentColor() Color {
	return ColorClass.ControlAccentColor()
}

func (cc _ColorClass) ControlColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("controlColor"))
	return rv
}

func Color_ControlColor() Color {
	return ColorClass.ControlColor()
}

func (cc _ColorClass) ControlBackgroundColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("controlBackgroundColor"))
	return rv
}

func Color_ControlBackgroundColor() Color {
	return ColorClass.ControlBackgroundColor()
}

func (cc _ColorClass) ControlTextColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("controlTextColor"))
	return rv
}

func Color_ControlTextColor() Color {
	return ColorClass.ControlTextColor()
}

func (cc _ColorClass) DisabledControlTextColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("disabledControlTextColor"))
	return rv
}

func Color_DisabledControlTextColor() Color {
	return ColorClass.DisabledControlTextColor()
}

func (cc _ColorClass) CurrentControlTint() ControlTint {
	rv := objc.CallMethod[ControlTint](cc, objc.GetSelector("currentControlTint"))
	return rv
}

func Color_CurrentControlTint() ControlTint {
	return ColorClass.CurrentControlTint()
}

func (cc _ColorClass) SelectedControlColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("selectedControlColor"))
	return rv
}

func Color_SelectedControlColor() Color {
	return ColorClass.SelectedControlColor()
}

func (cc _ColorClass) SelectedControlTextColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("selectedControlTextColor"))
	return rv
}

func Color_SelectedControlTextColor() Color {
	return ColorClass.SelectedControlTextColor()
}

func (cc _ColorClass) AlternateSelectedControlTextColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("alternateSelectedControlTextColor"))
	return rv
}

func Color_AlternateSelectedControlTextColor() Color {
	return ColorClass.AlternateSelectedControlTextColor()
}

func (cc _ColorClass) ScrubberTexturedBackgroundColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("scrubberTexturedBackgroundColor"))
	return rv
}

func Color_ScrubberTexturedBackgroundColor() Color {
	return ColorClass.ScrubberTexturedBackgroundColor()
}

func (cc _ColorClass) WindowBackgroundColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("windowBackgroundColor"))
	return rv
}

func Color_WindowBackgroundColor() Color {
	return ColorClass.WindowBackgroundColor()
}

func (cc _ColorClass) WindowFrameTextColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("windowFrameTextColor"))
	return rv
}

func Color_WindowFrameTextColor() Color {
	return ColorClass.WindowFrameTextColor()
}

func (cc _ColorClass) UnderPageBackgroundColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("underPageBackgroundColor"))
	return rv
}

func Color_UnderPageBackgroundColor() Color {
	return ColorClass.UnderPageBackgroundColor()
}

func (cc _ColorClass) FindHighlightColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("findHighlightColor"))
	return rv
}

func Color_FindHighlightColor() Color {
	return ColorClass.FindHighlightColor()
}

func (cc _ColorClass) HighlightColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("highlightColor"))
	return rv
}

func Color_HighlightColor() Color {
	return ColorClass.HighlightColor()
}

func (cc _ColorClass) ShadowColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("shadowColor"))
	return rv
}

func Color_ShadowColor() Color {
	return ColorClass.ShadowColor()
}

func (cc _ColorClass) SystemBlueColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("systemBlueColor"))
	return rv
}

func Color_SystemBlueColor() Color {
	return ColorClass.SystemBlueColor()
}

func (cc _ColorClass) SystemBrownColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("systemBrownColor"))
	return rv
}

func Color_SystemBrownColor() Color {
	return ColorClass.SystemBrownColor()
}

func (cc _ColorClass) SystemGrayColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("systemGrayColor"))
	return rv
}

func Color_SystemGrayColor() Color {
	return ColorClass.SystemGrayColor()
}

func (cc _ColorClass) SystemGreenColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("systemGreenColor"))
	return rv
}

func Color_SystemGreenColor() Color {
	return ColorClass.SystemGreenColor()
}

func (cc _ColorClass) SystemIndigoColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("systemIndigoColor"))
	return rv
}

func Color_SystemIndigoColor() Color {
	return ColorClass.SystemIndigoColor()
}

func (cc _ColorClass) SystemOrangeColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("systemOrangeColor"))
	return rv
}

func Color_SystemOrangeColor() Color {
	return ColorClass.SystemOrangeColor()
}

func (cc _ColorClass) SystemPinkColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("systemPinkColor"))
	return rv
}

func Color_SystemPinkColor() Color {
	return ColorClass.SystemPinkColor()
}

func (cc _ColorClass) SystemPurpleColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("systemPurpleColor"))
	return rv
}

func Color_SystemPurpleColor() Color {
	return ColorClass.SystemPurpleColor()
}

func (cc _ColorClass) SystemRedColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("systemRedColor"))
	return rv
}

func Color_SystemRedColor() Color {
	return ColorClass.SystemRedColor()
}

func (cc _ColorClass) SystemTealColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("systemTealColor"))
	return rv
}

func Color_SystemTealColor() Color {
	return ColorClass.SystemTealColor()
}

func (cc _ColorClass) SystemYellowColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("systemYellowColor"))
	return rv
}

func Color_SystemYellowColor() Color {
	return ColorClass.SystemYellowColor()
}

func (cc _ColorClass) ClearColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("clearColor"))
	return rv
}

func Color_ClearColor() Color {
	return ColorClass.ClearColor()
}

func (cc _ColorClass) BlackColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("blackColor"))
	return rv
}

func Color_BlackColor() Color {
	return ColorClass.BlackColor()
}

func (cc _ColorClass) BlueColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("blueColor"))
	return rv
}

func Color_BlueColor() Color {
	return ColorClass.BlueColor()
}

func (cc _ColorClass) BrownColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("brownColor"))
	return rv
}

func Color_BrownColor() Color {
	return ColorClass.BrownColor()
}

func (cc _ColorClass) CyanColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("cyanColor"))
	return rv
}

func Color_CyanColor() Color {
	return ColorClass.CyanColor()
}

func (cc _ColorClass) DarkGrayColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("darkGrayColor"))
	return rv
}

func Color_DarkGrayColor() Color {
	return ColorClass.DarkGrayColor()
}

func (cc _ColorClass) GrayColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("grayColor"))
	return rv
}

func Color_GrayColor() Color {
	return ColorClass.GrayColor()
}

func (cc _ColorClass) GreenColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("greenColor"))
	return rv
}

func Color_GreenColor() Color {
	return ColorClass.GreenColor()
}

func (cc _ColorClass) LightGrayColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("lightGrayColor"))
	return rv
}

func Color_LightGrayColor() Color {
	return ColorClass.LightGrayColor()
}

func (cc _ColorClass) MagentaColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("magentaColor"))
	return rv
}

func Color_MagentaColor() Color {
	return ColorClass.MagentaColor()
}

func (cc _ColorClass) OrangeColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("orangeColor"))
	return rv
}

func Color_OrangeColor() Color {
	return ColorClass.OrangeColor()
}

func (cc _ColorClass) PurpleColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("purpleColor"))
	return rv
}

func Color_PurpleColor() Color {
	return ColorClass.PurpleColor()
}

func (cc _ColorClass) RedColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("redColor"))
	return rv
}

func Color_RedColor() Color {
	return ColorClass.RedColor()
}

func (cc _ColorClass) WhiteColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("whiteColor"))
	return rv
}

func Color_WhiteColor() Color {
	return ColorClass.WhiteColor()
}

func (cc _ColorClass) YellowColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("yellowColor"))
	return rv
}

func Color_YellowColor() Color {
	return ColorClass.YellowColor()
}

func (c_ Color) PatternImage() Image {
	rv := objc.CallMethod[Image](c_, objc.GetSelector("patternImage"))
	return rv
}
