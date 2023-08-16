// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a blend with mask filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendwithmask?language=objc
type PBlendWithMask interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetMaskImage(value Image)
	HasSetMaskImage() bool

	// optional
	MaskImage() IImage
	HasMaskImage() bool

	// optional
	SetBackgroundImage(value Image)
	HasSetBackgroundImage() bool

	// optional
	BackgroundImage() IImage
	HasBackgroundImage() bool
}

// A concrete type wrapper for the [PBlendWithMask] protocol.
type BlendWithMaskWrapper struct {
	objc.Object
}

func (b_ BlendWithMaskWrapper) HasSetInputImage() bool {
	return b_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as a foreground image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendwithmask/3228081-inputimage?language=objc
func (b_ BlendWithMaskWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](b_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (b_ BlendWithMaskWrapper) HasInputImage() bool {
	return b_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as a foreground image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendwithmask/3228081-inputimage?language=objc
func (b_ BlendWithMaskWrapper) InputImage() Image {
	rv := objc.Call[Image](b_, objc.Sel("inputImage"))
	return rv
}

func (b_ BlendWithMaskWrapper) HasSetMaskImage() bool {
	return b_.RespondsToSelector(objc.Sel("setMaskImage:"))
}

// A grayscale mask that defines the blend. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendwithmask/3228082-maskimage?language=objc
func (b_ BlendWithMaskWrapper) SetMaskImage(value IImage) {
	objc.Call[objc.Void](b_, objc.Sel("setMaskImage:"), objc.Ptr(value))
}

func (b_ BlendWithMaskWrapper) HasMaskImage() bool {
	return b_.RespondsToSelector(objc.Sel("maskImage"))
}

// A grayscale mask that defines the blend. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendwithmask/3228082-maskimage?language=objc
func (b_ BlendWithMaskWrapper) MaskImage() Image {
	rv := objc.Call[Image](b_, objc.Sel("maskImage"))
	return rv
}

func (b_ BlendWithMaskWrapper) HasSetBackgroundImage() bool {
	return b_.RespondsToSelector(objc.Sel("setBackgroundImage:"))
}

// The image to use as a background image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendwithmask/3228080-backgroundimage?language=objc
func (b_ BlendWithMaskWrapper) SetBackgroundImage(value IImage) {
	objc.Call[objc.Void](b_, objc.Sel("setBackgroundImage:"), objc.Ptr(value))
}

func (b_ BlendWithMaskWrapper) HasBackgroundImage() bool {
	return b_.RespondsToSelector(objc.Sel("backgroundImage"))
}

// The image to use as a background image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciblendwithmask/3228080-backgroundimage?language=objc
func (b_ BlendWithMaskWrapper) BackgroundImage() Image {
	rv := objc.Call[Image](b_, objc.Sel("backgroundImage"))
	return rv
}
