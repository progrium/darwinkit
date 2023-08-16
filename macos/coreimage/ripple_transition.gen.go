// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a ripple transition filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirippletransition?language=objc
type PRippleTransition interface {
	// optional
	SetWidth(value float64)
	HasSetWidth() bool

	// optional
	Width() float64
	HasWidth() bool

	// optional
	SetScale(value float64)
	HasSetScale() bool

	// optional
	Scale() float64
	HasScale() bool

	// optional
	SetExtent(value coregraphics.Rect)
	HasSetExtent() bool

	// optional
	Extent() coregraphics.Rect
	HasExtent() bool

	// optional
	SetShadingImage(value Image)
	HasSetShadingImage() bool

	// optional
	ShadingImage() IImage
	HasShadingImage() bool

	// optional
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool
}

// A concrete type wrapper for the [PRippleTransition] protocol.
type RippleTransitionWrapper struct {
	objc.Object
}

func (r_ RippleTransitionWrapper) HasSetWidth() bool {
	return r_.RespondsToSelector(objc.Sel("setWidth:"))
}

// The width of the ripple. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirippletransition/3228696-width?language=objc
func (r_ RippleTransitionWrapper) SetWidth(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setWidth:"), value)
}

func (r_ RippleTransitionWrapper) HasWidth() bool {
	return r_.RespondsToSelector(objc.Sel("width"))
}

// The width of the ripple. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirippletransition/3228696-width?language=objc
func (r_ RippleTransitionWrapper) Width() float64 {
	rv := objc.Call[float64](r_, objc.Sel("width"))
	return rv
}

func (r_ RippleTransitionWrapper) HasSetScale() bool {
	return r_.RespondsToSelector(objc.Sel("setScale:"))
}

// A value that determines whether the ripple starts as a bulge (a higher value) or a dimple (a lower value). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirippletransition/3228694-scale?language=objc
func (r_ RippleTransitionWrapper) SetScale(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setScale:"), value)
}

func (r_ RippleTransitionWrapper) HasScale() bool {
	return r_.RespondsToSelector(objc.Sel("scale"))
}

// A value that determines whether the ripple starts as a bulge (a higher value) or a dimple (a lower value). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirippletransition/3228694-scale?language=objc
func (r_ RippleTransitionWrapper) Scale() float64 {
	rv := objc.Call[float64](r_, objc.Sel("scale"))
	return rv
}

func (r_ RippleTransitionWrapper) HasSetExtent() bool {
	return r_.RespondsToSelector(objc.Sel("setExtent:"))
}

// A rectangle that defines the extent of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirippletransition/3228693-extent?language=objc
func (r_ RippleTransitionWrapper) SetExtent(value coregraphics.Rect) {
	objc.Call[objc.Void](r_, objc.Sel("setExtent:"), value)
}

func (r_ RippleTransitionWrapper) HasExtent() bool {
	return r_.RespondsToSelector(objc.Sel("extent"))
}

// A rectangle that defines the extent of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirippletransition/3228693-extent?language=objc
func (r_ RippleTransitionWrapper) Extent() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](r_, objc.Sel("extent"))
	return rv
}

func (r_ RippleTransitionWrapper) HasSetShadingImage() bool {
	return r_.RespondsToSelector(objc.Sel("setShadingImage:"))
}

// An image that looks like a shaded sphere enclosed in a square. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirippletransition/3228695-shadingimage?language=objc
func (r_ RippleTransitionWrapper) SetShadingImage(value IImage) {
	objc.Call[objc.Void](r_, objc.Sel("setShadingImage:"), objc.Ptr(value))
}

func (r_ RippleTransitionWrapper) HasShadingImage() bool {
	return r_.RespondsToSelector(objc.Sel("shadingImage"))
}

// An image that looks like a shaded sphere enclosed in a square. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirippletransition/3228695-shadingimage?language=objc
func (r_ RippleTransitionWrapper) ShadingImage() Image {
	rv := objc.Call[Image](r_, objc.Sel("shadingImage"))
	return rv
}

func (r_ RippleTransitionWrapper) HasSetCenter() bool {
	return r_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirippletransition/3228692-center?language=objc
func (r_ RippleTransitionWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](r_, objc.Sel("setCenter:"), value)
}

func (r_ RippleTransitionWrapper) HasCenter() bool {
	return r_.RespondsToSelector(objc.Sel("center"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirippletransition/3228692-center?language=objc
func (r_ RippleTransitionWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](r_, objc.Sel("center"))
	return rv
}
