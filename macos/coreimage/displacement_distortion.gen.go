// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidisplacementdistortion?language=objc
type PDisplacementDistortion interface {
	// optional
	SetScale(value float64)
	HasSetScale() bool

	// optional
	Scale() float64
	HasScale() bool

	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetDisplacementImage(value Image)
	HasSetDisplacementImage() bool

	// optional
	DisplacementImage() IImage
	HasDisplacementImage() bool
}

// A concrete type wrapper for the [PDisplacementDistortion] protocol.
type DisplacementDistortionWrapper struct {
	objc.Object
}

func (d_ DisplacementDistortionWrapper) HasSetScale() bool {
	return d_.RespondsToSelector(objc.Sel("setScale:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidisplacementdistortion/3600129-scale?language=objc
func (d_ DisplacementDistortionWrapper) SetScale(value float64) {
	objc.Call[objc.Void](d_, objc.Sel("setScale:"), value)
}

func (d_ DisplacementDistortionWrapper) HasScale() bool {
	return d_.RespondsToSelector(objc.Sel("scale"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidisplacementdistortion/3600129-scale?language=objc
func (d_ DisplacementDistortionWrapper) Scale() float64 {
	rv := objc.Call[float64](d_, objc.Sel("scale"))
	return rv
}

func (d_ DisplacementDistortionWrapper) HasSetInputImage() bool {
	return d_.RespondsToSelector(objc.Sel("setInputImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidisplacementdistortion/3600128-inputimage?language=objc
func (d_ DisplacementDistortionWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](d_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (d_ DisplacementDistortionWrapper) HasInputImage() bool {
	return d_.RespondsToSelector(objc.Sel("inputImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidisplacementdistortion/3600128-inputimage?language=objc
func (d_ DisplacementDistortionWrapper) InputImage() Image {
	rv := objc.Call[Image](d_, objc.Sel("inputImage"))
	return rv
}

func (d_ DisplacementDistortionWrapper) HasSetDisplacementImage() bool {
	return d_.RespondsToSelector(objc.Sel("setDisplacementImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidisplacementdistortion/3600127-displacementimage?language=objc
func (d_ DisplacementDistortionWrapper) SetDisplacementImage(value IImage) {
	objc.Call[objc.Void](d_, objc.Sel("setDisplacementImage:"), objc.Ptr(value))
}

func (d_ DisplacementDistortionWrapper) HasDisplacementImage() bool {
	return d_.RespondsToSelector(objc.Sel("displacementImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidisplacementdistortion/3600127-displacementimage?language=objc
func (d_ DisplacementDistortionWrapper) DisplacementImage() Image {
	rv := objc.Call[Image](d_, objc.Sel("displacementImage"))
	return rv
}
