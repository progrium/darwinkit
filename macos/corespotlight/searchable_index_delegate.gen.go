// Code generated by DarwinKit. DO NOT EDIT.

package corespotlight

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// A protocol defining methods a delegate object or app extension uses to handle communication from the on-device index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindexdelegate?language=objc
type PSearchableIndexDelegate interface {
	// optional
	SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler(searchableIndex SearchableIndex, identifiers []string, acknowledgementHandler func())
	HasSearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler() bool

	// optional
	SearchableIndexDidFinishThrottle(searchableIndex SearchableIndex)
	HasSearchableIndexDidFinishThrottle() bool

	// optional
	FileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError(searchableIndex SearchableIndex, itemIdentifier string, typeIdentifier string, inPlace bool, outError unsafe.Pointer) foundation.URL
	HasFileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError() bool

	// optional
	DataForSearchableIndexItemIdentifierTypeIdentifierError(searchableIndex SearchableIndex, itemIdentifier string, typeIdentifier string, outError unsafe.Pointer) []byte
	HasDataForSearchableIndexItemIdentifierTypeIdentifierError() bool

	// optional
	SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler(searchableIndex SearchableIndex, acknowledgementHandler func())
	HasSearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler() bool

	// optional
	SearchableIndexDidThrottle(searchableIndex SearchableIndex)
	HasSearchableIndexDidThrottle() bool
}

// A delegate implementation builder for the [PSearchableIndexDelegate] protocol.
type SearchableIndexDelegate struct {
	_SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler func(searchableIndex SearchableIndex, identifiers []string, acknowledgementHandler func())
	_SearchableIndexDidFinishThrottle                                           func(searchableIndex SearchableIndex)
	_FileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError          func(searchableIndex SearchableIndex, itemIdentifier string, typeIdentifier string, inPlace bool, outError unsafe.Pointer) foundation.URL
	_DataForSearchableIndexItemIdentifierTypeIdentifierError                    func(searchableIndex SearchableIndex, itemIdentifier string, typeIdentifier string, outError unsafe.Pointer) []byte
	_SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler         func(searchableIndex SearchableIndex, acknowledgementHandler func())
	_SearchableIndexDidThrottle                                                 func(searchableIndex SearchableIndex)
}

func (di *SearchableIndexDelegate) HasSearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler() bool {
	return di._SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler != nil
}

// Tells the delegate to reindex the searchable items associated with the specified identifiers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindexdelegate/1620338-searchableindex?language=objc
func (di *SearchableIndexDelegate) SetSearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler(f func(searchableIndex SearchableIndex, identifiers []string, acknowledgementHandler func())) {
	di._SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler = f
}

