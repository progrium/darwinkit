// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a tone curve filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citonecurve?language=objc
type PToneCurve interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetPoint4(value coregraphics.Point)
	HasSetPoint4() bool

	// optional
	Point4() coregraphics.Point
	HasPoint4() bool

	// optional
	SetPoint3(value coregraphics.Point)
	HasSetPoint3() bool

	// optional
	Point3() coregraphics.Point
	HasPoint3() bool

	// optional
	SetPoint2(value coregraphics.Point)
	HasSetPoint2() bool

	// optional
	Point2() coregraphics.Point
	HasPoint2() bool

	// optional
	SetPoint1(value coregraphics.Point)
	HasSetPoint1() bool

	// optional
	Point1() coregraphics.Point
	HasPoint1() bool

	// optional
	SetPoint0(value coregraphics.Point)
	HasSetPoint0() bool

	// optional
	Point0() coregraphics.Point
	HasPoint0() bool
}

// A concrete type wrapper for the [PToneCurve] protocol.
type ToneCurveWrapper struct {
	objc.Object
}

func (t_ ToneCurveWrapper) HasSetInputImage() bool {
	return t_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citonecurve/3228792-inputimage?language=objc
func (t_ ToneCurveWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](t_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (t_ ToneCurveWrapper) HasInputImage() bool {
	return t_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citonecurve/3228792-inputimage?language=objc
func (t_ ToneCurveWrapper) InputImage() Image {
	rv := objc.Call[Image](t_, objc.Sel("inputImage"))
	return rv
}

func (t_ ToneCurveWrapper) HasSetPoint4() bool {
	return t_.RespondsToSelector(objc.Sel("setPoint4:"))
}

// A vector containing the position of the fifth point of the tone curve. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citonecurve/3228797-point4?language=objc
func (t_ ToneCurveWrapper) SetPoint4(value coregraphics.Point) {
	objc.Call[objc.Void](t_, objc.Sel("setPoint4:"), value)
}

func (t_ ToneCurveWrapper) HasPoint4() bool {
	return t_.RespondsToSelector(objc.Sel("point4"))
}

// A vector containing the position of the fifth point of the tone curve. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citonecurve/3228797-point4?language=objc
func (t_ ToneCurveWrapper) Point4() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](t_, objc.Sel("point4"))
	return rv
}

func (t_ ToneCurveWrapper) HasSetPoint3() bool {
	return t_.RespondsToSelector(objc.Sel("setPoint3:"))
}

// A vector containing the position of the fourth point of the tone curve. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citonecurve/3228796-point3?language=objc
func (t_ ToneCurveWrapper) SetPoint3(value coregraphics.Point) {
	objc.Call[objc.Void](t_, objc.Sel("setPoint3:"), value)
}

func (t_ ToneCurveWrapper) HasPoint3() bool {
	return t_.RespondsToSelector(objc.Sel("point3"))
}

// A vector containing the position of the fourth point of the tone curve. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citonecurve/3228796-point3?language=objc
func (t_ ToneCurveWrapper) Point3() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](t_, objc.Sel("point3"))
	return rv
}

func (t_ ToneCurveWrapper) HasSetPoint2() bool {
	return t_.RespondsToSelector(objc.Sel("setPoint2:"))
}

// A vector containing the position of the third point of the tone curve. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citonecurve/3228795-point2?language=objc
func (t_ ToneCurveWrapper) SetPoint2(value coregraphics.Point) {
	objc.Call[objc.Void](t_, objc.Sel("setPoint2:"), value)
}

func (t_ ToneCurveWrapper) HasPoint2() bool {
	return t_.RespondsToSelector(objc.Sel("point2"))
}

// A vector containing the position of the third point of the tone curve. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citonecurve/3228795-point2?language=objc
func (t_ ToneCurveWrapper) Point2() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](t_, objc.Sel("point2"))
	return rv
}

func (t_ ToneCurveWrapper) HasSetPoint1() bool {
	return t_.RespondsToSelector(objc.Sel("setPoint1:"))
}

// A vector containing the position of the second point of the tone curve. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citonecurve/3228794-point1?language=objc
func (t_ ToneCurveWrapper) SetPoint1(value coregraphics.Point) {
	objc.Call[objc.Void](t_, objc.Sel("setPoint1:"), value)
}

func (t_ ToneCurveWrapper) HasPoint1() bool {
	return t_.RespondsToSelector(objc.Sel("point1"))
}

// A vector containing the position of the second point of the tone curve. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citonecurve/3228794-point1?language=objc
func (t_ ToneCurveWrapper) Point1() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](t_, objc.Sel("point1"))
	return rv
}

func (t_ ToneCurveWrapper) HasSetPoint0() bool {
	return t_.RespondsToSelector(objc.Sel("setPoint0:"))
}

// A vector containing the position of the first point of the tone curve. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citonecurve/3228793-point0?language=objc
func (t_ ToneCurveWrapper) SetPoint0(value coregraphics.Point) {
	objc.Call[objc.Void](t_, objc.Sel("setPoint0:"), value)
}

func (t_ ToneCurveWrapper) HasPoint0() bool {
	return t_.RespondsToSelector(objc.Sel("point0"))
}

// A vector containing the position of the first point of the tone curve. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citonecurve/3228793-point0?language=objc
func (t_ ToneCurveWrapper) Point0() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](t_, objc.Sel("point0"))
	return rv
}
