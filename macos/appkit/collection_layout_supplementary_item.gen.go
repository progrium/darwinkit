// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var CollectionLayoutSupplementaryItemClass = _CollectionLayoutSupplementaryItemClass{objc.GetClass("NSCollectionLayoutSupplementaryItem")}

type _CollectionLayoutSupplementaryItemClass struct {
	objc.Class
}

type ICollectionLayoutSupplementaryItem interface {
	ICollectionLayoutItem
	ItemAnchor() CollectionLayoutAnchor
	ContainerAnchor() CollectionLayoutAnchor
	ElementKind() string
	ZIndex() int
	SetZIndex(value int)
}

type CollectionLayoutSupplementaryItem struct {
	CollectionLayoutItem
}

func MakeCollectionLayoutSupplementaryItem(ptr unsafe.Pointer) CollectionLayoutSupplementaryItem {
	return CollectionLayoutSupplementaryItem{
		CollectionLayoutItem: MakeCollectionLayoutItem(ptr),
	}
}

func (cc _CollectionLayoutSupplementaryItemClass) SupplementaryItemWithLayoutSizeElementKindContainerAnchor(layoutSize ICollectionLayoutSize, elementKind string, containerAnchor ICollectionLayoutAnchor) CollectionLayoutSupplementaryItem {
	rv := objc.CallMethod[CollectionLayoutSupplementaryItem](cc, objc.GetSelector("supplementaryItemWithLayoutSize:elementKind:containerAnchor:"), objc.ExtractPtr(layoutSize), elementKind, objc.ExtractPtr(containerAnchor))
	return rv
}

func CollectionLayoutSupplementaryItem_SupplementaryItemWithLayoutSizeElementKindContainerAnchor(layoutSize ICollectionLayoutSize, elementKind string, containerAnchor ICollectionLayoutAnchor) CollectionLayoutSupplementaryItem {
	return CollectionLayoutSupplementaryItemClass.SupplementaryItemWithLayoutSizeElementKindContainerAnchor(layoutSize, elementKind, containerAnchor)
}

func (cc _CollectionLayoutSupplementaryItemClass) SupplementaryItemWithLayoutSizeElementKindContainerAnchorItemAnchor(layoutSize ICollectionLayoutSize, elementKind string, containerAnchor ICollectionLayoutAnchor, itemAnchor ICollectionLayoutAnchor) CollectionLayoutSupplementaryItem {
	rv := objc.CallMethod[CollectionLayoutSupplementaryItem](cc, objc.GetSelector("supplementaryItemWithLayoutSize:elementKind:containerAnchor:itemAnchor:"), objc.ExtractPtr(layoutSize), elementKind, objc.ExtractPtr(containerAnchor), objc.ExtractPtr(itemAnchor))
	return rv
}

func CollectionLayoutSupplementaryItem_SupplementaryItemWithLayoutSizeElementKindContainerAnchorItemAnchor(layoutSize ICollectionLayoutSize, elementKind string, containerAnchor ICollectionLayoutAnchor, itemAnchor ICollectionLayoutAnchor) CollectionLayoutSupplementaryItem {
	return CollectionLayoutSupplementaryItemClass.SupplementaryItemWithLayoutSizeElementKindContainerAnchorItemAnchor(layoutSize, elementKind, containerAnchor, itemAnchor)
}

func (cc _CollectionLayoutSupplementaryItemClass) ItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutSupplementaryItem {
	rv := objc.CallMethod[CollectionLayoutSupplementaryItem](cc, objc.GetSelector("itemWithLayoutSize:"), objc.ExtractPtr(layoutSize))
	return rv
}

func CollectionLayoutSupplementaryItem_ItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutSupplementaryItem {
	return CollectionLayoutSupplementaryItemClass.ItemWithLayoutSize(layoutSize)
}

func (cc _CollectionLayoutSupplementaryItemClass) ItemWithLayoutSizeSupplementaryItems(layoutSize ICollectionLayoutSize, supplementaryItems []ICollectionLayoutSupplementaryItem) CollectionLayoutSupplementaryItem {
	rv := objc.CallMethod[CollectionLayoutSupplementaryItem](cc, objc.GetSelector("itemWithLayoutSize:supplementaryItems:"), objc.ExtractPtr(layoutSize), supplementaryItems)
	return rv
}

func CollectionLayoutSupplementaryItem_ItemWithLayoutSizeSupplementaryItems(layoutSize ICollectionLayoutSize, supplementaryItems []ICollectionLayoutSupplementaryItem) CollectionLayoutSupplementaryItem {
	return CollectionLayoutSupplementaryItemClass.ItemWithLayoutSizeSupplementaryItems(layoutSize, supplementaryItems)
}

func (cc _CollectionLayoutSupplementaryItemClass) Alloc() CollectionLayoutSupplementaryItem {
	rv := objc.CallMethod[CollectionLayoutSupplementaryItem](cc, objc.GetSelector("alloc"))
	return rv
}

func CollectionLayoutSupplementaryItem_Alloc() CollectionLayoutSupplementaryItem {
	return CollectionLayoutSupplementaryItemClass.Alloc()
}

func (cc _CollectionLayoutSupplementaryItemClass) New() CollectionLayoutSupplementaryItem {
	rv := objc.CallMethod[CollectionLayoutSupplementaryItem](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutSupplementaryItem() CollectionLayoutSupplementaryItem {
	return CollectionLayoutSupplementaryItemClass.New()
}

func CollectionLayoutSupplementaryItem_New() CollectionLayoutSupplementaryItem {
	return CollectionLayoutSupplementaryItemClass.New()
}

func (c_ CollectionLayoutSupplementaryItem) Init() CollectionLayoutSupplementaryItem {
	rv := objc.CallMethod[CollectionLayoutSupplementaryItem](c_, objc.GetSelector("init"))
	return rv
}

func CollectionLayoutSupplementaryItem_Init() CollectionLayoutSupplementaryItem {
	return CollectionLayoutSupplementaryItemClass.Alloc().Init()
}

func (c_ CollectionLayoutSupplementaryItem) ItemAnchor() CollectionLayoutAnchor {
	rv := objc.CallMethod[CollectionLayoutAnchor](c_, objc.GetSelector("itemAnchor"))
	return rv
}

func (c_ CollectionLayoutSupplementaryItem) ContainerAnchor() CollectionLayoutAnchor {
	rv := objc.CallMethod[CollectionLayoutAnchor](c_, objc.GetSelector("containerAnchor"))
	return rv
}

func (c_ CollectionLayoutSupplementaryItem) ElementKind() string {
	rv := objc.CallMethod[string](c_, objc.GetSelector("elementKind"))
	return rv
}

func (c_ CollectionLayoutSupplementaryItem) ZIndex() int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("zIndex"))
	return rv
}

func (c_ CollectionLayoutSupplementaryItem) SetZIndex(value int) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setZIndex:"), value)
}
