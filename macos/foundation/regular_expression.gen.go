// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var RegularExpressionClass = _RegularExpressionClass{objc.GetClass("NSRegularExpression")}

type _RegularExpressionClass struct {
	objc.Class
}

type IRegularExpression interface {
	objc.IObject
	NumberOfMatchesInStringOptionsRange(string_ string, options MatchingOptions, range_ Range) uint
	EnumerateMatchesInStringOptionsRangeUsingBlock(string_ string, options MatchingOptions, range_ Range, block func(result TextCheckingResult, flags MatchingFlags, stop *bool))
	MatchesInStringOptionsRange(string_ string, options MatchingOptions, range_ Range) []TextCheckingResult
	FirstMatchInStringOptionsRange(string_ string, options MatchingOptions, range_ Range) TextCheckingResult
	RangeOfFirstMatchInStringOptionsRange(string_ string, options MatchingOptions, range_ Range) Range
	StringByReplacingMatchesInStringOptionsRangeWithTemplate(string_ string, options MatchingOptions, range_ Range, templ string) string
	ReplacementStringForResultInStringOffsetTemplate(result ITextCheckingResult, string_ string, offset int, templ string) string
	Pattern() string
	Options() RegularExpressionOptions
	NumberOfCaptureGroups() uint
}

type RegularExpression struct {
	objc.Object
}

func MakeRegularExpression(ptr unsafe.Pointer) RegularExpression {
	return RegularExpression{
		Object: objc.MakeObject(ptr),
	}
}

func (r_ RegularExpression) InitWithPatternOptionsError(pattern string, options RegularExpressionOptions, error *Error) RegularExpression {
	rv := objc.CallMethod[RegularExpression](r_, objc.GetSelector("initWithPattern:options:error:"), pattern, options, unsafe.Pointer(error))
	return rv
}

func RegularExpression_InitWithPatternOptionsError(pattern string, options RegularExpressionOptions, error *Error) RegularExpression {
	return RegularExpressionClass.Alloc().InitWithPatternOptionsError(pattern, options, error)
}

func (rc _RegularExpressionClass) Alloc() RegularExpression {
	rv := objc.CallMethod[RegularExpression](rc, objc.GetSelector("alloc"))
	return rv
}

func RegularExpression_Alloc() RegularExpression {
	return RegularExpressionClass.Alloc()
}

func (rc _RegularExpressionClass) New() RegularExpression {
	rv := objc.CallMethod[RegularExpression](rc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewRegularExpression() RegularExpression {
	return RegularExpressionClass.New()
}

func RegularExpression_New() RegularExpression {
	return RegularExpressionClass.New()
}

func (r_ RegularExpression) Init() RegularExpression {
	rv := objc.CallMethod[RegularExpression](r_, objc.GetSelector("init"))
	return rv
}

func RegularExpression_Init() RegularExpression {
	return RegularExpressionClass.Alloc().Init()
}

func (rc _RegularExpressionClass) RegularExpressionWithPatternOptionsError(pattern string, options RegularExpressionOptions, error *Error) RegularExpression {
	rv := objc.CallMethod[RegularExpression](rc, objc.GetSelector("regularExpressionWithPattern:options:error:"), pattern, options, unsafe.Pointer(error))
	return rv
}

func RegularExpression_RegularExpressionWithPatternOptionsError(pattern string, options RegularExpressionOptions, error *Error) RegularExpression {
	return RegularExpressionClass.RegularExpressionWithPatternOptionsError(pattern, options, error)
}

func (r_ RegularExpression) NumberOfMatchesInStringOptionsRange(string_ string, options MatchingOptions, range_ Range) uint {
	rv := objc.CallMethod[uint](r_, objc.GetSelector("numberOfMatchesInString:options:range:"), string_, options, range_)
	return rv
}

func (r_ RegularExpression) EnumerateMatchesInStringOptionsRangeUsingBlock(string_ string, options MatchingOptions, range_ Range, block func(result TextCheckingResult, flags MatchingFlags, stop *bool)) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("enumerateMatchesInString:options:range:usingBlock:"), string_, options, range_, block)
}

func (r_ RegularExpression) MatchesInStringOptionsRange(string_ string, options MatchingOptions, range_ Range) []TextCheckingResult {
	rv := objc.CallMethod[[]TextCheckingResult](r_, objc.GetSelector("matchesInString:options:range:"), string_, options, range_)
	return rv
}

func (r_ RegularExpression) FirstMatchInStringOptionsRange(string_ string, options MatchingOptions, range_ Range) TextCheckingResult {
	rv := objc.CallMethod[TextCheckingResult](r_, objc.GetSelector("firstMatchInString:options:range:"), string_, options, range_)
	return rv
}

func (r_ RegularExpression) RangeOfFirstMatchInStringOptionsRange(string_ string, options MatchingOptions, range_ Range) Range {
	rv := objc.CallMethod[Range](r_, objc.GetSelector("rangeOfFirstMatchInString:options:range:"), string_, options, range_)
	return rv
}

func (r_ RegularExpression) StringByReplacingMatchesInStringOptionsRangeWithTemplate(string_ string, options MatchingOptions, range_ Range, templ string) string {
	rv := objc.CallMethod[string](r_, objc.GetSelector("stringByReplacingMatchesInString:options:range:withTemplate:"), string_, options, range_, templ)
	return rv
}

func (rc _RegularExpressionClass) EscapedTemplateForString(string_ string) string {
	rv := objc.CallMethod[string](rc, objc.GetSelector("escapedTemplateForString:"), string_)
	return rv
}

func RegularExpression_EscapedTemplateForString(string_ string) string {
	return RegularExpressionClass.EscapedTemplateForString(string_)
}

func (rc _RegularExpressionClass) EscapedPatternForString(string_ string) string {
	rv := objc.CallMethod[string](rc, objc.GetSelector("escapedPatternForString:"), string_)
	return rv
}

func RegularExpression_EscapedPatternForString(string_ string) string {
	return RegularExpressionClass.EscapedPatternForString(string_)
}

func (r_ RegularExpression) ReplacementStringForResultInStringOffsetTemplate(result ITextCheckingResult, string_ string, offset int, templ string) string {
	rv := objc.CallMethod[string](r_, objc.GetSelector("replacementStringForResult:inString:offset:template:"), objc.ExtractPtr(result), string_, offset, templ)
	return rv
}

func (r_ RegularExpression) Pattern() string {
	rv := objc.CallMethod[string](r_, objc.GetSelector("pattern"))
	return rv
}

func (r_ RegularExpression) Options() RegularExpressionOptions {
	rv := objc.CallMethod[RegularExpressionOptions](r_, objc.GetSelector("options"))
	return rv
}

func (r_ RegularExpression) NumberOfCaptureGroups() uint {
	rv := objc.CallMethod[uint](r_, objc.GetSelector("numberOfCaptureGroups"))
	return rv
}
