// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [BlitPassSampleBufferAttachmentDescriptor] class.
var BlitPassSampleBufferAttachmentDescriptorClass = _BlitPassSampleBufferAttachmentDescriptorClass{objc.GetClass("MTLBlitPassSampleBufferAttachmentDescriptor")}

type _BlitPassSampleBufferAttachmentDescriptorClass struct {
	objc.Class
}

// An interface definition for the [BlitPassSampleBufferAttachmentDescriptor] class.
type IBlitPassSampleBufferAttachmentDescriptor interface {
	objc.IObject
	StartOfEncoderSampleIndex() uint
	SetStartOfEncoderSampleIndex(value uint)
	EndOfEncoderSampleIndex() uint
	SetEndOfEncoderSampleIndex(value uint)
	SampleBuffer() CounterSampleBufferWrapper
	SetSampleBuffer(value PCounterSampleBuffer)
	SetSampleBufferObject(valueObject objc.IObject)
}

// A configuration that instructs the GPU where to store counter data from the beginning and end of a blit pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitpasssamplebufferattachmentdescriptor?language=objc
type BlitPassSampleBufferAttachmentDescriptor struct {
	objc.Object
}

func BlitPassSampleBufferAttachmentDescriptorFrom(ptr unsafe.Pointer) BlitPassSampleBufferAttachmentDescriptor {
	return BlitPassSampleBufferAttachmentDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (bc _BlitPassSampleBufferAttachmentDescriptorClass) Alloc() BlitPassSampleBufferAttachmentDescriptor {
	rv := objc.Call[BlitPassSampleBufferAttachmentDescriptor](bc, objc.Sel("alloc"))
	return rv
}

func BlitPassSampleBufferAttachmentDescriptor_Alloc() BlitPassSampleBufferAttachmentDescriptor {
	return BlitPassSampleBufferAttachmentDescriptorClass.Alloc()
}

func (bc _BlitPassSampleBufferAttachmentDescriptorClass) New() BlitPassSampleBufferAttachmentDescriptor {
	rv := objc.Call[BlitPassSampleBufferAttachmentDescriptor](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewBlitPassSampleBufferAttachmentDescriptor() BlitPassSampleBufferAttachmentDescriptor {
	return BlitPassSampleBufferAttachmentDescriptorClass.New()
}

func (b_ BlitPassSampleBufferAttachmentDescriptor) Init() BlitPassSampleBufferAttachmentDescriptor {
	rv := objc.Call[BlitPassSampleBufferAttachmentDescriptor](b_, objc.Sel("init"))
	return rv
}

// An index within a counter sample buffer that tells the GPU where to store counter data from the start of a blit pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitpasssamplebufferattachmentdescriptor/3564426-startofencodersampleindex?language=objc
func (b_ BlitPassSampleBufferAttachmentDescriptor) StartOfEncoderSampleIndex() uint {
	rv := objc.Call[uint](b_, objc.Sel("startOfEncoderSampleIndex"))
	return rv
}

// An index within a counter sample buffer that tells the GPU where to store counter data from the start of a blit pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitpasssamplebufferattachmentdescriptor/3564426-startofencodersampleindex?language=objc
func (b_ BlitPassSampleBufferAttachmentDescriptor) SetStartOfEncoderSampleIndex(value uint) {
	objc.Call[objc.Void](b_, objc.Sel("setStartOfEncoderSampleIndex:"), value)
}

// An index within a counter sample buffer that tells the GPU where to store counter data from the end of a blit pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitpasssamplebufferattachmentdescriptor/3564424-endofencodersampleindex?language=objc
func (b_ BlitPassSampleBufferAttachmentDescriptor) EndOfEncoderSampleIndex() uint {
	rv := objc.Call[uint](b_, objc.Sel("endOfEncoderSampleIndex"))
	return rv
}

// An index within a counter sample buffer that tells the GPU where to store counter data from the end of a blit pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitpasssamplebufferattachmentdescriptor/3564424-endofencodersampleindex?language=objc
func (b_ BlitPassSampleBufferAttachmentDescriptor) SetEndOfEncoderSampleIndex(value uint) {
	objc.Call[objc.Void](b_, objc.Sel("setEndOfEncoderSampleIndex:"), value)
}

// A specialized memory buffer that the GPU uses to store its counter data during the blit pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitpasssamplebufferattachmentdescriptor/3564425-samplebuffer?language=objc
func (b_ BlitPassSampleBufferAttachmentDescriptor) SampleBuffer() CounterSampleBufferWrapper {
	rv := objc.Call[CounterSampleBufferWrapper](b_, objc.Sel("sampleBuffer"))
	return rv
}

// A specialized memory buffer that the GPU uses to store its counter data during the blit pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitpasssamplebufferattachmentdescriptor/3564425-samplebuffer?language=objc
func (b_ BlitPassSampleBufferAttachmentDescriptor) SetSampleBuffer(value PCounterSampleBuffer) {
	po0 := objc.WrapAsProtocol("MTLCounterSampleBuffer", value)
	objc.Call[objc.Void](b_, objc.Sel("setSampleBuffer:"), po0)
}

// A specialized memory buffer that the GPU uses to store its counter data during the blit pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitpasssamplebufferattachmentdescriptor/3564425-samplebuffer?language=objc
func (b_ BlitPassSampleBufferAttachmentDescriptor) SetSampleBufferObject(valueObject objc.IObject) {
	objc.Call[objc.Void](b_, objc.Sel("setSampleBuffer:"), objc.Ptr(valueObject))
}
