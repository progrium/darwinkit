// AUTO-GENERATED CODE, DO NOT MODIFY

package mpsgraph

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/macos/mps"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Graph] class.
var GraphClass = _GraphClass{objc.GetClass("MPSGraph")}

type _GraphClass struct {
	objc.Class
}

// An interface definition for the [Graph] class.
type IGraph interface {
	objc.IObject
	MaxPooling2DWithSourceTensorDescriptorName(source ITensor, descriptor IPooling2DOpDescriptor, name string) Tensor
	AvgPooling4DWithSourceTensorDescriptorName(source ITensor, descriptor IPooling4DOpDescriptor, name string) Tensor
	LessThanOrEqualToWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor
	DepthwiseConvolution2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeDescriptorName(incomingGradient ITensor, weights ITensor, outputShape *foundation.Array, descriptor IDepthwiseConvolution2DOpDescriptor, name string) Tensor
	SparseTensorWithDescriptorTensorsShapeName(sparseDescriptor ICreateSparseOpDescriptor, inputTensorArray []ITensor, shape *foundation.Array, name string) Tensor
	LSTMWithSourceTensorRecurrentWeightInitStateInitCellDescriptorName(source ITensor, recurrentWeight ITensor, initState ITensor, initCell ITensor, descriptor ILSTMDescriptor, name string) []Tensor
	Convolution2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeTensorForwardConvolutionDescriptorName(gradient ITensor, weights ITensor, outputShapeTensor ITensor, forwardConvolutionDescriptor IConvolution2DOpDescriptor, name string) Tensor
	RunWithMTLCommandQueueFeedsTargetOperationsResultsDictionary(commandQueue metal.PCommandQueue, feeds *foundation.Dictionary, targetOperations []IOperation, resultsDictionary *foundation.Dictionary)
	RunWithMTLCommandQueueObjectFeedsTargetOperationsResultsDictionary(commandQueueObject objc.IObject, feeds *foundation.Dictionary, targetOperations []IOperation, resultsDictionary *foundation.Dictionary)
	ExpandDimsOfTensorAxesTensorName(tensor ITensor, axesTensor ITensor, name string) Tensor
	ReductionAndWithTensorAxisName(tensor ITensor, axis int, name string) Tensor
	ResizeTensorSizeTensorModeCenterResultAlignCornersLayoutName(imagesTensor ITensor, size ITensor, mode ResizeMode, centerResult bool, alignCorners bool, layout TensorNamedDataLayout, name string) Tensor
	TanWithTensorName(tensor ITensor, name string) Tensor
	FloorModuloWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor
	ConvolutionTranspose2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeTensorForwardConvolutionDescriptorName(incomingGradientTensor ITensor, source ITensor, outputShape ITensor, forwardConvolutionDescriptor IConvolution2DOpDescriptor, name string) Tensor
	CoordinateAlongAxisTensorWithShapeTensorName(axisTensor ITensor, shapeTensor ITensor, name string) Tensor
	ReductionProductWithTensorAxesName(tensor ITensor, axes []foundation.INumber, name string) Tensor
	AssignVariableWithValueOfTensorName(variable ITensor, tensor ITensor, name string) Operation
	RandomUniformTensorWithShapeTensorName(shapeTensor ITensor, name string) Tensor
	PadTensorWithPaddingModeLeftPaddingRightPaddingConstantValueName(tensor ITensor, paddingMode PaddingMode, leftPadding *foundation.Array, rightPadding *foundation.Array, constantValue float64, name string) Tensor
	ApplyStochasticGradientDescentWithLearningRateTensorVariableGradientTensorName(learningRateTensor ITensor, variable IVariableOp, gradientTensor ITensor, name string) Operation
	SelectWithPredicateTensorTruePredicateTensorFalsePredicateTensorName(predicateTensor ITensor, truePredicateTensor ITensor, falseSelectTensor ITensor, name string) Tensor
	DepthwiseConvolution3DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeDescriptorName(incomingGradient ITensor, source ITensor, outputShape *foundation.Array, descriptor IDepthwiseConvolution3DOpDescriptor, name string) Tensor
	LeakyReLUWithTensorAlphaName(tensor ITensor, alpha float64, name string) Tensor
	LogicalNORWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor
	ScatterNDWithDataTensorUpdatesTensorIndicesTensorBatchDimensionsModeName(dataTensor ITensor, updatesTensor ITensor, indicesTensor ITensor, batchDimensions uint, mode ScatterMode, name string) Tensor
	SquareWithTensorName(tensor ITensor, name string) Tensor
	ForLoopWithNumberOfIterationsInitialBodyArgumentsBodyName(numberOfIterations ITensor, initialBodyArguments []ITensor, body ForLoopBodyBlock, name string) []Tensor
	LogicalORWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor
	ReductionOrWithTensorAxisName(tensor ITensor, axis int, name string) Tensor
	DepthToSpace2DTensorWidthAxisHeightAxisDepthAxisBlockSizeUsePixelShuffleOrderName(tensor ITensor, widthAxis uint, heightAxis uint, depthAxis uint, blockSize uint, usePixelShuffleOrder bool, name string) Tensor
	MaxPooling4DGradientWithGradientTensorSourceTensorDescriptorName(gradient ITensor, source ITensor, descriptor IPooling4DOpDescriptor, name string) Tensor
	AsinWithTensorName(tensor ITensor, name string) Tensor
	SigmoidWithTensorName(tensor ITensor, name string) Tensor
	TileGradientWithIncomingGradientTensorSourceTensorWithMultiplierName(incomingGradientTensor ITensor, sourceTensor ITensor, multiplier *foundation.Array, name string) Tensor
	ScatterWithUpdatesTensorIndicesTensorShapeAxisModeName(updatesTensor ITensor, indicesTensor ITensor, shape *foundation.Array, axis int, mode ScatterMode, name string) Tensor
	ErfWithTensorName(tensor ITensor, name string) Tensor
	TanhWithTensorName(tensor ITensor, name string) Tensor
	BandPartWithTensorNumLowerTensorNumUpperTensorName(inputTensor ITensor, numLowerTensor ITensor, numUpperTensor ITensor, name string) Tensor
	DepthwiseConvolution2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeDescriptorName(incomingGradient ITensor, source ITensor, outputShape *foundation.Array, descriptor IDepthwiseConvolution2DOpDescriptor, name string) Tensor
	SoftMaxCrossEntropyWithSourceTensorLabelsTensorAxisReductionTypeName(sourceTensor ITensor, labelsTensor ITensor, axis int, reductionType LossReductionType, name string) Tensor
	SignbitWithTensorName(tensor ITensor, name string) Tensor
	StencilWithSourceTensorWeightsTensorDescriptorName(source ITensor, weights ITensor, descriptor IStencilOpDescriptor, name string) Tensor
	AvgPooling4DGradientWithGradientTensorSourceTensorDescriptorName(gradient ITensor, source ITensor, descriptor IPooling4DOpDescriptor, name string) Tensor
	Convolution2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeTensorForwardConvolutionDescriptorName(gradient ITensor, source ITensor, outputShapeTensor ITensor, forwardConvolutionDescriptor IConvolution2DOpDescriptor, name string) Tensor
	RandomTensorWithShapeDescriptorName(shape *foundation.Array, descriptor IRandomOpDescriptor, name string) Tensor
	AtanhWithTensorName(tensor ITensor, name string) Tensor
	ReductionArgMinimumWithTensorAxisName(tensor ITensor, axis int, name string) Tensor
	ConvolutionTranspose2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeTensorForwardConvolutionDescriptorName(incomingGradient ITensor, weights ITensor, outputShape ITensor, forwardConvolutionDescriptor IConvolution2DOpDescriptor, name string) Tensor
	IsNaNWithTensorName(tensor ITensor, name string) Tensor
	MaximumWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor
	MaxPooling2DReturnIndicesWithSourceTensorDescriptorName(source ITensor, descriptor IPooling2DOpDescriptor, name string) []Tensor
	IfWithPredicateTensorThenBlockElseBlockName(predicateTensor ITensor, thenBlock IfThenElseBlock, elseBlock IfThenElseBlock, name string) []Tensor
	LessThanWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor
	ScatterAlongAxisWithDataTensorUpdatesTensorIndicesTensorModeName(axis int, dataTensor ITensor, updatesTensor ITensor, indicesTensor ITensor, mode ScatterMode, name string) Tensor
	LogarithmBase10WithTensorName(tensor ITensor, name string) Tensor
	SpaceToDepth2DTensorWidthAxisHeightAxisDepthAxisBlockSizeUsePixelShuffleOrderName(tensor ITensor, widthAxis uint, heightAxis uint, depthAxis uint, blockSize uint, usePixelShuffleOrder bool, name string) Tensor
	AdditionWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor
	TileTensorWithMultiplierName(tensor ITensor, multiplier *foundation.Array, name string) Tensor
	ExponentWithTensorName(tensor ITensor, name string) Tensor
	CastTensorToTypeName(tensor ITensor, type_ mps.DataType, name string) Tensor
	MeanOfTensorAxesName(tensor ITensor, axes []foundation.INumber, name string) Tensor
	ReverseSquareRootWithTensorName(tensor ITensor, name string) Tensor
	ScatterNDWithUpdatesTensorIndicesTensorShapeBatchDimensionsModeName(updatesTensor ITensor, indicesTensor ITensor, shape *foundation.Array, batchDimensions uint, mode ScatterMode, name string) Tensor
	CoordinateAlongAxisWithShapeTensorName(axis int, shapeTensor ITensor, name string) Tensor
	GatherAlongAxisWithUpdatesTensorIndicesTensorName(axis int, updatesTensor ITensor, indicesTensor ITensor, name string) Tensor
	LogicalXORWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor
	FloorWithTensorName(tensor ITensor, name string) Tensor
	DepthwiseConvolution3DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeDescriptorName(incomingGradient ITensor, weights ITensor, outputShape *foundation.Array, descriptor IDepthwiseConvolution3DOpDescriptor, name string) Tensor
	LSTMGradientsWithSourceTensorRecurrentWeightSourceGradientZStateCellOutputFwdDescriptorName(source ITensor, recurrentWeight ITensor, sourceGradient ITensor, zState ITensor, cellOutputFwd ITensor, descriptor ILSTMDescriptor, name string) []Tensor
	SplitTensorSplitSizesTensorAxisName(tensor ITensor, splitSizesTensor ITensor, axis int, name string) []Tensor
	TopKWithGradientTensorSourceKTensorName(gradient ITensor, source ITensor, kTensor ITensor, name string) Tensor
	IsFiniteWithTensorName(tensor ITensor, name string) Tensor
	SubtractionWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor
	ReadVariableName(variable ITensor, name string) Tensor
	LeakyReLUGradientWithIncomingGradientSourceTensorAlphaTensorName(gradient ITensor, source ITensor, alphaTensor ITensor, name string) Tensor
	SqueezeTensorName(tensor ITensor, name string) Tensor
	NormalizationGradientWithIncomingGradientTensorSourceTensorMeanTensorVarianceTensorGammaTensorGammaGradientTensorBetaGradientTensorReductionAxesEpsilonName(incomingGradientTensor ITensor, sourceTensor ITensor, meanTensor ITensor, varianceTensor ITensor, gamma ITensor, gammaGradient ITensor, betaGradient ITensor, axes []foundation.INumber, epsilon float64, name string) Tensor
	Convolution2DWithSourceTensorWeightsTensorDescriptorName(source ITensor, weights ITensor, descriptor IConvolution2DOpDescriptor, name string) Tensor
	ReLUGradientWithIncomingGradientSourceTensorName(gradient ITensor, source ITensor, name string) Tensor
	Flatten2DTensorAxisTensorName(tensor ITensor, axisTensor ITensor, name string) Tensor
	OneHotWithIndicesTensorDepthName(indicesTensor ITensor, depth uint, name string) Tensor
	ConcatTensorsDimensionInterleaveName(tensors []ITensor, dimensionIndex int, interleave bool, name string) Tensor
	ReductionMinimumWithTensorAxesName(tensor ITensor, axes []foundation.INumber, name string) Tensor
	ReductionMaximumPropagateNaNWithTensorAxisName(tensor ITensor, axis int, name string) Tensor
	RintWithTensorName(tensor ITensor, name string) Tensor
	Atan2WithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor
	CosWithTensorName(tensor ITensor, name string) Tensor
	MaxPooling4DWithSourceTensorDescriptorName(source ITensor, descriptor IPooling4DOpDescriptor, name string) Tensor
	RunAsyncWithMTLCommandQueueFeedsTargetOperationsResultsDictionaryExecutionDescriptor(commandQueue metal.PCommandQueue, feeds *foundation.Dictionary, targetOperations []IOperation, resultsDictionary *foundation.Dictionary, executionDescriptor IExecutionDescriptor)
	RunAsyncWithMTLCommandQueueObjectFeedsTargetOperationsResultsDictionaryExecutionDescriptor(commandQueueObject objc.IObject, feeds *foundation.Dictionary, targetOperations []IOperation, resultsDictionary *foundation.Dictionary, executionDescriptor IExecutionDescriptor)
	ForLoopWithLowerBoundUpperBoundStepInitialBodyArgumentsBodyName(lowerBound ITensor, upperBound ITensor, step ITensor, initialBodyArguments []ITensor, body ForLoopBodyBlock, name string) []Tensor
	ReLUWithTensorName(tensor ITensor, name string) Tensor
	SigmoidGradientWithIncomingGradientSourceTensorName(gradient ITensor, source ITensor, name string) Tensor
	SliceGradientTensorFwdInShapeTensorStartsEndsStridesName(inputGradientTensor ITensor, fwdInShapeTensor ITensor, starts []foundation.INumber, ends []foundation.INumber, strides []foundation.INumber, name string) Tensor
	AvgPooling2DWithSourceTensorDescriptorName(source ITensor, descriptor IPooling2DOpDescriptor, name string) Tensor
	ConstantWithScalarDataType(scalar float64, dataType mps.DataType) Tensor
	ReshapeTensorWithShapeTensorName(tensor ITensor, shapeTensor ITensor, name string) Tensor
	MatrixMultiplicationWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor
	ControlDependencyWithOperationsDependentBlockName(operations []IOperation, dependentBlock ControlFlowDependencyBlock, name string) []Tensor
	LogicalXNORWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor
	GatherWithUpdatesTensorIndicesTensorAxisBatchDimensionsName(updatesTensor ITensor, indicesTensor ITensor, axis uint, batchDimensions uint, name string) Tensor
	CoshWithTensorName(tensor ITensor, name string) Tensor
	AvgPooling2DGradientWithGradientTensorSourceTensorDescriptorName(gradient ITensor, source ITensor, descriptor IPooling2DOpDescriptor, name string) Tensor
	ReductionMinimumPropagateNaNWithTensorAxisName(tensor ITensor, axis int, name string) Tensor
	RandomUniformTensorWithShapeName(shape *foundation.Array, name string) Tensor
	LogarithmWithTensorName(tensor ITensor, name string) Tensor
	SparseTensorWithTypeTensorsShapeDataTypeName(sparseStorageType SparseStorageType, inputTensorArray []ITensor, shape *foundation.Array, dataType mps.DataType, name string) Tensor
	AbsoluteWithTensorName(tensor ITensor, name string) Tensor
	ConvolutionTranspose2DWithSourceTensorWeightsTensorOutputShapeTensorDescriptorName(source ITensor, weights ITensor, outputShape ITensor, descriptor IConvolution2DOpDescriptor, name string) Tensor
	DivisionNoNaNWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor
	AcosWithTensorName(tensor ITensor, name string) Tensor
	MaxPooling4DReturnIndicesWithSourceTensorDescriptorName(source ITensor, descriptor IPooling4DOpDescriptor, name string) []Tensor
	RandomTensorWithShapeTensorDescriptorName(shapeTensor ITensor, descriptor IRandomOpDescriptor, name string) Tensor
	ConstantWithDataShapeDataType(data []byte, shape *foundation.Array, dataType mps.DataType) Tensor
	RandomPhiloxStateTensorWithSeedName(seed uint, name string) Tensor
	L2NormPooling4DGradientWithGradientTensorSourceTensorDescriptorName(gradient ITensor, source ITensor, descriptor IPooling4DOpDescriptor, name string) Tensor
	ResizeWithGradientTensorInputModeCenterResultAlignCornersLayoutName(gradient ITensor, input ITensor, mode ResizeMode, centerResult bool, alignCorners bool, layout TensorNamedDataLayout, name string) Tensor
	VariableWithDataShapeDataTypeName(data []byte, shape *foundation.Array, dataType mps.DataType, name string) Tensor
	GreaterThanOrEqualToWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor
	GradientForPrimaryTensorWithTensorsName(primaryTensor ITensor, tensors []ITensor, name string) foundation.Dictionary
	CeilWithTensorName(tensor ITensor, name string) Tensor
	ReverseTensorName(tensor ITensor, name string) Tensor
	ExponentBase2WithTensorName(tensor ITensor, name string) Tensor
	MinimumWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor
	IsInfiniteWithTensorName(tensor ITensor, name string) Tensor
	DepthwiseConvolution3DWithSourceTensorWeightsTensorDescriptorName(source ITensor, weights ITensor, descriptor IDepthwiseConvolution3DOpDescriptor, name string) Tensor
	TopKWithSourceTensorKTensorName(source ITensor, kTensor ITensor, name string) []Tensor
	GatherNDWithUpdatesTensorIndicesTensorBatchDimensionsName(updatesTensor ITensor, indicesTensor ITensor, batchDimensions uint, name string) Tensor
	RunAsyncWithFeedsTargetTensorsTargetOperationsExecutionDescriptor(feeds *foundation.Dictionary, targetTensors []ITensor, targetOperations []IOperation, executionDescriptor IExecutionDescriptor) *foundation.Dictionary
	MultiplicationWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor
	SinhWithTensorName(tensor ITensor, name string) Tensor
	DropoutTensorRateName(tensor ITensor, rate float64, name string) Tensor
	SingleGateRNNWithSourceTensorRecurrentWeightInitStateDescriptorName(source ITensor, recurrentWeight ITensor, initState ITensor, descriptor ISingleGateRNNDescriptor, name string) []Tensor
	MinimumWithNaNPropagationWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor
	LogarithmBase2WithTensorName(tensor ITensor, name string) Tensor
	PlaceholderWithShapeName(shape *foundation.Array, name string) Tensor
	NegativeWithTensorName(tensor ITensor, name string) Tensor
	ShapeOfTensorName(tensor ITensor, name string) Tensor
	ExponentBase10WithTensorName(tensor ITensor, name string) Tensor
	RunWithFeedsTargetTensorsTargetOperations(feeds *foundation.Dictionary, targetTensors []ITensor, targetOperations []IOperation) *foundation.Dictionary
	MaximumWithNaNPropagationWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor
	ScatterWithDataTensorUpdatesTensorIndicesTensorAxisModeName(dataTensor ITensor, updatesTensor ITensor, indicesTensor ITensor, axis int, mode ScatterMode, name string) Tensor
	GreaterThanWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor
	RoundWithTensorName(tensor ITensor, name string) Tensor
	LogicalANDWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor
	ReductionMaximumWithTensorAxesName(tensor ITensor, axes []foundation.INumber, name string) Tensor
	SignWithTensorName(tensor ITensor, name string) Tensor
	EqualWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor
	AsinhWithTensorName(tensor ITensor, name string) Tensor
	DepthwiseConvolution2DWithSourceTensorWeightsTensorDescriptorName(source ITensor, weights ITensor, descriptor IDepthwiseConvolution2DOpDescriptor, name string) Tensor
	VarianceOfTensorAxesName(tensor ITensor, axes []foundation.INumber, name string) Tensor
	PadGradientWithIncomingGradientTensorSourceTensorPaddingModeLeftPaddingRightPaddingName(incomingGradientTensor ITensor, sourceTensor ITensor, paddingMode PaddingMode, leftPadding *foundation.Array, rightPadding *foundation.Array, name string) Tensor
	NormalizationWithTensorMeanTensorVarianceTensorGammaTensorBetaTensorEpsilonName(tensor ITensor, mean ITensor, variance ITensor, gamma ITensor, beta ITensor, epsilon float64, name string) Tensor
	DivisionWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor
	SingleGateRNNGradientsWithSourceTensorRecurrentWeightSourceGradientZStateInitStateDescriptorName(source ITensor, recurrentWeight ITensor, sourceGradient ITensor, zState ITensor, initState ITensor, descriptor ISingleGateRNNDescriptor, name string) []Tensor
	EncodeToCommandBufferFeedsTargetOperationsResultsDictionaryExecutionDescriptor(commandBuffer mps.ICommandBuffer, feeds *foundation.Dictionary, targetOperations []IOperation, resultsDictionary *foundation.Dictionary, executionDescriptor IExecutionDescriptor)
	CompileWithDeviceFeedsTargetTensorsTargetOperationsCompilationDescriptor(device IDevice, feeds *foundation.Dictionary, targetTensors []ITensor, targetOperations []IOperation, compilationDescriptor ICompilationDescriptor) Executable
	StochasticGradientDescentWithLearningRateTensorValuesTensorGradientTensorName(learningRateTensor ITensor, valuesTensor ITensor, gradientTensor ITensor, name string) Tensor
	ReciprocalWithTensorName(tensor ITensor, name string) Tensor
	SquareRootWithTensorName(tensor ITensor, name string) Tensor
	NotWithTensorName(tensor ITensor, name string) Tensor
	SinWithTensorName(tensor ITensor, name string) Tensor
	NotEqualWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor
	SliceTensorStartsEndsStridesName(tensor ITensor, starts []foundation.INumber, ends []foundation.INumber, strides []foundation.INumber, name string) Tensor
	PowerWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor
	LogicalNANDWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor
	NormalizationBetaGradientWithIncomingGradientTensorSourceTensorReductionAxesName(incomingGradientTensor ITensor, sourceTensor ITensor, axes []foundation.INumber, name string) Tensor
	AcoshWithTensorName(tensor ITensor, name string) Tensor
	StackTensorsAxisName(inputTensors []ITensor, axis int, name string) Tensor
	SoftMaxCrossEntropyGradientWithIncomingGradientTensorSourceTensorLabelsTensorAxisReductionTypeName(gradientTensor ITensor, sourceTensor ITensor, labelsTensor ITensor, axis int, reductionType LossReductionType, name string) Tensor
	ReductionArgMaximumWithTensorAxisName(tensor ITensor, axis int, name string) Tensor
	ConcatTensorWithTensorDimensionName(tensor ITensor, tensor2 ITensor, dimensionIndex int, name string) Tensor
	L2NormPooling4DWithSourceTensorDescriptorName(source ITensor, descriptor IPooling4DOpDescriptor, name string) Tensor
	GatherAlongAxisTensorWithUpdatesTensorIndicesTensorName(axisTensor ITensor, updatesTensor ITensor, indicesTensor ITensor, name string) Tensor
	ScatterAlongAxisTensorWithDataTensorUpdatesTensorIndicesTensorModeName(axisTensor ITensor, dataTensor ITensor, updatesTensor ITensor, indicesTensor ITensor, mode ScatterMode, name string) Tensor
	TransposeTensorDimensionWithDimensionName(tensor ITensor, dimensionIndex uint, dimensionIndex2 uint, name string) Tensor
	BroadcastTensorToShapeTensorName(tensor ITensor, shapeTensor ITensor, name string) Tensor
	SoftMaxGradientWithIncomingGradientSourceTensorAxisName(gradient ITensor, source ITensor, axis int, name string) Tensor
	ClampWithTensorMinValueTensorMaxValueTensorName(tensor ITensor, minValueTensor ITensor, maxValueTensor ITensor, name string) Tensor
	MaxPooling2DGradientWithGradientTensorSourceTensorDescriptorName(gradient ITensor, source ITensor, descriptor IPooling2DOpDescriptor, name string) Tensor
	ReductionSumWithTensorAxesName(tensor ITensor, axes []foundation.INumber, name string) Tensor
	ModuloWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor
	RandomPhiloxStateTensorWithCounterLowCounterHighKeyName(counterLow uint, counterHigh uint, key uint, name string) Tensor
	IdentityWithTensorName(tensor ITensor, name string) Tensor
	WhileWithInitialInputsBeforeAfterName(initialInputs []ITensor, before WhileBeforeBlock, after WhileAfterBlock, name string) []Tensor
	AtanWithTensorName(tensor ITensor, name string) Tensor
	NormalizationGammaGradientWithIncomingGradientTensorSourceTensorMeanTensorVarianceTensorReductionAxesEpsilonName(incomingGradientTensor ITensor, sourceTensor ITensor, meanTensor ITensor, varianceTensor ITensor, axes []foundation.INumber, epsilon float64, name string) Tensor
	SoftMaxWithTensorAxisName(tensor ITensor, axis int, name string) Tensor
	Options() Options
	SetOptions(value Options)
	PlaceholderTensors() []Tensor
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph?language=objc
type Graph struct {
	objc.Object
}

func GraphFrom(ptr unsafe.Pointer) Graph {
	return Graph{
		Object: objc.ObjectFrom(ptr),
	}
}

func (gc _GraphClass) New() Graph {
	rv := objc.Call[Graph](gc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewGraph() Graph {
	return GraphClass.New()
}

func (g_ Graph) Init() Graph {
	rv := objc.Call[Graph](g_, objc.Sel("init"))
	return rv
}

func (gc _GraphClass) Alloc() Graph {
	rv := objc.Call[Graph](gc, objc.Sel("alloc"))
	return rv
}

func Graph_Alloc() Graph {
	return GraphClass.Alloc()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564708-maxpooling2dwithsourcetensor?language=objc
func (g_ Graph) MaxPooling2DWithSourceTensorDescriptorName(source ITensor, descriptor IPooling2DOpDescriptor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("maxPooling2DWithSourceTensor:descriptor:name:"), objc.Ptr(source), objc.Ptr(descriptor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3750693-avgpooling4dwithsourcetensor?language=objc
func (g_ Graph) AvgPooling4DWithSourceTensorDescriptorName(source ITensor, descriptor IPooling4DOpDescriptor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("avgPooling4DWithSourceTensor:descriptor:name:"), objc.Ptr(source), objc.Ptr(descriptor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564559-lessthanorequaltowithprimarytens?language=objc
func (g_ Graph) LessThanOrEqualToWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("lessThanOrEqualToWithPrimaryTensor:secondaryTensor:name:"), objc.Ptr(primaryTensor), objc.Ptr(secondaryTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3667487-depthwiseconvolution2ddatagradie?language=objc
func (g_ Graph) DepthwiseConvolution2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeDescriptorName(incomingGradient ITensor, weights ITensor, outputShape *foundation.Array, descriptor IDepthwiseConvolution2DOpDescriptor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("depthwiseConvolution2DDataGradientWithIncomingGradientTensor:weightsTensor:outputShape:descriptor:name:"), objc.Ptr(incomingGradient), objc.Ptr(weights), outputShape, objc.Ptr(descriptor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3763057-sparsetensorwithdescriptor?language=objc
func (g_ Graph) SparseTensorWithDescriptorTensorsShapeName(sparseDescriptor ICreateSparseOpDescriptor, inputTensorArray []ITensor, shape *foundation.Array, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("sparseTensorWithDescriptor:tensors:shape:name:"), objc.Ptr(sparseDescriptor), inputTensorArray, shape, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3925432-lstmwithsourcetensor?language=objc
func (g_ Graph) LSTMWithSourceTensorRecurrentWeightInitStateInitCellDescriptorName(source ITensor, recurrentWeight ITensor, initState ITensor, initCell ITensor, descriptor ILSTMDescriptor, name string) []Tensor {
	rv := objc.Call[[]Tensor](g_, objc.Sel("LSTMWithSourceTensor:recurrentWeight:initState:initCell:descriptor:name:"), objc.Ptr(source), objc.Ptr(recurrentWeight), objc.Ptr(initState), objc.Ptr(initCell), objc.Ptr(descriptor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3867067-convolution2ddatagradientwithinc?language=objc
func (g_ Graph) Convolution2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeTensorForwardConvolutionDescriptorName(gradient ITensor, weights ITensor, outputShapeTensor ITensor, forwardConvolutionDescriptor IConvolution2DOpDescriptor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("convolution2DDataGradientWithIncomingGradientTensor:weightsTensor:outputShapeTensor:forwardConvolutionDescriptor:name:"), objc.Ptr(gradient), objc.Ptr(weights), objc.Ptr(outputShapeTensor), objc.Ptr(forwardConvolutionDescriptor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564622-runwithmtlcommandqueue?language=objc
func (g_ Graph) RunWithMTLCommandQueueFeedsTargetOperationsResultsDictionary(commandQueue metal.PCommandQueue, feeds *foundation.Dictionary, targetOperations []IOperation, resultsDictionary *foundation.Dictionary) {
	po0 := objc.WrapAsProtocol("MTLCommandQueue", commandQueue)
	objc.Call[objc.Void](g_, objc.Sel("runWithMTLCommandQueue:feeds:targetOperations:resultsDictionary:"), po0, feeds, targetOperations, resultsDictionary)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564622-runwithmtlcommandqueue?language=objc
func (g_ Graph) RunWithMTLCommandQueueObjectFeedsTargetOperationsResultsDictionary(commandQueueObject objc.IObject, feeds *foundation.Dictionary, targetOperations []IOperation, resultsDictionary *foundation.Dictionary) {
	objc.Call[objc.Void](g_, objc.Sel("runWithMTLCommandQueue:feeds:targetOperations:resultsDictionary:"), objc.Ptr(commandQueueObject), feeds, targetOperations, resultsDictionary)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3925453-expanddimsoftensor?language=objc
func (g_ Graph) ExpandDimsOfTensorAxesTensorName(tensor ITensor, axesTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("expandDimsOfTensor:axesTensor:name:"), objc.Ptr(tensor), objc.Ptr(axesTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3919778-reductionandwithtensor?language=objc
func (g_ Graph) ReductionAndWithTensorAxisName(tensor ITensor, axis int, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("reductionAndWithTensor:axis:name:"), objc.Ptr(tensor), axis, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3867072-resizetensor?language=objc
func (g_ Graph) ResizeTensorSizeTensorModeCenterResultAlignCornersLayoutName(imagesTensor ITensor, size ITensor, mode ResizeMode, centerResult bool, alignCorners bool, layout TensorNamedDataLayout, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("resizeTensor:sizeTensor:mode:centerResult:alignCorners:layout:name:"), objc.Ptr(imagesTensor), objc.Ptr(size), mode, centerResult, alignCorners, layout, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564591-tanwithtensor?language=objc
func (g_ Graph) TanWithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("tanWithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3675594-floormodulowithprimarytensor?language=objc
func (g_ Graph) FloorModuloWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("floorModuloWithPrimaryTensor:secondaryTensor:name:"), objc.Ptr(primaryTensor), objc.Ptr(secondaryTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3867070-convolutiontranspose2dweightsgra?language=objc
func (g_ Graph) ConvolutionTranspose2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeTensorForwardConvolutionDescriptorName(incomingGradientTensor ITensor, source ITensor, outputShape ITensor, forwardConvolutionDescriptor IConvolution2DOpDescriptor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("convolutionTranspose2DWeightsGradientWithIncomingGradientTensor:sourceTensor:outputShapeTensor:forwardConvolutionDescriptor:name:"), objc.Ptr(incomingGradientTensor), objc.Ptr(source), objc.Ptr(outputShape), objc.Ptr(forwardConvolutionDescriptor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3925451-coordinatealongaxistensor?language=objc
func (g_ Graph) CoordinateAlongAxisTensorWithShapeTensorName(axisTensor ITensor, shapeTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("coordinateAlongAxisTensor:withShapeTensor:name:"), objc.Ptr(axisTensor), objc.Ptr(shapeTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3589368-reductionproductwithtensor?language=objc
func (g_ Graph) ReductionProductWithTensorAxesName(tensor ITensor, axes []foundation.INumber, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("reductionProductWithTensor:axes:name:"), objc.Ptr(tensor), axes, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564686-assignvariable?language=objc
func (g_ Graph) AssignVariableWithValueOfTensorName(variable ITensor, tensor ITensor, name string) Operation {
	rv := objc.Call[Operation](g_, objc.Sel("assignVariable:withValueOfTensor:name:"), objc.Ptr(variable), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3901487-randomuniformtensorwithshapetens?language=objc
func (g_ Graph) RandomUniformTensorWithShapeTensorName(shapeTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("randomUniformTensorWithShapeTensor:name:"), objc.Ptr(shapeTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3580502-padtensor?language=objc
func (g_ Graph) PadTensorWithPaddingModeLeftPaddingRightPaddingConstantValueName(tensor ITensor, paddingMode PaddingMode, leftPadding *foundation.Array, rightPadding *foundation.Array, constantValue float64, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("padTensor:withPaddingMode:leftPadding:rightPadding:constantValue:name:"), objc.Ptr(tensor), paddingMode, leftPadding, rightPadding, constantValue, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3618936-applystochasticgradientdescentwi?language=objc
func (g_ Graph) ApplyStochasticGradientDescentWithLearningRateTensorVariableGradientTensorName(learningRateTensor ITensor, variable IVariableOp, gradientTensor ITensor, name string) Operation {
	rv := objc.Call[Operation](g_, objc.Sel("applyStochasticGradientDescentWithLearningRateTensor:variable:gradientTensor:name:"), objc.Ptr(learningRateTensor), objc.Ptr(variable), objc.Ptr(gradientTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564583-selectwithpredicatetensor?language=objc
func (g_ Graph) SelectWithPredicateTensorTruePredicateTensorFalsePredicateTensorName(predicateTensor ITensor, truePredicateTensor ITensor, falseSelectTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("selectWithPredicateTensor:truePredicateTensor:falsePredicateTensor:name:"), objc.Ptr(predicateTensor), objc.Ptr(truePredicateTensor), objc.Ptr(falseSelectTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3750680-depthwiseconvolution3dweightsgra?language=objc
func (g_ Graph) DepthwiseConvolution3DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeDescriptorName(incomingGradient ITensor, source ITensor, outputShape *foundation.Array, descriptor IDepthwiseConvolution3DOpDescriptor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("depthwiseConvolution3DWeightsGradientWithIncomingGradientTensor:sourceTensor:outputShape:descriptor:name:"), objc.Ptr(incomingGradient), objc.Ptr(source), outputShape, objc.Ptr(descriptor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3867065-leakyreluwithtensor?language=objc
func (g_ Graph) LeakyReLUWithTensorAlphaName(tensor ITensor, alpha float64, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("leakyReLUWithTensor:alpha:name:"), objc.Ptr(tensor), alpha, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564566-logicalnorwithprimarytensor?language=objc
func (g_ Graph) LogicalNORWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("logicalNORWithPrimaryTensor:secondaryTensor:name:"), objc.Ptr(primaryTensor), objc.Ptr(secondaryTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3867073-scatterndwithdatatensor?language=objc
func (g_ Graph) ScatterNDWithDataTensorUpdatesTensorIndicesTensorBatchDimensionsModeName(dataTensor ITensor, updatesTensor ITensor, indicesTensor ITensor, batchDimensions uint, mode ScatterMode, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("scatterNDWithDataTensor:updatesTensor:indicesTensor:batchDimensions:mode:name:"), objc.Ptr(dataTensor), objc.Ptr(updatesTensor), objc.Ptr(indicesTensor), batchDimensions, mode, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564589-squarewithtensor?language=objc
func (g_ Graph) SquareWithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("squareWithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3816770-forloopwithnumberofiterations?language=objc
func (g_ Graph) ForLoopWithNumberOfIterationsInitialBodyArgumentsBodyName(numberOfIterations ITensor, initialBodyArguments []ITensor, body ForLoopBodyBlock, name string) []Tensor {
	rv := objc.Call[[]Tensor](g_, objc.Sel("forLoopWithNumberOfIterations:initialBodyArguments:body:name:"), objc.Ptr(numberOfIterations), initialBodyArguments, body, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564567-logicalorwithprimarytensor?language=objc
func (g_ Graph) LogicalORWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("logicalORWithPrimaryTensor:secondaryTensor:name:"), objc.Ptr(primaryTensor), objc.Ptr(secondaryTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3919780-reductionorwithtensor?language=objc
func (g_ Graph) ReductionOrWithTensorAxisName(tensor ITensor, axis int, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("reductionOrWithTensor:axis:name:"), objc.Ptr(tensor), axis, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3750709-depthtospace2dtensor?language=objc
func (g_ Graph) DepthToSpace2DTensorWidthAxisHeightAxisDepthAxisBlockSizeUsePixelShuffleOrderName(tensor ITensor, widthAxis uint, heightAxis uint, depthAxis uint, blockSize uint, usePixelShuffleOrder bool, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("depthToSpace2DTensor:widthAxis:heightAxis:depthAxis:blockSize:usePixelShuffleOrder:name:"), objc.Ptr(tensor), widthAxis, heightAxis, depthAxis, blockSize, usePixelShuffleOrder, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3750694-maxpooling4dgradientwithgradient?language=objc
func (g_ Graph) MaxPooling4DGradientWithGradientTensorSourceTensorDescriptorName(gradient ITensor, source ITensor, descriptor IPooling4DOpDescriptor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("maxPooling4DGradientWithGradientTensor:sourceTensor:descriptor:name:"), objc.Ptr(gradient), objc.Ptr(source), objc.Ptr(descriptor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3656156-asinwithtensor?language=objc
func (g_ Graph) AsinWithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("asinWithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564471-sigmoidwithtensor?language=objc
func (g_ Graph) SigmoidWithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("sigmoidWithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3618939-tilegradientwithincominggradient?language=objc
func (g_ Graph) TileGradientWithIncomingGradientTensorSourceTensorWithMultiplierName(incomingGradientTensor ITensor, sourceTensor ITensor, multiplier *foundation.Array, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("tileGradientWithIncomingGradientTensor:sourceTensor:withMultiplier:name:"), objc.Ptr(incomingGradientTensor), objc.Ptr(sourceTensor), multiplier, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3867075-scatterwithupdatestensor?language=objc
func (g_ Graph) ScatterWithUpdatesTensorIndicesTensorShapeAxisModeName(updatesTensor ITensor, indicesTensor ITensor, shape *foundation.Array, axis int, mode ScatterMode, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("scatterWithUpdatesTensor:indicesTensor:shape:axis:mode:name:"), objc.Ptr(updatesTensor), objc.Ptr(indicesTensor), shape, axis, mode, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564548-erfwithtensor?language=objc
func (g_ Graph) ErfWithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("erfWithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564592-tanhwithtensor?language=objc
func (g_ Graph) TanhWithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("tanhWithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3925427-bandpartwithtensor?language=objc
func (g_ Graph) BandPartWithTensorNumLowerTensorNumUpperTensorName(inputTensor ITensor, numLowerTensor ITensor, numUpperTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("bandPartWithTensor:numLowerTensor:numUpperTensor:name:"), objc.Ptr(inputTensor), objc.Ptr(numLowerTensor), objc.Ptr(numUpperTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3667488-depthwiseconvolution2dweightsgra?language=objc
func (g_ Graph) DepthwiseConvolution2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeDescriptorName(incomingGradient ITensor, source ITensor, outputShape *foundation.Array, descriptor IDepthwiseConvolution2DOpDescriptor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("depthwiseConvolution2DWeightsGradientWithIncomingGradientTensor:sourceTensor:outputShape:descriptor:name:"), objc.Ptr(incomingGradient), objc.Ptr(source), outputShape, objc.Ptr(descriptor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3656161-softmaxcrossentropywithsourceten?language=objc
func (g_ Graph) SoftMaxCrossEntropyWithSourceTensorLabelsTensorAxisReductionTypeName(sourceTensor ITensor, labelsTensor ITensor, axis int, reductionType LossReductionType, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("softMaxCrossEntropyWithSourceTensor:labelsTensor:axis:reductionType:name:"), objc.Ptr(sourceTensor), objc.Ptr(labelsTensor), axis, reductionType, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564585-signbitwithtensor?language=objc
func (g_ Graph) SignbitWithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("signbitWithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3787605-stencilwithsourcetensor?language=objc
func (g_ Graph) StencilWithSourceTensorWeightsTensorDescriptorName(source ITensor, weights ITensor, descriptor IStencilOpDescriptor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("stencilWithSourceTensor:weightsTensor:descriptor:name:"), objc.Ptr(source), objc.Ptr(weights), objc.Ptr(descriptor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3750692-avgpooling4dgradientwithgradient?language=objc
func (g_ Graph) AvgPooling4DGradientWithGradientTensorSourceTensorDescriptorName(gradient ITensor, source ITensor, descriptor IPooling4DOpDescriptor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("avgPooling4DGradientWithGradientTensor:sourceTensor:descriptor:name:"), objc.Ptr(gradient), objc.Ptr(source), objc.Ptr(descriptor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3867068-convolution2dweightsgradientwith?language=objc
func (g_ Graph) Convolution2DWeightsGradientWithIncomingGradientTensorSourceTensorOutputShapeTensorForwardConvolutionDescriptorName(gradient ITensor, source ITensor, outputShapeTensor ITensor, forwardConvolutionDescriptor IConvolution2DOpDescriptor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("convolution2DWeightsGradientWithIncomingGradientTensor:sourceTensor:outputShapeTensor:forwardConvolutionDescriptor:name:"), objc.Ptr(gradient), objc.Ptr(source), objc.Ptr(outputShapeTensor), objc.Ptr(forwardConvolutionDescriptor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3901478-randomtensorwithshape?language=objc
func (g_ Graph) RandomTensorWithShapeDescriptorName(shape *foundation.Array, descriptor IRandomOpDescriptor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("randomTensorWithShape:descriptor:name:"), shape, objc.Ptr(descriptor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3656160-atanhwithtensor?language=objc
func (g_ Graph) AtanhWithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("atanhWithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3750707-reductionargminimumwithtensor?language=objc
func (g_ Graph) ReductionArgMinimumWithTensorAxisName(tensor ITensor, axis int, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("reductionArgMinimumWithTensor:axis:name:"), objc.Ptr(tensor), axis, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3867069-convolutiontranspose2ddatagradie?language=objc
func (g_ Graph) ConvolutionTranspose2DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeTensorForwardConvolutionDescriptorName(incomingGradient ITensor, weights ITensor, outputShape ITensor, forwardConvolutionDescriptor IConvolution2DOpDescriptor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("convolutionTranspose2DDataGradientWithIncomingGradientTensor:weightsTensor:outputShapeTensor:forwardConvolutionDescriptor:name:"), objc.Ptr(incomingGradient), objc.Ptr(weights), objc.Ptr(outputShape), objc.Ptr(forwardConvolutionDescriptor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564558-isnanwithtensor?language=objc
func (g_ Graph) IsNaNWithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("isNaNWithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564571-maximumwithprimarytensor?language=objc
func (g_ Graph) MaximumWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("maximumWithPrimaryTensor:secondaryTensor:name:"), objc.Ptr(primaryTensor), objc.Ptr(secondaryTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3919742-maxpooling2dreturnindiceswithsou?language=objc
func (g_ Graph) MaxPooling2DReturnIndicesWithSourceTensorDescriptorName(source ITensor, descriptor IPooling2DOpDescriptor, name string) []Tensor {
	rv := objc.Call[[]Tensor](g_, objc.Sel("maxPooling2DReturnIndicesWithSourceTensor:descriptor:name:"), objc.Ptr(source), objc.Ptr(descriptor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3750669-ifwithpredicatetensor?language=objc
func (g_ Graph) IfWithPredicateTensorThenBlockElseBlockName(predicateTensor ITensor, thenBlock IfThenElseBlock, elseBlock IfThenElseBlock, name string) []Tensor {
	rv := objc.Call[[]Tensor](g_, objc.Sel("ifWithPredicateTensor:thenBlock:elseBlock:name:"), objc.Ptr(predicateTensor), thenBlock, elseBlock, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564560-lessthanwithprimarytensor?language=objc
func (g_ Graph) LessThanWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("lessThanWithPrimaryTensor:secondaryTensor:name:"), objc.Ptr(primaryTensor), objc.Ptr(secondaryTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3928155-scatteralongaxis?language=objc
func (g_ Graph) ScatterAlongAxisWithDataTensorUpdatesTensorIndicesTensorModeName(axis int, dataTensor ITensor, updatesTensor ITensor, indicesTensor ITensor, mode ScatterMode, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("scatterAlongAxis:withDataTensor:updatesTensor:indicesTensor:mode:name:"), axis, objc.Ptr(dataTensor), objc.Ptr(updatesTensor), objc.Ptr(indicesTensor), mode, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564561-logarithmbase10withtensor?language=objc
func (g_ Graph) LogarithmBase10WithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("logarithmBase10WithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3750716-spacetodepth2dtensor?language=objc
func (g_ Graph) SpaceToDepth2DTensorWidthAxisHeightAxisDepthAxisBlockSizeUsePixelShuffleOrderName(tensor ITensor, widthAxis uint, heightAxis uint, depthAxis uint, blockSize uint, usePixelShuffleOrder bool, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("spaceToDepth2DTensor:widthAxis:heightAxis:depthAxis:blockSize:usePixelShuffleOrder:name:"), objc.Ptr(tensor), widthAxis, heightAxis, depthAxis, blockSize, usePixelShuffleOrder, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564541-additionwithprimarytensor?language=objc
func (g_ Graph) AdditionWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("additionWithPrimaryTensor:secondaryTensor:name:"), objc.Ptr(primaryTensor), objc.Ptr(secondaryTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564735-tiletensor?language=objc
func (g_ Graph) TileTensorWithMultiplierName(tensor ITensor, multiplier *foundation.Array, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("tileTensor:withMultiplier:name:"), objc.Ptr(tensor), multiplier, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564551-exponentwithtensor?language=objc
func (g_ Graph) ExponentWithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("exponentWithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3867078-casttensor?language=objc
func (g_ Graph) CastTensorToTypeName(tensor ITensor, type_ mps.DataType, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("castTensor:toType:name:"), objc.Ptr(tensor), type_, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3656162-meanoftensor?language=objc
func (g_ Graph) MeanOfTensorAxesName(tensor ITensor, axes []foundation.INumber, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("meanOfTensor:axes:name:"), objc.Ptr(tensor), axes, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564580-reversesquarerootwithtensor?language=objc
func (g_ Graph) ReverseSquareRootWithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("reverseSquareRootWithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3728124-scatterndwithupdatestensor?language=objc
func (g_ Graph) ScatterNDWithUpdatesTensorIndicesTensorShapeBatchDimensionsModeName(updatesTensor ITensor, indicesTensor ITensor, shape *foundation.Array, batchDimensions uint, mode ScatterMode, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("scatterNDWithUpdatesTensor:indicesTensor:shape:batchDimensions:mode:name:"), objc.Ptr(updatesTensor), objc.Ptr(indicesTensor), shape, batchDimensions, mode, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3925449-coordinatealongaxis?language=objc
func (g_ Graph) CoordinateAlongAxisWithShapeTensorName(axis int, shapeTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("coordinateAlongAxis:withShapeTensor:name:"), axis, objc.Ptr(shapeTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3928153-gatheralongaxis?language=objc
func (g_ Graph) GatherAlongAxisWithUpdatesTensorIndicesTensorName(axis int, updatesTensor ITensor, indicesTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("gatherAlongAxis:withUpdatesTensor:indicesTensor:name:"), axis, objc.Ptr(updatesTensor), objc.Ptr(indicesTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564569-logicalxorwithprimarytensor?language=objc
func (g_ Graph) LogicalXORWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("logicalXORWithPrimaryTensor:secondaryTensor:name:"), objc.Ptr(primaryTensor), objc.Ptr(secondaryTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564552-floorwithtensor?language=objc
func (g_ Graph) FloorWithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("floorWithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3750679-depthwiseconvolution3ddatagradie?language=objc
func (g_ Graph) DepthwiseConvolution3DDataGradientWithIncomingGradientTensorWeightsTensorOutputShapeDescriptorName(incomingGradient ITensor, weights ITensor, outputShape *foundation.Array, descriptor IDepthwiseConvolution3DOpDescriptor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("depthwiseConvolution3DDataGradientWithIncomingGradientTensor:weightsTensor:outputShape:descriptor:name:"), objc.Ptr(incomingGradient), objc.Ptr(weights), outputShape, objc.Ptr(descriptor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3925428-lstmgradientswithsourcetensor?language=objc
func (g_ Graph) LSTMGradientsWithSourceTensorRecurrentWeightSourceGradientZStateCellOutputFwdDescriptorName(source ITensor, recurrentWeight ITensor, sourceGradient ITensor, zState ITensor, cellOutputFwd ITensor, descriptor ILSTMDescriptor, name string) []Tensor {
	rv := objc.Call[[]Tensor](g_, objc.Sel("LSTMGradientsWithSourceTensor:recurrentWeight:sourceGradient:zState:cellOutputFwd:descriptor:name:"), objc.Ptr(source), objc.Ptr(recurrentWeight), objc.Ptr(sourceGradient), objc.Ptr(zState), objc.Ptr(cellOutputFwd), objc.Ptr(descriptor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3923675-splittensor?language=objc
func (g_ Graph) SplitTensorSplitSizesTensorAxisName(tensor ITensor, splitSizesTensor ITensor, axis int, name string) []Tensor {
	rv := objc.Call[[]Tensor](g_, objc.Sel("splitTensor:splitSizesTensor:axis:name:"), objc.Ptr(tensor), objc.Ptr(splitSizesTensor), axis, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3867082-topkwithgradienttensor?language=objc
func (g_ Graph) TopKWithGradientTensorSourceKTensorName(gradient ITensor, source ITensor, kTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("topKWithGradientTensor:source:kTensor:name:"), objc.Ptr(gradient), objc.Ptr(source), objc.Ptr(kTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564556-isfinitewithtensor?language=objc
func (g_ Graph) IsFiniteWithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("isFiniteWithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564590-subtractionwithprimarytensor?language=objc
func (g_ Graph) SubtractionWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("subtractionWithPrimaryTensor:secondaryTensor:name:"), objc.Ptr(primaryTensor), objc.Ptr(secondaryTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564690-readvariable?language=objc
func (g_ Graph) ReadVariableName(variable ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("readVariable:name:"), objc.Ptr(variable), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3867064-leakyrelugradientwithincominggra?language=objc
func (g_ Graph) LeakyReLUGradientWithIncomingGradientSourceTensorAlphaTensorName(gradient ITensor, source ITensor, alphaTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("leakyReLUGradientWithIncomingGradient:sourceTensor:alphaTensor:name:"), objc.Ptr(gradient), objc.Ptr(source), objc.Ptr(alphaTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3925458-squeezetensor?language=objc
func (g_ Graph) SqueezeTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("squeezeTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3618935-normalizationgradientwithincomin?language=objc
func (g_ Graph) NormalizationGradientWithIncomingGradientTensorSourceTensorMeanTensorVarianceTensorGammaTensorGammaGradientTensorBetaGradientTensorReductionAxesEpsilonName(incomingGradientTensor ITensor, sourceTensor ITensor, meanTensor ITensor, varianceTensor ITensor, gamma ITensor, gammaGradient ITensor, betaGradient ITensor, axes []foundation.INumber, epsilon float64, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("normalizationGradientWithIncomingGradientTensor:sourceTensor:meanTensor:varianceTensor:gammaTensor:gammaGradientTensor:betaGradientTensor:reductionAxes:epsilon:name:"), objc.Ptr(incomingGradientTensor), objc.Ptr(sourceTensor), objc.Ptr(meanTensor), objc.Ptr(varianceTensor), objc.Ptr(gamma), objc.Ptr(gammaGradient), objc.Ptr(betaGradient), axes, epsilon, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564596-convolution2dwithsourcetensor?language=objc
func (g_ Graph) Convolution2DWithSourceTensorWeightsTensorDescriptorName(source ITensor, weights ITensor, descriptor IConvolution2DOpDescriptor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("convolution2DWithSourceTensor:weightsTensor:descriptor:name:"), objc.Ptr(source), objc.Ptr(weights), objc.Ptr(descriptor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564468-relugradientwithincominggradient?language=objc
func (g_ Graph) ReLUGradientWithIncomingGradientSourceTensorName(gradient ITensor, source ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("reLUGradientWithIncomingGradient:sourceTensor:name:"), objc.Ptr(gradient), objc.Ptr(source), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3750712-flatten2dtensor?language=objc
func (g_ Graph) Flatten2DTensorAxisTensorName(tensor ITensor, axisTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("flatten2DTensor:axisTensor:name:"), objc.Ptr(tensor), objc.Ptr(axisTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3580486-onehotwithindicestensor?language=objc
func (g_ Graph) OneHotWithIndicesTensorDepthName(indicesTensor ITensor, depth uint, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("oneHotWithIndicesTensor:depth:name:"), objc.Ptr(indicesTensor), depth, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3750708-concattensors?language=objc
func (g_ Graph) ConcatTensorsDimensionInterleaveName(tensors []ITensor, dimensionIndex int, interleave bool, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("concatTensors:dimension:interleave:name:"), tensors, dimensionIndex, interleave, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3589367-reductionminimumwithtensor?language=objc
func (g_ Graph) ReductionMinimumWithTensorAxesName(tensor ITensor, axes []foundation.INumber, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("reductionMinimumWithTensor:axes:name:"), objc.Ptr(tensor), axes, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3901509-reductionmaximumpropagatenanwith?language=objc
func (g_ Graph) ReductionMaximumPropagateNaNWithTensorAxisName(tensor ITensor, axis int, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("reductionMaximumPropagateNaNWithTensor:axis:name:"), objc.Ptr(tensor), axis, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564581-rintwithtensor?language=objc
func (g_ Graph) RintWithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("rintWithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3656158-atan2withprimarytensor?language=objc
func (g_ Graph) Atan2WithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("atan2WithPrimaryTensor:secondaryTensor:name:"), objc.Ptr(primaryTensor), objc.Ptr(secondaryTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564544-coswithtensor?language=objc
func (g_ Graph) CosWithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("cosWithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3750695-maxpooling4dwithsourcetensor?language=objc
func (g_ Graph) MaxPooling4DWithSourceTensorDescriptorName(source ITensor, descriptor IPooling4DOpDescriptor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("maxPooling4DWithSourceTensor:descriptor:name:"), objc.Ptr(source), objc.Ptr(descriptor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564619-runasyncwithmtlcommandqueue?language=objc
func (g_ Graph) RunAsyncWithMTLCommandQueueFeedsTargetOperationsResultsDictionaryExecutionDescriptor(commandQueue metal.PCommandQueue, feeds *foundation.Dictionary, targetOperations []IOperation, resultsDictionary *foundation.Dictionary, executionDescriptor IExecutionDescriptor) {
	po0 := objc.WrapAsProtocol("MTLCommandQueue", commandQueue)
	objc.Call[objc.Void](g_, objc.Sel("runAsyncWithMTLCommandQueue:feeds:targetOperations:resultsDictionary:executionDescriptor:"), po0, feeds, targetOperations, resultsDictionary, objc.Ptr(executionDescriptor))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564619-runasyncwithmtlcommandqueue?language=objc
func (g_ Graph) RunAsyncWithMTLCommandQueueObjectFeedsTargetOperationsResultsDictionaryExecutionDescriptor(commandQueueObject objc.IObject, feeds *foundation.Dictionary, targetOperations []IOperation, resultsDictionary *foundation.Dictionary, executionDescriptor IExecutionDescriptor) {
	objc.Call[objc.Void](g_, objc.Sel("runAsyncWithMTLCommandQueue:feeds:targetOperations:resultsDictionary:executionDescriptor:"), objc.Ptr(commandQueueObject), feeds, targetOperations, resultsDictionary, objc.Ptr(executionDescriptor))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3816769-forloopwithlowerbound?language=objc
func (g_ Graph) ForLoopWithLowerBoundUpperBoundStepInitialBodyArgumentsBodyName(lowerBound ITensor, upperBound ITensor, step ITensor, initialBodyArguments []ITensor, body ForLoopBodyBlock, name string) []Tensor {
	rv := objc.Call[[]Tensor](g_, objc.Sel("forLoopWithLowerBound:upperBound:step:initialBodyArguments:body:name:"), objc.Ptr(lowerBound), objc.Ptr(upperBound), objc.Ptr(step), initialBodyArguments, body, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564469-reluwithtensor?language=objc
func (g_ Graph) ReLUWithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("reLUWithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564470-sigmoidgradientwithincominggradi?language=objc
func (g_ Graph) SigmoidGradientWithIncomingGradientSourceTensorName(gradient ITensor, source ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("sigmoidGradientWithIncomingGradient:sourceTensor:name:"), objc.Ptr(gradient), objc.Ptr(source), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3728133-slicegradienttensor?language=objc
func (g_ Graph) SliceGradientTensorFwdInShapeTensorStartsEndsStridesName(inputGradientTensor ITensor, fwdInShapeTensor ITensor, starts []foundation.INumber, ends []foundation.INumber, strides []foundation.INumber, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("sliceGradientTensor:fwdInShapeTensor:starts:ends:strides:name:"), objc.Ptr(inputGradientTensor), objc.Ptr(fwdInShapeTensor), starts, ends, strides, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564706-avgpooling2dwithsourcetensor?language=objc
func (g_ Graph) AvgPooling2DWithSourceTensorDescriptorName(source ITensor, descriptor IPooling2DOpDescriptor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("avgPooling2DWithSourceTensor:descriptor:name:"), objc.Ptr(source), objc.Ptr(descriptor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3686986-constantwithscalar?language=objc
func (g_ Graph) ConstantWithScalarDataType(scalar float64, dataType mps.DataType) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("constantWithScalar:dataType:"), scalar, dataType)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3867079-reshapetensor?language=objc
func (g_ Graph) ReshapeTensorWithShapeTensorName(tensor ITensor, shapeTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("reshapeTensor:withShapeTensor:name:"), objc.Ptr(tensor), objc.Ptr(shapeTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564684-matrixmultiplicationwithprimaryt?language=objc
func (g_ Graph) MatrixMultiplicationWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("matrixMultiplicationWithPrimaryTensor:secondaryTensor:name:"), objc.Ptr(primaryTensor), objc.Ptr(secondaryTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3750667-controldependencywithoperations?language=objc
func (g_ Graph) ControlDependencyWithOperationsDependentBlockName(operations []IOperation, dependentBlock ControlFlowDependencyBlock, name string) []Tensor {
	rv := objc.Call[[]Tensor](g_, objc.Sel("controlDependencyWithOperations:dependentBlock:name:"), operations, dependentBlock, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564568-logicalxnorwithprimarytensor?language=objc
func (g_ Graph) LogicalXNORWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("logicalXNORWithPrimaryTensor:secondaryTensor:name:"), objc.Ptr(primaryTensor), objc.Ptr(secondaryTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3589364-gatherwithupdatestensor?language=objc
func (g_ Graph) GatherWithUpdatesTensorIndicesTensorAxisBatchDimensionsName(updatesTensor ITensor, indicesTensor ITensor, axis uint, batchDimensions uint, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("gatherWithUpdatesTensor:indicesTensor:axis:batchDimensions:name:"), objc.Ptr(updatesTensor), objc.Ptr(indicesTensor), axis, batchDimensions, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564545-coshwithtensor?language=objc
func (g_ Graph) CoshWithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("coshWithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564705-avgpooling2dgradientwithgradient?language=objc
func (g_ Graph) AvgPooling2DGradientWithGradientTensorSourceTensorDescriptorName(gradient ITensor, source ITensor, descriptor IPooling2DOpDescriptor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("avgPooling2DGradientWithGradientTensor:sourceTensor:descriptor:name:"), objc.Ptr(gradient), objc.Ptr(source), objc.Ptr(descriptor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3901511-reductionminimumpropagatenanwith?language=objc
func (g_ Graph) ReductionMinimumPropagateNaNWithTensorAxisName(tensor ITensor, axis int, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("reductionMinimumPropagateNaNWithTensor:axis:name:"), objc.Ptr(tensor), axis, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3901484-randomuniformtensorwithshape?language=objc
func (g_ Graph) RandomUniformTensorWithShapeName(shape *foundation.Array, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("randomUniformTensorWithShape:name:"), shape, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564563-logarithmwithtensor?language=objc
func (g_ Graph) LogarithmWithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("logarithmWithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3763058-sparsetensorwithtype?language=objc
func (g_ Graph) SparseTensorWithTypeTensorsShapeDataTypeName(sparseStorageType SparseStorageType, inputTensorArray []ITensor, shape *foundation.Array, dataType mps.DataType, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("sparseTensorWithType:tensors:shape:dataType:name:"), sparseStorageType, inputTensorArray, shape, dataType, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564540-absolutewithtensor?language=objc
func (g_ Graph) AbsoluteWithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("absoluteWithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3867071-convolutiontranspose2dwithsource?language=objc
func (g_ Graph) ConvolutionTranspose2DWithSourceTensorWeightsTensorOutputShapeTensorDescriptorName(source ITensor, weights ITensor, outputShape ITensor, descriptor IConvolution2DOpDescriptor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("convolutionTranspose2DWithSourceTensor:weightsTensor:outputShapeTensor:descriptor:name:"), objc.Ptr(source), objc.Ptr(weights), objc.Ptr(outputShape), objc.Ptr(descriptor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3675593-divisionnonanwithprimarytensor?language=objc
func (g_ Graph) DivisionNoNaNWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("divisionNoNaNWithPrimaryTensor:secondaryTensor:name:"), objc.Ptr(primaryTensor), objc.Ptr(secondaryTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3656154-acoswithtensor?language=objc
func (g_ Graph) AcosWithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("acosWithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3919743-maxpooling4dreturnindiceswithsou?language=objc
func (g_ Graph) MaxPooling4DReturnIndicesWithSourceTensorDescriptorName(source ITensor, descriptor IPooling4DOpDescriptor, name string) []Tensor {
	rv := objc.Call[[]Tensor](g_, objc.Sel("maxPooling4DReturnIndicesWithSourceTensor:descriptor:name:"), objc.Ptr(source), objc.Ptr(descriptor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3901481-randomtensorwithshapetensor?language=objc
func (g_ Graph) RandomTensorWithShapeTensorDescriptorName(shapeTensor ITensor, descriptor IRandomOpDescriptor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("randomTensorWithShapeTensor:descriptor:name:"), objc.Ptr(shapeTensor), objc.Ptr(descriptor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564687-constantwithdata?language=objc
func (g_ Graph) ConstantWithDataShapeDataType(data []byte, shape *foundation.Array, dataType mps.DataType) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("constantWithData:shape:dataType:"), data, shape, dataType)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3901477-randomphiloxstatetensorwithseed?language=objc
func (g_ Graph) RandomPhiloxStateTensorWithSeedName(seed uint, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("randomPhiloxStateTensorWithSeed:name:"), seed, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3750690-l2normpooling4dgradientwithgradi?language=objc
func (g_ Graph) L2NormPooling4DGradientWithGradientTensorSourceTensorDescriptorName(gradient ITensor, source ITensor, descriptor IPooling4DOpDescriptor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("L2NormPooling4DGradientWithGradientTensor:sourceTensor:descriptor:name:"), objc.Ptr(gradient), objc.Ptr(source), objc.Ptr(descriptor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3667506-resizewithgradienttensor?language=objc
func (g_ Graph) ResizeWithGradientTensorInputModeCenterResultAlignCornersLayoutName(gradient ITensor, input ITensor, mode ResizeMode, centerResult bool, alignCorners bool, layout TensorNamedDataLayout, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("resizeWithGradientTensor:input:mode:centerResult:alignCorners:layout:name:"), objc.Ptr(gradient), objc.Ptr(input), mode, centerResult, alignCorners, layout, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564691-variablewithdata?language=objc
func (g_ Graph) VariableWithDataShapeDataTypeName(data []byte, shape *foundation.Array, dataType mps.DataType, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("variableWithData:shape:dataType:name:"), data, shape, dataType, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564553-greaterthanorequaltowithprimaryt?language=objc
func (g_ Graph) GreaterThanOrEqualToWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("greaterThanOrEqualToWithPrimaryTensor:secondaryTensor:name:"), objc.Ptr(primaryTensor), objc.Ptr(secondaryTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3580451-gradientforprimarytensor?language=objc
func (g_ Graph) GradientForPrimaryTensorWithTensorsName(primaryTensor ITensor, tensors []ITensor, name string) foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](g_, objc.Sel("gradientForPrimaryTensor:withTensors:name:"), objc.Ptr(primaryTensor), tensors, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564542-ceilwithtensor?language=objc
func (g_ Graph) CeilWithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("ceilWithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3750715-reversetensor?language=objc
func (g_ Graph) ReverseTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("reverseTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564550-exponentbase2withtensor?language=objc
func (g_ Graph) ExponentBase2WithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("exponentBase2WithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564572-minimumwithprimarytensor?language=objc
func (g_ Graph) MinimumWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("minimumWithPrimaryTensor:secondaryTensor:name:"), objc.Ptr(primaryTensor), objc.Ptr(secondaryTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564557-isinfinitewithtensor?language=objc
func (g_ Graph) IsInfiniteWithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("isInfiniteWithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3750681-depthwiseconvolution3dwithsource?language=objc
func (g_ Graph) DepthwiseConvolution3DWithSourceTensorWeightsTensorDescriptorName(source ITensor, weights ITensor, descriptor IDepthwiseConvolution3DOpDescriptor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("depthwiseConvolution3DWithSourceTensor:weightsTensor:descriptor:name:"), objc.Ptr(source), objc.Ptr(weights), objc.Ptr(descriptor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3867084-topkwithsourcetensor?language=objc
func (g_ Graph) TopKWithSourceTensorKTensorName(source ITensor, kTensor ITensor, name string) []Tensor {
	rv := objc.Call[[]Tensor](g_, objc.Sel("topKWithSourceTensor:kTensor:name:"), objc.Ptr(source), objc.Ptr(kTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3589362-gatherndwithupdatestensor?language=objc
func (g_ Graph) GatherNDWithUpdatesTensorIndicesTensorBatchDimensionsName(updatesTensor ITensor, indicesTensor ITensor, batchDimensions uint, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("gatherNDWithUpdatesTensor:indicesTensor:batchDimensions:name:"), objc.Ptr(updatesTensor), objc.Ptr(indicesTensor), batchDimensions, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564618-runasyncwithfeeds?language=objc
func (g_ Graph) RunAsyncWithFeedsTargetTensorsTargetOperationsExecutionDescriptor(feeds *foundation.Dictionary, targetTensors []ITensor, targetOperations []IOperation, executionDescriptor IExecutionDescriptor) *foundation.Dictionary {
	rv := objc.Call[*foundation.Dictionary](g_, objc.Sel("runAsyncWithFeeds:targetTensors:targetOperations:executionDescriptor:"), feeds, targetTensors, targetOperations, objc.Ptr(executionDescriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564574-multiplicationwithprimarytensor?language=objc
func (g_ Graph) MultiplicationWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("multiplicationWithPrimaryTensor:secondaryTensor:name:"), objc.Ptr(primaryTensor), objc.Ptr(secondaryTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564587-sinhwithtensor?language=objc
func (g_ Graph) SinhWithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("sinhWithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3626198-dropouttensor?language=objc
func (g_ Graph) DropoutTensorRateName(tensor ITensor, rate float64, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("dropoutTensor:rate:name:"), objc.Ptr(tensor), rate, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3919762-singlegaternnwithsourcetensor?language=objc
func (g_ Graph) SingleGateRNNWithSourceTensorRecurrentWeightInitStateDescriptorName(source ITensor, recurrentWeight ITensor, initState ITensor, descriptor ISingleGateRNNDescriptor, name string) []Tensor {
	rv := objc.Call[[]Tensor](g_, objc.Sel("singleGateRNNWithSourceTensor:recurrentWeight:initState:descriptor:name:"), objc.Ptr(source), objc.Ptr(recurrentWeight), objc.Ptr(initState), objc.Ptr(descriptor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3857574-minimumwithnanpropagationwithpri?language=objc
func (g_ Graph) MinimumWithNaNPropagationWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("minimumWithNaNPropagationWithPrimaryTensor:secondaryTensor:name:"), objc.Ptr(primaryTensor), objc.Ptr(secondaryTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564562-logarithmbase2withtensor?language=objc
func (g_ Graph) LogarithmBase2WithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("logarithmBase2WithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564689-placeholderwithshape?language=objc
func (g_ Graph) PlaceholderWithShapeName(shape *foundation.Array, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("placeholderWithShape:name:"), shape, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564575-negativewithtensor?language=objc
func (g_ Graph) NegativeWithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("negativeWithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3867080-shapeoftensor?language=objc
func (g_ Graph) ShapeOfTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("shapeOfTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564549-exponentbase10withtensor?language=objc
func (g_ Graph) ExponentBase10WithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("exponentBase10WithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564621-runwithfeeds?language=objc
func (g_ Graph) RunWithFeedsTargetTensorsTargetOperations(feeds *foundation.Dictionary, targetTensors []ITensor, targetOperations []IOperation) *foundation.Dictionary {
	rv := objc.Call[*foundation.Dictionary](g_, objc.Sel("runWithFeeds:targetTensors:targetOperations:"), feeds, targetTensors, targetOperations)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3857573-maximumwithnanpropagationwithpri?language=objc
func (g_ Graph) MaximumWithNaNPropagationWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("maximumWithNaNPropagationWithPrimaryTensor:secondaryTensor:name:"), objc.Ptr(primaryTensor), objc.Ptr(secondaryTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3867074-scatterwithdatatensor?language=objc
func (g_ Graph) ScatterWithDataTensorUpdatesTensorIndicesTensorAxisModeName(dataTensor ITensor, updatesTensor ITensor, indicesTensor ITensor, axis int, mode ScatterMode, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("scatterWithDataTensor:updatesTensor:indicesTensor:axis:mode:name:"), objc.Ptr(dataTensor), objc.Ptr(updatesTensor), objc.Ptr(indicesTensor), axis, mode, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564554-greaterthanwithprimarytensor?language=objc
func (g_ Graph) GreaterThanWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("greaterThanWithPrimaryTensor:secondaryTensor:name:"), objc.Ptr(primaryTensor), objc.Ptr(secondaryTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564582-roundwithtensor?language=objc
func (g_ Graph) RoundWithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("roundWithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564564-logicalandwithprimarytensor?language=objc
func (g_ Graph) LogicalANDWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("logicalANDWithPrimaryTensor:secondaryTensor:name:"), objc.Ptr(primaryTensor), objc.Ptr(secondaryTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3589366-reductionmaximumwithtensor?language=objc
func (g_ Graph) ReductionMaximumWithTensorAxesName(tensor ITensor, axes []foundation.INumber, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("reductionMaximumWithTensor:axes:name:"), objc.Ptr(tensor), axes, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564584-signwithtensor?language=objc
func (g_ Graph) SignWithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("signWithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564547-equalwithprimarytensor?language=objc
func (g_ Graph) EqualWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("equalWithPrimaryTensor:secondaryTensor:name:"), objc.Ptr(primaryTensor), objc.Ptr(secondaryTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3656157-asinhwithtensor?language=objc
func (g_ Graph) AsinhWithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("asinhWithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3667489-depthwiseconvolution2dwithsource?language=objc
func (g_ Graph) DepthwiseConvolution2DWithSourceTensorWeightsTensorDescriptorName(source ITensor, weights ITensor, descriptor IDepthwiseConvolution2DOpDescriptor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("depthwiseConvolution2DWithSourceTensor:weightsTensor:descriptor:name:"), objc.Ptr(source), objc.Ptr(weights), objc.Ptr(descriptor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3656163-varianceoftensor?language=objc
func (g_ Graph) VarianceOfTensorAxesName(tensor ITensor, axes []foundation.INumber, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("varianceOfTensor:axes:name:"), objc.Ptr(tensor), axes, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3618938-padgradientwithincominggradientt?language=objc
func (g_ Graph) PadGradientWithIncomingGradientTensorSourceTensorPaddingModeLeftPaddingRightPaddingName(incomingGradientTensor ITensor, sourceTensor ITensor, paddingMode PaddingMode, leftPadding *foundation.Array, rightPadding *foundation.Array, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("padGradientWithIncomingGradientTensor:sourceTensor:paddingMode:leftPadding:rightPadding:name:"), objc.Ptr(incomingGradientTensor), objc.Ptr(sourceTensor), paddingMode, leftPadding, rightPadding, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564699-normalizationwithtensor?language=objc
func (g_ Graph) NormalizationWithTensorMeanTensorVarianceTensorGammaTensorBetaTensorEpsilonName(tensor ITensor, mean ITensor, variance ITensor, gamma ITensor, beta ITensor, epsilon float64, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("normalizationWithTensor:meanTensor:varianceTensor:gammaTensor:betaTensor:epsilon:name:"), objc.Ptr(tensor), objc.Ptr(mean), objc.Ptr(variance), objc.Ptr(gamma), objc.Ptr(beta), epsilon, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564546-divisionwithprimarytensor?language=objc
func (g_ Graph) DivisionWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("divisionWithPrimaryTensor:secondaryTensor:name:"), objc.Ptr(primaryTensor), objc.Ptr(secondaryTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3919758-singlegaternngradientswithsource?language=objc
func (g_ Graph) SingleGateRNNGradientsWithSourceTensorRecurrentWeightSourceGradientZStateInitStateDescriptorName(source ITensor, recurrentWeight ITensor, sourceGradient ITensor, zState ITensor, initState ITensor, descriptor ISingleGateRNNDescriptor, name string) []Tensor {
	rv := objc.Call[[]Tensor](g_, objc.Sel("singleGateRNNGradientsWithSourceTensor:recurrentWeight:sourceGradient:zState:initState:descriptor:name:"), objc.Ptr(source), objc.Ptr(recurrentWeight), objc.Ptr(sourceGradient), objc.Ptr(zState), objc.Ptr(initState), objc.Ptr(descriptor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3580383-encodetocommandbuffer?language=objc
func (g_ Graph) EncodeToCommandBufferFeedsTargetOperationsResultsDictionaryExecutionDescriptor(commandBuffer mps.ICommandBuffer, feeds *foundation.Dictionary, targetOperations []IOperation, resultsDictionary *foundation.Dictionary, executionDescriptor IExecutionDescriptor) {
	objc.Call[objc.Void](g_, objc.Sel("encodeToCommandBuffer:feeds:targetOperations:resultsDictionary:executionDescriptor:"), objc.Ptr(commandBuffer), feeds, targetOperations, resultsDictionary, objc.Ptr(executionDescriptor))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3787575-compilewithdevice?language=objc
func (g_ Graph) CompileWithDeviceFeedsTargetTensorsTargetOperationsCompilationDescriptor(device IDevice, feeds *foundation.Dictionary, targetTensors []ITensor, targetOperations []IOperation, compilationDescriptor ICompilationDescriptor) Executable {
	rv := objc.Call[Executable](g_, objc.Sel("compileWithDevice:feeds:targetTensors:targetOperations:compilationDescriptor:"), objc.Ptr(device), feeds, targetTensors, targetOperations, objc.Ptr(compilationDescriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3618937-stochasticgradientdescentwithlea?language=objc
func (g_ Graph) StochasticGradientDescentWithLearningRateTensorValuesTensorGradientTensorName(learningRateTensor ITensor, valuesTensor ITensor, gradientTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("stochasticGradientDescentWithLearningRateTensor:valuesTensor:gradientTensor:name:"), objc.Ptr(learningRateTensor), objc.Ptr(valuesTensor), objc.Ptr(gradientTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564579-reciprocalwithtensor?language=objc
func (g_ Graph) ReciprocalWithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("reciprocalWithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564588-squarerootwithtensor?language=objc
func (g_ Graph) SquareRootWithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("squareRootWithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564577-notwithtensor?language=objc
func (g_ Graph) NotWithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("notWithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564586-sinwithtensor?language=objc
func (g_ Graph) SinWithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("sinWithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564576-notequalwithprimarytensor?language=objc
func (g_ Graph) NotEqualWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("notEqualWithPrimaryTensor:secondaryTensor:name:"), objc.Ptr(primaryTensor), objc.Ptr(secondaryTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3728135-slicetensor?language=objc
func (g_ Graph) SliceTensorStartsEndsStridesName(tensor ITensor, starts []foundation.INumber, ends []foundation.INumber, strides []foundation.INumber, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("sliceTensor:starts:ends:strides:name:"), objc.Ptr(tensor), starts, ends, strides, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564578-powerwithprimarytensor?language=objc
func (g_ Graph) PowerWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("powerWithPrimaryTensor:secondaryTensor:name:"), objc.Ptr(primaryTensor), objc.Ptr(secondaryTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564565-logicalnandwithprimarytensor?language=objc
func (g_ Graph) LogicalNANDWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("logicalNANDWithPrimaryTensor:secondaryTensor:name:"), objc.Ptr(primaryTensor), objc.Ptr(secondaryTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3618933-normalizationbetagradientwithinc?language=objc
func (g_ Graph) NormalizationBetaGradientWithIncomingGradientTensorSourceTensorReductionAxesName(incomingGradientTensor ITensor, sourceTensor ITensor, axes []foundation.INumber, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("normalizationBetaGradientWithIncomingGradientTensor:sourceTensor:reductionAxes:name:"), objc.Ptr(incomingGradientTensor), objc.Ptr(sourceTensor), axes, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3656155-acoshwithtensor?language=objc
func (g_ Graph) AcoshWithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("acoshWithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3923676-stacktensors?language=objc
func (g_ Graph) StackTensorsAxisName(inputTensors []ITensor, axis int, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("stackTensors:axis:name:"), inputTensors, axis, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3580468-softmaxcrossentropygradientwithi?language=objc
func (g_ Graph) SoftMaxCrossEntropyGradientWithIncomingGradientTensorSourceTensorLabelsTensorAxisReductionTypeName(gradientTensor ITensor, sourceTensor ITensor, labelsTensor ITensor, axis int, reductionType LossReductionType, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("softMaxCrossEntropyGradientWithIncomingGradientTensor:sourceTensor:labelsTensor:axis:reductionType:name:"), objc.Ptr(gradientTensor), objc.Ptr(sourceTensor), objc.Ptr(labelsTensor), axis, reductionType, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3750706-reductionargmaximumwithtensor?language=objc
func (g_ Graph) ReductionArgMaximumWithTensorAxisName(tensor ITensor, axis int, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("reductionArgMaximumWithTensor:axis:name:"), objc.Ptr(tensor), axis, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564731-concattensor?language=objc
func (g_ Graph) ConcatTensorWithTensorDimensionName(tensor ITensor, tensor2 ITensor, dimensionIndex int, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("concatTensor:withTensor:dimension:name:"), objc.Ptr(tensor), objc.Ptr(tensor2), dimensionIndex, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3750691-l2normpooling4dwithsourcetensor?language=objc
func (g_ Graph) L2NormPooling4DWithSourceTensorDescriptorName(source ITensor, descriptor IPooling4DOpDescriptor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("L2NormPooling4DWithSourceTensor:descriptor:name:"), objc.Ptr(source), objc.Ptr(descriptor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3928154-gatheralongaxistensor?language=objc
func (g_ Graph) GatherAlongAxisTensorWithUpdatesTensorIndicesTensorName(axisTensor ITensor, updatesTensor ITensor, indicesTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("gatherAlongAxisTensor:withUpdatesTensor:indicesTensor:name:"), objc.Ptr(axisTensor), objc.Ptr(updatesTensor), objc.Ptr(indicesTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3928157-scatteralongaxistensor?language=objc
func (g_ Graph) ScatterAlongAxisTensorWithDataTensorUpdatesTensorIndicesTensorModeName(axisTensor ITensor, dataTensor ITensor, updatesTensor ITensor, indicesTensor ITensor, mode ScatterMode, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("scatterAlongAxisTensor:withDataTensor:updatesTensor:indicesTensor:mode:name:"), objc.Ptr(axisTensor), objc.Ptr(dataTensor), objc.Ptr(updatesTensor), objc.Ptr(indicesTensor), mode, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564736-transposetensor?language=objc
func (g_ Graph) TransposeTensorDimensionWithDimensionName(tensor ITensor, dimensionIndex uint, dimensionIndex2 uint, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("transposeTensor:dimension:withDimension:name:"), objc.Ptr(tensor), dimensionIndex, dimensionIndex2, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3867077-broadcasttensor?language=objc
func (g_ Graph) BroadcastTensorToShapeTensorName(tensor ITensor, shapeTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("broadcastTensor:toShapeTensor:name:"), objc.Ptr(tensor), objc.Ptr(shapeTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3580386-softmaxgradientwithincominggradi?language=objc
func (g_ Graph) SoftMaxGradientWithIncomingGradientSourceTensorAxisName(gradient ITensor, source ITensor, axis int, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("softMaxGradientWithIncomingGradient:sourceTensor:axis:name:"), objc.Ptr(gradient), objc.Ptr(source), axis, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564543-clampwithtensor?language=objc
func (g_ Graph) ClampWithTensorMinValueTensorMaxValueTensorName(tensor ITensor, minValueTensor ITensor, maxValueTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("clampWithTensor:minValueTensor:maxValueTensor:name:"), objc.Ptr(tensor), objc.Ptr(minValueTensor), objc.Ptr(maxValueTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564707-maxpooling2dgradientwithgradient?language=objc
func (g_ Graph) MaxPooling2DGradientWithGradientTensorSourceTensorDescriptorName(gradient ITensor, source ITensor, descriptor IPooling2DOpDescriptor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("maxPooling2DGradientWithGradientTensor:sourceTensor:descriptor:name:"), objc.Ptr(gradient), objc.Ptr(source), objc.Ptr(descriptor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3580494-reductionsumwithtensor?language=objc
func (g_ Graph) ReductionSumWithTensorAxesName(tensor ITensor, axes []foundation.INumber, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("reductionSumWithTensor:axes:name:"), objc.Ptr(tensor), axes, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564573-modulowithprimarytensor?language=objc
func (g_ Graph) ModuloWithPrimaryTensorSecondaryTensorName(primaryTensor ITensor, secondaryTensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("moduloWithPrimaryTensor:secondaryTensor:name:"), objc.Ptr(primaryTensor), objc.Ptr(secondaryTensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3901476-randomphiloxstatetensorwithcount?language=objc
func (g_ Graph) RandomPhiloxStateTensorWithCounterLowCounterHighKeyName(counterLow uint, counterHigh uint, key uint, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("randomPhiloxStateTensorWithCounterLow:counterHigh:key:name:"), counterLow, counterHigh, key, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564555-identitywithtensor?language=objc
func (g_ Graph) IdentityWithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("identityWithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3750670-whilewithinitialinputs?language=objc
func (g_ Graph) WhileWithInitialInputsBeforeAfterName(initialInputs []ITensor, before WhileBeforeBlock, after WhileAfterBlock, name string) []Tensor {
	rv := objc.Call[[]Tensor](g_, objc.Sel("whileWithInitialInputs:before:after:name:"), initialInputs, before, after, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3656159-atanwithtensor?language=objc
func (g_ Graph) AtanWithTensorName(tensor ITensor, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("atanWithTensor:name:"), objc.Ptr(tensor), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3618934-normalizationgammagradientwithin?language=objc
func (g_ Graph) NormalizationGammaGradientWithIncomingGradientTensorSourceTensorMeanTensorVarianceTensorReductionAxesEpsilonName(incomingGradientTensor ITensor, sourceTensor ITensor, meanTensor ITensor, varianceTensor ITensor, axes []foundation.INumber, epsilon float64, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("normalizationGammaGradientWithIncomingGradientTensor:sourceTensor:meanTensor:varianceTensor:reductionAxes:epsilon:name:"), objc.Ptr(incomingGradientTensor), objc.Ptr(sourceTensor), objc.Ptr(meanTensor), objc.Ptr(varianceTensor), axes, epsilon, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3580387-softmaxwithtensor?language=objc
func (g_ Graph) SoftMaxWithTensorAxisName(tensor ITensor, axis int, name string) Tensor {
	rv := objc.Call[Tensor](g_, objc.Sel("softMaxWithTensor:axis:name:"), objc.Ptr(tensor), axis, name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564616-options?language=objc
func (g_ Graph) Options() Options {
	rv := objc.Call[Options](g_, objc.Sel("options"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564616-options?language=objc
func (g_ Graph) SetOptions(value Options) {
	objc.Call[objc.Void](g_, objc.Sel("setOptions:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraph/3564617-placeholdertensors?language=objc
func (g_ Graph) PlaceholderTensors() []Tensor {
	rv := objc.Call[[]Tensor](g_, objc.Sel("placeholderTensors"))
	return rv
}
