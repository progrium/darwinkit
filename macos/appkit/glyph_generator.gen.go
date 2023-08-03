// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var GlyphGeneratorClass = _GlyphGeneratorClass{objc.GetClass("NSGlyphGenerator")}

type _GlyphGeneratorClass struct {
	objc.Class
}

type IGlyphGenerator interface {
	objc.IObject
	GenerateGlyphsForGlyphStorageDesiredNumberOfCharactersGlyphIndexCharacterIndex(glyphStorage IGlyphStorage, nChars uint, glyphIndex *uint, charIndex *uint)
	GenerateGlyphsForGlyphStorage0DesiredNumberOfCharactersGlyphIndexCharacterIndex(glyphStorage objc.IObject, nChars uint, glyphIndex *uint, charIndex *uint)
}

type GlyphGenerator struct {
	objc.Object
}

func MakeGlyphGenerator(ptr unsafe.Pointer) GlyphGenerator {
	return GlyphGenerator{
		Object: objc.MakeObject(ptr),
	}
}

func (gc _GlyphGeneratorClass) Alloc() GlyphGenerator {
	rv := objc.CallMethod[GlyphGenerator](gc, objc.GetSelector("alloc"))
	return rv
}

func GlyphGenerator_Alloc() GlyphGenerator {
	return GlyphGeneratorClass.Alloc()
}

func (gc _GlyphGeneratorClass) New() GlyphGenerator {
	rv := objc.CallMethod[GlyphGenerator](gc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewGlyphGenerator() GlyphGenerator {
	return GlyphGeneratorClass.New()
}

func GlyphGenerator_New() GlyphGenerator {
	return GlyphGeneratorClass.New()
}

func (g_ GlyphGenerator) Init() GlyphGenerator {
	rv := objc.CallMethod[GlyphGenerator](g_, objc.GetSelector("init"))
	return rv
}

func GlyphGenerator_Init() GlyphGenerator {
	return GlyphGeneratorClass.Alloc().Init()
}

func (g_ GlyphGenerator) GenerateGlyphsForGlyphStorageDesiredNumberOfCharactersGlyphIndexCharacterIndex(glyphStorage IGlyphStorage, nChars uint, glyphIndex *uint, charIndex *uint) {
	po := objc.WrapAsProtocol("NSGlyphStorage", glyphStorage)
	objc.CallMethod[objc.Void](g_, objc.GetSelector("generateGlyphsForGlyphStorage:desiredNumberOfCharacters:glyphIndex:characterIndex:"), po, nChars, glyphIndex, charIndex)
}

func (g_ GlyphGenerator) GenerateGlyphsForGlyphStorage0DesiredNumberOfCharactersGlyphIndexCharacterIndex(glyphStorage objc.IObject, nChars uint, glyphIndex *uint, charIndex *uint) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("generateGlyphsForGlyphStorage:desiredNumberOfCharacters:glyphIndex:characterIndex:"), objc.ExtractPtr(glyphStorage), nChars, glyphIndex, charIndex)
}

func (gc _GlyphGeneratorClass) SharedGlyphGenerator() GlyphGenerator {
	rv := objc.CallMethod[GlyphGenerator](gc, objc.GetSelector("sharedGlyphGenerator"))
	return rv
}

func GlyphGenerator_SharedGlyphGenerator() GlyphGenerator {
	return GlyphGeneratorClass.SharedGlyphGenerator()
}
