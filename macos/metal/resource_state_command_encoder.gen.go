// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/objc"
)

// An encoder that encodes commands that modify resource configurations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourcestatecommandencoder?language=objc
type PResourceStateCommandEncoder interface {
	// optional
	UpdateTextureMappingsModeRegionsMipLevelsSlicesNumRegions(texture TextureWrapper, mode SparseTextureMappingMode, regions *Region, mipLevels *uint, slices *uint, numRegions uint)
	HasUpdateTextureMappingsModeRegionsMipLevelsSlicesNumRegions() bool

	// optional
	UpdateTextureMappingModeIndirectBufferIndirectBufferOffset(texture TextureWrapper, mode SparseTextureMappingMode, indirectBuffer BufferWrapper, indirectBufferOffset uint)
	HasUpdateTextureMappingModeIndirectBufferIndirectBufferOffset() bool

	// optional
	UpdateFence(fence FenceWrapper)
	HasUpdateFence() bool

	// optional
	WaitForFence(fence FenceWrapper)
	HasWaitForFence() bool
}

// A concrete type wrapper for the [PResourceStateCommandEncoder] protocol.
type ResourceStateCommandEncoderWrapper struct {
	objc.Object
}

func (r_ ResourceStateCommandEncoderWrapper) HasUpdateTextureMappingsModeRegionsMipLevelsSlicesNumRegions() bool {
	return r_.RespondsToSelector(objc.Sel("updateTextureMappings:mode:regions:mipLevels:slices:numRegions:"))
}

// Encodes a command to update memory mappings for multiple regions inside a texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourcestatecommandencoder/3043995-updatetexturemappings?language=objc
func (r_ ResourceStateCommandEncoderWrapper) UpdateTextureMappingsModeRegionsMipLevelsSlicesNumRegions(texture PTexture, mode SparseTextureMappingMode, regions *Region, mipLevels *uint, slices *uint, numRegions uint) {
	po0 := objc.WrapAsProtocol("MTLTexture", texture)
	objc.Call[objc.Void](r_, objc.Sel("updateTextureMappings:mode:regions:mipLevels:slices:numRegions:"), po0, mode, regions, mipLevels, slices, numRegions)
}

func (r_ ResourceStateCommandEncoderWrapper) HasUpdateTextureMappingModeIndirectBufferIndirectBufferOffset() bool {
	return r_.RespondsToSelector(objc.Sel("updateTextureMapping:mode:indirectBuffer:indirectBufferOffset:"))
}

// Encodes a command to update a texture’s memory mappings, specifying the parameters indirectly. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourcestatecommandencoder/3043993-updatetexturemapping?language=objc
func (r_ ResourceStateCommandEncoderWrapper) UpdateTextureMappingModeIndirectBufferIndirectBufferOffset(texture PTexture, mode SparseTextureMappingMode, indirectBuffer PBuffer, indirectBufferOffset uint) {
	po0 := objc.WrapAsProtocol("MTLTexture", texture)
	po2 := objc.WrapAsProtocol("MTLBuffer", indirectBuffer)
	objc.Call[objc.Void](r_, objc.Sel("updateTextureMapping:mode:indirectBuffer:indirectBufferOffset:"), po0, mode, po2, indirectBufferOffset)
}

func (r_ ResourceStateCommandEncoderWrapper) HasUpdateFence() bool {
	return r_.RespondsToSelector(objc.Sel("updateFence:"))
}

// Updates the given fence to capture all GPU work that the command encoder has enqueued up to this  point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourcestatecommandencoder/3114014-updatefence?language=objc
func (r_ ResourceStateCommandEncoderWrapper) UpdateFence(fence PFence) {
	po0 := objc.WrapAsProtocol("MTLFence", fence)
	objc.Call[objc.Void](r_, objc.Sel("updateFence:"), po0)
}

func (r_ ResourceStateCommandEncoderWrapper) HasWaitForFence() bool {
	return r_.RespondsToSelector(objc.Sel("waitForFence:"))
}

// Prevents the command encoder from enqueuing further GPU commands until the given fence is reached. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourcestatecommandencoder/3114015-waitforfence?language=objc
func (r_ ResourceStateCommandEncoderWrapper) WaitForFence(fence PFence) {
	po0 := objc.WrapAsProtocol("MTLFence", fence)
	objc.Call[objc.Void](r_, objc.Sel("waitForFence:"), po0)
}
