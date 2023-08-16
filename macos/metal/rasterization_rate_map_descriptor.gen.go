// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RasterizationRateMapDescriptor] class.
var RasterizationRateMapDescriptorClass = _RasterizationRateMapDescriptorClass{objc.GetClass("MTLRasterizationRateMapDescriptor")}

type _RasterizationRateMapDescriptorClass struct {
	objc.Class
}

// An interface definition for the [RasterizationRateMapDescriptor] class.
type IRasterizationRateMapDescriptor interface {
	objc.IObject
	LayerAtIndex(layerIndex uint) RasterizationRateLayerDescriptor
	SetLayerAtIndex(layer IRasterizationRateLayerDescriptor, layerIndex uint)
	ScreenSize() Size
	SetScreenSize(value Size)
	Label() string
	SetLabel(value string)
	LayerCount() uint
	Layers() RasterizationRateLayerArray
}

// An object that you use to configure new rasterization rate maps. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratemapdescriptor?language=objc
type RasterizationRateMapDescriptor struct {
	objc.Object
}

func RasterizationRateMapDescriptorFrom(ptr unsafe.Pointer) RasterizationRateMapDescriptor {
	return RasterizationRateMapDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (rc _RasterizationRateMapDescriptorClass) Alloc() RasterizationRateMapDescriptor {
	rv := objc.Call[RasterizationRateMapDescriptor](rc, objc.Sel("alloc"))
	return rv
}

func RasterizationRateMapDescriptor_Alloc() RasterizationRateMapDescriptor {
	return RasterizationRateMapDescriptorClass.Alloc()
}

func (rc _RasterizationRateMapDescriptorClass) New() RasterizationRateMapDescriptor {
	rv := objc.Call[RasterizationRateMapDescriptor](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRasterizationRateMapDescriptor() RasterizationRateMapDescriptor {
	return RasterizationRateMapDescriptorClass.New()
}

func (r_ RasterizationRateMapDescriptor) Init() RasterizationRateMapDescriptor {
	rv := objc.Call[RasterizationRateMapDescriptor](r_, objc.Sel("init"))
	return rv
}

// Returns the layer description for a layer in the rate map. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratemapdescriptor/3131705-layeratindex?language=objc
func (r_ RasterizationRateMapDescriptor) LayerAtIndex(layerIndex uint) RasterizationRateLayerDescriptor {
	rv := objc.Call[RasterizationRateLayerDescriptor](r_, objc.Sel("layerAtIndex:"), layerIndex)
	return rv
}

// Sets a configuration for a layer rate map. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratemapdescriptor/3131712-setlayer?language=objc
func (r_ RasterizationRateMapDescriptor) SetLayerAtIndex(layer IRasterizationRateLayerDescriptor, layerIndex uint) {
	objc.Call[objc.Void](r_, objc.Sel("setLayer:atIndex:"), objc.Ptr(layer), layerIndex)
}

// Creates a rate map descriptor with a set of layer descriptors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratemapdescriptor/3131710-rasterizationratemapdescriptorwi?language=objc
func (rc _RasterizationRateMapDescriptorClass) RasterizationRateMapDescriptorWithScreenSizeLayerCountLayers(screenSize Size, layerCount uint, layers IRasterizationRateLayerDescriptor) RasterizationRateMapDescriptor {
	rv := objc.Call[RasterizationRateMapDescriptor](rc, objc.Sel("rasterizationRateMapDescriptorWithScreenSize:layerCount:layers:"), screenSize, layerCount, objc.Ptr(layers))
	return rv
}

// Creates a rate map descriptor with a set of layer descriptors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratemapdescriptor/3131710-rasterizationratemapdescriptorwi?language=objc
func RasterizationRateMapDescriptor_RasterizationRateMapDescriptorWithScreenSizeLayerCountLayers(screenSize Size, layerCount uint, layers IRasterizationRateLayerDescriptor) RasterizationRateMapDescriptor {
	return RasterizationRateMapDescriptorClass.RasterizationRateMapDescriptorWithScreenSizeLayerCountLayers(screenSize, layerCount, layers)
}

// The size of the viewport coordinate system, in logical pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratemapdescriptor/3131711-screensize?language=objc
func (r_ RasterizationRateMapDescriptor) ScreenSize() Size {
	rv := objc.Call[Size](r_, objc.Sel("screenSize"))
	return rv
}

// The size of the viewport coordinate system, in logical pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratemapdescriptor/3131711-screensize?language=objc
func (r_ RasterizationRateMapDescriptor) SetScreenSize(value Size) {
	objc.Call[objc.Void](r_, objc.Sel("setScreenSize:"), value)
}

// A string used to identify the rate map you create with the descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratemapdescriptor/3131704-label?language=objc
func (r_ RasterizationRateMapDescriptor) Label() string {
	rv := objc.Call[string](r_, objc.Sel("label"))
	return rv
}

// A string used to identify the rate map you create with the descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratemapdescriptor/3131704-label?language=objc
func (r_ RasterizationRateMapDescriptor) SetLabel(value string) {
	objc.Call[objc.Void](r_, objc.Sel("setLabel:"), value)
}

// The number of layers in the rate map. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratemapdescriptor/3131706-layercount?language=objc
func (r_ RasterizationRateMapDescriptor) LayerCount() uint {
	rv := objc.Call[uint](r_, objc.Sel("layerCount"))
	return rv
}

// The rasterization rates for one or more layers in the rate map. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratemapdescriptor/3131707-layers?language=objc
func (r_ RasterizationRateMapDescriptor) Layers() RasterizationRateLayerArray {
	rv := objc.Call[RasterizationRateLayerArray](r_, objc.Sel("layers"))
	return rv
}
