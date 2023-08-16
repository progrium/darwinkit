// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/dispatch"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [OperationQueue] class.
var OperationQueueClass = _OperationQueueClass{objc.GetClass("NSOperationQueue")}

type _OperationQueueClass struct {
	objc.Class
}

// An interface definition for the [OperationQueue] class.
type IOperationQueue interface {
	objc.IObject
	AddOperation(op IOperation)
	AddBarrierBlock(barrier func())
	WaitUntilAllOperationsAreFinished()
	AddOperationsWaitUntilFinished(ops []IOperation, wait bool)
	AddOperationWithBlock(block func())
	CancelAllOperations()
	Name() string
	SetName(value string)
	QualityOfService() QualityOfService
	SetQualityOfService(value QualityOfService)
	IsSuspended() bool
	SetSuspended(value bool)
	UnderlyingQueue() dispatch.Queue
	SetUnderlyingQueue(value dispatch.Queue)
	Progress() Progress
	MaxConcurrentOperationCount() int
	SetMaxConcurrentOperationCount(value int)
}

// A queue that regulates the execution of operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperationqueue?language=objc
type OperationQueue struct {
	objc.Object
}

func OperationQueueFrom(ptr unsafe.Pointer) OperationQueue {
	return OperationQueue{
		Object: objc.ObjectFrom(ptr),
	}
}

func (oc _OperationQueueClass) Alloc() OperationQueue {
	rv := objc.Call[OperationQueue](oc, objc.Sel("alloc"))
	return rv
}

func OperationQueue_Alloc() OperationQueue {
	return OperationQueueClass.Alloc()
}

func (oc _OperationQueueClass) New() OperationQueue {
	rv := objc.Call[OperationQueue](oc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewOperationQueue() OperationQueue {
	return OperationQueueClass.New()
}

func (o_ OperationQueue) Init() OperationQueue {
	rv := objc.Call[OperationQueue](o_, objc.Sel("init"))
	return rv
}

// Adds the specified operation to the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperationqueue/1410704-addoperation?language=objc
func (o_ OperationQueue) AddOperation(op IOperation) {
	objc.Call[objc.Void](o_, objc.Sel("addOperation:"), objc.Ptr(op))
}

// Invokes a block when the queue finishes all enqueued operations, and prevents subsequent operations from starting until the block has completed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperationqueue/3172534-addbarrierblock?language=objc
func (o_ OperationQueue) AddBarrierBlock(barrier func()) {
	objc.Call[objc.Void](o_, objc.Sel("addBarrierBlock:"), barrier)
}

// Blocks the current thread until all the receiver’s queued and executing operations finish executing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperationqueue/1407971-waituntilalloperationsarefinishe?language=objc
func (o_ OperationQueue) WaitUntilAllOperationsAreFinished() {
	objc.Call[objc.Void](o_, objc.Sel("waitUntilAllOperationsAreFinished"))
}

// Adds the specified operations to the queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperationqueue/1408358-addoperations?language=objc
func (o_ OperationQueue) AddOperationsWaitUntilFinished(ops []IOperation, wait bool) {
	objc.Call[objc.Void](o_, objc.Sel("addOperations:waitUntilFinished:"), ops, wait)
}

// Wraps the specified block in an operation and adds it to the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperationqueue/1412949-addoperationwithblock?language=objc
func (o_ OperationQueue) AddOperationWithBlock(block func()) {
	objc.Call[objc.Void](o_, objc.Sel("addOperationWithBlock:"), block)
}

// Cancels all queued and executing operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperationqueue/1417849-cancelalloperations?language=objc
func (o_ OperationQueue) CancelAllOperations() {
	objc.Call[objc.Void](o_, objc.Sel("cancelAllOperations"))
}

// The name of the operation queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperationqueue/1418063-name?language=objc
func (o_ OperationQueue) Name() string {
	rv := objc.Call[string](o_, objc.Sel("name"))
	return rv
}

// The name of the operation queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperationqueue/1418063-name?language=objc
func (o_ OperationQueue) SetName(value string) {
	objc.Call[objc.Void](o_, objc.Sel("setName:"), value)
}

// The default service level to apply to operations that the queue invokes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperationqueue/1417919-qualityofservice?language=objc
func (o_ OperationQueue) QualityOfService() QualityOfService {
	rv := objc.Call[QualityOfService](o_, objc.Sel("qualityOfService"))
	return rv
}

// The default service level to apply to operations that the queue invokes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperationqueue/1417919-qualityofservice?language=objc
func (o_ OperationQueue) SetQualityOfService(value QualityOfService) {
	objc.Call[objc.Void](o_, objc.Sel("setQualityOfService:"), value)
}

// Returns the operation queue associated with the main thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperationqueue/1409193-mainqueue?language=objc
func (oc _OperationQueueClass) MainQueue() OperationQueue {
	rv := objc.Call[OperationQueue](oc, objc.Sel("mainQueue"))
	return rv
}

// Returns the operation queue associated with the main thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperationqueue/1409193-mainqueue?language=objc
func OperationQueue_MainQueue() OperationQueue {
	return OperationQueueClass.MainQueue()
}

// A Boolean value indicating whether the queue is actively scheduling operations for execution. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperationqueue/1415909-suspended?language=objc
func (o_ OperationQueue) IsSuspended() bool {
	rv := objc.Call[bool](o_, objc.Sel("isSuspended"))
	return rv
}

// A Boolean value indicating whether the queue is actively scheduling operations for execution. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperationqueue/1415909-suspended?language=objc
func (o_ OperationQueue) SetSuspended(value bool) {
	objc.Call[objc.Void](o_, objc.Sel("setSuspended:"), value)
}

// The dispatch queue that the operation queue uses to invoke operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperationqueue/1415344-underlyingqueue?language=objc
func (o_ OperationQueue) UnderlyingQueue() dispatch.Queue {
	rv := objc.Call[dispatch.Queue](o_, objc.Sel("underlyingQueue"))
	return rv
}

// The dispatch queue that the operation queue uses to invoke operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperationqueue/1415344-underlyingqueue?language=objc
func (o_ OperationQueue) SetUnderlyingQueue(value dispatch.Queue) {
	objc.Call[objc.Void](o_, objc.Sel("setUnderlyingQueue:"), value)
}

// An object that represents the total progress of the operations executing in the queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperationqueue/3172535-progress?language=objc
func (o_ OperationQueue) Progress() Progress {
	rv := objc.Call[Progress](o_, objc.Sel("progress"))
	return rv
}

// The maximum number of queued operations that can run at the same time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperationqueue/1414982-maxconcurrentoperationcount?language=objc
func (o_ OperationQueue) MaxConcurrentOperationCount() int {
	rv := objc.Call[int](o_, objc.Sel("maxConcurrentOperationCount"))
	return rv
}

// The maximum number of queued operations that can run at the same time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperationqueue/1414982-maxconcurrentoperationcount?language=objc
func (o_ OperationQueue) SetMaxConcurrentOperationCount(value int) {
	objc.Call[objc.Void](o_, objc.Sel("setMaxConcurrentOperationCount:"), value)
}

// Returns the operation queue that launched the current operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperationqueue/1413097-currentqueue?language=objc
func (oc _OperationQueueClass) CurrentQueue() OperationQueue {
	rv := objc.Call[OperationQueue](oc, objc.Sel("currentQueue"))
	return rv
}

// Returns the operation queue that launched the current operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperationqueue/1413097-currentqueue?language=objc
func OperationQueue_CurrentQueue() OperationQueue {
	return OperationQueueClass.CurrentQueue()
}
