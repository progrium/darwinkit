// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nseditorregistration?language=objc
type PEditorRegistration interface {
	// optional
	ObjectDidBeginEditing(editor EditorWrapper)
	HasObjectDidBeginEditing() bool

	// optional
	ObjectDidEndEditing(editor EditorWrapper)
	HasObjectDidEndEditing() bool
}

// A concrete type wrapper for the [PEditorRegistration] protocol.
type EditorRegistrationWrapper struct {
	objc.Object
}

func (e_ EditorRegistrationWrapper) HasObjectDidBeginEditing() bool {
	return e_.RespondsToSelector(objc.Sel("objectDidBeginEditing:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nseditorregistration/3005188-objectdidbeginediting?language=objc
func (e_ EditorRegistrationWrapper) ObjectDidBeginEditing(editor PEditor) {
	po0 := objc.WrapAsProtocol("NSEditor", editor)
	objc.Call[objc.Void](e_, objc.Sel("objectDidBeginEditing:"), po0)
}

func (e_ EditorRegistrationWrapper) HasObjectDidEndEditing() bool {
	return e_.RespondsToSelector(objc.Sel("objectDidEndEditing:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nseditorregistration/3005189-objectdidendediting?language=objc
func (e_ EditorRegistrationWrapper) ObjectDidEndEditing(editor PEditor) {
	po0 := objc.WrapAsProtocol("NSEditor", editor)
	objc.Call[objc.Void](e_, objc.Sel("objectDidEndEditing:"), po0)
}
