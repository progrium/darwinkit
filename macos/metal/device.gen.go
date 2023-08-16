// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/dispatch"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The main Metal interface to a GPU that apps use to draw graphics and run computations in parallel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice?language=objc
type PDevice interface {
	// optional
	SupportsFamily(gpuFamily GPUFamily) bool
	HasSupportsFamily() bool

	// optional
	MinimumLinearTextureAlignmentForPixelFormat(format PixelFormat) uint
	HasMinimumLinearTextureAlignmentForPixelFormat() bool

	// optional
	NewDynamicLibraryError(library LibraryWrapper, error foundation.Error) PDynamicLibrary
	HasNewDynamicLibraryError() bool

	// optional
	SampleTimestampsGpuTimestamp(cpuTimestamp *Timestamp, gpuTimestamp *Timestamp)
	HasSampleTimestampsGpuTimestamp() bool

	// optional
	NewTextureWithDescriptor(descriptor TextureDescriptor) PTexture
	HasNewTextureWithDescriptor() bool

	// optional
	NewBinaryArchiveWithDescriptorError(descriptor BinaryArchiveDescriptor, error foundation.Error) PBinaryArchive
	HasNewBinaryArchiveWithDescriptorError() bool

	// optional
	NewAccelerationStructureWithSize(size uint) PAccelerationStructure
	HasNewAccelerationStructureWithSize() bool

	// optional
	SupportsRasterizationRateMapWithLayerCount(layerCount uint) bool
	HasSupportsRasterizationRateMapWithLayerCount() bool

	// optional
	NewLibraryWithSourceOptionsError(source string, options CompileOptions, error foundation.Error) PLibrary
	HasNewLibraryWithSourceOptionsError() bool

	// optional
	NewArgumentEncoderWithArguments(arguments []ArgumentDescriptor) PArgumentEncoder
	HasNewArgumentEncoderWithArguments() bool

	// optional
	NewComputePipelineStateWithDescriptorOptionsCompletionHandler(descriptor ComputePipelineDescriptor, options PipelineOption, completionHandler NewComputePipelineStateWithReflectionCompletionHandler)
	HasNewComputePipelineStateWithDescriptorOptionsCompletionHandler() bool

	// optional
	NewCommandQueue() PCommandQueue
	HasNewCommandQueue() bool

	// optional
	SupportsTextureSampleCount(sampleCount uint) bool
	HasSupportsTextureSampleCount() bool

	// optional
	MinimumTextureBufferAlignmentForPixelFormat(format PixelFormat) uint
	HasMinimumTextureBufferAlignmentForPixelFormat() bool

	// optional
	NewFence() PFence
	HasNewFence() bool

	// optional
	GetDefaultSamplePositionsCount(positions *SamplePosition, count uint)
	HasGetDefaultSamplePositionsCount() bool

	// optional
	NewRasterizationRateMapWithDescriptor(descriptor RasterizationRateMapDescriptor) PRasterizationRateMap
	HasNewRasterizationRateMapWithDescriptor() bool

	// optional
	NewCounterSampleBufferWithDescriptorError(descriptor CounterSampleBufferDescriptor, error foundation.Error) PCounterSampleBuffer
	HasNewCounterSampleBufferWithDescriptorError() bool

	// optional
	NewCommandQueueWithMaxCommandBufferCount(maxCommandBufferCount uint) PCommandQueue
	HasNewCommandQueueWithMaxCommandBufferCount() bool

	// optional
	NewHeapWithDescriptor(descriptor HeapDescriptor) PHeap
	HasNewHeapWithDescriptor() bool

	// optional
	NewSharedTextureWithDescriptor(descriptor TextureDescriptor) PTexture
	HasNewSharedTextureWithDescriptor() bool

	// optional
	ConvertSparsePixelRegionsToTileRegionsWithTileSizeAlignmentModeNumRegions(pixelRegions *Region, tileRegions *Region, tileSize Size, mode SparseTextureRegionAlignmentMode, numRegions uint)
	HasConvertSparsePixelRegionsToTileRegionsWithTileSizeAlignmentModeNumRegions() bool

	// optional
	HeapBufferSizeAndAlignWithLengthOptions(length uint, options ResourceOptions) SizeAndAlign
	HasHeapBufferSizeAndAlignWithLengthOptions() bool

	// optional
	NewIndirectCommandBufferWithDescriptorMaxCommandCountOptions(descriptor IndirectCommandBufferDescriptor, maxCount uint, options ResourceOptions) PIndirectCommandBuffer
	HasNewIndirectCommandBufferWithDescriptorMaxCommandCountOptions() bool

	// optional
	HeapTextureSizeAndAlignWithDescriptor(desc TextureDescriptor) SizeAndAlign
	HasHeapTextureSizeAndAlignWithDescriptor() bool

	// optional
	NewLibraryWithDataError(data dispatch.Data, error foundation.Error) PLibrary
	HasNewLibraryWithDataError() bool

	// optional
	NewDefaultLibraryWithBundleError(bundle foundation.Bundle, error foundation.Error) PLibrary
	HasNewDefaultLibraryWithBundleError() bool

	// optional
	NewLibraryWithStitchedDescriptorCompletionHandler(descriptor StitchedLibraryDescriptor, completionHandler NewLibraryCompletionHandler)
	HasNewLibraryWithStitchedDescriptorCompletionHandler() bool

	// optional
	NewAccelerationStructureWithDescriptor(descriptor AccelerationStructureDescriptor) PAccelerationStructure
	HasNewAccelerationStructureWithDescriptor() bool

	// optional
	NewBufferWithBytesNoCopyLengthOptionsDeallocator(pointer unsafe.Pointer, length uint, options ResourceOptions, deallocator func(pointer unsafe.Pointer, length uint)) PBuffer
	HasNewBufferWithBytesNoCopyLengthOptionsDeallocator() bool

	// optional
	NewEvent() PEvent
	HasNewEvent() bool

	// optional
	NewBufferWithLengthOptions(length uint, options ResourceOptions) PBuffer
	HasNewBufferWithLengthOptions() bool

	// optional
	NewBufferWithBytesLengthOptions(pointer unsafe.Pointer, length uint, options ResourceOptions) PBuffer
	HasNewBufferWithBytesLengthOptions() bool

	// optional
	NewSharedEvent() PSharedEvent
	HasNewSharedEvent() bool

	// optional
	SparseTileSizeWithTextureTypePixelFormatSampleCount(textureType TextureType, pixelFormat PixelFormat, sampleCount uint) Size
	HasSparseTileSizeWithTextureTypePixelFormatSampleCount() bool

	// optional
	NewSamplerStateWithDescriptor(descriptor SamplerDescriptor) PSamplerState
	HasNewSamplerStateWithDescriptor() bool

	// optional
	NewDynamicLibraryWithURLError(url foundation.URL, error foundation.Error) PDynamicLibrary
	HasNewDynamicLibraryWithURLError() bool

	// optional
	NewRenderPipelineStateWithTileDescriptorOptionsCompletionHandler(descriptor TileRenderPipelineDescriptor, options PipelineOption, completionHandler NewRenderPipelineStateWithReflectionCompletionHandler)
	HasNewRenderPipelineStateWithTileDescriptorOptionsCompletionHandler() bool

	// optional
	SupportsCounterSampling(samplingPoint CounterSamplingPoint) bool
	HasSupportsCounterSampling() bool

	// optional
	SupportsVertexAmplificationCount(count uint) bool
	HasSupportsVertexAmplificationCount() bool

	// optional
	NewComputePipelineStateWithFunctionCompletionHandler(computeFunction FunctionWrapper, completionHandler NewComputePipelineStateCompletionHandler)
	HasNewComputePipelineStateWithFunctionCompletionHandler() bool

	// optional
	AccelerationStructureSizesWithDescriptor(descriptor AccelerationStructureDescriptor) AccelerationStructureSizes
	HasAccelerationStructureSizesWithDescriptor() bool

	// optional
	NewDefaultLibrary() PLibrary
	HasNewDefaultLibrary() bool

	// optional
	NewDepthStencilStateWithDescriptor(descriptor DepthStencilDescriptor) PDepthStencilState
	HasNewDepthStencilStateWithDescriptor() bool

	// optional
	ConvertSparseTileRegionsToPixelRegionsWithTileSizeNumRegions(tileRegions *Region, pixelRegions *Region, tileSize Size, numRegions uint)
	HasConvertSparseTileRegionsToPixelRegionsWithTileSizeNumRegions() bool

	// optional
	NewRenderPipelineStateWithDescriptorCompletionHandler(descriptor RenderPipelineDescriptor, completionHandler NewRenderPipelineStateCompletionHandler)
	HasNewRenderPipelineStateWithDescriptorCompletionHandler() bool

	// optional
	NewSharedTextureWithHandle(sharedHandle SharedTextureHandle) PTexture
	HasNewSharedTextureWithHandle() bool

	// optional
	NewLibraryWithURLError(url foundation.URL, error foundation.Error) PLibrary
	HasNewLibraryWithURLError() bool

	// optional
	NewSharedEventWithHandle(sharedEventHandle SharedEventHandle) PSharedEvent
	HasNewSharedEventWithHandle() bool

	// optional
	SupportsPrimitiveMotionBlur() bool
	HasSupportsPrimitiveMotionBlur() bool

	// optional
	SupportsPullModelInterpolation() bool
	HasSupportsPullModelInterpolation() bool

	// optional
	SparseTileSizeInBytes() uint
	HasSparseTileSizeInBytes() bool

	// optional
	ReadWriteTextureSupport() ReadWriteTextureTier
	HasReadWriteTextureSupport() bool

	// optional
	Supports32BitFloatFiltering() bool
	HasSupports32BitFloatFiltering() bool

	// optional
	PeerCount() uint32
	HasPeerCount() bool

	// optional
	AreRasterOrderGroupsSupported() bool
	HasAreRasterOrderGroupsSupported() bool

	// optional
	Name() string
	HasName() bool

	// optional
	Location() DeviceLocation
	HasLocation() bool

	// optional
	IsLowPower() bool
	HasIsLowPower() bool

	// optional
	SupportsBCTextureCompression() bool
	HasSupportsBCTextureCompression() bool

	// optional
	SupportsRaytracing() bool
	HasSupportsRaytracing() bool

	// optional
	SupportsShaderBarycentricCoordinates() bool
	HasSupportsShaderBarycentricCoordinates() bool

	// optional
	RecommendedMaxWorkingSetSize() uint64
	HasRecommendedMaxWorkingSetSize() bool

	// optional
	MaxThreadsPerThreadgroup() Size
	HasMaxThreadsPerThreadgroup() bool

	// optional
	PeerIndex() uint32
	HasPeerIndex() bool

	// optional
	SupportsFunctionPointersFromRender() bool
	HasSupportsFunctionPointersFromRender() bool

	// optional
	SupportsFunctionPointers() bool
	HasSupportsFunctionPointers() bool

	// optional
	CounterSets() []PCounterSet
	HasCounterSets() bool

	// optional
	IsRemovable() bool
	HasIsRemovable() bool

	// optional
	AreProgrammableSamplePositionsSupported() bool
	HasAreProgrammableSamplePositionsSupported() bool

	// optional
	IsDepth24Stencil8PixelFormatSupported() bool
	HasIsDepth24Stencil8PixelFormatSupported() bool

	// optional
	LocationNumber() uint
	HasLocationNumber() bool

	// optional
	RegistryID() uint64
	HasRegistryID() bool

	// optional
	Supports32BitMSAA() bool
	HasSupports32BitMSAA() bool

	// optional
	MaxBufferLength() uint
	HasMaxBufferLength() bool

	// optional
	SupportsRenderDynamicLibraries() bool
	HasSupportsRenderDynamicLibraries() bool

	// optional
	PeerGroupID() uint64
	HasPeerGroupID() bool

	// optional
	MaxArgumentBufferSamplerCount() uint
	HasMaxArgumentBufferSamplerCount() bool

	// optional
	IsHeadless() bool
	HasIsHeadless() bool

	// optional
	SupportsRaytracingFromRender() bool
	HasSupportsRaytracingFromRender() bool

	// optional
	SupportsQueryTextureLOD() bool
	HasSupportsQueryTextureLOD() bool

	// optional
	MaxTransferRate() uint64
	HasMaxTransferRate() bool

	// optional
	MaxThreadgroupMemoryLength() uint
	HasMaxThreadgroupMemoryLength() bool

	// optional
	HasUnifiedMemory() bool
	HasHasUnifiedMemory() bool

	// optional
	SupportsDynamicLibraries() bool
	HasSupportsDynamicLibraries() bool

	// optional
	CurrentAllocatedSize() uint
	HasCurrentAllocatedSize() bool

	// optional
	ArgumentBuffersSupport() ArgumentBuffersTier
	HasArgumentBuffersSupport() bool
}

