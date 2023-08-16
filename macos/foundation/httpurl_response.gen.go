// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [HTTPURLResponse] class.
var HTTPURLResponseClass = _HTTPURLResponseClass{objc.GetClass("NSHTTPURLResponse")}

type _HTTPURLResponseClass struct {
	objc.Class
}

// An interface definition for the [HTTPURLResponse] class.
type IHTTPURLResponse interface {
	IURLResponse
	ValueForHTTPHeaderField(field string) string
	StatusCode() int
	AllHeaderFields() Dictionary
}

// The metadata associated with the response to an HTTP protocol URL load request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpurlresponse?language=objc
type HTTPURLResponse struct {
	URLResponse
}

func HTTPURLResponseFrom(ptr unsafe.Pointer) HTTPURLResponse {
	return HTTPURLResponse{
		URLResponse: URLResponseFrom(ptr),
	}
}

func (h_ HTTPURLResponse) InitWithURLStatusCodeHTTPVersionHeaderFields(url IURL, statusCode int, HTTPVersion string, headerFields map[string]string) HTTPURLResponse {
	rv := objc.Call[HTTPURLResponse](h_, objc.Sel("initWithURL:statusCode:HTTPVersion:headerFields:"), objc.Ptr(url), statusCode, HTTPVersion, headerFields)
	return rv
}

// Initializes an HTTP URL response object with a status code, protocol version, and response headers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpurlresponse/1415870-initwithurl?language=objc
func HTTPURLResponse_InitWithURLStatusCodeHTTPVersionHeaderFields(url IURL, statusCode int, HTTPVersion string, headerFields map[string]string) HTTPURLResponse {
	return HTTPURLResponseClass.Alloc().InitWithURLStatusCodeHTTPVersionHeaderFields(url, statusCode, HTTPVersion, headerFields)
}

func (hc _HTTPURLResponseClass) Alloc() HTTPURLResponse {
	rv := objc.Call[HTTPURLResponse](hc, objc.Sel("alloc"))
	return rv
}

func HTTPURLResponse_Alloc() HTTPURLResponse {
	return HTTPURLResponseClass.Alloc()
}

func (hc _HTTPURLResponseClass) New() HTTPURLResponse {
	rv := objc.Call[HTTPURLResponse](hc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewHTTPURLResponse() HTTPURLResponse {
	return HTTPURLResponseClass.New()
}

func (h_ HTTPURLResponse) Init() HTTPURLResponse {
	rv := objc.Call[HTTPURLResponse](h_, objc.Sel("init"))
	return rv
}

func (h_ HTTPURLResponse) InitWithURLMIMETypeExpectedContentLengthTextEncodingName(URL IURL, MIMEType string, length int, name string) HTTPURLResponse {
	rv := objc.Call[HTTPURLResponse](h_, objc.Sel("initWithURL:MIMEType:expectedContentLength:textEncodingName:"), objc.Ptr(URL), MIMEType, length, name)
	return rv
}

// Creates an initialized NSURLResponse object with the URL, MIME type, length, and text encoding set to given values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlresponse/1413566-initwithurl?language=objc
func HTTPURLResponse_InitWithURLMIMETypeExpectedContentLengthTextEncodingName(URL IURL, MIMEType string, length int, name string) HTTPURLResponse {
	return HTTPURLResponseClass.Alloc().InitWithURLMIMETypeExpectedContentLengthTextEncodingName(URL, MIMEType, length, name)
}

// Returns a localized string corresponding to a specified HTTP status code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpurlresponse/1409741-localizedstringforstatuscode?language=objc
func (hc _HTTPURLResponseClass) LocalizedStringForStatusCode(statusCode int) string {
	rv := objc.Call[string](hc, objc.Sel("localizedStringForStatusCode:"), statusCode)
	return rv
}

// Returns a localized string corresponding to a specified HTTP status code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpurlresponse/1409741-localizedstringforstatuscode?language=objc
func HTTPURLResponse_LocalizedStringForStatusCode(statusCode int) string {
	return HTTPURLResponseClass.LocalizedStringForStatusCode(statusCode)
}

// Returns the value that corresponds to the given header field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpurlresponse/3240613-valueforhttpheaderfield?language=objc
func (h_ HTTPURLResponse) ValueForHTTPHeaderField(field string) string {
	rv := objc.Call[string](h_, objc.Sel("valueForHTTPHeaderField:"), field)
	return rv
}

// The response’s HTTP status code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpurlresponse/1409395-statuscode?language=objc
func (h_ HTTPURLResponse) StatusCode() int {
	rv := objc.Call[int](h_, objc.Sel("statusCode"))
	return rv
}

// All HTTP header fields of the response. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshttpurlresponse/1417930-allheaderfields?language=objc
func (h_ HTTPURLResponse) AllHeaderFields() Dictionary {
	rv := objc.Call[Dictionary](h_, objc.Sel("allHeaderFields"))
	return rv
}
