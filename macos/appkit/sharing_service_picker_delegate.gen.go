// auto-generated code, do not modify
package appkit

import (
	"github.com/progrium/macdriver/objc"
)

type ISharingServicePickerDelegate interface {
	ImplementsSharingServicePickerSharingServicesForItemsProposedSharingServices() bool
	// optional
	SharingServicePickerSharingServicesForItemsProposedSharingServices(sharingServicePicker SharingServicePicker, items []objc.Object, proposedServices []SharingService) []ISharingService
	ImplementsSharingServicePickerDidChooseSharingService() bool
	// optional
	SharingServicePickerDidChooseSharingService(sharingServicePicker SharingServicePicker, service SharingService)
	ImplementsSharingServicePickerDelegateForSharingService() bool
	// optional
	SharingServicePickerDelegateForSharingService(sharingServicePicker SharingServicePicker, sharingService SharingService) ISharingServiceDelegate
}

type SharingServicePickerDelegate struct {
	_SharingServicePickerSharingServicesForItemsProposedSharingServices func(sharingServicePicker SharingServicePicker, items []objc.Object, proposedServices []SharingService) []ISharingService
	_SharingServicePickerDidChooseSharingService                        func(sharingServicePicker SharingServicePicker, service SharingService)
	_SharingServicePickerDelegateForSharingService                      func(sharingServicePicker SharingServicePicker, sharingService SharingService) ISharingServiceDelegate
}

func (di *SharingServicePickerDelegate) ImplementsSharingServicePickerSharingServicesForItemsProposedSharingServices() bool {
	return di._SharingServicePickerSharingServicesForItemsProposedSharingServices != nil
}

func (di *SharingServicePickerDelegate) SetSharingServicePickerSharingServicesForItemsProposedSharingServices(f func(sharingServicePicker SharingServicePicker, items []objc.Object, proposedServices []SharingService) []ISharingService) {
	di._SharingServicePickerSharingServicesForItemsProposedSharingServices = f
}

func (di *SharingServicePickerDelegate) SharingServicePickerSharingServicesForItemsProposedSharingServices(sharingServicePicker SharingServicePicker, items []objc.Object, proposedServices []SharingService) []ISharingService {
	return di._SharingServicePickerSharingServicesForItemsProposedSharingServices(sharingServicePicker, items, proposedServices)
}
func (di *SharingServicePickerDelegate) ImplementsSharingServicePickerDidChooseSharingService() bool {
	return di._SharingServicePickerDidChooseSharingService != nil
}

func (di *SharingServicePickerDelegate) SetSharingServicePickerDidChooseSharingService(f func(sharingServicePicker SharingServicePicker, service SharingService)) {
	di._SharingServicePickerDidChooseSharingService = f
}

func (di *SharingServicePickerDelegate) SharingServicePickerDidChooseSharingService(sharingServicePicker SharingServicePicker, service SharingService) {
	di._SharingServicePickerDidChooseSharingService(sharingServicePicker, service)
}
func (di *SharingServicePickerDelegate) ImplementsSharingServicePickerDelegateForSharingService() bool {
	return di._SharingServicePickerDelegateForSharingService != nil
}

func (di *SharingServicePickerDelegate) SetSharingServicePickerDelegateForSharingService(f func(sharingServicePicker SharingServicePicker, sharingService SharingService) ISharingServiceDelegate) {
	di._SharingServicePickerDelegateForSharingService = f
}

func (di *SharingServicePickerDelegate) SharingServicePickerDelegateForSharingService(sharingServicePicker SharingServicePicker, sharingService SharingService) ISharingServiceDelegate {
	return di._SharingServicePickerDelegateForSharingService(sharingServicePicker, sharingService)
}

type SharingServicePickerDelegateWrapper struct {
	objc.Object
}

func (s_ SharingServicePickerDelegateWrapper) ImplementsSharingServicePickerSharingServicesForItemsProposedSharingServices() bool {
	return s_.RespondsToSelector(objc.GetSelector("sharingServicePicker:sharingServicesForItems:proposedSharingServices:"))
}

func (s_ SharingServicePickerDelegateWrapper) SharingServicePickerSharingServicesForItemsProposedSharingServices(sharingServicePicker ISharingServicePicker, items []objc.IObject, proposedServices []ISharingService) []SharingService {
	rv := objc.CallMethod[[]SharingService](s_, objc.GetSelector("sharingServicePicker:sharingServicesForItems:proposedSharingServices:"), objc.ExtractPtr(sharingServicePicker), items, proposedServices)
	return rv
}

func (s_ SharingServicePickerDelegateWrapper) ImplementsSharingServicePickerDidChooseSharingService() bool {
	return s_.RespondsToSelector(objc.GetSelector("sharingServicePicker:didChooseSharingService:"))
}

func (s_ SharingServicePickerDelegateWrapper) SharingServicePickerDidChooseSharingService(sharingServicePicker ISharingServicePicker, service ISharingService) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("sharingServicePicker:didChooseSharingService:"), objc.ExtractPtr(sharingServicePicker), objc.ExtractPtr(service))
}

func (s_ SharingServicePickerDelegateWrapper) ImplementsSharingServicePickerDelegateForSharingService() bool {
	return s_.RespondsToSelector(objc.GetSelector("sharingServicePicker:delegateForSharingService:"))
}

func (s_ SharingServicePickerDelegateWrapper) SharingServicePickerDelegateForSharingService(sharingServicePicker ISharingServicePicker, sharingService ISharingService) SharingServiceDelegateWrapper {
	rv := objc.CallMethod[SharingServiceDelegateWrapper](s_, objc.GetSelector("sharingServicePicker:delegateForSharingService:"), objc.ExtractPtr(sharingServicePicker), objc.ExtractPtr(sharingService))
	return rv
}
