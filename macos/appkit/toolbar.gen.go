// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Toolbar] class.
var ToolbarClass = _ToolbarClass{objc.GetClass("NSToolbar")}

type _ToolbarClass struct {
	objc.Class
}

// An interface definition for the [Toolbar] class.
type IToolbar interface {
	objc.IObject
	SetConfigurationFromDictionary(configDict map[string]objc.IObject)
	InsertItemWithItemIdentifierAtIndex(itemIdentifier ToolbarItemIdentifier, index int)
	RemoveItemAtIndex(index int)
	ValidateVisibleItems()
	RunCustomizationPalette(sender objc.IObject)
	IsVisible() bool
	SetVisible(value bool)
	DisplayMode() ToolbarDisplayMode
	SetDisplayMode(value ToolbarDisplayMode)
	AllowsUserCustomization() bool
	SetAllowsUserCustomization(value bool)
	Items() []ToolbarItem
	ConfigurationDictionary() map[string]objc.Object
	Delegate() ToolbarDelegateWrapper
	SetDelegate(value PToolbarDelegate)
	SetDelegateObject(valueObject objc.IObject)
	ShowsBaselineSeparator() bool
	SetShowsBaselineSeparator(value bool)
	VisibleItems() []ToolbarItem
	AllowsExtensionItems() bool
	SetAllowsExtensionItems(value bool)
	CustomizationPaletteIsRunning() bool
	Identifier() ToolbarIdentifier
	SelectedItemIdentifier() ToolbarItemIdentifier
	SetSelectedItemIdentifier(value ToolbarItemIdentifier)
	AutosavesConfiguration() bool
	SetAutosavesConfiguration(value bool)
}

// An object that manages the space above your app’s custom content and either below or integrated with the window’s title bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar?language=objc
type Toolbar struct {
	objc.Object
}

func ToolbarFrom(ptr unsafe.Pointer) Toolbar {
	return Toolbar{
		Object: objc.ObjectFrom(ptr),
	}
}

func (t_ Toolbar) InitWithIdentifier(identifier ToolbarIdentifier) Toolbar {
	rv := objc.Call[Toolbar](t_, objc.Sel("initWithIdentifier:"), identifier)
	return rv
}

// Creates a newly allocated toolbar with the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/1516975-initwithidentifier?language=objc
func Toolbar_InitWithIdentifier(identifier ToolbarIdentifier) Toolbar {
	return ToolbarClass.Alloc().InitWithIdentifier(identifier)
}

func (t_ Toolbar) Init() Toolbar {
	rv := objc.Call[Toolbar](t_, objc.Sel("init"))
	return rv
}

func (tc _ToolbarClass) Alloc() Toolbar {
	rv := objc.Call[Toolbar](tc, objc.Sel("alloc"))
	return rv
}

func Toolbar_Alloc() Toolbar {
	return ToolbarClass.Alloc()
}

func (tc _ToolbarClass) New() Toolbar {
	rv := objc.Call[Toolbar](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewToolbar() Toolbar {
	return ToolbarClass.New()
}

// Specifies the new configuration details for the toolbar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/1516951-setconfigurationfromdictionary?language=objc
func (t_ Toolbar) SetConfigurationFromDictionary(configDict map[string]objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("setConfigurationFromDictionary:"), configDict)
}

// Inserts an item into the toolbar at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/1516941-insertitemwithitemidentifier?language=objc
func (t_ Toolbar) InsertItemWithItemIdentifierAtIndex(itemIdentifier ToolbarItemIdentifier, index int) {
	objc.Call[objc.Void](t_, objc.Sel("insertItemWithItemIdentifier:atIndex:"), itemIdentifier, index)
}

// Removes the item at the specified index in the toolbar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/1516980-removeitematindex?language=objc
func (t_ Toolbar) RemoveItemAtIndex(index int) {
	objc.Call[objc.Void](t_, objc.Sel("removeItemAtIndex:"), index)
}

// Validates the toolbar’s visible items during a window update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/1516947-validatevisibleitems?language=objc
func (t_ Toolbar) ValidateVisibleItems() {
	objc.Call[objc.Void](t_, objc.Sel("validateVisibleItems"))
}

// Displays the toolbar’s customization palette and handles any user-initiated customizations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/1516979-runcustomizationpalette?language=objc
func (t_ Toolbar) RunCustomizationPalette(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("runCustomizationPalette:"), sender)
}

// A Boolean value that indicates whether the toolbar is visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/1516935-visible?language=objc
func (t_ Toolbar) IsVisible() bool {
	rv := objc.Call[bool](t_, objc.Sel("isVisible"))
	return rv
}

// A Boolean value that indicates whether the toolbar is visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/1516935-visible?language=objc
func (t_ Toolbar) SetVisible(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setVisible:"), value)
}

// A value that indicates whether the toolbar displays items using a name, icon, or combination of elements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/1516937-displaymode?language=objc
func (t_ Toolbar) DisplayMode() ToolbarDisplayMode {
	rv := objc.Call[ToolbarDisplayMode](t_, objc.Sel("displayMode"))
	return rv
}

