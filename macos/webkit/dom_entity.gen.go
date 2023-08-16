// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMEntity] class.
var DOMEntityClass = _DOMEntityClass{objc.GetClass("DOMEntity")}

type _DOMEntityClass struct {
	objc.Class
}

// An interface definition for the [DOMEntity] class.
type IDOMEntity interface {
	IDOMNode
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domentity?language=objc
type DOMEntity struct {
	DOMNode
}

func DOMEntityFrom(ptr unsafe.Pointer) DOMEntity {
	return DOMEntity{
		DOMNode: DOMNodeFrom(ptr),
	}
}

func (dc _DOMEntityClass) Alloc() DOMEntity {
	rv := objc.Call[DOMEntity](dc, objc.Sel("alloc"))
	return rv
}

func DOMEntity_Alloc() DOMEntity {
	return DOMEntityClass.Alloc()
}

func (dc _DOMEntityClass) New() DOMEntity {
	rv := objc.Call[DOMEntity](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMEntity() DOMEntity {
	return DOMEntityClass.New()
}

func (d_ DOMEntity) Init() DOMEntity {
	rv := objc.Call[DOMEntity](d_, objc.Sel("init"))
	return rv
}
