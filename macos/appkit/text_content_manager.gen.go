// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TextContentManager] class.
var TextContentManagerClass = _TextContentManagerClass{objc.GetClass("NSTextContentManager")}

type _TextContentManagerClass struct {
	objc.Class
}

// An interface definition for the [TextContentManager] class.
type ITextContentManager interface {
	objc.IObject
	RemoveTextLayoutManager(textLayoutManager ITextLayoutManager)
	PerformEditingTransactionUsingBlock(transaction func())
	SynchronizeTextLayoutManagers(completionHandler func(error foundation.Error))
	RecordEditActionInRangeNewTextRange(originalTextRange ITextRange, newTextRange ITextRange)
	TextElementsForRange(range_ ITextRange) []TextElement
	AddTextLayoutManager(textLayoutManager ITextLayoutManager)
	AutomaticallySynchronizesToBackingStore() bool
	SetAutomaticallySynchronizesToBackingStore(value bool)
	AutomaticallySynchronizesTextLayoutManagers() bool
	SetAutomaticallySynchronizesTextLayoutManagers(value bool)
	Delegate() TextContentManagerDelegateWrapper
	SetDelegate(value PTextContentManagerDelegate)
	SetDelegateObject(valueObject objc.IObject)
	PrimaryTextLayoutManager() TextLayoutManager
	SetPrimaryTextLayoutManager(value ITextLayoutManager)
	TextLayoutManagers() []TextLayoutManager
	HasEditingTransaction() bool
}

// An abstract class that defines the interface and a default implementation for managing the text document contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentmanager?language=objc
type TextContentManager struct {
	objc.Object
}

func TextContentManagerFrom(ptr unsafe.Pointer) TextContentManager {
	return TextContentManager{
		Object: objc.ObjectFrom(ptr),
	}
}

func (t_ TextContentManager) Init() TextContentManager {
	rv := objc.Call[TextContentManager](t_, objc.Sel("init"))
	return rv
}

func (tc _TextContentManagerClass) Alloc() TextContentManager {
	rv := objc.Call[TextContentManager](tc, objc.Sel("alloc"))
	return rv
}

func TextContentManager_Alloc() TextContentManager {
	return TextContentManagerClass.Alloc()
}

