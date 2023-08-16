// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RenderPassSampleBufferAttachmentDescriptorArray] class.
var RenderPassSampleBufferAttachmentDescriptorArrayClass = _RenderPassSampleBufferAttachmentDescriptorArrayClass{objc.GetClass("MTLRenderPassSampleBufferAttachmentDescriptorArray")}

type _RenderPassSampleBufferAttachmentDescriptorArrayClass struct {
	objc.Class
}

// An interface definition for the [RenderPassSampleBufferAttachmentDescriptorArray] class.
type IRenderPassSampleBufferAttachmentDescriptorArray interface {
	objc.IObject
	ObjectAtIndexedSubscript(attachmentIndex uint) RenderPassSampleBufferAttachmentDescriptor
	SetObjectAtIndexedSubscript(attachment IRenderPassSampleBufferAttachmentDescriptor, attachmentIndex uint)
}

// An array of sample buffer attachments for a render pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpasssamplebufferattachmentdescriptorarray?language=objc
type RenderPassSampleBufferAttachmentDescriptorArray struct {
	objc.Object
}

func RenderPassSampleBufferAttachmentDescriptorArrayFrom(ptr unsafe.Pointer) RenderPassSampleBufferAttachmentDescriptorArray {
	return RenderPassSampleBufferAttachmentDescriptorArray{
		Object: objc.ObjectFrom(ptr),
	}
}

func (rc _RenderPassSampleBufferAttachmentDescriptorArrayClass) Alloc() RenderPassSampleBufferAttachmentDescriptorArray {
	rv := objc.Call[RenderPassSampleBufferAttachmentDescriptorArray](rc, objc.Sel("alloc"))
	return rv
}

func RenderPassSampleBufferAttachmentDescriptorArray_Alloc() RenderPassSampleBufferAttachmentDescriptorArray {
	return RenderPassSampleBufferAttachmentDescriptorArrayClass.Alloc()
}

func (rc _RenderPassSampleBufferAttachmentDescriptorArrayClass) New() RenderPassSampleBufferAttachmentDescriptorArray {
	rv := objc.Call[RenderPassSampleBufferAttachmentDescriptorArray](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRenderPassSampleBufferAttachmentDescriptorArray() RenderPassSampleBufferAttachmentDescriptorArray {
	return RenderPassSampleBufferAttachmentDescriptorArrayClass.New()
}

func (r_ RenderPassSampleBufferAttachmentDescriptorArray) Init() RenderPassSampleBufferAttachmentDescriptorArray {
	rv := objc.Call[RenderPassSampleBufferAttachmentDescriptorArray](r_, objc.Sel("init"))
	return rv
}

// Returns the descriptor object for the specified sample buffer attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpasssamplebufferattachmentdescriptorarray/3081754-objectatindexedsubscript?language=objc
func (r_ RenderPassSampleBufferAttachmentDescriptorArray) ObjectAtIndexedSubscript(attachmentIndex uint) RenderPassSampleBufferAttachmentDescriptor {
	rv := objc.Call[RenderPassSampleBufferAttachmentDescriptor](r_, objc.Sel("objectAtIndexedSubscript:"), attachmentIndex)
	return rv
}

// Sets the descriptor object for the specified sample buffer attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpasssamplebufferattachmentdescriptorarray/3081755-setobject?language=objc
func (r_ RenderPassSampleBufferAttachmentDescriptorArray) SetObjectAtIndexedSubscript(attachment IRenderPassSampleBufferAttachmentDescriptor, attachmentIndex uint) {
	objc.Call[objc.Void](r_, objc.Sel("setObject:atIndexedSubscript:"), objc.Ptr(attachment), attachmentIndex)
}
