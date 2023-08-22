// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNOptimizerAdam] class.
var NNOptimizerAdamClass = _NNOptimizerAdamClass{objc.GetClass("MPSNNOptimizerAdam")}

type _NNOptimizerAdamClass struct {
	objc.Class
}

// An interface definition for the [NNOptimizerAdam] class.
type INNOptimizerAdam interface {
	INNOptimizer
	EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixInputVelocityMatrixResultValuesMatrix(commandBuffer metal.PCommandBuffer, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputMomentumMatrix IMatrix, inputVelocityMatrix IMatrix, resultValuesMatrix IMatrix)
	EncodeToCommandBufferObjectInputGradientMatrixInputValuesMatrixInputMomentumMatrixInputVelocityMatrixResultValuesMatrix(commandBufferObject objc.IObject, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputMomentumMatrix IMatrix, inputVelocityMatrix IMatrix, resultValuesMatrix IMatrix)
	Epsilon() float64
	Beta2() float64
	Beta1() float64
	TimeStep() uint
	SetTimeStep(value uint)
}

// An optimization layer that performs an Adam pdate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam?language=objc
type NNOptimizerAdam struct {
	NNOptimizer
}

func NNOptimizerAdamFrom(ptr unsafe.Pointer) NNOptimizerAdam {
	return NNOptimizerAdam{
		NNOptimizer: NNOptimizerFrom(ptr),
	}
}

func (n_ NNOptimizerAdam) InitWithDeviceLearningRate(device metal.PDevice, learningRate float64) NNOptimizerAdam {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNOptimizerAdam](n_, objc.Sel("initWithDevice:learningRate:"), po0, learningRate)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/2966719-initwithdevice?language=objc
func NewNNOptimizerAdamWithDeviceLearningRate(device metal.PDevice, learningRate float64) NNOptimizerAdam {
	instance := NNOptimizerAdamClass.Alloc().InitWithDeviceLearningRate(device, learningRate)
	instance.Autorelease()
	return instance
}

func (nc _NNOptimizerAdamClass) Alloc() NNOptimizerAdam {
	rv := objc.Call[NNOptimizerAdam](nc, objc.Sel("alloc"))
	return rv
}

func NNOptimizerAdam_Alloc() NNOptimizerAdam {
	return NNOptimizerAdamClass.Alloc()
}

func (nc _NNOptimizerAdamClass) New() NNOptimizerAdam {
	rv := objc.Call[NNOptimizerAdam](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNOptimizerAdam() NNOptimizerAdam {
	return NNOptimizerAdamClass.New()
}

func (n_ NNOptimizerAdam) Init() NNOptimizerAdam {
	rv := objc.Call[NNOptimizerAdam](n_, objc.Sel("init"))
	return rv
}

func (n_ NNOptimizerAdam) InitWithDevice(device metal.PDevice) NNOptimizerAdam {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNOptimizerAdam](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewNNOptimizerAdamWithDevice(device metal.PDevice) NNOptimizerAdam {
	instance := NNOptimizerAdamClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (n_ NNOptimizerAdam) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNOptimizerAdam {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNOptimizerAdam](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNOptimizerAdam_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNOptimizerAdam {
	instance := NNOptimizerAdamClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/3197826-encodetocommandbuffer?language=objc
func (n_ NNOptimizerAdam) EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixInputVelocityMatrixResultValuesMatrix(commandBuffer metal.PCommandBuffer, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputMomentumMatrix IMatrix, inputVelocityMatrix IMatrix, resultValuesMatrix IMatrix) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:inputGradientMatrix:inputValuesMatrix:inputMomentumMatrix:inputVelocityMatrix:resultValuesMatrix:"), po0, objc.Ptr(inputGradientMatrix), objc.Ptr(inputValuesMatrix), objc.Ptr(inputMomentumMatrix), objc.Ptr(inputVelocityMatrix), objc.Ptr(resultValuesMatrix))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/3197826-encodetocommandbuffer?language=objc
func (n_ NNOptimizerAdam) EncodeToCommandBufferObjectInputGradientMatrixInputValuesMatrixInputMomentumMatrixInputVelocityMatrixResultValuesMatrix(commandBufferObject objc.IObject, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputMomentumMatrix IMatrix, inputVelocityMatrix IMatrix, resultValuesMatrix IMatrix) {
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:inputGradientMatrix:inputValuesMatrix:inputMomentumMatrix:inputVelocityMatrix:resultValuesMatrix:"), objc.Ptr(commandBufferObject), objc.Ptr(inputGradientMatrix), objc.Ptr(inputValuesMatrix), objc.Ptr(inputMomentumMatrix), objc.Ptr(inputVelocityMatrix), objc.Ptr(resultValuesMatrix))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/2966717-epsilon?language=objc
func (n_ NNOptimizerAdam) Epsilon() float64 {
	rv := objc.Call[float64](n_, objc.Sel("epsilon"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/2966715-beta2?language=objc
func (n_ NNOptimizerAdam) Beta2() float64 {
	rv := objc.Call[float64](n_, objc.Sel("beta2"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/2966714-beta1?language=objc
func (n_ NNOptimizerAdam) Beta1() float64 {
	rv := objc.Call[float64](n_, objc.Sel("beta1"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/2966720-timestep?language=objc
func (n_ NNOptimizerAdam) TimeStep() uint {
	rv := objc.Call[uint](n_, objc.Sel("timeStep"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizeradam/2966720-timestep?language=objc
func (n_ NNOptimizerAdam) SetTimeStep(value uint) {
	objc.Call[objc.Void](n_, objc.Sel("setTimeStep:"), value)
}
