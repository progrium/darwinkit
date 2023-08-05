// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/uniformtypeidentifiers"
	"github.com/progrium/macdriver/objc"
)

var SavePanelClass = _SavePanelClass{objc.GetClass("NSSavePanel")}

type _SavePanelClass struct {
	objc.Class
}

type ISavePanel interface {
	IPanel
	BeginSheetModalForWindowCompletionHandler(window IWindow, handler func(result ModalResponse))
	BeginWithCompletionHandler(handler func(result ModalResponse))
	RunModal() ModalResponse
	ValidateVisibleColumns()
	Ok(sender objc.IObject)
	Cancel(sender objc.IObject)
	URL() foundation.URL
	Prompt() string
	SetPrompt(value string)
	Message() string
	SetMessage(value string)
	NameFieldLabel() string
	SetNameFieldLabel(value string)
	NameFieldStringValue() string
	SetNameFieldStringValue(value string)
	DirectoryURL() foundation.URL
	SetDirectoryURL(value foundation.IURL)
	AccessoryView() View
	SetAccessoryView(value IView)
	ShowsTagField() bool
	SetShowsTagField(value bool)
	TagNames() []string
	SetTagNames(value []string)
	CanCreateDirectories() bool
	SetCanCreateDirectories(value bool)
	CanSelectHiddenExtension() bool
	SetCanSelectHiddenExtension(value bool)
	ShowsHiddenFiles() bool
	SetShowsHiddenFiles(value bool)
	IsExtensionHidden() bool
	SetExtensionHidden(value bool)
	IsExpanded() bool
	AllowedContentTypes() []uniformtypeidentifiers.Type
	SetAllowedContentTypes(value []uniformtypeidentifiers.IType)
	AllowsOtherFileTypes() bool
	SetAllowsOtherFileTypes(value bool)
	TreatsFilePackagesAsDirectories() bool
	SetTreatsFilePackagesAsDirectories(value bool)
}

type SavePanel struct {
	Panel
}

func MakeSavePanel(ptr unsafe.Pointer) SavePanel {
	return SavePanel{
		Panel: MakePanel(ptr),
	}
}

func (sc _SavePanelClass) WindowWithContentViewController(contentViewController IViewController) SavePanel {
	rv := objc.CallMethod[SavePanel](sc, objc.GetSelector("windowWithContentViewController:"), objc.ExtractPtr(contentViewController))
	return rv
}

func SavePanel_WindowWithContentViewController(contentViewController IViewController) SavePanel {
	return SavePanelClass.WindowWithContentViewController(contentViewController)
}

func (s_ SavePanel) InitWithContentRectStyleMaskBackingDefer(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) SavePanel {
	rv := objc.CallMethod[SavePanel](s_, objc.GetSelector("initWithContentRect:styleMask:backing:defer:"), contentRect, style, backingStoreType, flag)
	return rv
}

func SavePanel_InitWithContentRectStyleMaskBackingDefer(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) SavePanel {
	return SavePanelClass.Alloc().InitWithContentRectStyleMaskBackingDefer(contentRect, style, backingStoreType, flag)
}

func (s_ SavePanel) InitWithContentRectStyleMaskBackingDeferScreen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) SavePanel {
	rv := objc.CallMethod[SavePanel](s_, objc.GetSelector("initWithContentRect:styleMask:backing:defer:screen:"), contentRect, style, backingStoreType, flag, objc.ExtractPtr(screen))
	return rv
}

func SavePanel_InitWithContentRectStyleMaskBackingDeferScreen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) SavePanel {
	return SavePanelClass.Alloc().InitWithContentRectStyleMaskBackingDeferScreen(contentRect, style, backingStoreType, flag, screen)
}

func (s_ SavePanel) Init() SavePanel {
	rv := objc.CallMethod[SavePanel](s_, objc.GetSelector("init"))
	return rv
}

func SavePanel_Init() SavePanel {
	return SavePanelClass.Alloc().Init()
}

func (sc _SavePanelClass) Alloc() SavePanel {
	rv := objc.CallMethod[SavePanel](sc, objc.GetSelector("alloc"))
	return rv
}

func SavePanel_Alloc() SavePanel {
	return SavePanelClass.Alloc()
}

func (sc _SavePanelClass) New() SavePanel {
	rv := objc.CallMethod[SavePanel](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewSavePanel() SavePanel {
	return SavePanelClass.New()
}

func SavePanel_New() SavePanel {
	return SavePanelClass.New()
}

func (sc _SavePanelClass) SavePanel() SavePanel {
	rv := objc.CallMethod[SavePanel](sc, objc.GetSelector("savePanel"))
	return rv
}

func SavePanel_SavePanel() SavePanel {
	return SavePanelClass.SavePanel()
}

func (s_ SavePanel) BeginSheetModalForWindowCompletionHandler(window IWindow, handler func(result ModalResponse)) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("beginSheetModalForWindow:completionHandler:"), objc.ExtractPtr(window), handler)
}

