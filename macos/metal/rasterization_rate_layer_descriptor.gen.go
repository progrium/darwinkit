// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RasterizationRateLayerDescriptor] class.
var RasterizationRateLayerDescriptorClass = _RasterizationRateLayerDescriptorClass{objc.GetClass("MTLRasterizationRateLayerDescriptor")}

type _RasterizationRateLayerDescriptorClass struct {
	objc.Class
}

// An interface definition for the [RasterizationRateLayerDescriptor] class.
type IRasterizationRateLayerDescriptor interface {
	objc.IObject
	HorizontalSampleStorage() *float64
	VerticalSampleStorage() *float64
	SampleCount() Size
	MaxSampleCount() Size
	Horizontal() RasterizationRateSampleArray
	Vertical() RasterizationRateSampleArray
}

// The minimum rasterization rates to apply to sections of a layer in the render target. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor?language=objc
type RasterizationRateLayerDescriptor struct {
	objc.Object
}

func RasterizationRateLayerDescriptorFrom(ptr unsafe.Pointer) RasterizationRateLayerDescriptor {
	return RasterizationRateLayerDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (r_ RasterizationRateLayerDescriptor) InitWithSampleCount(sampleCount Size) RasterizationRateLayerDescriptor {
	rv := objc.Call[RasterizationRateLayerDescriptor](r_, objc.Sel("initWithSampleCount:"), sampleCount)
	return rv
}

// Initializes the layer map with an empty grid. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/3088860-initwithsamplecount?language=objc
func NewRasterizationRateLayerDescriptorWithSampleCount(sampleCount Size) RasterizationRateLayerDescriptor {
	instance := RasterizationRateLayerDescriptorClass.Alloc().InitWithSampleCount(sampleCount)
	instance.Autorelease()
	return instance
}

func (rc _RasterizationRateLayerDescriptorClass) Alloc() RasterizationRateLayerDescriptor {
	rv := objc.Call[RasterizationRateLayerDescriptor](rc, objc.Sel("alloc"))
	return rv
}

func RasterizationRateLayerDescriptor_Alloc() RasterizationRateLayerDescriptor {
	return RasterizationRateLayerDescriptorClass.Alloc()
}

func (rc _RasterizationRateLayerDescriptorClass) New() RasterizationRateLayerDescriptor {
	rv := objc.Call[RasterizationRateLayerDescriptor](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRasterizationRateLayerDescriptor() RasterizationRateLayerDescriptor {
	return RasterizationRateLayerDescriptorClass.New()
}

func (r_ RasterizationRateLayerDescriptor) Init() RasterizationRateLayerDescriptor {
	rv := objc.Call[RasterizationRateLayerDescriptor](r_, objc.Sel("init"))
	return rv
}

// A pointer to the storage for the layer map's horizontal rasterization rates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/3088859-horizontalsamplestorage?language=objc
func (r_ RasterizationRateLayerDescriptor) HorizontalSampleStorage() *float64 {
	rv := objc.Call[*float64](r_, objc.Sel("horizontalSampleStorage"))
	return rv
}

// A pointer to the storage for the layer map's vertical rasterization rates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/3088863-verticalsamplestorage?language=objc
func (r_ RasterizationRateLayerDescriptor) VerticalSampleStorage() *float64 {
	rv := objc.Call[*float64](r_, objc.Sel("verticalSampleStorage"))
	return rv
}

// The number of rows and columns in the layer map. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/3088861-samplecount?language=objc
func (r_ RasterizationRateLayerDescriptor) SampleCount() Size {
	rv := objc.Call[Size](r_, objc.Sel("sampleCount"))
	return rv
}

// The maximum number of rows and columns in the layer map. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/3750552-maxsamplecount?language=objc
func (r_ RasterizationRateLayerDescriptor) MaxSampleCount() Size {
	rv := objc.Call[Size](r_, objc.Sel("maxSampleCount"))
	return rv
}

// The horizontal rasterization rates for the layer map’s rows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/3088858-horizontal?language=objc
func (r_ RasterizationRateLayerDescriptor) Horizontal() RasterizationRateSampleArray {
	rv := objc.Call[RasterizationRateSampleArray](r_, objc.Sel("horizontal"))
	return rv
}

// The vertical rasterization rates for the layer map’s rows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/3088862-vertical?language=objc
func (r_ RasterizationRateLayerDescriptor) Vertical() RasterizationRateSampleArray {
	rv := objc.Call[RasterizationRateSampleArray](r_, objc.Sel("vertical"))
	return rv
}
