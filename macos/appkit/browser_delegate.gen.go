// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of methods that a browser delegate implements to manage selection, scrolling, sizing, and other behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowserdelegate?language=objc
type PBrowserDelegate interface {
	// optional
	BrowserColumnConfigurationDidChange(notification foundation.Notification)
	HasBrowserColumnConfigurationDidChange() bool

	// optional
	RootItemForBrowser(browser Browser) objc.IObject
	HasRootItemForBrowser() bool

	// optional
	BrowserDidScroll(sender Browser)
	HasBrowserDidScroll() bool

	// optional
	BrowserObjectValueForItem(browser Browser, item objc.Object) objc.IObject
	HasBrowserObjectValueForItem() bool

	// optional
	BrowserWillScroll(sender Browser)
	HasBrowserWillScroll() bool
}

// A delegate implementation builder for the [PBrowserDelegate] protocol.
type BrowserDelegate struct {
	_BrowserColumnConfigurationDidChange func(notification foundation.Notification)
	_RootItemForBrowser                  func(browser Browser) objc.IObject
	_BrowserDidScroll                    func(sender Browser)
	_BrowserObjectValueForItem           func(browser Browser, item objc.Object) objc.IObject
	_BrowserWillScroll                   func(sender Browser)
}

func (di *BrowserDelegate) HasBrowserColumnConfigurationDidChange() bool {
	return di._BrowserColumnConfigurationDidChange != nil
}

// Used by clients to implement their own column width persistence. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowserdelegate/1407542-browsercolumnconfigurationdidcha?language=objc
func (di *BrowserDelegate) SetBrowserColumnConfigurationDidChange(f func(notification foundation.Notification)) {
	di._BrowserColumnConfigurationDidChange = f
}

// Used by clients to implement their own column width persistence. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowserdelegate/1407542-browsercolumnconfigurationdidcha?language=objc
func (di *BrowserDelegate) BrowserColumnConfigurationDidChange(notification foundation.Notification) {
	di._BrowserColumnConfigurationDidChange(notification)
}
func (di *BrowserDelegate) HasRootItemForBrowser() bool {
	return di._RootItemForBrowser != nil
}

// Asks the delegate to return the root item of the browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowserdelegate/1407526-rootitemforbrowser?language=objc
func (di *BrowserDelegate) SetRootItemForBrowser(f func(browser Browser) objc.IObject) {
	di._RootItemForBrowser = f
}

// Asks the delegate to return the root item of the browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowserdelegate/1407526-rootitemforbrowser?language=objc
func (di *BrowserDelegate) RootItemForBrowser(browser Browser) objc.IObject {
	return di._RootItemForBrowser(browser)
}
func (di *BrowserDelegate) HasBrowserDidScroll() bool {
	return di._BrowserDidScroll != nil
}

// Notifies the delegate when the browser has scrolled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowserdelegate/1407566-browserdidscroll?language=objc
func (di *BrowserDelegate) SetBrowserDidScroll(f func(sender Browser)) {
	di._BrowserDidScroll = f
}

// Notifies the delegate when the browser has scrolled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowserdelegate/1407566-browserdidscroll?language=objc
func (di *BrowserDelegate) BrowserDidScroll(sender Browser) {
	di._BrowserDidScroll(sender)
}
func (di *BrowserDelegate) HasBrowserObjectValueForItem() bool {
	return di._BrowserObjectValueForItem != nil
}

// Returns the object that the specified item uses to draw its contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowserdelegate/1407594-browser?language=objc
func (di *BrowserDelegate) SetBrowserObjectValueForItem(f func(browser Browser, item objc.Object) objc.IObject) {
	di._BrowserObjectValueForItem = f
}

// Returns the object that the specified item uses to draw its contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowserdelegate/1407594-browser?language=objc
func (di *BrowserDelegate) BrowserObjectValueForItem(browser Browser, item objc.Object) objc.IObject {
	return di._BrowserObjectValueForItem(browser, item)
}
func (di *BrowserDelegate) HasBrowserWillScroll() bool {
	return di._BrowserWillScroll != nil
}

// Notifies the delegate when the browser will scroll. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowserdelegate/1407721-browserwillscroll?language=objc
func (di *BrowserDelegate) SetBrowserWillScroll(f func(sender Browser)) {
	di._BrowserWillScroll = f
}

// Notifies the delegate when the browser will scroll. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowserdelegate/1407721-browserwillscroll?language=objc
func (di *BrowserDelegate) BrowserWillScroll(sender Browser) {
	di._BrowserWillScroll(sender)
}

// A concrete type wrapper for the [PBrowserDelegate] protocol.
type BrowserDelegateWrapper struct {
	objc.Object
}

func (b_ BrowserDelegateWrapper) HasBrowserColumnConfigurationDidChange() bool {
	return b_.RespondsToSelector(objc.Sel("browserColumnConfigurationDidChange:"))
}

// Used by clients to implement their own column width persistence. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowserdelegate/1407542-browsercolumnconfigurationdidcha?language=objc
func (b_ BrowserDelegateWrapper) BrowserColumnConfigurationDidChange(notification foundation.INotification) {
	objc.Call[objc.Void](b_, objc.Sel("browserColumnConfigurationDidChange:"), objc.Ptr(notification))
}

func (b_ BrowserDelegateWrapper) HasRootItemForBrowser() bool {
	return b_.RespondsToSelector(objc.Sel("rootItemForBrowser:"))
}

// Asks the delegate to return the root item of the browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowserdelegate/1407526-rootitemforbrowser?language=objc
func (b_ BrowserDelegateWrapper) RootItemForBrowser(browser IBrowser) objc.Object {
	rv := objc.Call[objc.Object](b_, objc.Sel("rootItemForBrowser:"), objc.Ptr(browser))
	return rv
}

func (b_ BrowserDelegateWrapper) HasBrowserDidScroll() bool {
	return b_.RespondsToSelector(objc.Sel("browserDidScroll:"))
}

// Notifies the delegate when the browser has scrolled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowserdelegate/1407566-browserdidscroll?language=objc
func (b_ BrowserDelegateWrapper) BrowserDidScroll(sender IBrowser) {
	objc.Call[objc.Void](b_, objc.Sel("browserDidScroll:"), objc.Ptr(sender))
}

func (b_ BrowserDelegateWrapper) HasBrowserObjectValueForItem() bool {
	return b_.RespondsToSelector(objc.Sel("browser:objectValueForItem:"))
}

// Returns the object that the specified item uses to draw its contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowserdelegate/1407594-browser?language=objc
func (b_ BrowserDelegateWrapper) BrowserObjectValueForItem(browser IBrowser, item objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](b_, objc.Sel("browser:objectValueForItem:"), objc.Ptr(browser), item)
	return rv
}

func (b_ BrowserDelegateWrapper) HasBrowserWillScroll() bool {
	return b_.RespondsToSelector(objc.Sel("browserWillScroll:"))
}

// Notifies the delegate when the browser will scroll. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowserdelegate/1407721-browserwillscroll?language=objc
func (b_ BrowserDelegateWrapper) BrowserWillScroll(sender IBrowser) {
	objc.Call[objc.Void](b_, objc.Sel("browserWillScroll:"), objc.Ptr(sender))
}
