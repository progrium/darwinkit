// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of methods that a glyph storage object must implement to interact properly with NSGlyphGenerator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphstorage?language=objc
type PGlyphStorage interface {
	// optional
	SetIntAttributeValueForGlyphAtIndex(attributeTag int, val int, glyphIndex uint)
	HasSetIntAttributeValueForGlyphAtIndex() bool

	// optional
	AttributedString() foundation.IAttributedString
	HasAttributedString() bool

	// optional
	LayoutOptions() uint
	HasLayoutOptions() bool

	// optional
	InsertGlyphsLengthForStartingGlyphAtIndexCharacterIndex(glyphs *Glyph, length uint, glyphIndex uint, charIndex uint)
	HasInsertGlyphsLengthForStartingGlyphAtIndexCharacterIndex() bool
}

// A concrete type wrapper for the [PGlyphStorage] protocol.
type GlyphStorageWrapper struct {
	objc.Object
}

func (g_ GlyphStorageWrapper) HasSetIntAttributeValueForGlyphAtIndex() bool {
	return g_.RespondsToSelector(objc.Sel("setIntAttribute:value:forGlyphAtIndex:"))
}

// Sets a custom attribute value for a given glyph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphstorage/1425141-setintattribute?language=objc
func (g_ GlyphStorageWrapper) SetIntAttributeValueForGlyphAtIndex(attributeTag int, val int, glyphIndex uint) {
	objc.Call[objc.Void](g_, objc.Sel("setIntAttribute:value:forGlyphAtIndex:"), attributeTag, val, glyphIndex)
}

func (g_ GlyphStorageWrapper) HasAttributedString() bool {
	return g_.RespondsToSelector(objc.Sel("attributedString"))
}

// Returns the text storage object from which the NSGlyphGenerator object procures characters for glyph generation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphstorage/1425147-attributedstring?language=objc
func (g_ GlyphStorageWrapper) AttributedString() foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](g_, objc.Sel("attributedString"))
	return rv
}

func (g_ GlyphStorageWrapper) HasLayoutOptions() bool {
	return g_.RespondsToSelector(objc.Sel("layoutOptions"))
}

// Returns the current layout options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphstorage/1425149-layoutoptions?language=objc
func (g_ GlyphStorageWrapper) LayoutOptions() uint {
	rv := objc.Call[uint](g_, objc.Sel("layoutOptions"))
	return rv
}

func (g_ GlyphStorageWrapper) HasInsertGlyphsLengthForStartingGlyphAtIndexCharacterIndex() bool {
	return g_.RespondsToSelector(objc.Sel("insertGlyphs:length:forStartingGlyphAtIndex:characterIndex:"))
}

// Inserts the given glyphs into the glyph cache and maps them to the specified characters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphstorage/1425153-insertglyphs?language=objc
func (g_ GlyphStorageWrapper) InsertGlyphsLengthForStartingGlyphAtIndexCharacterIndex(glyphs *Glyph, length uint, glyphIndex uint, charIndex uint) {
	objc.Call[objc.Void](g_, objc.Sel("insertGlyphs:length:forStartingGlyphAtIndex:characterIndex:"), glyphs, length, glyphIndex, charIndex)
}
