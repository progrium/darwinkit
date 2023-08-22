// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of methods for responding to the life cycle events of the cloud-sharing service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscloudsharingservicedelegate?language=objc
type PCloudSharingServiceDelegate interface {
	// optional
	SharingServiceDidSaveShare(sharingService SharingService, share objc.Object)
	HasSharingServiceDidSaveShare() bool

	// optional
	OptionsForSharingServiceShareProvider(cloudKitSharingService SharingService, provider foundation.ItemProvider) CloudKitSharingServiceOptions
	HasOptionsForSharingServiceShareProvider() bool
}

// A delegate implementation builder for the [PCloudSharingServiceDelegate] protocol.
type CloudSharingServiceDelegate struct {
	_SharingServiceDidSaveShare            func(sharingService SharingService, share objc.Object)
	_OptionsForSharingServiceShareProvider func(cloudKitSharingService SharingService, provider foundation.ItemProvider) CloudKitSharingServiceOptions
}

func (di *CloudSharingServiceDelegate) HasSharingServiceDidSaveShare() bool {
	return di._SharingServiceDidSaveShare != nil
}

// Tells the delegate when the cloud-sharing service saves the CloudKit share. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscloudsharingservicedelegate/1644712-sharingservice?language=objc
func (di *CloudSharingServiceDelegate) SetSharingServiceDidSaveShare(f func(sharingService SharingService, share objc.Object)) {
	di._SharingServiceDidSaveShare = f
}

// Tells the delegate when the cloud-sharing service saves the CloudKit share. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscloudsharingservicedelegate/1644712-sharingservice?language=objc
func (di *CloudSharingServiceDelegate) SharingServiceDidSaveShare(sharingService SharingService, share objc.Object) {
	di._SharingServiceDidSaveShare(sharingService, share)
}
func (di *CloudSharingServiceDelegate) HasOptionsForSharingServiceShareProvider() bool {
	return di._OptionsForSharingServiceShareProvider != nil
}

// Asks the delegate for the participant options for the cloud-sharing service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscloudsharingservicedelegate/1644694-optionsforsharingservice?language=objc
func (di *CloudSharingServiceDelegate) SetOptionsForSharingServiceShareProvider(f func(cloudKitSharingService SharingService, provider foundation.ItemProvider) CloudKitSharingServiceOptions) {
	di._OptionsForSharingServiceShareProvider = f
}

// Asks the delegate for the participant options for the cloud-sharing service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscloudsharingservicedelegate/1644694-optionsforsharingservice?language=objc
func (di *CloudSharingServiceDelegate) OptionsForSharingServiceShareProvider(cloudKitSharingService SharingService, provider foundation.ItemProvider) CloudKitSharingServiceOptions {
	return di._OptionsForSharingServiceShareProvider(cloudKitSharingService, provider)
}

// A concrete type wrapper for the [PCloudSharingServiceDelegate] protocol.
type CloudSharingServiceDelegateWrapper struct {
	objc.Object
}

func (c_ CloudSharingServiceDelegateWrapper) HasSharingServiceDidSaveShare() bool {
	return c_.RespondsToSelector(objc.Sel("sharingService:didSaveShare:"))
}

// Tells the delegate when the cloud-sharing service saves the CloudKit share. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscloudsharingservicedelegate/1644712-sharingservice?language=objc
func (c_ CloudSharingServiceDelegateWrapper) SharingServiceDidSaveShare(sharingService ISharingService, share objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("sharingService:didSaveShare:"), objc.Ptr(sharingService), objc.Ptr(share))
}

func (c_ CloudSharingServiceDelegateWrapper) HasOptionsForSharingServiceShareProvider() bool {
	return c_.RespondsToSelector(objc.Sel("optionsForSharingService:shareProvider:"))
}

// Asks the delegate for the participant options for the cloud-sharing service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscloudsharingservicedelegate/1644694-optionsforsharingservice?language=objc
func (c_ CloudSharingServiceDelegateWrapper) OptionsForSharingServiceShareProvider(cloudKitSharingService ISharingService, provider foundation.IItemProvider) CloudKitSharingServiceOptions {
	rv := objc.Call[CloudKitSharingServiceOptions](c_, objc.Sel("optionsForSharingService:shareProvider:"), objc.Ptr(cloudKitSharingService), objc.Ptr(provider))
	return rv
}
