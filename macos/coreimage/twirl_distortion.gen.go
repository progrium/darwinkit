// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citwirldistortion?language=objc
type PTwirlDistortion interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetRadius(value float64)
	HasSetRadius() bool

	// optional
	Radius() float64
	HasRadius() bool

	// optional
	SetAngle(value float64)
	HasSetAngle() bool

	// optional
	Angle() float64
	HasAngle() bool

	// optional
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool
}

// A concrete type wrapper for the [PTwirlDistortion] protocol.
type TwirlDistortionWrapper struct {
	objc.Object
}

func (t_ TwirlDistortionWrapper) HasSetInputImage() bool {
	return t_.RespondsToSelector(objc.Sel("setInputImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citwirldistortion/3600205-inputimage?language=objc
func (t_ TwirlDistortionWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](t_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (t_ TwirlDistortionWrapper) HasInputImage() bool {
	return t_.RespondsToSelector(objc.Sel("inputImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citwirldistortion/3600205-inputimage?language=objc
func (t_ TwirlDistortionWrapper) InputImage() Image {
	rv := objc.Call[Image](t_, objc.Sel("inputImage"))
	return rv
}

func (t_ TwirlDistortionWrapper) HasSetRadius() bool {
	return t_.RespondsToSelector(objc.Sel("setRadius:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citwirldistortion/3600206-radius?language=objc
func (t_ TwirlDistortionWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](t_, objc.Sel("setRadius:"), value)
}

func (t_ TwirlDistortionWrapper) HasRadius() bool {
	return t_.RespondsToSelector(objc.Sel("radius"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citwirldistortion/3600206-radius?language=objc
func (t_ TwirlDistortionWrapper) Radius() float64 {
	rv := objc.Call[float64](t_, objc.Sel("radius"))
	return rv
}

func (t_ TwirlDistortionWrapper) HasSetAngle() bool {
	return t_.RespondsToSelector(objc.Sel("setAngle:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citwirldistortion/3600203-angle?language=objc
func (t_ TwirlDistortionWrapper) SetAngle(value float64) {
	objc.Call[objc.Void](t_, objc.Sel("setAngle:"), value)
}

func (t_ TwirlDistortionWrapper) HasAngle() bool {
	return t_.RespondsToSelector(objc.Sel("angle"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citwirldistortion/3600203-angle?language=objc
func (t_ TwirlDistortionWrapper) Angle() float64 {
	rv := objc.Call[float64](t_, objc.Sel("angle"))
	return rv
}

func (t_ TwirlDistortionWrapper) HasSetCenter() bool {
	return t_.RespondsToSelector(objc.Sel("setCenter:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citwirldistortion/3600204-center?language=objc
func (t_ TwirlDistortionWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](t_, objc.Sel("setCenter:"), value)
}

func (t_ TwirlDistortionWrapper) HasCenter() bool {
	return t_.RespondsToSelector(objc.Sel("center"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citwirldistortion/3600204-center?language=objc
func (t_ TwirlDistortionWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](t_, objc.Sel("center"))
	return rv
}