// A concrete type wrapper for the [PDevice] protocol.
type DeviceWrapper struct {
	objc.Object
}

func (d_ DeviceWrapper) HasSupportsFamily() bool {
	return d_.RespondsToSelector(objc.Sel("supportsFamily:"))
}

// Returns a Boolean value that indicates whether the GPU device supports the feature set of a specific GPU family. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3143473-supportsfamily?language=objc
func (d_ DeviceWrapper) SupportsFamily(gpuFamily GPUFamily) bool {
	rv := objc.Call[bool](d_, objc.Sel("supportsFamily:"), gpuFamily)
	return rv
}

func (d_ DeviceWrapper) HasMinimumLinearTextureAlignmentForPixelFormat() bool {
	return d_.RespondsToSelector(objc.Sel("minimumLinearTextureAlignmentForPixelFormat:"))
}

// Returns the minimum alignment the GPU device requires to create a linear texture from a buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/2866126-minimumlineartexturealignmentfor?language=objc
func (d_ DeviceWrapper) MinimumLinearTextureAlignmentForPixelFormat(format PixelFormat) uint {
	rv := objc.Call[uint](d_, objc.Sel("minimumLinearTextureAlignmentForPixelFormat:"), format)
	return rv
}

func (d_ DeviceWrapper) HasNewDynamicLibraryError() bool {
	return d_.RespondsToSelector(objc.Sel("newDynamicLibrary:error:"))
}

// Creates a Metal dynamic library instance from a Metal library instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3553974-newdynamiclibrary?language=objc
func (d_ DeviceWrapper) NewDynamicLibraryError(library PLibrary, error foundation.IError) DynamicLibraryWrapper {
	po0 := objc.WrapAsProtocol("MTLLibrary", library)
	rv := objc.Call[DynamicLibraryWrapper](d_, objc.Sel("newDynamicLibrary:error:"), po0, objc.Ptr(error))
	rv.Autorelease()
	return rv
}

func (d_ DeviceWrapper) HasSampleTimestampsGpuTimestamp() bool {
	return d_.RespondsToSelector(objc.Sel("sampleTimestamps:gpuTimestamp:"))
}

// Captures and returns a CPU timestamp and a GPU timestamp from the same moment in time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3194378-sampletimestamps?language=objc
func (d_ DeviceWrapper) SampleTimestampsGpuTimestamp(cpuTimestamp *Timestamp, gpuTimestamp *Timestamp) {
	objc.Call[objc.Void](d_, objc.Sel("sampleTimestamps:gpuTimestamp:"), cpuTimestamp, gpuTimestamp)
}

func (d_ DeviceWrapper) HasNewTextureWithDescriptor() bool {
	return d_.RespondsToSelector(objc.Sel("newTextureWithDescriptor:"))
}

