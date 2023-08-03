// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var CollectionLayoutBoundarySupplementaryItemClass = _CollectionLayoutBoundarySupplementaryItemClass{objc.GetClass("NSCollectionLayoutBoundarySupplementaryItem")}

type _CollectionLayoutBoundarySupplementaryItemClass struct {
	objc.Class
}

type ICollectionLayoutBoundarySupplementaryItem interface {
	ICollectionLayoutSupplementaryItem
	PinToVisibleBounds() bool
	SetPinToVisibleBounds(value bool)
	Offset() foundation.Point
	Alignment() RectAlignment
	ExtendsBoundary() bool
	SetExtendsBoundary(value bool)
}

type CollectionLayoutBoundarySupplementaryItem struct {
	CollectionLayoutSupplementaryItem
}

func MakeCollectionLayoutBoundarySupplementaryItem(ptr unsafe.Pointer) CollectionLayoutBoundarySupplementaryItem {
	return CollectionLayoutBoundarySupplementaryItem{
		CollectionLayoutSupplementaryItem: MakeCollectionLayoutSupplementaryItem(ptr),
	}
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) BoundarySupplementaryItemWithLayoutSizeElementKindAlignment(layoutSize ICollectionLayoutSize, elementKind string, alignment RectAlignment) CollectionLayoutBoundarySupplementaryItem {
	rv := objc.CallMethod[CollectionLayoutBoundarySupplementaryItem](cc, objc.GetSelector("boundarySupplementaryItemWithLayoutSize:elementKind:alignment:"), objc.ExtractPtr(layoutSize), elementKind, alignment)
	return rv
}

func CollectionLayoutBoundarySupplementaryItem_BoundarySupplementaryItemWithLayoutSizeElementKindAlignment(layoutSize ICollectionLayoutSize, elementKind string, alignment RectAlignment) CollectionLayoutBoundarySupplementaryItem {
	return CollectionLayoutBoundarySupplementaryItemClass.BoundarySupplementaryItemWithLayoutSizeElementKindAlignment(layoutSize, elementKind, alignment)
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) BoundarySupplementaryItemWithLayoutSizeElementKindAlignmentAbsoluteOffset(layoutSize ICollectionLayoutSize, elementKind string, alignment RectAlignment, absoluteOffset foundation.Point) CollectionLayoutBoundarySupplementaryItem {
	rv := objc.CallMethod[CollectionLayoutBoundarySupplementaryItem](cc, objc.GetSelector("boundarySupplementaryItemWithLayoutSize:elementKind:alignment:absoluteOffset:"), objc.ExtractPtr(layoutSize), elementKind, alignment, absoluteOffset)
	return rv
}

func CollectionLayoutBoundarySupplementaryItem_BoundarySupplementaryItemWithLayoutSizeElementKindAlignmentAbsoluteOffset(layoutSize ICollectionLayoutSize, elementKind string, alignment RectAlignment, absoluteOffset foundation.Point) CollectionLayoutBoundarySupplementaryItem {
	return CollectionLayoutBoundarySupplementaryItemClass.BoundarySupplementaryItemWithLayoutSizeElementKindAlignmentAbsoluteOffset(layoutSize, elementKind, alignment, absoluteOffset)
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) SupplementaryItemWithLayoutSizeElementKindContainerAnchor(layoutSize ICollectionLayoutSize, elementKind string, containerAnchor ICollectionLayoutAnchor) CollectionLayoutBoundarySupplementaryItem {
	rv := objc.CallMethod[CollectionLayoutBoundarySupplementaryItem](cc, objc.GetSelector("supplementaryItemWithLayoutSize:elementKind:containerAnchor:"), objc.ExtractPtr(layoutSize), elementKind, objc.ExtractPtr(containerAnchor))
	return rv
}

func CollectionLayoutBoundarySupplementaryItem_SupplementaryItemWithLayoutSizeElementKindContainerAnchor(layoutSize ICollectionLayoutSize, elementKind string, containerAnchor ICollectionLayoutAnchor) CollectionLayoutBoundarySupplementaryItem {
	return CollectionLayoutBoundarySupplementaryItemClass.SupplementaryItemWithLayoutSizeElementKindContainerAnchor(layoutSize, elementKind, containerAnchor)
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) SupplementaryItemWithLayoutSizeElementKindContainerAnchorItemAnchor(layoutSize ICollectionLayoutSize, elementKind string, containerAnchor ICollectionLayoutAnchor, itemAnchor ICollectionLayoutAnchor) CollectionLayoutBoundarySupplementaryItem {
	rv := objc.CallMethod[CollectionLayoutBoundarySupplementaryItem](cc, objc.GetSelector("supplementaryItemWithLayoutSize:elementKind:containerAnchor:itemAnchor:"), objc.ExtractPtr(layoutSize), elementKind, objc.ExtractPtr(containerAnchor), objc.ExtractPtr(itemAnchor))
	return rv
}

