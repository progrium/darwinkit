// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a color clamp filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorclamp?language=objc
type PColorClamp interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetMaxComponents(value Vector)
	HasSetMaxComponents() bool

	// optional
	MaxComponents() IVector
	HasMaxComponents() bool

	// optional
	SetMinComponents(value Vector)
	HasSetMinComponents() bool

	// optional
	MinComponents() IVector
	HasMinComponents() bool
}

// A concrete type wrapper for the [PColorClamp] protocol.
type ColorClampWrapper struct {
	objc.Object
}

func (c_ ColorClampWrapper) HasSetInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorclamp/3228120-inputimage?language=objc
func (c_ ColorClampWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](c_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (c_ ColorClampWrapper) HasInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorclamp/3228120-inputimage?language=objc
func (c_ ColorClampWrapper) InputImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("inputImage"))
	return rv
}

func (c_ ColorClampWrapper) HasSetMaxComponents() bool {
	return c_.RespondsToSelector(objc.Sel("setMaxComponents:"))
}

// A vector containing the higher clamping values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorclamp/3228121-maxcomponents?language=objc
func (c_ ColorClampWrapper) SetMaxComponents(value IVector) {
	objc.Call[objc.Void](c_, objc.Sel("setMaxComponents:"), objc.Ptr(value))
}

func (c_ ColorClampWrapper) HasMaxComponents() bool {
	return c_.RespondsToSelector(objc.Sel("maxComponents"))
}

// A vector containing the higher clamping values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorclamp/3228121-maxcomponents?language=objc
func (c_ ColorClampWrapper) MaxComponents() Vector {
	rv := objc.Call[Vector](c_, objc.Sel("maxComponents"))
	return rv
}

func (c_ ColorClampWrapper) HasSetMinComponents() bool {
	return c_.RespondsToSelector(objc.Sel("setMinComponents:"))
}

// A vector containing the lower clamping values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorclamp/3228122-mincomponents?language=objc
func (c_ ColorClampWrapper) SetMinComponents(value IVector) {
	objc.Call[objc.Void](c_, objc.Sel("setMinComponents:"), objc.Ptr(value))
}

func (c_ ColorClampWrapper) HasMinComponents() bool {
	return c_.RespondsToSelector(objc.Sel("minComponents"))
}

// A vector containing the lower clamping values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorclamp/3228122-mincomponents?language=objc
func (c_ ColorClampWrapper) MinComponents() Vector {
	rv := objc.Call[Vector](c_, objc.Sel("minComponents"))
	return rv
}
