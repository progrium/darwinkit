// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NavigationResponse] class.
var NavigationResponseClass = _NavigationResponseClass{objc.GetClass("WKNavigationResponse")}

type _NavigationResponseClass struct {
	objc.Class
}

// An interface definition for the [NavigationResponse] class.
type INavigationResponse interface {
	objc.IObject
	IsForMainFrame() bool
	Response() foundation.URLResponse
	CanShowMIMEType() bool
}

// An object that contains the response to a navigation request, and which you use to make navigation-related policy decisions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationresponse?language=objc
type NavigationResponse struct {
	objc.Object
}

func NavigationResponseFrom(ptr unsafe.Pointer) NavigationResponse {
	return NavigationResponse{
		Object: objc.ObjectFrom(ptr),
	}
}

func (nc _NavigationResponseClass) Alloc() NavigationResponse {
	rv := objc.Call[NavigationResponse](nc, objc.Sel("alloc"))
	return rv
}

func NavigationResponse_Alloc() NavigationResponse {
	return NavigationResponseClass.Alloc()
}

func (nc _NavigationResponseClass) New() NavigationResponse {
	rv := objc.Call[NavigationResponse](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNavigationResponse() NavigationResponse {
	return NavigationResponseClass.New()
}

func (n_ NavigationResponse) Init() NavigationResponse {
	rv := objc.Call[NavigationResponse](n_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether the response targets the web view’s main frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationresponse/1459482-formainframe?language=objc
func (n_ NavigationResponse) IsForMainFrame() bool {
	rv := objc.Call[bool](n_, objc.Sel("isForMainFrame"))
	return rv
}

// The frame’s response. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationresponse/1459484-response?language=objc
func (n_ NavigationResponse) Response() foundation.URLResponse {
	rv := objc.Call[foundation.URLResponse](n_, objc.Sel("response"))
	return rv
}

// A Boolean value that indicates whether WebKit is capable of displaying the response’s MIME type natively. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationresponse/1459480-canshowmimetype?language=objc
func (n_ NavigationResponse) CanShowMIMEType() bool {
	rv := objc.Call[bool](n_, objc.Sel("canShowMIMEType"))
	return rv
}
