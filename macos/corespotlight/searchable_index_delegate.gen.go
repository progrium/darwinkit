// AUTO-GENERATED CODE, DO NOT MODIFY

package corespotlight

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A protocol defining methods a delegate object or app extension uses to handle communication from the on-device index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindexdelegate?language=objc
type PSearchableIndexDelegate interface {
	// optional
	SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler(searchableIndex SearchableIndex, acknowledgementHandler func())
	HasSearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler() bool

	// optional
	SearchableIndexDidThrottle(searchableIndex SearchableIndex)
	HasSearchableIndexDidThrottle() bool

	// optional
	SearchableIndexDidFinishThrottle(searchableIndex SearchableIndex)
	HasSearchableIndexDidFinishThrottle() bool

	// optional
	DataForSearchableIndexItemIdentifierTypeIdentifierError(searchableIndex SearchableIndex, itemIdentifier string, typeIdentifier string, outError foundation.Error) []byte
	HasDataForSearchableIndexItemIdentifierTypeIdentifierError() bool

	// optional
	FileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError(searchableIndex SearchableIndex, itemIdentifier string, typeIdentifier string, inPlace bool, outError foundation.Error) foundation.IURL
	HasFileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError() bool
}

// A delegate implementation builder for the [PSearchableIndexDelegate] protocol.
type SearchableIndexDelegate struct {
	_SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler func(searchableIndex SearchableIndex, acknowledgementHandler func())
	_SearchableIndexDidThrottle                                         func(searchableIndex SearchableIndex)
	_SearchableIndexDidFinishThrottle                                   func(searchableIndex SearchableIndex)
	_DataForSearchableIndexItemIdentifierTypeIdentifierError            func(searchableIndex SearchableIndex, itemIdentifier string, typeIdentifier string, outError foundation.Error) []byte
	_FileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError  func(searchableIndex SearchableIndex, itemIdentifier string, typeIdentifier string, inPlace bool, outError foundation.Error) foundation.IURL
}

func (di *SearchableIndexDelegate) HasSearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler() bool {
	return di._SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler != nil
}

// Tells the delegate to reindex all searchable data and clear all local state information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindexdelegate/1620348-searchableindex?language=objc
func (di *SearchableIndexDelegate) SetSearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler(f func(searchableIndex SearchableIndex, acknowledgementHandler func())) {
	di._SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler = f
}

// Tells the delegate to reindex all searchable data and clear all local state information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindexdelegate/1620348-searchableindex?language=objc
func (di *SearchableIndexDelegate) SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler(searchableIndex SearchableIndex, acknowledgementHandler func()) {
	di._SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler(searchableIndex, acknowledgementHandler)
}
func (di *SearchableIndexDelegate) HasSearchableIndexDidThrottle() bool {
	return di._SearchableIndexDidThrottle != nil
}

// Tells the delegate that indexing is being throttled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindexdelegate/1620353-searchableindexdidthrottle?language=objc
func (di *SearchableIndexDelegate) SetSearchableIndexDidThrottle(f func(searchableIndex SearchableIndex)) {
	di._SearchableIndexDidThrottle = f
}

// Tells the delegate that indexing is being throttled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindexdelegate/1620353-searchableindexdidthrottle?language=objc
func (di *SearchableIndexDelegate) SearchableIndexDidThrottle(searchableIndex SearchableIndex) {
	di._SearchableIndexDidThrottle(searchableIndex)
}
func (di *SearchableIndexDelegate) HasSearchableIndexDidFinishThrottle() bool {
	return di._SearchableIndexDidFinishThrottle != nil
}

// Tells the delegate that the index throttling has finished. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindexdelegate/1620343-searchableindexdidfinishthrottle?language=objc
func (di *SearchableIndexDelegate) SetSearchableIndexDidFinishThrottle(f func(searchableIndex SearchableIndex)) {
	di._SearchableIndexDidFinishThrottle = f
}

