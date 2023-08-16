// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ChangeHistoryRemoveSubgroupFromGroupEvent] class.
var ChangeHistoryRemoveSubgroupFromGroupEventClass = _ChangeHistoryRemoveSubgroupFromGroupEventClass{objc.GetClass("CNChangeHistoryRemoveSubgroupFromGroupEvent")}

type _ChangeHistoryRemoveSubgroupFromGroupEventClass struct {
	objc.Class
}

// An interface definition for the [ChangeHistoryRemoveSubgroupFromGroupEvent] class.
type IChangeHistoryRemoveSubgroupFromGroupEvent interface {
	IChangeHistoryEvent
	Subgroup() Group
	Group() Group
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryremovesubgroupfromgroupevent?language=objc
type ChangeHistoryRemoveSubgroupFromGroupEvent struct {
	ChangeHistoryEvent
}

func ChangeHistoryRemoveSubgroupFromGroupEventFrom(ptr unsafe.Pointer) ChangeHistoryRemoveSubgroupFromGroupEvent {
	return ChangeHistoryRemoveSubgroupFromGroupEvent{
		ChangeHistoryEvent: ChangeHistoryEventFrom(ptr),
	}
}

func (cc _ChangeHistoryRemoveSubgroupFromGroupEventClass) Alloc() ChangeHistoryRemoveSubgroupFromGroupEvent {
	rv := objc.Call[ChangeHistoryRemoveSubgroupFromGroupEvent](cc, objc.Sel("alloc"))
	return rv
}

func ChangeHistoryRemoveSubgroupFromGroupEvent_Alloc() ChangeHistoryRemoveSubgroupFromGroupEvent {
	return ChangeHistoryRemoveSubgroupFromGroupEventClass.Alloc()
}

func (cc _ChangeHistoryRemoveSubgroupFromGroupEventClass) New() ChangeHistoryRemoveSubgroupFromGroupEvent {
	rv := objc.Call[ChangeHistoryRemoveSubgroupFromGroupEvent](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewChangeHistoryRemoveSubgroupFromGroupEvent() ChangeHistoryRemoveSubgroupFromGroupEvent {
	return ChangeHistoryRemoveSubgroupFromGroupEventClass.New()
}

func (c_ ChangeHistoryRemoveSubgroupFromGroupEvent) Init() ChangeHistoryRemoveSubgroupFromGroupEvent {
	rv := objc.Call[ChangeHistoryRemoveSubgroupFromGroupEvent](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryremovesubgroupfromgroupevent/3113285-subgroup?language=objc
func (c_ ChangeHistoryRemoveSubgroupFromGroupEvent) Subgroup() Group {
	rv := objc.Call[Group](c_, objc.Sel("subgroup"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryremovesubgroupfromgroupevent/3113284-group?language=objc
func (c_ ChangeHistoryRemoveSubgroupFromGroupEvent) Group() Group {
	rv := objc.Call[Group](c_, objc.Sel("group"))
	return rv
}