func CollectionLayoutBoundarySupplementaryItem_SupplementaryItemWithLayoutSizeElementKindContainerAnchorItemAnchor(layoutSize ICollectionLayoutSize, elementKind string, containerAnchor ICollectionLayoutAnchor, itemAnchor ICollectionLayoutAnchor) CollectionLayoutBoundarySupplementaryItem {
	return CollectionLayoutBoundarySupplementaryItemClass.SupplementaryItemWithLayoutSizeElementKindContainerAnchorItemAnchor(layoutSize, elementKind, containerAnchor, itemAnchor)
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) ItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutBoundarySupplementaryItem {
	rv := objc.CallMethod[CollectionLayoutBoundarySupplementaryItem](cc, objc.GetSelector("itemWithLayoutSize:"), objc.ExtractPtr(layoutSize))
	return rv
}

func CollectionLayoutBoundarySupplementaryItem_ItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutBoundarySupplementaryItem {
	return CollectionLayoutBoundarySupplementaryItemClass.ItemWithLayoutSize(layoutSize)
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) ItemWithLayoutSizeSupplementaryItems(layoutSize ICollectionLayoutSize, supplementaryItems []ICollectionLayoutSupplementaryItem) CollectionLayoutBoundarySupplementaryItem {
	rv := objc.CallMethod[CollectionLayoutBoundarySupplementaryItem](cc, objc.GetSelector("itemWithLayoutSize:supplementaryItems:"), objc.ExtractPtr(layoutSize), supplementaryItems)
	return rv
}

func CollectionLayoutBoundarySupplementaryItem_ItemWithLayoutSizeSupplementaryItems(layoutSize ICollectionLayoutSize, supplementaryItems []ICollectionLayoutSupplementaryItem) CollectionLayoutBoundarySupplementaryItem {
	return CollectionLayoutBoundarySupplementaryItemClass.ItemWithLayoutSizeSupplementaryItems(layoutSize, supplementaryItems)
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) Alloc() CollectionLayoutBoundarySupplementaryItem {
	rv := objc.CallMethod[CollectionLayoutBoundarySupplementaryItem](cc, objc.GetSelector("alloc"))
	return rv
}

func CollectionLayoutBoundarySupplementaryItem_Alloc() CollectionLayoutBoundarySupplementaryItem {
	return CollectionLayoutBoundarySupplementaryItemClass.Alloc()
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) New() CollectionLayoutBoundarySupplementaryItem {
	rv := objc.CallMethod[CollectionLayoutBoundarySupplementaryItem](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutBoundarySupplementaryItem() CollectionLayoutBoundarySupplementaryItem {
	return CollectionLayoutBoundarySupplementaryItemClass.New()
}

func CollectionLayoutBoundarySupplementaryItem_New() CollectionLayoutBoundarySupplementaryItem {
	return CollectionLayoutBoundarySupplementaryItemClass.New()
}

func (c_ CollectionLayoutBoundarySupplementaryItem) Init() CollectionLayoutBoundarySupplementaryItem {
	rv := objc.CallMethod[CollectionLayoutBoundarySupplementaryItem](c_, objc.GetSelector("init"))
	return rv
}

func CollectionLayoutBoundarySupplementaryItem_Init() CollectionLayoutBoundarySupplementaryItem {
	return CollectionLayoutBoundarySupplementaryItemClass.Alloc().Init()
}

func (c_ CollectionLayoutBoundarySupplementaryItem) PinToVisibleBounds() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("pinToVisibleBounds"))
	return rv
}

func (c_ CollectionLayoutBoundarySupplementaryItem) SetPinToVisibleBounds(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setPinToVisibleBounds:"), value)
}

func (c_ CollectionLayoutBoundarySupplementaryItem) Offset() foundation.Point {
	rv := objc.CallMethod[foundation.Point](c_, objc.GetSelector("offset"))
	return rv
}

func (c_ CollectionLayoutBoundarySupplementaryItem) Alignment() RectAlignment {
	rv := objc.CallMethod[RectAlignment](c_, objc.GetSelector("alignment"))
	return rv
}

func (c_ CollectionLayoutBoundarySupplementaryItem) ExtendsBoundary() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("extendsBoundary"))
	return rv
}

func (c_ CollectionLayoutBoundarySupplementaryItem) SetExtendsBoundary(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setExtendsBoundary:"), value)
}
