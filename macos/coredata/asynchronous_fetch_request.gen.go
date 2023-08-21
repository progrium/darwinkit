// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AsynchronousFetchRequest] class.
var AsynchronousFetchRequestClass = _AsynchronousFetchRequestClass{objc.GetClass("NSAsynchronousFetchRequest")}

type _AsynchronousFetchRequestClass struct {
	objc.Class
}

// An interface definition for the [AsynchronousFetchRequest] class.
type IAsynchronousFetchRequest interface {
	IPersistentStoreRequest
	FetchRequest() FetchRequest
	EstimatedResultCount() int
	SetEstimatedResultCount(value int)
	CompletionBlock() PersistentStoreAsynchronousFetchResultCompletionBlock
}

// A fetch request that retrieves results asynchronously and supports progress notification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsasynchronousfetchrequest?language=objc
type AsynchronousFetchRequest struct {
	PersistentStoreRequest
}

func AsynchronousFetchRequestFrom(ptr unsafe.Pointer) AsynchronousFetchRequest {
	return AsynchronousFetchRequest{
		PersistentStoreRequest: PersistentStoreRequestFrom(ptr),
	}
}

func (a_ AsynchronousFetchRequest) InitWithFetchRequestCompletionBlock(request IFetchRequest, blk func(arg0 AsynchronousFetchResult)) AsynchronousFetchRequest {
	rv := objc.Call[AsynchronousFetchRequest](a_, objc.Sel("initWithFetchRequest:completionBlock:"), objc.Ptr(request), blk)
	return rv
}

// Initializes a new asynchronous fetch request configured with the provided fetch request and completion block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsasynchronousfetchrequest/1506218-initwithfetchrequest?language=objc
func NewAsynchronousFetchRequestWithFetchRequestCompletionBlock(request IFetchRequest, blk func(arg0 AsynchronousFetchResult)) AsynchronousFetchRequest {
	instance := AsynchronousFetchRequestClass.Alloc().InitWithFetchRequestCompletionBlock(request, blk)
	instance.Autorelease()
	return instance
}

func (ac _AsynchronousFetchRequestClass) Alloc() AsynchronousFetchRequest {
	rv := objc.Call[AsynchronousFetchRequest](ac, objc.Sel("alloc"))
	return rv
}

func AsynchronousFetchRequest_Alloc() AsynchronousFetchRequest {
	return AsynchronousFetchRequestClass.Alloc()
}

func (ac _AsynchronousFetchRequestClass) New() AsynchronousFetchRequest {
	rv := objc.Call[AsynchronousFetchRequest](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAsynchronousFetchRequest() AsynchronousFetchRequest {
	return AsynchronousFetchRequestClass.New()
}

func (a_ AsynchronousFetchRequest) Init() AsynchronousFetchRequest {
	rv := objc.Call[AsynchronousFetchRequest](a_, objc.Sel("init"))
	return rv
}

// The underlying fetch request that is executed asynchronously. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsasynchronousfetchrequest/1506719-fetchrequest?language=objc
func (a_ AsynchronousFetchRequest) FetchRequest() FetchRequest {
	rv := objc.Call[FetchRequest](a_, objc.Sel("fetchRequest"))
	return rv
}

// A configuration parameter that assists Core Data with scheduling the asynchronous fetch request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsasynchronousfetchrequest/1506474-estimatedresultcount?language=objc
func (a_ AsynchronousFetchRequest) EstimatedResultCount() int {
	rv := objc.Call[int](a_, objc.Sel("estimatedResultCount"))
	return rv
}

// A configuration parameter that assists Core Data with scheduling the asynchronous fetch request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsasynchronousfetchrequest/1506474-estimatedresultcount?language=objc
func (a_ AsynchronousFetchRequest) SetEstimatedResultCount(value int) {
	objc.Call[objc.Void](a_, objc.Sel("setEstimatedResultCount:"), value)
}

// The block that is executed when the fetch request has completed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsasynchronousfetchrequest/1506815-completionblock?language=objc
func (a_ AsynchronousFetchRequest) CompletionBlock() PersistentStoreAsynchronousFetchResultCompletionBlock {
	rv := objc.Call[PersistentStoreAsynchronousFetchResultCompletionBlock](a_, objc.Sel("completionBlock"))
	return rv
}
