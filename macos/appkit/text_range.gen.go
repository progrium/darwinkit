// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TextRange] class.
var TextRangeClass = _TextRangeClass{objc.GetClass("NSTextRange")}

type _TextRangeClass struct {
	objc.Class
}

// An interface definition for the [TextRange] class.
type ITextRange interface {
	objc.IObject
	IsEqualToTextRange(textRange ITextRange) bool
	ContainsLocation(location PTextLocation) bool
	ContainsLocationObject(locationObject objc.IObject) bool
	ContainsRange(textRange ITextRange) bool
	IntersectsWithTextRange(textRange ITextRange) bool
	IsEmpty() bool
	Location() TextLocationWrapper
	EndLocation() TextLocationWrapper
}

// A class that represents a contiguous range between two locations inside document contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextrange?language=objc
type TextRange struct {
	objc.Object
}

func TextRangeFrom(ptr unsafe.Pointer) TextRange {
	return TextRange{
		Object: objc.ObjectFrom(ptr),
	}
}

func (t_ TextRange) TextRangeByFormingUnionWithTextRange(textRange ITextRange) TextRange {
	rv := objc.Call[TextRange](t_, objc.Sel("textRangeByFormingUnionWithTextRange:"), objc.Ptr(textRange))
	return rv
}

// Returns a new text range by forming the union with the text range you provide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextrange/3801813-textrangebyformingunionwithtextr?language=objc
func TextRange_TextRangeByFormingUnionWithTextRange(textRange ITextRange) TextRange {
	return TextRangeClass.Alloc().TextRangeByFormingUnionWithTextRange(textRange)
}

func (t_ TextRange) InitWithLocation(location PTextLocation) TextRange {
	po0 := objc.WrapAsProtocol("NSTextLocation", location)
	rv := objc.Call[TextRange](t_, objc.Sel("initWithLocation:"), po0)
	return rv
}

// Creates a new text range at the location you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextrange/3801808-initwithlocation?language=objc
func TextRange_InitWithLocation(location PTextLocation) TextRange {
	return TextRangeClass.Alloc().InitWithLocation(location)
}

func (t_ TextRange) TextRangeByIntersectingWithTextRange(textRange ITextRange) TextRange {
	rv := objc.Call[TextRange](t_, objc.Sel("textRangeByIntersectingWithTextRange:"), objc.Ptr(textRange))
	return rv
}

// Returns the range, if any, where two text ranges intersect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextrange/3801814-textrangebyintersectingwithtextr?language=objc
func TextRange_TextRangeByIntersectingWithTextRange(textRange ITextRange) TextRange {
	return TextRangeClass.Alloc().TextRangeByIntersectingWithTextRange(textRange)
}

func (tc _TextRangeClass) Alloc() TextRange {
	rv := objc.Call[TextRange](tc, objc.Sel("alloc"))
	return rv
}

func TextRange_Alloc() TextRange {
	return TextRangeClass.Alloc()
}

func (tc _TextRangeClass) New() TextRange {
	rv := objc.Call[TextRange](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTextRange() TextRange {
	return TextRangeClass.New()
}

func (t_ TextRange) Init() TextRange {
	rv := objc.Call[TextRange](t_, objc.Sel("init"))
	return rv
}

// Compares two text ranges. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextrange/3801811-isequaltotextrange?language=objc
func (t_ TextRange) IsEqualToTextRange(textRange ITextRange) bool {
	rv := objc.Call[bool](t_, objc.Sel("isEqualToTextRange:"), objc.Ptr(textRange))
	return rv
}

// Determines if the text location you specify is in the current text range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextrange/3801804-containslocation?language=objc
func (t_ TextRange) ContainsLocation(location PTextLocation) bool {
	po0 := objc.WrapAsProtocol("NSTextLocation", location)
	rv := objc.Call[bool](t_, objc.Sel("containsLocation:"), po0)
	return rv
}

// Determines if the text location you specify is in the current text range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextrange/3801804-containslocation?language=objc
func (t_ TextRange) ContainsLocationObject(locationObject objc.IObject) bool {
	rv := objc.Call[bool](t_, objc.Sel("containsLocation:"), objc.Ptr(locationObject))
	return rv
}

// Determines if the text range you specify is in the current text range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextrange/3801805-containsrange?language=objc
func (t_ TextRange) ContainsRange(textRange ITextRange) bool {
	rv := objc.Call[bool](t_, objc.Sel("containsRange:"), objc.Ptr(textRange))
	return rv
}

// Determines if two ranges intersect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextrange/3801810-intersectswithtextrange?language=objc
func (t_ TextRange) IntersectsWithTextRange(textRange ITextRange) bool {
	rv := objc.Call[bool](t_, objc.Sel("intersectsWithTextRange:"), objc.Ptr(textRange))
	return rv
}

// Returns whether the text range is empty. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextrange/3801806-empty?language=objc
func (t_ TextRange) IsEmpty() bool {
	rv := objc.Call[bool](t_, objc.Sel("isEmpty"))
	return rv
}

// Returns the starting location of the text range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextrange/3801812-location?language=objc
func (t_ TextRange) Location() TextLocationWrapper {
	rv := objc.Call[TextLocationWrapper](t_, objc.Sel("location"))
	return rv
}

// Returns the ending location of the text range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextrange/3801807-endlocation?language=objc
func (t_ TextRange) EndLocation() TextLocationWrapper {
	rv := objc.Call[TextLocationWrapper](t_, objc.Sel("endLocation"))
	return rv
}
