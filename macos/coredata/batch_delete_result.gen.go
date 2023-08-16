// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [BatchDeleteResult] class.
var BatchDeleteResultClass = _BatchDeleteResultClass{objc.GetClass("NSBatchDeleteResult")}

type _BatchDeleteResultClass struct {
	objc.Class
}

// An interface definition for the [BatchDeleteResult] class.
type IBatchDeleteResult interface {
	IPersistentStoreResult
	Result() objc.Object
	ResultType() BatchDeleteRequestResultType
}

// An object that describes the result of a batch delete request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchdeleteresult?language=objc
type BatchDeleteResult struct {
	PersistentStoreResult
}

func BatchDeleteResultFrom(ptr unsafe.Pointer) BatchDeleteResult {
	return BatchDeleteResult{
		PersistentStoreResult: PersistentStoreResultFrom(ptr),
	}
}

func (bc _BatchDeleteResultClass) Alloc() BatchDeleteResult {
	rv := objc.Call[BatchDeleteResult](bc, objc.Sel("alloc"))
	return rv
}

func BatchDeleteResult_Alloc() BatchDeleteResult {
	return BatchDeleteResultClass.Alloc()
}

func (bc _BatchDeleteResultClass) New() BatchDeleteResult {
	rv := objc.Call[BatchDeleteResult](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewBatchDeleteResult() BatchDeleteResult {
	return BatchDeleteResultClass.New()
}

func (b_ BatchDeleteResult) Init() BatchDeleteResult {
	rv := objc.Call[BatchDeleteResult](b_, objc.Sel("init"))
	return rv
}

// The value the request returns after it executes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchdeleteresult/1404922-result?language=objc
func (b_ BatchDeleteResult) Result() objc.Object {
	rv := objc.Call[objc.Object](b_, objc.Sel("result"))
	return rv
}

// The data type of the request’s result value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchdeleteresult/1404941-resulttype?language=objc
func (b_ BatchDeleteResult) ResultType() BatchDeleteRequestResultType {
	rv := objc.Call[BatchDeleteRequestResultType](b_, objc.Sel("resultType"))
	return rv
}
