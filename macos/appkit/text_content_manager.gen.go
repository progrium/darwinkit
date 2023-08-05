// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var TextContentManagerClass = _TextContentManagerClass{objc.GetClass("NSTextContentManager")}

type _TextContentManagerClass struct {
	objc.Class
}

type ITextContentManager interface {
	objc.IObject
	PerformEditingTransactionUsingBlock(transaction func())
	RecordEditActionInRangeNewTextRange(originalTextRange ITextRange, newTextRange ITextRange)
	AddTextLayoutManager(textLayoutManager ITextLayoutManager)
	RemoveTextLayoutManager(textLayoutManager ITextLayoutManager)
	SynchronizeTextLayoutManagers(completionHandler func(error foundation.Error))
	TextElementsForRange(range_ ITextRange) []TextElement
	AutomaticallySynchronizesToBackingStore() bool
	SetAutomaticallySynchronizesToBackingStore(value bool)
	HasEditingTransaction() bool
	PrimaryTextLayoutManager() TextLayoutManager
	SetPrimaryTextLayoutManager(value ITextLayoutManager)
	TextLayoutManagers() []TextLayoutManager
	AutomaticallySynchronizesTextLayoutManagers() bool
	SetAutomaticallySynchronizesTextLayoutManagers(value bool)
}

type TextContentManager struct {
	objc.Object
}

func MakeTextContentManager(ptr unsafe.Pointer) TextContentManager {
	return TextContentManager{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TextContentManager) Init() TextContentManager {
	rv := objc.CallMethod[TextContentManager](t_, objc.GetSelector("init"))
	return rv
}

func TextContentManager_Init() TextContentManager {
	return TextContentManagerClass.Alloc().Init()
}

func (tc _TextContentManagerClass) Alloc() TextContentManager {
	rv := objc.CallMethod[TextContentManager](tc, objc.GetSelector("alloc"))
	return rv
}

func TextContentManager_Alloc() TextContentManager {
	return TextContentManagerClass.Alloc()
}

func (tc _TextContentManagerClass) New() TextContentManager {
	rv := objc.CallMethod[TextContentManager](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTextContentManager() TextContentManager {
	return TextContentManagerClass.New()
}

func TextContentManager_New() TextContentManager {
	return TextContentManagerClass.New()
}

func (t_ TextContentManager) PerformEditingTransactionUsingBlock(transaction func()) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("performEditingTransactionUsingBlock:"), transaction)
}

func (t_ TextContentManager) RecordEditActionInRangeNewTextRange(originalTextRange ITextRange, newTextRange ITextRange) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("recordEditActionInRange:newTextRange:"), objc.ExtractPtr(originalTextRange), objc.ExtractPtr(newTextRange))
}

func (t_ TextContentManager) AddTextLayoutManager(textLayoutManager ITextLayoutManager) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("addTextLayoutManager:"), objc.ExtractPtr(textLayoutManager))
}

func (t_ TextContentManager) RemoveTextLayoutManager(textLayoutManager ITextLayoutManager) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("removeTextLayoutManager:"), objc.ExtractPtr(textLayoutManager))
}

func (t_ TextContentManager) SynchronizeTextLayoutManagers(completionHandler func(error foundation.Error)) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("synchronizeTextLayoutManagers:"), completionHandler)
}

func (t_ TextContentManager) TextElementsForRange(range_ ITextRange) []TextElement {
	rv := objc.CallMethod[[]TextElement](t_, objc.GetSelector("textElementsForRange:"), objc.ExtractPtr(range_))
	return rv
}

func (t_ TextContentManager) AutomaticallySynchronizesToBackingStore() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("automaticallySynchronizesToBackingStore"))
	return rv
}

func (t_ TextContentManager) SetAutomaticallySynchronizesToBackingStore(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAutomaticallySynchronizesToBackingStore:"), value)
}

func (t_ TextContentManager) HasEditingTransaction() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("hasEditingTransaction"))
	return rv
}

func (t_ TextContentManager) PrimaryTextLayoutManager() TextLayoutManager {
	rv := objc.CallMethod[TextLayoutManager](t_, objc.GetSelector("primaryTextLayoutManager"))
	return rv
}

func (t_ TextContentManager) SetPrimaryTextLayoutManager(value ITextLayoutManager) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setPrimaryTextLayoutManager:"), objc.ExtractPtr(value))
}

func (t_ TextContentManager) TextLayoutManagers() []TextLayoutManager {
	rv := objc.CallMethod[[]TextLayoutManager](t_, objc.GetSelector("textLayoutManagers"))
	return rv
}

func (t_ TextContentManager) AutomaticallySynchronizesTextLayoutManagers() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("automaticallySynchronizesTextLayoutManagers"))
	return rv
}

func (t_ TextContentManager) SetAutomaticallySynchronizesTextLayoutManagers(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAutomaticallySynchronizesTextLayoutManagers:"), value)
}
