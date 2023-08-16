// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// The interface through which a user activity instance notifies its delegate of updates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivitydelegate?language=objc
type PUserActivityDelegate interface {
	// optional
	UserActivityWasContinued(userActivity UserActivity)
	HasUserActivityWasContinued() bool

	// optional
	UserActivityWillSave(userActivity UserActivity)
	HasUserActivityWillSave() bool

	// optional
	UserActivityDidReceiveInputStreamOutputStream(userActivity UserActivity, inputStream InputStream, outputStream OutputStream)
	HasUserActivityDidReceiveInputStreamOutputStream() bool
}

// A delegate implementation builder for the [PUserActivityDelegate] protocol.
type UserActivityDelegate struct {
	_UserActivityWasContinued                      func(userActivity UserActivity)
	_UserActivityWillSave                          func(userActivity UserActivity)
	_UserActivityDidReceiveInputStreamOutputStream func(userActivity UserActivity, inputStream InputStream, outputStream OutputStream)
}

func (di *UserActivityDelegate) HasUserActivityWasContinued() bool {
	return di._UserActivityWasContinued != nil
}

// Notifies the delegate that the user activity was continued on another device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivitydelegate/1413276-useractivitywascontinued?language=objc
func (di *UserActivityDelegate) SetUserActivityWasContinued(f func(userActivity UserActivity)) {
	di._UserActivityWasContinued = f
}

// Notifies the delegate that the user activity was continued on another device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivitydelegate/1413276-useractivitywascontinued?language=objc
func (di *UserActivityDelegate) UserActivityWasContinued(userActivity UserActivity) {
	di._UserActivityWasContinued(userActivity)
}
func (di *UserActivityDelegate) HasUserActivityWillSave() bool {
	return di._UserActivityWillSave != nil
}

// Notifies the delegate that the user activity will be saved to be continued or persisted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivitydelegate/1414848-useractivitywillsave?language=objc
func (di *UserActivityDelegate) SetUserActivityWillSave(f func(userActivity UserActivity)) {
	di._UserActivityWillSave = f
}

// Notifies the delegate that the user activity will be saved to be continued or persisted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivitydelegate/1414848-useractivitywillsave?language=objc
func (di *UserActivityDelegate) UserActivityWillSave(userActivity UserActivity) {
	di._UserActivityWillSave(userActivity)
}
func (di *UserActivityDelegate) HasUserActivityDidReceiveInputStreamOutputStream() bool {
	return di._UserActivityDidReceiveInputStreamOutputStream != nil
}

// Notifies the user activity delegate that an input and output streams are available to open. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivitydelegate/1407386-useractivity?language=objc
func (di *UserActivityDelegate) SetUserActivityDidReceiveInputStreamOutputStream(f func(userActivity UserActivity, inputStream InputStream, outputStream OutputStream)) {
	di._UserActivityDidReceiveInputStreamOutputStream = f
}

// Notifies the user activity delegate that an input and output streams are available to open. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivitydelegate/1407386-useractivity?language=objc
func (di *UserActivityDelegate) UserActivityDidReceiveInputStreamOutputStream(userActivity UserActivity, inputStream InputStream, outputStream OutputStream) {
	di._UserActivityDidReceiveInputStreamOutputStream(userActivity, inputStream, outputStream)
}

// A concrete type wrapper for the [PUserActivityDelegate] protocol.
type UserActivityDelegateWrapper struct {
	objc.Object
}

func (u_ UserActivityDelegateWrapper) HasUserActivityWasContinued() bool {
	return u_.RespondsToSelector(objc.Sel("userActivityWasContinued:"))
}

// Notifies the delegate that the user activity was continued on another device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivitydelegate/1413276-useractivitywascontinued?language=objc
func (u_ UserActivityDelegateWrapper) UserActivityWasContinued(userActivity IUserActivity) {
	objc.Call[objc.Void](u_, objc.Sel("userActivityWasContinued:"), objc.Ptr(userActivity))
}

func (u_ UserActivityDelegateWrapper) HasUserActivityWillSave() bool {
	return u_.RespondsToSelector(objc.Sel("userActivityWillSave:"))
}

// Notifies the delegate that the user activity will be saved to be continued or persisted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivitydelegate/1414848-useractivitywillsave?language=objc
func (u_ UserActivityDelegateWrapper) UserActivityWillSave(userActivity IUserActivity) {
	objc.Call[objc.Void](u_, objc.Sel("userActivityWillSave:"), objc.Ptr(userActivity))
}

func (u_ UserActivityDelegateWrapper) HasUserActivityDidReceiveInputStreamOutputStream() bool {
	return u_.RespondsToSelector(objc.Sel("userActivity:didReceiveInputStream:outputStream:"))
}

// Notifies the user activity delegate that an input and output streams are available to open. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivitydelegate/1407386-useractivity?language=objc
func (u_ UserActivityDelegateWrapper) UserActivityDidReceiveInputStreamOutputStream(userActivity IUserActivity, inputStream IInputStream, outputStream IOutputStream) {
	objc.Call[objc.Void](u_, objc.Sel("userActivity:didReceiveInputStream:outputStream:"), objc.Ptr(userActivity), objc.Ptr(inputStream), objc.Ptr(outputStream))
}
