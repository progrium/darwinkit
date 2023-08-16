// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a crystalize filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicrystallize?language=objc
type PCrystallize interface {
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

// A concrete type wrapper for the [PCrystallize] protocol.
type CrystallizeWrapper struct {
	objc.Object
}

func (c_ CrystallizeWrapper) HasSetInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicrystallize/3228201-inputimage?language=objc
func (c_ CrystallizeWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](c_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (c_ CrystallizeWrapper) HasInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicrystallize/3228201-inputimage?language=objc
func (c_ CrystallizeWrapper) InputImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("inputImage"))
	return rv
}

func (c_ CrystallizeWrapper) HasSetRadius() bool {
	return c_.RespondsToSelector(objc.Sel("setRadius:"))
}

// The radius, in pixels, of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicrystallize/3228202-radius?language=objc
func (c_ CrystallizeWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setRadius:"), value)
}

func (c_ CrystallizeWrapper) HasRadius() bool {
	return c_.RespondsToSelector(objc.Sel("radius"))
}

// The radius, in pixels, of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicrystallize/3228202-radius?language=objc
func (c_ CrystallizeWrapper) Radius() float64 {
	rv := objc.Call[float64](c_, objc.Sel("radius"))
	return rv
}

func (c_ CrystallizeWrapper) HasSetCenter() bool {
	return c_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The center of the effect as x and y coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicrystallize/3228200-center?language=objc
func (c_ CrystallizeWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](c_, objc.Sel("setCenter:"), value)
}

func (c_ CrystallizeWrapper) HasCenter() bool {
	return c_.RespondsToSelector(objc.Sel("center"))
}

// The center of the effect as x and y coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicrystallize/3228200-center?language=objc
func (c_ CrystallizeWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](c_, objc.Sel("center"))
	return rv
}
