// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UnaryImageKernel] class.
var UnaryImageKernelClass = _UnaryImageKernelClass{objc.GetClass("MPSUnaryImageKernel")}

type _UnaryImageKernelClass struct {
	objc.Class
}

// An interface definition for the [UnaryImageKernel] class.
type IUnaryImageKernel interface {
	IKernel
	SourceRegionForDestinationSize(destinationSize metal.Size) Region
	EncodeToCommandBufferSourceImageDestinationImage(commandBuffer metal.PCommandBuffer, sourceImage IImage, destinationImage IImage)
	EncodeToCommandBufferObjectSourceImageDestinationImage(commandBufferObject objc.IObject, sourceImage IImage, destinationImage IImage)
	EdgeMode() ImageEdgeMode
	SetEdgeMode(value ImageEdgeMode)
	Offset() Offset
	SetOffset(value Offset)
	ClipRect() metal.Region
	SetClipRect(value metal.Region)
}

// A kernel that consumes one texture and produces one texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel?language=objc
type UnaryImageKernel struct {
	Kernel
}

func UnaryImageKernelFrom(ptr unsafe.Pointer) UnaryImageKernel {
	return UnaryImageKernel{
		Kernel: KernelFrom(ptr),
	}
}

func (u_ UnaryImageKernel) InitWithDevice(device metal.PDevice) UnaryImageKernel {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[UnaryImageKernel](u_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/2866332-initwithdevice?language=objc
func NewUnaryImageKernelWithDevice(device metal.PDevice) UnaryImageKernel {
	instance := UnaryImageKernelClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (uc _UnaryImageKernelClass) Alloc() UnaryImageKernel {
	rv := objc.Call[UnaryImageKernel](uc, objc.Sel("alloc"))
	return rv
}

func UnaryImageKernel_Alloc() UnaryImageKernel {
	return UnaryImageKernelClass.Alloc()
}

func (uc _UnaryImageKernelClass) New() UnaryImageKernel {
	rv := objc.Call[UnaryImageKernel](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUnaryImageKernel() UnaryImageKernel {
	return UnaryImageKernelClass.New()
}

func (u_ UnaryImageKernel) Init() UnaryImageKernel {
	rv := objc.Call[UnaryImageKernel](u_, objc.Sel("init"))
	return rv
}

func (u_ UnaryImageKernel) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) UnaryImageKernel {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[UnaryImageKernel](u_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func UnaryImageKernel_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) UnaryImageKernel {
	instance := UnaryImageKernelClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

// Determines the region of the source texture that will be read for an encode operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/1618754-sourceregionfordestinationsize?language=objc
func (u_ UnaryImageKernel) SourceRegionForDestinationSize(destinationSize metal.Size) Region {
	rv := objc.Call[Region](u_, objc.Sel("sourceRegionForDestinationSize:"), destinationSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/2866328-encodetocommandbuffer?language=objc
func (u_ UnaryImageKernel) EncodeToCommandBufferSourceImageDestinationImage(commandBuffer metal.PCommandBuffer, sourceImage IImage, destinationImage IImage) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](u_, objc.Sel("encodeToCommandBuffer:sourceImage:destinationImage:"), po0, objc.Ptr(sourceImage), objc.Ptr(destinationImage))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/2866328-encodetocommandbuffer?language=objc
func (u_ UnaryImageKernel) EncodeToCommandBufferObjectSourceImageDestinationImage(commandBufferObject objc.IObject, sourceImage IImage, destinationImage IImage) {
	objc.Call[objc.Void](u_, objc.Sel("encodeToCommandBuffer:sourceImage:destinationImage:"), objc.Ptr(commandBufferObject), objc.Ptr(sourceImage), objc.Ptr(destinationImage))
}

// The edge mode to use when texture reads stray off the edge of an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/1618812-edgemode?language=objc
func (u_ UnaryImageKernel) EdgeMode() ImageEdgeMode {
	rv := objc.Call[ImageEdgeMode](u_, objc.Sel("edgeMode"))
	return rv
}

// The edge mode to use when texture reads stray off the edge of an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/1618812-edgemode?language=objc
func (u_ UnaryImageKernel) SetEdgeMode(value ImageEdgeMode) {
	objc.Call[objc.Void](u_, objc.Sel("setEdgeMode:"), value)
}

// The position of the destination clip rectangle origin relative to the source buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/1618884-offset?language=objc
func (u_ UnaryImageKernel) Offset() Offset {
	rv := objc.Call[Offset](u_, objc.Sel("offset"))
	return rv
}

// The position of the destination clip rectangle origin relative to the source buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/1618884-offset?language=objc
func (u_ UnaryImageKernel) SetOffset(value Offset) {
	objc.Call[objc.Void](u_, objc.Sel("setOffset:"), value)
}

// An optional clip rectangle to use when writing data. Only the pixels in the rectangle will be overwritten. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/1618859-cliprect?language=objc
func (u_ UnaryImageKernel) ClipRect() metal.Region {
	rv := objc.Call[metal.Region](u_, objc.Sel("clipRect"))
	return rv
}

// An optional clip rectangle to use when writing data. Only the pixels in the rectangle will be overwritten. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/1618859-cliprect?language=objc
func (u_ UnaryImageKernel) SetClipRect(value metal.Region) {
	objc.Call[objc.Void](u_, objc.Sel("setClipRect:"), value)
}
