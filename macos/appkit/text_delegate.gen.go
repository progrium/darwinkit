// auto-generated code, do not modify
package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type ITextDelegate interface {
	ImplementsTextDidChange() bool
	// optional
	TextDidChange(notification foundation.Notification)
	ImplementsTextShouldBeginEditing() bool
	// optional
	TextShouldBeginEditing(textObject Text) bool
	ImplementsTextDidBeginEditing() bool
	// optional
	TextDidBeginEditing(notification foundation.Notification)
	ImplementsTextShouldEndEditing() bool
	// optional
	TextShouldEndEditing(textObject Text) bool
	ImplementsTextDidEndEditing() bool
	// optional
	TextDidEndEditing(notification foundation.Notification)
}

type TextDelegate struct {
	_TextDidChange          func(notification foundation.Notification)
	_TextShouldBeginEditing func(textObject Text) bool
	_TextDidBeginEditing    func(notification foundation.Notification)
	_TextShouldEndEditing   func(textObject Text) bool
	_TextDidEndEditing      func(notification foundation.Notification)
}

func (di *TextDelegate) ImplementsTextDidChange() bool {
	return di._TextDidChange != nil
}

func (di *TextDelegate) SetTextDidChange(f func(notification foundation.Notification)) {
	di._TextDidChange = f
}

func (di *TextDelegate) TextDidChange(notification foundation.Notification) {
	di._TextDidChange(notification)
}
func (di *TextDelegate) ImplementsTextShouldBeginEditing() bool {
	return di._TextShouldBeginEditing != nil
}

func (di *TextDelegate) SetTextShouldBeginEditing(f func(textObject Text) bool) {
	di._TextShouldBeginEditing = f
}

func (di *TextDelegate) TextShouldBeginEditing(textObject Text) bool {
	return di._TextShouldBeginEditing(textObject)
}
func (di *TextDelegate) ImplementsTextDidBeginEditing() bool {
	return di._TextDidBeginEditing != nil
}

func (di *TextDelegate) SetTextDidBeginEditing(f func(notification foundation.Notification)) {
	di._TextDidBeginEditing = f
}

func (di *TextDelegate) TextDidBeginEditing(notification foundation.Notification) {
	di._TextDidBeginEditing(notification)
}
func (di *TextDelegate) ImplementsTextShouldEndEditing() bool {
	return di._TextShouldEndEditing != nil
}

func (di *TextDelegate) SetTextShouldEndEditing(f func(textObject Text) bool) {
	di._TextShouldEndEditing = f
}

func (di *TextDelegate) TextShouldEndEditing(textObject Text) bool {
	return di._TextShouldEndEditing(textObject)
}
func (di *TextDelegate) ImplementsTextDidEndEditing() bool {
	return di._TextDidEndEditing != nil
}

func (di *TextDelegate) SetTextDidEndEditing(f func(notification foundation.Notification)) {
	di._TextDidEndEditing = f
}

func (di *TextDelegate) TextDidEndEditing(notification foundation.Notification) {
	di._TextDidEndEditing(notification)
}

type TextDelegateWrapper struct {
	objc.Object
}

func (t_ TextDelegateWrapper) ImplementsTextDidChange() bool {
	return t_.RespondsToSelector(objc.GetSelector("textDidChange:"))
}

func (t_ TextDelegateWrapper) TextDidChange(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("textDidChange:"), objc.ExtractPtr(notification))
}

func (t_ TextDelegateWrapper) ImplementsTextShouldBeginEditing() bool {
	return t_.RespondsToSelector(objc.GetSelector("textShouldBeginEditing:"))
}

func (t_ TextDelegateWrapper) TextShouldBeginEditing(textObject IText) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("textShouldBeginEditing:"), objc.ExtractPtr(textObject))
	return rv
}

func (t_ TextDelegateWrapper) ImplementsTextDidBeginEditing() bool {
	return t_.RespondsToSelector(objc.GetSelector("textDidBeginEditing:"))
}

func (t_ TextDelegateWrapper) TextDidBeginEditing(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("textDidBeginEditing:"), objc.ExtractPtr(notification))
}

func (t_ TextDelegateWrapper) ImplementsTextShouldEndEditing() bool {
	return t_.RespondsToSelector(objc.GetSelector("textShouldEndEditing:"))
}

func (t_ TextDelegateWrapper) TextShouldEndEditing(textObject IText) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("textShouldEndEditing:"), objc.ExtractPtr(textObject))
	return rv
}

func (t_ TextDelegateWrapper) ImplementsTextDidEndEditing() bool {
	return t_.RespondsToSelector(objc.GetSelector("textDidEndEditing:"))
}

func (t_ TextDelegateWrapper) TextDidEndEditing(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("textDidEndEditing:"), objc.ExtractPtr(notification))
}
