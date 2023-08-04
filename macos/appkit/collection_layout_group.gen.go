// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var CollectionLayoutGroupClass = _CollectionLayoutGroupClass{objc.GetClass("NSCollectionLayoutGroup")}

type _CollectionLayoutGroupClass struct {
	objc.Class
}

type ICollectionLayoutGroup interface {
	ICollectionLayoutItem
	VisualDescription() string
	Subitems() []CollectionLayoutItem
	SetSupplementaryItems(value []ICollectionLayoutSupplementaryItem)
	InterItemSpacing() CollectionLayoutSpacing
	SetInterItemSpacing(value ICollectionLayoutSpacing)
}

type CollectionLayoutGroup struct {
	CollectionLayoutItem
}

func MakeCollectionLayoutGroup(ptr unsafe.Pointer) CollectionLayoutGroup {
	return CollectionLayoutGroup{
		CollectionLayoutItem: MakeCollectionLayoutItem(ptr),
	}
}

func (cc _CollectionLayoutGroupClass) HorizontalGroupWithLayoutSizeSubitems(layoutSize ICollectionLayoutSize, subitems []ICollectionLayoutItem) CollectionLayoutGroup {
	rv := objc.CallMethod[CollectionLayoutGroup](cc, objc.GetSelector("horizontalGroupWithLayoutSize:subitems:"), objc.ExtractPtr(layoutSize), subitems)
	return rv
}

func CollectionLayoutGroup_HorizontalGroupWithLayoutSizeSubitems(layoutSize ICollectionLayoutSize, subitems []ICollectionLayoutItem) CollectionLayoutGroup {
	return CollectionLayoutGroupClass.HorizontalGroupWithLayoutSizeSubitems(layoutSize, subitems)
}

func (cc _CollectionLayoutGroupClass) VerticalGroupWithLayoutSizeSubitems(layoutSize ICollectionLayoutSize, subitems []ICollectionLayoutItem) CollectionLayoutGroup {
	rv := objc.CallMethod[CollectionLayoutGroup](cc, objc.GetSelector("verticalGroupWithLayoutSize:subitems:"), objc.ExtractPtr(layoutSize), subitems)
	return rv
}

func CollectionLayoutGroup_VerticalGroupWithLayoutSizeSubitems(layoutSize ICollectionLayoutSize, subitems []ICollectionLayoutItem) CollectionLayoutGroup {
	return CollectionLayoutGroupClass.VerticalGroupWithLayoutSizeSubitems(layoutSize, subitems)
}

func (cc _CollectionLayoutGroupClass) CustomGroupWithLayoutSizeItemProvider(layoutSize ICollectionLayoutSize, itemProvider func(layoutEnvironment CollectionLayoutEnvironmentWrapper) []CollectionLayoutGroupCustomItem) CollectionLayoutGroup {
	rv := objc.CallMethod[CollectionLayoutGroup](cc, objc.GetSelector("customGroupWithLayoutSize:itemProvider:"), objc.ExtractPtr(layoutSize), itemProvider)
	return rv
}

func CollectionLayoutGroup_CustomGroupWithLayoutSizeItemProvider(layoutSize ICollectionLayoutSize, itemProvider func(layoutEnvironment CollectionLayoutEnvironmentWrapper) []CollectionLayoutGroupCustomItem) CollectionLayoutGroup {
	return CollectionLayoutGroupClass.CustomGroupWithLayoutSizeItemProvider(layoutSize, itemProvider)
}

func (cc _CollectionLayoutGroupClass) HorizontalGroupWithLayoutSizeSubitemCount(layoutSize ICollectionLayoutSize, subitem ICollectionLayoutItem, count int) CollectionLayoutGroup {
	rv := objc.CallMethod[CollectionLayoutGroup](cc, objc.GetSelector("horizontalGroupWithLayoutSize:subitem:count:"), objc.ExtractPtr(layoutSize), objc.ExtractPtr(subitem), count)
	return rv
}

func CollectionLayoutGroup_HorizontalGroupWithLayoutSizeSubitemCount(layoutSize ICollectionLayoutSize, subitem ICollectionLayoutItem, count int) CollectionLayoutGroup {
	return CollectionLayoutGroupClass.HorizontalGroupWithLayoutSizeSubitemCount(layoutSize, subitem, count)
}

