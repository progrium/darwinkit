// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PersistentHistoryResult] class.
var PersistentHistoryResultClass = _PersistentHistoryResultClass{objc.GetClass("NSPersistentHistoryResult")}

type _PersistentHistoryResultClass struct {
	objc.Class
}

// An interface definition for the [PersistentHistoryResult] class.
type IPersistentHistoryResult interface {
	IPersistentStoreResult
	Result() objc.Object
	ResultType() PersistentHistoryResultType
}

// The result of a request to fetch persistent history. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistoryresult?language=objc
type PersistentHistoryResult struct {
	PersistentStoreResult
}

func PersistentHistoryResultFrom(ptr unsafe.Pointer) PersistentHistoryResult {
	return PersistentHistoryResult{
		PersistentStoreResult: PersistentStoreResultFrom(ptr),
	}
}

func (pc _PersistentHistoryResultClass) Alloc() PersistentHistoryResult {
	rv := objc.Call[PersistentHistoryResult](pc, objc.Sel("alloc"))
	return rv
}

func PersistentHistoryResult_Alloc() PersistentHistoryResult {
	return PersistentHistoryResultClass.Alloc()
}

func (pc _PersistentHistoryResultClass) New() PersistentHistoryResult {
	rv := objc.Call[PersistentHistoryResult](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPersistentHistoryResult() PersistentHistoryResult {
	return PersistentHistoryResultClass.New()
}

func (p_ PersistentHistoryResult) Init() PersistentHistoryResult {
	rv := objc.Call[PersistentHistoryResult](p_, objc.Sel("init"))
	return rv
}

// The result of the history request determined by the persistent history result type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistoryresult/2892350-result?language=objc
func (p_ PersistentHistoryResult) Result() objc.Object {
	rv := objc.Call[objc.Object](p_, objc.Sel("result"))
	return rv
}

// The type of result that the persistent history change request returns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistoryresult/2892356-resulttype?language=objc
func (p_ PersistentHistoryResult) ResultType() PersistentHistoryResultType {
	rv := objc.Call[PersistentHistoryResultType](p_, objc.Sel("resultType"))
	return rv
}
