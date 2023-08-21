// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [URLResponse] class.
var URLResponseClass = _URLResponseClass{objc.GetClass("NSURLResponse")}

type _URLResponseClass struct {
	objc.Class
}

// An interface definition for the [URLResponse] class.
type IURLResponse interface {
	objc.IObject
	ExpectedContentLength() int64
	MIMEType() string
	URL() URL
	TextEncodingName() string
	SuggestedFilename() string
}

// The metadata associated with the response to a URL load request, independent of protocol and URL scheme. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlresponse?language=objc
type URLResponse struct {
	objc.Object
}

func URLResponseFrom(ptr unsafe.Pointer) URLResponse {
	return URLResponse{
		Object: objc.ObjectFrom(ptr),
	}
}

func (u_ URLResponse) InitWithURLMIMETypeExpectedContentLengthTextEncodingName(URL IURL, MIMEType string, length int, name string) URLResponse {
	rv := objc.Call[URLResponse](u_, objc.Sel("initWithURL:MIMEType:expectedContentLength:textEncodingName:"), objc.Ptr(URL), MIMEType, length, name)
	return rv
}

// Creates an initialized NSURLResponse object with the URL, MIME type, length, and text encoding set to given values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlresponse/1413566-initwithurl?language=objc
func NewURLResponseWithURLMIMETypeExpectedContentLengthTextEncodingName(URL IURL, MIMEType string, length int, name string) URLResponse {
	instance := URLResponseClass.Alloc().InitWithURLMIMETypeExpectedContentLengthTextEncodingName(URL, MIMEType, length, name)
	instance.Autorelease()
	return instance
}

func (uc _URLResponseClass) Alloc() URLResponse {
	rv := objc.Call[URLResponse](uc, objc.Sel("alloc"))
	return rv
}

func URLResponse_Alloc() URLResponse {
	return URLResponseClass.Alloc()
}

func (uc _URLResponseClass) New() URLResponse {
	rv := objc.Call[URLResponse](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewURLResponse() URLResponse {
	return URLResponseClass.New()
}

func (u_ URLResponse) Init() URLResponse {
	rv := objc.Call[URLResponse](u_, objc.Sel("init"))
	return rv
}

// The expected length of the response’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlresponse/1413507-expectedcontentlength?language=objc
func (u_ URLResponse) ExpectedContentLength() int64 {
	rv := objc.Call[int64](u_, objc.Sel("expectedContentLength"))
	return rv
}

// The MIME type of the response. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlresponse/1411613-mimetype?language=objc
func (u_ URLResponse) MIMEType() string {
	rv := objc.Call[string](u_, objc.Sel("MIMEType"))
	return rv
}

// The URL for the response. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlresponse/1414219-url?language=objc
func (u_ URLResponse) URL() URL {
	rv := objc.Call[URL](u_, objc.Sel("URL"))
	return rv
}

// The name of the text encoding provided by the response’s originating source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlresponse/1408005-textencodingname?language=objc
func (u_ URLResponse) TextEncodingName() string {
	rv := objc.Call[string](u_, objc.Sel("textEncodingName"))
	return rv
}

// A suggested filename for the response data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlresponse/1415924-suggestedfilename?language=objc
func (u_ URLResponse) SuggestedFilename() string {
	rv := objc.Call[string](u_, objc.Sel("suggestedFilename"))
	return rv
}
