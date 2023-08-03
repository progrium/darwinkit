// auto-generated code, do not modify
package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type IControlTextEditingDelegate interface {
	ImplementsControlIsValidObject() bool
	// optional
	ControlIsValidObject(control Control, obj objc.Object) bool
	ImplementsControlDidFailToValidatePartialStringErrorDescription() bool
	// optional
	ControlDidFailToValidatePartialStringErrorDescription(control Control, string_ string, error string)
	ImplementsControlDidFailToFormatStringErrorDescription() bool
	// optional
	ControlDidFailToFormatStringErrorDescription(control Control, string_ string, error string) bool
	ImplementsControlTextShouldBeginEditing() bool
	// optional
	ControlTextShouldBeginEditing(control Control, fieldEditor Text) bool
	ImplementsControlTextShouldEndEditing() bool
	// optional
	ControlTextShouldEndEditing(control Control, fieldEditor Text) bool
	ImplementsControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem() bool
	// optional
	ControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem(control Control, textView TextView, words []string, charRange foundation.Range, index *int) []string
	ImplementsControlTextViewDoCommandBySelector() bool
	// optional
	ControlTextViewDoCommandBySelector(control Control, textView TextView, commandSelector objc.Selector) bool
	ImplementsControlTextDidBeginEditing() bool
	// optional
	ControlTextDidBeginEditing(obj foundation.Notification)
	ImplementsControlTextDidChange() bool
	// optional
	ControlTextDidChange(obj foundation.Notification)
	ImplementsControlTextDidEndEditing() bool
	// optional
	ControlTextDidEndEditing(obj foundation.Notification)
}

type ControlTextEditingDelegate struct {
	_ControlIsValidObject                                             func(control Control, obj objc.Object) bool
	_ControlDidFailToValidatePartialStringErrorDescription            func(control Control, string_ string, error string)
	_ControlDidFailToFormatStringErrorDescription                     func(control Control, string_ string, error string) bool
	_ControlTextShouldBeginEditing                                    func(control Control, fieldEditor Text) bool
	_ControlTextShouldEndEditing                                      func(control Control, fieldEditor Text) bool
	_ControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem func(control Control, textView TextView, words []string, charRange foundation.Range, index *int) []string
	_ControlTextViewDoCommandBySelector                               func(control Control, textView TextView, commandSelector objc.Selector) bool
	_ControlTextDidBeginEditing                                       func(obj foundation.Notification)
	_ControlTextDidChange                                             func(obj foundation.Notification)
	_ControlTextDidEndEditing                                         func(obj foundation.Notification)
}

func (di *ControlTextEditingDelegate) ImplementsControlIsValidObject() bool {
	return di._ControlIsValidObject != nil
}

func (di *ControlTextEditingDelegate) SetControlIsValidObject(f func(control Control, obj objc.Object) bool) {
	di._ControlIsValidObject = f
}

func (di *ControlTextEditingDelegate) ControlIsValidObject(control Control, obj objc.Object) bool {
	return di._ControlIsValidObject(control, obj)
}
func (di *ControlTextEditingDelegate) ImplementsControlDidFailToValidatePartialStringErrorDescription() bool {
	return di._ControlDidFailToValidatePartialStringErrorDescription != nil
}

func (di *ControlTextEditingDelegate) SetControlDidFailToValidatePartialStringErrorDescription(f func(control Control, string_ string, error string)) {
	di._ControlDidFailToValidatePartialStringErrorDescription = f
}

func (di *ControlTextEditingDelegate) ControlDidFailToValidatePartialStringErrorDescription(control Control, string_ string, error string) {
	di._ControlDidFailToValidatePartialStringErrorDescription(control, string_, error)
}
func (di *ControlTextEditingDelegate) ImplementsControlDidFailToFormatStringErrorDescription() bool {
	return di._ControlDidFailToFormatStringErrorDescription != nil
}

func (di *ControlTextEditingDelegate) SetControlDidFailToFormatStringErrorDescription(f func(control Control, string_ string, error string) bool) {
	di._ControlDidFailToFormatStringErrorDescription = f
}

