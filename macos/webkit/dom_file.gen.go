// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMFile] class.
var DOMFileClass = _DOMFileClass{objc.GetClass("DOMFile")}

type _DOMFileClass struct {
	objc.Class
}

// An interface definition for the [DOMFile] class.
type IDOMFile interface {
	IDOMBlob
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domfile?language=objc
type DOMFile struct {
	DOMBlob
}

func DOMFileFrom(ptr unsafe.Pointer) DOMFile {
	return DOMFile{
		DOMBlob: DOMBlobFrom(ptr),
	}
}

func (dc _DOMFileClass) Alloc() DOMFile {
	rv := objc.Call[DOMFile](dc, objc.Sel("alloc"))
	return rv
}

func DOMFile_Alloc() DOMFile {
	return DOMFileClass.Alloc()
}

func (dc _DOMFileClass) New() DOMFile {
	rv := objc.Call[DOMFile](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMFile() DOMFile {
	return DOMFileClass.New()
}

func (d_ DOMFile) Init() DOMFile {
	rv := objc.Call[DOMFile](d_, objc.Sel("init"))
	return rv
}
