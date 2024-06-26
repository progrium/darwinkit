// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// A protocol that a text field delegate can use to control its field editor action menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfielddelegate?language=objc
type PTextFieldDelegate interface {
	PControlTextEditingDelegate
	// optional
	TextFieldTextViewShouldSelectCandidateAtIndex(textField TextField, textView TextView, index uint) bool
	HasTextFieldTextViewShouldSelectCandidateAtIndex() bool

	// optional
	TextFieldTextViewCandidatesForSelectedRange_(textField TextField, textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.TextCheckingResult
	HasTextFieldTextViewCandidatesForSelectedRange_() bool

	// optional
	TextFieldTextViewCandidatesForSelectedRange(textField TextField, textView TextView, selectedRange foundation.Range) []objc.Object
	HasTextFieldTextViewCandidatesForSelectedRange() bool
}

// A delegate implementation builder for the [PTextFieldDelegate] protocol.
type TextFieldDelegate struct {
	ControlTextEditingDelegate
	_TextFieldTextViewShouldSelectCandidateAtIndex func(textField TextField, textView TextView, index uint) bool
	_TextFieldTextViewCandidatesForSelectedRange_  func(textField TextField, textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.TextCheckingResult
	_TextFieldTextViewCandidatesForSelectedRange   func(textField TextField, textView TextView, selectedRange foundation.Range) []objc.Object
}

func (di *TextFieldDelegate) HasTextFieldTextViewShouldSelectCandidateAtIndex() bool {
	return di._TextFieldTextViewShouldSelectCandidateAtIndex != nil
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfielddelegate/2539557-textfield?language=objc
func (di *TextFieldDelegate) SetTextFieldTextViewShouldSelectCandidateAtIndex(f func(textField TextField, textView TextView, index uint) bool) {
	di._TextFieldTextViewShouldSelectCandidateAtIndex = f
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfielddelegate/2539557-textfield?language=objc
func (di *TextFieldDelegate) TextFieldTextViewShouldSelectCandidateAtIndex(textField TextField, textView TextView, index uint) bool {
	return di._TextFieldTextViewShouldSelectCandidateAtIndex(textField, textView, index)
}
func (di *TextFieldDelegate) HasTextFieldTextViewCandidatesForSelectedRange_() bool {
	return di._TextFieldTextViewCandidatesForSelectedRange_ != nil
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfielddelegate/2539556-textfield?language=objc
func (di *TextFieldDelegate) SetTextFieldTextViewCandidatesForSelectedRange_(f func(textField TextField, textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.TextCheckingResult) {
	di._TextFieldTextViewCandidatesForSelectedRange_ = f
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfielddelegate/2539556-textfield?language=objc
func (di *TextFieldDelegate) TextFieldTextViewCandidatesForSelectedRange_(textField TextField, textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.TextCheckingResult {
	return di._TextFieldTextViewCandidatesForSelectedRange_(textField, textView, candidates, selectedRange)
}
func (di *TextFieldDelegate) HasTextFieldTextViewCandidatesForSelectedRange() bool {
	return di._TextFieldTextViewCandidatesForSelectedRange != nil
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfielddelegate/2539552-textfield?language=objc
func (di *TextFieldDelegate) SetTextFieldTextViewCandidatesForSelectedRange(f func(textField TextField, textView TextView, selectedRange foundation.Range) []objc.Object) {
	di._TextFieldTextViewCandidatesForSelectedRange = f
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfielddelegate/2539552-textfield?language=objc
func (di *TextFieldDelegate) TextFieldTextViewCandidatesForSelectedRange(textField TextField, textView TextView, selectedRange foundation.Range) []objc.Object {
	return di._TextFieldTextViewCandidatesForSelectedRange(textField, textView, selectedRange)
}

// ensure impl type implements protocol interface
var _ PTextFieldDelegate = (*TextFieldDelegateObject)(nil)

// A concrete type for the [PTextFieldDelegate] protocol.
type TextFieldDelegateObject struct {
	ControlTextEditingDelegateObject
}

func (t_ TextFieldDelegateObject) HasTextFieldTextViewShouldSelectCandidateAtIndex() bool {
	return t_.RespondsToSelector(objc.Sel("textField:textView:shouldSelectCandidateAtIndex:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfielddelegate/2539557-textfield?language=objc
func (t_ TextFieldDelegateObject) TextFieldTextViewShouldSelectCandidateAtIndex(textField TextField, textView TextView, index uint) bool {
	rv := objc.Call[bool](t_, objc.Sel("textField:textView:shouldSelectCandidateAtIndex:"), textField, textView, index)
	return rv
}

func (t_ TextFieldDelegateObject) HasTextFieldTextViewCandidatesForSelectedRange_() bool {
	return t_.RespondsToSelector(objc.Sel("textField:textView:candidates:forSelectedRange:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfielddelegate/2539556-textfield?language=objc
func (t_ TextFieldDelegateObject) TextFieldTextViewCandidatesForSelectedRange_(textField TextField, textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.TextCheckingResult {
	rv := objc.Call[[]foundation.TextCheckingResult](t_, objc.Sel("textField:textView:candidates:forSelectedRange:"), textField, textView, candidates, selectedRange)
	return rv
}

func (t_ TextFieldDelegateObject) HasTextFieldTextViewCandidatesForSelectedRange() bool {
	return t_.RespondsToSelector(objc.Sel("textField:textView:candidatesForSelectedRange:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfielddelegate/2539552-textfield?language=objc
func (t_ TextFieldDelegateObject) TextFieldTextViewCandidatesForSelectedRange(textField TextField, textView TextView, selectedRange foundation.Range) []objc.Object {
	rv := objc.Call[[]objc.Object](t_, objc.Sel("textField:textView:candidatesForSelectedRange:"), textField, textView, selectedRange)
	return rv
}
