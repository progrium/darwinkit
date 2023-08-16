// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NumberFormatter] class.
var NumberFormatterClass = _NumberFormatterClass{objc.GetClass("NSNumberFormatter")}

type _NumberFormatterClass struct {
	objc.Class
}

// An interface definition for the [NumberFormatter] class.
type INumberFormatter interface {
	IFormatter
	GetObjectValueForStringRangeError(obj objc.IObject, string_ string, rangep *Range, error IError) bool
	StringFromNumber(number INumber) string
	NumberFromString(string_ string) Number
	IsLenient() bool
	SetLenient(value bool)
	AlwaysShowsDecimalSeparator() bool
	SetAlwaysShowsDecimalSeparator(value bool)
	LocalizesFormat() bool
	SetLocalizesFormat(value bool)
	PlusSign() string
	SetPlusSign(value string)
	CurrencyGroupingSeparator() string
	SetCurrencyGroupingSeparator(value string)
	UsesSignificantDigits() bool
	SetUsesSignificantDigits(value bool)
	NegativePrefix() string
	SetNegativePrefix(value string)
	AllowsFloats() bool
	SetAllowsFloats(value bool)
	PositiveFormat() string
	SetPositiveFormat(value string)
	PositivePrefix() string
	SetPositivePrefix(value string)
	SecondaryGroupingSize() uint
	SetSecondaryGroupingSize(value uint)
	CurrencyCode() string
	SetCurrencyCode(value string)
	NegativeFormat() string
	SetNegativeFormat(value string)
	FormatWidth() uint
	SetFormatWidth(value uint)
	PaddingCharacter() string
	SetPaddingCharacter(value string)
	PositiveInfinitySymbol() string
	SetPositiveInfinitySymbol(value string)
	Multiplier() Number
	SetMultiplier(value INumber)
	ExponentSymbol() string
	SetExponentSymbol(value string)
	FormatterBehavior() NumberFormatterBehavior
	SetFormatterBehavior(value NumberFormatterBehavior)
	NumberStyle() NumberFormatterStyle
	SetNumberStyle(value NumberFormatterStyle)
	AttributedStringForNotANumber() AttributedString
	SetAttributedStringForNotANumber(value IAttributedString)
	InternationalCurrencySymbol() string
	SetInternationalCurrencySymbol(value string)
	MaximumIntegerDigits() uint
	SetMaximumIntegerDigits(value uint)
	TextAttributesForPositiveInfinity() map[string]objc.Object
	SetTextAttributesForPositiveInfinity(value map[string]objc.IObject)
	RoundingIncrement() Number
	SetRoundingIncrement(value INumber)
	TextAttributesForNegativeInfinity() map[string]objc.Object
	SetTextAttributesForNegativeInfinity(value map[string]objc.IObject)
	NegativeInfinitySymbol() string
	SetNegativeInfinitySymbol(value string)
	GeneratesDecimalNumbers() bool
	SetGeneratesDecimalNumbers(value bool)
	TextAttributesForZero() map[string]objc.Object
	SetTextAttributesForZero(value map[string]objc.IObject)
	MinimumSignificantDigits() uint
	SetMinimumSignificantDigits(value uint)
	Locale() Locale
	SetLocale(value ILocale)
	HasThousandSeparators() bool
	SetHasThousandSeparators(value bool)
	CurrencyDecimalSeparator() string
	SetCurrencyDecimalSeparator(value string)
	MaximumFractionDigits() uint
	SetMaximumFractionDigits(value uint)
	Minimum() Number
	SetMinimum(value INumber)
	AttributedStringForZero() AttributedString
	SetAttributedStringForZero(value IAttributedString)
	PerMillSymbol() string
	SetPerMillSymbol(value string)
	MaximumSignificantDigits() uint
	SetMaximumSignificantDigits(value uint)
	ZeroSymbol() string
	SetZeroSymbol(value string)
	PercentSymbol() string
	SetPercentSymbol(value string)
	NilSymbol() string
	SetNilSymbol(value string)
	MinusSign() string
	SetMinusSign(value string)
	NotANumberSymbol() string
	SetNotANumberSymbol(value string)
	CurrencySymbol() string
	SetCurrencySymbol(value string)
	FormattingContext() FormattingContext
	SetFormattingContext(value FormattingContext)
	DecimalSeparator() string
	SetDecimalSeparator(value string)
	UsesGroupingSeparator() bool
	SetUsesGroupingSeparator(value bool)
	IsPartialStringValidationEnabled() bool
	SetPartialStringValidationEnabled(value bool)
	AttributedStringForNil() AttributedString
	SetAttributedStringForNil(value IAttributedString)
	ThousandSeparator() string
	SetThousandSeparator(value string)
	TextAttributesForNotANumber() map[string]objc.Object
	SetTextAttributesForNotANumber(value map[string]objc.IObject)
	PositiveSuffix() string
	SetPositiveSuffix(value string)
	TextAttributesForPositiveValues() map[string]objc.Object
	SetTextAttributesForPositiveValues(value map[string]objc.IObject)
	RoundingBehavior() DecimalNumberHandler
	SetRoundingBehavior(value IDecimalNumberHandler)
	GroupingSeparator() string
	SetGroupingSeparator(value string)
	MinimumFractionDigits() uint
	SetMinimumFractionDigits(value uint)
	Format() string
	SetFormat(value string)
	TextAttributesForNil() map[string]objc.Object
	SetTextAttributesForNil(value map[string]objc.IObject)
	MinimumIntegerDigits() uint
	SetMinimumIntegerDigits(value uint)
	RoundingMode() NumberFormatterRoundingMode
	SetRoundingMode(value NumberFormatterRoundingMode)
	Maximum() Number
	SetMaximum(value INumber)
	PaddingPosition() NumberFormatterPadPosition
	SetPaddingPosition(value NumberFormatterPadPosition)
	NegativeSuffix() string
	SetNegativeSuffix(value string)
	TextAttributesForNegativeValues() map[string]objc.Object
	SetTextAttributesForNegativeValues(value map[string]objc.IObject)
	GroupingSize() uint
	SetGroupingSize(value uint)
}

