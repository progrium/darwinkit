// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of optional methods that a popover delegate can implement to provide additional or custom functionality. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate?language=objc
type PPopoverDelegate interface {
	// optional
	PopoverWillShow(notification foundation.Notification)
	HasPopoverWillShow() bool

	// optional
	PopoverWillClose(notification foundation.Notification)
	HasPopoverWillClose() bool

	// optional
	PopoverDidShow(notification foundation.Notification)
	HasPopoverDidShow() bool

	// optional
	PopoverDidDetach(popover Popover)
	HasPopoverDidDetach() bool

	// optional
	DetachableWindowForPopover(popover Popover) IWindow
	HasDetachableWindowForPopover() bool

	// optional
	PopoverDidClose(notification foundation.Notification)
	HasPopoverDidClose() bool

	// optional
	PopoverShouldClose(popover Popover) bool
	HasPopoverShouldClose() bool

	// optional
	PopoverShouldDetach(popover Popover) bool
	HasPopoverShouldDetach() bool
}

// A delegate implementation builder for the [PPopoverDelegate] protocol.
type PopoverDelegate struct {
	_PopoverWillShow            func(notification foundation.Notification)
	_PopoverWillClose           func(notification foundation.Notification)
	_PopoverDidShow             func(notification foundation.Notification)
	_PopoverDidDetach           func(popover Popover)
	_DetachableWindowForPopover func(popover Popover) IWindow
	_PopoverDidClose            func(notification foundation.Notification)
	_PopoverShouldClose         func(popover Popover) bool
	_PopoverShouldDetach        func(popover Popover) bool
}

func (di *PopoverDelegate) HasPopoverWillShow() bool {
	return di._PopoverWillShow != nil
}

// Invoked when the popover will show. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1532556-popoverwillshow?language=objc
func (di *PopoverDelegate) SetPopoverWillShow(f func(notification foundation.Notification)) {
	di._PopoverWillShow = f
}

// Invoked when the popover will show. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1532556-popoverwillshow?language=objc
func (di *PopoverDelegate) PopoverWillShow(notification foundation.Notification) {
	di._PopoverWillShow(notification)
}
func (di *PopoverDelegate) HasPopoverWillClose() bool {
	return di._PopoverWillClose != nil
}

// Invoked when the popover is about to close. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1535119-popoverwillclose?language=objc
func (di *PopoverDelegate) SetPopoverWillClose(f func(notification foundation.Notification)) {
	di._PopoverWillClose = f
}

// Invoked when the popover is about to close. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1535119-popoverwillclose?language=objc
func (di *PopoverDelegate) PopoverWillClose(notification foundation.Notification) {
	di._PopoverWillClose(notification)
}
func (di *PopoverDelegate) HasPopoverDidShow() bool {
	return di._PopoverDidShow != nil
}

// Invoked when the popover has been shown. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1533573-popoverdidshow?language=objc
func (di *PopoverDelegate) SetPopoverDidShow(f func(notification foundation.Notification)) {
	di._PopoverDidShow = f
}

// Invoked when the popover has been shown. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1533573-popoverdidshow?language=objc
func (di *PopoverDelegate) PopoverDidShow(notification foundation.Notification) {
	di._PopoverDidShow(notification)
}
func (di *PopoverDelegate) HasPopoverDidDetach() bool {
	return di._PopoverDidDetach != nil
}

// Indicates that a popover has been released while it's in an implicitly detached state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1524674-popoverdiddetach?language=objc
func (di *PopoverDelegate) SetPopoverDidDetach(f func(popover Popover)) {
	di._PopoverDidDetach = f
}

// Indicates that a popover has been released while it's in an implicitly detached state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1524674-popoverdiddetach?language=objc
func (di *PopoverDelegate) PopoverDidDetach(popover Popover) {
	di._PopoverDidDetach(popover)
}
func (di *PopoverDelegate) HasDetachableWindowForPopover() bool {
	return di._DetachableWindowForPopover != nil
}

// Detaches the popover creating a window containing the content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1534822-detachablewindowforpopover?language=objc
func (di *PopoverDelegate) SetDetachableWindowForPopover(f func(popover Popover) IWindow) {
	di._DetachableWindowForPopover = f
}

// Detaches the popover creating a window containing the content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1534822-detachablewindowforpopover?language=objc
func (di *PopoverDelegate) DetachableWindowForPopover(popover Popover) IWindow {
	return di._DetachableWindowForPopover(popover)
}
func (di *PopoverDelegate) HasPopoverDidClose() bool {
	return di._PopoverDidClose != nil
}

// Invoked when the popover did close. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1526581-popoverdidclose?language=objc
func (di *PopoverDelegate) SetPopoverDidClose(f func(notification foundation.Notification)) {
	di._PopoverDidClose = f
}

// Invoked when the popover did close. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1526581-popoverdidclose?language=objc
func (di *PopoverDelegate) PopoverDidClose(notification foundation.Notification) {
	di._PopoverDidClose(notification)
}
func (di *PopoverDelegate) HasPopoverShouldClose() bool {
	return di._PopoverShouldClose != nil
}

