// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [WebResource] class.
var WebResourceClass = _WebResourceClass{objc.GetClass("WebResource")}

type _WebResourceClass struct {
	objc.Class
}

// An interface definition for the [WebResource] class.
type IWebResource interface {
	objc.IObject
	MIMEType() string
	FrameName() string
	Data() []byte
	URL() foundation.URL
	TextEncodingName() string
}

// A WebResource object represents a downloaded URL. It encapsulates the data of the download as well as other resource properties such as the URL, MIME type, and frame name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webresource?language=objc
type WebResource struct {
	objc.Object
}

func WebResourceFrom(ptr unsafe.Pointer) WebResource {
	return WebResource{
		Object: objc.ObjectFrom(ptr),
	}
}

func (w_ WebResource) InitWithDataURLMIMETypeTextEncodingNameFrameName(data []byte, URL foundation.IURL, MIMEType string, textEncodingName string, frameName string) WebResource {
	rv := objc.Call[WebResource](w_, objc.Sel("initWithData:URL:MIMEType:textEncodingName:frameName:"), data, objc.Ptr(URL), MIMEType, textEncodingName, frameName)
	return rv
}

// Initializes and returns a web resource instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webresource/1392185-initwithdata?language=objc
func NewWebResourceWithDataURLMIMETypeTextEncodingNameFrameName(data []byte, URL foundation.IURL, MIMEType string, textEncodingName string, frameName string) WebResource {
	instance := WebResourceClass.Alloc().InitWithDataURLMIMETypeTextEncodingNameFrameName(data, URL, MIMEType, textEncodingName, frameName)
	instance.Autorelease()
	return instance
}

func (wc _WebResourceClass) Alloc() WebResource {
	rv := objc.Call[WebResource](wc, objc.Sel("alloc"))
	return rv
}

func WebResource_Alloc() WebResource {
	return WebResourceClass.Alloc()
}

func (wc _WebResourceClass) New() WebResource {
	rv := objc.Call[WebResource](wc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewWebResource() WebResource {
	return WebResourceClass.New()
}

func (w_ WebResource) Init() WebResource {
	rv := objc.Call[WebResource](w_, objc.Sel("init"))
	return rv
}

// The receiver’s MIME type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webresource/1392187-mimetype?language=objc
func (w_ WebResource) MIMEType() string {
	rv := objc.Call[string](w_, objc.Sel("MIMEType"))
	return rv
}

// The name of the frame. If the receiver does not represent the contents of an entire HTML frame, this is nil. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webresource/1392193-framename?language=objc
func (w_ WebResource) FrameName() string {
	rv := objc.Call[string](w_, objc.Sel("frameName"))
	return rv
}

// The receiver’s data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webresource/1392191-data?language=objc
func (w_ WebResource) Data() []byte {
	rv := objc.Call[[]byte](w_, objc.Sel("data"))
	return rv
}

// The receiver’s URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webresource/1392195-url?language=objc
func (w_ WebResource) URL() foundation.URL {
	rv := objc.Call[foundation.URL](w_, objc.Sel("URL"))
	return rv
}

// The receiver’s text encoding name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webresource/1392189-textencodingname?language=objc
func (w_ WebResource) TextEncodingName() string {
	rv := objc.Call[string](w_, objc.Sel("textEncodingName"))
	return rv
}