// A value that indicates whether the toolbar displays items using a name, icon, or combination of elements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/1516937-displaymode?language=objc
func (t_ Toolbar) SetDisplayMode(value ToolbarDisplayMode) {
	objc.Call[objc.Void](t_, objc.Sel("setDisplayMode:"), value)
}

// A Boolean value that indicates whether users can modify the contents of the toolbar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/1516962-allowsusercustomization?language=objc
func (t_ Toolbar) AllowsUserCustomization() bool {
	rv := objc.Call[bool](t_, objc.Sel("allowsUserCustomization"))
	return rv
}

// A Boolean value that indicates whether users can modify the contents of the toolbar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/1516962-allowsusercustomization?language=objc
func (t_ Toolbar) SetAllowsUserCustomization(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAllowsUserCustomization:"), value)
}

// An array containing the toolbar’s current items, in order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/1516946-items?language=objc
func (t_ Toolbar) Items() []ToolbarItem {
	rv := objc.Call[[]ToolbarItem](t_, objc.Sel("items"))
	return rv
}

// A dictionary containing the current configuration details for the toolbar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/1516956-configurationdictionary?language=objc
func (t_ Toolbar) ConfigurationDictionary() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](t_, objc.Sel("configurationDictionary"))
	return rv
}

// The object you use to customize the toolbar contents and configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/1516939-delegate?language=objc
func (t_ Toolbar) Delegate() ToolbarDelegateWrapper {
	rv := objc.Call[ToolbarDelegateWrapper](t_, objc.Sel("delegate"))
	return rv
}

// The object you use to customize the toolbar contents and configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/1516939-delegate?language=objc
func (t_ Toolbar) SetDelegate(value PToolbarDelegate) {
	po0 := objc.WrapAsProtocol("NSToolbarDelegate", value)
	objc.SetAssociatedObject(t_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](t_, objc.Sel("setDelegate:"), po0)
}

// The object you use to customize the toolbar contents and configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/1516939-delegate?language=objc
func (t_ Toolbar) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// A Boolean value that indicates whether the toolbar shows the separator between the toolbar and the main window contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/1516954-showsbaselineseparator?language=objc
func (t_ Toolbar) ShowsBaselineSeparator() bool {
	rv := objc.Call[bool](t_, objc.Sel("showsBaselineSeparator"))
	return rv
}

// A Boolean value that indicates whether the toolbar shows the separator between the toolbar and the main window contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/1516954-showsbaselineseparator?language=objc
func (t_ Toolbar) SetShowsBaselineSeparator(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setShowsBaselineSeparator:"), value)
}

// An array containing the toolbar’s currently visible items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/1516993-visibleitems?language=objc
func (t_ Toolbar) VisibleItems() []ToolbarItem {
	rv := objc.Call[[]ToolbarItem](t_, objc.Sel("visibleItems"))
	return rv
}

// A Boolean value that indicates whether the toolbar can add items for Action extensions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/1517005-allowsextensionitems?language=objc
func (t_ Toolbar) AllowsExtensionItems() bool {
	rv := objc.Call[bool](t_, objc.Sel("allowsExtensionItems"))
	return rv
}

// A Boolean value that indicates whether the toolbar can add items for Action extensions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/1517005-allowsextensionitems?language=objc
func (t_ Toolbar) SetAllowsExtensionItems(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAllowsExtensionItems:"), value)
}

// A Boolean value that indicates whether the toolbar’s customization palette is in use. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/1516987-customizationpaletteisrunning?language=objc
func (t_ Toolbar) CustomizationPaletteIsRunning() bool {
	rv := objc.Call[bool](t_, objc.Sel("customizationPaletteIsRunning"))
	return rv
}

// The value you use to identify the toolbar in your app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/1516953-identifier?language=objc
func (t_ Toolbar) Identifier() ToolbarIdentifier {
	rv := objc.Call[ToolbarIdentifier](t_, objc.Sel("identifier"))
	return rv
}

// The identifier of the toolbar’s currently selected item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/1516999-selecteditemidentifier?language=objc
func (t_ Toolbar) SelectedItemIdentifier() ToolbarItemIdentifier {
	rv := objc.Call[ToolbarItemIdentifier](t_, objc.Sel("selectedItemIdentifier"))
	return rv
}

// The identifier of the toolbar’s currently selected item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/1516999-selecteditemidentifier?language=objc
func (t_ Toolbar) SetSelectedItemIdentifier(value ToolbarItemIdentifier) {
	objc.Call[objc.Void](t_, objc.Sel("setSelectedItemIdentifier:"), value)
}

// A Boolean value that indicates whether the toolbar autosaves its configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/1516992-autosavesconfiguration?language=objc
func (t_ Toolbar) AutosavesConfiguration() bool {
	rv := objc.Call[bool](t_, objc.Sel("autosavesConfiguration"))
	return rv
}

// A Boolean value that indicates whether the toolbar autosaves its configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/1516992-autosavesconfiguration?language=objc
func (t_ Toolbar) SetAutosavesConfiguration(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAutosavesConfiguration:"), value)
}
