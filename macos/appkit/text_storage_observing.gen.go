// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
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
	TextStorage() ITextStorage
	HasTextStorage() bool
}

// A concrete type wrapper for the [PTextStorageObserving] protocol.
type TextStorageObservingWrapper struct {
	objc.Object
}

func (t_ TextStorageObservingWrapper) HasProcessEditingForTextStorageEditedRangeChangeInLengthInvalidatedRange() bool {
	return t_.RespondsToSelector(objc.Sel("processEditingForTextStorage:edited:range:changeInLength:invalidatedRange:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextstorageobserving/3810050-processeditingfortextstorage?language=objc
func (t_ TextStorageObservingWrapper) ProcessEditingForTextStorageEditedRangeChangeInLengthInvalidatedRange(textStorage ITextStorage, editMask TextStorageEditActions, newCharRange foundation.Range, delta int, invalidatedCharRange foundation.Range) {
	objc.Call[objc.Void](t_, objc.Sel("processEditingForTextStorage:edited:range:changeInLength:invalidatedRange:"), objc.Ptr(textStorage), editMask, newCharRange, delta, invalidatedCharRange)
}

func (t_ TextStorageObservingWrapper) HasPerformEditingTransactionForTextStorageUsingBlock() bool {
	return t_.RespondsToSelector(objc.Sel("performEditingTransactionForTextStorage:usingBlock:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextstorageobserving/3852588-performeditingtransactionfortext?language=objc
func (t_ TextStorageObservingWrapper) PerformEditingTransactionForTextStorageUsingBlock(textStorage ITextStorage, transaction func()) {
	objc.Call[objc.Void](t_, objc.Sel("performEditingTransactionForTextStorage:usingBlock:"), objc.Ptr(textStorage), transaction)
}

func (t_ TextStorageObservingWrapper) HasSetTextStorage() bool {
	return t_.RespondsToSelector(objc.Sel("setTextStorage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextstorageobserving/3810051-textstorage?language=objc
func (t_ TextStorageObservingWrapper) SetTextStorage(value ITextStorage) {
	objc.Call[objc.Void](t_, objc.Sel("setTextStorage:"), objc.Ptr(value))
}

func (t_ TextStorageObservingWrapper) HasTextStorage() bool {
	return t_.RespondsToSelector(objc.Sel("textStorage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextstorageobserving/3810051-textstorage?language=objc
func (t_ TextStorageObservingWrapper) TextStorage() TextStorage {
	rv := objc.Call[TextStorage](t_, objc.Sel("textStorage"))
	return rv
}
