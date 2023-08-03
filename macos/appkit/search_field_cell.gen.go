// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var SearchFieldCellClass = _SearchFieldCellClass{objc.GetClass("NSSearchFieldCell")}

type _SearchFieldCellClass struct {
	objc.Class
}

type ISearchFieldCell interface {
	ITextFieldCell
	ResetSearchButtonCell()
	ResetCancelButtonCell()
	SearchTextRectForBounds(rect foundation.Rect) foundation.Rect
	SearchButtonRectForBounds(rect foundation.Rect) foundation.Rect
	CancelButtonRectForBounds(rect foundation.Rect) foundation.Rect
	SearchButtonCell() ButtonCell
	SetSearchButtonCell(value IButtonCell)
	CancelButtonCell() ButtonCell
	SetCancelButtonCell(value IButtonCell)
	SearchMenuTemplate() Menu
	SetSearchMenuTemplate(value IMenu)
	SendsWholeSearchString() bool
	SetSendsWholeSearchString(value bool)
	SendsSearchStringImmediately() bool
	SetSendsSearchStringImmediately(value bool)
	MaximumRecents() int
	SetMaximumRecents(value int)
	RecentSearches() []string
	SetRecentSearches(value []string)
	RecentsAutosaveName() SearchFieldRecentsAutosaveName
	SetRecentsAutosaveName(value SearchFieldRecentsAutosaveName)
}

type SearchFieldCell struct {
	TextFieldCell
}

func MakeSearchFieldCell(ptr unsafe.Pointer) SearchFieldCell {
	return SearchFieldCell{
		TextFieldCell: MakeTextFieldCell(ptr),
	}
}

func (s_ SearchFieldCell) InitTextCell(string_ string) SearchFieldCell {
	rv := objc.CallMethod[SearchFieldCell](s_, objc.GetSelector("initTextCell:"), string_)
	return rv
}

func SearchFieldCell_InitTextCell(string_ string) SearchFieldCell {
	return SearchFieldCellClass.Alloc().InitTextCell(string_)
}

func (s_ SearchFieldCell) InitImageCell(image IImage) SearchFieldCell {
	rv := objc.CallMethod[SearchFieldCell](s_, objc.GetSelector("initImageCell:"), objc.ExtractPtr(image))
	return rv
}

func SearchFieldCell_InitImageCell(image IImage) SearchFieldCell {
	return SearchFieldCellClass.Alloc().InitImageCell(image)
}

func (s_ SearchFieldCell) Init() SearchFieldCell {
	rv := objc.CallMethod[SearchFieldCell](s_, objc.GetSelector("init"))
	return rv
}

func SearchFieldCell_Init() SearchFieldCell {
	return SearchFieldCellClass.Alloc().Init()
}

func (sc _SearchFieldCellClass) Alloc() SearchFieldCell {
	rv := objc.CallMethod[SearchFieldCell](sc, objc.GetSelector("alloc"))
	return rv
}

func SearchFieldCell_Alloc() SearchFieldCell {
	return SearchFieldCellClass.Alloc()
}

func (sc _SearchFieldCellClass) New() SearchFieldCell {
	rv := objc.CallMethod[SearchFieldCell](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewSearchFieldCell() SearchFieldCell {
	return SearchFieldCellClass.New()
}

func SearchFieldCell_New() SearchFieldCell {
	return SearchFieldCellClass.New()
}

func (s_ SearchFieldCell) ResetSearchButtonCell() {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("resetSearchButtonCell"))
}

func (s_ SearchFieldCell) ResetCancelButtonCell() {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("resetCancelButtonCell"))
}

func (s_ SearchFieldCell) SearchTextRectForBounds(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.GetSelector("searchTextRectForBounds:"), rect)
	return rv
}

func (s_ SearchFieldCell) SearchButtonRectForBounds(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.GetSelector("searchButtonRectForBounds:"), rect)
	return rv
}

func (s_ SearchFieldCell) CancelButtonRectForBounds(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.GetSelector("cancelButtonRectForBounds:"), rect)
	return rv
}

func (s_ SearchFieldCell) SearchButtonCell() ButtonCell {
	rv := objc.CallMethod[ButtonCell](s_, objc.GetSelector("searchButtonCell"))
	return rv
}

func (s_ SearchFieldCell) SetSearchButtonCell(value IButtonCell) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setSearchButtonCell:"), objc.ExtractPtr(value))
}

func (s_ SearchFieldCell) CancelButtonCell() ButtonCell {
	rv := objc.CallMethod[ButtonCell](s_, objc.GetSelector("cancelButtonCell"))
	return rv
}

func (s_ SearchFieldCell) SetCancelButtonCell(value IButtonCell) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setCancelButtonCell:"), objc.ExtractPtr(value))
}

func (s_ SearchFieldCell) SearchMenuTemplate() Menu {
	rv := objc.CallMethod[Menu](s_, objc.GetSelector("searchMenuTemplate"))
	return rv
}

func (s_ SearchFieldCell) SetSearchMenuTemplate(value IMenu) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setSearchMenuTemplate:"), objc.ExtractPtr(value))
}

func (s_ SearchFieldCell) SendsWholeSearchString() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("sendsWholeSearchString"))
	return rv
}

func (s_ SearchFieldCell) SetSendsWholeSearchString(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setSendsWholeSearchString:"), value)
}

func (s_ SearchFieldCell) SendsSearchStringImmediately() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("sendsSearchStringImmediately"))
	return rv
}

func (s_ SearchFieldCell) SetSendsSearchStringImmediately(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setSendsSearchStringImmediately:"), value)
}

func (s_ SearchFieldCell) MaximumRecents() int {
	rv := objc.CallMethod[int](s_, objc.GetSelector("maximumRecents"))
	return rv
}

func (s_ SearchFieldCell) SetMaximumRecents(value int) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setMaximumRecents:"), value)
}

func (s_ SearchFieldCell) RecentSearches() []string {
	rv := objc.CallMethod[[]string](s_, objc.GetSelector("recentSearches"))
	return rv
}

func (s_ SearchFieldCell) SetRecentSearches(value []string) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setRecentSearches:"), value)
}

func (s_ SearchFieldCell) RecentsAutosaveName() SearchFieldRecentsAutosaveName {
	rv := objc.CallMethod[SearchFieldRecentsAutosaveName](s_, objc.GetSelector("recentsAutosaveName"))
	return rv
}

func (s_ SearchFieldCell) SetRecentsAutosaveName(value SearchFieldRecentsAutosaveName) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setRecentsAutosaveName:"), value)
}
