// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// An object used to encode data into an argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder?language=objc
type PArgumentEncoder interface {
	// optional
	SetSamplerStateAtIndex(sampler SamplerStateWrapper, index uint)
	HasSetSamplerStateAtIndex() bool

	// optional
	SetComputePipelineStateAtIndex(pipeline ComputePipelineStateWrapper, index uint)
	HasSetComputePipelineStateAtIndex() bool

	// optional
	SetRenderPipelineStatesWithRange(pipelines RenderPipelineStateWrapper, range_ foundation.Range)
	HasSetRenderPipelineStatesWithRange() bool

	// optional
	SetIndirectCommandBufferAtIndex(indirectCommandBuffer IndirectCommandBufferWrapper, index uint)
	HasSetIndirectCommandBufferAtIndex() bool

	// optional
	NewArgumentEncoderForBufferAtIndex(index uint) PArgumentEncoder
	HasNewArgumentEncoderForBufferAtIndex() bool

	// optional
	SetBuffersOffsetsWithRange(buffers BufferWrapper, offsets *uint, range_ foundation.Range)
	HasSetBuffersOffsetsWithRange() bool

	// optional
	SetVisibleFunctionTableAtIndex(visibleFunctionTable VisibleFunctionTableWrapper, index uint)
	HasSetVisibleFunctionTableAtIndex() bool

	// optional
	SetSamplerStatesWithRange(samplers SamplerStateWrapper, range_ foundation.Range)
	HasSetSamplerStatesWithRange() bool

	// optional
	SetAccelerationStructureAtIndex(accelerationStructure AccelerationStructureWrapper, index uint)
	HasSetAccelerationStructureAtIndex() bool

	// optional
	SetRenderPipelineStateAtIndex(pipeline RenderPipelineStateWrapper, index uint)
	HasSetRenderPipelineStateAtIndex() bool

	// optional
	SetArgumentBufferOffset(argumentBuffer BufferWrapper, offset uint)
	HasSetArgumentBufferOffset() bool

	// optional
	SetVisibleFunctionTablesWithRange(visibleFunctionTables VisibleFunctionTableWrapper, range_ foundation.Range)
	HasSetVisibleFunctionTablesWithRange() bool

	// optional
	SetTexturesWithRange(textures TextureWrapper, range_ foundation.Range)
	HasSetTexturesWithRange() bool

	// optional
	SetComputePipelineStatesWithRange(pipelines ComputePipelineStateWrapper, range_ foundation.Range)
	HasSetComputePipelineStatesWithRange() bool

	// optional
	SetIntersectionFunctionTableAtIndex(intersectionFunctionTable IntersectionFunctionTableWrapper, index uint)
	HasSetIntersectionFunctionTableAtIndex() bool

	// optional
	SetTextureAtIndex(texture TextureWrapper, index uint)
	HasSetTextureAtIndex() bool

	// optional
	SetIntersectionFunctionTablesWithRange(intersectionFunctionTables IntersectionFunctionTableWrapper, range_ foundation.Range)
	HasSetIntersectionFunctionTablesWithRange() bool

	// optional
	SetBufferOffsetAtIndex(buffer BufferWrapper, offset uint, index uint)
	HasSetBufferOffsetAtIndex() bool

	// optional
	SetIndirectCommandBuffersWithRange(buffers IndirectCommandBufferWrapper, range_ foundation.Range)
	HasSetIndirectCommandBuffersWithRange() bool

	// optional
	ConstantDataAtIndex(index uint) unsafe.Pointer
	HasConstantDataAtIndex() bool

	// optional
	Alignment() uint
	HasAlignment() bool

	// optional
	Device() PDevice
	HasDevice() bool

	// optional
	EncodedLength() uint
	HasEncodedLength() bool

	// optional
	SetLabel(value string)
	HasSetLabel() bool

	// optional
	Label() string
	HasLabel() bool
}

// A concrete type wrapper for the [PArgumentEncoder] protocol.
type ArgumentEncoderWrapper struct {
	objc.Object
}

func (a_ ArgumentEncoderWrapper) HasSetSamplerStateAtIndex() bool {
	return a_.RespondsToSelector(objc.Sel("setSamplerState:atIndex:"))
}

// Encodes a sampler into the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2915770-setsamplerstate?language=objc
func (a_ ArgumentEncoderWrapper) SetSamplerStateAtIndex(sampler PSamplerState, index uint) {
	po0 := objc.WrapAsProtocol("MTLSamplerState", sampler)
	objc.Call[objc.Void](a_, objc.Sel("setSamplerState:atIndex:"), po0, index)
}

