// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilighttunnel?language=objc
type PLightTunnel interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetRotation(value float64)
	HasSetRotation() bool

	// optional
	Rotation() float64
	HasRotation() bool

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

// A concrete type wrapper for the [PLightTunnel] protocol.
type LightTunnelWrapper struct {
	objc.Object
}

func (l_ LightTunnelWrapper) HasSetInputImage() bool {
	return l_.RespondsToSelector(objc.Sel("setInputImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilighttunnel/3600172-inputimage?language=objc
func (l_ LightTunnelWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](l_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (l_ LightTunnelWrapper) HasInputImage() bool {
	return l_.RespondsToSelector(objc.Sel("inputImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilighttunnel/3600172-inputimage?language=objc
func (l_ LightTunnelWrapper) InputImage() Image {
	rv := objc.Call[Image](l_, objc.Sel("inputImage"))
	return rv
}

func (l_ LightTunnelWrapper) HasSetRotation() bool {
	return l_.RespondsToSelector(objc.Sel("setRotation:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilighttunnel/3600174-rotation?language=objc
func (l_ LightTunnelWrapper) SetRotation(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setRotation:"), value)
}

func (l_ LightTunnelWrapper) HasRotation() bool {
	return l_.RespondsToSelector(objc.Sel("rotation"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilighttunnel/3600174-rotation?language=objc
func (l_ LightTunnelWrapper) Rotation() float64 {
	rv := objc.Call[float64](l_, objc.Sel("rotation"))
	return rv
}

func (l_ LightTunnelWrapper) HasSetRadius() bool {
	return l_.RespondsToSelector(objc.Sel("setRadius:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilighttunnel/3600173-radius?language=objc
func (l_ LightTunnelWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setRadius:"), value)
}

func (l_ LightTunnelWrapper) HasRadius() bool {
	return l_.RespondsToSelector(objc.Sel("radius"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilighttunnel/3600173-radius?language=objc
func (l_ LightTunnelWrapper) Radius() float64 {
	rv := objc.Call[float64](l_, objc.Sel("radius"))
	return rv
}

func (l_ LightTunnelWrapper) HasSetCenter() bool {
	return l_.RespondsToSelector(objc.Sel("setCenter:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilighttunnel/3600171-center?language=objc
func (l_ LightTunnelWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](l_, objc.Sel("setCenter:"), value)
}

func (l_ LightTunnelWrapper) HasCenter() bool {
	return l_.RespondsToSelector(objc.Sel("center"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilighttunnel/3600171-center?language=objc
func (l_ LightTunnelWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](l_, objc.Sel("center"))
	return rv
}
