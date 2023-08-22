// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PlayerItemErrorLog] class.
var PlayerItemErrorLogClass = _PlayerItemErrorLogClass{objc.GetClass("AVPlayerItemErrorLog")}

type _PlayerItemErrorLogClass struct {
	objc.Class
}

// An interface definition for the [PlayerItemErrorLog] class.
type IPlayerItemErrorLog interface {
	objc.IObject
	ExtendedLogData() []byte
	Events() []PlayerItemErrorLogEvent
	ExtendedLogDataStringEncoding() foundation.StringEncoding
}

// The error log associated with a player item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemerrorlog?language=objc
type PlayerItemErrorLog struct {
	objc.Object
}

func PlayerItemErrorLogFrom(ptr unsafe.Pointer) PlayerItemErrorLog {
	return PlayerItemErrorLog{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PlayerItemErrorLogClass) Alloc() PlayerItemErrorLog {
	rv := objc.Call[PlayerItemErrorLog](pc, objc.Sel("alloc"))
	return rv
}

func PlayerItemErrorLog_Alloc() PlayerItemErrorLog {
	return PlayerItemErrorLogClass.Alloc()
}

func (pc _PlayerItemErrorLogClass) New() PlayerItemErrorLog {
	rv := objc.Call[PlayerItemErrorLog](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPlayerItemErrorLog() PlayerItemErrorLog {
	return PlayerItemErrorLogClass.New()
}

func (p_ PlayerItemErrorLog) Init() PlayerItemErrorLog {
	rv := objc.Call[PlayerItemErrorLog](p_, objc.Sel("init"))
	return rv
}

// Returns a serialized representation of the error log in the Extended Log File Format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemerrorlog/1389100-extendedlogdata?language=objc
func (p_ PlayerItemErrorLog) ExtendedLogData() []byte {
	rv := objc.Call[[]byte](p_, objc.Sel("extendedLogData"))
	return rv
}

// A chronologically ordered array of player item error log event objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemerrorlog/1387637-events?language=objc
func (p_ PlayerItemErrorLog) Events() []PlayerItemErrorLogEvent {
	rv := objc.Call[[]PlayerItemErrorLogEvent](p_, objc.Sel("events"))
	return rv
}

// The string encoding of the extended log data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemerrorlog/1387271-extendedlogdatastringencoding?language=objc
func (p_ PlayerItemErrorLog) ExtendedLogDataStringEncoding() foundation.StringEncoding {
	rv := objc.Call[foundation.StringEncoding](p_, objc.Sel("extendedLogDataStringEncoding"))
	return rv
}
