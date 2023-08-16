// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a transition filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citransitionfilter?language=objc
type PTransitionFilter interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetTargetImage(value Image)
	HasSetTargetImage() bool

	// optional
	TargetImage() IImage
	HasTargetImage() bool

	// optional
	SetTime(value float64)
	HasSetTime() bool

	// optional
	Time() float64
	HasTime() bool
}

// A concrete type wrapper for the [PTransitionFilter] protocol.
type TransitionFilterWrapper struct {
	objc.Object
}

func (t_ TransitionFilterWrapper) HasSetInputImage() bool {
	return t_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citransitionfilter/3228799-inputimage?language=objc
func (t_ TransitionFilterWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](t_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (t_ TransitionFilterWrapper) HasInputImage() bool {
	return t_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citransitionfilter/3228799-inputimage?language=objc
func (t_ TransitionFilterWrapper) InputImage() Image {
	rv := objc.Call[Image](t_, objc.Sel("inputImage"))
	return rv
}

func (t_ TransitionFilterWrapper) HasSetTargetImage() bool {
	return t_.RespondsToSelector(objc.Sel("setTargetImage:"))
}

// The target image for a transition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citransitionfilter/3228800-targetimage?language=objc
func (t_ TransitionFilterWrapper) SetTargetImage(value IImage) {
	objc.Call[objc.Void](t_, objc.Sel("setTargetImage:"), objc.Ptr(value))
}

func (t_ TransitionFilterWrapper) HasTargetImage() bool {
	return t_.RespondsToSelector(objc.Sel("targetImage"))
}

// The target image for a transition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citransitionfilter/3228800-targetimage?language=objc
func (t_ TransitionFilterWrapper) TargetImage() Image {
	rv := objc.Call[Image](t_, objc.Sel("targetImage"))
	return rv
}

func (t_ TransitionFilterWrapper) HasSetTime() bool {
	return t_.RespondsToSelector(objc.Sel("setTime:"))
}

// The parametric time of the transition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citransitionfilter/3228801-time?language=objc
func (t_ TransitionFilterWrapper) SetTime(value float64) {
	objc.Call[objc.Void](t_, objc.Sel("setTime:"), value)
}

func (t_ TransitionFilterWrapper) HasTime() bool {
	return t_.RespondsToSelector(objc.Sel("time"))
}

// The parametric time of the transition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citransitionfilter/3228801-time?language=objc
func (t_ TransitionFilterWrapper) Time() float64 {
	rv := objc.Call[float64](t_, objc.Sel("time"))
	return rv
}
