// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citoruslensdistortion?language=objc
type PTorusLensDistortion interface {
	// optional
	SetWidth(value float64)
	HasSetWidth() bool

	// optional
	Width() float64
	HasWidth() bool

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
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool
}

// A concrete type wrapper for the [PTorusLensDistortion] protocol.
type TorusLensDistortionWrapper struct {
	objc.Object
}

func (t_ TorusLensDistortionWrapper) HasSetWidth() bool {
	return t_.RespondsToSelector(objc.Sel("setWidth:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citoruslensdistortion/3600201-width?language=objc
func (t_ TorusLensDistortionWrapper) SetWidth(value float64) {
	objc.Call[objc.Void](t_, objc.Sel("setWidth:"), value)
}

func (t_ TorusLensDistortionWrapper) HasWidth() bool {
	return t_.RespondsToSelector(objc.Sel("width"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citoruslensdistortion/3600201-width?language=objc
func (t_ TorusLensDistortionWrapper) Width() float64 {
	rv := objc.Call[float64](t_, objc.Sel("width"))
	return rv
}

func (t_ TorusLensDistortionWrapper) HasSetInputImage() bool {
	return t_.RespondsToSelector(objc.Sel("setInputImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citoruslensdistortion/3600198-inputimage?language=objc
func (t_ TorusLensDistortionWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](t_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (t_ TorusLensDistortionWrapper) HasInputImage() bool {
	return t_.RespondsToSelector(objc.Sel("inputImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citoruslensdistortion/3600198-inputimage?language=objc
func (t_ TorusLensDistortionWrapper) InputImage() Image {
	rv := objc.Call[Image](t_, objc.Sel("inputImage"))
	return rv
}

func (t_ TorusLensDistortionWrapper) HasSetRefraction() bool {
	return t_.RespondsToSelector(objc.Sel("setRefraction:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citoruslensdistortion/3600200-refraction?language=objc
func (t_ TorusLensDistortionWrapper) SetRefraction(value float64) {
	objc.Call[objc.Void](t_, objc.Sel("setRefraction:"), value)
}

func (t_ TorusLensDistortionWrapper) HasRefraction() bool {
	return t_.RespondsToSelector(objc.Sel("refraction"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citoruslensdistortion/3600200-refraction?language=objc
func (t_ TorusLensDistortionWrapper) Refraction() float64 {
	rv := objc.Call[float64](t_, objc.Sel("refraction"))
	return rv
}

func (t_ TorusLensDistortionWrapper) HasSetRadius() bool {
	return t_.RespondsToSelector(objc.Sel("setRadius:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citoruslensdistortion/3600199-radius?language=objc
func (t_ TorusLensDistortionWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](t_, objc.Sel("setRadius:"), value)
}

func (t_ TorusLensDistortionWrapper) HasRadius() bool {
	return t_.RespondsToSelector(objc.Sel("radius"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citoruslensdistortion/3600199-radius?language=objc
func (t_ TorusLensDistortionWrapper) Radius() float64 {
	rv := objc.Call[float64](t_, objc.Sel("radius"))
	return rv
}

func (t_ TorusLensDistortionWrapper) HasSetCenter() bool {
	return t_.RespondsToSelector(objc.Sel("setCenter:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citoruslensdistortion/3600197-center?language=objc
func (t_ TorusLensDistortionWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](t_, objc.Sel("setCenter:"), value)
}

func (t_ TorusLensDistortionWrapper) HasCenter() bool {
	return t_.RespondsToSelector(objc.Sel("center"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citoruslensdistortion/3600197-center?language=objc
func (t_ TorusLensDistortionWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](t_, objc.Sel("center"))
	return rv
}
