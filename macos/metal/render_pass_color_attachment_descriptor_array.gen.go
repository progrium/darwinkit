// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RenderPassColorAttachmentDescriptorArray] class.
var RenderPassColorAttachmentDescriptorArrayClass = _RenderPassColorAttachmentDescriptorArrayClass{objc.GetClass("MTLRenderPassColorAttachmentDescriptorArray")}

type _RenderPassColorAttachmentDescriptorArrayClass struct {
	objc.Class
}

// An interface definition for the [RenderPassColorAttachmentDescriptorArray] class.
type IRenderPassColorAttachmentDescriptorArray interface {
	objc.IObject
	ObjectAtIndexedSubscript(attachmentIndex uint) RenderPassColorAttachmentDescriptor
	SetObjectAtIndexedSubscript(attachment IRenderPassColorAttachmentDescriptor, attachmentIndex uint)
}

// An array of render pass color attachment descriptor objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpasscolorattachmentdescriptorarray?language=objc
type RenderPassColorAttachmentDescriptorArray struct {
	objc.Object
}

func RenderPassColorAttachmentDescriptorArrayFrom(ptr unsafe.Pointer) RenderPassColorAttachmentDescriptorArray {
	return RenderPassColorAttachmentDescriptorArray{
		Object: objc.ObjectFrom(ptr),
	}
}

func (rc _RenderPassColorAttachmentDescriptorArrayClass) Alloc() RenderPassColorAttachmentDescriptorArray {
	rv := objc.Call[RenderPassColorAttachmentDescriptorArray](rc, objc.Sel("alloc"))
	return rv
}

func RenderPassColorAttachmentDescriptorArray_Alloc() RenderPassColorAttachmentDescriptorArray {
	return RenderPassColorAttachmentDescriptorArrayClass.Alloc()
}

func (rc _RenderPassColorAttachmentDescriptorArrayClass) New() RenderPassColorAttachmentDescriptorArray {
	rv := objc.Call[RenderPassColorAttachmentDescriptorArray](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRenderPassColorAttachmentDescriptorArray() RenderPassColorAttachmentDescriptorArray {
	return RenderPassColorAttachmentDescriptorArrayClass.New()
}

func (r_ RenderPassColorAttachmentDescriptorArray) Init() RenderPassColorAttachmentDescriptorArray {
	rv := objc.Call[RenderPassColorAttachmentDescriptorArray](r_, objc.Sel("init"))
	return rv
}

// Returns the descriptor object for the specified color attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpasscolorattachmentdescriptorarray/1437977-objectatindexedsubscript?language=objc
func (r_ RenderPassColorAttachmentDescriptorArray) ObjectAtIndexedSubscript(attachmentIndex uint) RenderPassColorAttachmentDescriptor {
	rv := objc.Call[RenderPassColorAttachmentDescriptor](r_, objc.Sel("objectAtIndexedSubscript:"), attachmentIndex)
	return rv
}

// Sets the descriptor for the specified color attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpasscolorattachmentdescriptorarray/1437982-setobject?language=objc
func (r_ RenderPassColorAttachmentDescriptorArray) SetObjectAtIndexedSubscript(attachment IRenderPassColorAttachmentDescriptor, attachmentIndex uint) {
	objc.Call[objc.Void](r_, objc.Sel("setObject:atIndexedSubscript:"), objc.Ptr(attachment), attachmentIndex)
}
