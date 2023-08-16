// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMCSSCharsetRule] class.
var DOMCSSCharsetRuleClass = _DOMCSSCharsetRuleClass{objc.GetClass("DOMCSSCharsetRule")}

type _DOMCSSCharsetRuleClass struct {
	objc.Class
}

// An interface definition for the [DOMCSSCharsetRule] class.
type IDOMCSSCharsetRule interface {
	IDOMCSSRule
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcsscharsetrule?language=objc
type DOMCSSCharsetRule struct {
	DOMCSSRule
}

func DOMCSSCharsetRuleFrom(ptr unsafe.Pointer) DOMCSSCharsetRule {
	return DOMCSSCharsetRule{
		DOMCSSRule: DOMCSSRuleFrom(ptr),
	}
}

func (dc _DOMCSSCharsetRuleClass) Alloc() DOMCSSCharsetRule {
	rv := objc.Call[DOMCSSCharsetRule](dc, objc.Sel("alloc"))
	return rv
}

func DOMCSSCharsetRule_Alloc() DOMCSSCharsetRule {
	return DOMCSSCharsetRuleClass.Alloc()
}

func (dc _DOMCSSCharsetRuleClass) New() DOMCSSCharsetRule {
	rv := objc.Call[DOMCSSCharsetRule](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMCSSCharsetRule() DOMCSSCharsetRule {
	return DOMCSSCharsetRuleClass.New()
}

func (d_ DOMCSSCharsetRule) Init() DOMCSSCharsetRule {
	rv := objc.Call[DOMCSSCharsetRule](d_, objc.Sel("init"))
	return rv
}
