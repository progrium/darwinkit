// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [QueryNotification] class.
var QueryNotificationClass = _QueryNotificationClass{objc.GetClass("CKQueryNotification")}

type _QueryNotificationClass struct {
	objc.Class
}

// An interface definition for the [QueryNotification] class.
type IQueryNotification interface {
	INotification
	RecordID() RecordID
	DatabaseScope() DatabaseScope
	QueryNotificationReason() QueryNotificationReason
	RecordFields() map[string]objc.Object
}

// A notification that triggers when a record that matches the subscription’s predicate changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckquerynotification?language=objc
type QueryNotification struct {
	Notification
}

func QueryNotificationFrom(ptr unsafe.Pointer) QueryNotification {
	return QueryNotification{
		Notification: NotificationFrom(ptr),
	}
}

func (qc _QueryNotificationClass) Alloc() QueryNotification {
	rv := objc.Call[QueryNotification](qc, objc.Sel("alloc"))
	return rv
}

func QueryNotification_Alloc() QueryNotification {
	return QueryNotificationClass.Alloc()
}

func (qc _QueryNotificationClass) New() QueryNotification {
	rv := objc.Call[QueryNotification](qc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewQueryNotification() QueryNotification {
	return QueryNotificationClass.New()
}

func (q_ QueryNotification) Init() QueryNotification {
	rv := objc.Call[QueryNotification](q_, objc.Sel("init"))
	return rv
}

func (qc _QueryNotificationClass) NotificationFromRemoteNotificationDictionary(notificationDictionary foundation.Dictionary) QueryNotification {
	rv := objc.Call[QueryNotification](qc, objc.Sel("notificationFromRemoteNotificationDictionary:"), notificationDictionary)
	return rv
}

// Creates a new notification using the specified payload data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/1428130-notificationfromremotenotificati?language=objc
func QueryNotification_NotificationFromRemoteNotificationDictionary(notificationDictionary foundation.Dictionary) QueryNotification {
	return QueryNotificationClass.NotificationFromRemoteNotificationDictionary(notificationDictionary)
}

// The ID of the record that CloudKit creates, updates, or deletes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckquerynotification/1428134-recordid?language=objc
func (q_ QueryNotification) RecordID() RecordID {
	rv := objc.Call[RecordID](q_, objc.Sel("recordID"))
	return rv
}

// The type of database for the record zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckquerynotification/1640449-databasescope?language=objc
func (q_ QueryNotification) DatabaseScope() DatabaseScope {
	rv := objc.Call[DatabaseScope](q_, objc.Sel("databaseScope"))
	return rv
}

// The event that triggers the push notification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckquerynotification/1428123-querynotificationreason?language=objc
func (q_ QueryNotification) QueryNotificationReason() QueryNotificationReason {
	rv := objc.Call[QueryNotificationReason](q_, objc.Sel("queryNotificationReason"))
	return rv
}

// A dictionary of fields that have changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckquerynotification/1428114-recordfields?language=objc
func (q_ QueryNotification) RecordFields() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](q_, objc.Sel("recordFields"))
	return rv
}
