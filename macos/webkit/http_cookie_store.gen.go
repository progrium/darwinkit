// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [HTTPCookieStore] class.
var HTTPCookieStoreClass = _HTTPCookieStoreClass{objc.GetClass("WKHTTPCookieStore")}

type _HTTPCookieStoreClass struct {
	objc.Class
}

// An interface definition for the [HTTPCookieStore] class.
type IHTTPCookieStore interface {
	objc.IObject
	DeleteCookieCompletionHandler(cookie foundation.IHTTPCookie, completionHandler func())
	GetAllCookies(completionHandler func(arg0 []foundation.HTTPCookie))
	RemoveObserver(observer PHTTPCookieStoreObserver)
	RemoveObserverObject(observerObject objc.IObject)
	SetCookieCompletionHandler(cookie foundation.IHTTPCookie, completionHandler func())
	AddObserver(observer PHTTPCookieStoreObserver)
	AddObserverObject(observerObject objc.IObject)
}

// An object that manages the HTTP cookies associated with a particular web view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkhttpcookiestore?language=objc
type HTTPCookieStore struct {
	objc.Object
}

func HTTPCookieStoreFrom(ptr unsafe.Pointer) HTTPCookieStore {
	return HTTPCookieStore{
		Object: objc.ObjectFrom(ptr),
	}
}

func (hc _HTTPCookieStoreClass) Alloc() HTTPCookieStore {
	rv := objc.Call[HTTPCookieStore](hc, objc.Sel("alloc"))
	return rv
}

func HTTPCookieStore_Alloc() HTTPCookieStore {
	return HTTPCookieStoreClass.Alloc()
}

func (hc _HTTPCookieStoreClass) New() HTTPCookieStore {
	rv := objc.Call[HTTPCookieStore](hc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewHTTPCookieStore() HTTPCookieStore {
	return HTTPCookieStoreClass.New()
}

func (h_ HTTPCookieStore) Init() HTTPCookieStore {
	rv := objc.Call[HTTPCookieStore](h_, objc.Sel("init"))
	return rv
}

// Deletes the specified cookie. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkhttpcookiestore/2882009-deletecookie?language=objc
func (h_ HTTPCookieStore) DeleteCookieCompletionHandler(cookie foundation.IHTTPCookie, completionHandler func()) {
	objc.Call[objc.Void](h_, objc.Sel("deleteCookie:completionHandler:"), objc.Ptr(cookie), completionHandler)
}

// Fetches all stored cookies asynchronously and delivers them to the specified completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkhttpcookiestore/2882005-getallcookies?language=objc
func (h_ HTTPCookieStore) GetAllCookies(completionHandler func(arg0 []foundation.HTTPCookie)) {
	objc.Call[objc.Void](h_, objc.Sel("getAllCookies:"), completionHandler)
}

// Removes an observer from the cookie store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkhttpcookiestore/2882004-removeobserver?language=objc
func (h_ HTTPCookieStore) RemoveObserver(observer PHTTPCookieStoreObserver) {
	po0 := objc.WrapAsProtocol("WKHTTPCookieStoreObserver", observer)
	objc.Call[objc.Void](h_, objc.Sel("removeObserver:"), po0)
}

// Removes an observer from the cookie store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkhttpcookiestore/2882004-removeobserver?language=objc
func (h_ HTTPCookieStore) RemoveObserverObject(observerObject objc.IObject) {
	objc.Call[objc.Void](h_, objc.Sel("removeObserver:"), objc.Ptr(observerObject))
}

// Adds a cookie to the cookie store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkhttpcookiestore/2882007-setcookie?language=objc
func (h_ HTTPCookieStore) SetCookieCompletionHandler(cookie foundation.IHTTPCookie, completionHandler func()) {
	objc.Call[objc.Void](h_, objc.Sel("setCookie:completionHandler:"), objc.Ptr(cookie), completionHandler)
}

// Adds an observer to the cookie store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkhttpcookiestore/2882010-addobserver?language=objc
func (h_ HTTPCookieStore) AddObserver(observer PHTTPCookieStoreObserver) {
	po0 := objc.WrapAsProtocol("WKHTTPCookieStoreObserver", observer)
	objc.Call[objc.Void](h_, objc.Sel("addObserver:"), po0)
}

// Adds an observer to the cookie store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkhttpcookiestore/2882010-addobserver?language=objc
func (h_ HTTPCookieStore) AddObserverObject(observerObject objc.IObject) {
	objc.Call[objc.Void](h_, objc.Sel("addObserver:"), objc.Ptr(observerObject))
}
