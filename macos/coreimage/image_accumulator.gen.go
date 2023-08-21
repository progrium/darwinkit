// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageAccumulator] class.
var ImageAccumulatorClass = _ImageAccumulatorClass{objc.GetClass("CIImageAccumulator")}

type _ImageAccumulatorClass struct {
	objc.Class
}

// An interface definition for the [ImageAccumulator] class.
type IImageAccumulator interface {
	objc.IObject
	SetImage(image IImage)
	Clear()
	Image() Image
	Extent() coregraphics.Rect
	Format() Format
}

// An object that manages feedback-based image processing for tasks such as painting or fluid simulation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageaccumulator?language=objc
type ImageAccumulator struct {
	objc.Object
}

func ImageAccumulatorFrom(ptr unsafe.Pointer) ImageAccumulator {
	return ImageAccumulator{
		Object: objc.ObjectFrom(ptr),
	}
}

func (i_ ImageAccumulator) InitWithExtentFormat(extent coregraphics.Rect, format Format) ImageAccumulator {
	rv := objc.Call[ImageAccumulator](i_, objc.Sel("initWithExtent:format:"), extent, format)
	return rv
}

// Initializes an image accumulator with the specified extent and pixel format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageaccumulator/1427718-initwithextent?language=objc
func NewImageAccumulatorWithExtentFormat(extent coregraphics.Rect, format Format) ImageAccumulator {
	instance := ImageAccumulatorClass.Alloc().InitWithExtentFormat(extent, format)
	instance.Autorelease()
	return instance
}

func (ic _ImageAccumulatorClass) ImageAccumulatorWithExtentFormat(extent coregraphics.Rect, format Format) ImageAccumulator {
	rv := objc.Call[ImageAccumulator](ic, objc.Sel("imageAccumulatorWithExtent:format:"), extent, format)
	return rv
}

// Creates an image accumulator with the specified extent and pixel format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageaccumulator/1427722-imageaccumulatorwithextent?language=objc
func ImageAccumulator_ImageAccumulatorWithExtentFormat(extent coregraphics.Rect, format Format) ImageAccumulator {
	return ImageAccumulatorClass.ImageAccumulatorWithExtentFormat(extent, format)
}

func (ic _ImageAccumulatorClass) Alloc() ImageAccumulator {
	rv := objc.Call[ImageAccumulator](ic, objc.Sel("alloc"))
	return rv
}

func ImageAccumulator_Alloc() ImageAccumulator {
	return ImageAccumulatorClass.Alloc()
}

func (ic _ImageAccumulatorClass) New() ImageAccumulator {
	rv := objc.Call[ImageAccumulator](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageAccumulator() ImageAccumulator {
	return ImageAccumulatorClass.New()
}

func (i_ ImageAccumulator) Init() ImageAccumulator {
	rv := objc.Call[ImageAccumulator](i_, objc.Sel("init"))
	return rv
}

// Sets the contents of the image accumulator to the contents of the specified image object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageaccumulator/1427702-setimage?language=objc
func (i_ ImageAccumulator) SetImage(image IImage) {
	objc.Call[objc.Void](i_, objc.Sel("setImage:"), objc.Ptr(image))
}

// Resets the accumulator, discarding any pending updates and the current content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageaccumulator/1427720-clear?language=objc
func (i_ ImageAccumulator) Clear() {
	objc.Call[objc.Void](i_, objc.Sel("clear"))
}

// Returns the current contents of the image accumulator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageaccumulator/1427704-image?language=objc
func (i_ ImageAccumulator) Image() Image {
	rv := objc.Call[Image](i_, objc.Sel("image"))
	return rv
}

// The extent of the image associated with the image accumulator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageaccumulator/1427714-extent?language=objc
func (i_ ImageAccumulator) Extent() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](i_, objc.Sel("extent"))
	return rv
}

// The pixel format of the image accumulator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageaccumulator/1427716-format?language=objc
func (i_ ImageAccumulator) Format() Format {
	rv := objc.Call[Format](i_, objc.Sel("format"))
	return rv
}
