// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var CollectionViewDiffableDataSourceClass = _CollectionViewDiffableDataSourceClass{objc.GetClass("NSCollectionViewDiffableDataSource")}

type _CollectionViewDiffableDataSourceClass struct {
	objc.Class
}

type ICollectionViewDiffableDataSource interface {
	objc.IObject
	ItemIdentifierForIndexPath(indexPath foundation.IIndexPath) objc.Object
	IndexPathForItemIdentifier(identifier objc.IObject) foundation.IndexPath
	Snapshot() DiffableDataSourceSnapshot
	ApplySnapshotAnimatingDifferences(snapshot IDiffableDataSourceSnapshot, animatingDifferences bool)
	SupplementaryViewProvider() func(param1 CollectionView, param2 string, param3 foundation.IndexPath) View
	SetSupplementaryViewProvider(value func(param1 CollectionView, param2 string, param3 foundation.IndexPath) View)
}

type CollectionViewDiffableDataSource struct {
	objc.Object
}

func MakeCollectionViewDiffableDataSource(ptr unsafe.Pointer) CollectionViewDiffableDataSource {
	return CollectionViewDiffableDataSource{
		Object: objc.MakeObject(ptr),
	}
}

func (c_ CollectionViewDiffableDataSource) InitWithCollectionViewItemProvider(collectionView ICollectionView, itemProvider func(param1 CollectionView, param2 foundation.IndexPath, param3 objc.Object) CollectionViewItem) CollectionViewDiffableDataSource {
	rv := objc.CallMethod[CollectionViewDiffableDataSource](c_, objc.GetSelector("initWithCollectionView:itemProvider:"), objc.ExtractPtr(collectionView), itemProvider)
	return rv
}

func CollectionViewDiffableDataSource_InitWithCollectionViewItemProvider(collectionView ICollectionView, itemProvider func(param1 CollectionView, param2 foundation.IndexPath, param3 objc.Object) CollectionViewItem) CollectionViewDiffableDataSource {
	return CollectionViewDiffableDataSourceClass.Alloc().InitWithCollectionViewItemProvider(collectionView, itemProvider)
}

func (cc _CollectionViewDiffableDataSourceClass) Alloc() CollectionViewDiffableDataSource {
	rv := objc.CallMethod[CollectionViewDiffableDataSource](cc, objc.GetSelector("alloc"))
	return rv
}

func CollectionViewDiffableDataSource_Alloc() CollectionViewDiffableDataSource {
	return CollectionViewDiffableDataSourceClass.Alloc()
}

func (cc _CollectionViewDiffableDataSourceClass) New() CollectionViewDiffableDataSource {
	rv := objc.CallMethod[CollectionViewDiffableDataSource](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionViewDiffableDataSource() CollectionViewDiffableDataSource {
	return CollectionViewDiffableDataSourceClass.New()
}

func CollectionViewDiffableDataSource_New() CollectionViewDiffableDataSource {
	return CollectionViewDiffableDataSourceClass.New()
}

func (c_ CollectionViewDiffableDataSource) Init() CollectionViewDiffableDataSource {
	rv := objc.CallMethod[CollectionViewDiffableDataSource](c_, objc.GetSelector("init"))
	return rv
}

func CollectionViewDiffableDataSource_Init() CollectionViewDiffableDataSource {
	return CollectionViewDiffableDataSourceClass.Alloc().Init()
}

func (c_ CollectionViewDiffableDataSource) ItemIdentifierForIndexPath(indexPath foundation.IIndexPath) objc.Object {
	rv := objc.CallMethod[objc.Object](c_, objc.GetSelector("itemIdentifierForIndexPath:"), objc.ExtractPtr(indexPath))
	return rv
}

func (c_ CollectionViewDiffableDataSource) IndexPathForItemIdentifier(identifier objc.IObject) foundation.IndexPath {
	rv := objc.CallMethod[foundation.IndexPath](c_, objc.GetSelector("indexPathForItemIdentifier:"), objc.ExtractPtr(identifier))
	return rv
}

func (c_ CollectionViewDiffableDataSource) Snapshot() DiffableDataSourceSnapshot {
	rv := objc.CallMethod[DiffableDataSourceSnapshot](c_, objc.GetSelector("snapshot"))
	return rv
}

func (c_ CollectionViewDiffableDataSource) ApplySnapshotAnimatingDifferences(snapshot IDiffableDataSourceSnapshot, animatingDifferences bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("applySnapshot:animatingDifferences:"), objc.ExtractPtr(snapshot), animatingDifferences)
}

func (c_ CollectionViewDiffableDataSource) SupplementaryViewProvider() func(param1 CollectionView, param2 string, param3 foundation.IndexPath) View {
	rv := objc.CallMethod[func(param1 CollectionView, param2 string, param3 foundation.IndexPath) View](c_, objc.GetSelector("supplementaryViewProvider"))
	return rv
}

func (c_ CollectionViewDiffableDataSource) SetSupplementaryViewProvider(value func(param1 CollectionView, param2 string, param3 foundation.IndexPath) View) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setSupplementaryViewProvider:"), value)
}
