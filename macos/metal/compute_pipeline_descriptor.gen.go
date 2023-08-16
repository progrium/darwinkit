// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ComputePipelineDescriptor] class.
var ComputePipelineDescriptorClass = _ComputePipelineDescriptorClass{objc.GetClass("MTLComputePipelineDescriptor")}

type _ComputePipelineDescriptorClass struct {
	objc.Class
}

// An interface definition for the [ComputePipelineDescriptor] class.
type IComputePipelineDescriptor interface {
	objc.IObject
	Reset()
	Buffers() PipelineBufferDescriptorArray
	SupportAddingBinaryFunctions() bool
	SetSupportAddingBinaryFunctions(value bool)
	MaxCallStackDepth() uint
	SetMaxCallStackDepth(value uint)
	LinkedFunctions() LinkedFunctions
	SetLinkedFunctions(value ILinkedFunctions)
	StageInputDescriptor() StageInputOutputDescriptor
	SetStageInputDescriptor(value IStageInputOutputDescriptor)
	MaxTotalThreadsPerThreadgroup() uint
	SetMaxTotalThreadsPerThreadgroup(value uint)
	ThreadGroupSizeIsMultipleOfThreadExecutionWidth() bool
	SetThreadGroupSizeIsMultipleOfThreadExecutionWidth(value bool)
	PreloadedLibraries() []DynamicLibraryWrapper
	SetPreloadedLibraries(value []PDynamicLibrary)
	Label() string
	SetLabel(value string)
	SupportIndirectCommandBuffers() bool
	SetSupportIndirectCommandBuffers(value bool)
	BinaryArchives() []BinaryArchiveWrapper
	SetBinaryArchives(value []PBinaryArchive)
	ComputeFunction() FunctionWrapper
	SetComputeFunction(value PFunction)
	SetComputeFunctionObject(valueObject objc.IObject)
}

// An object for customizing the compilation of a new compute pipeline state object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor?language=objc
type ComputePipelineDescriptor struct {
	objc.Object
}

func ComputePipelineDescriptorFrom(ptr unsafe.Pointer) ComputePipelineDescriptor {
	return ComputePipelineDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _ComputePipelineDescriptorClass) Alloc() ComputePipelineDescriptor {
	rv := objc.Call[ComputePipelineDescriptor](cc, objc.Sel("alloc"))
	return rv
}

func ComputePipelineDescriptor_Alloc() ComputePipelineDescriptor {
	return ComputePipelineDescriptorClass.Alloc()
}

