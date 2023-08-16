// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Host] class.
var HostClass = _HostClass{objc.GetClass("NSHost")}

type _HostClass struct {
	objc.Class
}

// An interface definition for the [Host] class.
type IHost interface {
	objc.IObject
}

// A representation of an individual host on the network. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshost?language=objc
type Host struct {
	objc.Object
}

func HostFrom(ptr unsafe.Pointer) Host {
	return Host{
		Object: objc.ObjectFrom(ptr),
	}
}

func (hc _HostClass) Alloc() Host {
	rv := objc.Call[Host](hc, objc.Sel("alloc"))
	return rv
}

func Host_Alloc() Host {
	return HostClass.Alloc()
}

func (hc _HostClass) New() Host {
	rv := objc.Call[Host](hc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewHost() Host {
	return HostClass.New()
}

func (h_ Host) Init() Host {
	rv := objc.Call[Host](h_, objc.Sel("init"))
	return rv
}
