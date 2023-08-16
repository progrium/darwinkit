// AUTO-GENERATED CODE, DO NOT MODIFY

package quartzcore

import (
	"github.com/progrium/macdriver/macos/corefoundation"
	"github.com/progrium/macdriver/objc"
)

// Methods that model a hierarchical timing system, allowing objects to map time between their parent and local time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/camediatiming?language=objc
type PMediaTiming interface {
	// optional
	SetTimeOffset(value corefoundation.TimeInterval)
	HasSetTimeOffset() bool

	// optional
	TimeOffset() corefoundation.TimeInterval
	HasTimeOffset() bool

	// optional
	SetAutoreverses(value bool)
	HasSetAutoreverses() bool

	// optional
	Autoreverses() bool
	HasAutoreverses() bool

	// optional
	SetRepeatDuration(value corefoundation.TimeInterval)
	HasSetRepeatDuration() bool

	// optional
	RepeatDuration() corefoundation.TimeInterval
	HasRepeatDuration() bool

	// optional
	SetFillMode(value MediaTimingFillMode)
	HasSetFillMode() bool

	// optional
	FillMode() MediaTimingFillMode
	HasFillMode() bool

	// optional
	SetSpeed(value float64)
	HasSetSpeed() bool

	// optional
	Speed() float64
	HasSpeed() bool

	// optional
	SetRepeatCount(value float64)
	HasSetRepeatCount() bool

	// optional
	RepeatCount() float64
	HasRepeatCount() bool

	// optional
	SetBeginTime(value corefoundation.TimeInterval)
	HasSetBeginTime() bool

	// optional
	BeginTime() corefoundation.TimeInterval
	HasBeginTime() bool

	// optional
	SetDuration(value corefoundation.TimeInterval)
	HasSetDuration() bool

	// optional
	Duration() corefoundation.TimeInterval
	HasDuration() bool
}

// A concrete type wrapper for the [PMediaTiming] protocol.
type MediaTimingWrapper struct {
	objc.Object
}

func (m_ MediaTimingWrapper) HasSetTimeOffset() bool {
	return m_.RespondsToSelector(objc.Sel("setTimeOffset:"))
}

// Specifies an additional time offset in active local time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/camediatiming/1427650-timeoffset?language=objc
func (m_ MediaTimingWrapper) SetTimeOffset(value corefoundation.TimeInterval) {
	objc.Call[objc.Void](m_, objc.Sel("setTimeOffset:"), value)
}

func (m_ MediaTimingWrapper) HasTimeOffset() bool {
	return m_.RespondsToSelector(objc.Sel("timeOffset"))
}

// Specifies an additional time offset in active local time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/camediatiming/1427650-timeoffset?language=objc
func (m_ MediaTimingWrapper) TimeOffset() corefoundation.TimeInterval {
	rv := objc.Call[corefoundation.TimeInterval](m_, objc.Sel("timeOffset"))
	return rv
}

func (m_ MediaTimingWrapper) HasSetAutoreverses() bool {
	return m_.RespondsToSelector(objc.Sel("setAutoreverses:"))
}

// Determines if the receiver plays in the reverse upon completion. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/camediatiming/1427645-autoreverses?language=objc
func (m_ MediaTimingWrapper) SetAutoreverses(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setAutoreverses:"), value)
}

func (m_ MediaTimingWrapper) HasAutoreverses() bool {
	return m_.RespondsToSelector(objc.Sel("autoreverses"))
}

// Determines if the receiver plays in the reverse upon completion. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/camediatiming/1427645-autoreverses?language=objc
func (m_ MediaTimingWrapper) Autoreverses() bool {
	rv := objc.Call[bool](m_, objc.Sel("autoreverses"))
	return rv
}

func (m_ MediaTimingWrapper) HasSetRepeatDuration() bool {
	return m_.RespondsToSelector(objc.Sel("setRepeatDuration:"))
}

// Determines how many seconds the animation will repeat for. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/camediatiming/1427643-repeatduration?language=objc
func (m_ MediaTimingWrapper) SetRepeatDuration(value corefoundation.TimeInterval) {
	objc.Call[objc.Void](m_, objc.Sel("setRepeatDuration:"), value)
}

func (m_ MediaTimingWrapper) HasRepeatDuration() bool {
	return m_.RespondsToSelector(objc.Sel("repeatDuration"))
}

