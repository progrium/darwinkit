// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a color curves filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcurves?language=objc
type PColorCurves interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetColorSpace(value coregraphics.ColorSpaceRef)
	HasSetColorSpace() bool

	// optional
	ColorSpace() coregraphics.ColorSpaceRef
	HasColorSpace() bool

	// optional
	SetCurvesData(value []byte)
	HasSetCurvesData() bool

	// optional
	CurvesData() []byte
	HasCurvesData() bool

	// optional
	SetCurvesDomain(value Vector)
	HasSetCurvesDomain() bool

	// optional
	CurvesDomain() IVector
	HasCurvesDomain() bool
}

// A concrete type wrapper for the [PColorCurves] protocol.
type ColorCurvesWrapper struct {
	objc.Object
}

func (c_ ColorCurvesWrapper) HasSetInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcurves/3228153-inputimage?language=objc
func (c_ ColorCurvesWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](c_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (c_ ColorCurvesWrapper) HasInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcurves/3228153-inputimage?language=objc
func (c_ ColorCurvesWrapper) InputImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("inputImage"))
	return rv
}

func (c_ ColorCurvesWrapper) HasSetColorSpace() bool {
	return c_.RespondsToSelector(objc.Sel("setColorSpace:"))
}

// The working color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcurves/3228150-colorspace?language=objc
func (c_ ColorCurvesWrapper) SetColorSpace(value coregraphics.ColorSpaceRef) {
	objc.Call[objc.Void](c_, objc.Sel("setColorSpace:"), value)
}

func (c_ ColorCurvesWrapper) HasColorSpace() bool {
	return c_.RespondsToSelector(objc.Sel("colorSpace"))
}

// The working color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcurves/3228150-colorspace?language=objc
func (c_ ColorCurvesWrapper) ColorSpace() coregraphics.ColorSpaceRef {
	rv := objc.Call[coregraphics.ColorSpaceRef](c_, objc.Sel("colorSpace"))
	return rv
}

func (c_ ColorCurvesWrapper) HasSetCurvesData() bool {
	return c_.RespondsToSelector(objc.Sel("setCurvesData:"))
}

// Color values that determine the color curves transform. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcurves/3228151-curvesdata?language=objc
func (c_ ColorCurvesWrapper) SetCurvesData(value []byte) {
	objc.Call[objc.Void](c_, objc.Sel("setCurvesData:"), value)
}

func (c_ ColorCurvesWrapper) HasCurvesData() bool {
	return c_.RespondsToSelector(objc.Sel("curvesData"))
}

// Color values that determine the color curves transform. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcurves/3228151-curvesdata?language=objc
func (c_ ColorCurvesWrapper) CurvesData() []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("curvesData"))
	return rv
}

func (c_ ColorCurvesWrapper) HasSetCurvesDomain() bool {
	return c_.RespondsToSelector(objc.Sel("setCurvesDomain:"))
}

// A two-element vector that defines the minimum and maximum values of the curve data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcurves/3228152-curvesdomain?language=objc
func (c_ ColorCurvesWrapper) SetCurvesDomain(value IVector) {
	objc.Call[objc.Void](c_, objc.Sel("setCurvesDomain:"), objc.Ptr(value))
}

func (c_ ColorCurvesWrapper) HasCurvesDomain() bool {
	return c_.RespondsToSelector(objc.Sel("curvesDomain"))
}

// A two-element vector that defines the minimum and maximum values of the curve data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcurves/3228152-curvesdomain?language=objc
func (c_ ColorCurvesWrapper) CurvesDomain() Vector {
	rv := objc.Call[Vector](c_, objc.Sel("curvesDomain"))
	return rv
}
