// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ByteCountFormatter] class.
var ByteCountFormatterClass = _ByteCountFormatterClass{objc.GetClass("NSByteCountFormatter")}

type _ByteCountFormatterClass struct {
	objc.Class
}

// An interface definition for the [ByteCountFormatter] class.
type IByteCountFormatter interface {
	IFormatter
	StringFromMeasurement(measurement IMeasurement) string
	StringFromByteCount(byteCount int64) string
	IsAdaptive() bool
	SetAdaptive(value bool)
	ZeroPadsFractionDigits() bool
	SetZeroPadsFractionDigits(value bool)
	IncludesCount() bool
	SetIncludesCount(value bool)
	AllowsNonnumericFormatting() bool
	SetAllowsNonnumericFormatting(value bool)
	AllowedUnits() ByteCountFormatterUnits
	SetAllowedUnits(value ByteCountFormatterUnits)
	FormattingContext() FormattingContext
	SetFormattingContext(value FormattingContext)
	CountStyle() ByteCountFormatterCountStyle
	SetCountStyle(value ByteCountFormatterCountStyle)
	IncludesActualByteCount() bool
	SetIncludesActualByteCount(value bool)
	IncludesUnit() bool
	SetIncludesUnit(value bool)
}

// A formatter that converts a byte count value into a localized description that is formatted with the appropriate byte modifier (KB, MB, GB and so on). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbytecountformatter?language=objc
type ByteCountFormatter struct {
	Formatter
}

func ByteCountFormatterFrom(ptr unsafe.Pointer) ByteCountFormatter {
	return ByteCountFormatter{
		Formatter: FormatterFrom(ptr),
	}
}

func (bc _ByteCountFormatterClass) Alloc() ByteCountFormatter {
	rv := objc.Call[ByteCountFormatter](bc, objc.Sel("alloc"))
	return rv
}

func ByteCountFormatter_Alloc() ByteCountFormatter {
	return ByteCountFormatterClass.Alloc()
}

