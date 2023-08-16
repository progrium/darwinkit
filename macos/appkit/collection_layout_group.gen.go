// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CollectionLayoutGroup] class.
var CollectionLayoutGroupClass = _CollectionLayoutGroupClass{objc.GetClass("NSCollectionLayoutGroup")}

type _CollectionLayoutGroupClass struct {
	objc.Class
}

// An interface definition for the [CollectionLayoutGroup] class.
type ICollectionLayoutGroup interface {
	ICollectionLayoutItem
	VisualDescription() string
	SetSupplementaryItems(value []ICollectionLayoutSupplementaryItem)
	Subitems() []CollectionLayoutItem
	InterItemSpacing() CollectionLayoutSpacing
	SetInterItemSpacing(value ICollectionLayoutSpacing)
}

// A container for a set of items that lays out the items along a path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutgroup?language=objc
type CollectionLayoutGroup struct {
	CollectionLayoutItem
}

func CollectionLayoutGroupFrom(ptr unsafe.Pointer) CollectionLayoutGroup {
	return CollectionLayoutGroup{
		CollectionLayoutItem: CollectionLayoutItemFrom(ptr),
	}
}

func (cc _CollectionLayoutGroupClass) VerticalGroupWithLayoutSizeSubitems(layoutSize ICollectionLayoutSize, subitems []ICollectionLayoutItem) CollectionLayoutGroup {
	rv := objc.Call[CollectionLayoutGroup](cc, objc.Sel("verticalGroupWithLayoutSize:subitems:"), objc.Ptr(layoutSize), subitems)
	return rv
}

// Creates a group of the specified size, containing an array of items arranged in a vertical line. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutgroup/3213860-verticalgroupwithlayoutsize?language=objc
func CollectionLayoutGroup_VerticalGroupWithLayoutSizeSubitems(layoutSize ICollectionLayoutSize, subitems []ICollectionLayoutItem) CollectionLayoutGroup {
	return CollectionLayoutGroupClass.VerticalGroupWithLayoutSizeSubitems(layoutSize, subitems)
}

func (cc _CollectionLayoutGroupClass) HorizontalGroupWithLayoutSizeSubitems(layoutSize ICollectionLayoutSize, subitems []ICollectionLayoutItem) CollectionLayoutGroup {
	rv := objc.Call[CollectionLayoutGroup](cc, objc.Sel("horizontalGroupWithLayoutSize:subitems:"), objc.Ptr(layoutSize), subitems)
	return rv
}

// Creates a group of the specified size, containing an array of items arranged in a horizontal line. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutgroup/3213855-horizontalgroupwithlayoutsize?language=objc
func CollectionLayoutGroup_HorizontalGroupWithLayoutSizeSubitems(layoutSize ICollectionLayoutSize, subitems []ICollectionLayoutItem) CollectionLayoutGroup {
	return CollectionLayoutGroupClass.HorizontalGroupWithLayoutSizeSubitems(layoutSize, subitems)
}

func (cc _CollectionLayoutGroupClass) CustomGroupWithLayoutSizeItemProvider(layoutSize ICollectionLayoutSize, itemProvider CollectionLayoutGroupCustomItemProvider) CollectionLayoutGroup {
	rv := objc.Call[CollectionLayoutGroup](cc, objc.Sel("customGroupWithLayoutSize:itemProvider:"), objc.Ptr(layoutSize), itemProvider)
	return rv
}

// Creates a group of the specified size, with an item provider that creates a custom arrangement for those items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutgroup/3213853-customgroupwithlayoutsize?language=objc
func CollectionLayoutGroup_CustomGroupWithLayoutSizeItemProvider(layoutSize ICollectionLayoutSize, itemProvider CollectionLayoutGroupCustomItemProvider) CollectionLayoutGroup {
	return CollectionLayoutGroupClass.CustomGroupWithLayoutSizeItemProvider(layoutSize, itemProvider)
}

func (cc _CollectionLayoutGroupClass) Alloc() CollectionLayoutGroup {
	rv := objc.Call[CollectionLayoutGroup](cc, objc.Sel("alloc"))
	return rv
}

func CollectionLayoutGroup_Alloc() CollectionLayoutGroup {
	return CollectionLayoutGroupClass.Alloc()
}

func (cc _CollectionLayoutGroupClass) New() CollectionLayoutGroup {
	rv := objc.Call[CollectionLayoutGroup](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutGroup() CollectionLayoutGroup {
	return CollectionLayoutGroupClass.New()
}

func (c_ CollectionLayoutGroup) Init() CollectionLayoutGroup {
	rv := objc.Call[CollectionLayoutGroup](c_, objc.Sel("init"))
	return rv
}

func (cc _CollectionLayoutGroupClass) ItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutGroup {
	rv := objc.Call[CollectionLayoutGroup](cc, objc.Sel("itemWithLayoutSize:"), objc.Ptr(layoutSize))
	return rv
}

// Creates an item of the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutitem/3213871-itemwithlayoutsize?language=objc
func CollectionLayoutGroup_ItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutGroup {
	return CollectionLayoutGroupClass.ItemWithLayoutSize(layoutSize)
}

// Returns a string with an ASCII representation of the group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutgroup/3199082-visualdescription?language=objc
func (c_ CollectionLayoutGroup) VisualDescription() string {
	rv := objc.Call[string](c_, objc.Sel("visualDescription"))
	return rv
}

// An array of the supplementary items that are anchored to the group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutgroup/3199079-supplementaryitems?language=objc
func (c_ CollectionLayoutGroup) SetSupplementaryItems(value []ICollectionLayoutSupplementaryItem) {
	objc.Call[objc.Void](c_, objc.Sel("setSupplementaryItems:"), value)
}

// An array of the items contained in the group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutgroup/3213857-subitems?language=objc
func (c_ CollectionLayoutGroup) Subitems() []CollectionLayoutItem {
	rv := objc.Call[[]CollectionLayoutItem](c_, objc.Sel("subitems"))
	return rv
}

// The amount of space between the items in the group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutgroup/3199078-interitemspacing?language=objc
func (c_ CollectionLayoutGroup) InterItemSpacing() CollectionLayoutSpacing {
	rv := objc.Call[CollectionLayoutSpacing](c_, objc.Sel("interItemSpacing"))
	return rv
}

// The amount of space between the items in the group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutgroup/3199078-interitemspacing?language=objc
func (c_ CollectionLayoutGroup) SetInterItemSpacing(value ICollectionLayoutSpacing) {
	objc.Call[objc.Void](c_, objc.Sel("setInterItemSpacing:"), objc.Ptr(value))
}
