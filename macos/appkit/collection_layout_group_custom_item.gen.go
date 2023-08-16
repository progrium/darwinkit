// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CollectionLayoutGroupCustomItem] class.
var CollectionLayoutGroupCustomItemClass = _CollectionLayoutGroupCustomItemClass{objc.GetClass("NSCollectionLayoutGroupCustomItem")}

type _CollectionLayoutGroupCustomItemClass struct {
	objc.Class
}

// An interface definition for the [CollectionLayoutGroupCustomItem] class.
type ICollectionLayoutGroupCustomItem interface {
	objc.IObject
	Frame() coregraphics.Rect
	ZIndex() int
}

// An item used in a group with a custom layout arrangement. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutgroupcustomitem?language=objc
type CollectionLayoutGroupCustomItem struct {
	objc.Object
}

func CollectionLayoutGroupCustomItemFrom(ptr unsafe.Pointer) CollectionLayoutGroupCustomItem {
	return CollectionLayoutGroupCustomItem{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CollectionLayoutGroupCustomItemClass) CustomItemWithFrame(frame coregraphics.Rect) CollectionLayoutGroupCustomItem {
	rv := objc.Call[CollectionLayoutGroupCustomItem](cc, objc.Sel("customItemWithFrame:"), frame)
	return rv
}

// Creates a custom item with the specified frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutgroupcustomitem/3213863-customitemwithframe?language=objc
func CollectionLayoutGroupCustomItem_CustomItemWithFrame(frame coregraphics.Rect) CollectionLayoutGroupCustomItem {
	return CollectionLayoutGroupCustomItemClass.CustomItemWithFrame(frame)
}

func (cc _CollectionLayoutGroupCustomItemClass) Alloc() CollectionLayoutGroupCustomItem {
	rv := objc.Call[CollectionLayoutGroupCustomItem](cc, objc.Sel("alloc"))
	return rv
}

func CollectionLayoutGroupCustomItem_Alloc() CollectionLayoutGroupCustomItem {
	return CollectionLayoutGroupCustomItemClass.Alloc()
}

func (cc _CollectionLayoutGroupCustomItemClass) New() CollectionLayoutGroupCustomItem {
	rv := objc.Call[CollectionLayoutGroupCustomItem](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutGroupCustomItem() CollectionLayoutGroupCustomItem {
	return CollectionLayoutGroupCustomItemClass.New()
}

func (c_ CollectionLayoutGroupCustomItem) Init() CollectionLayoutGroupCustomItem {
	rv := objc.Call[CollectionLayoutGroupCustomItem](c_, objc.Sel("init"))
	return rv
}

// The frame of the custom item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutgroupcustomitem/3213865-frame?language=objc
func (c_ CollectionLayoutGroupCustomItem) Frame() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](c_, objc.Sel("frame"))
	return rv
}

// The vertical stacking order of the custom item in relation to other items in the group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutgroupcustomitem/3213866-zindex?language=objc
func (c_ CollectionLayoutGroupCustomItem) ZIndex() int {
	rv := objc.Call[int](c_, objc.Sel("zIndex"))
	return rv
}
