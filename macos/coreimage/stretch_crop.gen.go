// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistretchcrop?language=objc
type PStretchCrop interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetCropAmount(value float64)
	HasSetCropAmount() bool

	// optional
	CropAmount() float64
	HasCropAmount() bool

	// optional
	SetCenterStretchAmount(value float64)
	HasSetCenterStretchAmount() bool

	// optional
	CenterStretchAmount() float64
	HasCenterStretchAmount() bool

	// optional
	SetSize(value coregraphics.Point)
	HasSetSize() bool

	// optional
	Size() coregraphics.Point
	HasSize() bool
}

// A concrete type wrapper for the [PStretchCrop] protocol.
type StretchCropWrapper struct {
	objc.Object
}

func (s_ StretchCropWrapper) HasSetInputImage() bool {
	return s_.RespondsToSelector(objc.Sel("setInputImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistretchcrop/3600194-inputimage?language=objc
func (s_ StretchCropWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](s_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (s_ StretchCropWrapper) HasInputImage() bool {
	return s_.RespondsToSelector(objc.Sel("inputImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistretchcrop/3600194-inputimage?language=objc
func (s_ StretchCropWrapper) InputImage() Image {
	rv := objc.Call[Image](s_, objc.Sel("inputImage"))
	return rv
}

func (s_ StretchCropWrapper) HasSetCropAmount() bool {
	return s_.RespondsToSelector(objc.Sel("setCropAmount:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistretchcrop/3600193-cropamount?language=objc
func (s_ StretchCropWrapper) SetCropAmount(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setCropAmount:"), value)
}

func (s_ StretchCropWrapper) HasCropAmount() bool {
	return s_.RespondsToSelector(objc.Sel("cropAmount"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistretchcrop/3600193-cropamount?language=objc
func (s_ StretchCropWrapper) CropAmount() float64 {
	rv := objc.Call[float64](s_, objc.Sel("cropAmount"))
	return rv
}

func (s_ StretchCropWrapper) HasSetCenterStretchAmount() bool {
	return s_.RespondsToSelector(objc.Sel("setCenterStretchAmount:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistretchcrop/3600192-centerstretchamount?language=objc
func (s_ StretchCropWrapper) SetCenterStretchAmount(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setCenterStretchAmount:"), value)
}

func (s_ StretchCropWrapper) HasCenterStretchAmount() bool {
	return s_.RespondsToSelector(objc.Sel("centerStretchAmount"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistretchcrop/3600192-centerstretchamount?language=objc
func (s_ StretchCropWrapper) CenterStretchAmount() float64 {
	rv := objc.Call[float64](s_, objc.Sel("centerStretchAmount"))
	return rv
}

func (s_ StretchCropWrapper) HasSetSize() bool {
	return s_.RespondsToSelector(objc.Sel("setSize:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistretchcrop/3600195-size?language=objc
func (s_ StretchCropWrapper) SetSize(value coregraphics.Point) {
	objc.Call[objc.Void](s_, objc.Sel("setSize:"), value)
}

func (s_ StretchCropWrapper) HasSize() bool {
	return s_.RespondsToSelector(objc.Sel("size"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistretchcrop/3600195-size?language=objc
func (s_ StretchCropWrapper) Size() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](s_, objc.Sel("size"))
	return rv
}
