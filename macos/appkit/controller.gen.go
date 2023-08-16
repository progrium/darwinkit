// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Controller] class.
var ControllerClass = _ControllerClass{objc.GetClass("NSController")}

type _ControllerClass struct {
	objc.Class
}

// An interface definition for the [Controller] class.
type IController interface {
	objc.IObject
	CommitEditing() bool
	ObjectDidBeginEditing(editor PEditor)
	ObjectDidBeginEditingObject(editorObject objc.IObject)
	ObjectDidEndEditing(editor PEditor)
	ObjectDidEndEditingObject(editorObject objc.IObject)
	CommitEditingWithDelegateDidCommitSelectorContextInfo(delegate objc.IObject, didCommitSelector objc.Selector, contextInfo unsafe.Pointer)
	DiscardEditing()
	IsEditing() bool
}

// An abstract class that implements the [appkit/deprecated_symbols/nseditor] and [appkit/deprecated_symbols/nseditorregistration] informal protocols required for controller classes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroller?language=objc
type Controller struct {
	objc.Object
}

func ControllerFrom(ptr unsafe.Pointer) Controller {
	return Controller{
		Object: objc.ObjectFrom(ptr),
	}
}

func (c_ Controller) Init() Controller {
	rv := objc.Call[Controller](c_, objc.Sel("init"))
	return rv
}

func (cc _ControllerClass) Alloc() Controller {
	rv := objc.Call[Controller](cc, objc.Sel("alloc"))
	return rv
}

func Controller_Alloc() Controller {
	return ControllerClass.Alloc()
}

func (cc _ControllerClass) New() Controller {
	rv := objc.Call[Controller](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewController() Controller {
	return ControllerClass.New()
}

// Causes the receiver to attempt to commit any pending edits, returning YES if successful or no edits were pending. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroller/1531472-commitediting?language=objc
func (c_ Controller) CommitEditing() bool {
	rv := objc.Call[bool](c_, objc.Sel("commitEditing"))
	return rv
}

// Invoked to inform the receiver that editor has uncommitted changes that can affect the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroller/1532991-objectdidbeginediting?language=objc
func (c_ Controller) ObjectDidBeginEditing(editor PEditor) {
	po0 := objc.WrapAsProtocol("NSEditor", editor)
	objc.Call[objc.Void](c_, objc.Sel("objectDidBeginEditing:"), po0)
}

// Invoked to inform the receiver that editor has uncommitted changes that can affect the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroller/1532991-objectdidbeginediting?language=objc
func (c_ Controller) ObjectDidBeginEditingObject(editorObject objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("objectDidBeginEditing:"), objc.Ptr(editorObject))
}

// Invoked to inform the receiver that editor has committed or discarded its changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroller/1524644-objectdidendediting?language=objc
func (c_ Controller) ObjectDidEndEditing(editor PEditor) {
	po0 := objc.WrapAsProtocol("NSEditor", editor)
	objc.Call[objc.Void](c_, objc.Sel("objectDidEndEditing:"), po0)
}

// Invoked to inform the receiver that editor has committed or discarded its changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroller/1524644-objectdidendediting?language=objc
func (c_ Controller) ObjectDidEndEditingObject(editorObject objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("objectDidEndEditing:"), objc.Ptr(editorObject))
}

// Attempts to commit any pending changes in known editors of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroller/1527876-commiteditingwithdelegate?language=objc
func (c_ Controller) CommitEditingWithDelegateDidCommitSelectorContextInfo(delegate objc.IObject, didCommitSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.Call[objc.Void](c_, objc.Sel("commitEditingWithDelegate:didCommitSelector:contextInfo:"), delegate, didCommitSelector, contextInfo)
}

// Discards any pending changes by registered editors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroller/1528652-discardediting?language=objc
func (c_ Controller) DiscardEditing() {
	objc.Call[objc.Void](c_, objc.Sel("discardEditing"))
}

// A Boolean value indicating if any editors are registered with the controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroller/1527136-editing?language=objc
func (c_ Controller) IsEditing() bool {
	rv := objc.Call[bool](c_, objc.Sel("isEditing"))
	return rv
}