func (di *ControlTextEditingDelegate) ControlDidFailToFormatStringErrorDescription(control Control, string_ string, error string) bool {
	return di._ControlDidFailToFormatStringErrorDescription(control, string_, error)
}
func (di *ControlTextEditingDelegate) ImplementsControlTextShouldBeginEditing() bool {
	return di._ControlTextShouldBeginEditing != nil
}

func (di *ControlTextEditingDelegate) SetControlTextShouldBeginEditing(f func(control Control, fieldEditor Text) bool) {
	di._ControlTextShouldBeginEditing = f
}

func (di *ControlTextEditingDelegate) ControlTextShouldBeginEditing(control Control, fieldEditor Text) bool {
	return di._ControlTextShouldBeginEditing(control, fieldEditor)
}
func (di *ControlTextEditingDelegate) ImplementsControlTextShouldEndEditing() bool {
	return di._ControlTextShouldEndEditing != nil
}

func (di *ControlTextEditingDelegate) SetControlTextShouldEndEditing(f func(control Control, fieldEditor Text) bool) {
	di._ControlTextShouldEndEditing = f
}

func (di *ControlTextEditingDelegate) ControlTextShouldEndEditing(control Control, fieldEditor Text) bool {
	return di._ControlTextShouldEndEditing(control, fieldEditor)
}
func (di *ControlTextEditingDelegate) ImplementsControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem() bool {
	return di._ControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem != nil
}

func (di *ControlTextEditingDelegate) SetControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem(f func(control Control, textView TextView, words []string, charRange foundation.Range, index *int) []string) {
	di._ControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem = f
}

func (di *ControlTextEditingDelegate) ControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem(control Control, textView TextView, words []string, charRange foundation.Range, index *int) []string {
	return di._ControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem(control, textView, words, charRange, index)
}
func (di *ControlTextEditingDelegate) ImplementsControlTextViewDoCommandBySelector() bool {
	return di._ControlTextViewDoCommandBySelector != nil
}

func (di *ControlTextEditingDelegate) SetControlTextViewDoCommandBySelector(f func(control Control, textView TextView, commandSelector objc.Selector) bool) {
	di._ControlTextViewDoCommandBySelector = f
}

func (di *ControlTextEditingDelegate) ControlTextViewDoCommandBySelector(control Control, textView TextView, commandSelector objc.Selector) bool {
	return di._ControlTextViewDoCommandBySelector(control, textView, commandSelector)
}
func (di *ControlTextEditingDelegate) ImplementsControlTextDidBeginEditing() bool {
	return di._ControlTextDidBeginEditing != nil
}

func (di *ControlTextEditingDelegate) SetControlTextDidBeginEditing(f func(obj foundation.Notification)) {
	di._ControlTextDidBeginEditing = f
}

func (di *ControlTextEditingDelegate) ControlTextDidBeginEditing(obj foundation.Notification) {
	di._ControlTextDidBeginEditing(obj)
}
func (di *ControlTextEditingDelegate) ImplementsControlTextDidChange() bool {
	return di._ControlTextDidChange != nil
}

func (di *ControlTextEditingDelegate) SetControlTextDidChange(f func(obj foundation.Notification)) {
	di._ControlTextDidChange = f
}

func (di *ControlTextEditingDelegate) ControlTextDidChange(obj foundation.Notification) {
	di._ControlTextDidChange(obj)
}
func (di *ControlTextEditingDelegate) ImplementsControlTextDidEndEditing() bool {
	return di._ControlTextDidEndEditing != nil
}

func (di *ControlTextEditingDelegate) SetControlTextDidEndEditing(f func(obj foundation.Notification)) {
	di._ControlTextDidEndEditing = f
}

func (di *ControlTextEditingDelegate) ControlTextDidEndEditing(obj foundation.Notification) {
	di._ControlTextDidEndEditing(obj)
}

type ControlTextEditingDelegateWrapper struct {
	objc.Object
}

func (c_ ControlTextEditingDelegateWrapper) ImplementsControlIsValidObject() bool {
	return c_.RespondsToSelector(objc.GetSelector("control:isValidObject:"))
}

func (c_ ControlTextEditingDelegateWrapper) ControlIsValidObject(control IControl, obj objc.IObject) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("control:isValidObject:"), objc.ExtractPtr(control), objc.ExtractPtr(obj))
	return rv
}

