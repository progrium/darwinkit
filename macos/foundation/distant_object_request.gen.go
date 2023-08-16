// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DistantObjectRequest] class.
var DistantObjectRequestClass = _DistantObjectRequestClass{objc.GetClass("NSDistantObjectRequest")}

type _DistantObjectRequestClass struct {
	objc.Class
}

// An interface definition for the [DistantObjectRequest] class.
type IDistantObjectRequest interface {
	objc.IObject
}

// An object used by the distributed objects system to help handle invocations between different processes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdistantobjectrequest?language=objc
type DistantObjectRequest struct {
	objc.Object
}

func DistantObjectRequestFrom(ptr unsafe.Pointer) DistantObjectRequest {
	return DistantObjectRequest{
		Object: objc.ObjectFrom(ptr),
	}
}

func (dc _DistantObjectRequestClass) Alloc() DistantObjectRequest {
	rv := objc.Call[DistantObjectRequest](dc, objc.Sel("alloc"))
	return rv
}

func DistantObjectRequest_Alloc() DistantObjectRequest {
	return DistantObjectRequestClass.Alloc()
}

func (dc _DistantObjectRequestClass) New() DistantObjectRequest {
	rv := objc.Call[DistantObjectRequest](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDistantObjectRequest() DistantObjectRequest {
	return DistantObjectRequestClass.New()
}

func (d_ DistantObjectRequest) Init() DistantObjectRequest {
	rv := objc.Call[DistantObjectRequest](d_, objc.Sel("init"))
	return rv
}
