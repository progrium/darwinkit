// AUTO-GENERATED CODE, DO NOT MODIFY
package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var NavigationResponseClass = _NavigationResponseClass{objc.GetClass("WKNavigationResponse")}

type _NavigationResponseClass struct {
	objc.Class
}

type INavigationResponse interface {
	objc.IObject
	Response() foundation.URLResponse
	CanShowMIMEType() bool
	IsForMainFrame() bool
}

type NavigationResponse struct {
	objc.Object
}

func MakeNavigationResponse(ptr unsafe.Pointer) NavigationResponse {
	return NavigationResponse{
		Object: objc.MakeObject(ptr),
	}
}

func (nc _NavigationResponseClass) Alloc() NavigationResponse {
	rv := objc.CallMethod[NavigationResponse](nc, objc.GetSelector("alloc"))
	return rv
}

func NavigationResponse_Alloc() NavigationResponse {
	return NavigationResponseClass.Alloc()
}

func (nc _NavigationResponseClass) New() NavigationResponse {
	rv := objc.CallMethod[NavigationResponse](nc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewNavigationResponse() NavigationResponse {
	return NavigationResponseClass.New()
}

func NavigationResponse_New() NavigationResponse {
	return NavigationResponseClass.New()
}

func (n_ NavigationResponse) Init() NavigationResponse {
	rv := objc.CallMethod[NavigationResponse](n_, objc.GetSelector("init"))
	return rv
}

func NavigationResponse_Init() NavigationResponse {
	return NavigationResponseClass.Alloc().Init()
}

func (n_ NavigationResponse) Response() foundation.URLResponse {
	rv := objc.CallMethod[foundation.URLResponse](n_, objc.GetSelector("response"))
	return rv
}

func (n_ NavigationResponse) CanShowMIMEType() bool {
	rv := objc.CallMethod[bool](n_, objc.GetSelector("canShowMIMEType"))
	return rv
}

func (n_ NavigationResponse) IsForMainFrame() bool {
	rv := objc.CallMethod[bool](n_, objc.GetSelector("isForMainFrame"))
	return rv
}
