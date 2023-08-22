// AUTO-GENERATED CODE, DO NOT MODIFY

package mpsgraph

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ExecutionDescriptor] class.
var ExecutionDescriptorClass = _ExecutionDescriptorClass{objc.GetClass("MPSGraphExecutionDescriptor")}

type _ExecutionDescriptorClass struct {
	objc.Class
}

// An interface definition for the [ExecutionDescriptor] class.
type IExecutionDescriptor interface {
	objc.IObject
	ScheduledHandler() ScheduledHandler
	SetScheduledHandler(value ScheduledHandler)
	WaitUntilCompleted() bool
	SetWaitUntilCompleted(value bool)
	CompilationDescriptor() CompilationDescriptor
	SetCompilationDescriptor(value ICompilationDescriptor)
	CompletionHandler() CompletionHandler
	SetCompletionHandler(value CompletionHandler)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutiondescriptor?language=objc
type ExecutionDescriptor struct {
	objc.Object
}

func ExecutionDescriptorFrom(ptr unsafe.Pointer) ExecutionDescriptor {
	return ExecutionDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ec _ExecutionDescriptorClass) Alloc() ExecutionDescriptor {
	rv := objc.Call[ExecutionDescriptor](ec, objc.Sel("alloc"))
	return rv
}

func ExecutionDescriptor_Alloc() ExecutionDescriptor {
	return ExecutionDescriptorClass.Alloc()
}

func (ec _ExecutionDescriptorClass) New() ExecutionDescriptor {
	rv := objc.Call[ExecutionDescriptor](ec, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewExecutionDescriptor() ExecutionDescriptor {
	return ExecutionDescriptorClass.New()
}

func (e_ ExecutionDescriptor) Init() ExecutionDescriptor {
	rv := objc.Call[ExecutionDescriptor](e_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutiondescriptor/3564627-scheduledhandler?language=objc
func (e_ ExecutionDescriptor) ScheduledHandler() ScheduledHandler {
	rv := objc.Call[ScheduledHandler](e_, objc.Sel("scheduledHandler"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutiondescriptor/3564627-scheduledhandler?language=objc
func (e_ ExecutionDescriptor) SetScheduledHandler(value ScheduledHandler) {
	objc.Call[objc.Void](e_, objc.Sel("setScheduledHandler:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutiondescriptor/3564628-waituntilcompleted?language=objc
func (e_ ExecutionDescriptor) WaitUntilCompleted() bool {
	rv := objc.Call[bool](e_, objc.Sel("waitUntilCompleted"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutiondescriptor/3564628-waituntilcompleted?language=objc
func (e_ ExecutionDescriptor) SetWaitUntilCompleted(value bool) {
	objc.Call[objc.Void](e_, objc.Sel("setWaitUntilCompleted:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutiondescriptor/3922625-compilationdescriptor?language=objc
func (e_ ExecutionDescriptor) CompilationDescriptor() CompilationDescriptor {
	rv := objc.Call[CompilationDescriptor](e_, objc.Sel("compilationDescriptor"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutiondescriptor/3922625-compilationdescriptor?language=objc
func (e_ ExecutionDescriptor) SetCompilationDescriptor(value ICompilationDescriptor) {
	objc.Call[objc.Void](e_, objc.Sel("setCompilationDescriptor:"), objc.Ptr(value))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutiondescriptor/3564626-completionhandler?language=objc
func (e_ ExecutionDescriptor) CompletionHandler() CompletionHandler {
	rv := objc.Call[CompletionHandler](e_, objc.Sel("completionHandler"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutiondescriptor/3564626-completionhandler?language=objc
func (e_ ExecutionDescriptor) SetCompletionHandler(value CompletionHandler) {
	objc.Call[objc.Void](e_, objc.Sel("setCompletionHandler:"), value)
}
