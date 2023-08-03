// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var ShapeLayerClass = _ShapeLayerClass{objc.GetClass("CAShapeLayer")}

type _ShapeLayerClass struct {
	objc.Class
}

type IShapeLayer interface {
	ILayer
	Path() coregraphics.PathRef
	SetPath(value coregraphics.PathRef)
	FillColor() coregraphics.ColorRef
	SetFillColor(value coregraphics.ColorRef)
	FillRule() ShapeLayerFillRule
	SetFillRule(value ShapeLayerFillRule)
	LineCap() ShapeLayerLineCap
	SetLineCap(value ShapeLayerLineCap)
	LineDashPattern() []foundation.Number
	SetLineDashPattern(value []foundation.INumber)
	LineDashPhase() float64
	SetLineDashPhase(value float64)
	LineJoin() ShapeLayerLineJoin
	SetLineJoin(value ShapeLayerLineJoin)
	LineWidth() float64
	SetLineWidth(value float64)
	MiterLimit() float64
	SetMiterLimit(value float64)
	StrokeColor() coregraphics.ColorRef
	SetStrokeColor(value coregraphics.ColorRef)
	StrokeStart() float64
	SetStrokeStart(value float64)
	StrokeEnd() float64
	SetStrokeEnd(value float64)
}

type ShapeLayer struct {
	Layer
}

func MakeShapeLayer(ptr unsafe.Pointer) ShapeLayer {
	return ShapeLayer{
		Layer: MakeLayer(ptr),
	}
}

func (sc _ShapeLayerClass) Layer() ShapeLayer {
	rv := objc.CallMethod[ShapeLayer](sc, objc.GetSelector("layer"))
	return rv
}

func ShapeLayer_Layer() ShapeLayer {
	return ShapeLayerClass.Layer()
}

func (s_ ShapeLayer) Init() ShapeLayer {
	rv := objc.CallMethod[ShapeLayer](s_, objc.GetSelector("init"))
	return rv
}

func ShapeLayer_Init() ShapeLayer {
	return ShapeLayerClass.Alloc().Init()
}

func (s_ ShapeLayer) InitWithLayer(layer objc.IObject) ShapeLayer {
	rv := objc.CallMethod[ShapeLayer](s_, objc.GetSelector("initWithLayer:"), objc.ExtractPtr(layer))
	return rv
}

func ShapeLayer_InitWithLayer(layer objc.IObject) ShapeLayer {
	return ShapeLayerClass.Alloc().InitWithLayer(layer)
}

func (s_ ShapeLayer) PresentationLayer() ShapeLayer {
	rv := objc.CallMethod[ShapeLayer](s_, objc.GetSelector("presentationLayer"))
	return rv
}

func ShapeLayer_PresentationLayer() ShapeLayer {
	return ShapeLayerClass.Alloc().PresentationLayer()
}

func (s_ ShapeLayer) ModelLayer() ShapeLayer {
	rv := objc.CallMethod[ShapeLayer](s_, objc.GetSelector("modelLayer"))
	return rv
}

func ShapeLayer_ModelLayer() ShapeLayer {
	return ShapeLayerClass.Alloc().ModelLayer()
}

func (sc _ShapeLayerClass) Alloc() ShapeLayer {
	rv := objc.CallMethod[ShapeLayer](sc, objc.GetSelector("alloc"))
	return rv
}

func ShapeLayer_Alloc() ShapeLayer {
	return ShapeLayerClass.Alloc()
}

func (sc _ShapeLayerClass) New() ShapeLayer {
	rv := objc.CallMethod[ShapeLayer](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewShapeLayer() ShapeLayer {
	return ShapeLayerClass.New()
}

func ShapeLayer_New() ShapeLayer {
	return ShapeLayerClass.New()
}

func (s_ ShapeLayer) Path() coregraphics.PathRef {
	rv := objc.CallMethod[coregraphics.PathRef](s_, objc.GetSelector("path"))
	return rv
}

func (s_ ShapeLayer) SetPath(value coregraphics.PathRef) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setPath:"), value)
}

func (s_ ShapeLayer) FillColor() coregraphics.ColorRef {
	rv := objc.CallMethod[coregraphics.ColorRef](s_, objc.GetSelector("fillColor"))
	return rv
}

func (s_ ShapeLayer) SetFillColor(value coregraphics.ColorRef) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setFillColor:"), value)
}

func (s_ ShapeLayer) FillRule() ShapeLayerFillRule {
	rv := objc.CallMethod[ShapeLayerFillRule](s_, objc.GetSelector("fillRule"))
	return rv
}

func (s_ ShapeLayer) SetFillRule(value ShapeLayerFillRule) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setFillRule:"), value)
}

func (s_ ShapeLayer) LineCap() ShapeLayerLineCap {
	rv := objc.CallMethod[ShapeLayerLineCap](s_, objc.GetSelector("lineCap"))
	return rv
}

func (s_ ShapeLayer) SetLineCap(value ShapeLayerLineCap) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setLineCap:"), value)
}

func (s_ ShapeLayer) LineDashPattern() []foundation.Number {
	rv := objc.CallMethod[[]foundation.Number](s_, objc.GetSelector("lineDashPattern"))
	// TODO: convert slice items...
	return rv
}

func (s_ ShapeLayer) SetLineDashPattern(value []foundation.INumber) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setLineDashPattern:"), value)
}

func (s_ ShapeLayer) LineDashPhase() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("lineDashPhase"))
	return rv
}

func (s_ ShapeLayer) SetLineDashPhase(value float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setLineDashPhase:"), value)
}

func (s_ ShapeLayer) LineJoin() ShapeLayerLineJoin {
	rv := objc.CallMethod[ShapeLayerLineJoin](s_, objc.GetSelector("lineJoin"))
	return rv
}

func (s_ ShapeLayer) SetLineJoin(value ShapeLayerLineJoin) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setLineJoin:"), value)
}

func (s_ ShapeLayer) LineWidth() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("lineWidth"))
	return rv
}

func (s_ ShapeLayer) SetLineWidth(value float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setLineWidth:"), value)
}

func (s_ ShapeLayer) MiterLimit() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("miterLimit"))
	return rv
}

func (s_ ShapeLayer) SetMiterLimit(value float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setMiterLimit:"), value)
}

func (s_ ShapeLayer) StrokeColor() coregraphics.ColorRef {
	rv := objc.CallMethod[coregraphics.ColorRef](s_, objc.GetSelector("strokeColor"))
	return rv
}

func (s_ ShapeLayer) SetStrokeColor(value coregraphics.ColorRef) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setStrokeColor:"), value)
}

func (s_ ShapeLayer) StrokeStart() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("strokeStart"))
	return rv
}

func (s_ ShapeLayer) SetStrokeStart(value float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setStrokeStart:"), value)
}

func (s_ ShapeLayer) StrokeEnd() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("strokeEnd"))
	return rv
}

func (s_ ShapeLayer) SetStrokeEnd(value float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setStrokeEnd:"), value)
}
