// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of methods that text views need to implement to interact properly with the text input management system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputclient?language=objc
type PTextInputClient interface {
	// optional
	InsertTextReplacementRange(string_ objc.Object, replacementRange foundation.Range)
	HasInsertTextReplacementRange() bool

	// optional
	HasMarkedText() bool
	HasHasMarkedText() bool

	// optional
	DrawsVerticallyForCharacterAtIndex(charIndex uint) bool
	HasDrawsVerticallyForCharacterAtIndex() bool

	// optional
	AttributedString() foundation.IAttributedString
	HasAttributedString() bool

	// optional
	CharacterIndexForPoint(point foundation.Point) uint
	HasCharacterIndexForPoint() bool

	// optional
	UnmarkText()
	HasUnmarkText() bool

	// optional
	SelectedRange() foundation.Range
	HasSelectedRange() bool

	// optional
	MarkedRange() foundation.Range
	HasMarkedRange() bool

	// optional
	SetMarkedTextSelectedRangeReplacementRange(string_ objc.Object, selectedRange foundation.Range, replacementRange foundation.Range)
	HasSetMarkedTextSelectedRangeReplacementRange() bool

	// optional
	WindowLevel() int
	HasWindowLevel() bool

	// optional
	BaselineDeltaForCharacterAtIndex(anIndex uint) float64
	HasBaselineDeltaForCharacterAtIndex() bool

	// optional
	AttributedSubstringForProposedRangeActualRange(range_ foundation.Range, actualRange foundation.RangePointer) foundation.IAttributedString
	HasAttributedSubstringForProposedRangeActualRange() bool

	// optional
	FirstRectForCharacterRangeActualRange(range_ foundation.Range, actualRange foundation.RangePointer) foundation.Rect
	HasFirstRectForCharacterRangeActualRange() bool

	// optional
	DoCommandBySelector(selector objc.Selector)
	HasDoCommandBySelector() bool

	// optional
	FractionOfDistanceThroughGlyphForPoint(point foundation.Point) float64
	HasFractionOfDistanceThroughGlyphForPoint() bool

	// optional
	ValidAttributesForMarkedText() []foundation.AttributedStringKey
	HasValidAttributesForMarkedText() bool
}

// A concrete type wrapper for the [PTextInputClient] protocol.
type TextInputClientWrapper struct {
	objc.Object
}

func (t_ TextInputClientWrapper) HasInsertTextReplacementRange() bool {
	return t_.RespondsToSelector(objc.Sel("insertText:replacementRange:"))
}

// Inserts the given string into the receiver, replacing the specified content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputclient/1438258-inserttext?language=objc
func (t_ TextInputClientWrapper) InsertTextReplacementRange(string_ objc.IObject, replacementRange foundation.Range) {
	objc.Call[objc.Void](t_, objc.Sel("insertText:replacementRange:"), string_, replacementRange)
}

func (t_ TextInputClientWrapper) HasHasMarkedText() bool {
	return t_.RespondsToSelector(objc.Sel("hasMarkedText"))
}

// Returns a Boolean value indicating whether the receiver has marked text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputclient/1438234-hasmarkedtext?language=objc
func (t_ TextInputClientWrapper) HasMarkedText() bool {
	rv := objc.Call[bool](t_, objc.Sel("hasMarkedText"))
	return rv
}

func (t_ TextInputClientWrapper) HasDrawsVerticallyForCharacterAtIndex() bool {
	return t_.RespondsToSelector(objc.Sel("drawsVerticallyForCharacterAtIndex:"))
}

// Informs the text input management system whether the protocol-conforming client renders the character at the given index vertically. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputclient/1438252-drawsverticallyforcharacteratind?language=objc
func (t_ TextInputClientWrapper) DrawsVerticallyForCharacterAtIndex(charIndex uint) bool {
	rv := objc.Call[bool](t_, objc.Sel("drawsVerticallyForCharacterAtIndex:"), charIndex)
	return rv
}

func (t_ TextInputClientWrapper) HasAttributedString() bool {
	return t_.RespondsToSelector(objc.Sel("attributedString"))
}

// Returns an attributed string representing the receiver's text storage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputclient/1438232-attributedstring?language=objc
func (t_ TextInputClientWrapper) AttributedString() foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](t_, objc.Sel("attributedString"))
	return rv
}

func (t_ TextInputClientWrapper) HasCharacterIndexForPoint() bool {
	return t_.RespondsToSelector(objc.Sel("characterIndexForPoint:"))
}

// Returns the index of the character whose bounding rectangle includes the given point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputclient/1438244-characterindexforpoint?language=objc
func (t_ TextInputClientWrapper) CharacterIndexForPoint(point foundation.Point) uint {
	rv := objc.Call[uint](t_, objc.Sel("characterIndexForPoint:"), point)
	return rv
}

func (t_ TextInputClientWrapper) HasUnmarkText() bool {
	return t_.RespondsToSelector(objc.Sel("unmarkText"))
}

// Unmarks the marked text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputclient/1438239-unmarktext?language=objc
func (t_ TextInputClientWrapper) UnmarkText() {
	objc.Call[objc.Void](t_, objc.Sel("unmarkText"))
}

func (t_ TextInputClientWrapper) HasSelectedRange() bool {
	return t_.RespondsToSelector(objc.Sel("selectedRange"))
}

