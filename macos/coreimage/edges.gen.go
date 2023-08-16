// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure an edges filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciedges?language=objc
type PEdges interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetIntensity(value float64)
	HasSetIntensity() bool

	// optional
	Intensity() float64
	HasIntensity() bool
}

// A concrete type wrapper for the [PEdges] protocol.
type EdgesWrapper struct {
	objc.Object
}

func (e_ EdgesWrapper) HasSetInputImage() bool {
	return e_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciedges/3228245-inputimage?language=objc
func (e_ EdgesWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](e_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (e_ EdgesWrapper) HasInputImage() bool {
	return e_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciedges/3228245-inputimage?language=objc
func (e_ EdgesWrapper) InputImage() Image {
	rv := objc.Call[Image](e_, objc.Sel("inputImage"))
	return rv
}

func (e_ EdgesWrapper) HasSetIntensity() bool {
	return e_.RespondsToSelector(objc.Sel("setIntensity:"))
}

// The intensity of the edges. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciedges/3228246-intensity?language=objc
func (e_ EdgesWrapper) SetIntensity(value float64) {
	objc.Call[objc.Void](e_, objc.Sel("setIntensity:"), value)
}

func (e_ EdgesWrapper) HasIntensity() bool {
	return e_.RespondsToSelector(objc.Sel("intensity"))
}

// The intensity of the edges. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciedges/3228246-intensity?language=objc
func (e_ EdgesWrapper) Intensity() float64 {
	rv := objc.Call[float64](e_, objc.Sel("intensity"))
	return rv
}