// Creates a new texture instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/1433425-newtexturewithdescriptor?language=objc
func (d_ DeviceWrapper) NewTextureWithDescriptor(descriptor ITextureDescriptor) TextureWrapper {
	rv := objc.Call[TextureWrapper](d_, objc.Sel("newTextureWithDescriptor:"), objc.Ptr(descriptor))
	rv.Autorelease()
	return rv
}

func (d_ DeviceWrapper) HasNewBinaryArchiveWithDescriptorError() bool {
	return d_.RespondsToSelector(objc.Sel("newBinaryArchiveWithDescriptor:error:"))
}

// Creates a Metal binary archive instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3553973-newbinaryarchivewithdescriptor?language=objc
func (d_ DeviceWrapper) NewBinaryArchiveWithDescriptorError(descriptor IBinaryArchiveDescriptor, error foundation.IError) BinaryArchiveWrapper {
	rv := objc.Call[BinaryArchiveWrapper](d_, objc.Sel("newBinaryArchiveWithDescriptor:error:"), objc.Ptr(descriptor), objc.Ptr(error))
	rv.Autorelease()
	return rv
}

func (d_ DeviceWrapper) HasNewAccelerationStructureWithSize() bool {
	return d_.RespondsToSelector(objc.Sel("newAccelerationStructureWithSize:"))
}

// Creates a new acceleration structure with a specific size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3553972-newaccelerationstructurewithsize?language=objc
func (d_ DeviceWrapper) NewAccelerationStructureWithSize(size uint) AccelerationStructureWrapper {
	rv := objc.Call[AccelerationStructureWrapper](d_, objc.Sel("newAccelerationStructureWithSize:"), size)
	rv.Autorelease()
	return rv
}

func (d_ DeviceWrapper) HasSupportsRasterizationRateMapWithLayerCount() bool {
	return d_.RespondsToSelector(objc.Sel("supportsRasterizationRateMapWithLayerCount:"))
}

// Returns a Boolean value that indicates whether the GPU can create a rasterization rate map with a specific number of layers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3131682-supportsrasterizationratemapwith?language=objc
func (d_ DeviceWrapper) SupportsRasterizationRateMapWithLayerCount(layerCount uint) bool {
	rv := objc.Call[bool](d_, objc.Sel("supportsRasterizationRateMapWithLayerCount:"), layerCount)
	return rv
}

func (d_ DeviceWrapper) HasNewLibraryWithSourceOptionsError() bool {
	return d_.RespondsToSelector(objc.Sel("newLibraryWithSource:options:error:"))
}

// Synchronously creates a Metal library instance by compiling the functions in a source string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/1433431-newlibrarywithsource?language=objc
func (d_ DeviceWrapper) NewLibraryWithSourceOptionsError(source string, options ICompileOptions, error foundation.IError) LibraryWrapper {
	rv := objc.Call[LibraryWrapper](d_, objc.Sel("newLibraryWithSource:options:error:"), source, objc.Ptr(options), objc.Ptr(error))
	rv.Autorelease()
	return rv
}

func (d_ DeviceWrapper) HasNewArgumentEncoderWithArguments() bool {
	return d_.RespondsToSelector(objc.Sel("newArgumentEncoderWithArguments:"))
}

// Creates a new argument encoder for an array of arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/2915744-newargumentencoderwitharguments?language=objc
func (d_ DeviceWrapper) NewArgumentEncoderWithArguments(arguments []IArgumentDescriptor) ArgumentEncoderWrapper {
	rv := objc.Call[ArgumentEncoderWrapper](d_, objc.Sel("newArgumentEncoderWithArguments:"), arguments)
	rv.Autorelease()
	return rv
}

func (d_ DeviceWrapper) HasNewComputePipelineStateWithDescriptorOptionsCompletionHandler() bool {
	return d_.RespondsToSelector(objc.Sel("newComputePipelineStateWithDescriptor:options:completionHandler:"))
}

// Asynchronously creates a compute pipeline state and reflection information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/1433403-newcomputepipelinestatewithdescr?language=objc
func (d_ DeviceWrapper) NewComputePipelineStateWithDescriptorOptionsCompletionHandler(descriptor IComputePipelineDescriptor, options PipelineOption, completionHandler NewComputePipelineStateWithReflectionCompletionHandler) {
	objc.Call[objc.Void](d_, objc.Sel("newComputePipelineStateWithDescriptor:options:completionHandler:"), objc.Ptr(descriptor), options, completionHandler)
}

func (d_ DeviceWrapper) HasNewCommandQueue() bool {
	return d_.RespondsToSelector(objc.Sel("newCommandQueue"))
}

// Creates a queue you use to submit rendering and computation commands to a GPU. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/1433388-newcommandqueue?language=objc
func (d_ DeviceWrapper) NewCommandQueue() CommandQueueWrapper {
	rv := objc.Call[CommandQueueWrapper](d_, objc.Sel("newCommandQueue"))
	rv.Autorelease()
	return rv
}

func (d_ DeviceWrapper) HasSupportsTextureSampleCount() bool {
	return d_.RespondsToSelector(objc.Sel("supportsTextureSampleCount:"))
}

// Returns a Boolean value that indicates whether the GPU can sample a texture with a specific number of sample points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/1433355-supportstexturesamplecount?language=objc
func (d_ DeviceWrapper) SupportsTextureSampleCount(sampleCount uint) bool {
	rv := objc.Call[bool](d_, objc.Sel("supportsTextureSampleCount:"), sampleCount)
	return rv
}

func (d_ DeviceWrapper) HasMinimumTextureBufferAlignmentForPixelFormat() bool {
	return d_.RespondsToSelector(objc.Sel("minimumTextureBufferAlignmentForPixelFormat:"))
}

// Returns the minimum alignment the GPU device requires to create a texture buffer from a buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/2966564-minimumtexturebufferalignmentfor?language=objc
func (d_ DeviceWrapper) MinimumTextureBufferAlignmentForPixelFormat(format PixelFormat) uint {
	rv := objc.Call[uint](d_, objc.Sel("minimumTextureBufferAlignmentForPixelFormat:"), format)
	return rv
}

func (d_ DeviceWrapper) HasNewFence() bool {
	return d_.RespondsToSelector(objc.Sel("newFence"))
}

// Creates a new memory fence instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/1649923-newfence?language=objc
func (d_ DeviceWrapper) NewFence() FenceWrapper {
	rv := objc.Call[FenceWrapper](d_, objc.Sel("newFence"))
	rv.Autorelease()
	return rv
}

func (d_ DeviceWrapper) HasGetDefaultSamplePositionsCount() bool {
	return d_.RespondsToSelector(objc.Sel("getDefaultSamplePositions:count:"))
}

// Retrieves the default sample positions for a specific sample count. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/2866120-getdefaultsamplepositions?language=objc
func (d_ DeviceWrapper) GetDefaultSamplePositionsCount(positions *SamplePosition, count uint) {
	objc.Call[objc.Void](d_, objc.Sel("getDefaultSamplePositions:count:"), positions, count)
}

func (d_ DeviceWrapper) HasNewRasterizationRateMapWithDescriptor() bool {
	return d_.RespondsToSelector(objc.Sel("newRasterizationRateMapWithDescriptor:"))
}

// Creates a rasterization rate map instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3131681-newrasterizationratemapwithdescr?language=objc
func (d_ DeviceWrapper) NewRasterizationRateMapWithDescriptor(descriptor IRasterizationRateMapDescriptor) RasterizationRateMapWrapper {
	rv := objc.Call[RasterizationRateMapWrapper](d_, objc.Sel("newRasterizationRateMapWithDescriptor:"), objc.Ptr(descriptor))
	rv.Autorelease()
	return rv
}

func (d_ DeviceWrapper) HasNewCounterSampleBufferWithDescriptorError() bool {
	return d_.RespondsToSelector(objc.Sel("newCounterSampleBufferWithDescriptor:error:"))
}

