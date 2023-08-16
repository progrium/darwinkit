// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TextCheckingResult] class.
var TextCheckingResultClass = _TextCheckingResultClass{objc.GetClass("NSTextCheckingResult")}

type _TextCheckingResultClass struct {
	objc.Class
}

// An interface definition for the [TextCheckingResult] class.
type ITextCheckingResult interface {
	objc.IObject
	RangeWithName(name string) Range
	ResultByAdjustingRangesWithOffset(offset int) TextCheckingResult
	RangeAtIndex(idx uint) Range
	PhoneNumber() string
	RegularExpression() RegularExpression
	Date() Date
	GrammarDetails() []map[string]objc.Object
	NumberOfRanges() uint
	Range() Range
	AddressComponents() map[TextCheckingKey]string
	Orthography() Orthography
	TimeZone() TimeZone
	URL() URL
	AlternativeStrings() []string
	Components() map[TextCheckingKey]string
	ReplacementString() string
	Duration() TimeInterval
	ResultType() TextCheckingType
}

// An occurrence of textual content found during the analysis of a block of text, such as when matching a regular expression. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult?language=objc
type TextCheckingResult struct {
	objc.Object
}

func TextCheckingResultFrom(ptr unsafe.Pointer) TextCheckingResult {
	return TextCheckingResult{
		Object: objc.ObjectFrom(ptr),
	}
}

func (tc _TextCheckingResultClass) Alloc() TextCheckingResult {
	rv := objc.Call[TextCheckingResult](tc, objc.Sel("alloc"))
	return rv
}

func TextCheckingResult_Alloc() TextCheckingResult {
	return TextCheckingResultClass.Alloc()
}

