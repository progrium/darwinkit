// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RenderPipelineColorAttachmentDescriptorArray] class.
var RenderPipelineColorAttachmentDescriptorArrayClass = _RenderPipelineColorAttachmentDescriptorArrayClass{objc.GetClass("MTLRenderPipelineColorAttachmentDescriptorArray")}

type _RenderPipelineColorAttachmentDescriptorArrayClass struct {
	objc.Class
}

// An interface definition for the [RenderPipelineColorAttachmentDescriptorArray] class.
type IRenderPipelineColorAttachmentDescriptorArray interface {
	objc.IObject
	ObjectAtIndexedSubscript(attachmentIndex uint) RenderPipelineColorAttachmentDescriptor
	SetObjectAtIndexedSubscript(attachment IRenderPipelineColorAttachmentDescriptor, attachmentIndex uint)
}

// An array of render pipeline color attachment descriptor objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinecolorattachmentdescriptorarray?language=objc
type RenderPipelineColorAttachmentDescriptorArray struct {
	objc.Object
}

func RenderPipelineColorAttachmentDescriptorArrayFrom(ptr unsafe.Pointer) RenderPipelineColorAttachmentDescriptorArray {
	return RenderPipelineColorAttachmentDescriptorArray{
		Object: objc.ObjectFrom(ptr),
	}
}

func (rc _RenderPipelineColorAttachmentDescriptorArrayClass) Alloc() RenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Call[RenderPipelineColorAttachmentDescriptorArray](rc, objc.Sel("alloc"))
	return rv
}

func RenderPipelineColorAttachmentDescriptorArray_Alloc() RenderPipelineColorAttachmentDescriptorArray {
	return RenderPipelineColorAttachmentDescriptorArrayClass.Alloc()
}

func (rc _RenderPipelineColorAttachmentDescriptorArrayClass) New() RenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Call[RenderPipelineColorAttachmentDescriptorArray](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRenderPipelineColorAttachmentDescriptorArray() RenderPipelineColorAttachmentDescriptorArray {
	return RenderPipelineColorAttachmentDescriptorArrayClass.New()
}

func (r_ RenderPipelineColorAttachmentDescriptorArray) Init() RenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Call[RenderPipelineColorAttachmentDescriptorArray](r_, objc.Sel("init"))
	return rv
}

// Returns the render pipeline state for the specified color attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinecolorattachmentdescriptorarray/1514673-objectatindexedsubscript?language=objc
func (r_ RenderPipelineColorAttachmentDescriptorArray) ObjectAtIndexedSubscript(attachmentIndex uint) RenderPipelineColorAttachmentDescriptor {
	rv := objc.Call[RenderPipelineColorAttachmentDescriptor](r_, objc.Sel("objectAtIndexedSubscript:"), attachmentIndex)
	return rv
}

// Sets the render pipeline state for a specified color attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinecolorattachmentdescriptorarray/1514675-setobject?language=objc
func (r_ RenderPipelineColorAttachmentDescriptorArray) SetObjectAtIndexedSubscript(attachment IRenderPipelineColorAttachmentDescriptor, attachmentIndex uint) {
	objc.Call[objc.Void](r_, objc.Sel("setObject:atIndexedSubscript:"), objc.Ptr(attachment), attachmentIndex)
}
