// auto-generated code, do not modify
package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type ICollectionViewDataSource interface {
	ImplementsNumberOfSectionsInCollectionView() bool
	// optional
	NumberOfSectionsInCollectionView(collectionView CollectionView) int
	// required
	CollectionViewNumberOfItemsInSection(collectionView CollectionView, section int) int
	// required
	CollectionViewItemForRepresentedObjectAtIndexPath(collectionView CollectionView, indexPath foundation.IndexPath) ICollectionViewItem
	ImplementsCollectionViewViewForSupplementaryElementOfKindAtIndexPath() bool
	// optional
	CollectionViewViewForSupplementaryElementOfKindAtIndexPath(collectionView CollectionView, kind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath) IView
}

type CollectionViewDataSourceWrapper struct {
	objc.Object
}

func (c_ CollectionViewDataSourceWrapper) ImplementsNumberOfSectionsInCollectionView() bool {
	return c_.RespondsToSelector(objc.GetSelector("numberOfSectionsInCollectionView:"))
}

func (c_ CollectionViewDataSourceWrapper) NumberOfSectionsInCollectionView(collectionView ICollectionView) int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("numberOfSectionsInCollectionView:"), objc.ExtractPtr(collectionView))
	return rv
}

func (c_ CollectionViewDataSourceWrapper) CollectionViewNumberOfItemsInSection(collectionView ICollectionView, section int) int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("collectionView:numberOfItemsInSection:"), objc.ExtractPtr(collectionView), section)
	return rv
}

func (c_ CollectionViewDataSourceWrapper) CollectionViewItemForRepresentedObjectAtIndexPath(collectionView ICollectionView, indexPath foundation.IIndexPath) CollectionViewItem {
	rv := objc.CallMethod[CollectionViewItem](c_, objc.GetSelector("collectionView:itemForRepresentedObjectAtIndexPath:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexPath))
	return rv
}

func (c_ CollectionViewDataSourceWrapper) ImplementsCollectionViewViewForSupplementaryElementOfKindAtIndexPath() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:viewForSupplementaryElementOfKind:atIndexPath:"))
}

func (c_ CollectionViewDataSourceWrapper) CollectionViewViewForSupplementaryElementOfKindAtIndexPath(collectionView ICollectionView, kind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) View {
	rv := objc.CallMethod[View](c_, objc.GetSelector("collectionView:viewForSupplementaryElementOfKind:atIndexPath:"), objc.ExtractPtr(collectionView), kind, objc.ExtractPtr(indexPath))
	return rv
}
