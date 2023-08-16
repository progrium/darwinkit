// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [BatchInsertResult] class.
var BatchInsertResultClass = _BatchInsertResultClass{objc.GetClass("NSBatchInsertResult")}

type _BatchInsertResultClass struct {
	objc.Class
}

// An interface definition for the [BatchInsertResult] class.
type IBatchInsertResult interface {
	IPersistentStoreResult
	Result() objc.Object
	ResultType() BatchInsertRequestResultType
}

// The result that Core Data returns when executing a batch-insertion request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchinsertresult?language=objc
type BatchInsertResult struct {
	PersistentStoreResult
}

func BatchInsertResultFrom(ptr unsafe.Pointer) BatchInsertResult {
	return BatchInsertResult{
		PersistentStoreResult: PersistentStoreResultFrom(ptr),
	}
}

func (bc _BatchInsertResultClass) Alloc() BatchInsertResult {
	rv := objc.Call[BatchInsertResult](bc, objc.Sel("alloc"))
	return rv
}

func BatchInsertResult_Alloc() BatchInsertResult {
	return BatchInsertResultClass.Alloc()
}

func (bc _BatchInsertResultClass) New() BatchInsertResult {
	rv := objc.Call[BatchInsertResult](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewBatchInsertResult() BatchInsertResult {
	return BatchInsertResultClass.New()
}

func (b_ BatchInsertResult) Init() BatchInsertResult {
	rv := objc.Call[BatchInsertResult](b_, objc.Sel("init"))
	return rv
}

// The result of a batch-insertion request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchinsertresult/3333246-result?language=objc
func (b_ BatchInsertResult) Result() objc.Object {
	rv := objc.Call[objc.Object](b_, objc.Sel("result"))
	return rv
}

// The type of result that Core Data returns from this request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchinsertresult/3333247-resulttype?language=objc
func (b_ BatchInsertResult) ResultType() BatchInsertRequestResultType {
	rv := objc.Call[BatchInsertRequestResultType](b_, objc.Sel("resultType"))
	return rv
}
