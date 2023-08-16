// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MenuToolbarItem] class.
var MenuToolbarItemClass = _MenuToolbarItemClass{objc.GetClass("NSMenuToolbarItem")}

type _MenuToolbarItemClass struct {
	objc.Class
}

// An interface definition for the [MenuToolbarItem] class.
type IMenuToolbarItem interface {
	IToolbarItem
	ShowsIndicator() bool
	SetShowsIndicator(value bool)
	Menu() Menu
	SetMenu(value IMenu)
}

// A control that presents a menu in a window’s toolbar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenutoolbaritem?language=objc
type MenuToolbarItem struct {
	ToolbarItem
}

func MenuToolbarItemFrom(ptr unsafe.Pointer) MenuToolbarItem {
	return MenuToolbarItem{
		ToolbarItem: ToolbarItemFrom(ptr),
	}
}

func (mc _MenuToolbarItemClass) Alloc() MenuToolbarItem {
	rv := objc.Call[MenuToolbarItem](mc, objc.Sel("alloc"))
	return rv
}

func MenuToolbarItem_Alloc() MenuToolbarItem {
	return MenuToolbarItemClass.Alloc()
}

func (mc _MenuToolbarItemClass) New() MenuToolbarItem {
	rv := objc.Call[MenuToolbarItem](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMenuToolbarItem() MenuToolbarItem {
	return MenuToolbarItemClass.New()
}

func (m_ MenuToolbarItem) Init() MenuToolbarItem {
	rv := objc.Call[MenuToolbarItem](m_, objc.Sel("init"))
	return rv
}

func (m_ MenuToolbarItem) InitWithItemIdentifier(itemIdentifier ToolbarItemIdentifier) MenuToolbarItem {
	rv := objc.Call[MenuToolbarItem](m_, objc.Sel("initWithItemIdentifier:"), itemIdentifier)
	return rv
}

// Creates a toolbar item with the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/1534084-initwithitemidentifier?language=objc
func MenuToolbarItem_InitWithItemIdentifier(itemIdentifier ToolbarItemIdentifier) MenuToolbarItem {
	return MenuToolbarItemClass.Alloc().InitWithItemIdentifier(itemIdentifier)
}

// A Boolean value that determines whether the toolbar item displays an indicator of additional functionality. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenutoolbaritem/3237192-showsindicator?language=objc
func (m_ MenuToolbarItem) ShowsIndicator() bool {
	rv := objc.Call[bool](m_, objc.Sel("showsIndicator"))
	return rv
}

// A Boolean value that determines whether the toolbar item displays an indicator of additional functionality. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenutoolbaritem/3237192-showsindicator?language=objc
func (m_ MenuToolbarItem) SetShowsIndicator(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setShowsIndicator:"), value)
}

// The menu presented from the toolbar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenutoolbaritem/3237191-menu?language=objc
func (m_ MenuToolbarItem) Menu() Menu {
	rv := objc.Call[Menu](m_, objc.Sel("menu"))
	return rv
}

// The menu presented from the toolbar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenutoolbaritem/3237191-menu?language=objc
func (m_ MenuToolbarItem) SetMenu(value IMenu) {
	objc.Call[objc.Void](m_, objc.Sel("setMenu:"), objc.Ptr(value))
}
