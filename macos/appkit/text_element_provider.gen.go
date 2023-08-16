// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A protocol the text content manager and its concrete subclasses conform to which defines the interface for interacting with custom content types of a text document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextelementprovider?language=objc
type PTextElementProvider interface {
	// optional
	ReplaceContentsInRangeWithTextElements(range_ TextRange, textElements []TextElement)
	HasReplaceContentsInRangeWithTextElements() bool

	// optional
	AdjustedRangeFromRangeForEditingTextSelection(textRange TextRange, forEditingTextSelection bool) ITextRange
	HasAdjustedRangeFromRangeForEditingTextSelection() bool

	// optional
	EnumerateTextElementsFromLocationOptionsUsingBlock(textLocation TextLocationWrapper, options TextContentManagerEnumerationOptions, block func(element TextElement) bool) PTextLocation
	HasEnumerateTextElementsFromLocationOptionsUsingBlock() bool

	// optional
	OffsetFromLocationToLocation(from TextLocationWrapper, to TextLocationWrapper) int
	HasOffsetFromLocationToLocation() bool

	// optional
	SynchronizeToBackingStore(completionHandler func(error foundation.Error))
	HasSynchronizeToBackingStore() bool

	// optional
	LocationFromLocationWithOffset(location TextLocationWrapper, offset int) PTextLocation
	HasLocationFromLocationWithOffset() bool

	// optional
	DocumentRange() ITextRange
	HasDocumentRange() bool
}

// A concrete type wrapper for the [PTextElementProvider] protocol.
type TextElementProviderWrapper struct {
	objc.Object
}

func (t_ TextElementProviderWrapper) HasReplaceContentsInRangeWithTextElements() bool {
	return t_.RespondsToSelector(objc.Sel("replaceContentsInRange:withTextElements:"))
}

// Replaces the characters specified by range with the text elements you provide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextelementprovider/3809948-replacecontentsinrange?language=objc
func (t_ TextElementProviderWrapper) ReplaceContentsInRangeWithTextElements(range_ ITextRange, textElements []ITextElement) {
	objc.Call[objc.Void](t_, objc.Sel("replaceContentsInRange:withTextElements:"), objc.Ptr(range_), textElements)
}

func (t_ TextElementProviderWrapper) HasAdjustedRangeFromRangeForEditingTextSelection() bool {
	return t_.RespondsToSelector(objc.Sel("adjustedRangeFromRange:forEditingTextSelection:"))
}

// A method you implement if the location backing store requires manual adjustment after editing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextelementprovider/3852573-adjustedrangefromrange?language=objc
func (t_ TextElementProviderWrapper) AdjustedRangeFromRangeForEditingTextSelection(textRange ITextRange, forEditingTextSelection bool) TextRange {
	rv := objc.Call[TextRange](t_, objc.Sel("adjustedRangeFromRange:forEditingTextSelection:"), objc.Ptr(textRange), forEditingTextSelection)
	return rv
}

func (t_ TextElementProviderWrapper) HasEnumerateTextElementsFromLocationOptionsUsingBlock() bool {
	return t_.RespondsToSelector(objc.Sel("enumerateTextElementsFromLocation:options:usingBlock:"))
}

// Enumerates text elements starting at the text location you provide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextelementprovider/3809945-enumeratetextelementsfromlocatio?language=objc
func (t_ TextElementProviderWrapper) EnumerateTextElementsFromLocationOptionsUsingBlock(textLocation PTextLocation, options TextContentManagerEnumerationOptions, block func(element TextElement) bool) TextLocationWrapper {
	po0 := objc.WrapAsProtocol("NSTextLocation", textLocation)
	rv := objc.Call[TextLocationWrapper](t_, objc.Sel("enumerateTextElementsFromLocation:options:usingBlock:"), po0, options, block)
	return rv
}

func (t_ TextElementProviderWrapper) HasOffsetFromLocationToLocation() bool {
	return t_.RespondsToSelector(objc.Sel("offsetFromLocation:toLocation:"))
}

// Returns the offset between the two specified locations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextelementprovider/3809947-offsetfromlocation?language=objc
func (t_ TextElementProviderWrapper) OffsetFromLocationToLocation(from PTextLocation, to PTextLocation) int {
	po0 := objc.WrapAsProtocol("NSTextLocation", from)
	po1 := objc.WrapAsProtocol("NSTextLocation", to)
	rv := objc.Call[int](t_, objc.Sel("offsetFromLocation:toLocation:"), po0, po1)
	return rv
}

func (t_ TextElementProviderWrapper) HasSynchronizeToBackingStore() bool {
	return t_.RespondsToSelector(objc.Sel("synchronizeToBackingStore:"))
}

// Synchronizes changes to the backing store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextelementprovider/3809949-synchronizetobackingstore?language=objc
func (t_ TextElementProviderWrapper) SynchronizeToBackingStore(completionHandler func(error foundation.Error)) {
	objc.Call[objc.Void](t_, objc.Sel("synchronizeToBackingStore:"), completionHandler)
}

func (t_ TextElementProviderWrapper) HasLocationFromLocationWithOffset() bool {
	return t_.RespondsToSelector(objc.Sel("locationFromLocation:withOffset:"))
}

// Returns a new location from location with offset you provide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextelementprovider/3809946-locationfromlocation?language=objc
func (t_ TextElementProviderWrapper) LocationFromLocationWithOffset(location PTextLocation, offset int) TextLocationWrapper {
	po0 := objc.WrapAsProtocol("NSTextLocation", location)
	rv := objc.Call[TextLocationWrapper](t_, objc.Sel("locationFromLocation:withOffset:"), po0, offset)
	return rv
}

func (t_ TextElementProviderWrapper) HasDocumentRange() bool {
	return t_.RespondsToSelector(objc.Sel("documentRange"))
}

// Describes the starting and ending locations for the document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextelementprovider/3809944-documentrange?language=objc
func (t_ TextElementProviderWrapper) DocumentRange() TextRange {
	rv := objc.Call[TextRange](t_, objc.Sel("documentRange"))
	return rv
}
