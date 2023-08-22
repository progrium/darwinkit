// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CoordinatedPlaybackSuspension] class.
var CoordinatedPlaybackSuspensionClass = _CoordinatedPlaybackSuspensionClass{objc.GetClass("AVCoordinatedPlaybackSuspension")}

type _CoordinatedPlaybackSuspensionClass struct {
	objc.Class
}

// An interface definition for the [CoordinatedPlaybackSuspension] class.
type ICoordinatedPlaybackSuspension interface {
	objc.IObject
	End()
	EndProposingNewTime(time coremedia.Time)
	BeginDate() foundation.Date
	Reason() CoordinatedPlaybackSuspensionReason
}

// An object that represents a temporary suspension of coordinated playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcoordinatedplaybacksuspension?language=objc
type CoordinatedPlaybackSuspension struct {
	objc.Object
}

func CoordinatedPlaybackSuspensionFrom(ptr unsafe.Pointer) CoordinatedPlaybackSuspension {
	return CoordinatedPlaybackSuspension{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CoordinatedPlaybackSuspensionClass) Alloc() CoordinatedPlaybackSuspension {
	rv := objc.Call[CoordinatedPlaybackSuspension](cc, objc.Sel("alloc"))
	return rv
}

func CoordinatedPlaybackSuspension_Alloc() CoordinatedPlaybackSuspension {
	return CoordinatedPlaybackSuspensionClass.Alloc()
}

func (cc _CoordinatedPlaybackSuspensionClass) New() CoordinatedPlaybackSuspension {
	rv := objc.Call[CoordinatedPlaybackSuspension](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCoordinatedPlaybackSuspension() CoordinatedPlaybackSuspension {
	return CoordinatedPlaybackSuspensionClass.New()
}

func (c_ CoordinatedPlaybackSuspension) Init() CoordinatedPlaybackSuspension {
	rv := objc.Call[CoordinatedPlaybackSuspension](c_, objc.Sel("init"))
	return rv
}

// Ends a suspension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcoordinatedplaybacksuspension/3750240-end?language=objc
func (c_ CoordinatedPlaybackSuspension) End() {
	objc.Call[objc.Void](c_, objc.Sel("end"))
}

// Ends a suspension and proposes a new playback time to the group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcoordinatedplaybacksuspension/3750241-endproposingnewtime?language=objc
func (c_ CoordinatedPlaybackSuspension) EndProposingNewTime(time coremedia.Time) {
	objc.Call[objc.Void](c_, objc.Sel("endProposingNewTime:"), time)
}

// The time the suspension begins. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcoordinatedplaybacksuspension/3750239-begindate?language=objc
func (c_ CoordinatedPlaybackSuspension) BeginDate() foundation.Date {
	rv := objc.Call[foundation.Date](c_, objc.Sel("beginDate"))
	return rv
}

// The reason for the suspension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcoordinatedplaybacksuspension/3750242-reason?language=objc
func (c_ CoordinatedPlaybackSuspension) Reason() CoordinatedPlaybackSuspensionReason {
	rv := objc.Call[CoordinatedPlaybackSuspensionReason](c_, objc.Sel("reason"))
	return rv
}
