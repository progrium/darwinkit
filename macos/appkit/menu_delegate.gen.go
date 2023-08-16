// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The optional methods implemented by delegates of NSMenu objects to manage menu display and handle some events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenudelegate?language=objc
type PMenuDelegate interface {
	// optional
	MenuHasKeyEquivalentForEventTargetAction(menu Menu, event Event, target objc.Object, action objc.Selector) bool
	HasMenuHasKeyEquivalentForEventTargetAction() bool

	// optional
	MenuWillHighlightItem(menu Menu, item MenuItem)
	HasMenuWillHighlightItem() bool

	// optional
	MenuDidClose(menu Menu)
	HasMenuDidClose() bool

	// optional
	NumberOfItemsInMenu(menu Menu) int
	HasNumberOfItemsInMenu() bool

	// optional
	MenuNeedsUpdate(menu Menu)
	HasMenuNeedsUpdate() bool

	// optional
	ConfinementRectForMenuOnScreen(menu Menu, screen Screen) foundation.Rect
	HasConfinementRectForMenuOnScreen() bool

	// optional
	MenuWillOpen(menu Menu)
	HasMenuWillOpen() bool
}

// A delegate implementation builder for the [PMenuDelegate] protocol.
type MenuDelegate struct {
	_MenuHasKeyEquivalentForEventTargetAction func(menu Menu, event Event, target objc.Object, action objc.Selector) bool
	_MenuWillHighlightItem                    func(menu Menu, item MenuItem)
	_MenuDidClose                             func(menu Menu)
	_NumberOfItemsInMenu                      func(menu Menu) int
	_MenuNeedsUpdate                          func(menu Menu)
	_ConfinementRectForMenuOnScreen           func(menu Menu, screen Screen) foundation.Rect
	_MenuWillOpen                             func(menu Menu)
}

func (di *MenuDelegate) HasMenuHasKeyEquivalentForEventTargetAction() bool {
	return di._MenuHasKeyEquivalentForEventTargetAction != nil
}

// Invoked to allow the delegate to return the target and action for a key-down event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenudelegate/1518211-menuhaskeyequivalent?language=objc
func (di *MenuDelegate) SetMenuHasKeyEquivalentForEventTargetAction(f func(menu Menu, event Event, target objc.Object, action objc.Selector) bool) {
	di._MenuHasKeyEquivalentForEventTargetAction = f
}

// Invoked to allow the delegate to return the target and action for a key-down event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenudelegate/1518211-menuhaskeyequivalent?language=objc
func (di *MenuDelegate) MenuHasKeyEquivalentForEventTargetAction(menu Menu, event Event, target objc.Object, action objc.Selector) bool {
	return di._MenuHasKeyEquivalentForEventTargetAction(menu, event, target, action)
}
func (di *MenuDelegate) HasMenuWillHighlightItem() bool {
	return di._MenuWillHighlightItem != nil
}

// Invoked to indicate that a menu is about to highlight a given item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenudelegate/1518260-menu?language=objc
func (di *MenuDelegate) SetMenuWillHighlightItem(f func(menu Menu, item MenuItem)) {
	di._MenuWillHighlightItem = f
}

// Invoked to indicate that a menu is about to highlight a given item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenudelegate/1518260-menu?language=objc
func (di *MenuDelegate) MenuWillHighlightItem(menu Menu, item MenuItem) {
	di._MenuWillHighlightItem(menu, item)
}
func (di *MenuDelegate) HasMenuDidClose() bool {
	return di._MenuDidClose != nil
}

// Invoked after a menu closed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenudelegate/1518167-menudidclose?language=objc
func (di *MenuDelegate) SetMenuDidClose(f func(menu Menu)) {
	di._MenuDidClose = f
}

// Invoked after a menu closed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenudelegate/1518167-menudidclose?language=objc
func (di *MenuDelegate) MenuDidClose(menu Menu) {
	di._MenuDidClose(menu)
}
func (di *MenuDelegate) HasNumberOfItemsInMenu() bool {
	return di._NumberOfItemsInMenu != nil
}

// Invoked when a menu is about to be displayed at the start of a tracking session so the delegate can specify the number of items in the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenudelegate/1518242-numberofitemsinmenu?language=objc
func (di *MenuDelegate) SetNumberOfItemsInMenu(f func(menu Menu) int) {
	di._NumberOfItemsInMenu = f
}

// Invoked when a menu is about to be displayed at the start of a tracking session so the delegate can specify the number of items in the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenudelegate/1518242-numberofitemsinmenu?language=objc
func (di *MenuDelegate) NumberOfItemsInMenu(menu Menu) int {
	return di._NumberOfItemsInMenu(menu)
}
func (di *MenuDelegate) HasMenuNeedsUpdate() bool {
	return di._MenuNeedsUpdate != nil
}

// Invoked when a menu is about to be displayed at the start of a tracking session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenudelegate/1518235-menuneedsupdate?language=objc
func (di *MenuDelegate) SetMenuNeedsUpdate(f func(menu Menu)) {
	di._MenuNeedsUpdate = f
}

// Invoked when a menu is about to be displayed at the start of a tracking session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenudelegate/1518235-menuneedsupdate?language=objc
func (di *MenuDelegate) MenuNeedsUpdate(menu Menu) {
	di._MenuNeedsUpdate(menu)
}
func (di *MenuDelegate) HasConfinementRectForMenuOnScreen() bool {
	return di._ConfinementRectForMenuOnScreen != nil
}

