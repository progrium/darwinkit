// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

// Constants that indicate an acceleration structure build state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaccelerationstructurestatus?language=objc
type AccelerationStructureStatus uint

const (
	AccelerationStructureStatusBuilt   AccelerationStructureStatus = 1
	AccelerationStructureStatusUnbuilt AccelerationStructureStatus = 0
)

// Options that describe how an acceleration structure will be used. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaccelerationstructureusage?language=objc
type AccelerationStructureUsage uint

const (
	AccelerationStructureUsageFrequentRebuild AccelerationStructureUsage = 2
	AccelerationStructureUsageNone            AccelerationStructureUsage = 0
	AccelerationStructureUsagePreferCPUBuild  AccelerationStructureUsage = 8
	AccelerationStructureUsagePreferGPUBuild  AccelerationStructureUsage = 4
	AccelerationStructureUsageRefit           AccelerationStructureUsage = 1
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaliasingstrategy?language=objc
type AliasingStrategy uint

const (
	AliasingStrategyAliasingReserved         AliasingStrategy = 3
	AliasingStrategyDefault                  AliasingStrategy = 0
	AliasingStrategyDontCare                 AliasingStrategy = 0
	AliasingStrategyPreferNonTemporaryMemory AliasingStrategy = 8
	AliasingStrategyPreferTemporaryMemory    AliasingStrategy = 4
	AliasingStrategyShallAlias               AliasingStrategy = 1
	AliasingStrategyShallNotAlias            AliasingStrategy = 2
)

// Premultiplication description for the color channels of an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsalphatype?language=objc
type AlphaType uint

const (
	AlphaTypeAlphaIsOne       AlphaType = 1
	AlphaTypeNonPremultiplied AlphaType = 0
	AlphaTypePremultiplied    AlphaType = 2
)

// Options for the intersection test type for a ray intersector bounding box. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsboundingboxintersectiontesttype?language=objc
type BoundingBoxIntersectionTestType uint

const (
	BoundingBoxIntersectionTestTypeAxisAligned BoundingBoxIntersectionTestType = 1
	BoundingBoxIntersectionTestTypeDefault     BoundingBoxIntersectionTestType = 0
	BoundingBoxIntersectionTestTypeFast        BoundingBoxIntersectionTestType = 2
)

// Options that define how statistics are calculated during batch normalization. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationflags?language=objc
type CNNBatchNormalizationFlags uint

const (
	CNNBatchNormalizationFlagsCalculateStatisticsAlways    CNNBatchNormalizationFlags = 1
	CNNBatchNormalizationFlagsCalculateStatisticsAutomatic CNNBatchNormalizationFlags = 0
	CNNBatchNormalizationFlagsCalculateStatisticsMask      CNNBatchNormalizationFlags = 3
	CNNBatchNormalizationFlagsCalculateStatisticsNever     CNNBatchNormalizationFlags = 2
	CNNBatchNormalizationFlagsDefault                      CNNBatchNormalizationFlags = 0
)

// Options used to control binary convolution kernels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryconvolutionflags?language=objc
type CNNBinaryConvolutionFlags uint

const (
	CNNBinaryConvolutionFlagsNone           CNNBinaryConvolutionFlags = 0
	CNNBinaryConvolutionFlagsUseBetaScaling CNNBinaryConvolutionFlags = 1
)

// Options that defines what operations are used to perform binary convolution. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryconvolutiontype?language=objc
type CNNBinaryConvolutionType uint

const (
	CNNBinaryConvolutionTypeAND           CNNBinaryConvolutionType = 2
	CNNBinaryConvolutionTypeBinaryWeights CNNBinaryConvolutionType = 0
	CNNBinaryConvolutionTypeXNOR          CNNBinaryConvolutionType = 1
)

// Options used to control how kernel weights are stored and used in the CNN kernels [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionflags?language=objc
type CNNConvolutionFlags uint

const (
	CNNConvolutionFlagsNone CNNConvolutionFlags = 0
)

// Options that control which gradient to compute during backward propagation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradientoption?language=objc
type CNNConvolutionGradientOption uint

const (
	CNNConvolutionGradientOptionAll                        CNNConvolutionGradientOption = 3
	CNNConvolutionGradientOptionGradientWithData           CNNConvolutionGradientOption = 1
	CNNConvolutionGradientOptionGradientWithWeightsAndBias CNNConvolutionGradientOption = 2
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionweightslayout?language=objc
type CNNConvolutionWeightsLayout uint32

const (
	CNNConvolutionWeightsLayoutOHWI CNNConvolutionWeightsLayout = 0
)

// Constants that indicate supported loss filter types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlosstype?language=objc
type CNNLossType uint32

const (
	CNNLossTypeCategoricalCrossEntropy   CNNLossType = 4
	CNNLossTypeCosineDistance            CNNLossType = 7
	CNNLossTypeCount                     CNNLossType = 10
	CNNLossTypeHinge                     CNNLossType = 5
	CNNLossTypeHuber                     CNNLossType = 6
	CNNLossTypeKullbackLeiblerDivergence CNNLossType = 9
	CNNLossTypeLog                       CNNLossType = 8
	CNNLossTypeMeanAbsoluteError         CNNLossType = 0
	CNNLossTypeMeanSquaredError          CNNLossType = 1
	CNNLossTypeSigmoidCrossEntropy       CNNLossType = 3
	CNNLossTypeSoftMaxCrossEntropy       CNNLossType = 2
)

// The types of neuron filter to append to a convolution. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurontype?language=objc
type CNNNeuronType int32

const (
	CNNNeuronTypeAbsolute    CNNNeuronType = 6
	CNNNeuronTypeCount       CNNNeuronType = 16
	CNNNeuronTypeELU         CNNNeuronType = 9
	CNNNeuronTypeExponential CNNNeuronType = 13
	CNNNeuronTypeGeLU        CNNNeuronType = 15
	CNNNeuronTypeHardSigmoid CNNNeuronType = 4
	CNNNeuronTypeLinear      CNNNeuronType = 2
	CNNNeuronTypeLogarithm   CNNNeuronType = 14
	CNNNeuronTypeNone        CNNNeuronType = 0
	CNNNeuronTypePReLU       CNNNeuronType = 10
	CNNNeuronTypePower       CNNNeuronType = 12
	CNNNeuronTypeReLU        CNNNeuronType = 1
	CNNNeuronTypeReLUN       CNNNeuronType = 11
	CNNNeuronTypeSigmoid     CNNNeuronType = 3
	CNNNeuronTypeSoftPlus    CNNNeuronType = 7
	CNNNeuronTypeSoftSign    CNNNeuronType = 8
	CNNNeuronTypeTanH        CNNNeuronType = 5
)

// Constants that indicate supported reduction types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnreductiontype?language=objc
type CNNReductionType int32

const (
	CNNReductionTypeCount               CNNReductionType = 4
	CNNReductionTypeMean                CNNReductionType = 2
	CNNReductionTypeNone                CNNReductionType = 0
	CNNReductionTypeSum                 CNNReductionType = 1
	CNNReductionTypeSumByNonZeroWeights CNNReductionType = 3
)

// Options that specify the type of quantization used to generate unsigned integer weights. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnweightsquantizationtype?language=objc
type CNNWeightsQuantizationType uint32

const (
	CNNWeightsQuantizationTypeLinear      CNNWeightsQuantizationType = 1
	CNNWeightsQuantizationTypeLookupTable CNNWeightsQuantizationType = 2
	CNNWeightsQuantizationTypeNone        CNNWeightsQuantizationType = 0
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscustomkernelindex?language=objc
type CustomKernelIndex int

const (
	CustomKernelIndexDestIndex     CustomKernelIndex = 0
	CustomKernelIndexSrc0Index     CustomKernelIndex = 0
	CustomKernelIndexSrc1Index     CustomKernelIndex = 1
	CustomKernelIndexSrc2Index     CustomKernelIndex = 2
	CustomKernelIndexSrc3Index     CustomKernelIndex = 3
	CustomKernelIndexSrc4Index     CustomKernelIndex = 4
	CustomKernelIndexUserDataIndex CustomKernelIndex = 30
)

// Options that define how buffer data is arranged. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsdatalayout?language=objc
type DataLayout uint

const (
	DataLayoutFeatureChannelsxHeightxWidth DataLayout = 1
	DataLayoutHeightxWidthxFeatureChannels DataLayout = 0
)

// A value to specify a type of data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsdatatype?language=objc
type DataType uint32

const (
	DataTypeAlternateEncodingBit DataType = 2147483648
	DataTypeBool                 DataType = 2147483656
	DataTypeFloat16              DataType = 268435472
	DataTypeFloat32              DataType = 268435488
	DataTypeFloatBit             DataType = 268435456
	DataTypeInt16                DataType = 536870928
	DataTypeInt32                DataType = 536870944
	DataTypeInt64                DataType = 536870976
	DataTypeInt8                 DataType = 536870920
	DataTypeIntBit               DataType = 536870912
	DataTypeInvalid              DataType = 0
	DataTypeNormalizedBit        DataType = 1073741824
	DataTypeSignedBit            DataType = 536870912
	DataTypeUInt16               DataType = 16
	DataTypeUInt32               DataType = 32
	DataTypeUInt64               DataType = 64
	DataTypeUInt8                DataType = 8
	DataTypeUnorm1               DataType = 1073741825
	DataTypeUnorm8               DataType = 1073741832
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsdevicecaps?language=objc
type DeviceCaps uint32

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsdevicecapsvalues?language=objc
type DeviceCapsValues uint32

const (
	DeviceCapsNull                        DeviceCapsValues = 0
	DeviceIsAppleDevice                   DeviceCapsValues = 1024
	DeviceSupportsFloat16BicubicFiltering DeviceCapsValues = 512
	DeviceSupportsFloat32Filtering        DeviceCapsValues = 128
	DeviceSupportsNorm16BicubicFiltering  DeviceCapsValues = 256
	DeviceSupportsQuadShuffle             DeviceCapsValues = 16
	DeviceSupportsReadWriteTextures       DeviceCapsValues = 4
	DeviceSupportsReadableArrayOfTextures DeviceCapsValues = 1
	DeviceSupportsSimdReduction           DeviceCapsValues = 64
	DeviceSupportsSimdShuffle             DeviceCapsValues = 32
	DeviceSupportsSimdgroupBarrier        DeviceCapsValues = 8
	DeviceSupportsWritableArrayOfTextures DeviceCapsValues = 2
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsdeviceoptions?language=objc
type DeviceOptions uint

const (
	DeviceOptionsDefault       DeviceOptions = 0
	DeviceOptionsLowPower      DeviceOptions = 1
	DeviceOptionsSkipRemovable DeviceOptions = 2
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsfunctionconstant?language=objc
type FunctionConstant int64

const (
	FunctionConstantNone FunctionConstant = -1
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsfunctionconstantinmetal?language=objc
type FunctionConstantInMetal uint32

// The options used to control the edge behavior of an image filter when it reads outside the bounds of a source texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedgemode?language=objc
type ImageEdgeMode uint

const (
	ImageEdgeModeClamp          ImageEdgeMode = 1
	ImageEdgeModeConstant       ImageEdgeMode = 4
	ImageEdgeModeMirror         ImageEdgeMode = 2
	ImageEdgeModeMirrorWithEdge ImageEdgeMode = 3
	ImageEdgeModeZero           ImageEdgeMode = 0
)

// Encodes the representation of a single channel within an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagefeaturechannelformat?language=objc
type ImageFeatureChannelFormat uint

const (
	ImageFeatureChannelFormatCount      ImageFeatureChannelFormat = 6
	ImageFeatureChannelFormatFloat16    ImageFeatureChannelFormat = 3
	ImageFeatureChannelFormatFloat32    ImageFeatureChannelFormat = 4
	ImageFeatureChannelFormatNone       ImageFeatureChannelFormat = 0
	ImageFeatureChannelFormatUnorm16    ImageFeatureChannelFormat = 2
	ImageFeatureChannelFormatUnorm8     ImageFeatureChannelFormat = 1
	ImageFeatureChannelFormat_reserved0 ImageFeatureChannelFormat = 5
)

// Options that define a Metal Performance Shaders image type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagetype?language=objc
type ImageType uint32

const (
	ImageType2d                    ImageType = 0
	ImageType2d_array              ImageType = 1
	ImageType2d_array_noAlpha      ImageType = 5
	ImageType2d_noAlpha            ImageType = 4
	ImageTypeArray2d               ImageType = 2
	ImageTypeArray2d_array         ImageType = 3
	ImageTypeArray2d_array_noAlpha ImageType = 7
	ImageTypeArray2d_noAlpha       ImageType = 6
	ImageType_ArrayMask            ImageType = 1
	ImageType_BatchMask            ImageType = 2
	ImageType_bitCount             ImageType = 6
	ImageType_mask                 ImageType = 63
	ImageType_noAlpha              ImageType = 4
	ImageType_texelFormatBFloat16  ImageType = 24
	ImageType_texelFormatFloat16   ImageType = 16
	ImageType_texelFormatMask      ImageType = 56
	ImageType_texelFormatShift     ImageType = 3
	ImageType_texelFormatStandard  ImageType = 0
	ImageType_texelFormatUnorm8    ImageType = 8
	ImageType_typeMask             ImageType = 3
)

// Options that determine the data contained in an intersection result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsintersectiondatatype?language=objc
type IntersectionDataType uint

const (
	IntersectionDataTypeDistance                                                  IntersectionDataType = 0
	IntersectionDataTypeDistancePrimitiveIndex                                    IntersectionDataType = 1
	IntersectionDataTypeDistancePrimitiveIndexBufferIndex                         IntersectionDataType = 5
	IntersectionDataTypeDistancePrimitiveIndexBufferIndexCoordinates              IntersectionDataType = 6
	IntersectionDataTypeDistancePrimitiveIndexBufferIndexInstanceIndex            IntersectionDataType = 7
	IntersectionDataTypeDistancePrimitiveIndexBufferIndexInstanceIndexCoordinates IntersectionDataType = 8
	IntersectionDataTypeDistancePrimitiveIndexCoordinates                         IntersectionDataType = 2
	IntersectionDataTypeDistancePrimitiveIndexInstanceIndex                       IntersectionDataType = 3
	IntersectionDataTypeDistancePrimitiveIndexInstanceIndexCoordinates            IntersectionDataType = 4
)

// Options that determine an intersection type for a ray intersector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsintersectiontype?language=objc
type IntersectionType uint

const (
	IntersectionTypeAny     IntersectionType = 1
	IntersectionTypeNearest IntersectionType = 0
)

// The options used when creating a kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskerneloptions?language=objc
type KernelOptions uint

const (
	KernelOptionsAllowReducedPrecision KernelOptions = 2
	KernelOptionsDisableInternalTiling KernelOptions = 4
	KernelOptionsInsertDebugGroups     KernelOptions = 8
	KernelOptionsNone                  KernelOptions = 0
	KernelOptionsSkipAPIValidation     KernelOptions = 1
	KernelOptionsVerbose               KernelOptions = 16
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdecompositionstatus?language=objc
type MatrixDecompositionStatus int

const (
	MatrixDecompositionStatusFailure             MatrixDecompositionStatus = -1
	MatrixDecompositionStatusNonPositiveDefinite MatrixDecompositionStatus = -3
	MatrixDecompositionStatusSingular            MatrixDecompositionStatus = -2
	MatrixDecompositionStatusSuccess             MatrixDecompositionStatus = 0
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistribution?language=objc
type MatrixRandomDistribution uint

const (
	MatrixRandomDistributionDefault MatrixRandomDistribution = 1
	MatrixRandomDistributionNormal  MatrixRandomDistribution = 3
	MatrixRandomDistributionUniform MatrixRandomDistribution = 2
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncomparisontype?language=objc
type NNComparisonType uint

const (
	NNComparisonTypeEqual          NNComparisonType = 0
	NNComparisonTypeGreater        NNComparisonType = 4
	NNComparisonTypeGreaterOrEqual NNComparisonType = 5
	NNComparisonTypeLess           NNComparisonType = 2
	NNComparisonTypeLessOrEqual    NNComparisonType = 3
	NNComparisonTypeNotEqual       NNComparisonType = 1
)

// Options that specify convolution accumulator precision. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnconvolutionaccumulatorprecisionoption?language=objc
type NNConvolutionAccumulatorPrecisionOption uint

const (
	NNConvolutionAccumulatorPrecisionOptionFloat NNConvolutionAccumulatorPrecisionOption = 1
	NNConvolutionAccumulatorPrecisionOptionHalf  NNConvolutionAccumulatorPrecisionOption = 0
)

// Options that define a graph's padding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpaddingmethod?language=objc
type NNPaddingMethod uint

const (
	NNPaddingMethodAddRemainderToBottomLeft     NNPaddingMethod = 8
	NNPaddingMethodAddRemainderToBottomRight    NNPaddingMethod = 12
	NNPaddingMethodAddRemainderToMask           NNPaddingMethod = 12
	NNPaddingMethodAddRemainderToTopLeft        NNPaddingMethod = 0
	NNPaddingMethodAddRemainderToTopRight       NNPaddingMethod = 4
	NNPaddingMethodAlignBottomRight             NNPaddingMethod = 2
	NNPaddingMethodAlignCentered                NNPaddingMethod = 0
	NNPaddingMethodAlignMask                    NNPaddingMethod = 3
	NNPaddingMethodAlignTopLeft                 NNPaddingMethod = 1
	NNPaddingMethodAlign_reserved               NNPaddingMethod = 3
	NNPaddingMethodCustom                       NNPaddingMethod = 16384
	NNPaddingMethodCustomAllowForNodeFusion     NNPaddingMethod = 8192
	NNPaddingMethodCustomWhitelistForNodeFusion NNPaddingMethod = 8192
	NNPaddingMethodExcludeEdges                 NNPaddingMethod = 32768
	NNPaddingMethodSizeFull                     NNPaddingMethod = 32
	NNPaddingMethodSizeMask                     NNPaddingMethod = 2032
	NNPaddingMethodSizeSame                     NNPaddingMethod = 16
	NNPaddingMethodSizeValidOnly                NNPaddingMethod = 0
	NNPaddingMethodSize_reserved                NNPaddingMethod = 48
)

// Options that define the regularization type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnregularizationtype?language=objc
type NNRegularizationType uint

const (
	NNRegularizationTypeL1   NNRegularizationType = 1
	NNRegularizationTypeL2   NNRegularizationType = 2
	NNRegularizationTypeNone NNRegularizationType = 0
)

// Options that control how graph nodes are trained. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnntrainingstyle?language=objc
type NNTrainingStyle uint

const (
	NNTrainingStyleUpdateDeviceCPU  NNTrainingStyle = 1
	NNTrainingStyleUpdateDeviceGPU  NNTrainingStyle = 2
	NNTrainingStyleUpdateDeviceNone NNTrainingStyle = 0
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygontype?language=objc
type PolygonType uint

const (
	PolygonTypeQuadrilateral PolygonType = 1
	PolygonTypeTriangle      PolygonType = 0
)

// The purgeable state of an image’s underlying texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspurgeablestate?language=objc
type PurgeableState uint

const (
	PurgeableStateAllocationDeferred PurgeableState = 0
	PurgeableStateEmpty              PurgeableState = 4
	PurgeableStateKeepCurrent        PurgeableState = 1
	PurgeableStateNonVolatile        PurgeableState = 2
	PurgeableStateVolatile           PurgeableState = 3
)

// Modes that define how two images or matrices are combined. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnbidirectionalcombinemode?language=objc
type RNNBidirectionalCombineMode uint

const (
	RNNBidirectionalCombineModeAdd         RNNBidirectionalCombineMode = 1
	RNNBidirectionalCombineModeConcatenate RNNBidirectionalCombineMode = 2
	RNNBidirectionalCombineModeNone        RNNBidirectionalCombineMode = 0
)

// Options that define which matrix is copied in or out of a trainable RNN layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixid?language=objc
type RNNMatrixId uint

const (
	RNNMatrixIdGRUInputGateBiasTerms            RNNMatrixId = 21
	RNNMatrixIdGRUInputGateInputWeights         RNNMatrixId = 19
	RNNMatrixIdGRUInputGateRecurrentWeights     RNNMatrixId = 20
	RNNMatrixIdGRUOutputGateBiasTerms           RNNMatrixId = 28
	RNNMatrixIdGRUOutputGateInputGateWeights    RNNMatrixId = 27
	RNNMatrixIdGRUOutputGateInputWeights        RNNMatrixId = 25
	RNNMatrixIdGRUOutputGateRecurrentWeights    RNNMatrixId = 26
	RNNMatrixIdGRURecurrentGateBiasTerms        RNNMatrixId = 24
	RNNMatrixIdGRURecurrentGateInputWeights     RNNMatrixId = 22
	RNNMatrixIdGRURecurrentGateRecurrentWeights RNNMatrixId = 23
	RNNMatrixIdLSTMForgetGateBiasTerms          RNNMatrixId = 10
	RNNMatrixIdLSTMForgetGateInputWeights       RNNMatrixId = 7
	RNNMatrixIdLSTMForgetGateMemoryWeights      RNNMatrixId = 9
	RNNMatrixIdLSTMForgetGateRecurrentWeights   RNNMatrixId = 8
	RNNMatrixIdLSTMInputGateBiasTerms           RNNMatrixId = 6
	RNNMatrixIdLSTMInputGateInputWeights        RNNMatrixId = 3
	RNNMatrixIdLSTMInputGateMemoryWeights       RNNMatrixId = 5
	RNNMatrixIdLSTMInputGateRecurrentWeights    RNNMatrixId = 4
	RNNMatrixIdLSTMMemoryGateBiasTerms          RNNMatrixId = 14
	RNNMatrixIdLSTMMemoryGateInputWeights       RNNMatrixId = 11
	RNNMatrixIdLSTMMemoryGateMemoryWeights      RNNMatrixId = 13
	RNNMatrixIdLSTMMemoryGateRecurrentWeights   RNNMatrixId = 12
	RNNMatrixIdLSTMOutputGateBiasTerms          RNNMatrixId = 18
	RNNMatrixIdLSTMOutputGateInputWeights       RNNMatrixId = 15
	RNNMatrixIdLSTMOutputGateMemoryWeights      RNNMatrixId = 17
	RNNMatrixIdLSTMOutputGateRecurrentWeights   RNNMatrixId = 16
	RNNMatrixIdSingleGateBiasTerms              RNNMatrixId = 2
	RNNMatrixIdSingleGateInputWeights           RNNMatrixId = 0
	RNNMatrixIdSingleGateRecurrentWeights       RNNMatrixId = 1
	RNNMatrixId_count                           RNNMatrixId = 29
)

// Directions that a sequence of inputs can be processed by a recurrent neural network layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnsequencedirection?language=objc
type RNNSequenceDirection uint

const (
	RNNSequenceDirectionBackward RNNSequenceDirection = 1
	RNNSequenceDirectionForward  RNNSequenceDirection = 0
)

// Options for the data type for an intersector ray. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsraydatatype?language=objc
type RayDataType uint

const (
	RayDataTypeOriginDirection                       RayDataType = 0
	RayDataTypeOriginMaskDirectionMaxDistance        RayDataType = 2
	RayDataTypeOriginMinDistanceDirectionMaxDistance RayDataType = 1
	RayDataTypePackedOriginDirection                 RayDataType = 3
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsraymaskoperator?language=objc
type RayMaskOperator uint

const (
	RayMaskOperatorAnd                  RayMaskOperator = 0
	RayMaskOperatorEqual                RayMaskOperator = 10
	RayMaskOperatorGreaterThan          RayMaskOperator = 8
	RayMaskOperatorGreaterThanOrEqualTo RayMaskOperator = 9
	RayMaskOperatorLessThan             RayMaskOperator = 6
	RayMaskOperatorLessThanOrEqualTo    RayMaskOperator = 7
	RayMaskOperatorNotAnd               RayMaskOperator = 1
	RayMaskOperatorNotEqual             RayMaskOperator = 11
	RayMaskOperatorNotOr                RayMaskOperator = 3
	RayMaskOperatorNotXor               RayMaskOperator = 5
	RayMaskOperatorOr                   RayMaskOperator = 2
	RayMaskOperatorXor                  RayMaskOperator = 4
)

// Options for ray intersector mask options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsraymaskoptions?language=objc
type RayMaskOptions uint

const (
	RayMaskOptionInstance  RayMaskOptions = 2
	RayMaskOptionNone      RayMaskOptions = 0
	RayMaskOptionPrimitive RayMaskOptions = 1
)

// Options for the underlying resource type for a state object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstateresourcetype?language=objc
type StateResourceType uint

const (
	StateResourceTypeBuffer  StateResourceType = 1
	StateResourceTypeNone    StateResourceType = 0
	StateResourceTypeTexture StateResourceType = 2
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporalweighting?language=objc
type TemporalWeighting uint

const (
	TemporalWeightingAverage                  TemporalWeighting = 0
	TemporalWeightingExponentialMovingAverage TemporalWeighting = 1
)

// Constants that indicate instance transformation types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstransformtype?language=objc
type TransformType uint

const (
	TransformTypeFloat4x4 TransformType = 0
	TransformTypeIdentity TransformType = 1
)

// Options for the ray-triangle intersection test. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstriangleintersectiontesttype?language=objc
type TriangleIntersectionTestType uint

const (
	TriangleIntersectionTestTypeDefault    TriangleIntersectionTestType = 0
	TriangleIntersectionTestTypeWatertight TriangleIntersectionTestType = 1
)
