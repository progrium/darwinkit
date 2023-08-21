// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TextLineFragment] class.
var TextLineFragmentClass = _TextLineFragmentClass{objc.GetClass("NSTextLineFragment")}

type _TextLineFragmentClass struct {
	objc.Class
}

// An interface definition for the [TextLineFragment] class.
type ITextLineFragment interface {
	objc.IObject
	DrawAtPointInContext(point coregraphics.Point, context coregraphics.ContextRef)
	CharacterIndexForPoint(point coregraphics.Point) int
	LocationForCharacterAtIndex(index int) coregraphics.Point
	FractionOfDistanceThroughGlyphForPoint(point coregraphics.Point) float64
	TypographicBounds() coregraphics.Rect
	CharacterRange() foundation.Range
	AttributedString() foundation.AttributedString
	GlyphOrigin() coregraphics.Point
}

// A class that represents a line fragment as a single textual layout and rendering unit inside a text layout fragment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlinefragment?language=objc
type TextLineFragment struct {
	objc.Object
}

func TextLineFragmentFrom(ptr unsafe.Pointer) TextLineFragment {
	return TextLineFragment{
		Object: objc.ObjectFrom(ptr),
	}
}

func (t_ TextLineFragment) InitWithAttributedStringRange(attributedString foundation.IAttributedString, range_ foundation.Range) TextLineFragment {
	rv := objc.Call[TextLineFragment](t_, objc.Sel("initWithAttributedString:range:"), objc.Ptr(attributedString), range_)
	return rv
}

// Creates a new line fragment from the attributed string for the range of characters you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlinefragment/3810042-initwithattributedstring?language=objc
func NewTextLineFragmentWithAttributedStringRange(attributedString foundation.IAttributedString, range_ foundation.Range) TextLineFragment {
	instance := TextLineFragmentClass.Alloc().InitWithAttributedStringRange(attributedString, range_)
	instance.Autorelease()
	return instance
}

func (t_ TextLineFragment) InitWithStringAttributesRange(string_ string, attributes map[foundation.AttributedStringKey]objc.IObject, range_ foundation.Range) TextLineFragment {
	rv := objc.Call[TextLineFragment](t_, objc.Sel("initWithString:attributes:range:"), string_, attributes, range_)
	return rv
}

// Creates a new line fragment using the string, attributes, and range you provide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlinefragment/3810044-initwithstring?language=objc
func NewTextLineFragmentWithStringAttributesRange(string_ string, attributes map[foundation.AttributedStringKey]objc.IObject, range_ foundation.Range) TextLineFragment {
	instance := TextLineFragmentClass.Alloc().InitWithStringAttributesRange(string_, attributes, range_)
	instance.Autorelease()
	return instance
}

func (tc _TextLineFragmentClass) Alloc() TextLineFragment {
	rv := objc.Call[TextLineFragment](tc, objc.Sel("alloc"))
	return rv
}

func TextLineFragment_Alloc() TextLineFragment {
	return TextLineFragmentClass.Alloc()
}

func (tc _TextLineFragmentClass) New() TextLineFragment {
	rv := objc.Call[TextLineFragment](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTextLineFragment() TextLineFragment {
	return TextLineFragmentClass.New()
}

func (t_ TextLineFragment) Init() TextLineFragment {
	rv := objc.Call[TextLineFragment](t_, objc.Sel("init"))
	return rv
}

// Renders the line fragment contents at the rendering origin. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlinefragment/3824754-drawatpoint?language=objc
func (t_ TextLineFragment) DrawAtPointInContext(point coregraphics.Point, context coregraphics.ContextRef) {
	objc.Call[objc.Void](t_, objc.Sel("drawAtPoint:inContext:"), point, context)
}

// Returns character index for a point inside the line fragment coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlinefragment/3852578-characterindexforpoint?language=objc
func (t_ TextLineFragment) CharacterIndexForPoint(point coregraphics.Point) int {
	rv := objc.Call[int](t_, objc.Sel("characterIndexForPoint:"), point)
	return rv
}

// Returns the location of the character at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlinefragment/3810045-locationforcharacteratindex?language=objc
func (t_ TextLineFragment) LocationForCharacterAtIndex(index int) coregraphics.Point {
	rv := objc.Call[coregraphics.Point](t_, objc.Sel("locationForCharacterAtIndex:"), index)
	return rv
}

// Returns character index for a point inside the line fragment coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlinefragment/3852579-fractionofdistancethroughglyphfo?language=objc
func (t_ TextLineFragment) FractionOfDistanceThroughGlyphForPoint(point coregraphics.Point) float64 {
	rv := objc.Call[float64](t_, objc.Sel("fractionOfDistanceThroughGlyphForPoint:"), point)
	return rv
}

// The typographic bounds that specifies the dimensions of the line fragment for laying out line fragments to each other. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlinefragment/3810046-typographicbounds?language=objc
func (t_ TextLineFragment) TypographicBounds() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](t_, objc.Sel("typographicBounds"))
	return rv
}

// The string range for the source attributed string that corresponds to this line fragment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlinefragment/3810039-characterrange?language=objc
func (t_ TextLineFragment) CharacterRange() foundation.Range {
	rv := objc.Call[foundation.Range](t_, objc.Sel("characterRange"))
	return rv
}

// The source attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlinefragment/3810037-attributedstring?language=objc
func (t_ TextLineFragment) AttributedString() foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](t_, objc.Sel("attributedString"))
	return rv
}

// Rendering origin for the left-most glyph in the line fragment coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlinefragment/3810041-glyphorigin?language=objc
func (t_ TextLineFragment) GlyphOrigin() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](t_, objc.Sel("glyphOrigin"))
	return rv
}
