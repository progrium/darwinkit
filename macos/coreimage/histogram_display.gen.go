// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihistogramdisplay?language=objc
type PHistogramDisplay interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetHeight(value float64)
	HasSetHeight() bool

	// optional
	Height() float64
	HasHeight() bool

	// optional
	SetHighLimit(value float64)
	HasSetHighLimit() bool

	// optional
	HighLimit() float64
	HasHighLimit() bool

	// optional
	SetLowLimit(value float64)
	HasSetLowLimit() bool

	// optional
	LowLimit() float64
	HasLowLimit() bool
}

// A concrete type wrapper for the [PHistogramDisplay] protocol.
type HistogramDisplayWrapper struct {
	objc.Object
}

func (h_ HistogramDisplayWrapper) HasSetInputImage() bool {
	return h_.RespondsToSelector(objc.Sel("setInputImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihistogramdisplay/3547127-inputimage?language=objc
func (h_ HistogramDisplayWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](h_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (h_ HistogramDisplayWrapper) HasInputImage() bool {
	return h_.RespondsToSelector(objc.Sel("inputImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihistogramdisplay/3547127-inputimage?language=objc
func (h_ HistogramDisplayWrapper) InputImage() Image {
	rv := objc.Call[Image](h_, objc.Sel("inputImage"))
	return rv
}

func (h_ HistogramDisplayWrapper) HasSetHeight() bool {
	return h_.RespondsToSelector(objc.Sel("setHeight:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihistogramdisplay/3547125-height?language=objc
func (h_ HistogramDisplayWrapper) SetHeight(value float64) {
	objc.Call[objc.Void](h_, objc.Sel("setHeight:"), value)
}

func (h_ HistogramDisplayWrapper) HasHeight() bool {
	return h_.RespondsToSelector(objc.Sel("height"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihistogramdisplay/3547125-height?language=objc
func (h_ HistogramDisplayWrapper) Height() float64 {
	rv := objc.Call[float64](h_, objc.Sel("height"))
	return rv
}

func (h_ HistogramDisplayWrapper) HasSetHighLimit() bool {
	return h_.RespondsToSelector(objc.Sel("setHighLimit:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihistogramdisplay/3547126-highlimit?language=objc
func (h_ HistogramDisplayWrapper) SetHighLimit(value float64) {
	objc.Call[objc.Void](h_, objc.Sel("setHighLimit:"), value)
}

func (h_ HistogramDisplayWrapper) HasHighLimit() bool {
	return h_.RespondsToSelector(objc.Sel("highLimit"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihistogramdisplay/3547126-highlimit?language=objc
func (h_ HistogramDisplayWrapper) HighLimit() float64 {
	rv := objc.Call[float64](h_, objc.Sel("highLimit"))
	return rv
}

func (h_ HistogramDisplayWrapper) HasSetLowLimit() bool {
	return h_.RespondsToSelector(objc.Sel("setLowLimit:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihistogramdisplay/3547128-lowlimit?language=objc
func (h_ HistogramDisplayWrapper) SetLowLimit(value float64) {
	objc.Call[objc.Void](h_, objc.Sel("setLowLimit:"), value)
}

func (h_ HistogramDisplayWrapper) HasLowLimit() bool {
	return h_.RespondsToSelector(objc.Sel("lowLimit"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihistogramdisplay/3547128-lowlimit?language=objc
func (h_ HistogramDisplayWrapper) LowLimit() float64 {
	rv := objc.Call[float64](h_, objc.Sel("lowLimit"))
	return rv
}
