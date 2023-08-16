// AUTO-GENERATED CODE, DO NOT MODIFY

package coreml

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageSize] class.
var ImageSizeClass = _ImageSizeClass{objc.GetClass("MLImageSize")}

type _ImageSizeClass struct {
	objc.Class
}

// An interface definition for the [ImageSize] class.
type IImageSize interface {
	objc.IObject
	PixelsWide() int
	PixelsHigh() int
}

// The width and height of an image feature size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimagesize?language=objc
type ImageSize struct {
	objc.Object
}

func ImageSizeFrom(ptr unsafe.Pointer) ImageSize {
	return ImageSize{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ic _ImageSizeClass) Alloc() ImageSize {
	rv := objc.Call[ImageSize](ic, objc.Sel("alloc"))
	return rv
}

func ImageSize_Alloc() ImageSize {
	return ImageSizeClass.Alloc()
}

func (ic _ImageSizeClass) New() ImageSize {
	rv := objc.Call[ImageSize](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageSize() ImageSize {
	return ImageSizeClass.New()
}

func (i_ ImageSize) Init() ImageSize {
	rv := objc.Call[ImageSize](i_, objc.Sel("init"))
	return rv
}

// The width of an image feature in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimagesize/2994302-pixelswide?language=objc
func (i_ ImageSize) PixelsWide() int {
	rv := objc.Call[int](i_, objc.Sel("pixelsWide"))
	return rv
}

// The height of an image feature in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimagesize/2994301-pixelshigh?language=objc
func (i_ ImageSize) PixelsHigh() int {
	rv := objc.Call[int](i_, objc.Sel("pixelsHigh"))
	return rv
}
