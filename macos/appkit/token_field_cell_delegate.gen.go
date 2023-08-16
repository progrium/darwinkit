// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A set of optional methods implemented by delegates of NSTokenFieldCell objects to work with tokenized strings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstokenfieldcelldelegate?language=objc
type PTokenFieldCellDelegate interface {
	// optional
	TokenFieldCellEditingStringForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) string
	HasTokenFieldCellEditingStringForRepresentedObject() bool
}

// A delegate implementation builder for the [PTokenFieldCellDelegate] protocol.
type TokenFieldCellDelegate struct {
	_TokenFieldCellEditingStringForRepresentedObject func(tokenFieldCell TokenFieldCell, representedObject objc.Object) string
}

func (di *TokenFieldCellDelegate) HasTokenFieldCellEditingStringForRepresentedObject() bool {
	return di._TokenFieldCellEditingStringForRepresentedObject != nil
}

// Allows the delegate to provide a string to be edited as a proxy for the represented object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstokenfieldcelldelegate/1523824-tokenfieldcell?language=objc
func (di *TokenFieldCellDelegate) SetTokenFieldCellEditingStringForRepresentedObject(f func(tokenFieldCell TokenFieldCell, representedObject objc.Object) string) {
	di._TokenFieldCellEditingStringForRepresentedObject = f
}

// Allows the delegate to provide a string to be edited as a proxy for the represented object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstokenfieldcelldelegate/1523824-tokenfieldcell?language=objc
func (di *TokenFieldCellDelegate) TokenFieldCellEditingStringForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) string {
	return di._TokenFieldCellEditingStringForRepresentedObject(tokenFieldCell, representedObject)
}

// A concrete type wrapper for the [PTokenFieldCellDelegate] protocol.
type TokenFieldCellDelegateWrapper struct {
	objc.Object
}

func (t_ TokenFieldCellDelegateWrapper) HasTokenFieldCellEditingStringForRepresentedObject() bool {
	return t_.RespondsToSelector(objc.Sel("tokenFieldCell:editingStringForRepresentedObject:"))
}

// Allows the delegate to provide a string to be edited as a proxy for the represented object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstokenfieldcelldelegate/1523824-tokenfieldcell?language=objc
func (t_ TokenFieldCellDelegateWrapper) TokenFieldCellEditingStringForRepresentedObject(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) string {
	rv := objc.Call[string](t_, objc.Sel("tokenFieldCell:editingStringForRepresentedObject:"), objc.Ptr(tokenFieldCell), representedObject)
	return rv
}
