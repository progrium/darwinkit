// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMMutationEvent] class.
var DOMMutationEventClass = _DOMMutationEventClass{objc.GetClass("DOMMutationEvent")}

type _DOMMutationEventClass struct {
	objc.Class
}

// An interface definition for the [DOMMutationEvent] class.
type IDOMMutationEvent interface {
	IDOMEvent
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/dommutationevent?language=objc
type DOMMutationEvent struct {
	DOMEvent
}

func DOMMutationEventFrom(ptr unsafe.Pointer) DOMMutationEvent {
	return DOMMutationEvent{
		DOMEvent: DOMEventFrom(ptr),
	}
}

func (dc _DOMMutationEventClass) Alloc() DOMMutationEvent {
	rv := objc.Call[DOMMutationEvent](dc, objc.Sel("alloc"))
	return rv
}

func DOMMutationEvent_Alloc() DOMMutationEvent {
	return DOMMutationEventClass.Alloc()
}

func (dc _DOMMutationEventClass) New() DOMMutationEvent {
	rv := objc.Call[DOMMutationEvent](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMMutationEvent() DOMMutationEvent {
	return DOMMutationEventClass.New()
}

func (d_ DOMMutationEvent) Init() DOMMutationEvent {
	rv := objc.Call[DOMMutationEvent](d_, objc.Sel("init"))
	return rv
}
