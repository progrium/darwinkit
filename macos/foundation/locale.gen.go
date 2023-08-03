// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var LocaleClass = _LocaleClass{objc.GetClass("NSLocale")}

type _LocaleClass struct {
	objc.Class
}

type ILocale interface {
	objc.IObject
	LocalizedStringForLocaleIdentifier(localeIdentifier string) string
	LocalizedStringForCountryCode(countryCode string) string
	LocalizedStringForLanguageCode(languageCode string) string
	LocalizedStringForScriptCode(scriptCode string) string
	LocalizedStringForVariantCode(variantCode string) string
	LocalizedStringForCollationIdentifier(collationIdentifier string) string
	LocalizedStringForCollatorIdentifier(collatorIdentifier string) string
	LocalizedStringForCurrencyCode(currencyCode string) string
	LocalizedStringForCalendarIdentifier(calendarIdentifier string) string
	ObjectForKey(key LocaleKey) objc.Object
	DisplayNameForKeyValue(key LocaleKey, value objc.IObject) string
	LocaleIdentifier() string
	CountryCode() string
	LanguageCode() string
	ScriptCode() string
	VariantCode() string
	ExemplarCharacterSet() CharacterSet
	CollationIdentifier() string
	CollatorIdentifier() string
	UsesMetricSystem() bool
	DecimalSeparator() string
	GroupingSeparator() string
	CurrencyCode() string
	CurrencySymbol() string
	CalendarIdentifier() string
	QuotationBeginDelimiter() string
	QuotationEndDelimiter() string
	AlternateQuotationBeginDelimiter() string
	AlternateQuotationEndDelimiter() string
}

type Locale struct {
	objc.Object
}

func MakeLocale(ptr unsafe.Pointer) Locale {
	return Locale{
		Object: objc.MakeObject(ptr),
	}
}

func (lc _LocaleClass) LocaleWithLocaleIdentifier(ident string) Locale {
	rv := objc.CallMethod[Locale](lc, objc.GetSelector("localeWithLocaleIdentifier:"), ident)
	return rv
}

func Locale_LocaleWithLocaleIdentifier(ident string) Locale {
	return LocaleClass.LocaleWithLocaleIdentifier(ident)
}

func (l_ Locale) InitWithLocaleIdentifier(string_ string) Locale {
	rv := objc.CallMethod[Locale](l_, objc.GetSelector("initWithLocaleIdentifier:"), string_)
	return rv
}

func Locale_InitWithLocaleIdentifier(string_ string) Locale {
	return LocaleClass.Alloc().InitWithLocaleIdentifier(string_)
}

func (lc _LocaleClass) Alloc() Locale {
	rv := objc.CallMethod[Locale](lc, objc.GetSelector("alloc"))
	return rv
}

func Locale_Alloc() Locale {
	return LocaleClass.Alloc()
}

