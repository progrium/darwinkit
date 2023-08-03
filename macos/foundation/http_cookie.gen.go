// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var HTTPCookieClass = _HTTPCookieClass{objc.GetClass("NSHTTPCookie")}

type _HTTPCookieClass struct {
	objc.Class
}

type IHTTPCookie interface {
	objc.IObject
	Domain() string
	Path() string
	PortList() []Number
	Name() string
	Value() string
	Version() uint
	ExpiresDate() Date
	IsSessionOnly() bool
	IsHTTPOnly() bool
	IsSecure() bool
	SameSitePolicy() HTTPCookieStringPolicy
	Properties() map[HTTPCookiePropertyKey]objc.Object
	Comment() string
	CommentURL() URL
}

type HTTPCookie struct {
	objc.Object
}

func MakeHTTPCookie(ptr unsafe.Pointer) HTTPCookie {
	return HTTPCookie{
		Object: objc.MakeObject(ptr),
	}
}

func (h_ HTTPCookie) InitWithProperties(properties map[HTTPCookiePropertyKey]objc.IObject) HTTPCookie {
	rv := objc.CallMethod[HTTPCookie](h_, objc.GetSelector("initWithProperties:"), properties)
	return rv
}

func HTTPCookie_InitWithProperties(properties map[HTTPCookiePropertyKey]objc.IObject) HTTPCookie {
	return HTTPCookieClass.Alloc().InitWithProperties(properties)
}

func (hc _HTTPCookieClass) Alloc() HTTPCookie {
	rv := objc.CallMethod[HTTPCookie](hc, objc.GetSelector("alloc"))
	return rv
}

func HTTPCookie_Alloc() HTTPCookie {
	return HTTPCookieClass.Alloc()
}

func (hc _HTTPCookieClass) New() HTTPCookie {
	rv := objc.CallMethod[HTTPCookie](hc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewHTTPCookie() HTTPCookie {
	return HTTPCookieClass.New()
}

func HTTPCookie_New() HTTPCookie {
	return HTTPCookieClass.New()
}

func (h_ HTTPCookie) Init() HTTPCookie {
	rv := objc.CallMethod[HTTPCookie](h_, objc.GetSelector("init"))
	return rv
}

func HTTPCookie_Init() HTTPCookie {
	return HTTPCookieClass.Alloc().Init()
}

func (hc _HTTPCookieClass) CookiesWithResponseHeaderFieldsForURL(headerFields map[string]string, URL IURL) []HTTPCookie {
	rv := objc.CallMethod[[]HTTPCookie](hc, objc.GetSelector("cookiesWithResponseHeaderFields:forURL:"), headerFields, objc.ExtractPtr(URL))
	// TODO: convert slice items...
	return rv
}

func HTTPCookie_CookiesWithResponseHeaderFieldsForURL(headerFields map[string]string, URL IURL) []HTTPCookie {
	return HTTPCookieClass.CookiesWithResponseHeaderFieldsForURL(headerFields, URL)
}

func (hc _HTTPCookieClass) CookieWithProperties(properties map[HTTPCookiePropertyKey]objc.IObject) HTTPCookie {
	rv := objc.CallMethod[HTTPCookie](hc, objc.GetSelector("cookieWithProperties:"), properties)
	return rv
}

func HTTPCookie_CookieWithProperties(properties map[HTTPCookiePropertyKey]objc.IObject) HTTPCookie {
	return HTTPCookieClass.CookieWithProperties(properties)
}

func (hc _HTTPCookieClass) RequestHeaderFieldsWithCookies(cookies []IHTTPCookie) map[string]string {
	rv := objc.CallMethod[map[string]string](hc, objc.GetSelector("requestHeaderFieldsWithCookies:"), cookies)
	return rv
}

func HTTPCookie_RequestHeaderFieldsWithCookies(cookies []IHTTPCookie) map[string]string {
	return HTTPCookieClass.RequestHeaderFieldsWithCookies(cookies)
}

func (h_ HTTPCookie) Domain() string {
	rv := objc.CallMethod[string](h_, objc.GetSelector("domain"))
	return rv
}

func (h_ HTTPCookie) Path() string {
	rv := objc.CallMethod[string](h_, objc.GetSelector("path"))
	return rv
}

func (h_ HTTPCookie) PortList() []Number {
	rv := objc.CallMethod[[]Number](h_, objc.GetSelector("portList"))
	// TODO: convert slice items...
	return rv
}

func (h_ HTTPCookie) Name() string {
	rv := objc.CallMethod[string](h_, objc.GetSelector("name"))
	return rv
}

func (h_ HTTPCookie) Value() string {
	rv := objc.CallMethod[string](h_, objc.GetSelector("value"))
	return rv
}

func (h_ HTTPCookie) Version() uint {
	rv := objc.CallMethod[uint](h_, objc.GetSelector("version"))
	return rv
}

func (h_ HTTPCookie) ExpiresDate() Date {
	rv := objc.CallMethod[Date](h_, objc.GetSelector("expiresDate"))
	return rv
}

func (h_ HTTPCookie) IsSessionOnly() bool {
	rv := objc.CallMethod[bool](h_, objc.GetSelector("isSessionOnly"))
	return rv
}

func (h_ HTTPCookie) IsHTTPOnly() bool {
	rv := objc.CallMethod[bool](h_, objc.GetSelector("isHTTPOnly"))
	return rv
}

func (h_ HTTPCookie) IsSecure() bool {
	rv := objc.CallMethod[bool](h_, objc.GetSelector("isSecure"))
	return rv
}

func (h_ HTTPCookie) SameSitePolicy() HTTPCookieStringPolicy {
	rv := objc.CallMethod[HTTPCookieStringPolicy](h_, objc.GetSelector("sameSitePolicy"))
	return rv
}

func (h_ HTTPCookie) Properties() map[HTTPCookiePropertyKey]objc.Object {
	rv := objc.CallMethod[map[HTTPCookiePropertyKey]objc.Object](h_, objc.GetSelector("properties"))
	return rv
}

func (h_ HTTPCookie) Comment() string {
	rv := objc.CallMethod[string](h_, objc.GetSelector("comment"))
	return rv
}

func (h_ HTTPCookie) CommentURL() URL {
	rv := objc.CallMethod[URL](h_, objc.GetSelector("commentURL"))
	return rv
}
