// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var MenuToolbarItemClass = _MenuToolbarItemClass{objc.GetClass("NSMenuToolbarItem")}

type _MenuToolbarItemClass struct {
	objc.Class
}

type IMenuToolbarItem interface {
	IToolbarItem
	ShowsIndicator() bool
	SetShowsIndicator(value bool)
	Menu() Menu
	SetMenu(value IMenu)
}

type MenuToolbarItem struct {
	ToolbarItem
}

func MakeMenuToolbarItem(ptr unsafe.Pointer) MenuToolbarItem {
	return MenuToolbarItem{
		ToolbarItem: MakeToolbarItem(ptr),
	}
}

func (m_ MenuToolbarItem) InitWithItemIdentifier(itemIdentifier ToolbarItemIdentifier) MenuToolbarItem {
	rv := objc.CallMethod[MenuToolbarItem](m_, objc.GetSelector("initWithItemIdentifier:"), itemIdentifier)
	return rv
}

func MenuToolbarItem_InitWithItemIdentifier(itemIdentifier ToolbarItemIdentifier) MenuToolbarItem {
	return MenuToolbarItemClass.Alloc().InitWithItemIdentifier(itemIdentifier)
}

func (mc _MenuToolbarItemClass) Alloc() MenuToolbarItem {
	rv := objc.CallMethod[MenuToolbarItem](mc, objc.GetSelector("alloc"))
	return rv
}

func MenuToolbarItem_Alloc() MenuToolbarItem {
	return MenuToolbarItemClass.Alloc()
}

func (mc _MenuToolbarItemClass) New() MenuToolbarItem {
	rv := objc.CallMethod[MenuToolbarItem](mc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewMenuToolbarItem() MenuToolbarItem {
	return MenuToolbarItemClass.New()
}

func MenuToolbarItem_New() MenuToolbarItem {
	return MenuToolbarItemClass.New()
}

func (m_ MenuToolbarItem) Init() MenuToolbarItem {
	rv := objc.CallMethod[MenuToolbarItem](m_, objc.GetSelector("init"))
	return rv
}

func MenuToolbarItem_Init() MenuToolbarItem {
	return MenuToolbarItemClass.Alloc().Init()
}

func (m_ MenuToolbarItem) ShowsIndicator() bool {
	rv := objc.CallMethod[bool](m_, objc.GetSelector("showsIndicator"))
	return rv
}

func (m_ MenuToolbarItem) SetShowsIndicator(value bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setShowsIndicator:"), value)
}

func (m_ MenuToolbarItem) Menu() Menu {
	rv := objc.CallMethod[Menu](m_, objc.GetSelector("menu"))
	return rv
}

func (m_ MenuToolbarItem) SetMenu(value IMenu) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setMenu:"), objc.ExtractPtr(value))
}
