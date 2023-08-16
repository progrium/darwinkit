// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// An interface that represents a graphics pipeline configuration for a render pass, which the pass applies to the draw commands you encode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinestate?language=objc
type PRenderPipelineState interface {
	// optional
	NewIntersectionFunctionTableWithDescriptorStage(descriptor IntersectionFunctionTableDescriptor, stage RenderStages) PIntersectionFunctionTable
	HasNewIntersectionFunctionTableWithDescriptorStage() bool

	// optional
	NewRenderPipelineStateWithAdditionalBinaryFunctionsError(additionalBinaryFunctions RenderPipelineFunctionsDescriptor, error foundation.Error) PRenderPipelineState
	HasNewRenderPipelineStateWithAdditionalBinaryFunctionsError() bool

	// optional
	NewVisibleFunctionTableWithDescriptorStage(descriptor VisibleFunctionTableDescriptor, stage RenderStages) PVisibleFunctionTable
	HasNewVisibleFunctionTableWithDescriptorStage() bool

	// optional
	FunctionHandleWithFunctionStage(function FunctionWrapper, stage RenderStages) PFunctionHandle
	HasFunctionHandleWithFunctionStage() bool

	// optional
	ImageblockMemoryLengthForDimensions(imageblockDimensions Size) uint
	HasImageblockMemoryLengthForDimensions() bool

	// optional
	Device() PDevice
	HasDevice() bool

	// optional
	MaxTotalThreadsPerThreadgroup() uint
	HasMaxTotalThreadsPerThreadgroup() bool

	// optional
	ThreadgroupSizeMatchesTileSize() bool
	HasThreadgroupSizeMatchesTileSize() bool

	// optional
	Label() string
	HasLabel() bool

	// optional
	SupportIndirectCommandBuffers() bool
	HasSupportIndirectCommandBuffers() bool

	// optional
	ImageblockSampleLength() uint
	HasImageblockSampleLength() bool
}

// A concrete type wrapper for the [PRenderPipelineState] protocol.
type RenderPipelineStateWrapper struct {
	objc.Object
}

func (r_ RenderPipelineStateWrapper) HasNewIntersectionFunctionTableWithDescriptorStage() bool {
	return r_.RespondsToSelector(objc.Sel("newIntersectionFunctionTableWithDescriptor:stage:"))
}

// Creates a new intersection function table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinestate/3750582-newintersectionfunctiontablewith?language=objc
func (r_ RenderPipelineStateWrapper) NewIntersectionFunctionTableWithDescriptorStage(descriptor IIntersectionFunctionTableDescriptor, stage RenderStages) IntersectionFunctionTableWrapper {
	rv := objc.Call[IntersectionFunctionTableWrapper](r_, objc.Sel("newIntersectionFunctionTableWithDescriptor:stage:"), objc.Ptr(descriptor), stage)
	rv.Autorelease()
	return rv
}

func (r_ RenderPipelineStateWrapper) HasNewRenderPipelineStateWithAdditionalBinaryFunctionsError() bool {
	return r_.RespondsToSelector(objc.Sel("newRenderPipelineStateWithAdditionalBinaryFunctions:error:"))
}

// Creates a new pipeline state that’s a copy of the current pipeline state with additional shaders. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinestate/3750583-newrenderpipelinestatewithadditi?language=objc
func (r_ RenderPipelineStateWrapper) NewRenderPipelineStateWithAdditionalBinaryFunctionsError(additionalBinaryFunctions IRenderPipelineFunctionsDescriptor, error foundation.IError) RenderPipelineStateWrapper {
	rv := objc.Call[RenderPipelineStateWrapper](r_, objc.Sel("newRenderPipelineStateWithAdditionalBinaryFunctions:error:"), objc.Ptr(additionalBinaryFunctions), objc.Ptr(error))
	rv.Autorelease()
	return rv
}

func (r_ RenderPipelineStateWrapper) HasNewVisibleFunctionTableWithDescriptorStage() bool {
	return r_.RespondsToSelector(objc.Sel("newVisibleFunctionTableWithDescriptor:stage:"))
}

// Creates a new visible function table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinestate/3750584-newvisiblefunctiontablewithdescr?language=objc
func (r_ RenderPipelineStateWrapper) NewVisibleFunctionTableWithDescriptorStage(descriptor IVisibleFunctionTableDescriptor, stage RenderStages) VisibleFunctionTableWrapper {
	rv := objc.Call[VisibleFunctionTableWrapper](r_, objc.Sel("newVisibleFunctionTableWithDescriptor:stage:"), objc.Ptr(descriptor), stage)
	rv.Autorelease()
	return rv
}

