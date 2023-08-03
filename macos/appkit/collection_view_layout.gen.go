// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var CollectionViewLayoutClass = _CollectionViewLayoutClass{objc.GetClass("NSCollectionViewLayout")}

type _CollectionViewLayoutClass struct {
	objc.Class
}

type ICollectionViewLayout interface {
	objc.IObject
	PrepareLayout()
	LayoutAttributesForElementsInRect(rect foundation.Rect) []CollectionViewLayoutAttributes
	LayoutAttributesForItemAtIndexPath(indexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	LayoutAttributesForSupplementaryViewOfKindAtIndexPath(elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	LayoutAttributesForDecorationViewOfKindAtIndexPath(elementKind CollectionViewDecorationElementKind, indexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	LayoutAttributesForDropTargetAtPoint(pointInCollectionView foundation.Point) CollectionViewLayoutAttributes
	LayoutAttributesForInterItemGapBeforeIndexPath(indexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	TargetContentOffsetForProposedContentOffset(proposedContentOffset foundation.Point) foundation.Point
	TargetContentOffsetForProposedContentOffsetWithScrollingVelocity(proposedContentOffset foundation.Point, velocity foundation.Point) foundation.Point
	PrepareForCollectionViewUpdates(updateItems []ICollectionViewUpdateItem)
	FinalizeCollectionViewUpdates()
	IndexPathsToInsertForSupplementaryViewOfKind(elementKind CollectionViewSupplementaryElementKind) foundation.Set
	IndexPathsToInsertForDecorationViewOfKind(elementKind CollectionViewDecorationElementKind) foundation.Set
	InitialLayoutAttributesForAppearingItemAtIndexPath(itemIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	InitialLayoutAttributesForAppearingSupplementaryElementOfKindAtIndexPath(elementKind CollectionViewSupplementaryElementKind, elementIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	InitialLayoutAttributesForAppearingDecorationElementOfKindAtIndexPath(elementKind CollectionViewDecorationElementKind, decorationIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	IndexPathsToDeleteForSupplementaryViewOfKind(elementKind CollectionViewSupplementaryElementKind) foundation.Set
	IndexPathsToDeleteForDecorationViewOfKind(elementKind CollectionViewDecorationElementKind) foundation.Set
	FinalLayoutAttributesForDisappearingItemAtIndexPath(itemIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	FinalLayoutAttributesForDisappearingSupplementaryElementOfKindAtIndexPath(elementKind CollectionViewSupplementaryElementKind, elementIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	FinalLayoutAttributesForDisappearingDecorationElementOfKindAtIndexPath(elementKind CollectionViewDecorationElementKind, decorationIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	InvalidateLayout()
	InvalidateLayoutWithContext(context ICollectionViewLayoutInvalidationContext)
	ShouldInvalidateLayoutForBoundsChange(newBounds foundation.Rect) bool
	ShouldInvalidateLayoutForPreferredLayoutAttributesWithOriginalAttributes(preferredAttributes ICollectionViewLayoutAttributes, originalAttributes ICollectionViewLayoutAttributes) bool
	InvalidationContextForBoundsChange(newBounds foundation.Rect) CollectionViewLayoutInvalidationContext
	InvalidationContextForPreferredLayoutAttributesWithOriginalAttributes(preferredAttributes ICollectionViewLayoutAttributes, originalAttributes ICollectionViewLayoutAttributes) CollectionViewLayoutInvalidationContext
	PrepareForAnimatedBoundsChange(oldBounds foundation.Rect)
	FinalizeAnimatedBoundsChange()
	RegisterClassForDecorationViewOfKind(viewClass objc.IClass, elementKind CollectionViewDecorationElementKind)
	RegisterNibForDecorationViewOfKind(nib INib, elementKind CollectionViewDecorationElementKind)
	PrepareForTransitionFromLayout(oldLayout ICollectionViewLayout)
	PrepareForTransitionToLayout(newLayout ICollectionViewLayout)
	FinalizeLayoutTransition()
	CollectionView() CollectionView
	CollectionViewContentSize() foundation.Size
}

type CollectionViewLayout struct {
	objc.Object
}

func MakeCollectionViewLayout(ptr unsafe.Pointer) CollectionViewLayout {
	return CollectionViewLayout{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionViewLayoutClass) Alloc() CollectionViewLayout {
	rv := objc.CallMethod[CollectionViewLayout](cc, objc.GetSelector("alloc"))
	return rv
}

func CollectionViewLayout_Alloc() CollectionViewLayout {
	return CollectionViewLayoutClass.Alloc()
}

func (cc _CollectionViewLayoutClass) New() CollectionViewLayout {
	rv := objc.CallMethod[CollectionViewLayout](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionViewLayout() CollectionViewLayout {
	return CollectionViewLayoutClass.New()
}

func CollectionViewLayout_New() CollectionViewLayout {
	return CollectionViewLayoutClass.New()
}

func (c_ CollectionViewLayout) Init() CollectionViewLayout {
	rv := objc.CallMethod[CollectionViewLayout](c_, objc.GetSelector("init"))
	return rv
}

func CollectionViewLayout_Init() CollectionViewLayout {
	return CollectionViewLayoutClass.Alloc().Init()
}

func (c_ CollectionViewLayout) PrepareLayout() {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("prepareLayout"))
}

func (c_ CollectionViewLayout) LayoutAttributesForElementsInRect(rect foundation.Rect) []CollectionViewLayoutAttributes {
	rv := objc.CallMethod[[]CollectionViewLayoutAttributes](c_, objc.GetSelector("layoutAttributesForElementsInRect:"), rect)
	// TODO: convert slice items...
	return rv
}

func (c_ CollectionViewLayout) LayoutAttributesForItemAtIndexPath(indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.CallMethod[CollectionViewLayoutAttributes](c_, objc.GetSelector("layoutAttributesForItemAtIndexPath:"), objc.ExtractPtr(indexPath))
	return rv
}

func (c_ CollectionViewLayout) LayoutAttributesForSupplementaryViewOfKindAtIndexPath(elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.CallMethod[CollectionViewLayoutAttributes](c_, objc.GetSelector("layoutAttributesForSupplementaryViewOfKind:atIndexPath:"), elementKind, objc.ExtractPtr(indexPath))
	return rv
}

func (c_ CollectionViewLayout) LayoutAttributesForDecorationViewOfKindAtIndexPath(elementKind CollectionViewDecorationElementKind, indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.CallMethod[CollectionViewLayoutAttributes](c_, objc.GetSelector("layoutAttributesForDecorationViewOfKind:atIndexPath:"), elementKind, objc.ExtractPtr(indexPath))
	return rv
}

func (c_ CollectionViewLayout) LayoutAttributesForDropTargetAtPoint(pointInCollectionView foundation.Point) CollectionViewLayoutAttributes {
	rv := objc.CallMethod[CollectionViewLayoutAttributes](c_, objc.GetSelector("layoutAttributesForDropTargetAtPoint:"), pointInCollectionView)
	return rv
}

func (c_ CollectionViewLayout) LayoutAttributesForInterItemGapBeforeIndexPath(indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.CallMethod[CollectionViewLayoutAttributes](c_, objc.GetSelector("layoutAttributesForInterItemGapBeforeIndexPath:"), objc.ExtractPtr(indexPath))
	return rv
}

func (c_ CollectionViewLayout) TargetContentOffsetForProposedContentOffset(proposedContentOffset foundation.Point) foundation.Point {
	rv := objc.CallMethod[foundation.Point](c_, objc.GetSelector("targetContentOffsetForProposedContentOffset:"), proposedContentOffset)
	return rv
}

func (c_ CollectionViewLayout) TargetContentOffsetForProposedContentOffsetWithScrollingVelocity(proposedContentOffset foundation.Point, velocity foundation.Point) foundation.Point {
	rv := objc.CallMethod[foundation.Point](c_, objc.GetSelector("targetContentOffsetForProposedContentOffset:withScrollingVelocity:"), proposedContentOffset, velocity)
	return rv
}

func (c_ CollectionViewLayout) PrepareForCollectionViewUpdates(updateItems []ICollectionViewUpdateItem) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("prepareForCollectionViewUpdates:"), updateItems)
}

func (c_ CollectionViewLayout) FinalizeCollectionViewUpdates() {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("finalizeCollectionViewUpdates"))
}

func (c_ CollectionViewLayout) IndexPathsToInsertForSupplementaryViewOfKind(elementKind CollectionViewSupplementaryElementKind) foundation.Set {
	rv := objc.CallMethod[foundation.Set](c_, objc.GetSelector("indexPathsToInsertForSupplementaryViewOfKind:"), elementKind)
	return rv
}

func (c_ CollectionViewLayout) IndexPathsToInsertForDecorationViewOfKind(elementKind CollectionViewDecorationElementKind) foundation.Set {
	rv := objc.CallMethod[foundation.Set](c_, objc.GetSelector("indexPathsToInsertForDecorationViewOfKind:"), elementKind)
	return rv
}

func (c_ CollectionViewLayout) InitialLayoutAttributesForAppearingItemAtIndexPath(itemIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.CallMethod[CollectionViewLayoutAttributes](c_, objc.GetSelector("initialLayoutAttributesForAppearingItemAtIndexPath:"), objc.ExtractPtr(itemIndexPath))
	return rv
}

func (c_ CollectionViewLayout) InitialLayoutAttributesForAppearingSupplementaryElementOfKindAtIndexPath(elementKind CollectionViewSupplementaryElementKind, elementIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.CallMethod[CollectionViewLayoutAttributes](c_, objc.GetSelector("initialLayoutAttributesForAppearingSupplementaryElementOfKind:atIndexPath:"), elementKind, objc.ExtractPtr(elementIndexPath))
	return rv
}

func (c_ CollectionViewLayout) InitialLayoutAttributesForAppearingDecorationElementOfKindAtIndexPath(elementKind CollectionViewDecorationElementKind, decorationIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.CallMethod[CollectionViewLayoutAttributes](c_, objc.GetSelector("initialLayoutAttributesForAppearingDecorationElementOfKind:atIndexPath:"), elementKind, objc.ExtractPtr(decorationIndexPath))
	return rv
}

func (c_ CollectionViewLayout) IndexPathsToDeleteForSupplementaryViewOfKind(elementKind CollectionViewSupplementaryElementKind) foundation.Set {
	rv := objc.CallMethod[foundation.Set](c_, objc.GetSelector("indexPathsToDeleteForSupplementaryViewOfKind:"), elementKind)
	return rv
}

func (c_ CollectionViewLayout) IndexPathsToDeleteForDecorationViewOfKind(elementKind CollectionViewDecorationElementKind) foundation.Set {
	rv := objc.CallMethod[foundation.Set](c_, objc.GetSelector("indexPathsToDeleteForDecorationViewOfKind:"), elementKind)
	return rv
}

func (c_ CollectionViewLayout) FinalLayoutAttributesForDisappearingItemAtIndexPath(itemIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.CallMethod[CollectionViewLayoutAttributes](c_, objc.GetSelector("finalLayoutAttributesForDisappearingItemAtIndexPath:"), objc.ExtractPtr(itemIndexPath))
	return rv
}

func (c_ CollectionViewLayout) FinalLayoutAttributesForDisappearingSupplementaryElementOfKindAtIndexPath(elementKind CollectionViewSupplementaryElementKind, elementIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.CallMethod[CollectionViewLayoutAttributes](c_, objc.GetSelector("finalLayoutAttributesForDisappearingSupplementaryElementOfKind:atIndexPath:"), elementKind, objc.ExtractPtr(elementIndexPath))
	return rv
}

func (c_ CollectionViewLayout) FinalLayoutAttributesForDisappearingDecorationElementOfKindAtIndexPath(elementKind CollectionViewDecorationElementKind, decorationIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.CallMethod[CollectionViewLayoutAttributes](c_, objc.GetSelector("finalLayoutAttributesForDisappearingDecorationElementOfKind:atIndexPath:"), elementKind, objc.ExtractPtr(decorationIndexPath))
	return rv
}

func (c_ CollectionViewLayout) InvalidateLayout() {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("invalidateLayout"))
}

func (c_ CollectionViewLayout) InvalidateLayoutWithContext(context ICollectionViewLayoutInvalidationContext) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("invalidateLayoutWithContext:"), objc.ExtractPtr(context))
}

func (c_ CollectionViewLayout) ShouldInvalidateLayoutForBoundsChange(newBounds foundation.Rect) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("shouldInvalidateLayoutForBoundsChange:"), newBounds)
	return rv
}

func (c_ CollectionViewLayout) ShouldInvalidateLayoutForPreferredLayoutAttributesWithOriginalAttributes(preferredAttributes ICollectionViewLayoutAttributes, originalAttributes ICollectionViewLayoutAttributes) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("shouldInvalidateLayoutForPreferredLayoutAttributes:withOriginalAttributes:"), objc.ExtractPtr(preferredAttributes), objc.ExtractPtr(originalAttributes))
	return rv
}

func (c_ CollectionViewLayout) InvalidationContextForBoundsChange(newBounds foundation.Rect) CollectionViewLayoutInvalidationContext {
	rv := objc.CallMethod[CollectionViewLayoutInvalidationContext](c_, objc.GetSelector("invalidationContextForBoundsChange:"), newBounds)
	return rv
}

func (c_ CollectionViewLayout) InvalidationContextForPreferredLayoutAttributesWithOriginalAttributes(preferredAttributes ICollectionViewLayoutAttributes, originalAttributes ICollectionViewLayoutAttributes) CollectionViewLayoutInvalidationContext {
	rv := objc.CallMethod[CollectionViewLayoutInvalidationContext](c_, objc.GetSelector("invalidationContextForPreferredLayoutAttributes:withOriginalAttributes:"), objc.ExtractPtr(preferredAttributes), objc.ExtractPtr(originalAttributes))
	return rv
}

func (c_ CollectionViewLayout) PrepareForAnimatedBoundsChange(oldBounds foundation.Rect) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("prepareForAnimatedBoundsChange:"), oldBounds)
}

