// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The optional methods that delegates of text storage objects implement to handle text-edit processing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextstoragedelegate?language=objc
type PTextStorageDelegate interface {
	// optional
	TextStorageDidProcessEditingRangeChangeInLength(textStorage TextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int)
	HasTextStorageDidProcessEditingRangeChangeInLength() bool
}

// A delegate implementation builder for the [PTextStorageDelegate] protocol.
type TextStorageDelegate struct {
	_TextStorageDidProcessEditingRangeChangeInLength func(textStorage TextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int)
}

func (di *TextStorageDelegate) HasTextStorageDidProcessEditingRangeChangeInLength() bool {
	return di._TextStorageDidProcessEditingRangeChangeInLength != nil
}

// The method the framework calls when a text storage object has finished processing edits. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextstoragedelegate/1534375-textstorage?language=objc
func (di *TextStorageDelegate) SetTextStorageDidProcessEditingRangeChangeInLength(f func(textStorage TextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int)) {
	di._TextStorageDidProcessEditingRangeChangeInLength = f
}

// The method the framework calls when a text storage object has finished processing edits. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextstoragedelegate/1534375-textstorage?language=objc
func (di *TextStorageDelegate) TextStorageDidProcessEditingRangeChangeInLength(textStorage TextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int) {
	di._TextStorageDidProcessEditingRangeChangeInLength(textStorage, editedMask, editedRange, delta)
}

// A concrete type wrapper for the [PTextStorageDelegate] protocol.
type TextStorageDelegateWrapper struct {
	objc.Object
}

func (t_ TextStorageDelegateWrapper) HasTextStorageDidProcessEditingRangeChangeInLength() bool {
	return t_.RespondsToSelector(objc.Sel("textStorage:didProcessEditing:range:changeInLength:"))
}

// The method the framework calls when a text storage object has finished processing edits. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextstoragedelegate/1534375-textstorage?language=objc
func (t_ TextStorageDelegateWrapper) TextStorageDidProcessEditingRangeChangeInLength(textStorage ITextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int) {
	objc.Call[objc.Void](t_, objc.Sel("textStorage:didProcessEditing:range:changeInLength:"), objc.Ptr(textStorage), editedMask, editedRange, delta)
}
