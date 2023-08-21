// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Notification] class.
var NotificationClass = _NotificationClass{objc.GetClass("NSNotification")}

type _NotificationClass struct {
	objc.Class
}

// An interface definition for the [Notification] class.
type INotification interface {
	objc.IObject
	Name() NotificationName
	UserInfo() Dictionary
	Object_() objc.Object
}

// A container for information broadcast through a notification center to all registered observers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotification?language=objc
type Notification struct {
	objc.Object
}

func NotificationFrom(ptr unsafe.Pointer) Notification {
	return Notification{
		Object: objc.ObjectFrom(ptr),
	}
}

func (n_ Notification) InitWithNameObjectUserInfo(name NotificationName, object objc.IObject, userInfo Dictionary) Notification {
	rv := objc.Call[Notification](n_, objc.Sel("initWithName:object:userInfo:"), name, object, userInfo)
	return rv
}

// Initializes a notification with a specified name, object, and user information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotification/1415764-initwithname?language=objc
func NewNotificationWithNameObjectUserInfo(name NotificationName, object objc.IObject, userInfo Dictionary) Notification {
	instance := NotificationClass.Alloc().InitWithNameObjectUserInfo(name, object, userInfo)
	instance.Autorelease()
	return instance
}

func (nc _NotificationClass) NotificationWithNameObject(aName NotificationName, anObject objc.IObject) Notification {
	rv := objc.Call[Notification](nc, objc.Sel("notificationWithName:object:"), aName, anObject)
	return rv
}

// Returns a new notification object with a specified name and object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotification/1417440-notificationwithname?language=objc
func Notification_NotificationWithNameObject(aName NotificationName, anObject objc.IObject) Notification {
	return NotificationClass.NotificationWithNameObject(aName, anObject)
}

func (n_ Notification) Init() Notification {
	rv := objc.Call[Notification](n_, objc.Sel("init"))
	return rv
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

// The name of the notification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotification/1416472-name?language=objc
func (n_ Notification) Name() NotificationName {
	rv := objc.Call[NotificationName](n_, objc.Sel("name"))
	return rv
}

// The user information dictionary associated with the notification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotification/1409222-userinfo?language=objc
func (n_ Notification) UserInfo() Dictionary {
	rv := objc.Call[Dictionary](n_, objc.Sel("userInfo"))
	return rv
}

// The object associated with the notification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotification/1414469-object?language=objc
func (n_ Notification) Object_() objc.Object {
	rv := objc.Call[objc.Object](n_, objc.Sel("object"))
	return rv
}