func (cc _ComputePipelineDescriptorClass) New() ComputePipelineDescriptor {
	rv := objc.Call[ComputePipelineDescriptor](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewComputePipelineDescriptor() ComputePipelineDescriptor {
	return ComputePipelineDescriptorClass.New()
}

func (c_ ComputePipelineDescriptor) Init() ComputePipelineDescriptor {
	rv := objc.Call[ComputePipelineDescriptor](c_, objc.Sel("init"))
	return rv
}

// Resets all compute pipeline descriptor properties to their default values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/1414923-reset?language=objc
func (c_ ComputePipelineDescriptor) Reset() {
	objc.Call[objc.Void](c_, objc.Sel("reset"))
}

// The buffer mutability options for a compute pipeline's kernel function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/2879269-buffers?language=objc
func (c_ ComputePipelineDescriptor) Buffers() PipelineBufferDescriptorArray {
	rv := objc.Call[PipelineBufferDescriptorArray](c_, objc.Sel("buffers"))
	return rv
}

// A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to its callable functions list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/3578132-supportaddingbinaryfunctions?language=objc
func (c_ ComputePipelineDescriptor) SupportAddingBinaryFunctions() bool {
	rv := objc.Call[bool](c_, objc.Sel("supportAddingBinaryFunctions"))
	return rv
}

// A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to its callable functions list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/3578132-supportaddingbinaryfunctions?language=objc
func (c_ ComputePipelineDescriptor) SetSupportAddingBinaryFunctions(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setSupportAddingBinaryFunctions:"), value)
}

// The maximum function call depth from the top-most shader function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/3650409-maxcallstackdepth?language=objc
func (c_ ComputePipelineDescriptor) MaxCallStackDepth() uint {
	rv := objc.Call[uint](c_, objc.Sel("maxCallStackDepth"))
	return rv
}

// The maximum function call depth from the top-most shader function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/3650409-maxcallstackdepth?language=objc
func (c_ ComputePipelineDescriptor) SetMaxCallStackDepth(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setMaxCallStackDepth:"), value)
}

// The functions that you can specify as function arguments when encoding commands that use this pipeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/3553963-linkedfunctions?language=objc
func (c_ ComputePipelineDescriptor) LinkedFunctions() LinkedFunctions {
	rv := objc.Call[LinkedFunctions](c_, objc.Sel("linkedFunctions"))
	return rv
}

// The functions that you can specify as function arguments when encoding commands that use this pipeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/3553963-linkedfunctions?language=objc
func (c_ ComputePipelineDescriptor) SetLinkedFunctions(value ILinkedFunctions) {
	objc.Call[objc.Void](c_, objc.Sel("setLinkedFunctions:"), objc.Ptr(value))
}

// The organization of input and output data for the compute function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/2092373-stageinputdescriptor?language=objc
func (c_ ComputePipelineDescriptor) StageInputDescriptor() StageInputOutputDescriptor {
	rv := objc.Call[StageInputOutputDescriptor](c_, objc.Sel("stageInputDescriptor"))
	return rv
}

// The organization of input and output data for the compute function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/2092373-stageinputdescriptor?language=objc
func (c_ ComputePipelineDescriptor) SetStageInputDescriptor(value IStageInputOutputDescriptor) {
	objc.Call[objc.Void](c_, objc.Sel("setStageInputDescriptor:"), objc.Ptr(value))
}

// The maximum number of threads in a threadgroup that you can dispatch to the compute function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/2966560-maxtotalthreadsperthreadgroup?language=objc
func (c_ ComputePipelineDescriptor) MaxTotalThreadsPerThreadgroup() uint {
	rv := objc.Call[uint](c_, objc.Sel("maxTotalThreadsPerThreadgroup"))
	return rv
}

// The maximum number of threads in a threadgroup that you can dispatch to the compute function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/2966560-maxtotalthreadsperthreadgroup?language=objc
func (c_ ComputePipelineDescriptor) SetMaxTotalThreadsPerThreadgroup(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setMaxTotalThreadsPerThreadgroup:"), value)
}

// A Boolean value that indicates whether the threadgroup size must always be a multiple of the thread execution width. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/1414915-threadgroupsizeismultipleofthrea?language=objc
func (c_ ComputePipelineDescriptor) ThreadGroupSizeIsMultipleOfThreadExecutionWidth() bool {
	rv := objc.Call[bool](c_, objc.Sel("threadGroupSizeIsMultipleOfThreadExecutionWidth"))
	return rv
}

// A Boolean value that indicates whether the threadgroup size must always be a multiple of the thread execution width. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/1414915-threadgroupsizeismultipleofthrea?language=objc
func (c_ ComputePipelineDescriptor) SetThreadGroupSizeIsMultipleOfThreadExecutionWidth(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setThreadGroupSizeIsMultipleOfThreadExecutionWidth:"), value)
}

// The dynamic libraries that contain any precompiled shader functions you want to link. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/3801722-preloadedlibraries?language=objc
func (c_ ComputePipelineDescriptor) PreloadedLibraries() []DynamicLibraryWrapper {
	rv := objc.Call[[]DynamicLibraryWrapper](c_, objc.Sel("preloadedLibraries"))
	return rv
}

// The dynamic libraries that contain any precompiled shader functions you want to link. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/3801722-preloadedlibraries?language=objc
func (c_ ComputePipelineDescriptor) SetPreloadedLibraries(value []PDynamicLibrary) {
	objc.Call[objc.Void](c_, objc.Sel("setPreloadedLibraries:"), value)
}

// A string that identifies the object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/1414918-label?language=objc
func (c_ ComputePipelineDescriptor) Label() string {
	rv := objc.Call[string](c_, objc.Sel("label"))
	return rv
}

// A string that identifies the object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/1414918-label?language=objc
func (c_ ComputePipelineDescriptor) SetLabel(value string) {
	objc.Call[objc.Void](c_, objc.Sel("setLabel:"), value)
}

// A Boolean value that indicates whether you can encode commands that reference the pipeline state object into an indirect command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/2966561-supportindirectcommandbuffers?language=objc
func (c_ ComputePipelineDescriptor) SupportIndirectCommandBuffers() bool {
	rv := objc.Call[bool](c_, objc.Sel("supportIndirectCommandBuffers"))
	return rv
}

// A Boolean value that indicates whether you can encode commands that reference the pipeline state object into an indirect command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/2966561-supportindirectcommandbuffers?language=objc
func (c_ ComputePipelineDescriptor) SetSupportIndirectCommandBuffers(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setSupportIndirectCommandBuffers:"), value)
}

// The binary archives that contain any precompiled shader functions you want to link. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/3553961-binaryarchives?language=objc
func (c_ ComputePipelineDescriptor) BinaryArchives() []BinaryArchiveWrapper {
	rv := objc.Call[[]BinaryArchiveWrapper](c_, objc.Sel("binaryArchives"))
	return rv
}

// The binary archives that contain any precompiled shader functions you want to link. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/3553961-binaryarchives?language=objc
func (c_ ComputePipelineDescriptor) SetBinaryArchives(value []PBinaryArchive) {
	objc.Call[objc.Void](c_, objc.Sel("setBinaryArchives:"), value)
}

// The compute kernel the pipeline calls. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/1414917-computefunction?language=objc
func (c_ ComputePipelineDescriptor) ComputeFunction() FunctionWrapper {
	rv := objc.Call[FunctionWrapper](c_, objc.Sel("computeFunction"))
	return rv
}

// The compute kernel the pipeline calls. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/1414917-computefunction?language=objc
func (c_ ComputePipelineDescriptor) SetComputeFunction(value PFunction) {
	po0 := objc.WrapAsProtocol("MTLFunction", value)
	objc.Call[objc.Void](c_, objc.Sel("setComputeFunction:"), po0)
}

// The compute kernel the pipeline calls. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/1414917-computefunction?language=objc
func (c_ ComputePipelineDescriptor) SetComputeFunctionObject(valueObject objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setComputeFunction:"), objc.Ptr(valueObject))
}
