// AUTO-GENERATED CODE, DO NOT MODIFY

package mediaplayer

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RatingCommand] class.
var RatingCommandClass = _RatingCommandClass{objc.GetClass("MPRatingCommand")}

type _RatingCommandClass struct {
	objc.Class
}

// An interface definition for the [RatingCommand] class.
type IRatingCommand interface {
	IRemoteCommand
	MinimumRating() float64
	SetMinimumRating(value float64)
	MaximumRating() float64
	SetMaximumRating(value float64)
}

// An object that provides a detailed rating for the playing item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpratingcommand?language=objc
type RatingCommand struct {
	RemoteCommand
}

func RatingCommandFrom(ptr unsafe.Pointer) RatingCommand {
	return RatingCommand{
		RemoteCommand: RemoteCommandFrom(ptr),
	}
}

func (rc _RatingCommandClass) Alloc() RatingCommand {
	rv := objc.Call[RatingCommand](rc, objc.Sel("alloc"))
	return rv
}

func RatingCommand_Alloc() RatingCommand {
	return RatingCommandClass.Alloc()
}

func (rc _RatingCommandClass) New() RatingCommand {
	rv := objc.Call[RatingCommand](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRatingCommand() RatingCommand {
	return RatingCommandClass.New()
}

func (r_ RatingCommand) Init() RatingCommand {
	rv := objc.Call[RatingCommand](r_, objc.Sel("init"))
	return rv
}

// The minimum rating for a command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpratingcommand/1622902-minimumrating?language=objc
func (r_ RatingCommand) MinimumRating() float64 {
	rv := objc.Call[float64](r_, objc.Sel("minimumRating"))
	return rv
}

// The minimum rating for a command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpratingcommand/1622902-minimumrating?language=objc
func (r_ RatingCommand) SetMinimumRating(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setMinimumRating:"), value)
}

// The maximum rating for a command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpratingcommand/1622898-maximumrating?language=objc
func (r_ RatingCommand) MaximumRating() float64 {
	rv := objc.Call[float64](r_, objc.Sel("maximumRating"))
	return rv
}

// The maximum rating for a command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpratingcommand/1622898-maximumrating?language=objc
func (r_ RatingCommand) SetMaximumRating(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setMaximumRating:"), value)
}
