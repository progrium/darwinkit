// auto-generated code, do not modify
package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type IMenuDelegate interface {
	ImplementsMenuUpdateItemAtIndexShouldCancel() bool
	// optional
	MenuUpdateItemAtIndexShouldCancel(menu Menu, item MenuItem, index int, shouldCancel bool) bool
	ImplementsConfinementRectForMenuOnScreen() bool
	// optional
	ConfinementRectForMenuOnScreen(menu Menu, screen Screen) foundation.Rect
	ImplementsMenuWillHighlightItem() bool
	// optional
	MenuWillHighlightItem(menu Menu, item MenuItem)
	ImplementsMenuWillOpen() bool
	// optional
	MenuWillOpen(menu Menu)
	ImplementsMenuDidClose() bool
	// optional
	MenuDidClose(menu Menu)
	ImplementsNumberOfItemsInMenu() bool
	// optional
	NumberOfItemsInMenu(menu Menu) int
	ImplementsMenuNeedsUpdate() bool
	// optional
	MenuNeedsUpdate(menu Menu)
}

type MenuDelegate struct {
	_MenuUpdateItemAtIndexShouldCancel func(menu Menu, item MenuItem, index int, shouldCancel bool) bool
	_ConfinementRectForMenuOnScreen    func(menu Menu, screen Screen) foundation.Rect
	_MenuWillHighlightItem             func(menu Menu, item MenuItem)
	_MenuWillOpen                      func(menu Menu)
	_MenuDidClose                      func(menu Menu)
	_NumberOfItemsInMenu               func(menu Menu) int
	_MenuNeedsUpdate                   func(menu Menu)
}

func (di *MenuDelegate) ImplementsMenuUpdateItemAtIndexShouldCancel() bool {
	return di._MenuUpdateItemAtIndexShouldCancel != nil
}

func (di *MenuDelegate) SetMenuUpdateItemAtIndexShouldCancel(f func(menu Menu, item MenuItem, index int, shouldCancel bool) bool) {
	di._MenuUpdateItemAtIndexShouldCancel = f
}

func (di *MenuDelegate) MenuUpdateItemAtIndexShouldCancel(menu Menu, item MenuItem, index int, shouldCancel bool) bool {
	return di._MenuUpdateItemAtIndexShouldCancel(menu, item, index, shouldCancel)
}
func (di *MenuDelegate) ImplementsConfinementRectForMenuOnScreen() bool {
	return di._ConfinementRectForMenuOnScreen != nil
}

func (di *MenuDelegate) SetConfinementRectForMenuOnScreen(f func(menu Menu, screen Screen) foundation.Rect) {
	di._ConfinementRectForMenuOnScreen = f
}

func (di *MenuDelegate) ConfinementRectForMenuOnScreen(menu Menu, screen Screen) foundation.Rect {
	return di._ConfinementRectForMenuOnScreen(menu, screen)
}
func (di *MenuDelegate) ImplementsMenuWillHighlightItem() bool {
	return di._MenuWillHighlightItem != nil
}

func (di *MenuDelegate) SetMenuWillHighlightItem(f func(menu Menu, item MenuItem)) {
	di._MenuWillHighlightItem = f
}

func (di *MenuDelegate) MenuWillHighlightItem(menu Menu, item MenuItem) {
	di._MenuWillHighlightItem(menu, item)
}
func (di *MenuDelegate) ImplementsMenuWillOpen() bool {
	return di._MenuWillOpen != nil
}

func (di *MenuDelegate) SetMenuWillOpen(f func(menu Menu)) {
	di._MenuWillOpen = f
}

func (di *MenuDelegate) MenuWillOpen(menu Menu) {
	di._MenuWillOpen(menu)
}
func (di *MenuDelegate) ImplementsMenuDidClose() bool {
	return di._MenuDidClose != nil
}

func (di *MenuDelegate) SetMenuDidClose(f func(menu Menu)) {
	di._MenuDidClose = f
}

