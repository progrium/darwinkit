// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Font] class.
var FontClass = _FontClass{objc.GetClass("NSFont")}

type _FontClass struct {
	objc.Class
}

// An interface definition for the [Font] class.
type IFont interface {
	objc.IObject
	GetBoundingRectsForCGGlyphsCount(bounds foundation.RectArray, glyphs *coregraphics.Glyph, glyphCount uint)
	BoundingRectForCGGlyph(glyph coregraphics.Glyph) foundation.Rect
	Set()
	FontWithSize(fontSize float64) Font
	GetAdvancementsForCGGlyphsCount(advancements foundation.SizeArray, glyphs *coregraphics.Glyph, glyphCount uint)
	SetInContext(graphicsContext IGraphicsContext)
	AdvancementForCGGlyph(glyph coregraphics.Glyph) foundation.Size
	FontDescriptor() FontDescriptor
	Ascender() float64
	CapHeight() float64
	IsFixedPitch() bool
	CoveredCharacterSet() foundation.CharacterSet
	TextTransform() foundation.AffineTransform
	UnderlinePosition() float64
	Leading() float64
	MaximumAdvancement() foundation.Size
	NumberOfGlyphs() uint
	XHeight() float64
	PointSize() float64
	VerticalFont() Font
	BoundingRectForFont() foundation.Rect
	Descender() float64
	FontName() string
	Matrix() *float64
	IsVertical() bool
	ItalicAngle() float64
	MostCompatibleStringEncoding() foundation.StringEncoding
	UnderlineThickness() float64
	DisplayName() string
	FamilyName() string
}

// The representation of a font in an app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont?language=objc
type Font struct {
	objc.Object
}

