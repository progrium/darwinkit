// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var ToolbarClass = _ToolbarClass{objc.GetClass("NSToolbar")}

type _ToolbarClass struct {
	objc.Class
}

type IToolbar interface {
	objc.IObject
	InsertItemWithItemIdentifierAtIndex(itemIdentifier ToolbarItemIdentifier, index int)
	RemoveItemAtIndex(index int)
	SetConfigurationFromDictionary(configDict map[string]objc.IObject)
	RunCustomizationPalette(sender objc.IObject)
	ValidateVisibleItems()
	Delegate() ToolbarDelegateWrapper
	SetDelegate(value IToolbarDelegate)
	SetDelegate0(value objc.IObject)
	Identifier() ToolbarIdentifier
	IsVisible() bool
	SetVisible(value bool)
	DisplayMode() ToolbarDisplayMode
	SetDisplayMode(value ToolbarDisplayMode)
	ShowsBaselineSeparator() bool
	SetShowsBaselineSeparator(value bool)
	AllowsUserCustomization() bool
	SetAllowsUserCustomization(value bool)
	AllowsExtensionItems() bool
	SetAllowsExtensionItems(value bool)
	Items() []ToolbarItem
	VisibleItems() []ToolbarItem
	CenteredItemIdentifiers() foundation.Set
	SetCenteredItemIdentifiers(value foundation.ISet)
	SelectedItemIdentifier() ToolbarItemIdentifier
	SetSelectedItemIdentifier(value ToolbarItemIdentifier)
	AutosavesConfiguration() bool
	SetAutosavesConfiguration(value bool)
	ConfigurationDictionary() map[string]objc.Object
	CustomizationPaletteIsRunning() bool
}

type Toolbar struct {
	objc.Object
}

func MakeToolbar(ptr unsafe.Pointer) Toolbar {
	return Toolbar{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ Toolbar) InitWithIdentifier(identifier ToolbarIdentifier) Toolbar {
	rv := objc.CallMethod[Toolbar](t_, objc.GetSelector("initWithIdentifier:"), identifier)
	return rv
}

func Toolbar_InitWithIdentifier(identifier ToolbarIdentifier) Toolbar {
	return ToolbarClass.Alloc().InitWithIdentifier(identifier)
}

func (t_ Toolbar) Init() Toolbar {
	rv := objc.CallMethod[Toolbar](t_, objc.GetSelector("init"))
	return rv
}

func Toolbar_Init() Toolbar {
	return ToolbarClass.Alloc().Init()
}

func (tc _ToolbarClass) Alloc() Toolbar {
	rv := objc.CallMethod[Toolbar](tc, objc.GetSelector("alloc"))
	return rv
}

func Toolbar_Alloc() Toolbar {
	return ToolbarClass.Alloc()
}

func (tc _ToolbarClass) New() Toolbar {
	rv := objc.CallMethod[Toolbar](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewToolbar() Toolbar {
	return ToolbarClass.New()
}

func Toolbar_New() Toolbar {
	return ToolbarClass.New()
}

func (t_ Toolbar) InsertItemWithItemIdentifierAtIndex(itemIdentifier ToolbarItemIdentifier, index int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("insertItemWithItemIdentifier:atIndex:"), itemIdentifier, index)
}

func (t_ Toolbar) RemoveItemAtIndex(index int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("removeItemAtIndex:"), index)
}

func (t_ Toolbar) SetConfigurationFromDictionary(configDict map[string]objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setConfigurationFromDictionary:"), configDict)
}

func (t_ Toolbar) RunCustomizationPalette(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("runCustomizationPalette:"), objc.ExtractPtr(sender))
}

func (t_ Toolbar) ValidateVisibleItems() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("validateVisibleItems"))
}

func (t_ Toolbar) Delegate() ToolbarDelegateWrapper {
	rv := objc.CallMethod[ToolbarDelegateWrapper](t_, objc.GetSelector("delegate"))
	return rv
}

func (t_ Toolbar) SetDelegate(value IToolbarDelegate) {
	po := objc.WrapAsProtocol("NSToolbarDelegate", value)
	objc.SetAssociatedObject(t_, objc.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDelegate:"), po)
}

func (t_ Toolbar) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDelegate:"), objc.ExtractPtr(value))
}

func (t_ Toolbar) Identifier() ToolbarIdentifier {
	rv := objc.CallMethod[ToolbarIdentifier](t_, objc.GetSelector("identifier"))
	return rv
}

func (t_ Toolbar) IsVisible() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isVisible"))
	return rv
}

func (t_ Toolbar) SetVisible(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setVisible:"), value)
}

func (t_ Toolbar) DisplayMode() ToolbarDisplayMode {
	rv := objc.CallMethod[ToolbarDisplayMode](t_, objc.GetSelector("displayMode"))
	return rv
}

func (t_ Toolbar) SetDisplayMode(value ToolbarDisplayMode) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDisplayMode:"), value)
}

func (t_ Toolbar) ShowsBaselineSeparator() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("showsBaselineSeparator"))
	return rv
}

func (t_ Toolbar) SetShowsBaselineSeparator(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setShowsBaselineSeparator:"), value)
}

func (t_ Toolbar) AllowsUserCustomization() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("allowsUserCustomization"))
	return rv
}

func (t_ Toolbar) SetAllowsUserCustomization(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAllowsUserCustomization:"), value)
}

func (t_ Toolbar) AllowsExtensionItems() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("allowsExtensionItems"))
	return rv
}

func (t_ Toolbar) SetAllowsExtensionItems(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAllowsExtensionItems:"), value)
}

func (t_ Toolbar) Items() []ToolbarItem {
	rv := objc.CallMethod[[]ToolbarItem](t_, objc.GetSelector("items"))
	return rv
}

func (t_ Toolbar) VisibleItems() []ToolbarItem {
	rv := objc.CallMethod[[]ToolbarItem](t_, objc.GetSelector("visibleItems"))
	return rv
}

func (t_ Toolbar) CenteredItemIdentifiers() foundation.Set {
	rv := objc.CallMethod[foundation.Set](t_, objc.GetSelector("centeredItemIdentifiers"))
	return rv
}

func (t_ Toolbar) SetCenteredItemIdentifiers(value foundation.ISet) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setCenteredItemIdentifiers:"), objc.ExtractPtr(value))
}

func (t_ Toolbar) SelectedItemIdentifier() ToolbarItemIdentifier {
	rv := objc.CallMethod[ToolbarItemIdentifier](t_, objc.GetSelector("selectedItemIdentifier"))
	return rv
}

func (t_ Toolbar) SetSelectedItemIdentifier(value ToolbarItemIdentifier) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setSelectedItemIdentifier:"), value)
}

func (t_ Toolbar) AutosavesConfiguration() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("autosavesConfiguration"))
	return rv
}

func (t_ Toolbar) SetAutosavesConfiguration(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAutosavesConfiguration:"), value)
}

func (t_ Toolbar) ConfigurationDictionary() map[string]objc.Object {
	rv := objc.CallMethod[map[string]objc.Object](t_, objc.GetSelector("configurationDictionary"))
	return rv
}

func (t_ Toolbar) CustomizationPaletteIsRunning() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("customizationPaletteIsRunning"))
	return rv
}
