// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMCSSPageRule] class.
var DOMCSSPageRuleClass = _DOMCSSPageRuleClass{objc.GetClass("DOMCSSPageRule")}

type _DOMCSSPageRuleClass struct {
	objc.Class
}

// An interface definition for the [DOMCSSPageRule] class.
type IDOMCSSPageRule interface {
	IDOMCSSRule
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcsspagerule?language=objc
type DOMCSSPageRule struct {
	DOMCSSRule
}

func DOMCSSPageRuleFrom(ptr unsafe.Pointer) DOMCSSPageRule {
	return DOMCSSPageRule{
		DOMCSSRule: DOMCSSRuleFrom(ptr),
	}
}

func (dc _DOMCSSPageRuleClass) Alloc() DOMCSSPageRule {
	rv := objc.Call[DOMCSSPageRule](dc, objc.Sel("alloc"))
	return rv
}

func DOMCSSPageRule_Alloc() DOMCSSPageRule {
	return DOMCSSPageRuleClass.Alloc()
}

func (dc _DOMCSSPageRuleClass) New() DOMCSSPageRule {
	rv := objc.Call[DOMCSSPageRule](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMCSSPageRule() DOMCSSPageRule {
	return DOMCSSPageRuleClass.New()
}

func (d_ DOMCSSPageRule) Init() DOMCSSPageRule {
	rv := objc.Call[DOMCSSPageRule](d_, objc.Sel("init"))
	return rv
}
