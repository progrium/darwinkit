// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CollectionLayoutAnchor] class.
var CollectionLayoutAnchorClass = _CollectionLayoutAnchorClass{objc.GetClass("NSCollectionLayoutAnchor")}

type _CollectionLayoutAnchorClass struct {
	objc.Class
}

// An interface definition for the [CollectionLayoutAnchor] class.
type ICollectionLayoutAnchor interface {
	objc.IObject
	Edges() DirectionalRectEdge
	IsFractionalOffset() bool
	IsAbsoluteOffset() bool
	Offset() coregraphics.Point
}

// An object that defines how to attach a supplementary item to an item in a collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutanchor?language=objc
type CollectionLayoutAnchor struct {
	objc.Object
}

func CollectionLayoutAnchorFrom(ptr unsafe.Pointer) CollectionLayoutAnchor {
	return CollectionLayoutAnchor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CollectionLayoutAnchorClass) LayoutAnchorWithEdges(edges DirectionalRectEdge) CollectionLayoutAnchor {
	rv := objc.Call[CollectionLayoutAnchor](cc, objc.Sel("layoutAnchorWithEdges:"), edges)
	return rv
}

// Creates an anchor with the specified edges to attach to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutanchor/3199034-layoutanchorwithedges?language=objc
func CollectionLayoutAnchor_LayoutAnchorWithEdges(edges DirectionalRectEdge) CollectionLayoutAnchor {
	return CollectionLayoutAnchorClass.LayoutAnchorWithEdges(edges)
}

func (cc _CollectionLayoutAnchorClass) Alloc() CollectionLayoutAnchor {
	rv := objc.Call[CollectionLayoutAnchor](cc, objc.Sel("alloc"))
	return rv
}

func CollectionLayoutAnchor_Alloc() CollectionLayoutAnchor {
	return CollectionLayoutAnchorClass.Alloc()
}

func (cc _CollectionLayoutAnchorClass) New() CollectionLayoutAnchor {
	rv := objc.Call[CollectionLayoutAnchor](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutAnchor() CollectionLayoutAnchor {
	return CollectionLayoutAnchorClass.New()
}

func (c_ CollectionLayoutAnchor) Init() CollectionLayoutAnchor {
	rv := objc.Call[CollectionLayoutAnchor](c_, objc.Sel("init"))
	return rv
}

// The edges of the item an anchor is attached to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutanchor/3199030-edges?language=objc
func (c_ CollectionLayoutAnchor) Edges() DirectionalRectEdge {
	rv := objc.Call[DirectionalRectEdge](c_, objc.Sel("edges"))
	return rv
}

// A Boolean value that indicates whether the anchor's offset is expressed as a fraction of its supplementary item's dimension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutanchor/3213812-isfractionaloffset?language=objc
func (c_ CollectionLayoutAnchor) IsFractionalOffset() bool {
	rv := objc.Call[bool](c_, objc.Sel("isFractionalOffset"))
	return rv
}

// A Boolean value that indicates whether the anchor's offset is expressed as an absolute value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutanchor/3213811-isabsoluteoffset?language=objc
func (c_ CollectionLayoutAnchor) IsAbsoluteOffset() bool {
	rv := objc.Call[bool](c_, objc.Sel("isAbsoluteOffset"))
	return rv
}

// The floating-point value of the anchor's offset from the item it's attached to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutanchor/3199037-offset?language=objc
func (c_ CollectionLayoutAnchor) Offset() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](c_, objc.Sel("offset"))
	return rv
}
