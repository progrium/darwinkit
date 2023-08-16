// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CollectionViewDiffableDataSource] class.
var CollectionViewDiffableDataSourceClass = _CollectionViewDiffableDataSourceClass{objc.GetClass("NSCollectionViewDiffableDataSource")}

type _CollectionViewDiffableDataSourceClass struct {
	objc.Class
}

// An interface definition for the [CollectionViewDiffableDataSource] class.
type ICollectionViewDiffableDataSource interface {
	objc.IObject
	ItemIdentifierForIndexPath(indexPath foundation.IIndexPath) objc.Object
	IndexPathForItemIdentifier(identifier objc.IObject) foundation.IndexPath
	SupplementaryViewProvider() CollectionViewDiffableDataSourceSupplementaryViewProvider
	SetSupplementaryViewProvider(value CollectionViewDiffableDataSourceSupplementaryViewProvider)
}

// The object you use to manage data and provide items for a collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewdiffabledatasource?language=objc
type CollectionViewDiffableDataSource struct {
	objc.Object
}

func CollectionViewDiffableDataSourceFrom(ptr unsafe.Pointer) CollectionViewDiffableDataSource {
	return CollectionViewDiffableDataSource{
		Object: objc.ObjectFrom(ptr),
	}
}

func (c_ CollectionViewDiffableDataSource) InitWithCollectionViewItemProvider(collectionView ICollectionView, itemProvider CollectionViewDiffableDataSourceItemProvider) CollectionViewDiffableDataSource {
	rv := objc.Call[CollectionViewDiffableDataSource](c_, objc.Sel("initWithCollectionView:itemProvider:"), objc.Ptr(collectionView), itemProvider)
	return rv
}

// Creates a diffable data source with the specified item provider, and connects it to the specified collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewdiffabledatasource/3281818-initwithcollectionview?language=objc
func CollectionViewDiffableDataSource_InitWithCollectionViewItemProvider(collectionView ICollectionView, itemProvider CollectionViewDiffableDataSourceItemProvider) CollectionViewDiffableDataSource {
	return CollectionViewDiffableDataSourceClass.Alloc().InitWithCollectionViewItemProvider(collectionView, itemProvider)
}

func (cc _CollectionViewDiffableDataSourceClass) Alloc() CollectionViewDiffableDataSource {
	rv := objc.Call[CollectionViewDiffableDataSource](cc, objc.Sel("alloc"))
	return rv
}

func CollectionViewDiffableDataSource_Alloc() CollectionViewDiffableDataSource {
	return CollectionViewDiffableDataSourceClass.Alloc()
}

func (cc _CollectionViewDiffableDataSourceClass) New() CollectionViewDiffableDataSource {
	rv := objc.Call[CollectionViewDiffableDataSource](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionViewDiffableDataSource() CollectionViewDiffableDataSource {
	return CollectionViewDiffableDataSourceClass.New()
}

func (c_ CollectionViewDiffableDataSource) Init() CollectionViewDiffableDataSource {
	rv := objc.Call[CollectionViewDiffableDataSource](c_, objc.Sel("init"))
	return rv
}

// Returns an identifier for the item at the specified index path in the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewdiffabledatasource/3281819-itemidentifierforindexpath?language=objc
func (c_ CollectionViewDiffableDataSource) ItemIdentifierForIndexPath(indexPath foundation.IIndexPath) objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("itemIdentifierForIndexPath:"), objc.Ptr(indexPath))
	return rv
}

// Returns an index path for the item with the specified identifier in the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewdiffabledatasource/3281817-indexpathforitemidentifier?language=objc
func (c_ CollectionViewDiffableDataSource) IndexPathForItemIdentifier(identifier objc.IObject) foundation.IndexPath {
	rv := objc.Call[foundation.IndexPath](c_, objc.Sel("indexPathForItemIdentifier:"), objc.Ptr(identifier))
	return rv
}

// The closure that configures and returns the collection view’s supplementary views, such as headers and footers, from the diffable data source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewdiffabledatasource/3281821-supplementaryviewprovider?language=objc
func (c_ CollectionViewDiffableDataSource) SupplementaryViewProvider() CollectionViewDiffableDataSourceSupplementaryViewProvider {
	rv := objc.Call[CollectionViewDiffableDataSourceSupplementaryViewProvider](c_, objc.Sel("supplementaryViewProvider"))
	return rv
}

// The closure that configures and returns the collection view’s supplementary views, such as headers and footers, from the diffable data source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewdiffabledatasource/3281821-supplementaryviewprovider?language=objc
func (c_ CollectionViewDiffableDataSource) SetSupplementaryViewProvider(value CollectionViewDiffableDataSourceSupplementaryViewProvider) {
	objc.Call[objc.Void](c_, objc.Sel("setSupplementaryViewProvider:"), value)
}
