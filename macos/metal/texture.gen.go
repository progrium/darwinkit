// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/iosurface"
	"github.com/progrium/macdriver/objc"
)

// A resource that holds formatted image data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture?language=objc
type PTexture interface {
	// optional
	ReplaceRegionMipmapLevelWithBytesBytesPerRow(region Region, level uint, pixelBytes unsafe.Pointer, bytesPerRow uint)
	HasReplaceRegionMipmapLevelWithBytesBytesPerRow() bool

	// optional
	NewRemoteTextureViewForDevice(device DeviceWrapper) PTexture
	HasNewRemoteTextureViewForDevice() bool

	// optional
	NewTextureViewWithPixelFormatTextureTypeLevelsSlicesSwizzle(pixelFormat PixelFormat, textureType TextureType, levelRange foundation.Range, sliceRange foundation.Range, swizzle TextureSwizzleChannels) PTexture
	HasNewTextureViewWithPixelFormatTextureTypeLevelsSlicesSwizzle() bool

	// optional
	GetBytesBytesPerRowBytesPerImageFromRegionMipmapLevelSlice(pixelBytes unsafe.Pointer, bytesPerRow uint, bytesPerImage uint, region Region, level uint, slice uint)
	HasGetBytesBytesPerRowBytesPerImageFromRegionMipmapLevelSlice() bool

	// optional
	NewSharedTextureHandle() ISharedTextureHandle
	HasNewSharedTextureHandle() bool

	// optional
	Width() uint
	HasWidth() bool

	// optional
	MipmapLevelCount() uint
	HasMipmapLevelCount() bool

	// optional
	ArrayLength() uint
	HasArrayLength() bool

	// optional
	RemoteStorageTexture() PTexture
	HasRemoteStorageTexture() bool

	// optional
	IsFramebufferOnly() bool
	HasIsFramebufferOnly() bool

	// optional
	Usage() TextureUsage
	HasUsage() bool

	// optional
	IosurfacePlane() uint
	HasIosurfacePlane() bool

	// optional
	ParentRelativeLevel() uint
	HasParentRelativeLevel() bool

	// optional
	IsShareable() bool
	HasIsShareable() bool

	// optional
	Buffer() PBuffer
	HasBuffer() bool

	// optional
	Swizzle() TextureSwizzleChannels
	HasSwizzle() bool

	// optional
	FirstMipmapInTail() uint
	HasFirstMipmapInTail() bool

	// optional
	ParentRelativeSlice() uint
	HasParentRelativeSlice() bool

	// optional
	SampleCount() uint
	HasSampleCount() bool

	// optional
	Height() uint
	HasHeight() bool

	// optional
	TextureType() TextureType
	HasTextureType() bool

	// optional
	BufferOffset() uint
	HasBufferOffset() bool

	// optional
	IsSparse() bool
	HasIsSparse() bool

	// optional
	AllowGPUOptimizedContents() bool
	HasAllowGPUOptimizedContents() bool

	// optional
	ParentTexture() PTexture
	HasParentTexture() bool

	// optional
	PixelFormat() PixelFormat
	HasPixelFormat() bool

	// optional
	Depth() uint
	HasDepth() bool

	// optional
	Iosurface() iosurface.Ref
	HasIosurface() bool

	// optional
	BufferBytesPerRow() uint
	HasBufferBytesPerRow() bool

	// optional
	TailSizeInBytes() uint
	HasTailSizeInBytes() bool
}

// A concrete type wrapper for the [PTexture] protocol.
type TextureWrapper struct {
	objc.Object
}

func (t_ TextureWrapper) HasReplaceRegionMipmapLevelWithBytesBytesPerRow() bool {
	return t_.RespondsToSelector(objc.Sel("replaceRegion:mipmapLevel:withBytes:bytesPerRow:"))
}

// Copies a block of pixels into a section of texture slice 0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1515464-replaceregion?language=objc
func (t_ TextureWrapper) ReplaceRegionMipmapLevelWithBytesBytesPerRow(region Region, level uint, pixelBytes unsafe.Pointer, bytesPerRow uint) {
	objc.Call[objc.Void](t_, objc.Sel("replaceRegion:mipmapLevel:withBytes:bytesPerRow:"), region, level, pixelBytes, bytesPerRow)
}

