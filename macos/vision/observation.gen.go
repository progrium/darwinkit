// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Observation] class.
var ObservationClass = _ObservationClass{objc.GetClass("VNObservation")}

type _ObservationClass struct {
	objc.Class
}

// An interface definition for the [Observation] class.
type IObservation interface {
	objc.IObject
	TimeRange() coremedia.TimeRange
	Uuid() foundation.UUID
	Confidence() Confidence
}

// The abstract superclass for analysis results. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnobservation?language=objc
type Observation struct {
	objc.Object
}

func ObservationFrom(ptr unsafe.Pointer) Observation {
	return Observation{
		Object: objc.ObjectFrom(ptr),
	}
}

func (oc _ObservationClass) Alloc() Observation {
	rv := objc.Call[Observation](oc, objc.Sel("alloc"))
	return rv
}

func Observation_Alloc() Observation {
	return ObservationClass.Alloc()
}

func (oc _ObservationClass) New() Observation {
	rv := objc.Call[Observation](oc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewObservation() Observation {
	return ObservationClass.New()
}

func (o_ Observation) Init() Observation {
	rv := objc.Call[Observation](o_, objc.Sel("init"))
	return rv
}

// The time range of the reported observation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnobservation/3548370-timerange?language=objc
func (o_ Observation) TimeRange() coremedia.TimeRange {
	rv := objc.Call[coremedia.TimeRange](o_, objc.Sel("timeRange"))
	return rv
}

// A unique identifier assigned to the Vision observation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnobservation/2879296-uuid?language=objc
func (o_ Observation) Uuid() foundation.UUID {
	rv := objc.Call[foundation.UUID](o_, objc.Sel("uuid"))
	return rv
}

// The level of confidence in the observation’s accuracy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnobservation/2867220-confidence?language=objc
func (o_ Observation) Confidence() Confidence {
	rv := objc.Call[Confidence](o_, objc.Sel("confidence"))
	return rv
}
