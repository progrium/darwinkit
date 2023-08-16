// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FetchResult] class.
var FetchResultClass = _FetchResultClass{objc.GetClass("CNFetchResult")}

type _FetchResultClass struct {
	objc.Class
}

// An interface definition for the [FetchResult] class.
type IFetchResult interface {
	objc.IObject
	Value() objc.Object
	CurrentHistoryToken() []byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnfetchresult?language=objc
type FetchResult struct {
	objc.Object
}

func FetchResultFrom(ptr unsafe.Pointer) FetchResult {
	return FetchResult{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FetchResultClass) Alloc() FetchResult {
	rv := objc.Call[FetchResult](fc, objc.Sel("alloc"))
	return rv
}

func FetchResult_Alloc() FetchResult {
	return FetchResultClass.Alloc()
}

func (fc _FetchResultClass) New() FetchResult {
	rv := objc.Call[FetchResult](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFetchResult() FetchResult {
	return FetchResultClass.New()
}

func (f_ FetchResult) Init() FetchResult {
	rv := objc.Call[FetchResult](f_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnfetchresult/3294194-value?language=objc
func (f_ FetchResult) Value() objc.Object {
	rv := objc.Call[objc.Object](f_, objc.Sel("value"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnfetchresult/3294193-currenthistorytoken?language=objc
func (f_ FetchResult) CurrentHistoryToken() []byte {
	rv := objc.Call[[]byte](f_, objc.Sel("currentHistoryToken"))
	return rv
}
