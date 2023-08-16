// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMMediaList] class.
var DOMMediaListClass = _DOMMediaListClass{objc.GetClass("DOMMediaList")}

type _DOMMediaListClass struct {
	objc.Class
}

// An interface definition for the [DOMMediaList] class.
type IDOMMediaList interface {
	IDOMObject
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/dommedialist?language=objc
type DOMMediaList struct {
	DOMObject
}

func DOMMediaListFrom(ptr unsafe.Pointer) DOMMediaList {
	return DOMMediaList{
		DOMObject: DOMObjectFrom(ptr),
	}
}

func (dc _DOMMediaListClass) Alloc() DOMMediaList {
	rv := objc.Call[DOMMediaList](dc, objc.Sel("alloc"))
	return rv
}

func DOMMediaList_Alloc() DOMMediaList {
	return DOMMediaListClass.Alloc()
}

func (dc _DOMMediaListClass) New() DOMMediaList {
	rv := objc.Call[DOMMediaList](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMMediaList() DOMMediaList {
	return DOMMediaListClass.New()
}

func (d_ DOMMediaList) Init() DOMMediaList {
	rv := objc.Call[DOMMediaList](d_, objc.Sel("init"))
	return rv
}
