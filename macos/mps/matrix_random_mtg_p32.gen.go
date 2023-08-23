// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MatrixRandomMTGP32] class.
var MatrixRandomMTGP32Class = _MatrixRandomMTGP32Class{objc.GetClass("MPSMatrixRandomMTGP32")}

type _MatrixRandomMTGP32Class struct {
	objc.Class
}

// An interface definition for the [MatrixRandomMTGP32] class.
type IMatrixRandomMTGP32 interface {
	IMatrixRandom
	SynchronizeStateOnCommandBuffer(commandBuffer metal.PCommandBuffer)
	SynchronizeStateOnCommandBufferObject(commandBufferObject objc.IObject)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandommtgp32?language=objc
type MatrixRandomMTGP32 struct {
	MatrixRandom
}

func MatrixRandomMTGP32From(ptr unsafe.Pointer) MatrixRandomMTGP32 {
	return MatrixRandomMTGP32{
		MatrixRandom: MatrixRandomFrom(ptr),
	}
}

func (m_ MatrixRandomMTGP32) InitWithDevice(device metal.PDevice) MatrixRandomMTGP32 {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixRandomMTGP32](m_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandommtgp32/3242865-initwithdevice?language=objc
func NewMatrixRandomMTGP32WithDevice(device metal.PDevice) MatrixRandomMTGP32 {
	instance := MatrixRandomMTGP32Class.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (mc _MatrixRandomMTGP32Class) Alloc() MatrixRandomMTGP32 {
	rv := objc.Call[MatrixRandomMTGP32](mc, objc.Sel("alloc"))
	return rv
}

func MatrixRandomMTGP32_Alloc() MatrixRandomMTGP32 {
	return MatrixRandomMTGP32Class.Alloc()
}

func (mc _MatrixRandomMTGP32Class) New() MatrixRandomMTGP32 {
	rv := objc.Call[MatrixRandomMTGP32](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMatrixRandomMTGP32() MatrixRandomMTGP32 {
	return MatrixRandomMTGP32Class.New()
}

func (m_ MatrixRandomMTGP32) Init() MatrixRandomMTGP32 {
	rv := objc.Call[MatrixRandomMTGP32](m_, objc.Sel("init"))
	return rv
}

func (m_ MatrixRandomMTGP32) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixRandomMTGP32 {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixRandomMTGP32](m_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func MatrixRandomMTGP32_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixRandomMTGP32 {
	instance := MatrixRandomMTGP32Class.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandommtgp32/3242868-synchronizestateoncommandbuffer?language=objc
func (m_ MatrixRandomMTGP32) SynchronizeStateOnCommandBuffer(commandBuffer metal.PCommandBuffer) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](m_, objc.Sel("synchronizeStateOnCommandBuffer:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandommtgp32/3242868-synchronizestateoncommandbuffer?language=objc
func (m_ MatrixRandomMTGP32) SynchronizeStateOnCommandBufferObject(commandBufferObject objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("synchronizeStateOnCommandBuffer:"), objc.Ptr(commandBufferObject))
}
