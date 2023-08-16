// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// The interface a net service browser uses to inform a delegate about the state of service discovery. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicebrowserdelegate?language=objc
type PNetServiceBrowserDelegate interface {
	// optional
	NetServiceBrowserWillSearch(browser NetServiceBrowser)
	HasNetServiceBrowserWillSearch() bool

	// optional
	NetServiceBrowserDidNotSearch(browser NetServiceBrowser, errorDict map[string]Number)
	HasNetServiceBrowserDidNotSearch() bool

	// optional
	NetServiceBrowserDidStopSearch(browser NetServiceBrowser)
	HasNetServiceBrowserDidStopSearch() bool
}

// A delegate implementation builder for the [PNetServiceBrowserDelegate] protocol.
type NetServiceBrowserDelegate struct {
	_NetServiceBrowserWillSearch    func(browser NetServiceBrowser)
	_NetServiceBrowserDidNotSearch  func(browser NetServiceBrowser, errorDict map[string]Number)
	_NetServiceBrowserDidStopSearch func(browser NetServiceBrowser)
}

func (di *NetServiceBrowserDelegate) HasNetServiceBrowserWillSearch() bool {
	return di._NetServiceBrowserWillSearch != nil
}

// Tells the delegate that a search is commencing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicebrowserdelegate/1408173-netservicebrowserwillsearch?language=objc
func (di *NetServiceBrowserDelegate) SetNetServiceBrowserWillSearch(f func(browser NetServiceBrowser)) {
	di._NetServiceBrowserWillSearch = f
}

// Tells the delegate that a search is commencing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicebrowserdelegate/1408173-netservicebrowserwillsearch?language=objc
func (di *NetServiceBrowserDelegate) NetServiceBrowserWillSearch(browser NetServiceBrowser) {
	di._NetServiceBrowserWillSearch(browser)
}
func (di *NetServiceBrowserDelegate) HasNetServiceBrowserDidNotSearch() bool {
	return di._NetServiceBrowserDidNotSearch != nil
}

// Tells the delegate that a search was not successful. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicebrowserdelegate/1410567-netservicebrowser?language=objc
func (di *NetServiceBrowserDelegate) SetNetServiceBrowserDidNotSearch(f func(browser NetServiceBrowser, errorDict map[string]Number)) {
	di._NetServiceBrowserDidNotSearch = f
}

// Tells the delegate that a search was not successful. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicebrowserdelegate/1410567-netservicebrowser?language=objc
func (di *NetServiceBrowserDelegate) NetServiceBrowserDidNotSearch(browser NetServiceBrowser, errorDict map[string]Number) {
	di._NetServiceBrowserDidNotSearch(browser, errorDict)
}
func (di *NetServiceBrowserDelegate) HasNetServiceBrowserDidStopSearch() bool {
	return di._NetServiceBrowserDidStopSearch != nil
}

// Tells the delegate that a search was stopped. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicebrowserdelegate/1418341-netservicebrowserdidstopsearch?language=objc
func (di *NetServiceBrowserDelegate) SetNetServiceBrowserDidStopSearch(f func(browser NetServiceBrowser)) {
	di._NetServiceBrowserDidStopSearch = f
}

// Tells the delegate that a search was stopped. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicebrowserdelegate/1418341-netservicebrowserdidstopsearch?language=objc
func (di *NetServiceBrowserDelegate) NetServiceBrowserDidStopSearch(browser NetServiceBrowser) {
	di._NetServiceBrowserDidStopSearch(browser)
}

// A concrete type wrapper for the [PNetServiceBrowserDelegate] protocol.
type NetServiceBrowserDelegateWrapper struct {
	objc.Object
}

func (n_ NetServiceBrowserDelegateWrapper) HasNetServiceBrowserWillSearch() bool {
	return n_.RespondsToSelector(objc.Sel("netServiceBrowserWillSearch:"))
}

// Tells the delegate that a search is commencing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicebrowserdelegate/1408173-netservicebrowserwillsearch?language=objc
func (n_ NetServiceBrowserDelegateWrapper) NetServiceBrowserWillSearch(browser INetServiceBrowser) {
	objc.Call[objc.Void](n_, objc.Sel("netServiceBrowserWillSearch:"), objc.Ptr(browser))
}

func (n_ NetServiceBrowserDelegateWrapper) HasNetServiceBrowserDidNotSearch() bool {
	return n_.RespondsToSelector(objc.Sel("netServiceBrowser:didNotSearch:"))
}

// Tells the delegate that a search was not successful. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicebrowserdelegate/1410567-netservicebrowser?language=objc
func (n_ NetServiceBrowserDelegateWrapper) NetServiceBrowserDidNotSearch(browser INetServiceBrowser, errorDict map[string]INumber) {
	objc.Call[objc.Void](n_, objc.Sel("netServiceBrowser:didNotSearch:"), objc.Ptr(browser), errorDict)
}

func (n_ NetServiceBrowserDelegateWrapper) HasNetServiceBrowserDidStopSearch() bool {
	return n_.RespondsToSelector(objc.Sel("netServiceBrowserDidStopSearch:"))
}

// Tells the delegate that a search was stopped. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicebrowserdelegate/1418341-netservicebrowserdidstopsearch?language=objc
func (n_ NetServiceBrowserDelegateWrapper) NetServiceBrowserDidStopSearch(browser INetServiceBrowser) {
	objc.Call[objc.Void](n_, objc.Sel("netServiceBrowserDidStopSearch:"), objc.Ptr(browser))
}
