// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var FontClass = _FontClass{objc.GetClass("NSFont")}

type _FontClass struct {
	objc.Class
}

type IFont interface {
	objc.IObject
	Set()
	SetInContext(graphicsContext IGraphicsContext)
	FontWithSize(fontSize float64) Font
	BoundingRectForCGGlyph(glyph coregraphics.Glyph) foundation.Rect
	GetBoundingRectsForCGGlyphsCount(bounds *foundation.Rect, glyphs *coregraphics.Glyph, glyphCount uint)
	AdvancementForCGGlyph(glyph coregraphics.Glyph) foundation.Size
	GetAdvancementsForCGGlyphsCount(advancements *foundation.Size, glyphs *coregraphics.Glyph, glyphCount uint)
	PointSize() float64
	CoveredCharacterSet() foundation.CharacterSet
	FontDescriptor() FontDescriptor
	IsFixedPitch() bool
	MostCompatibleStringEncoding() foundation.StringEncoding
	NumberOfGlyphs() uint
	DisplayName() string
	FamilyName() string
	FontName() string
	IsVertical() bool
	VerticalFont() Font
	Ascender() float64
	Descender() float64
	CapHeight() float64
	Leading() float64
	XHeight() float64
	ItalicAngle() float64
	UnderlinePosition() float64
	UnderlineThickness() float64
	BoundingRectForFont() foundation.Rect
	MaximumAdvancement() foundation.Size
	Matrix() *float64
	TextTransform() foundation.AffineTransform
}

type Font struct {
	objc.Object
}

func MakeFont(ptr unsafe.Pointer) Font {
	return Font{
		Object: objc.MakeObject(ptr),
	}
}

func (fc _FontClass) Alloc() Font {
	rv := objc.CallMethod[Font](fc, objc.GetSelector("alloc"))
	return rv
}

func Font_Alloc() Font {
	return FontClass.Alloc()
}

