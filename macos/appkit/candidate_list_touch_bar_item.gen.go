// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CandidateListTouchBarItem] class.
var CandidateListTouchBarItemClass = _CandidateListTouchBarItemClass{objc.GetClass("NSCandidateListTouchBarItem")}

type _CandidateListTouchBarItemClass struct {
	objc.Class
}

// An interface definition for the [CandidateListTouchBarItem] class.
type ICandidateListTouchBarItem interface {
	ITouchBarItem
	SetCandidatesForSelectedRangeInString(candidates []objc.IObject, selectedRange foundation.Range, originalString string)
	UpdateWithInsertionPointVisibility(isVisible bool)
	Candidates() []objc.Object
	AllowsTextInputContextCandidates() bool
	SetAllowsTextInputContextCandidates(value bool)
	IsCandidateListVisible() bool
	SetCustomizationLabel(value string)
	AllowsCollapsing() bool
	SetAllowsCollapsing(value bool)
	IsCollapsed() bool
	SetCollapsed(value bool)
	Delegate() CandidateListTouchBarItemDelegateWrapper
	SetDelegate(value PCandidateListTouchBarItemDelegate)
	SetDelegateObject(valueObject objc.IObject)
	Client() View
	SetClient(value IView)
	AttributedStringForCandidate() func(candidate objc.Object, index int) foundation.AttributedString
	SetAttributedStringForCandidate(value func(candidate objc.Object, index int) foundation.AttributedString)
}

// A bar item that, along with its delegate, provides a list of textual suggestions for the current text view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem?language=objc
type CandidateListTouchBarItem struct {
	TouchBarItem
}

func CandidateListTouchBarItemFrom(ptr unsafe.Pointer) CandidateListTouchBarItem {
	return CandidateListTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}

func (cc _CandidateListTouchBarItemClass) Alloc() CandidateListTouchBarItem {
	rv := objc.Call[CandidateListTouchBarItem](cc, objc.Sel("alloc"))
	return rv
}

func CandidateListTouchBarItem_Alloc() CandidateListTouchBarItem {
	return CandidateListTouchBarItemClass.Alloc()
}

