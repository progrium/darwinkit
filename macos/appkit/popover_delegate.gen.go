// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type IPopoverDelegate interface {
	ImplementsDetachableWindowForPopover() bool
	// optional
	DetachableWindowForPopover(popover Popover) IWindow
	ImplementsPopoverShouldClose() bool
	// optional
	PopoverShouldClose(popover Popover) bool
	ImplementsPopoverWillShow() bool
	// optional
	PopoverWillShow(notification foundation.Notification)
	ImplementsPopoverDidShow() bool
	// optional
	PopoverDidShow(notification foundation.Notification)
	ImplementsPopoverWillClose() bool
	// optional
	PopoverWillClose(notification foundation.Notification)
	ImplementsPopoverDidClose() bool
	// optional
	PopoverDidClose(notification foundation.Notification)
	ImplementsPopoverDidDetach() bool
	// optional
	PopoverDidDetach(popover Popover)
	ImplementsPopoverShouldDetach() bool
	// optional
	PopoverShouldDetach(popover Popover) bool
}

type PopoverDelegate struct {
	_DetachableWindowForPopover func(popover Popover) IWindow
	_PopoverShouldClose         func(popover Popover) bool
	_PopoverWillShow            func(notification foundation.Notification)
	_PopoverDidShow             func(notification foundation.Notification)
	_PopoverWillClose           func(notification foundation.Notification)
	_PopoverDidClose            func(notification foundation.Notification)
	_PopoverDidDetach           func(popover Popover)
	_PopoverShouldDetach        func(popover Popover) bool
}

func (di *PopoverDelegate) ImplementsDetachableWindowForPopover() bool {
	return di._DetachableWindowForPopover != nil
}

func (di *PopoverDelegate) SetDetachableWindowForPopover(f func(popover Popover) IWindow) {
	di._DetachableWindowForPopover = f
}

func (di *PopoverDelegate) DetachableWindowForPopover(popover Popover) IWindow {
	return di._DetachableWindowForPopover(popover)
}
func (di *PopoverDelegate) ImplementsPopoverShouldClose() bool {
	return di._PopoverShouldClose != nil
}

func (di *PopoverDelegate) SetPopoverShouldClose(f func(popover Popover) bool) {
	di._PopoverShouldClose = f
}

func (di *PopoverDelegate) PopoverShouldClose(popover Popover) bool {
	return di._PopoverShouldClose(popover)
}
func (di *PopoverDelegate) ImplementsPopoverWillShow() bool {
	return di._PopoverWillShow != nil
}

func (di *PopoverDelegate) SetPopoverWillShow(f func(notification foundation.Notification)) {
	di._PopoverWillShow = f
}

func (di *PopoverDelegate) PopoverWillShow(notification foundation.Notification) {
	di._PopoverWillShow(notification)
}
func (di *PopoverDelegate) ImplementsPopoverDidShow() bool {
	return di._PopoverDidShow != nil
}

func (di *PopoverDelegate) SetPopoverDidShow(f func(notification foundation.Notification)) {
	di._PopoverDidShow = f
}

func (di *PopoverDelegate) PopoverDidShow(notification foundation.Notification) {
	di._PopoverDidShow(notification)
}
func (di *PopoverDelegate) ImplementsPopoverWillClose() bool {
	return di._PopoverWillClose != nil
}

func (di *PopoverDelegate) SetPopoverWillClose(f func(notification foundation.Notification)) {
	di._PopoverWillClose = f
}

func (di *PopoverDelegate) PopoverWillClose(notification foundation.Notification) {
	di._PopoverWillClose(notification)
}
func (di *PopoverDelegate) ImplementsPopoverDidClose() bool {
	return di._PopoverDidClose != nil
}

func (di *PopoverDelegate) SetPopoverDidClose(f func(notification foundation.Notification)) {
	di._PopoverDidClose = f
}

