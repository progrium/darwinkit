// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var CollectionLayoutSpacingClass = _CollectionLayoutSpacingClass{objc.GetClass("NSCollectionLayoutSpacing")}

type _CollectionLayoutSpacingClass struct {
	objc.Class
}

type ICollectionLayoutSpacing interface {
	objc.IObject
	Spacing() float64
	IsFixedSpacing() bool
	IsFlexibleSpacing() bool
}

type CollectionLayoutSpacing struct {
	objc.Object
}

func MakeCollectionLayoutSpacing(ptr unsafe.Pointer) CollectionLayoutSpacing {
	return CollectionLayoutSpacing{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionLayoutSpacingClass) FixedSpacing(fixedSpacing float64) CollectionLayoutSpacing {
	rv := objc.CallMethod[CollectionLayoutSpacing](cc, objc.GetSelector("fixedSpacing:"), fixedSpacing)
	return rv
}

func CollectionLayoutSpacing_FixedSpacing(fixedSpacing float64) CollectionLayoutSpacing {
	return CollectionLayoutSpacingClass.FixedSpacing(fixedSpacing)
}

func (cc _CollectionLayoutSpacingClass) FlexibleSpacing(flexibleSpacing float64) CollectionLayoutSpacing {
	rv := objc.CallMethod[CollectionLayoutSpacing](cc, objc.GetSelector("flexibleSpacing:"), flexibleSpacing)
	return rv
}

func CollectionLayoutSpacing_FlexibleSpacing(flexibleSpacing float64) CollectionLayoutSpacing {
	return CollectionLayoutSpacingClass.FlexibleSpacing(flexibleSpacing)
}

func (cc _CollectionLayoutSpacingClass) Alloc() CollectionLayoutSpacing {
	rv := objc.CallMethod[CollectionLayoutSpacing](cc, objc.GetSelector("alloc"))
	return rv
}

func CollectionLayoutSpacing_Alloc() CollectionLayoutSpacing {
	return CollectionLayoutSpacingClass.Alloc()
}

func (cc _CollectionLayoutSpacingClass) New() CollectionLayoutSpacing {
	rv := objc.CallMethod[CollectionLayoutSpacing](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutSpacing() CollectionLayoutSpacing {
	return CollectionLayoutSpacingClass.New()
}

func CollectionLayoutSpacing_New() CollectionLayoutSpacing {
	return CollectionLayoutSpacingClass.New()
}

func (c_ CollectionLayoutSpacing) Init() CollectionLayoutSpacing {
	rv := objc.CallMethod[CollectionLayoutSpacing](c_, objc.GetSelector("init"))
	return rv
}

func CollectionLayoutSpacing_Init() CollectionLayoutSpacing {
	return CollectionLayoutSpacingClass.Alloc().Init()
}

func (c_ CollectionLayoutSpacing) Spacing() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("spacing"))
	return rv
}

func (c_ CollectionLayoutSpacing) IsFixedSpacing() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isFixedSpacing"))
	return rv
}

func (c_ CollectionLayoutSpacing) IsFlexibleSpacing() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isFlexibleSpacing"))
	return rv
}
