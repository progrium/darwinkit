// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMXPathResult] class.
var DOMXPathResultClass = _DOMXPathResultClass{objc.GetClass("DOMXPathResult")}

type _DOMXPathResultClass struct {
	objc.Class
}

// An interface definition for the [DOMXPathResult] class.
type IDOMXPathResult interface {
	IDOMObject
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domxpathresult?language=objc
type DOMXPathResult struct {
	DOMObject
}

func DOMXPathResultFrom(ptr unsafe.Pointer) DOMXPathResult {
	return DOMXPathResult{
		DOMObject: DOMObjectFrom(ptr),
	}
}

func (dc _DOMXPathResultClass) Alloc() DOMXPathResult {
	rv := objc.Call[DOMXPathResult](dc, objc.Sel("alloc"))
	return rv
}

func DOMXPathResult_Alloc() DOMXPathResult {
	return DOMXPathResultClass.Alloc()
}

func (dc _DOMXPathResultClass) New() DOMXPathResult {
	rv := objc.Call[DOMXPathResult](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMXPathResult() DOMXPathResult {
	return DOMXPathResultClass.New()
}

func (d_ DOMXPathResult) Init() DOMXPathResult {
	rv := objc.Call[DOMXPathResult](d_, objc.Sel("init"))
	return rv
}
