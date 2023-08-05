// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var CollectionLayoutGroupCustomItemClass = _CollectionLayoutGroupCustomItemClass{objc.GetClass("NSCollectionLayoutGroupCustomItem")}

type _CollectionLayoutGroupCustomItemClass struct {
	objc.Class
}

type ICollectionLayoutGroupCustomItem interface {
	objc.IObject
	Frame() foundation.Rect
	ZIndex() int
}

type CollectionLayoutGroupCustomItem struct {
	objc.Object
}

func MakeCollectionLayoutGroupCustomItem(ptr unsafe.Pointer) CollectionLayoutGroupCustomItem {
	return CollectionLayoutGroupCustomItem{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionLayoutGroupCustomItemClass) CustomItemWithFrame(frame foundation.Rect) CollectionLayoutGroupCustomItem {
	rv := objc.CallMethod[CollectionLayoutGroupCustomItem](cc, objc.GetSelector("customItemWithFrame:"), frame)
	return rv
}

func CollectionLayoutGroupCustomItem_CustomItemWithFrame(frame foundation.Rect) CollectionLayoutGroupCustomItem {
	return CollectionLayoutGroupCustomItemClass.CustomItemWithFrame(frame)
}

func (cc _CollectionLayoutGroupCustomItemClass) CustomItemWithFrameZIndex(frame foundation.Rect, zIndex int) CollectionLayoutGroupCustomItem {
	rv := objc.CallMethod[CollectionLayoutGroupCustomItem](cc, objc.GetSelector("customItemWithFrame:zIndex:"), frame, zIndex)
	return rv
}

func CollectionLayoutGroupCustomItem_CustomItemWithFrameZIndex(frame foundation.Rect, zIndex int) CollectionLayoutGroupCustomItem {
	return CollectionLayoutGroupCustomItemClass.CustomItemWithFrameZIndex(frame, zIndex)
}

func (cc _CollectionLayoutGroupCustomItemClass) Alloc() CollectionLayoutGroupCustomItem {
	rv := objc.CallMethod[CollectionLayoutGroupCustomItem](cc, objc.GetSelector("alloc"))
	return rv
}

func CollectionLayoutGroupCustomItem_Alloc() CollectionLayoutGroupCustomItem {
	return CollectionLayoutGroupCustomItemClass.Alloc()
}

func (cc _CollectionLayoutGroupCustomItemClass) New() CollectionLayoutGroupCustomItem {
	rv := objc.CallMethod[CollectionLayoutGroupCustomItem](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutGroupCustomItem() CollectionLayoutGroupCustomItem {
	return CollectionLayoutGroupCustomItemClass.New()
}

func CollectionLayoutGroupCustomItem_New() CollectionLayoutGroupCustomItem {
	return CollectionLayoutGroupCustomItemClass.New()
}

func (c_ CollectionLayoutGroupCustomItem) Init() CollectionLayoutGroupCustomItem {
	rv := objc.CallMethod[CollectionLayoutGroupCustomItem](c_, objc.GetSelector("init"))
	return rv
}

func CollectionLayoutGroupCustomItem_Init() CollectionLayoutGroupCustomItem {
	return CollectionLayoutGroupCustomItemClass.Alloc().Init()
}

func (c_ CollectionLayoutGroupCustomItem) Frame() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](c_, objc.GetSelector("frame"))
	return rv
}

func (c_ CollectionLayoutGroupCustomItem) ZIndex() int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("zIndex"))
	return rv
}
