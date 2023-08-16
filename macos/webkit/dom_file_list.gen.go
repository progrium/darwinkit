// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMFileList] class.
var DOMFileListClass = _DOMFileListClass{objc.GetClass("DOMFileList")}

type _DOMFileListClass struct {
	objc.Class
}

// An interface definition for the [DOMFileList] class.
type IDOMFileList interface {
	IDOMObject
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domfilelist?language=objc
type DOMFileList struct {
	DOMObject
}

func DOMFileListFrom(ptr unsafe.Pointer) DOMFileList {
	return DOMFileList{
		DOMObject: DOMObjectFrom(ptr),
	}
}

func (dc _DOMFileListClass) Alloc() DOMFileList {
	rv := objc.Call[DOMFileList](dc, objc.Sel("alloc"))
	return rv
}

func DOMFileList_Alloc() DOMFileList {
	return DOMFileListClass.Alloc()
}

func (dc _DOMFileListClass) New() DOMFileList {
	rv := objc.Call[DOMFileList](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMFileList() DOMFileList {
	return DOMFileListClass.New()
}

func (d_ DOMFileList) Init() DOMFileList {
	rv := objc.Call[DOMFileList](d_, objc.Sel("init"))
	return rv
}