// A formatter that converts between numeric values and their textual representations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter?language=objc
type NumberFormatter struct {
	Formatter
}

func NumberFormatterFrom(ptr unsafe.Pointer) NumberFormatter {
	return NumberFormatter{
		Formatter: FormatterFrom(ptr),
	}
}

func (nc _NumberFormatterClass) Alloc() NumberFormatter {
	rv := objc.Call[NumberFormatter](nc, objc.Sel("alloc"))
	return rv
}

func NumberFormatter_Alloc() NumberFormatter {
	return NumberFormatterClass.Alloc()
}

func (nc _NumberFormatterClass) New() NumberFormatter {
	rv := objc.Call[NumberFormatter](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNumberFormatter() NumberFormatter {
	return NumberFormatterClass.New()
}

func (n_ NumberFormatter) Init() NumberFormatter {
	rv := objc.Call[NumberFormatter](n_, objc.Sel("init"))
	return rv
}

// Sets the default formatter behavior for new instances of NSNumberFormatter . [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1415422-setdefaultformatterbehavior?language=objc
func (nc _NumberFormatterClass) SetDefaultFormatterBehavior(behavior NumberFormatterBehavior) {
	objc.Call[objc.Void](nc, objc.Sel("setDefaultFormatterBehavior:"), behavior)
}

// Sets the default formatter behavior for new instances of NSNumberFormatter . [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1415422-setdefaultformatterbehavior?language=objc
func NumberFormatter_SetDefaultFormatterBehavior(behavior NumberFormatterBehavior) {
	NumberFormatterClass.SetDefaultFormatterBehavior(behavior)
}

// Returns a localized number string with the specified style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1416418-localizedstringfromnumber?language=objc
func (nc _NumberFormatterClass) LocalizedStringFromNumberNumberStyle(num INumber, nstyle NumberFormatterStyle) string {
	rv := objc.Call[string](nc, objc.Sel("localizedStringFromNumber:numberStyle:"), objc.Ptr(num), nstyle)
	return rv
}

// Returns a localized number string with the specified style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1416418-localizedstringfromnumber?language=objc
func NumberFormatter_LocalizedStringFromNumberNumberStyle(num INumber, nstyle NumberFormatterStyle) string {
	return NumberFormatterClass.LocalizedStringFromNumberNumberStyle(num, nstyle)
}

// Returns by reference a cell-content object after creating it from a range of characters in a given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1412588-getobjectvalue?language=objc
func (n_ NumberFormatter) GetObjectValueForStringRangeError(obj objc.IObject, string_ string, rangep *Range, error IError) bool {
	rv := objc.Call[bool](n_, objc.Sel("getObjectValue:forString:range:error:"), obj, string_, rangep, objc.Ptr(error))
	return rv
}

// Returns a string containing the formatted value of the provided number object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1418046-stringfromnumber?language=objc
func (n_ NumberFormatter) StringFromNumber(number INumber) string {
	rv := objc.Call[string](n_, objc.Sel("stringFromNumber:"), objc.Ptr(number))
	return rv
}

// Returns an NSNumber object created by parsing a given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1408845-numberfromstring?language=objc
func (n_ NumberFormatter) NumberFromString(string_ string) Number {
	rv := objc.Call[Number](n_, objc.Sel("numberFromString:"), string_)
	return rv
}

// Returns an NSNumberFormatterBehavior constant that indicates default formatter behavior for new instances of NSNumberFormatter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1415137-defaultformatterbehavior?language=objc
func (nc _NumberFormatterClass) DefaultFormatterBehavior() NumberFormatterBehavior {
	rv := objc.Call[NumberFormatterBehavior](nc, objc.Sel("defaultFormatterBehavior"))
	return rv
}

// Returns an NSNumberFormatterBehavior constant that indicates default formatter behavior for new instances of NSNumberFormatter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1415137-defaultformatterbehavior?language=objc
func NumberFormatter_DefaultFormatterBehavior() NumberFormatterBehavior {
	return NumberFormatterClass.DefaultFormatterBehavior()
}

// Determines whether the receiver will use heuristics to guess at the number which is intended by a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1416953-lenient?language=objc
func (n_ NumberFormatter) IsLenient() bool {
	rv := objc.Call[bool](n_, objc.Sel("isLenient"))
	return rv
}

// Determines whether the receiver will use heuristics to guess at the number which is intended by a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1416953-lenient?language=objc
func (n_ NumberFormatter) SetLenient(value bool) {
	objc.Call[objc.Void](n_, objc.Sel("setLenient:"), value)
}

// Determines whether the receiver always shows the decimal separator, even for integer numbers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1408550-alwaysshowsdecimalseparator?language=objc
func (n_ NumberFormatter) AlwaysShowsDecimalSeparator() bool {
	rv := objc.Call[bool](n_, objc.Sel("alwaysShowsDecimalSeparator"))
	return rv
}

// Determines whether the receiver always shows the decimal separator, even for integer numbers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1408550-alwaysshowsdecimalseparator?language=objc
func (n_ NumberFormatter) SetAlwaysShowsDecimalSeparator(value bool) {
	objc.Call[objc.Void](n_, objc.Sel("setAlwaysShowsDecimalSeparator:"), value)
}

// Determines whether the dollar sign character ($), decimal separator character (.), and thousand separator character (,) are converted to appropriately localized characters as specified by the user’s localization preference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1408133-localizesformat?language=objc
func (n_ NumberFormatter) LocalizesFormat() bool {
	rv := objc.Call[bool](n_, objc.Sel("localizesFormat"))
	return rv
}

// Determines whether the dollar sign character ($), decimal separator character (.), and thousand separator character (,) are converted to appropriately localized characters as specified by the user’s localization preference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1408133-localizesformat?language=objc
func (n_ NumberFormatter) SetLocalizesFormat(value bool) {
	objc.Call[objc.Void](n_, objc.Sel("setLocalizesFormat:"), value)
}

// The string used to represent a plus sign. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1416423-plussign?language=objc
func (n_ NumberFormatter) PlusSign() string {
	rv := objc.Call[string](n_, objc.Sel("plusSign"))
	return rv
}

// The string used to represent a plus sign. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1416423-plussign?language=objc
func (n_ NumberFormatter) SetPlusSign(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setPlusSign:"), value)
}

// The currency grouping separator for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1416213-currencygroupingseparator?language=objc
func (n_ NumberFormatter) CurrencyGroupingSeparator() string {
	rv := objc.Call[string](n_, objc.Sel("currencyGroupingSeparator"))
	return rv
}

// The currency grouping separator for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1416213-currencygroupingseparator?language=objc
func (n_ NumberFormatter) SetCurrencyGroupingSeparator(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setCurrencyGroupingSeparator:"), value)
}

// A Boolean value indicating whether the formatter uses minimum and maximum significant digits when formatting numbers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1417793-usessignificantdigits?language=objc
func (n_ NumberFormatter) UsesSignificantDigits() bool {
	rv := objc.Call[bool](n_, objc.Sel("usesSignificantDigits"))
	return rv
}

// A Boolean value indicating whether the formatter uses minimum and maximum significant digits when formatting numbers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1417793-usessignificantdigits?language=objc
func (n_ NumberFormatter) SetUsesSignificantDigits(value bool) {
	objc.Call[objc.Void](n_, objc.Sel("setUsesSignificantDigits:"), value)
}

// The string the receiver uses as a prefix for negative values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1408096-negativeprefix?language=objc
func (n_ NumberFormatter) NegativePrefix() string {
	rv := objc.Call[string](n_, objc.Sel("negativePrefix"))
	return rv
}

// The string the receiver uses as a prefix for negative values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1408096-negativeprefix?language=objc
func (n_ NumberFormatter) SetNegativePrefix(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setNegativePrefix:"), value)
}

// Determines whether the receiver allows as input floating-point values (that is, values that include the period character [.]). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1416119-allowsfloats?language=objc
func (n_ NumberFormatter) AllowsFloats() bool {
	rv := objc.Call[bool](n_, objc.Sel("allowsFloats"))
	return rv
}

// Determines whether the receiver allows as input floating-point values (that is, values that include the period character [.]). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1416119-allowsfloats?language=objc
func (n_ NumberFormatter) SetAllowsFloats(value bool) {
	objc.Call[objc.Void](n_, objc.Sel("setAllowsFloats:"), value)
}

// The format the receiver uses to display positive values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1410737-positiveformat?language=objc
func (n_ NumberFormatter) PositiveFormat() string {
	rv := objc.Call[string](n_, objc.Sel("positiveFormat"))
	return rv
}

// The format the receiver uses to display positive values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1410737-positiveformat?language=objc
func (n_ NumberFormatter) SetPositiveFormat(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setPositiveFormat:"), value)
}

// The string the receiver uses as the prefix for positive values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1414204-positiveprefix?language=objc
func (n_ NumberFormatter) PositivePrefix() string {
	rv := objc.Call[string](n_, objc.Sel("positivePrefix"))
	return rv
}

// The string the receiver uses as the prefix for positive values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1414204-positiveprefix?language=objc
func (n_ NumberFormatter) SetPositivePrefix(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setPositivePrefix:"), value)
}

// The secondary grouping size of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1413348-secondarygroupingsize?language=objc
func (n_ NumberFormatter) SecondaryGroupingSize() uint {
	rv := objc.Call[uint](n_, objc.Sel("secondaryGroupingSize"))
	return rv
}

// The secondary grouping size of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1413348-secondarygroupingsize?language=objc
func (n_ NumberFormatter) SetSecondaryGroupingSize(value uint) {
	objc.Call[objc.Void](n_, objc.Sel("setSecondaryGroupingSize:"), value)
}

// The receiver’s currency code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1410463-currencycode?language=objc
func (n_ NumberFormatter) CurrencyCode() string {
	rv := objc.Call[string](n_, objc.Sel("currencyCode"))
	return rv
}

// The receiver’s currency code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1410463-currencycode?language=objc
func (n_ NumberFormatter) SetCurrencyCode(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setCurrencyCode:"), value)
}

// The format the receiver uses to display negative values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1414039-negativeformat?language=objc
func (n_ NumberFormatter) NegativeFormat() string {
	rv := objc.Call[string](n_, objc.Sel("negativeFormat"))
	return rv
}

// The format the receiver uses to display negative values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1414039-negativeformat?language=objc
func (n_ NumberFormatter) SetNegativeFormat(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setNegativeFormat:"), value)
}

// The format width used by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1411919-formatwidth?language=objc
func (n_ NumberFormatter) FormatWidth() uint {
	rv := objc.Call[uint](n_, objc.Sel("formatWidth"))
	return rv
}

// The format width used by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1411919-formatwidth?language=objc
func (n_ NumberFormatter) SetFormatWidth(value uint) {
	objc.Call[objc.Void](n_, objc.Sel("setFormatWidth:"), value)
}

// The string that the receiver uses to pad numbers in the formatted string representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1413690-paddingcharacter?language=objc
func (n_ NumberFormatter) PaddingCharacter() string {
	rv := objc.Call[string](n_, objc.Sel("paddingCharacter"))
	return rv
}

// The string that the receiver uses to pad numbers in the formatted string representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1413690-paddingcharacter?language=objc
func (n_ NumberFormatter) SetPaddingCharacter(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setPaddingCharacter:"), value)
}

// The string used to represent a positive infinity symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1412976-positiveinfinitysymbol?language=objc
func (n_ NumberFormatter) PositiveInfinitySymbol() string {
	rv := objc.Call[string](n_, objc.Sel("positiveInfinitySymbol"))
	return rv
}

// The string used to represent a positive infinity symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1412976-positiveinfinitysymbol?language=objc
func (n_ NumberFormatter) SetPositiveInfinitySymbol(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setPositiveInfinitySymbol:"), value)
}

// The multiplier of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1408601-multiplier?language=objc
func (n_ NumberFormatter) Multiplier() Number {
	rv := objc.Call[Number](n_, objc.Sel("multiplier"))
	return rv
}

// The multiplier of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1408601-multiplier?language=objc
func (n_ NumberFormatter) SetMultiplier(value INumber) {
	objc.Call[objc.Void](n_, objc.Sel("setMultiplier:"), objc.Ptr(value))
}

// The string used to represent an exponent symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1417223-exponentsymbol?language=objc
func (n_ NumberFormatter) ExponentSymbol() string {
	rv := objc.Call[string](n_, objc.Sel("exponentSymbol"))
	return rv
}

// The string used to represent an exponent symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1417223-exponentsymbol?language=objc
func (n_ NumberFormatter) SetExponentSymbol(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setExponentSymbol:"), value)
}

// The formatter behavior of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1417642-formatterbehavior?language=objc
func (n_ NumberFormatter) FormatterBehavior() NumberFormatterBehavior {
	rv := objc.Call[NumberFormatterBehavior](n_, objc.Sel("formatterBehavior"))
	return rv
}

// The formatter behavior of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1417642-formatterbehavior?language=objc
func (n_ NumberFormatter) SetFormatterBehavior(value NumberFormatterBehavior) {
	objc.Call[objc.Void](n_, objc.Sel("setFormatterBehavior:"), value)
}

// The number style used by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1416915-numberstyle?language=objc
func (n_ NumberFormatter) NumberStyle() NumberFormatterStyle {
	rv := objc.Call[NumberFormatterStyle](n_, objc.Sel("numberStyle"))
	return rv
}

// The number style used by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1416915-numberstyle?language=objc
func (n_ NumberFormatter) SetNumberStyle(value NumberFormatterStyle) {
	objc.Call[objc.Void](n_, objc.Sel("setNumberStyle:"), value)
}

// The attributed string the receiver uses to display “not a number” values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1416819-attributedstringfornotanumber?language=objc
func (n_ NumberFormatter) AttributedStringForNotANumber() AttributedString {
	rv := objc.Call[AttributedString](n_, objc.Sel("attributedStringForNotANumber"))
	return rv
}

// The attributed string the receiver uses to display “not a number” values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1416819-attributedstringfornotanumber?language=objc
func (n_ NumberFormatter) SetAttributedStringForNotANumber(value IAttributedString) {
	objc.Call[objc.Void](n_, objc.Sel("setAttributedStringForNotANumber:"), objc.Ptr(value))
}

// The international currency symbol used by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1412755-internationalcurrencysymbol?language=objc
func (n_ NumberFormatter) InternationalCurrencySymbol() string {
	rv := objc.Call[string](n_, objc.Sel("internationalCurrencySymbol"))
	return rv
}

// The international currency symbol used by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1412755-internationalcurrencysymbol?language=objc
func (n_ NumberFormatter) SetInternationalCurrencySymbol(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setInternationalCurrencySymbol:"), value)
}

// The maximum number of digits before the decimal separator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1407284-maximumintegerdigits?language=objc
func (n_ NumberFormatter) MaximumIntegerDigits() uint {
	rv := objc.Call[uint](n_, objc.Sel("maximumIntegerDigits"))
	return rv
}

// The maximum number of digits before the decimal separator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1407284-maximumintegerdigits?language=objc
func (n_ NumberFormatter) SetMaximumIntegerDigits(value uint) {
	objc.Call[objc.Void](n_, objc.Sel("setMaximumIntegerDigits:"), value)
}

// The text attributes used to display the positive infinity symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1408176-textattributesforpositiveinfinit?language=objc
func (n_ NumberFormatter) TextAttributesForPositiveInfinity() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](n_, objc.Sel("textAttributesForPositiveInfinity"))
	return rv
}

