// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicirclesplashdistortion?language=objc
type PCircleSplashDistortion interface {
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

// A concrete type wrapper for the [PCircleSplashDistortion] protocol.
type CircleSplashDistortionWrapper struct {
	objc.Object
}

func (c_ CircleSplashDistortionWrapper) HasSetInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("setInputImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicirclesplashdistortion/3600119-inputimage?language=objc
func (c_ CircleSplashDistortionWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](c_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (c_ CircleSplashDistortionWrapper) HasInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("inputImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicirclesplashdistortion/3600119-inputimage?language=objc
func (c_ CircleSplashDistortionWrapper) InputImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("inputImage"))
	return rv
}

func (c_ CircleSplashDistortionWrapper) HasSetRadius() bool {
	return c_.RespondsToSelector(objc.Sel("setRadius:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicirclesplashdistortion/3600120-radius?language=objc
func (c_ CircleSplashDistortionWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setRadius:"), value)
}

func (c_ CircleSplashDistortionWrapper) HasRadius() bool {
	return c_.RespondsToSelector(objc.Sel("radius"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicirclesplashdistortion/3600120-radius?language=objc
func (c_ CircleSplashDistortionWrapper) Radius() float64 {
	rv := objc.Call[float64](c_, objc.Sel("radius"))
	return rv
}

func (c_ CircleSplashDistortionWrapper) HasSetCenter() bool {
	return c_.RespondsToSelector(objc.Sel("setCenter:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicirclesplashdistortion/3600118-center?language=objc
func (c_ CircleSplashDistortionWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](c_, objc.Sel("setCenter:"), value)
}

func (c_ CircleSplashDistortionWrapper) HasCenter() bool {
	return c_.RespondsToSelector(objc.Sel("center"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicirclesplashdistortion/3600118-center?language=objc
func (c_ CircleSplashDistortionWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](c_, objc.Sel("center"))
	return rv
}
