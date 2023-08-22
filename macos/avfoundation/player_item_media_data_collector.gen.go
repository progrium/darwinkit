// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PlayerItemMediaDataCollector] class.
var PlayerItemMediaDataCollectorClass = _PlayerItemMediaDataCollectorClass{objc.GetClass("AVPlayerItemMediaDataCollector")}

type _PlayerItemMediaDataCollectorClass struct {
	objc.Class
}

// An interface definition for the [PlayerItemMediaDataCollector] class.
type IPlayerItemMediaDataCollector interface {
	objc.IObject
}

// The abstract base for media data collectors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemmediadatacollector?language=objc
type PlayerItemMediaDataCollector struct {
	objc.Object
}

func PlayerItemMediaDataCollectorFrom(ptr unsafe.Pointer) PlayerItemMediaDataCollector {
	return PlayerItemMediaDataCollector{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PlayerItemMediaDataCollectorClass) Alloc() PlayerItemMediaDataCollector {
	rv := objc.Call[PlayerItemMediaDataCollector](pc, objc.Sel("alloc"))
	return rv
}

func PlayerItemMediaDataCollector_Alloc() PlayerItemMediaDataCollector {
	return PlayerItemMediaDataCollectorClass.Alloc()
}

func (pc _PlayerItemMediaDataCollectorClass) New() PlayerItemMediaDataCollector {
	rv := objc.Call[PlayerItemMediaDataCollector](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPlayerItemMediaDataCollector() PlayerItemMediaDataCollector {
	return PlayerItemMediaDataCollectorClass.New()
}

func (p_ PlayerItemMediaDataCollector) Init() PlayerItemMediaDataCollector {
	rv := objc.Call[PlayerItemMediaDataCollector](p_, objc.Sel("init"))
	return rv
}
