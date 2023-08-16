// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [HTTPCookieStorage] class.
var HTTPCookieStorageClass = _HTTPCookieStorageClass{objc.GetClass("NSHTTPCookieStorage")}

type _HTTPCookieStorageClass struct {
	objc.Class
}

// An interface definition for the [HTTPCookieStorage] class.
type IHTTPCookieStorage interface {
	objc.IObject
	DeleteCookie(cookie IHTTPCookie)
	GetCookiesForTaskCompletionHandler(task IURLSessionTask, completionHandler func(cookies []HTTPCookie))
	StoreCookiesForTask(cookies []IHTTPCookie, task IURLSessionTask)
	SortedCookiesUsingDescriptors(sortOrder []ISortDescriptor) []HTTPCookie
	CookiesForURL(URL IURL) []HTTPCookie
	RemoveCookiesSinceDate(date IDate)
	SetCookie(cookie IHTTPCookie)
	SetCookiesForURLMainDocumentURL(cookies []IHTTPCookie, URL IURL, mainDocumentURL IURL)
	Cookies() []HTTPCookie
	CookieAcceptPolicy() HTTPCookieAcceptPolicy
	SetCookieAcceptPolicy(value HTTPCookieAcceptPolicy)
}

// A container that manages the storage of cookies. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookiestorage?language=objc
type HTTPCookieStorage struct {
	objc.Object
}

func HTTPCookieStorageFrom(ptr unsafe.Pointer) HTTPCookieStorage {
	return HTTPCookieStorage{
		Object: objc.ObjectFrom(ptr),
	}
}

func (hc _HTTPCookieStorageClass) Alloc() HTTPCookieStorage {
	rv := objc.Call[HTTPCookieStorage](hc, objc.Sel("alloc"))
	return rv
}

func HTTPCookieStorage_Alloc() HTTPCookieStorage {
	return HTTPCookieStorageClass.Alloc()
}

