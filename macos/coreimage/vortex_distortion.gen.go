// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civortexdistortion?language=objc
type PVortexDistortion interface {
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

// A concrete type wrapper for the [PVortexDistortion] protocol.
type VortexDistortionWrapper struct {
	objc.Object
}

func (v_ VortexDistortionWrapper) HasSetInputImage() bool {
	return v_.RespondsToSelector(objc.Sel("setInputImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civortexdistortion/3600210-inputimage?language=objc
func (v_ VortexDistortionWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](v_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (v_ VortexDistortionWrapper) HasInputImage() bool {
	return v_.RespondsToSelector(objc.Sel("inputImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civortexdistortion/3600210-inputimage?language=objc
func (v_ VortexDistortionWrapper) InputImage() Image {
	rv := objc.Call[Image](v_, objc.Sel("inputImage"))
	return rv
}

func (v_ VortexDistortionWrapper) HasSetRadius() bool {
	return v_.RespondsToSelector(objc.Sel("setRadius:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civortexdistortion/3600211-radius?language=objc
func (v_ VortexDistortionWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](v_, objc.Sel("setRadius:"), value)
}

func (v_ VortexDistortionWrapper) HasRadius() bool {
	return v_.RespondsToSelector(objc.Sel("radius"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civortexdistortion/3600211-radius?language=objc
func (v_ VortexDistortionWrapper) Radius() float64 {
	rv := objc.Call[float64](v_, objc.Sel("radius"))
	return rv
}

func (v_ VortexDistortionWrapper) HasSetAngle() bool {
	return v_.RespondsToSelector(objc.Sel("setAngle:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civortexdistortion/3600208-angle?language=objc
func (v_ VortexDistortionWrapper) SetAngle(value float64) {
	objc.Call[objc.Void](v_, objc.Sel("setAngle:"), value)
}

func (v_ VortexDistortionWrapper) HasAngle() bool {
	return v_.RespondsToSelector(objc.Sel("angle"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civortexdistortion/3600208-angle?language=objc
func (v_ VortexDistortionWrapper) Angle() float64 {
	rv := objc.Call[float64](v_, objc.Sel("angle"))
	return rv
}

func (v_ VortexDistortionWrapper) HasSetCenter() bool {
	return v_.RespondsToSelector(objc.Sel("setCenter:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civortexdistortion/3600209-center?language=objc
func (v_ VortexDistortionWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](v_, objc.Sel("setCenter:"), value)
}

func (v_ VortexDistortionWrapper) HasCenter() bool {
	return v_.RespondsToSelector(objc.Sel("center"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civortexdistortion/3600209-center?language=objc
func (v_ VortexDistortionWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](v_, objc.Sel("center"))
	return rv
}
