// AUTO-GENERATED CODE, DO NOT MODIFY

package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ShapeLayer] class.
var ShapeLayerClass = _ShapeLayerClass{objc.GetClass("CAShapeLayer")}

type _ShapeLayerClass struct {
	objc.Class
}

// An interface definition for the [ShapeLayer] class.
type IShapeLayer interface {
	ILayer
	Path() unsafe.Pointer
	SetPath(value unsafe.Pointer)
	LineWidth() float64
	SetLineWidth(value float64)
	StrokeStart() float64
	SetStrokeStart(value float64)
	LineDashPhase() float64
	SetLineDashPhase(value float64)
	MiterLimit() float64
	SetMiterLimit(value float64)
	StrokeEnd() float64
	SetStrokeEnd(value float64)
	LineJoin() ShapeLayerLineJoin
	SetLineJoin(value ShapeLayerLineJoin)
	LineDashPattern() []foundation.Number
	SetLineDashPattern(value []foundation.INumber)
	StrokeColor() coregraphics.ColorRef
	SetStrokeColor(value coregraphics.ColorRef)
	FillColor() coregraphics.ColorRef
	SetFillColor(value coregraphics.ColorRef)
	LineCap() ShapeLayerLineCap
	SetLineCap(value ShapeLayerLineCap)
	FillRule() ShapeLayerFillRule
	SetFillRule(value ShapeLayerFillRule)
}

// A layer that draws a cubic Bezier spline in its coordinate space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer?language=objc
type ShapeLayer struct {
	Layer
}

func ShapeLayerFrom(ptr unsafe.Pointer) ShapeLayer {
	return ShapeLayer{
		Layer: LayerFrom(ptr),
	}
}

func (sc _ShapeLayerClass) Alloc() ShapeLayer {
	rv := objc.Call[ShapeLayer](sc, objc.Sel("alloc"))
	return rv
}

func ShapeLayer_Alloc() ShapeLayer {
	return ShapeLayerClass.Alloc()
}

