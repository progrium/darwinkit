// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMEntityReference] class.
var DOMEntityReferenceClass = _DOMEntityReferenceClass{objc.GetClass("DOMEntityReference")}

type _DOMEntityReferenceClass struct {
	objc.Class
}

// An interface definition for the [DOMEntityReference] class.
type IDOMEntityReference interface {
	IDOMNode
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domentityreference?language=objc
type DOMEntityReference struct {
	DOMNode
}

func DOMEntityReferenceFrom(ptr unsafe.Pointer) DOMEntityReference {
	return DOMEntityReference{
		DOMNode: DOMNodeFrom(ptr),
	}
}

func (dc _DOMEntityReferenceClass) Alloc() DOMEntityReference {
	rv := objc.Call[DOMEntityReference](dc, objc.Sel("alloc"))
	return rv
}

func DOMEntityReference_Alloc() DOMEntityReference {
	return DOMEntityReferenceClass.Alloc()
}

func (dc _DOMEntityReferenceClass) New() DOMEntityReference {
	rv := objc.Call[DOMEntityReference](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMEntityReference() DOMEntityReference {
	return DOMEntityReferenceClass.New()
}

func (d_ DOMEntityReference) Init() DOMEntityReference {
	rv := objc.Call[DOMEntityReference](d_, objc.Sel("init"))
	return rv
}
