// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RenderPassSampleBufferAttachmentDescriptor] class.
var RenderPassSampleBufferAttachmentDescriptorClass = _RenderPassSampleBufferAttachmentDescriptorClass{objc.GetClass("MTLRenderPassSampleBufferAttachmentDescriptor")}

type _RenderPassSampleBufferAttachmentDescriptorClass struct {
	objc.Class
}

// An interface definition for the [RenderPassSampleBufferAttachmentDescriptor] class.
type IRenderPassSampleBufferAttachmentDescriptor interface {
	objc.IObject
	EndOfVertexSampleIndex() uint
	SetEndOfVertexSampleIndex(value uint)
	StartOfVertexSampleIndex() uint
	SetStartOfVertexSampleIndex(value uint)
	EndOfFragmentSampleIndex() uint
	SetEndOfFragmentSampleIndex(value uint)
	StartOfFragmentSampleIndex() uint
	SetStartOfFragmentSampleIndex(value uint)
	SampleBuffer() CounterSampleBufferWrapper
	SetSampleBuffer(value PCounterSampleBuffer)
	SetSampleBufferObject(valueObject objc.IObject)
}

// A description of where to store GPU counter information at the start and end of a render pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpasssamplebufferattachmentdescriptor?language=objc
type RenderPassSampleBufferAttachmentDescriptor struct {
	objc.Object
}

func RenderPassSampleBufferAttachmentDescriptorFrom(ptr unsafe.Pointer) RenderPassSampleBufferAttachmentDescriptor {
	return RenderPassSampleBufferAttachmentDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (rc _RenderPassSampleBufferAttachmentDescriptorClass) Alloc() RenderPassSampleBufferAttachmentDescriptor {
	rv := objc.Call[RenderPassSampleBufferAttachmentDescriptor](rc, objc.Sel("alloc"))
	return rv
}

func RenderPassSampleBufferAttachmentDescriptor_Alloc() RenderPassSampleBufferAttachmentDescriptor {
	return RenderPassSampleBufferAttachmentDescriptorClass.Alloc()
}

func (rc _RenderPassSampleBufferAttachmentDescriptorClass) New() RenderPassSampleBufferAttachmentDescriptor {
	rv := objc.Call[RenderPassSampleBufferAttachmentDescriptor](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRenderPassSampleBufferAttachmentDescriptor() RenderPassSampleBufferAttachmentDescriptor {
	return RenderPassSampleBufferAttachmentDescriptorClass.New()
}

func (r_ RenderPassSampleBufferAttachmentDescriptor) Init() RenderPassSampleBufferAttachmentDescriptor {
	rv := objc.Call[RenderPassSampleBufferAttachmentDescriptor](r_, objc.Sel("init"))
	return rv
}

// The index the Metal device object should use to store GPU counters when ending the render pass’s vertex stage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpasssamplebufferattachmentdescriptor/3081770-endofvertexsampleindex?language=objc
func (r_ RenderPassSampleBufferAttachmentDescriptor) EndOfVertexSampleIndex() uint {
	rv := objc.Call[uint](r_, objc.Sel("endOfVertexSampleIndex"))
	return rv
}

// The index the Metal device object should use to store GPU counters when ending the render pass’s vertex stage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpasssamplebufferattachmentdescriptor/3081770-endofvertexsampleindex?language=objc
func (r_ RenderPassSampleBufferAttachmentDescriptor) SetEndOfVertexSampleIndex(value uint) {
	objc.Call[objc.Void](r_, objc.Sel("setEndOfVertexSampleIndex:"), value)
}

// The index the Metal device object should use to store GPU counters when starting the render pass’s vertex stage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpasssamplebufferattachmentdescriptor/3081772-startofvertexsampleindex?language=objc
func (r_ RenderPassSampleBufferAttachmentDescriptor) StartOfVertexSampleIndex() uint {
	rv := objc.Call[uint](r_, objc.Sel("startOfVertexSampleIndex"))
	return rv
}

// The index the Metal device object should use to store GPU counters when starting the render pass’s vertex stage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpasssamplebufferattachmentdescriptor/3081772-startofvertexsampleindex?language=objc
func (r_ RenderPassSampleBufferAttachmentDescriptor) SetStartOfVertexSampleIndex(value uint) {
	objc.Call[objc.Void](r_, objc.Sel("setStartOfVertexSampleIndex:"), value)
}

// The index the Metal device object should use to store GPU counters when ending the render pass’s fragment stage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpasssamplebufferattachmentdescriptor/3081769-endoffragmentsampleindex?language=objc
func (r_ RenderPassSampleBufferAttachmentDescriptor) EndOfFragmentSampleIndex() uint {
	rv := objc.Call[uint](r_, objc.Sel("endOfFragmentSampleIndex"))
	return rv
}

// The index the Metal device object should use to store GPU counters when ending the render pass’s fragment stage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpasssamplebufferattachmentdescriptor/3081769-endoffragmentsampleindex?language=objc
func (r_ RenderPassSampleBufferAttachmentDescriptor) SetEndOfFragmentSampleIndex(value uint) {
	objc.Call[objc.Void](r_, objc.Sel("setEndOfFragmentSampleIndex:"), value)
}

// The index the Metal device object should use to store GPU counters when starting the render pass’s fragment stage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpasssamplebufferattachmentdescriptor/3081771-startoffragmentsampleindex?language=objc
func (r_ RenderPassSampleBufferAttachmentDescriptor) StartOfFragmentSampleIndex() uint {
	rv := objc.Call[uint](r_, objc.Sel("startOfFragmentSampleIndex"))
	return rv
}

// The index the Metal device object should use to store GPU counters when starting the render pass’s fragment stage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpasssamplebufferattachmentdescriptor/3081771-startoffragmentsampleindex?language=objc
func (r_ RenderPassSampleBufferAttachmentDescriptor) SetStartOfFragmentSampleIndex(value uint) {
	objc.Call[objc.Void](r_, objc.Sel("setStartOfFragmentSampleIndex:"), value)
}

// A specialized memory buffer that the GPU uses to store its counter data during the render pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpasssamplebufferattachmentdescriptor/3081752-samplebuffer?language=objc
func (r_ RenderPassSampleBufferAttachmentDescriptor) SampleBuffer() CounterSampleBufferWrapper {
	rv := objc.Call[CounterSampleBufferWrapper](r_, objc.Sel("sampleBuffer"))
	return rv
}

// A specialized memory buffer that the GPU uses to store its counter data during the render pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpasssamplebufferattachmentdescriptor/3081752-samplebuffer?language=objc
func (r_ RenderPassSampleBufferAttachmentDescriptor) SetSampleBuffer(value PCounterSampleBuffer) {
	po0 := objc.WrapAsProtocol("MTLCounterSampleBuffer", value)
	objc.Call[objc.Void](r_, objc.Sel("setSampleBuffer:"), po0)
}

// A specialized memory buffer that the GPU uses to store its counter data during the render pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpasssamplebufferattachmentdescriptor/3081752-samplebuffer?language=objc
func (r_ RenderPassSampleBufferAttachmentDescriptor) SetSampleBufferObject(valueObject objc.IObject) {
	objc.Call[objc.Void](r_, objc.Sel("setSampleBuffer:"), objc.Ptr(valueObject))
}
