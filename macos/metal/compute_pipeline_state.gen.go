// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// An object that contains a compiled compute pipeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinestate?language=objc
type PComputePipelineState interface {
	// optional
	NewIntersectionFunctionTableWithDescriptor(descriptor IntersectionFunctionTableDescriptor) PIntersectionFunctionTable
	HasNewIntersectionFunctionTableWithDescriptor() bool

	// optional
	NewVisibleFunctionTableWithDescriptor(descriptor VisibleFunctionTableDescriptor) PVisibleFunctionTable
	HasNewVisibleFunctionTableWithDescriptor() bool

	// optional
	FunctionHandleWithFunction(function FunctionWrapper) PFunctionHandle
	HasFunctionHandleWithFunction() bool

	// optional
	ImageblockMemoryLengthForDimensions(imageblockDimensions Size) uint
	HasImageblockMemoryLengthForDimensions() bool

	// optional
	NewComputePipelineStateWithAdditionalBinaryFunctionsError(functions []FunctionWrapper, error foundation.Error) PComputePipelineState
	HasNewComputePipelineStateWithAdditionalBinaryFunctionsError() bool

	// optional
	Device() PDevice
	HasDevice() bool

	// optional
	MaxTotalThreadsPerThreadgroup() uint
	HasMaxTotalThreadsPerThreadgroup() bool

	// optional
	ThreadExecutionWidth() uint
	HasThreadExecutionWidth() bool

	// optional
	StaticThreadgroupMemoryLength() uint
	HasStaticThreadgroupMemoryLength() bool

	// optional
	Label() string
	HasLabel() bool

	// optional
	SupportIndirectCommandBuffers() bool
	HasSupportIndirectCommandBuffers() bool
}

// A concrete type wrapper for the [PComputePipelineState] protocol.
type ComputePipelineStateWrapper struct {
	objc.Object
}

func (c_ ComputePipelineStateWrapper) HasNewIntersectionFunctionTableWithDescriptor() bool {
	return c_.RespondsToSelector(objc.Sel("newIntersectionFunctionTableWithDescriptor:"))
}

// Creates a new intersection function table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinestate/3580381-newintersectionfunctiontablewith?language=objc
func (c_ ComputePipelineStateWrapper) NewIntersectionFunctionTableWithDescriptor(descriptor IIntersectionFunctionTableDescriptor) IntersectionFunctionTableWrapper {
	rv := objc.Call[IntersectionFunctionTableWrapper](c_, objc.Sel("newIntersectionFunctionTableWithDescriptor:"), objc.Ptr(descriptor))
	rv.Autorelease()
	return rv
}

func (c_ ComputePipelineStateWrapper) HasNewVisibleFunctionTableWithDescriptor() bool {
	return c_.RespondsToSelector(objc.Sel("newVisibleFunctionTableWithDescriptor:"))
}

// Creates a new visible function table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinestate/3566543-newvisiblefunctiontablewithdescr?language=objc
func (c_ ComputePipelineStateWrapper) NewVisibleFunctionTableWithDescriptor(descriptor IVisibleFunctionTableDescriptor) VisibleFunctionTableWrapper {
	rv := objc.Call[VisibleFunctionTableWrapper](c_, objc.Sel("newVisibleFunctionTableWithDescriptor:"), objc.Ptr(descriptor))
	rv.Autorelease()
	return rv
}

func (c_ ComputePipelineStateWrapper) HasFunctionHandleWithFunction() bool {
	return c_.RespondsToSelector(objc.Sel("functionHandleWithFunction:"))
}

// Creates a function handle for a visible function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinestate/3553964-functionhandlewithfunction?language=objc
func (c_ ComputePipelineStateWrapper) FunctionHandleWithFunction(function PFunction) FunctionHandleWrapper {
	po0 := objc.WrapAsProtocol("MTLFunction", function)
	rv := objc.Call[FunctionHandleWrapper](c_, objc.Sel("functionHandleWithFunction:"), po0)
	return rv
}

func (c_ ComputePipelineStateWrapper) HasImageblockMemoryLengthForDimensions() bool {
	return c_.RespondsToSelector(objc.Sel("imageblockMemoryLengthForDimensions:"))
}

