// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [GlyphGenerator] class.
var GlyphGeneratorClass = _GlyphGeneratorClass{objc.GetClass("NSGlyphGenerator")}

type _GlyphGeneratorClass struct {
	objc.Class
}

// An interface definition for the [GlyphGenerator] class.
type IGlyphGenerator interface {
	objc.IObject
	GenerateGlyphsForGlyphStorageDesiredNumberOfCharactersGlyphIndexCharacterIndex(glyphStorage PGlyphStorage, nChars uint, glyphIndex *uint, charIndex *uint)
	GenerateGlyphsForGlyphStorageObjectDesiredNumberOfCharactersGlyphIndexCharacterIndex(glyphStorageObject objc.IObject, nChars uint, glyphIndex *uint, charIndex *uint)
}

// An object that performs the initial, nominal glyph generation phase in the layout process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphgenerator?language=objc
type GlyphGenerator struct {
	objc.Object
}

func GlyphGeneratorFrom(ptr unsafe.Pointer) GlyphGenerator {
	return GlyphGenerator{
		Object: objc.ObjectFrom(ptr),
	}
}

func (gc _GlyphGeneratorClass) Alloc() GlyphGenerator {
	rv := objc.Call[GlyphGenerator](gc, objc.Sel("alloc"))
	return rv
}

func GlyphGenerator_Alloc() GlyphGenerator {
	return GlyphGeneratorClass.Alloc()
}

func (gc _GlyphGeneratorClass) New() GlyphGenerator {
	rv := objc.Call[GlyphGenerator](gc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewGlyphGenerator() GlyphGenerator {
	return GlyphGeneratorClass.New()
}

func (g_ GlyphGenerator) Init() GlyphGenerator {
	rv := objc.Call[GlyphGenerator](g_, objc.Sel("init"))
	return rv
}

// Generates glyphs for the specified glyph storage object (NSLayoutManager by default). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphgenerator/1425139-generateglyphsforglyphstorage?language=objc
func (g_ GlyphGenerator) GenerateGlyphsForGlyphStorageDesiredNumberOfCharactersGlyphIndexCharacterIndex(glyphStorage PGlyphStorage, nChars uint, glyphIndex *uint, charIndex *uint) {
	po0 := objc.WrapAsProtocol("NSGlyphStorage", glyphStorage)
	objc.Call[objc.Void](g_, objc.Sel("generateGlyphsForGlyphStorage:desiredNumberOfCharacters:glyphIndex:characterIndex:"), po0, nChars, glyphIndex, charIndex)
}

// Generates glyphs for the specified glyph storage object (NSLayoutManager by default). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphgenerator/1425139-generateglyphsforglyphstorage?language=objc
func (g_ GlyphGenerator) GenerateGlyphsForGlyphStorageObjectDesiredNumberOfCharactersGlyphIndexCharacterIndex(glyphStorageObject objc.IObject, nChars uint, glyphIndex *uint, charIndex *uint) {
	objc.Call[objc.Void](g_, objc.Sel("generateGlyphsForGlyphStorage:desiredNumberOfCharacters:glyphIndex:characterIndex:"), objc.Ptr(glyphStorageObject), nChars, glyphIndex, charIndex)
}

// Returns a shared instance of NSGlyphGenerator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphgenerator/1425155-sharedglyphgenerator?language=objc
func (gc _GlyphGeneratorClass) SharedGlyphGenerator() GlyphGenerator {
	rv := objc.Call[GlyphGenerator](gc, objc.Sel("sharedGlyphGenerator"))
	return rv
}

// Returns a shared instance of NSGlyphGenerator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphgenerator/1425155-sharedglyphgenerator?language=objc
func GlyphGenerator_SharedGlyphGenerator() GlyphGenerator {
	return GlyphGeneratorClass.SharedGlyphGenerator()
}
