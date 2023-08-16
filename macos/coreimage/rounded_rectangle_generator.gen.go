// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a rounded rectangle generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciroundedrectanglegenerator?language=objc
type PRoundedRectangleGenerator interface {
	// optional
	SetColor(value Color)
	HasSetColor() bool

	// optional
	Color() IColor
	HasColor() bool

	// optional
	SetRadius(value float64)
	HasSetRadius() bool

	// optional
	Radius() float64
	HasRadius() bool

	// optional
	SetExtent(value coregraphics.Rect)
	HasSetExtent() bool

	// optional
	Extent() coregraphics.Rect
	HasExtent() bool
}

// A concrete type wrapper for the [PRoundedRectangleGenerator] protocol.
type RoundedRectangleGeneratorWrapper struct {
	objc.Object
}

func (r_ RoundedRectangleGeneratorWrapper) HasSetColor() bool {
	return r_.RespondsToSelector(objc.Sel("setColor:"))
}

// The color of the rounded rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciroundedrectanglegenerator/3338737-color?language=objc
func (r_ RoundedRectangleGeneratorWrapper) SetColor(value IColor) {
	objc.Call[objc.Void](r_, objc.Sel("setColor:"), objc.Ptr(value))
}

func (r_ RoundedRectangleGeneratorWrapper) HasColor() bool {
	return r_.RespondsToSelector(objc.Sel("color"))
}

// The color of the rounded rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciroundedrectanglegenerator/3338737-color?language=objc
func (r_ RoundedRectangleGeneratorWrapper) Color() Color {
	rv := objc.Call[Color](r_, objc.Sel("color"))
	return rv
}

func (r_ RoundedRectangleGeneratorWrapper) HasSetRadius() bool {
	return r_.RespondsToSelector(objc.Sel("setRadius:"))
}

// The distance from the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciroundedrectanglegenerator/3338739-radius?language=objc
func (r_ RoundedRectangleGeneratorWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setRadius:"), value)
}

func (r_ RoundedRectangleGeneratorWrapper) HasRadius() bool {
	return r_.RespondsToSelector(objc.Sel("radius"))
}

// The distance from the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciroundedrectanglegenerator/3338739-radius?language=objc
func (r_ RoundedRectangleGeneratorWrapper) Radius() float64 {
	rv := objc.Call[float64](r_, objc.Sel("radius"))
	return rv
}

func (r_ RoundedRectangleGeneratorWrapper) HasSetExtent() bool {
	return r_.RespondsToSelector(objc.Sel("setExtent:"))
}

// A rectangle that defines the extent of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciroundedrectanglegenerator/3338738-extent?language=objc
func (r_ RoundedRectangleGeneratorWrapper) SetExtent(value coregraphics.Rect) {
	objc.Call[objc.Void](r_, objc.Sel("setExtent:"), value)
}

func (r_ RoundedRectangleGeneratorWrapper) HasExtent() bool {
	return r_.RespondsToSelector(objc.Sel("extent"))
}

// A rectangle that defines the extent of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciroundedrectanglegenerator/3338738-extent?language=objc
func (r_ RoundedRectangleGeneratorWrapper) Extent() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](r_, objc.Sel("extent"))
	return rv
}
