// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNLogSoftMax] class.
var CNNLogSoftMaxClass = _CNNLogSoftMaxClass{objc.GetClass("MPSCNNLogSoftMax")}

type _CNNLogSoftMaxClass struct {
	objc.Class
}

// An interface definition for the [CNNLogSoftMax] class.
type ICNNLogSoftMax interface {
	ICNNKernel
}

// A neural transfer function that  is useful for constructing a loss function to be minimized when training neural networks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlogsoftmax?language=objc
type CNNLogSoftMax struct {
	CNNKernel
}

func CNNLogSoftMaxFrom(ptr unsafe.Pointer) CNNLogSoftMax {
	return CNNLogSoftMax{
		CNNKernel: CNNKernelFrom(ptr),
	}
}

func (cc _CNNLogSoftMaxClass) Alloc() CNNLogSoftMax {
	rv := objc.Call[CNNLogSoftMax](cc, objc.Sel("alloc"))
	return rv
}

func CNNLogSoftMax_Alloc() CNNLogSoftMax {
	return CNNLogSoftMaxClass.Alloc()
}

func (cc _CNNLogSoftMaxClass) New() CNNLogSoftMax {
	rv := objc.Call[CNNLogSoftMax](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNLogSoftMax() CNNLogSoftMax {
	return CNNLogSoftMaxClass.New()
}

func (c_ CNNLogSoftMax) Init() CNNLogSoftMax {
	rv := objc.Call[CNNLogSoftMax](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNLogSoftMax) InitWithDevice(device metal.PDevice) CNNLogSoftMax {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNLogSoftMax](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNLogSoftMaxWithDevice(device metal.PDevice) CNNLogSoftMax {
	instance := CNNLogSoftMaxClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNLogSoftMax) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNLogSoftMax {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNLogSoftMax](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNLogSoftMax_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNLogSoftMax {
	instance := CNNLogSoftMaxClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
