// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMCSSUnknownRule] class.
var DOMCSSUnknownRuleClass = _DOMCSSUnknownRuleClass{objc.GetClass("DOMCSSUnknownRule")}

type _DOMCSSUnknownRuleClass struct {
	objc.Class
}

// An interface definition for the [DOMCSSUnknownRule] class.
type IDOMCSSUnknownRule interface {
	IDOMCSSRule
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssunknownrule?language=objc
type DOMCSSUnknownRule struct {
	DOMCSSRule
}

func DOMCSSUnknownRuleFrom(ptr unsafe.Pointer) DOMCSSUnknownRule {
	return DOMCSSUnknownRule{
		DOMCSSRule: DOMCSSRuleFrom(ptr),
	}
}

func (dc _DOMCSSUnknownRuleClass) Alloc() DOMCSSUnknownRule {
	rv := objc.Call[DOMCSSUnknownRule](dc, objc.Sel("alloc"))
	return rv
}

func DOMCSSUnknownRule_Alloc() DOMCSSUnknownRule {
	return DOMCSSUnknownRuleClass.Alloc()
}

func (dc _DOMCSSUnknownRuleClass) New() DOMCSSUnknownRule {
	rv := objc.Call[DOMCSSUnknownRule](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMCSSUnknownRule() DOMCSSUnknownRule {
	return DOMCSSUnknownRuleClass.New()
}

func (d_ DOMCSSUnknownRule) Init() DOMCSSUnknownRule {
	rv := objc.Call[DOMCSSUnknownRule](d_, objc.Sel("init"))
	return rv
}
