// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var OperationQueueClass = _OperationQueueClass{objc.GetClass("NSOperationQueue")}

type _OperationQueueClass struct {
	objc.Class
}

type IOperationQueue interface {
	objc.IObject
	AddOperation(op IOperation)
	AddOperationsWaitUntilFinished(ops []IOperation, wait bool)
	AddOperationWithBlock(block func())
	AddBarrierBlock(barrier func())
	CancelAllOperations()
	WaitUntilAllOperationsAreFinished()
	QualityOfService() QualityOfService
	SetQualityOfService(value QualityOfService)
	MaxConcurrentOperationCount() int
	SetMaxConcurrentOperationCount(value int)
	Progress() Progress
	IsSuspended() bool
	SetSuspended(value bool)
	Name() string
	SetName(value string)
}

type OperationQueue struct {
	objc.Object
}

func MakeOperationQueue(ptr unsafe.Pointer) OperationQueue {
	return OperationQueue{
		Object: objc.MakeObject(ptr),
	}
}

func (oc _OperationQueueClass) Alloc() OperationQueue {
	rv := objc.CallMethod[OperationQueue](oc, objc.GetSelector("alloc"))
	return rv
}

func OperationQueue_Alloc() OperationQueue {
	return OperationQueueClass.Alloc()
}

func (oc _OperationQueueClass) New() OperationQueue {
	rv := objc.CallMethod[OperationQueue](oc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewOperationQueue() OperationQueue {
	return OperationQueueClass.New()
}

func OperationQueue_New() OperationQueue {
	return OperationQueueClass.New()
}

func (o_ OperationQueue) Init() OperationQueue {
	rv := objc.CallMethod[OperationQueue](o_, objc.GetSelector("init"))
	return rv
}

func OperationQueue_Init() OperationQueue {
	return OperationQueueClass.Alloc().Init()
}

func (o_ OperationQueue) AddOperation(op IOperation) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("addOperation:"), objc.ExtractPtr(op))
}

func (o_ OperationQueue) AddOperationsWaitUntilFinished(ops []IOperation, wait bool) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("addOperations:waitUntilFinished:"), ops, wait)
}

func (o_ OperationQueue) AddOperationWithBlock(block func()) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("addOperationWithBlock:"), block)
}

func (o_ OperationQueue) AddBarrierBlock(barrier func()) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("addBarrierBlock:"), barrier)
}

func (o_ OperationQueue) CancelAllOperations() {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("cancelAllOperations"))
}

func (o_ OperationQueue) WaitUntilAllOperationsAreFinished() {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("waitUntilAllOperationsAreFinished"))
}

func (oc _OperationQueueClass) MainQueue() OperationQueue {
	rv := objc.CallMethod[OperationQueue](oc, objc.GetSelector("mainQueue"))
	return rv
}

func OperationQueue_MainQueue() OperationQueue {
	return OperationQueueClass.MainQueue()
}

func (oc _OperationQueueClass) CurrentQueue() OperationQueue {
	rv := objc.CallMethod[OperationQueue](oc, objc.GetSelector("currentQueue"))
	return rv
}

func OperationQueue_CurrentQueue() OperationQueue {
	return OperationQueueClass.CurrentQueue()
}

func (o_ OperationQueue) QualityOfService() QualityOfService {
	rv := objc.CallMethod[QualityOfService](o_, objc.GetSelector("qualityOfService"))
	return rv
}

func (o_ OperationQueue) SetQualityOfService(value QualityOfService) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("setQualityOfService:"), value)
}

func (o_ OperationQueue) MaxConcurrentOperationCount() int {
	rv := objc.CallMethod[int](o_, objc.GetSelector("maxConcurrentOperationCount"))
	return rv
}

func (o_ OperationQueue) SetMaxConcurrentOperationCount(value int) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("setMaxConcurrentOperationCount:"), value)
}

func (o_ OperationQueue) Progress() Progress {
	rv := objc.CallMethod[Progress](o_, objc.GetSelector("progress"))
	return rv
}

func (o_ OperationQueue) IsSuspended() bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("isSuspended"))
	return rv
}

func (o_ OperationQueue) SetSuspended(value bool) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("setSuspended:"), value)
}

func (o_ OperationQueue) Name() string {
	rv := objc.CallMethod[string](o_, objc.GetSelector("name"))
	return rv
}

func (o_ OperationQueue) SetName(value string) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("setName:"), value)
}
