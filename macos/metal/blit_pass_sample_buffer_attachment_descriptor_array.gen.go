// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [BlitPassSampleBufferAttachmentDescriptorArray] class.
var BlitPassSampleBufferAttachmentDescriptorArrayClass = _BlitPassSampleBufferAttachmentDescriptorArrayClass{objc.GetClass("MTLBlitPassSampleBufferAttachmentDescriptorArray")}

type _BlitPassSampleBufferAttachmentDescriptorArrayClass struct {
	objc.Class
}

// An interface definition for the [BlitPassSampleBufferAttachmentDescriptorArray] class.
type IBlitPassSampleBufferAttachmentDescriptorArray interface {
	objc.IObject
	ObjectAtIndexedSubscript(attachmentIndex uint) BlitPassSampleBufferAttachmentDescriptor
	SetObjectAtIndexedSubscript(attachment IBlitPassSampleBufferAttachmentDescriptor, attachmentIndex uint)
}

// A container that stores an array of sample buffer attachments for a blit pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitpasssamplebufferattachmentdescriptorarray?language=objc
type BlitPassSampleBufferAttachmentDescriptorArray struct {
	objc.Object
}

func BlitPassSampleBufferAttachmentDescriptorArrayFrom(ptr unsafe.Pointer) BlitPassSampleBufferAttachmentDescriptorArray {
	return BlitPassSampleBufferAttachmentDescriptorArray{
		Object: objc.ObjectFrom(ptr),
	}
}

func (bc _BlitPassSampleBufferAttachmentDescriptorArrayClass) Alloc() BlitPassSampleBufferAttachmentDescriptorArray {
	rv := objc.Call[BlitPassSampleBufferAttachmentDescriptorArray](bc, objc.Sel("alloc"))
	return rv
}

func BlitPassSampleBufferAttachmentDescriptorArray_Alloc() BlitPassSampleBufferAttachmentDescriptorArray {
	return BlitPassSampleBufferAttachmentDescriptorArrayClass.Alloc()
}

func (bc _BlitPassSampleBufferAttachmentDescriptorArrayClass) New() BlitPassSampleBufferAttachmentDescriptorArray {
	rv := objc.Call[BlitPassSampleBufferAttachmentDescriptorArray](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewBlitPassSampleBufferAttachmentDescriptorArray() BlitPassSampleBufferAttachmentDescriptorArray {
	return BlitPassSampleBufferAttachmentDescriptorArrayClass.New()
}

func (b_ BlitPassSampleBufferAttachmentDescriptorArray) Init() BlitPassSampleBufferAttachmentDescriptorArray {
	rv := objc.Call[BlitPassSampleBufferAttachmentDescriptorArray](b_, objc.Sel("init"))
	return rv
}

// Returns one of the array’s blit pass sample buffer attachment descriptor instances. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitpasssamplebufferattachmentdescriptorarray/3564428-objectatindexedsubscript?language=objc
func (b_ BlitPassSampleBufferAttachmentDescriptorArray) ObjectAtIndexedSubscript(attachmentIndex uint) BlitPassSampleBufferAttachmentDescriptor {
	rv := objc.Call[BlitPassSampleBufferAttachmentDescriptor](b_, objc.Sel("objectAtIndexedSubscript:"), attachmentIndex)
	return rv
}

// Copies the properties of a blit pass sample buffer attachment descriptor instance to the properties of one of the array’s instances. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitpasssamplebufferattachmentdescriptorarray/3564429-setobject?language=objc
func (b_ BlitPassSampleBufferAttachmentDescriptorArray) SetObjectAtIndexedSubscript(attachment IBlitPassSampleBufferAttachmentDescriptor, attachmentIndex uint) {
	objc.Call[objc.Void](b_, objc.Sel("setObject:atIndexedSubscript:"), objc.Ptr(attachment), attachmentIndex)
}
