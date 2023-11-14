package metal

// The per-patch tessellation factors for a quad patch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlquadtessellationfactorshalf?language=objc
type QuadTessellationFactorsHalf struct {
	EdgeTessellationFactor   [4]uint16
	InsideTessellationFactor [2]uint16
}

// The data layout for mapping sparse texture regions when using indirect commands. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlmapindirectarguments?language=objc
type MapIndirectArguments struct {
	RegionOriginX    uint32
	RegionOriginY    uint32
	RegionOriginZ    uint32
	RegionSizeWidth  uint32
	RegionSizeHeight uint32
	RegionSizeDepth  uint32
	MipMapLevel      uint32
	SliceId          uint32
}

// The per-patch tessellation factors for a triangle patch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltriangletessellationfactorshalf?language=objc
type TriangleTessellationFactorsHalf struct {
	EdgeTessellationFactor   [3]uint16
	InsideTessellationFactor uint16
}

// An RGBA value used for a color pixel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlclearcolor?language=objc
type ClearColor struct {
	Red   float64
	Green float64
	Blue  float64
	Alpha float64
}

// A description of an instance in an instanced geometry acceleration structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructureinstancedescriptor?language=objc
type AccelerationStructureInstanceDescriptor struct {
	TransformationMatrix            PackedFloat4x3
	Options                         uint32
	Mask                            uint32
	IntersectionFunctionTableOffset uint32
	AccelerationStructureIndex      uint32
}

// A description of an instance in an instanced geometry acceleration structure, with the instance including a user identifier and motion data for the instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotioninstancedescriptor?language=objc
type AccelerationStructureMotionInstanceDescriptor struct {
	Options                         uint32
	Mask                            uint32
	IntersectionFunctionTableOffset uint32
	AccelerationStructureIndex      uint32
	UserID                          uint32
	MotionTransformsStartIndex      uint32
	MotionTransformsCount           uint32
	MotionStartBorderMode           uint32
	MotionEndBorderMode             uint32
	MotionStartTime                 float32
	MotionEndTime                   float32
}

// A description of an instance in an instanced geometry acceleration structure, with the instance including a user identifier for the instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructureuseridinstancedescriptor?language=objc
type AccelerationStructureUserIDInstanceDescriptor struct {
	TransformationMatrix            PackedFloat4x3
	Options                         uint32
	Mask                            uint32
	IntersectionFunctionTableOffset uint32
	AccelerationStructureIndex      uint32
	UserID                          uint32
}

// The dimensions of an object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsize?language=objc
type Size struct {
	Width  uint64
	Height uint64
	Depth  uint64
}

// The bounds for a subset of an object's elements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlregion?language=objc
type Region struct {
	Origin Origin
	Size   Size
}

// The data layout required for the arguments needed to specify the stage-in region. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstageinregionindirectarguments?language=objc
type StageInRegionIndirectArguments struct {
	StageInOrigin [3]uint32
	StageInSize   [3]uint32
}

// A range of commands in an indirect command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectcommandbufferexecutionrange?language=objc
type IndirectCommandBufferExecutionRange struct {
	Location uint32
	Length   uint32
}

// The data layout required for drawing patches via indirect buffer calls. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldrawpatchindirectarguments?language=objc
type DrawPatchIndirectArguments struct {
	PatchCount    uint32
	InstanceCount uint32
	PatchStart    uint32
	BaseInstance  uint32
}

// The data layout required for drawing primitives via indirect buffer calls. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldrawprimitivesindirectarguments?language=objc
type DrawPrimitivesIndirectArguments struct {
	VertexCount   uint32
	InstanceCount uint32
	VertexStart   uint32
	BaseInstance  uint32
}

// A structure that contains the top three rows of a 4x4 matrix of 32-bit floating-point values, in column-major order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlpackedfloat4x3?language=objc
type PackedFloat4x3 struct {
	Columns [4]PackedFloat3
}

// A 3D rectangular region for the viewport clipping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlviewport?language=objc
type Viewport struct {
	OriginX float64
	OriginY float64
	Width   float64
	Height  float64
	Znear   float64
	Zfar    float64
}

// The data structure for storing the data you resolve from a stage-utilization counter set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcounterresultstageutilization?language=objc
type CounterResultStageUtilization struct {
	TotalCycles                  uint64
	VertexCycles                 uint64
	TessellationCycles           uint64
	PostTessellationVertexCycles uint64
	FragmentCycles               uint64
	RenderTargetCycles           uint64
}

// A pattern that modifies the data read or sampled from a texture by rearranging or duplicating the elements of a vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltextureswizzlechannels?language=objc
type TextureSwizzleChannels struct {
	Red   uint8
	Green uint8
	Blue  uint8
	Alpha uint8
}

// The data layout required for the arguments needed to specify the size of threadgroups. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldispatchthreadgroupsindirectarguments?language=objc
type DispatchThreadgroupsIndirectArguments struct {
	ThreadgroupsPerGrid [3]uint32
}

// The bounds for an axis-aligned bounding box. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaxisalignedboundingbox?language=objc
type AxisAlignedBoundingBox struct {
	Min PackedFloat3
	Max PackedFloat3
}

// An offset applied to a render target index and viewport index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexamplificationviewmapping?language=objc
type VertexAmplificationViewMapping struct {
	ViewportArrayIndexOffset     uint32
	RenderTargetArrayIndexOffset uint32
}

// The size and alignment of a resource, in bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsizeandalign?language=objc
type SizeAndAlign struct {
	Size  uint64
	Align uint64
}

// A structure that contains three 32-bit floating-point values with no additional padding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlpackedfloat3?language=objc
type PackedFloat3 struct {
	Anon0 [12]byte
}

// The data structure for storing the data you resolve from a statistic counter set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcounterresultstatistic?language=objc
type CounterResultStatistic struct {
	TessellationInputPatches          uint64
	VertexInvocations                 uint64
	PostTessellationVertexInvocations uint64
	ClipperInvocations                uint64
	ClipperPrimitivesOut              uint64
	FragmentInvocations               uint64
	FragmentsPassed                   uint64
	ComputeKernelInvocations          uint64
}

// The coordinates for the front upper-left corner of a region. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlorigin?language=objc
type Origin struct {
	X uint64
	Y uint64
	Z uint64
}

// The data layout required for drawing indexed primitives via indirect buffer calls. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldrawindexedprimitivesindirectarguments?language=objc
type DrawIndexedPrimitivesIndirectArguments struct {
	IndexCount    uint32
	InstanceCount uint32
	IndexStart    uint32
	BaseVertex    int32
	BaseInstance  uint32
}

// The expected sizes for a ray-tracing acceleration structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuresizes?language=objc
type AccelerationStructureSizes struct {
	AccelerationStructureSize uint64
	BuildScratchBufferSize    uint64
	RefitScratchBufferSize    uint64
}

// A sample position on a subpixel grid. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsampleposition?language=objc
type SamplePosition struct {
	X float32
	Y float32
}

// The data structure for storing the data you resolve from a timestamp counter set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcounterresulttimestamp?language=objc
type CounterResultTimestamp struct {
	Timestamp uint64
}

// A rectangle for the scissor fragment test. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlscissorrect?language=objc
type ScissorRect struct {
	X      uint64
	Y      uint64
	Width  uint64
	Height uint64
}
