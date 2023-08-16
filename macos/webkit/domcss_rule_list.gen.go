// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMCSSRuleList] class.
var DOMCSSRuleListClass = _DOMCSSRuleListClass{objc.GetClass("DOMCSSRuleList")}

type _DOMCSSRuleListClass struct {
	objc.Class
}

// An interface definition for the [DOMCSSRuleList] class.
type IDOMCSSRuleList interface {
	IDOMObject
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssrulelist?language=objc
type DOMCSSRuleList struct {
	DOMObject
}

func DOMCSSRuleListFrom(ptr unsafe.Pointer) DOMCSSRuleList {
	return DOMCSSRuleList{
		DOMObject: DOMObjectFrom(ptr),
	}
}

func (dc _DOMCSSRuleListClass) Alloc() DOMCSSRuleList {
	rv := objc.Call[DOMCSSRuleList](dc, objc.Sel("alloc"))
	return rv
}

func DOMCSSRuleList_Alloc() DOMCSSRuleList {
	return DOMCSSRuleListClass.Alloc()
}

func (dc _DOMCSSRuleListClass) New() DOMCSSRuleList {
	rv := objc.Call[DOMCSSRuleList](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMCSSRuleList() DOMCSSRuleList {
	return DOMCSSRuleListClass.New()
}

func (d_ DOMCSSRuleList) Init() DOMCSSRuleList {
	rv := objc.Call[DOMCSSRuleList](d_, objc.Sel("init"))
	return rv
}
