// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that a custom class can adopt to manage the automatic enablement of a UI control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsvalidateduserinterfaceitem?language=objc
type PValidatedUserInterfaceItem interface {
	// optional
	Action() objc.Selector
	HasAction() bool

	// optional
	Tag() int
	HasTag() bool
}

// A concrete type wrapper for the [PValidatedUserInterfaceItem] protocol.
type ValidatedUserInterfaceItemWrapper struct {
	objc.Object
}

func (v_ ValidatedUserInterfaceItemWrapper) HasAction() bool {
	return v_.RespondsToSelector(objc.Sel("action"))
}

// Returns the selector of the receiver’s action method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsvalidateduserinterfaceitem/1527339-action?language=objc
func (v_ ValidatedUserInterfaceItemWrapper) Action() objc.Selector {
	rv := objc.Call[objc.Selector](v_, objc.Sel("action"))
	return rv
}

func (v_ ValidatedUserInterfaceItemWrapper) HasTag() bool {
	return v_.RespondsToSelector(objc.Sel("tag"))
}

// Returns the receiver’s tag integer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsvalidateduserinterfaceitem/1531030-tag?language=objc
func (v_ ValidatedUserInterfaceItemWrapper) Tag() int {
	rv := objc.Call[int](v_, objc.Sel("tag"))
	return rv
}