// Determines how many seconds the animation will repeat for. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/camediatiming/1427643-repeatduration?language=objc
func (m_ MediaTimingWrapper) RepeatDuration() corefoundation.TimeInterval {
	rv := objc.Call[corefoundation.TimeInterval](m_, objc.Sel("repeatDuration"))
	return rv
}

func (m_ MediaTimingWrapper) HasSetFillMode() bool {
	return m_.RespondsToSelector(objc.Sel("setFillMode:"))
}

// Determines if the receiver’s presentation is frozen or removed once its active duration has completed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/camediatiming/1427656-fillmode?language=objc
func (m_ MediaTimingWrapper) SetFillMode(value MediaTimingFillMode) {
	objc.Call[objc.Void](m_, objc.Sel("setFillMode:"), value)
}

func (m_ MediaTimingWrapper) HasFillMode() bool {
	return m_.RespondsToSelector(objc.Sel("fillMode"))
}

// Determines if the receiver’s presentation is frozen or removed once its active duration has completed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/camediatiming/1427656-fillmode?language=objc
func (m_ MediaTimingWrapper) FillMode() MediaTimingFillMode {
	rv := objc.Call[MediaTimingFillMode](m_, objc.Sel("fillMode"))
	return rv
}

func (m_ MediaTimingWrapper) HasSetSpeed() bool {
	return m_.RespondsToSelector(objc.Sel("setSpeed:"))
}

// Specifies how time is mapped to receiver’s time space from the parent time space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/camediatiming/1427647-speed?language=objc
func (m_ MediaTimingWrapper) SetSpeed(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setSpeed:"), value)
}

func (m_ MediaTimingWrapper) HasSpeed() bool {
	return m_.RespondsToSelector(objc.Sel("speed"))
}

// Specifies how time is mapped to receiver’s time space from the parent time space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/camediatiming/1427647-speed?language=objc
func (m_ MediaTimingWrapper) Speed() float64 {
	rv := objc.Call[float64](m_, objc.Sel("speed"))
	return rv
}

func (m_ MediaTimingWrapper) HasSetRepeatCount() bool {
	return m_.RespondsToSelector(objc.Sel("setRepeatCount:"))
}

// Determines the number of times the animation will repeat. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/camediatiming/1427666-repeatcount?language=objc
func (m_ MediaTimingWrapper) SetRepeatCount(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setRepeatCount:"), value)
}

func (m_ MediaTimingWrapper) HasRepeatCount() bool {
	return m_.RespondsToSelector(objc.Sel("repeatCount"))
}

// Determines the number of times the animation will repeat. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/camediatiming/1427666-repeatcount?language=objc
func (m_ MediaTimingWrapper) RepeatCount() float64 {
	rv := objc.Call[float64](m_, objc.Sel("repeatCount"))
	return rv
}

func (m_ MediaTimingWrapper) HasSetBeginTime() bool {
	return m_.RespondsToSelector(objc.Sel("setBeginTime:"))
}

// Specifies the begin time of the receiver in relation to its parent object, if applicable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/camediatiming/1427654-begintime?language=objc
func (m_ MediaTimingWrapper) SetBeginTime(value corefoundation.TimeInterval) {
	objc.Call[objc.Void](m_, objc.Sel("setBeginTime:"), value)
}

func (m_ MediaTimingWrapper) HasBeginTime() bool {
	return m_.RespondsToSelector(objc.Sel("beginTime"))
}

// Specifies the begin time of the receiver in relation to its parent object, if applicable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/camediatiming/1427654-begintime?language=objc
func (m_ MediaTimingWrapper) BeginTime() corefoundation.TimeInterval {
	rv := objc.Call[corefoundation.TimeInterval](m_, objc.Sel("beginTime"))
	return rv
}

func (m_ MediaTimingWrapper) HasSetDuration() bool {
	return m_.RespondsToSelector(objc.Sel("setDuration:"))
}

// Specifies the basic duration of the animation, in seconds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/camediatiming/1427652-duration?language=objc
func (m_ MediaTimingWrapper) SetDuration(value corefoundation.TimeInterval) {
	objc.Call[objc.Void](m_, objc.Sel("setDuration:"), value)
}

func (m_ MediaTimingWrapper) HasDuration() bool {
	return m_.RespondsToSelector(objc.Sel("duration"))
}

// Specifies the basic duration of the animation, in seconds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/camediatiming/1427652-duration?language=objc
func (m_ MediaTimingWrapper) Duration() corefoundation.TimeInterval {
	rv := objc.Call[corefoundation.TimeInterval](m_, objc.Sel("duration"))
	return rv
}
