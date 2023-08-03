// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var URLResponseClass = _URLResponseClass{objc.GetClass("NSURLResponse")}

type _URLResponseClass struct {
	objc.Class
}

type IURLResponse interface {
	objc.IObject
	ExpectedContentLength() int64
	SuggestedFilename() string
	MIMEType() string
	TextEncodingName() string
	URL() URL
}

type URLResponse struct {
	objc.Object
}

func MakeURLResponse(ptr unsafe.Pointer) URLResponse {
	return URLResponse{
		Object: objc.MakeObject(ptr),
	}
}

func (u_ URLResponse) InitWithURLMIMETypeExpectedContentLengthTextEncodingName(URL IURL, MIMEType string, length int, name string) URLResponse {
	rv := objc.CallMethod[URLResponse](u_, objc.GetSelector("initWithURL:MIMEType:expectedContentLength:textEncodingName:"), objc.ExtractPtr(URL), MIMEType, length, name)
	return rv
}

func URLResponse_InitWithURLMIMETypeExpectedContentLengthTextEncodingName(URL IURL, MIMEType string, length int, name string) URLResponse {
	return URLResponseClass.Alloc().InitWithURLMIMETypeExpectedContentLengthTextEncodingName(URL, MIMEType, length, name)
}

func (uc _URLResponseClass) Alloc() URLResponse {
	rv := objc.CallMethod[URLResponse](uc, objc.GetSelector("alloc"))
	return rv
}

func URLResponse_Alloc() URLResponse {
	return URLResponseClass.Alloc()
}

func (uc _URLResponseClass) New() URLResponse {
	rv := objc.CallMethod[URLResponse](uc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewURLResponse() URLResponse {
	return URLResponseClass.New()
}

func URLResponse_New() URLResponse {
	return URLResponseClass.New()
}

func (u_ URLResponse) Init() URLResponse {
	rv := objc.CallMethod[URLResponse](u_, objc.GetSelector("init"))
	return rv
}

func URLResponse_Init() URLResponse {
	return URLResponseClass.Alloc().Init()
}

func (u_ URLResponse) ExpectedContentLength() int64 {
	rv := objc.CallMethod[int64](u_, objc.GetSelector("expectedContentLength"))
	return rv
}

func (u_ URLResponse) SuggestedFilename() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("suggestedFilename"))
	return rv
}

func (u_ URLResponse) MIMEType() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("MIMEType"))
	return rv
}

func (u_ URLResponse) TextEncodingName() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("textEncodingName"))
	return rv
}

func (u_ URLResponse) URL() URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("URL"))
	return rv
}
