// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UserNotificationCenter] class.
var UserNotificationCenterClass = _UserNotificationCenterClass{objc.GetClass("NSUserNotificationCenter")}

type _UserNotificationCenterClass struct {
	objc.Class
}

// An interface definition for the [UserNotificationCenter] class.
type IUserNotificationCenter interface {
	objc.IObject
}

// An object that delivers notifications from apps to the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotificationcenter?language=objc
type UserNotificationCenter struct {
	objc.Object
}

func UserNotificationCenterFrom(ptr unsafe.Pointer) UserNotificationCenter {
	return UserNotificationCenter{
		Object: objc.ObjectFrom(ptr),
	}
}

func (uc _UserNotificationCenterClass) Alloc() UserNotificationCenter {
	rv := objc.Call[UserNotificationCenter](uc, objc.Sel("alloc"))
	return rv
}

func UserNotificationCenter_Alloc() UserNotificationCenter {
	return UserNotificationCenterClass.Alloc()
}

func (uc _UserNotificationCenterClass) New() UserNotificationCenter {
	rv := objc.Call[UserNotificationCenter](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUserNotificationCenter() UserNotificationCenter {
	return UserNotificationCenterClass.New()
}

func (u_ UserNotificationCenter) Init() UserNotificationCenter {
	rv := objc.Call[UserNotificationCenter](u_, objc.Sel("init"))
	return rv
}
