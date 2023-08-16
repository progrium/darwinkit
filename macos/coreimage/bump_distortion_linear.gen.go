// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibumpdistortionlinear?language=objc
type PBumpDistortionLinear interface {
	// optional
	SetScale(value float64)
	HasSetScale() bool

	// optional
	Scale() float64
	HasScale() bool

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

// A concrete type wrapper for the [PBumpDistortionLinear] protocol.
type BumpDistortionLinearWrapper struct {
	objc.Object
}

func (b_ BumpDistortionLinearWrapper) HasSetScale() bool {
	return b_.RespondsToSelector(objc.Sel("setScale:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibumpdistortionlinear/3600116-scale?language=objc
func (b_ BumpDistortionLinearWrapper) SetScale(value float64) {
	objc.Call[objc.Void](b_, objc.Sel("setScale:"), value)
}

func (b_ BumpDistortionLinearWrapper) HasScale() bool {
	return b_.RespondsToSelector(objc.Sel("scale"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibumpdistortionlinear/3600116-scale?language=objc
func (b_ BumpDistortionLinearWrapper) Scale() float64 {
	rv := objc.Call[float64](b_, objc.Sel("scale"))
	return rv
}

func (b_ BumpDistortionLinearWrapper) HasSetInputImage() bool {
	return b_.RespondsToSelector(objc.Sel("setInputImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibumpdistortionlinear/3600114-inputimage?language=objc
func (b_ BumpDistortionLinearWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](b_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (b_ BumpDistortionLinearWrapper) HasInputImage() bool {
	return b_.RespondsToSelector(objc.Sel("inputImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibumpdistortionlinear/3600114-inputimage?language=objc
func (b_ BumpDistortionLinearWrapper) InputImage() Image {
	rv := objc.Call[Image](b_, objc.Sel("inputImage"))
	return rv
}

func (b_ BumpDistortionLinearWrapper) HasSetRadius() bool {
	return b_.RespondsToSelector(objc.Sel("setRadius:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibumpdistortionlinear/3600115-radius?language=objc
func (b_ BumpDistortionLinearWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](b_, objc.Sel("setRadius:"), value)
}

func (b_ BumpDistortionLinearWrapper) HasRadius() bool {
	return b_.RespondsToSelector(objc.Sel("radius"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibumpdistortionlinear/3600115-radius?language=objc
func (b_ BumpDistortionLinearWrapper) Radius() float64 {
	rv := objc.Call[float64](b_, objc.Sel("radius"))
	return rv
}

func (b_ BumpDistortionLinearWrapper) HasSetAngle() bool {
	return b_.RespondsToSelector(objc.Sel("setAngle:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibumpdistortionlinear/3600112-angle?language=objc
func (b_ BumpDistortionLinearWrapper) SetAngle(value float64) {
	objc.Call[objc.Void](b_, objc.Sel("setAngle:"), value)
}

func (b_ BumpDistortionLinearWrapper) HasAngle() bool {
	return b_.RespondsToSelector(objc.Sel("angle"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibumpdistortionlinear/3600112-angle?language=objc
func (b_ BumpDistortionLinearWrapper) Angle() float64 {
	rv := objc.Call[float64](b_, objc.Sel("angle"))
	return rv
}

func (b_ BumpDistortionLinearWrapper) HasSetCenter() bool {
	return b_.RespondsToSelector(objc.Sel("setCenter:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibumpdistortionlinear/3600113-center?language=objc
func (b_ BumpDistortionLinearWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](b_, objc.Sel("setCenter:"), value)
}

func (b_ BumpDistortionLinearWrapper) HasCenter() bool {
	return b_.RespondsToSelector(objc.Sel("center"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibumpdistortionlinear/3600113-center?language=objc
func (b_ BumpDistortionLinearWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](b_, objc.Sel("center"))
	return rv
}
