// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var HTTPCookieStoreClass = _HTTPCookieStoreClass{objc.GetClass("WKHTTPCookieStore")}

type _HTTPCookieStoreClass struct {
	objc.Class
}

type IHTTPCookieStore interface {
	objc.IObject
	GetAllCookies(completionHandler func(param1 []foundation.HTTPCookie))
	SetCookieCompletionHandler(cookie foundation.IHTTPCookie, completionHandler func())
	DeleteCookieCompletionHandler(cookie foundation.IHTTPCookie, completionHandler func())
}

type HTTPCookieStore struct {
	objc.Object
}

func MakeHTTPCookieStore(ptr unsafe.Pointer) HTTPCookieStore {
	return HTTPCookieStore{
		Object: objc.MakeObject(ptr),
	}
}

func (hc _HTTPCookieStoreClass) Alloc() HTTPCookieStore {
	rv := objc.CallMethod[HTTPCookieStore](hc, objc.GetSelector("alloc"))
	return rv
}

func HTTPCookieStore_Alloc() HTTPCookieStore {
	return HTTPCookieStoreClass.Alloc()
}

func (hc _HTTPCookieStoreClass) New() HTTPCookieStore {
	rv := objc.CallMethod[HTTPCookieStore](hc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewHTTPCookieStore() HTTPCookieStore {
	return HTTPCookieStoreClass.New()
}

func HTTPCookieStore_New() HTTPCookieStore {
	return HTTPCookieStoreClass.New()
}

func (h_ HTTPCookieStore) Init() HTTPCookieStore {
	rv := objc.CallMethod[HTTPCookieStore](h_, objc.GetSelector("init"))
	return rv
}

func HTTPCookieStore_Init() HTTPCookieStore {
	return HTTPCookieStoreClass.Alloc().Init()
}

func (h_ HTTPCookieStore) GetAllCookies(completionHandler func(param1 []foundation.HTTPCookie)) {
	objc.CallMethod[objc.Void](h_, objc.GetSelector("getAllCookies:"), completionHandler)
}

func (h_ HTTPCookieStore) SetCookieCompletionHandler(cookie foundation.IHTTPCookie, completionHandler func()) {
	objc.CallMethod[objc.Void](h_, objc.GetSelector("setCookie:completionHandler:"), objc.ExtractPtr(cookie), completionHandler)
}

func (h_ HTTPCookieStore) DeleteCookieCompletionHandler(cookie foundation.IHTTPCookie, completionHandler func()) {
	objc.CallMethod[objc.Void](h_, objc.GetSelector("deleteCookie:completionHandler:"), objc.ExtractPtr(cookie), completionHandler)
}