// The text attributes used to display the positive infinity symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1408176-textattributesforpositiveinfinit?language=objc
func (n_ NumberFormatter) SetTextAttributesForPositiveInfinity(value map[string]objc.IObject) {
	objc.Call[objc.Void](n_, objc.Sel("setTextAttributesForPositiveInfinity:"), value)
}

// The rounding increment used by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1412561-roundingincrement?language=objc
func (n_ NumberFormatter) RoundingIncrement() Number {
	rv := objc.Call[Number](n_, objc.Sel("roundingIncrement"))
	return rv
}

// The rounding increment used by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1412561-roundingincrement?language=objc
func (n_ NumberFormatter) SetRoundingIncrement(value INumber) {
	objc.Call[objc.Void](n_, objc.Sel("setRoundingIncrement:"), objc.Ptr(value))
}

// The text attributes used to display the negative infinity symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1410417-textattributesfornegativeinfinit?language=objc
func (n_ NumberFormatter) TextAttributesForNegativeInfinity() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](n_, objc.Sel("textAttributesForNegativeInfinity"))
	return rv
}

// The text attributes used to display the negative infinity symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1410417-textattributesfornegativeinfinit?language=objc
func (n_ NumberFormatter) SetTextAttributesForNegativeInfinity(value map[string]objc.IObject) {
	objc.Call[objc.Void](n_, objc.Sel("setTextAttributesForNegativeInfinity:"), value)
}

