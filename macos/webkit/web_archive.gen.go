// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [WebArchive] class.
var WebArchiveClass = _WebArchiveClass{objc.GetClass("WebArchive")}

type _WebArchiveClass struct {
	objc.Class
}

// An interface definition for the [WebArchive] class.
type IWebArchive interface {
	objc.IObject
}

// A WebArchive object represents a webpage that can be archived—for example, archived on disk or on the pasteboard. A WebArchive object contains the main resource, as well as the subresources and subframes of the main resource. The main resource can be an entire webpage, a portion of a webpage, or some other kind of data such as an image. Use this class to archive webpages, or place a portion of a webpage on the pasteboard, or to represent rich web content in any application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webarchive?language=objc
type WebArchive struct {
	objc.Object
}

func WebArchiveFrom(ptr unsafe.Pointer) WebArchive {
	return WebArchive{
		Object: objc.ObjectFrom(ptr),
	}
}

func (wc _WebArchiveClass) Alloc() WebArchive {
	rv := objc.Call[WebArchive](wc, objc.Sel("alloc"))
	return rv
}

func WebArchive_Alloc() WebArchive {
	return WebArchiveClass.Alloc()
}

func (wc _WebArchiveClass) New() WebArchive {
	rv := objc.Call[WebArchive](wc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewWebArchive() WebArchive {
	return WebArchiveClass.New()
}

func (w_ WebArchive) Init() WebArchive {
	rv := objc.Call[WebArchive](w_, objc.Sel("init"))
	return rv
}
