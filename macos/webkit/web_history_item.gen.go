// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [WebHistoryItem] class.
var WebHistoryItemClass = _WebHistoryItemClass{objc.GetClass("WebHistoryItem")}

type _WebHistoryItemClass struct {
	objc.Class
}

// An interface definition for the [WebHistoryItem] class.
type IWebHistoryItem interface {
	objc.IObject
}

// WebHistoryItem objects encapsulate information about visiting a page so that users can return to that page. WebHistory and WebBackForwardList objects manage lists of WebHistoryItem objects. WebHistoryItem objects are created and added to these lists automatically when loading pages, so you do not need to create WebHistoryItem objects directly. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webhistoryitem?language=objc
type WebHistoryItem struct {
	objc.Object
}

func WebHistoryItemFrom(ptr unsafe.Pointer) WebHistoryItem {
	return WebHistoryItem{
		Object: objc.ObjectFrom(ptr),
	}
}

func (wc _WebHistoryItemClass) Alloc() WebHistoryItem {
	rv := objc.Call[WebHistoryItem](wc, objc.Sel("alloc"))
	return rv
}

func WebHistoryItem_Alloc() WebHistoryItem {
	return WebHistoryItemClass.Alloc()
}

func (wc _WebHistoryItemClass) New() WebHistoryItem {
	rv := objc.Call[WebHistoryItem](wc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewWebHistoryItem() WebHistoryItem {
	return WebHistoryItemClass.New()
}

func (w_ WebHistoryItem) Init() WebHistoryItem {
	rv := objc.Call[WebHistoryItem](w_, objc.Sel("init"))
	return rv
}