// Returns the imageblock memory length, in bytes, for the specified imageblock dimensions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinestate/2928195-imageblockmemorylengthfordimensi?language=objc
func (c_ ComputePipelineStateWrapper) ImageblockMemoryLengthForDimensions(imageblockDimensions Size) uint {
	rv := objc.Call[uint](c_, objc.Sel("imageblockMemoryLengthForDimensions:"), imageblockDimensions)
	return rv
}

func (c_ ComputePipelineStateWrapper) HasNewComputePipelineStateWithAdditionalBinaryFunctionsError() bool {
	return c_.RespondsToSelector(objc.Sel("newComputePipelineStateWithAdditionalBinaryFunctions:error:"))
}

// Creates a new pipeline state object with additional callable functions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinestate/3580380-newcomputepipelinestatewithaddit?language=objc
func (c_ ComputePipelineStateWrapper) NewComputePipelineStateWithAdditionalBinaryFunctionsError(functions []PFunction, error foundation.IError) ComputePipelineStateWrapper {
	rv := objc.Call[ComputePipelineStateWrapper](c_, objc.Sel("newComputePipelineStateWithAdditionalBinaryFunctions:error:"), functions, objc.Ptr(error))
	rv.Autorelease()
	return rv
}

func (c_ ComputePipelineStateWrapper) HasDevice() bool {
	return c_.RespondsToSelector(objc.Sel("device"))
}

// The Metal device object that created the pipeline state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinestate/1414925-device?language=objc
func (c_ ComputePipelineStateWrapper) Device() DeviceWrapper {
	rv := objc.Call[DeviceWrapper](c_, objc.Sel("device"))
	return rv
}

func (c_ ComputePipelineStateWrapper) HasMaxTotalThreadsPerThreadgroup() bool {
	return c_.RespondsToSelector(objc.Sel("maxTotalThreadsPerThreadgroup"))
}

// The maximum number of threads in a threadgroup that you can dispatch to the pipeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinestate/1414927-maxtotalthreadsperthreadgroup?language=objc
func (c_ ComputePipelineStateWrapper) MaxTotalThreadsPerThreadgroup() uint {
	rv := objc.Call[uint](c_, objc.Sel("maxTotalThreadsPerThreadgroup"))
	return rv
}

func (c_ ComputePipelineStateWrapper) HasThreadExecutionWidth() bool {
	return c_.RespondsToSelector(objc.Sel("threadExecutionWidth"))
}

// The number of threads that the GPU executes simultaneously. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinestate/1414911-threadexecutionwidth?language=objc
func (c_ ComputePipelineStateWrapper) ThreadExecutionWidth() uint {
	rv := objc.Call[uint](c_, objc.Sel("threadExecutionWidth"))
	return rv
}

func (c_ ComputePipelineStateWrapper) HasStaticThreadgroupMemoryLength() bool {
	return c_.RespondsToSelector(objc.Sel("staticThreadgroupMemoryLength"))
}

// The length, in bytes, of statically allocated threadgroup memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinestate/2877435-staticthreadgroupmemorylength?language=objc
func (c_ ComputePipelineStateWrapper) StaticThreadgroupMemoryLength() uint {
	rv := objc.Call[uint](c_, objc.Sel("staticThreadgroupMemoryLength"))
	return rv
}

func (c_ ComputePipelineStateWrapper) HasLabel() bool {
	return c_.RespondsToSelector(objc.Sel("label"))
}

// A string that identifies the compute pipeline state object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinestate/2880323-label?language=objc
func (c_ ComputePipelineStateWrapper) Label() string {
	rv := objc.Call[string](c_, objc.Sel("label"))
	return rv
}

func (c_ ComputePipelineStateWrapper) HasSupportIndirectCommandBuffers() bool {
	return c_.RespondsToSelector(objc.Sel("supportIndirectCommandBuffers"))
}

// A Boolean value that indicates whether the pipeline supports indirect command buffers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinestate/2966562-supportindirectcommandbuffers?language=objc
func (c_ ComputePipelineStateWrapper) SupportIndirectCommandBuffers() bool {
	rv := objc.Call[bool](c_, objc.Sel("supportIndirectCommandBuffers"))
	return rv
}
