// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var BackForwardListClass = _BackForwardListClass{objc.GetClass("WKBackForwardList")}

type _BackForwardListClass struct {
	objc.Class
}

type IBackForwardList interface {
	objc.IObject
	ItemAtIndex(index int) BackForwardListItem
	BackItem() BackForwardListItem
	CurrentItem() BackForwardListItem
	ForwardItem() BackForwardListItem
	BackList() []BackForwardListItem
	ForwardList() []BackForwardListItem
}

type BackForwardList struct {
	objc.Object
}

func MakeBackForwardList(ptr unsafe.Pointer) BackForwardList {
	return BackForwardList{
		Object: objc.MakeObject(ptr),
	}
}

func (bc _BackForwardListClass) Alloc() BackForwardList {
	rv := objc.CallMethod[BackForwardList](bc, objc.GetSelector("alloc"))
	return rv
}

func BackForwardList_Alloc() BackForwardList {
	return BackForwardListClass.Alloc()
}

func (bc _BackForwardListClass) New() BackForwardList {
	rv := objc.CallMethod[BackForwardList](bc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewBackForwardList() BackForwardList {
	return BackForwardListClass.New()
}

func BackForwardList_New() BackForwardList {
	return BackForwardListClass.New()
}

func (b_ BackForwardList) Init() BackForwardList {
	rv := objc.CallMethod[BackForwardList](b_, objc.GetSelector("init"))
	return rv
}

func BackForwardList_Init() BackForwardList {
	return BackForwardListClass.Alloc().Init()
}

func (b_ BackForwardList) ItemAtIndex(index int) BackForwardListItem {
	rv := objc.CallMethod[BackForwardListItem](b_, objc.GetSelector("itemAtIndex:"), index)
	return rv
}

func (b_ BackForwardList) BackItem() BackForwardListItem {
	rv := objc.CallMethod[BackForwardListItem](b_, objc.GetSelector("backItem"))
	return rv
}

func (b_ BackForwardList) CurrentItem() BackForwardListItem {
	rv := objc.CallMethod[BackForwardListItem](b_, objc.GetSelector("currentItem"))
	return rv
}

func (b_ BackForwardList) ForwardItem() BackForwardListItem {
	rv := objc.CallMethod[BackForwardListItem](b_, objc.GetSelector("forwardItem"))
	return rv
}

func (b_ BackForwardList) BackList() []BackForwardListItem {
	rv := objc.CallMethod[[]BackForwardListItem](b_, objc.GetSelector("backList"))
	return rv
}

func (b_ BackForwardList) ForwardList() []BackForwardListItem {
	rv := objc.CallMethod[[]BackForwardListItem](b_, objc.GetSelector("forwardList"))
	return rv
}
