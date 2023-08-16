// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a bars swipe transition filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibarsswipetransition?language=objc
type PBarsSwipeTransition interface {
	// optional
	SetWidth(value float64)
	HasSetWidth() bool

	// optional
	Width() float64
	HasWidth() bool

	// optional
	SetBarOffset(value float64)
	HasSetBarOffset() bool

	// optional
	BarOffset() float64
	HasBarOffset() bool

	// optional
	SetAngle(value float64)
	HasSetAngle() bool

	// optional
	Angle() float64
	HasAngle() bool
}

// A concrete type wrapper for the [PBarsSwipeTransition] protocol.
type BarsSwipeTransitionWrapper struct {
	objc.Object
}

func (b_ BarsSwipeTransitionWrapper) HasSetWidth() bool {
	return b_.RespondsToSelector(objc.Sel("setWidth:"))
}

// The width of each bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibarsswipetransition/3228072-width?language=objc
func (b_ BarsSwipeTransitionWrapper) SetWidth(value float64) {
	objc.Call[objc.Void](b_, objc.Sel("setWidth:"), value)
}

func (b_ BarsSwipeTransitionWrapper) HasWidth() bool {
	return b_.RespondsToSelector(objc.Sel("width"))
}

// The width of each bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibarsswipetransition/3228072-width?language=objc
func (b_ BarsSwipeTransitionWrapper) Width() float64 {
	rv := objc.Call[float64](b_, objc.Sel("width"))
	return rv
}

func (b_ BarsSwipeTransitionWrapper) HasSetBarOffset() bool {
	return b_.RespondsToSelector(objc.Sel("setBarOffset:"))
}

// The offset of one bar with respect to another. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibarsswipetransition/3228071-baroffset?language=objc
func (b_ BarsSwipeTransitionWrapper) SetBarOffset(value float64) {
	objc.Call[objc.Void](b_, objc.Sel("setBarOffset:"), value)
}

func (b_ BarsSwipeTransitionWrapper) HasBarOffset() bool {
	return b_.RespondsToSelector(objc.Sel("barOffset"))
}

// The offset of one bar with respect to another. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibarsswipetransition/3228071-baroffset?language=objc
func (b_ BarsSwipeTransitionWrapper) BarOffset() float64 {
	rv := objc.Call[float64](b_, objc.Sel("barOffset"))
	return rv
}

func (b_ BarsSwipeTransitionWrapper) HasSetAngle() bool {
	return b_.RespondsToSelector(objc.Sel("setAngle:"))
}

// The angle, in radians, of the bars. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibarsswipetransition/3228070-angle?language=objc
func (b_ BarsSwipeTransitionWrapper) SetAngle(value float64) {
	objc.Call[objc.Void](b_, objc.Sel("setAngle:"), value)
}

func (b_ BarsSwipeTransitionWrapper) HasAngle() bool {
	return b_.RespondsToSelector(objc.Sel("angle"))
}

// The angle, in radians, of the bars. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibarsswipetransition/3228070-angle?language=objc
func (b_ BarsSwipeTransitionWrapper) Angle() float64 {
	rv := objc.Call[float64](b_, objc.Sel("angle"))
	return rv
}