func (lc _LocaleClass) New() Locale {
	rv := objc.CallMethod[Locale](lc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewLocale() Locale {
	return LocaleClass.New()
}

func Locale_New() Locale {
	return LocaleClass.New()
}

func (l_ Locale) Init() Locale {
	rv := objc.CallMethod[Locale](l_, objc.GetSelector("init"))
	return rv
}

func Locale_Init() Locale {
	return LocaleClass.Alloc().Init()
}

func (lc _LocaleClass) CanonicalLocaleIdentifierFromString(string_ string) string {
	rv := objc.CallMethod[string](lc, objc.GetSelector("canonicalLocaleIdentifierFromString:"), string_)
	return rv
}

func Locale_CanonicalLocaleIdentifierFromString(string_ string) string {
	return LocaleClass.CanonicalLocaleIdentifierFromString(string_)
}

func (lc _LocaleClass) ComponentsFromLocaleIdentifier(string_ string) map[string]string {
	rv := objc.CallMethod[map[string]string](lc, objc.GetSelector("componentsFromLocaleIdentifier:"), string_)
	return rv
}

func Locale_ComponentsFromLocaleIdentifier(string_ string) map[string]string {
	return LocaleClass.ComponentsFromLocaleIdentifier(string_)
}

func (lc _LocaleClass) LocaleIdentifierFromComponents(dict map[string]string) string {
	rv := objc.CallMethod[string](lc, objc.GetSelector("localeIdentifierFromComponents:"), dict)
	return rv
}

func Locale_LocaleIdentifierFromComponents(dict map[string]string) string {
	return LocaleClass.LocaleIdentifierFromComponents(dict)
}

func (lc _LocaleClass) CanonicalLanguageIdentifierFromString(string_ string) string {
	rv := objc.CallMethod[string](lc, objc.GetSelector("canonicalLanguageIdentifierFromString:"), string_)
	return rv
}

func Locale_CanonicalLanguageIdentifierFromString(string_ string) string {
	return LocaleClass.CanonicalLanguageIdentifierFromString(string_)
}

func (lc _LocaleClass) LocaleIdentifierFromWindowsLocaleCode(lcid uint32) string {
	rv := objc.CallMethod[string](lc, objc.GetSelector("localeIdentifierFromWindowsLocaleCode:"), lcid)
	return rv
}

func Locale_LocaleIdentifierFromWindowsLocaleCode(lcid uint32) string {
	return LocaleClass.LocaleIdentifierFromWindowsLocaleCode(lcid)
}

func (lc _LocaleClass) WindowsLocaleCodeFromLocaleIdentifier(localeIdentifier string) uint32 {
	rv := objc.CallMethod[uint32](lc, objc.GetSelector("windowsLocaleCodeFromLocaleIdentifier:"), localeIdentifier)
	return rv
}

func Locale_WindowsLocaleCodeFromLocaleIdentifier(localeIdentifier string) uint32 {
	return LocaleClass.WindowsLocaleCodeFromLocaleIdentifier(localeIdentifier)
}

func (l_ Locale) LocalizedStringForLocaleIdentifier(localeIdentifier string) string {
	rv := objc.CallMethod[string](l_, objc.GetSelector("localizedStringForLocaleIdentifier:"), localeIdentifier)
	return rv
}

func (l_ Locale) LocalizedStringForCountryCode(countryCode string) string {
	rv := objc.CallMethod[string](l_, objc.GetSelector("localizedStringForCountryCode:"), countryCode)
	return rv
}

func (l_ Locale) LocalizedStringForLanguageCode(languageCode string) string {
	rv := objc.CallMethod[string](l_, objc.GetSelector("localizedStringForLanguageCode:"), languageCode)
	return rv
}

func (l_ Locale) LocalizedStringForScriptCode(scriptCode string) string {
	rv := objc.CallMethod[string](l_, objc.GetSelector("localizedStringForScriptCode:"), scriptCode)
	return rv
}

func (l_ Locale) LocalizedStringForVariantCode(variantCode string) string {
	rv := objc.CallMethod[string](l_, objc.GetSelector("localizedStringForVariantCode:"), variantCode)
	return rv
}

func (l_ Locale) LocalizedStringForCollationIdentifier(collationIdentifier string) string {
	rv := objc.CallMethod[string](l_, objc.GetSelector("localizedStringForCollationIdentifier:"), collationIdentifier)
	return rv
}

func (l_ Locale) LocalizedStringForCollatorIdentifier(collatorIdentifier string) string {
	rv := objc.CallMethod[string](l_, objc.GetSelector("localizedStringForCollatorIdentifier:"), collatorIdentifier)
	return rv
}

func (l_ Locale) LocalizedStringForCurrencyCode(currencyCode string) string {
	rv := objc.CallMethod[string](l_, objc.GetSelector("localizedStringForCurrencyCode:"), currencyCode)
	return rv
}

func (l_ Locale) LocalizedStringForCalendarIdentifier(calendarIdentifier string) string {
	rv := objc.CallMethod[string](l_, objc.GetSelector("localizedStringForCalendarIdentifier:"), calendarIdentifier)
	return rv
}

func (l_ Locale) ObjectForKey(key LocaleKey) objc.Object {
	rv := objc.CallMethod[objc.Object](l_, objc.GetSelector("objectForKey:"), key)
	return rv
}

func (l_ Locale) DisplayNameForKeyValue(key LocaleKey, value objc.IObject) string {
	rv := objc.CallMethod[string](l_, objc.GetSelector("displayNameForKey:value:"), key, objc.ExtractPtr(value))
	return rv
}

func (lc _LocaleClass) CharacterDirectionForLanguage(isoLangCode string) LocaleLanguageDirection {
	rv := objc.CallMethod[LocaleLanguageDirection](lc, objc.GetSelector("characterDirectionForLanguage:"), isoLangCode)
	return rv
}

func Locale_CharacterDirectionForLanguage(isoLangCode string) LocaleLanguageDirection {
	return LocaleClass.CharacterDirectionForLanguage(isoLangCode)
}

func (lc _LocaleClass) LineDirectionForLanguage(isoLangCode string) LocaleLanguageDirection {
	rv := objc.CallMethod[LocaleLanguageDirection](lc, objc.GetSelector("lineDirectionForLanguage:"), isoLangCode)
	return rv
}

func Locale_LineDirectionForLanguage(isoLangCode string) LocaleLanguageDirection {
	return LocaleClass.LineDirectionForLanguage(isoLangCode)
}

func (lc _LocaleClass) AutoupdatingCurrentLocale() Locale {
	rv := objc.CallMethod[Locale](lc, objc.GetSelector("autoupdatingCurrentLocale"))
	return rv
}

func Locale_AutoupdatingCurrentLocale() Locale {
	return LocaleClass.AutoupdatingCurrentLocale()
}

func (lc _LocaleClass) CurrentLocale() Locale {
	rv := objc.CallMethod[Locale](lc, objc.GetSelector("currentLocale"))
	return rv
}

func Locale_CurrentLocale() Locale {
	return LocaleClass.CurrentLocale()
}

func (lc _LocaleClass) SystemLocale() Locale {
	rv := objc.CallMethod[Locale](lc, objc.GetSelector("systemLocale"))
	return rv
}

func Locale_SystemLocale() Locale {
	return LocaleClass.SystemLocale()
}

func (lc _LocaleClass) AvailableLocaleIdentifiers() []string {
	rv := objc.CallMethod[[]string](lc, objc.GetSelector("availableLocaleIdentifiers"))
	return rv
}

func Locale_AvailableLocaleIdentifiers() []string {
	return LocaleClass.AvailableLocaleIdentifiers()
}

func (lc _LocaleClass) ISOCountryCodes() []string {
	rv := objc.CallMethod[[]string](lc, objc.GetSelector("ISOCountryCodes"))
	return rv
}

func Locale_ISOCountryCodes() []string {
	return LocaleClass.ISOCountryCodes()
}

func (lc _LocaleClass) ISOLanguageCodes() []string {
	rv := objc.CallMethod[[]string](lc, objc.GetSelector("ISOLanguageCodes"))
	return rv
}

func Locale_ISOLanguageCodes() []string {
	return LocaleClass.ISOLanguageCodes()
}

func (lc _LocaleClass) ISOCurrencyCodes() []string {
	rv := objc.CallMethod[[]string](lc, objc.GetSelector("ISOCurrencyCodes"))
	return rv
}

func Locale_ISOCurrencyCodes() []string {
	return LocaleClass.ISOCurrencyCodes()
}

func (lc _LocaleClass) CommonISOCurrencyCodes() []string {
	rv := objc.CallMethod[[]string](lc, objc.GetSelector("commonISOCurrencyCodes"))
	return rv
}

func Locale_CommonISOCurrencyCodes() []string {
	return LocaleClass.CommonISOCurrencyCodes()
}

func (l_ Locale) LocaleIdentifier() string {
	rv := objc.CallMethod[string](l_, objc.GetSelector("localeIdentifier"))
	return rv
}

func (l_ Locale) CountryCode() string {
	rv := objc.CallMethod[string](l_, objc.GetSelector("countryCode"))
	return rv
}

func (l_ Locale) LanguageCode() string {
	rv := objc.CallMethod[string](l_, objc.GetSelector("languageCode"))
	return rv
}

func (l_ Locale) ScriptCode() string {
	rv := objc.CallMethod[string](l_, objc.GetSelector("scriptCode"))
	return rv
}

func (l_ Locale) VariantCode() string {
	rv := objc.CallMethod[string](l_, objc.GetSelector("variantCode"))
	return rv
}

func (l_ Locale) ExemplarCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](l_, objc.GetSelector("exemplarCharacterSet"))
	return rv
}

