// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMNodeIterator] class.
var DOMNodeIteratorClass = _DOMNodeIteratorClass{objc.GetClass("DOMNodeIterator")}

type _DOMNodeIteratorClass struct {
	objc.Class
}

// An interface definition for the [DOMNodeIterator] class.
type IDOMNodeIterator interface {
	IDOMObject
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domnodeiterator?language=objc
type DOMNodeIterator struct {
	DOMObject
}

func DOMNodeIteratorFrom(ptr unsafe.Pointer) DOMNodeIterator {
	return DOMNodeIterator{
		DOMObject: DOMObjectFrom(ptr),
	}
}

func (dc _DOMNodeIteratorClass) Alloc() DOMNodeIterator {
	rv := objc.Call[DOMNodeIterator](dc, objc.Sel("alloc"))
	return rv
}

func DOMNodeIterator_Alloc() DOMNodeIterator {
	return DOMNodeIteratorClass.Alloc()
}

func (dc _DOMNodeIteratorClass) New() DOMNodeIterator {
	rv := objc.Call[DOMNodeIterator](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMNodeIterator() DOMNodeIterator {
	return DOMNodeIteratorClass.New()
}

func (d_ DOMNodeIterator) Init() DOMNodeIterator {
	rv := objc.Call[DOMNodeIterator](d_, objc.Sel("init"))
	return rv
}