func (fc _FontClass) New() Font {
	rv := objc.CallMethod[Font](fc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewFont() Font {
	return FontClass.New()
}

func Font_New() Font {
	return FontClass.New()
}

func (f_ Font) Init() Font {
	rv := objc.CallMethod[Font](f_, objc.GetSelector("init"))
	return rv
}

func Font_Init() Font {
	return FontClass.Alloc().Init()
}

func (fc _FontClass) FontWithNameSize(fontName string, fontSize float64) Font {
	rv := objc.CallMethod[Font](fc, objc.GetSelector("fontWithName:size:"), fontName, fontSize)
	return rv
}

func Font_FontWithNameSize(fontName string, fontSize float64) Font {
	return FontClass.FontWithNameSize(fontName, fontSize)
}

func (fc _FontClass) FontWithDescriptorSize(fontDescriptor IFontDescriptor, fontSize float64) Font {
	rv := objc.CallMethod[Font](fc, objc.GetSelector("fontWithDescriptor:size:"), objc.ExtractPtr(fontDescriptor), fontSize)
	return rv
}

func Font_FontWithDescriptorSize(fontDescriptor IFontDescriptor, fontSize float64) Font {
	return FontClass.FontWithDescriptorSize(fontDescriptor, fontSize)
}

func (fc _FontClass) FontWithDescriptorTextTransform(fontDescriptor IFontDescriptor, textTransform foundation.IAffineTransform) Font {
	rv := objc.CallMethod[Font](fc, objc.GetSelector("fontWithDescriptor:textTransform:"), objc.ExtractPtr(fontDescriptor), objc.ExtractPtr(textTransform))
	return rv
}

func Font_FontWithDescriptorTextTransform(fontDescriptor IFontDescriptor, textTransform foundation.IAffineTransform) Font {
	return FontClass.FontWithDescriptorTextTransform(fontDescriptor, textTransform)
}

func (fc _FontClass) FontWithNameMatrix(fontName string, fontMatrix *float64) Font {
	rv := objc.CallMethod[Font](fc, objc.GetSelector("fontWithName:matrix:"), fontName, fontMatrix)
	return rv
}

func Font_FontWithNameMatrix(fontName string, fontMatrix *float64) Font {
	return FontClass.FontWithNameMatrix(fontName, fontMatrix)
}

func (fc _FontClass) UserFontOfSize(fontSize float64) Font {
	rv := objc.CallMethod[Font](fc, objc.GetSelector("userFontOfSize:"), fontSize)
	return rv
}

func Font_UserFontOfSize(fontSize float64) Font {
	return FontClass.UserFontOfSize(fontSize)
}

func (fc _FontClass) UserFixedPitchFontOfSize(fontSize float64) Font {
	rv := objc.CallMethod[Font](fc, objc.GetSelector("userFixedPitchFontOfSize:"), fontSize)
	return rv
}

func Font_UserFixedPitchFontOfSize(fontSize float64) Font {
	return FontClass.UserFixedPitchFontOfSize(fontSize)
}

func (fc _FontClass) PreferredFontForTextStyleOptions(style FontTextStyle, options map[FontTextStyleOptionKey]objc.IObject) Font {
	rv := objc.CallMethod[Font](fc, objc.GetSelector("preferredFontForTextStyle:options:"), style, options)
	return rv
}

func Font_PreferredFontForTextStyleOptions(style FontTextStyle, options map[FontTextStyleOptionKey]objc.IObject) Font {
	return FontClass.PreferredFontForTextStyleOptions(style, options)
}

func (fc _FontClass) SystemFontOfSize(fontSize float64) Font {
	rv := objc.CallMethod[Font](fc, objc.GetSelector("systemFontOfSize:"), fontSize)
	return rv
}

func Font_SystemFontOfSize(fontSize float64) Font {
	return FontClass.SystemFontOfSize(fontSize)
}

func (fc _FontClass) SystemFontOfSizeWeight(fontSize float64, weight FontWeight) Font {
	rv := objc.CallMethod[Font](fc, objc.GetSelector("systemFontOfSize:weight:"), fontSize, weight)
	return rv
}

func Font_SystemFontOfSizeWeight(fontSize float64, weight FontWeight) Font {
	return FontClass.SystemFontOfSizeWeight(fontSize, weight)
}

func (fc _FontClass) BoldSystemFontOfSize(fontSize float64) Font {
	rv := objc.CallMethod[Font](fc, objc.GetSelector("boldSystemFontOfSize:"), fontSize)
	return rv
}

func Font_BoldSystemFontOfSize(fontSize float64) Font {
	return FontClass.BoldSystemFontOfSize(fontSize)
}

func (fc _FontClass) MonospacedSystemFontOfSizeWeight(fontSize float64, weight FontWeight) Font {
	rv := objc.CallMethod[Font](fc, objc.GetSelector("monospacedSystemFontOfSize:weight:"), fontSize, weight)
	return rv
}

func Font_MonospacedSystemFontOfSizeWeight(fontSize float64, weight FontWeight) Font {
	return FontClass.MonospacedSystemFontOfSizeWeight(fontSize, weight)
}

func (fc _FontClass) MonospacedDigitSystemFontOfSizeWeight(fontSize float64, weight FontWeight) Font {
	rv := objc.CallMethod[Font](fc, objc.GetSelector("monospacedDigitSystemFontOfSize:weight:"), fontSize, weight)
	return rv
}

func Font_MonospacedDigitSystemFontOfSizeWeight(fontSize float64, weight FontWeight) Font {
	return FontClass.MonospacedDigitSystemFontOfSizeWeight(fontSize, weight)
}

func (fc _FontClass) LabelFontOfSize(fontSize float64) Font {
	rv := objc.CallMethod[Font](fc, objc.GetSelector("labelFontOfSize:"), fontSize)
	return rv
}

func Font_LabelFontOfSize(fontSize float64) Font {
	return FontClass.LabelFontOfSize(fontSize)
}

func (fc _FontClass) MessageFontOfSize(fontSize float64) Font {
	rv := objc.CallMethod[Font](fc, objc.GetSelector("messageFontOfSize:"), fontSize)
	return rv
}

func Font_MessageFontOfSize(fontSize float64) Font {
	return FontClass.MessageFontOfSize(fontSize)
}

func (fc _FontClass) MenuBarFontOfSize(fontSize float64) Font {
	rv := objc.CallMethod[Font](fc, objc.GetSelector("menuBarFontOfSize:"), fontSize)
	return rv
}

func Font_MenuBarFontOfSize(fontSize float64) Font {
	return FontClass.MenuBarFontOfSize(fontSize)
}

func (fc _FontClass) MenuFontOfSize(fontSize float64) Font {
	rv := objc.CallMethod[Font](fc, objc.GetSelector("menuFontOfSize:"), fontSize)
	return rv
}

func Font_MenuFontOfSize(fontSize float64) Font {
	return FontClass.MenuFontOfSize(fontSize)
}

func (fc _FontClass) ControlContentFontOfSize(fontSize float64) Font {
	rv := objc.CallMethod[Font](fc, objc.GetSelector("controlContentFontOfSize:"), fontSize)
	return rv
}

func Font_ControlContentFontOfSize(fontSize float64) Font {
	return FontClass.ControlContentFontOfSize(fontSize)
}

func (fc _FontClass) TitleBarFontOfSize(fontSize float64) Font {
	rv := objc.CallMethod[Font](fc, objc.GetSelector("titleBarFontOfSize:"), fontSize)
	return rv
}

func Font_TitleBarFontOfSize(fontSize float64) Font {
	return FontClass.TitleBarFontOfSize(fontSize)
}

func (fc _FontClass) PaletteFontOfSize(fontSize float64) Font {
	rv := objc.CallMethod[Font](fc, objc.GetSelector("paletteFontOfSize:"), fontSize)
	return rv
}

func Font_PaletteFontOfSize(fontSize float64) Font {
	return FontClass.PaletteFontOfSize(fontSize)
}

func (fc _FontClass) ToolTipsFontOfSize(fontSize float64) Font {
	rv := objc.CallMethod[Font](fc, objc.GetSelector("toolTipsFontOfSize:"), fontSize)
	return rv
}

func Font_ToolTipsFontOfSize(fontSize float64) Font {
	return FontClass.ToolTipsFontOfSize(fontSize)
}

func (fc _FontClass) SystemFontSizeForControlSize(controlSize ControlSize) float64 {
	rv := objc.CallMethod[float64](fc, objc.GetSelector("systemFontSizeForControlSize:"), controlSize)
	return rv
}

func Font_SystemFontSizeForControlSize(controlSize ControlSize) float64 {
	return FontClass.SystemFontSizeForControlSize(controlSize)
}

func (f_ Font) Set() {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("set"))
}

func (f_ Font) SetInContext(graphicsContext IGraphicsContext) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("setInContext:"), objc.ExtractPtr(graphicsContext))
}

