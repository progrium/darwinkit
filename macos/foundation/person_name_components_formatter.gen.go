// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var PersonNameComponentsFormatterClass = _PersonNameComponentsFormatterClass{objc.GetClass("NSPersonNameComponentsFormatter")}

type _PersonNameComponentsFormatterClass struct {
	objc.Class
}

type IPersonNameComponentsFormatter interface {
	IFormatter
	StringFromPersonNameComponents(components IPersonNameComponents) string
	AnnotatedStringFromPersonNameComponents(components IPersonNameComponents) AttributedString
	PersonNameComponentsFromString(string_ string) PersonNameComponents
	Style() PersonNameComponentsFormatterStyle
	SetStyle(value PersonNameComponentsFormatterStyle)
	IsPhonetic() bool
	SetPhonetic(value bool)
	Locale() Locale
	SetLocale(value ILocale)
}

type PersonNameComponentsFormatter struct {
	Formatter
}

func MakePersonNameComponentsFormatter(ptr unsafe.Pointer) PersonNameComponentsFormatter {
	return PersonNameComponentsFormatter{
		Formatter: MakeFormatter(ptr),
	}
}

func (pc _PersonNameComponentsFormatterClass) Alloc() PersonNameComponentsFormatter {
	rv := objc.CallMethod[PersonNameComponentsFormatter](pc, objc.GetSelector("alloc"))
	return rv
}

func PersonNameComponentsFormatter_Alloc() PersonNameComponentsFormatter {
	return PersonNameComponentsFormatterClass.Alloc()
}

func (pc _PersonNameComponentsFormatterClass) New() PersonNameComponentsFormatter {
	rv := objc.CallMethod[PersonNameComponentsFormatter](pc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewPersonNameComponentsFormatter() PersonNameComponentsFormatter {
	return PersonNameComponentsFormatterClass.New()
}

func PersonNameComponentsFormatter_New() PersonNameComponentsFormatter {
	return PersonNameComponentsFormatterClass.New()
}

func (p_ PersonNameComponentsFormatter) Init() PersonNameComponentsFormatter {
	rv := objc.CallMethod[PersonNameComponentsFormatter](p_, objc.GetSelector("init"))
	return rv
}

func PersonNameComponentsFormatter_Init() PersonNameComponentsFormatter {
	return PersonNameComponentsFormatterClass.Alloc().Init()
}

func (pc _PersonNameComponentsFormatterClass) LocalizedStringFromPersonNameComponentsStyleOptions(components IPersonNameComponents, nameFormatStyle PersonNameComponentsFormatterStyle, nameOptions PersonNameComponentsFormatterOptions) string {
	rv := objc.CallMethod[string](pc, objc.GetSelector("localizedStringFromPersonNameComponents:style:options:"), objc.ExtractPtr(components), nameFormatStyle, nameOptions)
	return rv
}

func PersonNameComponentsFormatter_LocalizedStringFromPersonNameComponentsStyleOptions(components IPersonNameComponents, nameFormatStyle PersonNameComponentsFormatterStyle, nameOptions PersonNameComponentsFormatterOptions) string {
	return PersonNameComponentsFormatterClass.LocalizedStringFromPersonNameComponentsStyleOptions(components, nameFormatStyle, nameOptions)
}

func (p_ PersonNameComponentsFormatter) StringFromPersonNameComponents(components IPersonNameComponents) string {
	rv := objc.CallMethod[string](p_, objc.GetSelector("stringFromPersonNameComponents:"), objc.ExtractPtr(components))
	return rv
}

func (p_ PersonNameComponentsFormatter) AnnotatedStringFromPersonNameComponents(components IPersonNameComponents) AttributedString {
	rv := objc.CallMethod[AttributedString](p_, objc.GetSelector("annotatedStringFromPersonNameComponents:"), objc.ExtractPtr(components))
	return rv
}

func (p_ PersonNameComponentsFormatter) PersonNameComponentsFromString(string_ string) PersonNameComponents {
	rv := objc.CallMethod[PersonNameComponents](p_, objc.GetSelector("personNameComponentsFromString:"), string_)
	return rv
}

func (p_ PersonNameComponentsFormatter) Style() PersonNameComponentsFormatterStyle {
	rv := objc.CallMethod[PersonNameComponentsFormatterStyle](p_, objc.GetSelector("style"))
	return rv
}

func (p_ PersonNameComponentsFormatter) SetStyle(value PersonNameComponentsFormatterStyle) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setStyle:"), value)
}

func (p_ PersonNameComponentsFormatter) IsPhonetic() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("isPhonetic"))
	return rv
}

func (p_ PersonNameComponentsFormatter) SetPhonetic(value bool) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setPhonetic:"), value)
}

func (p_ PersonNameComponentsFormatter) Locale() Locale {
	rv := objc.CallMethod[Locale](p_, objc.GetSelector("locale"))
	return rv
}

func (p_ PersonNameComponentsFormatter) SetLocale(value ILocale) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setLocale:"), objc.ExtractPtr(value))
}
