// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UserNotificationAction] class.
var UserNotificationActionClass = _UserNotificationActionClass{objc.GetClass("NSUserNotificationAction")}

type _UserNotificationActionClass struct {
	objc.Class
}

// An interface definition for the [UserNotificationAction] class.
type IUserNotificationAction interface {
	objc.IObject
}

// An action that the user can take in response to receiving a notification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotificationaction?language=objc
type UserNotificationAction struct {
	objc.Object
}

func UserNotificationActionFrom(ptr unsafe.Pointer) UserNotificationAction {
	return UserNotificationAction{
		Object: objc.ObjectFrom(ptr),
	}
}

func (uc _UserNotificationActionClass) Alloc() UserNotificationAction {
	rv := objc.Call[UserNotificationAction](uc, objc.Sel("alloc"))
	return rv
}

func UserNotificationAction_Alloc() UserNotificationAction {
	return UserNotificationActionClass.Alloc()
}

func (uc _UserNotificationActionClass) New() UserNotificationAction {
	rv := objc.Call[UserNotificationAction](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUserNotificationAction() UserNotificationAction {
	return UserNotificationActionClass.New()
}

func (u_ UserNotificationAction) Init() UserNotificationAction {
	rv := objc.Call[UserNotificationAction](u_, objc.Sel("init"))
	return rv
}
