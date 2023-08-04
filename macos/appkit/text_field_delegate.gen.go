// auto-generated code, do not modify
package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type ITextFieldDelegate interface {
	IControlTextEditingDelegate
	ImplementsTextFieldTextViewCandidatesForSelectedRange_() bool
	// optional
	TextFieldTextViewCandidatesForSelectedRange_(textField TextField, textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.ITextCheckingResult
	ImplementsTextFieldTextViewCandidatesForSelectedRange() bool
	// optional
	TextFieldTextViewCandidatesForSelectedRange(textField TextField, textView TextView, selectedRange foundation.Range) []objc.IObject
	ImplementsTextFieldTextViewShouldSelectCandidateAtIndex() bool
	// optional
	TextFieldTextViewShouldSelectCandidateAtIndex(textField TextField, textView TextView, index uint) bool
}

type TextFieldDelegate struct {
	ControlTextEditingDelegate
	_TextFieldTextViewCandidatesForSelectedRange_  func(textField TextField, textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.ITextCheckingResult
	_TextFieldTextViewCandidatesForSelectedRange   func(textField TextField, textView TextView, selectedRange foundation.Range) []objc.IObject
	_TextFieldTextViewShouldSelectCandidateAtIndex func(textField TextField, textView TextView, index uint) bool
}

func (di *TextFieldDelegate) ImplementsTextFieldTextViewCandidatesForSelectedRange_() bool {
	return di._TextFieldTextViewCandidatesForSelectedRange_ != nil
}

func (di *TextFieldDelegate) SetTextFieldTextViewCandidatesForSelectedRange_(f func(textField TextField, textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.ITextCheckingResult) {
	di._TextFieldTextViewCandidatesForSelectedRange_ = f
}

func (di *TextFieldDelegate) TextFieldTextViewCandidatesForSelectedRange_(textField TextField, textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.ITextCheckingResult {
	return di._TextFieldTextViewCandidatesForSelectedRange_(textField, textView, candidates, selectedRange)
}
func (di *TextFieldDelegate) ImplementsTextFieldTextViewCandidatesForSelectedRange() bool {
	return di._TextFieldTextViewCandidatesForSelectedRange != nil
}

func (di *TextFieldDelegate) SetTextFieldTextViewCandidatesForSelectedRange(f func(textField TextField, textView TextView, selectedRange foundation.Range) []objc.IObject) {
	di._TextFieldTextViewCandidatesForSelectedRange = f
}

func (di *TextFieldDelegate) TextFieldTextViewCandidatesForSelectedRange(textField TextField, textView TextView, selectedRange foundation.Range) []objc.IObject {
	return di._TextFieldTextViewCandidatesForSelectedRange(textField, textView, selectedRange)
}
func (di *TextFieldDelegate) ImplementsTextFieldTextViewShouldSelectCandidateAtIndex() bool {
	return di._TextFieldTextViewShouldSelectCandidateAtIndex != nil
}

func (di *TextFieldDelegate) SetTextFieldTextViewShouldSelectCandidateAtIndex(f func(textField TextField, textView TextView, index uint) bool) {
	di._TextFieldTextViewShouldSelectCandidateAtIndex = f
}

func (di *TextFieldDelegate) TextFieldTextViewShouldSelectCandidateAtIndex(textField TextField, textView TextView, index uint) bool {
	return di._TextFieldTextViewShouldSelectCandidateAtIndex(textField, textView, index)
}

type TextFieldDelegateWrapper struct {
	ControlTextEditingDelegateWrapper
}

func (t_ TextFieldDelegateWrapper) ImplementsTextFieldTextViewCandidatesForSelectedRange_() bool {
	return t_.RespondsToSelector(objc.GetSelector("textField:textView:candidates:forSelectedRange:"))
}

func (t_ TextFieldDelegateWrapper) TextFieldTextViewCandidatesForSelectedRange_(textField ITextField, textView ITextView, candidates []foundation.ITextCheckingResult, selectedRange foundation.Range) []foundation.TextCheckingResult {
	rv := objc.CallMethod[[]foundation.TextCheckingResult](t_, objc.GetSelector("textField:textView:candidates:forSelectedRange:"), objc.ExtractPtr(textField), objc.ExtractPtr(textView), candidates, selectedRange)
	return rv
}

func (t_ TextFieldDelegateWrapper) ImplementsTextFieldTextViewCandidatesForSelectedRange() bool {
	return t_.RespondsToSelector(objc.GetSelector("textField:textView:candidatesForSelectedRange:"))
}

func (t_ TextFieldDelegateWrapper) TextFieldTextViewCandidatesForSelectedRange(textField ITextField, textView ITextView, selectedRange foundation.Range) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](t_, objc.GetSelector("textField:textView:candidatesForSelectedRange:"), objc.ExtractPtr(textField), objc.ExtractPtr(textView), selectedRange)
	return rv
}

func (t_ TextFieldDelegateWrapper) ImplementsTextFieldTextViewShouldSelectCandidateAtIndex() bool {
	return t_.RespondsToSelector(objc.GetSelector("textField:textView:shouldSelectCandidateAtIndex:"))
}

func (t_ TextFieldDelegateWrapper) TextFieldTextViewShouldSelectCandidateAtIndex(textField ITextField, textView ITextView, index uint) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("textField:textView:shouldSelectCandidateAtIndex:"), objc.ExtractPtr(textField), objc.ExtractPtr(textView), index)
	return rv
}
