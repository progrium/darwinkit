// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMMouseEvent] class.
var DOMMouseEventClass = _DOMMouseEventClass{objc.GetClass("DOMMouseEvent")}

type _DOMMouseEventClass struct {
	objc.Class
}

// An interface definition for the [DOMMouseEvent] class.
type IDOMMouseEvent interface {
	IDOMUIEvent
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/dommouseevent?language=objc
type DOMMouseEvent struct {
	DOMUIEvent
}

func DOMMouseEventFrom(ptr unsafe.Pointer) DOMMouseEvent {
	return DOMMouseEvent{
		DOMUIEvent: DOMUIEventFrom(ptr),
	}
}

func (dc _DOMMouseEventClass) Alloc() DOMMouseEvent {
	rv := objc.Call[DOMMouseEvent](dc, objc.Sel("alloc"))
	return rv
}

func DOMMouseEvent_Alloc() DOMMouseEvent {
	return DOMMouseEventClass.Alloc()
}

func (dc _DOMMouseEventClass) New() DOMMouseEvent {
	rv := objc.Call[DOMMouseEvent](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMMouseEvent() DOMMouseEvent {
	return DOMMouseEventClass.New()
}

func (d_ DOMMouseEvent) Init() DOMMouseEvent {
	rv := objc.Call[DOMMouseEvent](d_, objc.Sel("init"))
	return rv
}
