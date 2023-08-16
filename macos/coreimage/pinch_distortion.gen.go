// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipinchdistortion?language=objc
type PPinchDistortion interface {
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
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool
}

// A concrete type wrapper for the [PPinchDistortion] protocol.
type PinchDistortionWrapper struct {
	objc.Object
}

func (p_ PinchDistortionWrapper) HasSetScale() bool {
	return p_.RespondsToSelector(objc.Sel("setScale:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipinchdistortion/3600190-scale?language=objc
func (p_ PinchDistortionWrapper) SetScale(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setScale:"), value)
}

func (p_ PinchDistortionWrapper) HasScale() bool {
	return p_.RespondsToSelector(objc.Sel("scale"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipinchdistortion/3600190-scale?language=objc
func (p_ PinchDistortionWrapper) Scale() float64 {
	rv := objc.Call[float64](p_, objc.Sel("scale"))
	return rv
}

func (p_ PinchDistortionWrapper) HasSetInputImage() bool {
	return p_.RespondsToSelector(objc.Sel("setInputImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipinchdistortion/3600188-inputimage?language=objc
func (p_ PinchDistortionWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](p_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (p_ PinchDistortionWrapper) HasInputImage() bool {
	return p_.RespondsToSelector(objc.Sel("inputImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipinchdistortion/3600188-inputimage?language=objc
func (p_ PinchDistortionWrapper) InputImage() Image {
	rv := objc.Call[Image](p_, objc.Sel("inputImage"))
	return rv
}

func (p_ PinchDistortionWrapper) HasSetRadius() bool {
	return p_.RespondsToSelector(objc.Sel("setRadius:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipinchdistortion/3600189-radius?language=objc
func (p_ PinchDistortionWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setRadius:"), value)
}

func (p_ PinchDistortionWrapper) HasRadius() bool {
	return p_.RespondsToSelector(objc.Sel("radius"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipinchdistortion/3600189-radius?language=objc
func (p_ PinchDistortionWrapper) Radius() float64 {
	rv := objc.Call[float64](p_, objc.Sel("radius"))
	return rv
}

func (p_ PinchDistortionWrapper) HasSetCenter() bool {
	return p_.RespondsToSelector(objc.Sel("setCenter:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipinchdistortion/3600187-center?language=objc
func (p_ PinchDistortionWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](p_, objc.Sel("setCenter:"), value)
}

func (p_ PinchDistortionWrapper) HasCenter() bool {
	return p_.RespondsToSelector(objc.Sel("center"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipinchdistortion/3600187-center?language=objc
func (p_ PinchDistortionWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](p_, objc.Sel("center"))
	return rv
}
