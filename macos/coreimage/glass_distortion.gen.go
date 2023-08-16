// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglassdistortion?language=objc
type PGlassDistortion interface {
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
	SetTextureImage(value Image)
	HasSetTextureImage() bool

	// optional
	TextureImage() IImage
	HasTextureImage() bool

	// optional
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool
}

// A concrete type wrapper for the [PGlassDistortion] protocol.
type GlassDistortionWrapper struct {
	objc.Object
}

func (g_ GlassDistortionWrapper) HasSetScale() bool {
	return g_.RespondsToSelector(objc.Sel("setScale:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglassdistortion/3600158-scale?language=objc
func (g_ GlassDistortionWrapper) SetScale(value float64) {
	objc.Call[objc.Void](g_, objc.Sel("setScale:"), value)
}

func (g_ GlassDistortionWrapper) HasScale() bool {
	return g_.RespondsToSelector(objc.Sel("scale"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglassdistortion/3600158-scale?language=objc
func (g_ GlassDistortionWrapper) Scale() float64 {
	rv := objc.Call[float64](g_, objc.Sel("scale"))
	return rv
}

func (g_ GlassDistortionWrapper) HasSetInputImage() bool {
	return g_.RespondsToSelector(objc.Sel("setInputImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglassdistortion/3600157-inputimage?language=objc
func (g_ GlassDistortionWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](g_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (g_ GlassDistortionWrapper) HasInputImage() bool {
	return g_.RespondsToSelector(objc.Sel("inputImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglassdistortion/3600157-inputimage?language=objc
func (g_ GlassDistortionWrapper) InputImage() Image {
	rv := objc.Call[Image](g_, objc.Sel("inputImage"))
	return rv
}

func (g_ GlassDistortionWrapper) HasSetTextureImage() bool {
	return g_.RespondsToSelector(objc.Sel("setTextureImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglassdistortion/3600159-textureimage?language=objc
func (g_ GlassDistortionWrapper) SetTextureImage(value IImage) {
	objc.Call[objc.Void](g_, objc.Sel("setTextureImage:"), objc.Ptr(value))
}

func (g_ GlassDistortionWrapper) HasTextureImage() bool {
	return g_.RespondsToSelector(objc.Sel("textureImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglassdistortion/3600159-textureimage?language=objc
func (g_ GlassDistortionWrapper) TextureImage() Image {
	rv := objc.Call[Image](g_, objc.Sel("textureImage"))
	return rv
}

func (g_ GlassDistortionWrapper) HasSetCenter() bool {
	return g_.RespondsToSelector(objc.Sel("setCenter:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglassdistortion/3600156-center?language=objc
func (g_ GlassDistortionWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](g_, objc.Sel("setCenter:"), value)
}

func (g_ GlassDistortionWrapper) HasCenter() bool {
	return g_.RespondsToSelector(objc.Sel("center"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglassdistortion/3600156-center?language=objc
func (g_ GlassDistortionWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](g_, objc.Sel("center"))
	return rv
}