func (bc _ByteCountFormatterClass) New() ByteCountFormatter {
	rv := objc.Call[ByteCountFormatter](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewByteCountFormatter() ByteCountFormatter {
	return ByteCountFormatterClass.New()
}

func (b_ ByteCountFormatter) Init() ByteCountFormatter {
	rv := objc.Call[ByteCountFormatter](b_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbytecountformatter/3172523-stringfrommeasurement?language=objc
func (b_ ByteCountFormatter) StringFromMeasurement(measurement IMeasurement) string {
	rv := objc.Call[string](b_, objc.Sel("stringFromMeasurement:"), objc.Ptr(measurement))
	return rv
}

// Converts a byte count into a string without creating an NSNumber object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbytecountformatter/1415338-stringfrombytecount?language=objc
func (b_ ByteCountFormatter) StringFromByteCount(byteCount int64) string {
	rv := objc.Call[string](b_, objc.Sel("stringFromByteCount:"), byteCount)
	return rv
}

// Determines the display style of the size representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbytecountformatter/1417887-adaptive?language=objc
func (b_ ByteCountFormatter) IsAdaptive() bool {
	rv := objc.Call[bool](b_, objc.Sel("isAdaptive"))
	return rv
}

// Determines the display style of the size representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbytecountformatter/1417887-adaptive?language=objc
func (b_ ByteCountFormatter) SetAdaptive(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setAdaptive:"), value)
}

// Determines whether to zero pad fraction digits so a consistent number of characters is displayed in a representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbytecountformatter/1409630-zeropadsfractiondigits?language=objc
func (b_ ByteCountFormatter) ZeroPadsFractionDigits() bool {
	rv := objc.Call[bool](b_, objc.Sel("zeroPadsFractionDigits"))
	return rv
}

// Determines whether to zero pad fraction digits so a consistent number of characters is displayed in a representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbytecountformatter/1409630-zeropadsfractiondigits?language=objc
func (b_ ByteCountFormatter) SetZeroPadsFractionDigits(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setZeroPadsFractionDigits:"), value)
}

// Determines whether to include the count in the resulting formatted string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbytecountformatter/1409874-includescount?language=objc
func (b_ ByteCountFormatter) IncludesCount() bool {
	rv := objc.Call[bool](b_, objc.Sel("includesCount"))
	return rv
}

// Determines whether to include the count in the resulting formatted string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbytecountformatter/1409874-includescount?language=objc
func (b_ ByteCountFormatter) SetIncludesCount(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setIncludesCount:"), value)
}

// Determines whether to allow more natural display of some values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbytecountformatter/1408929-allowsnonnumericformatting?language=objc
func (b_ ByteCountFormatter) AllowsNonnumericFormatting() bool {
	rv := objc.Call[bool](b_, objc.Sel("allowsNonnumericFormatting"))
	return rv
}

// Determines whether to allow more natural display of some values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbytecountformatter/1408929-allowsnonnumericformatting?language=objc
func (b_ ByteCountFormatter) SetAllowsNonnumericFormatting(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setAllowsNonnumericFormatting:"), value)
}

// Specify the units that can be used in the output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbytecountformatter/1409137-allowedunits?language=objc
func (b_ ByteCountFormatter) AllowedUnits() ByteCountFormatterUnits {
	rv := objc.Call[ByteCountFormatterUnits](b_, objc.Sel("allowedUnits"))
	return rv
}

// Specify the units that can be used in the output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbytecountformatter/1409137-allowedunits?language=objc
func (b_ ByteCountFormatter) SetAllowedUnits(value ByteCountFormatterUnits) {
	objc.Call[objc.Void](b_, objc.Sel("setAllowedUnits:"), value)
}

// Specify the formatting context for the formatted string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbytecountformatter/1412185-formattingcontext?language=objc
func (b_ ByteCountFormatter) FormattingContext() FormattingContext {
	rv := objc.Call[FormattingContext](b_, objc.Sel("formattingContext"))
	return rv
}

// Specify the formatting context for the formatted string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbytecountformatter/1412185-formattingcontext?language=objc
func (b_ ByteCountFormatter) SetFormattingContext(value FormattingContext) {
	objc.Call[objc.Void](b_, objc.Sel("setFormattingContext:"), value)
}

// Specify the number of bytes to be used for kilobytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbytecountformatter/1417194-countstyle?language=objc
func (b_ ByteCountFormatter) CountStyle() ByteCountFormatterCountStyle {
	rv := objc.Call[ByteCountFormatterCountStyle](b_, objc.Sel("countStyle"))
	return rv
}

// Specify the number of bytes to be used for kilobytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbytecountformatter/1417194-countstyle?language=objc
func (b_ ByteCountFormatter) SetCountStyle(value ByteCountFormatterCountStyle) {
	objc.Call[objc.Void](b_, objc.Sel("setCountStyle:"), value)
}

// Determines whether to include the number of bytes after the formatted string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbytecountformatter/1411068-includesactualbytecount?language=objc
func (b_ ByteCountFormatter) IncludesActualByteCount() bool {
	rv := objc.Call[bool](b_, objc.Sel("includesActualByteCount"))
	return rv
}

// Determines whether to include the number of bytes after the formatted string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbytecountformatter/1411068-includesactualbytecount?language=objc
func (b_ ByteCountFormatter) SetIncludesActualByteCount(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setIncludesActualByteCount:"), value)
}

// Determines whether to include the units in the resulting formatted string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbytecountformatter/1415784-includesunit?language=objc
func (b_ ByteCountFormatter) IncludesUnit() bool {
	rv := objc.Call[bool](b_, objc.Sel("includesUnit"))
	return rv
}

// Determines whether to include the units in the resulting formatted string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbytecountformatter/1415784-includesunit?language=objc
func (b_ ByteCountFormatter) SetIncludesUnit(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setIncludesUnit:"), value)
}
