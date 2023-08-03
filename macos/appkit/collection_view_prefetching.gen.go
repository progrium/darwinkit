// auto-generated code, do not modify
package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type ICollectionViewPrefetching interface {
	ImplementsCollectionViewCancelPrefetchingForItemsAtIndexPaths() bool
	// optional
	CollectionViewCancelPrefetchingForItemsAtIndexPaths(collectionView CollectionView, indexPaths []foundation.IndexPath)
	// required
	CollectionViewPrefetchItemsAtIndexPaths(collectionView CollectionView, indexPaths []foundation.IndexPath)
}

type CollectionViewPrefetchingWrapper struct {
	objc.Object
}

func (c_ CollectionViewPrefetchingWrapper) ImplementsCollectionViewCancelPrefetchingForItemsAtIndexPaths() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:cancelPrefetchingForItemsAtIndexPaths:"))
}

func (c_ CollectionViewPrefetchingWrapper) CollectionViewCancelPrefetchingForItemsAtIndexPaths(collectionView ICollectionView, indexPaths []foundation.IIndexPath) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:cancelPrefetchingForItemsAtIndexPaths:"), objc.ExtractPtr(collectionView), indexPaths)
}

func (c_ CollectionViewPrefetchingWrapper) CollectionViewPrefetchItemsAtIndexPaths(collectionView ICollectionView, indexPaths []foundation.IIndexPath) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:prefetchItemsAtIndexPaths:"), objc.ExtractPtr(collectionView), indexPaths)
}
