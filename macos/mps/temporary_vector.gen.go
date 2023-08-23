// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TemporaryVector] class.
var TemporaryVectorClass = _TemporaryVectorClass{objc.GetClass("MPSTemporaryVector")}

type _TemporaryVectorClass struct {
	objc.Class
}

// An interface definition for the [TemporaryVector] class.
type ITemporaryVector interface {
	IVector
	ReadCount() uint
	SetReadCount(value uint)
}

// A vector allocated on GPU private memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryvector?language=objc
type TemporaryVector struct {
	Vector
}

func TemporaryVectorFrom(ptr unsafe.Pointer) TemporaryVector {
	return TemporaryVector{
		Vector: VectorFrom(ptr),
	}
}

func (tc _TemporaryVectorClass) TemporaryVectorWithCommandBufferDescriptor(commandBuffer metal.PCommandBuffer, descriptor IVectorDescriptor) TemporaryVector {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[TemporaryVector](tc, objc.Sel("temporaryVectorWithCommandBuffer:descriptor:"), po0, objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryvector/2935550-temporaryvectorwithcommandbuffer?language=objc
func TemporaryVector_TemporaryVectorWithCommandBufferDescriptor(commandBuffer metal.PCommandBuffer, descriptor IVectorDescriptor) TemporaryVector {
	return TemporaryVectorClass.TemporaryVectorWithCommandBufferDescriptor(commandBuffer, descriptor)
}

func (tc _TemporaryVectorClass) Alloc() TemporaryVector {
	rv := objc.Call[TemporaryVector](tc, objc.Sel("alloc"))
	return rv
}

func TemporaryVector_Alloc() TemporaryVector {
	return TemporaryVectorClass.Alloc()
}

func (tc _TemporaryVectorClass) New() TemporaryVector {
	rv := objc.Call[TemporaryVector](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTemporaryVector() TemporaryVector {
	return TemporaryVectorClass.New()
}

func (t_ TemporaryVector) Init() TemporaryVector {
	rv := objc.Call[TemporaryVector](t_, objc.Sel("init"))
	return rv
}

func (t_ TemporaryVector) InitWithDeviceDescriptor(device metal.PDevice, descriptor IVectorDescriptor) TemporaryVector {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[TemporaryVector](t_, objc.Sel("initWithDevice:descriptor:"), po0, objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector/2942566-initwithdevice?language=objc
func NewTemporaryVectorWithDeviceDescriptor(device metal.PDevice, descriptor IVectorDescriptor) TemporaryVector {
	instance := TemporaryVectorClass.Alloc().InitWithDeviceDescriptor(device, descriptor)
	instance.Autorelease()
	return instance
}

func (t_ TemporaryVector) InitWithBufferDescriptor(buffer metal.PBuffer, descriptor IVectorDescriptor) TemporaryVector {
	po0 := objc.WrapAsProtocol("MTLBuffer", buffer)
	rv := objc.Call[TemporaryVector](t_, objc.Sel("initWithBuffer:descriptor:"), po0, objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector/2873346-initwithbuffer?language=objc
func NewTemporaryVectorWithBufferDescriptor(buffer metal.PBuffer, descriptor IVectorDescriptor) TemporaryVector {
	instance := TemporaryVectorClass.Alloc().InitWithBufferDescriptor(buffer, descriptor)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryvector/2935544-prefetchstoragewithcommandbuffer?language=objc
func (tc _TemporaryVectorClass) PrefetchStorageWithCommandBufferDescriptorList(commandBuffer metal.PCommandBuffer, descriptorList []IVectorDescriptor) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](tc, objc.Sel("prefetchStorageWithCommandBuffer:descriptorList:"), po0, descriptorList)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryvector/2935544-prefetchstoragewithcommandbuffer?language=objc
func TemporaryVector_PrefetchStorageWithCommandBufferDescriptorList(commandBuffer metal.PCommandBuffer, descriptorList []IVectorDescriptor) {
	TemporaryVectorClass.PrefetchStorageWithCommandBufferDescriptorList(commandBuffer, descriptorList)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryvector/2935547-readcount?language=objc
func (t_ TemporaryVector) ReadCount() uint {
	rv := objc.Call[uint](t_, objc.Sel("readCount"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryvector/2935547-readcount?language=objc
func (t_ TemporaryVector) SetReadCount(value uint) {
	objc.Call[objc.Void](t_, objc.Sel("setReadCount:"), value)
}