func (l_ Locale) CollationIdentifier() string {
	rv := objc.CallMethod[string](l_, objc.GetSelector("collationIdentifier"))
	return rv
}

func (l_ Locale) CollatorIdentifier() string {
	rv := objc.CallMethod[string](l_, objc.GetSelector("collatorIdentifier"))
	return rv
}

func (l_ Locale) UsesMetricSystem() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("usesMetricSystem"))
	return rv
}

func (l_ Locale) DecimalSeparator() string {
	rv := objc.CallMethod[string](l_, objc.GetSelector("decimalSeparator"))
	return rv
}

func (l_ Locale) GroupingSeparator() string {
	rv := objc.CallMethod[string](l_, objc.GetSelector("groupingSeparator"))
	return rv
}

func (l_ Locale) CurrencyCode() string {
	rv := objc.CallMethod[string](l_, objc.GetSelector("currencyCode"))
	return rv
}

func (l_ Locale) CurrencySymbol() string {
	rv := objc.CallMethod[string](l_, objc.GetSelector("currencySymbol"))
	return rv
}

func (l_ Locale) CalendarIdentifier() string {
	rv := objc.CallMethod[string](l_, objc.GetSelector("calendarIdentifier"))
	return rv
}

func (l_ Locale) QuotationBeginDelimiter() string {
	rv := objc.CallMethod[string](l_, objc.GetSelector("quotationBeginDelimiter"))
	return rv
}

func (l_ Locale) QuotationEndDelimiter() string {
	rv := objc.CallMethod[string](l_, objc.GetSelector("quotationEndDelimiter"))
	return rv
}

func (l_ Locale) AlternateQuotationBeginDelimiter() string {
	rv := objc.CallMethod[string](l_, objc.GetSelector("alternateQuotationBeginDelimiter"))
	return rv
}

func (l_ Locale) AlternateQuotationEndDelimiter() string {
	rv := objc.CallMethod[string](l_, objc.GetSelector("alternateQuotationEndDelimiter"))
	return rv
}

func (lc _LocaleClass) PreferredLanguages() []string {
	rv := objc.CallMethod[[]string](lc, objc.GetSelector("preferredLanguages"))
	return rv
}

func Locale_PreferredLanguages() []string {
	return LocaleClass.PreferredLanguages()
}