func (tc _TextCheckingResultClass) New() TextCheckingResult {
	rv := objc.Call[TextCheckingResult](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTextCheckingResult() TextCheckingResult {
	return TextCheckingResultClass.New()
}

func (t_ TextCheckingResult) Init() TextCheckingResult {
	rv := objc.Call[TextCheckingResult](t_, objc.Sel("init"))
	return rv
}

// Creates and returns a text checking result with the specified orthography. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1415506-orthographycheckingresultwithran?language=objc
func (tc _TextCheckingResultClass) OrthographyCheckingResultWithRangeOrthography(range_ Range, orthography IOrthography) TextCheckingResult {
	rv := objc.Call[TextCheckingResult](tc, objc.Sel("orthographyCheckingResultWithRange:orthography:"), range_, objc.Ptr(orthography))
	return rv
}

// Creates and returns a text checking result with the specified orthography. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1415506-orthographycheckingresultwithran?language=objc
func TextCheckingResult_OrthographyCheckingResultWithRangeOrthography(range_ Range, orthography IOrthography) TextCheckingResult {
	return TextCheckingResultClass.OrthographyCheckingResultWithRangeOrthography(range_, orthography)
}

// Creates and returns a text checking result with the specified transit information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1408575-transitinformationcheckingresult?language=objc
func (tc _TextCheckingResultClass) TransitInformationCheckingResultWithRangeComponents(range_ Range, components map[TextCheckingKey]string) TextCheckingResult {
	rv := objc.Call[TextCheckingResult](tc, objc.Sel("transitInformationCheckingResultWithRange:components:"), range_, components)
	return rv
}

// Creates and returns a text checking result with the specified transit information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1408575-transitinformationcheckingresult?language=objc
func TextCheckingResult_TransitInformationCheckingResultWithRangeComponents(range_ Range, components map[TextCheckingKey]string) TextCheckingResult {
	return TextCheckingResultClass.TransitInformationCheckingResultWithRangeComponents(range_, components)
}

// Creates and returns a text checking result with the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1413056-linkcheckingresultwithrange?language=objc
func (tc _TextCheckingResultClass) LinkCheckingResultWithRangeURL(range_ Range, url IURL) TextCheckingResult {
	rv := objc.Call[TextCheckingResult](tc, objc.Sel("linkCheckingResultWithRange:URL:"), range_, objc.Ptr(url))
	return rv
}

// Creates and returns a text checking result with the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1413056-linkcheckingresultwithrange?language=objc
func TextCheckingResult_LinkCheckingResultWithRangeURL(range_ Range, url IURL) TextCheckingResult {
	return TextCheckingResultClass.LinkCheckingResultWithRangeURL(range_, url)
}

// Creates and returns a text checking result with the specified array of grammatical errors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1407190-grammarcheckingresultwithrange?language=objc
func (tc _TextCheckingResultClass) GrammarCheckingResultWithRangeDetails(range_ Range, details []map[string]objc.IObject) TextCheckingResult {
	rv := objc.Call[TextCheckingResult](tc, objc.Sel("grammarCheckingResultWithRange:details:"), range_, details)
	return rv
}

// Creates and returns a text checking result with the specified array of grammatical errors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1407190-grammarcheckingresultwithrange?language=objc
func TextCheckingResult_GrammarCheckingResultWithRangeDetails(range_ Range, details []map[string]objc.IObject) TextCheckingResult {
	return TextCheckingResultClass.GrammarCheckingResultWithRangeDetails(range_, details)
}

// Creates and returns a text checking result with the specified quote-balanced replacement string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1413783-quotecheckingresultwithrange?language=objc
func (tc _TextCheckingResultClass) QuoteCheckingResultWithRangeReplacementString(range_ Range, replacementString string) TextCheckingResult {
	rv := objc.Call[TextCheckingResult](tc, objc.Sel("quoteCheckingResultWithRange:replacementString:"), range_, replacementString)
	return rv
}

// Creates and returns a text checking result with the specified quote-balanced replacement string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1413783-quotecheckingresultwithrange?language=objc
func TextCheckingResult_QuoteCheckingResultWithRangeReplacementString(range_ Range, replacementString string) TextCheckingResult {
	return TextCheckingResultClass.QuoteCheckingResultWithRangeReplacementString(range_, replacementString)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/2915200-rangewithname?language=objc
func (t_ TextCheckingResult) RangeWithName(name string) Range {
	rv := objc.Call[Range](t_, objc.Sel("rangeWithName:"), name)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1416640-correctioncheckingresultwithrang?language=objc
func (tc _TextCheckingResultClass) CorrectionCheckingResultWithRangeReplacementStringAlternativeStrings(range_ Range, replacementString string, alternativeStrings []string) TextCheckingResult {
	rv := objc.Call[TextCheckingResult](tc, objc.Sel("correctionCheckingResultWithRange:replacementString:alternativeStrings:"), range_, replacementString, alternativeStrings)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1416640-correctioncheckingresultwithrang?language=objc
func TextCheckingResult_CorrectionCheckingResultWithRangeReplacementStringAlternativeStrings(range_ Range, replacementString string, alternativeStrings []string) TextCheckingResult {
	return TextCheckingResultClass.CorrectionCheckingResultWithRangeReplacementStringAlternativeStrings(range_, replacementString, alternativeStrings)
}

// Returns a new text checking result after adjusting the ranges as specified by the offset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1417597-resultbyadjustingrangeswithoffse?language=objc
func (t_ TextCheckingResult) ResultByAdjustingRangesWithOffset(offset int) TextCheckingResult {
	rv := objc.Call[TextCheckingResult](t_, objc.Sel("resultByAdjustingRangesWithOffset:"), offset)
	return rv
}

// Returns the result type that the range represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1416732-rangeatindex?language=objc
func (t_ TextCheckingResult) RangeAtIndex(idx uint) Range {
	rv := objc.Call[Range](t_, objc.Sel("rangeAtIndex:"), idx)
	return rv
}

// Creates and returns a text checking result with the specified dash corrected replacement string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1409525-dashcheckingresultwithrange?language=objc
func (tc _TextCheckingResultClass) DashCheckingResultWithRangeReplacementString(range_ Range, replacementString string) TextCheckingResult {
	rv := objc.Call[TextCheckingResult](tc, objc.Sel("dashCheckingResultWithRange:replacementString:"), range_, replacementString)
	return rv
}

// Creates and returns a text checking result with the specified dash corrected replacement string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1409525-dashcheckingresultwithrange?language=objc
func TextCheckingResult_DashCheckingResultWithRangeReplacementString(range_ Range, replacementString string) TextCheckingResult {
	return TextCheckingResultClass.DashCheckingResultWithRangeReplacementString(range_, replacementString)
}

// Creates and returns a text checking result with the range of a misspelled word. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1416255-spellcheckingresultwithrange?language=objc
func (tc _TextCheckingResultClass) SpellCheckingResultWithRange(range_ Range) TextCheckingResult {
	rv := objc.Call[TextCheckingResult](tc, objc.Sel("spellCheckingResultWithRange:"), range_)
	return rv
}

// Creates and returns a text checking result with the range of a misspelled word. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1416255-spellcheckingresultwithrange?language=objc
func TextCheckingResult_SpellCheckingResultWithRange(range_ Range) TextCheckingResult {
	return TextCheckingResultClass.SpellCheckingResultWithRange(range_)
}

// Creates and returns a text checking result with the specified date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1410401-datecheckingresultwithrange?language=objc
func (tc _TextCheckingResultClass) DateCheckingResultWithRangeDate(range_ Range, date IDate) TextCheckingResult {
	rv := objc.Call[TextCheckingResult](tc, objc.Sel("dateCheckingResultWithRange:date:"), range_, objc.Ptr(date))
	return rv
}

// Creates and returns a text checking result with the specified date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1410401-datecheckingresultwithrange?language=objc
func TextCheckingResult_DateCheckingResultWithRangeDate(range_ Range, date IDate) TextCheckingResult {
	return TextCheckingResultClass.DateCheckingResultWithRangeDate(range_, date)
}

// Creates and returns a text checking result with the specified phone number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1410389-phonenumbercheckingresultwithran?language=objc
func (tc _TextCheckingResultClass) PhoneNumberCheckingResultWithRangePhoneNumber(range_ Range, phoneNumber string) TextCheckingResult {
	rv := objc.Call[TextCheckingResult](tc, objc.Sel("phoneNumberCheckingResultWithRange:phoneNumber:"), range_, phoneNumber)
	return rv
}

// Creates and returns a text checking result with the specified phone number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1410389-phonenumbercheckingresultwithran?language=objc
func TextCheckingResult_PhoneNumberCheckingResultWithRangePhoneNumber(range_ Range, phoneNumber string) TextCheckingResult {
	return TextCheckingResultClass.PhoneNumberCheckingResultWithRangePhoneNumber(range_, phoneNumber)
}

// Creates and returns a type checking result with the specified regular expression data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1412104-regularexpressioncheckingresultw?language=objc
func (tc _TextCheckingResultClass) RegularExpressionCheckingResultWithRangesCountRegularExpression(ranges RangePointer, count uint, regularExpression IRegularExpression) TextCheckingResult {
	rv := objc.Call[TextCheckingResult](tc, objc.Sel("regularExpressionCheckingResultWithRanges:count:regularExpression:"), ranges, count, objc.Ptr(regularExpression))
	return rv
}

// Creates and returns a type checking result with the specified regular expression data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1412104-regularexpressioncheckingresultw?language=objc
func TextCheckingResult_RegularExpressionCheckingResultWithRangesCountRegularExpression(ranges RangePointer, count uint, regularExpression IRegularExpression) TextCheckingResult {
	return TextCheckingResultClass.RegularExpressionCheckingResultWithRangesCountRegularExpression(ranges, count, regularExpression)
}

// Creates and returns a text checking result with the specified address components. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1413828-addresscheckingresultwithrange?language=objc
func (tc _TextCheckingResultClass) AddressCheckingResultWithRangeComponents(range_ Range, components map[TextCheckingKey]string) TextCheckingResult {
	rv := objc.Call[TextCheckingResult](tc, objc.Sel("addressCheckingResultWithRange:components:"), range_, components)
	return rv
}

// Creates and returns a text checking result with the specified address components. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1413828-addresscheckingresultwithrange?language=objc
func TextCheckingResult_AddressCheckingResultWithRangeComponents(range_ Range, components map[TextCheckingKey]string) TextCheckingResult {
	return TextCheckingResultClass.AddressCheckingResultWithRangeComponents(range_, components)
}

// Creates and returns a text checking result with the specified replacement string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1416651-replacementcheckingresultwithran?language=objc
func (tc _TextCheckingResultClass) ReplacementCheckingResultWithRangeReplacementString(range_ Range, replacementString string) TextCheckingResult {
	rv := objc.Call[TextCheckingResult](tc, objc.Sel("replacementCheckingResultWithRange:replacementString:"), range_, replacementString)
	return rv
}

// Creates and returns a text checking result with the specified replacement string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1416651-replacementcheckingresultwithran?language=objc
func TextCheckingResult_ReplacementCheckingResultWithRangeReplacementString(range_ Range, replacementString string) TextCheckingResult {
	return TextCheckingResultClass.ReplacementCheckingResultWithRangeReplacementString(range_, replacementString)
}

// The phone number of a type checking result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1415511-phonenumber?language=objc
func (t_ TextCheckingResult) PhoneNumber() string {
	rv := objc.Call[string](t_, objc.Sel("phoneNumber"))
	return rv
}

// The regular expression of a type checking result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1417393-regularexpression?language=objc
func (t_ TextCheckingResult) RegularExpression() RegularExpression {
	rv := objc.Call[RegularExpression](t_, objc.Sel("regularExpression"))
	return rv
}

// The date component of a type checking result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1414289-date?language=objc
func (t_ TextCheckingResult) Date() Date {
	rv := objc.Call[Date](t_, objc.Sel("date"))
	return rv
}

// The details of a located grammatical type checking result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1408959-grammardetails?language=objc
func (t_ TextCheckingResult) GrammarDetails() []map[string]objc.Object {
	rv := objc.Call[[]map[string]objc.Object](t_, objc.Sel("grammarDetails"))
	return rv
}

// Returns the number of ranges. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1407475-numberofranges?language=objc
func (t_ TextCheckingResult) NumberOfRanges() uint {
	rv := objc.Call[uint](t_, objc.Sel("numberOfRanges"))
	return rv
}

// Returns the range of the result that the receiver represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1415852-range?language=objc
func (t_ TextCheckingResult) Range() Range {
	rv := objc.Call[Range](t_, objc.Sel("range"))
	return rv
}

// The address dictionary of a type checking result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1413728-addresscomponents?language=objc
func (t_ TextCheckingResult) AddressComponents() map[TextCheckingKey]string {
	rv := objc.Call[map[TextCheckingKey]string](t_, objc.Sel("addressComponents"))
	return rv
}

// The detected orthography of a type checking result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1414551-orthography?language=objc
func (t_ TextCheckingResult) Orthography() Orthography {
	rv := objc.Call[Orthography](t_, objc.Sel("orthography"))
	return rv
}

// The time zone component of a type checking result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1418476-timezone?language=objc
func (t_ TextCheckingResult) TimeZone() TimeZone {
	rv := objc.Call[TimeZone](t_, objc.Sel("timeZone"))
	return rv
}

// The URL of a type checking result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1417843-url?language=objc
func (t_ TextCheckingResult) URL() URL {
	rv := objc.Call[URL](t_, objc.Sel("URL"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1415454-alternativestrings?language=objc
func (t_ TextCheckingResult) AlternativeStrings() []string {
	rv := objc.Call[[]string](t_, objc.Sel("alternativeStrings"))
	return rv
}

// A dictionary containing the components of a type checking result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1407367-components?language=objc
func (t_ TextCheckingResult) Components() map[TextCheckingKey]string {
	rv := objc.Call[map[TextCheckingKey]string](t_, objc.Sel("components"))
	return rv
}

// A replacement string from one of a number of replacement checking results. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1412681-replacementstring?language=objc
func (t_ TextCheckingResult) ReplacementString() string {
	rv := objc.Call[string](t_, objc.Sel("replacementString"))
	return rv
}

// The duration component of a type checking result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1415044-duration?language=objc
func (t_ TextCheckingResult) Duration() TimeInterval {
	rv := objc.Call[TimeInterval](t_, objc.Sel("duration"))
	return rv
}

// Returns the text checking result type that the receiver represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/1407779-resulttype?language=objc
func (t_ TextCheckingResult) ResultType() TextCheckingType {
	rv := objc.Call[TextCheckingType](t_, objc.Sel("resultType"))
	return rv
}
