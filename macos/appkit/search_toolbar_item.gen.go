// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var SearchToolbarItemClass = _SearchToolbarItemClass{objc.GetClass("NSSearchToolbarItem")}

type _SearchToolbarItemClass struct {
	objc.Class
}

type ISearchToolbarItem interface {
	IToolbarItem
	BeginSearchInteraction()
	EndSearchInteraction()
	PreferredWidthForSearchField() float64
	SetPreferredWidthForSearchField(value float64)
	ResignsFirstResponderWithCancel() bool
	SetResignsFirstResponderWithCancel(value bool)
	SearchField() SearchField
	SetSearchField(value ISearchField)
}

type SearchToolbarItem struct {
	ToolbarItem
}

func MakeSearchToolbarItem(ptr unsafe.Pointer) SearchToolbarItem {
	return SearchToolbarItem{
		ToolbarItem: MakeToolbarItem(ptr),
	}
}

func (s_ SearchToolbarItem) InitWithItemIdentifier(itemIdentifier ToolbarItemIdentifier) SearchToolbarItem {
	rv := objc.CallMethod[SearchToolbarItem](s_, objc.GetSelector("initWithItemIdentifier:"), itemIdentifier)
	return rv
}

func SearchToolbarItem_InitWithItemIdentifier(itemIdentifier ToolbarItemIdentifier) SearchToolbarItem {
	return SearchToolbarItemClass.Alloc().InitWithItemIdentifier(itemIdentifier)
}

func (sc _SearchToolbarItemClass) Alloc() SearchToolbarItem {
	rv := objc.CallMethod[SearchToolbarItem](sc, objc.GetSelector("alloc"))
	return rv
}

func SearchToolbarItem_Alloc() SearchToolbarItem {
	return SearchToolbarItemClass.Alloc()
}

func (sc _SearchToolbarItemClass) New() SearchToolbarItem {
	rv := objc.CallMethod[SearchToolbarItem](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewSearchToolbarItem() SearchToolbarItem {
	return SearchToolbarItemClass.New()
}

func SearchToolbarItem_New() SearchToolbarItem {
	return SearchToolbarItemClass.New()
}

func (s_ SearchToolbarItem) Init() SearchToolbarItem {
	rv := objc.CallMethod[SearchToolbarItem](s_, objc.GetSelector("init"))
	return rv
}

func SearchToolbarItem_Init() SearchToolbarItem {
	return SearchToolbarItemClass.Alloc().Init()
}

func (s_ SearchToolbarItem) BeginSearchInteraction() {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("beginSearchInteraction"))
}

func (s_ SearchToolbarItem) EndSearchInteraction() {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("endSearchInteraction"))
}

func (s_ SearchToolbarItem) PreferredWidthForSearchField() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("preferredWidthForSearchField"))
	return rv
}

func (s_ SearchToolbarItem) SetPreferredWidthForSearchField(value float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setPreferredWidthForSearchField:"), value)
}

func (s_ SearchToolbarItem) ResignsFirstResponderWithCancel() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("resignsFirstResponderWithCancel"))
	return rv
}

func (s_ SearchToolbarItem) SetResignsFirstResponderWithCancel(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setResignsFirstResponderWithCancel:"), value)
}

func (s_ SearchToolbarItem) SearchField() SearchField {
	rv := objc.CallMethod[SearchField](s_, objc.GetSelector("searchField"))
	return rv
}

func (s_ SearchToolbarItem) SetSearchField(value ISearchField) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setSearchField:"), objc.ExtractPtr(value))
}
