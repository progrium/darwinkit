// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageEuclideanDistanceTransform] class.
var ImageEuclideanDistanceTransformClass = _ImageEuclideanDistanceTransformClass{objc.GetClass("MPSImageEuclideanDistanceTransform")}

type _ImageEuclideanDistanceTransformClass struct {
	objc.Class
}

// An interface definition for the [ImageEuclideanDistanceTransform] class.
type IImageEuclideanDistanceTransform interface {
	IUnaryImageKernel
	SearchLimitRadius() float64
	SetSearchLimitRadius(value float64)
}

// A filter that performs a Euclidean distance transform on an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageeuclideandistancetransform?language=objc
type ImageEuclideanDistanceTransform struct {
	UnaryImageKernel
}

func ImageEuclideanDistanceTransformFrom(ptr unsafe.Pointer) ImageEuclideanDistanceTransform {
	return ImageEuclideanDistanceTransform{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

func (i_ ImageEuclideanDistanceTransform) InitWithDevice(device metal.PDevice) ImageEuclideanDistanceTransform {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageEuclideanDistanceTransform](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Creates a Euclidean distance transform that runs on a specified device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageeuclideandistancetransform/2953973-initwithdevice?language=objc
func NewImageEuclideanDistanceTransformWithDevice(device metal.PDevice) ImageEuclideanDistanceTransform {
	instance := ImageEuclideanDistanceTransformClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (ic _ImageEuclideanDistanceTransformClass) Alloc() ImageEuclideanDistanceTransform {
	rv := objc.Call[ImageEuclideanDistanceTransform](ic, objc.Sel("alloc"))
	return rv
}

func ImageEuclideanDistanceTransform_Alloc() ImageEuclideanDistanceTransform {
	return ImageEuclideanDistanceTransformClass.Alloc()
}

func (ic _ImageEuclideanDistanceTransformClass) New() ImageEuclideanDistanceTransform {
	rv := objc.Call[ImageEuclideanDistanceTransform](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageEuclideanDistanceTransform() ImageEuclideanDistanceTransform {
	return ImageEuclideanDistanceTransformClass.New()
}

func (i_ ImageEuclideanDistanceTransform) Init() ImageEuclideanDistanceTransform {
	rv := objc.Call[ImageEuclideanDistanceTransform](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageEuclideanDistanceTransform) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageEuclideanDistanceTransform {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageEuclideanDistanceTransform](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageEuclideanDistanceTransform_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageEuclideanDistanceTransform {
	instance := ImageEuclideanDistanceTransformClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

// Limits the search in an image from a pixel to the closest nonzero pixel within a specified radius. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageeuclideandistancetransform/3547977-searchlimitradius?language=objc
func (i_ ImageEuclideanDistanceTransform) SearchLimitRadius() float64 {
	rv := objc.Call[float64](i_, objc.Sel("searchLimitRadius"))
	return rv
}

// Limits the search in an image from a pixel to the closest nonzero pixel within a specified radius. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageeuclideandistancetransform/3547977-searchlimitradius?language=objc
func (i_ ImageEuclideanDistanceTransform) SetSearchLimitRadius(value float64) {
	objc.Call[objc.Void](i_, objc.Sel("setSearchLimitRadius:"), value)
}
