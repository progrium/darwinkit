// AUTO-GENERATED CODE, DO NOT MODIFY
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var OperationClass = _OperationClass{objc.GetClass("NSOperation")}

type _OperationClass struct {
	objc.Class
}

type IOperation interface {
	objc.IObject
	Start()
	Main()
	Cancel()
	AddDependency(op IOperation)
	RemoveDependency(op IOperation)
	WaitUntilFinished()
	CompletionBlock() func()
	SetCompletionBlock(value func())
	IsCancelled() bool
	IsExecuting() bool
	IsFinished() bool
	IsConcurrent() bool
	IsAsynchronous() bool
	IsReady() bool
	Name() string
	SetName(value string)
	Dependencies() []Operation
	QualityOfService() QualityOfService
	SetQualityOfService(value QualityOfService)
	QueuePriority() OperationQueuePriority
	SetQueuePriority(value OperationQueuePriority)
}

type Operation struct {
	objc.Object
}

func MakeOperation(ptr unsafe.Pointer) Operation {
	return Operation{
		Object: objc.MakeObject(ptr),
	}
}

func (oc _OperationClass) Alloc() Operation {
	rv := objc.CallMethod[Operation](oc, objc.GetSelector("alloc"))
	return rv
}

func Operation_Alloc() Operation {
	return OperationClass.Alloc()
}

func (oc _OperationClass) New() Operation {
	rv := objc.CallMethod[Operation](oc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewOperation() Operation {
	return OperationClass.New()
}

func Operation_New() Operation {
	return OperationClass.New()
}

func (o_ Operation) Init() Operation {
	rv := objc.CallMethod[Operation](o_, objc.GetSelector("init"))
	return rv
}

func Operation_Init() Operation {
	return OperationClass.Alloc().Init()
}

func (o_ Operation) Start() {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("start"))
}

func (o_ Operation) Main() {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("main"))
}

func (o_ Operation) Cancel() {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("cancel"))
}

func (o_ Operation) AddDependency(op IOperation) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("addDependency:"), objc.ExtractPtr(op))
}

func (o_ Operation) RemoveDependency(op IOperation) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("removeDependency:"), objc.ExtractPtr(op))
}

func (o_ Operation) WaitUntilFinished() {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("waitUntilFinished"))
}

func (o_ Operation) CompletionBlock() func() {
	rv := objc.CallMethod[func()](o_, objc.GetSelector("completionBlock"))
	return rv
}

func (o_ Operation) SetCompletionBlock(value func()) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("setCompletionBlock:"), value)
}

func (o_ Operation) IsCancelled() bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("isCancelled"))
	return rv
}

func (o_ Operation) IsExecuting() bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("isExecuting"))
	return rv
}

func (o_ Operation) IsFinished() bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("isFinished"))
	return rv
}

func (o_ Operation) IsConcurrent() bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("isConcurrent"))
	return rv
}

func (o_ Operation) IsAsynchronous() bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("isAsynchronous"))
	return rv
}

func (o_ Operation) IsReady() bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("isReady"))
	return rv
}

func (o_ Operation) Name() string {
	rv := objc.CallMethod[string](o_, objc.GetSelector("name"))
	return rv
}

func (o_ Operation) SetName(value string) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("setName:"), value)
}

func (o_ Operation) Dependencies() []Operation {
	rv := objc.CallMethod[[]Operation](o_, objc.GetSelector("dependencies"))
	return rv
}

func (o_ Operation) QualityOfService() QualityOfService {
	rv := objc.CallMethod[QualityOfService](o_, objc.GetSelector("qualityOfService"))
	return rv
}

func (o_ Operation) SetQualityOfService(value QualityOfService) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("setQualityOfService:"), value)
}

func (o_ Operation) QueuePriority() OperationQueuePriority {
	rv := objc.CallMethod[OperationQueuePriority](o_, objc.GetSelector("queuePriority"))
	return rv
}

func (o_ Operation) SetQueuePriority(value OperationQueuePriority) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("setQueuePriority:"), value)
}