func FontFrom(ptr unsafe.Pointer) Font {
	return Font{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FontClass) Alloc() Font {
	rv := objc.Call[Font](fc, objc.Sel("alloc"))
	return rv
}

func Font_Alloc() Font {
	return FontClass.Alloc()
}

func (fc _FontClass) New() Font {
	rv := objc.Call[Font](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFont() Font {
	return FontClass.New()
}

func (f_ Font) Init() Font {
	rv := objc.Call[Font](f_, objc.Sel("init"))
	return rv
}

// Returns an array of the bounding rectangles for the specified glyphs rendered by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/2887175-getboundingrects?language=objc
func (f_ Font) GetBoundingRectsForCGGlyphsCount(bounds foundation.RectArray, glyphs *coregraphics.Glyph, glyphCount uint) {
	objc.Call[objc.Void](f_, objc.Sel("getBoundingRects:forCGGlyphs:count:"), bounds, glyphs, glyphCount)
}

// Returns the font used for palette window title bars, in the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1535462-palettefontofsize?language=objc
func (fc _FontClass) PaletteFontOfSize(fontSize float64) Font {
	rv := objc.Call[Font](fc, objc.Sel("paletteFontOfSize:"), fontSize)
	return rv
}

// Returns the font used for palette window title bars, in the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1535462-palettefontofsize?language=objc
func Font_PaletteFontOfSize(fontSize float64) Font {
	return FontClass.PaletteFontOfSize(fontSize)
}

// Returns the font used for window title bars, in the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1530200-titlebarfontofsize?language=objc
func (fc _FontClass) TitleBarFontOfSize(fontSize float64) Font {
	rv := objc.Call[Font](fc, objc.Sel("titleBarFontOfSize:"), fontSize)
	return rv
}

// Returns the font used for window title bars, in the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1530200-titlebarfontofsize?language=objc
func Font_TitleBarFontOfSize(fontSize float64) Font {
	return FontClass.TitleBarFontOfSize(fontSize)
}

// Returns the bounding rectangle for the specified glyph, scaled to the receiver’s size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/2887147-boundingrectforcgglyph?language=objc
func (f_ Font) BoundingRectForCGGlyph(glyph coregraphics.Glyph) foundation.Rect {
	rv := objc.Call[foundation.Rect](f_, objc.Sel("boundingRectForCGGlyph:"), glyph)
	return rv
}

// Sets this font as the font for the current graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1531373-set?language=objc
func (f_ Font) Set() {
	objc.Call[objc.Void](f_, objc.Sel("set"))
}

// Sets the font used by default for documents and other text under the user’s control to the specified font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1526068-setuserfont?language=objc
func (fc _FontClass) SetUserFont(font IFont) {
	objc.Call[objc.Void](fc, objc.Sel("setUserFont:"), objc.Ptr(font))
}

// Sets the font used by default for documents and other text under the user’s control to the specified font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1526068-setuserfont?language=objc
func Font_SetUserFont(font IFont) {
	FontClass.SetUserFont(font)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/3667454-fontwithsize?language=objc
func (f_ Font) FontWithSize(fontSize float64) Font {
	rv := objc.Call[Font](f_, objc.Sel("fontWithSize:"), fontSize)
	return rv
}

// Returns a version of the standard system font that contains monospaced digit glyphs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1528205-monospaceddigitsystemfontofsize?language=objc
func (fc _FontClass) MonospacedDigitSystemFontOfSizeWeight(fontSize float64, weight FontWeight) Font {
	rv := objc.Call[Font](fc, objc.Sel("monospacedDigitSystemFontOfSize:weight:"), fontSize, weight)
	return rv
}

// Returns a version of the standard system font that contains monospaced digit glyphs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1528205-monospaceddigitsystemfontofsize?language=objc
func Font_MonospacedDigitSystemFontOfSizeWeight(fontSize float64, weight FontWeight) Font {
	return FontClass.MonospacedDigitSystemFontOfSizeWeight(fontSize, weight)
}

// Returns the font associated with the text style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/3553195-preferredfontfortextstyle?language=objc
func (fc _FontClass) PreferredFontForTextStyleOptions(style FontTextStyle, options map[FontTextStyleOptionKey]objc.IObject) Font {
	rv := objc.Call[Font](fc, objc.Sel("preferredFontForTextStyle:options:"), style, options)
	return rv
}

// Returns the font associated with the text style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/3553195-preferredfontfortextstyle?language=objc
func Font_PreferredFontForTextStyleOptions(style FontTextStyle, options map[FontTextStyleOptionKey]objc.IObject) Font {
	return FontClass.PreferredFontForTextStyleOptions(style, options)
}

// Returns the font used for standard interface items, such as button labels, menu items, and so on, in the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1525777-messagefontofsize?language=objc
func (fc _FontClass) MessageFontOfSize(fontSize float64) Font {
	rv := objc.Call[Font](fc, objc.Sel("messageFontOfSize:"), fontSize)
	return rv
}

// Returns the font used for standard interface items, such as button labels, menu items, and so on, in the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1525777-messagefontofsize?language=objc
func Font_MessageFontOfSize(fontSize float64) Font {
	return FontClass.MessageFontOfSize(fontSize)
}

// Returns the standard system font in boldface type with the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1533549-boldsystemfontofsize?language=objc
func (fc _FontClass) BoldSystemFontOfSize(fontSize float64) Font {
	rv := objc.Call[Font](fc, objc.Sel("boldSystemFontOfSize:"), fontSize)
	return rv
}

// Returns the standard system font in boldface type with the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1533549-boldsystemfontofsize?language=objc
func Font_BoldSystemFontOfSize(fontSize float64) Font {
	return FontClass.BoldSystemFontOfSize(fontSize)
}

// Returns the font used for menu items, in the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1533068-menufontofsize?language=objc
func (fc _FontClass) MenuFontOfSize(fontSize float64) Font {
	rv := objc.Call[Font](fc, objc.Sel("menuFontOfSize:"), fontSize)
	return rv
}

// Returns the font used for menu items, in the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1533068-menufontofsize?language=objc
func Font_MenuFontOfSize(fontSize float64) Font {
	return FontClass.MenuFontOfSize(fontSize)
}

// Returns the standard system font with the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1530094-systemfontofsize?language=objc
func (fc _FontClass) SystemFontOfSize(fontSize float64) Font {
	rv := objc.Call[Font](fc, objc.Sel("systemFontOfSize:"), fontSize)
	return rv
}

// Returns the standard system font with the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1530094-systemfontofsize?language=objc
func Font_SystemFontOfSize(fontSize float64) Font {
	return FontClass.SystemFontOfSize(fontSize)
}

// Returns an array of the advancements for the specified glyphs rendered by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/2887171-getadvancements?language=objc
func (f_ Font) GetAdvancementsForCGGlyphsCount(advancements foundation.SizeArray, glyphs *coregraphics.Glyph, glyphCount uint) {
	objc.Call[objc.Void](f_, objc.Sel("getAdvancements:forCGGlyphs:count:"), advancements, glyphs, glyphCount)
}

// Returns the font used for tool tips labels, in the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1527704-tooltipsfontofsize?language=objc
func (fc _FontClass) ToolTipsFontOfSize(fontSize float64) Font {
	rv := objc.Call[Font](fc, objc.Sel("toolTipsFontOfSize:"), fontSize)
	return rv
}

// Returns the font used for tool tips labels, in the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1527704-tooltipsfontofsize?language=objc
func Font_ToolTipsFontOfSize(fontSize float64) Font {
	return FontClass.ToolTipsFontOfSize(fontSize)
}

// Returns a font object for the specified font name and matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1530751-fontwithname?language=objc
func (fc _FontClass) FontWithNameMatrix(fontName string, fontMatrix *float64) Font {
	rv := objc.Call[Font](fc, objc.Sel("fontWithName:matrix:"), fontName, fontMatrix)
	return rv
}

// Returns a font object for the specified font name and matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1530751-fontwithname?language=objc
func Font_FontWithNameMatrix(fontName string, fontMatrix *float64) Font {
	return FontClass.FontWithNameMatrix(fontName, fontMatrix)
}

// Returns a monospace version of the system font with the specified size and weight. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/3042659-monospacedsystemfontofsize?language=objc
func (fc _FontClass) MonospacedSystemFontOfSizeWeight(fontSize float64, weight FontWeight) Font {
	rv := objc.Call[Font](fc, objc.Sel("monospacedSystemFontOfSize:weight:"), fontSize, weight)
	return rv
}

// Returns a monospace version of the system font with the specified size and weight. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/3042659-monospacedsystemfontofsize?language=objc
func Font_MonospacedSystemFontOfSizeWeight(fontSize float64, weight FontWeight) Font {
	return FontClass.MonospacedSystemFontOfSizeWeight(fontSize, weight)
}

// Sets this font as the font for the specified graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1534538-setincontext?language=objc
func (f_ Font) SetInContext(graphicsContext IGraphicsContext) {
	objc.Call[objc.Void](f_, objc.Sel("setInContext:"), objc.Ptr(graphicsContext))
}

// Returns the nominal spacing for the given glyph—the distance the current point moves after showing the glyph—accounting for the receiver’s size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/2887191-advancementforcgglyph?language=objc
func (f_ Font) AdvancementForCGGlyph(glyph coregraphics.Glyph) foundation.Size {
	rv := objc.Call[foundation.Size](f_, objc.Sel("advancementForCGGlyph:"), glyph)
	return rv
}

// Returns the font used for menu bar items, in the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1534194-menubarfontofsize?language=objc
func (fc _FontClass) MenuBarFontOfSize(fontSize float64) Font {
	rv := objc.Call[Font](fc, objc.Sel("menuBarFontOfSize:"), fontSize)
	return rv
}

// Returns the font used for menu bar items, in the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1534194-menubarfontofsize?language=objc
func Font_MenuBarFontOfSize(fontSize float64) Font {
	return FontClass.MenuBarFontOfSize(fontSize)
}

// Returns a font object for the specified font descriptor and font size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1525386-fontwithdescriptor?language=objc
func (fc _FontClass) FontWithDescriptorSize(fontDescriptor IFontDescriptor, fontSize float64) Font {
	rv := objc.Call[Font](fc, objc.Sel("fontWithDescriptor:size:"), objc.Ptr(fontDescriptor), fontSize)
	return rv
}

// Returns a font object for the specified font descriptor and font size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1525386-fontwithdescriptor?language=objc
func Font_FontWithDescriptorSize(fontDescriptor IFontDescriptor, fontSize float64) Font {
	return FontClass.FontWithDescriptorSize(fontDescriptor, fontSize)
}

// Returns the font used by default for documents and other text under the user’s control (that is, text whose font the user can normally change), when that font should be fixed-pitch, in the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1531381-userfixedpitchfontofsize?language=objc
func (fc _FontClass) UserFixedPitchFontOfSize(fontSize float64) Font {
	rv := objc.Call[Font](fc, objc.Sel("userFixedPitchFontOfSize:"), fontSize)
	return rv
}

// Returns the font used by default for documents and other text under the user’s control (that is, text whose font the user can normally change), when that font should be fixed-pitch, in the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1531381-userfixedpitchfontofsize?language=objc
func Font_UserFixedPitchFontOfSize(fontSize float64) Font {
	return FontClass.UserFixedPitchFontOfSize(fontSize)
}

// Sets the font used by default for documents and other text under the user’s control, when that font should be fixed-pitch, to the specified font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1529050-setuserfixedpitchfont?language=objc
func (fc _FontClass) SetUserFixedPitchFont(font IFont) {
	objc.Call[objc.Void](fc, objc.Sel("setUserFixedPitchFont:"), objc.Ptr(font))
}

// Sets the font used by default for documents and other text under the user’s control, when that font should be fixed-pitch, to the specified font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1529050-setuserfixedpitchfont?language=objc
func Font_SetUserFixedPitchFont(font IFont) {
	FontClass.SetUserFixedPitchFont(font)
}

// Returns the font used for the content of controls in the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1527070-controlcontentfontofsize?language=objc
func (fc _FontClass) ControlContentFontOfSize(fontSize float64) Font {
	rv := objc.Call[Font](fc, objc.Sel("controlContentFontOfSize:"), fontSize)
	return rv
}

// Returns the font used for the content of controls in the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1527070-controlcontentfontofsize?language=objc
func Font_ControlContentFontOfSize(fontSize float64) Font {
	return FontClass.ControlContentFontOfSize(fontSize)
}

// Returns the font used for standard interface labels in the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1528213-labelfontofsize?language=objc
func (fc _FontClass) LabelFontOfSize(fontSize float64) Font {
	rv := objc.Call[Font](fc, objc.Sel("labelFontOfSize:"), fontSize)
	return rv
}

// Returns the font used for standard interface labels in the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1528213-labelfontofsize?language=objc
func Font_LabelFontOfSize(fontSize float64) Font {
	return FontClass.LabelFontOfSize(fontSize)
}

// Returns the font size used for the specified control size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1529747-systemfontsizeforcontrolsize?language=objc
func (fc _FontClass) SystemFontSizeForControlSize(controlSize ControlSize) float64 {
	rv := objc.Call[float64](fc, objc.Sel("systemFontSizeForControlSize:"), controlSize)
	return rv
}

// Returns the font size used for the specified control size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1529747-systemfontsizeforcontrolsize?language=objc
func Font_SystemFontSizeForControlSize(controlSize ControlSize) float64 {
	return FontClass.SystemFontSizeForControlSize(controlSize)
}

// Returns the font used by default for documents and other text under the user’s control (that is, text whose font the user can normally change), in the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1524559-userfontofsize?language=objc
func (fc _FontClass) UserFontOfSize(fontSize float64) Font {
	rv := objc.Call[Font](fc, objc.Sel("userFontOfSize:"), fontSize)
	return rv
}

// Returns the font used by default for documents and other text under the user’s control (that is, text whose font the user can normally change), in the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1524559-userfontofsize?language=objc
func Font_UserFontOfSize(fontSize float64) Font {
	return FontClass.UserFontOfSize(fontSize)
}

// The font descriptor object for the font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1530476-fontdescriptor?language=objc
func (f_ Font) FontDescriptor() FontDescriptor {
	rv := objc.Call[FontDescriptor](f_, objc.Sel("fontDescriptor"))
	return rv
}

// The top y-coordinate, offset from the baseline, of the font’s longest ascender. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1535420-ascender?language=objc
func (f_ Font) Ascender() float64 {
	rv := objc.Call[float64](f_, objc.Sel("ascender"))
	return rv
}

// The cap height of the font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1528292-capheight?language=objc
func (f_ Font) CapHeight() float64 {
	rv := objc.Call[float64](f_, objc.Sel("capHeight"))
	return rv
}

// Returns the size of the standard system font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1531931-systemfontsize?language=objc
func (fc _FontClass) SystemFontSize() float64 {
	rv := objc.Call[float64](fc, objc.Sel("systemFontSize"))
	return rv
}

// Returns the size of the standard system font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1531931-systemfontsize?language=objc
func Font_SystemFontSize() float64 {
	return FontClass.SystemFontSize()
}

// A Boolean value indicating whether all glyphs in the font have the same advancement. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1529210-fixedpitch?language=objc
func (f_ Font) IsFixedPitch() bool {
	rv := objc.Call[bool](f_, objc.Sel("isFixedPitch"))
	return rv
}

// The character set containing all of the nominal characters that the font can render. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1535912-coveredcharacterset?language=objc
func (f_ Font) CoveredCharacterSet() foundation.CharacterSet {
	rv := objc.Call[foundation.CharacterSet](f_, objc.Sel("coveredCharacterSet"))
	return rv
}

// The current transformation matrix of the font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1526270-texttransform?language=objc
func (f_ Font) TextTransform() foundation.AffineTransform {
	rv := objc.Call[foundation.AffineTransform](f_, objc.Sel("textTransform"))
	return rv
}

// The baseline offset to use when drawing underlines with the font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1533984-underlineposition?language=objc
func (f_ Font) UnderlinePosition() float64 {
	rv := objc.Call[float64](f_, objc.Sel("underlinePosition"))
	return rv
}

// The leading value of the font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1534083-leading?language=objc
func (f_ Font) Leading() float64 {
	rv := objc.Call[float64](f_, objc.Sel("leading"))
	return rv
}

// The maximum advance of any of the font’s glyphs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1526023-maximumadvancement?language=objc
func (f_ Font) MaximumAdvancement() foundation.Size {
	rv := objc.Call[foundation.Size](f_, objc.Sel("maximumAdvancement"))
	return rv
}

// The number of glyphs in the font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1533968-numberofglyphs?language=objc
func (f_ Font) NumberOfGlyphs() uint {
	rv := objc.Call[uint](f_, objc.Sel("numberOfGlyphs"))
	return rv
}

// The x-height of the font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1533428-xheight?language=objc
func (f_ Font) XHeight() float64 {
	rv := objc.Call[float64](f_, objc.Sel("xHeight"))
	return rv
}

// Returns the size of the standard small system font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1535612-smallsystemfontsize?language=objc
func (fc _FontClass) SmallSystemFontSize() float64 {
	rv := objc.Call[float64](fc, objc.Sel("smallSystemFontSize"))
	return rv
}

// Returns the size of the standard small system font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1535612-smallsystemfontsize?language=objc
func Font_SmallSystemFontSize() float64 {
	return FontClass.SmallSystemFontSize()
}

// The point size of the font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1524511-pointsize?language=objc
func (f_ Font) PointSize() float64 {
	rv := objc.Call[float64](f_, objc.Sel("pointSize"))
	return rv
}

// A vertical version of the font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1535152-verticalfont?language=objc
func (f_ Font) VerticalFont() Font {
	rv := objc.Call[Font](f_, objc.Sel("verticalFont"))
	return rv
}

// The font’s bounding rectangle, scaled to the font’s size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1527321-boundingrectforfont?language=objc
func (f_ Font) BoundingRectForFont() foundation.Rect {
	rv := objc.Call[foundation.Rect](f_, objc.Sel("boundingRectForFont"))
	return rv
}

// The bottom y-coordinate, offset from the baseline, of the font’s longest descender. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1532270-descender?language=objc
func (f_ Font) Descender() float64 {
	rv := objc.Call[float64](f_, objc.Sel("descender"))
	return rv
}

// The full name of the font, as used in PostScript language code—for example, “Times-Roman” or “Helvetica-Oblique.” [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1526183-fontname?language=objc
func (f_ Font) FontName() string {
	rv := objc.Call[string](f_, objc.Sel("fontName"))
	return rv
}

// The transformation matrix associated with the font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1531033-matrix?language=objc
func (f_ Font) Matrix() *float64 {
	rv := objc.Call[*float64](f_, objc.Sel("matrix"))
	return rv
}

// A Boolean value indicating whether the font is a vertical font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1534644-vertical?language=objc
func (f_ Font) IsVertical() bool {
	rv := objc.Call[bool](f_, objc.Sel("isVertical"))
	return rv
}

// Returns the size of the standard label font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1534629-labelfontsize?language=objc
func (fc _FontClass) LabelFontSize() float64 {
	rv := objc.Call[float64](fc, objc.Sel("labelFontSize"))
	return rv
}

// Returns the size of the standard label font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1534629-labelfontsize?language=objc
func Font_LabelFontSize() float64 {
	return FontClass.LabelFontSize()
}

// The number of degrees that the font is slanted counterclockwise from the vertical. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1535194-italicangle?language=objc
func (f_ Font) ItalicAngle() float64 {
	rv := objc.Call[float64](f_, objc.Sel("italicAngle"))
	return rv
}

// The string encoding that works best with the font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1527635-mostcompatiblestringencoding?language=objc
func (f_ Font) MostCompatibleStringEncoding() foundation.StringEncoding {
	rv := objc.Call[foundation.StringEncoding](f_, objc.Sel("mostCompatibleStringEncoding"))
	return rv
}

// The thickness to use when drawing underlines with the font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1531229-underlinethickness?language=objc
func (f_ Font) UnderlineThickness() float64 {
	rv := objc.Call[float64](f_, objc.Sel("underlineThickness"))
	return rv
}

// The name of the font, including family and face names, to use when displaying the font information to the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1531660-displayname?language=objc
func (f_ Font) DisplayName() string {
	rv := objc.Call[string](f_, objc.Sel("displayName"))
	return rv
}

// The family name of the font—for example, “Times” or “Helvetica.” [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/1529585-familyname?language=objc
func (f_ Font) FamilyName() string {
	rv := objc.Call[string](f_, objc.Sel("familyName"))
	return rv
}
