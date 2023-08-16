// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NotificationID] class.
var NotificationIDClass = _NotificationIDClass{objc.GetClass("CKNotificationID")}

type _NotificationIDClass struct {
	objc.Class
}

// An interface definition for the [NotificationID] class.
type INotificationID interface {
	objc.IObject
}

// An object that uniquely identifies a push notification that a container sends. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationid?language=objc
type NotificationID struct {
	objc.Object
}

func NotificationIDFrom(ptr unsafe.Pointer) NotificationID {
	return NotificationID{
		Object: objc.ObjectFrom(ptr),
	}
}

func (nc _NotificationIDClass) Alloc() NotificationID {
	rv := objc.Call[NotificationID](nc, objc.Sel("alloc"))
	return rv
}

func NotificationID_Alloc() NotificationID {
	return NotificationIDClass.Alloc()
}

func (nc _NotificationIDClass) New() NotificationID {
	rv := objc.Call[NotificationID](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNotificationID() NotificationID {
	return NotificationIDClass.New()
}

func (n_ NotificationID) Init() NotificationID {
	rv := objc.Call[NotificationID](n_, objc.Sel("init"))
	return rv
}