func (fc _FontClass) SetUserFont(font IFont) {
	objc.CallMethod[objc.Void](fc, objc.GetSelector("setUserFont:"), objc.ExtractPtr(font))
}

func Font_SetUserFont(font IFont) {
	FontClass.SetUserFont(font)
}

func (fc _FontClass) SetUserFixedPitchFont(font IFont) {
	objc.CallMethod[objc.Void](fc, objc.GetSelector("setUserFixedPitchFont:"), objc.ExtractPtr(font))
}

func Font_SetUserFixedPitchFont(font IFont) {
	FontClass.SetUserFixedPitchFont(font)
}

func (f_ Font) FontWithSize(fontSize float64) Font {
	rv := objc.CallMethod[Font](f_, objc.GetSelector("fontWithSize:"), fontSize)
	return rv
}

func (fc _FontClass) SystemFontOfSizeWeightWidth(fontSize float64, weight FontWeight, width FontWidth) Font {
	rv := objc.CallMethod[Font](fc, objc.GetSelector("systemFontOfSize:weight:width:"), fontSize, weight, width)
	return rv
}

func Font_SystemFontOfSizeWeightWidth(fontSize float64, weight FontWeight, width FontWidth) Font {
	return FontClass.SystemFontOfSizeWeightWidth(fontSize, weight, width)
}

func (f_ Font) BoundingRectForCGGlyph(glyph coregraphics.Glyph) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](f_, objc.GetSelector("boundingRectForCGGlyph:"), glyph)
	return rv
}

func (f_ Font) GetBoundingRectsForCGGlyphsCount(bounds *foundation.Rect, glyphs *coregraphics.Glyph, glyphCount uint) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("getBoundingRects:forCGGlyphs:count:"), bounds, glyphs, glyphCount)
}

func (f_ Font) AdvancementForCGGlyph(glyph coregraphics.Glyph) foundation.Size {
	rv := objc.CallMethod[foundation.Size](f_, objc.GetSelector("advancementForCGGlyph:"), glyph)
	return rv
}

func (f_ Font) GetAdvancementsForCGGlyphsCount(advancements *foundation.Size, glyphs *coregraphics.Glyph, glyphCount uint) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("getAdvancements:forCGGlyphs:count:"), advancements, glyphs, glyphCount)
}