// Tells the delegate that the index throttling has finished. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindexdelegate/1620343-searchableindexdidfinishthrottle?language=objc
func (di *SearchableIndexDelegate) SearchableIndexDidFinishThrottle(searchableIndex SearchableIndex) {
	di._SearchableIndexDidFinishThrottle(searchableIndex)
}
func (di *SearchableIndexDelegate) HasDataForSearchableIndexItemIdentifierTypeIdentifierError() bool {
	return di._DataForSearchableIndexItemIdentifierTypeIdentifierError != nil
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindexdelegate/2867899-dataforsearchableindex?language=objc
func (di *SearchableIndexDelegate) SetDataForSearchableIndexItemIdentifierTypeIdentifierError(f func(searchableIndex SearchableIndex, itemIdentifier string, typeIdentifier string, outError foundation.Error) []byte) {
	di._DataForSearchableIndexItemIdentifierTypeIdentifierError = f
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindexdelegate/2867899-dataforsearchableindex?language=objc
func (di *SearchableIndexDelegate) DataForSearchableIndexItemIdentifierTypeIdentifierError(searchableIndex SearchableIndex, itemIdentifier string, typeIdentifier string, outError foundation.Error) []byte {
	return di._DataForSearchableIndexItemIdentifierTypeIdentifierError(searchableIndex, itemIdentifier, typeIdentifier, outError)
}
func (di *SearchableIndexDelegate) HasFileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError() bool {
	return di._FileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError != nil
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindexdelegate/2867898-fileurlforsearchableindex?language=objc
func (di *SearchableIndexDelegate) SetFileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError(f func(searchableIndex SearchableIndex, itemIdentifier string, typeIdentifier string, inPlace bool, outError foundation.Error) foundation.IURL) {
	di._FileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError = f
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindexdelegate/2867898-fileurlforsearchableindex?language=objc
func (di *SearchableIndexDelegate) FileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError(searchableIndex SearchableIndex, itemIdentifier string, typeIdentifier string, inPlace bool, outError foundation.Error) foundation.IURL {
	return di._FileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError(searchableIndex, itemIdentifier, typeIdentifier, inPlace, outError)
}

// A concrete type wrapper for the [PSearchableIndexDelegate] protocol.
type SearchableIndexDelegateWrapper struct {
	objc.Object
}

func (s_ SearchableIndexDelegateWrapper) HasSearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler() bool {
	return s_.RespondsToSelector(objc.Sel("searchableIndex:reindexAllSearchableItemsWithAcknowledgementHandler:"))
}

// Tells the delegate to reindex all searchable data and clear all local state information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindexdelegate/1620348-searchableindex?language=objc
func (s_ SearchableIndexDelegateWrapper) SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler(searchableIndex ISearchableIndex, acknowledgementHandler func()) {
	objc.Call[objc.Void](s_, objc.Sel("searchableIndex:reindexAllSearchableItemsWithAcknowledgementHandler:"), objc.Ptr(searchableIndex), acknowledgementHandler)
}

func (s_ SearchableIndexDelegateWrapper) HasSearchableIndexDidThrottle() bool {
	return s_.RespondsToSelector(objc.Sel("searchableIndexDidThrottle:"))
}

// Tells the delegate that indexing is being throttled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindexdelegate/1620353-searchableindexdidthrottle?language=objc
func (s_ SearchableIndexDelegateWrapper) SearchableIndexDidThrottle(searchableIndex ISearchableIndex) {
	objc.Call[objc.Void](s_, objc.Sel("searchableIndexDidThrottle:"), objc.Ptr(searchableIndex))
}

func (s_ SearchableIndexDelegateWrapper) HasSearchableIndexDidFinishThrottle() bool {
	return s_.RespondsToSelector(objc.Sel("searchableIndexDidFinishThrottle:"))
}

// Tells the delegate that the index throttling has finished. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindexdelegate/1620343-searchableindexdidfinishthrottle?language=objc
func (s_ SearchableIndexDelegateWrapper) SearchableIndexDidFinishThrottle(searchableIndex ISearchableIndex) {
	objc.Call[objc.Void](s_, objc.Sel("searchableIndexDidFinishThrottle:"), objc.Ptr(searchableIndex))
}

func (s_ SearchableIndexDelegateWrapper) HasDataForSearchableIndexItemIdentifierTypeIdentifierError() bool {
	return s_.RespondsToSelector(objc.Sel("dataForSearchableIndex:itemIdentifier:typeIdentifier:error:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindexdelegate/2867899-dataforsearchableindex?language=objc
func (s_ SearchableIndexDelegateWrapper) DataForSearchableIndexItemIdentifierTypeIdentifierError(searchableIndex ISearchableIndex, itemIdentifier string, typeIdentifier string, outError foundation.IError) []byte {
	rv := objc.Call[[]byte](s_, objc.Sel("dataForSearchableIndex:itemIdentifier:typeIdentifier:error:"), objc.Ptr(searchableIndex), itemIdentifier, typeIdentifier, objc.Ptr(outError))
	return rv
}

func (s_ SearchableIndexDelegateWrapper) HasFileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError() bool {
	return s_.RespondsToSelector(objc.Sel("fileURLForSearchableIndex:itemIdentifier:typeIdentifier:inPlace:error:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindexdelegate/2867898-fileurlforsearchableindex?language=objc
func (s_ SearchableIndexDelegateWrapper) FileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError(searchableIndex ISearchableIndex, itemIdentifier string, typeIdentifier string, inPlace bool, outError foundation.IError) foundation.URL {
	rv := objc.Call[foundation.URL](s_, objc.Sel("fileURLForSearchableIndex:itemIdentifier:typeIdentifier:inPlace:error:"), objc.Ptr(searchableIndex), itemIdentifier, typeIdentifier, inPlace, objc.Ptr(outError))
	return rv
}
