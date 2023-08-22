// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReshape] class.
var NNReshapeClass = _NNReshapeClass{objc.GetClass("MPSNNReshape")}

type _NNReshapeClass struct {
	objc.Class
}

// An interface definition for the [NNReshape] class.
type INNReshape interface {
	ICNNKernel
	EncodeBatchToCommandBufferSourceImagesReshapedWidthReshapedHeightReshapedFeatureChannels(commandBuffer metal.PCommandBuffer, sourceImages *foundation.Array, reshapedWidth uint, reshapedHeight uint, reshapedFeatureChannels uint) *foundation.Array
	EncodeBatchToCommandBufferObjectSourceImagesReshapedWidthReshapedHeightReshapedFeatureChannels(commandBufferObject objc.IObject, sourceImages *foundation.Array, reshapedWidth uint, reshapedHeight uint, reshapedFeatureChannels uint) *foundation.Array
	EncodeToCommandBufferSourceImageReshapedWidthReshapedHeightReshapedFeatureChannels(commandBuffer metal.PCommandBuffer, sourceImage IImage, reshapedWidth uint, reshapedHeight uint, reshapedFeatureChannels uint) Image
	EncodeToCommandBufferObjectSourceImageReshapedWidthReshapedHeightReshapedFeatureChannels(commandBufferObject objc.IObject, sourceImage IImage, reshapedWidth uint, reshapedHeight uint, reshapedFeatureChannels uint) Image
}

// The base class for reshape operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshape?language=objc
type NNReshape struct {
	CNNKernel
}

func NNReshapeFrom(ptr unsafe.Pointer) NNReshape {
	return NNReshape{
		CNNKernel: CNNKernelFrom(ptr),
	}
}

func (n_ NNReshape) InitWithDevice(device metal.PDevice) NNReshape {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReshape](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshape/2951929-initwithdevice?language=objc
func NewNNReshapeWithDevice(device metal.PDevice) NNReshape {
	instance := NNReshapeClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (nc _NNReshapeClass) Alloc() NNReshape {
	rv := objc.Call[NNReshape](nc, objc.Sel("alloc"))
	return rv
}

func NNReshape_Alloc() NNReshape {
	return NNReshapeClass.Alloc()
}

func (nc _NNReshapeClass) New() NNReshape {
	rv := objc.Call[NNReshape](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReshape() NNReshape {
	return NNReshapeClass.New()
}

func (n_ NNReshape) Init() NNReshape {
	rv := objc.Call[NNReshape](n_, objc.Sel("init"))
	return rv
}

func (n_ NNReshape) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReshape {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReshape](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNReshape_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReshape {
	instance := NNReshapeClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshape/3547990-encodebatchtocommandbuffer?language=objc
func (n_ NNReshape) EncodeBatchToCommandBufferSourceImagesReshapedWidthReshapedHeightReshapedFeatureChannels(commandBuffer metal.PCommandBuffer, sourceImages *foundation.Array, reshapedWidth uint, reshapedHeight uint, reshapedFeatureChannels uint) *foundation.Array {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[*foundation.Array](n_, objc.Sel("encodeBatchToCommandBuffer:sourceImages:reshapedWidth:reshapedHeight:reshapedFeatureChannels:"), po0, sourceImages, reshapedWidth, reshapedHeight, reshapedFeatureChannels)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshape/3547990-encodebatchtocommandbuffer?language=objc
func (n_ NNReshape) EncodeBatchToCommandBufferObjectSourceImagesReshapedWidthReshapedHeightReshapedFeatureChannels(commandBufferObject objc.IObject, sourceImages *foundation.Array, reshapedWidth uint, reshapedHeight uint, reshapedFeatureChannels uint) *foundation.Array {
	rv := objc.Call[*foundation.Array](n_, objc.Sel("encodeBatchToCommandBuffer:sourceImages:reshapedWidth:reshapedHeight:reshapedFeatureChannels:"), objc.Ptr(commandBufferObject), sourceImages, reshapedWidth, reshapedHeight, reshapedFeatureChannels)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshape/3547992-encodetocommandbuffer?language=objc
func (n_ NNReshape) EncodeToCommandBufferSourceImageReshapedWidthReshapedHeightReshapedFeatureChannels(commandBuffer metal.PCommandBuffer, sourceImage IImage, reshapedWidth uint, reshapedHeight uint, reshapedFeatureChannels uint) Image {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[Image](n_, objc.Sel("encodeToCommandBuffer:sourceImage:reshapedWidth:reshapedHeight:reshapedFeatureChannels:"), po0, objc.Ptr(sourceImage), reshapedWidth, reshapedHeight, reshapedFeatureChannels)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshape/3547992-encodetocommandbuffer?language=objc
func (n_ NNReshape) EncodeToCommandBufferObjectSourceImageReshapedWidthReshapedHeightReshapedFeatureChannels(commandBufferObject objc.IObject, sourceImage IImage, reshapedWidth uint, reshapedHeight uint, reshapedFeatureChannels uint) Image {
	rv := objc.Call[Image](n_, objc.Sel("encodeToCommandBuffer:sourceImage:reshapedWidth:reshapedHeight:reshapedFeatureChannels:"), objc.Ptr(commandBufferObject), objc.Ptr(sourceImage), reshapedWidth, reshapedHeight, reshapedFeatureChannels)
	return rv
}
