// AUTO-GENERATED CODE, DO NOT MODIFY

package mpsgraph

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ExecutableExecutionDescriptor] class.
var ExecutableExecutionDescriptorClass = _ExecutableExecutionDescriptorClass{objc.GetClass("MPSGraphExecutableExecutionDescriptor")}

type _ExecutableExecutionDescriptorClass struct {
	objc.Class
}

// An interface definition for the [ExecutableExecutionDescriptor] class.
type IExecutableExecutionDescriptor interface {
	objc.IObject
	ScheduledHandler() ExecutableScheduledHandler
	SetScheduledHandler(value ExecutableScheduledHandler)
	WaitUntilCompleted() bool
	SetWaitUntilCompleted(value bool)
	CompletionHandler() ExecutableCompletionHandler
	SetCompletionHandler(value ExecutableCompletionHandler)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutableexecutiondescriptor?language=objc
type ExecutableExecutionDescriptor struct {
	objc.Object
}

func ExecutableExecutionDescriptorFrom(ptr unsafe.Pointer) ExecutableExecutionDescriptor {
	return ExecutableExecutionDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ec _ExecutableExecutionDescriptorClass) Alloc() ExecutableExecutionDescriptor {
	rv := objc.Call[ExecutableExecutionDescriptor](ec, objc.Sel("alloc"))
	return rv
}

func ExecutableExecutionDescriptor_Alloc() ExecutableExecutionDescriptor {
	return ExecutableExecutionDescriptorClass.Alloc()
}

func (ec _ExecutableExecutionDescriptorClass) New() ExecutableExecutionDescriptor {
	rv := objc.Call[ExecutableExecutionDescriptor](ec, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewExecutableExecutionDescriptor() ExecutableExecutionDescriptor {
	return ExecutableExecutionDescriptorClass.New()
}

func (e_ ExecutableExecutionDescriptor) Init() ExecutableExecutionDescriptor {
	rv := objc.Call[ExecutableExecutionDescriptor](e_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutableexecutiondescriptor/3787601-scheduledhandler?language=objc
func (e_ ExecutableExecutionDescriptor) ScheduledHandler() ExecutableScheduledHandler {
	rv := objc.Call[ExecutableScheduledHandler](e_, objc.Sel("scheduledHandler"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutableexecutiondescriptor/3787601-scheduledhandler?language=objc
func (e_ ExecutableExecutionDescriptor) SetScheduledHandler(value ExecutableScheduledHandler) {
	objc.Call[objc.Void](e_, objc.Sel("setScheduledHandler:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutableexecutiondescriptor/3787602-waituntilcompleted?language=objc
func (e_ ExecutableExecutionDescriptor) WaitUntilCompleted() bool {
	rv := objc.Call[bool](e_, objc.Sel("waitUntilCompleted"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutableexecutiondescriptor/3787602-waituntilcompleted?language=objc
func (e_ ExecutableExecutionDescriptor) SetWaitUntilCompleted(value bool) {
	objc.Call[objc.Void](e_, objc.Sel("setWaitUntilCompleted:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutableexecutiondescriptor/3787600-completionhandler?language=objc
func (e_ ExecutableExecutionDescriptor) CompletionHandler() ExecutableCompletionHandler {
	rv := objc.Call[ExecutableCompletionHandler](e_, objc.Sel("completionHandler"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutableexecutiondescriptor/3787600-completionhandler?language=objc
func (e_ ExecutableExecutionDescriptor) SetCompletionHandler(value ExecutableCompletionHandler) {
	objc.Call[objc.Void](e_, objc.Sel("setCompletionHandler:"), value)
}