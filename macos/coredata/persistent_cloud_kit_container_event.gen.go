// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PersistentCloudKitContainerEvent] class.
var PersistentCloudKitContainerEventClass = _PersistentCloudKitContainerEventClass{objc.GetClass("NSPersistentCloudKitContainerEvent")}

type _PersistentCloudKitContainerEventClass struct {
	objc.Class
}

// An interface definition for the [PersistentCloudKitContainerEvent] class.
type IPersistentCloudKitContainerEvent interface {
	objc.IObject
	Error() foundation.Error
	StoreIdentifier() string
	Succeeded() bool
	StartDate() foundation.Date
	Type() PersistentCloudKitContainerEventType
	EndDate() foundation.Date
	Identifier() foundation.UUID
}

// An object that represents activity in a persistent CloudKit container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainerevent?language=objc
type PersistentCloudKitContainerEvent struct {
	objc.Object
}

func PersistentCloudKitContainerEventFrom(ptr unsafe.Pointer) PersistentCloudKitContainerEvent {
	return PersistentCloudKitContainerEvent{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PersistentCloudKitContainerEventClass) Alloc() PersistentCloudKitContainerEvent {
	rv := objc.Call[PersistentCloudKitContainerEvent](pc, objc.Sel("alloc"))
	return rv
}

func PersistentCloudKitContainerEvent_Alloc() PersistentCloudKitContainerEvent {
	return PersistentCloudKitContainerEventClass.Alloc()
}

func (pc _PersistentCloudKitContainerEventClass) New() PersistentCloudKitContainerEvent {
	rv := objc.Call[PersistentCloudKitContainerEvent](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPersistentCloudKitContainerEvent() PersistentCloudKitContainerEvent {
	return PersistentCloudKitContainerEventClass.New()
}

func (p_ PersistentCloudKitContainerEvent) Init() PersistentCloudKitContainerEvent {
	rv := objc.Call[PersistentCloudKitContainerEvent](p_, objc.Sel("init"))
	return rv
}

// An error that indicates why an operation fails. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainerevent/3618802-error?language=objc
func (p_ PersistentCloudKitContainerEvent) Error() foundation.Error {
	rv := objc.Call[foundation.Error](p_, objc.Sel("error"))
	return rv
}

// The associated store identifier in the container for the event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainerevent/3618805-storeidentifier?language=objc
func (p_ PersistentCloudKitContainerEvent) StoreIdentifier() string {
	rv := objc.Call[string](p_, objc.Sel("storeIdentifier"))
	return rv
}

// A Boolean value that indicates whether the operation the event represents is successful. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainerevent/3618806-succeeded?language=objc
func (p_ PersistentCloudKitContainerEvent) Succeeded() bool {
	rv := objc.Call[bool](p_, objc.Sel("succeeded"))
	return rv
}

// The start date of the operation that the event represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainerevent/3618804-startdate?language=objc
func (p_ PersistentCloudKitContainerEvent) StartDate() foundation.Date {
	rv := objc.Call[foundation.Date](p_, objc.Sel("startDate"))
	return rv
}

// The type of event, either setup, import, or export. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainerevent/3618807-type?language=objc
func (p_ PersistentCloudKitContainerEvent) Type() PersistentCloudKitContainerEventType {
	rv := objc.Call[PersistentCloudKitContainerEventType](p_, objc.Sel("type"))
	return rv
}

// The end date of the operation that the event represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainerevent/3618801-enddate?language=objc
func (p_ PersistentCloudKitContainerEvent) EndDate() foundation.Date {
	rv := objc.Call[foundation.Date](p_, objc.Sel("endDate"))
	return rv
}

// A unique identifier for the event in a container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainerevent/3618803-identifier?language=objc
func (p_ PersistentCloudKitContainerEvent) Identifier() foundation.UUID {
	rv := objc.Call[foundation.UUID](p_, objc.Sel("identifier"))
	return rv
}
