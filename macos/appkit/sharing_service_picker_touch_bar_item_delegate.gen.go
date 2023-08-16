// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that a sharing service picker item delegate uses to provide a list of items eligible for sharing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertouchbaritemdelegate?language=objc
type PSharingServicePickerTouchBarItemDelegate interface {
	// optional
	ItemsForSharingServicePickerTouchBarItem(pickerTouchBarItem SharingServicePickerTouchBarItem) []objc.IObject
	HasItemsForSharingServicePickerTouchBarItem() bool
}

// A delegate implementation builder for the [PSharingServicePickerTouchBarItemDelegate] protocol.
type SharingServicePickerTouchBarItemDelegate struct {
	_ItemsForSharingServicePickerTouchBarItem func(pickerTouchBarItem SharingServicePickerTouchBarItem) []objc.IObject
}

func (di *SharingServicePickerTouchBarItemDelegate) HasItemsForSharingServicePickerTouchBarItem() bool {
	return di._ItemsForSharingServicePickerTouchBarItem != nil
}

// Asks the delegate for items that represent the objects to be shared. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertouchbaritemdelegate/2646997-itemsforsharingservicepickertouc?language=objc
func (di *SharingServicePickerTouchBarItemDelegate) SetItemsForSharingServicePickerTouchBarItem(f func(pickerTouchBarItem SharingServicePickerTouchBarItem) []objc.IObject) {
	di._ItemsForSharingServicePickerTouchBarItem = f
}

// Asks the delegate for items that represent the objects to be shared. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertouchbaritemdelegate/2646997-itemsforsharingservicepickertouc?language=objc
func (di *SharingServicePickerTouchBarItemDelegate) ItemsForSharingServicePickerTouchBarItem(pickerTouchBarItem SharingServicePickerTouchBarItem) []objc.IObject {
	return di._ItemsForSharingServicePickerTouchBarItem(pickerTouchBarItem)
}

// A concrete type wrapper for the [PSharingServicePickerTouchBarItemDelegate] protocol.
type SharingServicePickerTouchBarItemDelegateWrapper struct {
	objc.Object
}

func (s_ SharingServicePickerTouchBarItemDelegateWrapper) HasItemsForSharingServicePickerTouchBarItem() bool {
	return s_.RespondsToSelector(objc.Sel("itemsForSharingServicePickerTouchBarItem:"))
}

// Asks the delegate for items that represent the objects to be shared. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertouchbaritemdelegate/2646997-itemsforsharingservicepickertouc?language=objc
func (s_ SharingServicePickerTouchBarItemDelegateWrapper) ItemsForSharingServicePickerTouchBarItem(pickerTouchBarItem ISharingServicePickerTouchBarItem) []objc.Object {
	rv := objc.Call[[]objc.Object](s_, objc.Sel("itemsForSharingServicePickerTouchBarItem:"), objc.Ptr(pickerTouchBarItem))
	return rv
}
