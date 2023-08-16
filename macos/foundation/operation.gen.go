// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Operation] class.
var OperationClass = _OperationClass{objc.GetClass("NSOperation")}

type _OperationClass struct {
	objc.Class
}

// An interface definition for the [Operation] class.
type IOperation interface {
	objc.IObject
	RemoveDependency(op IOperation)
	AddDependency(op IOperation)
	Main()
	Start()
	WaitUntilFinished()
	Cancel()
	IsCancelled() bool
	IsExecuting() bool
	Dependencies() []Operation
	Name() string
	SetName(value string)
	QualityOfService() QualityOfService
	SetQualityOfService(value QualityOfService)
	IsConcurrent() bool
	QueuePriority() OperationQueuePriority
	SetQueuePriority(value OperationQueuePriority)
	CompletionBlock() func()
	SetCompletionBlock(value func())
	IsAsynchronous() bool
	IsFinished() bool
	IsReady() bool
}

// An abstract class that represents the code and data associated with a single task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperation?language=objc
type Operation struct {
	objc.Object
}

func OperationFrom(ptr unsafe.Pointer) Operation {
	return Operation{
		Object: objc.ObjectFrom(ptr),
	}
}

func (oc _OperationClass) Alloc() Operation {
	rv := objc.Call[Operation](oc, objc.Sel("alloc"))
	return rv
}

func Operation_Alloc() Operation {
	return OperationClass.Alloc()
}

func (oc _OperationClass) New() Operation {
	rv := objc.Call[Operation](oc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewOperation() Operation {
	return OperationClass.New()
}

func (o_ Operation) Init() Operation {
	rv := objc.Call[Operation](o_, objc.Sel("init"))
	return rv
}

// Removes the receiver’s dependence on the specified operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperation/1414945-removedependency?language=objc
func (o_ Operation) RemoveDependency(op IOperation) {
	objc.Call[objc.Void](o_, objc.Sel("removeDependency:"), objc.Ptr(op))
}

// Makes the receiver dependent on the completion of the specified operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperation/1412859-adddependency?language=objc
func (o_ Operation) AddDependency(op IOperation) {
	objc.Call[objc.Void](o_, objc.Sel("addDependency:"), objc.Ptr(op))
}

// Performs the receiver’s non-concurrent task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperation/1407732-main?language=objc
func (o_ Operation) Main() {
	objc.Call[objc.Void](o_, objc.Sel("main"))
}

// Begins the execution of the operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperation/1416837-start?language=objc
func (o_ Operation) Start() {
	objc.Call[objc.Void](o_, objc.Sel("start"))
}

// Blocks execution of the current thread until the operation object finishes its task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperation/1409256-waituntilfinished?language=objc
func (o_ Operation) WaitUntilFinished() {
	objc.Call[objc.Void](o_, objc.Sel("waitUntilFinished"))
}

// Advises the operation object that it should stop executing its task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperation/1411672-cancel?language=objc
func (o_ Operation) Cancel() {
	objc.Call[objc.Void](o_, objc.Sel("cancel"))
}

// A Boolean value indicating whether the operation has been cancelled [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperation/1408418-cancelled?language=objc
func (o_ Operation) IsCancelled() bool {
	rv := objc.Call[bool](o_, objc.Sel("isCancelled"))
	return rv
}

// A Boolean value indicating whether the operation is currently executing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperation/1415621-executing?language=objc
func (o_ Operation) IsExecuting() bool {
	rv := objc.Call[bool](o_, objc.Sel("isExecuting"))
	return rv
}

// An array of the operation objects that must finish executing before the current object can begin executing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperation/1416668-dependencies?language=objc
func (o_ Operation) Dependencies() []Operation {
	rv := objc.Call[[]Operation](o_, objc.Sel("dependencies"))
	return rv
}

// The name of the operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperation/1416089-name?language=objc
func (o_ Operation) Name() string {
	rv := objc.Call[string](o_, objc.Sel("name"))
	return rv
}

// The name of the operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperation/1416089-name?language=objc
func (o_ Operation) SetName(value string) {
	objc.Call[objc.Void](o_, objc.Sel("setName:"), value)
}

// The relative amount of importance for granting system resources to the operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperation/1413553-qualityofservice?language=objc
func (o_ Operation) QualityOfService() QualityOfService {
	rv := objc.Call[QualityOfService](o_, objc.Sel("qualityOfService"))
	return rv
}

// The relative amount of importance for granting system resources to the operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperation/1413553-qualityofservice?language=objc
func (o_ Operation) SetQualityOfService(value QualityOfService) {
	objc.Call[objc.Void](o_, objc.Sel("setQualityOfService:"), value)
}

// A Boolean value indicating whether the operation executes its task asynchronously. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperation/1411089-concurrent?language=objc
func (o_ Operation) IsConcurrent() bool {
	rv := objc.Call[bool](o_, objc.Sel("isConcurrent"))
	return rv
}

// The execution priority of the operation in an operation queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperation/1411204-queuepriority?language=objc
func (o_ Operation) QueuePriority() OperationQueuePriority {
	rv := objc.Call[OperationQueuePriority](o_, objc.Sel("queuePriority"))
	return rv
}

// The execution priority of the operation in an operation queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperation/1411204-queuepriority?language=objc
func (o_ Operation) SetQueuePriority(value OperationQueuePriority) {
	objc.Call[objc.Void](o_, objc.Sel("setQueuePriority:"), value)
}

// The block to execute after the operation’s main task is completed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperation/1408085-completionblock?language=objc
func (o_ Operation) CompletionBlock() func() {
	rv := objc.Call[func()](o_, objc.Sel("completionBlock"))
	return rv
}

// The block to execute after the operation’s main task is completed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperation/1408085-completionblock?language=objc
func (o_ Operation) SetCompletionBlock(value func()) {
	objc.Call[objc.Void](o_, objc.Sel("setCompletionBlock:"), value)
}

// A Boolean value indicating whether the operation executes its task asynchronously. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperation/1408275-asynchronous?language=objc
func (o_ Operation) IsAsynchronous() bool {
	rv := objc.Call[bool](o_, objc.Sel("isAsynchronous"))
	return rv
}

// A Boolean value indicating whether the operation has finished executing its task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperation/1413540-finished?language=objc
func (o_ Operation) IsFinished() bool {
	rv := objc.Call[bool](o_, objc.Sel("isFinished"))
	return rv
}

// A Boolean value indicating whether the operation can be performed now. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsoperation/1412992-ready?language=objc
func (o_ Operation) IsReady() bool {
	rv := objc.Call[bool](o_, objc.Sel("isReady"))
	return rv
}