// Creates a counter sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3081741-newcountersamplebufferwithdescri?language=objc
func (d_ DeviceWrapper) NewCounterSampleBufferWithDescriptorError(descriptor ICounterSampleBufferDescriptor, error foundation.IError) CounterSampleBufferWrapper {
	rv := objc.Call[CounterSampleBufferWrapper](d_, objc.Sel("newCounterSampleBufferWithDescriptor:error:"), objc.Ptr(descriptor), objc.Ptr(error))
	rv.Autorelease()
	return rv
}

func (d_ DeviceWrapper) HasNewCommandQueueWithMaxCommandBufferCount() bool {
	return d_.RespondsToSelector(objc.Sel("newCommandQueueWithMaxCommandBufferCount:"))
}

// Creates a queue you use to submit rendering and computation commands to a GPU that has a fixed number of uncompleted command buffers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/1433433-newcommandqueuewithmaxcommandbuf?language=objc
func (d_ DeviceWrapper) NewCommandQueueWithMaxCommandBufferCount(maxCommandBufferCount uint) CommandQueueWrapper {
	rv := objc.Call[CommandQueueWrapper](d_, objc.Sel("newCommandQueueWithMaxCommandBufferCount:"), maxCommandBufferCount)
	rv.Autorelease()
	return rv
}

func (d_ DeviceWrapper) HasNewHeapWithDescriptor() bool {
	return d_.RespondsToSelector(objc.Sel("newHeapWithDescriptor:"))
}

// Creates a new GPU heap instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/1649928-newheapwithdescriptor?language=objc
func (d_ DeviceWrapper) NewHeapWithDescriptor(descriptor IHeapDescriptor) HeapWrapper {
	rv := objc.Call[HeapWrapper](d_, objc.Sel("newHeapWithDescriptor:"), objc.Ptr(descriptor))
	rv.Autorelease()
	return rv
}

func (d_ DeviceWrapper) HasNewSharedTextureWithDescriptor() bool {
	return d_.RespondsToSelector(objc.Sel("newSharedTextureWithDescriptor:"))
}

// Creates a texture that you can share across process boundaries. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/2967421-newsharedtexturewithdescriptor?language=objc
func (d_ DeviceWrapper) NewSharedTextureWithDescriptor(descriptor ITextureDescriptor) TextureWrapper {
	rv := objc.Call[TextureWrapper](d_, objc.Sel("newSharedTextureWithDescriptor:"), objc.Ptr(descriptor))
	rv.Autorelease()
	return rv
}

func (d_ DeviceWrapper) HasConvertSparsePixelRegionsToTileRegionsWithTileSizeAlignmentModeNumRegions() bool {
	return d_.RespondsToSelector(objc.Sel("convertSparsePixelRegions:toTileRegions:withTileSize:alignmentMode:numRegions:"))
}

// Converts a list of sparse pixel regions to tile regions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3197982-convertsparsepixelregions?language=objc
func (d_ DeviceWrapper) ConvertSparsePixelRegionsToTileRegionsWithTileSizeAlignmentModeNumRegions(pixelRegions *Region, tileRegions *Region, tileSize Size, mode SparseTextureRegionAlignmentMode, numRegions uint) {
	objc.Call[objc.Void](d_, objc.Sel("convertSparsePixelRegions:toTileRegions:withTileSize:alignmentMode:numRegions:"), pixelRegions, tileRegions, tileSize, mode, numRegions)
}

func (d_ DeviceWrapper) HasHeapBufferSizeAndAlignWithLengthOptions() bool {
	return d_.RespondsToSelector(objc.Sel("heapBufferSizeAndAlignWithLength:options:"))
}

// Returns the size and alignment, in bytes, of a buffer if you create it from a heap. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/1649922-heapbuffersizeandalignwithlength?language=objc
func (d_ DeviceWrapper) HeapBufferSizeAndAlignWithLengthOptions(length uint, options ResourceOptions) SizeAndAlign {
	rv := objc.Call[SizeAndAlign](d_, objc.Sel("heapBufferSizeAndAlignWithLength:options:"), length, options)
	return rv
}

func (d_ DeviceWrapper) HasNewIndirectCommandBufferWithDescriptorMaxCommandCountOptions() bool {
	return d_.RespondsToSelector(objc.Sel("newIndirectCommandBufferWithDescriptor:maxCommandCount:options:"))
}

// Creates an indirect command buffer instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/2967420-newindirectcommandbufferwithdesc?language=objc
func (d_ DeviceWrapper) NewIndirectCommandBufferWithDescriptorMaxCommandCountOptions(descriptor IIndirectCommandBufferDescriptor, maxCount uint, options ResourceOptions) IndirectCommandBufferWrapper {
	rv := objc.Call[IndirectCommandBufferWrapper](d_, objc.Sel("newIndirectCommandBufferWithDescriptor:maxCommandCount:options:"), objc.Ptr(descriptor), maxCount, options)
	rv.Autorelease()
	return rv
}

func (d_ DeviceWrapper) HasHeapTextureSizeAndAlignWithDescriptor() bool {
	return d_.RespondsToSelector(objc.Sel("heapTextureSizeAndAlignWithDescriptor:"))
}

// Returns the size and alignment, in bytes, of a texture if you create it from a heap. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/1649927-heaptexturesizeandalignwithdescr?language=objc
func (d_ DeviceWrapper) HeapTextureSizeAndAlignWithDescriptor(desc ITextureDescriptor) SizeAndAlign {
	rv := objc.Call[SizeAndAlign](d_, objc.Sel("heapTextureSizeAndAlignWithDescriptor:"), objc.Ptr(desc))
	return rv
}

func (d_ DeviceWrapper) HasNewLibraryWithDataError() bool {
	return d_.RespondsToSelector(objc.Sel("newLibraryWithData:error:"))
}

// Creates a Metal library instance that contains the functions in a precompiled Metal library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/1433391-newlibrarywithdata?language=objc
func (d_ DeviceWrapper) NewLibraryWithDataError(data dispatch.Data, error foundation.IError) LibraryWrapper {
	rv := objc.Call[LibraryWrapper](d_, objc.Sel("newLibraryWithData:error:"), data, objc.Ptr(error))
	rv.Autorelease()
	return rv
}

func (d_ DeviceWrapper) HasNewDefaultLibraryWithBundleError() bool {
	return d_.RespondsToSelector(objc.Sel("newDefaultLibraryWithBundle:error:"))
}

// Creates a Metal library instance that contains the functions in a bundle’s default Metal library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/2177054-newdefaultlibrarywithbundle?language=objc
func (d_ DeviceWrapper) NewDefaultLibraryWithBundleError(bundle foundation.IBundle, error foundation.IError) LibraryWrapper {
	rv := objc.Call[LibraryWrapper](d_, objc.Sel("newDefaultLibraryWithBundle:error:"), objc.Ptr(bundle), objc.Ptr(error))
	rv.Autorelease()
	return rv
}

func (d_ DeviceWrapper) HasNewLibraryWithStitchedDescriptorCompletionHandler() bool {
	return d_.RespondsToSelector(objc.Sel("newLibraryWithStitchedDescriptor:completionHandler:"))
}

// Asynchronously creates a Metal library from the function stitching graphs in a descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3857571-newlibrarywithstitcheddescriptor?language=objc
func (d_ DeviceWrapper) NewLibraryWithStitchedDescriptorCompletionHandler(descriptor IStitchedLibraryDescriptor, completionHandler NewLibraryCompletionHandler) {
	objc.Call[objc.Void](d_, objc.Sel("newLibraryWithStitchedDescriptor:completionHandler:"), objc.Ptr(descriptor), completionHandler)
}

func (d_ DeviceWrapper) HasNewAccelerationStructureWithDescriptor() bool {
	return d_.RespondsToSelector(objc.Sel("newAccelerationStructureWithDescriptor:"))
}

