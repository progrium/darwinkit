// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ChangeHistoryDeleteContactEvent] class.
var ChangeHistoryDeleteContactEventClass = _ChangeHistoryDeleteContactEventClass{objc.GetClass("CNChangeHistoryDeleteContactEvent")}

type _ChangeHistoryDeleteContactEventClass struct {
	objc.Class
}

// An interface definition for the [ChangeHistoryDeleteContactEvent] class.
type IChangeHistoryDeleteContactEvent interface {
	IChangeHistoryEvent
	ContactIdentifier() string
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistorydeletecontactevent?language=objc
type ChangeHistoryDeleteContactEvent struct {
	ChangeHistoryEvent
}

func ChangeHistoryDeleteContactEventFrom(ptr unsafe.Pointer) ChangeHistoryDeleteContactEvent {
	return ChangeHistoryDeleteContactEvent{
		ChangeHistoryEvent: ChangeHistoryEventFrom(ptr),
	}
}

func (cc _ChangeHistoryDeleteContactEventClass) Alloc() ChangeHistoryDeleteContactEvent {
	rv := objc.Call[ChangeHistoryDeleteContactEvent](cc, objc.Sel("alloc"))
	return rv
}

func ChangeHistoryDeleteContactEvent_Alloc() ChangeHistoryDeleteContactEvent {
	return ChangeHistoryDeleteContactEventClass.Alloc()
}

func (cc _ChangeHistoryDeleteContactEventClass) New() ChangeHistoryDeleteContactEvent {
	rv := objc.Call[ChangeHistoryDeleteContactEvent](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewChangeHistoryDeleteContactEvent() ChangeHistoryDeleteContactEvent {
	return ChangeHistoryDeleteContactEventClass.New()
}

func (c_ ChangeHistoryDeleteContactEvent) Init() ChangeHistoryDeleteContactEvent {
	rv := objc.Call[ChangeHistoryDeleteContactEvent](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistorydeletecontactevent/3113262-contactidentifier?language=objc
func (c_ ChangeHistoryDeleteContactEvent) ContactIdentifier() string {
	rv := objc.Call[string](c_, objc.Sel("contactIdentifier"))
	return rv
}
