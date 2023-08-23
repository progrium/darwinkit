// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [BinaryImageKernel] class.
var BinaryImageKernelClass = _BinaryImageKernelClass{objc.GetClass("MPSBinaryImageKernel")}

type _BinaryImageKernelClass struct {
	objc.Class
}

// An interface definition for the [BinaryImageKernel] class.
type IBinaryImageKernel interface {
	IKernel
	SecondarySourceRegionForDestinationSize(destinationSize metal.Size) Region
	EncodeToCommandBufferPrimaryImageSecondaryImageDestinationImage(commandBuffer metal.PCommandBuffer, primaryImage IImage, secondaryImage IImage, destinationImage IImage)
	EncodeToCommandBufferObjectPrimaryImageSecondaryImageDestinationImage(commandBufferObject objc.IObject, primaryImage IImage, secondaryImage IImage, destinationImage IImage)
	PrimarySourceRegionForDestinationSize(destinationSize metal.Size) Region
	SecondaryEdgeMode() ImageEdgeMode
	SetSecondaryEdgeMode(value ImageEdgeMode)
	PrimaryEdgeMode() ImageEdgeMode
	SetPrimaryEdgeMode(value ImageEdgeMode)
	PrimaryOffset() Offset
	SetPrimaryOffset(value Offset)
	SecondaryOffset() Offset
	SetSecondaryOffset(value Offset)
	ClipRect() metal.Region
	SetClipRect(value metal.Region)
}

// A kernel that consumes two textures and produces one texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel?language=objc
type BinaryImageKernel struct {
	Kernel
}

func BinaryImageKernelFrom(ptr unsafe.Pointer) BinaryImageKernel {
	return BinaryImageKernel{
		Kernel: KernelFrom(ptr),
	}
}

