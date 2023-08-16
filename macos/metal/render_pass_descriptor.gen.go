// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RenderPassDescriptor] class.
var RenderPassDescriptorClass = _RenderPassDescriptorClass{objc.GetClass("MTLRenderPassDescriptor")}

type _RenderPassDescriptorClass struct {
	objc.Class
}

// An interface definition for the [RenderPassDescriptor] class.
type IRenderPassDescriptor interface {
	objc.IObject
	SetSamplePositionsCount(positions *SamplePosition, count uint)
	GetSamplePositionsCount(positions *SamplePosition, count uint) uint
	VisibilityResultBuffer() BufferWrapper
	SetVisibilityResultBuffer(value PBuffer)
	SetVisibilityResultBufferObject(valueObject objc.IObject)
	RasterizationRateMap() RasterizationRateMapWrapper
	SetRasterizationRateMap(value PRasterizationRateMap)
	SetRasterizationRateMapObject(valueObject objc.IObject)
	ColorAttachments() RenderPassColorAttachmentDescriptorArray
	RenderTargetArrayLength() uint
	SetRenderTargetArrayLength(value uint)
	DepthAttachment() RenderPassDepthAttachmentDescriptor
	SetDepthAttachment(value IRenderPassDepthAttachmentDescriptor)
	TileWidth() uint
	SetTileWidth(value uint)
	TileHeight() uint
	SetTileHeight(value uint)
	SampleBufferAttachments() RenderPassSampleBufferAttachmentDescriptorArray
	DefaultRasterSampleCount() uint
	SetDefaultRasterSampleCount(value uint)
	RenderTargetWidth() uint
	SetRenderTargetWidth(value uint)
	ThreadgroupMemoryLength() uint
	SetThreadgroupMemoryLength(value uint)
	RenderTargetHeight() uint
	SetRenderTargetHeight(value uint)
	StencilAttachment() RenderPassStencilAttachmentDescriptor
	SetStencilAttachment(value IRenderPassStencilAttachmentDescriptor)
	ImageblockSampleLength() uint
	SetImageblockSampleLength(value uint)
}

// A group of render targets that hold the results of a render pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassdescriptor?language=objc
type RenderPassDescriptor struct {
	objc.Object
}

