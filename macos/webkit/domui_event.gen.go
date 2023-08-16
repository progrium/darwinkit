// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMUIEvent] class.
var DOMUIEventClass = _DOMUIEventClass{objc.GetClass("DOMUIEvent")}

type _DOMUIEventClass struct {
	objc.Class
}

// An interface definition for the [DOMUIEvent] class.
type IDOMUIEvent interface {
	IDOMEvent
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domuievent?language=objc
type DOMUIEvent struct {
	DOMEvent
}

func DOMUIEventFrom(ptr unsafe.Pointer) DOMUIEvent {
	return DOMUIEvent{
		DOMEvent: DOMEventFrom(ptr),
	}
}

func (dc _DOMUIEventClass) Alloc() DOMUIEvent {
	rv := objc.Call[DOMUIEvent](dc, objc.Sel("alloc"))
	return rv
}

func DOMUIEvent_Alloc() DOMUIEvent {
	return DOMUIEventClass.Alloc()
}

func (dc _DOMUIEventClass) New() DOMUIEvent {
	rv := objc.Call[DOMUIEvent](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMUIEvent() DOMUIEvent {
	return DOMUIEventClass.New()
}

func (d_ DOMUIEvent) Init() DOMUIEvent {
	rv := objc.Call[DOMUIEvent](d_, objc.Sel("init"))
	return rv
}