// The string used to represent a negative infinity symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1417968-negativeinfinitysymbol?language=objc
func (n_ NumberFormatter) NegativeInfinitySymbol() string {
	rv := objc.Call[string](n_, objc.Sel("negativeInfinitySymbol"))
	return rv
}

// The string used to represent a negative infinity symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1417968-negativeinfinitysymbol?language=objc
func (n_ NumberFormatter) SetNegativeInfinitySymbol(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setNegativeInfinitySymbol:"), value)
}

// Determines whether the receiver creates instances of NSDecimalNumber when it converts strings to number objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1410503-generatesdecimalnumbers?language=objc
func (n_ NumberFormatter) GeneratesDecimalNumbers() bool {
	rv := objc.Call[bool](n_, objc.Sel("generatesDecimalNumbers"))
	return rv
}

// Determines whether the receiver creates instances of NSDecimalNumber when it converts strings to number objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1410503-generatesdecimalnumbers?language=objc
func (n_ NumberFormatter) SetGeneratesDecimalNumbers(value bool) {
	objc.Call[objc.Void](n_, objc.Sel("setGeneratesDecimalNumbers:"), value)
}

// The text attributes used to display a zero value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1415971-textattributesforzero?language=objc
func (n_ NumberFormatter) TextAttributesForZero() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](n_, objc.Sel("textAttributesForZero"))
	return rv
}

