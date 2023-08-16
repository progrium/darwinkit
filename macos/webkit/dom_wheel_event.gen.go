// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMWheelEvent] class.
var DOMWheelEventClass = _DOMWheelEventClass{objc.GetClass("DOMWheelEvent")}

type _DOMWheelEventClass struct {
	objc.Class
}

// An interface definition for the [DOMWheelEvent] class.
type IDOMWheelEvent interface {
	IDOMMouseEvent
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domwheelevent?language=objc
type DOMWheelEvent struct {
	DOMMouseEvent
}

func DOMWheelEventFrom(ptr unsafe.Pointer) DOMWheelEvent {
	return DOMWheelEvent{
		DOMMouseEvent: DOMMouseEventFrom(ptr),
	}
}

func (dc _DOMWheelEventClass) Alloc() DOMWheelEvent {
	rv := objc.Call[DOMWheelEvent](dc, objc.Sel("alloc"))
	return rv
}

func DOMWheelEvent_Alloc() DOMWheelEvent {
	return DOMWheelEventClass.Alloc()
}

func (dc _DOMWheelEventClass) New() DOMWheelEvent {
	rv := objc.Call[DOMWheelEvent](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMWheelEvent() DOMWheelEvent {
	return DOMWheelEventClass.New()
}

func (d_ DOMWheelEvent) Init() DOMWheelEvent {
	rv := objc.Call[DOMWheelEvent](d_, objc.Sel("init"))
	return rv
}
