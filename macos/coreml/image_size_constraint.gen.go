// Code generated by DarwinKit. DO NOT EDIT.

package coreml

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [ImageSizeConstraint] class.
var ImageSizeConstraintClass = _ImageSizeConstraintClass{objc.GetClass("MLImageSizeConstraint")}

type _ImageSizeConstraintClass struct {
	objc.Class
}

// An interface definition for the [ImageSizeConstraint] class.
type IImageSizeConstraint interface {
	objc.IObject
	Type() ImageSizeConstraintType
	EnumeratedImageSizes() []ImageSize
	PixelsWideRange() foundation.Range
	PixelsHighRange() foundation.Range
}

// A list or range of sizes that augment an image constraint's default size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimagesizeconstraint?language=objc
type ImageSizeConstraint struct {
	objc.Object
}

func ImageSizeConstraintFrom(ptr unsafe.Pointer) ImageSizeConstraint {
	return ImageSizeConstraint{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ic _ImageSizeConstraintClass) Alloc() ImageSizeConstraint {
	rv := objc.Call[ImageSizeConstraint](ic, objc.Sel("alloc"))
	return rv
}

func (ic _ImageSizeConstraintClass) New() ImageSizeConstraint {
	rv := objc.Call[ImageSizeConstraint](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageSizeConstraint() ImageSizeConstraint {
	return ImageSizeConstraintClass.New()
}

func (i_ ImageSizeConstraint) Init() ImageSizeConstraint {
	rv := objc.Call[ImageSizeConstraint](i_, objc.Sel("init"))
	return rv
}

// Indicator of which properties to inspect for this image size constraint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimagesizeconstraint/2994307-type?language=objc
func (i_ ImageSizeConstraint) Type() ImageSizeConstraintType {
	rv := objc.Call[ImageSizeConstraintType](i_, objc.Sel("type"))
	return rv
}

// An array of image sizes a model's image feature accepts as input or produces as output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimagesizeconstraint/2994304-enumeratedimagesizes?language=objc
func (i_ ImageSizeConstraint) EnumeratedImageSizes() []ImageSize {
	rv := objc.Call[[]ImageSize](i_, objc.Sel("enumeratedImageSizes"))
	return rv
}

// The range of widths a model's image feature accepts as input or produces as output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimagesizeconstraint/2994306-pixelswiderange?language=objc
func (i_ ImageSizeConstraint) PixelsWideRange() foundation.Range {
	rv := objc.Call[foundation.Range](i_, objc.Sel("pixelsWideRange"))
	return rv
}

// The range of heights a model's image feature accepts as input or produces as output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimagesizeconstraint/2994305-pixelshighrange?language=objc
func (i_ ImageSizeConstraint) PixelsHighRange() foundation.Range {
	rv := objc.Call[foundation.Range](i_, objc.Sel("pixelsHighRange"))
	return rv
}