// The text attributes used to display a zero value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1415971-textattributesforzero?language=objc
func (n_ NumberFormatter) SetTextAttributesForZero(value map[string]objc.IObject) {
	objc.Call[objc.Void](n_, objc.Sel("setTextAttributesForZero:"), value)
}

// The minimum number of significant digits for the number formatter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1410174-minimumsignificantdigits?language=objc
func (n_ NumberFormatter) MinimumSignificantDigits() uint {
	rv := objc.Call[uint](n_, objc.Sel("minimumSignificantDigits"))
	return rv
}

// The minimum number of significant digits for the number formatter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1410174-minimumsignificantdigits?language=objc
func (n_ NumberFormatter) SetMinimumSignificantDigits(value uint) {
	objc.Call[objc.Void](n_, objc.Sel("setMinimumSignificantDigits:"), value)
}

// The locale of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1416967-locale?language=objc
func (n_ NumberFormatter) Locale() Locale {
	rv := objc.Call[Locale](n_, objc.Sel("locale"))
	return rv
}

// The locale of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1416967-locale?language=objc
func (n_ NumberFormatter) SetLocale(value ILocale) {
	objc.Call[objc.Void](n_, objc.Sel("setLocale:"), objc.Ptr(value))
}

// Determines whether the receiver uses thousand separators. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1416451-hasthousandseparators?language=objc
func (n_ NumberFormatter) HasThousandSeparators() bool {
	rv := objc.Call[bool](n_, objc.Sel("hasThousandSeparators"))
	return rv
}

// Determines whether the receiver uses thousand separators. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1416451-hasthousandseparators?language=objc
func (n_ NumberFormatter) SetHasThousandSeparators(value bool) {
	objc.Call[objc.Void](n_, objc.Sel("setHasThousandSeparators:"), value)
}

