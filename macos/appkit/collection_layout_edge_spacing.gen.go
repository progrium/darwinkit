// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var CollectionLayoutEdgeSpacingClass = _CollectionLayoutEdgeSpacingClass{objc.GetClass("NSCollectionLayoutEdgeSpacing")}

type _CollectionLayoutEdgeSpacingClass struct {
	objc.Class
}

type ICollectionLayoutEdgeSpacing interface {
	objc.IObject
	Leading() CollectionLayoutSpacing
	Top() CollectionLayoutSpacing
	Trailing() CollectionLayoutSpacing
	Bottom() CollectionLayoutSpacing
}

type CollectionLayoutEdgeSpacing struct {
	objc.Object
}

func MakeCollectionLayoutEdgeSpacing(ptr unsafe.Pointer) CollectionLayoutEdgeSpacing {
	return CollectionLayoutEdgeSpacing{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionLayoutEdgeSpacingClass) SpacingForLeadingTopTrailingBottom(leading ICollectionLayoutSpacing, top ICollectionLayoutSpacing, trailing ICollectionLayoutSpacing, bottom ICollectionLayoutSpacing) CollectionLayoutEdgeSpacing {
	rv := objc.CallMethod[CollectionLayoutEdgeSpacing](cc, objc.GetSelector("spacingForLeading:top:trailing:bottom:"), objc.ExtractPtr(leading), objc.ExtractPtr(top), objc.ExtractPtr(trailing), objc.ExtractPtr(bottom))
	return rv
}

func CollectionLayoutEdgeSpacing_SpacingForLeadingTopTrailingBottom(leading ICollectionLayoutSpacing, top ICollectionLayoutSpacing, trailing ICollectionLayoutSpacing, bottom ICollectionLayoutSpacing) CollectionLayoutEdgeSpacing {
	return CollectionLayoutEdgeSpacingClass.SpacingForLeadingTopTrailingBottom(leading, top, trailing, bottom)
}

func (cc _CollectionLayoutEdgeSpacingClass) Alloc() CollectionLayoutEdgeSpacing {
	rv := objc.CallMethod[CollectionLayoutEdgeSpacing](cc, objc.GetSelector("alloc"))
	return rv
}

func CollectionLayoutEdgeSpacing_Alloc() CollectionLayoutEdgeSpacing {
	return CollectionLayoutEdgeSpacingClass.Alloc()
}

func (cc _CollectionLayoutEdgeSpacingClass) New() CollectionLayoutEdgeSpacing {
	rv := objc.CallMethod[CollectionLayoutEdgeSpacing](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutEdgeSpacing() CollectionLayoutEdgeSpacing {
	return CollectionLayoutEdgeSpacingClass.New()
}

func CollectionLayoutEdgeSpacing_New() CollectionLayoutEdgeSpacing {
	return CollectionLayoutEdgeSpacingClass.New()
}

func (c_ CollectionLayoutEdgeSpacing) Init() CollectionLayoutEdgeSpacing {
	rv := objc.CallMethod[CollectionLayoutEdgeSpacing](c_, objc.GetSelector("init"))
	return rv
}

func CollectionLayoutEdgeSpacing_Init() CollectionLayoutEdgeSpacing {
	return CollectionLayoutEdgeSpacingClass.Alloc().Init()
}

func (c_ CollectionLayoutEdgeSpacing) Leading() CollectionLayoutSpacing {
	rv := objc.CallMethod[CollectionLayoutSpacing](c_, objc.GetSelector("leading"))
	return rv
}

func (c_ CollectionLayoutEdgeSpacing) Top() CollectionLayoutSpacing {
	rv := objc.CallMethod[CollectionLayoutSpacing](c_, objc.GetSelector("top"))
	return rv
}

func (c_ CollectionLayoutEdgeSpacing) Trailing() CollectionLayoutSpacing {
	rv := objc.CallMethod[CollectionLayoutSpacing](c_, objc.GetSelector("trailing"))
	return rv
}

func (c_ CollectionLayoutEdgeSpacing) Bottom() CollectionLayoutSpacing {
	rv := objc.CallMethod[CollectionLayoutSpacing](c_, objc.GetSelector("bottom"))
	return rv
}
