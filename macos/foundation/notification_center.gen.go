// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NotificationCenter] class.
var NotificationCenterClass = _NotificationCenterClass{objc.GetClass("NSNotificationCenter")}

type _NotificationCenterClass struct {
	objc.Class
}

// An interface definition for the [NotificationCenter] class.
type INotificationCenter interface {
	objc.IObject
	PostNotificationNameObjectUserInfo(aName NotificationName, anObject objc.IObject, aUserInfo Dictionary)
	AddObserverForNameObjectQueueUsingBlock(name NotificationName, obj objc.IObject, queue IOperationQueue, block func(notification Notification)) objc.Object
	RemoveObserverNameObject(observer objc.IObject, aName NotificationName, anObject objc.IObject)
	AddObserverSelectorNameObject(observer objc.IObject, aSelector objc.Selector, aName NotificationName, anObject objc.IObject)
	PostNotification(notification INotification)
}

// A notification dispatch mechanism that enables the broadcast of information to registered observers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotificationcenter?language=objc
type NotificationCenter struct {
	objc.Object
}

func NotificationCenterFrom(ptr unsafe.Pointer) NotificationCenter {
	return NotificationCenter{
		Object: objc.ObjectFrom(ptr),
	}
}

func (nc _NotificationCenterClass) Alloc() NotificationCenter {
	rv := objc.Call[NotificationCenter](nc, objc.Sel("alloc"))
	return rv
}

func NotificationCenter_Alloc() NotificationCenter {
	return NotificationCenterClass.Alloc()
}

func (nc _NotificationCenterClass) New() NotificationCenter {
	rv := objc.Call[NotificationCenter](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNotificationCenter() NotificationCenter {
	return NotificationCenterClass.New()
}

func (n_ NotificationCenter) Init() NotificationCenter {
	rv := objc.Call[NotificationCenter](n_, objc.Sel("init"))
	return rv
}

// Creates a notification with a given name, sender, and information and posts it to the notification center. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotificationcenter/1410608-postnotificationname?language=objc
func (n_ NotificationCenter) PostNotificationNameObjectUserInfo(aName NotificationName, anObject objc.IObject, aUserInfo Dictionary) {
	objc.Call[objc.Void](n_, objc.Sel("postNotificationName:object:userInfo:"), aName, anObject, aUserInfo)
}

// Adds an entry to the notification center to receive notifications that passed to the provided block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotificationcenter/1411723-addobserverforname?language=objc
func (n_ NotificationCenter) AddObserverForNameObjectQueueUsingBlock(name NotificationName, obj objc.IObject, queue IOperationQueue, block func(notification Notification)) objc.Object {
	rv := objc.Call[objc.Object](n_, objc.Sel("addObserverForName:object:queue:usingBlock:"), name, obj, objc.Ptr(queue), block)
	return rv
}

// Removes matching entries from the notification center's dispatch table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotificationcenter/1407263-removeobserver?language=objc
func (n_ NotificationCenter) RemoveObserverNameObject(observer objc.IObject, aName NotificationName, anObject objc.IObject) {
	objc.Call[objc.Void](n_, objc.Sel("removeObserver:name:object:"), observer, aName, anObject)
}

// Adds an entry to the notification center to call the provided selector with the notification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotificationcenter/1415360-addobserver?language=objc
func (n_ NotificationCenter) AddObserverSelectorNameObject(observer objc.IObject, aSelector objc.Selector, aName NotificationName, anObject objc.IObject) {
	objc.Call[objc.Void](n_, objc.Sel("addObserver:selector:name:object:"), observer, aSelector, aName, anObject)
}

// Posts a given notification to the notification center. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotificationcenter/1410472-postnotification?language=objc
func (n_ NotificationCenter) PostNotification(notification INotification) {
	objc.Call[objc.Void](n_, objc.Sel("postNotification:"), objc.Ptr(notification))
}

// The app’s default notification center. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotificationcenter/1414169-defaultcenter?language=objc
func (nc _NotificationCenterClass) DefaultCenter() NotificationCenter {
	rv := objc.Call[NotificationCenter](nc, objc.Sel("defaultCenter"))
	return rv
}

// The app’s default notification center. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotificationcenter/1414169-defaultcenter?language=objc
func NotificationCenter_DefaultCenter() NotificationCenter {
	return NotificationCenterClass.DefaultCenter()
}
