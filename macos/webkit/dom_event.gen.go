// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMEvent] class.
var DOMEventClass = _DOMEventClass{objc.GetClass("DOMEvent")}

type _DOMEventClass struct {
	objc.Class
}

// An interface definition for the [DOMEvent] class.
type IDOMEvent interface {
	IDOMObject
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domevent?language=objc
type DOMEvent struct {
	DOMObject
}

func DOMEventFrom(ptr unsafe.Pointer) DOMEvent {
	return DOMEvent{
		DOMObject: DOMObjectFrom(ptr),
	}
}

func (dc _DOMEventClass) Alloc() DOMEvent {
	rv := objc.Call[DOMEvent](dc, objc.Sel("alloc"))
	return rv
}

func DOMEvent_Alloc() DOMEvent {
	return DOMEventClass.Alloc()
}

func (dc _DOMEventClass) New() DOMEvent {
	rv := objc.Call[DOMEvent](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMEvent() DOMEvent {
	return DOMEventClass.New()
}

func (d_ DOMEvent) Init() DOMEvent {
	rv := objc.Call[DOMEvent](d_, objc.Sel("init"))
	return rv
}
