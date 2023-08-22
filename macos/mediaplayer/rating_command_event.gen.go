// AUTO-GENERATED CODE, DO NOT MODIFY

package mediaplayer

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RatingCommandEvent] class.
var RatingCommandEventClass = _RatingCommandEventClass{objc.GetClass("MPRatingCommandEvent")}

type _RatingCommandEventClass struct {
	objc.Class
}

// An interface definition for the [RatingCommandEvent] class.
type IRatingCommandEvent interface {
	IRemoteCommandEvent
	Rating() float64
}

// An event requesting a change in the rating. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpratingcommandevent?language=objc
type RatingCommandEvent struct {
	RemoteCommandEvent
}

func RatingCommandEventFrom(ptr unsafe.Pointer) RatingCommandEvent {
	return RatingCommandEvent{
		RemoteCommandEvent: RemoteCommandEventFrom(ptr),
	}
}

func (rc _RatingCommandEventClass) Alloc() RatingCommandEvent {
	rv := objc.Call[RatingCommandEvent](rc, objc.Sel("alloc"))
	return rv
}

func RatingCommandEvent_Alloc() RatingCommandEvent {
	return RatingCommandEventClass.Alloc()
}

func (rc _RatingCommandEventClass) New() RatingCommandEvent {
	rv := objc.Call[RatingCommandEvent](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRatingCommandEvent() RatingCommandEvent {
	return RatingCommandEventClass.New()
}

func (r_ RatingCommandEvent) Init() RatingCommandEvent {
	rv := objc.Call[RatingCommandEvent](r_, objc.Sel("init"))
	return rv
}

// The rating for the command event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpratingcommandevent/1616764-rating?language=objc
func (r_ RatingCommandEvent) Rating() float64 {
	rv := objc.Call[float64](r_, objc.Sel("rating"))
	return rv
}
