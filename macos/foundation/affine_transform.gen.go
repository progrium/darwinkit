// AUTO-GENERATED CODE, DO NOT MODIFY
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var AffineTransformClass = _AffineTransformClass{objc.GetClass("NSAffineTransform")}

type _AffineTransformClass struct {
	objc.Class
}

type IAffineTransform interface {
	objc.IObject
	RotateByDegrees(angle float64)
	RotateByRadians(angle float64)
	ScaleBy(scale float64)
	ScaleXByYBy(scaleX float64, scaleY float64)
	TranslateXByYBy(deltaX float64, deltaY float64)
	AppendTransform(transform IAffineTransform)
	PrependTransform(transform IAffineTransform)
	Invert()
	TransformPoint(aPoint Point) Point
	TransformSize(aSize Size) Size
	TransformStruct() AffineTransformStruct
	SetTransformStruct(value AffineTransformStruct)
}

type AffineTransform struct {
	objc.Object
}

func MakeAffineTransform(ptr unsafe.Pointer) AffineTransform {
	return AffineTransform{
		Object: objc.MakeObject(ptr),
	}
}

func (a_ AffineTransform) Init() AffineTransform {
	rv := objc.CallMethod[AffineTransform](a_, objc.GetSelector("init"))
	return rv
}

func AffineTransform_Init() AffineTransform {
	return AffineTransformClass.Alloc().Init()
}

func (a_ AffineTransform) InitWithTransform(transform IAffineTransform) AffineTransform {
	rv := objc.CallMethod[AffineTransform](a_, objc.GetSelector("initWithTransform:"), objc.ExtractPtr(transform))
	return rv
}

func AffineTransform_InitWithTransform(transform IAffineTransform) AffineTransform {
	return AffineTransformClass.Alloc().InitWithTransform(transform)
}

func (ac _AffineTransformClass) Alloc() AffineTransform {
	rv := objc.CallMethod[AffineTransform](ac, objc.GetSelector("alloc"))
	return rv
}

func AffineTransform_Alloc() AffineTransform {
	return AffineTransformClass.Alloc()
}

func (ac _AffineTransformClass) New() AffineTransform {
	rv := objc.CallMethod[AffineTransform](ac, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewAffineTransform() AffineTransform {
	return AffineTransformClass.New()
}

func AffineTransform_New() AffineTransform {
	return AffineTransformClass.New()
}

func (ac _AffineTransformClass) Transform() AffineTransform {
	rv := objc.CallMethod[AffineTransform](ac, objc.GetSelector("transform"))
	return rv
}

func AffineTransform_Transform() AffineTransform {
	return AffineTransformClass.Transform()
}

func (a_ AffineTransform) RotateByDegrees(angle float64) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("rotateByDegrees:"), angle)
}

func (a_ AffineTransform) RotateByRadians(angle float64) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("rotateByRadians:"), angle)
}

func (a_ AffineTransform) ScaleBy(scale float64) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("scaleBy:"), scale)
}

func (a_ AffineTransform) ScaleXByYBy(scaleX float64, scaleY float64) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("scaleXBy:yBy:"), scaleX, scaleY)
}

func (a_ AffineTransform) TranslateXByYBy(deltaX float64, deltaY float64) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("translateXBy:yBy:"), deltaX, deltaY)
}

func (a_ AffineTransform) AppendTransform(transform IAffineTransform) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("appendTransform:"), objc.ExtractPtr(transform))
}

func (a_ AffineTransform) PrependTransform(transform IAffineTransform) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("prependTransform:"), objc.ExtractPtr(transform))
}

func (a_ AffineTransform) Invert() {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("invert"))
}

func (a_ AffineTransform) TransformPoint(aPoint Point) Point {
	rv := objc.CallMethod[Point](a_, objc.GetSelector("transformPoint:"), aPoint)
	return rv
}

func (a_ AffineTransform) TransformSize(aSize Size) Size {
	rv := objc.CallMethod[Size](a_, objc.GetSelector("transformSize:"), aSize)
	return rv
}

func (a_ AffineTransform) TransformStruct() AffineTransformStruct {
	rv := objc.CallMethod[AffineTransformStruct](a_, objc.GetSelector("transformStruct"))
	return rv
}

func (a_ AffineTransform) SetTransformStruct(value AffineTransformStruct) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setTransformStruct:"), value)
}
