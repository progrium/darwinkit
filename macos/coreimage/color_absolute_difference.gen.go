// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorabsolutedifference?language=objc
type PColorAbsoluteDifference interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetInputImage2(value Image)
	HasSetInputImage2() bool

	// optional
	InputImage2() IImage
	HasInputImage2() bool
}

// A concrete type wrapper for the [PColorAbsoluteDifference] protocol.
type ColorAbsoluteDifferenceWrapper struct {
	objc.Object
}

func (c_ ColorAbsoluteDifferenceWrapper) HasSetInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("setInputImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorabsolutedifference/3547104-inputimage?language=objc
func (c_ ColorAbsoluteDifferenceWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](c_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (c_ ColorAbsoluteDifferenceWrapper) HasInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("inputImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorabsolutedifference/3547104-inputimage?language=objc
func (c_ ColorAbsoluteDifferenceWrapper) InputImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("inputImage"))
	return rv
}

func (c_ ColorAbsoluteDifferenceWrapper) HasSetInputImage2() bool {
	return c_.RespondsToSelector(objc.Sel("setInputImage2:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorabsolutedifference/3547105-inputimage2?language=objc
func (c_ ColorAbsoluteDifferenceWrapper) SetInputImage2(value IImage) {
	objc.Call[objc.Void](c_, objc.Sel("setInputImage2:"), objc.Ptr(value))
}

func (c_ ColorAbsoluteDifferenceWrapper) HasInputImage2() bool {
	return c_.RespondsToSelector(objc.Sel("inputImage2"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorabsolutedifference/3547105-inputimage2?language=objc
func (c_ ColorAbsoluteDifferenceWrapper) InputImage2() Image {
	rv := objc.Call[Image](c_, objc.Sel("inputImage2"))
	return rv
}
