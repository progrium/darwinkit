// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SearchFieldCell] class.
var SearchFieldCellClass = _SearchFieldCellClass{objc.GetClass("NSSearchFieldCell")}

type _SearchFieldCellClass struct {
	objc.Class
}

// An interface definition for the [SearchFieldCell] class.
type ISearchFieldCell interface {
	ITextFieldCell
	SearchTextRectForBounds(rect foundation.Rect) foundation.Rect
	SearchButtonRectForBounds(rect foundation.Rect) foundation.Rect
	ResetSearchButtonCell()
	CancelButtonRectForBounds(rect foundation.Rect) foundation.Rect
	ResetCancelButtonCell()
	RecentsAutosaveName() SearchFieldRecentsAutosaveName
	SetRecentsAutosaveName(value SearchFieldRecentsAutosaveName)
	MaximumRecents() int
	SetMaximumRecents(value int)
	SearchButtonCell() ButtonCell
	SetSearchButtonCell(value IButtonCell)
	SendsWholeSearchString() bool
	SetSendsWholeSearchString(value bool)
	SearchMenuTemplate() Menu
	SetSearchMenuTemplate(value IMenu)
	CancelButtonCell() ButtonCell
	SetCancelButtonCell(value IButtonCell)
	SendsSearchStringImmediately() bool
	SetSendsSearchStringImmediately(value bool)
	RecentSearches() []string
	SetRecentSearches(value []string)
}

// The programmatic interface for text fields that are used for text-based searches. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfieldcell?language=objc
type SearchFieldCell struct {
	TextFieldCell
}

func SearchFieldCellFrom(ptr unsafe.Pointer) SearchFieldCell {
	return SearchFieldCell{
		TextFieldCell: TextFieldCellFrom(ptr),
	}
}

