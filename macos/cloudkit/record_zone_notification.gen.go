// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RecordZoneNotification] class.
var RecordZoneNotificationClass = _RecordZoneNotificationClass{objc.GetClass("CKRecordZoneNotification")}

type _RecordZoneNotificationClass struct {
	objc.Class
}

// An interface definition for the [RecordZoneNotification] class.
type IRecordZoneNotification interface {
	INotification
	DatabaseScope() DatabaseScope
	RecordZoneID() RecordZoneID
}

// A notification that triggers when the contents of a record zone change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordzonenotification?language=objc
type RecordZoneNotification struct {
	Notification
}

func RecordZoneNotificationFrom(ptr unsafe.Pointer) RecordZoneNotification {
	return RecordZoneNotification{
		Notification: NotificationFrom(ptr),
	}
}

func (rc _RecordZoneNotificationClass) Alloc() RecordZoneNotification {
	rv := objc.Call[RecordZoneNotification](rc, objc.Sel("alloc"))
	return rv
}

func RecordZoneNotification_Alloc() RecordZoneNotification {
	return RecordZoneNotificationClass.Alloc()
}

func (rc _RecordZoneNotificationClass) New() RecordZoneNotification {
	rv := objc.Call[RecordZoneNotification](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRecordZoneNotification() RecordZoneNotification {
	return RecordZoneNotificationClass.New()
}

func (r_ RecordZoneNotification) Init() RecordZoneNotification {
	rv := objc.Call[RecordZoneNotification](r_, objc.Sel("init"))
	return rv
}

func (rc _RecordZoneNotificationClass) NotificationFromRemoteNotificationDictionary(notificationDictionary foundation.Dictionary) RecordZoneNotification {
	rv := objc.Call[RecordZoneNotification](rc, objc.Sel("notificationFromRemoteNotificationDictionary:"), notificationDictionary)
	return rv
}

// Creates a new notification using the specified payload data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/1428130-notificationfromremotenotificati?language=objc
func RecordZoneNotification_NotificationFromRemoteNotificationDictionary(notificationDictionary foundation.Dictionary) RecordZoneNotification {
	return RecordZoneNotificationClass.NotificationFromRemoteNotificationDictionary(notificationDictionary)
}

// The type of database for the record zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordzonenotification/1640394-databasescope?language=objc
func (r_ RecordZoneNotification) DatabaseScope() DatabaseScope {
	rv := objc.Call[DatabaseScope](r_, objc.Sel("databaseScope"))
	return rv
}

// The ID of the record zone that has changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordzonenotification/1428086-recordzoneid?language=objc
func (r_ RecordZoneNotification) RecordZoneID() RecordZoneID {
	rv := objc.Call[RecordZoneID](r_, objc.Sel("recordZoneID"))
	return rv
}
