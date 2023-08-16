// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a line overlay filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineoverlay?language=objc
type PLineOverlay interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetEdgeIntensity(value float64)
	HasSetEdgeIntensity() bool

	// optional
	EdgeIntensity() float64
	HasEdgeIntensity() bool

	// optional
	SetNRNoiseLevel(value float64)
	HasSetNRNoiseLevel() bool

	// optional
	NRNoiseLevel() float64
	HasNRNoiseLevel() bool

	// optional
	SetNRSharpness(value float64)
	HasSetNRSharpness() bool

	// optional
	NRSharpness() float64
	HasNRSharpness() bool

	// optional
	SetContrast(value float64)
	HasSetContrast() bool

	// optional
	Contrast() float64
	HasContrast() bool

	// optional
	SetThreshold(value float64)
	HasSetThreshold() bool

	// optional
	Threshold() float64
	HasThreshold() bool
}

// A concrete type wrapper for the [PLineOverlay] protocol.
type LineOverlayWrapper struct {
	objc.Object
}

func (l_ LineOverlayWrapper) HasSetInputImage() bool {
	return l_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineoverlay/3228533-inputimage?language=objc
func (l_ LineOverlayWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](l_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (l_ LineOverlayWrapper) HasInputImage() bool {
	return l_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineoverlay/3228533-inputimage?language=objc
func (l_ LineOverlayWrapper) InputImage() Image {
	rv := objc.Call[Image](l_, objc.Sel("inputImage"))
	return rv
}

func (l_ LineOverlayWrapper) HasSetEdgeIntensity() bool {
	return l_.RespondsToSelector(objc.Sel("setEdgeIntensity:"))
}

// The accentuation factor of the Sobel gradient information when tracing the edges of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineoverlay/3228532-edgeintensity?language=objc
func (l_ LineOverlayWrapper) SetEdgeIntensity(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setEdgeIntensity:"), value)
}

func (l_ LineOverlayWrapper) HasEdgeIntensity() bool {
	return l_.RespondsToSelector(objc.Sel("edgeIntensity"))
}

// The accentuation factor of the Sobel gradient information when tracing the edges of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineoverlay/3228532-edgeintensity?language=objc
func (l_ LineOverlayWrapper) EdgeIntensity() float64 {
	rv := objc.Call[float64](l_, objc.Sel("edgeIntensity"))
	return rv
}

func (l_ LineOverlayWrapper) HasSetNRNoiseLevel() bool {
	return l_.RespondsToSelector(objc.Sel("setNRNoiseLevel:"))
}

// The noise level of the image, used with camera data, that's removed before tracing the edges of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineoverlay/3228529-nrnoiselevel?language=objc
func (l_ LineOverlayWrapper) SetNRNoiseLevel(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setNRNoiseLevel:"), value)
}

func (l_ LineOverlayWrapper) HasNRNoiseLevel() bool {
	return l_.RespondsToSelector(objc.Sel("NRNoiseLevel"))
}

// The noise level of the image, used with camera data, that's removed before tracing the edges of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineoverlay/3228529-nrnoiselevel?language=objc
func (l_ LineOverlayWrapper) NRNoiseLevel() float64 {
	rv := objc.Call[float64](l_, objc.Sel("NRNoiseLevel"))
	return rv
}

func (l_ LineOverlayWrapper) HasSetNRSharpness() bool {
	return l_.RespondsToSelector(objc.Sel("setNRSharpness:"))
}

// The amount of sharpening done when removing noise in the image before tracing the edges of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineoverlay/3228530-nrsharpness?language=objc
func (l_ LineOverlayWrapper) SetNRSharpness(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setNRSharpness:"), value)
}

func (l_ LineOverlayWrapper) HasNRSharpness() bool {
	return l_.RespondsToSelector(objc.Sel("NRSharpness"))
}

// The amount of sharpening done when removing noise in the image before tracing the edges of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineoverlay/3228530-nrsharpness?language=objc
func (l_ LineOverlayWrapper) NRSharpness() float64 {
	rv := objc.Call[float64](l_, objc.Sel("NRSharpness"))
	return rv
}

func (l_ LineOverlayWrapper) HasSetContrast() bool {
	return l_.RespondsToSelector(objc.Sel("setContrast:"))
}

// The amount of antialiasing to use on the edges produced by this filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineoverlay/3228531-contrast?language=objc
func (l_ LineOverlayWrapper) SetContrast(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setContrast:"), value)
}

func (l_ LineOverlayWrapper) HasContrast() bool {
	return l_.RespondsToSelector(objc.Sel("contrast"))
}

// The amount of antialiasing to use on the edges produced by this filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineoverlay/3228531-contrast?language=objc
func (l_ LineOverlayWrapper) Contrast() float64 {
	rv := objc.Call[float64](l_, objc.Sel("contrast"))
	return rv
}

func (l_ LineOverlayWrapper) HasSetThreshold() bool {
	return l_.RespondsToSelector(objc.Sel("setThreshold:"))
}

// A value that determines edge visibility. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineoverlay/3228534-threshold?language=objc
func (l_ LineOverlayWrapper) SetThreshold(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setThreshold:"), value)
}

func (l_ LineOverlayWrapper) HasThreshold() bool {
	return l_.RespondsToSelector(objc.Sel("threshold"))
}

// A value that determines edge visibility. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineoverlay/3228534-threshold?language=objc
func (l_ LineOverlayWrapper) Threshold() float64 {
	rv := objc.Call[float64](l_, objc.Sel("threshold"))
	return rv
}