// The string used by the receiver as a currency decimal separator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1407247-currencydecimalseparator?language=objc
func (n_ NumberFormatter) CurrencyDecimalSeparator() string {
	rv := objc.Call[string](n_, objc.Sel("currencyDecimalSeparator"))
	return rv
}

// The string used by the receiver as a currency decimal separator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1407247-currencydecimalseparator?language=objc
func (n_ NumberFormatter) SetCurrencyDecimalSeparator(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setCurrencyDecimalSeparator:"), value)
}

// The maximum number of digits after the decimal separator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1415364-maximumfractiondigits?language=objc
func (n_ NumberFormatter) MaximumFractionDigits() uint {
	rv := objc.Call[uint](n_, objc.Sel("maximumFractionDigits"))
	return rv
}

// The maximum number of digits after the decimal separator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1415364-maximumfractiondigits?language=objc
func (n_ NumberFormatter) SetMaximumFractionDigits(value uint) {
	objc.Call[objc.Void](n_, objc.Sel("setMaximumFractionDigits:"), value)
}

// The lowest number allowed as input by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1417228-minimum?language=objc
func (n_ NumberFormatter) Minimum() Number {
	rv := objc.Call[Number](n_, objc.Sel("minimum"))
	return rv
}

// The lowest number allowed as input by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1417228-minimum?language=objc
func (n_ NumberFormatter) SetMinimum(value INumber) {
	objc.Call[objc.Void](n_, objc.Sel("setMinimum:"), objc.Ptr(value))
}

// The attributed string that the receiver uses to display zero values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1415516-attributedstringforzero?language=objc
func (n_ NumberFormatter) AttributedStringForZero() AttributedString {
	rv := objc.Call[AttributedString](n_, objc.Sel("attributedStringForZero"))
	return rv
}

// The attributed string that the receiver uses to display zero values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1415516-attributedstringforzero?language=objc
func (n_ NumberFormatter) SetAttributedStringForZero(value IAttributedString) {
	objc.Call[objc.Void](n_, objc.Sel("setAttributedStringForZero:"), objc.Ptr(value))
}

// The string used to represent a per-mill (per-thousand) symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1412399-permillsymbol?language=objc
func (n_ NumberFormatter) PerMillSymbol() string {
	rv := objc.Call[string](n_, objc.Sel("perMillSymbol"))
	return rv
}

// The string used to represent a per-mill (per-thousand) symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1412399-permillsymbol?language=objc
func (n_ NumberFormatter) SetPerMillSymbol(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setPerMillSymbol:"), value)
}

// The maximum number of significant digits for the number formatter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1412008-maximumsignificantdigits?language=objc
func (n_ NumberFormatter) MaximumSignificantDigits() uint {
	rv := objc.Call[uint](n_, objc.Sel("maximumSignificantDigits"))
	return rv
}

// The maximum number of significant digits for the number formatter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1412008-maximumsignificantdigits?language=objc
func (n_ NumberFormatter) SetMaximumSignificantDigits(value uint) {
	objc.Call[objc.Void](n_, objc.Sel("setMaximumSignificantDigits:"), value)
}

// The string used to represent a zero value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1410405-zerosymbol?language=objc
func (n_ NumberFormatter) ZeroSymbol() string {
	rv := objc.Call[string](n_, objc.Sel("zeroSymbol"))
	return rv
}

// The string used to represent a zero value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1410405-zerosymbol?language=objc
func (n_ NumberFormatter) SetZeroSymbol(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setZeroSymbol:"), value)
}

// The string used to represent a percent symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1407789-percentsymbol?language=objc
func (n_ NumberFormatter) PercentSymbol() string {
	rv := objc.Call[string](n_, objc.Sel("percentSymbol"))
	return rv
}

// The string used to represent a percent symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1407789-percentsymbol?language=objc
func (n_ NumberFormatter) SetPercentSymbol(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setPercentSymbol:"), value)
}

// The string used to represent a nil value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1412699-nilsymbol?language=objc
func (n_ NumberFormatter) NilSymbol() string {
	rv := objc.Call[string](n_, objc.Sel("nilSymbol"))
	return rv
}

// The string used to represent a nil value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1412699-nilsymbol?language=objc
func (n_ NumberFormatter) SetNilSymbol(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setNilSymbol:"), value)
}

// The string used to represent a minus sign. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1409416-minussign?language=objc
func (n_ NumberFormatter) MinusSign() string {
	rv := objc.Call[string](n_, objc.Sel("minusSign"))
	return rv
}

// The string used to represent a minus sign. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1409416-minussign?language=objc
func (n_ NumberFormatter) SetMinusSign(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setMinusSign:"), value)
}

// The string used to represent a NaN (“not a number”) value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1416993-notanumbersymbol?language=objc
func (n_ NumberFormatter) NotANumberSymbol() string {
	rv := objc.Call[string](n_, objc.Sel("notANumberSymbol"))
	return rv
}

// The string used to represent a NaN (“not a number”) value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1416993-notanumbersymbol?language=objc
func (n_ NumberFormatter) SetNotANumberSymbol(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setNotANumberSymbol:"), value)
}

// The string used by the receiver as a local currency symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1414668-currencysymbol?language=objc
func (n_ NumberFormatter) CurrencySymbol() string {
	rv := objc.Call[string](n_, objc.Sel("currencySymbol"))
	return rv
}

