// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a color polynomial filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorpolynomial?language=objc
type PColorPolynomial interface {
	// optional
	SetRedCoefficients(value Vector)
	HasSetRedCoefficients() bool

	// optional
	RedCoefficients() IVector
	HasRedCoefficients() bool

	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetGreenCoefficients(value Vector)
	HasSetGreenCoefficients() bool

	// optional
	GreenCoefficients() IVector
	HasGreenCoefficients() bool

	// optional
	SetBlueCoefficients(value Vector)
	HasSetBlueCoefficients() bool

	// optional
	BlueCoefficients() IVector
	HasBlueCoefficients() bool

	// optional
	SetAlphaCoefficients(value Vector)
	HasSetAlphaCoefficients() bool

	// optional
	AlphaCoefficients() IVector
	HasAlphaCoefficients() bool
}

// A concrete type wrapper for the [PColorPolynomial] protocol.
type ColorPolynomialWrapper struct {
	objc.Object
}

func (c_ ColorPolynomialWrapper) HasSetRedCoefficients() bool {
	return c_.RespondsToSelector(objc.Sel("setRedCoefficients:"))
}

// Polynomial coefficients for the red channel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorpolynomial/3228175-redcoefficients?language=objc
func (c_ ColorPolynomialWrapper) SetRedCoefficients(value IVector) {
	objc.Call[objc.Void](c_, objc.Sel("setRedCoefficients:"), objc.Ptr(value))
}

func (c_ ColorPolynomialWrapper) HasRedCoefficients() bool {
	return c_.RespondsToSelector(objc.Sel("redCoefficients"))
}

// Polynomial coefficients for the red channel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorpolynomial/3228175-redcoefficients?language=objc
func (c_ ColorPolynomialWrapper) RedCoefficients() Vector {
	rv := objc.Call[Vector](c_, objc.Sel("redCoefficients"))
	return rv
}

func (c_ ColorPolynomialWrapper) HasSetInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorpolynomial/3228174-inputimage?language=objc
func (c_ ColorPolynomialWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](c_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (c_ ColorPolynomialWrapper) HasInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorpolynomial/3228174-inputimage?language=objc
func (c_ ColorPolynomialWrapper) InputImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("inputImage"))
	return rv
}

func (c_ ColorPolynomialWrapper) HasSetGreenCoefficients() bool {
	return c_.RespondsToSelector(objc.Sel("setGreenCoefficients:"))
}

// Polynomial coefficients for the green channel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorpolynomial/3228173-greencoefficients?language=objc
func (c_ ColorPolynomialWrapper) SetGreenCoefficients(value IVector) {
	objc.Call[objc.Void](c_, objc.Sel("setGreenCoefficients:"), objc.Ptr(value))
}

func (c_ ColorPolynomialWrapper) HasGreenCoefficients() bool {
	return c_.RespondsToSelector(objc.Sel("greenCoefficients"))
}

// Polynomial coefficients for the green channel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorpolynomial/3228173-greencoefficients?language=objc
func (c_ ColorPolynomialWrapper) GreenCoefficients() Vector {
	rv := objc.Call[Vector](c_, objc.Sel("greenCoefficients"))
	return rv
}

func (c_ ColorPolynomialWrapper) HasSetBlueCoefficients() bool {
	return c_.RespondsToSelector(objc.Sel("setBlueCoefficients:"))
}

// Polynomial coefficients for the blue channel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorpolynomial/3228172-bluecoefficients?language=objc
func (c_ ColorPolynomialWrapper) SetBlueCoefficients(value IVector) {
	objc.Call[objc.Void](c_, objc.Sel("setBlueCoefficients:"), objc.Ptr(value))
}

func (c_ ColorPolynomialWrapper) HasBlueCoefficients() bool {
	return c_.RespondsToSelector(objc.Sel("blueCoefficients"))
}

// Polynomial coefficients for the blue channel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorpolynomial/3228172-bluecoefficients?language=objc
func (c_ ColorPolynomialWrapper) BlueCoefficients() Vector {
	rv := objc.Call[Vector](c_, objc.Sel("blueCoefficients"))
	return rv
}

func (c_ ColorPolynomialWrapper) HasSetAlphaCoefficients() bool {
	return c_.RespondsToSelector(objc.Sel("setAlphaCoefficients:"))
}

// Polynomial coefficients for the alpha channel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorpolynomial/3228171-alphacoefficients?language=objc
func (c_ ColorPolynomialWrapper) SetAlphaCoefficients(value IVector) {
	objc.Call[objc.Void](c_, objc.Sel("setAlphaCoefficients:"), objc.Ptr(value))
}

func (c_ ColorPolynomialWrapper) HasAlphaCoefficients() bool {
	return c_.RespondsToSelector(objc.Sel("alphaCoefficients"))
}

// Polynomial coefficients for the alpha channel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorpolynomial/3228171-alphacoefficients?language=objc
func (c_ ColorPolynomialWrapper) AlphaCoefficients() Vector {
	rv := objc.Call[Vector](c_, objc.Sel("alphaCoefficients"))
	return rv
}
