// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RasterizationRateLayerArray] class.
var RasterizationRateLayerArrayClass = _RasterizationRateLayerArrayClass{objc.GetClass("MTLRasterizationRateLayerArray")}

type _RasterizationRateLayerArrayClass struct {
	objc.Class
}

// An interface definition for the [RasterizationRateLayerArray] class.
type IRasterizationRateLayerArray interface {
	objc.IObject
	ObjectAtIndexedSubscript(layerIndex uint) RasterizationRateLayerDescriptor
	SetObjectAtIndexedSubscript(layer IRasterizationRateLayerDescriptor, layerIndex uint)
}

// Descriptions for the rasterization rates to apply to the set of layers in a rate map. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerarray?language=objc
type RasterizationRateLayerArray struct {
	objc.Object
}

func RasterizationRateLayerArrayFrom(ptr unsafe.Pointer) RasterizationRateLayerArray {
	return RasterizationRateLayerArray{
		Object: objc.ObjectFrom(ptr),
	}
}

func (rc _RasterizationRateLayerArrayClass) Alloc() RasterizationRateLayerArray {
	rv := objc.Call[RasterizationRateLayerArray](rc, objc.Sel("alloc"))
	return rv
}

func RasterizationRateLayerArray_Alloc() RasterizationRateLayerArray {
	return RasterizationRateLayerArrayClass.Alloc()
}

func (rc _RasterizationRateLayerArrayClass) New() RasterizationRateLayerArray {
	rv := objc.Call[RasterizationRateLayerArray](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRasterizationRateLayerArray() RasterizationRateLayerArray {
	return RasterizationRateLayerArrayClass.New()
}

func (r_ RasterizationRateLayerArray) Init() RasterizationRateLayerArray {
	rv := objc.Call[RasterizationRateLayerArray](r_, objc.Sel("init"))
	return rv
}

// Retrieves the sample value at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerarray/3131699-objectatindexedsubscript?language=objc
func (r_ RasterizationRateLayerArray) ObjectAtIndexedSubscript(layerIndex uint) RasterizationRateLayerDescriptor {
	rv := objc.Call[RasterizationRateLayerDescriptor](r_, objc.Sel("objectAtIndexedSubscript:"), layerIndex)
	return rv
}

// Stores a sample value at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerarray/3131700-setobject?language=objc
func (r_ RasterizationRateLayerArray) SetObjectAtIndexedSubscript(layer IRasterizationRateLayerDescriptor, layerIndex uint) {
	objc.Call[objc.Void](r_, objc.Sel("setObject:atIndexedSubscript:"), objc.Ptr(layer), layerIndex)
}
