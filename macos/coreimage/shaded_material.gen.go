// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a shaded material filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cishadedmaterial?language=objc
type PShadedMaterial interface {
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
	SetShadingImage(value Image)
	HasSetShadingImage() bool

	// optional
	ShadingImage() IImage
	HasShadingImage() bool
}

// A concrete type wrapper for the [PShadedMaterial] protocol.
type ShadedMaterialWrapper struct {
	objc.Object
}

func (s_ ShadedMaterialWrapper) HasSetScale() bool {
	return s_.RespondsToSelector(objc.Sel("setScale:"))
}

// The scale of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cishadedmaterial/3228706-scale?language=objc
func (s_ ShadedMaterialWrapper) SetScale(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setScale:"), value)
}

func (s_ ShadedMaterialWrapper) HasScale() bool {
	return s_.RespondsToSelector(objc.Sel("scale"))
}

// The scale of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cishadedmaterial/3228706-scale?language=objc
func (s_ ShadedMaterialWrapper) Scale() float64 {
	rv := objc.Call[float64](s_, objc.Sel("scale"))
	return rv
}

func (s_ ShadedMaterialWrapper) HasSetInputImage() bool {
	return s_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cishadedmaterial/3228705-inputimage?language=objc
func (s_ ShadedMaterialWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](s_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (s_ ShadedMaterialWrapper) HasInputImage() bool {
	return s_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cishadedmaterial/3228705-inputimage?language=objc
func (s_ ShadedMaterialWrapper) InputImage() Image {
	rv := objc.Call[Image](s_, objc.Sel("inputImage"))
	return rv
}

func (s_ ShadedMaterialWrapper) HasSetShadingImage() bool {
	return s_.RespondsToSelector(objc.Sel("setShadingImage:"))
}

// The image to use as the height field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cishadedmaterial/3228707-shadingimage?language=objc
func (s_ ShadedMaterialWrapper) SetShadingImage(value IImage) {
	objc.Call[objc.Void](s_, objc.Sel("setShadingImage:"), objc.Ptr(value))
}

func (s_ ShadedMaterialWrapper) HasShadingImage() bool {
	return s_.RespondsToSelector(objc.Sel("shadingImage"))
}

// The image to use as the height field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cishadedmaterial/3228707-shadingimage?language=objc
func (s_ ShadedMaterialWrapper) ShadingImage() Image {
	rv := objc.Call[Image](s_, objc.Sel("shadingImage"))
	return rv
}