func (cc _CandidateListTouchBarItemClass) New() CandidateListTouchBarItem {
	rv := objc.Call[CandidateListTouchBarItem](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCandidateListTouchBarItem() CandidateListTouchBarItem {
	return CandidateListTouchBarItemClass.New()
}

func (c_ CandidateListTouchBarItem) Init() CandidateListTouchBarItem {
	rv := objc.Call[CandidateListTouchBarItem](c_, objc.Sel("init"))
	return rv
}

func (c_ CandidateListTouchBarItem) InitWithIdentifier(identifier TouchBarItemIdentifier) CandidateListTouchBarItem {
	rv := objc.Call[CandidateListTouchBarItem](c_, objc.Sel("initWithIdentifier:"), identifier)
	return rv
}

// Creates a new item with the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/2544755-initwithidentifier?language=objc
func NewCandidateListTouchBarItemWithIdentifier(identifier TouchBarItemIdentifier) CandidateListTouchBarItem {
	instance := CandidateListTouchBarItemClass.Alloc().InitWithIdentifier(identifier)
	instance.Autorelease()
	return instance
}

// Sets an array of candidate objects to be displayed in the candidate list bar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/2544701-setcandidates?language=objc
func (c_ CandidateListTouchBarItem) SetCandidatesForSelectedRangeInString(candidates []objc.IObject, selectedRange foundation.Range, originalString string) {
	objc.Call[objc.Void](c_, objc.Sel("setCandidates:forSelectedRange:inString:"), candidates, selectedRange, originalString)
}

// Updates the candidate list visibility configuration based on the client's insertion point state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/2544658-updatewithinsertionpointvisibili?language=objc
func (c_ CandidateListTouchBarItem) UpdateWithInsertionPointVisibility(isVisible bool) {
	objc.Call[objc.Void](c_, objc.Sel("updateWithInsertionPointVisibility:"), isVisible)
}

// The array of candidate objects previously set by setCandidates:forSelectedRange:inString:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/2544728-candidates?language=objc
func (c_ CandidateListTouchBarItem) Candidates() []objc.Object {
	rv := objc.Call[[]objc.Object](c_, objc.Sel("candidates"))
	return rv
}

// A Boolean value that specifies whether a candidate list item displays candidates from text input providers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/2544668-allowstextinputcontextcandidates?language=objc
func (c_ CandidateListTouchBarItem) AllowsTextInputContextCandidates() bool {
	rv := objc.Call[bool](c_, objc.Sel("allowsTextInputContextCandidates"))
	return rv
}

// A Boolean value that specifies whether a candidate list item displays candidates from text input providers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/2544668-allowstextinputcontextcandidates?language=objc
func (c_ CandidateListTouchBarItem) SetAllowsTextInputContextCandidates(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setAllowsTextInputContextCandidates:"), value)
}

// A Boolean value that represents the visibility of this item's candidate list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/2544870-candidatelistvisible?language=objc
func (c_ CandidateListTouchBarItem) IsCandidateListVisible() bool {
	rv := objc.Call[bool](c_, objc.Sel("isCandidateListVisible"))
	return rv
}

// The user-visible string identifying this item during bar customization. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/2544821-customizationlabel?language=objc
func (c_ CandidateListTouchBarItem) SetCustomizationLabel(value string) {
	objc.Call[objc.Void](c_, objc.Sel("setCustomizationLabel:"), value)
}

// A Boolean value that specifies whether the item can be collapsed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/2544718-allowscollapsing?language=objc
func (c_ CandidateListTouchBarItem) AllowsCollapsing() bool {
	rv := objc.Call[bool](c_, objc.Sel("allowsCollapsing"))
	return rv
}

// A Boolean value that specifies whether the item can be collapsed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/2544718-allowscollapsing?language=objc
func (c_ CandidateListTouchBarItem) SetAllowsCollapsing(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setAllowsCollapsing:"), value)
}

// A Boolean value that controls the visibility of the candidate list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/2544733-collapsed?language=objc
func (c_ CandidateListTouchBarItem) IsCollapsed() bool {
	rv := objc.Call[bool](c_, objc.Sel("isCollapsed"))
	return rv
}

// A Boolean value that controls the visibility of the candidate list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/2544733-collapsed?language=objc
func (c_ CandidateListTouchBarItem) SetCollapsed(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setCollapsed:"), value)
}

// The delegate of the candidate list item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/2544820-delegate?language=objc
func (c_ CandidateListTouchBarItem) Delegate() CandidateListTouchBarItemDelegateWrapper {
	rv := objc.Call[CandidateListTouchBarItemDelegateWrapper](c_, objc.Sel("delegate"))
	return rv
}

// The delegate of the candidate list item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/2544820-delegate?language=objc
func (c_ CandidateListTouchBarItem) SetDelegate(value PCandidateListTouchBarItemDelegate) {
	po0 := objc.WrapAsProtocol("NSCandidateListTouchBarItemDelegate", value)
	objc.SetAssociatedObject(c_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](c_, objc.Sel("setDelegate:"), po0)
}

// The delegate of the candidate list item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/2544820-delegate?language=objc
func (c_ CandidateListTouchBarItem) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The client object for the candidate list item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/2544873-client?language=objc
func (c_ CandidateListTouchBarItem) Client() View {
	rv := objc.Call[View](c_, objc.Sel("client"))
	return rv
}

// The client object for the candidate list item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/2544873-client?language=objc
func (c_ CandidateListTouchBarItem) SetClient(value IView) {
	objc.Call[objc.Void](c_, objc.Sel("setClient:"), objc.Ptr(value))
}

// A block that converts a candidate object into an attributed string for display in the candidate list item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/2544823-attributedstringforcandidate?language=objc
func (c_ CandidateListTouchBarItem) AttributedStringForCandidate() func(candidate objc.Object, index int) foundation.AttributedString {
	rv := objc.Call[func(candidate objc.Object, index int) foundation.AttributedString](c_, objc.Sel("attributedStringForCandidate"))
	return rv
}

// A block that converts a candidate object into an attributed string for display in the candidate list item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/2544823-attributedstringforcandidate?language=objc
func (c_ CandidateListTouchBarItem) SetAttributedStringForCandidate(value func(candidate objc.Object, index int) foundation.AttributedString) {
	objc.Call[objc.Void](c_, objc.Sel("setAttributedStringForCandidate:"), value)
}
