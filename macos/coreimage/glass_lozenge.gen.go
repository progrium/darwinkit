// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglasslozenge?language=objc
type PGlassLozenge interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetRefraction(value float64)
	HasSetRefraction() bool

	// optional
	Refraction() float64
	HasRefraction() bool

	// optional
	SetRadius(value float64)
	HasSetRadius() bool

	// optional
	Radius() float64
	HasRadius() bool

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

// A concrete type wrapper for the [PGlassLozenge] protocol.
type GlassLozengeWrapper struct {
	objc.Object
}

func (g_ GlassLozengeWrapper) HasSetInputImage() bool {
	return g_.RespondsToSelector(objc.Sel("setInputImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglasslozenge/3600161-inputimage?language=objc
func (g_ GlassLozengeWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](g_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (g_ GlassLozengeWrapper) HasInputImage() bool {
	return g_.RespondsToSelector(objc.Sel("inputImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglasslozenge/3600161-inputimage?language=objc
func (g_ GlassLozengeWrapper) InputImage() Image {
	rv := objc.Call[Image](g_, objc.Sel("inputImage"))
	return rv
}

func (g_ GlassLozengeWrapper) HasSetRefraction() bool {
	return g_.RespondsToSelector(objc.Sel("setRefraction:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglasslozenge/3600165-refraction?language=objc
func (g_ GlassLozengeWrapper) SetRefraction(value float64) {
	objc.Call[objc.Void](g_, objc.Sel("setRefraction:"), value)
}

func (g_ GlassLozengeWrapper) HasRefraction() bool {
	return g_.RespondsToSelector(objc.Sel("refraction"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglasslozenge/3600165-refraction?language=objc
func (g_ GlassLozengeWrapper) Refraction() float64 {
	rv := objc.Call[float64](g_, objc.Sel("refraction"))
	return rv
}

func (g_ GlassLozengeWrapper) HasSetRadius() bool {
	return g_.RespondsToSelector(objc.Sel("setRadius:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglasslozenge/3600164-radius?language=objc
func (g_ GlassLozengeWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](g_, objc.Sel("setRadius:"), value)
}

func (g_ GlassLozengeWrapper) HasRadius() bool {
	return g_.RespondsToSelector(objc.Sel("radius"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglasslozenge/3600164-radius?language=objc
func (g_ GlassLozengeWrapper) Radius() float64 {
	rv := objc.Call[float64](g_, objc.Sel("radius"))
	return rv
}

func (g_ GlassLozengeWrapper) HasSetPoint1() bool {
	return g_.RespondsToSelector(objc.Sel("setPoint1:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglasslozenge/3600163-point1?language=objc
func (g_ GlassLozengeWrapper) SetPoint1(value coregraphics.Point) {
	objc.Call[objc.Void](g_, objc.Sel("setPoint1:"), value)
}

func (g_ GlassLozengeWrapper) HasPoint1() bool {
	return g_.RespondsToSelector(objc.Sel("point1"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglasslozenge/3600163-point1?language=objc
func (g_ GlassLozengeWrapper) Point1() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](g_, objc.Sel("point1"))
	return rv
}

func (g_ GlassLozengeWrapper) HasSetPoint0() bool {
	return g_.RespondsToSelector(objc.Sel("setPoint0:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglasslozenge/3600162-point0?language=objc
func (g_ GlassLozengeWrapper) SetPoint0(value coregraphics.Point) {
	objc.Call[objc.Void](g_, objc.Sel("setPoint0:"), value)
}

func (g_ GlassLozengeWrapper) HasPoint0() bool {
	return g_.RespondsToSelector(objc.Sel("point0"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglasslozenge/3600162-point0?language=objc
func (g_ GlassLozengeWrapper) Point0() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](g_, objc.Sel("point0"))
	return rv
}
