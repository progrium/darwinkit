// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var CollectionLayoutItemClass = _CollectionLayoutItemClass{objc.GetClass("NSCollectionLayoutItem")}

type _CollectionLayoutItemClass struct {
	objc.Class
}

type ICollectionLayoutItem interface {
	objc.IObject
	LayoutSize() CollectionLayoutSize
	SupplementaryItems() []CollectionLayoutSupplementaryItem
	EdgeSpacing() CollectionLayoutEdgeSpacing
	SetEdgeSpacing(value ICollectionLayoutEdgeSpacing)
	ContentInsets() DirectionalEdgeInsets
	SetContentInsets(value DirectionalEdgeInsets)
}

type CollectionLayoutItem struct {
	objc.Object
}

func MakeCollectionLayoutItem(ptr unsafe.Pointer) CollectionLayoutItem {
	return CollectionLayoutItem{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionLayoutItemClass) ItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutItem {
	rv := objc.CallMethod[CollectionLayoutItem](cc, objc.GetSelector("itemWithLayoutSize:"), objc.ExtractPtr(layoutSize))
	return rv
}

func CollectionLayoutItem_ItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutItem {
	return CollectionLayoutItemClass.ItemWithLayoutSize(layoutSize)
}

func (cc _CollectionLayoutItemClass) ItemWithLayoutSizeSupplementaryItems(layoutSize ICollectionLayoutSize, supplementaryItems []ICollectionLayoutSupplementaryItem) CollectionLayoutItem {
	rv := objc.CallMethod[CollectionLayoutItem](cc, objc.GetSelector("itemWithLayoutSize:supplementaryItems:"), objc.ExtractPtr(layoutSize), supplementaryItems)
	return rv
}

func CollectionLayoutItem_ItemWithLayoutSizeSupplementaryItems(layoutSize ICollectionLayoutSize, supplementaryItems []ICollectionLayoutSupplementaryItem) CollectionLayoutItem {
	return CollectionLayoutItemClass.ItemWithLayoutSizeSupplementaryItems(layoutSize, supplementaryItems)
}

func (cc _CollectionLayoutItemClass) Alloc() CollectionLayoutItem {
	rv := objc.CallMethod[CollectionLayoutItem](cc, objc.GetSelector("alloc"))
	return rv
}

func CollectionLayoutItem_Alloc() CollectionLayoutItem {
	return CollectionLayoutItemClass.Alloc()
}

func (cc _CollectionLayoutItemClass) New() CollectionLayoutItem {
	rv := objc.CallMethod[CollectionLayoutItem](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutItem() CollectionLayoutItem {
	return CollectionLayoutItemClass.New()
}

func CollectionLayoutItem_New() CollectionLayoutItem {
	return CollectionLayoutItemClass.New()
}

func (c_ CollectionLayoutItem) Init() CollectionLayoutItem {
	rv := objc.CallMethod[CollectionLayoutItem](c_, objc.GetSelector("init"))
	return rv
}

func CollectionLayoutItem_Init() CollectionLayoutItem {
	return CollectionLayoutItemClass.Alloc().Init()
}

func (c_ CollectionLayoutItem) LayoutSize() CollectionLayoutSize {
	rv := objc.CallMethod[CollectionLayoutSize](c_, objc.GetSelector("layoutSize"))
	return rv
}

func (c_ CollectionLayoutItem) SupplementaryItems() []CollectionLayoutSupplementaryItem {
	rv := objc.CallMethod[[]CollectionLayoutSupplementaryItem](c_, objc.GetSelector("supplementaryItems"))
	return rv
}

func (c_ CollectionLayoutItem) EdgeSpacing() CollectionLayoutEdgeSpacing {
	rv := objc.CallMethod[CollectionLayoutEdgeSpacing](c_, objc.GetSelector("edgeSpacing"))
	return rv
}

func (c_ CollectionLayoutItem) SetEdgeSpacing(value ICollectionLayoutEdgeSpacing) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setEdgeSpacing:"), objc.ExtractPtr(value))
}

func (c_ CollectionLayoutItem) ContentInsets() DirectionalEdgeInsets {
	rv := objc.CallMethod[DirectionalEdgeInsets](c_, objc.GetSelector("contentInsets"))
	return rv
}

func (c_ CollectionLayoutItem) SetContentInsets(value DirectionalEdgeInsets) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setContentInsets:"), value)
}
