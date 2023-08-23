// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNInstanceNormalization] class.
var CNNInstanceNormalizationClass = _CNNInstanceNormalizationClass{objc.GetClass("MPSCNNInstanceNormalization")}

type _CNNInstanceNormalizationClass struct {
	objc.Class
}

// An interface definition for the [CNNInstanceNormalization] class.
type ICNNInstanceNormalization interface {
	ICNNKernel
	ReloadGammaAndBetaFromDataSource()
	ReloadGammaAndBetaWithCommandBufferGammaAndBetaState(commandBuffer metal.PCommandBuffer, gammaAndBetaState ICNNNormalizationGammaAndBetaState)
	ReloadGammaAndBetaWithCommandBufferObjectGammaAndBetaState(commandBufferObject objc.IObject, gammaAndBetaState ICNNNormalizationGammaAndBetaState)
	DataSource() CNNInstanceNormalizationDataSourceWrapper
	Epsilon() float64
	SetEpsilon(value float64)
}

// An instance normalization kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalization?language=objc
type CNNInstanceNormalization struct {
	CNNKernel
}

func CNNInstanceNormalizationFrom(ptr unsafe.Pointer) CNNInstanceNormalization {
	return CNNInstanceNormalization{
		CNNKernel: CNNKernelFrom(ptr),
	}
}

func (c_ CNNInstanceNormalization) InitWithDeviceDataSource(device metal.PDevice, dataSource PCNNInstanceNormalizationDataSource) CNNInstanceNormalization {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	po1 := objc.WrapAsProtocol("MPSCNNInstanceNormalizationDataSource", dataSource)
	rv := objc.Call[CNNInstanceNormalization](c_, objc.Sel("initWithDevice:dataSource:"), po0, po1)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalization/2947947-initwithdevice?language=objc
func NewCNNInstanceNormalizationWithDeviceDataSource(device metal.PDevice, dataSource PCNNInstanceNormalizationDataSource) CNNInstanceNormalization {
	instance := CNNInstanceNormalizationClass.Alloc().InitWithDeviceDataSource(device, dataSource)
	instance.Autorelease()
	return instance
}

func (cc _CNNInstanceNormalizationClass) Alloc() CNNInstanceNormalization {
	rv := objc.Call[CNNInstanceNormalization](cc, objc.Sel("alloc"))
	return rv
}

func CNNInstanceNormalization_Alloc() CNNInstanceNormalization {
	return CNNInstanceNormalizationClass.Alloc()
}

func (cc _CNNInstanceNormalizationClass) New() CNNInstanceNormalization {
	rv := objc.Call[CNNInstanceNormalization](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNInstanceNormalization() CNNInstanceNormalization {
	return CNNInstanceNormalizationClass.New()
}

func (c_ CNNInstanceNormalization) Init() CNNInstanceNormalization {
	rv := objc.Call[CNNInstanceNormalization](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNInstanceNormalization) InitWithDevice(device metal.PDevice) CNNInstanceNormalization {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNInstanceNormalization](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNInstanceNormalizationWithDevice(device metal.PDevice) CNNInstanceNormalization {
	instance := CNNInstanceNormalizationClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNInstanceNormalization) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNInstanceNormalization {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNInstanceNormalization](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNInstanceNormalization_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNInstanceNormalization {
	instance := CNNInstanceNormalizationClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalization/2976471-reloadgammaandbetafromdatasource?language=objc
func (c_ CNNInstanceNormalization) ReloadGammaAndBetaFromDataSource() {
	objc.Call[objc.Void](c_, objc.Sel("reloadGammaAndBetaFromDataSource"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalization/2953921-reloadgammaandbetawithcommandbuf?language=objc
func (c_ CNNInstanceNormalization) ReloadGammaAndBetaWithCommandBufferGammaAndBetaState(commandBuffer metal.PCommandBuffer, gammaAndBetaState ICNNNormalizationGammaAndBetaState) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](c_, objc.Sel("reloadGammaAndBetaWithCommandBuffer:gammaAndBetaState:"), po0, objc.Ptr(gammaAndBetaState))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalization/2953921-reloadgammaandbetawithcommandbuf?language=objc
func (c_ CNNInstanceNormalization) ReloadGammaAndBetaWithCommandBufferObjectGammaAndBetaState(commandBufferObject objc.IObject, gammaAndBetaState ICNNNormalizationGammaAndBetaState) {
	objc.Call[objc.Void](c_, objc.Sel("reloadGammaAndBetaWithCommandBuffer:gammaAndBetaState:"), objc.Ptr(commandBufferObject), objc.Ptr(gammaAndBetaState))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalization/2953927-datasource?language=objc
func (c_ CNNInstanceNormalization) DataSource() CNNInstanceNormalizationDataSourceWrapper {
	rv := objc.Call[CNNInstanceNormalizationDataSourceWrapper](c_, objc.Sel("dataSource"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalization/2947943-epsilon?language=objc
func (c_ CNNInstanceNormalization) Epsilon() float64 {
	rv := objc.Call[float64](c_, objc.Sel("epsilon"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalization/2947943-epsilon?language=objc
func (c_ CNNInstanceNormalization) SetEpsilon(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setEpsilon:"), value)
}
