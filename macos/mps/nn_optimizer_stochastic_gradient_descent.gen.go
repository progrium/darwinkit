// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNOptimizerStochasticGradientDescent] class.
var NNOptimizerStochasticGradientDescentClass = _NNOptimizerStochasticGradientDescentClass{objc.GetClass("MPSNNOptimizerStochasticGradientDescent")}

type _NNOptimizerStochasticGradientDescentClass struct {
	objc.Class
}

// An interface definition for the [NNOptimizerStochasticGradientDescent] class.
type INNOptimizerStochasticGradientDescent interface {
	INNOptimizer
	EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixResultValuesMatrix(commandBuffer metal.PCommandBuffer, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputMomentumMatrix IMatrix, resultValuesMatrix IMatrix)
	EncodeToCommandBufferObjectInputGradientMatrixInputValuesMatrixInputMomentumMatrixResultValuesMatrix(commandBufferObject objc.IObject, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputMomentumMatrix IMatrix, resultValuesMatrix IMatrix)
	UseNesterovMomentum() bool
	MomentumScale() float64
	UseNestrovMomentum() bool
}

// An optimization layer that performs a gradient descent with an optional momentum update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent?language=objc
type NNOptimizerStochasticGradientDescent struct {
	NNOptimizer
}

func NNOptimizerStochasticGradientDescentFrom(ptr unsafe.Pointer) NNOptimizerStochasticGradientDescent {
	return NNOptimizerStochasticGradientDescent{
		NNOptimizer: NNOptimizerFrom(ptr),
	}
}

func (n_ NNOptimizerStochasticGradientDescent) InitWithDeviceLearningRate(device metal.PDevice, learningRate float64) NNOptimizerStochasticGradientDescent {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNOptimizerStochasticGradientDescent](n_, objc.Sel("initWithDevice:learningRate:"), po0, learningRate)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/2966741-initwithdevice?language=objc
func NewNNOptimizerStochasticGradientDescentWithDeviceLearningRate(device metal.PDevice, learningRate float64) NNOptimizerStochasticGradientDescent {
	instance := NNOptimizerStochasticGradientDescentClass.Alloc().InitWithDeviceLearningRate(device, learningRate)
	instance.Autorelease()
	return instance
}

func (nc _NNOptimizerStochasticGradientDescentClass) Alloc() NNOptimizerStochasticGradientDescent {
	rv := objc.Call[NNOptimizerStochasticGradientDescent](nc, objc.Sel("alloc"))
	return rv
}

func NNOptimizerStochasticGradientDescent_Alloc() NNOptimizerStochasticGradientDescent {
	return NNOptimizerStochasticGradientDescentClass.Alloc()
}

func (nc _NNOptimizerStochasticGradientDescentClass) New() NNOptimizerStochasticGradientDescent {
	rv := objc.Call[NNOptimizerStochasticGradientDescent](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNOptimizerStochasticGradientDescent() NNOptimizerStochasticGradientDescent {
	return NNOptimizerStochasticGradientDescentClass.New()
}

func (n_ NNOptimizerStochasticGradientDescent) Init() NNOptimizerStochasticGradientDescent {
	rv := objc.Call[NNOptimizerStochasticGradientDescent](n_, objc.Sel("init"))
	return rv
}

func (n_ NNOptimizerStochasticGradientDescent) InitWithDevice(device metal.PDevice) NNOptimizerStochasticGradientDescent {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNOptimizerStochasticGradientDescent](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewNNOptimizerStochasticGradientDescentWithDevice(device metal.PDevice) NNOptimizerStochasticGradientDescent {
	instance := NNOptimizerStochasticGradientDescentClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (n_ NNOptimizerStochasticGradientDescent) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNOptimizerStochasticGradientDescent {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNOptimizerStochasticGradientDescent](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNOptimizerStochasticGradientDescent_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNOptimizerStochasticGradientDescent {
	instance := NNOptimizerStochasticGradientDescentClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/3197828-encodetocommandbuffer?language=objc
func (n_ NNOptimizerStochasticGradientDescent) EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixResultValuesMatrix(commandBuffer metal.PCommandBuffer, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputMomentumMatrix IMatrix, resultValuesMatrix IMatrix) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:inputGradientMatrix:inputValuesMatrix:inputMomentumMatrix:resultValuesMatrix:"), po0, objc.Ptr(inputGradientMatrix), objc.Ptr(inputValuesMatrix), objc.Ptr(inputMomentumMatrix), objc.Ptr(resultValuesMatrix))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/3197828-encodetocommandbuffer?language=objc
func (n_ NNOptimizerStochasticGradientDescent) EncodeToCommandBufferObjectInputGradientMatrixInputValuesMatrixInputMomentumMatrixResultValuesMatrix(commandBufferObject objc.IObject, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputMomentumMatrix IMatrix, resultValuesMatrix IMatrix) {
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:inputGradientMatrix:inputValuesMatrix:inputMomentumMatrix:resultValuesMatrix:"), objc.Ptr(commandBufferObject), objc.Ptr(inputGradientMatrix), objc.Ptr(inputValuesMatrix), objc.Ptr(inputMomentumMatrix), objc.Ptr(resultValuesMatrix))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/3675592-usenesterovmomentum?language=objc
func (n_ NNOptimizerStochasticGradientDescent) UseNesterovMomentum() bool {
	rv := objc.Call[bool](n_, objc.Sel("useNesterovMomentum"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/2966743-momentumscale?language=objc
func (n_ NNOptimizerStochasticGradientDescent) MomentumScale() float64 {
	rv := objc.Call[float64](n_, objc.Sel("momentumScale"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/2966744-usenestrovmomentum?language=objc
func (n_ NNOptimizerStochasticGradientDescent) UseNestrovMomentum() bool {
	rv := objc.Call[bool](n_, objc.Sel("useNestrovMomentum"))
	return rv
}
