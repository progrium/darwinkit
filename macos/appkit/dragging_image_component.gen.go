// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var DraggingImageComponentClass = _DraggingImageComponentClass{objc.GetClass("NSDraggingImageComponent")}

type _DraggingImageComponentClass struct {
	objc.Class
}

type IDraggingImageComponent interface {
	objc.IObject
	Key() DraggingImageComponentKey
	SetKey(value DraggingImageComponentKey)
	Contents() objc.Object
	SetContents(value objc.IObject)
	Frame() foundation.Rect
	SetFrame(value foundation.Rect)
}

type DraggingImageComponent struct {
	objc.Object
}

func MakeDraggingImageComponent(ptr unsafe.Pointer) DraggingImageComponent {
	return DraggingImageComponent{
		Object: objc.MakeObject(ptr),
	}
}

func (d_ DraggingImageComponent) InitWithKey(key DraggingImageComponentKey) DraggingImageComponent {
	rv := objc.CallMethod[DraggingImageComponent](d_, objc.GetSelector("initWithKey:"), key)
	return rv
}

func DraggingImageComponent_InitWithKey(key DraggingImageComponentKey) DraggingImageComponent {
	return DraggingImageComponentClass.Alloc().InitWithKey(key)
}

func (dc _DraggingImageComponentClass) Alloc() DraggingImageComponent {
	rv := objc.CallMethod[DraggingImageComponent](dc, objc.GetSelector("alloc"))
	return rv
}

func DraggingImageComponent_Alloc() DraggingImageComponent {
	return DraggingImageComponentClass.Alloc()
}

func (dc _DraggingImageComponentClass) New() DraggingImageComponent {
	rv := objc.CallMethod[DraggingImageComponent](dc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewDraggingImageComponent() DraggingImageComponent {
	return DraggingImageComponentClass.New()
}

func DraggingImageComponent_New() DraggingImageComponent {
	return DraggingImageComponentClass.New()
}

func (d_ DraggingImageComponent) Init() DraggingImageComponent {
	rv := objc.CallMethod[DraggingImageComponent](d_, objc.GetSelector("init"))
	return rv
}

func DraggingImageComponent_Init() DraggingImageComponent {
	return DraggingImageComponentClass.Alloc().Init()
}

func (dc _DraggingImageComponentClass) DraggingImageComponentWithKey(key DraggingImageComponentKey) DraggingImageComponent {
	rv := objc.CallMethod[DraggingImageComponent](dc, objc.GetSelector("draggingImageComponentWithKey:"), key)
	return rv
}

func DraggingImageComponent_DraggingImageComponentWithKey(key DraggingImageComponentKey) DraggingImageComponent {
	return DraggingImageComponentClass.DraggingImageComponentWithKey(key)
}

func (d_ DraggingImageComponent) Key() DraggingImageComponentKey {
	rv := objc.CallMethod[DraggingImageComponentKey](d_, objc.GetSelector("key"))
	return rv
}

func (d_ DraggingImageComponent) SetKey(value DraggingImageComponentKey) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setKey:"), value)
}

func (d_ DraggingImageComponent) Contents() objc.Object {
	rv := objc.CallMethod[objc.Object](d_, objc.GetSelector("contents"))
	return rv
}

func (d_ DraggingImageComponent) SetContents(value objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setContents:"), objc.ExtractPtr(value))
}

func (d_ DraggingImageComponent) Frame() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](d_, objc.GetSelector("frame"))
	return rv
}

func (d_ DraggingImageComponent) SetFrame(value foundation.Rect) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setFrame:"), value)
}
