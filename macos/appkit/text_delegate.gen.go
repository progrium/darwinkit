// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of optional methods implemented by the delegate of an NSText object to edit text and change text formats. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextdelegate?language=objc
type PTextDelegate interface {
	// optional
	TextShouldBeginEditing(textObject Text) bool
	HasTextShouldBeginEditing() bool

	// optional
	TextDidChange(notification foundation.Notification)
	HasTextDidChange() bool

	// optional
	TextDidBeginEditing(notification foundation.Notification)
	HasTextDidBeginEditing() bool

	// optional
	TextShouldEndEditing(textObject Text) bool
	HasTextShouldEndEditing() bool

	// optional
	TextDidEndEditing(notification foundation.Notification)
	HasTextDidEndEditing() bool
}

// A delegate implementation builder for the [PTextDelegate] protocol.
type TextDelegate struct {
	_TextShouldBeginEditing func(textObject Text) bool
	_TextDidChange          func(notification foundation.Notification)
	_TextDidBeginEditing    func(notification foundation.Notification)
	_TextShouldEndEditing   func(textObject Text) bool
	_TextDidEndEditing      func(notification foundation.Notification)
}

func (di *TextDelegate) HasTextShouldBeginEditing() bool {
	return di._TextShouldBeginEditing != nil
}

// Invoked when a text object begins to change its text, this method requests permission for aTextObject to begin editing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextdelegate/1533298-textshouldbeginediting?language=objc
func (di *TextDelegate) SetTextShouldBeginEditing(f func(textObject Text) bool) {
	di._TextShouldBeginEditing = f
}

// Invoked when a text object begins to change its text, this method requests permission for aTextObject to begin editing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextdelegate/1533298-textshouldbeginediting?language=objc
func (di *TextDelegate) TextShouldBeginEditing(textObject Text) bool {
	return di._TextShouldBeginEditing(textObject)
}
func (di *TextDelegate) HasTextDidChange() bool {
	return di._TextDidChange != nil
}

// Informs the delegate that the text object has changed its characters or formatting attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextdelegate/1526982-textdidchange?language=objc
func (di *TextDelegate) SetTextDidChange(f func(notification foundation.Notification)) {
	di._TextDidChange = f
}

// Informs the delegate that the text object has changed its characters or formatting attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextdelegate/1526982-textdidchange?language=objc
func (di *TextDelegate) TextDidChange(notification foundation.Notification) {
	di._TextDidChange(notification)
}
func (di *TextDelegate) HasTextDidBeginEditing() bool {
	return di._TextDidBeginEditing != nil
}

// Informs the delegate that the text object has begun editing (that the user has begun changing it). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextdelegate/1535575-textdidbeginediting?language=objc
func (di *TextDelegate) SetTextDidBeginEditing(f func(notification foundation.Notification)) {
	di._TextDidBeginEditing = f
}

// Informs the delegate that the text object has begun editing (that the user has begun changing it). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextdelegate/1535575-textdidbeginediting?language=objc
func (di *TextDelegate) TextDidBeginEditing(notification foundation.Notification) {
	di._TextDidBeginEditing(notification)
}
func (di *TextDelegate) HasTextShouldEndEditing() bool {
	return di._TextShouldEndEditing != nil
}

// Invoked from a text object’s implementation of resignFirstResponder, this method requests permission for aTextObject to end editing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextdelegate/1525992-textshouldendediting?language=objc
func (di *TextDelegate) SetTextShouldEndEditing(f func(textObject Text) bool) {
	di._TextShouldEndEditing = f
}

// Invoked from a text object’s implementation of resignFirstResponder, this method requests permission for aTextObject to end editing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextdelegate/1525992-textshouldendediting?language=objc
func (di *TextDelegate) TextShouldEndEditing(textObject Text) bool {
	return di._TextShouldEndEditing(textObject)
}
func (di *TextDelegate) HasTextDidEndEditing() bool {
	return di._TextDidEndEditing != nil
}

// Informs the delegate that the text object has finished editing (that it has resigned first responder status). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextdelegate/1529016-textdidendediting?language=objc
func (di *TextDelegate) SetTextDidEndEditing(f func(notification foundation.Notification)) {
	di._TextDidEndEditing = f
}

// Informs the delegate that the text object has finished editing (that it has resigned first responder status). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextdelegate/1529016-textdidendediting?language=objc
func (di *TextDelegate) TextDidEndEditing(notification foundation.Notification) {
	di._TextDidEndEditing(notification)
}

// A concrete type wrapper for the [PTextDelegate] protocol.
type TextDelegateWrapper struct {
	objc.Object
}

func (t_ TextDelegateWrapper) HasTextShouldBeginEditing() bool {
	return t_.RespondsToSelector(objc.Sel("textShouldBeginEditing:"))
}

// Invoked when a text object begins to change its text, this method requests permission for aTextObject to begin editing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextdelegate/1533298-textshouldbeginediting?language=objc
func (t_ TextDelegateWrapper) TextShouldBeginEditing(textObject IText) bool {
	rv := objc.Call[bool](t_, objc.Sel("textShouldBeginEditing:"), objc.Ptr(textObject))
	return rv
}

func (t_ TextDelegateWrapper) HasTextDidChange() bool {
	return t_.RespondsToSelector(objc.Sel("textDidChange:"))
}

// Informs the delegate that the text object has changed its characters or formatting attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextdelegate/1526982-textdidchange?language=objc
func (t_ TextDelegateWrapper) TextDidChange(notification foundation.INotification) {
	objc.Call[objc.Void](t_, objc.Sel("textDidChange:"), objc.Ptr(notification))
}

func (t_ TextDelegateWrapper) HasTextDidBeginEditing() bool {
	return t_.RespondsToSelector(objc.Sel("textDidBeginEditing:"))
}

// Informs the delegate that the text object has begun editing (that the user has begun changing it). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextdelegate/1535575-textdidbeginediting?language=objc
func (t_ TextDelegateWrapper) TextDidBeginEditing(notification foundation.INotification) {
	objc.Call[objc.Void](t_, objc.Sel("textDidBeginEditing:"), objc.Ptr(notification))
}

func (t_ TextDelegateWrapper) HasTextShouldEndEditing() bool {
	return t_.RespondsToSelector(objc.Sel("textShouldEndEditing:"))
}

// Invoked from a text object’s implementation of resignFirstResponder, this method requests permission for aTextObject to end editing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextdelegate/1525992-textshouldendediting?language=objc
func (t_ TextDelegateWrapper) TextShouldEndEditing(textObject IText) bool {
	rv := objc.Call[bool](t_, objc.Sel("textShouldEndEditing:"), objc.Ptr(textObject))
	return rv
}

func (t_ TextDelegateWrapper) HasTextDidEndEditing() bool {
	return t_.RespondsToSelector(objc.Sel("textDidEndEditing:"))
}

// Informs the delegate that the text object has finished editing (that it has resigned first responder status). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextdelegate/1529016-textdidendediting?language=objc
func (t_ TextDelegateWrapper) TextDidEndEditing(notification foundation.INotification) {
	objc.Call[objc.Void](t_, objc.Sel("textDidEndEditing:"), objc.Ptr(notification))
}
