// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMProgressEvent] class.
var DOMProgressEventClass = _DOMProgressEventClass{objc.GetClass("DOMProgressEvent")}

type _DOMProgressEventClass struct {
	objc.Class
}

// An interface definition for the [DOMProgressEvent] class.
type IDOMProgressEvent interface {
	IDOMEvent
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domprogressevent?language=objc
type DOMProgressEvent struct {
	DOMEvent
}

func DOMProgressEventFrom(ptr unsafe.Pointer) DOMProgressEvent {
	return DOMProgressEvent{
		DOMEvent: DOMEventFrom(ptr),
	}
}

func (dc _DOMProgressEventClass) Alloc() DOMProgressEvent {
	rv := objc.Call[DOMProgressEvent](dc, objc.Sel("alloc"))
	return rv
}

func DOMProgressEvent_Alloc() DOMProgressEvent {
	return DOMProgressEventClass.Alloc()
}

func (dc _DOMProgressEventClass) New() DOMProgressEvent {
	rv := objc.Call[DOMProgressEvent](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMProgressEvent() DOMProgressEvent {
	return DOMProgressEventClass.New()
}

func (d_ DOMProgressEvent) Init() DOMProgressEvent {
	rv := objc.Call[DOMProgressEvent](d_, objc.Sel("init"))
	return rv
}
