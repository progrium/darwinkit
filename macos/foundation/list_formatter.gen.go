// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ListFormatter] class.
var ListFormatterClass = _ListFormatterClass{objc.GetClass("NSListFormatter")}

type _ListFormatterClass struct {
	objc.Class
}

// An interface definition for the [ListFormatter] class.
type IListFormatter interface {
	IFormatter
	StringFromItems(items []objc.IObject) string
	Locale() Locale
	SetLocale(value ILocale)
	ItemFormatter() Formatter
	SetItemFormatter(value IFormatter)
}

// An object that provides locale-correct formatting of a list of items using the appropriate separator and conjunction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslistformatter?language=objc
type ListFormatter struct {
	Formatter
}

func ListFormatterFrom(ptr unsafe.Pointer) ListFormatter {
	return ListFormatter{
		Formatter: FormatterFrom(ptr),
	}
}

func (lc _ListFormatterClass) Alloc() ListFormatter {
	rv := objc.Call[ListFormatter](lc, objc.Sel("alloc"))
	return rv
}

func ListFormatter_Alloc() ListFormatter {
	return ListFormatterClass.Alloc()
}

func (lc _ListFormatterClass) New() ListFormatter {
	rv := objc.Call[ListFormatter](lc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewListFormatter() ListFormatter {
	return ListFormatterClass.New()
}

func (l_ ListFormatter) Init() ListFormatter {
	rv := objc.Call[ListFormatter](l_, objc.Sel("init"))
	return rv
}

// Creates a formatted string for an array of items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslistformatter/3130992-stringfromitems?language=objc
func (l_ ListFormatter) StringFromItems(items []objc.IObject) string {
	rv := objc.Call[string](l_, objc.Sel("stringFromItems:"), items)
	return rv
}

// Constructs a formatted string from an array of strings that uses the list format specific to the current locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslistformatter/3130990-localizedstringbyjoiningstrings?language=objc
func (lc _ListFormatterClass) LocalizedStringByJoiningStrings(strings []string) string {
	rv := objc.Call[string](lc, objc.Sel("localizedStringByJoiningStrings:"), strings)
	return rv
}

// Constructs a formatted string from an array of strings that uses the list format specific to the current locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslistformatter/3130990-localizedstringbyjoiningstrings?language=objc
func ListFormatter_LocalizedStringByJoiningStrings(strings []string) string {
	return ListFormatterClass.LocalizedStringByJoiningStrings(strings)
}

// The locale to use when formatting items in the list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslistformatter/3130989-locale?language=objc
func (l_ ListFormatter) Locale() Locale {
	rv := objc.Call[Locale](l_, objc.Sel("locale"))
	return rv
}

// The locale to use when formatting items in the list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslistformatter/3130989-locale?language=objc
func (l_ ListFormatter) SetLocale(value ILocale) {
	objc.Call[objc.Void](l_, objc.Sel("setLocale:"), objc.Ptr(value))
}

// An object that formats each item in the list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslistformatter/3130988-itemformatter?language=objc
func (l_ ListFormatter) ItemFormatter() Formatter {
	rv := objc.Call[Formatter](l_, objc.Sel("itemFormatter"))
	return rv
}

// An object that formats each item in the list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslistformatter/3130988-itemformatter?language=objc
func (l_ ListFormatter) SetItemFormatter(value IFormatter) {
	objc.Call[objc.Void](l_, objc.Sel("setItemFormatter:"), objc.Ptr(value))
}
