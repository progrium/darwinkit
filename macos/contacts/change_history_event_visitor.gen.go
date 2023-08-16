// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryeventvisitor?language=objc
type PChangeHistoryEventVisitor interface {
	// optional
	VisitAddMemberToGroupEvent(event ChangeHistoryAddMemberToGroupEvent)
	HasVisitAddMemberToGroupEvent() bool

	// optional
	VisitDeleteContactEvent(event ChangeHistoryDeleteContactEvent)
	HasVisitDeleteContactEvent() bool

	// optional
	VisitAddContactEvent(event ChangeHistoryAddContactEvent)
	HasVisitAddContactEvent() bool

	// optional
	VisitAddGroupEvent(event ChangeHistoryAddGroupEvent)
	HasVisitAddGroupEvent() bool

	// optional
	VisitDeleteGroupEvent(event ChangeHistoryDeleteGroupEvent)
	HasVisitDeleteGroupEvent() bool

	// optional
	VisitRemoveSubgroupFromGroupEvent(event ChangeHistoryRemoveSubgroupFromGroupEvent)
	HasVisitRemoveSubgroupFromGroupEvent() bool

	// optional
	VisitUpdateContactEvent(event ChangeHistoryUpdateContactEvent)
	HasVisitUpdateContactEvent() bool

	// optional
	VisitDropEverythingEvent(event ChangeHistoryDropEverythingEvent)
	HasVisitDropEverythingEvent() bool

	// optional
	VisitUpdateGroupEvent(event ChangeHistoryUpdateGroupEvent)
	HasVisitUpdateGroupEvent() bool

	// optional
	VisitAddSubgroupToGroupEvent(event ChangeHistoryAddSubgroupToGroupEvent)
	HasVisitAddSubgroupToGroupEvent() bool

	// optional
	VisitRemoveMemberFromGroupEvent(event ChangeHistoryRemoveMemberFromGroupEvent)
	HasVisitRemoveMemberFromGroupEvent() bool
}

// A concrete type wrapper for the [PChangeHistoryEventVisitor] protocol.
type ChangeHistoryEventVisitorWrapper struct {
	objc.Object
}

