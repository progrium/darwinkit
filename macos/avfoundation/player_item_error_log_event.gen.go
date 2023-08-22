// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PlayerItemErrorLogEvent] class.
var PlayerItemErrorLogEventClass = _PlayerItemErrorLogEventClass{objc.GetClass("AVPlayerItemErrorLogEvent")}

type _PlayerItemErrorLogEventClass struct {
	objc.Class
}

// An interface definition for the [PlayerItemErrorLogEvent] class.
type IPlayerItemErrorLogEvent interface {
	objc.IObject
	Date() foundation.Date
	ErrorStatusCode() int
	URI() string
	ServerAddress() string
	ErrorComment() string
	PlaybackSessionID() string
	ErrorDomain() string
}

// A single item in a player item’s error log. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemerrorlogevent?language=objc
type PlayerItemErrorLogEvent struct {
	objc.Object
}

func PlayerItemErrorLogEventFrom(ptr unsafe.Pointer) PlayerItemErrorLogEvent {
	return PlayerItemErrorLogEvent{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PlayerItemErrorLogEventClass) Alloc() PlayerItemErrorLogEvent {
	rv := objc.Call[PlayerItemErrorLogEvent](pc, objc.Sel("alloc"))
	return rv
}

func PlayerItemErrorLogEvent_Alloc() PlayerItemErrorLogEvent {
	return PlayerItemErrorLogEventClass.Alloc()
}

func (pc _PlayerItemErrorLogEventClass) New() PlayerItemErrorLogEvent {
	rv := objc.Call[PlayerItemErrorLogEvent](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPlayerItemErrorLogEvent() PlayerItemErrorLogEvent {
	return PlayerItemErrorLogEventClass.New()
}

func (p_ PlayerItemErrorLogEvent) Init() PlayerItemErrorLogEvent {
	rv := objc.Call[PlayerItemErrorLogEvent](p_, objc.Sel("init"))
	return rv
}

// The date and time when the error occurred. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemerrorlogevent/1388416-date?language=objc
func (p_ PlayerItemErrorLogEvent) Date() foundation.Date {
	rv := objc.Call[foundation.Date](p_, objc.Sel("date"))
	return rv
}

// A unique error code identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemerrorlogevent/1387875-errorstatuscode?language=objc
func (p_ PlayerItemErrorLogEvent) ErrorStatusCode() int {
	rv := objc.Call[int](p_, objc.Sel("errorStatusCode"))
	return rv
}

// The URI of the playback item that had an error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemerrorlogevent/1389302-uri?language=objc
func (p_ PlayerItemErrorLogEvent) URI() string {
	rv := objc.Call[string](p_, objc.Sel("URI"))
	return rv
}

// The IP address of the server that was the source of the error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemerrorlogevent/1385797-serveraddress?language=objc
func (p_ PlayerItemErrorLogEvent) ServerAddress() string {
	rv := objc.Call[string](p_, objc.Sel("serverAddress"))
	return rv
}

// A description of the error encountered [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemerrorlogevent/1388011-errorcomment?language=objc
func (p_ PlayerItemErrorLogEvent) ErrorComment() string {
	rv := objc.Call[string](p_, objc.Sel("errorComment"))
	return rv
}

// A GUID that identifies the playback session that had an error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemerrorlogevent/1385934-playbacksessionid?language=objc
func (p_ PlayerItemErrorLogEvent) PlaybackSessionID() string {
	rv := objc.Call[string](p_, objc.Sel("playbackSessionID"))
	return rv
}

// The domain of the error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemerrorlogevent/1388603-errordomain?language=objc
func (p_ PlayerItemErrorLogEvent) ErrorDomain() string {
	rv := objc.Call[string](p_, objc.Sel("errorDomain"))
	return rv
}
