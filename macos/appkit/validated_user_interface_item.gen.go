// auto-generated code, do not modify
package appkit

import (
	"github.com/progrium/macdriver/objc"
)

type IValidatedUserInterfaceItem interface {
	ImplementsAction() bool
	// optional
	Action() objc.Selector
	ImplementsTag() bool
	// optional
	Tag() int
}

type ValidatedUserInterfaceItemWrapper struct {
	objc.Object
}

func (v_ ValidatedUserInterfaceItemWrapper) ImplementsAction() bool {
	return v_.RespondsToSelector(objc.GetSelector("action"))
}

func (v_ ValidatedUserInterfaceItemWrapper) Action() objc.Selector {
	rv := objc.CallMethod[objc.Selector](v_, objc.GetSelector("action"))
	return rv
}

func (v_ ValidatedUserInterfaceItemWrapper) ImplementsTag() bool {
	return v_.RespondsToSelector(objc.GetSelector("tag"))
}

func (v_ ValidatedUserInterfaceItemWrapper) Tag() int {
	rv := objc.CallMethod[int](v_, objc.GetSelector("tag"))
	return rv
}
