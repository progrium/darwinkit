// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [BackForwardList] class.
var BackForwardListClass = _BackForwardListClass{objc.GetClass("WKBackForwardList")}

type _BackForwardListClass struct {
	objc.Class
}

// An interface definition for the [BackForwardList] class.
type IBackForwardList interface {
	objc.IObject
	ItemAtIndex(index int) BackForwardListItem
	BackList() []BackForwardListItem
	ForwardItem() BackForwardListItem
	CurrentItem() BackForwardListItem
	BackItem() BackForwardListItem
	ForwardList() []BackForwardListItem
}

// An object that manages the list of previously loaded webpages, which the web view uses for forward and backward navigation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkbackforwardlist?language=objc
type BackForwardList struct {
	objc.Object
}

func BackForwardListFrom(ptr unsafe.Pointer) BackForwardList {
	return BackForwardList{
		Object: objc.ObjectFrom(ptr),
	}
}

func (bc _BackForwardListClass) Alloc() BackForwardList {
	rv := objc.Call[BackForwardList](bc, objc.Sel("alloc"))
	return rv
}

func BackForwardList_Alloc() BackForwardList {
	return BackForwardListClass.Alloc()
}

func (bc _BackForwardListClass) New() BackForwardList {
	rv := objc.Call[BackForwardList](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewBackForwardList() BackForwardList {
	return BackForwardListClass.New()
}

func (b_ BackForwardList) Init() BackForwardList {
	rv := objc.Call[BackForwardList](b_, objc.Sel("init"))
	return rv
}

// Returns the item at the relative offset from the current item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkbackforwardlist/1516694-itematindex?language=objc
func (b_ BackForwardList) ItemAtIndex(index int) BackForwardListItem {
	rv := objc.Call[BackForwardListItem](b_, objc.Sel("itemAtIndex:"), index)
	return rv
}

// The array of items that precede the current item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkbackforwardlist/1516698-backlist?language=objc
func (b_ BackForwardList) BackList() []BackForwardListItem {
	rv := objc.Call[[]BackForwardListItem](b_, objc.Sel("backList"))
	return rv
}

// The item immediately following the current item, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkbackforwardlist/1516700-forwarditem?language=objc
func (b_ BackForwardList) ForwardItem() BackForwardListItem {
	rv := objc.Call[BackForwardListItem](b_, objc.Sel("forwardItem"))
	return rv
}

// The current item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkbackforwardlist/1516703-currentitem?language=objc
func (b_ BackForwardList) CurrentItem() BackForwardListItem {
	rv := objc.Call[BackForwardListItem](b_, objc.Sel("currentItem"))
	return rv
}

// The item immediately preceding the current item, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkbackforwardlist/1516693-backitem?language=objc
func (b_ BackForwardList) BackItem() BackForwardListItem {
	rv := objc.Call[BackForwardListItem](b_, objc.Sel("backItem"))
	return rv
}

// The array of items that follow the current item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkbackforwardlist/1516701-forwardlist?language=objc
func (b_ BackForwardList) ForwardList() []BackForwardListItem {
	rv := objc.Call[[]BackForwardListItem](b_, objc.Sel("forwardList"))
	return rv
}
