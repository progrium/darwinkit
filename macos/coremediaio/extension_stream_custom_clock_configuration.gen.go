// AUTO-GENERATED CODE, DO NOT MODIFY

package coremediaio

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ExtensionStreamCustomClockConfiguration] class.
var ExtensionStreamCustomClockConfigurationClass = _ExtensionStreamCustomClockConfigurationClass{objc.GetClass("CMIOExtensionStreamCustomClockConfiguration")}

type _ExtensionStreamCustomClockConfigurationClass struct {
	objc.Class
}

// An interface definition for the [ExtensionStreamCustomClockConfiguration] class.
type IExtensionStreamCustomClockConfiguration interface {
	objc.IObject
	NumberOfEventsForRateSmoothing() uint32
	NumberOfAveragesForRateSmoothing() uint32
	SourceIdentifier() foundation.UUID
	GetTimeCallMinimumInterval() coremedia.Time
	ClockName() string
}

// An object that describes the parameters to create a custom clock on the host side. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamcustomclockconfiguration?language=objc
type ExtensionStreamCustomClockConfiguration struct {
	objc.Object
}

func ExtensionStreamCustomClockConfigurationFrom(ptr unsafe.Pointer) ExtensionStreamCustomClockConfiguration {
	return ExtensionStreamCustomClockConfiguration{
		Object: objc.ObjectFrom(ptr),
	}
}

func (e_ ExtensionStreamCustomClockConfiguration) InitWithClockNameSourceIdentifierGetTimeCallMinimumIntervalNumberOfEventsForRateSmoothingNumberOfAveragesForRateSmoothing(clockName string, sourceIdentifier foundation.IUUID, getTimeCallMinimumInterval coremedia.Time, numberOfEventsForRateSmoothing uint32, numberOfAveragesForRateSmoothing uint32) ExtensionStreamCustomClockConfiguration {
	rv := objc.Call[ExtensionStreamCustomClockConfiguration](e_, objc.Sel("initWithClockName:sourceIdentifier:getTimeCallMinimumInterval:numberOfEventsForRateSmoothing:numberOfAveragesForRateSmoothing:"), clockName, objc.Ptr(sourceIdentifier), getTimeCallMinimumInterval, numberOfEventsForRateSmoothing, numberOfAveragesForRateSmoothing)
	return rv
}

// Creates a custom clock configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamcustomclockconfiguration/3915892-initwithclockname?language=objc
func NewExtensionStreamCustomClockConfigurationWithClockNameSourceIdentifierGetTimeCallMinimumIntervalNumberOfEventsForRateSmoothingNumberOfAveragesForRateSmoothing(clockName string, sourceIdentifier foundation.IUUID, getTimeCallMinimumInterval coremedia.Time, numberOfEventsForRateSmoothing uint32, numberOfAveragesForRateSmoothing uint32) ExtensionStreamCustomClockConfiguration {
	instance := ExtensionStreamCustomClockConfigurationClass.Alloc().InitWithClockNameSourceIdentifierGetTimeCallMinimumIntervalNumberOfEventsForRateSmoothingNumberOfAveragesForRateSmoothing(clockName, sourceIdentifier, getTimeCallMinimumInterval, numberOfEventsForRateSmoothing, numberOfAveragesForRateSmoothing)
	instance.Autorelease()
	return instance
}

func (ec _ExtensionStreamCustomClockConfigurationClass) CustomClockConfigurationWithClockNameSourceIdentifierGetTimeCallMinimumIntervalNumberOfEventsForRateSmoothingNumberOfAveragesForRateSmoothing(clockName string, sourceIdentifier foundation.IUUID, getTimeCallMinimumInterval coremedia.Time, numberOfEventsForRateSmoothing uint32, numberOfAveragesForRateSmoothing uint32) ExtensionStreamCustomClockConfiguration {
	rv := objc.Call[ExtensionStreamCustomClockConfiguration](ec, objc.Sel("customClockConfigurationWithClockName:sourceIdentifier:getTimeCallMinimumInterval:numberOfEventsForRateSmoothing:numberOfAveragesForRateSmoothing:"), clockName, objc.Ptr(sourceIdentifier), getTimeCallMinimumInterval, numberOfEventsForRateSmoothing, numberOfAveragesForRateSmoothing)
	return rv
}

