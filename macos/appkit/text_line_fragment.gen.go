// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var TextLineFragmentClass = _TextLineFragmentClass{objc.GetClass("NSTextLineFragment")}

type _TextLineFragmentClass struct {
	objc.Class
}

type ITextLineFragment interface {
	objc.IObject
	CharacterIndexForPoint(point coregraphics.Point) int
	FractionOfDistanceThroughGlyphForPoint(point coregraphics.Point) float64
	LocationForCharacterAtIndex(index int) coregraphics.Point
	DrawAtPointInContext(point coregraphics.Point, context coregraphics.ContextRef)
	AttributedString() foundation.AttributedString
	CharacterRange() foundation.Range
	GlyphOrigin() coregraphics.Point
	TypographicBounds() coregraphics.Rect
}

type TextLineFragment struct {
	objc.Object
}

func MakeTextLineFragment(ptr unsafe.Pointer) TextLineFragment {
	return TextLineFragment{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TextLineFragment) InitWithAttributedStringRange(attributedString foundation.IAttributedString, range_ foundation.Range) TextLineFragment {
	rv := objc.CallMethod[TextLineFragment](t_, objc.GetSelector("initWithAttributedString:range:"), objc.ExtractPtr(attributedString), range_)
	return rv
}

func TextLineFragment_InitWithAttributedStringRange(attributedString foundation.IAttributedString, range_ foundation.Range) TextLineFragment {
	return TextLineFragmentClass.Alloc().InitWithAttributedStringRange(attributedString, range_)
}

func (t_ TextLineFragment) InitWithStringAttributesRange(string_ string, attributes map[foundation.AttributedStringKey]objc.IObject, range_ foundation.Range) TextLineFragment {
	rv := objc.CallMethod[TextLineFragment](t_, objc.GetSelector("initWithString:attributes:range:"), string_, attributes, range_)
	return rv
}

func TextLineFragment_InitWithStringAttributesRange(string_ string, attributes map[foundation.AttributedStringKey]objc.IObject, range_ foundation.Range) TextLineFragment {
	return TextLineFragmentClass.Alloc().InitWithStringAttributesRange(string_, attributes, range_)
}

func (tc _TextLineFragmentClass) Alloc() TextLineFragment {
	rv := objc.CallMethod[TextLineFragment](tc, objc.GetSelector("alloc"))
	return rv
}

func TextLineFragment_Alloc() TextLineFragment {
	return TextLineFragmentClass.Alloc()
}

func (tc _TextLineFragmentClass) New() TextLineFragment {
	rv := objc.CallMethod[TextLineFragment](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTextLineFragment() TextLineFragment {
	return TextLineFragmentClass.New()
}

func TextLineFragment_New() TextLineFragment {
	return TextLineFragmentClass.New()
}

func (t_ TextLineFragment) Init() TextLineFragment {
	rv := objc.CallMethod[TextLineFragment](t_, objc.GetSelector("init"))
	return rv
}

func TextLineFragment_Init() TextLineFragment {
	return TextLineFragmentClass.Alloc().Init()
}

func (t_ TextLineFragment) CharacterIndexForPoint(point coregraphics.Point) int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("characterIndexForPoint:"), point)
	return rv
}

func (t_ TextLineFragment) FractionOfDistanceThroughGlyphForPoint(point coregraphics.Point) float64 {
	rv := objc.CallMethod[float64](t_, objc.GetSelector("fractionOfDistanceThroughGlyphForPoint:"), point)
	return rv
}

func (t_ TextLineFragment) LocationForCharacterAtIndex(index int) coregraphics.Point {
	rv := objc.CallMethod[coregraphics.Point](t_, objc.GetSelector("locationForCharacterAtIndex:"), index)
	return rv
}

func (t_ TextLineFragment) DrawAtPointInContext(point coregraphics.Point, context coregraphics.ContextRef) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("drawAtPoint:inContext:"), point, context)
}

func (t_ TextLineFragment) AttributedString() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](t_, objc.GetSelector("attributedString"))
	return rv
}

func (t_ TextLineFragment) CharacterRange() foundation.Range {
	rv := objc.CallMethod[foundation.Range](t_, objc.GetSelector("characterRange"))
	return rv
}

func (t_ TextLineFragment) GlyphOrigin() coregraphics.Point {
	rv := objc.CallMethod[coregraphics.Point](t_, objc.GetSelector("glyphOrigin"))
	return rv
}

func (t_ TextLineFragment) TypographicBounds() coregraphics.Rect {
	rv := objc.CallMethod[coregraphics.Rect](t_, objc.GetSelector("typographicBounds"))
	return rv
}
