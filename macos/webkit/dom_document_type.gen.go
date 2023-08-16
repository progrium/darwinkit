// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMDocumentType] class.
var DOMDocumentTypeClass = _DOMDocumentTypeClass{objc.GetClass("DOMDocumentType")}

type _DOMDocumentTypeClass struct {
	objc.Class
}

// An interface definition for the [DOMDocumentType] class.
type IDOMDocumentType interface {
	IDOMNode
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domdocumenttype?language=objc
type DOMDocumentType struct {
	DOMNode
}

func DOMDocumentTypeFrom(ptr unsafe.Pointer) DOMDocumentType {
	return DOMDocumentType{
		DOMNode: DOMNodeFrom(ptr),
	}
}

func (dc _DOMDocumentTypeClass) Alloc() DOMDocumentType {
	rv := objc.Call[DOMDocumentType](dc, objc.Sel("alloc"))
	return rv
}

func DOMDocumentType_Alloc() DOMDocumentType {
	return DOMDocumentTypeClass.Alloc()
}

func (dc _DOMDocumentTypeClass) New() DOMDocumentType {
	rv := objc.Call[DOMDocumentType](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMDocumentType() DOMDocumentType {
	return DOMDocumentTypeClass.New()
}

func (d_ DOMDocumentType) Init() DOMDocumentType {
	rv := objc.Call[DOMDocumentType](d_, objc.Sel("init"))
	return rv
}
