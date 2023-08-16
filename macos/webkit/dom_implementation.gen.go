// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMImplementation] class.
var DOMImplementationClass = _DOMImplementationClass{objc.GetClass("DOMImplementation")}

type _DOMImplementationClass struct {
	objc.Class
}

// An interface definition for the [DOMImplementation] class.
type IDOMImplementation interface {
	IDOMObject
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domimplementation?language=objc
type DOMImplementation struct {
	DOMObject
}

func DOMImplementationFrom(ptr unsafe.Pointer) DOMImplementation {
	return DOMImplementation{
		DOMObject: DOMObjectFrom(ptr),
	}
}

func (dc _DOMImplementationClass) Alloc() DOMImplementation {
	rv := objc.Call[DOMImplementation](dc, objc.Sel("alloc"))
	return rv
}

func DOMImplementation_Alloc() DOMImplementation {
	return DOMImplementationClass.Alloc()
}

func (dc _DOMImplementationClass) New() DOMImplementation {
	rv := objc.Call[DOMImplementation](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMImplementation() DOMImplementation {
	return DOMImplementationClass.New()
}

func (d_ DOMImplementation) Init() DOMImplementation {
	rv := objc.Call[DOMImplementation](d_, objc.Sel("init"))
	return rv
}
