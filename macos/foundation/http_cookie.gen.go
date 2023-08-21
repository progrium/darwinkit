// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [HTTPCookie] class.
var HTTPCookieClass = _HTTPCookieClass{objc.GetClass("NSHTTPCookie")}

type _HTTPCookieClass struct {
	objc.Class
}

// An interface definition for the [HTTPCookie] class.
type IHTTPCookie interface {
	objc.IObject
	Domain() string
	Path() string
	Version() uint
	CommentURL() URL
	Value() string
	Name() string
	IsHTTPOnly() bool
	Comment() string
	Properties() map[HTTPCookiePropertyKey]objc.Object
	IsSecure() bool
	IsSessionOnly() bool
	ExpiresDate() Date
	SameSitePolicy() HTTPCookieStringPolicy
	PortList() []Number
}

// A representation of an HTTP cookie. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookie?language=objc
type HTTPCookie struct {
	objc.Object
}

func HTTPCookieFrom(ptr unsafe.Pointer) HTTPCookie {
	return HTTPCookie{
		Object: objc.ObjectFrom(ptr),
	}
}

func (h_ HTTPCookie) InitWithProperties(properties map[HTTPCookiePropertyKey]objc.IObject) HTTPCookie {
	rv := objc.Call[HTTPCookie](h_, objc.Sel("initWithProperties:"), properties)
	return rv
}

// Initializes an HTTP cookie object with the given cookie properties. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookie/1392975-initwithproperties?language=objc
func NewHTTPCookieWithProperties(properties map[HTTPCookiePropertyKey]objc.IObject) HTTPCookie {
	instance := HTTPCookieClass.Alloc().InitWithProperties(properties)
	instance.Autorelease()
	return instance
}

func (hc _HTTPCookieClass) Alloc() HTTPCookie {
	rv := objc.Call[HTTPCookie](hc, objc.Sel("alloc"))
	return rv
}

func HTTPCookie_Alloc() HTTPCookie {
	return HTTPCookieClass.Alloc()
}

