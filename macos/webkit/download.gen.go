// AUTO-GENERATED CODE, DO NOT MODIFY
package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var DownloadClass = _DownloadClass{objc.GetClass("WKDownload")}

type _DownloadClass struct {
	objc.Class
}

type IDownload interface {
	objc.IObject
	Cancel(completionHandler func(resumeData []byte))
	OriginalRequest() foundation.URLRequest
	WebView() WebView
}

type Download struct {
	objc.Object
}

func MakeDownload(ptr unsafe.Pointer) Download {
	return Download{
		Object: objc.MakeObject(ptr),
	}
}

func (dc _DownloadClass) Alloc() Download {
	rv := objc.CallMethod[Download](dc, objc.GetSelector("alloc"))
	return rv
}

func Download_Alloc() Download {
	return DownloadClass.Alloc()
}

func (dc _DownloadClass) New() Download {
	rv := objc.CallMethod[Download](dc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewDownload() Download {
	return DownloadClass.New()
}

func Download_New() Download {
	return DownloadClass.New()
}

func (d_ Download) Init() Download {
	rv := objc.CallMethod[Download](d_, objc.GetSelector("init"))
	return rv
}

func Download_Init() Download {
	return DownloadClass.Alloc().Init()
}

func (d_ Download) Cancel(completionHandler func(resumeData []byte)) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("cancel:"), completionHandler)
}

func (d_ Download) OriginalRequest() foundation.URLRequest {
	rv := objc.CallMethod[foundation.URLRequest](d_, objc.GetSelector("originalRequest"))
	return rv
}

func (d_ Download) WebView() WebView {
	rv := objc.CallMethod[WebView](d_, objc.GetSelector("webView"))
	return rv
}
