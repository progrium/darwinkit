// Code generated by DarwinKit. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/coregraphics"
	"github.com/progrium/darwinkit/macos/coremedia"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [MutableCaption] class.
var MutableCaptionClass = _MutableCaptionClass{objc.GetClass("AVMutableCaption")}

type _MutableCaptionClass struct {
	objc.Class
}

// An interface definition for the [MutableCaption] class.
type IMutableCaption interface {
	ICaption
	SetBackgroundColorInRange(color coregraphics.ColorRef, range_ foundation.Range)
	SetFontWeightInRange(fontWeight CaptionFontWeight, range_ foundation.Range)
	RemoveRubyInRange(range_ foundation.Range)
	RemoveBackgroundColorInRange(range_ foundation.Range)
	RemoveFontWeightInRange(range_ foundation.Range)
	SetTextCombineInRange(textCombine CaptionTextCombine, range_ foundation.Range)
	SetTextColorInRange(color coregraphics.ColorRef, range_ foundation.Range)
	RemoveTextColorInRange(range_ foundation.Range)
	RemoveTextCombineInRange(range_ foundation.Range)
	SetFontStyleInRange(fontStyle CaptionFontStyle, range_ foundation.Range)
	RemoveFontStyleInRange(range_ foundation.Range)
	SetRubyInRange(ruby ICaptionRuby, range_ foundation.Range)
	RemoveDecorationInRange(range_ foundation.Range)
	SetDecorationInRange(decoration CaptionDecoration, range_ foundation.Range)
	SetTimeRange(value coremedia.TimeRange)
	SetRegion(value ICaptionRegion)
	SetText(value string)
	SetTextAlignment(value CaptionTextAlignment)
	SetAnimation(value CaptionAnimation)
}

// A mutable caption subclass that you use to create new captions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecaption?language=objc
type MutableCaption struct {
	Caption
}

func MutableCaptionFrom(ptr unsafe.Pointer) MutableCaption {
	return MutableCaption{
		Caption: CaptionFrom(ptr),
	}
}

func (mc _MutableCaptionClass) Alloc() MutableCaption {
	rv := objc.Call[MutableCaption](mc, objc.Sel("alloc"))
	return rv
}

