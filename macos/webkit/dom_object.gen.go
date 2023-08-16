// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMObject] class.
var DOMObjectClass = _DOMObjectClass{objc.GetClass("DOMObject")}

type _DOMObjectClass struct {
	objc.Class
}

// An interface definition for the [DOMObject] class.
type IDOMObject interface {
	IWebScriptObject
	Sheet() DOMStyleSheet
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domobject?language=objc
type DOMObject struct {
	WebScriptObject
}

func DOMObjectFrom(ptr unsafe.Pointer) DOMObject {
	return DOMObject{
		WebScriptObject: WebScriptObjectFrom(ptr),
	}
}

func (dc _DOMObjectClass) Alloc() DOMObject {
	rv := objc.Call[DOMObject](dc, objc.Sel("alloc"))
	return rv
}

func DOMObject_Alloc() DOMObject {
	return DOMObjectClass.Alloc()
}

func (dc _DOMObjectClass) New() DOMObject {
	rv := objc.Call[DOMObject](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMObject() DOMObject {
	return DOMObjectClass.New()
}

func (d_ DOMObject) Init() DOMObject {
	rv := objc.Call[DOMObject](d_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domobject/1521125-sheet?language=objc
func (d_ DOMObject) Sheet() DOMStyleSheet {
	rv := objc.Call[DOMStyleSheet](d_, objc.Sel("sheet"))
	return rv
}
