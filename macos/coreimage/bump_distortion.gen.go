// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibumpdistortion?language=objc
type PBumpDistortion interface {
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

// A concrete type wrapper for the [PBumpDistortion] protocol.
type BumpDistortionWrapper struct {
	objc.Object
}

func (b_ BumpDistortionWrapper) HasSetScale() bool {
	return b_.RespondsToSelector(objc.Sel("setScale:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibumpdistortion/3600110-scale?language=objc
func (b_ BumpDistortionWrapper) SetScale(value float64) {
	objc.Call[objc.Void](b_, objc.Sel("setScale:"), value)
}

func (b_ BumpDistortionWrapper) HasScale() bool {
	return b_.RespondsToSelector(objc.Sel("scale"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibumpdistortion/3600110-scale?language=objc
func (b_ BumpDistortionWrapper) Scale() float64 {
	rv := objc.Call[float64](b_, objc.Sel("scale"))
	return rv
}

func (b_ BumpDistortionWrapper) HasSetInputImage() bool {
	return b_.RespondsToSelector(objc.Sel("setInputImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibumpdistortion/3600108-inputimage?language=objc
func (b_ BumpDistortionWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](b_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (b_ BumpDistortionWrapper) HasInputImage() bool {
	return b_.RespondsToSelector(objc.Sel("inputImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibumpdistortion/3600108-inputimage?language=objc
func (b_ BumpDistortionWrapper) InputImage() Image {
	rv := objc.Call[Image](b_, objc.Sel("inputImage"))
	return rv
}

func (b_ BumpDistortionWrapper) HasSetRadius() bool {
	return b_.RespondsToSelector(objc.Sel("setRadius:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibumpdistortion/3600109-radius?language=objc
func (b_ BumpDistortionWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](b_, objc.Sel("setRadius:"), value)
}

func (b_ BumpDistortionWrapper) HasRadius() bool {
	return b_.RespondsToSelector(objc.Sel("radius"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibumpdistortion/3600109-radius?language=objc
func (b_ BumpDistortionWrapper) Radius() float64 {
	rv := objc.Call[float64](b_, objc.Sel("radius"))
	return rv
}

func (b_ BumpDistortionWrapper) HasSetCenter() bool {
	return b_.RespondsToSelector(objc.Sel("setCenter:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibumpdistortion/3600107-center?language=objc
func (b_ BumpDistortionWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](b_, objc.Sel("setCenter:"), value)
}

func (b_ BumpDistortionWrapper) HasCenter() bool {
	return b_.RespondsToSelector(objc.Sel("center"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibumpdistortion/3600107-center?language=objc
func (b_ BumpDistortionWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](b_, objc.Sel("center"))
	return rv
}
