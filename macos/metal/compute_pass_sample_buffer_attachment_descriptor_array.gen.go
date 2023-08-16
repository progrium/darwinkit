// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ComputePassSampleBufferAttachmentDescriptorArray] class.
var ComputePassSampleBufferAttachmentDescriptorArrayClass = _ComputePassSampleBufferAttachmentDescriptorArrayClass{objc.GetClass("MTLComputePassSampleBufferAttachmentDescriptorArray")}

type _ComputePassSampleBufferAttachmentDescriptorArrayClass struct {
	objc.Class
}

// An interface definition for the [ComputePassSampleBufferAttachmentDescriptorArray] class.
type IComputePassSampleBufferAttachmentDescriptorArray interface {
	objc.IObject
	ObjectAtIndexedSubscript(attachmentIndex uint) ComputePassSampleBufferAttachmentDescriptor
	SetObjectAtIndexedSubscript(attachment IComputePassSampleBufferAttachmentDescriptor, attachmentIndex uint)
}

// An array of sample buffer attachments for a compute pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepasssamplebufferattachmentdescriptorarray?language=objc
type ComputePassSampleBufferAttachmentDescriptorArray struct {
	objc.Object
}

func ComputePassSampleBufferAttachmentDescriptorArrayFrom(ptr unsafe.Pointer) ComputePassSampleBufferAttachmentDescriptorArray {
	return ComputePassSampleBufferAttachmentDescriptorArray{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _ComputePassSampleBufferAttachmentDescriptorArrayClass) Alloc() ComputePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Call[ComputePassSampleBufferAttachmentDescriptorArray](cc, objc.Sel("alloc"))
	return rv
}

func ComputePassSampleBufferAttachmentDescriptorArray_Alloc() ComputePassSampleBufferAttachmentDescriptorArray {
	return ComputePassSampleBufferAttachmentDescriptorArrayClass.Alloc()
}

func (cc _ComputePassSampleBufferAttachmentDescriptorArrayClass) New() ComputePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Call[ComputePassSampleBufferAttachmentDescriptorArray](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewComputePassSampleBufferAttachmentDescriptorArray() ComputePassSampleBufferAttachmentDescriptorArray {
	return ComputePassSampleBufferAttachmentDescriptorArrayClass.New()
}

func (c_ ComputePassSampleBufferAttachmentDescriptorArray) Init() ComputePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Call[ComputePassSampleBufferAttachmentDescriptorArray](c_, objc.Sel("init"))
	return rv
}

// Returns the descriptor object for the specified sample buffer attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepasssamplebufferattachmentdescriptorarray/3564442-objectatindexedsubscript?language=objc
func (c_ ComputePassSampleBufferAttachmentDescriptorArray) ObjectAtIndexedSubscript(attachmentIndex uint) ComputePassSampleBufferAttachmentDescriptor {
	rv := objc.Call[ComputePassSampleBufferAttachmentDescriptor](c_, objc.Sel("objectAtIndexedSubscript:"), attachmentIndex)
	return rv
}

// Sets the descriptor object for the specified sample buffer attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepasssamplebufferattachmentdescriptorarray/3564443-setobject?language=objc
func (c_ ComputePassSampleBufferAttachmentDescriptorArray) SetObjectAtIndexedSubscript(attachment IComputePassSampleBufferAttachmentDescriptor, attachmentIndex uint) {
	objc.Call[objc.Void](c_, objc.Sel("setObject:atIndexedSubscript:"), objc.Ptr(attachment), attachmentIndex)
}
