// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMOverflowEvent] class.
var DOMOverflowEventClass = _DOMOverflowEventClass{objc.GetClass("DOMOverflowEvent")}

type _DOMOverflowEventClass struct {
	objc.Class
}

// An interface definition for the [DOMOverflowEvent] class.
type IDOMOverflowEvent interface {
	IDOMEvent
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domoverflowevent?language=objc
type DOMOverflowEvent struct {
	DOMEvent
}

func DOMOverflowEventFrom(ptr unsafe.Pointer) DOMOverflowEvent {
	return DOMOverflowEvent{
		DOMEvent: DOMEventFrom(ptr),
	}
}

func (dc _DOMOverflowEventClass) Alloc() DOMOverflowEvent {
	rv := objc.Call[DOMOverflowEvent](dc, objc.Sel("alloc"))
	return rv
}

func DOMOverflowEvent_Alloc() DOMOverflowEvent {
	return DOMOverflowEventClass.Alloc()
}

func (dc _DOMOverflowEventClass) New() DOMOverflowEvent {
	rv := objc.Call[DOMOverflowEvent](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMOverflowEvent() DOMOverflowEvent {
	return DOMOverflowEventClass.New()
}

func (d_ DOMOverflowEvent) Init() DOMOverflowEvent {
	rv := objc.Call[DOMOverflowEvent](d_, objc.Sel("init"))
	return rv
}