// Returns a new a custom clock configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamcustomclockconfiguration/3915890-customclockconfigurationwithcloc?language=objc
func ExtensionStreamCustomClockConfiguration_CustomClockConfigurationWithClockNameSourceIdentifierGetTimeCallMinimumIntervalNumberOfEventsForRateSmoothingNumberOfAveragesForRateSmoothing(clockName string, sourceIdentifier foundation.IUUID, getTimeCallMinimumInterval coremedia.Time, numberOfEventsForRateSmoothing uint32, numberOfAveragesForRateSmoothing uint32) ExtensionStreamCustomClockConfiguration {
	return ExtensionStreamCustomClockConfigurationClass.CustomClockConfigurationWithClockNameSourceIdentifierGetTimeCallMinimumIntervalNumberOfEventsForRateSmoothingNumberOfAveragesForRateSmoothing(clockName, sourceIdentifier, getTimeCallMinimumInterval, numberOfEventsForRateSmoothing, numberOfAveragesForRateSmoothing)
}

func (ec _ExtensionStreamCustomClockConfigurationClass) Alloc() ExtensionStreamCustomClockConfiguration {
	rv := objc.Call[ExtensionStreamCustomClockConfiguration](ec, objc.Sel("alloc"))
	return rv
}

func ExtensionStreamCustomClockConfiguration_Alloc() ExtensionStreamCustomClockConfiguration {
	return ExtensionStreamCustomClockConfigurationClass.Alloc()
}

func (ec _ExtensionStreamCustomClockConfigurationClass) New() ExtensionStreamCustomClockConfiguration {
	rv := objc.Call[ExtensionStreamCustomClockConfiguration](ec, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewExtensionStreamCustomClockConfiguration() ExtensionStreamCustomClockConfiguration {
	return ExtensionStreamCustomClockConfigurationClass.New()
}

func (e_ ExtensionStreamCustomClockConfiguration) Init() ExtensionStreamCustomClockConfiguration {
	rv := objc.Call[ExtensionStreamCustomClockConfiguration](e_, objc.Sel("init"))
	return rv
}

// The number of events to use for rate smoothing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamcustomclockconfiguration/3915894-numberofeventsforratesmoothing?language=objc
func (e_ ExtensionStreamCustomClockConfiguration) NumberOfEventsForRateSmoothing() uint32 {
	rv := objc.Call[uint32](e_, objc.Sel("numberOfEventsForRateSmoothing"))
	return rv
}

// The number of averages to use for rate smoothing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamcustomclockconfiguration/3915893-numberofaveragesforratesmoothing?language=objc
func (e_ ExtensionStreamCustomClockConfiguration) NumberOfAveragesForRateSmoothing() uint32 {
	rv := objc.Call[uint32](e_, objc.Sel("numberOfAveragesForRateSmoothing"))
	return rv
}

// A universally unique identifier for the clock. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamcustomclockconfiguration/3915895-sourceidentifier?language=objc
func (e_ ExtensionStreamCustomClockConfiguration) SourceIdentifier() foundation.UUID {
	rv := objc.Call[foundation.UUID](e_, objc.Sel("sourceIdentifier"))
	return rv
}

// A minimum call time interval for the clock. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamcustomclockconfiguration/3915891-gettimecallminimuminterval?language=objc
func (e_ ExtensionStreamCustomClockConfiguration) GetTimeCallMinimumInterval() coremedia.Time {
	rv := objc.Call[coremedia.Time](e_, objc.Sel("getTimeCallMinimumInterval"))
	return rv
}

// The name of the clock. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionstreamcustomclockconfiguration/3915889-clockname?language=objc
func (e_ ExtensionStreamCustomClockConfiguration) ClockName() string {
	rv := objc.Call[string](e_, objc.Sel("clockName"))
	return rv
}
