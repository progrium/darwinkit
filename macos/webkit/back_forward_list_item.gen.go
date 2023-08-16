// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [BackForwardListItem] class.
var BackForwardListItemClass = _BackForwardListItemClass{objc.GetClass("WKBackForwardListItem")}

type _BackForwardListItemClass struct {
	objc.Class
}

// An interface definition for the [BackForwardListItem] class.
type IBackForwardListItem interface {
	objc.IObject
	URL() foundation.URL
	Title() string
	InitialURL() foundation.URL
}

// A representation of a webpage that the web view previously visited. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkbackforwardlistitem?language=objc
type BackForwardListItem struct {
	objc.Object
}

func BackForwardListItemFrom(ptr unsafe.Pointer) BackForwardListItem {
	return BackForwardListItem{
		Object: objc.ObjectFrom(ptr),
	}
}

func (bc _BackForwardListItemClass) Alloc() BackForwardListItem {
	rv := objc.Call[BackForwardListItem](bc, objc.Sel("alloc"))
	return rv
}

func BackForwardListItem_Alloc() BackForwardListItem {
	return BackForwardListItemClass.Alloc()
}

func (bc _BackForwardListItemClass) New() BackForwardListItem {
	rv := objc.Call[BackForwardListItem](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewBackForwardListItem() BackForwardListItem {
	return BackForwardListItemClass.New()
}

func (b_ BackForwardListItem) Init() BackForwardListItem {
	rv := objc.Call[BackForwardListItem](b_, objc.Sel("init"))
	return rv
}

// The URL of the webpage this item represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkbackforwardlistitem/1455513-url?language=objc
func (b_ BackForwardListItem) URL() foundation.URL {
	rv := objc.Call[foundation.URL](b_, objc.Sel("URL"))
	return rv
}

// The title of the webpage this item represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkbackforwardlistitem/1455511-title?language=objc
func (b_ BackForwardListItem) Title() string {
	rv := objc.Call[string](b_, objc.Sel("title"))
	return rv
}

// The source URL that originally asked the web view to load this page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkbackforwardlistitem/1455507-initialurl?language=objc
func (b_ BackForwardListItem) InitialURL() foundation.URL {
	rv := objc.Call[foundation.URL](b_, objc.Sel("initialURL"))
	return rv
}
