// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ISO8601DateFormatter] class.
var ISO8601DateFormatterClass = _ISO8601DateFormatterClass{objc.GetClass("NSISO8601DateFormatter")}

type _ISO8601DateFormatterClass struct {
	objc.Class
}

// An interface definition for the [ISO8601DateFormatter] class.
type IISO8601DateFormatter interface {
	IFormatter
	DateFromString(string_ string) Date
	StringFromDate(date IDate) string
	TimeZone() TimeZone
	SetTimeZone(value ITimeZone)
	FormatOptions() ISO8601DateFormatOptions
	SetFormatOptions(value ISO8601DateFormatOptions)
}

// A formatter that converts between dates and their ISO 8601 string representations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsiso8601dateformatter?language=objc
type ISO8601DateFormatter struct {
	Formatter
}

func ISO8601DateFormatterFrom(ptr unsafe.Pointer) ISO8601DateFormatter {
	return ISO8601DateFormatter{
		Formatter: FormatterFrom(ptr),
	}
}

func (i_ ISO8601DateFormatter) Init() ISO8601DateFormatter {
	rv := objc.Call[ISO8601DateFormatter](i_, objc.Sel("init"))
	return rv
}

func (ic _ISO8601DateFormatterClass) Alloc() ISO8601DateFormatter {
	rv := objc.Call[ISO8601DateFormatter](ic, objc.Sel("alloc"))
	return rv
}

func ISO8601DateFormatter_Alloc() ISO8601DateFormatter {
	return ISO8601DateFormatterClass.Alloc()
}

func (ic _ISO8601DateFormatterClass) New() ISO8601DateFormatter {
	rv := objc.Call[ISO8601DateFormatter](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewISO8601DateFormatter() ISO8601DateFormatter {
	return ISO8601DateFormatterClass.New()
}

// Creates and returns a date object from the specified ISO 8601 formatted string representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsiso8601dateformatter/1643127-datefromstring?language=objc
func (i_ ISO8601DateFormatter) DateFromString(string_ string) Date {
	rv := objc.Call[Date](i_, objc.Sel("dateFromString:"), string_)
	return rv
}

// Creates and returns an ISO 8601 formatted string representation of the specified date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsiso8601dateformatter/1643076-stringfromdate?language=objc
func (i_ ISO8601DateFormatter) StringFromDate(date IDate) string {
	rv := objc.Call[string](i_, objc.Sel("stringFromDate:"), objc.Ptr(date))
	return rv
}

// The time zone used to create and parse date representations. When unspecified, GMT is used. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsiso8601dateformatter/1643185-timezone?language=objc
func (i_ ISO8601DateFormatter) TimeZone() TimeZone {
	rv := objc.Call[TimeZone](i_, objc.Sel("timeZone"))
	return rv
}

// The time zone used to create and parse date representations. When unspecified, GMT is used. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsiso8601dateformatter/1643185-timezone?language=objc
func (i_ ISO8601DateFormatter) SetTimeZone(value ITimeZone) {
	objc.Call[objc.Void](i_, objc.Sel("setTimeZone:"), objc.Ptr(value))
}

// Options for generating and parsing ISO 8601 date representations. See NSISO8601DateFormatOptions for possible values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsiso8601dateformatter/1643324-formatoptions?language=objc
func (i_ ISO8601DateFormatter) FormatOptions() ISO8601DateFormatOptions {
	rv := objc.Call[ISO8601DateFormatOptions](i_, objc.Sel("formatOptions"))
	return rv
}

// Options for generating and parsing ISO 8601 date representations. See NSISO8601DateFormatOptions for possible values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsiso8601dateformatter/1643324-formatoptions?language=objc
func (i_ ISO8601DateFormatter) SetFormatOptions(value ISO8601DateFormatOptions) {
	objc.Call[objc.Void](i_, objc.Sel("setFormatOptions:"), value)
}
