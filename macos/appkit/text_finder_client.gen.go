// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of methods implemented by objects that support searching using the NSTextFinder class and the in-window text find bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient?language=objc
type PTextFinderClient interface {
	// optional
	DidReplaceCharacters()
	HasDidReplaceCharacters() bool

	// optional
	ScrollRangeToVisible(range_ foundation.Range)
	HasScrollRangeToVisible() bool

	// optional
	DrawCharactersInRangeForContentView(range_ foundation.Range, view View)
	HasDrawCharactersInRangeForContentView() bool

	// optional
	StringAtIndexEffectiveRangeEndsWithSearchBoundary(characterIndex uint, outRange foundation.RangePointer, outFlag *bool) string
	HasStringAtIndexEffectiveRangeEndsWithSearchBoundary() bool

	// optional
	StringLength() uint
	HasStringLength() bool

	// optional
	ShouldReplaceCharactersInRangesWithStrings(ranges []foundation.Value, strings []string) bool
	HasShouldReplaceCharactersInRangesWithStrings() bool

	// optional
	ReplaceCharactersInRangeWithString(range_ foundation.Range, string_ string)
	HasReplaceCharactersInRangeWithString() bool

	// optional
	RectsForCharacterRange(range_ foundation.Range) []foundation.IValue
	HasRectsForCharacterRange() bool

	// optional
	ContentViewAtIndexEffectiveCharacterRange(index uint, outRange foundation.RangePointer) IView
	HasContentViewAtIndexEffectiveCharacterRange() bool

	// optional
	IsEditable() bool
	HasIsEditable() bool

	// optional
	FirstSelectedRange() foundation.Range
	HasFirstSelectedRange() bool

	// optional
	VisibleCharacterRanges() []foundation.IValue
	HasVisibleCharacterRanges() bool

	// optional
	IsSelectable() bool
	HasIsSelectable() bool

	// optional
	String() string
	HasString() bool

	// optional
	AllowsMultipleSelection() bool
	HasAllowsMultipleSelection() bool

	// optional
	SetSelectedRanges(value []foundation.Value)
	HasSetSelectedRanges() bool

	// optional
	SelectedRanges() []foundation.IValue
	HasSelectedRanges() bool
}

// A concrete type wrapper for the [PTextFinderClient] protocol.
type TextFinderClientWrapper struct {
	objc.Object
}

func (t_ TextFinderClientWrapper) HasDidReplaceCharacters() bool {
	return t_.RespondsToSelector(objc.Sel("didReplaceCharacters"))
}

// Specifies whether text characters were replaced. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/1534301-didreplacecharacters?language=objc
func (t_ TextFinderClientWrapper) DidReplaceCharacters() {
	objc.Call[objc.Void](t_, objc.Sel("didReplaceCharacters"))
}

func (t_ TextFinderClientWrapper) HasScrollRangeToVisible() bool {
	return t_.RespondsToSelector(objc.Sel("scrollRangeToVisible:"))
}

// Scrolls the specified range such that it is visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/1526989-scrollrangetovisible?language=objc
func (t_ TextFinderClientWrapper) ScrollRangeToVisible(range_ foundation.Range) {
	objc.Call[objc.Void](t_, objc.Sel("scrollRangeToVisible:"), range_)
}

func (t_ TextFinderClientWrapper) HasDrawCharactersInRangeForContentView() bool {
	return t_.RespondsToSelector(objc.Sel("drawCharactersInRange:forContentView:"))
}

// Draw the glyphs for the requested character range as they are drawn in the given content view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/1533760-drawcharactersinrange?language=objc
func (t_ TextFinderClientWrapper) DrawCharactersInRangeForContentView(range_ foundation.Range, view IView) {
	objc.Call[objc.Void](t_, objc.Sel("drawCharactersInRange:forContentView:"), range_, objc.Ptr(view))
}

func (t_ TextFinderClientWrapper) HasStringAtIndexEffectiveRangeEndsWithSearchBoundary() bool {
	return t_.RespondsToSelector(objc.Sel("stringAtIndex:effectiveRange:endsWithSearchBoundary:"))
}

// Returns the found string that is created by conceptually mapping its content to a single string, which is composed of a concatenation of all its substrings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/1529466-stringatindex?language=objc
func (t_ TextFinderClientWrapper) StringAtIndexEffectiveRangeEndsWithSearchBoundary(characterIndex uint, outRange foundation.RangePointer, outFlag *bool) string {
	rv := objc.Call[string](t_, objc.Sel("stringAtIndex:effectiveRange:endsWithSearchBoundary:"), characterIndex, outRange, outFlag)
	return rv
}

func (t_ TextFinderClientWrapper) HasStringLength() bool {
	return t_.RespondsToSelector(objc.Sel("stringLength"))
}

// Returns the full length of the conceptually concatenated string return by the stringAtIndex:effectiveRange:endsWithSearchBoundary: method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/1534333-stringlength?language=objc
func (t_ TextFinderClientWrapper) StringLength() uint {
	rv := objc.Call[uint](t_, objc.Sel("stringLength"))
	return rv
}

func (t_ TextFinderClientWrapper) HasShouldReplaceCharactersInRangesWithStrings() bool {
	return t_.RespondsToSelector(objc.Sel("shouldReplaceCharactersInRanges:withStrings:"))
}

