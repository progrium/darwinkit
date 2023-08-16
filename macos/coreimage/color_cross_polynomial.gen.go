// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a color cross-polynomial filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcrosspolynomial?language=objc
type PColorCrossPolynomial interface {
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
}

// A concrete type wrapper for the [PColorCrossPolynomial] protocol.
type ColorCrossPolynomialWrapper struct {
	objc.Object
}

func (c_ ColorCrossPolynomialWrapper) HasSetRedCoefficients() bool {
	return c_.RespondsToSelector(objc.Sel("setRedCoefficients:"))
}

// Polynomial coefficients for the red channel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcrosspolynomial/3228132-redcoefficients?language=objc
func (c_ ColorCrossPolynomialWrapper) SetRedCoefficients(value IVector) {
	objc.Call[objc.Void](c_, objc.Sel("setRedCoefficients:"), objc.Ptr(value))
}

func (c_ ColorCrossPolynomialWrapper) HasRedCoefficients() bool {
	return c_.RespondsToSelector(objc.Sel("redCoefficients"))
}

// Polynomial coefficients for the red channel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcrosspolynomial/3228132-redcoefficients?language=objc
func (c_ ColorCrossPolynomialWrapper) RedCoefficients() Vector {
	rv := objc.Call[Vector](c_, objc.Sel("redCoefficients"))
	return rv
}

func (c_ ColorCrossPolynomialWrapper) HasSetInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcrosspolynomial/3228131-inputimage?language=objc
func (c_ ColorCrossPolynomialWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](c_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (c_ ColorCrossPolynomialWrapper) HasInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcrosspolynomial/3228131-inputimage?language=objc
func (c_ ColorCrossPolynomialWrapper) InputImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("inputImage"))
	return rv
}

func (c_ ColorCrossPolynomialWrapper) HasSetGreenCoefficients() bool {
	return c_.RespondsToSelector(objc.Sel("setGreenCoefficients:"))
}

// Polynomial coefficients for the green channel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcrosspolynomial/3228130-greencoefficients?language=objc
func (c_ ColorCrossPolynomialWrapper) SetGreenCoefficients(value IVector) {
	objc.Call[objc.Void](c_, objc.Sel("setGreenCoefficients:"), objc.Ptr(value))
}

func (c_ ColorCrossPolynomialWrapper) HasGreenCoefficients() bool {
	return c_.RespondsToSelector(objc.Sel("greenCoefficients"))
}

// Polynomial coefficients for the green channel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcrosspolynomial/3228130-greencoefficients?language=objc
func (c_ ColorCrossPolynomialWrapper) GreenCoefficients() Vector {
	rv := objc.Call[Vector](c_, objc.Sel("greenCoefficients"))
	return rv
}

func (c_ ColorCrossPolynomialWrapper) HasSetBlueCoefficients() bool {
	return c_.RespondsToSelector(objc.Sel("setBlueCoefficients:"))
}

// Polynomial coefficients for the blue channel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcrosspolynomial/3228129-bluecoefficients?language=objc
func (c_ ColorCrossPolynomialWrapper) SetBlueCoefficients(value IVector) {
	objc.Call[objc.Void](c_, objc.Sel("setBlueCoefficients:"), objc.Ptr(value))
}

func (c_ ColorCrossPolynomialWrapper) HasBlueCoefficients() bool {
	return c_.RespondsToSelector(objc.Sel("blueCoefficients"))
}

// Polynomial coefficients for the blue channel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcrosspolynomial/3228129-bluecoefficients?language=objc
func (c_ ColorCrossPolynomialWrapper) BlueCoefficients() Vector {
	rv := objc.Call[Vector](c_, objc.Sel("blueCoefficients"))
	return rv
}
