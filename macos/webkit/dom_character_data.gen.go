// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMCharacterData] class.
var DOMCharacterDataClass = _DOMCharacterDataClass{objc.GetClass("DOMCharacterData")}

type _DOMCharacterDataClass struct {
	objc.Class
}

// An interface definition for the [DOMCharacterData] class.
type IDOMCharacterData interface {
	IDOMNode
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcharacterdata?language=objc
type DOMCharacterData struct {
	DOMNode
}

func DOMCharacterDataFrom(ptr unsafe.Pointer) DOMCharacterData {
	return DOMCharacterData{
		DOMNode: DOMNodeFrom(ptr),
	}
}

func (dc _DOMCharacterDataClass) Alloc() DOMCharacterData {
	rv := objc.Call[DOMCharacterData](dc, objc.Sel("alloc"))
	return rv
}

func DOMCharacterData_Alloc() DOMCharacterData {
	return DOMCharacterDataClass.Alloc()
}

func (dc _DOMCharacterDataClass) New() DOMCharacterData {
	rv := objc.Call[DOMCharacterData](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMCharacterData() DOMCharacterData {
	return DOMCharacterDataClass.New()
}

func (d_ DOMCharacterData) Init() DOMCharacterData {
	rv := objc.Call[DOMCharacterData](d_, objc.Sel("init"))
	return rv
}
