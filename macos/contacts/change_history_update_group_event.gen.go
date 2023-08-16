// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ChangeHistoryUpdateGroupEvent] class.
var ChangeHistoryUpdateGroupEventClass = _ChangeHistoryUpdateGroupEventClass{objc.GetClass("CNChangeHistoryUpdateGroupEvent")}

type _ChangeHistoryUpdateGroupEventClass struct {
	objc.Class
}

// An interface definition for the [ChangeHistoryUpdateGroupEvent] class.
type IChangeHistoryUpdateGroupEvent interface {
	IChangeHistoryEvent
	Group() Group
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryupdategroupevent?language=objc
type ChangeHistoryUpdateGroupEvent struct {
	ChangeHistoryEvent
}

func ChangeHistoryUpdateGroupEventFrom(ptr unsafe.Pointer) ChangeHistoryUpdateGroupEvent {
	return ChangeHistoryUpdateGroupEvent{
		ChangeHistoryEvent: ChangeHistoryEventFrom(ptr),
	}
}

func (cc _ChangeHistoryUpdateGroupEventClass) Alloc() ChangeHistoryUpdateGroupEvent {
	rv := objc.Call[ChangeHistoryUpdateGroupEvent](cc, objc.Sel("alloc"))
	return rv
}

func ChangeHistoryUpdateGroupEvent_Alloc() ChangeHistoryUpdateGroupEvent {
	return ChangeHistoryUpdateGroupEventClass.Alloc()
}

func (cc _ChangeHistoryUpdateGroupEventClass) New() ChangeHistoryUpdateGroupEvent {
	rv := objc.Call[ChangeHistoryUpdateGroupEvent](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewChangeHistoryUpdateGroupEvent() ChangeHistoryUpdateGroupEvent {
	return ChangeHistoryUpdateGroupEventClass.New()
}

func (c_ ChangeHistoryUpdateGroupEvent) Init() ChangeHistoryUpdateGroupEvent {
	rv := objc.Call[ChangeHistoryUpdateGroupEvent](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryupdategroupevent/3113289-group?language=objc
func (c_ ChangeHistoryUpdateGroupEvent) Group() Group {
	rv := objc.Call[Group](c_, objc.Sel("group"))
	return rv
}