func RenderPassDescriptorFrom(ptr unsafe.Pointer) RenderPassDescriptor {
	return RenderPassDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (rc _RenderPassDescriptorClass) Alloc() RenderPassDescriptor {
	rv := objc.Call[RenderPassDescriptor](rc, objc.Sel("alloc"))
	return rv
}

func RenderPassDescriptor_Alloc() RenderPassDescriptor {
	return RenderPassDescriptorClass.Alloc()
}

func (rc _RenderPassDescriptorClass) New() RenderPassDescriptor {
	rv := objc.Call[RenderPassDescriptor](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRenderPassDescriptor() RenderPassDescriptor {
	return RenderPassDescriptorClass.New()
}

func (r_ RenderPassDescriptor) Init() RenderPassDescriptor {
	rv := objc.Call[RenderPassDescriptor](r_, objc.Sel("init"))
	return rv
}

// Sets the programmable sample positions for a render pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassdescriptor/2866524-setsamplepositions?language=objc
func (r_ RenderPassDescriptor) SetSamplePositionsCount(positions *SamplePosition, count uint) {
	objc.Call[objc.Void](r_, objc.Sel("setSamplePositions:count:"), positions, count)
}

// Retrieves the programmable sample positions set for a render pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassdescriptor/2866342-getsamplepositions?language=objc
func (r_ RenderPassDescriptor) GetSamplePositionsCount(positions *SamplePosition, count uint) uint {
	rv := objc.Call[uint](r_, objc.Sel("getSamplePositions:count:"), positions, count)
	return rv
}

// Creates a default render pass descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassdescriptor/1437979-renderpassdescriptor?language=objc
func (rc _RenderPassDescriptorClass) RenderPassDescriptor() RenderPassDescriptor {
	rv := objc.Call[RenderPassDescriptor](rc, objc.Sel("renderPassDescriptor"))
	return rv
}

// Creates a default render pass descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassdescriptor/1437979-renderpassdescriptor?language=objc
func RenderPassDescriptor_RenderPassDescriptor() RenderPassDescriptor {
	return RenderPassDescriptorClass.RenderPassDescriptor()
}

// A buffer where the GPU writes visibility test results when fragments pass depth and stencil tests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassdescriptor/1437942-visibilityresultbuffer?language=objc
func (r_ RenderPassDescriptor) VisibilityResultBuffer() BufferWrapper {
	rv := objc.Call[BufferWrapper](r_, objc.Sel("visibilityResultBuffer"))
	return rv
}

// A buffer where the GPU writes visibility test results when fragments pass depth and stencil tests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassdescriptor/1437942-visibilityresultbuffer?language=objc
func (r_ RenderPassDescriptor) SetVisibilityResultBuffer(value PBuffer) {
	po0 := objc.WrapAsProtocol("MTLBuffer", value)
	objc.Call[objc.Void](r_, objc.Sel("setVisibilityResultBuffer:"), po0)
}

// A buffer where the GPU writes visibility test results when fragments pass depth and stencil tests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassdescriptor/1437942-visibilityresultbuffer?language=objc
func (r_ RenderPassDescriptor) SetVisibilityResultBufferObject(valueObject objc.IObject) {
	objc.Call[objc.Void](r_, objc.Sel("setVisibilityResultBuffer:"), objc.Ptr(valueObject))
}

// The rasterization rate map to use when executing the render pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassdescriptor/3088853-rasterizationratemap?language=objc
func (r_ RenderPassDescriptor) RasterizationRateMap() RasterizationRateMapWrapper {
	rv := objc.Call[RasterizationRateMapWrapper](r_, objc.Sel("rasterizationRateMap"))
	return rv
}

// The rasterization rate map to use when executing the render pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassdescriptor/3088853-rasterizationratemap?language=objc
func (r_ RenderPassDescriptor) SetRasterizationRateMap(value PRasterizationRateMap) {
	po0 := objc.WrapAsProtocol("MTLRasterizationRateMap", value)
	objc.Call[objc.Void](r_, objc.Sel("setRasterizationRateMap:"), po0)
}

// The rasterization rate map to use when executing the render pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassdescriptor/3088853-rasterizationratemap?language=objc
func (r_ RenderPassDescriptor) SetRasterizationRateMapObject(valueObject objc.IObject) {
	objc.Call[objc.Void](r_, objc.Sel("setRasterizationRateMap:"), objc.Ptr(valueObject))
}

// An array of state information for attachments that store color data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassdescriptor/1437970-colorattachments?language=objc
func (r_ RenderPassDescriptor) ColorAttachments() RenderPassColorAttachmentDescriptorArray {
	rv := objc.Call[RenderPassColorAttachmentDescriptorArray](r_, objc.Sel("colorAttachments"))
	return rv
}

// The number of active layers that all attachments must have for layered rendering. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassdescriptor/1437975-rendertargetarraylength?language=objc
func (r_ RenderPassDescriptor) RenderTargetArrayLength() uint {
	rv := objc.Call[uint](r_, objc.Sel("renderTargetArrayLength"))
	return rv
}

// The number of active layers that all attachments must have for layered rendering. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassdescriptor/1437975-rendertargetarraylength?language=objc
func (r_ RenderPassDescriptor) SetRenderTargetArrayLength(value uint) {
	objc.Call[objc.Void](r_, objc.Sel("setRenderTargetArrayLength:"), value)
}

// State information for an attachment that stores depth data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassdescriptor/1437973-depthattachment?language=objc
func (r_ RenderPassDescriptor) DepthAttachment() RenderPassDepthAttachmentDescriptor {
	rv := objc.Call[RenderPassDepthAttachmentDescriptor](r_, objc.Sel("depthAttachment"))
	return rv
}

// State information for an attachment that stores depth data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassdescriptor/1437973-depthattachment?language=objc
func (r_ RenderPassDescriptor) SetDepthAttachment(value IRenderPassDepthAttachmentDescriptor) {
	objc.Call[objc.Void](r_, objc.Sel("setDepthAttachment:"), objc.Ptr(value))
}

// The tile width, in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassdescriptor/2866529-tilewidth?language=objc
func (r_ RenderPassDescriptor) TileWidth() uint {
	rv := objc.Call[uint](r_, objc.Sel("tileWidth"))
	return rv
}

// The tile width, in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassdescriptor/2866529-tilewidth?language=objc
func (r_ RenderPassDescriptor) SetTileWidth(value uint) {
	objc.Call[objc.Void](r_, objc.Sel("setTileWidth:"), value)
}

// The tile height, in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassdescriptor/2866356-tileheight?language=objc
func (r_ RenderPassDescriptor) TileHeight() uint {
	rv := objc.Call[uint](r_, objc.Sel("tileHeight"))
	return rv
}

// The tile height, in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassdescriptor/2866356-tileheight?language=objc
func (r_ RenderPassDescriptor) SetTileHeight(value uint) {
	objc.Call[objc.Void](r_, objc.Sel("setTileHeight:"), value)
}

// The array of sample buffers that the render pass can access. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassdescriptor/3564463-samplebufferattachments?language=objc
func (r_ RenderPassDescriptor) SampleBufferAttachments() RenderPassSampleBufferAttachmentDescriptorArray {
	rv := objc.Call[RenderPassSampleBufferAttachmentDescriptorArray](r_, objc.Sel("sampleBufferAttachments"))
	return rv
}

// The raster sample count for the render pass when the render pass doesn’t have explicit attachments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassdescriptor/2890260-defaultrastersamplecount?language=objc
func (r_ RenderPassDescriptor) DefaultRasterSampleCount() uint {
	rv := objc.Call[uint](r_, objc.Sel("defaultRasterSampleCount"))
	return rv
}

// The raster sample count for the render pass when the render pass doesn’t have explicit attachments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassdescriptor/2890260-defaultrastersamplecount?language=objc
func (r_ RenderPassDescriptor) SetDefaultRasterSampleCount(value uint) {
	objc.Call[objc.Void](r_, objc.Sel("setDefaultRasterSampleCount:"), value)
}

// The width, in pixels, to constrain the render target to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassdescriptor/2879276-rendertargetwidth?language=objc
func (r_ RenderPassDescriptor) RenderTargetWidth() uint {
	rv := objc.Call[uint](r_, objc.Sel("renderTargetWidth"))
	return rv
}

// The width, in pixels, to constrain the render target to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassdescriptor/2879276-rendertargetwidth?language=objc
func (r_ RenderPassDescriptor) SetRenderTargetWidth(value uint) {
	objc.Call[objc.Void](r_, objc.Sel("setRenderTargetWidth:"), value)
}

// The per-tile size, in bytes, of the persistent threadgroup memory allocation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassdescriptor/2866527-threadgroupmemorylength?language=objc
func (r_ RenderPassDescriptor) ThreadgroupMemoryLength() uint {
	rv := objc.Call[uint](r_, objc.Sel("threadgroupMemoryLength"))
	return rv
}

// The per-tile size, in bytes, of the persistent threadgroup memory allocation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassdescriptor/2866527-threadgroupmemorylength?language=objc
func (r_ RenderPassDescriptor) SetThreadgroupMemoryLength(value uint) {
	objc.Call[objc.Void](r_, objc.Sel("setThreadgroupMemoryLength:"), value)
}

// The height, in pixels, to constrain the render target to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassdescriptor/2879279-rendertargetheight?language=objc
func (r_ RenderPassDescriptor) RenderTargetHeight() uint {
	rv := objc.Call[uint](r_, objc.Sel("renderTargetHeight"))
	return rv
}

// The height, in pixels, to constrain the render target to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassdescriptor/2879279-rendertargetheight?language=objc
func (r_ RenderPassDescriptor) SetRenderTargetHeight(value uint) {
	objc.Call[objc.Void](r_, objc.Sel("setRenderTargetHeight:"), value)
}

// State information for an attachment that stores stencil data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassdescriptor/1437950-stencilattachment?language=objc
func (r_ RenderPassDescriptor) StencilAttachment() RenderPassStencilAttachmentDescriptor {
	rv := objc.Call[RenderPassStencilAttachmentDescriptor](r_, objc.Sel("stencilAttachment"))
	return rv
}

// State information for an attachment that stores stencil data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassdescriptor/1437950-stencilattachment?language=objc
func (r_ RenderPassDescriptor) SetStencilAttachment(value IRenderPassStencilAttachmentDescriptor) {
	objc.Call[objc.Void](r_, objc.Sel("setStencilAttachment:"), objc.Ptr(value))
}

// The per-sample size, in bytes, of the largest explicit imageblock layout in the render pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassdescriptor/2928281-imageblocksamplelength?language=objc
func (r_ RenderPassDescriptor) ImageblockSampleLength() uint {
	rv := objc.Call[uint](r_, objc.Sel("imageblockSampleLength"))
	return rv
}

// The per-sample size, in bytes, of the largest explicit imageblock layout in the render pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpassdescriptor/2928281-imageblocksamplelength?language=objc
func (r_ RenderPassDescriptor) SetImageblockSampleLength(value uint) {
	objc.Call[objc.Void](r_, objc.Sel("setImageblockSampleLength:"), value)
}
