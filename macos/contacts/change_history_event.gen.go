// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ChangeHistoryEvent] class.
var ChangeHistoryEventClass = _ChangeHistoryEventClass{objc.GetClass("CNChangeHistoryEvent")}

type _ChangeHistoryEventClass struct {
	objc.Class
}

// An interface definition for the [ChangeHistoryEvent] class.
type IChangeHistoryEvent interface {
	objc.IObject
	AcceptEventVisitor(visitor PChangeHistoryEventVisitor)
	AcceptEventVisitorObject(visitorObject objc.IObject)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryevent?language=objc
type ChangeHistoryEvent struct {
	objc.Object
}

func ChangeHistoryEventFrom(ptr unsafe.Pointer) ChangeHistoryEvent {
	return ChangeHistoryEvent{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _ChangeHistoryEventClass) Alloc() ChangeHistoryEvent {
	rv := objc.Call[ChangeHistoryEvent](cc, objc.Sel("alloc"))
	return rv
}

func ChangeHistoryEvent_Alloc() ChangeHistoryEvent {
	return ChangeHistoryEventClass.Alloc()
}

func (cc _ChangeHistoryEventClass) New() ChangeHistoryEvent {
	rv := objc.Call[ChangeHistoryEvent](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewChangeHistoryEvent() ChangeHistoryEvent {
	return ChangeHistoryEventClass.New()
}

func (c_ ChangeHistoryEvent) Init() ChangeHistoryEvent {
	rv := objc.Call[ChangeHistoryEvent](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryevent/3113267-accepteventvisitor?language=objc
func (c_ ChangeHistoryEvent) AcceptEventVisitor(visitor PChangeHistoryEventVisitor) {
	po0 := objc.WrapAsProtocol("CNChangeHistoryEventVisitor", visitor)
	objc.Call[objc.Void](c_, objc.Sel("acceptEventVisitor:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryevent/3113267-accepteventvisitor?language=objc
func (c_ ChangeHistoryEvent) AcceptEventVisitorObject(visitorObject objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("acceptEventVisitor:"), objc.Ptr(visitorObject))
}