// Creates a new ray-tracing acceleration structure from a descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3553971-newaccelerationstructurewithdesc?language=objc
func (d_ DeviceWrapper) NewAccelerationStructureWithDescriptor(descriptor IAccelerationStructureDescriptor) AccelerationStructureWrapper {
	rv := objc.Call[AccelerationStructureWrapper](d_, objc.Sel("newAccelerationStructureWithDescriptor:"), objc.Ptr(descriptor))
	rv.Autorelease()
	return rv
}

func (d_ DeviceWrapper) HasNewBufferWithBytesNoCopyLengthOptionsDeallocator() bool {
	return d_.RespondsToSelector(objc.Sel("newBufferWithBytesNoCopy:length:options:deallocator:"))
}

// Creates a buffer that wraps an existing contiguous memory allocation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/1433382-newbufferwithbytesnocopy?language=objc
func (d_ DeviceWrapper) NewBufferWithBytesNoCopyLengthOptionsDeallocator(pointer unsafe.Pointer, length uint, options ResourceOptions, deallocator func(pointer unsafe.Pointer, length uint)) BufferWrapper {
	rv := objc.Call[BufferWrapper](d_, objc.Sel("newBufferWithBytesNoCopy:length:options:deallocator:"), pointer, length, options, deallocator)
	rv.Autorelease()
	return rv
}

func (d_ DeviceWrapper) HasNewEvent() bool {
	return d_.RespondsToSelector(objc.Sel("newEvent"))
}

// Creates a new event instance that you can use to synchronize commands and resources within the same GPU device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/2966565-newevent?language=objc
func (d_ DeviceWrapper) NewEvent() EventWrapper {
	rv := objc.Call[EventWrapper](d_, objc.Sel("newEvent"))
	rv.Autorelease()
	return rv
}

func (d_ DeviceWrapper) HasNewBufferWithLengthOptions() bool {
	return d_.RespondsToSelector(objc.Sel("newBufferWithLength:options:"))
}

// Creates a buffer the method clears with zero values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/1433375-newbufferwithlength?language=objc
func (d_ DeviceWrapper) NewBufferWithLengthOptions(length uint, options ResourceOptions) BufferWrapper {
	rv := objc.Call[BufferWrapper](d_, objc.Sel("newBufferWithLength:options:"), length, options)
	rv.Autorelease()
	return rv
}

func (d_ DeviceWrapper) HasNewBufferWithBytesLengthOptions() bool {
	return d_.RespondsToSelector(objc.Sel("newBufferWithBytes:length:options:"))
}

// Allocates a new buffer of a given length and initializes its contents by copying existing data into it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/1433429-newbufferwithbytes?language=objc
func (d_ DeviceWrapper) NewBufferWithBytesLengthOptions(pointer unsafe.Pointer, length uint, options ResourceOptions) BufferWrapper {
	rv := objc.Call[BufferWrapper](d_, objc.Sel("newBufferWithBytes:length:options:"), pointer, length, options)
	rv.Autorelease()
	return rv
}

func (d_ DeviceWrapper) HasNewSharedEvent() bool {
	return d_.RespondsToSelector(objc.Sel("newSharedEvent"))
}

// Creates a new shared event instance that you can use to synchronize commands and resources across different GPU devices. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/2966569-newsharedevent?language=objc
func (d_ DeviceWrapper) NewSharedEvent() SharedEventWrapper {
	rv := objc.Call[SharedEventWrapper](d_, objc.Sel("newSharedEvent"))
	rv.Autorelease()
	return rv
}

func (d_ DeviceWrapper) HasSparseTileSizeWithTextureTypePixelFormatSampleCount() bool {
	return d_.RespondsToSelector(objc.Sel("sparseTileSizeWithTextureType:pixelFormat:sampleCount:"))
}

// Returns the dimensions of a sparse tile for a texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3377856-sparsetilesizewithtexturetype?language=objc
func (d_ DeviceWrapper) SparseTileSizeWithTextureTypePixelFormatSampleCount(textureType TextureType, pixelFormat PixelFormat, sampleCount uint) Size {
	rv := objc.Call[Size](d_, objc.Sel("sparseTileSizeWithTextureType:pixelFormat:sampleCount:"), textureType, pixelFormat, sampleCount)
	return rv
}

func (d_ DeviceWrapper) HasNewSamplerStateWithDescriptor() bool {
	return d_.RespondsToSelector(objc.Sel("newSamplerStateWithDescriptor:"))
}

// Creates a sampler state instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/1433408-newsamplerstatewithdescriptor?language=objc
func (d_ DeviceWrapper) NewSamplerStateWithDescriptor(descriptor ISamplerDescriptor) SamplerStateWrapper {
	rv := objc.Call[SamplerStateWrapper](d_, objc.Sel("newSamplerStateWithDescriptor:"), objc.Ptr(descriptor))
	rv.Autorelease()
	return rv
}

func (d_ DeviceWrapper) HasNewDynamicLibraryWithURLError() bool {
	return d_.RespondsToSelector(objc.Sel("newDynamicLibraryWithURL:error:"))
}

// Creates a Metal dynamic library instance that contains the functions in the Metal library file at a URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3564458-newdynamiclibrarywithurl?language=objc
func (d_ DeviceWrapper) NewDynamicLibraryWithURLError(url foundation.IURL, error foundation.IError) DynamicLibraryWrapper {
	rv := objc.Call[DynamicLibraryWrapper](d_, objc.Sel("newDynamicLibraryWithURL:error:"), objc.Ptr(url), objc.Ptr(error))
	rv.Autorelease()
	return rv
}

func (d_ DeviceWrapper) HasNewRenderPipelineStateWithTileDescriptorOptionsCompletionHandler() bool {
	return d_.RespondsToSelector(objc.Sel("newRenderPipelineStateWithTileDescriptor:options:completionHandler:"))
}

// Asynchronously creates a tile shader’s render pipeline state and reflection information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/2866129-newrenderpipelinestatewithtilede?language=objc
func (d_ DeviceWrapper) NewRenderPipelineStateWithTileDescriptorOptionsCompletionHandler(descriptor ITileRenderPipelineDescriptor, options PipelineOption, completionHandler NewRenderPipelineStateWithReflectionCompletionHandler) {
	objc.Call[objc.Void](d_, objc.Sel("newRenderPipelineStateWithTileDescriptor:options:completionHandler:"), objc.Ptr(descriptor), options, completionHandler)
}

func (d_ DeviceWrapper) HasSupportsCounterSampling() bool {
	return d_.RespondsToSelector(objc.Sel("supportsCounterSampling:"))
}

// Returns a Boolean value that indicates whether you can read GPU counters at the specified command boundary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3564459-supportscountersampling?language=objc
func (d_ DeviceWrapper) SupportsCounterSampling(samplingPoint CounterSamplingPoint) bool {
	rv := objc.Call[bool](d_, objc.Sel("supportsCounterSampling:"), samplingPoint)
	return rv
}

func (d_ DeviceWrapper) HasSupportsVertexAmplificationCount() bool {
	return d_.RespondsToSelector(objc.Sel("supportsVertexAmplificationCount:"))
}

// Returns a Boolean value that indicates whether the GPU supports an amplification factor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3197984-supportsvertexamplificationcount?language=objc
func (d_ DeviceWrapper) SupportsVertexAmplificationCount(count uint) bool {
	rv := objc.Call[bool](d_, objc.Sel("supportsVertexAmplificationCount:"), count)
	return rv
}

func (d_ DeviceWrapper) HasNewComputePipelineStateWithFunctionCompletionHandler() bool {
	return d_.RespondsToSelector(objc.Sel("newComputePipelineStateWithFunction:completionHandler:"))
}

