// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CollectionViewLayoutInvalidationContext] class.
var CollectionViewLayoutInvalidationContextClass = _CollectionViewLayoutInvalidationContextClass{objc.GetClass("NSCollectionViewLayoutInvalidationContext")}

type _CollectionViewLayoutInvalidationContextClass struct {
	objc.Class
}

// An interface definition for the [CollectionViewLayoutInvalidationContext] class.
type ICollectionViewLayoutInvalidationContext interface {
	objc.IObject
	InvalidateSupplementaryElementsOfKindAtIndexPaths(elementKind CollectionViewSupplementaryElementKind, indexPaths foundation.ISet)
	InvalidateDecorationElementsOfKindAtIndexPaths(elementKind CollectionViewDecorationElementKind, indexPaths foundation.ISet)
	InvalidateItemsAtIndexPaths(indexPaths foundation.ISet)
	InvalidatedDecorationIndexPaths() map[CollectionViewDecorationElementKind]foundation.Set
	InvalidateEverything() bool
	InvalidatedItemIndexPaths() foundation.Set
	InvalidateDataSourceCounts() bool
	InvalidatedSupplementaryIndexPaths() map[CollectionViewSupplementaryElementKind]foundation.Set
	ContentSizeAdjustment() foundation.Size
	SetContentSizeAdjustment(value foundation.Size)
	ContentOffsetAdjustment() foundation.Point
	SetContentOffsetAdjustment(value foundation.Point)
}

// An object that identifies the portions of your layout that need to be updated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutinvalidationcontext?language=objc
type CollectionViewLayoutInvalidationContext struct {
	objc.Object
}

