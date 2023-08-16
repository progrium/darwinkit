// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMComment] class.
var DOMCommentClass = _DOMCommentClass{objc.GetClass("DOMComment")}

type _DOMCommentClass struct {
	objc.Class
}

// An interface definition for the [DOMComment] class.
type IDOMComment interface {
	IDOMCharacterData
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcomment?language=objc
type DOMComment struct {
	DOMCharacterData
}

func DOMCommentFrom(ptr unsafe.Pointer) DOMComment {
	return DOMComment{
		DOMCharacterData: DOMCharacterDataFrom(ptr),
	}
}

func (dc _DOMCommentClass) Alloc() DOMComment {
	rv := objc.Call[DOMComment](dc, objc.Sel("alloc"))
	return rv
}

func DOMComment_Alloc() DOMComment {
	return DOMCommentClass.Alloc()
}

func (dc _DOMCommentClass) New() DOMComment {
	rv := objc.Call[DOMComment](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMComment() DOMComment {
	return DOMCommentClass.New()
}

func (d_ DOMComment) Init() DOMComment {
	rv := objc.Call[DOMComment](d_, objc.Sel("init"))
	return rv
}
