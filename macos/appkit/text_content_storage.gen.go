// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TextContentStorage] class.
var TextContentStorageClass = _TextContentStorageClass{objc.GetClass("NSTextContentStorage")}

type _TextContentStorageClass struct {
	objc.Class
}

// An interface definition for the [TextContentStorage] class.
type ITextContentStorage interface {
	ITextContentManager
	AdjustedRangeFromRangeForEditingTextSelection(textRange ITextRange, forEditingTextSelection bool) TextRange
	OffsetFromLocationToLocation(from PTextLocation, to PTextLocation) int
	OffsetFromLocationObjectToLocationObject(fromObject objc.IObject, toObject objc.IObject) int
	AttributedStringForTextElement(textElement ITextElement) foundation.AttributedString
	TextElementForAttributedString(attributedString foundation.IAttributedString) TextElement
	LocationFromLocationWithOffset(location PTextLocation, offset int) TextLocationWrapper
	LocationFromLocationObjectWithOffset(locationObject objc.IObject, offset int) TextLocationWrapper
	AttributedString() foundation.AttributedString
	SetAttributedString(value foundation.IAttributedString)
}

// A concrete subclass of the text content manager that provides support for attributed strings that you use as a backing store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentstorage?language=objc
type TextContentStorage struct {
	TextContentManager
}

func TextContentStorageFrom(ptr unsafe.Pointer) TextContentStorage {
	return TextContentStorage{
		TextContentManager: TextContentManagerFrom(ptr),
	}
}

func (tc _TextContentStorageClass) Alloc() TextContentStorage {
	rv := objc.Call[TextContentStorage](tc, objc.Sel("alloc"))
	return rv
}

func TextContentStorage_Alloc() TextContentStorage {
	return TextContentStorageClass.Alloc()
}

func (tc _TextContentStorageClass) New() TextContentStorage {
	rv := objc.Call[TextContentStorage](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTextContentStorage() TextContentStorage {
	return TextContentStorageClass.New()
}

func (t_ TextContentStorage) Init() TextContentStorage {
	rv := objc.Call[TextContentStorage](t_, objc.Sel("init"))
	return rv
}

// Returns the text range, if any, in the backing store that required manual adjustment after editing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentstorage/3852570-adjustedrangefromrange?language=objc
func (t_ TextContentStorage) AdjustedRangeFromRangeForEditingTextSelection(textRange ITextRange, forEditingTextSelection bool) TextRange {
	rv := objc.Call[TextRange](t_, objc.Sel("adjustedRangeFromRange:forEditingTextSelection:"), objc.Ptr(textRange), forEditingTextSelection)
	return rv
}

// Returns the offset between the two specified locations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentstorage/3852572-offsetfromlocation?language=objc
func (t_ TextContentStorage) OffsetFromLocationToLocation(from PTextLocation, to PTextLocation) int {
	po0 := objc.WrapAsProtocol("NSTextLocation", from)
	po1 := objc.WrapAsProtocol("NSTextLocation", to)
	rv := objc.Call[int](t_, objc.Sel("offsetFromLocation:toLocation:"), po0, po1)
	return rv
}

// Returns the offset between the two specified locations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentstorage/3852572-offsetfromlocation?language=objc
func (t_ TextContentStorage) OffsetFromLocationObjectToLocationObject(fromObject objc.IObject, toObject objc.IObject) int {
	rv := objc.Call[int](t_, objc.Sel("offsetFromLocation:toLocation:"), objc.Ptr(fromObject), objc.Ptr(toObject))
	return rv
}

// Returns a new attributed string for the text element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentstorage/3809935-attributedstringfortextelement?language=objc
func (t_ TextContentStorage) AttributedStringForTextElement(textElement ITextElement) foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](t_, objc.Sel("attributedStringForTextElement:"), objc.Ptr(textElement))
	return rv
}

// Returns the text element corresponding to object’s attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentstorage/3809937-textelementforattributedstring?language=objc
func (t_ TextContentStorage) TextElementForAttributedString(attributedString foundation.IAttributedString) TextElement {
	rv := objc.Call[TextElement](t_, objc.Sel("textElementForAttributedString:"), objc.Ptr(attributedString))
	return rv
}

// Returns a new text location from an existing location and offset you provide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentstorage/3852571-locationfromlocation?language=objc
func (t_ TextContentStorage) LocationFromLocationWithOffset(location PTextLocation, offset int) TextLocationWrapper {
	po0 := objc.WrapAsProtocol("NSTextLocation", location)
	rv := objc.Call[TextLocationWrapper](t_, objc.Sel("locationFromLocation:withOffset:"), po0, offset)
	return rv
}

// Returns a new text location from an existing location and offset you provide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentstorage/3852571-locationfromlocation?language=objc
func (t_ TextContentStorage) LocationFromLocationObjectWithOffset(locationObject objc.IObject, offset int) TextLocationWrapper {
	rv := objc.Call[TextLocationWrapper](t_, objc.Sel("locationFromLocation:withOffset:"), objc.Ptr(locationObject), offset)
	return rv
}

// The contents of a document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentstorage/3809934-attributedstring?language=objc
func (t_ TextContentStorage) AttributedString() foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](t_, objc.Sel("attributedString"))
	return rv
}

// The contents of a document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentstorage/3809934-attributedstring?language=objc
func (t_ TextContentStorage) SetAttributedString(value foundation.IAttributedString) {
	objc.Call[objc.Void](t_, objc.Sel("setAttributedString:"), objc.Ptr(value))
}
