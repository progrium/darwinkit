// auto-generated code, do not modify
package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type ICollectionViewDelegateFlowLayout interface {
	ICollectionViewDelegate
	ImplementsCollectionViewLayoutSizeForItemAtIndexPath() bool
	// optional
	CollectionViewLayoutSizeForItemAtIndexPath(collectionView CollectionView, collectionViewLayout CollectionViewLayout, indexPath foundation.IndexPath) foundation.Size
	ImplementsCollectionViewLayoutInsetForSectionAtIndex() bool
	// optional
	CollectionViewLayoutInsetForSectionAtIndex(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.EdgeInsets
	ImplementsCollectionViewLayoutMinimumLineSpacingForSectionAtIndex() bool
	// optional
	CollectionViewLayoutMinimumLineSpacingForSectionAtIndex(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) float64
	ImplementsCollectionViewLayoutMinimumInteritemSpacingForSectionAtIndex() bool
	// optional
	CollectionViewLayoutMinimumInteritemSpacingForSectionAtIndex(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) float64
	ImplementsCollectionViewLayoutReferenceSizeForHeaderInSection() bool
	// optional
	CollectionViewLayoutReferenceSizeForHeaderInSection(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.Size
	ImplementsCollectionViewLayoutReferenceSizeForFooterInSection() bool
	// optional
	CollectionViewLayoutReferenceSizeForFooterInSection(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.Size
}

type CollectionViewDelegateFlowLayout struct {
	CollectionViewDelegate
	_CollectionViewLayoutSizeForItemAtIndexPath                   func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, indexPath foundation.IndexPath) foundation.Size
	_CollectionViewLayoutInsetForSectionAtIndex                   func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.EdgeInsets
	_CollectionViewLayoutMinimumLineSpacingForSectionAtIndex      func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) float64
	_CollectionViewLayoutMinimumInteritemSpacingForSectionAtIndex func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) float64
	_CollectionViewLayoutReferenceSizeForHeaderInSection          func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.Size
	_CollectionViewLayoutReferenceSizeForFooterInSection          func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.Size
}

func (di *CollectionViewDelegateFlowLayout) ImplementsCollectionViewLayoutSizeForItemAtIndexPath() bool {
	return di._CollectionViewLayoutSizeForItemAtIndexPath != nil
}

func (di *CollectionViewDelegateFlowLayout) SetCollectionViewLayoutSizeForItemAtIndexPath(f func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, indexPath foundation.IndexPath) foundation.Size) {
	di._CollectionViewLayoutSizeForItemAtIndexPath = f
}

func (di *CollectionViewDelegateFlowLayout) CollectionViewLayoutSizeForItemAtIndexPath(collectionView CollectionView, collectionViewLayout CollectionViewLayout, indexPath foundation.IndexPath) foundation.Size {
	return di._CollectionViewLayoutSizeForItemAtIndexPath(collectionView, collectionViewLayout, indexPath)
}
func (di *CollectionViewDelegateFlowLayout) ImplementsCollectionViewLayoutInsetForSectionAtIndex() bool {
	return di._CollectionViewLayoutInsetForSectionAtIndex != nil
}

func (di *CollectionViewDelegateFlowLayout) SetCollectionViewLayoutInsetForSectionAtIndex(f func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.EdgeInsets) {
	di._CollectionViewLayoutInsetForSectionAtIndex = f
}

func (di *CollectionViewDelegateFlowLayout) CollectionViewLayoutInsetForSectionAtIndex(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.EdgeInsets {
	return di._CollectionViewLayoutInsetForSectionAtIndex(collectionView, collectionViewLayout, section)
}
func (di *CollectionViewDelegateFlowLayout) ImplementsCollectionViewLayoutMinimumLineSpacingForSectionAtIndex() bool {
	return di._CollectionViewLayoutMinimumLineSpacingForSectionAtIndex != nil
}

func (di *CollectionViewDelegateFlowLayout) SetCollectionViewLayoutMinimumLineSpacingForSectionAtIndex(f func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) float64) {
	di._CollectionViewLayoutMinimumLineSpacingForSectionAtIndex = f
}

func (di *CollectionViewDelegateFlowLayout) CollectionViewLayoutMinimumLineSpacingForSectionAtIndex(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) float64 {
	return di._CollectionViewLayoutMinimumLineSpacingForSectionAtIndex(collectionView, collectionViewLayout, section)
}
func (di *CollectionViewDelegateFlowLayout) ImplementsCollectionViewLayoutMinimumInteritemSpacingForSectionAtIndex() bool {
	return di._CollectionViewLayoutMinimumInteritemSpacingForSectionAtIndex != nil
}

