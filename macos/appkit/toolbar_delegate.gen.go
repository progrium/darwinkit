// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of optional methods you use to configure the toolbar and respond to changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbardelegate?language=objc
type PToolbarDelegate interface {
	// optional
	ToolbarAllowedItemIdentifiers(toolbar Toolbar) []ToolbarItemIdentifier
	HasToolbarAllowedItemIdentifiers() bool

	// optional
	ToolbarWillAddItem(notification foundation.Notification)
	HasToolbarWillAddItem() bool

	// optional
	ToolbarItemForItemIdentifierWillBeInsertedIntoToolbar(toolbar Toolbar, itemIdentifier ToolbarItemIdentifier, flag bool) IToolbarItem
	HasToolbarItemForItemIdentifierWillBeInsertedIntoToolbar() bool

	// optional
	ToolbarDefaultItemIdentifiers(toolbar Toolbar) []ToolbarItemIdentifier
	HasToolbarDefaultItemIdentifiers() bool

	// optional
	ToolbarSelectableItemIdentifiers(toolbar Toolbar) []ToolbarItemIdentifier
	HasToolbarSelectableItemIdentifiers() bool

	// optional
	ToolbarDidRemoveItem(notification foundation.Notification)
	HasToolbarDidRemoveItem() bool
}

// A delegate implementation builder for the [PToolbarDelegate] protocol.
type ToolbarDelegate struct {
	_ToolbarAllowedItemIdentifiers                         func(toolbar Toolbar) []ToolbarItemIdentifier
	_ToolbarWillAddItem                                    func(notification foundation.Notification)
	_ToolbarItemForItemIdentifierWillBeInsertedIntoToolbar func(toolbar Toolbar, itemIdentifier ToolbarItemIdentifier, flag bool) IToolbarItem
	_ToolbarDefaultItemIdentifiers                         func(toolbar Toolbar) []ToolbarItemIdentifier
	_ToolbarSelectableItemIdentifiers                      func(toolbar Toolbar) []ToolbarItemIdentifier
	_ToolbarDidRemoveItem                                  func(notification foundation.Notification)
}

func (di *ToolbarDelegate) HasToolbarAllowedItemIdentifiers() bool {
	return di._ToolbarAllowedItemIdentifiers != nil
}

// Asks the delegate to provide the items allowed on the toolbar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbardelegate/1516995-toolbaralloweditemidentifiers?language=objc
func (di *ToolbarDelegate) SetToolbarAllowedItemIdentifiers(f func(toolbar Toolbar) []ToolbarItemIdentifier) {
	di._ToolbarAllowedItemIdentifiers = f
}

// Asks the delegate to provide the items allowed on the toolbar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbardelegate/1516995-toolbaralloweditemidentifiers?language=objc
func (di *ToolbarDelegate) ToolbarAllowedItemIdentifiers(toolbar Toolbar) []ToolbarItemIdentifier {
	return di._ToolbarAllowedItemIdentifiers(toolbar)
}
func (di *ToolbarDelegate) HasToolbarWillAddItem() bool {
	return di._ToolbarWillAddItem != nil
}

// Tells the delegate that the toolbar is about to add the specified item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbardelegate/1516964-toolbarwilladditem?language=objc
func (di *ToolbarDelegate) SetToolbarWillAddItem(f func(notification foundation.Notification)) {
	di._ToolbarWillAddItem = f
}

// Tells the delegate that the toolbar is about to add the specified item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbardelegate/1516964-toolbarwilladditem?language=objc
func (di *ToolbarDelegate) ToolbarWillAddItem(notification foundation.Notification) {
	di._ToolbarWillAddItem(notification)
}
func (di *ToolbarDelegate) HasToolbarItemForItemIdentifierWillBeInsertedIntoToolbar() bool {
	return di._ToolbarItemForItemIdentifierWillBeInsertedIntoToolbar != nil
}

// Asks the delegate for the toolbar item associated with the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbardelegate/1516985-toolbar?language=objc
func (di *ToolbarDelegate) SetToolbarItemForItemIdentifierWillBeInsertedIntoToolbar(f func(toolbar Toolbar, itemIdentifier ToolbarItemIdentifier, flag bool) IToolbarItem) {
	di._ToolbarItemForItemIdentifierWillBeInsertedIntoToolbar = f
}

// Asks the delegate for the toolbar item associated with the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbardelegate/1516985-toolbar?language=objc
func (di *ToolbarDelegate) ToolbarItemForItemIdentifierWillBeInsertedIntoToolbar(toolbar Toolbar, itemIdentifier ToolbarItemIdentifier, flag bool) IToolbarItem {
	return di._ToolbarItemForItemIdentifierWillBeInsertedIntoToolbar(toolbar, itemIdentifier, flag)
}
func (di *ToolbarDelegate) HasToolbarDefaultItemIdentifiers() bool {
	return di._ToolbarDefaultItemIdentifiers != nil
}

// Asks the delegate to provide the default items to display on the toolbar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbardelegate/1516944-toolbardefaultitemidentifiers?language=objc
func (di *ToolbarDelegate) SetToolbarDefaultItemIdentifiers(f func(toolbar Toolbar) []ToolbarItemIdentifier) {
	di._ToolbarDefaultItemIdentifiers = f
}

// Asks the delegate to provide the default items to display on the toolbar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbardelegate/1516944-toolbardefaultitemidentifiers?language=objc
func (di *ToolbarDelegate) ToolbarDefaultItemIdentifiers(toolbar Toolbar) []ToolbarItemIdentifier {
	return di._ToolbarDefaultItemIdentifiers(toolbar)
}
func (di *ToolbarDelegate) HasToolbarSelectableItemIdentifiers() bool {
	return di._ToolbarSelectableItemIdentifiers != nil
}