func (di *PopoverDelegate) PopoverDidClose(notification foundation.Notification) {
	di._PopoverDidClose(notification)
}
func (di *PopoverDelegate) ImplementsPopoverDidDetach() bool {
	return di._PopoverDidDetach != nil
}

func (di *PopoverDelegate) SetPopoverDidDetach(f func(popover Popover)) {
	di._PopoverDidDetach = f
}

func (di *PopoverDelegate) PopoverDidDetach(popover Popover) {
	di._PopoverDidDetach(popover)
}
func (di *PopoverDelegate) ImplementsPopoverShouldDetach() bool {
	return di._PopoverShouldDetach != nil
}

func (di *PopoverDelegate) SetPopoverShouldDetach(f func(popover Popover) bool) {
	di._PopoverShouldDetach = f
}

func (di *PopoverDelegate) PopoverShouldDetach(popover Popover) bool {
	return di._PopoverShouldDetach(popover)
}

type PopoverDelegateWrapper struct {
	objc.Object
}

func (p_ PopoverDelegateWrapper) ImplementsDetachableWindowForPopover() bool {
	return p_.RespondsToSelector(objc.GetSelector("detachableWindowForPopover:"))
}

func (p_ PopoverDelegateWrapper) DetachableWindowForPopover(popover IPopover) Window {
	rv := objc.CallMethod[Window](p_, objc.GetSelector("detachableWindowForPopover:"), objc.ExtractPtr(popover))
	return rv
}

func (p_ PopoverDelegateWrapper) ImplementsPopoverShouldClose() bool {
	return p_.RespondsToSelector(objc.GetSelector("popoverShouldClose:"))
}

func (p_ PopoverDelegateWrapper) PopoverShouldClose(popover IPopover) bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("popoverShouldClose:"), objc.ExtractPtr(popover))
	return rv
}

func (p_ PopoverDelegateWrapper) ImplementsPopoverWillShow() bool {
	return p_.RespondsToSelector(objc.GetSelector("popoverWillShow:"))
}

func (p_ PopoverDelegateWrapper) PopoverWillShow(notification foundation.INotification) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("popoverWillShow:"), objc.ExtractPtr(notification))
}

func (p_ PopoverDelegateWrapper) ImplementsPopoverDidShow() bool {
	return p_.RespondsToSelector(objc.GetSelector("popoverDidShow:"))
}

func (p_ PopoverDelegateWrapper) PopoverDidShow(notification foundation.INotification) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("popoverDidShow:"), objc.ExtractPtr(notification))
}

func (p_ PopoverDelegateWrapper) ImplementsPopoverWillClose() bool {
	return p_.RespondsToSelector(objc.GetSelector("popoverWillClose:"))
}

func (p_ PopoverDelegateWrapper) PopoverWillClose(notification foundation.INotification) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("popoverWillClose:"), objc.ExtractPtr(notification))
}

func (p_ PopoverDelegateWrapper) ImplementsPopoverDidClose() bool {
	return p_.RespondsToSelector(objc.GetSelector("popoverDidClose:"))
}

func (p_ PopoverDelegateWrapper) PopoverDidClose(notification foundation.INotification) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("popoverDidClose:"), objc.ExtractPtr(notification))
}

func (p_ PopoverDelegateWrapper) ImplementsPopoverDidDetach() bool {
	return p_.RespondsToSelector(objc.GetSelector("popoverDidDetach:"))
}

func (p_ PopoverDelegateWrapper) PopoverDidDetach(popover IPopover) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("popoverDidDetach:"), objc.ExtractPtr(popover))
}

func (p_ PopoverDelegateWrapper) ImplementsPopoverShouldDetach() bool {
	return p_.RespondsToSelector(objc.GetSelector("popoverShouldDetach:"))
}

func (p_ PopoverDelegateWrapper) PopoverShouldDetach(popover IPopover) bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("popoverShouldDetach:"), objc.ExtractPtr(popover))
	return rv
}
