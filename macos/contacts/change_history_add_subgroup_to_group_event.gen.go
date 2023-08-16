// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ChangeHistoryAddSubgroupToGroupEvent] class.
var ChangeHistoryAddSubgroupToGroupEventClass = _ChangeHistoryAddSubgroupToGroupEventClass{objc.GetClass("CNChangeHistoryAddSubgroupToGroupEvent")}

type _ChangeHistoryAddSubgroupToGroupEventClass struct {
	objc.Class
}

// An interface definition for the [ChangeHistoryAddSubgroupToGroupEvent] class.
type IChangeHistoryAddSubgroupToGroupEvent interface {
	IChangeHistoryEvent
	Subgroup() Group
	Group() Group
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryaddsubgrouptogroupevent?language=objc
type ChangeHistoryAddSubgroupToGroupEvent struct {
	ChangeHistoryEvent
}

func ChangeHistoryAddSubgroupToGroupEventFrom(ptr unsafe.Pointer) ChangeHistoryAddSubgroupToGroupEvent {
	return ChangeHistoryAddSubgroupToGroupEvent{
		ChangeHistoryEvent: ChangeHistoryEventFrom(ptr),
	}
}

func (cc _ChangeHistoryAddSubgroupToGroupEventClass) Alloc() ChangeHistoryAddSubgroupToGroupEvent {
	rv := objc.Call[ChangeHistoryAddSubgroupToGroupEvent](cc, objc.Sel("alloc"))
	return rv
}

func ChangeHistoryAddSubgroupToGroupEvent_Alloc() ChangeHistoryAddSubgroupToGroupEvent {
	return ChangeHistoryAddSubgroupToGroupEventClass.Alloc()
}

func (cc _ChangeHistoryAddSubgroupToGroupEventClass) New() ChangeHistoryAddSubgroupToGroupEvent {
	rv := objc.Call[ChangeHistoryAddSubgroupToGroupEvent](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewChangeHistoryAddSubgroupToGroupEvent() ChangeHistoryAddSubgroupToGroupEvent {
	return ChangeHistoryAddSubgroupToGroupEventClass.New()
}

func (c_ ChangeHistoryAddSubgroupToGroupEvent) Init() ChangeHistoryAddSubgroupToGroupEvent {
	rv := objc.Call[ChangeHistoryAddSubgroupToGroupEvent](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryaddsubgrouptogroupevent/3113260-subgroup?language=objc
func (c_ ChangeHistoryAddSubgroupToGroupEvent) Subgroup() Group {
	rv := objc.Call[Group](c_, objc.Sel("subgroup"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryaddsubgrouptogroupevent/3113259-group?language=objc
func (c_ ChangeHistoryAddSubgroupToGroupEvent) Group() Group {
	rv := objc.Call[Group](c_, objc.Sel("group"))
	return rv
}
