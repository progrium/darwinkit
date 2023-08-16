// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMHTMLOptionsCollection] class.
var DOMHTMLOptionsCollectionClass = _DOMHTMLOptionsCollectionClass{objc.GetClass("DOMHTMLOptionsCollection")}

type _DOMHTMLOptionsCollectionClass struct {
	objc.Class
}

// An interface definition for the [DOMHTMLOptionsCollection] class.
type IDOMHTMLOptionsCollection interface {
	IDOMObject
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domhtmloptionscollection?language=objc
type DOMHTMLOptionsCollection struct {
	DOMObject
}

func DOMHTMLOptionsCollectionFrom(ptr unsafe.Pointer) DOMHTMLOptionsCollection {
	return DOMHTMLOptionsCollection{
		DOMObject: DOMObjectFrom(ptr),
	}
}

func (dc _DOMHTMLOptionsCollectionClass) Alloc() DOMHTMLOptionsCollection {
	rv := objc.Call[DOMHTMLOptionsCollection](dc, objc.Sel("alloc"))
	return rv
}

func DOMHTMLOptionsCollection_Alloc() DOMHTMLOptionsCollection {
	return DOMHTMLOptionsCollectionClass.Alloc()
}

func (dc _DOMHTMLOptionsCollectionClass) New() DOMHTMLOptionsCollection {
	rv := objc.Call[DOMHTMLOptionsCollection](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMHTMLOptionsCollection() DOMHTMLOptionsCollection {
	return DOMHTMLOptionsCollectionClass.New()
}

func (d_ DOMHTMLOptionsCollection) Init() DOMHTMLOptionsCollection {
	rv := objc.Call[DOMHTMLOptionsCollection](d_, objc.Sel("init"))
	return rv
}
