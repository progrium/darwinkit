// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TemporaryNDArray] class.
var TemporaryNDArrayClass = _TemporaryNDArrayClass{objc.GetClass("MPSTemporaryNDArray")}

type _TemporaryNDArrayClass struct {
	objc.Class
}

// An interface definition for the [TemporaryNDArray] class.
type ITemporaryNDArray interface {
	INDArray
	ReadCount() uint
	SetReadCount(value uint)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryndarray?language=objc
type TemporaryNDArray struct {
	NDArray
}

func TemporaryNDArrayFrom(ptr unsafe.Pointer) TemporaryNDArray {
	return TemporaryNDArray{
		NDArray: NDArrayFrom(ptr),
	}
}

func (tc _TemporaryNDArrayClass) TemporaryNDArrayWithCommandBufferDescriptor(commandBuffer metal.PCommandBuffer, descriptor INDArrayDescriptor) TemporaryNDArray {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[TemporaryNDArray](tc, objc.Sel("temporaryNDArrayWithCommandBuffer:descriptor:"), po0, objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryndarray/3114075-temporaryndarraywithcommandbuffe?language=objc
func TemporaryNDArray_TemporaryNDArrayWithCommandBufferDescriptor(commandBuffer metal.PCommandBuffer, descriptor INDArrayDescriptor) TemporaryNDArray {
	return TemporaryNDArrayClass.TemporaryNDArrayWithCommandBufferDescriptor(commandBuffer, descriptor)
}

func (tc _TemporaryNDArrayClass) Alloc() TemporaryNDArray {
	rv := objc.Call[TemporaryNDArray](tc, objc.Sel("alloc"))
	return rv
}

func TemporaryNDArray_Alloc() TemporaryNDArray {
	return TemporaryNDArrayClass.Alloc()
}

func (tc _TemporaryNDArrayClass) New() TemporaryNDArray {
	rv := objc.Call[TemporaryNDArray](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTemporaryNDArray() TemporaryNDArray {
	return TemporaryNDArrayClass.New()
}

func (t_ TemporaryNDArray) Init() TemporaryNDArray {
	rv := objc.Call[TemporaryNDArray](t_, objc.Sel("init"))
	return rv
}

func (t_ TemporaryNDArray) InitWithDeviceScalar(device metal.PDevice, value float64) TemporaryNDArray {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[TemporaryNDArray](t_, objc.Sel("initWithDevice:scalar:"), po0, value)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114051-initwithdevice?language=objc
func NewTemporaryNDArrayWithDeviceScalar(device metal.PDevice, value float64) TemporaryNDArray {
	instance := TemporaryNDArrayClass.Alloc().InitWithDeviceScalar(device, value)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryndarray/3114074-readcount?language=objc
func (t_ TemporaryNDArray) ReadCount() uint {
	rv := objc.Call[uint](t_, objc.Sel("readCount"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryndarray/3114074-readcount?language=objc
func (t_ TemporaryNDArray) SetReadCount(value uint) {
	objc.Call[objc.Void](t_, objc.Sel("setReadCount:"), value)
}
