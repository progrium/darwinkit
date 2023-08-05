// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var OpenPanelClass = _OpenPanelClass{objc.GetClass("NSOpenPanel")}

type _OpenPanelClass struct {
	objc.Class
}

type IOpenPanel interface {
	ISavePanel
	CanChooseFiles() bool
	SetCanChooseFiles(value bool)
	CanChooseDirectories() bool
	SetCanChooseDirectories(value bool)
	ResolvesAliases() bool
	SetResolvesAliases(value bool)
	AllowsMultipleSelection() bool
	SetAllowsMultipleSelection(value bool)
	IsAccessoryViewDisclosed() bool
	SetAccessoryViewDisclosed(value bool)
	URLs() []foundation.URL
	CanDownloadUbiquitousContents() bool
	SetCanDownloadUbiquitousContents(value bool)
	CanResolveUbiquitousConflicts() bool
	SetCanResolveUbiquitousConflicts(value bool)
}

type OpenPanel struct {
	SavePanel
}

func MakeOpenPanel(ptr unsafe.Pointer) OpenPanel {
	return OpenPanel{
		SavePanel: MakeSavePanel(ptr),
	}
}

func (oc _OpenPanelClass) WindowWithContentViewController(contentViewController IViewController) OpenPanel {
	rv := objc.CallMethod[OpenPanel](oc, objc.GetSelector("windowWithContentViewController:"), objc.ExtractPtr(contentViewController))
	return rv
}

func OpenPanel_WindowWithContentViewController(contentViewController IViewController) OpenPanel {
	return OpenPanelClass.WindowWithContentViewController(contentViewController)
}

func (o_ OpenPanel) InitWithContentRectStyleMaskBackingDefer(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) OpenPanel {
	rv := objc.CallMethod[OpenPanel](o_, objc.GetSelector("initWithContentRect:styleMask:backing:defer:"), contentRect, style, backingStoreType, flag)
	return rv
}

func OpenPanel_InitWithContentRectStyleMaskBackingDefer(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) OpenPanel {
	return OpenPanelClass.Alloc().InitWithContentRectStyleMaskBackingDefer(contentRect, style, backingStoreType, flag)
}

func (o_ OpenPanel) InitWithContentRectStyleMaskBackingDeferScreen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) OpenPanel {
	rv := objc.CallMethod[OpenPanel](o_, objc.GetSelector("initWithContentRect:styleMask:backing:defer:screen:"), contentRect, style, backingStoreType, flag, objc.ExtractPtr(screen))
	return rv
}

func OpenPanel_InitWithContentRectStyleMaskBackingDeferScreen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) OpenPanel {
	return OpenPanelClass.Alloc().InitWithContentRectStyleMaskBackingDeferScreen(contentRect, style, backingStoreType, flag, screen)
}

func (o_ OpenPanel) Init() OpenPanel {
	rv := objc.CallMethod[OpenPanel](o_, objc.GetSelector("init"))
	return rv
}

func OpenPanel_Init() OpenPanel {
	return OpenPanelClass.Alloc().Init()
}

func (oc _OpenPanelClass) Alloc() OpenPanel {
	rv := objc.CallMethod[OpenPanel](oc, objc.GetSelector("alloc"))
	return rv
}

func OpenPanel_Alloc() OpenPanel {
	return OpenPanelClass.Alloc()
}

func (oc _OpenPanelClass) New() OpenPanel {
	rv := objc.CallMethod[OpenPanel](oc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewOpenPanel() OpenPanel {
	return OpenPanelClass.New()
}

func OpenPanel_New() OpenPanel {
	return OpenPanelClass.New()
}

func (oc _OpenPanelClass) OpenPanel() OpenPanel {
	rv := objc.CallMethod[OpenPanel](oc, objc.GetSelector("openPanel"))
	return rv
}

func OpenPanel_OpenPanel() OpenPanel {
	return OpenPanelClass.OpenPanel()
}

func (o_ OpenPanel) CanChooseFiles() bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("canChooseFiles"))
	return rv
}

func (o_ OpenPanel) SetCanChooseFiles(value bool) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("setCanChooseFiles:"), value)
}

func (o_ OpenPanel) CanChooseDirectories() bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("canChooseDirectories"))
	return rv
}

func (o_ OpenPanel) SetCanChooseDirectories(value bool) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("setCanChooseDirectories:"), value)
}

func (o_ OpenPanel) ResolvesAliases() bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("resolvesAliases"))
	return rv
}

func (o_ OpenPanel) SetResolvesAliases(value bool) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("setResolvesAliases:"), value)
}

func (o_ OpenPanel) AllowsMultipleSelection() bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("allowsMultipleSelection"))
	return rv
}

func (o_ OpenPanel) SetAllowsMultipleSelection(value bool) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("setAllowsMultipleSelection:"), value)
}

func (o_ OpenPanel) IsAccessoryViewDisclosed() bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("isAccessoryViewDisclosed"))
	return rv
}

func (o_ OpenPanel) SetAccessoryViewDisclosed(value bool) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("setAccessoryViewDisclosed:"), value)
}

func (o_ OpenPanel) URLs() []foundation.URL {
	rv := objc.CallMethod[[]foundation.URL](o_, objc.GetSelector("URLs"))
	return rv
}

func (o_ OpenPanel) CanDownloadUbiquitousContents() bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("canDownloadUbiquitousContents"))
	return rv
}

func (o_ OpenPanel) SetCanDownloadUbiquitousContents(value bool) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("setCanDownloadUbiquitousContents:"), value)
}

func (o_ OpenPanel) CanResolveUbiquitousConflicts() bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("canResolveUbiquitousConflicts"))
	return rv
}

func (o_ OpenPanel) SetCanResolveUbiquitousConflicts(value bool) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("setCanResolveUbiquitousConflicts:"), value)
}
