// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var CollectionLayoutSizeClass = _CollectionLayoutSizeClass{objc.GetClass("NSCollectionLayoutSize")}

type _CollectionLayoutSizeClass struct {
	objc.Class
}

type ICollectionLayoutSize interface {
	objc.IObject
	WidthDimension() CollectionLayoutDimension
	HeightDimension() CollectionLayoutDimension
}

type CollectionLayoutSize struct {
	objc.Object
}

func MakeCollectionLayoutSize(ptr unsafe.Pointer) CollectionLayoutSize {
	return CollectionLayoutSize{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionLayoutSizeClass) SizeWithWidthDimensionHeightDimension(width ICollectionLayoutDimension, height ICollectionLayoutDimension) CollectionLayoutSize {
	rv := objc.CallMethod[CollectionLayoutSize](cc, objc.GetSelector("sizeWithWidthDimension:heightDimension:"), objc.ExtractPtr(width), objc.ExtractPtr(height))
	return rv
}

func CollectionLayoutSize_SizeWithWidthDimensionHeightDimension(width ICollectionLayoutDimension, height ICollectionLayoutDimension) CollectionLayoutSize {
	return CollectionLayoutSizeClass.SizeWithWidthDimensionHeightDimension(width, height)
}

func (cc _CollectionLayoutSizeClass) Alloc() CollectionLayoutSize {
	rv := objc.CallMethod[CollectionLayoutSize](cc, objc.GetSelector("alloc"))
	return rv
}

func CollectionLayoutSize_Alloc() CollectionLayoutSize {
	return CollectionLayoutSizeClass.Alloc()
}

func (cc _CollectionLayoutSizeClass) New() CollectionLayoutSize {
	rv := objc.CallMethod[CollectionLayoutSize](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutSize() CollectionLayoutSize {
	return CollectionLayoutSizeClass.New()
}

func CollectionLayoutSize_New() CollectionLayoutSize {
	return CollectionLayoutSizeClass.New()
}

func (c_ CollectionLayoutSize) Init() CollectionLayoutSize {
	rv := objc.CallMethod[CollectionLayoutSize](c_, objc.GetSelector("init"))
	return rv
}

func CollectionLayoutSize_Init() CollectionLayoutSize {
	return CollectionLayoutSizeClass.Alloc().Init()
}

func (c_ CollectionLayoutSize) WidthDimension() CollectionLayoutDimension {
	rv := objc.CallMethod[CollectionLayoutDimension](c_, objc.GetSelector("widthDimension"))
	return rv
}

func (c_ CollectionLayoutSize) HeightDimension() CollectionLayoutDimension {
	rv := objc.CallMethod[CollectionLayoutDimension](c_, objc.GetSelector("heightDimension"))
	return rv
}
