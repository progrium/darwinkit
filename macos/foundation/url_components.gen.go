// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [URLComponents] class.
var URLComponentsClass = _URLComponentsClass{objc.GetClass("NSURLComponents")}

type _URLComponentsClass struct {
	objc.Class
}

// An interface definition for the [URLComponents] class.
type IURLComponents interface {
	objc.IObject
	URLRelativeToURL(baseURL IURL) URL
	Password() string
	SetPassword(value string)
	Path() string
	SetPath(value string)
	RangeOfPort() Range
	PercentEncodedUser() string
	SetPercentEncodedUser(value string)
	PercentEncodedPassword() string
	SetPercentEncodedPassword(value string)
	Host() string
	SetHost(value string)
	RangeOfPassword() Range
	RangeOfScheme() Range
	Scheme() string
	SetScheme(value string)
	RangeOfUser() Range
	QueryItems() []URLQueryItem
	SetQueryItems(value []IURLQueryItem)
	Query() string
	SetQuery(value string)
	RangeOfFragment() Range
	PercentEncodedQuery() string
	SetPercentEncodedQuery(value string)
	PercentEncodedFragment() string
	SetPercentEncodedFragment(value string)
	Fragment() string
	SetFragment(value string)
	PercentEncodedQueryItems() []URLQueryItem
	SetPercentEncodedQueryItems(value []IURLQueryItem)
	String() string
	PercentEncodedPath() string
	SetPercentEncodedPath(value string)
	RangeOfQuery() Range
	User() string
	SetUser(value string)
	URL() URL
	RangeOfHost() Range
	Port() Number
	SetPort(value INumber)
	RangeOfPath() Range
}

// An object that parses URLs into and constructs URLs from their constituent parts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents?language=objc
type URLComponents struct {
	objc.Object
}

func URLComponentsFrom(ptr unsafe.Pointer) URLComponents {
	return URLComponents{
		Object: objc.ObjectFrom(ptr),
	}
}

func (uc _URLComponentsClass) ComponentsWithURLResolvingAgainstBaseURL(url IURL, resolve bool) URLComponents {
	rv := objc.Call[URLComponents](uc, objc.Sel("componentsWithURL:resolvingAgainstBaseURL:"), objc.Ptr(url), resolve)
	return rv
}

// Returns a URL components object by parsing the URL from an NSURL object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1572050-componentswithurl?language=objc
func URLComponents_ComponentsWithURLResolvingAgainstBaseURL(url IURL, resolve bool) URLComponents {
	return URLComponentsClass.ComponentsWithURLResolvingAgainstBaseURL(url, resolve)
}

func (uc _URLComponentsClass) ComponentsWithString(URLString string) URLComponents {
	rv := objc.Call[URLComponents](uc, objc.Sel("componentsWithString:"), URLString)
	return rv
}

// Returns a URL components object by parsing a URL in string form. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1572054-componentswithstring?language=objc
func URLComponents_ComponentsWithString(URLString string) URLComponents {
	return URLComponentsClass.ComponentsWithString(URLString)
}

func (u_ URLComponents) InitWithURLResolvingAgainstBaseURL(url IURL, resolve bool) URLComponents {
	rv := objc.Call[URLComponents](u_, objc.Sel("initWithURL:resolvingAgainstBaseURL:"), objc.Ptr(url), resolve)
	return rv
}

// Initializes a URL components object by parsing the URL from an NSURL object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1416476-initwithurl?language=objc
func NewURLComponentsWithURLResolvingAgainstBaseURL(url IURL, resolve bool) URLComponents {
	instance := URLComponentsClass.Alloc().InitWithURLResolvingAgainstBaseURL(url, resolve)
	instance.Autorelease()
	return instance
}

func (u_ URLComponents) InitWithString(URLString string) URLComponents {
	rv := objc.Call[URLComponents](u_, objc.Sel("initWithString:"), URLString)
	return rv
}

// Initializes a URL components object by parsing a URL in string form. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1410784-initwithstring?language=objc
func NewURLComponentsWithString(URLString string) URLComponents {
	instance := URLComponentsClass.Alloc().InitWithString(URLString)
	instance.Autorelease()
	return instance
}

func (u_ URLComponents) Init() URLComponents {
	rv := objc.Call[URLComponents](u_, objc.Sel("init"))
	return rv
}

func (uc _URLComponentsClass) Alloc() URLComponents {
	rv := objc.Call[URLComponents](uc, objc.Sel("alloc"))
	return rv
}

func URLComponents_Alloc() URLComponents {
	return URLComponentsClass.Alloc()
}

