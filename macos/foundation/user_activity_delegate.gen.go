// auto-generated code, do not modify
package foundation

import (
	"github.com/progrium/macdriver/objc"
)

type IUserActivityDelegate interface {
	ImplementsUserActivityDidReceiveInputStreamOutputStream() bool
	// optional
	UserActivityDidReceiveInputStreamOutputStream(userActivity UserActivity, inputStream InputStream, outputStream OutputStream)
	ImplementsUserActivityWasContinued() bool
	// optional
	UserActivityWasContinued(userActivity UserActivity)
	ImplementsUserActivityWillSave() bool
	// optional
	UserActivityWillSave(userActivity UserActivity)
}

type UserActivityDelegate struct {
	_UserActivityDidReceiveInputStreamOutputStream func(userActivity UserActivity, inputStream InputStream, outputStream OutputStream)
	_UserActivityWasContinued                      func(userActivity UserActivity)
	_UserActivityWillSave                          func(userActivity UserActivity)
}

func (di *UserActivityDelegate) ImplementsUserActivityDidReceiveInputStreamOutputStream() bool {
	return di._UserActivityDidReceiveInputStreamOutputStream != nil
}

func (di *UserActivityDelegate) SetUserActivityDidReceiveInputStreamOutputStream(f func(userActivity UserActivity, inputStream InputStream, outputStream OutputStream)) {
	di._UserActivityDidReceiveInputStreamOutputStream = f
}

func (di *UserActivityDelegate) UserActivityDidReceiveInputStreamOutputStream(userActivity UserActivity, inputStream InputStream, outputStream OutputStream) {
	di._UserActivityDidReceiveInputStreamOutputStream(userActivity, inputStream, outputStream)
}
func (di *UserActivityDelegate) ImplementsUserActivityWasContinued() bool {
	return di._UserActivityWasContinued != nil
}

func (di *UserActivityDelegate) SetUserActivityWasContinued(f func(userActivity UserActivity)) {
	di._UserActivityWasContinued = f
}

func (di *UserActivityDelegate) UserActivityWasContinued(userActivity UserActivity) {
	di._UserActivityWasContinued(userActivity)
}
func (di *UserActivityDelegate) ImplementsUserActivityWillSave() bool {
	return di._UserActivityWillSave != nil
}

func (di *UserActivityDelegate) SetUserActivityWillSave(f func(userActivity UserActivity)) {
	di._UserActivityWillSave = f
}

func (di *UserActivityDelegate) UserActivityWillSave(userActivity UserActivity) {
	di._UserActivityWillSave(userActivity)
}

type UserActivityDelegateWrapper struct {
	objc.Object
}

func (u_ UserActivityDelegateWrapper) ImplementsUserActivityDidReceiveInputStreamOutputStream() bool {
	return u_.RespondsToSelector(objc.GetSelector("userActivity:didReceiveInputStream:outputStream:"))
}

func (u_ UserActivityDelegateWrapper) UserActivityDidReceiveInputStreamOutputStream(userActivity IUserActivity, inputStream IInputStream, outputStream IOutputStream) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("userActivity:didReceiveInputStream:outputStream:"), objc.ExtractPtr(userActivity), objc.ExtractPtr(inputStream), objc.ExtractPtr(outputStream))
}

func (u_ UserActivityDelegateWrapper) ImplementsUserActivityWasContinued() bool {
	return u_.RespondsToSelector(objc.GetSelector("userActivityWasContinued:"))
}

func (u_ UserActivityDelegateWrapper) UserActivityWasContinued(userActivity IUserActivity) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("userActivityWasContinued:"), objc.ExtractPtr(userActivity))
}

func (u_ UserActivityDelegateWrapper) ImplementsUserActivityWillSave() bool {
	return u_.RespondsToSelector(objc.GetSelector("userActivityWillSave:"))
}

func (u_ UserActivityDelegateWrapper) UserActivityWillSave(userActivity IUserActivity) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("userActivityWillSave:"), objc.ExtractPtr(userActivity))
}
