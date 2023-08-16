// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure an edge preserve upsample filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciedgepreserveupsample?language=objc
type PEdgePreserveUpsample interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetSpatialSigma(value float64)
	HasSetSpatialSigma() bool

	// optional
	SpatialSigma() float64
	HasSpatialSigma() bool

	// optional
	SetSmallImage(value Image)
	HasSetSmallImage() bool

	// optional
	SmallImage() IImage
	HasSmallImage() bool

	// optional
	SetLumaSigma(value float64)
	HasSetLumaSigma() bool

	// optional
	LumaSigma() float64
	HasLumaSigma() bool
}

// A concrete type wrapper for the [PEdgePreserveUpsample] protocol.
type EdgePreserveUpsampleWrapper struct {
	objc.Object
}

func (e_ EdgePreserveUpsampleWrapper) HasSetInputImage() bool {
	return e_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciedgepreserveupsample/3228237-inputimage?language=objc
func (e_ EdgePreserveUpsampleWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](e_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (e_ EdgePreserveUpsampleWrapper) HasInputImage() bool {
	return e_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciedgepreserveupsample/3228237-inputimage?language=objc
func (e_ EdgePreserveUpsampleWrapper) InputImage() Image {
	rv := objc.Call[Image](e_, objc.Sel("inputImage"))
	return rv
}

func (e_ EdgePreserveUpsampleWrapper) HasSetSpatialSigma() bool {
	return e_.RespondsToSelector(objc.Sel("setSpatialSigma:"))
}

// A value that specifies the influence of the input image’s spatial information on the upsampling operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciedgepreserveupsample/3228240-spatialsigma?language=objc
func (e_ EdgePreserveUpsampleWrapper) SetSpatialSigma(value float64) {
	objc.Call[objc.Void](e_, objc.Sel("setSpatialSigma:"), value)
}

func (e_ EdgePreserveUpsampleWrapper) HasSpatialSigma() bool {
	return e_.RespondsToSelector(objc.Sel("spatialSigma"))
}

// A value that specifies the influence of the input image’s spatial information on the upsampling operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciedgepreserveupsample/3228240-spatialsigma?language=objc
func (e_ EdgePreserveUpsampleWrapper) SpatialSigma() float64 {
	rv := objc.Call[float64](e_, objc.Sel("spatialSigma"))
	return rv
}

func (e_ EdgePreserveUpsampleWrapper) HasSetSmallImage() bool {
	return e_.RespondsToSelector(objc.Sel("setSmallImage:"))
}

// The image that the filter upsamples. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciedgepreserveupsample/3228239-smallimage?language=objc
func (e_ EdgePreserveUpsampleWrapper) SetSmallImage(value IImage) {
	objc.Call[objc.Void](e_, objc.Sel("setSmallImage:"), objc.Ptr(value))
}

func (e_ EdgePreserveUpsampleWrapper) HasSmallImage() bool {
	return e_.RespondsToSelector(objc.Sel("smallImage"))
}

// The image that the filter upsamples. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciedgepreserveupsample/3228239-smallimage?language=objc
func (e_ EdgePreserveUpsampleWrapper) SmallImage() Image {
	rv := objc.Call[Image](e_, objc.Sel("smallImage"))
	return rv
}

func (e_ EdgePreserveUpsampleWrapper) HasSetLumaSigma() bool {
	return e_.RespondsToSelector(objc.Sel("setLumaSigma:"))
}

// A value that specifies the influence of the input image’s luma information on the upsampling operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciedgepreserveupsample/3228238-lumasigma?language=objc
func (e_ EdgePreserveUpsampleWrapper) SetLumaSigma(value float64) {
	objc.Call[objc.Void](e_, objc.Sel("setLumaSigma:"), value)
}

func (e_ EdgePreserveUpsampleWrapper) HasLumaSigma() bool {
	return e_.RespondsToSelector(objc.Sel("lumaSigma"))
}

// A value that specifies the influence of the input image’s luma information on the upsampling operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciedgepreserveupsample/3228238-lumasigma?language=objc
func (e_ EdgePreserveUpsampleWrapper) LumaSigma() float64 {
	rv := objc.Call[float64](e_, objc.Sel("lumaSigma"))
	return rv
}