// Returns whether the specified strings should be replaced. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/1529811-shouldreplacecharactersinranges?language=objc
func (t_ TextFinderClientWrapper) ShouldReplaceCharactersInRangesWithStrings(ranges []foundation.IValue, strings []string) bool {
	rv := objc.Call[bool](t_, objc.Sel("shouldReplaceCharactersInRanges:withStrings:"), ranges, strings)
	return rv
}

func (t_ TextFinderClientWrapper) HasReplaceCharactersInRangeWithString() bool {
	return t_.RespondsToSelector(objc.Sel("replaceCharactersInRange:withString:"))
}

// Replaces the text in the specified range with the new string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/1527702-replacecharactersinrange?language=objc
func (t_ TextFinderClientWrapper) ReplaceCharactersInRangeWithString(range_ foundation.Range, string_ string) {
	objc.Call[objc.Void](t_, objc.Sel("replaceCharactersInRange:withString:"), range_, string_)
}

func (t_ TextFinderClientWrapper) HasRectsForCharacterRange() bool {
	return t_.RespondsToSelector(objc.Sel("rectsForCharacterRange:"))
}

// An array containing the located text in the content view’s coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/1529980-rectsforcharacterrange?language=objc
func (t_ TextFinderClientWrapper) RectsForCharacterRange(range_ foundation.Range) []foundation.Value {
	rv := objc.Call[[]foundation.Value](t_, objc.Sel("rectsForCharacterRange:"), range_)
	return rv
}

func (t_ TextFinderClientWrapper) HasContentViewAtIndexEffectiveCharacterRange() bool {
	return t_.RespondsToSelector(objc.Sel("contentViewAtIndex:effectiveCharacterRange:"))
}

// Returns the view the context is displayed in. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/1524830-contentviewatindex?language=objc
func (t_ TextFinderClientWrapper) ContentViewAtIndexEffectiveCharacterRange(index uint, outRange foundation.RangePointer) View {
	rv := objc.Call[View](t_, objc.Sel("contentViewAtIndex:effectiveCharacterRange:"), index, outRange)
	return rv
}

func (t_ TextFinderClientWrapper) HasIsEditable() bool {
	return t_.RespondsToSelector(objc.Sel("isEditable"))
}

// Returns whether the text is editable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/1528515-editable?language=objc
func (t_ TextFinderClientWrapper) IsEditable() bool {
	rv := objc.Call[bool](t_, objc.Sel("isEditable"))
	return rv
}

func (t_ TextFinderClientWrapper) HasFirstSelectedRange() bool {
	return t_.RespondsToSelector(objc.Sel("firstSelectedRange"))
}

// Returns the currently selected range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/1526936-firstselectedrange?language=objc
func (t_ TextFinderClientWrapper) FirstSelectedRange() foundation.Range {
	rv := objc.Call[foundation.Range](t_, objc.Sel("firstSelectedRange"))
	return rv
}

func (t_ TextFinderClientWrapper) HasVisibleCharacterRanges() bool {
	return t_.RespondsToSelector(objc.Sel("visibleCharacterRanges"))
}

// An array of visible character ranges. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/1524834-visiblecharacterranges?language=objc
func (t_ TextFinderClientWrapper) VisibleCharacterRanges() []foundation.Value {
	rv := objc.Call[[]foundation.Value](t_, objc.Sel("visibleCharacterRanges"))
	return rv
}

func (t_ TextFinderClientWrapper) HasIsSelectable() bool {
	return t_.RespondsToSelector(objc.Sel("isSelectable"))
}

// Returns whether the text is selectable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/1533766-selectable?language=objc
func (t_ TextFinderClientWrapper) IsSelectable() bool {
	rv := objc.Call[bool](t_, objc.Sel("isSelectable"))
	return rv
}

func (t_ TextFinderClientWrapper) HasString() bool {
	return t_.RespondsToSelector(objc.Sel("string"))
}

// Allows the client to specify a single string for searching. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/1529462-string?language=objc
func (t_ TextFinderClientWrapper) String() string {
	rv := objc.Call[string](t_, objc.Sel("string"))
	return rv
}

func (t_ TextFinderClientWrapper) HasAllowsMultipleSelection() bool {
	return t_.RespondsToSelector(objc.Sel("allowsMultipleSelection"))
}

// Returns whether multiple items can be selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/1530815-allowsmultipleselection?language=objc
func (t_ TextFinderClientWrapper) AllowsMultipleSelection() bool {
	rv := objc.Call[bool](t_, objc.Sel("allowsMultipleSelection"))
	return rv
}

func (t_ TextFinderClientWrapper) HasSetSelectedRanges() bool {
	return t_.RespondsToSelector(objc.Sel("setSelectedRanges:"))
}

// Returns an array of selected ranges. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/1524696-selectedranges?language=objc
func (t_ TextFinderClientWrapper) SetSelectedRanges(value []foundation.IValue) {
	objc.Call[objc.Void](t_, objc.Sel("setSelectedRanges:"), value)
}

func (t_ TextFinderClientWrapper) HasSelectedRanges() bool {
	return t_.RespondsToSelector(objc.Sel("selectedRanges"))
}

// Returns an array of selected ranges. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/1524696-selectedranges?language=objc
func (t_ TextFinderClientWrapper) SelectedRanges() []foundation.Value {
	rv := objc.Call[[]foundation.Value](t_, objc.Sel("selectedRanges"))
	return rv
}
