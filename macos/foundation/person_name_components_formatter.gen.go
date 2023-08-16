// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PersonNameComponentsFormatter] class.
var PersonNameComponentsFormatterClass = _PersonNameComponentsFormatterClass{objc.GetClass("NSPersonNameComponentsFormatter")}

type _PersonNameComponentsFormatterClass struct {
	objc.Class
}

// An interface definition for the [PersonNameComponentsFormatter] class.
type IPersonNameComponentsFormatter interface {
	IFormatter
	AnnotatedStringFromPersonNameComponents(components IPersonNameComponents) AttributedString
	StringFromPersonNameComponents(components IPersonNameComponents) string
	PersonNameComponentsFromString(string_ string) PersonNameComponents
	Style() PersonNameComponentsFormatterStyle
	SetStyle(value PersonNameComponentsFormatterStyle)
	Locale() Locale
	SetLocale(value ILocale)
	IsPhonetic() bool
	SetPhonetic(value bool)
}

// A formatter that provides localized representations of the components of a person’s name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspersonnamecomponentsformatter?language=objc
type PersonNameComponentsFormatter struct {
	Formatter
}

func PersonNameComponentsFormatterFrom(ptr unsafe.Pointer) PersonNameComponentsFormatter {
	return PersonNameComponentsFormatter{
		Formatter: FormatterFrom(ptr),
	}
}

func (pc _PersonNameComponentsFormatterClass) Alloc() PersonNameComponentsFormatter {
	rv := objc.Call[PersonNameComponentsFormatter](pc, objc.Sel("alloc"))
	return rv
}

func PersonNameComponentsFormatter_Alloc() PersonNameComponentsFormatter {
	return PersonNameComponentsFormatterClass.Alloc()
}

func (pc _PersonNameComponentsFormatterClass) New() PersonNameComponentsFormatter {
	rv := objc.Call[PersonNameComponentsFormatter](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPersonNameComponentsFormatter() PersonNameComponentsFormatter {
	return PersonNameComponentsFormatterClass.New()
}

func (p_ PersonNameComponentsFormatter) Init() PersonNameComponentsFormatter {
	rv := objc.Call[PersonNameComponentsFormatter](p_, objc.Sel("init"))
	return rv
}

// Returns an attributed string formatted for a given NSPersonNameComponents object, with attribute annotations for each component. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspersonnamecomponentsformatter/1408241-annotatedstringfrompersonnamecom?language=objc
func (p_ PersonNameComponentsFormatter) AnnotatedStringFromPersonNameComponents(components IPersonNameComponents) AttributedString {
	rv := objc.Call[AttributedString](p_, objc.Sel("annotatedStringFromPersonNameComponents:"), objc.Ptr(components))
	return rv
}

// Returns a string formatted for a given NSPersonNameComponents object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspersonnamecomponentsformatter/1408243-stringfrompersonnamecomponents?language=objc
func (p_ PersonNameComponentsFormatter) StringFromPersonNameComponents(components IPersonNameComponents) string {
	rv := objc.Call[string](p_, objc.Sel("stringFromPersonNameComponents:"), objc.Ptr(components))
	return rv
}

// Returns a string formatted for a given NSPersonNameComponents object using the provided style and options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspersonnamecomponentsformatter/1408258-localizedstringfrompersonnamecom?language=objc
func (pc _PersonNameComponentsFormatterClass) LocalizedStringFromPersonNameComponentsStyleOptions(components IPersonNameComponents, nameFormatStyle PersonNameComponentsFormatterStyle, nameOptions PersonNameComponentsFormatterOptions) string {
	rv := objc.Call[string](pc, objc.Sel("localizedStringFromPersonNameComponents:style:options:"), objc.Ptr(components), nameFormatStyle, nameOptions)
	return rv
}

// Returns a string formatted for a given NSPersonNameComponents object using the provided style and options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspersonnamecomponentsformatter/1408258-localizedstringfrompersonnamecom?language=objc
func PersonNameComponentsFormatter_LocalizedStringFromPersonNameComponentsStyleOptions(components IPersonNameComponents, nameFormatStyle PersonNameComponentsFormatterStyle, nameOptions PersonNameComponentsFormatterOptions) string {
	return PersonNameComponentsFormatterClass.LocalizedStringFromPersonNameComponentsStyleOptions(components, nameFormatStyle, nameOptions)
}

// Returns a person name components object from a given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspersonnamecomponentsformatter/1642979-personnamecomponentsfromstring?language=objc
func (p_ PersonNameComponentsFormatter) PersonNameComponentsFromString(string_ string) PersonNameComponents {
	rv := objc.Call[PersonNameComponents](p_, objc.Sel("personNameComponentsFromString:"), string_)
	return rv
}

// The formatting style of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspersonnamecomponentsformatter/1408260-style?language=objc
func (p_ PersonNameComponentsFormatter) Style() PersonNameComponentsFormatterStyle {
	rv := objc.Call[PersonNameComponentsFormatterStyle](p_, objc.Sel("style"))
	return rv
}

// The formatting style of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspersonnamecomponentsformatter/1408260-style?language=objc
func (p_ PersonNameComponentsFormatter) SetStyle(value PersonNameComponentsFormatterStyle) {
	objc.Call[objc.Void](p_, objc.Sel("setStyle:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspersonnamecomponentsformatter/3850507-locale?language=objc
func (p_ PersonNameComponentsFormatter) Locale() Locale {
	rv := objc.Call[Locale](p_, objc.Sel("locale"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspersonnamecomponentsformatter/3850507-locale?language=objc
func (p_ PersonNameComponentsFormatter) SetLocale(value ILocale) {
	objc.Call[objc.Void](p_, objc.Sel("setLocale:"), objc.Ptr(value))
}

// A Boolean value that specifies whether the receiver should use only the phonetic representations of name components. NO by default. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspersonnamecomponentsformatter/1408242-phonetic?language=objc
func (p_ PersonNameComponentsFormatter) IsPhonetic() bool {
	rv := objc.Call[bool](p_, objc.Sel("isPhonetic"))
	return rv
}

// A Boolean value that specifies whether the receiver should use only the phonetic representations of name components. NO by default. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspersonnamecomponentsformatter/1408242-phonetic?language=objc
func (p_ PersonNameComponentsFormatter) SetPhonetic(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setPhonetic:"), value)
}
