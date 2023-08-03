// auto-generated code, do not modify
package appkit

import (
	"github.com/progrium/macdriver/objc"
)

type IUserInterfaceItemIdentification interface {
	ImplementsSetIdentifier() bool
	// optional
	SetIdentifier(value UserInterfaceItemIdentifier)
	ImplementsIdentifier() bool
	// optional
	Identifier() UserInterfaceItemIdentifier
}

type UserInterfaceItemIdentificationWrapper struct {
	objc.Object
}

func (u_ UserInterfaceItemIdentificationWrapper) ImplementsSetIdentifier() bool {
	return u_.RespondsToSelector(objc.GetSelector("setIdentifier:"))
}

func (u_ UserInterfaceItemIdentificationWrapper) SetIdentifier(value UserInterfaceItemIdentifier) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("setIdentifier:"), value)
}

func (u_ UserInterfaceItemIdentificationWrapper) ImplementsIdentifier() bool {
	return u_.RespondsToSelector(objc.GetSelector("identifier"))
}

func (u_ UserInterfaceItemIdentificationWrapper) Identifier() UserInterfaceItemIdentifier {
	rv := objc.CallMethod[UserInterfaceItemIdentifier](u_, objc.GetSelector("identifier"))
	return rv
}
