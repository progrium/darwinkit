// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RenderPipelineDescriptor] class.
var RenderPipelineDescriptorClass = _RenderPipelineDescriptorClass{objc.GetClass("MTLRenderPipelineDescriptor")}

type _RenderPipelineDescriptorClass struct {
	objc.Class
}

// An interface definition for the [RenderPipelineDescriptor] class.
type IRenderPipelineDescriptor interface {
	objc.IObject
	Reset()
	VertexLinkedFunctions() LinkedFunctions
	SetVertexLinkedFunctions(value ILinkedFunctions)
	FragmentFunction() FunctionWrapper
	SetFragmentFunction(value PFunction)
	SetFragmentFunctionObject(valueObject objc.IObject)
	TessellationFactorFormat() TessellationFactorFormat
	SetTessellationFactorFormat(value TessellationFactorFormat)
	IsRasterizationEnabled() bool
	SetRasterizationEnabled(value bool)
	FragmentLinkedFunctions() LinkedFunctions
	SetFragmentLinkedFunctions(value ILinkedFunctions)
	FragmentBuffers() PipelineBufferDescriptorArray
	MaxFragmentCallStackDepth() uint
	SetMaxFragmentCallStackDepth(value uint)
	InputPrimitiveTopology() PrimitiveTopologyClass
	SetInputPrimitiveTopology(value PrimitiveTopologyClass)
	ColorAttachments() RenderPipelineColorAttachmentDescriptorArray
	IsAlphaToOneEnabled() bool
	SetAlphaToOneEnabled(value bool)
	TessellationOutputWindingOrder() Winding
	SetTessellationOutputWindingOrder(value Winding)
	VertexBuffers() PipelineBufferDescriptorArray
	SupportAddingVertexBinaryFunctions() bool
	SetSupportAddingVertexBinaryFunctions(value bool)
	IsAlphaToCoverageEnabled() bool
	SetAlphaToCoverageEnabled(value bool)
	MaxVertexCallStackDepth() uint
	SetMaxVertexCallStackDepth(value uint)
	VertexDescriptor() VertexDescriptor
	SetVertexDescriptor(value IVertexDescriptor)
	FragmentPreloadedLibraries() []DynamicLibraryWrapper
	SetFragmentPreloadedLibraries(value []PDynamicLibrary)
	TessellationControlPointIndexType() TessellationControlPointIndexType
	SetTessellationControlPointIndexType(value TessellationControlPointIndexType)
	IsTessellationFactorScaleEnabled() bool
	SetTessellationFactorScaleEnabled(value bool)
	RasterSampleCount() uint
	SetRasterSampleCount(value uint)
	StencilAttachmentPixelFormat() PixelFormat
	SetStencilAttachmentPixelFormat(value PixelFormat)
	TessellationPartitionMode() TessellationPartitionMode
	SetTessellationPartitionMode(value TessellationPartitionMode)
	MaxTessellationFactor() uint
	SetMaxTessellationFactor(value uint)
	VertexPreloadedLibraries() []DynamicLibraryWrapper
	SetVertexPreloadedLibraries(value []PDynamicLibrary)
	VertexFunction() FunctionWrapper
	SetVertexFunction(value PFunction)
	SetVertexFunctionObject(valueObject objc.IObject)
	MaxVertexAmplificationCount() uint
	SetMaxVertexAmplificationCount(value uint)
	Label() string
	SetLabel(value string)
	TessellationFactorStepFunction() TessellationFactorStepFunction
	SetTessellationFactorStepFunction(value TessellationFactorStepFunction)
	SupportAddingFragmentBinaryFunctions() bool
	SetSupportAddingFragmentBinaryFunctions(value bool)
	DepthAttachmentPixelFormat() PixelFormat
	SetDepthAttachmentPixelFormat(value PixelFormat)
	SupportIndirectCommandBuffers() bool
	SetSupportIndirectCommandBuffers(value bool)
	BinaryArchives() []BinaryArchiveWrapper
	SetBinaryArchives(value []PBinaryArchive)
}

// An argument of options you pass to a GPU device to get a render pipeline state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor?language=objc
type RenderPipelineDescriptor struct {
	objc.Object
}

