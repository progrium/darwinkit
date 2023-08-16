// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [WebDownload] class.
var WebDownloadClass = _WebDownloadClass{objc.GetClass("WebDownload")}

type _WebDownloadClass struct {
	objc.Class
}

// An interface definition for the [WebDownload] class.
type IWebDownload interface {
	foundation.IURLDownload
}

// WebDownload objects initiate download client requests on behalf of a delegate. A download request involves loading the data, decoding it (if necessary), and saving it to a file. Instances of this class behave similar to NSURLDownload except delegates of WebDownload may implement an additional delegate method. The method allows the delegate to specify the window to be used for authentication sheets. If the delegate does not implement this method, the WebDownload object will prompt the user for authentication using the standard WebKit authentication panel, as either a sheet or window. There are no additional methods defined in this class. See WebDownloadDelegate for the delegate method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webdownload?language=objc
type WebDownload struct {
	foundation.URLDownload
}

func WebDownloadFrom(ptr unsafe.Pointer) WebDownload {
	return WebDownload{
		URLDownload: foundation.URLDownloadFrom(ptr),
	}
}

func (wc _WebDownloadClass) Alloc() WebDownload {
	rv := objc.Call[WebDownload](wc, objc.Sel("alloc"))
	return rv
}

func WebDownload_Alloc() WebDownload {
	return WebDownloadClass.Alloc()
}

func (wc _WebDownloadClass) New() WebDownload {
	rv := objc.Call[WebDownload](wc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewWebDownload() WebDownload {
	return WebDownloadClass.New()
}

func (w_ WebDownload) Init() WebDownload {
	rv := objc.Call[WebDownload](w_, objc.Sel("init"))
	return rv
}
