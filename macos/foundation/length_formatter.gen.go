// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [LengthFormatter] class.
var LengthFormatterClass = _LengthFormatterClass{objc.GetClass("NSLengthFormatter")}

type _LengthFormatterClass struct {
	objc.Class
}

// An interface definition for the [LengthFormatter] class.
type ILengthFormatter interface {
	IFormatter
	StringFromValueUnit(value float64, unit LengthFormatterUnit) string
	StringFromMeters(numberInMeters float64) string
	UnitStringFromValueUnit(value float64, unit LengthFormatterUnit) string
	UnitStringFromMetersUsedUnit(numberInMeters float64, unitp *LengthFormatterUnit) string
	IsForPersonHeightUse() bool
	SetForPersonHeightUse(value bool)
	UnitStyle() FormattingUnitStyle
	SetUnitStyle(value FormattingUnitStyle)
	NumberFormatter() NumberFormatter
	SetNumberFormatter(value INumberFormatter)
}

// A formatter that provides localized descriptions of linear distances, such as length and height measurements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslengthformatter?language=objc
type LengthFormatter struct {
	Formatter
}

func LengthFormatterFrom(ptr unsafe.Pointer) LengthFormatter {
	return LengthFormatter{
		Formatter: FormatterFrom(ptr),
	}
}

func (lc _LengthFormatterClass) Alloc() LengthFormatter {
	rv := objc.Call[LengthFormatter](lc, objc.Sel("alloc"))
	return rv
}

func LengthFormatter_Alloc() LengthFormatter {
	return LengthFormatterClass.Alloc()
}

func (lc _LengthFormatterClass) New() LengthFormatter {
	rv := objc.Call[LengthFormatter](lc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewLengthFormatter() LengthFormatter {
	return LengthFormatterClass.New()
}

func (l_ LengthFormatter) Init() LengthFormatter {
	rv := objc.Call[LengthFormatter](l_, objc.Sel("init"))
	return rv
}

// Returns a properly formatted length string for the given value and unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslengthformatter/1418018-stringfromvalue?language=objc
func (l_ LengthFormatter) StringFromValueUnit(value float64, unit LengthFormatterUnit) string {
	rv := objc.Call[string](l_, objc.Sel("stringFromValue:unit:"), value, unit)
	return rv
}

// Returns a length string for the provided value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslengthformatter/1416111-stringfrommeters?language=objc
func (l_ LengthFormatter) StringFromMeters(numberInMeters float64) string {
	rv := objc.Call[string](l_, objc.Sel("stringFromMeters:"), numberInMeters)
	return rv
}

// Returns the unit string based on the provided value and unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslengthformatter/1416076-unitstringfromvalue?language=objc
func (l_ LengthFormatter) UnitStringFromValueUnit(value float64, unit LengthFormatterUnit) string {
	rv := objc.Call[string](l_, objc.Sel("unitStringFromValue:unit:"), value, unit)
	return rv
}

// Returns the unit string for the provided value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslengthformatter/1407661-unitstringfrommeters?language=objc
func (l_ LengthFormatter) UnitStringFromMetersUsedUnit(numberInMeters float64, unitp *LengthFormatterUnit) string {
	rv := objc.Call[string](l_, objc.Sel("unitStringFromMeters:usedUnit:"), numberInMeters, unitp)
	return rv
}

// A Boolean value that indicates whether the resulting string represents a person’s height. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslengthformatter/1416517-forpersonheightuse?language=objc
func (l_ LengthFormatter) IsForPersonHeightUse() bool {
	rv := objc.Call[bool](l_, objc.Sel("isForPersonHeightUse"))
	return rv
}

// A Boolean value that indicates whether the resulting string represents a person’s height. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslengthformatter/1416517-forpersonheightuse?language=objc
func (l_ LengthFormatter) SetForPersonHeightUse(value bool) {
	objc.Call[objc.Void](l_, objc.Sel("setForPersonHeightUse:"), value)
}

// The unit style used by this formatter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslengthformatter/1409965-unitstyle?language=objc
func (l_ LengthFormatter) UnitStyle() FormattingUnitStyle {
	rv := objc.Call[FormattingUnitStyle](l_, objc.Sel("unitStyle"))
	return rv
}

// The unit style used by this formatter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslengthformatter/1409965-unitstyle?language=objc
func (l_ LengthFormatter) SetUnitStyle(value FormattingUnitStyle) {
	objc.Call[objc.Void](l_, objc.Sel("setUnitStyle:"), value)
}

// The number formatter used to format the numbers in length strings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslengthformatter/1417778-numberformatter?language=objc
func (l_ LengthFormatter) NumberFormatter() NumberFormatter {
	rv := objc.Call[NumberFormatter](l_, objc.Sel("numberFormatter"))
	return rv
}

// The number formatter used to format the numbers in length strings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslengthformatter/1417778-numberformatter?language=objc
func (l_ LengthFormatter) SetNumberFormatter(value INumberFormatter) {
	objc.Call[objc.Void](l_, objc.Sel("setNumberFormatter:"), objc.Ptr(value))
}
