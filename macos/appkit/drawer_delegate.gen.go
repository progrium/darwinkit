// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type IDrawerDelegate interface {
	ImplementsDrawerShouldOpen() bool
	// optional
	// deprecated
	DrawerShouldOpen(sender Drawer) bool
	ImplementsDrawerWillOpen() bool
	// optional
	// deprecated
	DrawerWillOpen(notification foundation.Notification)
	ImplementsDrawerDidOpen() bool
	// optional
	// deprecated
	DrawerDidOpen(notification foundation.Notification)
	ImplementsDrawerShouldClose() bool
	// optional
	// deprecated
	DrawerShouldClose(sender Drawer) bool
	ImplementsDrawerWillClose() bool
	// optional
	// deprecated
	DrawerWillClose(notification foundation.Notification)
	ImplementsDrawerDidClose() bool
	// optional
	// deprecated
	DrawerDidClose(notification foundation.Notification)
	ImplementsDrawerWillResizeContentsToSize() bool
	// optional
	// deprecated
	DrawerWillResizeContentsToSize(sender Drawer, contentSize foundation.Size) foundation.Size
}

type DrawerDelegate struct {
	_DrawerShouldOpen               func(sender Drawer) bool
	_DrawerWillOpen                 func(notification foundation.Notification)
	_DrawerDidOpen                  func(notification foundation.Notification)
	_DrawerShouldClose              func(sender Drawer) bool
	_DrawerWillClose                func(notification foundation.Notification)
	_DrawerDidClose                 func(notification foundation.Notification)
	_DrawerWillResizeContentsToSize func(sender Drawer, contentSize foundation.Size) foundation.Size
}

func (di *DrawerDelegate) ImplementsDrawerShouldOpen() bool {
	return di._DrawerShouldOpen != nil
}

// deprecated
func (di *DrawerDelegate) SetDrawerShouldOpen(f func(sender Drawer) bool) {
	di._DrawerShouldOpen = f
}

// deprecated
func (di *DrawerDelegate) DrawerShouldOpen(sender Drawer) bool {
	return di._DrawerShouldOpen(sender)
}
func (di *DrawerDelegate) ImplementsDrawerWillOpen() bool {
	return di._DrawerWillOpen != nil
}

// deprecated
func (di *DrawerDelegate) SetDrawerWillOpen(f func(notification foundation.Notification)) {
	di._DrawerWillOpen = f
}

// deprecated
func (di *DrawerDelegate) DrawerWillOpen(notification foundation.Notification) {
	di._DrawerWillOpen(notification)
}
func (di *DrawerDelegate) ImplementsDrawerDidOpen() bool {
	return di._DrawerDidOpen != nil
}

// deprecated
func (di *DrawerDelegate) SetDrawerDidOpen(f func(notification foundation.Notification)) {
	di._DrawerDidOpen = f
}

// deprecated
func (di *DrawerDelegate) DrawerDidOpen(notification foundation.Notification) {
	di._DrawerDidOpen(notification)
}
func (di *DrawerDelegate) ImplementsDrawerShouldClose() bool {
	return di._DrawerShouldClose != nil
}

// deprecated
func (di *DrawerDelegate) SetDrawerShouldClose(f func(sender Drawer) bool) {
	di._DrawerShouldClose = f
}

// deprecated
func (di *DrawerDelegate) DrawerShouldClose(sender Drawer) bool {
	return di._DrawerShouldClose(sender)
}
func (di *DrawerDelegate) ImplementsDrawerWillClose() bool {
	return di._DrawerWillClose != nil
}

// deprecated
func (di *DrawerDelegate) SetDrawerWillClose(f func(notification foundation.Notification)) {
	di._DrawerWillClose = f
}

// deprecated
func (di *DrawerDelegate) DrawerWillClose(notification foundation.Notification) {
	di._DrawerWillClose(notification)
}
func (di *DrawerDelegate) ImplementsDrawerDidClose() bool {
	return di._DrawerDidClose != nil
}

// deprecated
func (di *DrawerDelegate) SetDrawerDidClose(f func(notification foundation.Notification)) {
	di._DrawerDidClose = f
}

// deprecated
func (di *DrawerDelegate) DrawerDidClose(notification foundation.Notification) {
	di._DrawerDidClose(notification)
}
func (di *DrawerDelegate) ImplementsDrawerWillResizeContentsToSize() bool {
	return di._DrawerWillResizeContentsToSize != nil
}

// deprecated
func (di *DrawerDelegate) SetDrawerWillResizeContentsToSize(f func(sender Drawer, contentSize foundation.Size) foundation.Size) {
	di._DrawerWillResizeContentsToSize = f
}

// deprecated
func (di *DrawerDelegate) DrawerWillResizeContentsToSize(sender Drawer, contentSize foundation.Size) foundation.Size {
	return di._DrawerWillResizeContentsToSize(sender, contentSize)
}

type DrawerDelegateWrapper struct {
	objc.Object
}

func (d_ DrawerDelegateWrapper) ImplementsDrawerShouldOpen() bool {
	return d_.RespondsToSelector(objc.GetSelector("drawerShouldOpen:"))
}

func (d_ DrawerDelegateWrapper) ImplementsDrawerWillOpen() bool {
	return d_.RespondsToSelector(objc.GetSelector("drawerWillOpen:"))
}

func (d_ DrawerDelegateWrapper) ImplementsDrawerDidOpen() bool {
	return d_.RespondsToSelector(objc.GetSelector("drawerDidOpen:"))
}

func (d_ DrawerDelegateWrapper) ImplementsDrawerShouldClose() bool {
	return d_.RespondsToSelector(objc.GetSelector("drawerShouldClose:"))
}

func (d_ DrawerDelegateWrapper) ImplementsDrawerWillClose() bool {
	return d_.RespondsToSelector(objc.GetSelector("drawerWillClose:"))
}

func (d_ DrawerDelegateWrapper) ImplementsDrawerDidClose() bool {
	return d_.RespondsToSelector(objc.GetSelector("drawerDidClose:"))
}

func (d_ DrawerDelegateWrapper) ImplementsDrawerWillResizeContentsToSize() bool {
	return d_.RespondsToSelector(objc.GetSelector("drawerWillResizeContents:toSize:"))
}
