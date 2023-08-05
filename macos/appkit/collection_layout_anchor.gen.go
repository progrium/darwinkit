// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var CollectionLayoutAnchorClass = _CollectionLayoutAnchorClass{objc.GetClass("NSCollectionLayoutAnchor")}

type _CollectionLayoutAnchorClass struct {
	objc.Class
}

type ICollectionLayoutAnchor interface {
	objc.IObject
	Edges() DirectionalRectEdge
	Offset() foundation.Point
	IsAbsoluteOffset() bool
	IsFractionalOffset() bool
}

type CollectionLayoutAnchor struct {
	objc.Object
}

func MakeCollectionLayoutAnchor(ptr unsafe.Pointer) CollectionLayoutAnchor {
	return CollectionLayoutAnchor{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionLayoutAnchorClass) LayoutAnchorWithEdges(edges DirectionalRectEdge) CollectionLayoutAnchor {
	rv := objc.CallMethod[CollectionLayoutAnchor](cc, objc.GetSelector("layoutAnchorWithEdges:"), edges)
	return rv
}

func CollectionLayoutAnchor_LayoutAnchorWithEdges(edges DirectionalRectEdge) CollectionLayoutAnchor {
	return CollectionLayoutAnchorClass.LayoutAnchorWithEdges(edges)
}

func (cc _CollectionLayoutAnchorClass) LayoutAnchorWithEdgesAbsoluteOffset(edges DirectionalRectEdge, absoluteOffset foundation.Point) CollectionLayoutAnchor {
	rv := objc.CallMethod[CollectionLayoutAnchor](cc, objc.GetSelector("layoutAnchorWithEdges:absoluteOffset:"), edges, absoluteOffset)
	return rv
}

func CollectionLayoutAnchor_LayoutAnchorWithEdgesAbsoluteOffset(edges DirectionalRectEdge, absoluteOffset foundation.Point) CollectionLayoutAnchor {
	return CollectionLayoutAnchorClass.LayoutAnchorWithEdgesAbsoluteOffset(edges, absoluteOffset)
}

func (cc _CollectionLayoutAnchorClass) LayoutAnchorWithEdgesFractionalOffset(edges DirectionalRectEdge, fractionalOffset foundation.Point) CollectionLayoutAnchor {
	rv := objc.CallMethod[CollectionLayoutAnchor](cc, objc.GetSelector("layoutAnchorWithEdges:fractionalOffset:"), edges, fractionalOffset)
	return rv
}

func CollectionLayoutAnchor_LayoutAnchorWithEdgesFractionalOffset(edges DirectionalRectEdge, fractionalOffset foundation.Point) CollectionLayoutAnchor {
	return CollectionLayoutAnchorClass.LayoutAnchorWithEdgesFractionalOffset(edges, fractionalOffset)
}

func (cc _CollectionLayoutAnchorClass) Alloc() CollectionLayoutAnchor {
	rv := objc.CallMethod[CollectionLayoutAnchor](cc, objc.GetSelector("alloc"))
	return rv
}

func CollectionLayoutAnchor_Alloc() CollectionLayoutAnchor {
	return CollectionLayoutAnchorClass.Alloc()
}

func (cc _CollectionLayoutAnchorClass) New() CollectionLayoutAnchor {
	rv := objc.CallMethod[CollectionLayoutAnchor](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutAnchor() CollectionLayoutAnchor {
	return CollectionLayoutAnchorClass.New()
}

func CollectionLayoutAnchor_New() CollectionLayoutAnchor {
	return CollectionLayoutAnchorClass.New()
}

func (c_ CollectionLayoutAnchor) Init() CollectionLayoutAnchor {
	rv := objc.CallMethod[CollectionLayoutAnchor](c_, objc.GetSelector("init"))
	return rv
}

func CollectionLayoutAnchor_Init() CollectionLayoutAnchor {
	return CollectionLayoutAnchorClass.Alloc().Init()
}

func (c_ CollectionLayoutAnchor) Edges() DirectionalRectEdge {
	rv := objc.CallMethod[DirectionalRectEdge](c_, objc.GetSelector("edges"))
	return rv
}

func (c_ CollectionLayoutAnchor) Offset() foundation.Point {
	rv := objc.CallMethod[foundation.Point](c_, objc.GetSelector("offset"))
	return rv
}

func (c_ CollectionLayoutAnchor) IsAbsoluteOffset() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isAbsoluteOffset"))
	return rv
}

func (c_ CollectionLayoutAnchor) IsFractionalOffset() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isFractionalOffset"))
	return rv
}