func (uc _URLComponentsClass) New() URLComponents {
	rv := objc.Call[URLComponents](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewURLComponents() URLComponents {
	return URLComponentsClass.New()
}

// Returns a URL object derived from the components object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1408378-urlrelativetourl?language=objc
func (u_ URLComponents) URLRelativeToURL(baseURL IURL) URL {
	rv := objc.Call[URL](u_, objc.Sel("URLRelativeToURL:"), objc.Ptr(baseURL))
	return rv
}

// The password URL subcomponent, or nil if not present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1415604-password?language=objc
func (u_ URLComponents) Password() string {
	rv := objc.Call[string](u_, objc.Sel("password"))
	return rv
}

// The password URL subcomponent, or nil if not present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1415604-password?language=objc
func (u_ URLComponents) SetPassword(value string) {
	objc.Call[objc.Void](u_, objc.Sel("setPassword:"), value)
}

// The path URL component, or nil if not present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1409650-path?language=objc
func (u_ URLComponents) Path() string {
	rv := objc.Call[string](u_, objc.Sel("path"))
	return rv
}

// The path URL component, or nil if not present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1409650-path?language=objc
func (u_ URLComponents) SetPath(value string) {
	objc.Call[objc.Void](u_, objc.Sel("setPath:"), value)
}

// Returns the character range of the port in the string returned by the string property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1411790-rangeofport?language=objc
func (u_ URLComponents) RangeOfPort() Range {
	rv := objc.Call[Range](u_, objc.Sel("rangeOfPort"))
	return rv
}

// The username URL subcomponent expressed as a URL-encoded string, or nil if not present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1417767-percentencodeduser?language=objc
func (u_ URLComponents) PercentEncodedUser() string {
	rv := objc.Call[string](u_, objc.Sel("percentEncodedUser"))
	return rv
}

// The username URL subcomponent expressed as a URL-encoded string, or nil if not present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1417767-percentencodeduser?language=objc
func (u_ URLComponents) SetPercentEncodedUser(value string) {
	objc.Call[objc.Void](u_, objc.Sel("setPercentEncodedUser:"), value)
}

// The password URL subcomponent expressed as a URL-encoded string, or nil if not present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1410319-percentencodedpassword?language=objc
func (u_ URLComponents) PercentEncodedPassword() string {
	rv := objc.Call[string](u_, objc.Sel("percentEncodedPassword"))
	return rv
}

// The password URL subcomponent expressed as a URL-encoded string, or nil if not present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1410319-percentencodedpassword?language=objc
func (u_ URLComponents) SetPercentEncodedPassword(value string) {
	objc.Call[objc.Void](u_, objc.Sel("setPercentEncodedPassword:"), value)
}

// The host URL subcomponent, or nil if not present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1411178-host?language=objc
func (u_ URLComponents) Host() string {
	rv := objc.Call[string](u_, objc.Sel("host"))
	return rv
}

// The host URL subcomponent, or nil if not present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1411178-host?language=objc
func (u_ URLComponents) SetHost(value string) {
	objc.Call[objc.Void](u_, objc.Sel("setHost:"), value)
}

// Returns the character range of the password in the string returned by the string property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1415024-rangeofpassword?language=objc
func (u_ URLComponents) RangeOfPassword() Range {
	rv := objc.Call[Range](u_, objc.Sel("rangeOfPassword"))
	return rv
}

// Returns the character range of the scheme in the string returned by the string property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1410099-rangeofscheme?language=objc
func (u_ URLComponents) RangeOfScheme() Range {
	rv := objc.Call[Range](u_, objc.Sel("rangeOfScheme"))
	return rv
}

// The scheme URL component, or nil if not present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1407517-scheme?language=objc
func (u_ URLComponents) Scheme() string {
	rv := objc.Call[string](u_, objc.Sel("scheme"))
	return rv
}

// The scheme URL component, or nil if not present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1407517-scheme?language=objc
func (u_ URLComponents) SetScheme(value string) {
	objc.Call[objc.Void](u_, objc.Sel("setScheme:"), value)
}

// Returns the character range of the user in the string returned by the string property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1414961-rangeofuser?language=objc
func (u_ URLComponents) RangeOfUser() Range {
	rv := objc.Call[Range](u_, objc.Sel("rangeOfUser"))
	return rv
}

// The query URL component as an array of name/value pairs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1407752-queryitems?language=objc
func (u_ URLComponents) QueryItems() []URLQueryItem {
	rv := objc.Call[[]URLQueryItem](u_, objc.Sel("queryItems"))
	return rv
}

// The query URL component as an array of name/value pairs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1407752-queryitems?language=objc
func (u_ URLComponents) SetQueryItems(value []IURLQueryItem) {
	objc.Call[objc.Void](u_, objc.Sel("setQueryItems:"), value)
}

// The query URL component as a string, or nil if not present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1415452-query?language=objc
func (u_ URLComponents) Query() string {
	rv := objc.Call[string](u_, objc.Sel("query"))
	return rv
}

// The query URL component as a string, or nil if not present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1415452-query?language=objc
func (u_ URLComponents) SetQuery(value string) {
	objc.Call[objc.Void](u_, objc.Sel("setQuery:"), value)
}

// Returns the character range of the fragment in the string returned by the string property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1415180-rangeoffragment?language=objc
func (u_ URLComponents) RangeOfFragment() Range {
	rv := objc.Call[Range](u_, objc.Sel("rangeOfFragment"))
	return rv
}

// The query URL component expressed as a URL-encoded string, or nil if not present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1410395-percentencodedquery?language=objc
func (u_ URLComponents) PercentEncodedQuery() string {
	rv := objc.Call[string](u_, objc.Sel("percentEncodedQuery"))
	return rv
}

// The query URL component expressed as a URL-encoded string, or nil if not present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1410395-percentencodedquery?language=objc
func (u_ URLComponents) SetPercentEncodedQuery(value string) {
	objc.Call[objc.Void](u_, objc.Sel("setPercentEncodedQuery:"), value)
}

// The fragment URL component (the part after a # symbol) expressed as a URL-encoded string, or nil if not present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1418392-percentencodedfragment?language=objc
func (u_ URLComponents) PercentEncodedFragment() string {
	rv := objc.Call[string](u_, objc.Sel("percentEncodedFragment"))
	return rv
}

// The fragment URL component (the part after a # symbol) expressed as a URL-encoded string, or nil if not present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1418392-percentencodedfragment?language=objc
func (u_ URLComponents) SetPercentEncodedFragment(value string) {
	objc.Call[objc.Void](u_, objc.Sel("setPercentEncodedFragment:"), value)
}

// The fragment URL component (the part after a # symbol), or nil if not present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1417638-fragment?language=objc
func (u_ URLComponents) Fragment() string {
	rv := objc.Call[string](u_, objc.Sel("fragment"))
	return rv
}

// The fragment URL component (the part after a # symbol), or nil if not present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1417638-fragment?language=objc
func (u_ URLComponents) SetFragment(value string) {
	objc.Call[objc.Void](u_, objc.Sel("setFragment:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/2865531-percentencodedqueryitems?language=objc
func (u_ URLComponents) PercentEncodedQueryItems() []URLQueryItem {
	rv := objc.Call[[]URLQueryItem](u_, objc.Sel("percentEncodedQueryItems"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/2865531-percentencodedqueryitems?language=objc
func (u_ URLComponents) SetPercentEncodedQueryItems(value []IURLQueryItem) {
	objc.Call[objc.Void](u_, objc.Sel("setPercentEncodedQueryItems:"), value)
}

// A URL derived from the components object, in string form. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1417970-string?language=objc
func (u_ URLComponents) String() string {
	rv := objc.Call[string](u_, objc.Sel("string"))
	return rv
}

// The path URL component expressed as a URL-encoded string, or nil if not present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1408161-percentencodedpath?language=objc
func (u_ URLComponents) PercentEncodedPath() string {
	rv := objc.Call[string](u_, objc.Sel("percentEncodedPath"))
	return rv
}

// The path URL component expressed as a URL-encoded string, or nil if not present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1408161-percentencodedpath?language=objc
func (u_ URLComponents) SetPercentEncodedPath(value string) {
	objc.Call[objc.Void](u_, objc.Sel("setPercentEncodedPath:"), value)
}

// Returns the character range of the query in the string returned by the string property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1409456-rangeofquery?language=objc
func (u_ URLComponents) RangeOfQuery() Range {
	rv := objc.Call[Range](u_, objc.Sel("rangeOfQuery"))
	return rv
}

// The username URL subcomponent, or nil if not present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1415026-user?language=objc
func (u_ URLComponents) User() string {
	rv := objc.Call[string](u_, objc.Sel("user"))
	return rv
}

// The username URL subcomponent, or nil if not present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1415026-user?language=objc
func (u_ URLComponents) SetUser(value string) {
	objc.Call[objc.Void](u_, objc.Sel("setUser:"), value)
}

// A URL object derived from the components object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1413469-url?language=objc
func (u_ URLComponents) URL() URL {
	rv := objc.Call[URL](u_, objc.Sel("URL"))
	return rv
}

// Returns the character range of the host in the string returned by the string property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1408894-rangeofhost?language=objc
func (u_ URLComponents) RangeOfHost() Range {
	rv := objc.Call[Range](u_, objc.Sel("rangeOfHost"))
	return rv
}

// The port number URL component, or nil if not present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1413451-port?language=objc
func (u_ URLComponents) Port() Number {
	rv := objc.Call[Number](u_, objc.Sel("port"))
	return rv
}

// The port number URL component, or nil if not present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1413451-port?language=objc
func (u_ URLComponents) SetPort(value INumber) {
	objc.Call[objc.Void](u_, objc.Sel("setPort:"), objc.Ptr(value))
}

// Returns the character range of the path in the string returned by the string property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/1418459-rangeofpath?language=objc
func (u_ URLComponents) RangeOfPath() Range {
	rv := objc.Call[Range](u_, objc.Sel("rangeOfPath"))
	return rv
}
