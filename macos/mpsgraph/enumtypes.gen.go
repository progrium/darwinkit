// AUTO-GENERATED CODE, DO NOT MODIFY

package mpsgraph

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdevicetype?language=objc
type DeviceType uint32

const (
	DeviceTypeMetal DeviceType = 0
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlossreductiontype?language=objc
type LossReductionType uint64

const (
	LossReductionTypeAxis LossReductionType = 0
	LossReductionTypeMean LossReductionType = 2
	LossReductionTypeSum  LossReductionType = 1
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphoptimization?language=objc
type Optimization uint64

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphoptimizationprofile?language=objc
type OptimizationProfile uint64

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphoptions?language=objc
type Options uint64

const (
	OptionsDefault            Options = 1
	OptionsNone               Options = 0
	OptionsSynchronizeResults Options = 1
	OptionsVerbose            Options = 2
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpaddingmode?language=objc
type PaddingMode int

const (
	PaddingModeAntiPeriodic PaddingMode = 6
	PaddingModeClampToEdge  PaddingMode = 3
	PaddingModeConstant     PaddingMode = 0
	PaddingModePeriodic     PaddingMode = 5
	PaddingModeReflect      PaddingMode = 1
	PaddingModeSymmetric    PaddingMode = 2
	PaddingModeZero         PaddingMode = 4
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpaddingstyle?language=objc
type PaddingStyle uint

const (
	PaddingStyleExplicit       PaddingStyle = 0
	PaddingStyleExplicitOffset PaddingStyle = 3
	PaddingStyleTF_SAME        PaddingStyle = 2
	PaddingStyleTF_VALID       PaddingStyle = 1
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpoolingreturnindicesmode?language=objc
type PoolingReturnIndicesMode uint

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrnnactivation?language=objc
type RNNActivation uint

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomdistribution?language=objc
type RandomDistribution uint64

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomnormalsamplingmethod?language=objc
type RandomNormalSamplingMethod uint64

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphreductionmode?language=objc
type ReductionMode uint

const (
	ReductionModeArgumentMax ReductionMode = 5
	ReductionModeArgumentMin ReductionMode = 4
	ReductionModeMax         ReductionMode = 1
	ReductionModeMin         ReductionMode = 0
	ReductionModeProduct     ReductionMode = 3
	ReductionModeSum         ReductionMode = 2
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphresizemode?language=objc
type ResizeMode uint

const (
	ResizeBilinear ResizeMode = 1
	ResizeNearest  ResizeMode = 0
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphscattermode?language=objc
type ScatterMode int

const (
	ScatterModeAdd ScatterMode = 0
	ScatterModeDiv ScatterMode = 3
	ScatterModeMax ScatterMode = 5
	ScatterModeMin ScatterMode = 4
	ScatterModeMul ScatterMode = 2
	ScatterModeSet ScatterMode = 6
	ScatterModeSub ScatterMode = 1
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphsparsestoragetype?language=objc
type SparseStorageType uint64

const (
	SparseStorageCOO SparseStorageType = 0
	SparseStorageCSC SparseStorageType = 1
	SparseStorageCSR SparseStorageType = 2
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphtensornameddatalayout?language=objc
type TensorNamedDataLayout uint

const (
	TensorNamedDataLayoutCHW  TensorNamedDataLayout = 4
	TensorNamedDataLayoutHW   TensorNamedDataLayout = 6
	TensorNamedDataLayoutHWC  TensorNamedDataLayout = 5
	TensorNamedDataLayoutHWIO TensorNamedDataLayout = 3
	TensorNamedDataLayoutNCHW TensorNamedDataLayout = 0
	TensorNamedDataLayoutNHWC TensorNamedDataLayout = 1
	TensorNamedDataLayoutOIHW TensorNamedDataLayout = 2
)
