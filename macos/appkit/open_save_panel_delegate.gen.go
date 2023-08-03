// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type IOpenSavePanelDelegate interface {
	ImplementsPanelUserEnteredFilenameConfirmed() bool
	// optional
	PanelUserEnteredFilenameConfirmed(sender objc.Object, filename string, okFlag bool) string
	ImplementsPanelSelectionDidChange() bool
	// optional
	PanelSelectionDidChange(sender objc.Object)
	ImplementsPanelDidChangeToDirectoryURL() bool
	// optional
	PanelDidChangeToDirectoryURL(sender objc.Object, url foundation.URL)
	ImplementsPanelWillExpand() bool
	// optional
	PanelWillExpand(sender objc.Object, expanding bool)
	ImplementsPanelShouldEnableURL() bool
	// optional
	PanelShouldEnableURL(sender objc.Object, url foundation.URL) bool
	ImplementsPanelValidateURLError() bool
	// optional
	PanelValidateURLError(sender objc.Object, url foundation.URL, outError *foundation.Error) bool
}

type OpenSavePanelDelegate struct {
	_PanelUserEnteredFilenameConfirmed func(sender objc.Object, filename string, okFlag bool) string
	_PanelSelectionDidChange           func(sender objc.Object)
	_PanelDidChangeToDirectoryURL      func(sender objc.Object, url foundation.URL)
	_PanelWillExpand                   func(sender objc.Object, expanding bool)
	_PanelShouldEnableURL              func(sender objc.Object, url foundation.URL) bool
	_PanelValidateURLError             func(sender objc.Object, url foundation.URL, outError *foundation.Error) bool
}

func (di *OpenSavePanelDelegate) ImplementsPanelUserEnteredFilenameConfirmed() bool {
	return di._PanelUserEnteredFilenameConfirmed != nil
}

func (di *OpenSavePanelDelegate) SetPanelUserEnteredFilenameConfirmed(f func(sender objc.Object, filename string, okFlag bool) string) {
	di._PanelUserEnteredFilenameConfirmed = f
}

func (di *OpenSavePanelDelegate) PanelUserEnteredFilenameConfirmed(sender objc.Object, filename string, okFlag bool) string {
	return di._PanelUserEnteredFilenameConfirmed(sender, filename, okFlag)
}
func (di *OpenSavePanelDelegate) ImplementsPanelSelectionDidChange() bool {
	return di._PanelSelectionDidChange != nil
}

func (di *OpenSavePanelDelegate) SetPanelSelectionDidChange(f func(sender objc.Object)) {
	di._PanelSelectionDidChange = f
}

func (di *OpenSavePanelDelegate) PanelSelectionDidChange(sender objc.Object) {
	di._PanelSelectionDidChange(sender)
}
func (di *OpenSavePanelDelegate) ImplementsPanelDidChangeToDirectoryURL() bool {
	return di._PanelDidChangeToDirectoryURL != nil
}

func (di *OpenSavePanelDelegate) SetPanelDidChangeToDirectoryURL(f func(sender objc.Object, url foundation.URL)) {
	di._PanelDidChangeToDirectoryURL = f
}

func (di *OpenSavePanelDelegate) PanelDidChangeToDirectoryURL(sender objc.Object, url foundation.URL) {
	di._PanelDidChangeToDirectoryURL(sender, url)
}
func (di *OpenSavePanelDelegate) ImplementsPanelWillExpand() bool {
	return di._PanelWillExpand != nil
}

func (di *OpenSavePanelDelegate) SetPanelWillExpand(f func(sender objc.Object, expanding bool)) {
	di._PanelWillExpand = f
}

func (di *OpenSavePanelDelegate) PanelWillExpand(sender objc.Object, expanding bool) {
	di._PanelWillExpand(sender, expanding)
}
func (di *OpenSavePanelDelegate) ImplementsPanelShouldEnableURL() bool {
	return di._PanelShouldEnableURL != nil
}

