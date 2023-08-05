// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var TouchBarClass = _TouchBarClass{objc.GetClass("NSTouchBar")}

type _TouchBarClass struct {
	objc.Class
}

type ITouchBar interface {
	objc.IObject
	ItemForIdentifier(identifier TouchBarItemIdentifier) TouchBarItem
	Delegate() TouchBarDelegateWrapper
	SetDelegate(value ITouchBarDelegate)
	SetDelegate0(value objc.IObject)
	TemplateItems() foundation.Set
	SetTemplateItems(value foundation.ISet)
	DefaultItemIdentifiers() []TouchBarItemIdentifier
	SetDefaultItemIdentifiers(value []TouchBarItemIdentifier)
	PrincipalItemIdentifier() TouchBarItemIdentifier
	SetPrincipalItemIdentifier(value TouchBarItemIdentifier)
	EscapeKeyReplacementItemIdentifier() TouchBarItemIdentifier
	SetEscapeKeyReplacementItemIdentifier(value TouchBarItemIdentifier)
	IsVisible() bool
	ItemIdentifiers() []TouchBarItemIdentifier
	CustomizationIdentifier() TouchBarCustomizationIdentifier
	SetCustomizationIdentifier(value TouchBarCustomizationIdentifier)
	CustomizationAllowedItemIdentifiers() []TouchBarItemIdentifier
	SetCustomizationAllowedItemIdentifiers(value []TouchBarItemIdentifier)
	CustomizationRequiredItemIdentifiers() []TouchBarItemIdentifier
	SetCustomizationRequiredItemIdentifiers(value []TouchBarItemIdentifier)
}

type TouchBar struct {
	objc.Object
}

func MakeTouchBar(ptr unsafe.Pointer) TouchBar {
	return TouchBar{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TouchBar) Init() TouchBar {
	rv := objc.CallMethod[TouchBar](t_, objc.GetSelector("init"))
	return rv
}

func TouchBar_Init() TouchBar {
	return TouchBarClass.Alloc().Init()
}

func (tc _TouchBarClass) Alloc() TouchBar {
	rv := objc.CallMethod[TouchBar](tc, objc.GetSelector("alloc"))
	return rv
}

func TouchBar_Alloc() TouchBar {
	return TouchBarClass.Alloc()
}

func (tc _TouchBarClass) New() TouchBar {
	rv := objc.CallMethod[TouchBar](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTouchBar() TouchBar {
	return TouchBarClass.New()
}

func TouchBar_New() TouchBar {
	return TouchBarClass.New()
}

func (t_ TouchBar) ItemForIdentifier(identifier TouchBarItemIdentifier) TouchBarItem {
	rv := objc.CallMethod[TouchBarItem](t_, objc.GetSelector("itemForIdentifier:"), identifier)
	return rv
}

func (t_ TouchBar) Delegate() TouchBarDelegateWrapper {
	rv := objc.CallMethod[TouchBarDelegateWrapper](t_, objc.GetSelector("delegate"))
	return rv
}

func (t_ TouchBar) SetDelegate(value ITouchBarDelegate) {
	po := objc.WrapAsProtocol("NSTouchBarDelegate", value)
	objc.SetAssociatedObject(t_, objc.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDelegate:"), po)
}

func (t_ TouchBar) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDelegate:"), objc.ExtractPtr(value))
}

func (t_ TouchBar) TemplateItems() foundation.Set {
	rv := objc.CallMethod[foundation.Set](t_, objc.GetSelector("templateItems"))
	return rv
}

func (t_ TouchBar) SetTemplateItems(value foundation.ISet) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTemplateItems:"), objc.ExtractPtr(value))
}

func (t_ TouchBar) DefaultItemIdentifiers() []TouchBarItemIdentifier {
	rv := objc.CallMethod[[]TouchBarItemIdentifier](t_, objc.GetSelector("defaultItemIdentifiers"))
	return rv
}

func (t_ TouchBar) SetDefaultItemIdentifiers(value []TouchBarItemIdentifier) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDefaultItemIdentifiers:"), value)
}

func (t_ TouchBar) PrincipalItemIdentifier() TouchBarItemIdentifier {
	rv := objc.CallMethod[TouchBarItemIdentifier](t_, objc.GetSelector("principalItemIdentifier"))
	return rv
}

func (t_ TouchBar) SetPrincipalItemIdentifier(value TouchBarItemIdentifier) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setPrincipalItemIdentifier:"), value)
}

func (t_ TouchBar) EscapeKeyReplacementItemIdentifier() TouchBarItemIdentifier {
	rv := objc.CallMethod[TouchBarItemIdentifier](t_, objc.GetSelector("escapeKeyReplacementItemIdentifier"))
	return rv
}

func (t_ TouchBar) SetEscapeKeyReplacementItemIdentifier(value TouchBarItemIdentifier) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setEscapeKeyReplacementItemIdentifier:"), value)
}

func (t_ TouchBar) IsVisible() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isVisible"))
	return rv
}

func (t_ TouchBar) ItemIdentifiers() []TouchBarItemIdentifier {
	rv := objc.CallMethod[[]TouchBarItemIdentifier](t_, objc.GetSelector("itemIdentifiers"))
	return rv
}

func (t_ TouchBar) CustomizationIdentifier() TouchBarCustomizationIdentifier {
	rv := objc.CallMethod[TouchBarCustomizationIdentifier](t_, objc.GetSelector("customizationIdentifier"))
	return rv
}

func (t_ TouchBar) SetCustomizationIdentifier(value TouchBarCustomizationIdentifier) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setCustomizationIdentifier:"), value)
}

func (t_ TouchBar) CustomizationAllowedItemIdentifiers() []TouchBarItemIdentifier {
	rv := objc.CallMethod[[]TouchBarItemIdentifier](t_, objc.GetSelector("customizationAllowedItemIdentifiers"))
	return rv
}

func (t_ TouchBar) SetCustomizationAllowedItemIdentifiers(value []TouchBarItemIdentifier) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setCustomizationAllowedItemIdentifiers:"), value)
}

func (t_ TouchBar) CustomizationRequiredItemIdentifiers() []TouchBarItemIdentifier {
	rv := objc.CallMethod[[]TouchBarItemIdentifier](t_, objc.GetSelector("customizationRequiredItemIdentifiers"))
	return rv
}

func (t_ TouchBar) SetCustomizationRequiredItemIdentifiers(value []TouchBarItemIdentifier) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setCustomizationRequiredItemIdentifiers:"), value)
}

func (tc _TouchBarClass) IsAutomaticCustomizeTouchBarMenuItemEnabled() bool {
	rv := objc.CallMethod[bool](tc, objc.GetSelector("isAutomaticCustomizeTouchBarMenuItemEnabled"))
	return rv
}

func TouchBar_IsAutomaticCustomizeTouchBarMenuItemEnabled() bool {
	return TouchBarClass.IsAutomaticCustomizeTouchBarMenuItemEnabled()
}

func (tc _TouchBarClass) SetAutomaticCustomizeTouchBarMenuItemEnabled(value bool) {
	objc.CallMethod[objc.Void](tc, objc.GetSelector("setAutomaticCustomizeTouchBarMenuItemEnabled:"), value)
}

func TouchBar_SetAutomaticCustomizeTouchBarMenuItemEnabled(value bool) {
	TouchBarClass.SetAutomaticCustomizeTouchBarMenuItemEnabled(value)
}
