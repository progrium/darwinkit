// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageDescriptor] class.
var ImageDescriptorClass = _ImageDescriptorClass{objc.GetClass("MPSImageDescriptor")}

type _ImageDescriptorClass struct {
	objc.Class
}

// An interface definition for the [ImageDescriptor] class.
type IImageDescriptor interface {
	objc.IObject
	Width() uint
	SetWidth(value uint)
	Usage() metal.TextureUsage
	SetUsage(value metal.TextureUsage)
	NumberOfImages() uint
	SetNumberOfImages(value uint)
	Height() uint
	SetHeight(value uint)
	CpuCacheMode() metal.CPUCacheMode
	SetCpuCacheMode(value metal.CPUCacheMode)
	PixelFormat() metal.PixelFormat
	StorageMode() metal.StorageMode
	SetStorageMode(value metal.StorageMode)
	FeatureChannels() uint
	SetFeatureChannels(value uint)
	ChannelFormat() ImageFeatureChannelFormat
	SetChannelFormat(value ImageFeatureChannelFormat)
}

// A description of the attributes used to create an MPSImage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor?language=objc
type ImageDescriptor struct {
	objc.Object
}

func ImageDescriptorFrom(ptr unsafe.Pointer) ImageDescriptor {
	return ImageDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ic _ImageDescriptorClass) ImageDescriptorWithChannelFormatWidthHeightFeatureChannels(channelFormat ImageFeatureChannelFormat, width uint, height uint, featureChannels uint) ImageDescriptor {
	rv := objc.Call[ImageDescriptor](ic, objc.Sel("imageDescriptorWithChannelFormat:width:height:featureChannels:"), channelFormat, width, height, featureChannels)
	return rv
}

// Creates an image descriptor for a single image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648819-imagedescriptorwithchannelformat?language=objc
func ImageDescriptor_ImageDescriptorWithChannelFormatWidthHeightFeatureChannels(channelFormat ImageFeatureChannelFormat, width uint, height uint, featureChannels uint) ImageDescriptor {
	return ImageDescriptorClass.ImageDescriptorWithChannelFormatWidthHeightFeatureChannels(channelFormat, width, height, featureChannels)
}

func (i_ ImageDescriptor) CopyWithZone(zone unsafe.Pointer) ImageDescriptor {
	rv := objc.Call[ImageDescriptor](i_, objc.Sel("copyWithZone:"), zone)
	rv.Autorelease()
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/3020686-copywithzone?language=objc
func ImageDescriptor_CopyWithZone(zone unsafe.Pointer) ImageDescriptor {
	instance := ImageDescriptorClass.Alloc().CopyWithZone(zone)
	instance.Autorelease()
	return instance
}

func (ic _ImageDescriptorClass) Alloc() ImageDescriptor {
	rv := objc.Call[ImageDescriptor](ic, objc.Sel("alloc"))
	return rv
}

func ImageDescriptor_Alloc() ImageDescriptor {
	return ImageDescriptorClass.Alloc()
}

func (ic _ImageDescriptorClass) New() ImageDescriptor {
	rv := objc.Call[ImageDescriptor](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageDescriptor() ImageDescriptor {
	return ImageDescriptorClass.New()
}

func (i_ ImageDescriptor) Init() ImageDescriptor {
	rv := objc.Call[ImageDescriptor](i_, objc.Sel("init"))
	return rv
}

// The width of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648830-width?language=objc
func (i_ ImageDescriptor) Width() uint {
	rv := objc.Call[uint](i_, objc.Sel("width"))
	return rv
}

// The width of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648830-width?language=objc
func (i_ ImageDescriptor) SetWidth(value uint) {
	objc.Call[objc.Void](i_, objc.Sel("setWidth:"), value)
}

// Options to specify the intended usage of the underlying texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648937-usage?language=objc
func (i_ ImageDescriptor) Usage() metal.TextureUsage {
	rv := objc.Call[metal.TextureUsage](i_, objc.Sel("usage"))
	return rv
}

// Options to specify the intended usage of the underlying texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648937-usage?language=objc
func (i_ ImageDescriptor) SetUsage(value metal.TextureUsage) {
	objc.Call[objc.Void](i_, objc.Sel("setUsage:"), value)
}

// The number of images for batch processing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648846-numberofimages?language=objc
func (i_ ImageDescriptor) NumberOfImages() uint {
	rv := objc.Call[uint](i_, objc.Sel("numberOfImages"))
	return rv
}

// The number of images for batch processing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648846-numberofimages?language=objc
func (i_ ImageDescriptor) SetNumberOfImages(value uint) {
	objc.Call[objc.Void](i_, objc.Sel("setNumberOfImages:"), value)
}

// The height of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648947-height?language=objc
func (i_ ImageDescriptor) Height() uint {
	rv := objc.Call[uint](i_, objc.Sel("height"))
	return rv
}

// The height of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648947-height?language=objc
func (i_ ImageDescriptor) SetHeight(value uint) {
	objc.Call[objc.Void](i_, objc.Sel("setHeight:"), value)
}

// The CPU cache mode of the underlying texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648930-cpucachemode?language=objc
func (i_ ImageDescriptor) CpuCacheMode() metal.CPUCacheMode {
	rv := objc.Call[metal.CPUCacheMode](i_, objc.Sel("cpuCacheMode"))
	return rv
}

// The CPU cache mode of the underlying texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648930-cpucachemode?language=objc
func (i_ ImageDescriptor) SetCpuCacheMode(value metal.CPUCacheMode) {
	objc.Call[objc.Void](i_, objc.Sel("setCpuCacheMode:"), value)
}

// The pixel format for the underlying texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648913-pixelformat?language=objc
func (i_ ImageDescriptor) PixelFormat() metal.PixelFormat {
	rv := objc.Call[metal.PixelFormat](i_, objc.Sel("pixelFormat"))
	return rv
}

// The storage mode of underlying texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648955-storagemode?language=objc
func (i_ ImageDescriptor) StorageMode() metal.StorageMode {
	rv := objc.Call[metal.StorageMode](i_, objc.Sel("storageMode"))
	return rv
}

// The storage mode of underlying texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648955-storagemode?language=objc
func (i_ ImageDescriptor) SetStorageMode(value metal.StorageMode) {
	objc.Call[objc.Void](i_, objc.Sel("setStorageMode:"), value)
}

// The number of feature channels per pixel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648918-featurechannels?language=objc
func (i_ ImageDescriptor) FeatureChannels() uint {
	rv := objc.Call[uint](i_, objc.Sel("featureChannels"))
	return rv
}

// The number of feature channels per pixel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648918-featurechannels?language=objc
func (i_ ImageDescriptor) SetFeatureChannels(value uint) {
	objc.Call[objc.Void](i_, objc.Sel("setFeatureChannels:"), value)
}

// The storage format to use for each channel in the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648818-channelformat?language=objc
func (i_ ImageDescriptor) ChannelFormat() ImageFeatureChannelFormat {
	rv := objc.Call[ImageFeatureChannelFormat](i_, objc.Sel("channelFormat"))
	return rv
}

// The storage format to use for each channel in the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedescriptor/1648818-channelformat?language=objc
func (i_ ImageDescriptor) SetChannelFormat(value ImageFeatureChannelFormat) {
	objc.Call[objc.Void](i_, objc.Sel("setChannelFormat:"), value)
}
