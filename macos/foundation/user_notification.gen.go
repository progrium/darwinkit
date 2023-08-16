// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UserNotification] class.
var UserNotificationClass = _UserNotificationClass{objc.GetClass("NSUserNotification")}

type _UserNotificationClass struct {
	objc.Class
}

// An interface definition for the [UserNotification] class.
type IUserNotification interface {
	objc.IObject
}

// A notification that can be scheduled for display in the notification center. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification?language=objc
type UserNotification struct {
	objc.Object
}

func UserNotificationFrom(ptr unsafe.Pointer) UserNotification {
	return UserNotification{
		Object: objc.ObjectFrom(ptr),
	}
}

func (uc _UserNotificationClass) Alloc() UserNotification {
	rv := objc.Call[UserNotification](uc, objc.Sel("alloc"))
	return rv
}

func UserNotification_Alloc() UserNotification {
	return UserNotificationClass.Alloc()
}

func (uc _UserNotificationClass) New() UserNotification {
	rv := objc.Call[UserNotification](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUserNotification() UserNotification {
	return UserNotificationClass.New()
}

func (u_ UserNotification) Init() UserNotification {
	rv := objc.Call[UserNotification](u_, objc.Sel("init"))
	return rv
}