func (t_ TextureWrapper) HasNewRemoteTextureViewForDevice() bool {
	return t_.RespondsToSelector(objc.Sel("newRemoteTextureViewForDevice:"))
}

// Creates a remote texture view for another GPU in the same peer group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/2967449-newremotetextureviewfordevice?language=objc
func (t_ TextureWrapper) NewRemoteTextureViewForDevice(device PDevice) TextureWrapper {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[TextureWrapper](t_, objc.Sel("newRemoteTextureViewForDevice:"), po0)
	rv.Autorelease()
	return rv
}

func (t_ TextureWrapper) HasNewTextureViewWithPixelFormatTextureTypeLevelsSlicesSwizzle() bool {
	return t_.RespondsToSelector(objc.Sel("newTextureViewWithPixelFormat:textureType:levels:slices:swizzle:"))
}

// Creates a new view of the texture, reinterpreting a subset of its data using a different type, pixel format, and swizzle pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/3114303-newtextureviewwithpixelformat?language=objc
func (t_ TextureWrapper) NewTextureViewWithPixelFormatTextureTypeLevelsSlicesSwizzle(pixelFormat PixelFormat, textureType TextureType, levelRange foundation.Range, sliceRange foundation.Range, swizzle TextureSwizzleChannels) TextureWrapper {
	rv := objc.Call[TextureWrapper](t_, objc.Sel("newTextureViewWithPixelFormat:textureType:levels:slices:swizzle:"), pixelFormat, textureType, levelRange, sliceRange, swizzle)
	rv.Autorelease()
	return rv
}

func (t_ TextureWrapper) HasGetBytesBytesPerRowBytesPerImageFromRegionMipmapLevelSlice() bool {
	return t_.RespondsToSelector(objc.Sel("getBytes:bytesPerRow:bytesPerImage:fromRegion:mipmapLevel:slice:"))
}

// Copies pixel data from the texture to a buffer in system memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1516318-getbytes?language=objc
func (t_ TextureWrapper) GetBytesBytesPerRowBytesPerImageFromRegionMipmapLevelSlice(pixelBytes unsafe.Pointer, bytesPerRow uint, bytesPerImage uint, region Region, level uint, slice uint) {
	objc.Call[objc.Void](t_, objc.Sel("getBytes:bytesPerRow:bytesPerImage:fromRegion:mipmapLevel:slice:"), pixelBytes, bytesPerRow, bytesPerImage, region, level, slice)
}

func (t_ TextureWrapper) HasNewSharedTextureHandle() bool {
	return t_.RespondsToSelector(objc.Sel("newSharedTextureHandle"))
}

// Creates a new texture handle from a shareable texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/2981032-newsharedtexturehandle?language=objc
func (t_ TextureWrapper) NewSharedTextureHandle() SharedTextureHandle {
	rv := objc.Call[SharedTextureHandle](t_, objc.Sel("newSharedTextureHandle"))
	rv.Autorelease()
	return rv
}

func (t_ TextureWrapper) HasWidth() bool {
	return t_.RespondsToSelector(objc.Sel("width"))
}

// The width of the texture image for the base level mipmap, in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1515339-width?language=objc
func (t_ TextureWrapper) Width() uint {
	rv := objc.Call[uint](t_, objc.Sel("width"))
	return rv
}

func (t_ TextureWrapper) HasMipmapLevelCount() bool {
	return t_.RespondsToSelector(objc.Sel("mipmapLevelCount"))
}

// The number of mipmap levels in the texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1515677-mipmaplevelcount?language=objc
func (t_ TextureWrapper) MipmapLevelCount() uint {
	rv := objc.Call[uint](t_, objc.Sel("mipmapLevelCount"))
	return rv
}

func (t_ TextureWrapper) HasArrayLength() bool {
	return t_.RespondsToSelector(objc.Sel("arrayLength"))
}

// The number of slices in the texture array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1515382-arraylength?language=objc
func (t_ TextureWrapper) ArrayLength() uint {
	rv := objc.Call[uint](t_, objc.Sel("arrayLength"))
	return rv
}

func (t_ TextureWrapper) HasRemoteStorageTexture() bool {
	return t_.RespondsToSelector(objc.Sel("remoteStorageTexture"))
}

// The texture on another GPU that the texture was created from, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/2967451-remotestoragetexture?language=objc
func (t_ TextureWrapper) RemoteStorageTexture() TextureWrapper {
	rv := objc.Call[TextureWrapper](t_, objc.Sel("remoteStorageTexture"))
	return rv
}

