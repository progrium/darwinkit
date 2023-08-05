// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type IToolbarDelegate interface {
	ImplementsToolbarItemForItemIdentifierWillBeInsertedIntoToolbar() bool
	// optional
	ToolbarItemForItemIdentifierWillBeInsertedIntoToolbar(toolbar Toolbar, itemIdentifier ToolbarItemIdentifier, flag bool) IToolbarItem
	ImplementsToolbarWillAddItem() bool
	// optional
	ToolbarWillAddItem(notification foundation.Notification)
	ImplementsToolbarDidRemoveItem() bool
	// optional
	ToolbarDidRemoveItem(notification foundation.Notification)
	ImplementsToolbarAllowedItemIdentifiers() bool
	// optional
	ToolbarAllowedItemIdentifiers(toolbar Toolbar) []ToolbarItemIdentifier
	ImplementsToolbarDefaultItemIdentifiers() bool
	// optional
	ToolbarDefaultItemIdentifiers(toolbar Toolbar) []ToolbarItemIdentifier
	ImplementsToolbarImmovableItemIdentifiers() bool
	// optional
	ToolbarImmovableItemIdentifiers(toolbar Toolbar) foundation.ISet
	ImplementsToolbarSelectableItemIdentifiers() bool
	// optional
	ToolbarSelectableItemIdentifiers(toolbar Toolbar) []ToolbarItemIdentifier
	ImplementsToolbarItemIdentifierCanBeInsertedAtIndex() bool
	// optional
	ToolbarItemIdentifierCanBeInsertedAtIndex(toolbar Toolbar, itemIdentifier ToolbarItemIdentifier, index int) bool
}

type ToolbarDelegate struct {
	_ToolbarItemForItemIdentifierWillBeInsertedIntoToolbar func(toolbar Toolbar, itemIdentifier ToolbarItemIdentifier, flag bool) IToolbarItem
	_ToolbarWillAddItem                                    func(notification foundation.Notification)
	_ToolbarDidRemoveItem                                  func(notification foundation.Notification)
	_ToolbarAllowedItemIdentifiers                         func(toolbar Toolbar) []ToolbarItemIdentifier
	_ToolbarDefaultItemIdentifiers                         func(toolbar Toolbar) []ToolbarItemIdentifier
	_ToolbarImmovableItemIdentifiers                       func(toolbar Toolbar) foundation.ISet
	_ToolbarSelectableItemIdentifiers                      func(toolbar Toolbar) []ToolbarItemIdentifier
	_ToolbarItemIdentifierCanBeInsertedAtIndex             func(toolbar Toolbar, itemIdentifier ToolbarItemIdentifier, index int) bool
}

func (di *ToolbarDelegate) ImplementsToolbarItemForItemIdentifierWillBeInsertedIntoToolbar() bool {
	return di._ToolbarItemForItemIdentifierWillBeInsertedIntoToolbar != nil
}

func (di *ToolbarDelegate) SetToolbarItemForItemIdentifierWillBeInsertedIntoToolbar(f func(toolbar Toolbar, itemIdentifier ToolbarItemIdentifier, flag bool) IToolbarItem) {
	di._ToolbarItemForItemIdentifierWillBeInsertedIntoToolbar = f
}

func (di *ToolbarDelegate) ToolbarItemForItemIdentifierWillBeInsertedIntoToolbar(toolbar Toolbar, itemIdentifier ToolbarItemIdentifier, flag bool) IToolbarItem {
	return di._ToolbarItemForItemIdentifierWillBeInsertedIntoToolbar(toolbar, itemIdentifier, flag)
}
func (di *ToolbarDelegate) ImplementsToolbarWillAddItem() bool {
	return di._ToolbarWillAddItem != nil
}

func (di *ToolbarDelegate) SetToolbarWillAddItem(f func(notification foundation.Notification)) {
	di._ToolbarWillAddItem = f
}

func (di *ToolbarDelegate) ToolbarWillAddItem(notification foundation.Notification) {
	di._ToolbarWillAddItem(notification)
}
func (di *ToolbarDelegate) ImplementsToolbarDidRemoveItem() bool {
	return di._ToolbarDidRemoveItem != nil
}