func (di *MenuDelegate) MenuDidClose(menu Menu) {
	di._MenuDidClose(menu)
}
func (di *MenuDelegate) ImplementsNumberOfItemsInMenu() bool {
	return di._NumberOfItemsInMenu != nil
}

func (di *MenuDelegate) SetNumberOfItemsInMenu(f func(menu Menu) int) {
	di._NumberOfItemsInMenu = f
}

func (di *MenuDelegate) NumberOfItemsInMenu(menu Menu) int {
	return di._NumberOfItemsInMenu(menu)
}
func (di *MenuDelegate) ImplementsMenuNeedsUpdate() bool {
	return di._MenuNeedsUpdate != nil
}

func (di *MenuDelegate) SetMenuNeedsUpdate(f func(menu Menu)) {
	di._MenuNeedsUpdate = f
}

func (di *MenuDelegate) MenuNeedsUpdate(menu Menu) {
	di._MenuNeedsUpdate(menu)
}

type MenuDelegateWrapper struct {
	objc.Object
}

func (m_ MenuDelegateWrapper) ImplementsMenuUpdateItemAtIndexShouldCancel() bool {
	return m_.RespondsToSelector(objc.GetSelector("menu:updateItem:atIndex:shouldCancel:"))
}

func (m_ MenuDelegateWrapper) MenuUpdateItemAtIndexShouldCancel(menu IMenu, item IMenuItem, index int, shouldCancel bool) bool {
	rv := objc.CallMethod[bool](m_, objc.GetSelector("menu:updateItem:atIndex:shouldCancel:"), objc.ExtractPtr(menu), objc.ExtractPtr(item), index, shouldCancel)
	return rv
}

func (m_ MenuDelegateWrapper) ImplementsConfinementRectForMenuOnScreen() bool {
	return m_.RespondsToSelector(objc.GetSelector("confinementRectForMenu:onScreen:"))
}

func (m_ MenuDelegateWrapper) ConfinementRectForMenuOnScreen(menu IMenu, screen IScreen) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](m_, objc.GetSelector("confinementRectForMenu:onScreen:"), objc.ExtractPtr(menu), objc.ExtractPtr(screen))
	return rv
}

func (m_ MenuDelegateWrapper) ImplementsMenuWillHighlightItem() bool {
	return m_.RespondsToSelector(objc.GetSelector("menu:willHighlightItem:"))
}

func (m_ MenuDelegateWrapper) MenuWillHighlightItem(menu IMenu, item IMenuItem) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("menu:willHighlightItem:"), objc.ExtractPtr(menu), objc.ExtractPtr(item))
}

func (m_ MenuDelegateWrapper) ImplementsMenuWillOpen() bool {
	return m_.RespondsToSelector(objc.GetSelector("menuWillOpen:"))
}

func (m_ MenuDelegateWrapper) MenuWillOpen(menu IMenu) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("menuWillOpen:"), objc.ExtractPtr(menu))
}

func (m_ MenuDelegateWrapper) ImplementsMenuDidClose() bool {
	return m_.RespondsToSelector(objc.GetSelector("menuDidClose:"))
}

func (m_ MenuDelegateWrapper) MenuDidClose(menu IMenu) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("menuDidClose:"), objc.ExtractPtr(menu))
}

func (m_ MenuDelegateWrapper) ImplementsNumberOfItemsInMenu() bool {
	return m_.RespondsToSelector(objc.GetSelector("numberOfItemsInMenu:"))
}

func (m_ MenuDelegateWrapper) NumberOfItemsInMenu(menu IMenu) int {
	rv := objc.CallMethod[int](m_, objc.GetSelector("numberOfItemsInMenu:"), objc.ExtractPtr(menu))
	return rv
}

func (m_ MenuDelegateWrapper) ImplementsMenuNeedsUpdate() bool {
	return m_.RespondsToSelector(objc.GetSelector("menuNeedsUpdate:"))
}

func (m_ MenuDelegateWrapper) MenuNeedsUpdate(menu IMenu) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("menuNeedsUpdate:"), objc.ExtractPtr(menu))
}
