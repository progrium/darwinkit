// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FetchRequest] class.
var FetchRequestClass = _FetchRequestClass{objc.GetClass("CNFetchRequest")}

type _FetchRequestClass struct {
	objc.Class
}

// An interface definition for the [FetchRequest] class.
type IFetchRequest interface {
	objc.IObject
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnfetchrequest?language=objc
type FetchRequest struct {
	objc.Object
}

func FetchRequestFrom(ptr unsafe.Pointer) FetchRequest {
	return FetchRequest{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FetchRequestClass) Alloc() FetchRequest {
	rv := objc.Call[FetchRequest](fc, objc.Sel("alloc"))
	return rv
}

func FetchRequest_Alloc() FetchRequest {
	return FetchRequestClass.Alloc()
}

func (fc _FetchRequestClass) New() FetchRequest {
	rv := objc.Call[FetchRequest](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFetchRequest() FetchRequest {
	return FetchRequestClass.New()
}

func (f_ FetchRequest) Init() FetchRequest {
	rv := objc.Call[FetchRequest](f_, objc.Sel("init"))
	return rv
}
