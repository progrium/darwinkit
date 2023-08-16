// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// An interface for managing content for the macOS share sheet. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickerdelegate?language=objc
type PSharingServicePickerDelegate interface {
	// optional
	SharingServicePickerDidChooseSharingService(sharingServicePicker SharingServicePicker, service SharingService)
	HasSharingServicePickerDidChooseSharingService() bool
}

// A delegate implementation builder for the [PSharingServicePickerDelegate] protocol.
type SharingServicePickerDelegate struct {
	_SharingServicePickerDidChooseSharingService func(sharingServicePicker SharingServicePicker, service SharingService)
}

func (di *SharingServicePickerDelegate) HasSharingServicePickerDidChooseSharingService() bool {
	return di._SharingServicePickerDidChooseSharingService != nil
}

// Tells the delegate that the person selected a sharing service for the current item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickerdelegate/1402610-sharingservicepicker?language=objc
func (di *SharingServicePickerDelegate) SetSharingServicePickerDidChooseSharingService(f func(sharingServicePicker SharingServicePicker, service SharingService)) {
	di._SharingServicePickerDidChooseSharingService = f
}

// Tells the delegate that the person selected a sharing service for the current item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickerdelegate/1402610-sharingservicepicker?language=objc
func (di *SharingServicePickerDelegate) SharingServicePickerDidChooseSharingService(sharingServicePicker SharingServicePicker, service SharingService) {
	di._SharingServicePickerDidChooseSharingService(sharingServicePicker, service)
}

// A concrete type wrapper for the [PSharingServicePickerDelegate] protocol.
type SharingServicePickerDelegateWrapper struct {
	objc.Object
}

func (s_ SharingServicePickerDelegateWrapper) HasSharingServicePickerDidChooseSharingService() bool {
	return s_.RespondsToSelector(objc.Sel("sharingServicePicker:didChooseSharingService:"))
}

// Tells the delegate that the person selected a sharing service for the current item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickerdelegate/1402610-sharingservicepicker?language=objc
func (s_ SharingServicePickerDelegateWrapper) SharingServicePickerDidChooseSharingService(sharingServicePicker ISharingServicePicker, service ISharingService) {
	objc.Call[objc.Void](s_, objc.Sel("sharingServicePicker:didChooseSharingService:"), objc.Ptr(sharingServicePicker), objc.Ptr(service))
}