// Asynchronously creates a new compute pipeline state with a function instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/1433427-newcomputepipelinestatewithfunct?language=objc
func (d_ DeviceWrapper) NewComputePipelineStateWithFunctionCompletionHandler(computeFunction PFunction, completionHandler NewComputePipelineStateCompletionHandler) {
	po0 := objc.WrapAsProtocol("MTLFunction", computeFunction)
	objc.Call[objc.Void](d_, objc.Sel("newComputePipelineStateWithFunction:completionHandler:"), po0, completionHandler)
}

func (d_ DeviceWrapper) HasAccelerationStructureSizesWithDescriptor() bool {
	return d_.RespondsToSelector(objc.Sel("accelerationStructureSizesWithDescriptor:"))
}

// Returns the buffer sizes the GPU device needs to build, refit, and store an acceleration structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3553970-accelerationstructuresizeswithde?language=objc
func (d_ DeviceWrapper) AccelerationStructureSizesWithDescriptor(descriptor IAccelerationStructureDescriptor) AccelerationStructureSizes {
	rv := objc.Call[AccelerationStructureSizes](d_, objc.Sel("accelerationStructureSizesWithDescriptor:"), objc.Ptr(descriptor))
	return rv
}

func (d_ DeviceWrapper) HasNewDefaultLibrary() bool {
	return d_.RespondsToSelector(objc.Sel("newDefaultLibrary"))
}

// Creates a Metal library instance that contains the functions from your app’s default Metal library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/1433380-newdefaultlibrary?language=objc
func (d_ DeviceWrapper) NewDefaultLibrary() LibraryWrapper {
	rv := objc.Call[LibraryWrapper](d_, objc.Sel("newDefaultLibrary"))
	rv.Autorelease()
	return rv
}

func (d_ DeviceWrapper) HasNewDepthStencilStateWithDescriptor() bool {
	return d_.RespondsToSelector(objc.Sel("newDepthStencilStateWithDescriptor:"))
}

// Creates a depth-stencil state instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/1433412-newdepthstencilstatewithdescript?language=objc
func (d_ DeviceWrapper) NewDepthStencilStateWithDescriptor(descriptor IDepthStencilDescriptor) DepthStencilStateWrapper {
	rv := objc.Call[DepthStencilStateWrapper](d_, objc.Sel("newDepthStencilStateWithDescriptor:"), objc.Ptr(descriptor))
	rv.Autorelease()
	return rv
}

func (d_ DeviceWrapper) HasConvertSparseTileRegionsToPixelRegionsWithTileSizeNumRegions() bool {
	return d_.RespondsToSelector(objc.Sel("convertSparseTileRegions:toPixelRegions:withTileSize:numRegions:"))
}

// Converts a list of sparse tile regions to pixel regions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3197983-convertsparsetileregions?language=objc
func (d_ DeviceWrapper) ConvertSparseTileRegionsToPixelRegionsWithTileSizeNumRegions(tileRegions *Region, pixelRegions *Region, tileSize Size, numRegions uint) {
	objc.Call[objc.Void](d_, objc.Sel("convertSparseTileRegions:toPixelRegions:withTileSize:numRegions:"), tileRegions, pixelRegions, tileSize, numRegions)
}

func (d_ DeviceWrapper) HasNewRenderPipelineStateWithDescriptorCompletionHandler() bool {
	return d_.RespondsToSelector(objc.Sel("newRenderPipelineStateWithDescriptor:completionHandler:"))
}

// Asynchronously creates a render pipeline state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/1433363-newrenderpipelinestatewithdescri?language=objc
func (d_ DeviceWrapper) NewRenderPipelineStateWithDescriptorCompletionHandler(descriptor IRenderPipelineDescriptor, completionHandler NewRenderPipelineStateCompletionHandler) {
	objc.Call[objc.Void](d_, objc.Sel("newRenderPipelineStateWithDescriptor:completionHandler:"), objc.Ptr(descriptor), completionHandler)
}

func (d_ DeviceWrapper) HasNewSharedTextureWithHandle() bool {
	return d_.RespondsToSelector(objc.Sel("newSharedTextureWithHandle:"))
}

// Creates a texture that references a shared texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/2967422-newsharedtexturewithhandle?language=objc
func (d_ DeviceWrapper) NewSharedTextureWithHandle(sharedHandle ISharedTextureHandle) TextureWrapper {
	rv := objc.Call[TextureWrapper](d_, objc.Sel("newSharedTextureWithHandle:"), objc.Ptr(sharedHandle))
	rv.Autorelease()
	return rv
}

func (d_ DeviceWrapper) HasNewLibraryWithURLError() bool {
	return d_.RespondsToSelector(objc.Sel("newLibraryWithURL:error:"))
}

// Creates a Metal library instance that contains the functions in the Metal library file at a URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/2877432-newlibrarywithurl?language=objc
func (d_ DeviceWrapper) NewLibraryWithURLError(url foundation.IURL, error foundation.IError) LibraryWrapper {
	rv := objc.Call[LibraryWrapper](d_, objc.Sel("newLibraryWithURL:error:"), objc.Ptr(url), objc.Ptr(error))
	rv.Autorelease()
	return rv
}

func (d_ DeviceWrapper) HasNewSharedEventWithHandle() bool {
	return d_.RespondsToSelector(objc.Sel("newSharedEventWithHandle:"))
}

// Recreates a shared event from a handle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/2981024-newsharedeventwithhandle?language=objc
func (d_ DeviceWrapper) NewSharedEventWithHandle(sharedEventHandle ISharedEventHandle) SharedEventWrapper {
	rv := objc.Call[SharedEventWrapper](d_, objc.Sel("newSharedEventWithHandle:"), objc.Ptr(sharedEventHandle))
	rv.Autorelease()
	return rv
}

func (d_ DeviceWrapper) HasSupportsPrimitiveMotionBlur() bool {
	return d_.RespondsToSelector(objc.Sel("supportsPrimitiveMotionBlur"))
}

// A Boolean value that indicates whether the GPU device supports motion blur for ray tracing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3850519-supportsprimitivemotionblur?language=objc
func (d_ DeviceWrapper) SupportsPrimitiveMotionBlur() bool {
	rv := objc.Call[bool](d_, objc.Sel("supportsPrimitiveMotionBlur"))
	return rv
}

func (d_ DeviceWrapper) HasSupportsPullModelInterpolation() bool {
	return d_.RespondsToSelector(objc.Sel("supportsPullModelInterpolation"))
}

// A Boolean value that indicates whether the GPU can compute multiple interpolations of a fragment function’s input. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3564460-supportspullmodelinterpolation?language=objc
func (d_ DeviceWrapper) SupportsPullModelInterpolation() bool {
	rv := objc.Call[bool](d_, objc.Sel("supportsPullModelInterpolation"))
	return rv
}

func (d_ DeviceWrapper) HasSparseTileSizeInBytes() bool {
	return d_.RespondsToSelector(objc.Sel("sparseTileSizeInBytes"))
}

// Returns the size, in bytes, of a sparse tile the GPU device creates using a default page size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3377855-sparsetilesizeinbytes?language=objc
func (d_ DeviceWrapper) SparseTileSizeInBytes() uint {
	rv := objc.Call[uint](d_, objc.Sel("sparseTileSizeInBytes"))
	return rv
}

func (d_ DeviceWrapper) HasReadWriteTextureSupport() bool {
	return d_.RespondsToSelector(objc.Sel("readWriteTextureSupport"))
}

// The GPU device’s texture support tier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/2887289-readwritetexturesupport?language=objc
func (d_ DeviceWrapper) ReadWriteTextureSupport() ReadWriteTextureTier {
	rv := objc.Call[ReadWriteTextureTier](d_, objc.Sel("readWriteTextureSupport"))
	return rv
}

func (d_ DeviceWrapper) HasSupports32BitFloatFiltering() bool {
	return d_.RespondsToSelector(objc.Sel("supports32BitFloatFiltering"))
}

