// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Caption] class.
var CaptionClass = _CaptionClass{objc.GetClass("AVCaption")}

type _CaptionClass struct {
	objc.Class
}

// An interface definition for the [Caption] class.
type ICaption interface {
	objc.IObject
	DecorationAtIndexRange(index int, outRange *foundation.Range) CaptionDecoration
	TextCombineAtIndexRange(index int, outRange *foundation.Range) CaptionTextCombine
	TextColorAtIndexRange(index int, outRange *foundation.Range) coregraphics.ColorRef
	FontWeightAtIndexRange(index int, outRange *foundation.Range) CaptionFontWeight
	FontStyleAtIndexRange(index int, outRange *foundation.Range) CaptionFontStyle
	RubyAtIndexRange(index int, outRange *foundation.Range) CaptionRuby
	BackgroundColorAtIndexRange(index int, outRange *foundation.Range) coregraphics.ColorRef
	Animation() CaptionAnimation
	TextAlignment() CaptionTextAlignment
	TimeRange() coremedia.TimeRange
	Region() CaptionRegion
	Text() string
}

// An object that represents text to present over a time range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaption?language=objc
type Caption struct {
	objc.Object
}

func CaptionFrom(ptr unsafe.Pointer) Caption {
	return Caption{
		Object: objc.ObjectFrom(ptr),
	}
}

func (c_ Caption) InitWithTextTimeRange(text string, timeRange coremedia.TimeRange) Caption {
	rv := objc.Call[Caption](c_, objc.Sel("initWithText:timeRange:"), text, timeRange)
	return rv
}

// Creates a caption that contains text and a time range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaption/3752820-initwithtext?language=objc
func NewCaptionWithTextTimeRange(text string, timeRange coremedia.TimeRange) Caption {
	instance := CaptionClass.Alloc().InitWithTextTimeRange(text, timeRange)
	instance.Autorelease()
	return instance
}

func (cc _CaptionClass) Alloc() Caption {
	rv := objc.Call[Caption](cc, objc.Sel("alloc"))
	return rv
}

func Caption_Alloc() Caption {
	return CaptionClass.Alloc()
}

func (cc _CaptionClass) New() Caption {
	rv := objc.Call[Caption](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaption() Caption {
	return CaptionClass.New()
}

func (c_ Caption) Init() Caption {
	rv := objc.Call[Caption](c_, objc.Sel("init"))
	return rv
}

// Returns the text decoration at the index position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaption/3752817-decorationatindex?language=objc
func (c_ Caption) DecorationAtIndexRange(index int, outRange *foundation.Range) CaptionDecoration {
	rv := objc.Call[CaptionDecoration](c_, objc.Sel("decorationAtIndex:range:"), index, outRange)
	return rv
}

// Returns the text combine at the index position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaption/3752826-textcombineatindex?language=objc
func (c_ Caption) TextCombineAtIndexRange(index int, outRange *foundation.Range) CaptionTextCombine {
	rv := objc.Call[CaptionTextCombine](c_, objc.Sel("textCombineAtIndex:range:"), index, outRange)
	return rv
}

// Returns the text color at the index position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaption/3752825-textcoloratindex?language=objc
func (c_ Caption) TextColorAtIndexRange(index int, outRange *foundation.Range) coregraphics.ColorRef {
	rv := objc.Call[coregraphics.ColorRef](c_, objc.Sel("textColorAtIndex:range:"), index, outRange)
	return rv
}

// Returns the font weight and range at the index position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaption/3752819-fontweightatindex?language=objc
func (c_ Caption) FontWeightAtIndexRange(index int, outRange *foundation.Range) CaptionFontWeight {
	rv := objc.Call[CaptionFontWeight](c_, objc.Sel("fontWeightAtIndex:range:"), index, outRange)
	return rv
}

// Returns the font style and range at the index position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaption/3752818-fontstyleatindex?language=objc
func (c_ Caption) FontStyleAtIndexRange(index int, outRange *foundation.Range) CaptionFontStyle {
	rv := objc.Call[CaptionFontStyle](c_, objc.Sel("fontStyleAtIndex:range:"), index, outRange)
	return rv
}

// Returns the ruby text at the index position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaption/3752822-rubyatindex?language=objc
func (c_ Caption) RubyAtIndexRange(index int, outRange *foundation.Range) CaptionRuby {
	rv := objc.Call[CaptionRuby](c_, objc.Sel("rubyAtIndex:range:"), index, outRange)
	return rv
}

// Returns the background color at the index position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaption/3752816-backgroundcoloratindex?language=objc
func (c_ Caption) BackgroundColorAtIndexRange(index int, outRange *foundation.Range) coregraphics.ColorRef {
	rv := objc.Call[coregraphics.ColorRef](c_, objc.Sel("backgroundColorAtIndex:range:"), index, outRange)
	return rv
}

// The animation that the system applies to this caption. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaption/3752815-animation?language=objc
func (c_ Caption) Animation() CaptionAnimation {
	rv := objc.Call[CaptionAnimation](c_, objc.Sel("animation"))
	return rv
}

// The alignment for the caption text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaption/3752824-textalignment?language=objc
func (c_ Caption) TextAlignment() CaptionTextAlignment {
	rv := objc.Call[CaptionTextAlignment](c_, objc.Sel("textAlignment"))
	return rv
}

// The time range over which the system presents the caption. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaption/3752827-timerange?language=objc
func (c_ Caption) TimeRange() coremedia.TimeRange {
	rv := objc.Call[coremedia.TimeRange](c_, objc.Sel("timeRange"))
	return rv
}

// The region in which the caption exists. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaption/3752821-region?language=objc
func (c_ Caption) Region() CaptionRegion {
	rv := objc.Call[CaptionRegion](c_, objc.Sel("region"))
	return rv
}

// The caption text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaption/3752823-text?language=objc
func (c_ Caption) Text() string {
	rv := objc.Call[string](c_, objc.Sel("text"))
	return rv
}
