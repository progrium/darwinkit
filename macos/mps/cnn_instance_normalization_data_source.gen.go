// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// A protocol that defines methods that an instance normalization uses to initialize scale factors and bias terms. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationdatasource?language=objc
type PCNNInstanceNormalizationDataSource interface {
	// optional
	UpdateGammaAndBetaWithCommandBufferInstanceNormalizationStateBatch(commandBuffer metal.CommandBufferWrapper, instanceNormalizationStateBatch *foundation.Array) ICNNNormalizationGammaAndBetaState
	HasUpdateGammaAndBetaWithCommandBufferInstanceNormalizationStateBatch() bool

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
	UpdateGammaAndBetaWithInstanceNormalizationStateBatch(instanceNormalizationStateBatch *foundation.Array) bool
	HasUpdateGammaAndBetaWithInstanceNormalizationStateBatch() bool

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
	CopyWithZoneDevice(zone unsafe.Pointer, device metal.DeviceWrapper) objc.IObject
	HasCopyWithZoneDevice() bool

	// optional
	Load() bool
	HasLoad() bool

	// optional
	Label() string
	HasLabel() bool

	// optional
	NumberOfFeatureChannels() uint
	HasNumberOfFeatureChannels() bool
}

// A concrete type wrapper for the [PCNNInstanceNormalizationDataSource] protocol.
type CNNInstanceNormalizationDataSourceWrapper struct {
	objc.Object
}

func (c_ CNNInstanceNormalizationDataSourceWrapper) HasUpdateGammaAndBetaWithCommandBufferInstanceNormalizationStateBatch() bool {
	return c_.RespondsToSelector(objc.Sel("updateGammaAndBetaWithCommandBuffer:instanceNormalizationStateBatch:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationdatasource/2953926-updategammaandbetawithcommandbuf?language=objc
func (c_ CNNInstanceNormalizationDataSourceWrapper) UpdateGammaAndBetaWithCommandBufferInstanceNormalizationStateBatch(commandBuffer metal.PCommandBuffer, instanceNormalizationStateBatch *foundation.Array) CNNNormalizationGammaAndBetaState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[CNNNormalizationGammaAndBetaState](c_, objc.Sel("updateGammaAndBetaWithCommandBuffer:instanceNormalizationStateBatch:"), po0, instanceNormalizationStateBatch)
	return rv
}

func (c_ CNNInstanceNormalizationDataSourceWrapper) HasBeta() bool {
	return c_.RespondsToSelector(objc.Sel("beta"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationdatasource/2953922-beta?language=objc
func (c_ CNNInstanceNormalizationDataSourceWrapper) Beta() *float64 {
	rv := objc.Call[*float64](c_, objc.Sel("beta"))
	return rv
}

func (c_ CNNInstanceNormalizationDataSourceWrapper) HasEpsilon() bool {
	return c_.RespondsToSelector(objc.Sel("epsilon"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationdatasource/2953925-epsilon?language=objc
func (c_ CNNInstanceNormalizationDataSourceWrapper) Epsilon() float64 {
	rv := objc.Call[float64](c_, objc.Sel("epsilon"))
	return rv
}

func (c_ CNNInstanceNormalizationDataSourceWrapper) HasEncodeWithCoder() bool {
	return c_.RespondsToSelector(objc.Sel("encodeWithCoder:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationdatasource/2947953-encodewithcoder?language=objc
func (c_ CNNInstanceNormalizationDataSourceWrapper) EncodeWithCoder(aCoder foundation.ICoder) {
	objc.Call[objc.Void](c_, objc.Sel("encodeWithCoder:"), objc.Ptr(aCoder))
}

func (c_ CNNInstanceNormalizationDataSourceWrapper) HasUpdateGammaAndBetaWithInstanceNormalizationStateBatch() bool {
	return c_.RespondsToSelector(objc.Sel("updateGammaAndBetaWithInstanceNormalizationStateBatch:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationdatasource/2953931-updategammaandbetawithinstanceno?language=objc
func (c_ CNNInstanceNormalizationDataSourceWrapper) UpdateGammaAndBetaWithInstanceNormalizationStateBatch(instanceNormalizationStateBatch *foundation.Array) bool {
	rv := objc.Call[bool](c_, objc.Sel("updateGammaAndBetaWithInstanceNormalizationStateBatch:"), instanceNormalizationStateBatch)
	return rv
}

func (c_ CNNInstanceNormalizationDataSourceWrapper) HasInitWithCoder() bool {
	return c_.RespondsToSelector(objc.Sel("initWithCoder:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationdatasource/2947957-initwithcoder?language=objc
func (c_ CNNInstanceNormalizationDataSourceWrapper) InitWithCoder(aDecoder foundation.ICoder) objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("initWithCoder:"), objc.Ptr(aDecoder))
	return rv
}

func (c_ CNNInstanceNormalizationDataSourceWrapper) HasGamma() bool {
	return c_.RespondsToSelector(objc.Sel("gamma"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationdatasource/2953923-gamma?language=objc
func (c_ CNNInstanceNormalizationDataSourceWrapper) Gamma() *float64 {
	rv := objc.Call[*float64](c_, objc.Sel("gamma"))
	return rv
}

func (c_ CNNInstanceNormalizationDataSourceWrapper) HasPurge() bool {
	return c_.RespondsToSelector(objc.Sel("purge"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationdatasource/3088879-purge?language=objc
func (c_ CNNInstanceNormalizationDataSourceWrapper) Purge() {
	objc.Call[objc.Void](c_, objc.Sel("purge"))
}

func (c_ CNNInstanceNormalizationDataSourceWrapper) HasCopyWithZoneDevice() bool {
	return c_.RespondsToSelector(objc.Sel("copyWithZone:device:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationdatasource/3013780-copywithzone?language=objc
func (c_ CNNInstanceNormalizationDataSourceWrapper) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) objc.Object {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[objc.Object](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

func (c_ CNNInstanceNormalizationDataSourceWrapper) HasLoad() bool {
	return c_.RespondsToSelector(objc.Sel("load"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationdatasource/3088878-load?language=objc
func (c_ CNNInstanceNormalizationDataSourceWrapper) Load() bool {
	rv := objc.Call[bool](c_, objc.Sel("load"))
	return rv
}

func (c_ CNNInstanceNormalizationDataSourceWrapper) HasLabel() bool {
	return c_.RespondsToSelector(objc.Sel("label"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationdatasource/2952998-label?language=objc
func (c_ CNNInstanceNormalizationDataSourceWrapper) Label() string {
	rv := objc.Call[string](c_, objc.Sel("label"))
	return rv
}

func (c_ CNNInstanceNormalizationDataSourceWrapper) HasNumberOfFeatureChannels() bool {
	return c_.RespondsToSelector(objc.Sel("numberOfFeatureChannels"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationdatasource/2947961-numberoffeaturechannels?language=objc
func (c_ CNNInstanceNormalizationDataSourceWrapper) NumberOfFeatureChannels() uint {
	rv := objc.Call[uint](c_, objc.Sel("numberOfFeatureChannels"))
	return rv
}
