// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SearchField] class.
var SearchFieldClass = _SearchFieldClass{objc.GetClass("NSSearchField")}

type _SearchFieldClass struct {
	objc.Class
}

// An interface definition for the [SearchField] class.
type ISearchField interface {
	ITextField
	RecentsAutosaveName() SearchFieldRecentsAutosaveName
	SetRecentsAutosaveName(value SearchFieldRecentsAutosaveName)
	MaximumRecents() int
	SetMaximumRecents(value int)
	CancelButtonBounds() foundation.Rect
	SendsWholeSearchString() bool
	SetSendsWholeSearchString(value bool)
	SearchMenuTemplate() Menu
	SetSearchMenuTemplate(value IMenu)
	SendsSearchStringImmediately() bool
	SetSendsSearchStringImmediately(value bool)
	RecentSearches() []string
	SetRecentSearches(value []string)
	SearchTextBounds() foundation.Rect
	SearchButtonBounds() foundation.Rect
}

// A text field optimized for performing text-based searches. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfield?language=objc
type SearchField struct {
	TextField
}

func SearchFieldFrom(ptr unsafe.Pointer) SearchField {
	return SearchField{
		TextField: TextFieldFrom(ptr),
	}
}

func (sc _SearchFieldClass) Alloc() SearchField {
	rv := objc.Call[SearchField](sc, objc.Sel("alloc"))
	return rv
}

func SearchField_Alloc() SearchField {
	return SearchFieldClass.Alloc()
}

