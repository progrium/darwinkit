// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ChangeHistoryDeleteGroupEvent] class.
var ChangeHistoryDeleteGroupEventClass = _ChangeHistoryDeleteGroupEventClass{objc.GetClass("CNChangeHistoryDeleteGroupEvent")}

type _ChangeHistoryDeleteGroupEventClass struct {
	objc.Class
}

// An interface definition for the [ChangeHistoryDeleteGroupEvent] class.
type IChangeHistoryDeleteGroupEvent interface {
	IChangeHistoryEvent
	GroupIdentifier() string
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistorydeletegroupevent?language=objc
type ChangeHistoryDeleteGroupEvent struct {
	ChangeHistoryEvent
}

func ChangeHistoryDeleteGroupEventFrom(ptr unsafe.Pointer) ChangeHistoryDeleteGroupEvent {
	return ChangeHistoryDeleteGroupEvent{
		ChangeHistoryEvent: ChangeHistoryEventFrom(ptr),
	}
}

func (cc _ChangeHistoryDeleteGroupEventClass) Alloc() ChangeHistoryDeleteGroupEvent {
	rv := objc.Call[ChangeHistoryDeleteGroupEvent](cc, objc.Sel("alloc"))
	return rv
}

func ChangeHistoryDeleteGroupEvent_Alloc() ChangeHistoryDeleteGroupEvent {
	return ChangeHistoryDeleteGroupEventClass.Alloc()
}

func (cc _ChangeHistoryDeleteGroupEventClass) New() ChangeHistoryDeleteGroupEvent {
	rv := objc.Call[ChangeHistoryDeleteGroupEvent](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewChangeHistoryDeleteGroupEvent() ChangeHistoryDeleteGroupEvent {
	return ChangeHistoryDeleteGroupEventClass.New()
}

func (c_ ChangeHistoryDeleteGroupEvent) Init() ChangeHistoryDeleteGroupEvent {
	rv := objc.Call[ChangeHistoryDeleteGroupEvent](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistorydeletegroupevent/3113264-groupidentifier?language=objc
func (c_ ChangeHistoryDeleteGroupEvent) GroupIdentifier() string {
	rv := objc.Call[string](c_, objc.Sel("groupIdentifier"))
	return rv
}
