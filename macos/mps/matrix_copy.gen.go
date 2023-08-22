// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MatrixCopy] class.
var MatrixCopyClass = _MatrixCopyClass{objc.GetClass("MPSMatrixCopy")}

type _MatrixCopyClass struct {
	objc.Class
}

// An interface definition for the [MatrixCopy] class.
type IMatrixCopy interface {
	IKernel
	EncodeToCommandBufferCopyDescriptor(commandBuffer metal.PCommandBuffer, copyDescriptor IMatrixCopyDescriptor)
	EncodeToCommandBufferObjectCopyDescriptor(commandBufferObject objc.IObject, copyDescriptor IMatrixCopyDescriptor)
	SourcesAreTransposed() bool
	DestinationsAreTransposed() bool
	CopyRows() uint
	CopyColumns() uint
}

// A class that can perform multiple matrix copy operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopy?language=objc
type MatrixCopy struct {
	Kernel
}

func MatrixCopyFrom(ptr unsafe.Pointer) MatrixCopy {
	return MatrixCopy{
		Kernel: KernelFrom(ptr),
	}
}

func (m_ MatrixCopy) InitWithDeviceCopyRowsCopyColumnsSourcesAreTransposedDestinationsAreTransposed(device metal.PDevice, copyRows uint, copyColumns uint, sourcesAreTransposed bool, destinationsAreTransposed bool) MatrixCopy {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixCopy](m_, objc.Sel("initWithDevice:copyRows:copyColumns:sourcesAreTransposed:destinationsAreTransposed:"), po0, copyRows, copyColumns, sourcesAreTransposed, destinationsAreTransposed)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopy/2915345-initwithdevice?language=objc
func NewMatrixCopyWithDeviceCopyRowsCopyColumnsSourcesAreTransposedDestinationsAreTransposed(device metal.PDevice, copyRows uint, copyColumns uint, sourcesAreTransposed bool, destinationsAreTransposed bool) MatrixCopy {
	instance := MatrixCopyClass.Alloc().InitWithDeviceCopyRowsCopyColumnsSourcesAreTransposedDestinationsAreTransposed(device, copyRows, copyColumns, sourcesAreTransposed, destinationsAreTransposed)
	instance.Autorelease()
	return instance
}

func (mc _MatrixCopyClass) Alloc() MatrixCopy {
	rv := objc.Call[MatrixCopy](mc, objc.Sel("alloc"))
	return rv
}

func MatrixCopy_Alloc() MatrixCopy {
	return MatrixCopyClass.Alloc()
}

func (mc _MatrixCopyClass) New() MatrixCopy {
	rv := objc.Call[MatrixCopy](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMatrixCopy() MatrixCopy {
	return MatrixCopyClass.New()
}

func (m_ MatrixCopy) Init() MatrixCopy {
	rv := objc.Call[MatrixCopy](m_, objc.Sel("init"))
	return rv
}

func (m_ MatrixCopy) InitWithDevice(device metal.PDevice) MatrixCopy {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixCopy](m_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewMatrixCopyWithDevice(device metal.PDevice) MatrixCopy {
	instance := MatrixCopyClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (m_ MatrixCopy) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixCopy {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixCopy](m_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func MatrixCopy_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixCopy {
	instance := MatrixCopyClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopy/2915341-encodetocommandbuffer?language=objc
func (m_ MatrixCopy) EncodeToCommandBufferCopyDescriptor(commandBuffer metal.PCommandBuffer, copyDescriptor IMatrixCopyDescriptor) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:copyDescriptor:"), po0, objc.Ptr(copyDescriptor))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopy/2915341-encodetocommandbuffer?language=objc
func (m_ MatrixCopy) EncodeToCommandBufferObjectCopyDescriptor(commandBufferObject objc.IObject, copyDescriptor IMatrixCopyDescriptor) {
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:copyDescriptor:"), objc.Ptr(commandBufferObject), objc.Ptr(copyDescriptor))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopy/2915340-sourcesaretransposed?language=objc
func (m_ MatrixCopy) SourcesAreTransposed() bool {
	rv := objc.Call[bool](m_, objc.Sel("sourcesAreTransposed"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopy/2915326-destinationsaretransposed?language=objc
func (m_ MatrixCopy) DestinationsAreTransposed() bool {
	rv := objc.Call[bool](m_, objc.Sel("destinationsAreTransposed"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopy/2915342-copyrows?language=objc
func (m_ MatrixCopy) CopyRows() uint {
	rv := objc.Call[uint](m_, objc.Sel("copyRows"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopy/2915325-copycolumns?language=objc
func (m_ MatrixCopy) CopyColumns() uint {
	rv := objc.Call[uint](m_, objc.Sel("copyColumns"))
	return rv
}
