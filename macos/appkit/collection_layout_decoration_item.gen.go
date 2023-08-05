// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var CollectionLayoutDecorationItemClass = _CollectionLayoutDecorationItemClass{objc.GetClass("NSCollectionLayoutDecorationItem")}

type _CollectionLayoutDecorationItemClass struct {
	objc.Class
}

type ICollectionLayoutDecorationItem interface {
	ICollectionLayoutItem
	ElementKind() string
	ZIndex() int
	SetZIndex(value int)
}

type CollectionLayoutDecorationItem struct {
	CollectionLayoutItem
}

func MakeCollectionLayoutDecorationItem(ptr unsafe.Pointer) CollectionLayoutDecorationItem {
	return CollectionLayoutDecorationItem{
		CollectionLayoutItem: MakeCollectionLayoutItem(ptr),
	}
}

func (cc _CollectionLayoutDecorationItemClass) BackgroundDecorationItemWithElementKind(elementKind string) CollectionLayoutDecorationItem {
	rv := objc.CallMethod[CollectionLayoutDecorationItem](cc, objc.GetSelector("backgroundDecorationItemWithElementKind:"), elementKind)
	return rv
}

func CollectionLayoutDecorationItem_BackgroundDecorationItemWithElementKind(elementKind string) CollectionLayoutDecorationItem {
	return CollectionLayoutDecorationItemClass.BackgroundDecorationItemWithElementKind(elementKind)
}

func (cc _CollectionLayoutDecorationItemClass) ItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutDecorationItem {
	rv := objc.CallMethod[CollectionLayoutDecorationItem](cc, objc.GetSelector("itemWithLayoutSize:"), objc.ExtractPtr(layoutSize))
	return rv
}

func CollectionLayoutDecorationItem_ItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutDecorationItem {
	return CollectionLayoutDecorationItemClass.ItemWithLayoutSize(layoutSize)
}

func (cc _CollectionLayoutDecorationItemClass) ItemWithLayoutSizeSupplementaryItems(layoutSize ICollectionLayoutSize, supplementaryItems []ICollectionLayoutSupplementaryItem) CollectionLayoutDecorationItem {
	rv := objc.CallMethod[CollectionLayoutDecorationItem](cc, objc.GetSelector("itemWithLayoutSize:supplementaryItems:"), objc.ExtractPtr(layoutSize), supplementaryItems)
	return rv
}

func CollectionLayoutDecorationItem_ItemWithLayoutSizeSupplementaryItems(layoutSize ICollectionLayoutSize, supplementaryItems []ICollectionLayoutSupplementaryItem) CollectionLayoutDecorationItem {
	return CollectionLayoutDecorationItemClass.ItemWithLayoutSizeSupplementaryItems(layoutSize, supplementaryItems)
}

func (cc _CollectionLayoutDecorationItemClass) Alloc() CollectionLayoutDecorationItem {
	rv := objc.CallMethod[CollectionLayoutDecorationItem](cc, objc.GetSelector("alloc"))
	return rv
}

func CollectionLayoutDecorationItem_Alloc() CollectionLayoutDecorationItem {
	return CollectionLayoutDecorationItemClass.Alloc()
}

func (cc _CollectionLayoutDecorationItemClass) New() CollectionLayoutDecorationItem {
	rv := objc.CallMethod[CollectionLayoutDecorationItem](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutDecorationItem() CollectionLayoutDecorationItem {
	return CollectionLayoutDecorationItemClass.New()
}

func CollectionLayoutDecorationItem_New() CollectionLayoutDecorationItem {
	return CollectionLayoutDecorationItemClass.New()
}

func (c_ CollectionLayoutDecorationItem) Init() CollectionLayoutDecorationItem {
	rv := objc.CallMethod[CollectionLayoutDecorationItem](c_, objc.GetSelector("init"))
	return rv
}

func CollectionLayoutDecorationItem_Init() CollectionLayoutDecorationItem {
	return CollectionLayoutDecorationItemClass.Alloc().Init()
}

func (c_ CollectionLayoutDecorationItem) ElementKind() string {
	rv := objc.CallMethod[string](c_, objc.GetSelector("elementKind"))
	return rv
}

func (c_ CollectionLayoutDecorationItem) ZIndex() int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("zIndex"))
	return rv
}

func (c_ CollectionLayoutDecorationItem) SetZIndex(value int) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setZIndex:"), value)
}
