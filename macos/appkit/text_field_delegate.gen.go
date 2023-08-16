// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that a text field delegate can use to control its field editor action menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfielddelegate?language=objc
type PTextFieldDelegate interface {
	PControlTextEditingDelegate
	// optional
	TextFieldTextViewShouldSelectCandidateAtIndex(textField TextField, textView TextView, index uint) bool
	HasTextFieldTextViewShouldSelectCandidateAtIndex() bool
}

// A delegate implementation builder for the [PTextFieldDelegate] protocol.
type TextFieldDelegate struct {
	ControlTextEditingDelegate
	_TextFieldTextViewShouldSelectCandidateAtIndex func(textField TextField, textView TextView, index uint) bool
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

// A concrete type wrapper for the [PTextFieldDelegate] protocol.
type TextFieldDelegateWrapper struct {
	ControlTextEditingDelegateWrapper
}

func (t_ TextFieldDelegateWrapper) HasTextFieldTextViewShouldSelectCandidateAtIndex() bool {
	return t_.RespondsToSelector(objc.Sel("textField:textView:shouldSelectCandidateAtIndex:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfielddelegate/2539557-textfield?language=objc
func (t_ TextFieldDelegateWrapper) TextFieldTextViewShouldSelectCandidateAtIndex(textField ITextField, textView ITextView, index uint) bool {
	rv := objc.Call[bool](t_, objc.Sel("textField:textView:shouldSelectCandidateAtIndex:"), objc.Ptr(textField), objc.Ptr(textView), index)
	return rv
}
