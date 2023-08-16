// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NetService] class.
var NetServiceClass = _NetServiceClass{objc.GetClass("NSNetService")}

type _NetServiceClass struct {
	objc.Class
}

// An interface definition for the [NetService] class.
type INetService interface {
	objc.IObject
}

// A network service that broadcasts its availability using multicast DNS. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservice?language=objc
type NetService struct {
	objc.Object
}

func NetServiceFrom(ptr unsafe.Pointer) NetService {
	return NetService{
		Object: objc.ObjectFrom(ptr),
	}
}

func (nc _NetServiceClass) Alloc() NetService {
	rv := objc.Call[NetService](nc, objc.Sel("alloc"))
	return rv
}

func NetService_Alloc() NetService {
	return NetServiceClass.Alloc()
}

func (nc _NetServiceClass) New() NetService {
	rv := objc.Call[NetService](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNetService() NetService {
	return NetServiceClass.New()
}

func (n_ NetService) Init() NetService {
	rv := objc.Call[NetService](n_, objc.Sel("init"))
	return rv
}
