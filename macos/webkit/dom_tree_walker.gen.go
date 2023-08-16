// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMTreeWalker] class.
var DOMTreeWalkerClass = _DOMTreeWalkerClass{objc.GetClass("DOMTreeWalker")}

type _DOMTreeWalkerClass struct {
	objc.Class
}

// An interface definition for the [DOMTreeWalker] class.
type IDOMTreeWalker interface {
	IDOMObject
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domtreewalker?language=objc
type DOMTreeWalker struct {
	DOMObject
}

func DOMTreeWalkerFrom(ptr unsafe.Pointer) DOMTreeWalker {
	return DOMTreeWalker{
		DOMObject: DOMObjectFrom(ptr),
	}
}

func (dc _DOMTreeWalkerClass) Alloc() DOMTreeWalker {
	rv := objc.Call[DOMTreeWalker](dc, objc.Sel("alloc"))
	return rv
}

func DOMTreeWalker_Alloc() DOMTreeWalker {
	return DOMTreeWalkerClass.Alloc()
}

func (dc _DOMTreeWalkerClass) New() DOMTreeWalker {
	rv := objc.Call[DOMTreeWalker](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMTreeWalker() DOMTreeWalker {
	return DOMTreeWalkerClass.New()
}

func (d_ DOMTreeWalker) Init() DOMTreeWalker {
	rv := objc.Call[DOMTreeWalker](d_, objc.Sel("init"))
	return rv
}
