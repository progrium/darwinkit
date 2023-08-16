// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AsynchronousFetchResult] class.
var AsynchronousFetchResultClass = _AsynchronousFetchResultClass{objc.GetClass("NSAsynchronousFetchResult")}

type _AsynchronousFetchResultClass struct {
	objc.Class
}

// An interface definition for the [AsynchronousFetchResult] class.
type IAsynchronousFetchResult interface {
	IPersistentStoreAsynchronousResult
	FetchRequest() AsynchronousFetchRequest
	FinalResult() []objc.Object
}

// A fetch result object that encompasses the response from an executed asynchronous fetch request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsasynchronousfetchresult?language=objc
type AsynchronousFetchResult struct {
	PersistentStoreAsynchronousResult
}

func AsynchronousFetchResultFrom(ptr unsafe.Pointer) AsynchronousFetchResult {
	return AsynchronousFetchResult{
		PersistentStoreAsynchronousResult: PersistentStoreAsynchronousResultFrom(ptr),
	}
}

func (ac _AsynchronousFetchResultClass) Alloc() AsynchronousFetchResult {
	rv := objc.Call[AsynchronousFetchResult](ac, objc.Sel("alloc"))
	return rv
}

func AsynchronousFetchResult_Alloc() AsynchronousFetchResult {
	return AsynchronousFetchResultClass.Alloc()
}

func (ac _AsynchronousFetchResultClass) New() AsynchronousFetchResult {
	rv := objc.Call[AsynchronousFetchResult](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAsynchronousFetchResult() AsynchronousFetchResult {
	return AsynchronousFetchResultClass.New()
}

func (a_ AsynchronousFetchResult) Init() AsynchronousFetchResult {
	rv := objc.Call[AsynchronousFetchResult](a_, objc.Sel("init"))
	return rv
}

// The underlying fetch request that was executed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsasynchronousfetchresult/1404906-fetchrequest?language=objc
func (a_ AsynchronousFetchResult) FetchRequest() AsynchronousFetchRequest {
	rv := objc.Call[AsynchronousFetchRequest](a_, objc.Sel("fetchRequest"))
	return rv
}

// The results that were received from the fetch request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsasynchronousfetchresult/1404930-finalresult?language=objc
func (a_ AsynchronousFetchResult) FinalResult() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("finalResult"))
	return rv
}
