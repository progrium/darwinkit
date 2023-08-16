// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Download] class.
var DownloadClass = _DownloadClass{objc.GetClass("WKDownload")}

type _DownloadClass struct {
	objc.Class
}

// An interface definition for the [Download] class.
type IDownload interface {
	objc.IObject
	Cancel(completionHandler func(resumeData []byte))
	Delegate() DownloadDelegateWrapper
	SetDelegate(value PDownloadDelegate)
	SetDelegateObject(valueObject objc.IObject)
	OriginalRequest() foundation.URLRequest
}

// An object that represents the download of a web resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownload?language=objc
type Download struct {
	objc.Object
}

func DownloadFrom(ptr unsafe.Pointer) Download {
	return Download{
		Object: objc.ObjectFrom(ptr),
	}
}

func (dc _DownloadClass) Alloc() Download {
	rv := objc.Call[Download](dc, objc.Sel("alloc"))
	return rv
}

func Download_Alloc() Download {
	return DownloadClass.Alloc()
}

func (dc _DownloadClass) New() Download {
	rv := objc.Call[Download](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDownload() Download {
	return DownloadClass.New()
}

func (d_ Download) Init() Download {
	rv := objc.Call[Download](d_, objc.Sel("init"))
	return rv
}

// Cancels the download, and optionally captures data so that you can resume the download later. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownload/3727339-cancel?language=objc
func (d_ Download) Cancel(completionHandler func(resumeData []byte)) {
	objc.Call[objc.Void](d_, objc.Sel("cancel:"), completionHandler)
}

// An object you use to track download progress and handle redirects, authentication challenges, and failures. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownload/3727340-delegate?language=objc
func (d_ Download) Delegate() DownloadDelegateWrapper {
	rv := objc.Call[DownloadDelegateWrapper](d_, objc.Sel("delegate"))
	return rv
}

// An object you use to track download progress and handle redirects, authentication challenges, and failures. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownload/3727340-delegate?language=objc
func (d_ Download) SetDelegate(value PDownloadDelegate) {
	po0 := objc.WrapAsProtocol("WKDownloadDelegate", value)
	objc.SetAssociatedObject(d_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](d_, objc.Sel("setDelegate:"), po0)
}

// An object you use to track download progress and handle redirects, authentication challenges, and failures. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownload/3727340-delegate?language=objc
func (d_ Download) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](d_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// An object that represents the request that initiated the download. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownload/3727341-originalrequest?language=objc
func (d_ Download) OriginalRequest() foundation.URLRequest {
	rv := objc.Call[foundation.URLRequest](d_, objc.Sel("originalRequest"))
	return rv
}
