// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RenderPipelineColorAttachmentDescriptor] class.
var RenderPipelineColorAttachmentDescriptorClass = _RenderPipelineColorAttachmentDescriptorClass{objc.GetClass("MTLRenderPipelineColorAttachmentDescriptor")}

type _RenderPipelineColorAttachmentDescriptorClass struct {
	objc.Class
}

// An interface definition for the [RenderPipelineColorAttachmentDescriptor] class.
type IRenderPipelineColorAttachmentDescriptor interface {
	objc.IObject
	DestinationRGBBlendFactor() BlendFactor
	SetDestinationRGBBlendFactor(value BlendFactor)
	SourceRGBBlendFactor() BlendFactor
	SetSourceRGBBlendFactor(value BlendFactor)
	RgbBlendOperation() BlendOperation
	SetRgbBlendOperation(value BlendOperation)
	AlphaBlendOperation() BlendOperation
	SetAlphaBlendOperation(value BlendOperation)
	WriteMask() ColorWriteMask
	SetWriteMask(value ColorWriteMask)
	IsBlendingEnabled() bool
	SetBlendingEnabled(value bool)
	PixelFormat() PixelFormat
	SetPixelFormat(value PixelFormat)
	SourceAlphaBlendFactor() BlendFactor
	SetSourceAlphaBlendFactor(value BlendFactor)
	DestinationAlphaBlendFactor() BlendFactor
	SetDestinationAlphaBlendFactor(value BlendFactor)
}

// A color render target that specifies the color configuration and color operations for a render pipeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinecolorattachmentdescriptor?language=objc
type RenderPipelineColorAttachmentDescriptor struct {
	objc.Object
}

