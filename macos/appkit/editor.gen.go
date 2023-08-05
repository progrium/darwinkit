// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type IEditor interface {
	// required
	CommitEditing() bool
	// required
	CommitEditingWithDelegateDidCommitSelectorContextInfo(delegate objc.Object, didCommitSelector objc.Selector, contextInfo unsafe.Pointer)
	// required
	CommitEditingAndReturnError(error *foundation.Error) bool
	// required
	DiscardEditing()
}

type EditorWrapper struct {
	objc.Object
}

func (e_ EditorWrapper) CommitEditing() bool {
	rv := objc.CallMethod[bool](e_, objc.GetSelector("commitEditing"))
	return rv
}

func (e_ EditorWrapper) CommitEditingWithDelegateDidCommitSelectorContextInfo(delegate objc.IObject, didCommitSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](e_, objc.GetSelector("commitEditingWithDelegate:didCommitSelector:contextInfo:"), objc.ExtractPtr(delegate), didCommitSelector, contextInfo)
}

func (e_ EditorWrapper) CommitEditingAndReturnError(error *foundation.Error) bool {
	rv := objc.CallMethod[bool](e_, objc.GetSelector("commitEditingAndReturnError:"), unsafe.Pointer(error))
	return rv
}

func (e_ EditorWrapper) DiscardEditing() {
	objc.CallMethod[objc.Void](e_, objc.GetSelector("discardEditing"))
}
