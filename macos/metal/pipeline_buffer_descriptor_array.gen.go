// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PipelineBufferDescriptorArray] class.
var PipelineBufferDescriptorArrayClass = _PipelineBufferDescriptorArrayClass{objc.GetClass("MTLPipelineBufferDescriptorArray")}

type _PipelineBufferDescriptorArrayClass struct {
	objc.Class
}

// An interface definition for the [PipelineBufferDescriptorArray] class.
type IPipelineBufferDescriptorArray interface {
	objc.IObject
	ObjectAtIndexedSubscript(bufferIndex uint) PipelineBufferDescriptor
	SetObjectAtIndexedSubscript(buffer IPipelineBufferDescriptor, bufferIndex uint)
}

// An array of pipeline buffer descriptor objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlpipelinebufferdescriptorarray?language=objc
type PipelineBufferDescriptorArray struct {
	objc.Object
}

func PipelineBufferDescriptorArrayFrom(ptr unsafe.Pointer) PipelineBufferDescriptorArray {
	return PipelineBufferDescriptorArray{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PipelineBufferDescriptorArrayClass) Alloc() PipelineBufferDescriptorArray {
	rv := objc.Call[PipelineBufferDescriptorArray](pc, objc.Sel("alloc"))
	return rv
}

func PipelineBufferDescriptorArray_Alloc() PipelineBufferDescriptorArray {
	return PipelineBufferDescriptorArrayClass.Alloc()
}

func (pc _PipelineBufferDescriptorArrayClass) New() PipelineBufferDescriptorArray {
	rv := objc.Call[PipelineBufferDescriptorArray](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPipelineBufferDescriptorArray() PipelineBufferDescriptorArray {
	return PipelineBufferDescriptorArrayClass.New()
}

func (p_ PipelineBufferDescriptorArray) Init() PipelineBufferDescriptorArray {
	rv := objc.Call[PipelineBufferDescriptorArray](p_, objc.Sel("init"))
	return rv
}

// Returns the pipeline buffer descriptor at the specified array index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlpipelinebufferdescriptorarray/2879272-objectatindexedsubscript?language=objc
func (p_ PipelineBufferDescriptorArray) ObjectAtIndexedSubscript(bufferIndex uint) PipelineBufferDescriptor {
	rv := objc.Call[PipelineBufferDescriptor](p_, objc.Sel("objectAtIndexedSubscript:"), bufferIndex)
	return rv
}

// Sets a pipeline buffer descriptor at the specified array index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlpipelinebufferdescriptorarray/2879310-setobject?language=objc
func (p_ PipelineBufferDescriptorArray) SetObjectAtIndexedSubscript(buffer IPipelineBufferDescriptor, bufferIndex uint) {
	objc.Call[objc.Void](p_, objc.Sel("setObject:atIndexedSubscript:"), objc.Ptr(buffer), bufferIndex)
}
