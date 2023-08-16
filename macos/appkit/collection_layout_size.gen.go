// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CollectionLayoutSize] class.
var CollectionLayoutSizeClass = _CollectionLayoutSizeClass{objc.GetClass("NSCollectionLayoutSize")}

type _CollectionLayoutSizeClass struct {
	objc.Class
}

// An interface definition for the [CollectionLayoutSize] class.
type ICollectionLayoutSize interface {
	objc.IObject
	HeightDimension() CollectionLayoutDimension
	WidthDimension() CollectionLayoutDimension
}

// The width and the height of an item in a collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutsize?language=objc
type CollectionLayoutSize struct {
	objc.Object
}

func CollectionLayoutSizeFrom(ptr unsafe.Pointer) CollectionLayoutSize {
	return CollectionLayoutSize{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CollectionLayoutSizeClass) SizeWithWidthDimensionHeightDimension(width ICollectionLayoutDimension, height ICollectionLayoutDimension) CollectionLayoutSize {
	rv := objc.Call[CollectionLayoutSize](cc, objc.Sel("sizeWithWidthDimension:heightDimension:"), objc.Ptr(width), objc.Ptr(height))
	return rv
}

// Creates a size with the specified width and height dimensions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutsize/3213887-sizewithwidthdimension?language=objc
func CollectionLayoutSize_SizeWithWidthDimensionHeightDimension(width ICollectionLayoutDimension, height ICollectionLayoutDimension) CollectionLayoutSize {
	return CollectionLayoutSizeClass.SizeWithWidthDimensionHeightDimension(width, height)
}

func (cc _CollectionLayoutSizeClass) Alloc() CollectionLayoutSize {
	rv := objc.Call[CollectionLayoutSize](cc, objc.Sel("alloc"))
	return rv
}

func CollectionLayoutSize_Alloc() CollectionLayoutSize {
	return CollectionLayoutSizeClass.Alloc()
}

func (cc _CollectionLayoutSizeClass) New() CollectionLayoutSize {
	rv := objc.Call[CollectionLayoutSize](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutSize() CollectionLayoutSize {
	return CollectionLayoutSizeClass.New()
}

func (c_ CollectionLayoutSize) Init() CollectionLayoutSize {
	rv := objc.Call[CollectionLayoutSize](c_, objc.Sel("init"))
	return rv
}

// The height dimension of an item in a collection view layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutsize/3213886-heightdimension?language=objc
func (c_ CollectionLayoutSize) HeightDimension() CollectionLayoutDimension {
	rv := objc.Call[CollectionLayoutDimension](c_, objc.Sel("heightDimension"))
	return rv
}

// The width dimension of an item in a collection view layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutsize/3213888-widthdimension?language=objc
func (c_ CollectionLayoutSize) WidthDimension() CollectionLayoutDimension {
	rv := objc.Call[CollectionLayoutDimension](c_, objc.Sel("widthDimension"))
	return rv
}