// The string used by the receiver as a local currency symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1414668-currencysymbol?language=objc
func (n_ NumberFormatter) SetCurrencySymbol(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setCurrencySymbol:"), value)
}

// The capitalization formatting context used when formatting a number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1408543-formattingcontext?language=objc
func (n_ NumberFormatter) FormattingContext() FormattingContext {
	rv := objc.Call[FormattingContext](n_, objc.Sel("formattingContext"))
	return rv
}

// The capitalization formatting context used when formatting a number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1408543-formattingcontext?language=objc
func (n_ NumberFormatter) SetFormattingContext(value FormattingContext) {
	objc.Call[objc.Void](n_, objc.Sel("setFormattingContext:"), value)
}

// The character the receiver uses as a decimal separator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1408029-decimalseparator?language=objc
func (n_ NumberFormatter) DecimalSeparator() string {
	rv := objc.Call[string](n_, objc.Sel("decimalSeparator"))
	return rv
}

// The character the receiver uses as a decimal separator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1408029-decimalseparator?language=objc
func (n_ NumberFormatter) SetDecimalSeparator(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setDecimalSeparator:"), value)
}

// Determines whether the receiver displays the group separator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1409880-usesgroupingseparator?language=objc
func (n_ NumberFormatter) UsesGroupingSeparator() bool {
	rv := objc.Call[bool](n_, objc.Sel("usesGroupingSeparator"))
	return rv
}

// Determines whether the receiver displays the group separator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1409880-usesgroupingseparator?language=objc
func (n_ NumberFormatter) SetUsesGroupingSeparator(value bool) {
	objc.Call[objc.Void](n_, objc.Sel("setUsesGroupingSeparator:"), value)
}

// Determines whether partial string validation is enabled for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1412244-partialstringvalidationenabled?language=objc
func (n_ NumberFormatter) IsPartialStringValidationEnabled() bool {
	rv := objc.Call[bool](n_, objc.Sel("isPartialStringValidationEnabled"))
	return rv
}

// Determines whether partial string validation is enabled for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1412244-partialstringvalidationenabled?language=objc
func (n_ NumberFormatter) SetPartialStringValidationEnabled(value bool) {
	objc.Call[objc.Void](n_, objc.Sel("setPartialStringValidationEnabled:"), value)
}

// The attributed string the receiver uses to display nil values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1416184-attributedstringfornil?language=objc
func (n_ NumberFormatter) AttributedStringForNil() AttributedString {
	rv := objc.Call[AttributedString](n_, objc.Sel("attributedStringForNil"))
	return rv
}

// The attributed string the receiver uses to display nil values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1416184-attributedstringfornil?language=objc
func (n_ NumberFormatter) SetAttributedStringForNil(value IAttributedString) {
	objc.Call[objc.Void](n_, objc.Sel("setAttributedStringForNil:"), objc.Ptr(value))
}

// The character the receiver uses as a thousand separator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1412771-thousandseparator?language=objc
func (n_ NumberFormatter) ThousandSeparator() string {
	rv := objc.Call[string](n_, objc.Sel("thousandSeparator"))
	return rv
}

// The character the receiver uses as a thousand separator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1412771-thousandseparator?language=objc
func (n_ NumberFormatter) SetThousandSeparator(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setThousandSeparator:"), value)
}

// The text attributes used to display the NaN (“not a number”) string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1410959-textattributesfornotanumber?language=objc
func (n_ NumberFormatter) TextAttributesForNotANumber() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](n_, objc.Sel("textAttributesForNotANumber"))
	return rv
}

// The text attributes used to display the NaN (“not a number”) string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1410959-textattributesfornotanumber?language=objc
func (n_ NumberFormatter) SetTextAttributesForNotANumber(value map[string]objc.IObject) {
	objc.Call[objc.Void](n_, objc.Sel("setTextAttributesForNotANumber:"), value)
}

// The string the receiver uses as the suffix for positive values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1415071-positivesuffix?language=objc
func (n_ NumberFormatter) PositiveSuffix() string {
	rv := objc.Call[string](n_, objc.Sel("positiveSuffix"))
	return rv
}

// The string the receiver uses as the suffix for positive values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1415071-positivesuffix?language=objc
func (n_ NumberFormatter) SetPositiveSuffix(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setPositiveSuffix:"), value)
}

// The text attributes to be used in displaying positive values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1409563-textattributesforpositivevalues?language=objc
func (n_ NumberFormatter) TextAttributesForPositiveValues() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](n_, objc.Sel("textAttributesForPositiveValues"))
	return rv
}

// The text attributes to be used in displaying positive values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1409563-textattributesforpositivevalues?language=objc
func (n_ NumberFormatter) SetTextAttributesForPositiveValues(value map[string]objc.IObject) {
	objc.Call[objc.Void](n_, objc.Sel("setTextAttributesForPositiveValues:"), value)
}

// The rounding behavior used by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1415880-roundingbehavior?language=objc
func (n_ NumberFormatter) RoundingBehavior() DecimalNumberHandler {
	rv := objc.Call[DecimalNumberHandler](n_, objc.Sel("roundingBehavior"))
	return rv
}

// The rounding behavior used by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1415880-roundingbehavior?language=objc
func (n_ NumberFormatter) SetRoundingBehavior(value IDecimalNumberHandler) {
	objc.Call[objc.Void](n_, objc.Sel("setRoundingBehavior:"), objc.Ptr(value))
}

