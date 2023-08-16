// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciholedistortion?language=objc
type PHoleDistortion interface {
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

// A concrete type wrapper for the [PHoleDistortion] protocol.
type HoleDistortionWrapper struct {
	objc.Object
}

func (h_ HoleDistortionWrapper) HasSetInputImage() bool {
	return h_.RespondsToSelector(objc.Sel("setInputImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciholedistortion/3600168-inputimage?language=objc
func (h_ HoleDistortionWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](h_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (h_ HoleDistortionWrapper) HasInputImage() bool {
	return h_.RespondsToSelector(objc.Sel("inputImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciholedistortion/3600168-inputimage?language=objc
func (h_ HoleDistortionWrapper) InputImage() Image {
	rv := objc.Call[Image](h_, objc.Sel("inputImage"))
	return rv
}

func (h_ HoleDistortionWrapper) HasSetRadius() bool {
	return h_.RespondsToSelector(objc.Sel("setRadius:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciholedistortion/3600169-radius?language=objc
func (h_ HoleDistortionWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](h_, objc.Sel("setRadius:"), value)
}

func (h_ HoleDistortionWrapper) HasRadius() bool {
	return h_.RespondsToSelector(objc.Sel("radius"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciholedistortion/3600169-radius?language=objc
func (h_ HoleDistortionWrapper) Radius() float64 {
	rv := objc.Call[float64](h_, objc.Sel("radius"))
	return rv
}

func (h_ HoleDistortionWrapper) HasSetCenter() bool {
	return h_.RespondsToSelector(objc.Sel("setCenter:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciholedistortion/3600167-center?language=objc
func (h_ HoleDistortionWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](h_, objc.Sel("setCenter:"), value)
}

func (h_ HoleDistortionWrapper) HasCenter() bool {
	return h_.RespondsToSelector(objc.Sel("center"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciholedistortion/3600167-center?language=objc
func (h_ HoleDistortionWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](h_, objc.Sel("center"))
	return rv
}
