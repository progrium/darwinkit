// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MeasurementFormatter] class.
var MeasurementFormatterClass = _MeasurementFormatterClass{objc.GetClass("NSMeasurementFormatter")}

type _MeasurementFormatterClass struct {
	objc.Class
}

// An interface definition for the [MeasurementFormatter] class.
type IMeasurementFormatter interface {
	IFormatter
	StringFromUnit(unit IUnit) string
	StringFromMeasurement(measurement IMeasurement) string
	UnitOptions() MeasurementFormatterUnitOptions
	SetUnitOptions(value MeasurementFormatterUnitOptions)
	Locale() Locale
	SetLocale(value ILocale)
	UnitStyle() FormattingUnitStyle
	SetUnitStyle(value FormattingUnitStyle)
	NumberFormatter() NumberFormatter
	SetNumberFormatter(value INumberFormatter)
}

// A formatter that provides localized representations of units and measurements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmeasurementformatter?language=objc
type MeasurementFormatter struct {
	Formatter
}

func MeasurementFormatterFrom(ptr unsafe.Pointer) MeasurementFormatter {
	return MeasurementFormatter{
		Formatter: FormatterFrom(ptr),
	}
}

func (mc _MeasurementFormatterClass) Alloc() MeasurementFormatter {
	rv := objc.Call[MeasurementFormatter](mc, objc.Sel("alloc"))
	return rv
}

func MeasurementFormatter_Alloc() MeasurementFormatter {
	return MeasurementFormatterClass.Alloc()
}

func (mc _MeasurementFormatterClass) New() MeasurementFormatter {
	rv := objc.Call[MeasurementFormatter](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMeasurementFormatter() MeasurementFormatter {
	return MeasurementFormatterClass.New()
}

func (m_ MeasurementFormatter) Init() MeasurementFormatter {
	rv := objc.Call[MeasurementFormatter](m_, objc.Sel("init"))
	return rv
}

// Creates and returns a localized string representation of the provided unit of measure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmeasurementformatter/1642059-stringfromunit?language=objc
func (m_ MeasurementFormatter) StringFromUnit(unit IUnit) string {
	rv := objc.Call[string](m_, objc.Sel("stringFromUnit:"), objc.Ptr(unit))
	return rv
}

// Creates and returns a localized string representation of the provided measurement. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmeasurementformatter/1642057-stringfrommeasurement?language=objc
func (m_ MeasurementFormatter) StringFromMeasurement(measurement IMeasurement) string {
	rv := objc.Call[string](m_, objc.Sel("stringFromMeasurement:"), objc.Ptr(measurement))
	return rv
}

// The options for how the unit is formatted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmeasurementformatter/1642066-unitoptions?language=objc
func (m_ MeasurementFormatter) UnitOptions() MeasurementFormatterUnitOptions {
	rv := objc.Call[MeasurementFormatterUnitOptions](m_, objc.Sel("unitOptions"))
	return rv
}

// The options for how the unit is formatted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmeasurementformatter/1642066-unitoptions?language=objc
func (m_ MeasurementFormatter) SetUnitOptions(value MeasurementFormatterUnitOptions) {
	objc.Call[objc.Void](m_, objc.Sel("setUnitOptions:"), value)
}

// The locale of the formatter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmeasurementformatter/1642061-locale?language=objc
func (m_ MeasurementFormatter) Locale() Locale {
	rv := objc.Call[Locale](m_, objc.Sel("locale"))
	return rv
}

// The locale of the formatter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmeasurementformatter/1642061-locale?language=objc
func (m_ MeasurementFormatter) SetLocale(value ILocale) {
	objc.Call[objc.Void](m_, objc.Sel("setLocale:"), objc.Ptr(value))
}

// The unit style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmeasurementformatter/1642067-unitstyle?language=objc
func (m_ MeasurementFormatter) UnitStyle() FormattingUnitStyle {
	rv := objc.Call[FormattingUnitStyle](m_, objc.Sel("unitStyle"))
	return rv
}

// The unit style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmeasurementformatter/1642067-unitstyle?language=objc
func (m_ MeasurementFormatter) SetUnitStyle(value FormattingUnitStyle) {
	objc.Call[objc.Void](m_, objc.Sel("setUnitStyle:"), value)
}

// The number formatter used to format the quantity of a measurement. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmeasurementformatter/1642056-numberformatter?language=objc
func (m_ MeasurementFormatter) NumberFormatter() NumberFormatter {
	rv := objc.Call[NumberFormatter](m_, objc.Sel("numberFormatter"))
	return rv
}

// The number formatter used to format the quantity of a measurement. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmeasurementformatter/1642056-numberformatter?language=objc
func (m_ MeasurementFormatter) SetNumberFormatter(value INumberFormatter) {
	objc.Call[objc.Void](m_, objc.Sel("setNumberFormatter:"), objc.Ptr(value))
}
