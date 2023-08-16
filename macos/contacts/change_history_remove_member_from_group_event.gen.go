// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ChangeHistoryRemoveMemberFromGroupEvent] class.
var ChangeHistoryRemoveMemberFromGroupEventClass = _ChangeHistoryRemoveMemberFromGroupEventClass{objc.GetClass("CNChangeHistoryRemoveMemberFromGroupEvent")}

type _ChangeHistoryRemoveMemberFromGroupEventClass struct {
	objc.Class
}

// An interface definition for the [ChangeHistoryRemoveMemberFromGroupEvent] class.
type IChangeHistoryRemoveMemberFromGroupEvent interface {
	IChangeHistoryEvent
	Member() Contact
	Group() Group
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryremovememberfromgroupevent?language=objc
type ChangeHistoryRemoveMemberFromGroupEvent struct {
	ChangeHistoryEvent
}

func ChangeHistoryRemoveMemberFromGroupEventFrom(ptr unsafe.Pointer) ChangeHistoryRemoveMemberFromGroupEvent {
	return ChangeHistoryRemoveMemberFromGroupEvent{
		ChangeHistoryEvent: ChangeHistoryEventFrom(ptr),
	}
}

func (cc _ChangeHistoryRemoveMemberFromGroupEventClass) Alloc() ChangeHistoryRemoveMemberFromGroupEvent {
	rv := objc.Call[ChangeHistoryRemoveMemberFromGroupEvent](cc, objc.Sel("alloc"))
	return rv
}

func ChangeHistoryRemoveMemberFromGroupEvent_Alloc() ChangeHistoryRemoveMemberFromGroupEvent {
	return ChangeHistoryRemoveMemberFromGroupEventClass.Alloc()
}

func (cc _ChangeHistoryRemoveMemberFromGroupEventClass) New() ChangeHistoryRemoveMemberFromGroupEvent {
	rv := objc.Call[ChangeHistoryRemoveMemberFromGroupEvent](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewChangeHistoryRemoveMemberFromGroupEvent() ChangeHistoryRemoveMemberFromGroupEvent {
	return ChangeHistoryRemoveMemberFromGroupEventClass.New()
}

func (c_ ChangeHistoryRemoveMemberFromGroupEvent) Init() ChangeHistoryRemoveMemberFromGroupEvent {
	rv := objc.Call[ChangeHistoryRemoveMemberFromGroupEvent](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryremovememberfromgroupevent/3113282-member?language=objc
func (c_ ChangeHistoryRemoveMemberFromGroupEvent) Member() Contact {
	rv := objc.Call[Contact](c_, objc.Sel("member"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryremovememberfromgroupevent/3113281-group?language=objc
func (c_ ChangeHistoryRemoveMemberFromGroupEvent) Group() Group {
	rv := objc.Call[Group](c_, objc.Sel("group"))
	return rv
}