func (di *CollectionViewDelegateFlowLayout) SetCollectionViewLayoutMinimumInteritemSpacingForSectionAtIndex(f func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) float64) {
	di._CollectionViewLayoutMinimumInteritemSpacingForSectionAtIndex = f
}

func (di *CollectionViewDelegateFlowLayout) CollectionViewLayoutMinimumInteritemSpacingForSectionAtIndex(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) float64 {
	return di._CollectionViewLayoutMinimumInteritemSpacingForSectionAtIndex(collectionView, collectionViewLayout, section)
}
func (di *CollectionViewDelegateFlowLayout) ImplementsCollectionViewLayoutReferenceSizeForHeaderInSection() bool {
	return di._CollectionViewLayoutReferenceSizeForHeaderInSection != nil
}

func (di *CollectionViewDelegateFlowLayout) SetCollectionViewLayoutReferenceSizeForHeaderInSection(f func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.Size) {
	di._CollectionViewLayoutReferenceSizeForHeaderInSection = f
}

func (di *CollectionViewDelegateFlowLayout) CollectionViewLayoutReferenceSizeForHeaderInSection(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.Size {
	return di._CollectionViewLayoutReferenceSizeForHeaderInSection(collectionView, collectionViewLayout, section)
}
func (di *CollectionViewDelegateFlowLayout) ImplementsCollectionViewLayoutReferenceSizeForFooterInSection() bool {
	return di._CollectionViewLayoutReferenceSizeForFooterInSection != nil
}

func (di *CollectionViewDelegateFlowLayout) SetCollectionViewLayoutReferenceSizeForFooterInSection(f func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.Size) {
	di._CollectionViewLayoutReferenceSizeForFooterInSection = f
}

func (di *CollectionViewDelegateFlowLayout) CollectionViewLayoutReferenceSizeForFooterInSection(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.Size {
	return di._CollectionViewLayoutReferenceSizeForFooterInSection(collectionView, collectionViewLayout, section)
}

type CollectionViewDelegateFlowLayoutWrapper struct {
	CollectionViewDelegateWrapper
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) ImplementsCollectionViewLayoutSizeForItemAtIndexPath() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:layout:sizeForItemAtIndexPath:"))
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionViewLayoutSizeForItemAtIndexPath(collectionView ICollectionView, collectionViewLayout ICollectionViewLayout, indexPath foundation.IIndexPath) foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.GetSelector("collectionView:layout:sizeForItemAtIndexPath:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(collectionViewLayout), objc.ExtractPtr(indexPath))
	return rv
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) ImplementsCollectionViewLayoutInsetForSectionAtIndex() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:layout:insetForSectionAtIndex:"))
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionViewLayoutInsetForSectionAtIndex(collectionView ICollectionView, collectionViewLayout ICollectionViewLayout, section int) foundation.EdgeInsets {
	rv := objc.CallMethod[foundation.EdgeInsets](c_, objc.GetSelector("collectionView:layout:insetForSectionAtIndex:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(collectionViewLayout), section)
	return rv
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) ImplementsCollectionViewLayoutMinimumLineSpacingForSectionAtIndex() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:layout:minimumLineSpacingForSectionAtIndex:"))
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionViewLayoutMinimumLineSpacingForSectionAtIndex(collectionView ICollectionView, collectionViewLayout ICollectionViewLayout, section int) float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("collectionView:layout:minimumLineSpacingForSectionAtIndex:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(collectionViewLayout), section)
	return rv
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) ImplementsCollectionViewLayoutMinimumInteritemSpacingForSectionAtIndex() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:layout:minimumInteritemSpacingForSectionAtIndex:"))
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionViewLayoutMinimumInteritemSpacingForSectionAtIndex(collectionView ICollectionView, collectionViewLayout ICollectionViewLayout, section int) float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("collectionView:layout:minimumInteritemSpacingForSectionAtIndex:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(collectionViewLayout), section)
	return rv
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) ImplementsCollectionViewLayoutReferenceSizeForHeaderInSection() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:layout:referenceSizeForHeaderInSection:"))
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionViewLayoutReferenceSizeForHeaderInSection(collectionView ICollectionView, collectionViewLayout ICollectionViewLayout, section int) foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.GetSelector("collectionView:layout:referenceSizeForHeaderInSection:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(collectionViewLayout), section)
	return rv
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) ImplementsCollectionViewLayoutReferenceSizeForFooterInSection() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:layout:referenceSizeForFooterInSection:"))
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionViewLayoutReferenceSizeForFooterInSection(collectionView ICollectionView, collectionViewLayout ICollectionViewLayout, section int) foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.GetSelector("collectionView:layout:referenceSizeForFooterInSection:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(collectionViewLayout), section)
	return rv
}