func (r_ RenderPipelineStateWrapper) HasFunctionHandleWithFunctionStage() bool {
	return r_.RespondsToSelector(objc.Sel("functionHandleWithFunction:stage:"))
}

// Creates a function handle for a shader. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinestate/3750581-functionhandlewithfunction?language=objc
func (r_ RenderPipelineStateWrapper) FunctionHandleWithFunctionStage(function PFunction, stage RenderStages) FunctionHandleWrapper {
	po0 := objc.WrapAsProtocol("MTLFunction", function)
	rv := objc.Call[FunctionHandleWrapper](r_, objc.Sel("functionHandleWithFunction:stage:"), po0, stage)
	return rv
}

func (r_ RenderPipelineStateWrapper) HasImageblockMemoryLengthForDimensions() bool {
	return r_.RespondsToSelector(objc.Sel("imageblockMemoryLengthForDimensions:"))
}

// Returns the length of an imageblock’s memory for the specified imageblock dimensions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinestate/2928216-imageblockmemorylengthfordimensi?language=objc
func (r_ RenderPipelineStateWrapper) ImageblockMemoryLengthForDimensions(imageblockDimensions Size) uint {
	rv := objc.Call[uint](r_, objc.Sel("imageblockMemoryLengthForDimensions:"), imageblockDimensions)
	return rv
}

func (r_ RenderPipelineStateWrapper) HasDevice() bool {
	return r_.RespondsToSelector(objc.Sel("device"))
}

// The device instance that creates the pipeline state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinestate/1514706-device?language=objc
func (r_ RenderPipelineStateWrapper) Device() DeviceWrapper {
	rv := objc.Call[DeviceWrapper](r_, objc.Sel("device"))
	return rv
}

func (r_ RenderPipelineStateWrapper) HasMaxTotalThreadsPerThreadgroup() bool {
	return r_.RespondsToSelector(objc.Sel("maxTotalThreadsPerThreadgroup"))
}

// The largest number of threads the pipeline state can have in a single tile shader threadgroup. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinestate/2866352-maxtotalthreadsperthreadgroup?language=objc
func (r_ RenderPipelineStateWrapper) MaxTotalThreadsPerThreadgroup() uint {
	rv := objc.Call[uint](r_, objc.Sel("maxTotalThreadsPerThreadgroup"))
	return rv
}

func (r_ RenderPipelineStateWrapper) HasThreadgroupSizeMatchesTileSize() bool {
	return r_.RespondsToSelector(objc.Sel("threadgroupSizeMatchesTileSize"))
}

// A Boolean value that indicates whether the pipeline state needs a threadgroup’s size to equal a tile’s size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinestate/2866353-threadgroupsizematchestilesize?language=objc
func (r_ RenderPipelineStateWrapper) ThreadgroupSizeMatchesTileSize() bool {
	rv := objc.Call[bool](r_, objc.Sel("threadgroupSizeMatchesTileSize"))
	return rv
}

func (r_ RenderPipelineStateWrapper) HasLabel() bool {
	return r_.RespondsToSelector(objc.Sel("label"))
}

// A string that helps you identify the render pipeline state during debugging. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinestate/1514621-label?language=objc
func (r_ RenderPipelineStateWrapper) Label() string {
	rv := objc.Call[string](r_, objc.Sel("label"))
	return rv
}

func (r_ RenderPipelineStateWrapper) HasSupportIndirectCommandBuffers() bool {
	return r_.RespondsToSelector(objc.Sel("supportIndirectCommandBuffers"))
}

// A Boolean value that indicates whether the render pipeline supports encoding commands into an indirect command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinestate/2966639-supportindirectcommandbuffers?language=objc
func (r_ RenderPipelineStateWrapper) SupportIndirectCommandBuffers() bool {
	rv := objc.Call[bool](r_, objc.Sel("supportIndirectCommandBuffers"))
	return rv
}

func (r_ RenderPipelineStateWrapper) HasImageblockSampleLength() bool {
	return r_.RespondsToSelector(objc.Sel("imageblockSampleLength"))
}

// The memory size, in byes, of the render pipeline’s imageblock for a single sample. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinestate/2928215-imageblocksamplelength?language=objc
func (r_ RenderPipelineStateWrapper) ImageblockSampleLength() uint {
	rv := objc.Call[uint](r_, objc.Sel("imageblockSampleLength"))
	return rv
}
