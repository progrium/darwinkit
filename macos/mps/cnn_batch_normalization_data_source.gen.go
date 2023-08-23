// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// A protocol that defines methods that a batch normalization state uses to initialize scale factors, bias terms, and batch statistics. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationdatasource?language=objc
type PCNNBatchNormalizationDataSource interface {
	// optional
	UpdateGammaAndBetaWithCommandBufferBatchNormalizationState(commandBuffer metal.CommandBufferWrapper, batchNormalizationState CNNBatchNormalizationState) ICNNNormalizationGammaAndBetaState
	HasUpdateGammaAndBetaWithCommandBufferBatchNormalizationState() bool

	// optional
	Beta() *float64
	HasBeta() bool

	// optional
	Epsilon() float64
	HasEpsilon() bool

	// optional
	EncodeWithCoder(aCoder foundation.Coder)
	HasEncodeWithCoder() bool

	// optional
	UpdateGammaAndBetaWithBatchNormalizationState(batchNormalizationState CNNBatchNormalizationState) bool
	HasUpdateGammaAndBetaWithBatchNormalizationState() bool

	// optional
	InitWithCoder(aDecoder foundation.Coder) objc.IObject
	HasInitWithCoder() bool

	// optional
	Gamma() *float64
	HasGamma() bool

	// optional
	Purge()
	HasPurge() bool

	// optional
	UpdateMeanAndVarianceWithBatchNormalizationState(batchNormalizationState CNNBatchNormalizationState) bool
	HasUpdateMeanAndVarianceWithBatchNormalizationState() bool

	// optional
	Variance() *float64
	HasVariance() bool

	// optional
	CopyWithZoneDevice(zone unsafe.Pointer, device metal.DeviceWrapper) objc.IObject
	HasCopyWithZoneDevice() bool

	// optional
	Load() bool
	HasLoad() bool

	// optional
	Label() string
	HasLabel() bool

	// optional
	UpdateMeanAndVarianceWithCommandBufferBatchNormalizationState(commandBuffer metal.CommandBufferWrapper, batchNormalizationState CNNBatchNormalizationState) ICNNNormalizationMeanAndVarianceState
	HasUpdateMeanAndVarianceWithCommandBufferBatchNormalizationState() bool

	// optional
	Mean() *float64
	HasMean() bool

	// optional
	NumberOfFeatureChannels() uint
	HasNumberOfFeatureChannels() bool
}

// A concrete type wrapper for the [PCNNBatchNormalizationDataSource] protocol.
type CNNBatchNormalizationDataSourceWrapper struct {
	objc.Object
}

