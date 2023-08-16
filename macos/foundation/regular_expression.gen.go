// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RegularExpression] class.
var RegularExpressionClass = _RegularExpressionClass{objc.GetClass("NSRegularExpression")}

type _RegularExpressionClass struct {
	objc.Class
}

// An interface definition for the [RegularExpression] class.
type IRegularExpression interface {
	objc.IObject
	ReplaceMatchesInStringOptionsRangeWithTemplate(string_ IMutableString, options MatchingOptions, range_ Range, templ string) uint
	RangeOfFirstMatchInStringOptionsRange(string_ string, options MatchingOptions, range_ Range) Range
	ReplacementStringForResultInStringOffsetTemplate(result ITextCheckingResult, string_ string, offset int, templ string) string
	FirstMatchInStringOptionsRange(string_ string, options MatchingOptions, range_ Range) TextCheckingResult
	NumberOfMatchesInStringOptionsRange(string_ string, options MatchingOptions, range_ Range) uint
	StringByReplacingMatchesInStringOptionsRangeWithTemplate(string_ string, options MatchingOptions, range_ Range, templ string) string
	MatchesInStringOptionsRange(string_ string, options MatchingOptions, range_ Range) []TextCheckingResult
	EnumerateMatchesInStringOptionsRangeUsingBlock(string_ string, options MatchingOptions, range_ Range, block func(result TextCheckingResult, flags MatchingFlags, stop *bool))
	Options() RegularExpressionOptions
	NumberOfCaptureGroups() uint
	Pattern() string
}

// An immutable representation of a compiled regular expression that you apply to Unicode strings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsregularexpression?language=objc
type RegularExpression struct {
	objc.Object
}

func RegularExpressionFrom(ptr unsafe.Pointer) RegularExpression {
	return RegularExpression{
		Object: objc.ObjectFrom(ptr),
	}
}

func (r_ RegularExpression) InitWithPatternOptionsError(pattern string, options RegularExpressionOptions, error IError) RegularExpression {
	rv := objc.Call[RegularExpression](r_, objc.Sel("initWithPattern:options:error:"), pattern, options, objc.Ptr(error))
	return rv
}

// Returns an initialized NSRegularExpression instance with the specified regular expression pattern and options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsregularexpression/1410900-initwithpattern?language=objc
func RegularExpression_InitWithPatternOptionsError(pattern string, options RegularExpressionOptions, error IError) RegularExpression {
	return RegularExpressionClass.Alloc().InitWithPatternOptionsError(pattern, options, error)
}

func (rc _RegularExpressionClass) Alloc() RegularExpression {
	rv := objc.Call[RegularExpression](rc, objc.Sel("alloc"))
	return rv
}

func RegularExpression_Alloc() RegularExpression {
	return RegularExpressionClass.Alloc()
}