func RenderPipelineDescriptorFrom(ptr unsafe.Pointer) RenderPipelineDescriptor {
	return RenderPipelineDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (rc _RenderPipelineDescriptorClass) Alloc() RenderPipelineDescriptor {
	rv := objc.Call[RenderPipelineDescriptor](rc, objc.Sel("alloc"))
	return rv
}

func RenderPipelineDescriptor_Alloc() RenderPipelineDescriptor {
	return RenderPipelineDescriptorClass.Alloc()
}

func (rc _RenderPipelineDescriptorClass) New() RenderPipelineDescriptor {
	rv := objc.Call[RenderPipelineDescriptor](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRenderPipelineDescriptor() RenderPipelineDescriptor {
	return RenderPipelineDescriptorClass.New()
}

func (r_ RenderPipelineDescriptor) Init() RenderPipelineDescriptor {
	rv := objc.Call[RenderPipelineDescriptor](r_, objc.Sel("init"))
	return rv
}

// Specifies the default rendering pipeline state values for the descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1514688-reset?language=objc
func (r_ RenderPipelineDescriptor) Reset() {
	objc.Call[objc.Void](r_, objc.Sel("reset"))
}

// Functions that you can specify as function arguments for the vertex shader when encoding commands that use the pipeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/3750576-vertexlinkedfunctions?language=objc
func (r_ RenderPipelineDescriptor) VertexLinkedFunctions() LinkedFunctions {
	rv := objc.Call[LinkedFunctions](r_, objc.Sel("vertexLinkedFunctions"))
	return rv
}

// Functions that you can specify as function arguments for the vertex shader when encoding commands that use the pipeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/3750576-vertexlinkedfunctions?language=objc
func (r_ RenderPipelineDescriptor) SetVertexLinkedFunctions(value ILinkedFunctions) {
	objc.Call[objc.Void](r_, objc.Sel("setVertexLinkedFunctions:"), objc.Ptr(value))
}

// The fragment function the pipeline calls to process fragments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1514600-fragmentfunction?language=objc
func (r_ RenderPipelineDescriptor) FragmentFunction() FunctionWrapper {
	rv := objc.Call[FunctionWrapper](r_, objc.Sel("fragmentFunction"))
	return rv
}

// The fragment function the pipeline calls to process fragments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1514600-fragmentfunction?language=objc
func (r_ RenderPipelineDescriptor) SetFragmentFunction(value PFunction) {
	po0 := objc.WrapAsProtocol("MTLFunction", value)
	objc.Call[objc.Void](r_, objc.Sel("setFragmentFunction:"), po0)
}

// The fragment function the pipeline calls to process fragments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1514600-fragmentfunction?language=objc
func (r_ RenderPipelineDescriptor) SetFragmentFunctionObject(valueObject objc.IObject) {
	objc.Call[objc.Void](r_, objc.Sel("setFragmentFunction:"), objc.Ptr(valueObject))
}

// The format of the tessellation factors in the tessellation factor buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1639951-tessellationfactorformat?language=objc
func (r_ RenderPipelineDescriptor) TessellationFactorFormat() TessellationFactorFormat {
	rv := objc.Call[TessellationFactorFormat](r_, objc.Sel("tessellationFactorFormat"))
	return rv
}

// The format of the tessellation factors in the tessellation factor buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1639951-tessellationfactorformat?language=objc
func (r_ RenderPipelineDescriptor) SetTessellationFactorFormat(value TessellationFactorFormat) {
	objc.Call[objc.Void](r_, objc.Sel("setTessellationFactorFormat:"), value)
}

// A Boolean value that determines whether the pipeline rasterizes primitives. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1514708-rasterizationenabled?language=objc
func (r_ RenderPipelineDescriptor) IsRasterizationEnabled() bool {
	rv := objc.Call[bool](r_, objc.Sel("isRasterizationEnabled"))
	return rv
}

// A Boolean value that determines whether the pipeline rasterizes primitives. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1514708-rasterizationenabled?language=objc
func (r_ RenderPipelineDescriptor) SetRasterizationEnabled(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setRasterizationEnabled:"), value)
}

// Functions that you can specify as function arguments for the fragment shader when encoding commands that use the pipeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/3750569-fragmentlinkedfunctions?language=objc
func (r_ RenderPipelineDescriptor) FragmentLinkedFunctions() LinkedFunctions {
	rv := objc.Call[LinkedFunctions](r_, objc.Sel("fragmentLinkedFunctions"))
	return rv
}

// Functions that you can specify as function arguments for the fragment shader when encoding commands that use the pipeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/3750569-fragmentlinkedfunctions?language=objc
func (r_ RenderPipelineDescriptor) SetFragmentLinkedFunctions(value ILinkedFunctions) {
	objc.Call[objc.Void](r_, objc.Sel("setFragmentLinkedFunctions:"), objc.Ptr(value))
}

// An array that contains the buffer mutability options for a render pipeline's fragment function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/2879275-fragmentbuffers?language=objc
func (r_ RenderPipelineDescriptor) FragmentBuffers() PipelineBufferDescriptorArray {
	rv := objc.Call[PipelineBufferDescriptorArray](r_, objc.Sel("fragmentBuffers"))
	return rv
}

// The maximum function call depth from the top-most fragment shader function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/3750572-maxfragmentcallstackdepth?language=objc
func (r_ RenderPipelineDescriptor) MaxFragmentCallStackDepth() uint {
	rv := objc.Call[uint](r_, objc.Sel("maxFragmentCallStackDepth"))
	return rv
}

// The maximum function call depth from the top-most fragment shader function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/3750572-maxfragmentcallstackdepth?language=objc
func (r_ RenderPipelineDescriptor) SetMaxFragmentCallStackDepth(value uint) {
	objc.Call[objc.Void](r_, objc.Sel("setMaxFragmentCallStackDepth:"), value)
}

// The type of primitive topology the pipeline renders. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1514684-inputprimitivetopology?language=objc
func (r_ RenderPipelineDescriptor) InputPrimitiveTopology() PrimitiveTopologyClass {
	rv := objc.Call[PrimitiveTopologyClass](r_, objc.Sel("inputPrimitiveTopology"))
	return rv
}

// The type of primitive topology the pipeline renders. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1514684-inputprimitivetopology?language=objc
func (r_ RenderPipelineDescriptor) SetInputPrimitiveTopology(value PrimitiveTopologyClass) {
	objc.Call[objc.Void](r_, objc.Sel("setInputPrimitiveTopology:"), value)
}

// An array of attachments that store color data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1514712-colorattachments?language=objc
func (r_ RenderPipelineDescriptor) ColorAttachments() RenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Call[RenderPipelineColorAttachmentDescriptorArray](r_, objc.Sel("colorAttachments"))
	return rv
}

// A Boolean value that indicates whether to force alpha channel values for color attachments to the largest representable value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1514697-alphatooneenabled?language=objc
func (r_ RenderPipelineDescriptor) IsAlphaToOneEnabled() bool {
	rv := objc.Call[bool](r_, objc.Sel("isAlphaToOneEnabled"))
	return rv
}

// A Boolean value that indicates whether to force alpha channel values for color attachments to the largest representable value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1514697-alphatooneenabled?language=objc
func (r_ RenderPipelineDescriptor) SetAlphaToOneEnabled(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setAlphaToOneEnabled:"), value)
}

// The winding order of triangles from the tessellator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1639911-tessellationoutputwindingorder?language=objc
func (r_ RenderPipelineDescriptor) TessellationOutputWindingOrder() Winding {
	rv := objc.Call[Winding](r_, objc.Sel("tessellationOutputWindingOrder"))
	return rv
}

// The winding order of triangles from the tessellator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1639911-tessellationoutputwindingorder?language=objc
func (r_ RenderPipelineDescriptor) SetTessellationOutputWindingOrder(value Winding) {
	objc.Call[objc.Void](r_, objc.Sel("setTessellationOutputWindingOrder:"), value)
}

// An array that contains the buffer mutability options for a render pipeline's vertex function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/2879280-vertexbuffers?language=objc
func (r_ RenderPipelineDescriptor) VertexBuffers() PipelineBufferDescriptorArray {
	rv := objc.Call[PipelineBufferDescriptorArray](r_, objc.Sel("vertexBuffers"))
	return rv
}

// A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to the vertex shader’s callable functions list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/3750575-supportaddingvertexbinaryfunctio?language=objc
func (r_ RenderPipelineDescriptor) SupportAddingVertexBinaryFunctions() bool {
	rv := objc.Call[bool](r_, objc.Sel("supportAddingVertexBinaryFunctions"))
	return rv
}

// A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to the vertex shader’s callable functions list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/3750575-supportaddingvertexbinaryfunctio?language=objc
func (r_ RenderPipelineDescriptor) SetSupportAddingVertexBinaryFunctions(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setSupportAddingVertexBinaryFunctions:"), value)
}

// A Boolean value that indicates whether to read and use the alpha channel fragment output for color attachments to compute a sample coverage mask. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1514624-alphatocoverageenabled?language=objc
func (r_ RenderPipelineDescriptor) IsAlphaToCoverageEnabled() bool {
	rv := objc.Call[bool](r_, objc.Sel("isAlphaToCoverageEnabled"))
	return rv
}

// A Boolean value that indicates whether to read and use the alpha channel fragment output for color attachments to compute a sample coverage mask. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1514624-alphatocoverageenabled?language=objc
func (r_ RenderPipelineDescriptor) SetAlphaToCoverageEnabled(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setAlphaToCoverageEnabled:"), value)
}

// The maximum function call depth from the top-most vertex shader function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/3750573-maxvertexcallstackdepth?language=objc
func (r_ RenderPipelineDescriptor) MaxVertexCallStackDepth() uint {
	rv := objc.Call[uint](r_, objc.Sel("maxVertexCallStackDepth"))
	return rv
}

// The maximum function call depth from the top-most vertex shader function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/3750573-maxvertexcallstackdepth?language=objc
func (r_ RenderPipelineDescriptor) SetMaxVertexCallStackDepth(value uint) {
	objc.Call[objc.Void](r_, objc.Sel("setMaxVertexCallStackDepth:"), value)
}

// The organization of vertex data in an attribute’s argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1514681-vertexdescriptor?language=objc
func (r_ RenderPipelineDescriptor) VertexDescriptor() VertexDescriptor {
	rv := objc.Call[VertexDescriptor](r_, objc.Sel("vertexDescriptor"))
	return rv
}

// The organization of vertex data in an attribute’s argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1514681-vertexdescriptor?language=objc
func (r_ RenderPipelineDescriptor) SetVertexDescriptor(value IVertexDescriptor) {
	objc.Call[objc.Void](r_, objc.Sel("setVertexDescriptor:"), objc.Ptr(value))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/3801723-fragmentpreloadedlibraries?language=objc
func (r_ RenderPipelineDescriptor) FragmentPreloadedLibraries() []DynamicLibraryWrapper {
	rv := objc.Call[[]DynamicLibraryWrapper](r_, objc.Sel("fragmentPreloadedLibraries"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/3801723-fragmentpreloadedlibraries?language=objc
func (r_ RenderPipelineDescriptor) SetFragmentPreloadedLibraries(value []PDynamicLibrary) {
	objc.Call[objc.Void](r_, objc.Sel("setFragmentPreloadedLibraries:"), value)
}

// The size of the control point indices in a control point index buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1640059-tessellationcontrolpointindextyp?language=objc
func (r_ RenderPipelineDescriptor) TessellationControlPointIndexType() TessellationControlPointIndexType {
	rv := objc.Call[TessellationControlPointIndexType](r_, objc.Sel("tessellationControlPointIndexType"))
	return rv
}

// The size of the control point indices in a control point index buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1640059-tessellationcontrolpointindextyp?language=objc
func (r_ RenderPipelineDescriptor) SetTessellationControlPointIndexType(value TessellationControlPointIndexType) {
	objc.Call[objc.Void](r_, objc.Sel("setTessellationControlPointIndexType:"), value)
}

// A Boolean value that determines whether the pipeline scales the tessellation factor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1640045-tessellationfactorscaleenabled?language=objc
func (r_ RenderPipelineDescriptor) IsTessellationFactorScaleEnabled() bool {
	rv := objc.Call[bool](r_, objc.Sel("isTessellationFactorScaleEnabled"))
	return rv
}

// A Boolean value that determines whether the pipeline scales the tessellation factor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1640045-tessellationfactorscaleenabled?language=objc
func (r_ RenderPipelineDescriptor) SetTessellationFactorScaleEnabled(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setTessellationFactorScaleEnabled:"), value)
}

// The number of samples the pipeline applies for each fragment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/2890271-rastersamplecount?language=objc
func (r_ RenderPipelineDescriptor) RasterSampleCount() uint {
	rv := objc.Call[uint](r_, objc.Sel("rasterSampleCount"))
	return rv
}

// The number of samples the pipeline applies for each fragment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/2890271-rastersamplecount?language=objc
func (r_ RenderPipelineDescriptor) SetRasterSampleCount(value uint) {
	objc.Call[objc.Void](r_, objc.Sel("setRasterSampleCount:"), value)
}

// The pixel format of the attachment that stores stencil data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1514650-stencilattachmentpixelformat?language=objc
func (r_ RenderPipelineDescriptor) StencilAttachmentPixelFormat() PixelFormat {
	rv := objc.Call[PixelFormat](r_, objc.Sel("stencilAttachmentPixelFormat"))
	return rv
}

// The pixel format of the attachment that stores stencil data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1514650-stencilattachmentpixelformat?language=objc
func (r_ RenderPipelineDescriptor) SetStencilAttachmentPixelFormat(value PixelFormat) {
	objc.Call[objc.Void](r_, objc.Sel("setStencilAttachmentPixelFormat:"), value)
}

// The partitioning mode that the tessellator uses to derive the number and spacing of segments for subdividing a corresponding edge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1639979-tessellationpartitionmode?language=objc
func (r_ RenderPipelineDescriptor) TessellationPartitionMode() TessellationPartitionMode {
	rv := objc.Call[TessellationPartitionMode](r_, objc.Sel("tessellationPartitionMode"))
	return rv
}

// The partitioning mode that the tessellator uses to derive the number and spacing of segments for subdividing a corresponding edge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1639979-tessellationpartitionmode?language=objc
func (r_ RenderPipelineDescriptor) SetTessellationPartitionMode(value TessellationPartitionMode) {
	objc.Call[objc.Void](r_, objc.Sel("setTessellationPartitionMode:"), value)
}

// The maximum tessellation factor that the tessellator uses when tessellating patches. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1640060-maxtessellationfactor?language=objc
func (r_ RenderPipelineDescriptor) MaxTessellationFactor() uint {
	rv := objc.Call[uint](r_, objc.Sel("maxTessellationFactor"))
	return rv
}

// The maximum tessellation factor that the tessellator uses when tessellating patches. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1640060-maxtessellationfactor?language=objc
func (r_ RenderPipelineDescriptor) SetMaxTessellationFactor(value uint) {
	objc.Call[objc.Void](r_, objc.Sel("setMaxTessellationFactor:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/3801724-vertexpreloadedlibraries?language=objc
func (r_ RenderPipelineDescriptor) VertexPreloadedLibraries() []DynamicLibraryWrapper {
	rv := objc.Call[[]DynamicLibraryWrapper](r_, objc.Sel("vertexPreloadedLibraries"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/3801724-vertexpreloadedlibraries?language=objc
func (r_ RenderPipelineDescriptor) SetVertexPreloadedLibraries(value []PDynamicLibrary) {
	objc.Call[objc.Void](r_, objc.Sel("setVertexPreloadedLibraries:"), value)
}

// The vertex function the pipeline calls to process vertices. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1514679-vertexfunction?language=objc
func (r_ RenderPipelineDescriptor) VertexFunction() FunctionWrapper {
	rv := objc.Call[FunctionWrapper](r_, objc.Sel("vertexFunction"))
	return rv
}

// The vertex function the pipeline calls to process vertices. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1514679-vertexfunction?language=objc
func (r_ RenderPipelineDescriptor) SetVertexFunction(value PFunction) {
	po0 := objc.WrapAsProtocol("MTLFunction", value)
	objc.Call[objc.Void](r_, objc.Sel("setVertexFunction:"), po0)
}

// The vertex function the pipeline calls to process vertices. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1514679-vertexfunction?language=objc
func (r_ RenderPipelineDescriptor) SetVertexFunctionObject(valueObject objc.IObject) {
	objc.Call[objc.Void](r_, objc.Sel("setVertexFunction:"), objc.Ptr(valueObject))
}

// The maximum vertex amplification count you can set when encoding render commands. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/3088854-maxvertexamplificationcount?language=objc
func (r_ RenderPipelineDescriptor) MaxVertexAmplificationCount() uint {
	rv := objc.Call[uint](r_, objc.Sel("maxVertexAmplificationCount"))
	return rv
}

// The maximum vertex amplification count you can set when encoding render commands. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/3088854-maxvertexamplificationcount?language=objc
func (r_ RenderPipelineDescriptor) SetMaxVertexAmplificationCount(value uint) {
	objc.Call[objc.Void](r_, objc.Sel("setMaxVertexAmplificationCount:"), value)
}

// A string that identifies the render pipeline descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1514705-label?language=objc
func (r_ RenderPipelineDescriptor) Label() string {
	rv := objc.Call[string](r_, objc.Sel("label"))
	return rv
}

// A string that identifies the render pipeline descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1514705-label?language=objc
func (r_ RenderPipelineDescriptor) SetLabel(value string) {
	objc.Call[objc.Void](r_, objc.Sel("setLabel:"), value)
}

// The step function for determining the tessellation factors for a patch from the tessellation factor buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1640062-tessellationfactorstepfunction?language=objc
func (r_ RenderPipelineDescriptor) TessellationFactorStepFunction() TessellationFactorStepFunction {
	rv := objc.Call[TessellationFactorStepFunction](r_, objc.Sel("tessellationFactorStepFunction"))
	return rv
}

// The step function for determining the tessellation factors for a patch from the tessellation factor buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1640062-tessellationfactorstepfunction?language=objc
func (r_ RenderPipelineDescriptor) SetTessellationFactorStepFunction(value TessellationFactorStepFunction) {
	objc.Call[objc.Void](r_, objc.Sel("setTessellationFactorStepFunction:"), value)
}

// A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to the fragment shader’s callable functions list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/3750574-supportaddingfragmentbinaryfunct?language=objc
func (r_ RenderPipelineDescriptor) SupportAddingFragmentBinaryFunctions() bool {
	rv := objc.Call[bool](r_, objc.Sel("supportAddingFragmentBinaryFunctions"))
	return rv
}

// A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to the fragment shader’s callable functions list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/3750574-supportaddingfragmentbinaryfunct?language=objc
func (r_ RenderPipelineDescriptor) SetSupportAddingFragmentBinaryFunctions(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setSupportAddingFragmentBinaryFunctions:"), value)
}

// The pixel format of the attachment that stores depth data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1514608-depthattachmentpixelformat?language=objc
func (r_ RenderPipelineDescriptor) DepthAttachmentPixelFormat() PixelFormat {
	rv := objc.Call[PixelFormat](r_, objc.Sel("depthAttachmentPixelFormat"))
	return rv
}

// The pixel format of the attachment that stores depth data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/1514608-depthattachmentpixelformat?language=objc
func (r_ RenderPipelineDescriptor) SetDepthAttachmentPixelFormat(value PixelFormat) {
	objc.Call[objc.Void](r_, objc.Sel("setDepthAttachmentPixelFormat:"), value)
}

// A Boolean value that determines whether you can encode commands into an indirect command buffer using the render pipeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/2966638-supportindirectcommandbuffers?language=objc
func (r_ RenderPipelineDescriptor) SupportIndirectCommandBuffers() bool {
	rv := objc.Call[bool](r_, objc.Sel("supportIndirectCommandBuffers"))
	return rv
}

// A Boolean value that determines whether you can encode commands into an indirect command buffer using the render pipeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/2966638-supportindirectcommandbuffers?language=objc
func (r_ RenderPipelineDescriptor) SetSupportIndirectCommandBuffers(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setSupportIndirectCommandBuffers:"), value)
}

// An array of binary archives to search for precompiled versions of the shader. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/3554050-binaryarchives?language=objc
func (r_ RenderPipelineDescriptor) BinaryArchives() []BinaryArchiveWrapper {
	rv := objc.Call[[]BinaryArchiveWrapper](r_, objc.Sel("binaryArchives"))
	return rv
}

// An array of binary archives to search for precompiled versions of the shader. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/3554050-binaryarchives?language=objc
func (r_ RenderPipelineDescriptor) SetBinaryArchives(value []PBinaryArchive) {
	objc.Call[objc.Void](r_, objc.Sel("setBinaryArchives:"), value)
}
