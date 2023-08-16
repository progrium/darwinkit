// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of methods that a delegate implements to provide layout information to a flow layout object in a collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewdelegateflowlayout?language=objc
type PCollectionViewDelegateFlowLayout interface {
	// optional
	CollectionViewLayoutInsetForSectionAtIndex(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.EdgeInsets
	HasCollectionViewLayoutInsetForSectionAtIndex() bool
}

// A delegate implementation builder for the [PCollectionViewDelegateFlowLayout] protocol.
type CollectionViewDelegateFlowLayout struct {
	_CollectionViewLayoutInsetForSectionAtIndex func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.EdgeInsets
}

func (di *CollectionViewDelegateFlowLayout) HasCollectionViewLayoutInsetForSectionAtIndex() bool {
	return di._CollectionViewLayoutInsetForSectionAtIndex != nil
}

// Asks the delegate for the margins to apply to content in the specified section. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewdelegateflowlayout/1402874-collectionview?language=objc
func (di *CollectionViewDelegateFlowLayout) SetCollectionViewLayoutInsetForSectionAtIndex(f func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.EdgeInsets) {
	di._CollectionViewLayoutInsetForSectionAtIndex = f
}

// Asks the delegate for the margins to apply to content in the specified section. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewdelegateflowlayout/1402874-collectionview?language=objc
func (di *CollectionViewDelegateFlowLayout) CollectionViewLayoutInsetForSectionAtIndex(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.EdgeInsets {
	return di._CollectionViewLayoutInsetForSectionAtIndex(collectionView, collectionViewLayout, section)
}

// A concrete type wrapper for the [PCollectionViewDelegateFlowLayout] protocol.
type CollectionViewDelegateFlowLayoutWrapper struct {
	objc.Object
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) HasCollectionViewLayoutInsetForSectionAtIndex() bool {
	return c_.RespondsToSelector(objc.Sel("collectionView:layout:insetForSectionAtIndex:"))
}

// Asks the delegate for the margins to apply to content in the specified section. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewdelegateflowlayout/1402874-collectionview?language=objc
func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionViewLayoutInsetForSectionAtIndex(collectionView ICollectionView, collectionViewLayout ICollectionViewLayout, section int) foundation.EdgeInsets {
	rv := objc.Call[foundation.EdgeInsets](c_, objc.Sel("collectionView:layout:insetForSectionAtIndex:"), objc.Ptr(collectionView), objc.Ptr(collectionViewLayout), section)
	return rv
}
