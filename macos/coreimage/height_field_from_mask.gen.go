// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a height-field-from-mask filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciheightfieldfrommask?language=objc
type PHeightFieldFromMask interface {
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

// A concrete type wrapper for the [PHeightFieldFromMask] protocol.
type HeightFieldFromMaskWrapper struct {
	objc.Object
}

func (h_ HeightFieldFromMaskWrapper) HasSetInputImage() bool {
	return h_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciheightfieldfrommask/3228487-inputimage?language=objc
func (h_ HeightFieldFromMaskWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](h_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (h_ HeightFieldFromMaskWrapper) HasInputImage() bool {
	return h_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciheightfieldfrommask/3228487-inputimage?language=objc
func (h_ HeightFieldFromMaskWrapper) InputImage() Image {
	rv := objc.Call[Image](h_, objc.Sel("inputImage"))
	return rv
}

func (h_ HeightFieldFromMaskWrapper) HasSetRadius() bool {
	return h_.RespondsToSelector(objc.Sel("setRadius:"))
}

// The length of the height-field transition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciheightfieldfrommask/3228488-radius?language=objc
func (h_ HeightFieldFromMaskWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](h_, objc.Sel("setRadius:"), value)
}

func (h_ HeightFieldFromMaskWrapper) HasRadius() bool {
	return h_.RespondsToSelector(objc.Sel("radius"))
}

// The length of the height-field transition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciheightfieldfrommask/3228488-radius?language=objc
func (h_ HeightFieldFromMaskWrapper) Radius() float64 {
	rv := objc.Call[float64](h_, objc.Sel("radius"))
	return rv
}