func (a_ ArgumentEncoderWrapper) HasSetComputePipelineStateAtIndex() bool {
	return a_.RespondsToSelector(objc.Sel("setComputePipelineState:atIndex:"))
}

// Encodes a reference to a compute pipeline state into the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2966533-setcomputepipelinestate?language=objc
func (a_ ArgumentEncoderWrapper) SetComputePipelineStateAtIndex(pipeline PComputePipelineState, index uint) {
	po0 := objc.WrapAsProtocol("MTLComputePipelineState", pipeline)
	objc.Call[objc.Void](a_, objc.Sel("setComputePipelineState:atIndex:"), po0, index)
}

func (a_ ArgumentEncoderWrapper) HasSetRenderPipelineStatesWithRange() bool {
	return a_.RespondsToSelector(objc.Sel("setRenderPipelineStates:withRange:"))
}

// Encodes references to an array of render pipeline states into the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2966536-setrenderpipelinestates?language=objc
func (a_ ArgumentEncoderWrapper) SetRenderPipelineStatesWithRange(pipelines PRenderPipelineState, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLRenderPipelineState", pipelines)
	objc.Call[objc.Void](a_, objc.Sel("setRenderPipelineStates:withRange:"), po0, range_)
}

func (a_ ArgumentEncoderWrapper) HasSetIndirectCommandBufferAtIndex() bool {
	return a_.RespondsToSelector(objc.Sel("setIndirectCommandBuffer:atIndex:"))
}

// Encodes a reference to an indirect command buffer into the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2967410-setindirectcommandbuffer?language=objc
func (a_ ArgumentEncoderWrapper) SetIndirectCommandBufferAtIndex(indirectCommandBuffer PIndirectCommandBuffer, index uint) {
	po0 := objc.WrapAsProtocol("MTLIndirectCommandBuffer", indirectCommandBuffer)
	objc.Call[objc.Void](a_, objc.Sel("setIndirectCommandBuffer:atIndex:"), po0, index)
}

func (a_ ArgumentEncoderWrapper) HasNewArgumentEncoderForBufferAtIndex() bool {
	return a_.RespondsToSelector(objc.Sel("newArgumentEncoderForBufferAtIndex:"))
}

// Creates a new argument encoder for a nested argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2915783-newargumentencoderforbufferatind?language=objc
func (a_ ArgumentEncoderWrapper) NewArgumentEncoderForBufferAtIndex(index uint) ArgumentEncoderWrapper {
	rv := objc.Call[ArgumentEncoderWrapper](a_, objc.Sel("newArgumentEncoderForBufferAtIndex:"), index)
	rv.Autorelease()
	return rv
}

func (a_ ArgumentEncoderWrapper) HasSetBuffersOffsetsWithRange() bool {
	return a_.RespondsToSelector(objc.Sel("setBuffers:offsets:withRange:"))
}

// Encodes references to an array of buffers into the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2915772-setbuffers?language=objc
func (a_ ArgumentEncoderWrapper) SetBuffersOffsetsWithRange(buffers PBuffer, offsets *uint, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLBuffer", buffers)
	objc.Call[objc.Void](a_, objc.Sel("setBuffers:offsets:withRange:"), po0, offsets, range_)
}

func (a_ ArgumentEncoderWrapper) HasSetVisibleFunctionTableAtIndex() bool {
	return a_.RespondsToSelector(objc.Sel("setVisibleFunctionTable:atIndex:"))
}

// Encodes a reference to a function table into the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/3608168-setvisiblefunctiontable?language=objc
func (a_ ArgumentEncoderWrapper) SetVisibleFunctionTableAtIndex(visibleFunctionTable PVisibleFunctionTable, index uint) {
	po0 := objc.WrapAsProtocol("MTLVisibleFunctionTable", visibleFunctionTable)
	objc.Call[objc.Void](a_, objc.Sel("setVisibleFunctionTable:atIndex:"), po0, index)
}

func (a_ ArgumentEncoderWrapper) HasSetSamplerStatesWithRange() bool {
	return a_.RespondsToSelector(objc.Sel("setSamplerStates:withRange:"))
}

// Encodes an array of samplers into the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2915778-setsamplerstates?language=objc
func (a_ ArgumentEncoderWrapper) SetSamplerStatesWithRange(samplers PSamplerState, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLSamplerState", samplers)
	objc.Call[objc.Void](a_, objc.Sel("setSamplerStates:withRange:"), po0, range_)
}

func (a_ ArgumentEncoderWrapper) HasSetAccelerationStructureAtIndex() bool {
	return a_.RespondsToSelector(objc.Sel("setAccelerationStructure:atIndex:"))
}

