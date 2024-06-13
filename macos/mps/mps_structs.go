package mps

// A 3D ray with an origin, a direction, and a mask to filter out intersections. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayoriginmaskdirectionmaxdistance?language=objc
type RayOriginMaskDirectionMaxDistance struct {
	Origin      PackedFloat3
	Mask        uint32
	Direction   PackedFloat3
	MaxDistance float32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsdimensionslice?language=objc
type DimensionSlice struct {
	Start  uint64
	Length uint64
}

// A region of an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsregion?language=objc
type Region struct {
	Origin Origin
	Size   Size
}

// A structure that specifies keypoint information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagekeypointdata?language=objc
type ImageKeypointData struct {
	KeypointCoordinate [2]uint16
	KeypointColorValue float32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscustomkernelsourceinfo?language=objc
type CustomKernelSourceInfo struct {
	KernelOrigin         [2]int16
	KernelPhase          [2]uint16
	KernelSize           [2]uint16
	Offset               [2]int16
	Stride               [2]uint16
	DilationRate         [2]uint16
	FeatureChannelOffset uint16
	FeatureChannels      uint16
	ImageArrayOffset     uint16
	ImageArraySize       uint16
}

// An intersection result that contains the origin-intersection distance, and intersected primitive and instance indices. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsintersectiondistanceprimitiveindexinstanceindex?language=objc
type IntersectionDistancePrimitiveIndexInstanceIndex struct {
	Distance       float32
	PrimitiveIndex uint32
	InstanceIndex  uint32
}

// Parameters that control reading and writing of a particular set of feature channels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereadwriteparams?language=objc
type ImageReadWriteParams struct {
	FeatureChannelOffset               uint64
	NumberOfFeatureChannelsToReadWrite uint64
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraysizes?language=objc
type NDArraySizes struct {
	Dimensions [16]uint64
}

// Parameters that define the parts of a division operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsintegerdivisionparams?language=objc
type IntegerDivisionParams struct {
	Divisor uint16
	Recip   uint16
	Addend  uint16
	Shift   uint16
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsintersectiondistanceprimitiveindexbufferindexinstanceindexcoordinates?language=objc
type IntersectionDistancePrimitiveIndexBufferIndexInstanceIndexCoordinates struct {
	Distance       float32
	PrimitiveIndex uint32
	BufferIndex    uint32
	InstanceIndex  uint32
	Coordinates    [2]float32
}

// An intersection result that contains the distance from the ray origin to the intersection point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsintersectiondistance?language=objc
type IntersectionDistance struct {
	Distance float32
}

// An axis-aligned bounding box. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaxisalignedboundingbox?language=objc
type AxisAlignedBoundingBox struct {
	Min       [3]float32
	Pad_cgo_0 [4]byte
	Max       [3]float32
	Pad_cgo_1 [4]byte
}

// A structure that specifies information to find the keypoints in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagekeypointrangeinfo?language=objc
type ImageKeypointRangeInfo struct {
	MaximumKeypoints      uint64
	MinimumThresholdValue float32
	Pad_cgo_0             [4]byte
}

// A 3D ray with an origin and a direction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayorigindirection?language=objc
type RayOriginDirection struct {
	Origin    [3]float32
	Pad_cgo_0 [4]byte
	Direction [3]float32
	Pad_cgo_1 [4]byte
}

// A signed coordinate with x, y, and z components. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsoffset?language=objc
type Offset struct {
	X int64
	Y int64
	Z int64
}

// An encapsulation of a texture's dimensions, format, type, and usage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstatetextureinfo?language=objc
type StateTextureInfo struct {
	Width       uint64
	Height      uint64
	Depth       uint64
	ArrayLength uint64
	PixelFormat uint64
	TextureType uint64
	Usage       uint64
	X_reserved  [4]uint64
}

// An intersection result that contains the distance from the ray origin to the intersection point, and the index of the intersected primitive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsintersectiondistanceprimitiveindex?language=objc
type IntersectionDistancePrimitiveIndex struct {
	Distance       float32
	PrimitiveIndex uint32
}

// A transform matrix for explicit resampling control with a Lanczos kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsscaletransform?language=objc
type ScaleTransform struct {
	ScaleX     float64
	ScaleY     float64
	TranslateX float64
	TranslateY float64
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsintersectiondistanceprimitiveindexbufferindexinstanceindex?language=objc
type IntersectionDistancePrimitiveIndexBufferIndexInstanceIndex struct {
	Distance       float32
	PrimitiveIndex uint32
	BufferIndex    uint32
	InstanceIndex  uint32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsintersectiondistanceprimitiveindexbufferindexcoordinates?language=objc
type IntersectionDistancePrimitiveIndexBufferIndexCoordinates struct {
	Distance       float32
	PrimitiveIndex uint32
	BufferIndex    uint32
	Pad_cgo_0      [4]byte
	Coordinates    [2]float32
}

// A structure that contains the number of destination, source, and broadcaset textures used by a custom kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscustomkernelargumentcount?language=objc
type CustomKernelArgumentCount struct {
	DestinationTextureCount uint64
	SourceTextureCount      uint64
	BroadcastTextureCount   uint64
}

// An intersection result that contains the origin-intersection distance, intersected primitive and instance indices, and intersection point coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsintersectiondistanceprimitiveindexinstanceindexcoordinates?language=objc
type IntersectionDistancePrimitiveIndexInstanceIndexCoordinates struct {
	Distance       float32
	PrimitiveIndex uint32
	InstanceIndex  uint32
	Pad_cgo_0      [4]byte
	Coordinates    [2]float32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageregion?language=objc
type ImageRegion struct {
	Offset ImageCoordinate
	Size   ImageCoordinate
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsraypackedorigindirection?language=objc
type RayPackedOriginDirection struct {
	Origin    PackedFloat3
	Direction PackedFloat3
}

// A packed three-element vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspackedfloat3?language=objc
type PackedFloat3 struct {
	Anon0 [12]byte
}

// A description of row and column offsets into a matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixoffset?language=objc
type MatrixOffset struct {
	RowOffset    uint32
	ColumnOffset uint32
}

// An intersection result that contains the origin-intersection distance, intersected primitive index, and intersection point coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsintersectiondistanceprimitiveindexcoordinates?language=objc
type IntersectionDistancePrimitiveIndexCoordinates struct {
	Distance       float32
	PrimitiveIndex uint32
	Coordinates    [2]float32
}

// The information used to compute the histogram channels of an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistograminfo?language=objc
type ImageHistogramInfo struct {
	NumberOfHistogramEntries uint64
	HistogramForAlpha        bool
	Pad_cgo_0                [4]byte
	MinPixelValue            [4]float32
	MaxPixelValue            [4]float32
}

// A size of a region in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssize?language=objc
type Size struct {
	Width  float64
	Height float64
	Depth  float64
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsintersectiondistanceprimitiveindexbufferindex?language=objc
type IntersectionDistancePrimitiveIndexBufferIndex struct {
	Distance       float32
	PrimitiveIndex uint32
	BufferIndex    uint32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscustomkernelinfo?language=objc
type CustomKernelInfo struct {
	ClipOrigin                 [4]uint16
	ClipSize                   [4]uint16
	DestinationFeatureChannels uint16
	DestImageArraySize         uint16
	SourceImageCount           uint16
	ThreadgroupSize            uint16
	SubbatchIndex              uint16
	SubbatchStride             uint16
	Idiv                       IntegerDivisionParams
	Pad_cgo_0                  [4]byte
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayoffsets?language=objc
type NDArrayOffsets struct {
	Dimensions [16]int64
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecoordinate?language=objc
type ImageCoordinate struct {
	X       uint64
	Y       uint64
	Channel uint64
}

// A 3D ray with an origin, a direction, and an intersection distance range from the origin. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayoriginmindistancedirectionmaxdistance?language=objc
type RayOriginMinDistanceDirectionMaxDistance struct {
	Origin      PackedFloat3
	MinDistance float32
	Direction   PackedFloat3
	MaxDistance float32
}

// A description of matrix copy operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopyoffsets?language=objc
type MatrixCopyOffsets struct {
	SourceRowOffset         uint32
	SourceColumnOffset      uint32
	DestinationRowOffset    uint32
	DestinationColumnOffset uint32
}

// A position in an image used as the source origin. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsorigin?language=objc
type Origin struct {
	X float64
	Y float64
	Z float64
}
