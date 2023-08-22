// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PlayerItemAccessLog] class.
var PlayerItemAccessLogClass = _PlayerItemAccessLogClass{objc.GetClass("AVPlayerItemAccessLog")}

type _PlayerItemAccessLogClass struct {
	objc.Class
}

// An interface definition for the [PlayerItemAccessLog] class.
type IPlayerItemAccessLog interface {
	objc.IObject
	ExtendedLogData() []byte
	Events() []PlayerItemAccessLogEvent
	ExtendedLogDataStringEncoding() foundation.StringEncoding
}

// An object used to retrieve the access log associated with a player item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslog?language=objc
type PlayerItemAccessLog struct {
	objc.Object
}

func PlayerItemAccessLogFrom(ptr unsafe.Pointer) PlayerItemAccessLog {
	return PlayerItemAccessLog{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PlayerItemAccessLogClass) Alloc() PlayerItemAccessLog {
	rv := objc.Call[PlayerItemAccessLog](pc, objc.Sel("alloc"))
	return rv
}

func PlayerItemAccessLog_Alloc() PlayerItemAccessLog {
	return PlayerItemAccessLogClass.Alloc()
}

func (pc _PlayerItemAccessLogClass) New() PlayerItemAccessLog {
	rv := objc.Call[PlayerItemAccessLog](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPlayerItemAccessLog() PlayerItemAccessLog {
	return PlayerItemAccessLogClass.New()
}

func (p_ PlayerItemAccessLog) Init() PlayerItemAccessLog {
	rv := objc.Call[PlayerItemAccessLog](p_, objc.Sel("init"))
	return rv
}

// Returns a serialized representation of the access log in the Extended Log File Format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslog/1386892-extendedlogdata?language=objc
func (p_ PlayerItemAccessLog) ExtendedLogData() []byte {
	rv := objc.Call[[]byte](p_, objc.Sel("extendedLogData"))
	return rv
}

// A chronologically ordered array of player item access log events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslog/1387406-events?language=objc
func (p_ PlayerItemAccessLog) Events() []PlayerItemAccessLogEvent {
	rv := objc.Call[[]PlayerItemAccessLogEvent](p_, objc.Sel("events"))
	return rv
}

// The string encoding of the extended log data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslog/1390863-extendedlogdatastringencoding?language=objc
func (p_ PlayerItemAccessLog) ExtendedLogDataStringEncoding() foundation.StringEncoding {
	rv := objc.Call[foundation.StringEncoding](p_, objc.Sel("extendedLogDataStringEncoding"))
	return rv
}
