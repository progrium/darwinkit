// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var FrameInfoClass = _FrameInfoClass{objc.GetClass("WKFrameInfo")}

type _FrameInfoClass struct {
	objc.Class
}

type IFrameInfo interface {
	objc.IObject
	IsMainFrame() bool
	Request() foundation.URLRequest
	SecurityOrigin() SecurityOrigin
	WebView() WebView
}

type FrameInfo struct {
	objc.Object
}

func MakeFrameInfo(ptr unsafe.Pointer) FrameInfo {
	return FrameInfo{
		Object: objc.MakeObject(ptr),
	}
}

func (fc _FrameInfoClass) Alloc() FrameInfo {
	rv := objc.CallMethod[FrameInfo](fc, objc.GetSelector("alloc"))
	return rv
}

func FrameInfo_Alloc() FrameInfo {
	return FrameInfoClass.Alloc()
}

func (fc _FrameInfoClass) New() FrameInfo {
	rv := objc.CallMethod[FrameInfo](fc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewFrameInfo() FrameInfo {
	return FrameInfoClass.New()
}

func FrameInfo_New() FrameInfo {
	return FrameInfoClass.New()
}

func (f_ FrameInfo) Init() FrameInfo {
	rv := objc.CallMethod[FrameInfo](f_, objc.GetSelector("init"))
	return rv
}

func FrameInfo_Init() FrameInfo {
	return FrameInfoClass.Alloc().Init()
}

func (f_ FrameInfo) IsMainFrame() bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("isMainFrame"))
	return rv
}

func (f_ FrameInfo) Request() foundation.URLRequest {
	rv := objc.CallMethod[foundation.URLRequest](f_, objc.GetSelector("request"))
	return rv
}

func (f_ FrameInfo) SecurityOrigin() SecurityOrigin {
	rv := objc.CallMethod[SecurityOrigin](f_, objc.GetSelector("securityOrigin"))
	return rv
}

func (f_ FrameInfo) WebView() WebView {
	rv := objc.CallMethod[WebView](f_, objc.GetSelector("webView"))
	return rv
}
