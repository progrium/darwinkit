// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TimeZone] class.
var TimeZoneClass = _TimeZoneClass{objc.GetClass("NSTimeZone")}

type _TimeZoneClass struct {
	objc.Class
}

// An interface definition for the [TimeZone] class.
type ITimeZone interface {
	objc.IObject
	IsDaylightSavingTimeForDate(aDate IDate) bool
	DaylightSavingTimeOffsetForDate(aDate IDate) TimeInterval
	LocalizedNameLocale(style TimeZoneNameStyle, locale ILocale) string
	SecondsFromGMTForDate(aDate IDate) int
	AbbreviationForDate(aDate IDate) string
	IsEqualToTimeZone(aTimeZone ITimeZone) bool
	IsDaylightSavingTime() bool
	DaylightSavingTimeOffset() TimeInterval
	Name() string
	SecondsFromGMT() int
	NextDaylightSavingTimeTransition() Date
	Data() []byte
	Description() string
	Abbreviation() string
}

// Information about standard time conventions associated with a specific geopolitical region. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone?language=objc
type TimeZone struct {
	objc.Object
}

func TimeZoneFrom(ptr unsafe.Pointer) TimeZone {
	return TimeZone{
		Object: objc.ObjectFrom(ptr),
	}
}

func (tc _TimeZoneClass) TimeZoneWithAbbreviation(abbreviation string) TimeZone {
	rv := objc.Call[TimeZone](tc, objc.Sel("timeZoneWithAbbreviation:"), abbreviation)
	return rv
}

// Returns the time zone object identified by a given abbreviation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387248-timezonewithabbreviation?language=objc
func TimeZone_TimeZoneWithAbbreviation(abbreviation string) TimeZone {
	return TimeZoneClass.TimeZoneWithAbbreviation(abbreviation)
}

func (tc _TimeZoneClass) TimeZoneWithName(tzName string) TimeZone {
	rv := objc.Call[TimeZone](tc, objc.Sel("timeZoneWithName:"), tzName)
	return rv
}

// Returns the time zone object identified by a given identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387227-timezonewithname?language=objc
func TimeZone_TimeZoneWithName(tzName string) TimeZone {
	return TimeZoneClass.TimeZoneWithName(tzName)
}

func (t_ TimeZone) InitWithNameData(tzName string, aData []byte) TimeZone {
	rv := objc.Call[TimeZone](t_, objc.Sel("initWithName:data:"), tzName, aData)
	return rv
}

// Initializes a time zone with a given identifier and time zone data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387250-initwithname?language=objc
func NewTimeZoneWithNameData(tzName string, aData []byte) TimeZone {
	instance := TimeZoneClass.Alloc().InitWithNameData(tzName, aData)
	instance.Autorelease()
	return instance
}

func (tc _TimeZoneClass) TimeZoneForSecondsFromGMT(seconds int) TimeZone {
	rv := objc.Call[TimeZone](tc, objc.Sel("timeZoneForSecondsFromGMT:"), seconds)
	return rv
}

// Returns a time zone object offset from Greenwich Mean Time by a given number of seconds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387199-timezoneforsecondsfromgmt?language=objc
func TimeZone_TimeZoneForSecondsFromGMT(seconds int) TimeZone {
	return TimeZoneClass.TimeZoneForSecondsFromGMT(seconds)
}

func (tc _TimeZoneClass) Alloc() TimeZone {
	rv := objc.Call[TimeZone](tc, objc.Sel("alloc"))
	return rv
}

func TimeZone_Alloc() TimeZone {
	return TimeZoneClass.Alloc()
}

