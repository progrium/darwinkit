// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidroste?language=objc
type PDroste interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetPeriodicity(value float64)
	HasSetPeriodicity() bool

	// optional
	Periodicity() float64
	HasPeriodicity() bool

	// optional
	SetRotation(value float64)
	HasSetRotation() bool

	// optional
	Rotation() float64
	HasRotation() bool

	// optional
	SetInsetPoint0(value coregraphics.Point)
	HasSetInsetPoint0() bool

	// optional
	InsetPoint0() coregraphics.Point
	HasInsetPoint0() bool

	// optional
	SetZoom(value float64)
	HasSetZoom() bool

	// optional
	Zoom() float64
	HasZoom() bool

	// optional
	SetStrands(value float64)
	HasSetStrands() bool

	// optional
	Strands() float64
	HasStrands() bool

	// optional
	SetInsetPoint1(value coregraphics.Point)
	HasSetInsetPoint1() bool

	// optional
	InsetPoint1() coregraphics.Point
	HasInsetPoint1() bool
}

// A concrete type wrapper for the [PDroste] protocol.
type DrosteWrapper struct {
	objc.Object
}

func (d_ DrosteWrapper) HasSetInputImage() bool {
	return d_.RespondsToSelector(objc.Sel("setInputImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidroste/3600131-inputimage?language=objc
func (d_ DrosteWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](d_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (d_ DrosteWrapper) HasInputImage() bool {
	return d_.RespondsToSelector(objc.Sel("inputImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidroste/3600131-inputimage?language=objc
func (d_ DrosteWrapper) InputImage() Image {
	rv := objc.Call[Image](d_, objc.Sel("inputImage"))
	return rv
}

func (d_ DrosteWrapper) HasSetPeriodicity() bool {
	return d_.RespondsToSelector(objc.Sel("setPeriodicity:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidroste/3600134-periodicity?language=objc
func (d_ DrosteWrapper) SetPeriodicity(value float64) {
	objc.Call[objc.Void](d_, objc.Sel("setPeriodicity:"), value)
}

func (d_ DrosteWrapper) HasPeriodicity() bool {
	return d_.RespondsToSelector(objc.Sel("periodicity"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidroste/3600134-periodicity?language=objc
func (d_ DrosteWrapper) Periodicity() float64 {
	rv := objc.Call[float64](d_, objc.Sel("periodicity"))
	return rv
}

func (d_ DrosteWrapper) HasSetRotation() bool {
	return d_.RespondsToSelector(objc.Sel("setRotation:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidroste/3600135-rotation?language=objc
func (d_ DrosteWrapper) SetRotation(value float64) {
	objc.Call[objc.Void](d_, objc.Sel("setRotation:"), value)
}

func (d_ DrosteWrapper) HasRotation() bool {
	return d_.RespondsToSelector(objc.Sel("rotation"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidroste/3600135-rotation?language=objc
func (d_ DrosteWrapper) Rotation() float64 {
	rv := objc.Call[float64](d_, objc.Sel("rotation"))
	return rv
}

func (d_ DrosteWrapper) HasSetInsetPoint0() bool {
	return d_.RespondsToSelector(objc.Sel("setInsetPoint0:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidroste/3600132-insetpoint0?language=objc
func (d_ DrosteWrapper) SetInsetPoint0(value coregraphics.Point) {
	objc.Call[objc.Void](d_, objc.Sel("setInsetPoint0:"), value)
}

func (d_ DrosteWrapper) HasInsetPoint0() bool {
	return d_.RespondsToSelector(objc.Sel("insetPoint0"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidroste/3600132-insetpoint0?language=objc
func (d_ DrosteWrapper) InsetPoint0() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](d_, objc.Sel("insetPoint0"))
	return rv
}

func (d_ DrosteWrapper) HasSetZoom() bool {
	return d_.RespondsToSelector(objc.Sel("setZoom:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidroste/3600137-zoom?language=objc
func (d_ DrosteWrapper) SetZoom(value float64) {
	objc.Call[objc.Void](d_, objc.Sel("setZoom:"), value)
}

func (d_ DrosteWrapper) HasZoom() bool {
	return d_.RespondsToSelector(objc.Sel("zoom"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidroste/3600137-zoom?language=objc
func (d_ DrosteWrapper) Zoom() float64 {
	rv := objc.Call[float64](d_, objc.Sel("zoom"))
	return rv
}

func (d_ DrosteWrapper) HasSetStrands() bool {
	return d_.RespondsToSelector(objc.Sel("setStrands:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidroste/3600136-strands?language=objc
func (d_ DrosteWrapper) SetStrands(value float64) {
	objc.Call[objc.Void](d_, objc.Sel("setStrands:"), value)
}

func (d_ DrosteWrapper) HasStrands() bool {
	return d_.RespondsToSelector(objc.Sel("strands"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidroste/3600136-strands?language=objc
func (d_ DrosteWrapper) Strands() float64 {
	rv := objc.Call[float64](d_, objc.Sel("strands"))
	return rv
}

func (d_ DrosteWrapper) HasSetInsetPoint1() bool {
	return d_.RespondsToSelector(objc.Sel("setInsetPoint1:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidroste/3600133-insetpoint1?language=objc
func (d_ DrosteWrapper) SetInsetPoint1(value coregraphics.Point) {
	objc.Call[objc.Void](d_, objc.Sel("setInsetPoint1:"), value)
}

func (d_ DrosteWrapper) HasInsetPoint1() bool {
	return d_.RespondsToSelector(objc.Sel("insetPoint1"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidroste/3600133-insetpoint1?language=objc
func (d_ DrosteWrapper) InsetPoint1() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](d_, objc.Sel("insetPoint1"))
	return rv
}
