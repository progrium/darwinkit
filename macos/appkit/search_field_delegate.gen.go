// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"github.com/progrium/macdriver/objc"
)

type ISearchFieldDelegate interface {
	ITextFieldDelegate
	ImplementsSearchFieldDidStartSearching() bool
	// optional
	SearchFieldDidStartSearching(sender SearchField)
	ImplementsSearchFieldDidEndSearching() bool
	// optional
	SearchFieldDidEndSearching(sender SearchField)
}

type SearchFieldDelegate struct {
	TextFieldDelegate
	_SearchFieldDidStartSearching func(sender SearchField)
	_SearchFieldDidEndSearching   func(sender SearchField)
}

func (di *SearchFieldDelegate) ImplementsSearchFieldDidStartSearching() bool {
	return di._SearchFieldDidStartSearching != nil
}

func (di *SearchFieldDelegate) SetSearchFieldDidStartSearching(f func(sender SearchField)) {
	di._SearchFieldDidStartSearching = f
}

func (di *SearchFieldDelegate) SearchFieldDidStartSearching(sender SearchField) {
	di._SearchFieldDidStartSearching(sender)
}
func (di *SearchFieldDelegate) ImplementsSearchFieldDidEndSearching() bool {
	return di._SearchFieldDidEndSearching != nil
}

func (di *SearchFieldDelegate) SetSearchFieldDidEndSearching(f func(sender SearchField)) {
	di._SearchFieldDidEndSearching = f
}

func (di *SearchFieldDelegate) SearchFieldDidEndSearching(sender SearchField) {
	di._SearchFieldDidEndSearching(sender)
}

type SearchFieldDelegateWrapper struct {
	TextFieldDelegateWrapper
}

func (s_ SearchFieldDelegateWrapper) ImplementsSearchFieldDidStartSearching() bool {
	return s_.RespondsToSelector(objc.GetSelector("searchFieldDidStartSearching:"))
}

func (s_ SearchFieldDelegateWrapper) SearchFieldDidStartSearching(sender ISearchField) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("searchFieldDidStartSearching:"), objc.ExtractPtr(sender))
}

func (s_ SearchFieldDelegateWrapper) ImplementsSearchFieldDidEndSearching() bool {
	return s_.RespondsToSelector(objc.GetSelector("searchFieldDidEndSearching:"))
}

func (s_ SearchFieldDelegateWrapper) SearchFieldDidEndSearching(sender ISearchField) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("searchFieldDidEndSearching:"), objc.ExtractPtr(sender))
}
