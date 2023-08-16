// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMBlob] class.
var DOMBlobClass = _DOMBlobClass{objc.GetClass("DOMBlob")}

type _DOMBlobClass struct {
	objc.Class
}

// An interface definition for the [DOMBlob] class.
type IDOMBlob interface {
	IDOMObject
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domblob?language=objc
type DOMBlob struct {
	DOMObject
}

func DOMBlobFrom(ptr unsafe.Pointer) DOMBlob {
	return DOMBlob{
		DOMObject: DOMObjectFrom(ptr),
	}
}

func (dc _DOMBlobClass) Alloc() DOMBlob {
	rv := objc.Call[DOMBlob](dc, objc.Sel("alloc"))
	return rv
}

func DOMBlob_Alloc() DOMBlob {
	return DOMBlobClass.Alloc()
}

func (dc _DOMBlobClass) New() DOMBlob {
	rv := objc.Call[DOMBlob](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMBlob() DOMBlob {
	return DOMBlobClass.New()
}

func (d_ DOMBlob) Init() DOMBlob {
	rv := objc.Call[DOMBlob](d_, objc.Sel("init"))
	return rv
}
