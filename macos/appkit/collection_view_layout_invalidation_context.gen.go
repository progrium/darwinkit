// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var CollectionViewLayoutInvalidationContextClass = _CollectionViewLayoutInvalidationContextClass{objc.GetClass("NSCollectionViewLayoutInvalidationContext")}

type _CollectionViewLayoutInvalidationContextClass struct {
	objc.Class
}

type ICollectionViewLayoutInvalidationContext interface {
	objc.IObject
	InvalidateItemsAtIndexPaths(indexPaths foundation.ISet)
	InvalidateSupplementaryElementsOfKindAtIndexPaths(elementKind CollectionViewSupplementaryElementKind, indexPaths foundation.ISet)
	InvalidateDecorationElementsOfKindAtIndexPaths(elementKind CollectionViewDecorationElementKind, indexPaths foundation.ISet)
	InvalidateEverything() bool
	InvalidateDataSourceCounts() bool
	ContentOffsetAdjustment() foundation.Point
	SetContentOffsetAdjustment(value foundation.Point)
	ContentSizeAdjustment() foundation.Size
	SetContentSizeAdjustment(value foundation.Size)
	InvalidatedItemIndexPaths() foundation.Set
	InvalidatedSupplementaryIndexPaths() map[CollectionViewSupplementaryElementKind]foundation.Set
	InvalidatedDecorationIndexPaths() map[CollectionViewDecorationElementKind]foundation.Set
}

type CollectionViewLayoutInvalidationContext struct {
	objc.Object
}

func MakeCollectionViewLayoutInvalidationContext(ptr unsafe.Pointer) CollectionViewLayoutInvalidationContext {
	return CollectionViewLayoutInvalidationContext{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionViewLayoutInvalidationContextClass) Alloc() CollectionViewLayoutInvalidationContext {
	rv := objc.CallMethod[CollectionViewLayoutInvalidationContext](cc, objc.GetSelector("alloc"))
	return rv
}

func CollectionViewLayoutInvalidationContext_Alloc() CollectionViewLayoutInvalidationContext {
	return CollectionViewLayoutInvalidationContextClass.Alloc()
}

func (cc _CollectionViewLayoutInvalidationContextClass) New() CollectionViewLayoutInvalidationContext {
	rv := objc.CallMethod[CollectionViewLayoutInvalidationContext](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionViewLayoutInvalidationContext() CollectionViewLayoutInvalidationContext {
	return CollectionViewLayoutInvalidationContextClass.New()
}

func CollectionViewLayoutInvalidationContext_New() CollectionViewLayoutInvalidationContext {
	return CollectionViewLayoutInvalidationContextClass.New()
}

func (c_ CollectionViewLayoutInvalidationContext) Init() CollectionViewLayoutInvalidationContext {
	rv := objc.CallMethod[CollectionViewLayoutInvalidationContext](c_, objc.GetSelector("init"))
	return rv
}

func CollectionViewLayoutInvalidationContext_Init() CollectionViewLayoutInvalidationContext {
	return CollectionViewLayoutInvalidationContextClass.Alloc().Init()
}

func (c_ CollectionViewLayoutInvalidationContext) InvalidateItemsAtIndexPaths(indexPaths foundation.ISet) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("invalidateItemsAtIndexPaths:"), objc.ExtractPtr(indexPaths))
}

func (c_ CollectionViewLayoutInvalidationContext) InvalidateSupplementaryElementsOfKindAtIndexPaths(elementKind CollectionViewSupplementaryElementKind, indexPaths foundation.ISet) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("invalidateSupplementaryElementsOfKind:atIndexPaths:"), elementKind, objc.ExtractPtr(indexPaths))
}

func (c_ CollectionViewLayoutInvalidationContext) InvalidateDecorationElementsOfKindAtIndexPaths(elementKind CollectionViewDecorationElementKind, indexPaths foundation.ISet) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("invalidateDecorationElementsOfKind:atIndexPaths:"), elementKind, objc.ExtractPtr(indexPaths))
}

func (c_ CollectionViewLayoutInvalidationContext) InvalidateEverything() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("invalidateEverything"))
	return rv
}

func (c_ CollectionViewLayoutInvalidationContext) InvalidateDataSourceCounts() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("invalidateDataSourceCounts"))
	return rv
}

func (c_ CollectionViewLayoutInvalidationContext) ContentOffsetAdjustment() foundation.Point {
	rv := objc.CallMethod[foundation.Point](c_, objc.GetSelector("contentOffsetAdjustment"))
	return rv
}

func (c_ CollectionViewLayoutInvalidationContext) SetContentOffsetAdjustment(value foundation.Point) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setContentOffsetAdjustment:"), value)
}

func (c_ CollectionViewLayoutInvalidationContext) ContentSizeAdjustment() foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.GetSelector("contentSizeAdjustment"))
	return rv
}

func (c_ CollectionViewLayoutInvalidationContext) SetContentSizeAdjustment(value foundation.Size) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setContentSizeAdjustment:"), value)
}

func (c_ CollectionViewLayoutInvalidationContext) InvalidatedItemIndexPaths() foundation.Set {
	rv := objc.CallMethod[foundation.Set](c_, objc.GetSelector("invalidatedItemIndexPaths"))
	return rv
}

func (c_ CollectionViewLayoutInvalidationContext) InvalidatedSupplementaryIndexPaths() map[CollectionViewSupplementaryElementKind]foundation.Set {
	rv := objc.CallMethod[map[CollectionViewSupplementaryElementKind]foundation.Set](c_, objc.GetSelector("invalidatedSupplementaryIndexPaths"))
	return rv
}

func (c_ CollectionViewLayoutInvalidationContext) InvalidatedDecorationIndexPaths() map[CollectionViewDecorationElementKind]foundation.Set {
	rv := objc.CallMethod[map[CollectionViewDecorationElementKind]foundation.Set](c_, objc.GetSelector("invalidatedDecorationIndexPaths"))
	return rv
}
