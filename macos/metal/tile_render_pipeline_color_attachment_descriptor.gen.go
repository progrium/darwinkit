// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TileRenderPipelineColorAttachmentDescriptor] class.
var TileRenderPipelineColorAttachmentDescriptorClass = _TileRenderPipelineColorAttachmentDescriptorClass{objc.GetClass("MTLTileRenderPipelineColorAttachmentDescriptor")}

type _TileRenderPipelineColorAttachmentDescriptorClass struct {
	objc.Class
}

// An interface definition for the [TileRenderPipelineColorAttachmentDescriptor] class.
type ITileRenderPipelineColorAttachmentDescriptor interface {
	objc.IObject
	PixelFormat() PixelFormat
	SetPixelFormat(value PixelFormat)
}

// A description of a tile-shading render pipeline’s color render target. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltilerenderpipelinecolorattachmentdescriptor?language=objc
type TileRenderPipelineColorAttachmentDescriptor struct {
	objc.Object
}

func TileRenderPipelineColorAttachmentDescriptorFrom(ptr unsafe.Pointer) TileRenderPipelineColorAttachmentDescriptor {
	return TileRenderPipelineColorAttachmentDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (tc _TileRenderPipelineColorAttachmentDescriptorClass) Alloc() TileRenderPipelineColorAttachmentDescriptor {
	rv := objc.Call[TileRenderPipelineColorAttachmentDescriptor](tc, objc.Sel("alloc"))
	return rv
}

func TileRenderPipelineColorAttachmentDescriptor_Alloc() TileRenderPipelineColorAttachmentDescriptor {
	return TileRenderPipelineColorAttachmentDescriptorClass.Alloc()
}

func (tc _TileRenderPipelineColorAttachmentDescriptorClass) New() TileRenderPipelineColorAttachmentDescriptor {
	rv := objc.Call[TileRenderPipelineColorAttachmentDescriptor](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTileRenderPipelineColorAttachmentDescriptor() TileRenderPipelineColorAttachmentDescriptor {
	return TileRenderPipelineColorAttachmentDescriptorClass.New()
}

func (t_ TileRenderPipelineColorAttachmentDescriptor) Init() TileRenderPipelineColorAttachmentDescriptor {
	rv := objc.Call[TileRenderPipelineColorAttachmentDescriptor](t_, objc.Sel("init"))
	return rv
}

// The pixel format associated with the tile shading render pipeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltilerenderpipelinecolorattachmentdescriptor/2866351-pixelformat?language=objc
func (t_ TileRenderPipelineColorAttachmentDescriptor) PixelFormat() PixelFormat {
	rv := objc.Call[PixelFormat](t_, objc.Sel("pixelFormat"))
	return rv
}

// The pixel format associated with the tile shading render pipeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltilerenderpipelinecolorattachmentdescriptor/2866351-pixelformat?language=objc
func (t_ TileRenderPipelineColorAttachmentDescriptor) SetPixelFormat(value PixelFormat) {
	objc.Call[objc.Void](t_, objc.Sel("setPixelFormat:"), value)
}