func (rc _RegularExpressionClass) New() RegularExpression {
	rv := objc.Call[RegularExpression](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRegularExpression() RegularExpression {
	return RegularExpressionClass.New()
}

func (r_ RegularExpression) Init() RegularExpression {
	rv := objc.Call[RegularExpression](r_, objc.Sel("init"))
	return rv
}

// Replaces regular expression matches within the mutable string using the template string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsregularexpression/1411139-replacematchesinstring?language=objc
func (r_ RegularExpression) ReplaceMatchesInStringOptionsRangeWithTemplate(string_ IMutableString, options MatchingOptions, range_ Range, templ string) uint {
	rv := objc.Call[uint](r_, objc.Sel("replaceMatchesInString:options:range:withTemplate:"), objc.Ptr(string_), options, range_, templ)
	return rv
}

// Returns the range of the first match of the regular expression within the specified range of the string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsregularexpression/1411510-rangeoffirstmatchinstring?language=objc
func (r_ RegularExpression) RangeOfFirstMatchInStringOptionsRange(string_ string, options MatchingOptions, range_ Range) Range {
	rv := objc.Call[Range](r_, objc.Sel("rangeOfFirstMatchInString:options:range:"), string_, options, range_)
	return rv
}

// Returns a string by adding backslash escapes as necessary to protect any characters that would match as pattern metacharacters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsregularexpression/1408386-escapedpatternforstring?language=objc
func (rc _RegularExpressionClass) EscapedPatternForString(string_ string) string {
	rv := objc.Call[string](rc, objc.Sel("escapedPatternForString:"), string_)
	return rv
}

// Returns a string by adding backslash escapes as necessary to protect any characters that would match as pattern metacharacters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsregularexpression/1408386-escapedpatternforstring?language=objc
func RegularExpression_EscapedPatternForString(string_ string) string {
	return RegularExpressionClass.EscapedPatternForString(string_)
}

// Used to perform template substitution for a single result for clients implementing their own replace functionality. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsregularexpression/1414859-replacementstringforresult?language=objc
func (r_ RegularExpression) ReplacementStringForResultInStringOffsetTemplate(result ITextCheckingResult, string_ string, offset int, templ string) string {
	rv := objc.Call[string](r_, objc.Sel("replacementStringForResult:inString:offset:template:"), objc.Ptr(result), string_, offset, templ)
	return rv
}

// Returns the first match of the regular expression within the specified range of the string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsregularexpression/1409838-firstmatchinstring?language=objc
func (r_ RegularExpression) FirstMatchInStringOptionsRange(string_ string, options MatchingOptions, range_ Range) TextCheckingResult {
	rv := objc.Call[TextCheckingResult](r_, objc.Sel("firstMatchInString:options:range:"), string_, options, range_)
	return rv
}

// Returns the number of matches of the regular expression within the specified range of the string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsregularexpression/1414308-numberofmatchesinstring?language=objc
func (r_ RegularExpression) NumberOfMatchesInStringOptionsRange(string_ string, options MatchingOptions, range_ Range) uint {
	rv := objc.Call[uint](r_, objc.Sel("numberOfMatchesInString:options:range:"), string_, options, range_)
	return rv
}

// Returns a new string containing matching regular expressions replaced with the template string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsregularexpression/1413740-stringbyreplacingmatchesinstring?language=objc
func (r_ RegularExpression) StringByReplacingMatchesInStringOptionsRangeWithTemplate(string_ string, options MatchingOptions, range_ Range, templ string) string {
	rv := objc.Call[string](r_, objc.Sel("stringByReplacingMatchesInString:options:range:withTemplate:"), string_, options, range_, templ)
	return rv
}

// Returns a template string by adding backslash escapes as necessary to protect any characters that would match as pattern metacharacters [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsregularexpression/1410085-escapedtemplateforstring?language=objc
func (rc _RegularExpressionClass) EscapedTemplateForString(string_ string) string {
	rv := objc.Call[string](rc, objc.Sel("escapedTemplateForString:"), string_)
	return rv
}

// Returns a template string by adding backslash escapes as necessary to protect any characters that would match as pattern metacharacters [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsregularexpression/1410085-escapedtemplateforstring?language=objc
func RegularExpression_EscapedTemplateForString(string_ string) string {
	return RegularExpressionClass.EscapedTemplateForString(string_)
}

// Returns an array containing all the matches of the regular expression in the string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsregularexpression/1412446-matchesinstring?language=objc
func (r_ RegularExpression) MatchesInStringOptionsRange(string_ string, options MatchingOptions, range_ Range) []TextCheckingResult {
	rv := objc.Call[[]TextCheckingResult](r_, objc.Sel("matchesInString:options:range:"), string_, options, range_)
	return rv
}

// Creates an NSRegularExpression instance with the specified regular expression pattern and options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsregularexpression/1557374-regularexpressionwithpattern?language=objc
func (rc _RegularExpressionClass) RegularExpressionWithPatternOptionsError(pattern string, options RegularExpressionOptions, error IError) RegularExpression {
	rv := objc.Call[RegularExpression](rc, objc.Sel("regularExpressionWithPattern:options:error:"), pattern, options, objc.Ptr(error))
	return rv
}

// Creates an NSRegularExpression instance with the specified regular expression pattern and options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsregularexpression/1557374-regularexpressionwithpattern?language=objc
func RegularExpression_RegularExpressionWithPatternOptionsError(pattern string, options RegularExpressionOptions, error IError) RegularExpression {
	return RegularExpressionClass.RegularExpressionWithPatternOptionsError(pattern, options, error)
}

// Enumerates the string allowing the Block to handle each regular expression match. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsregularexpression/1409687-enumeratematchesinstring?language=objc
func (r_ RegularExpression) EnumerateMatchesInStringOptionsRangeUsingBlock(string_ string, options MatchingOptions, range_ Range, block func(result TextCheckingResult, flags MatchingFlags, stop *bool)) {
	objc.Call[objc.Void](r_, objc.Sel("enumerateMatchesInString:options:range:usingBlock:"), string_, options, range_, block)
}

// Returns the options used when the regular expression option was created. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsregularexpression/1408265-options?language=objc
func (r_ RegularExpression) Options() RegularExpressionOptions {
	rv := objc.Call[RegularExpressionOptions](r_, objc.Sel("options"))
	return rv
}

// Returns the number of capture groups in the regular expression. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsregularexpression/1415052-numberofcapturegroups?language=objc
func (r_ RegularExpression) NumberOfCaptureGroups() uint {
	rv := objc.Call[uint](r_, objc.Sel("numberOfCaptureGroups"))
	return rv
}

// Returns the regular expression pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsregularexpression/1414932-pattern?language=objc
func (r_ RegularExpression) Pattern() string {
	rv := objc.Call[string](r_, objc.Sel("pattern"))
	return rv
}
