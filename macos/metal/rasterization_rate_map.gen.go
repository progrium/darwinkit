// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/objc"
)

// A compiled read-only object that determines how to apply variable rasterization rates when rendering. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratemap?language=objc
type PRasterizationRateMap interface {
	// optional
	MapPhysicalToScreenCoordinatesForLayer(physicalCoordinates Coordinate2D, layerIndex uint) Coordinate2D
	HasMapPhysicalToScreenCoordinatesForLayer() bool

	// optional
	PhysicalSizeForLayer(layerIndex uint) Size
	HasPhysicalSizeForLayer() bool

	// optional
	CopyParameterDataToBufferOffset(buffer BufferWrapper, offset uint)
	HasCopyParameterDataToBufferOffset() bool

	// optional
	MapScreenToPhysicalCoordinatesForLayer(screenCoordinates Coordinate2D, layerIndex uint) Coordinate2D
	HasMapScreenToPhysicalCoordinatesForLayer() bool

	// optional
	ParameterBufferSizeAndAlign() SizeAndAlign
	HasParameterBufferSizeAndAlign() bool

	// optional
	ScreenSize() Size
	HasScreenSize() bool

	// optional
	Device() PDevice
	HasDevice() bool

	// optional
	PhysicalGranularity() Size
	HasPhysicalGranularity() bool

	// optional
	Label() string
	HasLabel() bool

	// optional
	LayerCount() uint
	HasLayerCount() bool
}

// A concrete type wrapper for the [PRasterizationRateMap] protocol.
type RasterizationRateMapWrapper struct {
	objc.Object
}

func (r_ RasterizationRateMapWrapper) HasMapPhysicalToScreenCoordinatesForLayer() bool {
	return r_.RespondsToSelector(objc.Sel("mapPhysicalToScreenCoordinates:forLayer:"))
}

// Converts a point in physical coordinates inside a layer to its corresponding logical viewport coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratemap/3173179-mapphysicaltoscreencoordinates?language=objc
func (r_ RasterizationRateMapWrapper) MapPhysicalToScreenCoordinatesForLayer(physicalCoordinates Coordinate2D, layerIndex uint) Coordinate2D {
	rv := objc.Call[Coordinate2D](r_, objc.Sel("mapPhysicalToScreenCoordinates:forLayer:"), physicalCoordinates, layerIndex)
	return rv
}

func (r_ RasterizationRateMapWrapper) HasPhysicalSizeForLayer() bool {
	return r_.RespondsToSelector(objc.Sel("physicalSizeForLayer:"))
}

// Returns the dimensions, in pixels, of the area in the render target affected by the rasterization rate map. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratemap/3088871-physicalsizeforlayer?language=objc
func (r_ RasterizationRateMapWrapper) PhysicalSizeForLayer(layerIndex uint) Size {
	rv := objc.Call[Size](r_, objc.Sel("physicalSizeForLayer:"), layerIndex)
	return rv
}

func (r_ RasterizationRateMapWrapper) HasCopyParameterDataToBufferOffset() bool {
	return r_.RespondsToSelector(objc.Sel("copyParameterDataToBuffer:offset:"))
}

// Copies the parameter data into the provided buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratemap/3153125-copyparameterdatatobuffer?language=objc
func (r_ RasterizationRateMapWrapper) CopyParameterDataToBufferOffset(buffer PBuffer, offset uint) {
	po0 := objc.WrapAsProtocol("MTLBuffer", buffer)
	objc.Call[objc.Void](r_, objc.Sel("copyParameterDataToBuffer:offset:"), po0, offset)
}

func (r_ RasterizationRateMapWrapper) HasMapScreenToPhysicalCoordinatesForLayer() bool {
	return r_.RespondsToSelector(objc.Sel("mapScreenToPhysicalCoordinates:forLayer:"))
}

// Converts a point in logical viewport coordinates to the corresponding physical coordinates in a render layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratemap/3173180-mapscreentophysicalcoordinates?language=objc
func (r_ RasterizationRateMapWrapper) MapScreenToPhysicalCoordinatesForLayer(screenCoordinates Coordinate2D, layerIndex uint) Coordinate2D {
	rv := objc.Call[Coordinate2D](r_, objc.Sel("mapScreenToPhysicalCoordinates:forLayer:"), screenCoordinates, layerIndex)
	return rv
}

func (r_ RasterizationRateMapWrapper) HasParameterBufferSizeAndAlign() bool {
	return r_.RespondsToSelector(objc.Sel("parameterBufferSizeAndAlign"))
}

// The size and alignment requirements to contain the coordinate transformation information in this rate map. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratemap/3088868-parameterbuffersizeandalign?language=objc
func (r_ RasterizationRateMapWrapper) ParameterBufferSizeAndAlign() SizeAndAlign {
	rv := objc.Call[SizeAndAlign](r_, objc.Sel("parameterBufferSizeAndAlign"))
	return rv
}

func (r_ RasterizationRateMapWrapper) HasScreenSize() bool {
	return r_.RespondsToSelector(objc.Sel("screenSize"))
}

// The logical size, in pixels, of the viewport coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratemap/3088873-screensize?language=objc
func (r_ RasterizationRateMapWrapper) ScreenSize() Size {
	rv := objc.Call[Size](r_, objc.Sel("screenSize"))
	return rv
}

func (r_ RasterizationRateMapWrapper) HasDevice() bool {
	return r_.RespondsToSelector(objc.Sel("device"))
}

// The device object that created the rate map. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratemap/3088866-device?language=objc
func (r_ RasterizationRateMapWrapper) Device() DeviceWrapper {
	rv := objc.Call[DeviceWrapper](r_, objc.Sel("device"))
	return rv
}

func (r_ RasterizationRateMapWrapper) HasPhysicalGranularity() bool {
	return r_.RespondsToSelector(objc.Sel("physicalGranularity"))
}

// The granularity, in physical pixels, at which the rasterization rate varies. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratemap/3088870-physicalgranularity?language=objc
func (r_ RasterizationRateMapWrapper) PhysicalGranularity() Size {
	rv := objc.Call[Size](r_, objc.Sel("physicalGranularity"))
	return rv
}

func (r_ RasterizationRateMapWrapper) HasLabel() bool {
	return r_.RespondsToSelector(objc.Sel("label"))
}

// A string that identifies the rate map. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratemap/3131702-label?language=objc
func (r_ RasterizationRateMapWrapper) Label() string {
	rv := objc.Call[string](r_, objc.Sel("label"))
	return rv
}

func (r_ RasterizationRateMapWrapper) HasLayerCount() bool {
	return r_.RespondsToSelector(objc.Sel("layerCount"))
}

// The number of layers in the rate map. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratemap/3088867-layercount?language=objc
func (r_ RasterizationRateMapWrapper) LayerCount() uint {
	rv := objc.Call[uint](r_, objc.Sel("layerCount"))
	return rv
}