func CollectionViewLayoutInvalidationContextFrom(ptr unsafe.Pointer) CollectionViewLayoutInvalidationContext {
	return CollectionViewLayoutInvalidationContext{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CollectionViewLayoutInvalidationContextClass) Alloc() CollectionViewLayoutInvalidationContext {
	rv := objc.Call[CollectionViewLayoutInvalidationContext](cc, objc.Sel("alloc"))
	return rv
}

func CollectionViewLayoutInvalidationContext_Alloc() CollectionViewLayoutInvalidationContext {
	return CollectionViewLayoutInvalidationContextClass.Alloc()
}

func (cc _CollectionViewLayoutInvalidationContextClass) New() CollectionViewLayoutInvalidationContext {
	rv := objc.Call[CollectionViewLayoutInvalidationContext](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionViewLayoutInvalidationContext() CollectionViewLayoutInvalidationContext {
	return CollectionViewLayoutInvalidationContextClass.New()
}

func (c_ CollectionViewLayoutInvalidationContext) Init() CollectionViewLayoutInvalidationContext {
	rv := objc.Call[CollectionViewLayoutInvalidationContext](c_, objc.Sel("init"))
	return rv
}

// Marks the specified supplementary views as invalid so that their layout information can be updated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutinvalidationcontext/1533990-invalidatesupplementaryelementso?language=objc
func (c_ CollectionViewLayoutInvalidationContext) InvalidateSupplementaryElementsOfKindAtIndexPaths(elementKind CollectionViewSupplementaryElementKind, indexPaths foundation.ISet) {
	objc.Call[objc.Void](c_, objc.Sel("invalidateSupplementaryElementsOfKind:atIndexPaths:"), elementKind, objc.Ptr(indexPaths))
}

// Marks the specified decoration views as invalid so that their layout information can be updated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutinvalidationcontext/1535363-invalidatedecorationelementsofki?language=objc
func (c_ CollectionViewLayoutInvalidationContext) InvalidateDecorationElementsOfKindAtIndexPaths(elementKind CollectionViewDecorationElementKind, indexPaths foundation.ISet) {
	objc.Call[objc.Void](c_, objc.Sel("invalidateDecorationElementsOfKind:atIndexPaths:"), elementKind, objc.Ptr(indexPaths))
}

// Marks the specified items as invalid so that their layout information can be updated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutinvalidationcontext/1526592-invalidateitemsatindexpaths?language=objc
func (c_ CollectionViewLayoutInvalidationContext) InvalidateItemsAtIndexPaths(indexPaths foundation.ISet) {
	objc.Call[objc.Void](c_, objc.Sel("invalidateItemsAtIndexPaths:"), objc.Ptr(indexPaths))
}

// A dictionary containing the decoration views whose layout attributes are invalid. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutinvalidationcontext/1530481-invalidateddecorationindexpaths?language=objc
func (c_ CollectionViewLayoutInvalidationContext) InvalidatedDecorationIndexPaths() map[CollectionViewDecorationElementKind]foundation.Set {
	rv := objc.Call[map[CollectionViewDecorationElementKind]foundation.Set](c_, objc.Sel("invalidatedDecorationIndexPaths"))
	return rv
}

// A Boolean that indicates whether all layout data should be marked as invalid. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutinvalidationcontext/1525613-invalidateeverything?language=objc
func (c_ CollectionViewLayoutInvalidationContext) InvalidateEverything() bool {
	rv := objc.Call[bool](c_, objc.Sel("invalidateEverything"))
	return rv
}

// The set of items whose layout attributes are invalid. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutinvalidationcontext/1525361-invalidateditemindexpaths?language=objc
func (c_ CollectionViewLayoutInvalidationContext) InvalidatedItemIndexPaths() foundation.Set {
	rv := objc.Call[foundation.Set](c_, objc.Sel("invalidatedItemIndexPaths"))
	return rv
}

// A Boolean that indicates whether the layout object should ask for new section and item counts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutinvalidationcontext/1526189-invalidatedatasourcecounts?language=objc
func (c_ CollectionViewLayoutInvalidationContext) InvalidateDataSourceCounts() bool {
	rv := objc.Call[bool](c_, objc.Sel("invalidateDataSourceCounts"))
	return rv
}

// A dictionary containing the supplementary views whose layout attributes are invalid. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutinvalidationcontext/1531052-invalidatedsupplementaryindexpat?language=objc
func (c_ CollectionViewLayoutInvalidationContext) InvalidatedSupplementaryIndexPaths() map[CollectionViewSupplementaryElementKind]foundation.Set {
	rv := objc.Call[map[CollectionViewSupplementaryElementKind]foundation.Set](c_, objc.Sel("invalidatedSupplementaryIndexPaths"))
	return rv
}

// The delta value to add to the collection view’s content size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutinvalidationcontext/1527883-contentsizeadjustment?language=objc
func (c_ CollectionViewLayoutInvalidationContext) ContentSizeAdjustment() foundation.Size {
	rv := objc.Call[foundation.Size](c_, objc.Sel("contentSizeAdjustment"))
	return rv
}

// The delta value to add to the collection view’s content size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutinvalidationcontext/1527883-contentsizeadjustment?language=objc
func (c_ CollectionViewLayoutInvalidationContext) SetContentSizeAdjustment(value foundation.Size) {
	objc.Call[objc.Void](c_, objc.Sel("setContentSizeAdjustment:"), value)
}

// The delta value to add to the collection view’s content offset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutinvalidationcontext/1533465-contentoffsetadjustment?language=objc
func (c_ CollectionViewLayoutInvalidationContext) ContentOffsetAdjustment() foundation.Point {
	rv := objc.Call[foundation.Point](c_, objc.Sel("contentOffsetAdjustment"))
	return rv
}

// The delta value to add to the collection view’s content offset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutinvalidationcontext/1533465-contentoffsetadjustment?language=objc
func (c_ CollectionViewLayoutInvalidationContext) SetContentOffsetAdjustment(value foundation.Point) {
	objc.Call[objc.Void](c_, objc.Sel("setContentOffsetAdjustment:"), value)
}