func (cc _CollectionLayoutGroupClass) VerticalGroupWithLayoutSizeSubitemCount(layoutSize ICollectionLayoutSize, subitem ICollectionLayoutItem, count int) CollectionLayoutGroup {
	rv := objc.CallMethod[CollectionLayoutGroup](cc, objc.GetSelector("verticalGroupWithLayoutSize:subitem:count:"), objc.ExtractPtr(layoutSize), objc.ExtractPtr(subitem), count)
	return rv
}

func CollectionLayoutGroup_VerticalGroupWithLayoutSizeSubitemCount(layoutSize ICollectionLayoutSize, subitem ICollectionLayoutItem, count int) CollectionLayoutGroup {
	return CollectionLayoutGroupClass.VerticalGroupWithLayoutSizeSubitemCount(layoutSize, subitem, count)
}

func (cc _CollectionLayoutGroupClass) ItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutGroup {
	rv := objc.CallMethod[CollectionLayoutGroup](cc, objc.GetSelector("itemWithLayoutSize:"), objc.ExtractPtr(layoutSize))
	return rv
}

func CollectionLayoutGroup_ItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutGroup {
	return CollectionLayoutGroupClass.ItemWithLayoutSize(layoutSize)
}

func (cc _CollectionLayoutGroupClass) ItemWithLayoutSizeSupplementaryItems(layoutSize ICollectionLayoutSize, supplementaryItems []ICollectionLayoutSupplementaryItem) CollectionLayoutGroup {
	rv := objc.CallMethod[CollectionLayoutGroup](cc, objc.GetSelector("itemWithLayoutSize:supplementaryItems:"), objc.ExtractPtr(layoutSize), supplementaryItems)
	return rv
}

func CollectionLayoutGroup_ItemWithLayoutSizeSupplementaryItems(layoutSize ICollectionLayoutSize, supplementaryItems []ICollectionLayoutSupplementaryItem) CollectionLayoutGroup {
	return CollectionLayoutGroupClass.ItemWithLayoutSizeSupplementaryItems(layoutSize, supplementaryItems)
}

func (cc _CollectionLayoutGroupClass) Alloc() CollectionLayoutGroup {
	rv := objc.CallMethod[CollectionLayoutGroup](cc, objc.GetSelector("alloc"))
	return rv
}

func CollectionLayoutGroup_Alloc() CollectionLayoutGroup {
	return CollectionLayoutGroupClass.Alloc()
}

func (cc _CollectionLayoutGroupClass) New() CollectionLayoutGroup {
	rv := objc.CallMethod[CollectionLayoutGroup](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutGroup() CollectionLayoutGroup {
	return CollectionLayoutGroupClass.New()
}

func CollectionLayoutGroup_New() CollectionLayoutGroup {
	return CollectionLayoutGroupClass.New()
}

func (c_ CollectionLayoutGroup) Init() CollectionLayoutGroup {
	rv := objc.CallMethod[CollectionLayoutGroup](c_, objc.GetSelector("init"))
	return rv
}

func CollectionLayoutGroup_Init() CollectionLayoutGroup {
	return CollectionLayoutGroupClass.Alloc().Init()
}

func (c_ CollectionLayoutGroup) VisualDescription() string {
	rv := objc.CallMethod[string](c_, objc.GetSelector("visualDescription"))
	return rv
}

func (c_ CollectionLayoutGroup) Subitems() []CollectionLayoutItem {
	rv := objc.CallMethod[[]CollectionLayoutItem](c_, objc.GetSelector("subitems"))
	return rv
}

func (c_ CollectionLayoutGroup) SetSupplementaryItems(value []ICollectionLayoutSupplementaryItem) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setSupplementaryItems:"), value)
}

func (c_ CollectionLayoutGroup) InterItemSpacing() CollectionLayoutSpacing {
	rv := objc.CallMethod[CollectionLayoutSpacing](c_, objc.GetSelector("interItemSpacing"))
	return rv
}

func (c_ CollectionLayoutGroup) SetInterItemSpacing(value ICollectionLayoutSpacing) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setInterItemSpacing:"), objc.ExtractPtr(value))
}