func (c_ ControlTextEditingDelegateWrapper) ImplementsControlDidFailToValidatePartialStringErrorDescription() bool {
	return c_.RespondsToSelector(objc.GetSelector("control:didFailToValidatePartialString:errorDescription:"))
}

func (c_ ControlTextEditingDelegateWrapper) ControlDidFailToValidatePartialStringErrorDescription(control IControl, string_ string, error string) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("control:didFailToValidatePartialString:errorDescription:"), objc.ExtractPtr(control), string_, error)
}

func (c_ ControlTextEditingDelegateWrapper) ImplementsControlDidFailToFormatStringErrorDescription() bool {
	return c_.RespondsToSelector(objc.GetSelector("control:didFailToFormatString:errorDescription:"))
}

func (c_ ControlTextEditingDelegateWrapper) ControlDidFailToFormatStringErrorDescription(control IControl, string_ string, error string) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("control:didFailToFormatString:errorDescription:"), objc.ExtractPtr(control), string_, error)
	return rv
}

func (c_ ControlTextEditingDelegateWrapper) ImplementsControlTextShouldBeginEditing() bool {
	return c_.RespondsToSelector(objc.GetSelector("control:textShouldBeginEditing:"))
}

func (c_ ControlTextEditingDelegateWrapper) ControlTextShouldBeginEditing(control IControl, fieldEditor IText) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("control:textShouldBeginEditing:"), objc.ExtractPtr(control), objc.ExtractPtr(fieldEditor))
	return rv
}

func (c_ ControlTextEditingDelegateWrapper) ImplementsControlTextShouldEndEditing() bool {
	return c_.RespondsToSelector(objc.GetSelector("control:textShouldEndEditing:"))
}

func (c_ ControlTextEditingDelegateWrapper) ControlTextShouldEndEditing(control IControl, fieldEditor IText) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("control:textShouldEndEditing:"), objc.ExtractPtr(control), objc.ExtractPtr(fieldEditor))
	return rv
}

func (c_ ControlTextEditingDelegateWrapper) ImplementsControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem() bool {
	return c_.RespondsToSelector(objc.GetSelector("control:textView:completions:forPartialWordRange:indexOfSelectedItem:"))
}

func (c_ ControlTextEditingDelegateWrapper) ControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem(control IControl, textView ITextView, words []string, charRange foundation.Range, index *int) []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("control:textView:completions:forPartialWordRange:indexOfSelectedItem:"), objc.ExtractPtr(control), objc.ExtractPtr(textView), words, charRange, index)
	return rv
}

func (c_ ControlTextEditingDelegateWrapper) ImplementsControlTextViewDoCommandBySelector() bool {
	return c_.RespondsToSelector(objc.GetSelector("control:textView:doCommandBySelector:"))
}

func (c_ ControlTextEditingDelegateWrapper) ControlTextViewDoCommandBySelector(control IControl, textView ITextView, commandSelector objc.Selector) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("control:textView:doCommandBySelector:"), objc.ExtractPtr(control), objc.ExtractPtr(textView), commandSelector)
	return rv
}

func (c_ ControlTextEditingDelegateWrapper) ImplementsControlTextDidBeginEditing() bool {
	return c_.RespondsToSelector(objc.GetSelector("controlTextDidBeginEditing:"))
}

func (c_ ControlTextEditingDelegateWrapper) ControlTextDidBeginEditing(obj foundation.INotification) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("controlTextDidBeginEditing:"), objc.ExtractPtr(obj))
}

func (c_ ControlTextEditingDelegateWrapper) ImplementsControlTextDidChange() bool {
	return c_.RespondsToSelector(objc.GetSelector("controlTextDidChange:"))
}

func (c_ ControlTextEditingDelegateWrapper) ControlTextDidChange(obj foundation.INotification) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("controlTextDidChange:"), objc.ExtractPtr(obj))
}

func (c_ ControlTextEditingDelegateWrapper) ImplementsControlTextDidEndEditing() bool {
	return c_.RespondsToSelector(objc.GetSelector("controlTextDidEndEditing:"))
}

func (c_ ControlTextEditingDelegateWrapper) ControlTextDidEndEditing(obj foundation.INotification) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("controlTextDidEndEditing:"), objc.ExtractPtr(obj))
}