// The string used by the receiver for a grouping separator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1412157-groupingseparator?language=objc
func (n_ NumberFormatter) GroupingSeparator() string {
	rv := objc.Call[string](n_, objc.Sel("groupingSeparator"))
	return rv
}

// The string used by the receiver for a grouping separator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1412157-groupingseparator?language=objc
func (n_ NumberFormatter) SetGroupingSeparator(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setGroupingSeparator:"), value)
}

// The minimum number of digits after the decimal separator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1410459-minimumfractiondigits?language=objc
func (n_ NumberFormatter) MinimumFractionDigits() uint {
	rv := objc.Call[uint](n_, objc.Sel("minimumFractionDigits"))
	return rv
}

// The minimum number of digits after the decimal separator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1410459-minimumfractiondigits?language=objc
func (n_ NumberFormatter) SetMinimumFractionDigits(value uint) {
	objc.Call[objc.Void](n_, objc.Sel("setMinimumFractionDigits:"), value)
}

// The receiver’s format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1410565-format?language=objc
func (n_ NumberFormatter) Format() string {
	rv := objc.Call[string](n_, objc.Sel("format"))
	return rv
}

// The receiver’s format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1410565-format?language=objc
func (n_ NumberFormatter) SetFormat(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setFormat:"), value)
}

// The text attributes used to display the nil symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1408943-textattributesfornil?language=objc
func (n_ NumberFormatter) TextAttributesForNil() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](n_, objc.Sel("textAttributesForNil"))
	return rv
}

// The text attributes used to display the nil symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1408943-textattributesfornil?language=objc
func (n_ NumberFormatter) SetTextAttributesForNil(value map[string]objc.IObject) {
	objc.Call[objc.Void](n_, objc.Sel("setTextAttributesForNil:"), value)
}

// The minimum number of digits before the decimal separator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1410052-minimumintegerdigits?language=objc
func (n_ NumberFormatter) MinimumIntegerDigits() uint {
	rv := objc.Call[uint](n_, objc.Sel("minimumIntegerDigits"))
	return rv
}

// The minimum number of digits before the decimal separator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1410052-minimumintegerdigits?language=objc
func (n_ NumberFormatter) SetMinimumIntegerDigits(value uint) {
	objc.Call[objc.Void](n_, objc.Sel("setMinimumIntegerDigits:"), value)
}

// The rounding mode used by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1411156-roundingmode?language=objc
func (n_ NumberFormatter) RoundingMode() NumberFormatterRoundingMode {
	rv := objc.Call[NumberFormatterRoundingMode](n_, objc.Sel("roundingMode"))
	return rv
}

// The rounding mode used by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1411156-roundingmode?language=objc
func (n_ NumberFormatter) SetRoundingMode(value NumberFormatterRoundingMode) {
	objc.Call[objc.Void](n_, objc.Sel("setRoundingMode:"), value)
}

// The highest number allowed as input by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1417787-maximum?language=objc
func (n_ NumberFormatter) Maximum() Number {
	rv := objc.Call[Number](n_, objc.Sel("maximum"))
	return rv
}

// The highest number allowed as input by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1417787-maximum?language=objc
func (n_ NumberFormatter) SetMaximum(value INumber) {
	objc.Call[objc.Void](n_, objc.Sel("setMaximum:"), objc.Ptr(value))
}

// The padding position used by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1411127-paddingposition?language=objc
func (n_ NumberFormatter) PaddingPosition() NumberFormatterPadPosition {
	rv := objc.Call[NumberFormatterPadPosition](n_, objc.Sel("paddingPosition"))
	return rv
}

// The padding position used by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1411127-paddingposition?language=objc
func (n_ NumberFormatter) SetPaddingPosition(value NumberFormatterPadPosition) {
	objc.Call[objc.Void](n_, objc.Sel("setPaddingPosition:"), value)
}

// The string the receiver uses as a suffix for negative values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1413927-negativesuffix?language=objc
func (n_ NumberFormatter) NegativeSuffix() string {
	rv := objc.Call[string](n_, objc.Sel("negativeSuffix"))
	return rv
}

// The string the receiver uses as a suffix for negative values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1413927-negativesuffix?language=objc
func (n_ NumberFormatter) SetNegativeSuffix(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setNegativeSuffix:"), value)
}

// The text attributes to be used in displaying negative values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1414530-textattributesfornegativevalues?language=objc
func (n_ NumberFormatter) TextAttributesForNegativeValues() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](n_, objc.Sel("textAttributesForNegativeValues"))
	return rv
}

// The text attributes to be used in displaying negative values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1414530-textattributesfornegativevalues?language=objc
func (n_ NumberFormatter) SetTextAttributesForNegativeValues(value map[string]objc.IObject) {
	objc.Call[objc.Void](n_, objc.Sel("setTextAttributesForNegativeValues:"), value)
}

// The grouping size of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1416167-groupingsize?language=objc
func (n_ NumberFormatter) GroupingSize() uint {
	rv := objc.Call[uint](n_, objc.Sel("groupingSize"))
	return rv
}

// The grouping size of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumberformatter/1416167-groupingsize?language=objc
func (n_ NumberFormatter) SetGroupingSize(value uint) {
	objc.Call[objc.Void](n_, objc.Sel("setGroupingSize:"), value)
}
