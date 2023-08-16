// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a color cube with color space filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcubewithcolorspace?language=objc
type PColorCubeWithColorSpace interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetCubeDimension(value float64)
	HasSetCubeDimension() bool

	// optional
	CubeDimension() float64
	HasCubeDimension() bool

	// optional
	SetCubeData(value []byte)
	HasSetCubeData() bool

	// optional
	CubeData() []byte
	HasCubeData() bool

	// optional
	SetColorSpace(value coregraphics.ColorSpaceRef)
	HasSetColorSpace() bool

	// optional
	ColorSpace() coregraphics.ColorSpaceRef
	HasColorSpace() bool
}

// A concrete type wrapper for the [PColorCubeWithColorSpace] protocol.
type ColorCubeWithColorSpaceWrapper struct {
	objc.Object
}

func (c_ ColorCubeWithColorSpaceWrapper) HasSetInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcubewithcolorspace/3228141-inputimage?language=objc
func (c_ ColorCubeWithColorSpaceWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](c_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (c_ ColorCubeWithColorSpaceWrapper) HasInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcubewithcolorspace/3228141-inputimage?language=objc
func (c_ ColorCubeWithColorSpaceWrapper) InputImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("inputImage"))
	return rv
}

func (c_ ColorCubeWithColorSpaceWrapper) HasSetCubeDimension() bool {
	return c_.RespondsToSelector(objc.Sel("setCubeDimension:"))
}

// The length, in texels, of each side of the cube texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcubewithcolorspace/3228140-cubedimension?language=objc
func (c_ ColorCubeWithColorSpaceWrapper) SetCubeDimension(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setCubeDimension:"), value)
}

func (c_ ColorCubeWithColorSpaceWrapper) HasCubeDimension() bool {
	return c_.RespondsToSelector(objc.Sel("cubeDimension"))
}

// The length, in texels, of each side of the cube texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcubewithcolorspace/3228140-cubedimension?language=objc
func (c_ ColorCubeWithColorSpaceWrapper) CubeDimension() float64 {
	rv := objc.Call[float64](c_, objc.Sel("cubeDimension"))
	return rv
}

func (c_ ColorCubeWithColorSpaceWrapper) HasSetCubeData() bool {
	return c_.RespondsToSelector(objc.Sel("setCubeData:"))
}

// The cube texture data to use as a color lookup table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcubewithcolorspace/3228139-cubedata?language=objc
func (c_ ColorCubeWithColorSpaceWrapper) SetCubeData(value []byte) {
	objc.Call[objc.Void](c_, objc.Sel("setCubeData:"), value)
}

func (c_ ColorCubeWithColorSpaceWrapper) HasCubeData() bool {
	return c_.RespondsToSelector(objc.Sel("cubeData"))
}

// The cube texture data to use as a color lookup table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcubewithcolorspace/3228139-cubedata?language=objc
func (c_ ColorCubeWithColorSpaceWrapper) CubeData() []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("cubeData"))
	return rv
}

func (c_ ColorCubeWithColorSpaceWrapper) HasSetColorSpace() bool {
	return c_.RespondsToSelector(objc.Sel("setColorSpace:"))
}

// The working color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcubewithcolorspace/3228138-colorspace?language=objc
func (c_ ColorCubeWithColorSpaceWrapper) SetColorSpace(value coregraphics.ColorSpaceRef) {
	objc.Call[objc.Void](c_, objc.Sel("setColorSpace:"), value)
}

func (c_ ColorCubeWithColorSpaceWrapper) HasColorSpace() bool {
	return c_.RespondsToSelector(objc.Sel("colorSpace"))
}

// The working color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcubewithcolorspace/3228138-colorspace?language=objc
func (c_ ColorCubeWithColorSpaceWrapper) ColorSpace() coregraphics.ColorSpaceRef {
	rv := objc.Call[coregraphics.ColorSpaceRef](c_, objc.Sel("colorSpace"))
	return rv
}