func (t_ TextureWrapper) HasIsFramebufferOnly() bool {
	return t_.RespondsToSelector(objc.Sel("isFramebufferOnly"))
}

// A Boolean value that indicates whether the texture can only be used as a render target. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1515749-framebufferonly?language=objc
func (t_ TextureWrapper) IsFramebufferOnly() bool {
	rv := objc.Call[bool](t_, objc.Sel("isFramebufferOnly"))
	return rv
}

func (t_ TextureWrapper) HasUsage() bool {
	return t_.RespondsToSelector(objc.Sel("usage"))
}

// Options that determine how you can use the texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1515763-usage?language=objc
func (t_ TextureWrapper) Usage() TextureUsage {
	rv := objc.Call[TextureUsage](t_, objc.Sel("usage"))
	return rv
}

func (t_ TextureWrapper) HasIosurfacePlane() bool {
	return t_.RespondsToSelector(objc.Sel("iosurfacePlane"))
}

// The plane of the IOSurface to reference if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1515986-iosurfaceplane?language=objc
func (t_ TextureWrapper) IosurfacePlane() uint {
	rv := objc.Call[uint](t_, objc.Sel("iosurfacePlane"))
	return rv
}

func (t_ TextureWrapper) HasParentRelativeLevel() bool {
	return t_.RespondsToSelector(objc.Sel("parentRelativeLevel"))
}

// The base level of the parent texture that the texture was created from, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1516265-parentrelativelevel?language=objc
func (t_ TextureWrapper) ParentRelativeLevel() uint {
	rv := objc.Call[uint](t_, objc.Sel("parentRelativeLevel"))
	return rv
}

func (t_ TextureWrapper) HasIsShareable() bool {
	return t_.RespondsToSelector(objc.Sel("isShareable"))
}

// A Boolean indicating whether this texture can be shared with other processes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/2998889-shareable?language=objc
func (t_ TextureWrapper) IsShareable() bool {
	rv := objc.Call[bool](t_, objc.Sel("isShareable"))
	return rv
}

func (t_ TextureWrapper) HasBuffer() bool {
	return t_.RespondsToSelector(objc.Sel("buffer"))
}

// The source buffer that the texture was created from, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1619090-buffer?language=objc
func (t_ TextureWrapper) Buffer() BufferWrapper {
	rv := objc.Call[BufferWrapper](t_, objc.Sel("buffer"))
	return rv
}

func (t_ TextureWrapper) HasSwizzle() bool {
	return t_.RespondsToSelector(objc.Sel("swizzle"))
}

// The pattern that the GPU applies to pixels when you read or sample pixels from the texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/3114304-swizzle?language=objc
func (t_ TextureWrapper) Swizzle() TextureSwizzleChannels {
	rv := objc.Call[TextureSwizzleChannels](t_, objc.Sel("swizzle"))
	return rv
}

func (t_ TextureWrapper) HasFirstMipmapInTail() bool {
	return t_.RespondsToSelector(objc.Sel("firstMipmapInTail"))
}

// The index of the first mipmap in the tail. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/3043999-firstmipmapintail?language=objc
func (t_ TextureWrapper) FirstMipmapInTail() uint {
	rv := objc.Call[uint](t_, objc.Sel("firstMipmapInTail"))
	return rv
}

func (t_ TextureWrapper) HasParentRelativeSlice() bool {
	return t_.RespondsToSelector(objc.Sel("parentRelativeSlice"))
}

// The base slice of the parent texture that the texture was created from, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1516221-parentrelativeslice?language=objc
func (t_ TextureWrapper) ParentRelativeSlice() uint {
	rv := objc.Call[uint](t_, objc.Sel("parentRelativeSlice"))
	return rv
}

func (t_ TextureWrapper) HasSampleCount() bool {
	return t_.RespondsToSelector(objc.Sel("sampleCount"))
}

// The number of samples in each pixel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1515443-samplecount?language=objc
func (t_ TextureWrapper) SampleCount() uint {
	rv := objc.Call[uint](t_, objc.Sel("sampleCount"))
	return rv
}

func (t_ TextureWrapper) HasHeight() bool {
	return t_.RespondsToSelector(objc.Sel("height"))
}

