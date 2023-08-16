// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ChangeHistoryAddContactEvent] class.
var ChangeHistoryAddContactEventClass = _ChangeHistoryAddContactEventClass{objc.GetClass("CNChangeHistoryAddContactEvent")}

type _ChangeHistoryAddContactEventClass struct {
	objc.Class
}

// An interface definition for the [ChangeHistoryAddContactEvent] class.
type IChangeHistoryAddContactEvent interface {
	IChangeHistoryEvent
	Contact() Contact
	ContainerIdentifier() string
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryaddcontactevent?language=objc
type ChangeHistoryAddContactEvent struct {
	ChangeHistoryEvent
}

func ChangeHistoryAddContactEventFrom(ptr unsafe.Pointer) ChangeHistoryAddContactEvent {
	return ChangeHistoryAddContactEvent{
		ChangeHistoryEvent: ChangeHistoryEventFrom(ptr),
	}
}

func (cc _ChangeHistoryAddContactEventClass) Alloc() ChangeHistoryAddContactEvent {
	rv := objc.Call[ChangeHistoryAddContactEvent](cc, objc.Sel("alloc"))
	return rv
}

func ChangeHistoryAddContactEvent_Alloc() ChangeHistoryAddContactEvent {
	return ChangeHistoryAddContactEventClass.Alloc()
}

func (cc _ChangeHistoryAddContactEventClass) New() ChangeHistoryAddContactEvent {
	rv := objc.Call[ChangeHistoryAddContactEvent](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewChangeHistoryAddContactEvent() ChangeHistoryAddContactEvent {
	return ChangeHistoryAddContactEventClass.New()
}

func (c_ ChangeHistoryAddContactEvent) Init() ChangeHistoryAddContactEvent {
	rv := objc.Call[ChangeHistoryAddContactEvent](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryaddcontactevent/3113250-contact?language=objc
func (c_ ChangeHistoryAddContactEvent) Contact() Contact {
	rv := objc.Call[Contact](c_, objc.Sel("contact"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryaddcontactevent/3113251-containeridentifier?language=objc
func (c_ ChangeHistoryAddContactEvent) ContainerIdentifier() string {
	rv := objc.Call[string](c_, objc.Sel("containerIdentifier"))
	return rv
}
