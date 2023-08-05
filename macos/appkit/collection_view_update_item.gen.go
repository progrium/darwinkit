// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var CollectionViewUpdateItemClass = _CollectionViewUpdateItemClass{objc.GetClass("NSCollectionViewUpdateItem")}

type _CollectionViewUpdateItemClass struct {
	objc.Class
}

type ICollectionViewUpdateItem interface {
	objc.IObject
	IndexPathBeforeUpdate() foundation.IndexPath
	IndexPathAfterUpdate() foundation.IndexPath
	UpdateAction() CollectionUpdateAction
}

type CollectionViewUpdateItem struct {
	objc.Object
}

func MakeCollectionViewUpdateItem(ptr unsafe.Pointer) CollectionViewUpdateItem {
	return CollectionViewUpdateItem{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionViewUpdateItemClass) Alloc() CollectionViewUpdateItem {
	rv := objc.CallMethod[CollectionViewUpdateItem](cc, objc.GetSelector("alloc"))
	return rv
}

func CollectionViewUpdateItem_Alloc() CollectionViewUpdateItem {
	return CollectionViewUpdateItemClass.Alloc()
}

func (cc _CollectionViewUpdateItemClass) New() CollectionViewUpdateItem {
	rv := objc.CallMethod[CollectionViewUpdateItem](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionViewUpdateItem() CollectionViewUpdateItem {
	return CollectionViewUpdateItemClass.New()
}

func CollectionViewUpdateItem_New() CollectionViewUpdateItem {
	return CollectionViewUpdateItemClass.New()
}

func (c_ CollectionViewUpdateItem) Init() CollectionViewUpdateItem {
	rv := objc.CallMethod[CollectionViewUpdateItem](c_, objc.GetSelector("init"))
	return rv
}

func CollectionViewUpdateItem_Init() CollectionViewUpdateItem {
	return CollectionViewUpdateItemClass.Alloc().Init()
}

func (c_ CollectionViewUpdateItem) IndexPathBeforeUpdate() foundation.IndexPath {
	rv := objc.CallMethod[foundation.IndexPath](c_, objc.GetSelector("indexPathBeforeUpdate"))
	return rv
}

func (c_ CollectionViewUpdateItem) IndexPathAfterUpdate() foundation.IndexPath {
	rv := objc.CallMethod[foundation.IndexPath](c_, objc.GetSelector("indexPathAfterUpdate"))
	return rv
}

func (c_ CollectionViewUpdateItem) UpdateAction() CollectionUpdateAction {
	rv := objc.CallMethod[CollectionUpdateAction](c_, objc.GetSelector("updateAction"))
	return rv
}
