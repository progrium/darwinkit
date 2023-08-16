// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a highlight-shadow adjust filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihighlightshadowadjust?language=objc
type PHighlightShadowAdjust interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetShadowAmount(value float64)
	HasSetShadowAmount() bool

	// optional
	ShadowAmount() float64
	HasShadowAmount() bool

	// optional
	SetRadius(value float64)
	HasSetRadius() bool

	// optional
	Radius() float64
	HasRadius() bool

	// optional
	SetHighlightAmount(value float64)
	HasSetHighlightAmount() bool

	// optional
	HighlightAmount() float64
	HasHighlightAmount() bool
}

// A concrete type wrapper for the [PHighlightShadowAdjust] protocol.
type HighlightShadowAdjustWrapper struct {
	objc.Object
}

func (h_ HighlightShadowAdjustWrapper) HasSetInputImage() bool {
	return h_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihighlightshadowadjust/3228495-inputimage?language=objc
func (h_ HighlightShadowAdjustWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](h_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (h_ HighlightShadowAdjustWrapper) HasInputImage() bool {
	return h_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihighlightshadowadjust/3228495-inputimage?language=objc
func (h_ HighlightShadowAdjustWrapper) InputImage() Image {
	rv := objc.Call[Image](h_, objc.Sel("inputImage"))
	return rv
}

func (h_ HighlightShadowAdjustWrapper) HasSetShadowAmount() bool {
	return h_.RespondsToSelector(objc.Sel("setShadowAmount:"))
}

// The amount of adjustment to the shadows in the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihighlightshadowadjust/3228497-shadowamount?language=objc
func (h_ HighlightShadowAdjustWrapper) SetShadowAmount(value float64) {
	objc.Call[objc.Void](h_, objc.Sel("setShadowAmount:"), value)
}

func (h_ HighlightShadowAdjustWrapper) HasShadowAmount() bool {
	return h_.RespondsToSelector(objc.Sel("shadowAmount"))
}

// The amount of adjustment to the shadows in the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihighlightshadowadjust/3228497-shadowamount?language=objc
func (h_ HighlightShadowAdjustWrapper) ShadowAmount() float64 {
	rv := objc.Call[float64](h_, objc.Sel("shadowAmount"))
	return rv
}

func (h_ HighlightShadowAdjustWrapper) HasSetRadius() bool {
	return h_.RespondsToSelector(objc.Sel("setRadius:"))
}

// The shadow highlight radius. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihighlightshadowadjust/3228496-radius?language=objc
func (h_ HighlightShadowAdjustWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](h_, objc.Sel("setRadius:"), value)
}

func (h_ HighlightShadowAdjustWrapper) HasRadius() bool {
	return h_.RespondsToSelector(objc.Sel("radius"))
}

// The shadow highlight radius. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihighlightshadowadjust/3228496-radius?language=objc
func (h_ HighlightShadowAdjustWrapper) Radius() float64 {
	rv := objc.Call[float64](h_, objc.Sel("radius"))
	return rv
}

func (h_ HighlightShadowAdjustWrapper) HasSetHighlightAmount() bool {
	return h_.RespondsToSelector(objc.Sel("setHighlightAmount:"))
}

// The amount of adjustment to the highlights in the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihighlightshadowadjust/3228494-highlightamount?language=objc
func (h_ HighlightShadowAdjustWrapper) SetHighlightAmount(value float64) {
	objc.Call[objc.Void](h_, objc.Sel("setHighlightAmount:"), value)
}

func (h_ HighlightShadowAdjustWrapper) HasHighlightAmount() bool {
	return h_.RespondsToSelector(objc.Sel("highlightAmount"))
}

// The amount of adjustment to the highlights in the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihighlightshadowadjust/3228494-highlightamount?language=objc
func (h_ HighlightShadowAdjustWrapper) HighlightAmount() float64 {
	rv := objc.Call[float64](h_, objc.Sel("highlightAmount"))
	return rv
}
