// AUTO-GENERATED CODE, DO NOT MODIFY

package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [GradientLayer] class.
var GradientLayerClass = _GradientLayerClass{objc.GetClass("CAGradientLayer")}

type _GradientLayerClass struct {
	objc.Class
}

// An interface definition for the [GradientLayer] class.
type IGradientLayer interface {
	ILayer
	Locations() []foundation.Number
	SetLocations(value []foundation.INumber)
	EndPoint() coregraphics.Point
	SetEndPoint(value coregraphics.Point)
	StartPoint() coregraphics.Point
	SetStartPoint(value coregraphics.Point)
	Type() GradientLayerType
	SetType(value GradientLayerType)
	Colors() []objc.Object
	SetColors(value []objc.IObject)
}

// A layer that draws a color gradient over its background color, filling the shape of the layer (including rounded corners) [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cagradientlayer?language=objc
type GradientLayer struct {
	Layer
}

func GradientLayerFrom(ptr unsafe.Pointer) GradientLayer {
	return GradientLayer{
		Layer: LayerFrom(ptr),
	}
}

func (gc _GradientLayerClass) Alloc() GradientLayer {
	rv := objc.Call[GradientLayer](gc, objc.Sel("alloc"))
	return rv
}

func GradientLayer_Alloc() GradientLayer {
	return GradientLayerClass.Alloc()
}

func (gc _GradientLayerClass) New() GradientLayer {
	rv := objc.Call[GradientLayer](gc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewGradientLayer() GradientLayer {
	return GradientLayerClass.New()
}

func (g_ GradientLayer) Init() GradientLayer {
	rv := objc.Call[GradientLayer](g_, objc.Sel("init"))
	return rv
}

func (gc _GradientLayerClass) Layer() GradientLayer {
	rv := objc.Call[GradientLayer](gc, objc.Sel("layer"))
	return rv
}

// Creates and returns an instance of the layer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410793-layer?language=objc
func GradientLayer_Layer() GradientLayer {
	return GradientLayerClass.Layer()
}

func (g_ GradientLayer) InitWithLayer(layer objc.IObject) GradientLayer {
	rv := objc.Call[GradientLayer](g_, objc.Sel("initWithLayer:"), layer)
	return rv
}

// Override to copy or initialize custom fields of the specified layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410842-initwithlayer?language=objc
func GradientLayer_InitWithLayer(layer objc.IObject) GradientLayer {
	return GradientLayerClass.Alloc().InitWithLayer(layer)
}

func (g_ GradientLayer) ModelLayer() GradientLayer {
	rv := objc.Call[GradientLayer](g_, objc.Sel("modelLayer"))
	return rv
}

// Returns the model layer object associated with the receiver, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410853-modellayer?language=objc
func GradientLayer_ModelLayer() GradientLayer {
	return GradientLayerClass.Alloc().ModelLayer()
}

func (g_ GradientLayer) PresentationLayer() GradientLayer {
	rv := objc.Call[GradientLayer](g_, objc.Sel("presentationLayer"))
	return rv
}

// Returns a copy of the presentation layer object that represents the state of the layer as it currently appears onscreen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410744-presentationlayer?language=objc
func GradientLayer_PresentationLayer() GradientLayer {
	return GradientLayerClass.Alloc().PresentationLayer()
}

// An optional array of NSNumber objects defining the location of each gradient stop. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cagradientlayer/1462410-locations?language=objc
func (g_ GradientLayer) Locations() []foundation.Number {
	rv := objc.Call[[]foundation.Number](g_, objc.Sel("locations"))
	return rv
}

// An optional array of NSNumber objects defining the location of each gradient stop. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cagradientlayer/1462410-locations?language=objc
func (g_ GradientLayer) SetLocations(value []foundation.INumber) {
	objc.Call[objc.Void](g_, objc.Sel("setLocations:"), value)
}

// The end point of the gradient when drawn in the layer’s coordinate space. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cagradientlayer/1462412-endpoint?language=objc
func (g_ GradientLayer) EndPoint() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](g_, objc.Sel("endPoint"))
	return rv
}

// The end point of the gradient when drawn in the layer’s coordinate space. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cagradientlayer/1462412-endpoint?language=objc
func (g_ GradientLayer) SetEndPoint(value coregraphics.Point) {
	objc.Call[objc.Void](g_, objc.Sel("setEndPoint:"), value)
}

// The start point of the gradient when drawn in the layer’s coordinate space. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cagradientlayer/1462408-startpoint?language=objc
func (g_ GradientLayer) StartPoint() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](g_, objc.Sel("startPoint"))
	return rv
}

// The start point of the gradient when drawn in the layer’s coordinate space. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cagradientlayer/1462408-startpoint?language=objc
func (g_ GradientLayer) SetStartPoint(value coregraphics.Point) {
	objc.Call[objc.Void](g_, objc.Sel("setStartPoint:"), value)
}

// Style of gradient drawn by the layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cagradientlayer/1462413-type?language=objc
func (g_ GradientLayer) Type() GradientLayerType {
	rv := objc.Call[GradientLayerType](g_, objc.Sel("type"))
	return rv
}

// Style of gradient drawn by the layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cagradientlayer/1462413-type?language=objc
func (g_ GradientLayer) SetType(value GradientLayerType) {
	objc.Call[objc.Void](g_, objc.Sel("setType:"), value)
}

// An array of CGColorRef objects defining the color of each gradient stop. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cagradientlayer/1462403-colors?language=objc
func (g_ GradientLayer) Colors() []objc.Object {
	rv := objc.Call[[]objc.Object](g_, objc.Sel("colors"))
	return rv
}

// An array of CGColorRef objects defining the color of each gradient stop. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cagradientlayer/1462403-colors?language=objc
func (g_ GradientLayer) SetColors(value []objc.IObject) {
	objc.Call[objc.Void](g_, objc.Sel("setColors:"), value)
}
