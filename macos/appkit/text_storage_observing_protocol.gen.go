// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// Optional methods that delegates implement to handle editing and transaction processing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextstorageobserving?language=objc
type PTextStorageObserving interface {
	// optional
	ProcessEditingForTextStorageEditedRangeChangeInLengthInvalidatedRange(textStorage TextStorage, editMask TextStorageEditActions, newCharRange foundation.Range, delta int, invalidatedCharRange foundation.Range)
	HasProcessEditingForTextStorageEditedRangeChangeInLengthInvalidatedRange() bool

	// optional
	PerformEditingTransactionForTextStorageUsingBlock(textStorage TextStorage, transaction func())
	HasPerformEditingTransactionForTextStorageUsingBlock() bool

	// optional
	SetTextStorage(value TextStorage)
	HasSetTextStorage() bool

	// optional
	TextStorage() TextStorage
	HasTextStorage() bool
}

// ensure impl type implements protocol interface
var _ PTextStorageObserving = (*TextStorageObservingObject)(nil)

// A concrete type for the [PTextStorageObserving] protocol.
type TextStorageObservingObject struct {
	objc.Object
}

func (t_ TextStorageObservingObject) HasProcessEditingForTextStorageEditedRangeChangeInLengthInvalidatedRange() bool {
	return t_.RespondsToSelector(objc.Sel("processEditingForTextStorage:edited:range:changeInLength:invalidatedRange:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextstorageobserving/3810050-processeditingfortextstorage?language=objc
func (t_ TextStorageObservingObject) ProcessEditingForTextStorageEditedRangeChangeInLengthInvalidatedRange(textStorage TextStorage, editMask TextStorageEditActions, newCharRange foundation.Range, delta int, invalidatedCharRange foundation.Range) {
	objc.Call[objc.Void](t_, objc.Sel("processEditingForTextStorage:edited:range:changeInLength:invalidatedRange:"), textStorage, editMask, newCharRange, delta, invalidatedCharRange)
}

func (t_ TextStorageObservingObject) HasPerformEditingTransactionForTextStorageUsingBlock() bool {
	return t_.RespondsToSelector(objc.Sel("performEditingTransactionForTextStorage:usingBlock:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextstorageobserving/3852588-performeditingtransactionfortext?language=objc
func (t_ TextStorageObservingObject) PerformEditingTransactionForTextStorageUsingBlock(textStorage TextStorage, transaction func()) {
	objc.Call[objc.Void](t_, objc.Sel("performEditingTransactionForTextStorage:usingBlock:"), textStorage, transaction)
}

func (t_ TextStorageObservingObject) HasSetTextStorage() bool {
	return t_.RespondsToSelector(objc.Sel("setTextStorage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextstorageobserving/3810051-textstorage?language=objc
func (t_ TextStorageObservingObject) SetTextStorage(value TextStorage) {
	objc.Call[objc.Void](t_, objc.Sel("setTextStorage:"), value)
}

func (t_ TextStorageObservingObject) HasTextStorage() bool {
	return t_.RespondsToSelector(objc.Sel("textStorage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextstorageobserving/3810051-textstorage?language=objc
func (t_ TextStorageObservingObject) TextStorage() TextStorage {
	rv := objc.Call[TextStorage](t_, objc.Sel("textStorage"))
	return rv
}
