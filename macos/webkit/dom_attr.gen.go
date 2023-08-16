// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMAttr] class.
var DOMAttrClass = _DOMAttrClass{objc.GetClass("DOMAttr")}

type _DOMAttrClass struct {
	objc.Class
}

// An interface definition for the [DOMAttr] class.
type IDOMAttr interface {
	IDOMNode
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domattr?language=objc
type DOMAttr struct {
	DOMNode
}

func DOMAttrFrom(ptr unsafe.Pointer) DOMAttr {
	return DOMAttr{
		DOMNode: DOMNodeFrom(ptr),
	}
}

func (dc _DOMAttrClass) Alloc() DOMAttr {
	rv := objc.Call[DOMAttr](dc, objc.Sel("alloc"))
	return rv
}

func DOMAttr_Alloc() DOMAttr {
	return DOMAttrClass.Alloc()
}

func (dc _DOMAttrClass) New() DOMAttr {
	rv := objc.Call[DOMAttr](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMAttr() DOMAttr {
	return DOMAttrClass.New()
}

func (d_ DOMAttr) Init() DOMAttr {
	rv := objc.Call[DOMAttr](d_, objc.Sel("init"))
	return rv
}