// A Boolean value that indicates whether the GPU can filter a texture with a 32-bit floating-point format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3566545-supports32bitfloatfiltering?language=objc
func (d_ DeviceWrapper) Supports32BitFloatFiltering() bool {
	rv := objc.Call[bool](d_, objc.Sel("supports32BitFloatFiltering"))
	return rv
}

func (d_ DeviceWrapper) HasPeerCount() bool {
	return d_.RespondsToSelector(objc.Sel("peerCount"))
}

// The total number of GPUs in the peer group, if applicable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/2967423-peercount?language=objc
func (d_ DeviceWrapper) PeerCount() uint32 {
	rv := objc.Call[uint32](d_, objc.Sel("peerCount"))
	return rv
}

func (d_ DeviceWrapper) HasAreRasterOrderGroupsSupported() bool {
	return d_.RespondsToSelector(objc.Sel("areRasterOrderGroupsSupported"))
}

// A Boolean value that indicates whether the GPU supports raster order groups. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/2887285-rasterordergroupssupported?language=objc
func (d_ DeviceWrapper) AreRasterOrderGroupsSupported() bool {
	rv := objc.Call[bool](d_, objc.Sel("areRasterOrderGroupsSupported"))
	return rv
}

func (d_ DeviceWrapper) HasName() bool {
	return d_.RespondsToSelector(objc.Sel("name"))
}

// The full name of the GPU device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/1433359-name?language=objc
func (d_ DeviceWrapper) Name() string {
	rv := objc.Call[string](d_, objc.Sel("name"))
	return rv
}

func (d_ DeviceWrapper) HasLocation() bool {
	return d_.RespondsToSelector(objc.Sel("location"))
}

// The physical location of the GPU relative to system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3114005-location?language=objc
func (d_ DeviceWrapper) Location() DeviceLocation {
	rv := objc.Call[DeviceLocation](d_, objc.Sel("location"))
	return rv
}

func (d_ DeviceWrapper) HasIsLowPower() bool {
	return d_.RespondsToSelector(objc.Sel("isLowPower"))
}

// A Boolean value that indicates whether the GPU lowers its performance to conserve energy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/1433409-lowpower?language=objc
func (d_ DeviceWrapper) IsLowPower() bool {
	rv := objc.Call[bool](d_, objc.Sel("isLowPower"))
	return rv
}

func (d_ DeviceWrapper) HasSupportsBCTextureCompression() bool {
	return d_.RespondsToSelector(objc.Sel("supportsBCTextureCompression"))
}

// A Boolean value that indicates whether you can use textures that use BC compression. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3566547-supportsbctexturecompression?language=objc
func (d_ DeviceWrapper) SupportsBCTextureCompression() bool {
	rv := objc.Call[bool](d_, objc.Sel("supportsBCTextureCompression"))
	return rv
}

func (d_ DeviceWrapper) HasSupportsRaytracing() bool {
	return d_.RespondsToSelector(objc.Sel("supportsRaytracing"))
}

// A Boolean value that indicates whether the GPU device supports ray tracing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3580382-supportsraytracing?language=objc
func (d_ DeviceWrapper) SupportsRaytracing() bool {
	rv := objc.Call[bool](d_, objc.Sel("supportsRaytracing"))
	return rv
}

func (d_ DeviceWrapper) HasSupportsShaderBarycentricCoordinates() bool {
	return d_.RespondsToSelector(objc.Sel("supportsShaderBarycentricCoordinates"))
}

// A Boolean value that indicates whether the GPU supports barycentric coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3325837-supportsshaderbarycentriccoordin?language=objc
func (d_ DeviceWrapper) SupportsShaderBarycentricCoordinates() bool {
	rv := objc.Call[bool](d_, objc.Sel("supportsShaderBarycentricCoordinates"))
	return rv
}

func (d_ DeviceWrapper) HasRecommendedMaxWorkingSetSize() bool {
	return d_.RespondsToSelector(objc.Sel("recommendedMaxWorkingSetSize"))
}

// An approximation of how much memory, in bytes, this GPU device can allocate without affecting its runtime performance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/2369280-recommendedmaxworkingsetsize?language=objc
func (d_ DeviceWrapper) RecommendedMaxWorkingSetSize() uint64 {
	rv := objc.Call[uint64](d_, objc.Sel("recommendedMaxWorkingSetSize"))
	return rv
}

func (d_ DeviceWrapper) HasMaxThreadsPerThreadgroup() bool {
	return d_.RespondsToSelector(objc.Sel("maxThreadsPerThreadgroup"))
}

// The maximum number of threads along each dimension of a threadgroup. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/1433393-maxthreadsperthreadgroup?language=objc
func (d_ DeviceWrapper) MaxThreadsPerThreadgroup() Size {
	rv := objc.Call[Size](d_, objc.Sel("maxThreadsPerThreadgroup"))
	return rv
}

func (d_ DeviceWrapper) HasPeerIndex() bool {
	return d_.RespondsToSelector(objc.Sel("peerIndex"))
}

// The unique identifier for a GPU in a peer group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/2967425-peerindex?language=objc
func (d_ DeviceWrapper) PeerIndex() uint32 {
	rv := objc.Call[uint32](d_, objc.Sel("peerIndex"))
	return rv
}

func (d_ DeviceWrapper) HasSupportsFunctionPointersFromRender() bool {
	return d_.RespondsToSelector(objc.Sel("supportsFunctionPointersFromRender"))
}

// A Boolean value that indicates whether the GPU device supports pointers to render functions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3750527-supportsfunctionpointersfromrend?language=objc
func (d_ DeviceWrapper) SupportsFunctionPointersFromRender() bool {
	rv := objc.Call[bool](d_, objc.Sel("supportsFunctionPointersFromRender"))
	return rv
}

func (d_ DeviceWrapper) HasSupportsFunctionPointers() bool {
	return d_.RespondsToSelector(objc.Sel("supportsFunctionPointers"))
}

// A Boolean value that indicates whether the GPU device supports pointers to compute kernel functions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3589487-supportsfunctionpointers?language=objc
func (d_ DeviceWrapper) SupportsFunctionPointers() bool {
	rv := objc.Call[bool](d_, objc.Sel("supportsFunctionPointers"))
	return rv
}

func (d_ DeviceWrapper) HasCounterSets() bool {
	return d_.RespondsToSelector(objc.Sel("counterSets"))
}

// The counter sets supported by the device object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3081739-countersets?language=objc
func (d_ DeviceWrapper) CounterSets() []CounterSetWrapper {
	rv := objc.Call[[]CounterSetWrapper](d_, objc.Sel("counterSets"))
	return rv
}

func (d_ DeviceWrapper) HasIsRemovable() bool {
	return d_.RespondsToSelector(objc.Sel("isRemovable"))
}

// A Boolean value that indicates whether the GPU is removable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/2889851-removable?language=objc
func (d_ DeviceWrapper) IsRemovable() bool {
	rv := objc.Call[bool](d_, objc.Sel("isRemovable"))
	return rv
}

func (d_ DeviceWrapper) HasAreProgrammableSamplePositionsSupported() bool {
	return d_.RespondsToSelector(objc.Sel("areProgrammableSamplePositionsSupported"))
}

// A Boolean value that indicates whether the GPU supports programmable sample positions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/2866117-programmablesamplepositionssuppo?language=objc
func (d_ DeviceWrapper) AreProgrammableSamplePositionsSupported() bool {
	rv := objc.Call[bool](d_, objc.Sel("areProgrammableSamplePositionsSupported"))
	return rv
}

func (d_ DeviceWrapper) HasIsDepth24Stencil8PixelFormatSupported() bool {
	return d_.RespondsToSelector(objc.Sel("isDepth24Stencil8PixelFormatSupported"))
}

