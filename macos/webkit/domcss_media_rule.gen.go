// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMCSSMediaRule] class.
var DOMCSSMediaRuleClass = _DOMCSSMediaRuleClass{objc.GetClass("DOMCSSMediaRule")}

type _DOMCSSMediaRuleClass struct {
	objc.Class
}

// An interface definition for the [DOMCSSMediaRule] class.
type IDOMCSSMediaRule interface {
	IDOMCSSRule
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssmediarule?language=objc
type DOMCSSMediaRule struct {
	DOMCSSRule
}

func DOMCSSMediaRuleFrom(ptr unsafe.Pointer) DOMCSSMediaRule {
	return DOMCSSMediaRule{
		DOMCSSRule: DOMCSSRuleFrom(ptr),
	}
}

func (dc _DOMCSSMediaRuleClass) Alloc() DOMCSSMediaRule {
	rv := objc.Call[DOMCSSMediaRule](dc, objc.Sel("alloc"))
	return rv
}

func DOMCSSMediaRule_Alloc() DOMCSSMediaRule {
	return DOMCSSMediaRuleClass.Alloc()
}

func (dc _DOMCSSMediaRuleClass) New() DOMCSSMediaRule {
	rv := objc.Call[DOMCSSMediaRule](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMCSSMediaRule() DOMCSSMediaRule {
	return DOMCSSMediaRuleClass.New()
}

func (d_ DOMCSSMediaRule) Init() DOMCSSMediaRule {
	rv := objc.Call[DOMCSSMediaRule](d_, objc.Sel("init"))
	return rv
}