func (hc _HTTPCookieClass) New() HTTPCookie {
	rv := objc.Call[HTTPCookie](hc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewHTTPCookie() HTTPCookie {
	return HTTPCookieClass.New()
}

func (h_ HTTPCookie) Init() HTTPCookie {
	rv := objc.Call[HTTPCookie](h_, objc.Sel("init"))
	return rv
}

// Creates an array of HTTP cookies that corresponds to the provided response header fields for the provided URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookie/1393011-cookieswithresponseheaderfields?language=objc
func (hc _HTTPCookieClass) CookiesWithResponseHeaderFieldsForURL(headerFields map[string]string, URL IURL) []HTTPCookie {
	rv := objc.Call[[]HTTPCookie](hc, objc.Sel("cookiesWithResponseHeaderFields:forURL:"), headerFields, objc.Ptr(URL))
	return rv
}

// Creates an array of HTTP cookies that corresponds to the provided response header fields for the provided URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookie/1393011-cookieswithresponseheaderfields?language=objc
func HTTPCookie_CookiesWithResponseHeaderFieldsForURL(headerFields map[string]string, URL IURL) []HTTPCookie {
	return HTTPCookieClass.CookiesWithResponseHeaderFieldsForURL(headerFields, URL)
}

// Converts an array of cookies to a dictionary of header fields. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookie/1393021-requestheaderfieldswithcookies?language=objc
func (hc _HTTPCookieClass) RequestHeaderFieldsWithCookies(cookies []IHTTPCookie) map[string]string {
	rv := objc.Call[map[string]string](hc, objc.Sel("requestHeaderFieldsWithCookies:"), cookies)
	return rv
}

// Converts an array of cookies to a dictionary of header fields. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookie/1393021-requestheaderfieldswithcookies?language=objc
func HTTPCookie_RequestHeaderFieldsWithCookies(cookies []IHTTPCookie) map[string]string {
	return HTTPCookieClass.RequestHeaderFieldsWithCookies(cookies)
}

// Creates and initializes an HTTP cookie object using the provided properties. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookie/1392967-cookiewithproperties?language=objc
func (hc _HTTPCookieClass) CookieWithProperties(properties map[HTTPCookiePropertyKey]objc.IObject) HTTPCookie {
	rv := objc.Call[HTTPCookie](hc, objc.Sel("cookieWithProperties:"), properties)
	return rv
}

// Creates and initializes an HTTP cookie object using the provided properties. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookie/1392967-cookiewithproperties?language=objc
func HTTPCookie_CookieWithProperties(properties map[HTTPCookiePropertyKey]objc.IObject) HTTPCookie {
	return HTTPCookieClass.CookieWithProperties(properties)
}

// The domain of the cookie. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookie/1393015-domain?language=objc
func (h_ HTTPCookie) Domain() string {
	rv := objc.Call[string](h_, objc.Sel("domain"))
	return rv
}

// The cookie’s path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookie/1392981-path?language=objc
func (h_ HTTPCookie) Path() string {
	rv := objc.Call[string](h_, objc.Sel("path"))
	return rv
}

// The cookie’s version. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookie/1392993-version?language=objc
func (h_ HTTPCookie) Version() uint {
	rv := objc.Call[uint](h_, objc.Sel("version"))
	return rv
}

// The cookie’s comment URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookie/1392987-commenturl?language=objc
func (h_ HTTPCookie) CommentURL() URL {
	rv := objc.Call[URL](h_, objc.Sel("commentURL"))
	return rv
}

// The cookie‘s string value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookie/1392995-value?language=objc
func (h_ HTTPCookie) Value() string {
	rv := objc.Call[string](h_, objc.Sel("value"))
	return rv
}

// The cookie’s name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookie/1393013-name?language=objc
func (h_ HTTPCookie) Name() string {
	rv := objc.Call[string](h_, objc.Sel("name"))
	return rv
}

// A Boolean value that indicates whether the cookie should only be sent to HTTP servers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookie/1392969-httponly?language=objc
func (h_ HTTPCookie) IsHTTPOnly() bool {
	rv := objc.Call[bool](h_, objc.Sel("isHTTPOnly"))
	return rv
}

// The cookie’s comment string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookie/1392997-comment?language=objc
func (h_ HTTPCookie) Comment() string {
	rv := objc.Call[string](h_, objc.Sel("comment"))
	return rv
}

// The cookie’s properties. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookie/1393017-properties?language=objc
func (h_ HTTPCookie) Properties() map[HTTPCookiePropertyKey]objc.Object {
	rv := objc.Call[map[HTTPCookiePropertyKey]objc.Object](h_, objc.Sel("properties"))
	return rv
}

// A Boolean value that indicates whether the cookie may only be sent over secure channels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookie/1393025-secure?language=objc
func (h_ HTTPCookie) IsSecure() bool {
	rv := objc.Call[bool](h_, objc.Sel("isSecure"))
	return rv
}

// A Boolean value that indicates whether the cookie should be discarded at the end of the session (regardless of expiration date). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookie/1392991-sessiononly?language=objc
func (h_ HTTPCookie) IsSessionOnly() bool {
	rv := objc.Call[bool](h_, objc.Sel("isSessionOnly"))
	return rv
}

// The cookie’s expiration date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookie/1393019-expiresdate?language=objc
func (h_ HTTPCookie) ExpiresDate() Date {
	rv := objc.Call[Date](h_, objc.Sel("expiresDate"))
	return rv
}

// A Boolean value that indicates whether to restrict the cookie to requests sent back to the same site that created it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookie/3042420-samesitepolicy?language=objc
func (h_ HTTPCookie) SameSitePolicy() HTTPCookieStringPolicy {
	rv := objc.Call[HTTPCookieStringPolicy](h_, objc.Sel("sameSitePolicy"))
	return rv
}

// The cookie’s port list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpcookie/1393027-portlist?language=objc
func (h_ HTTPCookie) PortList() []Number {
	rv := objc.Call[[]Number](h_, objc.Sel("portList"))
	return rv
}
