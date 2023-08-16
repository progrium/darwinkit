// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a kaleidoscope filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikaleidoscope?language=objc
type PKaleidoscope interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetCount(value int)
	HasSetCount() bool

	// optional
	Count() int
	HasCount() bool

	// optional
	SetAngle(value float64)
	HasSetAngle() bool

	// optional
	Angle() float64
	HasAngle() bool

	// optional
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool
}

// A concrete type wrapper for the [PKaleidoscope] protocol.
type KaleidoscopeWrapper struct {
	objc.Object
}

func (k_ KaleidoscopeWrapper) HasSetInputImage() bool {
	return k_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikaleidoscope/3228511-inputimage?language=objc
func (k_ KaleidoscopeWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](k_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (k_ KaleidoscopeWrapper) HasInputImage() bool {
	return k_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikaleidoscope/3228511-inputimage?language=objc
func (k_ KaleidoscopeWrapper) InputImage() Image {
	rv := objc.Call[Image](k_, objc.Sel("inputImage"))
	return rv
}

func (k_ KaleidoscopeWrapper) HasSetCount() bool {
	return k_.RespondsToSelector(objc.Sel("setCount:"))
}

// The number of reflections in the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikaleidoscope/3228510-count?language=objc
func (k_ KaleidoscopeWrapper) SetCount(value int) {
	objc.Call[objc.Void](k_, objc.Sel("setCount:"), value)
}

func (k_ KaleidoscopeWrapper) HasCount() bool {
	return k_.RespondsToSelector(objc.Sel("count"))
}

// The number of reflections in the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikaleidoscope/3228510-count?language=objc
func (k_ KaleidoscopeWrapper) Count() int {
	rv := objc.Call[int](k_, objc.Sel("count"))
	return rv
}

func (k_ KaleidoscopeWrapper) HasSetAngle() bool {
	return k_.RespondsToSelector(objc.Sel("setAngle:"))
}

// The angle of the reflection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikaleidoscope/3228508-angle?language=objc
func (k_ KaleidoscopeWrapper) SetAngle(value float64) {
	objc.Call[objc.Void](k_, objc.Sel("setAngle:"), value)
}

func (k_ KaleidoscopeWrapper) HasAngle() bool {
	return k_.RespondsToSelector(objc.Sel("angle"))
}

// The angle of the reflection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikaleidoscope/3228508-angle?language=objc
func (k_ KaleidoscopeWrapper) Angle() float64 {
	rv := objc.Call[float64](k_, objc.Sel("angle"))
	return rv
}

func (k_ KaleidoscopeWrapper) HasSetCenter() bool {
	return k_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikaleidoscope/3228509-center?language=objc
func (k_ KaleidoscopeWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](k_, objc.Sel("setCenter:"), value)
}

func (k_ KaleidoscopeWrapper) HasCenter() bool {
	return k_.RespondsToSelector(objc.Sel("center"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikaleidoscope/3228509-center?language=objc
func (k_ KaleidoscopeWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](k_, objc.Sel("center"))
	return rv
}
