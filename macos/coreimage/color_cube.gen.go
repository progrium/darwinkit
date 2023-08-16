// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a color cube filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcube?language=objc
type PColorCube interface {
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
}

// A concrete type wrapper for the [PColorCube] protocol.
type ColorCubeWrapper struct {
	objc.Object
}

func (c_ ColorCubeWrapper) HasSetInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcube/3228136-inputimage?language=objc
func (c_ ColorCubeWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](c_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (c_ ColorCubeWrapper) HasInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcube/3228136-inputimage?language=objc
func (c_ ColorCubeWrapper) InputImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("inputImage"))
	return rv
}

func (c_ ColorCubeWrapper) HasSetCubeDimension() bool {
	return c_.RespondsToSelector(objc.Sel("setCubeDimension:"))
}

// The length, in texels, of each side of the cube texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcube/3228135-cubedimension?language=objc
func (c_ ColorCubeWrapper) SetCubeDimension(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setCubeDimension:"), value)
}

func (c_ ColorCubeWrapper) HasCubeDimension() bool {
	return c_.RespondsToSelector(objc.Sel("cubeDimension"))
}

// The length, in texels, of each side of the cube texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcube/3228135-cubedimension?language=objc
func (c_ ColorCubeWrapper) CubeDimension() float64 {
	rv := objc.Call[float64](c_, objc.Sel("cubeDimension"))
	return rv
}

func (c_ ColorCubeWrapper) HasSetCubeData() bool {
	return c_.RespondsToSelector(objc.Sel("setCubeData:"))
}

// The cube texture data to use as a color lookup table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcube/3228134-cubedata?language=objc
func (c_ ColorCubeWrapper) SetCubeData(value []byte) {
	objc.Call[objc.Void](c_, objc.Sel("setCubeData:"), value)
}

func (c_ ColorCubeWrapper) HasCubeData() bool {
	return c_.RespondsToSelector(objc.Sel("cubeData"))
}

// The cube texture data to use as a color lookup table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcube/3228134-cubedata?language=objc
func (c_ ColorCubeWrapper) CubeData() []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("cubeData"))
	return rv
}
