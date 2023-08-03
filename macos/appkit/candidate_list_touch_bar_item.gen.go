// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var CandidateListTouchBarItemClass = _CandidateListTouchBarItemClass{objc.GetClass("NSCandidateListTouchBarItem")}

type _CandidateListTouchBarItemClass struct {
	objc.Class
}

type ICandidateListTouchBarItem interface {
	ITouchBarItem
	UpdateWithInsertionPointVisibility(isVisible bool)
	Client() View
	SetClient(value IView)
	Delegate() CandidateListTouchBarItemDelegateWrapper
	SetDelegate(value ICandidateListTouchBarItemDelegate)
	SetDelegate0(value objc.IObject)
	AllowsTextInputContextCandidates() bool
	SetAllowsTextInputContextCandidates(value bool)
	AllowsCollapsing() bool
	SetAllowsCollapsing(value bool)
	IsCollapsed() bool
	SetCollapsed(value bool)
	IsCandidateListVisible() bool
	SetCustomizationLabel(value string)
}

type CandidateListTouchBarItem struct {
	TouchBarItem
}

func MakeCandidateListTouchBarItem(ptr unsafe.Pointer) CandidateListTouchBarItem {
	return CandidateListTouchBarItem{
		TouchBarItem: MakeTouchBarItem(ptr),
	}
}

func (c_ CandidateListTouchBarItem) InitWithIdentifier(identifier TouchBarItemIdentifier) CandidateListTouchBarItem {
	rv := objc.CallMethod[CandidateListTouchBarItem](c_, objc.GetSelector("initWithIdentifier:"), identifier)
	return rv
}

func CandidateListTouchBarItem_InitWithIdentifier(identifier TouchBarItemIdentifier) CandidateListTouchBarItem {
	return CandidateListTouchBarItemClass.Alloc().InitWithIdentifier(identifier)
}

func (cc _CandidateListTouchBarItemClass) Alloc() CandidateListTouchBarItem {
	rv := objc.CallMethod[CandidateListTouchBarItem](cc, objc.GetSelector("alloc"))
	return rv
}

func CandidateListTouchBarItem_Alloc() CandidateListTouchBarItem {
	return CandidateListTouchBarItemClass.Alloc()
}

func (cc _CandidateListTouchBarItemClass) New() CandidateListTouchBarItem {
	rv := objc.CallMethod[CandidateListTouchBarItem](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCandidateListTouchBarItem() CandidateListTouchBarItem {
	return CandidateListTouchBarItemClass.New()
}

func CandidateListTouchBarItem_New() CandidateListTouchBarItem {
	return CandidateListTouchBarItemClass.New()
}

func (c_ CandidateListTouchBarItem) Init() CandidateListTouchBarItem {
	rv := objc.CallMethod[CandidateListTouchBarItem](c_, objc.GetSelector("init"))
	return rv
}

func CandidateListTouchBarItem_Init() CandidateListTouchBarItem {
	return CandidateListTouchBarItemClass.Alloc().Init()
}

func (c_ CandidateListTouchBarItem) UpdateWithInsertionPointVisibility(isVisible bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("updateWithInsertionPointVisibility:"), isVisible)
}

func (c_ CandidateListTouchBarItem) Client() View {
	rv := objc.CallMethod[View](c_, objc.GetSelector("client"))
	return rv
}

func (c_ CandidateListTouchBarItem) SetClient(value IView) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setClient:"), objc.ExtractPtr(value))
}

func (c_ CandidateListTouchBarItem) Delegate() CandidateListTouchBarItemDelegateWrapper {
	rv := objc.CallMethod[CandidateListTouchBarItemDelegateWrapper](c_, objc.GetSelector("delegate"))
	return rv
}

func (c_ CandidateListTouchBarItem) SetDelegate(value ICandidateListTouchBarItemDelegate) {
	po := objc.WrapAsProtocol("NSCandidateListTouchBarItemDelegate", value)
	objc.SetAssociatedObject(c_, objc.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setDelegate:"), po)
}

func (c_ CandidateListTouchBarItem) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setDelegate:"), objc.ExtractPtr(value))
}

func (c_ CandidateListTouchBarItem) AllowsTextInputContextCandidates() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("allowsTextInputContextCandidates"))
	return rv
}

func (c_ CandidateListTouchBarItem) SetAllowsTextInputContextCandidates(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setAllowsTextInputContextCandidates:"), value)
}

func (c_ CandidateListTouchBarItem) AllowsCollapsing() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("allowsCollapsing"))
	return rv
}

func (c_ CandidateListTouchBarItem) SetAllowsCollapsing(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setAllowsCollapsing:"), value)
}

func (c_ CandidateListTouchBarItem) IsCollapsed() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isCollapsed"))
	return rv
}

func (c_ CandidateListTouchBarItem) SetCollapsed(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setCollapsed:"), value)
}

func (c_ CandidateListTouchBarItem) IsCandidateListVisible() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isCandidateListVisible"))
	return rv
}

func (c_ CandidateListTouchBarItem) SetCustomizationLabel(value string) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setCustomizationLabel:"), value)
}
