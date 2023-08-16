// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// An interface that provides the content to share from the macOS share sheet. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertoolbaritemdelegate?language=objc
type PSharingServicePickerToolbarItemDelegate interface {
	// optional
	ItemsForSharingServicePickerToolbarItem(pickerToolbarItem SharingServicePickerToolbarItem) []objc.IObject
	HasItemsForSharingServicePickerToolbarItem() bool
}

// A delegate implementation builder for the [PSharingServicePickerToolbarItemDelegate] protocol.
type SharingServicePickerToolbarItemDelegate struct {
	_ItemsForSharingServicePickerToolbarItem func(pickerToolbarItem SharingServicePickerToolbarItem) []objc.IObject
}

func (di *SharingServicePickerToolbarItemDelegate) HasItemsForSharingServicePickerToolbarItem() bool {
	return di._ItemsForSharingServicePickerToolbarItem != nil
}

// Asks the delegate for the items to share. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertoolbaritemdelegate/3365983-itemsforsharingservicepickertool?language=objc
func (di *SharingServicePickerToolbarItemDelegate) SetItemsForSharingServicePickerToolbarItem(f func(pickerToolbarItem SharingServicePickerToolbarItem) []objc.IObject) {
	di._ItemsForSharingServicePickerToolbarItem = f
}

// Asks the delegate for the items to share. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertoolbaritemdelegate/3365983-itemsforsharingservicepickertool?language=objc
func (di *SharingServicePickerToolbarItemDelegate) ItemsForSharingServicePickerToolbarItem(pickerToolbarItem SharingServicePickerToolbarItem) []objc.IObject {
	return di._ItemsForSharingServicePickerToolbarItem(pickerToolbarItem)
}

// A concrete type wrapper for the [PSharingServicePickerToolbarItemDelegate] protocol.
type SharingServicePickerToolbarItemDelegateWrapper struct {
	objc.Object
}

func (s_ SharingServicePickerToolbarItemDelegateWrapper) HasItemsForSharingServicePickerToolbarItem() bool {
	return s_.RespondsToSelector(objc.Sel("itemsForSharingServicePickerToolbarItem:"))
}

// Asks the delegate for the items to share. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertoolbaritemdelegate/3365983-itemsforsharingservicepickertool?language=objc
func (s_ SharingServicePickerToolbarItemDelegateWrapper) ItemsForSharingServicePickerToolbarItem(pickerToolbarItem ISharingServicePickerToolbarItem) []objc.Object {
	rv := objc.Call[[]objc.Object](s_, objc.Sel("itemsForSharingServicePickerToolbarItem:"), objc.Ptr(pickerToolbarItem))
	return rv
}
