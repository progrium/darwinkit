// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Notification] class.
var NotificationClass = _NotificationClass{objc.GetClass("CKNotification")}

type _NotificationClass struct {
	objc.Class
}

// An interface definition for the [Notification] class.
type INotification interface {
	objc.IObject
	NotificationID() NotificationID
	IsPruned() bool
	NotificationType() NotificationType
	SubscriptionID() SubscriptionID
	ContainerIdentifier() string
	SubscriptionOwnerUserRecordID() RecordID
}

// The abstract base class for CloudKit notifications. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification?language=objc
type Notification struct {
	objc.Object
}

func NotificationFrom(ptr unsafe.Pointer) Notification {
	return Notification{
		Object: objc.ObjectFrom(ptr),
	}
}

func (nc _NotificationClass) NotificationFromRemoteNotificationDictionary(notificationDictionary foundation.Dictionary) Notification {
	rv := objc.Call[Notification](nc, objc.Sel("notificationFromRemoteNotificationDictionary:"), notificationDictionary)
	return rv
}

// Creates a new notification using the specified payload data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/1428130-notificationfromremotenotificati?language=objc
func Notification_NotificationFromRemoteNotificationDictionary(notificationDictionary foundation.Dictionary) Notification {
	return NotificationClass.NotificationFromRemoteNotificationDictionary(notificationDictionary)
}

func (nc _NotificationClass) Alloc() Notification {
	rv := objc.Call[Notification](nc, objc.Sel("alloc"))
	return rv
}

func Notification_Alloc() Notification {
	return NotificationClass.Alloc()
}

func (nc _NotificationClass) New() Notification {
	rv := objc.Call[Notification](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNotification() Notification {
	return NotificationClass.New()
}

func (n_ Notification) Init() Notification {
	rv := objc.Call[Notification](n_, objc.Sel("init"))
	return rv
}

// The notification’s ID. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/1428080-notificationid?language=objc
func (n_ Notification) NotificationID() NotificationID {
	rv := objc.Call[NotificationID](n_, objc.Sel("notificationID"))
	return rv
}

// A Boolean value that indicates whether the system removes some push notification content before delivery. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/1428088-ispruned?language=objc
func (n_ Notification) IsPruned() bool {
	rv := objc.Call[bool](n_, objc.Sel("isPruned"))
	return rv
}

// The type of event that generates the notification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/1428078-notificationtype?language=objc
func (n_ Notification) NotificationType() NotificationType {
	rv := objc.Call[NotificationType](n_, objc.Sel("notificationType"))
	return rv
}

// The ID of the subscription that triggers the notification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/1428118-subscriptionid?language=objc
func (n_ Notification) SubscriptionID() SubscriptionID {
	rv := objc.Call[SubscriptionID](n_, objc.Sel("subscriptionID"))
	return rv
}

// The ID of the container with the content that triggers the notification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/1428119-containeridentifier?language=objc
func (n_ Notification) ContainerIdentifier() string {
	rv := objc.Call[string](n_, objc.Sel("containerIdentifier"))
	return rv
}

// The ID of the user record that creates the subscription that generates the push notification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/3577533-subscriptionowneruserrecordid?language=objc
func (n_ Notification) SubscriptionOwnerUserRecordID() RecordID {
	rv := objc.Call[RecordID](n_, objc.Sel("subscriptionOwnerUserRecordID"))
	return rv
}
