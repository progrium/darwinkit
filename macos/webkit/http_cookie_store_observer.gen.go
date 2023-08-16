// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"github.com/progrium/macdriver/objc"
)

// The methods to adopt in an object that monitors changes to a webpage’s cookies. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkhttpcookiestoreobserver?language=objc
type PHTTPCookieStoreObserver interface {
	// optional
	CookiesDidChangeInCookieStore(cookieStore HTTPCookieStore)
	HasCookiesDidChangeInCookieStore() bool
}

// A concrete type wrapper for the [PHTTPCookieStoreObserver] protocol.
type HTTPCookieStoreObserverWrapper struct {
	objc.Object
}

func (h_ HTTPCookieStoreObserverWrapper) HasCookiesDidChangeInCookieStore() bool {
	return h_.RespondsToSelector(objc.Sel("cookiesDidChangeInCookieStore:"))
}

// Tells the delegate that the cookies in the specified cookie store changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkhttpcookiestoreobserver/2882008-cookiesdidchangeincookiestore?language=objc
func (h_ HTTPCookieStoreObserverWrapper) CookiesDidChangeInCookieStore(cookieStore IHTTPCookieStore) {
	objc.Call[objc.Void](h_, objc.Sel("cookiesDidChangeInCookieStore:"), objc.Ptr(cookieStore))
}
