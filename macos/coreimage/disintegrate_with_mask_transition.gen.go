// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a disintegrate-with-mask transition filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidisintegratewithmasktransition?language=objc
type PDisintegrateWithMaskTransition interface {
	// optional
	SetShadowDensity(value float64)
	HasSetShadowDensity() bool

	// optional
	ShadowDensity() float64
	HasShadowDensity() bool

	// optional
	SetShadowOffset(value coregraphics.Point)
	HasSetShadowOffset() bool

	// optional
	ShadowOffset() coregraphics.Point
	HasShadowOffset() bool

	// optional
	SetMaskImage(value Image)
	HasSetMaskImage() bool

	// optional
	MaskImage() IImage
	HasMaskImage() bool

	// optional
	SetShadowRadius(value float64)
	HasSetShadowRadius() bool

	// optional
	ShadowRadius() float64
	HasShadowRadius() bool
}

// A concrete type wrapper for the [PDisintegrateWithMaskTransition] protocol.
type DisintegrateWithMaskTransitionWrapper struct {
	objc.Object
}

func (d_ DisintegrateWithMaskTransitionWrapper) HasSetShadowDensity() bool {
	return d_.RespondsToSelector(objc.Sel("setShadowDensity:"))
}

// The density of the shadow the mask creates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidisintegratewithmasktransition/3228218-shadowdensity?language=objc
func (d_ DisintegrateWithMaskTransitionWrapper) SetShadowDensity(value float64) {
	objc.Call[objc.Void](d_, objc.Sel("setShadowDensity:"), value)
}

func (d_ DisintegrateWithMaskTransitionWrapper) HasShadowDensity() bool {
	return d_.RespondsToSelector(objc.Sel("shadowDensity"))
}

// The density of the shadow the mask creates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidisintegratewithmasktransition/3228218-shadowdensity?language=objc
func (d_ DisintegrateWithMaskTransitionWrapper) ShadowDensity() float64 {
	rv := objc.Call[float64](d_, objc.Sel("shadowDensity"))
	return rv
}

func (d_ DisintegrateWithMaskTransitionWrapper) HasSetShadowOffset() bool {
	return d_.RespondsToSelector(objc.Sel("setShadowOffset:"))
}

// The offset of the shadow the mask creates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidisintegratewithmasktransition/3228219-shadowoffset?language=objc
func (d_ DisintegrateWithMaskTransitionWrapper) SetShadowOffset(value coregraphics.Point) {
	objc.Call[objc.Void](d_, objc.Sel("setShadowOffset:"), value)
}

func (d_ DisintegrateWithMaskTransitionWrapper) HasShadowOffset() bool {
	return d_.RespondsToSelector(objc.Sel("shadowOffset"))
}

// The offset of the shadow the mask creates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidisintegratewithmasktransition/3228219-shadowoffset?language=objc
func (d_ DisintegrateWithMaskTransitionWrapper) ShadowOffset() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](d_, objc.Sel("shadowOffset"))
	return rv
}

func (d_ DisintegrateWithMaskTransitionWrapper) HasSetMaskImage() bool {
	return d_.RespondsToSelector(objc.Sel("setMaskImage:"))
}

// An image that defines the shape to use when disintegrating from the source to the target image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidisintegratewithmasktransition/3228217-maskimage?language=objc
func (d_ DisintegrateWithMaskTransitionWrapper) SetMaskImage(value IImage) {
	objc.Call[objc.Void](d_, objc.Sel("setMaskImage:"), objc.Ptr(value))
}

func (d_ DisintegrateWithMaskTransitionWrapper) HasMaskImage() bool {
	return d_.RespondsToSelector(objc.Sel("maskImage"))
}

// An image that defines the shape to use when disintegrating from the source to the target image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidisintegratewithmasktransition/3228217-maskimage?language=objc
func (d_ DisintegrateWithMaskTransitionWrapper) MaskImage() Image {
	rv := objc.Call[Image](d_, objc.Sel("maskImage"))
	return rv
}

func (d_ DisintegrateWithMaskTransitionWrapper) HasSetShadowRadius() bool {
	return d_.RespondsToSelector(objc.Sel("setShadowRadius:"))
}

// The radius of the shadow the mask creates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidisintegratewithmasktransition/3228220-shadowradius?language=objc
func (d_ DisintegrateWithMaskTransitionWrapper) SetShadowRadius(value float64) {
	objc.Call[objc.Void](d_, objc.Sel("setShadowRadius:"), value)
}

func (d_ DisintegrateWithMaskTransitionWrapper) HasShadowRadius() bool {
	return d_.RespondsToSelector(objc.Sel("shadowRadius"))
}

// The radius of the shadow the mask creates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidisintegratewithmasktransition/3228220-shadowradius?language=objc
func (d_ DisintegrateWithMaskTransitionWrapper) ShadowRadius() float64 {
	rv := objc.Call[float64](d_, objc.Sel("shadowRadius"))
	return rv
}
