// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ChangeHistoryUpdateContactEvent] class.
var ChangeHistoryUpdateContactEventClass = _ChangeHistoryUpdateContactEventClass{objc.GetClass("CNChangeHistoryUpdateContactEvent")}

type _ChangeHistoryUpdateContactEventClass struct {
	objc.Class
}

// An interface definition for the [ChangeHistoryUpdateContactEvent] class.
type IChangeHistoryUpdateContactEvent interface {
	IChangeHistoryEvent
	Contact() Contact
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryupdatecontactevent?language=objc
type ChangeHistoryUpdateContactEvent struct {
	ChangeHistoryEvent
}

func ChangeHistoryUpdateContactEventFrom(ptr unsafe.Pointer) ChangeHistoryUpdateContactEvent {
	return ChangeHistoryUpdateContactEvent{
		ChangeHistoryEvent: ChangeHistoryEventFrom(ptr),
	}
}

func (cc _ChangeHistoryUpdateContactEventClass) Alloc() ChangeHistoryUpdateContactEvent {
	rv := objc.Call[ChangeHistoryUpdateContactEvent](cc, objc.Sel("alloc"))
	return rv
}

func ChangeHistoryUpdateContactEvent_Alloc() ChangeHistoryUpdateContactEvent {
	return ChangeHistoryUpdateContactEventClass.Alloc()
}

func (cc _ChangeHistoryUpdateContactEventClass) New() ChangeHistoryUpdateContactEvent {
	rv := objc.Call[ChangeHistoryUpdateContactEvent](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewChangeHistoryUpdateContactEvent() ChangeHistoryUpdateContactEvent {
	return ChangeHistoryUpdateContactEventClass.New()
}

func (c_ ChangeHistoryUpdateContactEvent) Init() ChangeHistoryUpdateContactEvent {
	rv := objc.Call[ChangeHistoryUpdateContactEvent](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryupdatecontactevent/3113287-contact?language=objc
func (c_ ChangeHistoryUpdateContactEvent) Contact() Contact {
	rv := objc.Call[Contact](c_, objc.Sel("contact"))
	return rv
}
