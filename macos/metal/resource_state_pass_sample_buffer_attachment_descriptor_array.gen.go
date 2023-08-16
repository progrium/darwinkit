// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ResourceStatePassSampleBufferAttachmentDescriptorArray] class.
var ResourceStatePassSampleBufferAttachmentDescriptorArrayClass = _ResourceStatePassSampleBufferAttachmentDescriptorArrayClass{objc.GetClass("MTLResourceStatePassSampleBufferAttachmentDescriptorArray")}

type _ResourceStatePassSampleBufferAttachmentDescriptorArrayClass struct {
	objc.Class
}

// An interface definition for the [ResourceStatePassSampleBufferAttachmentDescriptorArray] class.
type IResourceStatePassSampleBufferAttachmentDescriptorArray interface {
	objc.IObject
	ObjectAtIndexedSubscript(attachmentIndex uint) ResourceStatePassSampleBufferAttachmentDescriptor
	SetObjectAtIndexedSubscript(attachment IResourceStatePassSampleBufferAttachmentDescriptor, attachmentIndex uint)
}

// An array of sample buffer attachments for a resource state pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourcestatepasssamplebufferattachmentdescriptorarray?language=objc
type ResourceStatePassSampleBufferAttachmentDescriptorArray struct {
	objc.Object
}

func ResourceStatePassSampleBufferAttachmentDescriptorArrayFrom(ptr unsafe.Pointer) ResourceStatePassSampleBufferAttachmentDescriptorArray {
	return ResourceStatePassSampleBufferAttachmentDescriptorArray{
		Object: objc.ObjectFrom(ptr),
	}
}

func (rc _ResourceStatePassSampleBufferAttachmentDescriptorArrayClass) Alloc() ResourceStatePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Call[ResourceStatePassSampleBufferAttachmentDescriptorArray](rc, objc.Sel("alloc"))
	return rv
}

func ResourceStatePassSampleBufferAttachmentDescriptorArray_Alloc() ResourceStatePassSampleBufferAttachmentDescriptorArray {
	return ResourceStatePassSampleBufferAttachmentDescriptorArrayClass.Alloc()
}

func (rc _ResourceStatePassSampleBufferAttachmentDescriptorArrayClass) New() ResourceStatePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Call[ResourceStatePassSampleBufferAttachmentDescriptorArray](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewResourceStatePassSampleBufferAttachmentDescriptorArray() ResourceStatePassSampleBufferAttachmentDescriptorArray {
	return ResourceStatePassSampleBufferAttachmentDescriptorArrayClass.New()
}

func (r_ ResourceStatePassSampleBufferAttachmentDescriptorArray) Init() ResourceStatePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Call[ResourceStatePassSampleBufferAttachmentDescriptorArray](r_, objc.Sel("init"))
	return rv
}

// Returns the descriptor object for the specified sample buffer attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourcestatepasssamplebufferattachmentdescriptorarray/3566570-objectatindexedsubscript?language=objc
func (r_ ResourceStatePassSampleBufferAttachmentDescriptorArray) ObjectAtIndexedSubscript(attachmentIndex uint) ResourceStatePassSampleBufferAttachmentDescriptor {
	rv := objc.Call[ResourceStatePassSampleBufferAttachmentDescriptor](r_, objc.Sel("objectAtIndexedSubscript:"), attachmentIndex)
	return rv
}

// Sets the descriptor object for the specified sample buffer attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourcestatepasssamplebufferattachmentdescriptorarray/3566571-setobject?language=objc
func (r_ ResourceStatePassSampleBufferAttachmentDescriptorArray) SetObjectAtIndexedSubscript(attachment IResourceStatePassSampleBufferAttachmentDescriptor, attachmentIndex uint) {
	objc.Call[objc.Void](r_, objc.Sel("setObject:atIndexedSubscript:"), objc.Ptr(attachment), attachmentIndex)
}
