// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMCSSStyleRule] class.
var DOMCSSStyleRuleClass = _DOMCSSStyleRuleClass{objc.GetClass("DOMCSSStyleRule")}

type _DOMCSSStyleRuleClass struct {
	objc.Class
}

// An interface definition for the [DOMCSSStyleRule] class.
type IDOMCSSStyleRule interface {
	IDOMCSSRule
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssstylerule?language=objc
type DOMCSSStyleRule struct {
	DOMCSSRule
}

func DOMCSSStyleRuleFrom(ptr unsafe.Pointer) DOMCSSStyleRule {
	return DOMCSSStyleRule{
		DOMCSSRule: DOMCSSRuleFrom(ptr),
	}
}

func (dc _DOMCSSStyleRuleClass) Alloc() DOMCSSStyleRule {
	rv := objc.Call[DOMCSSStyleRule](dc, objc.Sel("alloc"))
	return rv
}

func DOMCSSStyleRule_Alloc() DOMCSSStyleRule {
	return DOMCSSStyleRuleClass.Alloc()
}

func (dc _DOMCSSStyleRuleClass) New() DOMCSSStyleRule {
	rv := objc.Call[DOMCSSStyleRule](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMCSSStyleRule() DOMCSSStyleRule {
	return DOMCSSStyleRuleClass.New()
}

func (d_ DOMCSSStyleRule) Init() DOMCSSStyleRule {
	rv := objc.Call[DOMCSSStyleRule](d_, objc.Sel("init"))
	return rv
}