func (hc _HTTPCookieStorageClass) New() HTTPCookieStorage {
	rv := objc.Call[HTTPCookieStorage](hc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewHTTPCookieStorage() HTTPCookieStorage {
	return HTTPCookieStorageClass.New()
}

func (h_ HTTPCookieStorage) Init() HTTPCookieStorage {
	rv := objc.Call[HTTPCookieStorage](h_, objc.Sel("init"))
	return rv
}

// Deletes the specified cookie from the cookie storage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookiestorage/1409218-deletecookie?language=objc
func (h_ HTTPCookieStorage) DeleteCookie(cookie IHTTPCookie) {
	objc.Call[objc.Void](h_, objc.Sel("deleteCookie:"), objc.Ptr(cookie))
}

// Fetches cookies relevant to the specified task and passes them to the completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookiestorage/1408517-getcookiesfortask?language=objc
func (h_ HTTPCookieStorage) GetCookiesForTaskCompletionHandler(task IURLSessionTask, completionHandler func(cookies []HTTPCookie)) {
	objc.Call[objc.Void](h_, objc.Sel("getCookiesForTask:completionHandler:"), objc.Ptr(task), completionHandler)
}

// Stores an array of cookies in the cookie storage, on behalf of the provided task, if the cookie accept policy permits. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookiestorage/1415381-storecookies?language=objc
func (h_ HTTPCookieStorage) StoreCookiesForTask(cookies []IHTTPCookie, task IURLSessionTask) {
	objc.Call[objc.Void](h_, objc.Sel("storeCookies:forTask:"), cookies, objc.Ptr(task))
}

// Returns the cookie storage instance for the container associated with the specified app group identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookiestorage/1411361-sharedcookiestorageforgroupconta?language=objc
func (hc _HTTPCookieStorageClass) SharedCookieStorageForGroupContainerIdentifier(identifier string) HTTPCookieStorage {
	rv := objc.Call[HTTPCookieStorage](hc, objc.Sel("sharedCookieStorageForGroupContainerIdentifier:"), identifier)
	return rv
}

// Returns the cookie storage instance for the container associated with the specified app group identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookiestorage/1411361-sharedcookiestorageforgroupconta?language=objc
func HTTPCookieStorage_SharedCookieStorageForGroupContainerIdentifier(identifier string) HTTPCookieStorage {
	return HTTPCookieStorageClass.SharedCookieStorageForGroupContainerIdentifier(identifier)
}

// Returns all of the cookie storage’s cookies, sorted according to a given set of sort descriptors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookiestorage/1413730-sortedcookiesusingdescriptors?language=objc
func (h_ HTTPCookieStorage) SortedCookiesUsingDescriptors(sortOrder []ISortDescriptor) []HTTPCookie {
	rv := objc.Call[[]HTTPCookie](h_, objc.Sel("sortedCookiesUsingDescriptors:"), sortOrder)
	return rv
}

// Returns all the cookie storage’s cookies that are sent to a specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookiestorage/1412100-cookiesforurl?language=objc
func (h_ HTTPCookieStorage) CookiesForURL(URL IURL) []HTTPCookie {
	rv := objc.Call[[]HTTPCookie](h_, objc.Sel("cookiesForURL:"), objc.Ptr(URL))
	return rv
}

// Removes cookies that were stored after a given date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookiestorage/1407256-removecookiessincedate?language=objc
func (h_ HTTPCookieStorage) RemoveCookiesSinceDate(date IDate) {
	objc.Call[objc.Void](h_, objc.Sel("removeCookiesSinceDate:"), objc.Ptr(date))
}

// Stores a specified cookie in the cookie storage if the cookie accept policy permits. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookiestorage/1418356-setcookie?language=objc
func (h_ HTTPCookieStorage) SetCookie(cookie IHTTPCookie) {
	objc.Call[objc.Void](h_, objc.Sel("setCookie:"), objc.Ptr(cookie))
}

// Adds an array of cookies to the cookie storage if the storage’s cookie acceptance policy permits. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookiestorage/1412510-setcookies?language=objc
func (h_ HTTPCookieStorage) SetCookiesForURLMainDocumentURL(cookies []IHTTPCookie, URL IURL, mainDocumentURL IURL) {
	objc.Call[objc.Void](h_, objc.Sel("setCookies:forURL:mainDocumentURL:"), cookies, objc.Ptr(URL), objc.Ptr(mainDocumentURL))
}

// The cookie storage’s cookies. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookiestorage/1418390-cookies?language=objc
func (h_ HTTPCookieStorage) Cookies() []HTTPCookie {
	rv := objc.Call[[]HTTPCookie](h_, objc.Sel("cookies"))
	return rv
}

// The cookie storage’s cookie accept policy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookiestorage/1410415-cookieacceptpolicy?language=objc
func (h_ HTTPCookieStorage) CookieAcceptPolicy() HTTPCookieAcceptPolicy {
	rv := objc.Call[HTTPCookieAcceptPolicy](h_, objc.Sel("cookieAcceptPolicy"))
	return rv
}

// The cookie storage’s cookie accept policy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookiestorage/1410415-cookieacceptpolicy?language=objc
func (h_ HTTPCookieStorage) SetCookieAcceptPolicy(value HTTPCookieAcceptPolicy) {
	objc.Call[objc.Void](h_, objc.Sel("setCookieAcceptPolicy:"), value)
}

// The shared cookie storage instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookiestorage/1416095-sharedhttpcookiestorage?language=objc
func (hc _HTTPCookieStorageClass) SharedHTTPCookieStorage() HTTPCookieStorage {
	rv := objc.Call[HTTPCookieStorage](hc, objc.Sel("sharedHTTPCookieStorage"))
	return rv
}

// The shared cookie storage instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookiestorage/1416095-sharedhttpcookiestorage?language=objc
func HTTPCookieStorage_SharedHTTPCookieStorage() HTTPCookieStorage {
	return HTTPCookieStorageClass.SharedHTTPCookieStorage()
}
