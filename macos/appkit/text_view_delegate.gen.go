// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of optional methods that text view delegates can use to manage selection, set text attributes, work with the spell checker, and more. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextviewdelegate?language=objc
type PTextViewDelegate interface {
	PTextDelegate
	// optional
	TextViewDidChangeTypingAttributes(notification foundation.Notification)
	HasTextViewDidChangeTypingAttributes() bool

	// optional
	UndoManagerForTextView(view TextView) foundation.IUndoManager
	HasUndoManagerForTextView() bool

	// optional
	TextViewCandidatesForSelectedRange(textView TextView, selectedRange foundation.Range) []objc.IObject
	HasTextViewCandidatesForSelectedRange() bool

	// optional
	TextViewDidChangeSelection(notification foundation.Notification)
	HasTextViewDidChangeSelection() bool
}

// A delegate implementation builder for the [PTextViewDelegate] protocol.
type TextViewDelegate struct {
	TextDelegate
	_TextViewDidChangeTypingAttributes  func(notification foundation.Notification)
	_UndoManagerForTextView             func(view TextView) foundation.IUndoManager
	_TextViewCandidatesForSelectedRange func(textView TextView, selectedRange foundation.Range) []objc.IObject
	_TextViewDidChangeSelection         func(notification foundation.Notification)
}

func (di *TextViewDelegate) HasTextViewDidChangeTypingAttributes() bool {
	return di._TextViewDidChangeTypingAttributes != nil
}

// Sent when a text view’s typing attributes change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextviewdelegate/1449303-textviewdidchangetypingattribute?language=objc
func (di *TextViewDelegate) SetTextViewDidChangeTypingAttributes(f func(notification foundation.Notification)) {
	di._TextViewDidChangeTypingAttributes = f
}

// Sent when a text view’s typing attributes change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextviewdelegate/1449303-textviewdidchangetypingattribute?language=objc
func (di *TextViewDelegate) TextViewDidChangeTypingAttributes(notification foundation.Notification) {
	di._TextViewDidChangeTypingAttributes(notification)
}
func (di *TextViewDelegate) HasUndoManagerForTextView() bool {
	return di._UndoManagerForTextView != nil
}

// Returns the undo manager for the specified text view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextviewdelegate/1449225-undomanagerfortextview?language=objc
func (di *TextViewDelegate) SetUndoManagerForTextView(f func(view TextView) foundation.IUndoManager) {
	di._UndoManagerForTextView = f
}

// Returns the undo manager for the specified text view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextviewdelegate/1449225-undomanagerfortextview?language=objc
func (di *TextViewDelegate) UndoManagerForTextView(view TextView) foundation.IUndoManager {
	return di._UndoManagerForTextView(view)
}
func (di *TextViewDelegate) HasTextViewCandidatesForSelectedRange() bool {
	return di._TextViewCandidatesForSelectedRange != nil
}

// Returns an array of objects that represent the elements of a selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextviewdelegate/2544744-textview?language=objc
func (di *TextViewDelegate) SetTextViewCandidatesForSelectedRange(f func(textView TextView, selectedRange foundation.Range) []objc.IObject) {
	di._TextViewCandidatesForSelectedRange = f
}

// Returns an array of objects that represent the elements of a selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextviewdelegate/2544744-textview?language=objc
func (di *TextViewDelegate) TextViewCandidatesForSelectedRange(textView TextView, selectedRange foundation.Range) []objc.IObject {
	return di._TextViewCandidatesForSelectedRange(textView, selectedRange)
}
func (di *TextViewDelegate) HasTextViewDidChangeSelection() bool {
	return di._TextViewDidChangeSelection != nil
}

// Sent when the selection changes in the text view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextviewdelegate/1449363-textviewdidchangeselection?language=objc
func (di *TextViewDelegate) SetTextViewDidChangeSelection(f func(notification foundation.Notification)) {
	di._TextViewDidChangeSelection = f
}

// Sent when the selection changes in the text view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextviewdelegate/1449363-textviewdidchangeselection?language=objc
func (di *TextViewDelegate) TextViewDidChangeSelection(notification foundation.Notification) {
	di._TextViewDidChangeSelection(notification)
}

// A concrete type wrapper for the [PTextViewDelegate] protocol.
type TextViewDelegateWrapper struct {
	TextDelegateWrapper
}

func (t_ TextViewDelegateWrapper) HasTextViewDidChangeTypingAttributes() bool {
	return t_.RespondsToSelector(objc.Sel("textViewDidChangeTypingAttributes:"))
}

// Sent when a text view’s typing attributes change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextviewdelegate/1449303-textviewdidchangetypingattribute?language=objc
func (t_ TextViewDelegateWrapper) TextViewDidChangeTypingAttributes(notification foundation.INotification) {
	objc.Call[objc.Void](t_, objc.Sel("textViewDidChangeTypingAttributes:"), objc.Ptr(notification))
}

func (t_ TextViewDelegateWrapper) HasUndoManagerForTextView() bool {
	return t_.RespondsToSelector(objc.Sel("undoManagerForTextView:"))
}

// Returns the undo manager for the specified text view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextviewdelegate/1449225-undomanagerfortextview?language=objc
func (t_ TextViewDelegateWrapper) UndoManagerForTextView(view ITextView) foundation.UndoManager {
	rv := objc.Call[foundation.UndoManager](t_, objc.Sel("undoManagerForTextView:"), objc.Ptr(view))
	return rv
}

func (t_ TextViewDelegateWrapper) HasTextViewCandidatesForSelectedRange() bool {
	return t_.RespondsToSelector(objc.Sel("textView:candidatesForSelectedRange:"))
}

// Returns an array of objects that represent the elements of a selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextviewdelegate/2544744-textview?language=objc
func (t_ TextViewDelegateWrapper) TextViewCandidatesForSelectedRange(textView ITextView, selectedRange foundation.Range) []objc.Object {
	rv := objc.Call[[]objc.Object](t_, objc.Sel("textView:candidatesForSelectedRange:"), objc.Ptr(textView), selectedRange)
	return rv
}

func (t_ TextViewDelegateWrapper) HasTextViewDidChangeSelection() bool {
	return t_.RespondsToSelector(objc.Sel("textViewDidChangeSelection:"))
}

// Sent when the selection changes in the text view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextviewdelegate/1449363-textviewdidchangeselection?language=objc
func (t_ TextViewDelegateWrapper) TextViewDidChangeSelection(notification foundation.INotification) {
	objc.Call[objc.Void](t_, objc.Sel("textViewDidChangeSelection:"), objc.Ptr(notification))
}
