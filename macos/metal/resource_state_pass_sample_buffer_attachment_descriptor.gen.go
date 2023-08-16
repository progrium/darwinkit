// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ResourceStatePassSampleBufferAttachmentDescriptor] class.
var ResourceStatePassSampleBufferAttachmentDescriptorClass = _ResourceStatePassSampleBufferAttachmentDescriptorClass{objc.GetClass("MTLResourceStatePassSampleBufferAttachmentDescriptor")}

type _ResourceStatePassSampleBufferAttachmentDescriptorClass struct {
	objc.Class
}

// An interface definition for the [ResourceStatePassSampleBufferAttachmentDescriptor] class.
type IResourceStatePassSampleBufferAttachmentDescriptor interface {
	objc.IObject
	StartOfEncoderSampleIndex() uint
	SetStartOfEncoderSampleIndex(value uint)
	EndOfEncoderSampleIndex() uint
	SetEndOfEncoderSampleIndex(value uint)
	SampleBuffer() CounterSampleBufferWrapper
	SetSampleBuffer(value PCounterSampleBuffer)
	SetSampleBufferObject(valueObject objc.IObject)
}

// A description of where to store GPU counter information at the start and end of a resource state pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourcestatepasssamplebufferattachmentdescriptor?language=objc
type ResourceStatePassSampleBufferAttachmentDescriptor struct {
	objc.Object
}

func ResourceStatePassSampleBufferAttachmentDescriptorFrom(ptr unsafe.Pointer) ResourceStatePassSampleBufferAttachmentDescriptor {
	return ResourceStatePassSampleBufferAttachmentDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (rc _ResourceStatePassSampleBufferAttachmentDescriptorClass) Alloc() ResourceStatePassSampleBufferAttachmentDescriptor {
	rv := objc.Call[ResourceStatePassSampleBufferAttachmentDescriptor](rc, objc.Sel("alloc"))
	return rv
}

func ResourceStatePassSampleBufferAttachmentDescriptor_Alloc() ResourceStatePassSampleBufferAttachmentDescriptor {
	return ResourceStatePassSampleBufferAttachmentDescriptorClass.Alloc()
}

func (rc _ResourceStatePassSampleBufferAttachmentDescriptorClass) New() ResourceStatePassSampleBufferAttachmentDescriptor {
	rv := objc.Call[ResourceStatePassSampleBufferAttachmentDescriptor](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewResourceStatePassSampleBufferAttachmentDescriptor() ResourceStatePassSampleBufferAttachmentDescriptor {
	return ResourceStatePassSampleBufferAttachmentDescriptorClass.New()
}

func (r_ ResourceStatePassSampleBufferAttachmentDescriptor) Init() ResourceStatePassSampleBufferAttachmentDescriptor {
	rv := objc.Call[ResourceStatePassSampleBufferAttachmentDescriptor](r_, objc.Sel("init"))
	return rv
}

// The index the Metal device object should use to store GPU counters when starting the resource state pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourcestatepasssamplebufferattachmentdescriptor/3566568-startofencodersampleindex?language=objc
func (r_ ResourceStatePassSampleBufferAttachmentDescriptor) StartOfEncoderSampleIndex() uint {
	rv := objc.Call[uint](r_, objc.Sel("startOfEncoderSampleIndex"))
	return rv
}

// The index the Metal device object should use to store GPU counters when starting the resource state pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourcestatepasssamplebufferattachmentdescriptor/3566568-startofencodersampleindex?language=objc
func (r_ ResourceStatePassSampleBufferAttachmentDescriptor) SetStartOfEncoderSampleIndex(value uint) {
	objc.Call[objc.Void](r_, objc.Sel("setStartOfEncoderSampleIndex:"), value)
}

// The index the Metal device object should use to store GPU counters when ending the resource state pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourcestatepasssamplebufferattachmentdescriptor/3566566-endofencodersampleindex?language=objc
func (r_ ResourceStatePassSampleBufferAttachmentDescriptor) EndOfEncoderSampleIndex() uint {
	rv := objc.Call[uint](r_, objc.Sel("endOfEncoderSampleIndex"))
	return rv
}

// The index the Metal device object should use to store GPU counters when ending the resource state pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourcestatepasssamplebufferattachmentdescriptor/3566566-endofencodersampleindex?language=objc
func (r_ ResourceStatePassSampleBufferAttachmentDescriptor) SetEndOfEncoderSampleIndex(value uint) {
	objc.Call[objc.Void](r_, objc.Sel("setEndOfEncoderSampleIndex:"), value)
}

// A specialized memory buffer that the GPU uses to store its counter data during the resource state pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourcestatepasssamplebufferattachmentdescriptor/3566567-samplebuffer?language=objc
func (r_ ResourceStatePassSampleBufferAttachmentDescriptor) SampleBuffer() CounterSampleBufferWrapper {
	rv := objc.Call[CounterSampleBufferWrapper](r_, objc.Sel("sampleBuffer"))
	return rv
}

// A specialized memory buffer that the GPU uses to store its counter data during the resource state pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourcestatepasssamplebufferattachmentdescriptor/3566567-samplebuffer?language=objc
func (r_ ResourceStatePassSampleBufferAttachmentDescriptor) SetSampleBuffer(value PCounterSampleBuffer) {
	po0 := objc.WrapAsProtocol("MTLCounterSampleBuffer", value)
	objc.Call[objc.Void](r_, objc.Sel("setSampleBuffer:"), po0)
}

// A specialized memory buffer that the GPU uses to store its counter data during the resource state pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourcestatepasssamplebufferattachmentdescriptor/3566567-samplebuffer?language=objc
func (r_ ResourceStatePassSampleBufferAttachmentDescriptor) SetSampleBufferObject(valueObject objc.IObject) {
	objc.Call[objc.Void](r_, objc.Sel("setSampleBuffer:"), objc.Ptr(valueObject))
}
