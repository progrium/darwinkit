// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationdatasource?language=objc
type PCNNGroupNormalizationDataSource interface {
	// optional
	UpdateGammaAndBetaWithCommandBufferGroupNormalizationStateBatch(commandBuffer metal.CommandBufferWrapper, groupNormalizationStateBatch *foundation.Array) ICNNNormalizationGammaAndBetaState
	HasUpdateGammaAndBetaWithCommandBufferGroupNormalizationStateBatch() bool

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
	UpdateGammaAndBetaWithGroupNormalizationStateBatch(groupNormalizationStateBatch *foundation.Array) bool
	HasUpdateGammaAndBetaWithGroupNormalizationStateBatch() bool

	// optional
	InitWithCoder(aDecoder foundation.Coder) objc.IObject
	HasInitWithCoder() bool

	// optional
	Gamma() *float64
	HasGamma() bool

	// optional
	CopyWithZoneDevice(zone unsafe.Pointer, device metal.DeviceWrapper) objc.IObject
	HasCopyWithZoneDevice() bool

	// optional
	Label() string
	HasLabel() bool

	// optional
	SetNumberOfGroups(value uint)
	HasSetNumberOfGroups() bool

	// optional
	NumberOfGroups() uint
	HasNumberOfGroups() bool

	// optional
	NumberOfFeatureChannels() uint
	HasNumberOfFeatureChannels() bool
}

// A concrete type wrapper for the [PCNNGroupNormalizationDataSource] protocol.
type CNNGroupNormalizationDataSourceWrapper struct {
	objc.Object
}

func (c_ CNNGroupNormalizationDataSourceWrapper) HasUpdateGammaAndBetaWithCommandBufferGroupNormalizationStateBatch() bool {
	return c_.RespondsToSelector(objc.Sel("updateGammaAndBetaWithCommandBuffer:groupNormalizationStateBatch:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationdatasource/3152553-updategammaandbetawithcommandbuf?language=objc
func (c_ CNNGroupNormalizationDataSourceWrapper) UpdateGammaAndBetaWithCommandBufferGroupNormalizationStateBatch(commandBuffer metal.PCommandBuffer, groupNormalizationStateBatch *foundation.Array) CNNNormalizationGammaAndBetaState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[CNNNormalizationGammaAndBetaState](c_, objc.Sel("updateGammaAndBetaWithCommandBuffer:groupNormalizationStateBatch:"), po0, groupNormalizationStateBatch)
	return rv
}

func (c_ CNNGroupNormalizationDataSourceWrapper) HasBeta() bool {
	return c_.RespondsToSelector(objc.Sel("beta"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationdatasource/3152543-beta?language=objc
func (c_ CNNGroupNormalizationDataSourceWrapper) Beta() *float64 {
	rv := objc.Call[*float64](c_, objc.Sel("beta"))
	return rv
}

func (c_ CNNGroupNormalizationDataSourceWrapper) HasEpsilon() bool {
	return c_.RespondsToSelector(objc.Sel("epsilon"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationdatasource/3152546-epsilon?language=objc
func (c_ CNNGroupNormalizationDataSourceWrapper) Epsilon() float64 {
	rv := objc.Call[float64](c_, objc.Sel("epsilon"))
	return rv
}

func (c_ CNNGroupNormalizationDataSourceWrapper) HasEncodeWithCoder() bool {
	return c_.RespondsToSelector(objc.Sel("encodeWithCoder:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationdatasource/3152545-encodewithcoder?language=objc
func (c_ CNNGroupNormalizationDataSourceWrapper) EncodeWithCoder(aCoder foundation.ICoder) {
	objc.Call[objc.Void](c_, objc.Sel("encodeWithCoder:"), objc.Ptr(aCoder))
}

func (c_ CNNGroupNormalizationDataSourceWrapper) HasUpdateGammaAndBetaWithGroupNormalizationStateBatch() bool {
	return c_.RespondsToSelector(objc.Sel("updateGammaAndBetaWithGroupNormalizationStateBatch:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationdatasource/3152554-updategammaandbetawithgroupnorma?language=objc
func (c_ CNNGroupNormalizationDataSourceWrapper) UpdateGammaAndBetaWithGroupNormalizationStateBatch(groupNormalizationStateBatch *foundation.Array) bool {
	rv := objc.Call[bool](c_, objc.Sel("updateGammaAndBetaWithGroupNormalizationStateBatch:"), groupNormalizationStateBatch)
	return rv
}

func (c_ CNNGroupNormalizationDataSourceWrapper) HasInitWithCoder() bool {
	return c_.RespondsToSelector(objc.Sel("initWithCoder:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationdatasource/3152548-initwithcoder?language=objc
func (c_ CNNGroupNormalizationDataSourceWrapper) InitWithCoder(aDecoder foundation.ICoder) objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("initWithCoder:"), objc.Ptr(aDecoder))
	return rv
}

func (c_ CNNGroupNormalizationDataSourceWrapper) HasGamma() bool {
	return c_.RespondsToSelector(objc.Sel("gamma"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationdatasource/3152547-gamma?language=objc
func (c_ CNNGroupNormalizationDataSourceWrapper) Gamma() *float64 {
	rv := objc.Call[*float64](c_, objc.Sel("gamma"))
	return rv
}

func (c_ CNNGroupNormalizationDataSourceWrapper) HasCopyWithZoneDevice() bool {
	return c_.RespondsToSelector(objc.Sel("copyWithZone:device:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationdatasource/3152544-copywithzone?language=objc
func (c_ CNNGroupNormalizationDataSourceWrapper) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) objc.Object {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[objc.Object](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

func (c_ CNNGroupNormalizationDataSourceWrapper) HasLabel() bool {
	return c_.RespondsToSelector(objc.Sel("label"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationdatasource/3152549-label?language=objc
func (c_ CNNGroupNormalizationDataSourceWrapper) Label() string {
	rv := objc.Call[string](c_, objc.Sel("label"))
	return rv
}

func (c_ CNNGroupNormalizationDataSourceWrapper) HasSetNumberOfGroups() bool {
	return c_.RespondsToSelector(objc.Sel("setNumberOfGroups:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationdatasource/3152551-numberofgroups?language=objc
func (c_ CNNGroupNormalizationDataSourceWrapper) SetNumberOfGroups(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setNumberOfGroups:"), value)
}

func (c_ CNNGroupNormalizationDataSourceWrapper) HasNumberOfGroups() bool {
	return c_.RespondsToSelector(objc.Sel("numberOfGroups"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationdatasource/3152551-numberofgroups?language=objc
func (c_ CNNGroupNormalizationDataSourceWrapper) NumberOfGroups() uint {
	rv := objc.Call[uint](c_, objc.Sel("numberOfGroups"))
	return rv
}

func (c_ CNNGroupNormalizationDataSourceWrapper) HasNumberOfFeatureChannels() bool {
	return c_.RespondsToSelector(objc.Sel("numberOfFeatureChannels"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationdatasource/3152550-numberoffeaturechannels?language=objc
func (c_ CNNGroupNormalizationDataSourceWrapper) NumberOfFeatureChannels() uint {
	rv := objc.Call[uint](c_, objc.Sel("numberOfFeatureChannels"))
	return rv
}
