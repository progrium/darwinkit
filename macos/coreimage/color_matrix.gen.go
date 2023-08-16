// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a color matrix filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolormatrix?language=objc
type PColorMatrix interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetGVector(value Vector)
	HasSetGVector() bool

	// optional
	GVector() IVector
	HasGVector() bool

	// optional
	SetBiasVector(value Vector)
	HasSetBiasVector() bool

	// optional
	BiasVector() IVector
	HasBiasVector() bool

	// optional
	SetBVector(value Vector)
	HasSetBVector() bool

	// optional
	BVector() IVector
	HasBVector() bool

	// optional
	SetRVector(value Vector)
	HasSetRVector() bool

	// optional
	RVector() IVector
	HasRVector() bool

	// optional
	SetAVector(value Vector)
	HasSetAVector() bool

	// optional
	AVector() IVector
	HasAVector() bool
}

// A concrete type wrapper for the [PColorMatrix] protocol.
type ColorMatrixWrapper struct {
	objc.Object
}

func (c_ ColorMatrixWrapper) HasSetInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolormatrix/3228165-inputimage?language=objc
func (c_ ColorMatrixWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](c_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (c_ ColorMatrixWrapper) HasInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolormatrix/3228165-inputimage?language=objc
func (c_ ColorMatrixWrapper) InputImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("inputImage"))
	return rv
}

func (c_ ColorMatrixWrapper) HasSetGVector() bool {
	return c_.RespondsToSelector(objc.Sel("setGVector:"))
}

// The amount of green to multiply the source color values by. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolormatrix/3228162-gvector?language=objc
func (c_ ColorMatrixWrapper) SetGVector(value IVector) {
	objc.Call[objc.Void](c_, objc.Sel("setGVector:"), objc.Ptr(value))
}

func (c_ ColorMatrixWrapper) HasGVector() bool {
	return c_.RespondsToSelector(objc.Sel("GVector"))
}

// The amount of green to multiply the source color values by. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolormatrix/3228162-gvector?language=objc
func (c_ ColorMatrixWrapper) GVector() Vector {
	rv := objc.Call[Vector](c_, objc.Sel("GVector"))
	return rv
}

func (c_ ColorMatrixWrapper) HasSetBiasVector() bool {
	return c_.RespondsToSelector(objc.Sel("setBiasVector:"))
}

// A vector that’s added to each color component. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolormatrix/3228164-biasvector?language=objc
func (c_ ColorMatrixWrapper) SetBiasVector(value IVector) {
	objc.Call[objc.Void](c_, objc.Sel("setBiasVector:"), objc.Ptr(value))
}

func (c_ ColorMatrixWrapper) HasBiasVector() bool {
	return c_.RespondsToSelector(objc.Sel("biasVector"))
}

// A vector that’s added to each color component. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolormatrix/3228164-biasvector?language=objc
func (c_ ColorMatrixWrapper) BiasVector() Vector {
	rv := objc.Call[Vector](c_, objc.Sel("biasVector"))
	return rv
}

func (c_ ColorMatrixWrapper) HasSetBVector() bool {
	return c_.RespondsToSelector(objc.Sel("setBVector:"))
}

// The amount of blue to multiply the source color values by. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolormatrix/3228161-bvector?language=objc
func (c_ ColorMatrixWrapper) SetBVector(value IVector) {
	objc.Call[objc.Void](c_, objc.Sel("setBVector:"), objc.Ptr(value))
}

func (c_ ColorMatrixWrapper) HasBVector() bool {
	return c_.RespondsToSelector(objc.Sel("BVector"))
}

// The amount of blue to multiply the source color values by. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolormatrix/3228161-bvector?language=objc
func (c_ ColorMatrixWrapper) BVector() Vector {
	rv := objc.Call[Vector](c_, objc.Sel("BVector"))
	return rv
}

func (c_ ColorMatrixWrapper) HasSetRVector() bool {
	return c_.RespondsToSelector(objc.Sel("setRVector:"))
}

// The amount of red to multiply the source color values by. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolormatrix/3228163-rvector?language=objc
func (c_ ColorMatrixWrapper) SetRVector(value IVector) {
	objc.Call[objc.Void](c_, objc.Sel("setRVector:"), objc.Ptr(value))
}

func (c_ ColorMatrixWrapper) HasRVector() bool {
	return c_.RespondsToSelector(objc.Sel("RVector"))
}

// The amount of red to multiply the source color values by. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolormatrix/3228163-rvector?language=objc
func (c_ ColorMatrixWrapper) RVector() Vector {
	rv := objc.Call[Vector](c_, objc.Sel("RVector"))
	return rv
}

func (c_ ColorMatrixWrapper) HasSetAVector() bool {
	return c_.RespondsToSelector(objc.Sel("setAVector:"))
}

// The amount of alpha to multiply the source color values by. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolormatrix/3228160-avector?language=objc
func (c_ ColorMatrixWrapper) SetAVector(value IVector) {
	objc.Call[objc.Void](c_, objc.Sel("setAVector:"), objc.Ptr(value))
}

func (c_ ColorMatrixWrapper) HasAVector() bool {
	return c_.RespondsToSelector(objc.Sel("AVector"))
}

// The amount of alpha to multiply the source color values by. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolormatrix/3228160-avector?language=objc
func (c_ ColorMatrixWrapper) AVector() Vector {
	rv := objc.Call[Vector](c_, objc.Sel("AVector"))
	return rv
}