// Encodes a reference to an acceleration structure into the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/3553920-setaccelerationstructure?language=objc
func (a_ ArgumentEncoderWrapper) SetAccelerationStructureAtIndex(accelerationStructure PAccelerationStructure, index uint) {
	po0 := objc.WrapAsProtocol("MTLAccelerationStructure", accelerationStructure)
	objc.Call[objc.Void](a_, objc.Sel("setAccelerationStructure:atIndex:"), po0, index)
}

func (a_ ArgumentEncoderWrapper) HasSetRenderPipelineStateAtIndex() bool {
	return a_.RespondsToSelector(objc.Sel("setRenderPipelineState:atIndex:"))
}

// Encodes a reference to a render pipeline state into the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2966535-setrenderpipelinestate?language=objc
func (a_ ArgumentEncoderWrapper) SetRenderPipelineStateAtIndex(pipeline PRenderPipelineState, index uint) {
	po0 := objc.WrapAsProtocol("MTLRenderPipelineState", pipeline)
	objc.Call[objc.Void](a_, objc.Sel("setRenderPipelineState:atIndex:"), po0, index)
}

func (a_ ArgumentEncoderWrapper) HasSetArgumentBufferOffset() bool {
	return a_.RespondsToSelector(objc.Sel("setArgumentBuffer:offset:"))
}

// Specifies the position in a buffer where the encoder writes argument data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2915777-setargumentbuffer?language=objc
func (a_ ArgumentEncoderWrapper) SetArgumentBufferOffset(argumentBuffer PBuffer, offset uint) {
	po0 := objc.WrapAsProtocol("MTLBuffer", argumentBuffer)
	objc.Call[objc.Void](a_, objc.Sel("setArgumentBuffer:offset:"), po0, offset)
}

func (a_ ArgumentEncoderWrapper) HasSetVisibleFunctionTablesWithRange() bool {
	return a_.RespondsToSelector(objc.Sel("setVisibleFunctionTables:withRange:"))
}

// Encodes references to an array of visible function tables into the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/3608169-setvisiblefunctiontables?language=objc
func (a_ ArgumentEncoderWrapper) SetVisibleFunctionTablesWithRange(visibleFunctionTables PVisibleFunctionTable, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLVisibleFunctionTable", visibleFunctionTables)
	objc.Call[objc.Void](a_, objc.Sel("setVisibleFunctionTables:withRange:"), po0, range_)
}

func (a_ ArgumentEncoderWrapper) HasSetTexturesWithRange() bool {
	return a_.RespondsToSelector(objc.Sel("setTextures:withRange:"))
}

// Encodes references to an array of textures into the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2915786-settextures?language=objc
func (a_ ArgumentEncoderWrapper) SetTexturesWithRange(textures PTexture, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLTexture", textures)
	objc.Call[objc.Void](a_, objc.Sel("setTextures:withRange:"), po0, range_)
}

func (a_ ArgumentEncoderWrapper) HasSetComputePipelineStatesWithRange() bool {
	return a_.RespondsToSelector(objc.Sel("setComputePipelineStates:withRange:"))
}

// Encodes references to an array of compute pipeline states into the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2966534-setcomputepipelinestates?language=objc
func (a_ ArgumentEncoderWrapper) SetComputePipelineStatesWithRange(pipelines PComputePipelineState, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLComputePipelineState", pipelines)
	objc.Call[objc.Void](a_, objc.Sel("setComputePipelineStates:withRange:"), po0, range_)
}

func (a_ ArgumentEncoderWrapper) HasSetIntersectionFunctionTableAtIndex() bool {
	return a_.RespondsToSelector(objc.Sel("setIntersectionFunctionTable:atIndex:"))
}

// Encodes a reference to a ray-tracing intersection function table into the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/3608166-setintersectionfunctiontable?language=objc
func (a_ ArgumentEncoderWrapper) SetIntersectionFunctionTableAtIndex(intersectionFunctionTable PIntersectionFunctionTable, index uint) {
	po0 := objc.WrapAsProtocol("MTLIntersectionFunctionTable", intersectionFunctionTable)
	objc.Call[objc.Void](a_, objc.Sel("setIntersectionFunctionTable:atIndex:"), po0, index)
}

func (a_ ArgumentEncoderWrapper) HasSetTextureAtIndex() bool {
	return a_.RespondsToSelector(objc.Sel("setTexture:atIndex:"))
}

