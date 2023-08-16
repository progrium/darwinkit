// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of methods for managing interactions with an open or save panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopensavepaneldelegate?language=objc
type POpenSavePanelDelegate interface {
	// optional
	PanelShouldEnableURL(sender objc.Object, url foundation.URL) bool
	HasPanelShouldEnableURL() bool

	// optional
	PanelSelectionDidChange(sender objc.Object)
	HasPanelSelectionDidChange() bool
}

// A delegate implementation builder for the [POpenSavePanelDelegate] protocol.
type OpenSavePanelDelegate struct {
	_PanelShouldEnableURL    func(sender objc.Object, url foundation.URL) bool
	_PanelSelectionDidChange func(sender objc.Object)
}

func (di *OpenSavePanelDelegate) HasPanelShouldEnableURL() bool {
	return di._PanelShouldEnableURL != nil
}

// Asks the delegate whether the specified URL should be enabled in the Open panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopensavepaneldelegate/1535200-panel?language=objc
func (di *OpenSavePanelDelegate) SetPanelShouldEnableURL(f func(sender objc.Object, url foundation.URL) bool) {
	di._PanelShouldEnableURL = f
}

// Asks the delegate whether the specified URL should be enabled in the Open panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopensavepaneldelegate/1535200-panel?language=objc
func (di *OpenSavePanelDelegate) PanelShouldEnableURL(sender objc.Object, url foundation.URL) bool {
	return di._PanelShouldEnableURL(sender, url)
}
func (di *OpenSavePanelDelegate) HasPanelSelectionDidChange() bool {
	return di._PanelSelectionDidChange != nil
}

// Tells the delegate that the user changed the selection in the specified Save panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopensavepaneldelegate/1533556-panelselectiondidchange?language=objc
func (di *OpenSavePanelDelegate) SetPanelSelectionDidChange(f func(sender objc.Object)) {
	di._PanelSelectionDidChange = f
}

// Tells the delegate that the user changed the selection in the specified Save panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopensavepaneldelegate/1533556-panelselectiondidchange?language=objc
func (di *OpenSavePanelDelegate) PanelSelectionDidChange(sender objc.Object) {
	di._PanelSelectionDidChange(sender)
}

// A concrete type wrapper for the [POpenSavePanelDelegate] protocol.
type OpenSavePanelDelegateWrapper struct {
	objc.Object
}

func (o_ OpenSavePanelDelegateWrapper) HasPanelShouldEnableURL() bool {
	return o_.RespondsToSelector(objc.Sel("panel:shouldEnableURL:"))
}

// Asks the delegate whether the specified URL should be enabled in the Open panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopensavepaneldelegate/1535200-panel?language=objc
func (o_ OpenSavePanelDelegateWrapper) PanelShouldEnableURL(sender objc.IObject, url foundation.IURL) bool {
	rv := objc.Call[bool](o_, objc.Sel("panel:shouldEnableURL:"), sender, objc.Ptr(url))
	return rv
}

func (o_ OpenSavePanelDelegateWrapper) HasPanelSelectionDidChange() bool {
	return o_.RespondsToSelector(objc.Sel("panelSelectionDidChange:"))
}

// Tells the delegate that the user changed the selection in the specified Save panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopensavepaneldelegate/1533556-panelselectiondidchange?language=objc
func (o_ OpenSavePanelDelegateWrapper) PanelSelectionDidChange(sender objc.IObject) {
	objc.Call[objc.Void](o_, objc.Sel("panelSelectionDidChange:"), sender)
}
