// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DistributedNotificationCenter] class.
var DistributedNotificationCenterClass = _DistributedNotificationCenterClass{objc.GetClass("NSDistributedNotificationCenter")}

type _DistributedNotificationCenterClass struct {
	objc.Class
}

// An interface definition for the [DistributedNotificationCenter] class.
type IDistributedNotificationCenter interface {
	INotificationCenter
	PostNotificationNameObject(aName NotificationName, anObject string)
	Suspended() bool
	SetSuspended(value bool)
}

// A notification dispatch mechanism that enables the broadcast of notifications across task boundaries. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdistributednotificationcenter?language=objc
type DistributedNotificationCenter struct {
	NotificationCenter
}

func DistributedNotificationCenterFrom(ptr unsafe.Pointer) DistributedNotificationCenter {
	return DistributedNotificationCenter{
		NotificationCenter: NotificationCenterFrom(ptr),
	}
}

func (dc _DistributedNotificationCenterClass) Alloc() DistributedNotificationCenter {
	rv := objc.Call[DistributedNotificationCenter](dc, objc.Sel("alloc"))
	return rv
}

func DistributedNotificationCenter_Alloc() DistributedNotificationCenter {
	return DistributedNotificationCenterClass.Alloc()
}

func (dc _DistributedNotificationCenterClass) New() DistributedNotificationCenter {
	rv := objc.Call[DistributedNotificationCenter](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDistributedNotificationCenter() DistributedNotificationCenter {
	return DistributedNotificationCenterClass.New()
}

func (d_ DistributedNotificationCenter) Init() DistributedNotificationCenter {
	rv := objc.Call[DistributedNotificationCenter](d_, objc.Sel("init"))
	return rv
}

// Creates a notification, and posts it to the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdistributednotificationcenter/1410991-postnotificationname?language=objc
func (d_ DistributedNotificationCenter) PostNotificationNameObject(aName NotificationName, anObject string) {
	objc.Call[objc.Void](d_, objc.Sel("postNotificationName:object:"), aName, anObject)
}

// Returns the distributed notification center for a particular notification center type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdistributednotificationcenter/1415403-notificationcenterfortype?language=objc
func (dc _DistributedNotificationCenterClass) NotificationCenterForType(notificationCenterType DistributedNotificationCenterType) DistributedNotificationCenter {
	rv := objc.Call[DistributedNotificationCenter](dc, objc.Sel("notificationCenterForType:"), notificationCenterType)
	return rv
}

// Returns the distributed notification center for a particular notification center type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdistributednotificationcenter/1415403-notificationcenterfortype?language=objc
func DistributedNotificationCenter_NotificationCenterForType(notificationCenterType DistributedNotificationCenterType) DistributedNotificationCenter {
	return DistributedNotificationCenterClass.NotificationCenterForType(notificationCenterType)
}

// Suspends or resumes notification delivery. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdistributednotificationcenter/1407301-suspended?language=objc
func (d_ DistributedNotificationCenter) Suspended() bool {
	rv := objc.Call[bool](d_, objc.Sel("suspended"))
	return rv
}

// Suspends or resumes notification delivery. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdistributednotificationcenter/1407301-suspended?language=objc
func (d_ DistributedNotificationCenter) SetSuspended(value bool) {
	objc.Call[objc.Void](d_, objc.Sel("setSuspended:"), value)
}
