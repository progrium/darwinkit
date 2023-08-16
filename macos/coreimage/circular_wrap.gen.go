// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicircularwrap?language=objc
type PCircularWrap interface {
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

// A concrete type wrapper for the [PCircularWrap] protocol.
type CircularWrapWrapper struct {
	objc.Object
}

func (c_ CircularWrapWrapper) HasSetInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("setInputImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicircularwrap/3600124-inputimage?language=objc
func (c_ CircularWrapWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](c_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (c_ CircularWrapWrapper) HasInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("inputImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicircularwrap/3600124-inputimage?language=objc
func (c_ CircularWrapWrapper) InputImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("inputImage"))
	return rv
}

func (c_ CircularWrapWrapper) HasSetRadius() bool {
	return c_.RespondsToSelector(objc.Sel("setRadius:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicircularwrap/3600125-radius?language=objc
func (c_ CircularWrapWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setRadius:"), value)
}

func (c_ CircularWrapWrapper) HasRadius() bool {
	return c_.RespondsToSelector(objc.Sel("radius"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicircularwrap/3600125-radius?language=objc
func (c_ CircularWrapWrapper) Radius() float64 {
	rv := objc.Call[float64](c_, objc.Sel("radius"))
	return rv
}

func (c_ CircularWrapWrapper) HasSetAngle() bool {
	return c_.RespondsToSelector(objc.Sel("setAngle:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicircularwrap/3600122-angle?language=objc
func (c_ CircularWrapWrapper) SetAngle(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setAngle:"), value)
}

func (c_ CircularWrapWrapper) HasAngle() bool {
	return c_.RespondsToSelector(objc.Sel("angle"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicircularwrap/3600122-angle?language=objc
func (c_ CircularWrapWrapper) Angle() float64 {
	rv := objc.Call[float64](c_, objc.Sel("angle"))
	return rv
}

func (c_ CircularWrapWrapper) HasSetCenter() bool {
	return c_.RespondsToSelector(objc.Sel("setCenter:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicircularwrap/3600123-center?language=objc
func (c_ CircularWrapWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](c_, objc.Sel("setCenter:"), value)
}

func (c_ CircularWrapWrapper) HasCenter() bool {
	return c_.RespondsToSelector(objc.Sel("center"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicircularwrap/3600123-center?language=objc
func (c_ CircularWrapWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](c_, objc.Sel("center"))
	return rv
}
