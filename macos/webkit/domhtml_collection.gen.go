// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLCollection] class.
var DOMHTMLCollectionClass = _DOMHTMLCollectionClass{objc.GetClass("DOMHTMLCollection")}

type _DOMHTMLCollectionClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLCollection] class.
type IDOMHTMLCollection interface {
	IDOMObject
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmlcollection?language=objc
type DOMHTMLCollection struct {
	DOMObject
}

func DOMHTMLCollectionFrom(ptr unsafe.Pointer) DOMHTMLCollection {
	return DOMHTMLCollection{
		DOMObject: DOMObjectFrom(ptr),
	}
}

func (dc _DOMHTMLCollectionClass) Alloc() DOMHTMLCollection {
	rv := objc.Call[DOMHTMLCollection](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLCollection_Alloc() DOMHTMLCollection {
	return DOMHTMLCollectionClass.Alloc()
}

func (dc _DOMHTMLCollectionClass) New() DOMHTMLCollection {
	rv := objc.Call[DOMHTMLCollection](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLCollection() DOMHTMLCollection {
	return DOMHTMLCollectionClass.New()
}

func (d_ DOMHTMLCollection) Init() DOMHTMLCollection {
	rv := objc.Call[DOMHTMLCollection](d_, objc.Sel("init"))
	return rv
}
