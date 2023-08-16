// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [WebBackForwardList] class.
var WebBackForwardListClass = _WebBackForwardListClass{objc.GetClass("WebBackForwardList")}

type _WebBackForwardListClass struct {
	objc.Class
}

// An interface definition for the [WebBackForwardList] class.
type IWebBackForwardList interface {
	objc.IObject
}

// A WebBackForwardList object maintains a list of visited pages used to go back and forward to the most recent page. A WebBackForwardList object maintains only the list data—it does not perform actual page loads (in other words, it does not make any client requests). If you need to perform a page load, see the loadRequest: method in WebFrame to find out how to do this. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webbackforwardlist?language=objc
type WebBackForwardList struct {
	objc.Object
}

func WebBackForwardListFrom(ptr unsafe.Pointer) WebBackForwardList {
	return WebBackForwardList{
		Object: objc.ObjectFrom(ptr),
	}
}

func (wc _WebBackForwardListClass) Alloc() WebBackForwardList {
	rv := objc.Call[WebBackForwardList](wc, objc.Sel("alloc"))
	return rv
}

func WebBackForwardList_Alloc() WebBackForwardList {
	return WebBackForwardListClass.Alloc()
}

func (wc _WebBackForwardListClass) New() WebBackForwardList {
	rv := objc.Call[WebBackForwardList](wc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewWebBackForwardList() WebBackForwardList {
	return WebBackForwardListClass.New()
}

func (w_ WebBackForwardList) Init() WebBackForwardList {
	rv := objc.Call[WebBackForwardList](w_, objc.Sel("init"))
	return rv
}
