// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nseditor?language=objc
type PEditor interface {
	// optional
	CommitEditing() bool
	HasCommitEditing() bool

	// optional
	CommitEditingWithDelegateDidCommitSelectorContextInfo(delegate objc.Object, didCommitSelector objc.Selector, contextInfo unsafe.Pointer)
	HasCommitEditingWithDelegateDidCommitSelectorContextInfo() bool

	// optional
	CommitEditingAndReturnError(error foundation.Error) bool
	HasCommitEditingAndReturnError() bool

	// optional
	DiscardEditing()
	HasDiscardEditing() bool
}

// A concrete type wrapper for the [PEditor] protocol.
type EditorWrapper struct {
	objc.Object
}

func (e_ EditorWrapper) HasCommitEditing() bool {
	return e_.RespondsToSelector(objc.Sel("commitEditing"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nseditor/3005183-commitediting?language=objc
func (e_ EditorWrapper) CommitEditing() bool {
	rv := objc.Call[bool](e_, objc.Sel("commitEditing"))
	return rv
}

func (e_ EditorWrapper) HasCommitEditingWithDelegateDidCommitSelectorContextInfo() bool {
	return e_.RespondsToSelector(objc.Sel("commitEditingWithDelegate:didCommitSelector:contextInfo:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nseditor/3005185-commiteditingwithdelegate?language=objc
func (e_ EditorWrapper) CommitEditingWithDelegateDidCommitSelectorContextInfo(delegate objc.IObject, didCommitSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.Call[objc.Void](e_, objc.Sel("commitEditingWithDelegate:didCommitSelector:contextInfo:"), delegate, didCommitSelector, contextInfo)
}

func (e_ EditorWrapper) HasCommitEditingAndReturnError() bool {
	return e_.RespondsToSelector(objc.Sel("commitEditingAndReturnError:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nseditor/3005184-commiteditingandreturnerror?language=objc
func (e_ EditorWrapper) CommitEditingAndReturnError(error foundation.IError) bool {
	rv := objc.Call[bool](e_, objc.Sel("commitEditingAndReturnError:"), objc.Ptr(error))
	return rv
}

func (e_ EditorWrapper) HasDiscardEditing() bool {
	return e_.RespondsToSelector(objc.Sel("discardEditing"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nseditor/3005186-discardediting?language=objc
func (e_ EditorWrapper) DiscardEditing() {
	objc.Call[objc.Void](e_, objc.Sel("discardEditing"))
}
