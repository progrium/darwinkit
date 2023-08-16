// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A protocol that marks classes to restore the state of your app to continue a user activity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuseractivityrestoring?language=objc
type PUserActivityRestoring interface {
	// optional
	RestoreUserActivityState(userActivity foundation.UserActivity)
	HasRestoreUserActivityState() bool
}

// A concrete type wrapper for the [PUserActivityRestoring] protocol.
type UserActivityRestoringWrapper struct {
	objc.Object
}

func (u_ UserActivityRestoringWrapper) HasRestoreUserActivityState() bool {
	return u_.RespondsToSelector(objc.Sel("restoreUserActivityState:"))
}

// Restores the state necessary to continue the specified user activity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuseractivityrestoring/3022485-restoreuseractivitystate?language=objc
func (u_ UserActivityRestoringWrapper) RestoreUserActivityState(userActivity foundation.IUserActivity) {
	objc.Call[objc.Void](u_, objc.Sel("restoreUserActivityState:"), objc.Ptr(userActivity))
}
