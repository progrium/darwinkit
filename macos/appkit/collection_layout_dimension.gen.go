// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CollectionLayoutDimension] class.
var CollectionLayoutDimensionClass = _CollectionLayoutDimensionClass{objc.GetClass("NSCollectionLayoutDimension")}

type _CollectionLayoutDimensionClass struct {
	objc.Class
}

// An interface definition for the [CollectionLayoutDimension] class.
type ICollectionLayoutDimension interface {
	objc.IObject
	IsEstimated() bool
	Dimension() float64
	IsFractionalWidth() bool
	IsAbsolute() bool
	IsFractionalHeight() bool
}

// An individual dimension representing an item’s width or height in a collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutdimension?language=objc
type CollectionLayoutDimension struct {
	objc.Object
}

func CollectionLayoutDimensionFrom(ptr unsafe.Pointer) CollectionLayoutDimension {
	return CollectionLayoutDimension{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CollectionLayoutDimensionClass) FractionalHeightDimension(fractionalHeight float64) CollectionLayoutDimension {
	rv := objc.Call[CollectionLayoutDimension](cc, objc.Sel("fractionalHeightDimension:"), fractionalHeight)
	return rv
}

// Creates a dimension that is computed as a fraction of the height of the containing group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutdimension/3199058-fractionalheightdimension?language=objc
func CollectionLayoutDimension_FractionalHeightDimension(fractionalHeight float64) CollectionLayoutDimension {
	return CollectionLayoutDimensionClass.FractionalHeightDimension(fractionalHeight)
}

func (cc _CollectionLayoutDimensionClass) AbsoluteDimension(absoluteDimension float64) CollectionLayoutDimension {
	rv := objc.Call[CollectionLayoutDimension](cc, objc.Sel("absoluteDimension:"), absoluteDimension)
	return rv
}

// Creates a dimension with an absolute point value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutdimension/3199055-absolutedimension?language=objc
func CollectionLayoutDimension_AbsoluteDimension(absoluteDimension float64) CollectionLayoutDimension {
	return CollectionLayoutDimensionClass.AbsoluteDimension(absoluteDimension)
}

func (cc _CollectionLayoutDimensionClass) FractionalWidthDimension(fractionalWidth float64) CollectionLayoutDimension {
	rv := objc.Call[CollectionLayoutDimension](cc, objc.Sel("fractionalWidthDimension:"), fractionalWidth)
	return rv
}

// Creates a dimension that is computed as a fraction of the width of the containing group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutdimension/3199059-fractionalwidthdimension?language=objc
func CollectionLayoutDimension_FractionalWidthDimension(fractionalWidth float64) CollectionLayoutDimension {
	return CollectionLayoutDimensionClass.FractionalWidthDimension(fractionalWidth)
}

func (cc _CollectionLayoutDimensionClass) EstimatedDimension(estimatedDimension float64) CollectionLayoutDimension {
	rv := objc.Call[CollectionLayoutDimension](cc, objc.Sel("estimatedDimension:"), estimatedDimension)
	return rv
}

// Creates a dimension with an estimated point value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutdimension/3199057-estimateddimension?language=objc
func CollectionLayoutDimension_EstimatedDimension(estimatedDimension float64) CollectionLayoutDimension {
	return CollectionLayoutDimensionClass.EstimatedDimension(estimatedDimension)
}

func (cc _CollectionLayoutDimensionClass) Alloc() CollectionLayoutDimension {
	rv := objc.Call[CollectionLayoutDimension](cc, objc.Sel("alloc"))
	return rv
}

func CollectionLayoutDimension_Alloc() CollectionLayoutDimension {
	return CollectionLayoutDimensionClass.Alloc()
}

func (cc _CollectionLayoutDimensionClass) New() CollectionLayoutDimension {
	rv := objc.Call[CollectionLayoutDimension](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutDimension() CollectionLayoutDimension {
	return CollectionLayoutDimensionClass.New()
}

func (c_ CollectionLayoutDimension) Init() CollectionLayoutDimension {
	rv := objc.Call[CollectionLayoutDimension](c_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether the dimension is expressed as an estimated value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutdimension/3199061-isestimated?language=objc
func (c_ CollectionLayoutDimension) IsEstimated() bool {
	rv := objc.Call[bool](c_, objc.Sel("isEstimated"))
	return rv
}

// The floating-point value of the dimension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutdimension/3199056-dimension?language=objc
func (c_ CollectionLayoutDimension) Dimension() float64 {
	rv := objc.Call[float64](c_, objc.Sel("dimension"))
	return rv
}

// A Boolean value that indicates whether the dimension is expressed as a fraction of its container’s width. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutdimension/3199063-isfractionalwidth?language=objc
func (c_ CollectionLayoutDimension) IsFractionalWidth() bool {
	rv := objc.Call[bool](c_, objc.Sel("isFractionalWidth"))
	return rv
}

// A Boolean value that indicates whether the dimension is expressed as an absolute value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutdimension/3199060-isabsolute?language=objc
func (c_ CollectionLayoutDimension) IsAbsolute() bool {
	rv := objc.Call[bool](c_, objc.Sel("isAbsolute"))
	return rv
}

// A Boolean value that indicates whether the dimension is expressed as a fraction of its container’s height. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutdimension/3199062-isfractionalheight?language=objc
func (c_ CollectionLayoutDimension) IsFractionalHeight() bool {
	rv := objc.Call[bool](c_, objc.Sel("isFractionalHeight"))
	return rv
}
