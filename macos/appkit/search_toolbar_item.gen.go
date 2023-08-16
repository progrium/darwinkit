// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SearchToolbarItem] class.
var SearchToolbarItemClass = _SearchToolbarItemClass{objc.GetClass("NSSearchToolbarItem")}

type _SearchToolbarItemClass struct {
	objc.Class
}

// An interface definition for the [SearchToolbarItem] class.
type ISearchToolbarItem interface {
	IToolbarItem
	EndSearchInteraction()
	BeginSearchInteraction()
	SearchField() SearchField
	SetSearchField(value ISearchField)
	ResignsFirstResponderWithCancel() bool
	SetResignsFirstResponderWithCancel(value bool)
	PreferredWidthForSearchField() float64
	SetPreferredWidthForSearchField(value float64)
}

// A toolbar item that contains a search field optimized for performing text-based searches. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchtoolbaritem?language=objc
type SearchToolbarItem struct {
	ToolbarItem
}

func SearchToolbarItemFrom(ptr unsafe.Pointer) SearchToolbarItem {
	return SearchToolbarItem{
		ToolbarItem: ToolbarItemFrom(ptr),
	}
}

func (sc _SearchToolbarItemClass) Alloc() SearchToolbarItem {
	rv := objc.Call[SearchToolbarItem](sc, objc.Sel("alloc"))
	return rv
}

func SearchToolbarItem_Alloc() SearchToolbarItem {
	return SearchToolbarItemClass.Alloc()
}

func (sc _SearchToolbarItemClass) New() SearchToolbarItem {
	rv := objc.Call[SearchToolbarItem](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSearchToolbarItem() SearchToolbarItem {
	return SearchToolbarItemClass.New()
}

func (s_ SearchToolbarItem) Init() SearchToolbarItem {
	rv := objc.Call[SearchToolbarItem](s_, objc.Sel("init"))
	return rv
}

func (s_ SearchToolbarItem) InitWithItemIdentifier(itemIdentifier ToolbarItemIdentifier) SearchToolbarItem {
	rv := objc.Call[SearchToolbarItem](s_, objc.Sel("initWithItemIdentifier:"), itemIdentifier)
	return rv
}

// Creates a toolbar item with the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/1534084-initwithitemidentifier?language=objc
func SearchToolbarItem_InitWithItemIdentifier(itemIdentifier ToolbarItemIdentifier) SearchToolbarItem {
	return SearchToolbarItemClass.Alloc().InitWithItemIdentifier(itemIdentifier)
}

// Ends a search interaction by giving up the first responder and adjusting the size of the search field to the available width for the toolbar item if necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchtoolbaritem/3674527-endsearchinteraction?language=objc
func (s_ SearchToolbarItem) EndSearchInteraction() {
	objc.Call[objc.Void](s_, objc.Sel("endSearchInteraction"))
}

// Starts a search interaction and moves the keyboard focus to the search field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchtoolbaritem/3674526-beginsearchinteraction?language=objc
func (s_ SearchToolbarItem) BeginSearchInteraction() {
	objc.Call[objc.Void](s_, objc.Sel("beginSearchInteraction"))
}

// The search field inside the toolbar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchtoolbaritem/3634330-searchfield?language=objc
func (s_ SearchToolbarItem) SearchField() SearchField {
	rv := objc.Call[SearchField](s_, objc.Sel("searchField"))
	return rv
}

// The search field inside the toolbar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchtoolbaritem/3634330-searchfield?language=objc
func (s_ SearchToolbarItem) SetSearchField(value ISearchField) {
	objc.Call[objc.Void](s_, objc.Sel("setSearchField:"), objc.Ptr(value))
}

// A Boolean value that enables the cancel button in the search field to resign the first responder in addition to clearing the contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchtoolbaritem/3634329-resignsfirstresponderwithcancel?language=objc
func (s_ SearchToolbarItem) ResignsFirstResponderWithCancel() bool {
	rv := objc.Call[bool](s_, objc.Sel("resignsFirstResponderWithCancel"))
	return rv
}

// A Boolean value that enables the cancel button in the search field to resign the first responder in addition to clearing the contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchtoolbaritem/3634329-resignsfirstresponderwithcancel?language=objc
func (s_ SearchToolbarItem) SetResignsFirstResponderWithCancel(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setResignsFirstResponderWithCancel:"), value)
}

// The preferred width for the toolbar item when it has keyboard focus. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchtoolbaritem/3674528-preferredwidthforsearchfield?language=objc
func (s_ SearchToolbarItem) PreferredWidthForSearchField() float64 {
	rv := objc.Call[float64](s_, objc.Sel("preferredWidthForSearchField"))
	return rv
}

// The preferred width for the toolbar item when it has keyboard focus. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchtoolbaritem/3674528-preferredwidthforsearchfield?language=objc
func (s_ SearchToolbarItem) SetPreferredWidthForSearchField(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setPreferredWidthForSearchField:"), value)
}