func (di *ToolbarDelegate) SetToolbarDidRemoveItem(f func(notification foundation.Notification)) {
	di._ToolbarDidRemoveItem = f
}

func (di *ToolbarDelegate) ToolbarDidRemoveItem(notification foundation.Notification) {
	di._ToolbarDidRemoveItem(notification)
}
func (di *ToolbarDelegate) ImplementsToolbarAllowedItemIdentifiers() bool {
	return di._ToolbarAllowedItemIdentifiers != nil
}

func (di *ToolbarDelegate) SetToolbarAllowedItemIdentifiers(f func(toolbar Toolbar) []ToolbarItemIdentifier) {
	di._ToolbarAllowedItemIdentifiers = f
}

func (di *ToolbarDelegate) ToolbarAllowedItemIdentifiers(toolbar Toolbar) []ToolbarItemIdentifier {
	return di._ToolbarAllowedItemIdentifiers(toolbar)
}
func (di *ToolbarDelegate) ImplementsToolbarDefaultItemIdentifiers() bool {
	return di._ToolbarDefaultItemIdentifiers != nil
}

func (di *ToolbarDelegate) SetToolbarDefaultItemIdentifiers(f func(toolbar Toolbar) []ToolbarItemIdentifier) {
	di._ToolbarDefaultItemIdentifiers = f
}

func (di *ToolbarDelegate) ToolbarDefaultItemIdentifiers(toolbar Toolbar) []ToolbarItemIdentifier {
	return di._ToolbarDefaultItemIdentifiers(toolbar)
}
func (di *ToolbarDelegate) ImplementsToolbarImmovableItemIdentifiers() bool {
	return di._ToolbarImmovableItemIdentifiers != nil
}

func (di *ToolbarDelegate) SetToolbarImmovableItemIdentifiers(f func(toolbar Toolbar) foundation.ISet) {
	di._ToolbarImmovableItemIdentifiers = f
}

func (di *ToolbarDelegate) ToolbarImmovableItemIdentifiers(toolbar Toolbar) foundation.ISet {
	return di._ToolbarImmovableItemIdentifiers(toolbar)
}
func (di *ToolbarDelegate) ImplementsToolbarSelectableItemIdentifiers() bool {
	return di._ToolbarSelectableItemIdentifiers != nil
}

func (di *ToolbarDelegate) SetToolbarSelectableItemIdentifiers(f func(toolbar Toolbar) []ToolbarItemIdentifier) {
	di._ToolbarSelectableItemIdentifiers = f
}

func (di *ToolbarDelegate) ToolbarSelectableItemIdentifiers(toolbar Toolbar) []ToolbarItemIdentifier {
	return di._ToolbarSelectableItemIdentifiers(toolbar)
}
func (di *ToolbarDelegate) ImplementsToolbarItemIdentifierCanBeInsertedAtIndex() bool {
	return di._ToolbarItemIdentifierCanBeInsertedAtIndex != nil
}

func (di *ToolbarDelegate) SetToolbarItemIdentifierCanBeInsertedAtIndex(f func(toolbar Toolbar, itemIdentifier ToolbarItemIdentifier, index int) bool) {
	di._ToolbarItemIdentifierCanBeInsertedAtIndex = f
}

func (di *ToolbarDelegate) ToolbarItemIdentifierCanBeInsertedAtIndex(toolbar Toolbar, itemIdentifier ToolbarItemIdentifier, index int) bool {
	return di._ToolbarItemIdentifierCanBeInsertedAtIndex(toolbar, itemIdentifier, index)
}

type ToolbarDelegateWrapper struct {
	objc.Object
}

func (t_ ToolbarDelegateWrapper) ImplementsToolbarItemForItemIdentifierWillBeInsertedIntoToolbar() bool {
	return t_.RespondsToSelector(objc.GetSelector("toolbar:itemForItemIdentifier:willBeInsertedIntoToolbar:"))
}