func (s_ SearchFieldCell) InitTextCell(string_ string) SearchFieldCell {
	rv := objc.Call[SearchFieldCell](s_, objc.Sel("initTextCell:"), string_)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfieldcell/1643722-inittextcell?language=objc
func SearchFieldCell_InitTextCell(string_ string) SearchFieldCell {
	return SearchFieldCellClass.Alloc().InitTextCell(string_)
}

func (sc _SearchFieldCellClass) Alloc() SearchFieldCell {
	rv := objc.Call[SearchFieldCell](sc, objc.Sel("alloc"))
	return rv
}

func SearchFieldCell_Alloc() SearchFieldCell {
	return SearchFieldCellClass.Alloc()
}

func (sc _SearchFieldCellClass) New() SearchFieldCell {
	rv := objc.Call[SearchFieldCell](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSearchFieldCell() SearchFieldCell {
	return SearchFieldCellClass.New()
}

func (s_ SearchFieldCell) Init() SearchFieldCell {
	rv := objc.Call[SearchFieldCell](s_, objc.Sel("init"))
	return rv
}

func (s_ SearchFieldCell) InitImageCell(image IImage) SearchFieldCell {
	rv := objc.Call[SearchFieldCell](s_, objc.Sel("initImageCell:"), objc.Ptr(image))
	return rv
}

// Returns an NSCell object initialized with the specified image and set to have the cell’s default menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1533898-initimagecell?language=objc
func SearchFieldCell_InitImageCell(image IImage) SearchFieldCell {
	return SearchFieldCellClass.Alloc().InitImageCell(image)
}

// Modifies the bounding rectangle for the search-text field cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfieldcell/1399455-searchtextrectforbounds?language=objc
func (s_ SearchFieldCell) SearchTextRectForBounds(rect foundation.Rect) foundation.Rect {
	rv := objc.Call[foundation.Rect](s_, objc.Sel("searchTextRectForBounds:"), rect)
	return rv
}

// Modifies the bounding rectangle for the search button cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfieldcell/1399450-searchbuttonrectforbounds?language=objc
func (s_ SearchFieldCell) SearchButtonRectForBounds(rect foundation.Rect) foundation.Rect {
	rv := objc.Call[foundation.Rect](s_, objc.Sel("searchButtonRectForBounds:"), rect)
	return rv
}

// Resets the search button cell to its default attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfieldcell/1399461-resetsearchbuttoncell?language=objc
func (s_ SearchFieldCell) ResetSearchButtonCell() {
	objc.Call[objc.Void](s_, objc.Sel("resetSearchButtonCell"))
}

// Modifies the bounding rectangle for the cancel button cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfieldcell/1399472-cancelbuttonrectforbounds?language=objc
func (s_ SearchFieldCell) CancelButtonRectForBounds(rect foundation.Rect) foundation.Rect {
	rv := objc.Call[foundation.Rect](s_, objc.Sel("cancelButtonRectForBounds:"), rect)
	return rv
}

// Resets the cancel button cell to its default attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfieldcell/1399442-resetcancelbuttoncell?language=objc
func (s_ SearchFieldCell) ResetCancelButtonCell() {
	objc.Call[objc.Void](s_, objc.Sel("resetCancelButtonCell"))
}

// The autosave name under which the search field automatically saves the list of recent search strings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfieldcell/1399463-recentsautosavename?language=objc
func (s_ SearchFieldCell) RecentsAutosaveName() SearchFieldRecentsAutosaveName {
	rv := objc.Call[SearchFieldRecentsAutosaveName](s_, objc.Sel("recentsAutosaveName"))
	return rv
}

// The autosave name under which the search field automatically saves the list of recent search strings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfieldcell/1399463-recentsautosavename?language=objc
func (s_ SearchFieldCell) SetRecentsAutosaveName(value SearchFieldRecentsAutosaveName) {
	objc.Call[objc.Void](s_, objc.Sel("setRecentsAutosaveName:"), value)
}

// The maximum number of search strings that can appear in the search menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfieldcell/1399468-maximumrecents?language=objc
func (s_ SearchFieldCell) MaximumRecents() int {
	rv := objc.Call[int](s_, objc.Sel("maximumRecents"))
	return rv
}

// The maximum number of search strings that can appear in the search menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfieldcell/1399468-maximumrecents?language=objc
func (s_ SearchFieldCell) SetMaximumRecents(value int) {
	objc.Call[objc.Void](s_, objc.Sel("setMaximumRecents:"), value)
}

// The button cell used to display the search-button image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfieldcell/1399457-searchbuttoncell?language=objc
func (s_ SearchFieldCell) SearchButtonCell() ButtonCell {
	rv := objc.Call[ButtonCell](s_, objc.Sel("searchButtonCell"))
	return rv
}

// The button cell used to display the search-button image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfieldcell/1399457-searchbuttoncell?language=objc
func (s_ SearchFieldCell) SetSearchButtonCell(value IButtonCell) {
	objc.Call[objc.Void](s_, objc.Sel("setSearchButtonCell:"), objc.Ptr(value))
}

// A Boolean value indicating whether the cell calls its search action method when the user clicks the search button (or presses Return) or after each keystroke. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfieldcell/1399453-sendswholesearchstring?language=objc
func (s_ SearchFieldCell) SendsWholeSearchString() bool {
	rv := objc.Call[bool](s_, objc.Sel("sendsWholeSearchString"))
	return rv
}

// A Boolean value indicating whether the cell calls its search action method when the user clicks the search button (or presses Return) or after each keystroke. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfieldcell/1399453-sendswholesearchstring?language=objc
func (s_ SearchFieldCell) SetSendsWholeSearchString(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setSendsWholeSearchString:"), value)
}

// The menu object used to dynamically construct the search field’s pop-up icon menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfieldcell/1399452-searchmenutemplate?language=objc
func (s_ SearchFieldCell) SearchMenuTemplate() Menu {
	rv := objc.Call[Menu](s_, objc.Sel("searchMenuTemplate"))
	return rv
}

// The menu object used to dynamically construct the search field’s pop-up icon menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfieldcell/1399452-searchmenutemplate?language=objc
func (s_ SearchFieldCell) SetSearchMenuTemplate(value IMenu) {
	objc.Call[objc.Void](s_, objc.Sel("setSearchMenuTemplate:"), objc.Ptr(value))
}

// The button cell used to display the cancel-button image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfieldcell/1399440-cancelbuttoncell?language=objc
func (s_ SearchFieldCell) CancelButtonCell() ButtonCell {
	rv := objc.Call[ButtonCell](s_, objc.Sel("cancelButtonCell"))
	return rv
}

// The button cell used to display the cancel-button image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfieldcell/1399440-cancelbuttoncell?language=objc
func (s_ SearchFieldCell) SetCancelButtonCell(value IButtonCell) {
	objc.Call[objc.Void](s_, objc.Sel("setCancelButtonCell:"), objc.Ptr(value))
}

// A Boolean value indicating whether the cell calls its action method immediately when an appropriate action occurs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfieldcell/1399444-sendssearchstringimmediately?language=objc
func (s_ SearchFieldCell) SendsSearchStringImmediately() bool {
	rv := objc.Call[bool](s_, objc.Sel("sendsSearchStringImmediately"))
	return rv
}

// A Boolean value indicating whether the cell calls its action method immediately when an appropriate action occurs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfieldcell/1399444-sendssearchstringimmediately?language=objc
func (s_ SearchFieldCell) SetSendsSearchStringImmediately(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setSendsSearchStringImmediately:"), value)
}

// An array of the recent search strings to display in the pop-up icon menu of the search field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfieldcell/1399446-recentsearches?language=objc
func (s_ SearchFieldCell) RecentSearches() []string {
	rv := objc.Call[[]string](s_, objc.Sel("recentSearches"))
	return rv
}

// An array of the recent search strings to display in the pop-up icon menu of the search field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfieldcell/1399446-recentsearches?language=objc
func (s_ SearchFieldCell) SetRecentSearches(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setRecentSearches:"), value)
}
