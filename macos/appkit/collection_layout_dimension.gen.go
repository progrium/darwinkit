// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var CollectionLayoutDimensionClass = _CollectionLayoutDimensionClass{objc.GetClass("NSCollectionLayoutDimension")}

type _CollectionLayoutDimensionClass struct {
	objc.Class
}

type ICollectionLayoutDimension interface {
	objc.IObject
	Dimension() float64
	IsAbsolute() bool
	IsEstimated() bool
	IsFractionalHeight() bool
	IsFractionalWidth() bool
}

type CollectionLayoutDimension struct {
	objc.Object
}

func MakeCollectionLayoutDimension(ptr unsafe.Pointer) CollectionLayoutDimension {
	return CollectionLayoutDimension{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionLayoutDimensionClass) AbsoluteDimension(absoluteDimension float64) CollectionLayoutDimension {
	rv := objc.CallMethod[CollectionLayoutDimension](cc, objc.GetSelector("absoluteDimension:"), absoluteDimension)
	return rv
}

func CollectionLayoutDimension_AbsoluteDimension(absoluteDimension float64) CollectionLayoutDimension {
	return CollectionLayoutDimensionClass.AbsoluteDimension(absoluteDimension)
}

func (cc _CollectionLayoutDimensionClass) EstimatedDimension(estimatedDimension float64) CollectionLayoutDimension {
	rv := objc.CallMethod[CollectionLayoutDimension](cc, objc.GetSelector("estimatedDimension:"), estimatedDimension)
	return rv
}

func CollectionLayoutDimension_EstimatedDimension(estimatedDimension float64) CollectionLayoutDimension {
	return CollectionLayoutDimensionClass.EstimatedDimension(estimatedDimension)
}

func (cc _CollectionLayoutDimensionClass) FractionalHeightDimension(fractionalHeight float64) CollectionLayoutDimension {
	rv := objc.CallMethod[CollectionLayoutDimension](cc, objc.GetSelector("fractionalHeightDimension:"), fractionalHeight)
	return rv
}

func CollectionLayoutDimension_FractionalHeightDimension(fractionalHeight float64) CollectionLayoutDimension {
	return CollectionLayoutDimensionClass.FractionalHeightDimension(fractionalHeight)
}

func (cc _CollectionLayoutDimensionClass) FractionalWidthDimension(fractionalWidth float64) CollectionLayoutDimension {
	rv := objc.CallMethod[CollectionLayoutDimension](cc, objc.GetSelector("fractionalWidthDimension:"), fractionalWidth)
	return rv
}

func CollectionLayoutDimension_FractionalWidthDimension(fractionalWidth float64) CollectionLayoutDimension {
	return CollectionLayoutDimensionClass.FractionalWidthDimension(fractionalWidth)
}

func (cc _CollectionLayoutDimensionClass) Alloc() CollectionLayoutDimension {
	rv := objc.CallMethod[CollectionLayoutDimension](cc, objc.GetSelector("alloc"))
	return rv
}

func CollectionLayoutDimension_Alloc() CollectionLayoutDimension {
	return CollectionLayoutDimensionClass.Alloc()
}

func (cc _CollectionLayoutDimensionClass) New() CollectionLayoutDimension {
	rv := objc.CallMethod[CollectionLayoutDimension](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutDimension() CollectionLayoutDimension {
	return CollectionLayoutDimensionClass.New()
}

func CollectionLayoutDimension_New() CollectionLayoutDimension {
	return CollectionLayoutDimensionClass.New()
}

func (c_ CollectionLayoutDimension) Init() CollectionLayoutDimension {
	rv := objc.CallMethod[CollectionLayoutDimension](c_, objc.GetSelector("init"))
	return rv
}

func CollectionLayoutDimension_Init() CollectionLayoutDimension {
	return CollectionLayoutDimensionClass.Alloc().Init()
}

func (c_ CollectionLayoutDimension) Dimension() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("dimension"))
	return rv
}

func (c_ CollectionLayoutDimension) IsAbsolute() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isAbsolute"))
	return rv
}

func (c_ CollectionLayoutDimension) IsEstimated() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isEstimated"))
	return rv
}

func (c_ CollectionLayoutDimension) IsFractionalHeight() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isFractionalHeight"))
	return rv
}

func (c_ CollectionLayoutDimension) IsFractionalWidth() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isFractionalWidth"))
	return rv
}
