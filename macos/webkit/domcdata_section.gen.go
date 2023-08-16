// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMCDATASection] class.
var DOMCDATASectionClass = _DOMCDATASectionClass{objc.GetClass("DOMCDATASection")}

type _DOMCDATASectionClass struct {
	objc.Class
}

// An interface definition for the [DOMCDATASection] class.
type IDOMCDATASection interface {
	IDOMText
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcdatasection?language=objc
type DOMCDATASection struct {
	DOMText
}

func DOMCDATASectionFrom(ptr unsafe.Pointer) DOMCDATASection {
	return DOMCDATASection{
		DOMText: DOMTextFrom(ptr),
	}
}

func (dc _DOMCDATASectionClass) Alloc() DOMCDATASection {
	rv := objc.Call[DOMCDATASection](dc, objc.Sel("alloc"))
	return rv
}

func DOMCDATASection_Alloc() DOMCDATASection {
	return DOMCDATASectionClass.Alloc()
}

func (dc _DOMCDATASectionClass) New() DOMCDATASection {
	rv := objc.Call[DOMCDATASection](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMCDATASection() DOMCDATASection {
	return DOMCDATASectionClass.New()
}

func (d_ DOMCDATASection) Init() DOMCDATASection {
	rv := objc.Call[DOMCDATASection](d_, objc.Sel("init"))
	return rv
}
