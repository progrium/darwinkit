// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PipelineBufferDescriptor] class.
var PipelineBufferDescriptorClass = _PipelineBufferDescriptorClass{objc.GetClass("MTLPipelineBufferDescriptor")}

type _PipelineBufferDescriptorClass struct {
	objc.Class
}

// An interface definition for the [PipelineBufferDescriptor] class.
type IPipelineBufferDescriptor interface {
	objc.IObject
	Mutability() Mutability
	SetMutability(value Mutability)
}

// An object that contains the mutability options for a buffer that you pass to a render or compute pipeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlpipelinebufferdescriptor?language=objc
type PipelineBufferDescriptor struct {
	objc.Object
}

func PipelineBufferDescriptorFrom(ptr unsafe.Pointer) PipelineBufferDescriptor {
	return PipelineBufferDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PipelineBufferDescriptorClass) Alloc() PipelineBufferDescriptor {
	rv := objc.Call[PipelineBufferDescriptor](pc, objc.Sel("alloc"))
	return rv
}

func PipelineBufferDescriptor_Alloc() PipelineBufferDescriptor {
	return PipelineBufferDescriptorClass.Alloc()
}

func (pc _PipelineBufferDescriptorClass) New() PipelineBufferDescriptor {
	rv := objc.Call[PipelineBufferDescriptor](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPipelineBufferDescriptor() PipelineBufferDescriptor {
	return PipelineBufferDescriptorClass.New()
}

func (p_ PipelineBufferDescriptor) Init() PipelineBufferDescriptor {
	rv := objc.Call[PipelineBufferDescriptor](p_, objc.Sel("init"))
	return rv
}

// A mutability option that determines whether you can update a buffer’s contents before related commands use the buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlpipelinebufferdescriptor/2879274-mutability?language=objc
func (p_ PipelineBufferDescriptor) Mutability() Mutability {
	rv := objc.Call[Mutability](p_, objc.Sel("mutability"))
	return rv
}

// A mutability option that determines whether you can update a buffer’s contents before related commands use the buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlpipelinebufferdescriptor/2879274-mutability?language=objc
func (p_ PipelineBufferDescriptor) SetMutability(value Mutability) {
	objc.Call[objc.Void](p_, objc.Sel("setMutability:"), value)
}