func (sc _ShapeLayerClass) New() ShapeLayer {
	rv := objc.Call[ShapeLayer](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewShapeLayer() ShapeLayer {
	return ShapeLayerClass.New()
}

func (s_ ShapeLayer) Init() ShapeLayer {
	rv := objc.Call[ShapeLayer](s_, objc.Sel("init"))
	return rv
}

func (sc _ShapeLayerClass) Layer() ShapeLayer {
	rv := objc.Call[ShapeLayer](sc, objc.Sel("layer"))
	return rv
}

// Creates and returns an instance of the layer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410793-layer?language=objc
func ShapeLayer_Layer() ShapeLayer {
	return ShapeLayerClass.Layer()
}

func (s_ ShapeLayer) InitWithLayer(layer objc.IObject) ShapeLayer {
	rv := objc.Call[ShapeLayer](s_, objc.Sel("initWithLayer:"), layer)
	return rv
}

// Override to copy or initialize custom fields of the specified layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410842-initwithlayer?language=objc
func NewShapeLayerWithLayer(layer objc.IObject) ShapeLayer {
	instance := ShapeLayerClass.Alloc().InitWithLayer(layer)
	instance.Autorelease()
	return instance
}

func (s_ ShapeLayer) ModelLayer() ShapeLayer {
	rv := objc.Call[ShapeLayer](s_, objc.Sel("modelLayer"))
	return rv
}

// Returns the model layer object associated with the receiver, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410853-modellayer?language=objc
func ShapeLayer_ModelLayer() ShapeLayer {
	instance := ShapeLayerClass.Alloc().ModelLayer()
	instance.Autorelease()
	return instance
}

func (s_ ShapeLayer) PresentationLayer() ShapeLayer {
	rv := objc.Call[ShapeLayer](s_, objc.Sel("presentationLayer"))
	return rv
}

// Returns a copy of the presentation layer object that represents the state of the layer as it currently appears onscreen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410744-presentationlayer?language=objc
func ShapeLayer_PresentationLayer() ShapeLayer {
	instance := ShapeLayerClass.Alloc().PresentationLayer()
	instance.Autorelease()
	return instance
}

// The path defining the shape to be rendered. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/1521904-path?language=objc
func (s_ ShapeLayer) Path() unsafe.Pointer {
	rv := objc.Call[unsafe.Pointer](s_, objc.Sel("path"))
	return rv
}

// The path defining the shape to be rendered. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/1521904-path?language=objc
func (s_ ShapeLayer) SetPath(value unsafe.Pointer) {
	objc.Call[objc.Void](s_, objc.Sel("setPath:"), value)
}

// Specifies the line width of the shape’s path. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/1521890-linewidth?language=objc
func (s_ ShapeLayer) LineWidth() float64 {
	rv := objc.Call[float64](s_, objc.Sel("lineWidth"))
	return rv
}

// Specifies the line width of the shape’s path. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/1521890-linewidth?language=objc
func (s_ ShapeLayer) SetLineWidth(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setLineWidth:"), value)
}

// The relative location at which to begin stroking the path. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/1521929-strokestart?language=objc
func (s_ ShapeLayer) StrokeStart() float64 {
	rv := objc.Call[float64](s_, objc.Sel("strokeStart"))
	return rv
}

// The relative location at which to begin stroking the path. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/1521929-strokestart?language=objc
func (s_ ShapeLayer) SetStrokeStart(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setStrokeStart:"), value)
}

// The dash phase applied to the shape’s path when stroked. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/1521856-linedashphase?language=objc
func (s_ ShapeLayer) LineDashPhase() float64 {
	rv := objc.Call[float64](s_, objc.Sel("lineDashPhase"))
	return rv
}

// The dash phase applied to the shape’s path when stroked. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/1521856-linedashphase?language=objc
func (s_ ShapeLayer) SetLineDashPhase(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setLineDashPhase:"), value)
}

// The miter limit used when stroking the shape’s path. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/1521870-miterlimit?language=objc
func (s_ ShapeLayer) MiterLimit() float64 {
	rv := objc.Call[float64](s_, objc.Sel("miterLimit"))
	return rv
}

// The miter limit used when stroking the shape’s path. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/1521870-miterlimit?language=objc
func (s_ ShapeLayer) SetMiterLimit(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setMiterLimit:"), value)
}

// The relative location at which to stop stroking the path. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/1522252-strokeend?language=objc
func (s_ ShapeLayer) StrokeEnd() float64 {
	rv := objc.Call[float64](s_, objc.Sel("strokeEnd"))
	return rv
}

// The relative location at which to stop stroking the path. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/1522252-strokeend?language=objc
func (s_ ShapeLayer) SetStrokeEnd(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setStrokeEnd:"), value)
}

// Specifies the line join style for the shape’s path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/1522147-linejoin?language=objc
func (s_ ShapeLayer) LineJoin() ShapeLayerLineJoin {
	rv := objc.Call[ShapeLayerLineJoin](s_, objc.Sel("lineJoin"))
	return rv
}

// Specifies the line join style for the shape’s path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/1522147-linejoin?language=objc
func (s_ ShapeLayer) SetLineJoin(value ShapeLayerLineJoin) {
	objc.Call[objc.Void](s_, objc.Sel("setLineJoin:"), value)
}

// The dash pattern applied to the shape’s path when stroked. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/1521921-linedashpattern?language=objc
func (s_ ShapeLayer) LineDashPattern() []foundation.Number {
	rv := objc.Call[[]foundation.Number](s_, objc.Sel("lineDashPattern"))
	return rv
}

// The dash pattern applied to the shape’s path when stroked. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/1521921-linedashpattern?language=objc
func (s_ ShapeLayer) SetLineDashPattern(value []foundation.INumber) {
	objc.Call[objc.Void](s_, objc.Sel("setLineDashPattern:"), value)
}

// The color used to stroke the shape’s path. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/1521897-strokecolor?language=objc
func (s_ ShapeLayer) StrokeColor() coregraphics.ColorRef {
	rv := objc.Call[coregraphics.ColorRef](s_, objc.Sel("strokeColor"))
	return rv
}

// The color used to stroke the shape’s path. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/1521897-strokecolor?language=objc
func (s_ ShapeLayer) SetStrokeColor(value coregraphics.ColorRef) {
	objc.Call[objc.Void](s_, objc.Sel("setStrokeColor:"), value)
}

// The color used to fill the shape’s path. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/1522248-fillcolor?language=objc
func (s_ ShapeLayer) FillColor() coregraphics.ColorRef {
	rv := objc.Call[coregraphics.ColorRef](s_, objc.Sel("fillColor"))
	return rv
}

// The color used to fill the shape’s path. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/1522248-fillcolor?language=objc
func (s_ ShapeLayer) SetFillColor(value coregraphics.ColorRef) {
	objc.Call[objc.Void](s_, objc.Sel("setFillColor:"), value)
}

// Specifies the line cap style for the shape’s path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/1521905-linecap?language=objc
func (s_ ShapeLayer) LineCap() ShapeLayerLineCap {
	rv := objc.Call[ShapeLayerLineCap](s_, objc.Sel("lineCap"))
	return rv
}

// Specifies the line cap style for the shape’s path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/1521905-linecap?language=objc
func (s_ ShapeLayer) SetLineCap(value ShapeLayerLineCap) {
	objc.Call[objc.Void](s_, objc.Sel("setLineCap:"), value)
}

// The fill rule used when filling the shape’s path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/1522146-fillrule?language=objc
func (s_ ShapeLayer) FillRule() ShapeLayerFillRule {
	rv := objc.Call[ShapeLayerFillRule](s_, objc.Sel("fillRule"))
	return rv
}

// The fill rule used when filling the shape’s path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/1522146-fillrule?language=objc
func (s_ ShapeLayer) SetFillRule(value ShapeLayerFillRule) {
	objc.Call[objc.Void](s_, objc.Sel("setFillRule:"), value)
}
