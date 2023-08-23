// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/corefoundation"
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/corevideo"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PlayerItemOutput] class.
var PlayerItemOutputClass = _PlayerItemOutputClass{objc.GetClass("AVPlayerItemOutput")}

type _PlayerItemOutputClass struct {
	objc.Class
}

// An interface definition for the [PlayerItemOutput] class.
type IPlayerItemOutput interface {
	objc.IObject
	ItemTimeForCVTimeStamp(timestamp corevideo.TimeStamp) coremedia.Time
	ItemTimeForHostTime(hostTimeInSeconds corefoundation.TimeInterval) coremedia.Time
	ItemTimeForMachAbsoluteTime(machAbsoluteTime int64) coremedia.Time
	SuppressesPlayerRendering() bool
	SetSuppressesPlayerRendering(value bool)
}

// An abstract class that defines the common interface to output media data from a player item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemoutput?language=objc
type PlayerItemOutput struct {
	objc.Object
}

func PlayerItemOutputFrom(ptr unsafe.Pointer) PlayerItemOutput {
	return PlayerItemOutput{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PlayerItemOutputClass) Alloc() PlayerItemOutput {
	rv := objc.Call[PlayerItemOutput](pc, objc.Sel("alloc"))
	return rv
}

func PlayerItemOutput_Alloc() PlayerItemOutput {
	return PlayerItemOutputClass.Alloc()
}

func (pc _PlayerItemOutputClass) New() PlayerItemOutput {
	rv := objc.Call[PlayerItemOutput](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPlayerItemOutput() PlayerItemOutput {
	return PlayerItemOutputClass.New()
}

func (p_ PlayerItemOutput) Init() PlayerItemOutput {
	rv := objc.Call[PlayerItemOutput](p_, objc.Sel("init"))
	return rv
}

// Converts a Core Video timestamp to the item’s timebase. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemoutput/1388333-itemtimeforcvtimestamp?language=objc
func (p_ PlayerItemOutput) ItemTimeForCVTimeStamp(timestamp corevideo.TimeStamp) coremedia.Time {
	rv := objc.Call[coremedia.Time](p_, objc.Sel("itemTimeForCVTimeStamp:"), timestamp)
	return rv
}

// Converts a host time, specified in seconds, to the item’s timebase. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemoutput/1386538-itemtimeforhosttime?language=objc
func (p_ PlayerItemOutput) ItemTimeForHostTime(hostTimeInSeconds corefoundation.TimeInterval) coremedia.Time {
	rv := objc.Call[coremedia.Time](p_, objc.Sel("itemTimeForHostTime:"), hostTimeInSeconds)
	return rv
}

// Converts a Mach host time to the item’s timebase. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemoutput/1386962-itemtimeformachabsolutetime?language=objc
func (p_ PlayerItemOutput) ItemTimeForMachAbsoluteTime(machAbsoluteTime int64) coremedia.Time {
	rv := objc.Call[coremedia.Time](p_, objc.Sel("itemTimeForMachAbsoluteTime:"), machAbsoluteTime)
	return rv
}

// A Boolean value that indicates whether the player object renders the receiver’s output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemoutput/1386133-suppressesplayerrendering?language=objc
func (p_ PlayerItemOutput) SuppressesPlayerRendering() bool {
	rv := objc.Call[bool](p_, objc.Sel("suppressesPlayerRendering"))
	return rv
}

// A Boolean value that indicates whether the player object renders the receiver’s output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemoutput/1386133-suppressesplayerrendering?language=objc
func (p_ PlayerItemOutput) SetSuppressesPlayerRendering(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setSuppressesPlayerRendering:"), value)
}
