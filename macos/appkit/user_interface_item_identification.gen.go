// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A set of methods used to associate a unique identifier with objects in your user interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfaceitemidentification?language=objc
type PUserInterfaceItemIdentification interface {
	// optional
	SetIdentifier(value UserInterfaceItemIdentifier)
	HasSetIdentifier() bool

	// optional
	Identifier() UserInterfaceItemIdentifier
	HasIdentifier() bool
}

// A concrete type wrapper for the [PUserInterfaceItemIdentification] protocol.
type UserInterfaceItemIdentificationWrapper struct {
	objc.Object
}

func (u_ UserInterfaceItemIdentificationWrapper) HasSetIdentifier() bool {
	return u_.RespondsToSelector(objc.Sel("setIdentifier:"))
}

// A string that identifies the user interface item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfaceitemidentification/1396829-identifier?language=objc
func (u_ UserInterfaceItemIdentificationWrapper) SetIdentifier(value UserInterfaceItemIdentifier) {
	objc.Call[objc.Void](u_, objc.Sel("setIdentifier:"), value)
}

func (u_ UserInterfaceItemIdentificationWrapper) HasIdentifier() bool {
	return u_.RespondsToSelector(objc.Sel("identifier"))
}

// A string that identifies the user interface item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfaceitemidentification/1396829-identifier?language=objc
func (u_ UserInterfaceItemIdentificationWrapper) Identifier() UserInterfaceItemIdentifier {
	rv := objc.Call[UserInterfaceItemIdentifier](u_, objc.Sel("identifier"))
	return rv
}
