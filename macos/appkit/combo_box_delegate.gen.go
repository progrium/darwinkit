// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of optional methods implemented by delegates of combo box objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxdelegate?language=objc
type PComboBoxDelegate interface {
	PTextFieldDelegate
	// optional
	ComboBoxSelectionDidChange(notification foundation.Notification)
	HasComboBoxSelectionDidChange() bool

	// optional
	ComboBoxWillPopUp(notification foundation.Notification)
	HasComboBoxWillPopUp() bool

	// optional
	ComboBoxWillDismiss(notification foundation.Notification)
	HasComboBoxWillDismiss() bool

	// optional
	ComboBoxSelectionIsChanging(notification foundation.Notification)
	HasComboBoxSelectionIsChanging() bool
}

// A delegate implementation builder for the [PComboBoxDelegate] protocol.
type ComboBoxDelegate struct {
	TextFieldDelegate
	_ComboBoxSelectionDidChange  func(notification foundation.Notification)
	_ComboBoxWillPopUp           func(notification foundation.Notification)
	_ComboBoxWillDismiss         func(notification foundation.Notification)
	_ComboBoxSelectionIsChanging func(notification foundation.Notification)
}

func (di *ComboBoxDelegate) HasComboBoxSelectionDidChange() bool {
	return di._ComboBoxSelectionDidChange != nil
}

// Informs the delegate that the pop-up list selection has finished changing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxdelegate/1436769-comboboxselectiondidchange?language=objc
func (di *ComboBoxDelegate) SetComboBoxSelectionDidChange(f func(notification foundation.Notification)) {
	di._ComboBoxSelectionDidChange = f
}

// Informs the delegate that the pop-up list selection has finished changing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxdelegate/1436769-comboboxselectiondidchange?language=objc
func (di *ComboBoxDelegate) ComboBoxSelectionDidChange(notification foundation.Notification) {
	di._ComboBoxSelectionDidChange(notification)
}
func (di *ComboBoxDelegate) HasComboBoxWillPopUp() bool {
	return di._ComboBoxWillPopUp != nil
}

// Informs the delegate that the pop-up list is about to be displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxdelegate/1436784-comboboxwillpopup?language=objc
func (di *ComboBoxDelegate) SetComboBoxWillPopUp(f func(notification foundation.Notification)) {
	di._ComboBoxWillPopUp = f
}

// Informs the delegate that the pop-up list is about to be displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxdelegate/1436784-comboboxwillpopup?language=objc
func (di *ComboBoxDelegate) ComboBoxWillPopUp(notification foundation.Notification) {
	di._ComboBoxWillPopUp(notification)
}
func (di *ComboBoxDelegate) HasComboBoxWillDismiss() bool {
	return di._ComboBoxWillDismiss != nil
}

// Informs the delegate that the pop-up list is about to be dismissed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxdelegate/1436763-comboboxwilldismiss?language=objc
func (di *ComboBoxDelegate) SetComboBoxWillDismiss(f func(notification foundation.Notification)) {
	di._ComboBoxWillDismiss = f
}

// Informs the delegate that the pop-up list is about to be dismissed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxdelegate/1436763-comboboxwilldismiss?language=objc
func (di *ComboBoxDelegate) ComboBoxWillDismiss(notification foundation.Notification) {
	di._ComboBoxWillDismiss(notification)
}
func (di *ComboBoxDelegate) HasComboBoxSelectionIsChanging() bool {
	return di._ComboBoxSelectionIsChanging != nil
}

// Informs the delegate that the pop-up list selection is changing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxdelegate/1436715-comboboxselectionischanging?language=objc
func (di *ComboBoxDelegate) SetComboBoxSelectionIsChanging(f func(notification foundation.Notification)) {
	di._ComboBoxSelectionIsChanging = f
}

// Informs the delegate that the pop-up list selection is changing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxdelegate/1436715-comboboxselectionischanging?language=objc
func (di *ComboBoxDelegate) ComboBoxSelectionIsChanging(notification foundation.Notification) {
	di._ComboBoxSelectionIsChanging(notification)
}

// A concrete type wrapper for the [PComboBoxDelegate] protocol.
type ComboBoxDelegateWrapper struct {
	TextFieldDelegateWrapper
}

func (c_ ComboBoxDelegateWrapper) HasComboBoxSelectionDidChange() bool {
	return c_.RespondsToSelector(objc.Sel("comboBoxSelectionDidChange:"))
}

// Informs the delegate that the pop-up list selection has finished changing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxdelegate/1436769-comboboxselectiondidchange?language=objc
func (c_ ComboBoxDelegateWrapper) ComboBoxSelectionDidChange(notification foundation.INotification) {
	objc.Call[objc.Void](c_, objc.Sel("comboBoxSelectionDidChange:"), objc.Ptr(notification))
}

func (c_ ComboBoxDelegateWrapper) HasComboBoxWillPopUp() bool {
	return c_.RespondsToSelector(objc.Sel("comboBoxWillPopUp:"))
}

// Informs the delegate that the pop-up list is about to be displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxdelegate/1436784-comboboxwillpopup?language=objc
func (c_ ComboBoxDelegateWrapper) ComboBoxWillPopUp(notification foundation.INotification) {
	objc.Call[objc.Void](c_, objc.Sel("comboBoxWillPopUp:"), objc.Ptr(notification))
}

func (c_ ComboBoxDelegateWrapper) HasComboBoxWillDismiss() bool {
	return c_.RespondsToSelector(objc.Sel("comboBoxWillDismiss:"))
}

// Informs the delegate that the pop-up list is about to be dismissed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxdelegate/1436763-comboboxwilldismiss?language=objc
func (c_ ComboBoxDelegateWrapper) ComboBoxWillDismiss(notification foundation.INotification) {
	objc.Call[objc.Void](c_, objc.Sel("comboBoxWillDismiss:"), objc.Ptr(notification))
}

func (c_ ComboBoxDelegateWrapper) HasComboBoxSelectionIsChanging() bool {
	return c_.RespondsToSelector(objc.Sel("comboBoxSelectionIsChanging:"))
}

// Informs the delegate that the pop-up list selection is changing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxdelegate/1436715-comboboxselectionischanging?language=objc
func (c_ ComboBoxDelegateWrapper) ComboBoxSelectionIsChanging(notification foundation.INotification) {
	objc.Call[objc.Void](c_, objc.Sel("comboBoxSelectionIsChanging:"), objc.Ptr(notification))
}
