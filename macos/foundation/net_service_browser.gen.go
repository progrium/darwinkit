// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NetServiceBrowser] class.
var NetServiceBrowserClass = _NetServiceBrowserClass{objc.GetClass("NSNetServiceBrowser")}

type _NetServiceBrowserClass struct {
	objc.Class
}

// An interface definition for the [NetServiceBrowser] class.
type INetServiceBrowser interface {
	objc.IObject
}

// A network service browser that finds published services on a network using multicast DNS. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicebrowser?language=objc
type NetServiceBrowser struct {
	objc.Object
}

func NetServiceBrowserFrom(ptr unsafe.Pointer) NetServiceBrowser {
	return NetServiceBrowser{
		Object: objc.ObjectFrom(ptr),
	}
}

func (nc _NetServiceBrowserClass) Alloc() NetServiceBrowser {
	rv := objc.Call[NetServiceBrowser](nc, objc.Sel("alloc"))
	return rv
}

func NetServiceBrowser_Alloc() NetServiceBrowser {
	return NetServiceBrowserClass.Alloc()
}

func (nc _NetServiceBrowserClass) New() NetServiceBrowser {
	rv := objc.Call[NetServiceBrowser](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNetServiceBrowser() NetServiceBrowser {
	return NetServiceBrowserClass.New()
}

func (n_ NetServiceBrowser) Init() NetServiceBrowser {
	rv := objc.Call[NetServiceBrowser](n_, objc.Sel("init"))
	return rv
}
