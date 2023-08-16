// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FetchRecordZoneChangesOptions] class.
var FetchRecordZoneChangesOptionsClass = _FetchRecordZoneChangesOptionsClass{objc.GetClass("CKFetchRecordZoneChangesOptions")}

type _FetchRecordZoneChangesOptionsClass struct {
	objc.Class
}

// An interface definition for the [FetchRecordZoneChangesOptions] class.
type IFetchRecordZoneChangesOptions interface {
	objc.IObject
}

// A configuration object that describes the information to fetch from a record zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoptions?language=objc
type FetchRecordZoneChangesOptions struct {
	objc.Object
}

func FetchRecordZoneChangesOptionsFrom(ptr unsafe.Pointer) FetchRecordZoneChangesOptions {
	return FetchRecordZoneChangesOptions{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FetchRecordZoneChangesOptionsClass) Alloc() FetchRecordZoneChangesOptions {
	rv := objc.Call[FetchRecordZoneChangesOptions](fc, objc.Sel("alloc"))
	return rv
}

func FetchRecordZoneChangesOptions_Alloc() FetchRecordZoneChangesOptions {
	return FetchRecordZoneChangesOptionsClass.Alloc()
}

func (fc _FetchRecordZoneChangesOptionsClass) New() FetchRecordZoneChangesOptions {
	rv := objc.Call[FetchRecordZoneChangesOptions](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFetchRecordZoneChangesOptions() FetchRecordZoneChangesOptions {
	return FetchRecordZoneChangesOptionsClass.New()
}

func (f_ FetchRecordZoneChangesOptions) Init() FetchRecordZoneChangesOptions {
	rv := objc.Call[FetchRecordZoneChangesOptions](f_, objc.Sel("init"))
	return rv
}
