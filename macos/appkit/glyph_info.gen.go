// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [GlyphInfo] class.
var GlyphInfoClass = _GlyphInfoClass{objc.GetClass("NSGlyphInfo")}

type _GlyphInfoClass struct {
	objc.Class
}

// An interface definition for the [GlyphInfo] class.
type IGlyphInfo interface {
	objc.IObject
	BaseString() string
	GlyphID() coregraphics.Glyph
}

// A glyph attribute in an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphinfo?language=objc
type GlyphInfo struct {
	objc.Object
}

func GlyphInfoFrom(ptr unsafe.Pointer) GlyphInfo {
	return GlyphInfo{
		Object: objc.ObjectFrom(ptr),
	}
}

func (gc _GlyphInfoClass) Alloc() GlyphInfo {
	rv := objc.Call[GlyphInfo](gc, objc.Sel("alloc"))
	return rv
}

func GlyphInfo_Alloc() GlyphInfo {
	return GlyphInfoClass.Alloc()
}

func (gc _GlyphInfoClass) New() GlyphInfo {
	rv := objc.Call[GlyphInfo](gc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewGlyphInfo() GlyphInfo {
	return GlyphInfoClass.New()
}

func (g_ GlyphInfo) Init() GlyphInfo {
	rv := objc.Call[GlyphInfo](g_, objc.Sel("init"))
	return rv
}

// Creates a glyph info object from the specified glyph identifier and font informaton. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphinfo/2887154-glyphinfowithcgglyph?language=objc
func (gc _GlyphInfoClass) GlyphInfoWithCGGlyphForFontBaseString(glyph coregraphics.Glyph, font IFont, string_ string) GlyphInfo {
	rv := objc.Call[GlyphInfo](gc, objc.Sel("glyphInfoWithCGGlyph:forFont:baseString:"), glyph, objc.Ptr(font), string_)
	return rv
}

// Creates a glyph info object from the specified glyph identifier and font informaton. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphinfo/2887154-glyphinfowithcgglyph?language=objc
func GlyphInfo_GlyphInfoWithCGGlyphForFontBaseString(glyph coregraphics.Glyph, font IFont, string_ string) GlyphInfo {
	return GlyphInfoClass.GlyphInfoWithCGGlyphForFontBaseString(glyph, font, string_)
}

// The string containing the character represented by the glyph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphinfo/2887163-basestring?language=objc
func (g_ GlyphInfo) BaseString() string {
	rv := objc.Call[string](g_, objc.Sel("baseString"))
	return rv
}

// The glyph identifier, specified as the index into the internal glyph table of the font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphinfo/2887152-glyphid?language=objc
func (g_ GlyphInfo) GlyphID() coregraphics.Glyph {
	rv := objc.Call[coregraphics.Glyph](g_, objc.Sel("glyphID"))
	return rv
}