func (fc _FontClass) SystemFontSize() float64 {
	rv := objc.CallMethod[float64](fc, objc.GetSelector("systemFontSize"))
	return rv
}

func Font_SystemFontSize() float64 {
	return FontClass.SystemFontSize()
}

func (fc _FontClass) SmallSystemFontSize() float64 {
	rv := objc.CallMethod[float64](fc, objc.GetSelector("smallSystemFontSize"))
	return rv
}

func Font_SmallSystemFontSize() float64 {
	return FontClass.SmallSystemFontSize()
}

func (fc _FontClass) LabelFontSize() float64 {
	rv := objc.CallMethod[float64](fc, objc.GetSelector("labelFontSize"))
	return rv
}

func Font_LabelFontSize() float64 {
	return FontClass.LabelFontSize()
}

func (f_ Font) PointSize() float64 {
	rv := objc.CallMethod[float64](f_, objc.GetSelector("pointSize"))
	return rv
}

func (f_ Font) CoveredCharacterSet() foundation.CharacterSet {
	rv := objc.CallMethod[foundation.CharacterSet](f_, objc.GetSelector("coveredCharacterSet"))
	return rv
}

func (f_ Font) FontDescriptor() FontDescriptor {
	rv := objc.CallMethod[FontDescriptor](f_, objc.GetSelector("fontDescriptor"))
	return rv
}

func (f_ Font) IsFixedPitch() bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("isFixedPitch"))
	return rv
}

func (f_ Font) MostCompatibleStringEncoding() foundation.StringEncoding {
	rv := objc.CallMethod[foundation.StringEncoding](f_, objc.GetSelector("mostCompatibleStringEncoding"))
	return rv
}

func (f_ Font) NumberOfGlyphs() uint {
	rv := objc.CallMethod[uint](f_, objc.GetSelector("numberOfGlyphs"))
	return rv
}

func (f_ Font) DisplayName() string {
	rv := objc.CallMethod[string](f_, objc.GetSelector("displayName"))
	return rv
}

func (f_ Font) FamilyName() string {
	rv := objc.CallMethod[string](f_, objc.GetSelector("familyName"))
	return rv
}

func (f_ Font) FontName() string {
	rv := objc.CallMethod[string](f_, objc.GetSelector("fontName"))
	return rv
}

func (f_ Font) IsVertical() bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("isVertical"))
	return rv
}

func (f_ Font) VerticalFont() Font {
	rv := objc.CallMethod[Font](f_, objc.GetSelector("verticalFont"))
	return rv
}

func (f_ Font) Ascender() float64 {
	rv := objc.CallMethod[float64](f_, objc.GetSelector("ascender"))
	return rv
}

func (f_ Font) Descender() float64 {
	rv := objc.CallMethod[float64](f_, objc.GetSelector("descender"))
	return rv
}

func (f_ Font) CapHeight() float64 {
	rv := objc.CallMethod[float64](f_, objc.GetSelector("capHeight"))
	return rv
}

func (f_ Font) Leading() float64 {
	rv := objc.CallMethod[float64](f_, objc.GetSelector("leading"))
	return rv
}

func (f_ Font) XHeight() float64 {
	rv := objc.CallMethod[float64](f_, objc.GetSelector("xHeight"))
	return rv
}

func (f_ Font) ItalicAngle() float64 {
	rv := objc.CallMethod[float64](f_, objc.GetSelector("italicAngle"))
	return rv
}

func (f_ Font) UnderlinePosition() float64 {
	rv := objc.CallMethod[float64](f_, objc.GetSelector("underlinePosition"))
	return rv
}

func (f_ Font) UnderlineThickness() float64 {
	rv := objc.CallMethod[float64](f_, objc.GetSelector("underlineThickness"))
	return rv
}

func (f_ Font) BoundingRectForFont() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](f_, objc.GetSelector("boundingRectForFont"))
	return rv
}

func (f_ Font) MaximumAdvancement() foundation.Size {
	rv := objc.CallMethod[foundation.Size](f_, objc.GetSelector("maximumAdvancement"))
	return rv
}

func (f_ Font) Matrix() *float64 {
	rv := objc.CallMethod[*float64](f_, objc.GetSelector("matrix"))
	return rv
}

func (f_ Font) TextTransform() foundation.AffineTransform {
	rv := objc.CallMethod[foundation.AffineTransform](f_, objc.GetSelector("textTransform"))
	return rv
}
