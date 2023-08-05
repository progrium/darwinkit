// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type ITextStorageDelegate interface {
	ImplementsTextStorageWillProcessEditingRangeChangeInLength() bool
	// optional
	TextStorageWillProcessEditingRangeChangeInLength(textStorage TextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int)
	ImplementsTextStorageDidProcessEditingRangeChangeInLength() bool
	// optional
	TextStorageDidProcessEditingRangeChangeInLength(textStorage TextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int)
}

type TextStorageDelegate struct {
	_TextStorageWillProcessEditingRangeChangeInLength func(textStorage TextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int)
	_TextStorageDidProcessEditingRangeChangeInLength  func(textStorage TextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int)
}

func (di *TextStorageDelegate) ImplementsTextStorageWillProcessEditingRangeChangeInLength() bool {
	return di._TextStorageWillProcessEditingRangeChangeInLength != nil
}

func (di *TextStorageDelegate) SetTextStorageWillProcessEditingRangeChangeInLength(f func(textStorage TextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int)) {
	di._TextStorageWillProcessEditingRangeChangeInLength = f
}

func (di *TextStorageDelegate) TextStorageWillProcessEditingRangeChangeInLength(textStorage TextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int) {
	di._TextStorageWillProcessEditingRangeChangeInLength(textStorage, editedMask, editedRange, delta)
}
func (di *TextStorageDelegate) ImplementsTextStorageDidProcessEditingRangeChangeInLength() bool {
	return di._TextStorageDidProcessEditingRangeChangeInLength != nil
}

func (di *TextStorageDelegate) SetTextStorageDidProcessEditingRangeChangeInLength(f func(textStorage TextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int)) {
	di._TextStorageDidProcessEditingRangeChangeInLength = f
}

func (di *TextStorageDelegate) TextStorageDidProcessEditingRangeChangeInLength(textStorage TextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int) {
	di._TextStorageDidProcessEditingRangeChangeInLength(textStorage, editedMask, editedRange, delta)
}

type TextStorageDelegateWrapper struct {
	objc.Object
}

func (t_ TextStorageDelegateWrapper) ImplementsTextStorageWillProcessEditingRangeChangeInLength() bool {
	return t_.RespondsToSelector(objc.GetSelector("textStorage:willProcessEditing:range:changeInLength:"))
}

func (t_ TextStorageDelegateWrapper) TextStorageWillProcessEditingRangeChangeInLength(textStorage ITextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("textStorage:willProcessEditing:range:changeInLength:"), objc.ExtractPtr(textStorage), editedMask, editedRange, delta)
}

func (t_ TextStorageDelegateWrapper) ImplementsTextStorageDidProcessEditingRangeChangeInLength() bool {
	return t_.RespondsToSelector(objc.GetSelector("textStorage:didProcessEditing:range:changeInLength:"))
}

func (t_ TextStorageDelegateWrapper) TextStorageDidProcessEditingRangeChangeInLength(textStorage ITextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("textStorage:didProcessEditing:range:changeInLength:"), objc.ExtractPtr(textStorage), editedMask, editedRange, delta)
}