func (sc _SearchFieldClass) New() SearchField {
	rv := objc.Call[SearchField](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSearchField() SearchField {
	return SearchFieldClass.New()
}

func (s_ SearchField) Init() SearchField {
	rv := objc.Call[SearchField](s_, objc.Sel("init"))
	return rv
}

func (sc _SearchFieldClass) LabelWithAttributedString(attributedStringValue foundation.IAttributedString) SearchField {
	rv := objc.Call[SearchField](sc, objc.Sel("labelWithAttributedString:"), objc.Ptr(attributedStringValue))
	return rv
}

// Creates a text field for use as a static label that displays styled text, doesn’t wrap, and doesn’t have selectable text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1644658-labelwithattributedstring?language=objc
func SearchField_LabelWithAttributedString(attributedStringValue foundation.IAttributedString) SearchField {
	return SearchFieldClass.LabelWithAttributedString(attributedStringValue)
}

func (sc _SearchFieldClass) LabelWithString(stringValue string) SearchField {
	rv := objc.Call[SearchField](sc, objc.Sel("labelWithString:"), stringValue)
	return rv
}

// Initializes a text field for use as a static label that uses the system default font, doesn’t wrap, and doesn’t have selectable text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1644377-labelwithstring?language=objc
func SearchField_LabelWithString(stringValue string) SearchField {
	return SearchFieldClass.LabelWithString(stringValue)
}

func (sc _SearchFieldClass) WrappingLabelWithString(stringValue string) SearchField {
	rv := objc.Call[SearchField](sc, objc.Sel("wrappingLabelWithString:"), stringValue)
	return rv
}

// Initializes a text field for use as a multiline static label with selectable text that uses the system default font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1644543-wrappinglabelwithstring?language=objc
func SearchField_WrappingLabelWithString(stringValue string) SearchField {
	return SearchFieldClass.WrappingLabelWithString(stringValue)
}

func (sc _SearchFieldClass) TextFieldWithString(stringValue string) SearchField {
	rv := objc.Call[SearchField](sc, objc.Sel("textFieldWithString:"), stringValue)
	return rv
}

// Initializes a single-line editable text field for user input using the system default font and standard visual appearance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1644706-textfieldwithstring?language=objc
func SearchField_TextFieldWithString(stringValue string) SearchField {
	return SearchFieldClass.TextFieldWithString(stringValue)
}

func (s_ SearchField) InitWithFrame(frameRect foundation.Rect) SearchField {
	rv := objc.Call[SearchField](s_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes a control with the specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428900-initwithframe?language=objc
func NewSearchFieldWithFrame(frameRect foundation.Rect) SearchField {
	instance := SearchFieldClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

// The name under which the search field automatically archives the list of recent search strings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfield/1530035-recentsautosavename?language=objc
func (s_ SearchField) RecentsAutosaveName() SearchFieldRecentsAutosaveName {
	rv := objc.Call[SearchFieldRecentsAutosaveName](s_, objc.Sel("recentsAutosaveName"))
	return rv
}

// The name under which the search field automatically archives the list of recent search strings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfield/1530035-recentsautosavename?language=objc
func (s_ SearchField) SetRecentsAutosaveName(value SearchFieldRecentsAutosaveName) {
	objc.Call[objc.Void](s_, objc.Sel("setRecentsAutosaveName:"), value)
}

// The maximum number of search strings that can appear in the search menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfield/1533938-maximumrecents?language=objc
func (s_ SearchField) MaximumRecents() int {
	rv := objc.Call[int](s_, objc.Sel("maximumRecents"))
	return rv
}

// The maximum number of search strings that can appear in the search menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfield/1533938-maximumrecents?language=objc
func (s_ SearchField) SetMaximumRecents(value int) {
	objc.Call[objc.Void](s_, objc.Sel("setMaximumRecents:"), value)
}

// The rectangle for the cancel button within the bounds of the search field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfield/3634323-cancelbuttonbounds?language=objc
func (s_ SearchField) CancelButtonBounds() foundation.Rect {
	rv := objc.Call[foundation.Rect](s_, objc.Sel("cancelButtonBounds"))
	return rv
}

// A Boolean value indicating whether the cell calls its search action method when the user clicks the search button or presses Return, or after each keystroke. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfield/1533976-sendswholesearchstring?language=objc
func (s_ SearchField) SendsWholeSearchString() bool {
	rv := objc.Call[bool](s_, objc.Sel("sendsWholeSearchString"))
	return rv
}

// A Boolean value indicating whether the cell calls its search action method when the user clicks the search button or presses Return, or after each keystroke. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfield/1533976-sendswholesearchstring?language=objc
func (s_ SearchField) SetSendsWholeSearchString(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setSendsWholeSearchString:"), value)
}

// The menu object used to dynamically construct the search field’s pop-up icon menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfield/1529467-searchmenutemplate?language=objc
func (s_ SearchField) SearchMenuTemplate() Menu {
	rv := objc.Call[Menu](s_, objc.Sel("searchMenuTemplate"))
	return rv
}

// The menu object used to dynamically construct the search field’s pop-up icon menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfield/1529467-searchmenutemplate?language=objc
func (s_ SearchField) SetSearchMenuTemplate(value IMenu) {
	objc.Call[objc.Void](s_, objc.Sel("setSearchMenuTemplate:"), objc.Ptr(value))
}

// A Boolean value indicating whether the cell calls its action method immediately when an appropriate action occurs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfield/1529081-sendssearchstringimmediately?language=objc
func (s_ SearchField) SendsSearchStringImmediately() bool {
	rv := objc.Call[bool](s_, objc.Sel("sendsSearchStringImmediately"))
	return rv
}

// A Boolean value indicating whether the cell calls its action method immediately when an appropriate action occurs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfield/1529081-sendssearchstringimmediately?language=objc
func (s_ SearchField) SetSendsSearchStringImmediately(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setSendsSearchStringImmediately:"), value)
}

// The list of recent search strings for the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfield/1531413-recentsearches?language=objc
func (s_ SearchField) RecentSearches() []string {
	rv := objc.Call[[]string](s_, objc.Sel("recentSearches"))
	return rv
}

// The list of recent search strings for the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfield/1531413-recentsearches?language=objc
func (s_ SearchField) SetRecentSearches(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setRecentSearches:"), value)
}

// The rectangle for the search text within the bounds of the search field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfield/3634325-searchtextbounds?language=objc
func (s_ SearchField) SearchTextBounds() foundation.Rect {
	rv := objc.Call[foundation.Rect](s_, objc.Sel("searchTextBounds"))
	return rv
}

// The rectangle for the search button within the bounds of the search field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssearchfield/3634324-searchbuttonbounds?language=objc
func (s_ SearchField) SearchButtonBounds() foundation.Rect {
	rv := objc.Call[foundation.Rect](s_, objc.Sel("searchButtonBounds"))
	return rv
}