func (c_ ChangeHistoryEventVisitorWrapper) HasVisitAddMemberToGroupEvent() bool {
	return c_.RespondsToSelector(objc.Sel("visitAddMemberToGroupEvent:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryeventvisitor/3113271-visitaddmembertogroupevent?language=objc
func (c_ ChangeHistoryEventVisitorWrapper) VisitAddMemberToGroupEvent(event IChangeHistoryAddMemberToGroupEvent) {
	objc.Call[objc.Void](c_, objc.Sel("visitAddMemberToGroupEvent:"), objc.Ptr(event))
}

func (c_ ChangeHistoryEventVisitorWrapper) HasVisitDeleteContactEvent() bool {
	return c_.RespondsToSelector(objc.Sel("visitDeleteContactEvent:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryeventvisitor/3113273-visitdeletecontactevent?language=objc
func (c_ ChangeHistoryEventVisitorWrapper) VisitDeleteContactEvent(event IChangeHistoryDeleteContactEvent) {
	objc.Call[objc.Void](c_, objc.Sel("visitDeleteContactEvent:"), objc.Ptr(event))
}

func (c_ ChangeHistoryEventVisitorWrapper) HasVisitAddContactEvent() bool {
	return c_.RespondsToSelector(objc.Sel("visitAddContactEvent:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryeventvisitor/3113269-visitaddcontactevent?language=objc
func (c_ ChangeHistoryEventVisitorWrapper) VisitAddContactEvent(event IChangeHistoryAddContactEvent) {
	objc.Call[objc.Void](c_, objc.Sel("visitAddContactEvent:"), objc.Ptr(event))
}

func (c_ ChangeHistoryEventVisitorWrapper) HasVisitAddGroupEvent() bool {
	return c_.RespondsToSelector(objc.Sel("visitAddGroupEvent:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryeventvisitor/3113270-visitaddgroupevent?language=objc
func (c_ ChangeHistoryEventVisitorWrapper) VisitAddGroupEvent(event IChangeHistoryAddGroupEvent) {
	objc.Call[objc.Void](c_, objc.Sel("visitAddGroupEvent:"), objc.Ptr(event))
}

func (c_ ChangeHistoryEventVisitorWrapper) HasVisitDeleteGroupEvent() bool {
	return c_.RespondsToSelector(objc.Sel("visitDeleteGroupEvent:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryeventvisitor/3113274-visitdeletegroupevent?language=objc
func (c_ ChangeHistoryEventVisitorWrapper) VisitDeleteGroupEvent(event IChangeHistoryDeleteGroupEvent) {
	objc.Call[objc.Void](c_, objc.Sel("visitDeleteGroupEvent:"), objc.Ptr(event))
}

func (c_ ChangeHistoryEventVisitorWrapper) HasVisitRemoveSubgroupFromGroupEvent() bool {
	return c_.RespondsToSelector(objc.Sel("visitRemoveSubgroupFromGroupEvent:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryeventvisitor/3113277-visitremovesubgroupfromgroupeven?language=objc
func (c_ ChangeHistoryEventVisitorWrapper) VisitRemoveSubgroupFromGroupEvent(event IChangeHistoryRemoveSubgroupFromGroupEvent) {
	objc.Call[objc.Void](c_, objc.Sel("visitRemoveSubgroupFromGroupEvent:"), objc.Ptr(event))
}

func (c_ ChangeHistoryEventVisitorWrapper) HasVisitUpdateContactEvent() bool {
	return c_.RespondsToSelector(objc.Sel("visitUpdateContactEvent:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryeventvisitor/3113278-visitupdatecontactevent?language=objc
func (c_ ChangeHistoryEventVisitorWrapper) VisitUpdateContactEvent(event IChangeHistoryUpdateContactEvent) {
	objc.Call[objc.Void](c_, objc.Sel("visitUpdateContactEvent:"), objc.Ptr(event))
}

func (c_ ChangeHistoryEventVisitorWrapper) HasVisitDropEverythingEvent() bool {
	return c_.RespondsToSelector(objc.Sel("visitDropEverythingEvent:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryeventvisitor/3113275-visitdropeverythingevent?language=objc
func (c_ ChangeHistoryEventVisitorWrapper) VisitDropEverythingEvent(event IChangeHistoryDropEverythingEvent) {
	objc.Call[objc.Void](c_, objc.Sel("visitDropEverythingEvent:"), objc.Ptr(event))
}

func (c_ ChangeHistoryEventVisitorWrapper) HasVisitUpdateGroupEvent() bool {
	return c_.RespondsToSelector(objc.Sel("visitUpdateGroupEvent:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryeventvisitor/3113279-visitupdategroupevent?language=objc
func (c_ ChangeHistoryEventVisitorWrapper) VisitUpdateGroupEvent(event IChangeHistoryUpdateGroupEvent) {
	objc.Call[objc.Void](c_, objc.Sel("visitUpdateGroupEvent:"), objc.Ptr(event))
}

func (c_ ChangeHistoryEventVisitorWrapper) HasVisitAddSubgroupToGroupEvent() bool {
	return c_.RespondsToSelector(objc.Sel("visitAddSubgroupToGroupEvent:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryeventvisitor/3113272-visitaddsubgrouptogroupevent?language=objc
func (c_ ChangeHistoryEventVisitorWrapper) VisitAddSubgroupToGroupEvent(event IChangeHistoryAddSubgroupToGroupEvent) {
	objc.Call[objc.Void](c_, objc.Sel("visitAddSubgroupToGroupEvent:"), objc.Ptr(event))
}

func (c_ ChangeHistoryEventVisitorWrapper) HasVisitRemoveMemberFromGroupEvent() bool {
	return c_.RespondsToSelector(objc.Sel("visitRemoveMemberFromGroupEvent:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryeventvisitor/3113276-visitremovememberfromgroupevent?language=objc
func (c_ ChangeHistoryEventVisitorWrapper) VisitRemoveMemberFromGroupEvent(event IChangeHistoryRemoveMemberFromGroupEvent) {
	objc.Call[objc.Void](c_, objc.Sel("visitRemoveMemberFromGroupEvent:"), objc.Ptr(event))
}
