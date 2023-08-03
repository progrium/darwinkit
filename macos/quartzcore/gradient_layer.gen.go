// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var GradientLayerClass = _GradientLayerClass{objc.GetClass("CAGradientLayer")}

type _GradientLayerClass struct {
	objc.Class
}

type IGradientLayer interface {
	ILayer
	Colors() []objc.Object
	SetColors(value []objc.IObject)
	Locations() []foundation.Number
	SetLocations(value []foundation.INumber)
	EndPoint() coregraphics.Point
	SetEndPoint(value coregraphics.Point)
	StartPoint() coregraphics.Point
	SetStartPoint(value coregraphics.Point)
	Type() GradientLayerType
	SetType(value GradientLayerType)
}

type GradientLayer struct {
	Layer
}

func MakeGradientLayer(ptr unsafe.Pointer) GradientLayer {
	return GradientLayer{
		Layer: MakeLayer(ptr),
	}
}

func (gc _GradientLayerClass) Layer() GradientLayer {
	rv := objc.CallMethod[GradientLayer](gc, objc.GetSelector("layer"))
	return rv
}

func GradientLayer_Layer() GradientLayer {
	return GradientLayerClass.Layer()
}

func (g_ GradientLayer) Init() GradientLayer {
	rv := objc.CallMethod[GradientLayer](g_, objc.GetSelector("init"))
	return rv
}

func GradientLayer_Init() GradientLayer {
	return GradientLayerClass.Alloc().Init()
}

func (g_ GradientLayer) InitWithLayer(layer objc.IObject) GradientLayer {
	rv := objc.CallMethod[GradientLayer](g_, objc.GetSelector("initWithLayer:"), objc.ExtractPtr(layer))
	return rv
}

func GradientLayer_InitWithLayer(layer objc.IObject) GradientLayer {
	return GradientLayerClass.Alloc().InitWithLayer(layer)
}

func (g_ GradientLayer) PresentationLayer() GradientLayer {
	rv := objc.CallMethod[GradientLayer](g_, objc.GetSelector("presentationLayer"))
	return rv
}

func GradientLayer_PresentationLayer() GradientLayer {
	return GradientLayerClass.Alloc().PresentationLayer()
}

func (g_ GradientLayer) ModelLayer() GradientLayer {
	rv := objc.CallMethod[GradientLayer](g_, objc.GetSelector("modelLayer"))
	return rv
}

func GradientLayer_ModelLayer() GradientLayer {
	return GradientLayerClass.Alloc().ModelLayer()
}

func (gc _GradientLayerClass) Alloc() GradientLayer {
	rv := objc.CallMethod[GradientLayer](gc, objc.GetSelector("alloc"))
	return rv
}

func GradientLayer_Alloc() GradientLayer {
	return GradientLayerClass.Alloc()
}

func (gc _GradientLayerClass) New() GradientLayer {
	rv := objc.CallMethod[GradientLayer](gc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewGradientLayer() GradientLayer {
	return GradientLayerClass.New()
}

func GradientLayer_New() GradientLayer {
	return GradientLayerClass.New()
}

func (g_ GradientLayer) Colors() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](g_, objc.GetSelector("colors"))
	// TODO: convert slice items...
	return rv
}

func (g_ GradientLayer) SetColors(value []objc.IObject) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setColors:"), value)
}

func (g_ GradientLayer) Locations() []foundation.Number {
	rv := objc.CallMethod[[]foundation.Number](g_, objc.GetSelector("locations"))
	// TODO: convert slice items...
	return rv
}

func (g_ GradientLayer) SetLocations(value []foundation.INumber) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setLocations:"), value)
}

func (g_ GradientLayer) EndPoint() coregraphics.Point {
	rv := objc.CallMethod[coregraphics.Point](g_, objc.GetSelector("endPoint"))
	return rv
}

func (g_ GradientLayer) SetEndPoint(value coregraphics.Point) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setEndPoint:"), value)
}

func (g_ GradientLayer) StartPoint() coregraphics.Point {
	rv := objc.CallMethod[coregraphics.Point](g_, objc.GetSelector("startPoint"))
	return rv
}

func (g_ GradientLayer) SetStartPoint(value coregraphics.Point) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setStartPoint:"), value)
}

func (g_ GradientLayer) Type() GradientLayerType {
	rv := objc.CallMethod[GradientLayerType](g_, objc.GetSelector("type"))
	return rv
}

func (g_ GradientLayer) SetType(value GradientLayerType) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setType:"), value)
}
