// AUTO-GENERATED CODE, DO NOT MODIFY

package coreml

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageConstraint] class.
var ImageConstraintClass = _ImageConstraintClass{objc.GetClass("MLImageConstraint")}

type _ImageConstraintClass struct {
	objc.Class
}

// An interface definition for the [ImageConstraint] class.
type IImageConstraint interface {
	objc.IObject
	SizeConstraint() ImageSizeConstraint
	PixelFormatType() uint
	PixelsWide() int
	PixelsHigh() int
}

// The width, height, and pixel format constraints of an image feature. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimageconstraint?language=objc
type ImageConstraint struct {
	objc.Object
}

func ImageConstraintFrom(ptr unsafe.Pointer) ImageConstraint {
	return ImageConstraint{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ic _ImageConstraintClass) Alloc() ImageConstraint {
	rv := objc.Call[ImageConstraint](ic, objc.Sel("alloc"))
	return rv
}

func ImageConstraint_Alloc() ImageConstraint {
	return ImageConstraintClass.Alloc()
}

func (ic _ImageConstraintClass) New() ImageConstraint {
	rv := objc.Call[ImageConstraint](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageConstraint() ImageConstraint {
	return ImageConstraintClass.New()
}

func (i_ ImageConstraint) Init() ImageConstraint {
	rv := objc.Call[ImageConstraint](i_, objc.Sel("init"))
	return rv
}

// Additional sizes this image feature supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimageconstraint/2994299-sizeconstraint?language=objc
func (i_ ImageConstraint) SizeConstraint() ImageSizeConstraint {
	rv := objc.Call[ImageSizeConstraint](i_, objc.Sel("sizeConstraint"))
	return rv
}

// The model's pixel format for an image feature. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimageconstraint/2921267-pixelformattype?language=objc
func (i_ ImageConstraint) PixelFormatType() uint {
	rv := objc.Call[uint](i_, objc.Sel("pixelFormatType"))
	return rv
}

// The model's default width for an image feature. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimageconstraint/2921270-pixelswide?language=objc
func (i_ ImageConstraint) PixelsWide() int {
	rv := objc.Call[int](i_, objc.Sel("pixelsWide"))
	return rv
}

// The model's default height for an image feature. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimageconstraint/2921268-pixelshigh?language=objc
func (i_ ImageConstraint) PixelsHigh() int {
	rv := objc.Call[int](i_, objc.Sel("pixelsHigh"))
	return rv
}