// Invoked to allow the delegate to specify a display location for the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenudelegate/1518247-confinementrectformenu?language=objc
func (di *MenuDelegate) SetConfinementRectForMenuOnScreen(f func(menu Menu, screen Screen) foundation.Rect) {
	di._ConfinementRectForMenuOnScreen = f
}

// Invoked to allow the delegate to specify a display location for the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenudelegate/1518247-confinementrectformenu?language=objc
func (di *MenuDelegate) ConfinementRectForMenuOnScreen(menu Menu, screen Screen) foundation.Rect {
	return di._ConfinementRectForMenuOnScreen(menu, screen)
}
func (di *MenuDelegate) HasMenuWillOpen() bool {
	return di._MenuWillOpen != nil
}

// Invoked when a menu is about to open. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenudelegate/1518156-menuwillopen?language=objc
func (di *MenuDelegate) SetMenuWillOpen(f func(menu Menu)) {
	di._MenuWillOpen = f
}

// Invoked when a menu is about to open. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenudelegate/1518156-menuwillopen?language=objc
func (di *MenuDelegate) MenuWillOpen(menu Menu) {
	di._MenuWillOpen(menu)
}

// A concrete type wrapper for the [PMenuDelegate] protocol.
type MenuDelegateWrapper struct {
	objc.Object
}

func (m_ MenuDelegateWrapper) HasMenuHasKeyEquivalentForEventTargetAction() bool {
	return m_.RespondsToSelector(objc.Sel("menuHasKeyEquivalent:forEvent:target:action:"))
}

// Invoked to allow the delegate to return the target and action for a key-down event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenudelegate/1518211-menuhaskeyequivalent?language=objc
func (m_ MenuDelegateWrapper) MenuHasKeyEquivalentForEventTargetAction(menu IMenu, event IEvent, target objc.IObject, action objc.Selector) bool {
	rv := objc.Call[bool](m_, objc.Sel("menuHasKeyEquivalent:forEvent:target:action:"), objc.Ptr(menu), objc.Ptr(event), target, action)
	return rv
}

func (m_ MenuDelegateWrapper) HasMenuWillHighlightItem() bool {
	return m_.RespondsToSelector(objc.Sel("menu:willHighlightItem:"))
}

// Invoked to indicate that a menu is about to highlight a given item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenudelegate/1518260-menu?language=objc
func (m_ MenuDelegateWrapper) MenuWillHighlightItem(menu IMenu, item IMenuItem) {
	objc.Call[objc.Void](m_, objc.Sel("menu:willHighlightItem:"), objc.Ptr(menu), objc.Ptr(item))
}

func (m_ MenuDelegateWrapper) HasMenuDidClose() bool {
	return m_.RespondsToSelector(objc.Sel("menuDidClose:"))
}

// Invoked after a menu closed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenudelegate/1518167-menudidclose?language=objc
func (m_ MenuDelegateWrapper) MenuDidClose(menu IMenu) {
	objc.Call[objc.Void](m_, objc.Sel("menuDidClose:"), objc.Ptr(menu))
}

func (m_ MenuDelegateWrapper) HasNumberOfItemsInMenu() bool {
	return m_.RespondsToSelector(objc.Sel("numberOfItemsInMenu:"))
}

// Invoked when a menu is about to be displayed at the start of a tracking session so the delegate can specify the number of items in the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenudelegate/1518242-numberofitemsinmenu?language=objc
func (m_ MenuDelegateWrapper) NumberOfItemsInMenu(menu IMenu) int {
	rv := objc.Call[int](m_, objc.Sel("numberOfItemsInMenu:"), objc.Ptr(menu))
	return rv
}

func (m_ MenuDelegateWrapper) HasMenuNeedsUpdate() bool {
	return m_.RespondsToSelector(objc.Sel("menuNeedsUpdate:"))
}

// Invoked when a menu is about to be displayed at the start of a tracking session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenudelegate/1518235-menuneedsupdate?language=objc
func (m_ MenuDelegateWrapper) MenuNeedsUpdate(menu IMenu) {
	objc.Call[objc.Void](m_, objc.Sel("menuNeedsUpdate:"), objc.Ptr(menu))
}

func (m_ MenuDelegateWrapper) HasConfinementRectForMenuOnScreen() bool {
	return m_.RespondsToSelector(objc.Sel("confinementRectForMenu:onScreen:"))
}

// Invoked to allow the delegate to specify a display location for the menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenudelegate/1518247-confinementrectformenu?language=objc
func (m_ MenuDelegateWrapper) ConfinementRectForMenuOnScreen(menu IMenu, screen IScreen) foundation.Rect {
	rv := objc.Call[foundation.Rect](m_, objc.Sel("confinementRectForMenu:onScreen:"), objc.Ptr(menu), objc.Ptr(screen))
	return rv
}

func (m_ MenuDelegateWrapper) HasMenuWillOpen() bool {
	return m_.RespondsToSelector(objc.Sel("menuWillOpen:"))
}

// Invoked when a menu is about to open. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenudelegate/1518156-menuwillopen?language=objc
func (m_ MenuDelegateWrapper) MenuWillOpen(menu IMenu) {
	objc.Call[objc.Void](m_, objc.Sel("menuWillOpen:"), objc.Ptr(menu))
}
