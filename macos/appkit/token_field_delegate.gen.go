// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A set of optional methods implemented by delegates of NSTokenField objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstokenfielddelegate?language=objc
type PTokenFieldDelegate interface {
	PTextFieldDelegate
	// optional
	TokenFieldEditingStringForRepresentedObject(tokenField TokenField, representedObject objc.Object) string
	HasTokenFieldEditingStringForRepresentedObject() bool
}

// A delegate implementation builder for the [PTokenFieldDelegate] protocol.
type TokenFieldDelegate struct {
	TextFieldDelegate
	_TokenFieldEditingStringForRepresentedObject func(tokenField TokenField, representedObject objc.Object) string
}

func (di *TokenFieldDelegate) HasTokenFieldEditingStringForRepresentedObject() bool {
	return di._TokenFieldEditingStringForRepresentedObject != nil
}

// Allows the delegate to provide a string to be edited as a proxy for a represented object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstokenfielddelegate/1524432-tokenfield?language=objc
func (di *TokenFieldDelegate) SetTokenFieldEditingStringForRepresentedObject(f func(tokenField TokenField, representedObject objc.Object) string) {
	di._TokenFieldEditingStringForRepresentedObject = f
}

// Allows the delegate to provide a string to be edited as a proxy for a represented object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstokenfielddelegate/1524432-tokenfield?language=objc
func (di *TokenFieldDelegate) TokenFieldEditingStringForRepresentedObject(tokenField TokenField, representedObject objc.Object) string {
	return di._TokenFieldEditingStringForRepresentedObject(tokenField, representedObject)
}

// A concrete type wrapper for the [PTokenFieldDelegate] protocol.
type TokenFieldDelegateWrapper struct {
	TextFieldDelegateWrapper
}

func (t_ TokenFieldDelegateWrapper) HasTokenFieldEditingStringForRepresentedObject() bool {
	return t_.RespondsToSelector(objc.Sel("tokenField:editingStringForRepresentedObject:"))
}

// Allows the delegate to provide a string to be edited as a proxy for a represented object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstokenfielddelegate/1524432-tokenfield?language=objc
func (t_ TokenFieldDelegateWrapper) TokenFieldEditingStringForRepresentedObject(tokenField ITokenField, representedObject objc.IObject) string {
	rv := objc.Call[string](t_, objc.Sel("tokenField:editingStringForRepresentedObject:"), objc.Ptr(tokenField), representedObject)
	return rv
}
