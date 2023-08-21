// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Locale] class.
var LocaleClass = _LocaleClass{objc.GetClass("NSLocale")}

type _LocaleClass struct {
	objc.Class
}

// An interface definition for the [Locale] class.
type ILocale interface {
	objc.IObject
	ObjectForKey(key LocaleKey) objc.Object
	LocalizedStringForCollatorIdentifier(collatorIdentifier string) string
	DisplayNameForKeyValue(key LocaleKey, value objc.IObject) string
	LocalizedStringForScriptCode(scriptCode string) string
	LocalizedStringForLanguageCode(languageCode string) string
	LocalizedStringForLocaleIdentifier(localeIdentifier string) string
	LocalizedStringForCalendarIdentifier(calendarIdentifier string) string
	LocalizedStringForCurrencyCode(currencyCode string) string
	LocalizedStringForCollationIdentifier(collationIdentifier string) string
	LocalizedStringForVariantCode(variantCode string) string
	LocalizedStringForCountryCode(countryCode string) string
	AlternateQuotationEndDelimiter() string
	CurrencyCode() string
	UsesMetricSystem() bool
	CalendarIdentifier() string
	AlternateQuotationBeginDelimiter() string
	QuotationEndDelimiter() string
	CollatorIdentifier() string
	VariantCode() string
	ExemplarCharacterSet() CharacterSet
	CollationIdentifier() string
	CurrencySymbol() string
	ScriptCode() string
	QuotationBeginDelimiter() string
	LocaleIdentifier() string
	DecimalSeparator() string
	LanguageCode() string
	GroupingSeparator() string
}

// Information about linguistic, cultural, and technological conventions for use in formatting data for presentation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale?language=objc
type Locale struct {
	objc.Object
}

func LocaleFrom(ptr unsafe.Pointer) Locale {
	return Locale{
		Object: objc.ObjectFrom(ptr),
	}
}

func (lc _LocaleClass) LocaleWithLocaleIdentifier(ident string) Locale {
	rv := objc.Call[Locale](lc, objc.Sel("localeWithLocaleIdentifier:"), ident)
	return rv
}

// Returns a locale initialized using the given locale identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1488627-localewithlocaleidentifier?language=objc
func Locale_LocaleWithLocaleIdentifier(ident string) Locale {
	return LocaleClass.LocaleWithLocaleIdentifier(ident)
}

func (l_ Locale) InitWithLocaleIdentifier(string_ string) Locale {
	rv := objc.Call[Locale](l_, objc.Sel("initWithLocaleIdentifier:"), string_)
	return rv
}

// Initializes a locale using a given locale identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1414217-initwithlocaleidentifier?language=objc
func NewLocaleWithLocaleIdentifier(string_ string) Locale {
	instance := LocaleClass.Alloc().InitWithLocaleIdentifier(string_)
	instance.Autorelease()
	return instance
}

func (lc _LocaleClass) Alloc() Locale {
	rv := objc.Call[Locale](lc, objc.Sel("alloc"))
	return rv
}

func Locale_Alloc() Locale {
	return LocaleClass.Alloc()
}

