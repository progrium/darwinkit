// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that a search field delegate can use to determine when a search started or ended. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfielddelegate?language=objc
type PSearchFieldDelegate interface {
	PTextFieldDelegate
	// optional
	SearchFieldDidStartSearching(sender SearchField)
	HasSearchFieldDidStartSearching() bool

	// optional
	SearchFieldDidEndSearching(sender SearchField)
	HasSearchFieldDidEndSearching() bool
}

// A delegate implementation builder for the [PSearchFieldDelegate] protocol.
type SearchFieldDelegate struct {
	TextFieldDelegate
	_SearchFieldDidStartSearching func(sender SearchField)
	_SearchFieldDidEndSearching   func(sender SearchField)
}

func (di *SearchFieldDelegate) HasSearchFieldDidStartSearching() bool {
	return di._SearchFieldDidStartSearching != nil
}

// The method that is called when the search field begins searching for content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfielddelegate/1535275-searchfielddidstartsearching?language=objc
func (di *SearchFieldDelegate) SetSearchFieldDidStartSearching(f func(sender SearchField)) {
	di._SearchFieldDidStartSearching = f
}

// The method that is called when the search field begins searching for content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfielddelegate/1535275-searchfielddidstartsearching?language=objc
func (di *SearchFieldDelegate) SearchFieldDidStartSearching(sender SearchField) {
	di._SearchFieldDidStartSearching(sender)
}
func (di *SearchFieldDelegate) HasSearchFieldDidEndSearching() bool {
	return di._SearchFieldDidEndSearching != nil
}

// The method that is called when the search field has ended its search for content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfielddelegate/1529867-searchfielddidendsearching?language=objc
func (di *SearchFieldDelegate) SetSearchFieldDidEndSearching(f func(sender SearchField)) {
	di._SearchFieldDidEndSearching = f
}

// The method that is called when the search field has ended its search for content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfielddelegate/1529867-searchfielddidendsearching?language=objc
func (di *SearchFieldDelegate) SearchFieldDidEndSearching(sender SearchField) {
	di._SearchFieldDidEndSearching(sender)
}

// A concrete type wrapper for the [PSearchFieldDelegate] protocol.
type SearchFieldDelegateWrapper struct {
	TextFieldDelegateWrapper
}

func (s_ SearchFieldDelegateWrapper) HasSearchFieldDidStartSearching() bool {
	return s_.RespondsToSelector(objc.Sel("searchFieldDidStartSearching:"))
}

// The method that is called when the search field begins searching for content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfielddelegate/1535275-searchfielddidstartsearching?language=objc
func (s_ SearchFieldDelegateWrapper) SearchFieldDidStartSearching(sender ISearchField) {
	objc.Call[objc.Void](s_, objc.Sel("searchFieldDidStartSearching:"), objc.Ptr(sender))
}

func (s_ SearchFieldDelegateWrapper) HasSearchFieldDidEndSearching() bool {
	return s_.RespondsToSelector(objc.Sel("searchFieldDidEndSearching:"))
}

// The method that is called when the search field has ended its search for content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfielddelegate/1529867-searchfielddidendsearching?language=objc
func (s_ SearchFieldDelegateWrapper) SearchFieldDidEndSearching(sender ISearchField) {
	objc.Call[objc.Void](s_, objc.Sel("searchFieldDidEndSearching:"), objc.Ptr(sender))
}
