// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMNamedNodeMap] class.
var DOMNamedNodeMapClass = _DOMNamedNodeMapClass{objc.GetClass("DOMNamedNodeMap")}

type _DOMNamedNodeMapClass struct {
	objc.Class
}

// An interface definition for the [DOMNamedNodeMap] class.
type IDOMNamedNodeMap interface {
	IDOMObject
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domnamednodemap?language=objc
type DOMNamedNodeMap struct {
	DOMObject
}

func DOMNamedNodeMapFrom(ptr unsafe.Pointer) DOMNamedNodeMap {
	return DOMNamedNodeMap{
		DOMObject: DOMObjectFrom(ptr),
	}
}

func (dc _DOMNamedNodeMapClass) Alloc() DOMNamedNodeMap {
	rv := objc.Call[DOMNamedNodeMap](dc, objc.Sel("alloc"))
	return rv
}

func DOMNamedNodeMap_Alloc() DOMNamedNodeMap {
	return DOMNamedNodeMapClass.Alloc()
}

func (dc _DOMNamedNodeMapClass) New() DOMNamedNodeMap {
	rv := objc.Call[DOMNamedNodeMap](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMNamedNodeMap() DOMNamedNodeMap {
	return DOMNamedNodeMapClass.New()
}

func (d_ DOMNamedNodeMap) Init() DOMNamedNodeMap {
	rv := objc.Call[DOMNamedNodeMap](d_, objc.Sel("init"))
	return rv
}
