// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A set of methods that a data source object implements to provide the information and view objects that a collection view requires to present content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewdatasource?language=objc
type PCollectionViewDataSource interface {
	// optional
	NumberOfSectionsInCollectionView(collectionView CollectionView) int
	HasNumberOfSectionsInCollectionView() bool

	// optional
	CollectionViewNumberOfItemsInSection(collectionView CollectionView, section int) int
	HasCollectionViewNumberOfItemsInSection() bool
}

// A concrete type wrapper for the [PCollectionViewDataSource] protocol.
type CollectionViewDataSourceWrapper struct {
	objc.Object
}

func (c_ CollectionViewDataSourceWrapper) HasNumberOfSectionsInCollectionView() bool {
	return c_.RespondsToSelector(objc.Sel("numberOfSectionsInCollectionView:"))
}

// Asks your data source object to provide the total number of sections. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewdatasource/1525901-numberofsectionsincollectionview?language=objc
func (c_ CollectionViewDataSourceWrapper) NumberOfSectionsInCollectionView(collectionView ICollectionView) int {
	rv := objc.Call[int](c_, objc.Sel("numberOfSectionsInCollectionView:"), objc.Ptr(collectionView))
	return rv
}

func (c_ CollectionViewDataSourceWrapper) HasCollectionViewNumberOfItemsInSection() bool {
	return c_.RespondsToSelector(objc.Sel("collectionView:numberOfItemsInSection:"))
}

// Asks your data source object to provide the number of items in the specified section. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewdatasource/1525594-collectionview?language=objc
func (c_ CollectionViewDataSourceWrapper) CollectionViewNumberOfItemsInSection(collectionView ICollectionView, section int) int {
	rv := objc.Call[int](c_, objc.Sel("collectionView:numberOfItemsInSection:"), objc.Ptr(collectionView), section)
	return rv
}
