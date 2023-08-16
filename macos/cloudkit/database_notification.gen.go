// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DatabaseNotification] class.
var DatabaseNotificationClass = _DatabaseNotificationClass{objc.GetClass("CKDatabaseNotification")}

type _DatabaseNotificationClass struct {
	objc.Class
}

// An interface definition for the [DatabaseNotification] class.
type IDatabaseNotification interface {
	INotification
	DatabaseScope() DatabaseScope
}

// A notification that triggers when the contents of a database change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdatabasenotification?language=objc
type DatabaseNotification struct {
	Notification
}

func DatabaseNotificationFrom(ptr unsafe.Pointer) DatabaseNotification {
	return DatabaseNotification{
		Notification: NotificationFrom(ptr),
	}
}

func (dc _DatabaseNotificationClass) Alloc() DatabaseNotification {
	rv := objc.Call[DatabaseNotification](dc, objc.Sel("alloc"))
	return rv
}

func DatabaseNotification_Alloc() DatabaseNotification {
	return DatabaseNotificationClass.Alloc()
}

func (dc _DatabaseNotificationClass) New() DatabaseNotification {
	rv := objc.Call[DatabaseNotification](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDatabaseNotification() DatabaseNotification {
	return DatabaseNotificationClass.New()
}

func (d_ DatabaseNotification) Init() DatabaseNotification {
	rv := objc.Call[DatabaseNotification](d_, objc.Sel("init"))
	return rv
}

func (dc _DatabaseNotificationClass) NotificationFromRemoteNotificationDictionary(notificationDictionary foundation.Dictionary) DatabaseNotification {
	rv := objc.Call[DatabaseNotification](dc, objc.Sel("notificationFromRemoteNotificationDictionary:"), notificationDictionary)
	return rv
}

// Creates a new notification using the specified payload data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/1428130-notificationfromremotenotificati?language=objc
func DatabaseNotification_NotificationFromRemoteNotificationDictionary(notificationDictionary foundation.Dictionary) DatabaseNotification {
	return DatabaseNotificationClass.NotificationFromRemoteNotificationDictionary(notificationDictionary)
}

// The type of database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdatabasenotification/1640510-databasescope?language=objc
func (d_ DatabaseNotification) DatabaseScope() DatabaseScope {
	rv := objc.Call[DatabaseScope](d_, objc.Sel("databaseScope"))
	return rv
}
