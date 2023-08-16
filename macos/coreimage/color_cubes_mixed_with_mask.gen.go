// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a color cube mixed with mask filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcubesmixedwithmask?language=objc
type PColorCubesMixedWithMask interface {
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
	SetMaskImage(value Image)
	HasSetMaskImage() bool

	// optional
	MaskImage() IImage
	HasMaskImage() bool

	// optional
	SetColorSpace(value coregraphics.ColorSpaceRef)
	HasSetColorSpace() bool

	// optional
	ColorSpace() coregraphics.ColorSpaceRef
	HasColorSpace() bool

	// optional
	SetCube0Data(value []byte)
	HasSetCube0Data() bool

	// optional
	Cube0Data() []byte
	HasCube0Data() bool

	// optional
	SetCube1Data(value []byte)
	HasSetCube1Data() bool

	// optional
	Cube1Data() []byte
	HasCube1Data() bool
}

// A concrete type wrapper for the [PColorCubesMixedWithMask] protocol.
type ColorCubesMixedWithMaskWrapper struct {
	objc.Object
}

func (c_ ColorCubesMixedWithMaskWrapper) HasSetInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcubesmixedwithmask/3228147-inputimage?language=objc
func (c_ ColorCubesMixedWithMaskWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](c_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (c_ ColorCubesMixedWithMaskWrapper) HasInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcubesmixedwithmask/3228147-inputimage?language=objc
func (c_ ColorCubesMixedWithMaskWrapper) InputImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("inputImage"))
	return rv
}

func (c_ ColorCubesMixedWithMaskWrapper) HasSetCubeDimension() bool {
	return c_.RespondsToSelector(objc.Sel("setCubeDimension:"))
}

// The length, in texels, of each side of the cube texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcubesmixedwithmask/3228146-cubedimension?language=objc
func (c_ ColorCubesMixedWithMaskWrapper) SetCubeDimension(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setCubeDimension:"), value)
}

func (c_ ColorCubesMixedWithMaskWrapper) HasCubeDimension() bool {
	return c_.RespondsToSelector(objc.Sel("cubeDimension"))
}

// The length, in texels, of each side of the cube texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcubesmixedwithmask/3228146-cubedimension?language=objc
func (c_ ColorCubesMixedWithMaskWrapper) CubeDimension() float64 {
	rv := objc.Call[float64](c_, objc.Sel("cubeDimension"))
	return rv
}

func (c_ ColorCubesMixedWithMaskWrapper) HasSetMaskImage() bool {
	return c_.RespondsToSelector(objc.Sel("setMaskImage:"))
}

// A masking image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcubesmixedwithmask/3228148-maskimage?language=objc
func (c_ ColorCubesMixedWithMaskWrapper) SetMaskImage(value IImage) {
	objc.Call[objc.Void](c_, objc.Sel("setMaskImage:"), objc.Ptr(value))
}

func (c_ ColorCubesMixedWithMaskWrapper) HasMaskImage() bool {
	return c_.RespondsToSelector(objc.Sel("maskImage"))
}

// A masking image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcubesmixedwithmask/3228148-maskimage?language=objc
func (c_ ColorCubesMixedWithMaskWrapper) MaskImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("maskImage"))
	return rv
}

func (c_ ColorCubesMixedWithMaskWrapper) HasSetColorSpace() bool {
	return c_.RespondsToSelector(objc.Sel("setColorSpace:"))
}

// The working color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcubesmixedwithmask/3228143-colorspace?language=objc
func (c_ ColorCubesMixedWithMaskWrapper) SetColorSpace(value coregraphics.ColorSpaceRef) {
	objc.Call[objc.Void](c_, objc.Sel("setColorSpace:"), value)
}

func (c_ ColorCubesMixedWithMaskWrapper) HasColorSpace() bool {
	return c_.RespondsToSelector(objc.Sel("colorSpace"))
}

// The working color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcubesmixedwithmask/3228143-colorspace?language=objc
func (c_ ColorCubesMixedWithMaskWrapper) ColorSpace() coregraphics.ColorSpaceRef {
	rv := objc.Call[coregraphics.ColorSpaceRef](c_, objc.Sel("colorSpace"))
	return rv
}

func (c_ ColorCubesMixedWithMaskWrapper) HasSetCube0Data() bool {
	return c_.RespondsToSelector(objc.Sel("setCube0Data:"))
}

// The cube texture data to use as a color lookup table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcubesmixedwithmask/3228144-cube0data?language=objc
func (c_ ColorCubesMixedWithMaskWrapper) SetCube0Data(value []byte) {
	objc.Call[objc.Void](c_, objc.Sel("setCube0Data:"), value)
}

func (c_ ColorCubesMixedWithMaskWrapper) HasCube0Data() bool {
	return c_.RespondsToSelector(objc.Sel("cube0Data"))
}

// The cube texture data to use as a color lookup table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcubesmixedwithmask/3228144-cube0data?language=objc
func (c_ ColorCubesMixedWithMaskWrapper) Cube0Data() []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("cube0Data"))
	return rv
}

func (c_ ColorCubesMixedWithMaskWrapper) HasSetCube1Data() bool {
	return c_.RespondsToSelector(objc.Sel("setCube1Data:"))
}

// The cube texture data to use as a color lookup table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcubesmixedwithmask/3228145-cube1data?language=objc
func (c_ ColorCubesMixedWithMaskWrapper) SetCube1Data(value []byte) {
	objc.Call[objc.Void](c_, objc.Sel("setCube1Data:"), value)
}

func (c_ ColorCubesMixedWithMaskWrapper) HasCube1Data() bool {
	return c_.RespondsToSelector(objc.Sel("cube1Data"))
}

// The cube texture data to use as a color lookup table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcubesmixedwithmask/3228145-cube1data?language=objc
func (c_ ColorCubesMixedWithMaskWrapper) Cube1Data() []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("cube1Data"))
	return rv
}