func (lc _LocaleClass) New() Locale {
	rv := objc.Call[Locale](lc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewLocale() Locale {
	return LocaleClass.New()
}

func (l_ Locale) Init() Locale {
	rv := objc.Call[Locale](l_, objc.Sel("init"))
	return rv
}

// Returns the value of the component corresponding to the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1418430-objectforkey?language=objc
func (l_ Locale) ObjectForKey(key LocaleKey) objc.Object {
	rv := objc.Call[objc.Object](l_, objc.Sel("objectForKey:"), key)
	return rv
}

// Returns the localized string for the specified collator identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1643004-localizedstringforcollatoridenti?language=objc
func (l_ Locale) LocalizedStringForCollatorIdentifier(collatorIdentifier string) string {
	rv := objc.Call[string](l_, objc.Sel("localizedStringForCollatorIdentifier:"), collatorIdentifier)
	return rv
}

// Returns the display name for the given locale component value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1415931-displaynameforkey?language=objc
func (l_ Locale) DisplayNameForKeyValue(key LocaleKey, value objc.IObject) string {
	rv := objc.Call[string](l_, objc.Sel("displayNameForKey:value:"), key, value)
	return rv
}

// Returns a locale identifier from the components specified in a given dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1412439-localeidentifierfromcomponents?language=objc
func (lc _LocaleClass) LocaleIdentifierFromComponents(dict map[string]string) string {
	rv := objc.Call[string](lc, objc.Sel("localeIdentifierFromComponents:"), dict)
	return rv
}

// Returns a locale identifier from the components specified in a given dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1412439-localeidentifierfromcomponents?language=objc
func Locale_LocaleIdentifierFromComponents(dict map[string]string) string {
	return LocaleClass.LocaleIdentifierFromComponents(dict)
}

// Returns the localized string for the specified script code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1643126-localizedstringforscriptcode?language=objc
func (l_ Locale) LocalizedStringForScriptCode(scriptCode string) string {
	rv := objc.Call[string](l_, objc.Sel("localizedStringForScriptCode:"), scriptCode)
	return rv
}

// Returns the localized string for the specified language code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1643226-localizedstringforlanguagecode?language=objc
func (l_ Locale) LocalizedStringForLanguageCode(languageCode string) string {
	rv := objc.Call[string](l_, objc.Sel("localizedStringForLanguageCode:"), languageCode)
	return rv
}

// Returns the direction of the sequence of lines for the specified ISO language code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1414007-linedirectionforlanguage?language=objc
func (lc _LocaleClass) LineDirectionForLanguage(isoLangCode string) LocaleLanguageDirection {
	rv := objc.Call[LocaleLanguageDirection](lc, objc.Sel("lineDirectionForLanguage:"), isoLangCode)
	return rv
}

// Returns the direction of the sequence of lines for the specified ISO language code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1414007-linedirectionforlanguage?language=objc
func Locale_LineDirectionForLanguage(isoLangCode string) LocaleLanguageDirection {
	return LocaleClass.LineDirectionForLanguage(isoLangCode)
}

// Returns the localized string for the specified locale identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1642864-localizedstringforlocaleidentifi?language=objc
func (l_ Locale) LocalizedStringForLocaleIdentifier(localeIdentifier string) string {
	rv := objc.Call[string](l_, objc.Sel("localizedStringForLocaleIdentifier:"), localeIdentifier)
	return rv
}

// Returns the localized string for the specified calendar identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/2242780-localizedstringforcalendaridenti?language=objc
func (l_ Locale) LocalizedStringForCalendarIdentifier(calendarIdentifier string) string {
	rv := objc.Call[string](l_, objc.Sel("localizedStringForCalendarIdentifier:"), calendarIdentifier)
	return rv
}

// Returns the canonical identifier for a given locale identification string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1411483-canonicallocaleidentifierfromstr?language=objc
func (lc _LocaleClass) CanonicalLocaleIdentifierFromString(string_ string) string {
	rv := objc.Call[string](lc, objc.Sel("canonicalLocaleIdentifierFromString:"), string_)
	return rv
}

// Returns the canonical identifier for a given locale identification string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1411483-canonicallocaleidentifierfromstr?language=objc
func Locale_CanonicalLocaleIdentifierFromString(string_ string) string {
	return LocaleClass.CanonicalLocaleIdentifierFromString(string_)
}

// Returns a canonical language identifier by mapping an arbitrary locale identification string to the canonical identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1413069-canonicallanguageidentifierfroms?language=objc
func (lc _LocaleClass) CanonicalLanguageIdentifierFromString(string_ string) string {
	rv := objc.Call[string](lc, objc.Sel("canonicalLanguageIdentifierFromString:"), string_)
	return rv
}

// Returns a canonical language identifier by mapping an arbitrary locale identification string to the canonical identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1413069-canonicallanguageidentifierfroms?language=objc
func Locale_CanonicalLanguageIdentifierFromString(string_ string) string {
	return LocaleClass.CanonicalLanguageIdentifierFromString(string_)
}

// Returns the localized string for the specified currency code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1643179-localizedstringforcurrencycode?language=objc
func (l_ Locale) LocalizedStringForCurrencyCode(currencyCode string) string {
	rv := objc.Call[string](l_, objc.Sel("localizedStringForCurrencyCode:"), currencyCode)
	return rv
}

// Returns the localized string for the specified collation identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1642875-localizedstringforcollationident?language=objc
func (l_ Locale) LocalizedStringForCollationIdentifier(collationIdentifier string) string {
	rv := objc.Call[string](l_, objc.Sel("localizedStringForCollationIdentifier:"), collationIdentifier)
	return rv
}

// Returns the localized string for the specified variant code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1643264-localizedstringforvariantcode?language=objc
func (l_ Locale) LocalizedStringForVariantCode(variantCode string) string {
	rv := objc.Call[string](l_, objc.Sel("localizedStringForVariantCode:"), variantCode)
	return rv
}

// Returns a dictionary that is the result of parsing a locale ID. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1409220-componentsfromlocaleidentifier?language=objc
func (lc _LocaleClass) ComponentsFromLocaleIdentifier(string_ string) map[string]string {
	rv := objc.Call[map[string]string](lc, objc.Sel("componentsFromLocaleIdentifier:"), string_)
	return rv
}

// Returns a dictionary that is the result of parsing a locale ID. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1409220-componentsfromlocaleidentifier?language=objc
func Locale_ComponentsFromLocaleIdentifier(string_ string) map[string]string {
	return LocaleClass.ComponentsFromLocaleIdentifier(string_)
}

// Returns a Window locale code from the locale identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1412214-windowslocalecodefromlocaleident?language=objc
func (lc _LocaleClass) WindowsLocaleCodeFromLocaleIdentifier(localeIdentifier string) uint32 {
	rv := objc.Call[uint32](lc, objc.Sel("windowsLocaleCodeFromLocaleIdentifier:"), localeIdentifier)
	return rv
}

// Returns a Window locale code from the locale identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1412214-windowslocalecodefromlocaleident?language=objc
func Locale_WindowsLocaleCodeFromLocaleIdentifier(localeIdentifier string) uint32 {
	return LocaleClass.WindowsLocaleCodeFromLocaleIdentifier(localeIdentifier)
}

// Returns the direction of the sequence of characters in a line for the specified ISO language code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1417681-characterdirectionforlanguage?language=objc
func (lc _LocaleClass) CharacterDirectionForLanguage(isoLangCode string) LocaleLanguageDirection {
	rv := objc.Call[LocaleLanguageDirection](lc, objc.Sel("characterDirectionForLanguage:"), isoLangCode)
	return rv
}

// Returns the direction of the sequence of characters in a line for the specified ISO language code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1417681-characterdirectionforlanguage?language=objc
func Locale_CharacterDirectionForLanguage(isoLangCode string) LocaleLanguageDirection {
	return LocaleClass.CharacterDirectionForLanguage(isoLangCode)
}

// Returns a locale identifier from a Windows locale code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1412710-localeidentifierfromwindowslocal?language=objc
func (lc _LocaleClass) LocaleIdentifierFromWindowsLocaleCode(lcid uint32) string {
	rv := objc.Call[string](lc, objc.Sel("localeIdentifierFromWindowsLocaleCode:"), lcid)
	return rv
}

// Returns a locale identifier from a Windows locale code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1412710-localeidentifierfromwindowslocal?language=objc
func Locale_LocaleIdentifierFromWindowsLocaleCode(lcid uint32) string {
	return LocaleClass.LocaleIdentifierFromWindowsLocaleCode(lcid)
}

// Returns the localized string for a country or region code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1642920-localizedstringforcountrycode?language=objc
func (l_ Locale) LocalizedStringForCountryCode(countryCode string) string {
	rv := objc.Call[string](l_, objc.Sel("localizedStringForCountryCode:"), countryCode)
	return rv
}

// The list of locale identifiers available on the system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1410448-availablelocaleidentifiers?language=objc
func (lc _LocaleClass) AvailableLocaleIdentifiers() []string {
	rv := objc.Call[[]string](lc, objc.Sel("availableLocaleIdentifiers"))
	return rv
}

// The list of locale identifiers available on the system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1410448-availablelocaleidentifiers?language=objc
func Locale_AvailableLocaleIdentifiers() []string {
	return LocaleClass.AvailableLocaleIdentifiers()
}

// The alternate end quotation symbol for the locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1642885-alternatequotationenddelimiter?language=objc
func (l_ Locale) AlternateQuotationEndDelimiter() string {
	rv := objc.Call[string](l_, objc.Sel("alternateQuotationEndDelimiter"))
	return rv
}

// The currency code for the locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1642836-currencycode?language=objc
func (l_ Locale) CurrencyCode() string {
	rv := objc.Call[string](l_, objc.Sel("currencyCode"))
	return rv
}

// An ordered list of the user's preferred languages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1415614-preferredlanguages?language=objc
func (lc _LocaleClass) PreferredLanguages() []string {
	rv := objc.Call[[]string](lc, objc.Sel("preferredLanguages"))
	return rv
}

// An ordered list of the user's preferred languages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1415614-preferredlanguages?language=objc
func Locale_PreferredLanguages() []string {
	return LocaleClass.PreferredLanguages()
}

// A Boolean value that indicates whether the locale uses the metric system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1643225-usesmetricsystem?language=objc
func (l_ Locale) UsesMetricSystem() bool {
	rv := objc.Call[bool](l_, objc.Sel("usesMetricSystem"))
	return rv
}

// The calendar identifier for the locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/2242779-calendaridentifier?language=objc
func (l_ Locale) CalendarIdentifier() string {
	rv := objc.Call[string](l_, objc.Sel("calendarIdentifier"))
	return rv
}

// A locale representing the generic root values with little localization. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1407691-systemlocale?language=objc
func (lc _LocaleClass) SystemLocale() Locale {
	rv := objc.Call[Locale](lc, objc.Sel("systemLocale"))
	return rv
}

// A locale representing the generic root values with little localization. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1407691-systemlocale?language=objc
func Locale_SystemLocale() Locale {
	return LocaleClass.SystemLocale()
}

// The alternate begin quotation symbol for the locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1643238-alternatequotationbegindelimiter?language=objc
func (l_ Locale) AlternateQuotationBeginDelimiter() string {
	rv := objc.Call[string](l_, objc.Sel("alternateQuotationBeginDelimiter"))
	return rv
}

// A locale which tracks the user’s current preferences. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1414388-autoupdatingcurrentlocale?language=objc
func (lc _LocaleClass) AutoupdatingCurrentLocale() Locale {
	rv := objc.Call[Locale](lc, objc.Sel("autoupdatingCurrentLocale"))
	return rv
}

// A locale which tracks the user’s current preferences. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1414388-autoupdatingcurrentlocale?language=objc
func Locale_AutoupdatingCurrentLocale() Locale {
	return LocaleClass.AutoupdatingCurrentLocale()
}

// The end quotation symbol for the locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1643162-quotationenddelimiter?language=objc
func (l_ Locale) QuotationEndDelimiter() string {
	rv := objc.Call[string](l_, objc.Sel("quotationEndDelimiter"))
	return rv
}

// The collator identifier for the locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1643195-collatoridentifier?language=objc
func (l_ Locale) CollatorIdentifier() string {
	rv := objc.Call[string](l_, objc.Sel("collatorIdentifier"))
	return rv
}

// The variant code for the locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1643152-variantcode?language=objc
func (l_ Locale) VariantCode() string {
	rv := objc.Call[string](l_, objc.Sel("variantCode"))
	return rv
}

// The list of known language codes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1418015-isolanguagecodes?language=objc
func (lc _LocaleClass) ISOLanguageCodes() []string {
	rv := objc.Call[[]string](lc, objc.Sel("ISOLanguageCodes"))
	return rv
}

// The list of known language codes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1418015-isolanguagecodes?language=objc
func Locale_ISOLanguageCodes() []string {
	return LocaleClass.ISOLanguageCodes()
}

// A locale representing the user's region settings at the time the property is read. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1409990-currentlocale?language=objc
func (lc _LocaleClass) CurrentLocale() Locale {
	rv := objc.Call[Locale](lc, objc.Sel("currentLocale"))
	return rv
}

// A locale representing the user's region settings at the time the property is read. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1409990-currentlocale?language=objc
func Locale_CurrentLocale() Locale {
	return LocaleClass.CurrentLocale()
}

// The exemplar character set for the locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1643019-exemplarcharacterset?language=objc
func (l_ Locale) ExemplarCharacterSet() CharacterSet {
	rv := objc.Call[CharacterSet](l_, objc.Sel("exemplarCharacterSet"))
	return rv
}

// The collation identifier for the locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1643092-collationidentifier?language=objc
func (l_ Locale) CollationIdentifier() string {
	rv := objc.Call[string](l_, objc.Sel("collationIdentifier"))
	return rv
}

// The list of known country or region codes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1413869-isocountrycodes?language=objc
func (lc _LocaleClass) ISOCountryCodes() []string {
	rv := objc.Call[[]string](lc, objc.Sel("ISOCountryCodes"))
	return rv
}

// The list of known country or region codes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1413869-isocountrycodes?language=objc
func Locale_ISOCountryCodes() []string {
	return LocaleClass.ISOCountryCodes()
}

// The currency symbol for the locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1642814-currencysymbol?language=objc
func (l_ Locale) CurrencySymbol() string {
	rv := objc.Call[string](l_, objc.Sel("currencySymbol"))
	return rv
}

// The script code for the locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1643213-scriptcode?language=objc
func (l_ Locale) ScriptCode() string {
	rv := objc.Call[string](l_, objc.Sel("scriptCode"))
	return rv
}

// The begin quotation symbol for the locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1643155-quotationbegindelimiter?language=objc
func (l_ Locale) QuotationBeginDelimiter() string {
	rv := objc.Call[string](l_, objc.Sel("quotationBeginDelimiter"))
	return rv
}

// The identifier for the locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1416263-localeidentifier?language=objc
func (l_ Locale) LocaleIdentifier() string {
	rv := objc.Call[string](l_, objc.Sel("localeIdentifier"))
	return rv
}

// The decimal separator for the locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1643064-decimalseparator?language=objc
func (l_ Locale) DecimalSeparator() string {
	rv := objc.Call[string](l_, objc.Sel("decimalSeparator"))
	return rv
}

// The list of known currency codes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1417834-isocurrencycodes?language=objc
func (lc _LocaleClass) ISOCurrencyCodes() []string {
	rv := objc.Call[[]string](lc, objc.Sel("ISOCurrencyCodes"))
	return rv
}

// The list of known currency codes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1417834-isocurrencycodes?language=objc
func Locale_ISOCurrencyCodes() []string {
	return LocaleClass.ISOCurrencyCodes()
}

// The language code for the locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1643026-languagecode?language=objc
func (l_ Locale) LanguageCode() string {
	rv := objc.Call[string](l_, objc.Sel("languageCode"))
	return rv
}

// The grouping separator for the locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1643096-groupingseparator?language=objc
func (l_ Locale) GroupingSeparator() string {
	rv := objc.Call[string](l_, objc.Sel("groupingSeparator"))
	return rv
}

// A list of commonly encountered currency codes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1407272-commonisocurrencycodes?language=objc
func (lc _LocaleClass) CommonISOCurrencyCodes() []string {
	rv := objc.Call[[]string](lc, objc.Sel("commonISOCurrencyCodes"))
	return rv
}

// A list of commonly encountered currency codes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/1407272-commonisocurrencycodes?language=objc
func Locale_CommonISOCurrencyCodes() []string {
	return LocaleClass.CommonISOCurrencyCodes()
}