func RenderPipelineColorAttachmentDescriptorFrom(ptr unsafe.Pointer) RenderPipelineColorAttachmentDescriptor {
	return RenderPipelineColorAttachmentDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (rc _RenderPipelineColorAttachmentDescriptorClass) Alloc() RenderPipelineColorAttachmentDescriptor {
	rv := objc.Call[RenderPipelineColorAttachmentDescriptor](rc, objc.Sel("alloc"))
	return rv
}

func RenderPipelineColorAttachmentDescriptor_Alloc() RenderPipelineColorAttachmentDescriptor {
	return RenderPipelineColorAttachmentDescriptorClass.Alloc()
}

func (rc _RenderPipelineColorAttachmentDescriptorClass) New() RenderPipelineColorAttachmentDescriptor {
	rv := objc.Call[RenderPipelineColorAttachmentDescriptor](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRenderPipelineColorAttachmentDescriptor() RenderPipelineColorAttachmentDescriptor {
	return RenderPipelineColorAttachmentDescriptorClass.New()
}

func (r_ RenderPipelineColorAttachmentDescriptor) Init() RenderPipelineColorAttachmentDescriptor {
	rv := objc.Call[RenderPipelineColorAttachmentDescriptor](r_, objc.Sel("init"))
	return rv
}

// The destination blend factor (DBF) used by the RGB blend operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinecolorattachmentdescriptor/1514626-destinationrgbblendfactor?language=objc
func (r_ RenderPipelineColorAttachmentDescriptor) DestinationRGBBlendFactor() BlendFactor {
	rv := objc.Call[BlendFactor](r_, objc.Sel("destinationRGBBlendFactor"))
	return rv
}

// The destination blend factor (DBF) used by the RGB blend operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinecolorattachmentdescriptor/1514626-destinationrgbblendfactor?language=objc
func (r_ RenderPipelineColorAttachmentDescriptor) SetDestinationRGBBlendFactor(value BlendFactor) {
	objc.Call[objc.Void](r_, objc.Sel("setDestinationRGBBlendFactor:"), value)
}

// The source blend factor (SBF) used by the RGB blend operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinecolorattachmentdescriptor/1514615-sourcergbblendfactor?language=objc
func (r_ RenderPipelineColorAttachmentDescriptor) SourceRGBBlendFactor() BlendFactor {
	rv := objc.Call[BlendFactor](r_, objc.Sel("sourceRGBBlendFactor"))
	return rv
}

// The source blend factor (SBF) used by the RGB blend operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinecolorattachmentdescriptor/1514615-sourcergbblendfactor?language=objc
func (r_ RenderPipelineColorAttachmentDescriptor) SetSourceRGBBlendFactor(value BlendFactor) {
	objc.Call[objc.Void](r_, objc.Sel("setSourceRGBBlendFactor:"), value)
}

// The blend operation assigned for the RGB data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinecolorattachmentdescriptor/1514659-rgbblendoperation?language=objc
func (r_ RenderPipelineColorAttachmentDescriptor) RgbBlendOperation() BlendOperation {
	rv := objc.Call[BlendOperation](r_, objc.Sel("rgbBlendOperation"))
	return rv
}

// The blend operation assigned for the RGB data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinecolorattachmentdescriptor/1514659-rgbblendoperation?language=objc
func (r_ RenderPipelineColorAttachmentDescriptor) SetRgbBlendOperation(value BlendOperation) {
	objc.Call[objc.Void](r_, objc.Sel("setRgbBlendOperation:"), value)
}

// The blend operation assigned for the alpha data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinecolorattachmentdescriptor/1514666-alphablendoperation?language=objc
func (r_ RenderPipelineColorAttachmentDescriptor) AlphaBlendOperation() BlendOperation {
	rv := objc.Call[BlendOperation](r_, objc.Sel("alphaBlendOperation"))
	return rv
}

// The blend operation assigned for the alpha data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinecolorattachmentdescriptor/1514666-alphablendoperation?language=objc
func (r_ RenderPipelineColorAttachmentDescriptor) SetAlphaBlendOperation(value BlendOperation) {
	objc.Call[objc.Void](r_, objc.Sel("setAlphaBlendOperation:"), value)
}

// A bitmask that restricts which color channels are written into the texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinecolorattachmentdescriptor/1514619-writemask?language=objc
func (r_ RenderPipelineColorAttachmentDescriptor) WriteMask() ColorWriteMask {
	rv := objc.Call[ColorWriteMask](r_, objc.Sel("writeMask"))
	return rv
}

// A bitmask that restricts which color channels are written into the texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinecolorattachmentdescriptor/1514619-writemask?language=objc
func (r_ RenderPipelineColorAttachmentDescriptor) SetWriteMask(value ColorWriteMask) {
	objc.Call[objc.Void](r_, objc.Sel("setWriteMask:"), value)
}

// A Boolean value that determines whether blending is enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinecolorattachmentdescriptor/1514642-blendingenabled?language=objc
func (r_ RenderPipelineColorAttachmentDescriptor) IsBlendingEnabled() bool {
	rv := objc.Call[bool](r_, objc.Sel("isBlendingEnabled"))
	return rv
}

// A Boolean value that determines whether blending is enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinecolorattachmentdescriptor/1514642-blendingenabled?language=objc
func (r_ RenderPipelineColorAttachmentDescriptor) SetBlendingEnabled(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setBlendingEnabled:"), value)
}

// The pixel format of the color attachment’s texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinecolorattachmentdescriptor/1514651-pixelformat?language=objc
func (r_ RenderPipelineColorAttachmentDescriptor) PixelFormat() PixelFormat {
	rv := objc.Call[PixelFormat](r_, objc.Sel("pixelFormat"))
	return rv
}

// The pixel format of the color attachment’s texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinecolorattachmentdescriptor/1514651-pixelformat?language=objc
func (r_ RenderPipelineColorAttachmentDescriptor) SetPixelFormat(value PixelFormat) {
	objc.Call[objc.Void](r_, objc.Sel("setPixelFormat:"), value)
}

// The source blend factor (SBF) used by the alpha blend operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinecolorattachmentdescriptor/1514660-sourcealphablendfactor?language=objc
func (r_ RenderPipelineColorAttachmentDescriptor) SourceAlphaBlendFactor() BlendFactor {
	rv := objc.Call[BlendFactor](r_, objc.Sel("sourceAlphaBlendFactor"))
	return rv
}

// The source blend factor (SBF) used by the alpha blend operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinecolorattachmentdescriptor/1514660-sourcealphablendfactor?language=objc
func (r_ RenderPipelineColorAttachmentDescriptor) SetSourceAlphaBlendFactor(value BlendFactor) {
	objc.Call[objc.Void](r_, objc.Sel("setSourceAlphaBlendFactor:"), value)
}

// The destination blend factor (DBF) used by the alpha blend operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinecolorattachmentdescriptor/1514657-destinationalphablendfactor?language=objc
func (r_ RenderPipelineColorAttachmentDescriptor) DestinationAlphaBlendFactor() BlendFactor {
	rv := objc.Call[BlendFactor](r_, objc.Sel("destinationAlphaBlendFactor"))
	return rv
}

// The destination blend factor (DBF) used by the alpha blend operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinecolorattachmentdescriptor/1514657-destinationalphablendfactor?language=objc
func (r_ RenderPipelineColorAttachmentDescriptor) SetDestinationAlphaBlendFactor(value BlendFactor) {
	objc.Call[objc.Void](r_, objc.Sel("setDestinationAlphaBlendFactor:"), value)
}