// The height of the texture image for the base level mipmap, in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1515938-height?language=objc
func (t_ TextureWrapper) Height() uint {
	rv := objc.Call[uint](t_, objc.Sel("height"))
	return rv
}

func (t_ TextureWrapper) HasTextureType() bool {
	return t_.RespondsToSelector(objc.Sel("textureType"))
}

// The dimension and arrangement of the texture image data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1515517-texturetype?language=objc
func (t_ TextureWrapper) TextureType() TextureType {
	rv := objc.Call[TextureType](t_, objc.Sel("textureType"))
	return rv
}

func (t_ TextureWrapper) HasBufferOffset() bool {
	return t_.RespondsToSelector(objc.Sel("bufferOffset"))
}

// The offset in the source buffer where the texture's data comes from. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1619019-bufferoffset?language=objc
func (t_ TextureWrapper) BufferOffset() uint {
	rv := objc.Call[uint](t_, objc.Sel("bufferOffset"))
	return rv
}

func (t_ TextureWrapper) HasIsSparse() bool {
	return t_.RespondsToSelector(objc.Sel("isSparse"))
}

// A Boolean value that indicates whether this is a sparse texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/3153124-issparse?language=objc
func (t_ TextureWrapper) IsSparse() bool {
	rv := objc.Call[bool](t_, objc.Sel("isSparse"))
	return rv
}

func (t_ TextureWrapper) HasAllowGPUOptimizedContents() bool {
	return t_.RespondsToSelector(objc.Sel("allowGPUOptimizedContents"))
}

// A Boolean value indicating whether the GPU is allowed to adjust the contents of the texture to improve GPU performance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/2966640-allowgpuoptimizedcontents?language=objc
func (t_ TextureWrapper) AllowGPUOptimizedContents() bool {
	rv := objc.Call[bool](t_, objc.Sel("allowGPUOptimizedContents"))
	return rv
}

func (t_ TextureWrapper) HasParentTexture() bool {
	return t_.RespondsToSelector(objc.Sel("parentTexture"))
}

// The parent texture that the texture was created from, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1515372-parenttexture?language=objc
func (t_ TextureWrapper) ParentTexture() TextureWrapper {
	rv := objc.Call[TextureWrapper](t_, objc.Sel("parentTexture"))
	return rv
}

func (t_ TextureWrapper) HasPixelFormat() bool {
	return t_.RespondsToSelector(objc.Sel("pixelFormat"))
}

// The format of pixels in the texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1515344-pixelformat?language=objc
func (t_ TextureWrapper) PixelFormat() PixelFormat {
	rv := objc.Call[PixelFormat](t_, objc.Sel("pixelFormat"))
	return rv
}

func (t_ TextureWrapper) HasDepth() bool {
	return t_.RespondsToSelector(objc.Sel("depth"))
}

// The depth of the texture image for the base level mipmap, in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1515942-depth?language=objc
func (t_ TextureWrapper) Depth() uint {
	rv := objc.Call[uint](t_, objc.Sel("depth"))
	return rv
}

func (t_ TextureWrapper) HasIosurface() bool {
	return t_.RespondsToSelector(objc.Sel("iosurface"))
}

// A reference to the IOSurface the texture was created from, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1516104-iosurface?language=objc
func (t_ TextureWrapper) Iosurface() iosurface.Ref {
	rv := objc.Call[iosurface.Ref](t_, objc.Sel("iosurface"))
	return rv
}

func (t_ TextureWrapper) HasBufferBytesPerRow() bool {
	return t_.RespondsToSelector(objc.Sel("bufferBytesPerRow"))
}

// The number of bytes in each row of the texture’s source buffer, if applicable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/1619175-bufferbytesperrow?language=objc
func (t_ TextureWrapper) BufferBytesPerRow() uint {
	rv := objc.Call[uint](t_, objc.Sel("bufferBytesPerRow"))
	return rv
}

func (t_ TextureWrapper) HasTailSizeInBytes() bool {
	return t_.RespondsToSelector(objc.Sel("tailSizeInBytes"))
}

// The size of the sparse texture tail, in bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexture/3044002-tailsizeinbytes?language=objc
func (t_ TextureWrapper) TailSizeInBytes() uint {
	rv := objc.Call[uint](t_, objc.Sel("tailSizeInBytes"))
	return rv
}
