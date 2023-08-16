// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorthreshold?language=objc
type PColorThreshold interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetThreshold(value float64)
	HasSetThreshold() bool

	// optional
	Threshold() float64
	HasThreshold() bool
}

// A concrete type wrapper for the [PColorThreshold] protocol.
type ColorThresholdWrapper struct {
	objc.Object
}

func (c_ ColorThresholdWrapper) HasSetInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("setInputImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorthreshold/3547107-inputimage?language=objc
func (c_ ColorThresholdWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](c_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (c_ ColorThresholdWrapper) HasInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("inputImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorthreshold/3547107-inputimage?language=objc
func (c_ ColorThresholdWrapper) InputImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("inputImage"))
	return rv
}

func (c_ ColorThresholdWrapper) HasSetThreshold() bool {
	return c_.RespondsToSelector(objc.Sel("setThreshold:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorthreshold/3547108-threshold?language=objc
func (c_ ColorThresholdWrapper) SetThreshold(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setThreshold:"), value)
}

func (c_ ColorThresholdWrapper) HasThreshold() bool {
	return c_.RespondsToSelector(objc.Sel("threshold"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorthreshold/3547108-threshold?language=objc
func (c_ ColorThresholdWrapper) Threshold() float64 {
	rv := objc.Call[float64](c_, objc.Sel("threshold"))
	return rv
}
