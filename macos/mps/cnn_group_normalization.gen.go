// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNGroupNormalization] class.
var CNNGroupNormalizationClass = _CNNGroupNormalizationClass{objc.GetClass("MPSCNNGroupNormalization")}

type _CNNGroupNormalizationClass struct {
	objc.Class
}

// An interface definition for the [CNNGroupNormalization] class.
type ICNNGroupNormalization interface {
	ICNNKernel
	ReloadGammaAndBetaFromDataSource()
	ReloadGammaAndBetaWithCommandBufferGammaAndBetaState(commandBuffer metal.PCommandBuffer, gammaAndBetaState ICNNNormalizationGammaAndBetaState)
	ReloadGammaAndBetaWithCommandBufferObjectGammaAndBetaState(commandBufferObject objc.IObject, gammaAndBetaState ICNNNormalizationGammaAndBetaState)
	DataSource() CNNGroupNormalizationDataSourceWrapper
	Epsilon() float64
	SetEpsilon(value float64)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalization?language=objc
type CNNGroupNormalization struct {
	CNNKernel
}

func CNNGroupNormalizationFrom(ptr unsafe.Pointer) CNNGroupNormalization {
	return CNNGroupNormalization{
		CNNKernel: CNNKernelFrom(ptr),
	}
}

func (c_ CNNGroupNormalization) InitWithDeviceDataSource(device metal.PDevice, dataSource PCNNGroupNormalizationDataSource) CNNGroupNormalization {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	po1 := objc.WrapAsProtocol("MPSCNNGroupNormalizationDataSource", dataSource)
	rv := objc.Call[CNNGroupNormalization](c_, objc.Sel("initWithDevice:dataSource:"), po0, po1)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalization/3152537-initwithdevice?language=objc
func NewCNNGroupNormalizationWithDeviceDataSource(device metal.PDevice, dataSource PCNNGroupNormalizationDataSource) CNNGroupNormalization {
	instance := CNNGroupNormalizationClass.Alloc().InitWithDeviceDataSource(device, dataSource)
	instance.Autorelease()
	return instance
}

func (cc _CNNGroupNormalizationClass) Alloc() CNNGroupNormalization {
	rv := objc.Call[CNNGroupNormalization](cc, objc.Sel("alloc"))
	return rv
}

func CNNGroupNormalization_Alloc() CNNGroupNormalization {
	return CNNGroupNormalizationClass.Alloc()
}

func (cc _CNNGroupNormalizationClass) New() CNNGroupNormalization {
	rv := objc.Call[CNNGroupNormalization](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNGroupNormalization() CNNGroupNormalization {
	return CNNGroupNormalizationClass.New()
}

func (c_ CNNGroupNormalization) Init() CNNGroupNormalization {
	rv := objc.Call[CNNGroupNormalization](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNGroupNormalization) InitWithDevice(device metal.PDevice) CNNGroupNormalization {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNGroupNormalization](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNGroupNormalizationWithDevice(device metal.PDevice) CNNGroupNormalization {
	instance := CNNGroupNormalizationClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNGroupNormalization) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNGroupNormalization {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNGroupNormalization](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNGroupNormalization_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNGroupNormalization {
	instance := CNNGroupNormalizationClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalization/3152538-reloadgammaandbetafromdatasource?language=objc
func (c_ CNNGroupNormalization) ReloadGammaAndBetaFromDataSource() {
	objc.Call[objc.Void](c_, objc.Sel("reloadGammaAndBetaFromDataSource"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalization/3152539-reloadgammaandbetawithcommandbuf?language=objc
func (c_ CNNGroupNormalization) ReloadGammaAndBetaWithCommandBufferGammaAndBetaState(commandBuffer metal.PCommandBuffer, gammaAndBetaState ICNNNormalizationGammaAndBetaState) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](c_, objc.Sel("reloadGammaAndBetaWithCommandBuffer:gammaAndBetaState:"), po0, objc.Ptr(gammaAndBetaState))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalization/3152539-reloadgammaandbetawithcommandbuf?language=objc
func (c_ CNNGroupNormalization) ReloadGammaAndBetaWithCommandBufferObjectGammaAndBetaState(commandBufferObject objc.IObject, gammaAndBetaState ICNNNormalizationGammaAndBetaState) {
	objc.Call[objc.Void](c_, objc.Sel("reloadGammaAndBetaWithCommandBuffer:gammaAndBetaState:"), objc.Ptr(commandBufferObject), objc.Ptr(gammaAndBetaState))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalization/3152534-datasource?language=objc
func (c_ CNNGroupNormalization) DataSource() CNNGroupNormalizationDataSourceWrapper {
	rv := objc.Call[CNNGroupNormalizationDataSourceWrapper](c_, objc.Sel("dataSource"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalization/3152535-epsilon?language=objc
func (c_ CNNGroupNormalization) Epsilon() float64 {
	rv := objc.Call[float64](c_, objc.Sel("epsilon"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalization/3152535-epsilon?language=objc
func (c_ CNNGroupNormalization) SetEpsilon(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setEpsilon:"), value)
}
