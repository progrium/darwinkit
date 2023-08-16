// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a pointillize filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipointillize?language=objc
type PPointillize interface {
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

// A concrete type wrapper for the [PPointillize] protocol.
type PointillizeWrapper struct {
	objc.Object
}

func (p_ PointillizeWrapper) HasSetInputImage() bool {
	return p_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipointillize/3228679-inputimage?language=objc
func (p_ PointillizeWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](p_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (p_ PointillizeWrapper) HasInputImage() bool {
	return p_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipointillize/3228679-inputimage?language=objc
func (p_ PointillizeWrapper) InputImage() Image {
	rv := objc.Call[Image](p_, objc.Sel("inputImage"))
	return rv
}

func (p_ PointillizeWrapper) HasSetRadius() bool {
	return p_.RespondsToSelector(objc.Sel("setRadius:"))
}

// The radius of the circles in the resulting pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipointillize/3228680-radius?language=objc
func (p_ PointillizeWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setRadius:"), value)
}

func (p_ PointillizeWrapper) HasRadius() bool {
	return p_.RespondsToSelector(objc.Sel("radius"))
}

// The radius of the circles in the resulting pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipointillize/3228680-radius?language=objc
func (p_ PointillizeWrapper) Radius() float64 {
	rv := objc.Call[float64](p_, objc.Sel("radius"))
	return rv
}

func (p_ PointillizeWrapper) HasSetCenter() bool {
	return p_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipointillize/3228678-center?language=objc
func (p_ PointillizeWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](p_, objc.Sel("setCenter:"), value)
}

func (p_ PointillizeWrapper) HasCenter() bool {
	return p_.RespondsToSelector(objc.Sel("center"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipointillize/3228678-center?language=objc
func (p_ PointillizeWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](p_, objc.Sel("center"))
	return rv
}
