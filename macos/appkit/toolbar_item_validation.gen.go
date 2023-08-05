// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"github.com/progrium/macdriver/objc"
)

type IToolbarItemValidation interface {
	// required
	ValidateToolbarItem(item ToolbarItem) bool
}

type ToolbarItemValidationWrapper struct {
	objc.Object
}

func (t_ ToolbarItemValidationWrapper) ValidateToolbarItem(item IToolbarItem) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("validateToolbarItem:"), objc.ExtractPtr(item))
	return rv
}
