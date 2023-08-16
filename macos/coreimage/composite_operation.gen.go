// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a composite operation filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicompositeoperation?language=objc
type PCompositeOperation interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetBackgroundImage(value Image)
	HasSetBackgroundImage() bool

	// optional
	BackgroundImage() IImage
	HasBackgroundImage() bool
}

// A concrete type wrapper for the [PCompositeOperation] protocol.
type CompositeOperationWrapper struct {
	objc.Object
}

func (c_ CompositeOperationWrapper) HasSetInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as a foreground image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicompositeoperation/3228183-inputimage?language=objc
func (c_ CompositeOperationWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](c_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (c_ CompositeOperationWrapper) HasInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as a foreground image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicompositeoperation/3228183-inputimage?language=objc
func (c_ CompositeOperationWrapper) InputImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("inputImage"))
	return rv
}

func (c_ CompositeOperationWrapper) HasSetBackgroundImage() bool {
	return c_.RespondsToSelector(objc.Sel("setBackgroundImage:"))
}

// The image to use as a background image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicompositeoperation/3228182-backgroundimage?language=objc
func (c_ CompositeOperationWrapper) SetBackgroundImage(value IImage) {
	objc.Call[objc.Void](c_, objc.Sel("setBackgroundImage:"), objc.Ptr(value))
}

func (c_ CompositeOperationWrapper) HasBackgroundImage() bool {
	return c_.RespondsToSelector(objc.Sel("backgroundImage"))
}

// The image to use as a background image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicompositeoperation/3228182-backgroundimage?language=objc
func (c_ CompositeOperationWrapper) BackgroundImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("backgroundImage"))
	return rv
}
