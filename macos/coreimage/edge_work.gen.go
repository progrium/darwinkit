// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure an edge-work filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciedgework?language=objc
type PEdgeWork interface {
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
}

// A concrete type wrapper for the [PEdgeWork] protocol.
type EdgeWorkWrapper struct {
	objc.Object
}

func (e_ EdgeWorkWrapper) HasSetInputImage() bool {
	return e_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciedgework/3228242-inputimage?language=objc
func (e_ EdgeWorkWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](e_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (e_ EdgeWorkWrapper) HasInputImage() bool {
	return e_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciedgework/3228242-inputimage?language=objc
func (e_ EdgeWorkWrapper) InputImage() Image {
	rv := objc.Call[Image](e_, objc.Sel("inputImage"))
	return rv
}

func (e_ EdgeWorkWrapper) HasSetRadius() bool {
	return e_.RespondsToSelector(objc.Sel("setRadius:"))
}

// The thickness of the edges. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciedgework/3228243-radius?language=objc
func (e_ EdgeWorkWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](e_, objc.Sel("setRadius:"), value)
}

func (e_ EdgeWorkWrapper) HasRadius() bool {
	return e_.RespondsToSelector(objc.Sel("radius"))
}

// The thickness of the edges. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciedgework/3228243-radius?language=objc
func (e_ EdgeWorkWrapper) Radius() float64 {
	rv := objc.Call[float64](e_, objc.Sel("radius"))
	return rv
}