func (c_ CollectionViewLayout) FinalizeAnimatedBoundsChange() {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("finalizeAnimatedBoundsChange"))
}

func (c_ CollectionViewLayout) RegisterClassForDecorationViewOfKind(viewClass objc.IClass, elementKind CollectionViewDecorationElementKind) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("registerClass:forDecorationViewOfKind:"), objc.ExtractPtr(viewClass), elementKind)
}

func (c_ CollectionViewLayout) RegisterNibForDecorationViewOfKind(nib INib, elementKind CollectionViewDecorationElementKind) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("registerNib:forDecorationViewOfKind:"), objc.ExtractPtr(nib), elementKind)
}

func (c_ CollectionViewLayout) PrepareForTransitionFromLayout(oldLayout ICollectionViewLayout) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("prepareForTransitionFromLayout:"), objc.ExtractPtr(oldLayout))
}

func (c_ CollectionViewLayout) PrepareForTransitionToLayout(newLayout ICollectionViewLayout) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("prepareForTransitionToLayout:"), objc.ExtractPtr(newLayout))
}

func (c_ CollectionViewLayout) FinalizeLayoutTransition() {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("finalizeLayoutTransition"))
}

func (c_ CollectionViewLayout) CollectionView() CollectionView {
	rv := objc.CallMethod[CollectionView](c_, objc.GetSelector("collectionView"))
	return rv
}

func (cc _CollectionViewLayoutClass) LayoutAttributesClass() objc.Class {
	rv := objc.CallMethod[objc.Class](cc, objc.GetSelector("layoutAttributesClass"))
	return rv
}

func CollectionViewLayout_LayoutAttributesClass() objc.Class {
	return CollectionViewLayoutClass.LayoutAttributesClass()
}

func (c_ CollectionViewLayout) CollectionViewContentSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.GetSelector("collectionViewContentSize"))
	return rv
}

func (cc _CollectionViewLayoutClass) InvalidationContextClass() objc.Class {
	rv := objc.CallMethod[objc.Class](cc, objc.GetSelector("invalidationContextClass"))
	return rv
}

func CollectionViewLayout_InvalidationContextClass() objc.Class {
	return CollectionViewLayoutClass.InvalidationContextClass()
}
