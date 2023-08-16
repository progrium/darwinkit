// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMCSSRule] class.
var DOMCSSRuleClass = _DOMCSSRuleClass{objc.GetClass("DOMCSSRule")}

type _DOMCSSRuleClass struct {
	objc.Class
}

// An interface definition for the [DOMCSSRule] class.
type IDOMCSSRule interface {
	IDOMObject
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssrule?language=objc
type DOMCSSRule struct {
	DOMObject
}

func DOMCSSRuleFrom(ptr unsafe.Pointer) DOMCSSRule {
	return DOMCSSRule{
		DOMObject: DOMObjectFrom(ptr),
	}
}

func (dc _DOMCSSRuleClass) Alloc() DOMCSSRule {
	rv := objc.Call[DOMCSSRule](dc, objc.Sel("alloc"))
	return rv
}

func DOMCSSRule_Alloc() DOMCSSRule {
	return DOMCSSRuleClass.Alloc()
}

func (dc _DOMCSSRuleClass) New() DOMCSSRule {
	rv := objc.Call[DOMCSSRule](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMCSSRule() DOMCSSRule {
	return DOMCSSRuleClass.New()
}

func (d_ DOMCSSRule) Init() DOMCSSRule {
	rv := objc.Call[DOMCSSRule](d_, objc.Sel("init"))
	return rv
}
