// AUTO-GENERATED CODE, DO NOT MODIFY
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var MeasurementFormatterClass = _MeasurementFormatterClass{objc.GetClass("NSMeasurementFormatter")}

type _MeasurementFormatterClass struct {
	objc.Class
}

type IMeasurementFormatter interface {
	IFormatter
	StringFromMeasurement(measurement IMeasurement) string
	StringFromUnit(unit IUnit) string
	UnitOptions() MeasurementFormatterUnitOptions
	SetUnitOptions(value MeasurementFormatterUnitOptions)
	UnitStyle() FormattingUnitStyle
	SetUnitStyle(value FormattingUnitStyle)
	Locale() Locale
	SetLocale(value ILocale)
	NumberFormatter() NumberFormatter
	SetNumberFormatter(value INumberFormatter)
}

type MeasurementFormatter struct {
	Formatter
}

func MakeMeasurementFormatter(ptr unsafe.Pointer) MeasurementFormatter {
	return MeasurementFormatter{
		Formatter: MakeFormatter(ptr),
	}
}

func (mc _MeasurementFormatterClass) Alloc() MeasurementFormatter {
	rv := objc.CallMethod[MeasurementFormatter](mc, objc.GetSelector("alloc"))
	return rv
}

func MeasurementFormatter_Alloc() MeasurementFormatter {
	return MeasurementFormatterClass.Alloc()
}

func (mc _MeasurementFormatterClass) New() MeasurementFormatter {
	rv := objc.CallMethod[MeasurementFormatter](mc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewMeasurementFormatter() MeasurementFormatter {
	return MeasurementFormatterClass.New()
}

func MeasurementFormatter_New() MeasurementFormatter {
	return MeasurementFormatterClass.New()
}

func (m_ MeasurementFormatter) Init() MeasurementFormatter {
	rv := objc.CallMethod[MeasurementFormatter](m_, objc.GetSelector("init"))
	return rv
}

func MeasurementFormatter_Init() MeasurementFormatter {
	return MeasurementFormatterClass.Alloc().Init()
}

func (m_ MeasurementFormatter) StringFromMeasurement(measurement IMeasurement) string {
	rv := objc.CallMethod[string](m_, objc.GetSelector("stringFromMeasurement:"), objc.ExtractPtr(measurement))
	return rv
}

func (m_ MeasurementFormatter) StringFromUnit(unit IUnit) string {
	rv := objc.CallMethod[string](m_, objc.GetSelector("stringFromUnit:"), objc.ExtractPtr(unit))
	return rv
}

func (m_ MeasurementFormatter) UnitOptions() MeasurementFormatterUnitOptions {
	rv := objc.CallMethod[MeasurementFormatterUnitOptions](m_, objc.GetSelector("unitOptions"))
	return rv
}

func (m_ MeasurementFormatter) SetUnitOptions(value MeasurementFormatterUnitOptions) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setUnitOptions:"), value)
}

func (m_ MeasurementFormatter) UnitStyle() FormattingUnitStyle {
	rv := objc.CallMethod[FormattingUnitStyle](m_, objc.GetSelector("unitStyle"))
	return rv
}

func (m_ MeasurementFormatter) SetUnitStyle(value FormattingUnitStyle) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setUnitStyle:"), value)
}

func (m_ MeasurementFormatter) Locale() Locale {
	rv := objc.CallMethod[Locale](m_, objc.GetSelector("locale"))
	return rv
}

func (m_ MeasurementFormatter) SetLocale(value ILocale) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setLocale:"), objc.ExtractPtr(value))
}

func (m_ MeasurementFormatter) NumberFormatter() NumberFormatter {
	rv := objc.CallMethod[NumberFormatter](m_, objc.GetSelector("numberFormatter"))
	return rv
}

func (m_ MeasurementFormatter) SetNumberFormatter(value INumberFormatter) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setNumberFormatter:"), objc.ExtractPtr(value))
}