func (di *OpenSavePanelDelegate) SetPanelShouldEnableURL(f func(sender objc.Object, url foundation.URL) bool) {
	di._PanelShouldEnableURL = f
}

func (di *OpenSavePanelDelegate) PanelShouldEnableURL(sender objc.Object, url foundation.URL) bool {
	return di._PanelShouldEnableURL(sender, url)
}
func (di *OpenSavePanelDelegate) ImplementsPanelValidateURLError() bool {
	return di._PanelValidateURLError != nil
}

func (di *OpenSavePanelDelegate) SetPanelValidateURLError(f func(sender objc.Object, url foundation.URL, outError *foundation.Error) bool) {
	di._PanelValidateURLError = f
}

func (di *OpenSavePanelDelegate) PanelValidateURLError(sender objc.Object, url foundation.URL, outError *foundation.Error) bool {
	return di._PanelValidateURLError(sender, url, outError)
}

type OpenSavePanelDelegateWrapper struct {
	objc.Object
}

func (o_ OpenSavePanelDelegateWrapper) ImplementsPanelUserEnteredFilenameConfirmed() bool {
	return o_.RespondsToSelector(objc.GetSelector("panel:userEnteredFilename:confirmed:"))
}

func (o_ OpenSavePanelDelegateWrapper) PanelUserEnteredFilenameConfirmed(sender objc.IObject, filename string, okFlag bool) string {
	rv := objc.CallMethod[string](o_, objc.GetSelector("panel:userEnteredFilename:confirmed:"), objc.ExtractPtr(sender), filename, okFlag)
	return rv
}

func (o_ OpenSavePanelDelegateWrapper) ImplementsPanelSelectionDidChange() bool {
	return o_.RespondsToSelector(objc.GetSelector("panelSelectionDidChange:"))
}

func (o_ OpenSavePanelDelegateWrapper) PanelSelectionDidChange(sender objc.IObject) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("panelSelectionDidChange:"), objc.ExtractPtr(sender))
}

func (o_ OpenSavePanelDelegateWrapper) ImplementsPanelDidChangeToDirectoryURL() bool {
	return o_.RespondsToSelector(objc.GetSelector("panel:didChangeToDirectoryURL:"))
}

func (o_ OpenSavePanelDelegateWrapper) PanelDidChangeToDirectoryURL(sender objc.IObject, url foundation.IURL) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("panel:didChangeToDirectoryURL:"), objc.ExtractPtr(sender), objc.ExtractPtr(url))
}

func (o_ OpenSavePanelDelegateWrapper) ImplementsPanelWillExpand() bool {
	return o_.RespondsToSelector(objc.GetSelector("panel:willExpand:"))
}

func (o_ OpenSavePanelDelegateWrapper) PanelWillExpand(sender objc.IObject, expanding bool) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("panel:willExpand:"), objc.ExtractPtr(sender), expanding)
}

func (o_ OpenSavePanelDelegateWrapper) ImplementsPanelShouldEnableURL() bool {
	return o_.RespondsToSelector(objc.GetSelector("panel:shouldEnableURL:"))
}

func (o_ OpenSavePanelDelegateWrapper) PanelShouldEnableURL(sender objc.IObject, url foundation.IURL) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("panel:shouldEnableURL:"), objc.ExtractPtr(sender), objc.ExtractPtr(url))
	return rv
}

func (o_ OpenSavePanelDelegateWrapper) ImplementsPanelValidateURLError() bool {
	return o_.RespondsToSelector(objc.GetSelector("panel:validateURL:error:"))
}

func (o_ OpenSavePanelDelegateWrapper) PanelValidateURLError(sender objc.IObject, url foundation.IURL, outError *foundation.Error) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("panel:validateURL:error:"), objc.ExtractPtr(sender), objc.ExtractPtr(url), unsafe.Pointer(outError))
	return rv
}