func (mc _MutableCaptionClass) New() MutableCaption {
	rv := objc.Call[MutableCaption](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMutableCaption() MutableCaption {
	return MutableCaptionClass.New()
}

func (m_ MutableCaption) Init() MutableCaption {
	rv := objc.Call[MutableCaption](m_, objc.Sel("init"))
	return rv
}

func (m_ MutableCaption) InitWithTextTimeRange(text string, timeRange coremedia.TimeRange) MutableCaption {
	rv := objc.Call[MutableCaption](m_, objc.Sel("initWithText:timeRange:"), text, timeRange)
	return rv
}

// Creates a caption that contains text and a time range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaption/3752820-initwithtext?language=objc
func NewMutableCaptionWithTextTimeRange(text string, timeRange coremedia.TimeRange) MutableCaption {
	instance := MutableCaptionClass.Alloc().InitWithTextTimeRange(text, timeRange)
	instance.Autorelease()
	return instance
}

// Sets the background color for a range of text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecaption/3752908-setbackgroundcolor?language=objc
func (m_ MutableCaption) SetBackgroundColorInRange(color coregraphics.ColorRef, range_ foundation.Range) {
	objc.Call[objc.Void](m_, objc.Sel("setBackgroundColor:inRange:"), color, range_)
}

// Sets the font weight for a range of text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecaption/3752911-setfontweight?language=objc
func (m_ MutableCaption) SetFontWeightInRange(fontWeight CaptionFontWeight, range_ foundation.Range) {
	objc.Call[objc.Void](m_, objc.Sel("setFontWeight:inRange:"), fontWeight, range_)
}

// Removes ruby text from a range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecaption/3752905-removerubyinrange?language=objc
func (m_ MutableCaption) RemoveRubyInRange(range_ foundation.Range) {
	objc.Call[objc.Void](m_, objc.Sel("removeRubyInRange:"), range_)
}

// Removes a background color from a range of text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecaption/3752901-removebackgroundcolorinrange?language=objc
func (m_ MutableCaption) RemoveBackgroundColorInRange(range_ foundation.Range) {
	objc.Call[objc.Void](m_, objc.Sel("removeBackgroundColorInRange:"), range_)
}

// Removes a font weight from a range of text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecaption/3752904-removefontweightinrange?language=objc
func (m_ MutableCaption) RemoveFontWeightInRange(range_ foundation.Range) {
	objc.Call[objc.Void](m_, objc.Sel("removeFontWeightInRange:"), range_)
}

// Sets text combine for a range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecaption/3752914-settextcombine?language=objc
func (m_ MutableCaption) SetTextCombineInRange(textCombine CaptionTextCombine, range_ foundation.Range) {
	objc.Call[objc.Void](m_, objc.Sel("setTextCombine:inRange:"), textCombine, range_)
}

// Sets the text color for a range of text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecaption/3752913-settextcolor?language=objc
func (m_ MutableCaption) SetTextColorInRange(color coregraphics.ColorRef, range_ foundation.Range) {
	objc.Call[objc.Void](m_, objc.Sel("setTextColor:inRange:"), color, range_)
}

// Removes the text color for a range of text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecaption/3752906-removetextcolorinrange?language=objc
func (m_ MutableCaption) RemoveTextColorInRange(range_ foundation.Range) {
	objc.Call[objc.Void](m_, objc.Sel("removeTextColorInRange:"), range_)
}

// Removes text combine from a range of text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecaption/3752907-removetextcombineinrange?language=objc
func (m_ MutableCaption) RemoveTextCombineInRange(range_ foundation.Range) {
	objc.Call[objc.Void](m_, objc.Sel("removeTextCombineInRange:"), range_)
}

// Sets the font style for a range of text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecaption/3752910-setfontstyle?language=objc
func (m_ MutableCaption) SetFontStyleInRange(fontStyle CaptionFontStyle, range_ foundation.Range) {
	objc.Call[objc.Void](m_, objc.Sel("setFontStyle:inRange:"), fontStyle, range_)
}

// Removes a font style from a range of text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecaption/3752903-removefontstyleinrange?language=objc
func (m_ MutableCaption) RemoveFontStyleInRange(range_ foundation.Range) {
	objc.Call[objc.Void](m_, objc.Sel("removeFontStyleInRange:"), range_)
}

// Sets ruby text for a range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecaption/3752912-setruby?language=objc
func (m_ MutableCaption) SetRubyInRange(ruby ICaptionRuby, range_ foundation.Range) {
	objc.Call[objc.Void](m_, objc.Sel("setRuby:inRange:"), ruby, range_)
}

// Removes a decoration from a range of text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecaption/3752902-removedecorationinrange?language=objc
func (m_ MutableCaption) RemoveDecorationInRange(range_ foundation.Range) {
	objc.Call[objc.Void](m_, objc.Sel("removeDecorationInRange:"), range_)
}

// Sets a decoration for a range of text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecaption/3752909-setdecoration?language=objc
func (m_ MutableCaption) SetDecorationInRange(decoration CaptionDecoration, range_ foundation.Range) {
	objc.Call[objc.Void](m_, objc.Sel("setDecoration:inRange:"), decoration, range_)
}

// The time range over which the system presents the caption. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecaption/3752917-timerange?language=objc
func (m_ MutableCaption) SetTimeRange(value coremedia.TimeRange) {
	objc.Call[objc.Void](m_, objc.Sel("setTimeRange:"), value)
}

// The region in which the caption exists. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecaption/3752900-region?language=objc
func (m_ MutableCaption) SetRegion(value ICaptionRegion) {
	objc.Call[objc.Void](m_, objc.Sel("setRegion:"), value)
}

// The caption text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecaption/3752915-text?language=objc
func (m_ MutableCaption) SetText(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setText:"), value)
}

// The alignment of the caption text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecaption/3752916-textalignment?language=objc
func (m_ MutableCaption) SetTextAlignment(value CaptionTextAlignment) {
	objc.Call[objc.Void](m_, objc.Sel("setTextAlignment:"), value)
}

// Animations to apply to the caption text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecaption/3752899-animation?language=objc
func (m_ MutableCaption) SetAnimation(value CaptionAnimation) {
	objc.Call[objc.Void](m_, objc.Sel("setAnimation:"), value)
}
