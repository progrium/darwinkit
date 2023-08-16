// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of methods that you use to manage the behavior of a collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewdelegate?language=objc
type PCollectionViewDelegate interface {
	// optional
	CollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath(collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath)
	HasCollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath() bool
}

// A delegate implementation builder for the [PCollectionViewDelegate] protocol.
type CollectionViewDelegate struct {
	_CollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath func(collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath)
}

func (di *CollectionViewDelegate) HasCollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath() bool {
	return di._CollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath != nil
}

// Notifies the delegate that the specified item was removed from the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewdelegate/1528183-collectionview?language=objc
func (di *CollectionViewDelegate) SetCollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath(f func(collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath)) {
	di._CollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath = f
}

// Notifies the delegate that the specified item was removed from the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewdelegate/1528183-collectionview?language=objc
func (di *CollectionViewDelegate) CollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath(collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath) {
	di._CollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath(collectionView, item, indexPath)
}

// A concrete type wrapper for the [PCollectionViewDelegate] protocol.
type CollectionViewDelegateWrapper struct {
	objc.Object
}

func (c_ CollectionViewDelegateWrapper) HasCollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath() bool {
	return c_.RespondsToSelector(objc.Sel("collectionView:didEndDisplayingItem:forRepresentedObjectAtIndexPath:"))
}

// Notifies the delegate that the specified item was removed from the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewdelegate/1528183-collectionview?language=objc
func (c_ CollectionViewDelegateWrapper) CollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath(collectionView ICollectionView, item ICollectionViewItem, indexPath foundation.IIndexPath) {
	objc.Call[objc.Void](c_, objc.Sel("collectionView:didEndDisplayingItem:forRepresentedObjectAtIndexPath:"), objc.Ptr(collectionView), objc.Ptr(item), objc.Ptr(indexPath))
}
