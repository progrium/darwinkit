// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DraggingImageComponent] class.
var DraggingImageComponentClass = _DraggingImageComponentClass{objc.GetClass("NSDraggingImageComponent")}

type _DraggingImageComponentClass struct {
	objc.Class
}

// An interface definition for the [DraggingImageComponent] class.
type IDraggingImageComponent interface {
	objc.IObject
	Key() DraggingImageComponentKey
	SetKey(value DraggingImageComponentKey)
	Contents() objc.Object
	SetContents(value objc.IObject)
	Frame() foundation.Rect
	SetFrame(value foundation.Rect)
}

// A single object in a dragging item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingimagecomponent?language=objc
type DraggingImageComponent struct {
	objc.Object
}

func DraggingImageComponentFrom(ptr unsafe.Pointer) DraggingImageComponent {
	return DraggingImageComponent{
		Object: objc.ObjectFrom(ptr),
	}
}

func (d_ DraggingImageComponent) InitWithKey(key DraggingImageComponentKey) DraggingImageComponent {
	rv := objc.Call[DraggingImageComponent](d_, objc.Sel("initWithKey:"), key)
	return rv
}

// Initializes and returns a dragging image component with the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingimagecomponent/1534187-initwithkey?language=objc
func DraggingImageComponent_InitWithKey(key DraggingImageComponentKey) DraggingImageComponent {
	return DraggingImageComponentClass.Alloc().InitWithKey(key)
}

func (dc _DraggingImageComponentClass) Alloc() DraggingImageComponent {
	rv := objc.Call[DraggingImageComponent](dc, objc.Sel("alloc"))
	return rv
}

func DraggingImageComponent_Alloc() DraggingImageComponent {
	return DraggingImageComponentClass.Alloc()
}

func (dc _DraggingImageComponentClass) New() DraggingImageComponent {
	rv := objc.Call[DraggingImageComponent](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDraggingImageComponent() DraggingImageComponent {
	return DraggingImageComponentClass.New()
}

func (d_ DraggingImageComponent) Init() DraggingImageComponent {
	rv := objc.Call[DraggingImageComponent](d_, objc.Sel("init"))
	return rv
}

// Creates and returns a dragging image component with the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingimagecomponent/1557913-draggingimagecomponentwithkey?language=objc
func (dc _DraggingImageComponentClass) DraggingImageComponentWithKey(key DraggingImageComponentKey) DraggingImageComponent {
	rv := objc.Call[DraggingImageComponent](dc, objc.Sel("draggingImageComponentWithKey:"), key)
	return rv
}

// Creates and returns a dragging image component with the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingimagecomponent/1557913-draggingimagecomponentwithkey?language=objc
func DraggingImageComponent_DraggingImageComponentWithKey(key DraggingImageComponentKey) DraggingImageComponent {
	return DraggingImageComponentClass.DraggingImageComponentWithKey(key)
}

// The unique name of this image component instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingimagecomponent/1535631-key?language=objc
func (d_ DraggingImageComponent) Key() DraggingImageComponentKey {
	rv := objc.Call[DraggingImageComponentKey](d_, objc.Sel("key"))
	return rv
}

// The unique name of this image component instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingimagecomponent/1535631-key?language=objc
func (d_ DraggingImageComponent) SetKey(value DraggingImageComponentKey) {
	objc.Call[objc.Void](d_, objc.Sel("setKey:"), value)
}

// An object providing the image contents of the component. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingimagecomponent/1529426-contents?language=objc
func (d_ DraggingImageComponent) Contents() objc.Object {
	rv := objc.Call[objc.Object](d_, objc.Sel("contents"))
	return rv
}

// An object providing the image contents of the component. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingimagecomponent/1529426-contents?language=objc
func (d_ DraggingImageComponent) SetContents(value objc.IObject) {
	objc.Call[objc.Void](d_, objc.Sel("setContents:"), value)
}

// The coordinate space is the bounds of the parent dragging item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingimagecomponent/1535507-frame?language=objc
func (d_ DraggingImageComponent) Frame() foundation.Rect {
	rv := objc.Call[foundation.Rect](d_, objc.Sel("frame"))
	return rv
}

// The coordinate space is the bounds of the parent dragging item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingimagecomponent/1535507-frame?language=objc
func (d_ DraggingImageComponent) SetFrame(value foundation.Rect) {
	objc.Call[objc.Void](d_, objc.Sel("setFrame:"), value)
}
