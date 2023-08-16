// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DOMKeyboardEvent] class.
var DOMKeyboardEventClass = _DOMKeyboardEventClass{objc.GetClass("DOMKeyboardEvent")}

type _DOMKeyboardEventClass struct {
	objc.Class
}

// An interface definition for the [DOMKeyboardEvent] class.
type IDOMKeyboardEvent interface {
	IDOMUIEvent
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/domkeyboardevent?language=objc
type DOMKeyboardEvent struct {
	DOMUIEvent
}

func DOMKeyboardEventFrom(ptr unsafe.Pointer) DOMKeyboardEvent {
	return DOMKeyboardEvent{
		DOMUIEvent: DOMUIEventFrom(ptr),
	}
}

func (dc _DOMKeyboardEventClass) Alloc() DOMKeyboardEvent {
	rv := objc.Call[DOMKeyboardEvent](dc, objc.Sel("alloc"))
	return rv
}

func DOMKeyboardEvent_Alloc() DOMKeyboardEvent {
	return DOMKeyboardEventClass.Alloc()
}

func (dc _DOMKeyboardEventClass) New() DOMKeyboardEvent {
	rv := objc.Call[DOMKeyboardEvent](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDOMKeyboardEvent() DOMKeyboardEvent {
	return DOMKeyboardEventClass.New()
}

func (d_ DOMKeyboardEvent) Init() DOMKeyboardEvent {
	rv := objc.Call[DOMKeyboardEvent](d_, objc.Sel("init"))
	return rv
}
