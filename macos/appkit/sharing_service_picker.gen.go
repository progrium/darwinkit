// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SharingServicePicker] class.
var SharingServicePickerClass = _SharingServicePickerClass{objc.GetClass("NSSharingServicePicker")}

type _SharingServicePickerClass struct {
	objc.Class
}

// An interface definition for the [SharingServicePicker] class.
type ISharingServicePicker interface {
	objc.IObject
	ShowRelativeToRectOfViewPreferredEdge(rect foundation.Rect, view IView, preferredEdge foundation.RectEdge)
	Delegate() SharingServicePickerDelegateWrapper
	SetDelegate(value PSharingServicePickerDelegate)
	SetDelegateObject(valueObject objc.IObject)
}

// A list of sharing services that the user can choose from. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepicker?language=objc
type SharingServicePicker struct {
	objc.Object
}

func SharingServicePickerFrom(ptr unsafe.Pointer) SharingServicePicker {
	return SharingServicePicker{
		Object: objc.ObjectFrom(ptr),
	}
}

func (s_ SharingServicePicker) InitWithItems(items []objc.IObject) SharingServicePicker {
	rv := objc.Call[SharingServicePicker](s_, objc.Sel("initWithItems:"), items)
	return rv
}

// Creates a new sharing service picker for the selected items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepicker/1402691-initwithitems?language=objc
func SharingServicePicker_InitWithItems(items []objc.IObject) SharingServicePicker {
	return SharingServicePickerClass.Alloc().InitWithItems(items)
}

func (sc _SharingServicePickerClass) Alloc() SharingServicePicker {
	rv := objc.Call[SharingServicePicker](sc, objc.Sel("alloc"))
	return rv
}

func SharingServicePicker_Alloc() SharingServicePicker {
	return SharingServicePickerClass.Alloc()
}

func (sc _SharingServicePickerClass) New() SharingServicePicker {
	rv := objc.Call[SharingServicePicker](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSharingServicePicker() SharingServicePicker {
	return SharingServicePickerClass.New()
}

func (s_ SharingServicePicker) Init() SharingServicePicker {
	rv := objc.Call[SharingServicePicker](s_, objc.Sel("init"))
	return rv
}

// Shows the picker interface and populates it with the relevant sharing services. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepicker/1402706-showrelativetorect?language=objc
func (s_ SharingServicePicker) ShowRelativeToRectOfViewPreferredEdge(rect foundation.Rect, view IView, preferredEdge foundation.RectEdge) {
	objc.Call[objc.Void](s_, objc.Sel("showRelativeToRect:ofView:preferredEdge:"), rect, objc.Ptr(view), preferredEdge)
}

// The object for managing the sharing service picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepicker/1402687-delegate?language=objc
func (s_ SharingServicePicker) Delegate() SharingServicePickerDelegateWrapper {
	rv := objc.Call[SharingServicePickerDelegateWrapper](s_, objc.Sel("delegate"))
	return rv
}

// The object for managing the sharing service picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepicker/1402687-delegate?language=objc
func (s_ SharingServicePicker) SetDelegate(value PSharingServicePickerDelegate) {
	po0 := objc.WrapAsProtocol("NSSharingServicePickerDelegate", value)
	objc.SetAssociatedObject(s_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](s_, objc.Sel("setDelegate:"), po0)
}

// The object for managing the sharing service picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepicker/1402687-delegate?language=objc
func (s_ SharingServicePicker) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}
