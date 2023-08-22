// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PlayerLooper] class.
var PlayerLooperClass = _PlayerLooperClass{objc.GetClass("AVPlayerLooper")}

type _PlayerLooperClass struct {
	objc.Class
}

// An interface definition for the [PlayerLooper] class.
type IPlayerLooper interface {
	objc.IObject
	DisableLooping()
	Error() foundation.Error
	LoopCount() int
	Status() PlayerLooperStatus
	LoopingPlayerItems() []PlayerItem
}

// An object that loops media content using a queue player. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerlooper?language=objc
type PlayerLooper struct {
	objc.Object
}

func PlayerLooperFrom(ptr unsafe.Pointer) PlayerLooper {
	return PlayerLooper{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PlayerLooperClass) PlayerLooperWithPlayerTemplateItem(player IQueuePlayer, itemToLoop IPlayerItem) PlayerLooper {
	rv := objc.Call[PlayerLooper](pc, objc.Sel("playerLooperWithPlayer:templateItem:"), objc.Ptr(player), objc.Ptr(itemToLoop))
	return rv
}

// Creates a player looper that continuously plays the full duration of a player item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerlooper/1643625-playerlooperwithplayer?language=objc
func PlayerLooper_PlayerLooperWithPlayerTemplateItem(player IQueuePlayer, itemToLoop IPlayerItem) PlayerLooper {
	return PlayerLooperClass.PlayerLooperWithPlayerTemplateItem(player, itemToLoop)
}

func (p_ PlayerLooper) InitWithPlayerTemplateItemTimeRange(player IQueuePlayer, itemToLoop IPlayerItem, loopRange coremedia.TimeRange) PlayerLooper {
	rv := objc.Call[PlayerLooper](p_, objc.Sel("initWithPlayer:templateItem:timeRange:"), objc.Ptr(player), objc.Ptr(itemToLoop), loopRange)
	return rv
}

// Creates a player looper that continuously plays the specified time range of a player item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerlooper/1643626-initwithplayer?language=objc
func NewPlayerLooperWithPlayerTemplateItemTimeRange(player IQueuePlayer, itemToLoop IPlayerItem, loopRange coremedia.TimeRange) PlayerLooper {
	instance := PlayerLooperClass.Alloc().InitWithPlayerTemplateItemTimeRange(player, itemToLoop, loopRange)
	instance.Autorelease()
	return instance
}

func (pc _PlayerLooperClass) Alloc() PlayerLooper {
	rv := objc.Call[PlayerLooper](pc, objc.Sel("alloc"))
	return rv
}

func PlayerLooper_Alloc() PlayerLooper {
	return PlayerLooperClass.Alloc()
}

func (pc _PlayerLooperClass) New() PlayerLooper {
	rv := objc.Call[PlayerLooper](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPlayerLooper() PlayerLooper {
	return PlayerLooperClass.New()
}

func (p_ PlayerLooper) Init() PlayerLooper {
	rv := objc.Call[PlayerLooper](p_, objc.Sel("init"))
	return rv
}

// Disables looping for the player queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerlooper/1643629-disablelooping?language=objc
func (p_ PlayerLooper) DisableLooping() {
	objc.Call[objc.Void](p_, objc.Sel("disableLooping"))
}

// An error that describes the reason looping failed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerlooper/2177064-error?language=objc
func (p_ PlayerLooper) Error() foundation.Error {
	rv := objc.Call[foundation.Error](p_, objc.Sel("error"))
	return rv
}

// The number of times the object played the media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerlooper/1643648-loopcount?language=objc
func (p_ PlayerLooper) LoopCount() int {
	rv := objc.Call[int](p_, objc.Sel("loopCount"))
	return rv
}

// A status that indicates the object’s ability to loop playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerlooper/2177060-status?language=objc
func (p_ PlayerLooper) Status() PlayerLooperStatus {
	rv := objc.Call[PlayerLooperStatus](p_, objc.Sel("status"))
	return rv
}

// An array containing replicas of the template player item used to accomplish the looping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerlooper/1643631-loopingplayeritems?language=objc
func (p_ PlayerLooper) LoopingPlayerItems() []PlayerItem {
	rv := objc.Call[[]PlayerItem](p_, objc.Sel("loopingPlayerItems"))
	return rv
}
