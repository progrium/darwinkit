// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ComputePassDescriptor] class.
var ComputePassDescriptorClass = _ComputePassDescriptorClass{objc.GetClass("MTLComputePassDescriptor")}

type _ComputePassDescriptorClass struct {
	objc.Class
}

// An interface definition for the [ComputePassDescriptor] class.
type IComputePassDescriptor interface {
	objc.IObject
	SampleBufferAttachments() ComputePassSampleBufferAttachmentDescriptorArray
	DispatchType() DispatchType
	SetDispatchType(value DispatchType)
}

// A configuration for a compute pass for creating a compute command encoder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepassdescriptor?language=objc
type ComputePassDescriptor struct {
	objc.Object
}

func ComputePassDescriptorFrom(ptr unsafe.Pointer) ComputePassDescriptor {
	return ComputePassDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _ComputePassDescriptorClass) Alloc() ComputePassDescriptor {
	rv := objc.Call[ComputePassDescriptor](cc, objc.Sel("alloc"))
	return rv
}

func ComputePassDescriptor_Alloc() ComputePassDescriptor {
	return ComputePassDescriptorClass.Alloc()
}

func (cc _ComputePassDescriptorClass) New() ComputePassDescriptor {
	rv := objc.Call[ComputePassDescriptor](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewComputePassDescriptor() ComputePassDescriptor {
	return ComputePassDescriptorClass.New()
}

func (c_ ComputePassDescriptor) Init() ComputePassDescriptor {
	rv := objc.Call[ComputePassDescriptor](c_, objc.Sel("init"))
	return rv
}

// Creates a default compute pass descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepassdescriptor/3564434-computepassdescriptor?language=objc
func (cc _ComputePassDescriptorClass) ComputePassDescriptor() ComputePassDescriptor {
	rv := objc.Call[ComputePassDescriptor](cc, objc.Sel("computePassDescriptor"))
	return rv
}

// Creates a default compute pass descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepassdescriptor/3564434-computepassdescriptor?language=objc
func ComputePassDescriptor_ComputePassDescriptor() ComputePassDescriptor {
	return ComputePassDescriptorClass.ComputePassDescriptor()
}

// The array of sample buffers that the compute pass can access. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepassdescriptor/3564436-samplebufferattachments?language=objc
func (c_ ComputePassDescriptor) SampleBufferAttachments() ComputePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Call[ComputePassSampleBufferAttachmentDescriptorArray](c_, objc.Sel("sampleBufferAttachments"))
	return rv
}

// The strategy to use for dispatching any compute commands encoded in the compute pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepassdescriptor/3564435-dispatchtype?language=objc
func (c_ ComputePassDescriptor) DispatchType() DispatchType {
	rv := objc.Call[DispatchType](c_, objc.Sel("dispatchType"))
	return rv
}

// The strategy to use for dispatching any compute commands encoded in the compute pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepassdescriptor/3564435-dispatchtype?language=objc
func (c_ ComputePassDescriptor) SetDispatchType(value DispatchType) {
	objc.Call[objc.Void](c_, objc.Sel("setDispatchType:"), value)
}