func (b_ BinaryImageKernel) InitWithDevice(device metal.PDevice) BinaryImageKernel {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[BinaryImageKernel](b_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/2866331-initwithdevice?language=objc
func NewBinaryImageKernelWithDevice(device metal.PDevice) BinaryImageKernel {
	instance := BinaryImageKernelClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (bc _BinaryImageKernelClass) Alloc() BinaryImageKernel {
	rv := objc.Call[BinaryImageKernel](bc, objc.Sel("alloc"))
	return rv
}

func BinaryImageKernel_Alloc() BinaryImageKernel {
	return BinaryImageKernelClass.Alloc()
}

func (bc _BinaryImageKernelClass) New() BinaryImageKernel {
	rv := objc.Call[BinaryImageKernel](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewBinaryImageKernel() BinaryImageKernel {
	return BinaryImageKernelClass.New()
}

func (b_ BinaryImageKernel) Init() BinaryImageKernel {
	rv := objc.Call[BinaryImageKernel](b_, objc.Sel("init"))
	return rv
}

func (b_ BinaryImageKernel) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) BinaryImageKernel {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[BinaryImageKernel](b_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func BinaryImageKernel_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) BinaryImageKernel {
	instance := BinaryImageKernelClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

// Determines the region of the secondary source texture that will be read for an encode operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/1618838-secondarysourceregionfordestinat?language=objc
func (b_ BinaryImageKernel) SecondarySourceRegionForDestinationSize(destinationSize metal.Size) Region {
	rv := objc.Call[Region](b_, objc.Sel("secondarySourceRegionForDestinationSize:"), destinationSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/2866330-encodetocommandbuffer?language=objc
func (b_ BinaryImageKernel) EncodeToCommandBufferPrimaryImageSecondaryImageDestinationImage(commandBuffer metal.PCommandBuffer, primaryImage IImage, secondaryImage IImage, destinationImage IImage) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](b_, objc.Sel("encodeToCommandBuffer:primaryImage:secondaryImage:destinationImage:"), po0, objc.Ptr(primaryImage), objc.Ptr(secondaryImage), objc.Ptr(destinationImage))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/2866330-encodetocommandbuffer?language=objc
func (b_ BinaryImageKernel) EncodeToCommandBufferObjectPrimaryImageSecondaryImageDestinationImage(commandBufferObject objc.IObject, primaryImage IImage, secondaryImage IImage, destinationImage IImage) {
	objc.Call[objc.Void](b_, objc.Sel("encodeToCommandBuffer:primaryImage:secondaryImage:destinationImage:"), objc.Ptr(commandBufferObject), objc.Ptr(primaryImage), objc.Ptr(secondaryImage), objc.Ptr(destinationImage))
}

// Determines the region of the primary source texture that will be read for an encode operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/1618900-primarysourceregionfordestinatio?language=objc
func (b_ BinaryImageKernel) PrimarySourceRegionForDestinationSize(destinationSize metal.Size) Region {
	rv := objc.Call[Region](b_, objc.Sel("primarySourceRegionForDestinationSize:"), destinationSize)
	return rv
}

// The edge mode to use when texture reads stray off the edge of the secondary source image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/1618848-secondaryedgemode?language=objc
func (b_ BinaryImageKernel) SecondaryEdgeMode() ImageEdgeMode {
	rv := objc.Call[ImageEdgeMode](b_, objc.Sel("secondaryEdgeMode"))
	return rv
}

// The edge mode to use when texture reads stray off the edge of the secondary source image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/1618848-secondaryedgemode?language=objc
func (b_ BinaryImageKernel) SetSecondaryEdgeMode(value ImageEdgeMode) {
	objc.Call[objc.Void](b_, objc.Sel("setSecondaryEdgeMode:"), value)
}

// The edge mode to use when texture reads stray off the edge of the primary source image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/1618782-primaryedgemode?language=objc
func (b_ BinaryImageKernel) PrimaryEdgeMode() ImageEdgeMode {
	rv := objc.Call[ImageEdgeMode](b_, objc.Sel("primaryEdgeMode"))
	return rv
}

// The edge mode to use when texture reads stray off the edge of the primary source image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/1618782-primaryedgemode?language=objc
func (b_ BinaryImageKernel) SetPrimaryEdgeMode(value ImageEdgeMode) {
	objc.Call[objc.Void](b_, objc.Sel("setPrimaryEdgeMode:"), value)
}

// The position of the destination clip rectangle origin relative to the primary source buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/1618880-primaryoffset?language=objc
func (b_ BinaryImageKernel) PrimaryOffset() Offset {
	rv := objc.Call[Offset](b_, objc.Sel("primaryOffset"))
	return rv
}

// The position of the destination clip rectangle origin relative to the primary source buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/1618880-primaryoffset?language=objc
func (b_ BinaryImageKernel) SetPrimaryOffset(value Offset) {
	objc.Call[objc.Void](b_, objc.Sel("setPrimaryOffset:"), value)
}

// The position of the destination clip rectangle origin relative to the secondary source buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/1618755-secondaryoffset?language=objc
func (b_ BinaryImageKernel) SecondaryOffset() Offset {
	rv := objc.Call[Offset](b_, objc.Sel("secondaryOffset"))
	return rv
}

// The position of the destination clip rectangle origin relative to the secondary source buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/1618755-secondaryoffset?language=objc
func (b_ BinaryImageKernel) SetSecondaryOffset(value Offset) {
	objc.Call[objc.Void](b_, objc.Sel("setSecondaryOffset:"), value)
}

// An optional clip rectangle to use when writing data. Only the pixels in the rectangle will be overwritten. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/1618879-cliprect?language=objc
func (b_ BinaryImageKernel) ClipRect() metal.Region {
	rv := objc.Call[metal.Region](b_, objc.Sel("clipRect"))
	return rv
}

// An optional clip rectangle to use when writing data. Only the pixels in the rectangle will be overwritten. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/1618879-cliprect?language=objc
func (b_ BinaryImageKernel) SetClipRect(value metal.Region) {
	objc.Call[objc.Void](b_, objc.Sel("setClipRect:"), value)
}