// Tells the delegate to reindex the searchable items associated with the specified identifiers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindexdelegate/1620338-searchableindex?language=objc
func (di *SearchableIndexDelegate) SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler(searchableIndex SearchableIndex, identifiers []string, acknowledgementHandler func()) {
	di._SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler(searchableIndex, identifiers, acknowledgementHandler)
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
func (di *SearchableIndexDelegate) HasFileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError() bool {
	return di._FileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError != nil
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindexdelegate/2867898-fileurlforsearchableindex?language=objc
func (di *SearchableIndexDelegate) SetFileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError(f func(searchableIndex SearchableIndex, itemIdentifier string, typeIdentifier string, inPlace bool, outError unsafe.Pointer) foundation.URL) {
	di._FileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError = f
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindexdelegate/2867898-fileurlforsearchableindex?language=objc
func (di *SearchableIndexDelegate) FileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError(searchableIndex SearchableIndex, itemIdentifier string, typeIdentifier string, inPlace bool, outError unsafe.Pointer) foundation.URL {
	return di._FileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError(searchableIndex, itemIdentifier, typeIdentifier, inPlace, outError)
}
func (di *SearchableIndexDelegate) HasDataForSearchableIndexItemIdentifierTypeIdentifierError() bool {
	return di._DataForSearchableIndexItemIdentifierTypeIdentifierError != nil
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindexdelegate/2867899-dataforsearchableindex?language=objc
func (di *SearchableIndexDelegate) SetDataForSearchableIndexItemIdentifierTypeIdentifierError(f func(searchableIndex SearchableIndex, itemIdentifier string, typeIdentifier string, outError unsafe.Pointer) []byte) {
	di._DataForSearchableIndexItemIdentifierTypeIdentifierError = f
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindexdelegate/2867899-dataforsearchableindex?language=objc
func (di *SearchableIndexDelegate) DataForSearchableIndexItemIdentifierTypeIdentifierError(searchableIndex SearchableIndex, itemIdentifier string, typeIdentifier string, outError unsafe.Pointer) []byte {
	return di._DataForSearchableIndexItemIdentifierTypeIdentifierError(searchableIndex, itemIdentifier, typeIdentifier, outError)
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

// ensure impl type implements protocol interface
var _ PSearchableIndexDelegate = (*SearchableIndexDelegateObject)(nil)

// A concrete type for the [PSearchableIndexDelegate] protocol.
type SearchableIndexDelegateObject struct {
	objc.Object
}

func (s_ SearchableIndexDelegateObject) HasSearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler() bool {
	return s_.RespondsToSelector(objc.Sel("searchableIndex:reindexSearchableItemsWithIdentifiers:acknowledgementHandler:"))
}

// Tells the delegate to reindex the searchable items associated with the specified identifiers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindexdelegate/1620338-searchableindex?language=objc
func (s_ SearchableIndexDelegateObject) SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler(searchableIndex SearchableIndex, identifiers []string, acknowledgementHandler func()) {
	objc.Call[objc.Void](s_, objc.Sel("searchableIndex:reindexSearchableItemsWithIdentifiers:acknowledgementHandler:"), searchableIndex, identifiers, acknowledgementHandler)
}

func (s_ SearchableIndexDelegateObject) HasSearchableIndexDidFinishThrottle() bool {
	return s_.RespondsToSelector(objc.Sel("searchableIndexDidFinishThrottle:"))
}

// Tells the delegate that the index throttling has finished. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindexdelegate/1620343-searchableindexdidfinishthrottle?language=objc
func (s_ SearchableIndexDelegateObject) SearchableIndexDidFinishThrottle(searchableIndex SearchableIndex) {
	objc.Call[objc.Void](s_, objc.Sel("searchableIndexDidFinishThrottle:"), searchableIndex)
}

func (s_ SearchableIndexDelegateObject) HasFileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError() bool {
	return s_.RespondsToSelector(objc.Sel("fileURLForSearchableIndex:itemIdentifier:typeIdentifier:inPlace:error:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindexdelegate/2867898-fileurlforsearchableindex?language=objc
func (s_ SearchableIndexDelegateObject) FileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError(searchableIndex SearchableIndex, itemIdentifier string, typeIdentifier string, inPlace bool, outError unsafe.Pointer) foundation.URL {
	rv := objc.Call[foundation.URL](s_, objc.Sel("fileURLForSearchableIndex:itemIdentifier:typeIdentifier:inPlace:error:"), searchableIndex, itemIdentifier, typeIdentifier, inPlace, outError)
	return rv
}

func (s_ SearchableIndexDelegateObject) HasDataForSearchableIndexItemIdentifierTypeIdentifierError() bool {
	return s_.RespondsToSelector(objc.Sel("dataForSearchableIndex:itemIdentifier:typeIdentifier:error:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindexdelegate/2867899-dataforsearchableindex?language=objc
func (s_ SearchableIndexDelegateObject) DataForSearchableIndexItemIdentifierTypeIdentifierError(searchableIndex SearchableIndex, itemIdentifier string, typeIdentifier string, outError unsafe.Pointer) []byte {
	rv := objc.Call[[]byte](s_, objc.Sel("dataForSearchableIndex:itemIdentifier:typeIdentifier:error:"), searchableIndex, itemIdentifier, typeIdentifier, outError)
	return rv
}

func (s_ SearchableIndexDelegateObject) HasSearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler() bool {
	return s_.RespondsToSelector(objc.Sel("searchableIndex:reindexAllSearchableItemsWithAcknowledgementHandler:"))
}

// Tells the delegate to reindex all searchable data and clear all local state information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindexdelegate/1620348-searchableindex?language=objc
func (s_ SearchableIndexDelegateObject) SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler(searchableIndex SearchableIndex, acknowledgementHandler func()) {
	objc.Call[objc.Void](s_, objc.Sel("searchableIndex:reindexAllSearchableItemsWithAcknowledgementHandler:"), searchableIndex, acknowledgementHandler)
}

func (s_ SearchableIndexDelegateObject) HasSearchableIndexDidThrottle() bool {
	return s_.RespondsToSelector(objc.Sel("searchableIndexDidThrottle:"))
}

// Tells the delegate that indexing is being throttled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindexdelegate/1620353-searchableindexdidthrottle?language=objc
func (s_ SearchableIndexDelegateObject) SearchableIndexDidThrottle(searchableIndex SearchableIndex) {
	objc.Call[objc.Void](s_, objc.Sel("searchableIndexDidThrottle:"), searchableIndex)
}
