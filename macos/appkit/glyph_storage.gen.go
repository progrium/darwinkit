// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type IGlyphStorage interface {
	// required
	AttributedString() foundation.IAttributedString
	// required
	LayoutOptions() uint
	// required
	InsertGlyphsLengthForStartingGlyphAtIndexCharacterIndex(glyphs *Glyph, length uint, glyphIndex uint, charIndex uint)
	// required
	SetIntAttributeValueForGlyphAtIndex(attributeTag int, val int, glyphIndex uint)
}

type GlyphStorageWrapper struct {
	objc.Object
}

func (g_ GlyphStorageWrapper) AttributedString() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](g_, objc.GetSelector("attributedString"))
	return rv
}

func (g_ GlyphStorageWrapper) LayoutOptions() uint {
	rv := objc.CallMethod[uint](g_, objc.GetSelector("layoutOptions"))
	return rv
}

func (g_ GlyphStorageWrapper) InsertGlyphsLengthForStartingGlyphAtIndexCharacterIndex(glyphs *Glyph, length uint, glyphIndex uint, charIndex uint) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("insertGlyphs:length:forStartingGlyphAtIndex:characterIndex:"), glyphs, length, glyphIndex, charIndex)
}

func (g_ GlyphStorageWrapper) SetIntAttributeValueForGlyphAtIndex(attributeTag int, val int, glyphIndex uint) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setIntAttribute:value:forGlyphAtIndex:"), attributeTag, val, glyphIndex)
}
