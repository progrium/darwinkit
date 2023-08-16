// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a copy machine transition filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicopymachinetransition?language=objc
type PCopyMachineTransition interface {
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

// A concrete type wrapper for the [PCopyMachineTransition] protocol.
type CopyMachineTransitionWrapper struct {
	objc.Object
}

func (c_ CopyMachineTransitionWrapper) HasSetColor() bool {
	return c_.RespondsToSelector(objc.Sel("setColor:"))
}

// The color of the copier light. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicopymachinetransition/3228190-color?language=objc
func (c_ CopyMachineTransitionWrapper) SetColor(value IColor) {
	objc.Call[objc.Void](c_, objc.Sel("setColor:"), objc.Ptr(value))
}

func (c_ CopyMachineTransitionWrapper) HasColor() bool {
	return c_.RespondsToSelector(objc.Sel("color"))
}

// The color of the copier light. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicopymachinetransition/3228190-color?language=objc
func (c_ CopyMachineTransitionWrapper) Color() Color {
	rv := objc.Call[Color](c_, objc.Sel("color"))
	return rv
}

func (c_ CopyMachineTransitionWrapper) HasSetWidth() bool {
	return c_.RespondsToSelector(objc.Sel("setWidth:"))
}

// The width of the copier light. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicopymachinetransition/3228193-width?language=objc
func (c_ CopyMachineTransitionWrapper) SetWidth(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setWidth:"), value)
}

func (c_ CopyMachineTransitionWrapper) HasWidth() bool {
	return c_.RespondsToSelector(objc.Sel("width"))
}

// The width of the copier light. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicopymachinetransition/3228193-width?language=objc
func (c_ CopyMachineTransitionWrapper) Width() float64 {
	rv := objc.Call[float64](c_, objc.Sel("width"))
	return rv
}

func (c_ CopyMachineTransitionWrapper) HasSetOpacity() bool {
	return c_.RespondsToSelector(objc.Sel("setOpacity:"))
}

// The opacity of the copier light. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicopymachinetransition/3228192-opacity?language=objc
func (c_ CopyMachineTransitionWrapper) SetOpacity(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setOpacity:"), value)
}

func (c_ CopyMachineTransitionWrapper) HasOpacity() bool {
	return c_.RespondsToSelector(objc.Sel("opacity"))
}

// The opacity of the copier light. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicopymachinetransition/3228192-opacity?language=objc
func (c_ CopyMachineTransitionWrapper) Opacity() float64 {
	rv := objc.Call[float64](c_, objc.Sel("opacity"))
	return rv
}

func (c_ CopyMachineTransitionWrapper) HasSetExtent() bool {
	return c_.RespondsToSelector(objc.Sel("setExtent:"))
}

// A rectangle that defines the extent of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicopymachinetransition/3228191-extent?language=objc
func (c_ CopyMachineTransitionWrapper) SetExtent(value coregraphics.Rect) {
	objc.Call[objc.Void](c_, objc.Sel("setExtent:"), value)
}

func (c_ CopyMachineTransitionWrapper) HasExtent() bool {
	return c_.RespondsToSelector(objc.Sel("extent"))
}

// A rectangle that defines the extent of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicopymachinetransition/3228191-extent?language=objc
func (c_ CopyMachineTransitionWrapper) Extent() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](c_, objc.Sel("extent"))
	return rv
}

func (c_ CopyMachineTransitionWrapper) HasSetAngle() bool {
	return c_.RespondsToSelector(objc.Sel("setAngle:"))
}

// The angle of the copier light. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicopymachinetransition/3228189-angle?language=objc
func (c_ CopyMachineTransitionWrapper) SetAngle(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setAngle:"), value)
}

func (c_ CopyMachineTransitionWrapper) HasAngle() bool {
	return c_.RespondsToSelector(objc.Sel("angle"))
}

// The angle of the copier light. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicopymachinetransition/3228189-angle?language=objc
func (c_ CopyMachineTransitionWrapper) Angle() float64 {
	rv := objc.Call[float64](c_, objc.Sel("angle"))
	return rv
}
