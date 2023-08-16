// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The optional methods that delegates of content storage objects implement to handle content processing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentstoragedelegate?language=objc
type PTextContentStorageDelegate interface {
	// optional
	TextContentStorageTextParagraphWithRange(textContentStorage TextContentStorage, range_ foundation.Range) ITextParagraph
	HasTextContentStorageTextParagraphWithRange() bool
}

// A delegate implementation builder for the [PTextContentStorageDelegate] protocol.
type TextContentStorageDelegate struct {
	_TextContentStorageTextParagraphWithRange func(textContentStorage TextContentStorage, range_ foundation.Range) ITextParagraph
}

func (di *TextContentStorageDelegate) HasTextContentStorageTextParagraphWithRange() bool {
	return di._TextContentStorageTextParagraphWithRange != nil
}

// Returns a custom paragraph for a range that you provide from the object’s attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentstoragedelegate/3809940-textcontentstorage?language=objc
func (di *TextContentStorageDelegate) SetTextContentStorageTextParagraphWithRange(f func(textContentStorage TextContentStorage, range_ foundation.Range) ITextParagraph) {
	di._TextContentStorageTextParagraphWithRange = f
}

// Returns a custom paragraph for a range that you provide from the object’s attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentstoragedelegate/3809940-textcontentstorage?language=objc
func (di *TextContentStorageDelegate) TextContentStorageTextParagraphWithRange(textContentStorage TextContentStorage, range_ foundation.Range) ITextParagraph {
	return di._TextContentStorageTextParagraphWithRange(textContentStorage, range_)
}

// A concrete type wrapper for the [PTextContentStorageDelegate] protocol.
type TextContentStorageDelegateWrapper struct {
	objc.Object
}

func (t_ TextContentStorageDelegateWrapper) HasTextContentStorageTextParagraphWithRange() bool {
	return t_.RespondsToSelector(objc.Sel("textContentStorage:textParagraphWithRange:"))
}

// Returns a custom paragraph for a range that you provide from the object’s attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentstoragedelegate/3809940-textcontentstorage?language=objc
func (t_ TextContentStorageDelegateWrapper) TextContentStorageTextParagraphWithRange(textContentStorage ITextContentStorage, range_ foundation.Range) TextParagraph {
	rv := objc.Call[TextParagraph](t_, objc.Sel("textContentStorage:textParagraphWithRange:"), objc.Ptr(textContentStorage), range_)
	return rv
}
