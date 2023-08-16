// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewprefetching?language=objc
type PCollectionViewPrefetching interface {
	// optional
	CollectionViewPrefetchItemsAtIndexPaths(collectionView CollectionView, indexPaths []foundation.IndexPath)
	HasCollectionViewPrefetchItemsAtIndexPaths() bool
}

// A concrete type wrapper for the [PCollectionViewPrefetching] protocol.
type CollectionViewPrefetchingWrapper struct {
	objc.Object
}

func (c_ CollectionViewPrefetchingWrapper) HasCollectionViewPrefetchItemsAtIndexPaths() bool {
	return c_.RespondsToSelector(objc.Sel("collectionView:prefetchItemsAtIndexPaths:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewprefetching/2879293-collectionview?language=objc
func (c_ CollectionViewPrefetchingWrapper) CollectionViewPrefetchItemsAtIndexPaths(collectionView ICollectionView, indexPaths []foundation.IIndexPath) {
	objc.Call[objc.Void](c_, objc.Sel("collectionView:prefetchItemsAtIndexPaths:"), objc.Ptr(collectionView), indexPaths)
}