func (s_ SavePanel) BeginWithCompletionHandler(handler func(result ModalResponse)) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("beginWithCompletionHandler:"), handler)
}

func (s_ SavePanel) RunModal() ModalResponse {
	rv := objc.CallMethod[ModalResponse](s_, objc.GetSelector("runModal"))
	return rv
}

func (s_ SavePanel) ValidateVisibleColumns() {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("validateVisibleColumns"))
}

func (s_ SavePanel) Ok(sender objc.IObject) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("ok:"), objc.ExtractPtr(sender))
}

func (s_ SavePanel) Cancel(sender objc.IObject) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("cancel:"), objc.ExtractPtr(sender))
}

func (s_ SavePanel) URL() foundation.URL {
	rv := objc.CallMethod[foundation.URL](s_, objc.GetSelector("URL"))
	return rv
}

func (s_ SavePanel) Prompt() string {
	rv := objc.CallMethod[string](s_, objc.GetSelector("prompt"))
	return rv
}

func (s_ SavePanel) SetPrompt(value string) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setPrompt:"), value)
}

func (s_ SavePanel) Message() string {
	rv := objc.CallMethod[string](s_, objc.GetSelector("message"))
	return rv
}

func (s_ SavePanel) SetMessage(value string) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setMessage:"), value)
}

func (s_ SavePanel) NameFieldLabel() string {
	rv := objc.CallMethod[string](s_, objc.GetSelector("nameFieldLabel"))
	return rv
}

func (s_ SavePanel) SetNameFieldLabel(value string) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setNameFieldLabel:"), value)
}

func (s_ SavePanel) NameFieldStringValue() string {
	rv := objc.CallMethod[string](s_, objc.GetSelector("nameFieldStringValue"))
	return rv
}

func (s_ SavePanel) SetNameFieldStringValue(value string) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setNameFieldStringValue:"), value)
}

func (s_ SavePanel) DirectoryURL() foundation.URL {
	rv := objc.CallMethod[foundation.URL](s_, objc.GetSelector("directoryURL"))
	return rv
}

func (s_ SavePanel) SetDirectoryURL(value foundation.IURL) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setDirectoryURL:"), objc.ExtractPtr(value))
}

func (s_ SavePanel) AccessoryView() View {
	rv := objc.CallMethod[View](s_, objc.GetSelector("accessoryView"))
	return rv
}

func (s_ SavePanel) SetAccessoryView(value IView) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setAccessoryView:"), objc.ExtractPtr(value))
}

func (s_ SavePanel) ShowsTagField() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("showsTagField"))
	return rv
}

func (s_ SavePanel) SetShowsTagField(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setShowsTagField:"), value)
}

func (s_ SavePanel) TagNames() []string {
	rv := objc.CallMethod[[]string](s_, objc.GetSelector("tagNames"))
	return rv
}

func (s_ SavePanel) SetTagNames(value []string) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setTagNames:"), value)
}

func (s_ SavePanel) CanCreateDirectories() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("canCreateDirectories"))
	return rv
}

func (s_ SavePanel) SetCanCreateDirectories(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setCanCreateDirectories:"), value)
}

func (s_ SavePanel) CanSelectHiddenExtension() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("canSelectHiddenExtension"))
	return rv
}

func (s_ SavePanel) SetCanSelectHiddenExtension(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setCanSelectHiddenExtension:"), value)
}

func (s_ SavePanel) ShowsHiddenFiles() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("showsHiddenFiles"))
	return rv
}

func (s_ SavePanel) SetShowsHiddenFiles(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setShowsHiddenFiles:"), value)
}

func (s_ SavePanel) IsExtensionHidden() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("isExtensionHidden"))
	return rv
}

func (s_ SavePanel) SetExtensionHidden(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setExtensionHidden:"), value)
}

func (s_ SavePanel) IsExpanded() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("isExpanded"))
	return rv
}

func (s_ SavePanel) AllowedContentTypes() []uniformtypeidentifiers.Type {
	rv := objc.CallMethod[[]uniformtypeidentifiers.Type](s_, objc.GetSelector("allowedContentTypes"))
	return rv
}

func (s_ SavePanel) SetAllowedContentTypes(value []uniformtypeidentifiers.IType) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setAllowedContentTypes:"), value)
}

func (s_ SavePanel) AllowsOtherFileTypes() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("allowsOtherFileTypes"))
	return rv
}

func (s_ SavePanel) SetAllowsOtherFileTypes(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setAllowsOtherFileTypes:"), value)
}

func (s_ SavePanel) TreatsFilePackagesAsDirectories() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("treatsFilePackagesAsDirectories"))
	return rv
}

func (s_ SavePanel) SetTreatsFilePackagesAsDirectories(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setTreatsFilePackagesAsDirectories:"), value)
}