// Encodes a reference to a texture into the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2915776-settexture?language=objc
func (a_ ArgumentEncoderWrapper) SetTextureAtIndex(texture PTexture, index uint) {
	po0 := objc.WrapAsProtocol("MTLTexture", texture)
	objc.Call[objc.Void](a_, objc.Sel("setTexture:atIndex:"), po0, index)
}

func (a_ ArgumentEncoderWrapper) HasSetIntersectionFunctionTablesWithRange() bool {
	return a_.RespondsToSelector(objc.Sel("setIntersectionFunctionTables:withRange:"))
}

// Encodes references to an array of intersection function tables into the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/3608167-setintersectionfunctiontables?language=objc
func (a_ ArgumentEncoderWrapper) SetIntersectionFunctionTablesWithRange(intersectionFunctionTables PIntersectionFunctionTable, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLIntersectionFunctionTable", intersectionFunctionTables)
	objc.Call[objc.Void](a_, objc.Sel("setIntersectionFunctionTables:withRange:"), po0, range_)
}

func (a_ ArgumentEncoderWrapper) HasSetBufferOffsetAtIndex() bool {
	return a_.RespondsToSelector(objc.Sel("setBuffer:offset:atIndex:"))
}

// Encodes a reference to a buffer into the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2915785-setbuffer?language=objc
func (a_ ArgumentEncoderWrapper) SetBufferOffsetAtIndex(buffer PBuffer, offset uint, index uint) {
	po0 := objc.WrapAsProtocol("MTLBuffer", buffer)
	objc.Call[objc.Void](a_, objc.Sel("setBuffer:offset:atIndex:"), po0, offset, index)
}

func (a_ ArgumentEncoderWrapper) HasSetIndirectCommandBuffersWithRange() bool {
	return a_.RespondsToSelector(objc.Sel("setIndirectCommandBuffers:withRange:"))
}

// Encodes an array of indirect command buffers into the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2967411-setindirectcommandbuffers?language=objc
func (a_ ArgumentEncoderWrapper) SetIndirectCommandBuffersWithRange(buffers PIndirectCommandBuffer, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLIndirectCommandBuffer", buffers)
	objc.Call[objc.Void](a_, objc.Sel("setIndirectCommandBuffers:withRange:"), po0, range_)
}

func (a_ ArgumentEncoderWrapper) HasConstantDataAtIndex() bool {
	return a_.RespondsToSelector(objc.Sel("constantDataAtIndex:"))
}

// Returns a pointer for an inlined constant data argument in the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2915771-constantdataatindex?language=objc
func (a_ ArgumentEncoderWrapper) ConstantDataAtIndex(index uint) unsafe.Pointer {
	rv := objc.Call[unsafe.Pointer](a_, objc.Sel("constantDataAtIndex:"), index)
	return rv
}

func (a_ ArgumentEncoderWrapper) HasAlignment() bool {
	return a_.RespondsToSelector(objc.Sel("alignment"))
}

// The alignment, in bytes, required for storing the encoded resources of an argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2915775-alignment?language=objc
func (a_ ArgumentEncoderWrapper) Alignment() uint {
	rv := objc.Call[uint](a_, objc.Sel("alignment"))
	return rv
}

func (a_ ArgumentEncoderWrapper) HasDevice() bool {
	return a_.RespondsToSelector(objc.Sel("device"))
}

// The device object that created the argument encoder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2915779-device?language=objc
func (a_ ArgumentEncoderWrapper) Device() DeviceWrapper {
	rv := objc.Call[DeviceWrapper](a_, objc.Sel("device"))
	return rv
}

func (a_ ArgumentEncoderWrapper) HasEncodedLength() bool {
	return a_.RespondsToSelector(objc.Sel("encodedLength"))
}

// The number of bytes required to store the encoded resources of an argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2915784-encodedlength?language=objc
func (a_ ArgumentEncoderWrapper) EncodedLength() uint {
	rv := objc.Call[uint](a_, objc.Sel("encodedLength"))
	return rv
}

func (a_ ArgumentEncoderWrapper) HasSetLabel() bool {
	return a_.RespondsToSelector(objc.Sel("setLabel:"))
}

// A string that identifies the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2915773-label?language=objc
func (a_ ArgumentEncoderWrapper) SetLabel(value string) {
	objc.Call[objc.Void](a_, objc.Sel("setLabel:"), value)
}

func (a_ ArgumentEncoderWrapper) HasLabel() bool {
	return a_.RespondsToSelector(objc.Sel("label"))
}

// A string that identifies the argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentencoder/2915773-label?language=objc
func (a_ ArgumentEncoderWrapper) Label() string {
	rv := objc.Call[string](a_, objc.Sel("label"))
	return rv
}
