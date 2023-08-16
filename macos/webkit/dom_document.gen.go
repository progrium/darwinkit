// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMDocument] class.
var DOMDocumentClass = _DOMDocumentClass{objc.GetClass("DOMDocument")}

type _DOMDocumentClass struct {
	objc.Class
}

// An interface definition for the [DOMDocument] class.
type IDOMDocument interface {
	IDOMNode
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domdocument?language=objc
type DOMDocument struct {
	DOMNode
}

func DOMDocumentFrom(ptr unsafe.Pointer) DOMDocument {
	return DOMDocument{
		DOMNode: DOMNodeFrom(ptr),
	}
}

func (dc _DOMDocumentClass) Alloc() DOMDocument {
	rv := objc.Call[DOMDocument](dc, objc.Sel("alloc"))
	return rv
}

func DOMDocument_Alloc() DOMDocument {
	return DOMDocumentClass.Alloc()
}

func (dc _DOMDocumentClass) New() DOMDocument {
	rv := objc.Call[DOMDocument](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMDocument() DOMDocument {
	return DOMDocumentClass.New()
}

func (d_ DOMDocument) Init() DOMDocument {
	rv := objc.Call[DOMDocument](d_, objc.Sel("init"))
	return rv
}
