// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [WebHistory] class.
var WebHistoryClass = _WebHistoryClass{objc.GetClass("WebHistory")}

type _WebHistoryClass struct {
	objc.Class
}

// An interface definition for the [WebHistory] class.
type IWebHistory interface {
	objc.IObject
}

// WebHistory objects are used to maintain the pages visited by users. Visited pages are represented by WebHistoryItem objects. You add and remove history items using the addItems: and removeItems: methods. These methods post appropriate notifications when items are added or removed so you can update the display. WebHistory organizes the WebHistoryItem objects by the day they were visited, ordered from most recent to oldest. You can request all the days that contain history items using the orderedLastVisitedDays method or request the items visited on a particular day using the orderedItemsLastVisitedOnDay: method. WebHistory objects can be loaded and saved by specifying a file URL (see loadFromURL:error:). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webhistory?language=objc
type WebHistory struct {
	objc.Object
}

func WebHistoryFrom(ptr unsafe.Pointer) WebHistory {
	return WebHistory{
		Object: objc.ObjectFrom(ptr),
	}
}

func (wc _WebHistoryClass) Alloc() WebHistory {
	rv := objc.Call[WebHistory](wc, objc.Sel("alloc"))
	return rv
}

func WebHistory_Alloc() WebHistory {
	return WebHistoryClass.Alloc()
}

func (wc _WebHistoryClass) New() WebHistory {
	rv := objc.Call[WebHistory](wc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewWebHistory() WebHistory {
	return WebHistoryClass.New()
}

func (w_ WebHistory) Init() WebHistory {
	rv := objc.Call[WebHistory](w_, objc.Sel("init"))
	return rv
}
