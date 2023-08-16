// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TileRenderPipelineColorAttachmentDescriptorArray] class.
var TileRenderPipelineColorAttachmentDescriptorArrayClass = _TileRenderPipelineColorAttachmentDescriptorArrayClass{objc.GetClass("MTLTileRenderPipelineColorAttachmentDescriptorArray")}

type _TileRenderPipelineColorAttachmentDescriptorArrayClass struct {
	objc.Class
}

// An interface definition for the [TileRenderPipelineColorAttachmentDescriptorArray] class.
type ITileRenderPipelineColorAttachmentDescriptorArray interface {
	objc.IObject
	ObjectAtIndexedSubscript(attachmentIndex uint) TileRenderPipelineColorAttachmentDescriptor
	SetObjectAtIndexedSubscript(attachment ITileRenderPipelineColorAttachmentDescriptor, attachmentIndex uint)
}

// An array of color attachment descriptors for the tile render pipeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltilerenderpipelinecolorattachmentdescriptorarray?language=objc
type TileRenderPipelineColorAttachmentDescriptorArray struct {
	objc.Object
}

func TileRenderPipelineColorAttachmentDescriptorArrayFrom(ptr unsafe.Pointer) TileRenderPipelineColorAttachmentDescriptorArray {
	return TileRenderPipelineColorAttachmentDescriptorArray{
		Object: objc.ObjectFrom(ptr),
	}
}

func (tc _TileRenderPipelineColorAttachmentDescriptorArrayClass) Alloc() TileRenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Call[TileRenderPipelineColorAttachmentDescriptorArray](tc, objc.Sel("alloc"))
	return rv
}

func TileRenderPipelineColorAttachmentDescriptorArray_Alloc() TileRenderPipelineColorAttachmentDescriptorArray {
	return TileRenderPipelineColorAttachmentDescriptorArrayClass.Alloc()
}

func (tc _TileRenderPipelineColorAttachmentDescriptorArrayClass) New() TileRenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Call[TileRenderPipelineColorAttachmentDescriptorArray](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTileRenderPipelineColorAttachmentDescriptorArray() TileRenderPipelineColorAttachmentDescriptorArray {
	return TileRenderPipelineColorAttachmentDescriptorArrayClass.New()
}

func (t_ TileRenderPipelineColorAttachmentDescriptorArray) Init() TileRenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Call[TileRenderPipelineColorAttachmentDescriptorArray](t_, objc.Sel("init"))
	return rv
}

// Returns the render pipeline state for the specified color attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltilerenderpipelinecolorattachmentdescriptorarray/2866392-objectatindexedsubscript?language=objc
func (t_ TileRenderPipelineColorAttachmentDescriptorArray) ObjectAtIndexedSubscript(attachmentIndex uint) TileRenderPipelineColorAttachmentDescriptor {
	rv := objc.Call[TileRenderPipelineColorAttachmentDescriptor](t_, objc.Sel("objectAtIndexedSubscript:"), attachmentIndex)
	return rv
}

// Sets the render pipeline state for a specified color attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltilerenderpipelinecolorattachmentdescriptorarray/2867636-setobject?language=objc
func (t_ TileRenderPipelineColorAttachmentDescriptorArray) SetObjectAtIndexedSubscript(attachment ITileRenderPipelineColorAttachmentDescriptor, attachmentIndex uint) {
	objc.Call[objc.Void](t_, objc.Sel("setObject:atIndexedSubscript:"), objc.Ptr(attachment), attachmentIndex)
}
