// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a swipe transition filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciswipetransition?language=objc
type PSwipeTransition interface {
	// optional
	SetColor(value Color)
	HasSetColor() bool

	// optional
	Color() IColor
	HasColor() bool

	// optional
	SetWidth(value float64)
	HasSetWidth() bool

	// optional
	Width() float64
	HasWidth() bool

	// optional
	SetOpacity(value float64)
	HasSetOpacity() bool

	// optional
	Opacity() float64
	HasOpacity() bool

	// optional
	SetExtent(value coregraphics.Rect)
	HasSetExtent() bool

	// optional
	Extent() coregraphics.Rect
	HasExtent() bool

	// optional
	SetAngle(value float64)
	HasSetAngle() bool

	// optional
	Angle() float64
	HasAngle() bool
}

// A concrete type wrapper for the [PSwipeTransition] protocol.
type SwipeTransitionWrapper struct {
	objc.Object
}

func (s_ SwipeTransitionWrapper) HasSetColor() bool {
	return s_.RespondsToSelector(objc.Sel("setColor:"))
}

// The color of the swipe. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciswipetransition/3228776-color?language=objc
func (s_ SwipeTransitionWrapper) SetColor(value IColor) {
	objc.Call[objc.Void](s_, objc.Sel("setColor:"), objc.Ptr(value))
}

func (s_ SwipeTransitionWrapper) HasColor() bool {
	return s_.RespondsToSelector(objc.Sel("color"))
}

// The color of the swipe. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciswipetransition/3228776-color?language=objc
func (s_ SwipeTransitionWrapper) Color() Color {
	rv := objc.Call[Color](s_, objc.Sel("color"))
	return rv
}

func (s_ SwipeTransitionWrapper) HasSetWidth() bool {
	return s_.RespondsToSelector(objc.Sel("setWidth:"))
}

// The width of the swipe. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciswipetransition/3228779-width?language=objc
func (s_ SwipeTransitionWrapper) SetWidth(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setWidth:"), value)
}

func (s_ SwipeTransitionWrapper) HasWidth() bool {
	return s_.RespondsToSelector(objc.Sel("width"))
}

// The width of the swipe. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciswipetransition/3228779-width?language=objc
func (s_ SwipeTransitionWrapper) Width() float64 {
	rv := objc.Call[float64](s_, objc.Sel("width"))
	return rv
}

func (s_ SwipeTransitionWrapper) HasSetOpacity() bool {
	return s_.RespondsToSelector(objc.Sel("setOpacity:"))
}

// The opacity of the swipe. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciswipetransition/3228778-opacity?language=objc
func (s_ SwipeTransitionWrapper) SetOpacity(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setOpacity:"), value)
}

func (s_ SwipeTransitionWrapper) HasOpacity() bool {
	return s_.RespondsToSelector(objc.Sel("opacity"))
}

// The opacity of the swipe. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciswipetransition/3228778-opacity?language=objc
func (s_ SwipeTransitionWrapper) Opacity() float64 {
	rv := objc.Call[float64](s_, objc.Sel("opacity"))
	return rv
}

func (s_ SwipeTransitionWrapper) HasSetExtent() bool {
	return s_.RespondsToSelector(objc.Sel("setExtent:"))
}

// The extent of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciswipetransition/3228777-extent?language=objc
func (s_ SwipeTransitionWrapper) SetExtent(value coregraphics.Rect) {
	objc.Call[objc.Void](s_, objc.Sel("setExtent:"), value)
}

func (s_ SwipeTransitionWrapper) HasExtent() bool {
	return s_.RespondsToSelector(objc.Sel("extent"))
}

// The extent of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciswipetransition/3228777-extent?language=objc
func (s_ SwipeTransitionWrapper) Extent() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](s_, objc.Sel("extent"))
	return rv
}

func (s_ SwipeTransitionWrapper) HasSetAngle() bool {
	return s_.RespondsToSelector(objc.Sel("setAngle:"))
}

// The angle of the swipe. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciswipetransition/3228775-angle?language=objc
func (s_ SwipeTransitionWrapper) SetAngle(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setAngle:"), value)
}

func (s_ SwipeTransitionWrapper) HasAngle() bool {
	return s_.RespondsToSelector(objc.Sel("angle"))
}

// The angle of the swipe. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciswipetransition/3228775-angle?language=objc
func (s_ SwipeTransitionWrapper) Angle() float64 {
	rv := objc.Call[float64](s_, objc.Sel("angle"))
	return rv
}
