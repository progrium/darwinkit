// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMCSSImportRule] class.
var DOMCSSImportRuleClass = _DOMCSSImportRuleClass{objc.GetClass("DOMCSSImportRule")}

type _DOMCSSImportRuleClass struct {
	objc.Class
}

// An interface definition for the [DOMCSSImportRule] class.
type IDOMCSSImportRule interface {
	IDOMCSSRule
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssimportrule?language=objc
type DOMCSSImportRule struct {
	DOMCSSRule
}

func DOMCSSImportRuleFrom(ptr unsafe.Pointer) DOMCSSImportRule {
	return DOMCSSImportRule{
		DOMCSSRule: DOMCSSRuleFrom(ptr),
	}
}

func (dc _DOMCSSImportRuleClass) Alloc() DOMCSSImportRule {
	rv := objc.Call[DOMCSSImportRule](dc, objc.Sel("alloc"))
	return rv
}

func DOMCSSImportRule_Alloc() DOMCSSImportRule {
	return DOMCSSImportRuleClass.Alloc()
}

func (dc _DOMCSSImportRuleClass) New() DOMCSSImportRule {
	rv := objc.Call[DOMCSSImportRule](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMCSSImportRule() DOMCSSImportRule {
	return DOMCSSImportRuleClass.New()
}

func (d_ DOMCSSImportRule) Init() DOMCSSImportRule {
	rv := objc.Call[DOMCSSImportRule](d_, objc.Sel("init"))
	return rv
}
