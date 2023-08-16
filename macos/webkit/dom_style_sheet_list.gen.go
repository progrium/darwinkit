// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMStyleSheetList] class.
var DOMStyleSheetListClass = _DOMStyleSheetListClass{objc.GetClass("DOMStyleSheetList")}

type _DOMStyleSheetListClass struct {
	objc.Class
}

// An interface definition for the [DOMStyleSheetList] class.
type IDOMStyleSheetList interface {
	IDOMObject
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domstylesheetlist?language=objc
type DOMStyleSheetList struct {
	DOMObject
}

func DOMStyleSheetListFrom(ptr unsafe.Pointer) DOMStyleSheetList {
	return DOMStyleSheetList{
		DOMObject: DOMObjectFrom(ptr),
	}
}

func (dc _DOMStyleSheetListClass) Alloc() DOMStyleSheetList {
	rv := objc.Call[DOMStyleSheetList](dc, objc.Sel("alloc"))
	return rv
}

func DOMStyleSheetList_Alloc() DOMStyleSheetList {
	return DOMStyleSheetListClass.Alloc()
}

func (dc _DOMStyleSheetListClass) New() DOMStyleSheetList {
	rv := objc.Call[DOMStyleSheetList](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMStyleSheetList() DOMStyleSheetList {
	return DOMStyleSheetListClass.New()
}

func (d_ DOMStyleSheetList) Init() DOMStyleSheetList {
	rv := objc.Call[DOMStyleSheetList](d_, objc.Sel("init"))
	return rv
}