func (c_ CNNBatchNormalizationDataSourceWrapper) HasUpdateGammaAndBetaWithCommandBufferBatchNormalizationState() bool {
	return c_.RespondsToSelector(objc.Sel("updateGammaAndBetaWithCommandBuffer:batchNormalizationState:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationdatasource/2951891-updategammaandbetawithcommandbuf?language=objc
func (c_ CNNBatchNormalizationDataSourceWrapper) UpdateGammaAndBetaWithCommandBufferBatchNormalizationState(commandBuffer metal.PCommandBuffer, batchNormalizationState ICNNBatchNormalizationState) CNNNormalizationGammaAndBetaState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[CNNNormalizationGammaAndBetaState](c_, objc.Sel("updateGammaAndBetaWithCommandBuffer:batchNormalizationState:"), po0, objc.Ptr(batchNormalizationState))
	return rv
}

func (c_ CNNBatchNormalizationDataSourceWrapper) HasBeta() bool {
	return c_.RespondsToSelector(objc.Sel("beta"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationdatasource/2942586-beta?language=objc
func (c_ CNNBatchNormalizationDataSourceWrapper) Beta() *float64 {
	rv := objc.Call[*float64](c_, objc.Sel("beta"))
	return rv
}

func (c_ CNNBatchNormalizationDataSourceWrapper) HasEpsilon() bool {
	return c_.RespondsToSelector(objc.Sel("epsilon"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationdatasource/2947917-epsilon?language=objc
func (c_ CNNBatchNormalizationDataSourceWrapper) Epsilon() float64 {
	rv := objc.Call[float64](c_, objc.Sel("epsilon"))
	return rv
}

func (c_ CNNBatchNormalizationDataSourceWrapper) HasEncodeWithCoder() bool {
	return c_.RespondsToSelector(objc.Sel("encodeWithCoder:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationdatasource/2951882-encodewithcoder?language=objc
func (c_ CNNBatchNormalizationDataSourceWrapper) EncodeWithCoder(aCoder foundation.ICoder) {
	objc.Call[objc.Void](c_, objc.Sel("encodeWithCoder:"), objc.Ptr(aCoder))
}

func (c_ CNNBatchNormalizationDataSourceWrapper) HasUpdateGammaAndBetaWithBatchNormalizationState() bool {
	return c_.RespondsToSelector(objc.Sel("updateGammaAndBetaWithBatchNormalizationState:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationdatasource/2953129-updategammaandbetawithbatchnorma?language=objc
func (c_ CNNBatchNormalizationDataSourceWrapper) UpdateGammaAndBetaWithBatchNormalizationState(batchNormalizationState ICNNBatchNormalizationState) bool {
	rv := objc.Call[bool](c_, objc.Sel("updateGammaAndBetaWithBatchNormalizationState:"), objc.Ptr(batchNormalizationState))
	return rv
}

func (c_ CNNBatchNormalizationDataSourceWrapper) HasInitWithCoder() bool {
	return c_.RespondsToSelector(objc.Sel("initWithCoder:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationdatasource/2951886-initwithcoder?language=objc
func (c_ CNNBatchNormalizationDataSourceWrapper) InitWithCoder(aDecoder foundation.ICoder) objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("initWithCoder:"), objc.Ptr(aDecoder))
	return rv
}

func (c_ CNNBatchNormalizationDataSourceWrapper) HasGamma() bool {
	return c_.RespondsToSelector(objc.Sel("gamma"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationdatasource/2942605-gamma?language=objc
func (c_ CNNBatchNormalizationDataSourceWrapper) Gamma() *float64 {
	rv := objc.Call[*float64](c_, objc.Sel("gamma"))
	return rv
}

func (c_ CNNBatchNormalizationDataSourceWrapper) HasPurge() bool {
	return c_.RespondsToSelector(objc.Sel("purge"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationdatasource/2942607-purge?language=objc
func (c_ CNNBatchNormalizationDataSourceWrapper) Purge() {
	objc.Call[objc.Void](c_, objc.Sel("purge"))
}

func (c_ CNNBatchNormalizationDataSourceWrapper) HasUpdateMeanAndVarianceWithBatchNormalizationState() bool {
	return c_.RespondsToSelector(objc.Sel("updateMeanAndVarianceWithBatchNormalizationState:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationdatasource/3002360-updatemeanandvariancewithbatchno?language=objc
func (c_ CNNBatchNormalizationDataSourceWrapper) UpdateMeanAndVarianceWithBatchNormalizationState(batchNormalizationState ICNNBatchNormalizationState) bool {
	rv := objc.Call[bool](c_, objc.Sel("updateMeanAndVarianceWithBatchNormalizationState:"), objc.Ptr(batchNormalizationState))
	return rv
}

func (c_ CNNBatchNormalizationDataSourceWrapper) HasVariance() bool {
	return c_.RespondsToSelector(objc.Sel("variance"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationdatasource/2942592-variance?language=objc
func (c_ CNNBatchNormalizationDataSourceWrapper) Variance() *float64 {
	rv := objc.Call[*float64](c_, objc.Sel("variance"))
	return rv
}

func (c_ CNNBatchNormalizationDataSourceWrapper) HasCopyWithZoneDevice() bool {
	return c_.RespondsToSelector(objc.Sel("copyWithZone:device:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationdatasource/3013773-copywithzone?language=objc
func (c_ CNNBatchNormalizationDataSourceWrapper) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) objc.Object {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[objc.Object](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

func (c_ CNNBatchNormalizationDataSourceWrapper) HasLoad() bool {
	return c_.RespondsToSelector(objc.Sel("load"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationdatasource/2942579-load?language=objc
func (c_ CNNBatchNormalizationDataSourceWrapper) Load() bool {
	rv := objc.Call[bool](c_, objc.Sel("load"))
	return rv
}

func (c_ CNNBatchNormalizationDataSourceWrapper) HasLabel() bool {
	return c_.RespondsToSelector(objc.Sel("label"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationdatasource/2953128-label?language=objc
func (c_ CNNBatchNormalizationDataSourceWrapper) Label() string {
	rv := objc.Call[string](c_, objc.Sel("label"))
	return rv
}

func (c_ CNNBatchNormalizationDataSourceWrapper) HasUpdateMeanAndVarianceWithCommandBufferBatchNormalizationState() bool {
	return c_.RespondsToSelector(objc.Sel("updateMeanAndVarianceWithCommandBuffer:batchNormalizationState:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationdatasource/3002361-updatemeanandvariancewithcommand?language=objc
func (c_ CNNBatchNormalizationDataSourceWrapper) UpdateMeanAndVarianceWithCommandBufferBatchNormalizationState(commandBuffer metal.PCommandBuffer, batchNormalizationState ICNNBatchNormalizationState) CNNNormalizationMeanAndVarianceState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[CNNNormalizationMeanAndVarianceState](c_, objc.Sel("updateMeanAndVarianceWithCommandBuffer:batchNormalizationState:"), po0, objc.Ptr(batchNormalizationState))
	return rv
}

func (c_ CNNBatchNormalizationDataSourceWrapper) HasMean() bool {
	return c_.RespondsToSelector(objc.Sel("mean"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationdatasource/2942589-mean?language=objc
func (c_ CNNBatchNormalizationDataSourceWrapper) Mean() *float64 {
	rv := objc.Call[*float64](c_, objc.Sel("mean"))
	return rv
}

func (c_ CNNBatchNormalizationDataSourceWrapper) HasNumberOfFeatureChannels() bool {
	return c_.RespondsToSelector(objc.Sel("numberOfFeatureChannels"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationdatasource/2942596-numberoffeaturechannels?language=objc
func (c_ CNNBatchNormalizationDataSourceWrapper) NumberOfFeatureChannels() uint {
	rv := objc.Call[uint](c_, objc.Sel("numberOfFeatureChannels"))
	return rv
}
