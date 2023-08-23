// AUTO-GENERATED CODE, DO NOT MODIFY

package contactsui

import (
	"github.com/progrium/macdriver/objc"
)

// The methods that you implement to respond to contact-picker user events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contactsui/cncontactpickerdelegate?language=objc
type PContactPickerDelegate interface {
	// optional
	ContactPickerDidClose(picker ContactPicker)
	HasContactPickerDidClose() bool

	// optional
	ContactPickerWillClose(picker ContactPicker)
	HasContactPickerWillClose() bool
}

// A delegate implementation builder for the [PContactPickerDelegate] protocol.
type ContactPickerDelegate struct {
	_ContactPickerDidClose  func(picker ContactPicker)
	_ContactPickerWillClose func(picker ContactPicker)
}

func (di *ContactPickerDelegate) HasContactPickerDidClose() bool {
	return di._ContactPickerDidClose != nil
}

// In macOS, called when the contact picker’s popover has closed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contactsui/cncontactpickerdelegate/1522584-contactpickerdidclose?language=objc
func (di *ContactPickerDelegate) SetContactPickerDidClose(f func(picker ContactPicker)) {
	di._ContactPickerDidClose = f
}

// In macOS, called when the contact picker’s popover has closed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contactsui/cncontactpickerdelegate/1522584-contactpickerdidclose?language=objc
func (di *ContactPickerDelegate) ContactPickerDidClose(picker ContactPicker) {
	di._ContactPickerDidClose(picker)
}
func (di *ContactPickerDelegate) HasContactPickerWillClose() bool {
	return di._ContactPickerWillClose != nil
}

// In macOS, called when the contact picker’s popover is about to close. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contactsui/cncontactpickerdelegate/1522594-contactpickerwillclose?language=objc
func (di *ContactPickerDelegate) SetContactPickerWillClose(f func(picker ContactPicker)) {
	di._ContactPickerWillClose = f
}

// In macOS, called when the contact picker’s popover is about to close. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contactsui/cncontactpickerdelegate/1522594-contactpickerwillclose?language=objc
func (di *ContactPickerDelegate) ContactPickerWillClose(picker ContactPicker) {
	di._ContactPickerWillClose(picker)
}

// A concrete type wrapper for the [PContactPickerDelegate] protocol.
type ContactPickerDelegateWrapper struct {
	objc.Object
}

func (c_ ContactPickerDelegateWrapper) HasContactPickerDidClose() bool {
	return c_.RespondsToSelector(objc.Sel("contactPickerDidClose:"))
}

// In macOS, called when the contact picker’s popover has closed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contactsui/cncontactpickerdelegate/1522584-contactpickerdidclose?language=objc
func (c_ ContactPickerDelegateWrapper) ContactPickerDidClose(picker IContactPicker) {
	objc.Call[objc.Void](c_, objc.Sel("contactPickerDidClose:"), objc.Ptr(picker))
}

func (c_ ContactPickerDelegateWrapper) HasContactPickerWillClose() bool {
	return c_.RespondsToSelector(objc.Sel("contactPickerWillClose:"))
}

// In macOS, called when the contact picker’s popover is about to close. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contactsui/cncontactpickerdelegate/1522594-contactpickerwillclose?language=objc
func (c_ ContactPickerDelegateWrapper) ContactPickerWillClose(picker IContactPicker) {
	objc.Call[objc.Void](c_, objc.Sel("contactPickerWillClose:"), objc.Ptr(picker))
}