// Allows a delegate to override a close request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1532593-popovershouldclose?language=objc
func (di *PopoverDelegate) SetPopoverShouldClose(f func(popover Popover) bool) {
	di._PopoverShouldClose = f
}

// Allows a delegate to override a close request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1532593-popovershouldclose?language=objc
func (di *PopoverDelegate) PopoverShouldClose(popover Popover) bool {
	return di._PopoverShouldClose(popover)
}
func (di *PopoverDelegate) HasPopoverShouldDetach() bool {
	return di._PopoverShouldDetach != nil
}

// Returns a Boolean value that indicates whether a popover should detach from its positioning view and become a separate window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1529911-popovershoulddetach?language=objc
func (di *PopoverDelegate) SetPopoverShouldDetach(f func(popover Popover) bool) {
	di._PopoverShouldDetach = f
}

// Returns a Boolean value that indicates whether a popover should detach from its positioning view and become a separate window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1529911-popovershoulddetach?language=objc
func (di *PopoverDelegate) PopoverShouldDetach(popover Popover) bool {
	return di._PopoverShouldDetach(popover)
}

// A concrete type wrapper for the [PPopoverDelegate] protocol.
type PopoverDelegateWrapper struct {
	objc.Object
}

func (p_ PopoverDelegateWrapper) HasPopoverWillShow() bool {
	return p_.RespondsToSelector(objc.Sel("popoverWillShow:"))
}

// Invoked when the popover will show. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1532556-popoverwillshow?language=objc
func (p_ PopoverDelegateWrapper) PopoverWillShow(notification foundation.INotification) {
	objc.Call[objc.Void](p_, objc.Sel("popoverWillShow:"), objc.Ptr(notification))
}

func (p_ PopoverDelegateWrapper) HasPopoverWillClose() bool {
	return p_.RespondsToSelector(objc.Sel("popoverWillClose:"))
}

// Invoked when the popover is about to close. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1535119-popoverwillclose?language=objc
func (p_ PopoverDelegateWrapper) PopoverWillClose(notification foundation.INotification) {
	objc.Call[objc.Void](p_, objc.Sel("popoverWillClose:"), objc.Ptr(notification))
}

func (p_ PopoverDelegateWrapper) HasPopoverDidShow() bool {
	return p_.RespondsToSelector(objc.Sel("popoverDidShow:"))
}

// Invoked when the popover has been shown. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1533573-popoverdidshow?language=objc
func (p_ PopoverDelegateWrapper) PopoverDidShow(notification foundation.INotification) {
	objc.Call[objc.Void](p_, objc.Sel("popoverDidShow:"), objc.Ptr(notification))
}

func (p_ PopoverDelegateWrapper) HasPopoverDidDetach() bool {
	return p_.RespondsToSelector(objc.Sel("popoverDidDetach:"))
}

// Indicates that a popover has been released while it's in an implicitly detached state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1524674-popoverdiddetach?language=objc
func (p_ PopoverDelegateWrapper) PopoverDidDetach(popover IPopover) {
	objc.Call[objc.Void](p_, objc.Sel("popoverDidDetach:"), objc.Ptr(popover))
}

func (p_ PopoverDelegateWrapper) HasDetachableWindowForPopover() bool {
	return p_.RespondsToSelector(objc.Sel("detachableWindowForPopover:"))
}

// Detaches the popover creating a window containing the content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1534822-detachablewindowforpopover?language=objc
func (p_ PopoverDelegateWrapper) DetachableWindowForPopover(popover IPopover) Window {
	rv := objc.Call[Window](p_, objc.Sel("detachableWindowForPopover:"), objc.Ptr(popover))
	return rv
}

func (p_ PopoverDelegateWrapper) HasPopoverDidClose() bool {
	return p_.RespondsToSelector(objc.Sel("popoverDidClose:"))
}

// Invoked when the popover did close. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1526581-popoverdidclose?language=objc
func (p_ PopoverDelegateWrapper) PopoverDidClose(notification foundation.INotification) {
	objc.Call[objc.Void](p_, objc.Sel("popoverDidClose:"), objc.Ptr(notification))
}

func (p_ PopoverDelegateWrapper) HasPopoverShouldClose() bool {
	return p_.RespondsToSelector(objc.Sel("popoverShouldClose:"))
}

// Allows a delegate to override a close request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1532593-popovershouldclose?language=objc
func (p_ PopoverDelegateWrapper) PopoverShouldClose(popover IPopover) bool {
	rv := objc.Call[bool](p_, objc.Sel("popoverShouldClose:"), objc.Ptr(popover))
	return rv
}

func (p_ PopoverDelegateWrapper) HasPopoverShouldDetach() bool {
	return p_.RespondsToSelector(objc.Sel("popoverShouldDetach:"))
}

// Returns a Boolean value that indicates whether a popover should detach from its positioning view and become a separate window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopoverdelegate/1529911-popovershoulddetach?language=objc
func (p_ PopoverDelegateWrapper) PopoverShouldDetach(popover IPopover) bool {
	rv := objc.Call[bool](p_, objc.Sel("popoverShouldDetach:"), objc.Ptr(popover))
	return rv
}
