// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMNodeList] class.
var DOMNodeListClass = _DOMNodeListClass{objc.GetClass("DOMNodeList")}

type _DOMNodeListClass struct {
	objc.Class
}

// An interface definition for the [DOMNodeList] class.
type IDOMNodeList interface {
	IDOMObject
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domnodelist?language=objc
type DOMNodeList struct {
	DOMObject
}

func DOMNodeListFrom(ptr unsafe.Pointer) DOMNodeList {
	return DOMNodeList{
		DOMObject: DOMObjectFrom(ptr),
	}
}

func (dc _DOMNodeListClass) Alloc() DOMNodeList {
	rv := objc.Call[DOMNodeList](dc, objc.Sel("alloc"))
	return rv
}

func DOMNodeList_Alloc() DOMNodeList {
	return DOMNodeListClass.Alloc()
}

func (dc _DOMNodeListClass) New() DOMNodeList {
	rv := objc.Call[DOMNodeList](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMNodeList() DOMNodeList {
	return DOMNodeListClass.New()
}

func (d_ DOMNodeList) Init() DOMNodeList {
	rv := objc.Call[DOMNodeList](d_, objc.Sel("init"))
	return rv
}
