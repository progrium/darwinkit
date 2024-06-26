// Code generated by DarwinKit. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [ComputePassSampleBufferAttachmentDescriptor] class.
var ComputePassSampleBufferAttachmentDescriptorClass = _ComputePassSampleBufferAttachmentDescriptorClass{objc.GetClass("MTLComputePassSampleBufferAttachmentDescriptor")}

type _ComputePassSampleBufferAttachmentDescriptorClass struct {
	objc.Class
}

// An interface definition for the [ComputePassSampleBufferAttachmentDescriptor] class.
type IComputePassSampleBufferAttachmentDescriptor interface {
	objc.IObject
	EndOfEncoderSampleIndex() uint
	SetEndOfEncoderSampleIndex(value uint)
	StartOfEncoderSampleIndex() uint
	SetStartOfEncoderSampleIndex(value uint)
	SampleBuffer() CounterSampleBufferObject
	SetSampleBuffer(value PCounterSampleBuffer)
	SetSampleBufferObject(valueObject objc.IObject)
}

// A description of where to store GPU counter information at the start and end of a compute pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepasssamplebufferattachmentdescriptor?language=objc
type ComputePassSampleBufferAttachmentDescriptor struct {
	objc.Object
}

func ComputePassSampleBufferAttachmentDescriptorFrom(ptr unsafe.Pointer) ComputePassSampleBufferAttachmentDescriptor {
	return ComputePassSampleBufferAttachmentDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _ComputePassSampleBufferAttachmentDescriptorClass) Alloc() ComputePassSampleBufferAttachmentDescriptor {
	rv := objc.Call[ComputePassSampleBufferAttachmentDescriptor](cc, objc.Sel("alloc"))
	return rv
}

func (cc _ComputePassSampleBufferAttachmentDescriptorClass) New() ComputePassSampleBufferAttachmentDescriptor {
	rv := objc.Call[ComputePassSampleBufferAttachmentDescriptor](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewComputePassSampleBufferAttachmentDescriptor() ComputePassSampleBufferAttachmentDescriptor {
	return ComputePassSampleBufferAttachmentDescriptorClass.New()
}

func (c_ ComputePassSampleBufferAttachmentDescriptor) Init() ComputePassSampleBufferAttachmentDescriptor {
	rv := objc.Call[ComputePassSampleBufferAttachmentDescriptor](c_, objc.Sel("init"))
	return rv
}

// The index the Metal device object uses to store GPU counters when ending the compute pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepasssamplebufferattachmentdescriptor/3564438-endofencodersampleindex?language=objc
func (c_ ComputePassSampleBufferAttachmentDescriptor) EndOfEncoderSampleIndex() uint {
	rv := objc.Call[uint](c_, objc.Sel("endOfEncoderSampleIndex"))
	return rv
}

// The index the Metal device object uses to store GPU counters when ending the compute pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepasssamplebufferattachmentdescriptor/3564438-endofencodersampleindex?language=objc
func (c_ ComputePassSampleBufferAttachmentDescriptor) SetEndOfEncoderSampleIndex(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setEndOfEncoderSampleIndex:"), value)
}

// The index the Metal device object uses to store GPU counters when starting the compute pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepasssamplebufferattachmentdescriptor/3564440-startofencodersampleindex?language=objc
func (c_ ComputePassSampleBufferAttachmentDescriptor) StartOfEncoderSampleIndex() uint {
	rv := objc.Call[uint](c_, objc.Sel("startOfEncoderSampleIndex"))
	return rv
}

// The index the Metal device object uses to store GPU counters when starting the compute pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepasssamplebufferattachmentdescriptor/3564440-startofencodersampleindex?language=objc
func (c_ ComputePassSampleBufferAttachmentDescriptor) SetStartOfEncoderSampleIndex(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setStartOfEncoderSampleIndex:"), value)
}

// A specialized memory buffer that the GPU uses to store its counter data during the compute pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepasssamplebufferattachmentdescriptor/3564439-samplebuffer?language=objc
func (c_ ComputePassSampleBufferAttachmentDescriptor) SampleBuffer() CounterSampleBufferObject {
	rv := objc.Call[CounterSampleBufferObject](c_, objc.Sel("sampleBuffer"))
	return rv
}

// A specialized memory buffer that the GPU uses to store its counter data during the compute pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepasssamplebufferattachmentdescriptor/3564439-samplebuffer?language=objc
func (c_ ComputePassSampleBufferAttachmentDescriptor) SetSampleBuffer(value PCounterSampleBuffer) {
	po0 := objc.WrapAsProtocol("MTLCounterSampleBuffer", value)
	objc.Call[objc.Void](c_, objc.Sel("setSampleBuffer:"), po0)
}

// A specialized memory buffer that the GPU uses to store its counter data during the compute pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepasssamplebufferattachmentdescriptor/3564439-samplebuffer?language=objc
func (c_ ComputePassSampleBufferAttachmentDescriptor) SetSampleBufferObject(valueObject objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setSampleBuffer:"), valueObject)
}