func (tc _TextContentManagerClass) New() TextContentManager {
	rv := objc.Call[TextContentManager](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTextContentManager() TextContentManager {
	return TextContentManagerClass.New()
}

// Removes the layout manager you specifiy from the list of layout managers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentmanager/3809924-removetextlayoutmanager?language=objc
func (t_ TextContentManager) RemoveTextLayoutManager(textLayoutManager ITextLayoutManager) {
	objc.Call[objc.Void](t_, objc.Sel("removeTextLayoutManager:"), objc.Ptr(textLayoutManager))
}

// Performs an editing transaction and invokes a block upon completion. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentmanager/3809921-performeditingtransactionusingbl?language=objc
func (t_ TextContentManager) PerformEditingTransactionUsingBlock(transaction func()) {
	objc.Call[objc.Void](t_, objc.Sel("performEditingTransactionUsingBlock:"), transaction)
}

// Synchronizes changes to all nonprimary text layout managers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentmanager/3809925-synchronizetextlayoutmanagers?language=objc
func (t_ TextContentManager) SynchronizeTextLayoutManagers(completionHandler func(error foundation.Error)) {
	objc.Call[objc.Void](t_, objc.Sel("synchronizeTextLayoutManagers:"), completionHandler)
}

// Records information about an edit action to the transaction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentmanager/3809923-recordeditactioninrange?language=objc
func (t_ TextContentManager) RecordEditActionInRangeNewTextRange(originalTextRange ITextRange, newTextRange ITextRange) {
	objc.Call[objc.Void](t_, objc.Sel("recordEditActionInRange:newTextRange:"), objc.Ptr(originalTextRange), objc.Ptr(newTextRange))
}

// Returns an array of text elements that intersect with the range you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentmanager/3809928-textelementsforrange?language=objc
func (t_ TextContentManager) TextElementsForRange(range_ ITextRange) []TextElement {
	rv := objc.Call[[]TextElement](t_, objc.Sel("textElementsForRange:"), objc.Ptr(range_))
	return rv
}

// Adds the layout manager you provide to the list of layout managers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentmanager/3809916-addtextlayoutmanager?language=objc
func (t_ TextContentManager) AddTextLayoutManager(textLayoutManager ITextLayoutManager) {
	objc.Call[objc.Void](t_, objc.Sel("addTextLayoutManager:"), objc.Ptr(textLayoutManager))
}

// Determines whether to automatically synchronize with the backing store when an editing transaction finishes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentmanager/3852565-automaticallysynchronizestobacki?language=objc
func (t_ TextContentManager) AutomaticallySynchronizesToBackingStore() bool {
	rv := objc.Call[bool](t_, objc.Sel("automaticallySynchronizesToBackingStore"))
	return rv
}

// Determines whether to automatically synchronize with the backing store when an editing transaction finishes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentmanager/3852565-automaticallysynchronizestobacki?language=objc
func (t_ TextContentManager) SetAutomaticallySynchronizesToBackingStore(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAutomaticallySynchronizesToBackingStore:"), value)
}

// Determines if the framework should automatically synchronize all text layout managers when exiting an editing transaction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentmanager/3852564-automaticallysynchronizestextlay?language=objc
func (t_ TextContentManager) AutomaticallySynchronizesTextLayoutManagers() bool {
	rv := objc.Call[bool](t_, objc.Sel("automaticallySynchronizesTextLayoutManagers"))
	return rv
}

// Determines if the framework should automatically synchronize all text layout managers when exiting an editing transaction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentmanager/3852564-automaticallysynchronizestextlay?language=objc
func (t_ TextContentManager) SetAutomaticallySynchronizesTextLayoutManagers(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAutomaticallySynchronizesTextLayoutManagers:"), value)
}

// The delegate for the content manager object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentmanager/3809917-delegate?language=objc
func (t_ TextContentManager) Delegate() TextContentManagerDelegateWrapper {
	rv := objc.Call[TextContentManagerDelegateWrapper](t_, objc.Sel("delegate"))
	return rv
}

// The delegate for the content manager object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentmanager/3809917-delegate?language=objc
func (t_ TextContentManager) SetDelegate(value PTextContentManagerDelegate) {
	po0 := objc.WrapAsProtocol("NSTextContentManagerDelegate", value)
	objc.SetAssociatedObject(t_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](t_, objc.Sel("setDelegate:"), po0)
}

// The delegate for the content manager object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentmanager/3809917-delegate?language=objc
func (t_ TextContentManager) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The primary text layout manager for this content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentmanager/3809922-primarytextlayoutmanager?language=objc
func (t_ TextContentManager) PrimaryTextLayoutManager() TextLayoutManager {
	rv := objc.Call[TextLayoutManager](t_, objc.Sel("primaryTextLayoutManager"))
	return rv
}

// The primary text layout manager for this content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentmanager/3809922-primarytextlayoutmanager?language=objc
func (t_ TextContentManager) SetPrimaryTextLayoutManager(value ITextLayoutManager) {
	objc.Call[objc.Void](t_, objc.Sel("setPrimaryTextLayoutManager:"), objc.Ptr(value))
}

// The array of text layout managers associated with this text content manager. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentmanager/3809929-textlayoutmanagers?language=objc
func (t_ TextContentManager) TextLayoutManagers() []TextLayoutManager {
	rv := objc.Call[[]TextLayoutManager](t_, objc.Sel("textLayoutManagers"))
	return rv
}

// Indicates there’s an active editing transaction from the primary text layout manager. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentmanager/3809918-haseditingtransaction?language=objc
func (t_ TextContentManager) HasEditingTransaction() bool {
	rv := objc.Call[bool](t_, objc.Sel("hasEditingTransaction"))
	return rv
}
