// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ChangeHistoryAddMemberToGroupEvent] class.
var ChangeHistoryAddMemberToGroupEventClass = _ChangeHistoryAddMemberToGroupEventClass{objc.GetClass("CNChangeHistoryAddMemberToGroupEvent")}

type _ChangeHistoryAddMemberToGroupEventClass struct {
	objc.Class
}

// An interface definition for the [ChangeHistoryAddMemberToGroupEvent] class.
type IChangeHistoryAddMemberToGroupEvent interface {
	IChangeHistoryEvent
	Member() Contact
	Group() Group
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryaddmembertogroupevent?language=objc
type ChangeHistoryAddMemberToGroupEvent struct {
	ChangeHistoryEvent
}

func ChangeHistoryAddMemberToGroupEventFrom(ptr unsafe.Pointer) ChangeHistoryAddMemberToGroupEvent {
	return ChangeHistoryAddMemberToGroupEvent{
		ChangeHistoryEvent: ChangeHistoryEventFrom(ptr),
	}
}

func (cc _ChangeHistoryAddMemberToGroupEventClass) Alloc() ChangeHistoryAddMemberToGroupEvent {
	rv := objc.Call[ChangeHistoryAddMemberToGroupEvent](cc, objc.Sel("alloc"))
	return rv
}

func ChangeHistoryAddMemberToGroupEvent_Alloc() ChangeHistoryAddMemberToGroupEvent {
	return ChangeHistoryAddMemberToGroupEventClass.Alloc()
}

func (cc _ChangeHistoryAddMemberToGroupEventClass) New() ChangeHistoryAddMemberToGroupEvent {
	rv := objc.Call[ChangeHistoryAddMemberToGroupEvent](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewChangeHistoryAddMemberToGroupEvent() ChangeHistoryAddMemberToGroupEvent {
	return ChangeHistoryAddMemberToGroupEventClass.New()
}

func (c_ ChangeHistoryAddMemberToGroupEvent) Init() ChangeHistoryAddMemberToGroupEvent {
	rv := objc.Call[ChangeHistoryAddMemberToGroupEvent](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryaddmembertogroupevent/3113257-member?language=objc
func (c_ ChangeHistoryAddMemberToGroupEvent) Member() Contact {
	rv := objc.Call[Contact](c_, objc.Sel("member"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryaddmembertogroupevent/3113256-group?language=objc
func (c_ ChangeHistoryAddMemberToGroupEvent) Group() Group {
	rv := objc.Call[Group](c_, objc.Sel("group"))
	return rv
}