// Returns the range of selected text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputclient/1438242-selectedrange?language=objc
func (t_ TextInputClientWrapper) SelectedRange() foundation.Range {
	rv := objc.Call[foundation.Range](t_, objc.Sel("selectedRange"))
	return rv
}

func (t_ TextInputClientWrapper) HasMarkedRange() bool {
	return t_.RespondsToSelector(objc.Sel("markedRange"))
}

// Returns the range of the marked text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputclient/1438250-markedrange?language=objc
func (t_ TextInputClientWrapper) MarkedRange() foundation.Range {
	rv := objc.Call[foundation.Range](t_, objc.Sel("markedRange"))
	return rv
}

func (t_ TextInputClientWrapper) HasSetMarkedTextSelectedRangeReplacementRange() bool {
	return t_.RespondsToSelector(objc.Sel("setMarkedText:selectedRange:replacementRange:"))
}

// Replaces a specified range in the receiver’s text storage with the given string and sets the selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputclient/1438246-setmarkedtext?language=objc
func (t_ TextInputClientWrapper) SetMarkedTextSelectedRangeReplacementRange(string_ objc.IObject, selectedRange foundation.Range, replacementRange foundation.Range) {
	objc.Call[objc.Void](t_, objc.Sel("setMarkedText:selectedRange:replacementRange:"), string_, selectedRange, replacementRange)
}

func (t_ TextInputClientWrapper) HasWindowLevel() bool {
	return t_.RespondsToSelector(objc.Sel("windowLevel"))
}

// Returns the window level of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputclient/1438248-windowlevel?language=objc
func (t_ TextInputClientWrapper) WindowLevel() int {
	rv := objc.Call[int](t_, objc.Sel("windowLevel"))
	return rv
}

func (t_ TextInputClientWrapper) HasBaselineDeltaForCharacterAtIndex() bool {
	return t_.RespondsToSelector(objc.Sel("baselineDeltaForCharacterAtIndex:"))
}

// Returns the baseline position of a given character relative to the origin of rectangle returned by firstRectForCharacterRange:actualRange:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputclient/1438254-baselinedeltaforcharacteratindex?language=objc
func (t_ TextInputClientWrapper) BaselineDeltaForCharacterAtIndex(anIndex uint) float64 {
	rv := objc.Call[float64](t_, objc.Sel("baselineDeltaForCharacterAtIndex:"), anIndex)
	return rv
}

func (t_ TextInputClientWrapper) HasAttributedSubstringForProposedRangeActualRange() bool {
	return t_.RespondsToSelector(objc.Sel("attributedSubstringForProposedRange:actualRange:"))
}

// Returns an attributed string derived from the given range in the receiver's text storage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputclient/1438238-attributedsubstringforproposedra?language=objc
func (t_ TextInputClientWrapper) AttributedSubstringForProposedRangeActualRange(range_ foundation.Range, actualRange foundation.RangePointer) foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](t_, objc.Sel("attributedSubstringForProposedRange:actualRange:"), range_, actualRange)
	return rv
}

func (t_ TextInputClientWrapper) HasFirstRectForCharacterRangeActualRange() bool {
	return t_.RespondsToSelector(objc.Sel("firstRectForCharacterRange:actualRange:"))
}

// Returns the first logical boundary rectangle for characters in the given range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputclient/1438240-firstrectforcharacterrange?language=objc
func (t_ TextInputClientWrapper) FirstRectForCharacterRangeActualRange(range_ foundation.Range, actualRange foundation.RangePointer) foundation.Rect {
	rv := objc.Call[foundation.Rect](t_, objc.Sel("firstRectForCharacterRange:actualRange:"), range_, actualRange)
	return rv
}

func (t_ TextInputClientWrapper) HasDoCommandBySelector() bool {
	return t_.RespondsToSelector(objc.Sel("doCommandBySelector:"))
}

// Invokes the action specified by the given selector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputclient/1438256-docommandbyselector?language=objc
func (t_ TextInputClientWrapper) DoCommandBySelector(selector objc.Selector) {
	objc.Call[objc.Void](t_, objc.Sel("doCommandBySelector:"), selector)
}

func (t_ TextInputClientWrapper) HasFractionOfDistanceThroughGlyphForPoint() bool {
	return t_.RespondsToSelector(objc.Sel("fractionOfDistanceThroughGlyphForPoint:"))
}

// Returns the fraction of the distance from the left side of the character to the right side that a given point lies. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputclient/1438236-fractionofdistancethroughglyphfo?language=objc
func (t_ TextInputClientWrapper) FractionOfDistanceThroughGlyphForPoint(point foundation.Point) float64 {
	rv := objc.Call[float64](t_, objc.Sel("fractionOfDistanceThroughGlyphForPoint:"), point)
	return rv
}

func (t_ TextInputClientWrapper) HasValidAttributesForMarkedText() bool {
	return t_.RespondsToSelector(objc.Sel("validAttributesForMarkedText"))
}

// Returns an array of attribute names recognized by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputclient/1438228-validattributesformarkedtext?language=objc
func (t_ TextInputClientWrapper) ValidAttributesForMarkedText() []foundation.AttributedStringKey {
	rv := objc.Call[[]foundation.AttributedStringKey](t_, objc.Sel("validAttributesForMarkedText"))
	return rv
}
