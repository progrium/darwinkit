// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CommandBufferDescriptor] class.
var CommandBufferDescriptorClass = _CommandBufferDescriptorClass{objc.GetClass("MTLCommandBufferDescriptor")}

type _CommandBufferDescriptorClass struct {
	objc.Class
}

// An interface definition for the [CommandBufferDescriptor] class.
type ICommandBufferDescriptor interface {
	objc.IObject
	RetainedReferences() bool
	SetRetainedReferences(value bool)
	ErrorOptions() CommandBufferErrorOption
	SetErrorOptions(value CommandBufferErrorOption)
}

// A configuration that customizes the behavior for a new command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbufferdescriptor?language=objc
type CommandBufferDescriptor struct {
	objc.Object
}

func CommandBufferDescriptorFrom(ptr unsafe.Pointer) CommandBufferDescriptor {
	return CommandBufferDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CommandBufferDescriptorClass) Alloc() CommandBufferDescriptor {
	rv := objc.Call[CommandBufferDescriptor](cc, objc.Sel("alloc"))
	return rv
}

func CommandBufferDescriptor_Alloc() CommandBufferDescriptor {
	return CommandBufferDescriptorClass.Alloc()
}

func (cc _CommandBufferDescriptorClass) New() CommandBufferDescriptor {
	rv := objc.Call[CommandBufferDescriptor](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCommandBufferDescriptor() CommandBufferDescriptor {
	return CommandBufferDescriptorClass.New()
}

func (c_ CommandBufferDescriptor) Init() CommandBufferDescriptor {
	rv := objc.Call[CommandBufferDescriptor](c_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether the command buffer the descriptor creates maintains strong references to the resources it uses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbufferdescriptor/3553942-retainedreferences?language=objc
func (c_ CommandBufferDescriptor) RetainedReferences() bool {
	rv := objc.Call[bool](c_, objc.Sel("retainedReferences"))
	return rv
}

// A Boolean value that indicates whether the command buffer the descriptor creates maintains strong references to the resources it uses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbufferdescriptor/3553942-retainedreferences?language=objc
func (c_ CommandBufferDescriptor) SetRetainedReferences(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setRetainedReferences:"), value)
}

// The reporting configuration that indicates which information the GPU driver stores in a command buffer’s error property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbufferdescriptor/3553941-erroroptions?language=objc
func (c_ CommandBufferDescriptor) ErrorOptions() CommandBufferErrorOption {
	rv := objc.Call[CommandBufferErrorOption](c_, objc.Sel("errorOptions"))
	return rv
}

// The reporting configuration that indicates which information the GPU driver stores in a command buffer’s error property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbufferdescriptor/3553941-erroroptions?language=objc
func (c_ CommandBufferDescriptor) SetErrorOptions(value CommandBufferErrorOption) {
	objc.Call[objc.Void](c_, objc.Sel("setErrorOptions:"), value)
}
