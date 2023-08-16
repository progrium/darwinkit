// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMCSSFontFaceRule] class.
var DOMCSSFontFaceRuleClass = _DOMCSSFontFaceRuleClass{objc.GetClass("DOMCSSFontFaceRule")}

type _DOMCSSFontFaceRuleClass struct {
	objc.Class
}

// An interface definition for the [DOMCSSFontFaceRule] class.
type IDOMCSSFontFaceRule interface {
	IDOMCSSRule
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssfontfacerule?language=objc
type DOMCSSFontFaceRule struct {
	DOMCSSRule
}

func DOMCSSFontFaceRuleFrom(ptr unsafe.Pointer) DOMCSSFontFaceRule {
	return DOMCSSFontFaceRule{
		DOMCSSRule: DOMCSSRuleFrom(ptr),
	}
}

func (dc _DOMCSSFontFaceRuleClass) Alloc() DOMCSSFontFaceRule {
	rv := objc.Call[DOMCSSFontFaceRule](dc, objc.Sel("alloc"))
	return rv
}

func DOMCSSFontFaceRule_Alloc() DOMCSSFontFaceRule {
	return DOMCSSFontFaceRuleClass.Alloc()
}

func (dc _DOMCSSFontFaceRuleClass) New() DOMCSSFontFaceRule {
	rv := objc.Call[DOMCSSFontFaceRule](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMCSSFontFaceRule() DOMCSSFontFaceRule {
	return DOMCSSFontFaceRuleClass.New()
}

func (d_ DOMCSSFontFaceRule) Init() DOMCSSFontFaceRule {
	rv := objc.Call[DOMCSSFontFaceRule](d_, objc.Sel("init"))
	return rv
}