// Asks the delegate to provide the set of selectable items in the toolbar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbardelegate/1516981-toolbarselectableitemidentifiers?language=objc
func (di *ToolbarDelegate) SetToolbarSelectableItemIdentifiers(f func(toolbar Toolbar) []ToolbarItemIdentifier) {
	di._ToolbarSelectableItemIdentifiers = f
}

// Asks the delegate to provide the set of selectable items in the toolbar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbardelegate/1516981-toolbarselectableitemidentifiers?language=objc
func (di *ToolbarDelegate) ToolbarSelectableItemIdentifiers(toolbar Toolbar) []ToolbarItemIdentifier {
	return di._ToolbarSelectableItemIdentifiers(toolbar)
}
func (di *ToolbarDelegate) HasToolbarDidRemoveItem() bool {
	return di._ToolbarDidRemoveItem != nil
}

// Tells the delegate that the toolbar removed the specified item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbardelegate/1516970-toolbardidremoveitem?language=objc
func (di *ToolbarDelegate) SetToolbarDidRemoveItem(f func(notification foundation.Notification)) {
	di._ToolbarDidRemoveItem = f
}

// Tells the delegate that the toolbar removed the specified item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbardelegate/1516970-toolbardidremoveitem?language=objc
func (di *ToolbarDelegate) ToolbarDidRemoveItem(notification foundation.Notification) {
	di._ToolbarDidRemoveItem(notification)
}

// A concrete type wrapper for the [PToolbarDelegate] protocol.
type ToolbarDelegateWrapper struct {
	objc.Object
}

func (t_ ToolbarDelegateWrapper) HasToolbarAllowedItemIdentifiers() bool {
	return t_.RespondsToSelector(objc.Sel("toolbarAllowedItemIdentifiers:"))
}

// Asks the delegate to provide the items allowed on the toolbar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbardelegate/1516995-toolbaralloweditemidentifiers?language=objc
func (t_ ToolbarDelegateWrapper) ToolbarAllowedItemIdentifiers(toolbar IToolbar) []ToolbarItemIdentifier {
	rv := objc.Call[[]ToolbarItemIdentifier](t_, objc.Sel("toolbarAllowedItemIdentifiers:"), objc.Ptr(toolbar))
	return rv
}

func (t_ ToolbarDelegateWrapper) HasToolbarWillAddItem() bool {
	return t_.RespondsToSelector(objc.Sel("toolbarWillAddItem:"))
}

// Tells the delegate that the toolbar is about to add the specified item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbardelegate/1516964-toolbarwilladditem?language=objc
func (t_ ToolbarDelegateWrapper) ToolbarWillAddItem(notification foundation.INotification) {
	objc.Call[objc.Void](t_, objc.Sel("toolbarWillAddItem:"), objc.Ptr(notification))
}

func (t_ ToolbarDelegateWrapper) HasToolbarItemForItemIdentifierWillBeInsertedIntoToolbar() bool {
	return t_.RespondsToSelector(objc.Sel("toolbar:itemForItemIdentifier:willBeInsertedIntoToolbar:"))
}

// Asks the delegate for the toolbar item associated with the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbardelegate/1516985-toolbar?language=objc
func (t_ ToolbarDelegateWrapper) ToolbarItemForItemIdentifierWillBeInsertedIntoToolbar(toolbar IToolbar, itemIdentifier ToolbarItemIdentifier, flag bool) ToolbarItem {
	rv := objc.Call[ToolbarItem](t_, objc.Sel("toolbar:itemForItemIdentifier:willBeInsertedIntoToolbar:"), objc.Ptr(toolbar), itemIdentifier, flag)
	return rv
}

func (t_ ToolbarDelegateWrapper) HasToolbarDefaultItemIdentifiers() bool {
	return t_.RespondsToSelector(objc.Sel("toolbarDefaultItemIdentifiers:"))
}

// Asks the delegate to provide the default items to display on the toolbar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbardelegate/1516944-toolbardefaultitemidentifiers?language=objc
func (t_ ToolbarDelegateWrapper) ToolbarDefaultItemIdentifiers(toolbar IToolbar) []ToolbarItemIdentifier {
	rv := objc.Call[[]ToolbarItemIdentifier](t_, objc.Sel("toolbarDefaultItemIdentifiers:"), objc.Ptr(toolbar))
	return rv
}

func (t_ ToolbarDelegateWrapper) HasToolbarSelectableItemIdentifiers() bool {
	return t_.RespondsToSelector(objc.Sel("toolbarSelectableItemIdentifiers:"))
}

// Asks the delegate to provide the set of selectable items in the toolbar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbardelegate/1516981-toolbarselectableitemidentifiers?language=objc
func (t_ ToolbarDelegateWrapper) ToolbarSelectableItemIdentifiers(toolbar IToolbar) []ToolbarItemIdentifier {
	rv := objc.Call[[]ToolbarItemIdentifier](t_, objc.Sel("toolbarSelectableItemIdentifiers:"), objc.Ptr(toolbar))
	return rv
}

func (t_ ToolbarDelegateWrapper) HasToolbarDidRemoveItem() bool {
	return t_.RespondsToSelector(objc.Sel("toolbarDidRemoveItem:"))
}

// Tells the delegate that the toolbar removed the specified item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbardelegate/1516970-toolbardidremoveitem?language=objc
func (t_ ToolbarDelegateWrapper) ToolbarDidRemoveItem(notification foundation.INotification) {
	objc.Call[objc.Void](t_, objc.Sel("toolbarDidRemoveItem:"), objc.Ptr(notification))
}