// A Boolean value that indicates whether a device supports a packed depth-and-stencil pixel format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/1433371-depth24stencil8pixelformatsuppor?language=objc
func (d_ DeviceWrapper) IsDepth24Stencil8PixelFormatSupported() bool {
	rv := objc.Call[bool](d_, objc.Sel("isDepth24Stencil8PixelFormatSupported"))
	return rv
}

func (d_ DeviceWrapper) HasLocationNumber() bool {
	return d_.RespondsToSelector(objc.Sel("locationNumber"))
}

// A specific GPU position based on its general location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3114006-locationnumber?language=objc
func (d_ DeviceWrapper) LocationNumber() uint {
	rv := objc.Call[uint](d_, objc.Sel("locationNumber"))
	return rv
}

func (d_ DeviceWrapper) HasRegistryID() bool {
	return d_.RespondsToSelector(objc.Sel("registryID"))
}

// The GPU device’s registry identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/2915737-registryid?language=objc
func (d_ DeviceWrapper) RegistryID() uint64 {
	rv := objc.Call[uint64](d_, objc.Sel("registryID"))
	return rv
}

func (d_ DeviceWrapper) HasSupports32BitMSAA() bool {
	return d_.RespondsToSelector(objc.Sel("supports32BitMSAA"))
}

// A Boolean value that indicates whether the GPU can allocate 32-bit integer texture formats and resolve to 32-bit floating-point texture formats. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3566546-supports32bitmsaa?language=objc
func (d_ DeviceWrapper) Supports32BitMSAA() bool {
	rv := objc.Call[bool](d_, objc.Sel("supports32BitMSAA"))
	return rv
}

func (d_ DeviceWrapper) HasMaxBufferLength() bool {
	return d_.RespondsToSelector(objc.Sel("maxBufferLength"))
}

// The largest amount of memory, in bytes, that a GPU device can allocate to a buffer instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/2966563-maxbufferlength?language=objc
func (d_ DeviceWrapper) MaxBufferLength() uint {
	rv := objc.Call[uint](d_, objc.Sel("maxBufferLength"))
	return rv
}

func (d_ DeviceWrapper) HasSupportsRenderDynamicLibraries() bool {
	return d_.RespondsToSelector(objc.Sel("supportsRenderDynamicLibraries"))
}

// A Boolean value that indicates whether the GPU device can create and use dynamic libraries in render pipelines. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3816767-supportsrenderdynamiclibraries?language=objc
func (d_ DeviceWrapper) SupportsRenderDynamicLibraries() bool {
	rv := objc.Call[bool](d_, objc.Sel("supportsRenderDynamicLibraries"))
	return rv
}

func (d_ DeviceWrapper) HasPeerGroupID() bool {
	return d_.RespondsToSelector(objc.Sel("peerGroupID"))
}

// The peer group ID the GPU belongs to, if applicable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/2967424-peergroupid?language=objc
func (d_ DeviceWrapper) PeerGroupID() uint64 {
	rv := objc.Call[uint64](d_, objc.Sel("peerGroupID"))
	return rv
}

func (d_ DeviceWrapper) HasMaxArgumentBufferSamplerCount() bool {
	return d_.RespondsToSelector(objc.Sel("maxArgumentBufferSamplerCount"))
}

// The maximum number of unique argument buffer samplers per app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/2977322-maxargumentbuffersamplercount?language=objc
func (d_ DeviceWrapper) MaxArgumentBufferSamplerCount() uint {
	rv := objc.Call[uint](d_, objc.Sel("maxArgumentBufferSamplerCount"))
	return rv
}

func (d_ DeviceWrapper) HasIsHeadless() bool {
	return d_.RespondsToSelector(objc.Sel("isHeadless"))
}

// A Boolean value that indicates whether a GPU device doesn’t have a connection to a display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/1433377-headless?language=objc
func (d_ DeviceWrapper) IsHeadless() bool {
	rv := objc.Call[bool](d_, objc.Sel("isHeadless"))
	return rv
}

func (d_ DeviceWrapper) HasSupportsRaytracingFromRender() bool {
	return d_.RespondsToSelector(objc.Sel("supportsRaytracingFromRender"))
}

// A Boolean value that indicates whether you can call ray-tracing functions from a vertex or fragment shader. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3750528-supportsraytracingfromrender?language=objc
func (d_ DeviceWrapper) SupportsRaytracingFromRender() bool {
	rv := objc.Call[bool](d_, objc.Sel("supportsRaytracingFromRender"))
	return rv
}

func (d_ DeviceWrapper) HasSupportsQueryTextureLOD() bool {
	return d_.RespondsToSelector(objc.Sel("supportsQueryTextureLOD"))
}

// A Boolean value that indicates whether you can query the texture level of detail from within a shader. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3566548-supportsquerytexturelod?language=objc
func (d_ DeviceWrapper) SupportsQueryTextureLOD() bool {
	rv := objc.Call[bool](d_, objc.Sel("supportsQueryTextureLOD"))
	return rv
}

func (d_ DeviceWrapper) HasMaxTransferRate() bool {
	return d_.RespondsToSelector(objc.Sel("maxTransferRate"))
}

// The highest theoretical rate, in bytes per second, the system can copy between system memory and the GPU’s dedicated memory (VRAM). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3114007-maxtransferrate?language=objc
func (d_ DeviceWrapper) MaxTransferRate() uint64 {
	rv := objc.Call[uint64](d_, objc.Sel("maxTransferRate"))
	return rv
}

func (d_ DeviceWrapper) HasMaxThreadgroupMemoryLength() bool {
	return d_.RespondsToSelector(objc.Sel("maxThreadgroupMemoryLength"))
}

// The maximum threadgroup memory available to a compute kernel, in bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/2877429-maxthreadgroupmemorylength?language=objc
func (d_ DeviceWrapper) MaxThreadgroupMemoryLength() uint {
	rv := objc.Call[uint](d_, objc.Sel("maxThreadgroupMemoryLength"))
	return rv
}

func (d_ DeviceWrapper) HasHasUnifiedMemory() bool {
	return d_.RespondsToSelector(objc.Sel("hasUnifiedMemory"))
}

// A Boolean value that indicates whether the GPU shares all of its memory with the CPU. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3229025-hasunifiedmemory?language=objc
func (d_ DeviceWrapper) HasUnifiedMemory() bool {
	rv := objc.Call[bool](d_, objc.Sel("hasUnifiedMemory"))
	return rv
}

func (d_ DeviceWrapper) HasSupportsDynamicLibraries() bool {
	return d_.RespondsToSelector(objc.Sel("supportsDynamicLibraries"))
}

// A Boolean value that indicates whether the GPU device can create and use dynamic libraries in compute pipelines. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/3553977-supportsdynamiclibraries?language=objc
func (d_ DeviceWrapper) SupportsDynamicLibraries() bool {
	rv := objc.Call[bool](d_, objc.Sel("supportsDynamicLibraries"))
	return rv
}

func (d_ DeviceWrapper) HasCurrentAllocatedSize() bool {
	return d_.RespondsToSelector(objc.Sel("currentAllocatedSize"))
}

// The total amount of memory, in bytes, the GPU device is using for all of its resources. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/2915745-currentallocatedsize?language=objc
func (d_ DeviceWrapper) CurrentAllocatedSize() uint {
	rv := objc.Call[uint](d_, objc.Sel("currentAllocatedSize"))
	return rv
}

func (d_ DeviceWrapper) HasArgumentBuffersSupport() bool {
	return d_.RespondsToSelector(objc.Sel("argumentBuffersSupport"))
}

// Returns the GPU device’s support tier for argument buffers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/2915742-argumentbufferssupport?language=objc
func (d_ DeviceWrapper) ArgumentBuffersSupport() ArgumentBuffersTier {
	rv := objc.Call[ArgumentBuffersTier](d_, objc.Sel("argumentBuffersSupport"))
	return rv
}