func (t_ ToolbarDelegateWrapper) ToolbarItemForItemIdentifierWillBeInsertedIntoToolbar(toolbar IToolbar, itemIdentifier ToolbarItemIdentifier, flag bool) ToolbarItem {
	rv := objc.CallMethod[ToolbarItem](t_, objc.GetSelector("toolbar:itemForItemIdentifier:willBeInsertedIntoToolbar:"), objc.ExtractPtr(toolbar), itemIdentifier, flag)
	return rv
}

func (t_ ToolbarDelegateWrapper) ImplementsToolbarWillAddItem() bool {
	return t_.RespondsToSelector(objc.GetSelector("toolbarWillAddItem:"))
}

func (t_ ToolbarDelegateWrapper) ToolbarWillAddItem(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("toolbarWillAddItem:"), objc.ExtractPtr(notification))
}

func (t_ ToolbarDelegateWrapper) ImplementsToolbarDidRemoveItem() bool {
	return t_.RespondsToSelector(objc.GetSelector("toolbarDidRemoveItem:"))
}

func (t_ ToolbarDelegateWrapper) ToolbarDidRemoveItem(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("toolbarDidRemoveItem:"), objc.ExtractPtr(notification))
}

func (t_ ToolbarDelegateWrapper) ImplementsToolbarAllowedItemIdentifiers() bool {
	return t_.RespondsToSelector(objc.GetSelector("toolbarAllowedItemIdentifiers:"))
}

func (t_ ToolbarDelegateWrapper) ToolbarAllowedItemIdentifiers(toolbar IToolbar) []ToolbarItemIdentifier {
	rv := objc.CallMethod[[]ToolbarItemIdentifier](t_, objc.GetSelector("toolbarAllowedItemIdentifiers:"), objc.ExtractPtr(toolbar))
	return rv
}

func (t_ ToolbarDelegateWrapper) ImplementsToolbarDefaultItemIdentifiers() bool {
	return t_.RespondsToSelector(objc.GetSelector("toolbarDefaultItemIdentifiers:"))
}

func (t_ ToolbarDelegateWrapper) ToolbarDefaultItemIdentifiers(toolbar IToolbar) []ToolbarItemIdentifier {
	rv := objc.CallMethod[[]ToolbarItemIdentifier](t_, objc.GetSelector("toolbarDefaultItemIdentifiers:"), objc.ExtractPtr(toolbar))
	return rv
}

func (t_ ToolbarDelegateWrapper) ImplementsToolbarImmovableItemIdentifiers() bool {
	return t_.RespondsToSelector(objc.GetSelector("toolbarImmovableItemIdentifiers:"))
}

func (t_ ToolbarDelegateWrapper) ToolbarImmovableItemIdentifiers(toolbar IToolbar) foundation.Set {
	rv := objc.CallMethod[foundation.Set](t_, objc.GetSelector("toolbarImmovableItemIdentifiers:"), objc.ExtractPtr(toolbar))
	return rv
}

func (t_ ToolbarDelegateWrapper) ImplementsToolbarSelectableItemIdentifiers() bool {
	return t_.RespondsToSelector(objc.GetSelector("toolbarSelectableItemIdentifiers:"))
}

func (t_ ToolbarDelegateWrapper) ToolbarSelectableItemIdentifiers(toolbar IToolbar) []ToolbarItemIdentifier {
	rv := objc.CallMethod[[]ToolbarItemIdentifier](t_, objc.GetSelector("toolbarSelectableItemIdentifiers:"), objc.ExtractPtr(toolbar))
	return rv
}

func (t_ ToolbarDelegateWrapper) ImplementsToolbarItemIdentifierCanBeInsertedAtIndex() bool {
	return t_.RespondsToSelector(objc.GetSelector("toolbar:itemIdentifier:canBeInsertedAtIndex:"))
}

func (t_ ToolbarDelegateWrapper) ToolbarItemIdentifierCanBeInsertedAtIndex(toolbar IToolbar, itemIdentifier ToolbarItemIdentifier, index int) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("toolbar:itemIdentifier:canBeInsertedAtIndex:"), objc.ExtractPtr(toolbar), itemIdentifier, index)
	return rv
}