func (tc _TimeZoneClass) New() TimeZone {
	rv := objc.Call[TimeZone](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTimeZone() TimeZone {
	return TimeZoneClass.New()
}

func (t_ TimeZone) Init() TimeZone {
	rv := objc.Call[TimeZone](t_, objc.Sel("init"))
	return rv
}

// Indicates whether the receiver uses daylight saving time on a given date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387201-isdaylightsavingtimefordate?language=objc
func (t_ TimeZone) IsDaylightSavingTimeForDate(aDate IDate) bool {
	rv := objc.Call[bool](t_, objc.Sel("isDaylightSavingTimeForDate:"), objc.Ptr(aDate))
	return rv
}

// Returns the daylight saving time offset for a given date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387207-daylightsavingtimeoffsetfordate?language=objc
func (t_ TimeZone) DaylightSavingTimeOffsetForDate(aDate IDate) TimeInterval {
	rv := objc.Call[TimeInterval](t_, objc.Sel("daylightSavingTimeOffsetForDate:"), objc.Ptr(aDate))
	return rv
}

// Returns the localized name of the time zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387193-localizedname?language=objc
func (t_ TimeZone) LocalizedNameLocale(style TimeZoneNameStyle, locale ILocale) string {
	rv := objc.Call[string](t_, objc.Sel("localizedName:locale:"), style, objc.Ptr(locale))
	return rv
}

// Returns the difference in seconds between the receiver and Greenwich Mean Time at a given date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387242-secondsfromgmtfordate?language=objc
func (t_ TimeZone) SecondsFromGMTForDate(aDate IDate) int {
	rv := objc.Call[int](t_, objc.Sel("secondsFromGMTForDate:"), objc.Ptr(aDate))
	return rv
}

// Clears any time zone value cached for the systemTimeZone property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387189-resetsystemtimezone?language=objc
func (tc _TimeZoneClass) ResetSystemTimeZone() {
	objc.Call[objc.Void](tc, objc.Sel("resetSystemTimeZone"))
}

// Clears any time zone value cached for the systemTimeZone property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387189-resetsystemtimezone?language=objc
func TimeZone_ResetSystemTimeZone() {
	TimeZoneClass.ResetSystemTimeZone()
}

// Returns the abbreviation for the receiver at a given date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387237-abbreviationfordate?language=objc
func (t_ TimeZone) AbbreviationForDate(aDate IDate) string {
	rv := objc.Call[string](t_, objc.Sel("abbreviationForDate:"), objc.Ptr(aDate))
	return rv
}

// Indicates whether the receiver has the same name and data as the specified time zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387211-isequaltotimezone?language=objc
func (t_ TimeZone) IsEqualToTimeZone(aTimeZone ITimeZone) bool {
	rv := objc.Call[bool](t_, objc.Sel("isEqualToTimeZone:"), objc.Ptr(aTimeZone))
	return rv
}

// The default time zone for the current app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387244-defaulttimezone?language=objc
func (tc _TimeZoneClass) DefaultTimeZone() TimeZone {
	rv := objc.Call[TimeZone](tc, objc.Sel("defaultTimeZone"))
	return rv
}

// The default time zone for the current app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387244-defaulttimezone?language=objc
func TimeZone_DefaultTimeZone() TimeZone {
	return TimeZoneClass.DefaultTimeZone()
}

// The default time zone for the current app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387244-defaulttimezone?language=objc
func (tc _TimeZoneClass) SetDefaultTimeZone(value ITimeZone) {
	objc.Call[objc.Void](tc, objc.Sel("setDefaultTimeZone:"), objc.Ptr(value))
}

// The default time zone for the current app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387244-defaulttimezone?language=objc
func TimeZone_SetDefaultTimeZone(value ITimeZone) {
	TimeZoneClass.SetDefaultTimeZone(value)
}

// An object that tracks the current system time zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387209-localtimezone?language=objc
func (tc _TimeZoneClass) LocalTimeZone() TimeZone {
	rv := objc.Call[TimeZone](tc, objc.Sel("localTimeZone"))
	return rv
}

// An object that tracks the current system time zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387209-localtimezone?language=objc
func TimeZone_LocalTimeZone() TimeZone {
	return TimeZoneClass.LocalTimeZone()
}

// A Boolean value that indicates whether the receiver is currently using daylight saving time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387191-daylightsavingtime?language=objc
func (t_ TimeZone) IsDaylightSavingTime() bool {
	rv := objc.Call[bool](t_, objc.Sel("isDaylightSavingTime"))
	return rv
}

// The current daylight saving time offset of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387235-daylightsavingtimeoffset?language=objc
func (t_ TimeZone) DaylightSavingTimeOffset() TimeInterval {
	rv := objc.Call[TimeInterval](t_, objc.Sel("daylightSavingTimeOffset"))
	return rv
}

// The geopolitical region ID that identifies the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387233-name?language=objc
func (t_ TimeZone) Name() string {
	rv := objc.Call[string](t_, objc.Sel("name"))
	return rv
}

// Returns an array of strings listing the IDs of all the time zones known to the system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387223-knowntimezonenames?language=objc
func (tc _TimeZoneClass) KnownTimeZoneNames() []string {
	rv := objc.Call[[]string](tc, objc.Sel("knownTimeZoneNames"))
	return rv
}

// Returns an array of strings listing the IDs of all the time zones known to the system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387223-knowntimezonenames?language=objc
func TimeZone_KnownTimeZoneNames() []string {
	return TimeZoneClass.KnownTimeZoneNames()
}

// The current difference in seconds between the receiver and Greenwich Mean Time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387221-secondsfromgmt?language=objc
func (t_ TimeZone) SecondsFromGMT() int {
	rv := objc.Call[int](t_, objc.Sel("secondsFromGMT"))
	return rv
}

// The date of the next daylight saving time transition for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387183-nextdaylightsavingtimetransition?language=objc
func (t_ TimeZone) NextDaylightSavingTimeTransition() Date {
	rv := objc.Call[Date](t_, objc.Sel("nextDaylightSavingTimeTransition"))
	return rv
}

// The data that stores the information used by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387213-data?language=objc
func (t_ TimeZone) Data() []byte {
	rv := objc.Call[[]byte](t_, objc.Sel("data"))
	return rv
}

// Returns a dictionary holding the mappings of time zone abbreviations to time zone names. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387258-abbreviationdictionary?language=objc
func (tc _TimeZoneClass) AbbreviationDictionary() map[string]string {
	rv := objc.Call[map[string]string](tc, objc.Sel("abbreviationDictionary"))
	return rv
}

// Returns a dictionary holding the mappings of time zone abbreviations to time zone names. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387258-abbreviationdictionary?language=objc
func TimeZone_AbbreviationDictionary() map[string]string {
	return TimeZoneClass.AbbreviationDictionary()
}

// Returns a dictionary holding the mappings of time zone abbreviations to time zone names. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387258-abbreviationdictionary?language=objc
func (tc _TimeZoneClass) SetAbbreviationDictionary(value map[string]string) {
	objc.Call[objc.Void](tc, objc.Sel("setAbbreviationDictionary:"), value)
}

// Returns a dictionary holding the mappings of time zone abbreviations to time zone names. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387258-abbreviationdictionary?language=objc
func TimeZone_SetAbbreviationDictionary(value map[string]string) {
	TimeZoneClass.SetAbbreviationDictionary(value)
}

// Returns the time zone data version. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387187-timezonedataversion?language=objc
func (tc _TimeZoneClass) TimeZoneDataVersion() string {
	rv := objc.Call[string](tc, objc.Sel("timeZoneDataVersion"))
	return rv
}

// Returns the time zone data version. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387187-timezonedataversion?language=objc
func TimeZone_TimeZoneDataVersion() string {
	return TimeZoneClass.TimeZoneDataVersion()
}

// A textual description of the time zone including the name, abbreviation, offset from GMT, and whether or not daylight saving time is currently in effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387217-description?language=objc
func (t_ TimeZone) Description() string {
	rv := objc.Call[string](t_, objc.Sel("description"))
	return rv
}

// The time zone currently used by the system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387231-systemtimezone?language=objc
func (tc _TimeZoneClass) SystemTimeZone() TimeZone {
	rv := objc.Call[TimeZone](tc, objc.Sel("systemTimeZone"))
	return rv
}

// The time zone currently used by the system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387231-systemtimezone?language=objc
func TimeZone_SystemTimeZone() TimeZone {
	return TimeZoneClass.SystemTimeZone()
}

// The abbreviation for the receiver, such as “EDT” (Eastern Daylight Time). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstimezone/1387195-abbreviation?language=objc
func (t_ TimeZone) Abbreviation() string {
	rv := objc.Call[string](t_, objc.Sel("abbreviation"))
	return rv
}
