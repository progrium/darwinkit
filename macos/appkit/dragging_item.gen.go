// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var DraggingItemClass = _DraggingItemClass{objc.GetClass("NSDraggingItem")}

type _DraggingItemClass struct {
	objc.Class
}

type IDraggingItem interface {
	objc.IObject
	SetDraggingFrameContents(frame foundation.Rect, contents objc.IObject)
	DraggingFrame() foundation.Rect
	SetDraggingFrame(value foundation.Rect)
	ImageComponents() []DraggingImageComponent
	ImageComponentsProvider() func() []DraggingImageComponent
	SetImageComponentsProvider(value func() []DraggingImageComponent)
	Item() objc.Object
}

type DraggingItem struct {
	objc.Object
}

func MakeDraggingItem(ptr unsafe.Pointer) DraggingItem {
	return DraggingItem{
		Object: objc.MakeObject(ptr),
	}
}

func (d_ DraggingItem) InitWithPasteboardWriter(pasteboardWriter IPasteboardWriting) DraggingItem {
	po := objc.WrapAsProtocol("NSPasteboardWriting", pasteboardWriter)
	rv := objc.CallMethod[DraggingItem](d_, objc.GetSelector("initWithPasteboardWriter:"), po)
	return rv
}

func DraggingItem_InitWithPasteboardWriter(pasteboardWriter IPasteboardWriting) DraggingItem {
	return DraggingItemClass.Alloc().InitWithPasteboardWriter(pasteboardWriter)
}

func (dc _DraggingItemClass) Alloc() DraggingItem {
	rv := objc.CallMethod[DraggingItem](dc, objc.GetSelector("alloc"))
	return rv
}

func DraggingItem_Alloc() DraggingItem {
	return DraggingItemClass.Alloc()
}

func (dc _DraggingItemClass) New() DraggingItem {
	rv := objc.CallMethod[DraggingItem](dc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewDraggingItem() DraggingItem {
	return DraggingItemClass.New()
}

func DraggingItem_New() DraggingItem {
	return DraggingItemClass.New()
}

func (d_ DraggingItem) Init() DraggingItem {
	rv := objc.CallMethod[DraggingItem](d_, objc.GetSelector("init"))
	return rv
}

func DraggingItem_Init() DraggingItem {
	return DraggingItemClass.Alloc().Init()
}

func (d_ DraggingItem) SetDraggingFrameContents(frame foundation.Rect, contents objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setDraggingFrame:contents:"), frame, objc.ExtractPtr(contents))
}

func (d_ DraggingItem) DraggingFrame() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](d_, objc.GetSelector("draggingFrame"))
	return rv
}

func (d_ DraggingItem) SetDraggingFrame(value foundation.Rect) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setDraggingFrame:"), value)
}

func (d_ DraggingItem) ImageComponents() []DraggingImageComponent {
	rv := objc.CallMethod[[]DraggingImageComponent](d_, objc.GetSelector("imageComponents"))
	return rv
}

func (d_ DraggingItem) ImageComponentsProvider() func() []DraggingImageComponent {
	rv := objc.CallMethod[func() []DraggingImageComponent](d_, objc.GetSelector("imageComponentsProvider"))
	return rv
}

func (d_ DraggingItem) SetImageComponentsProvider(value func() []DraggingImageComponent) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setImageComponentsProvider:"), value)
}

func (d_ DraggingItem) Item() objc.Object {
	rv := objc.CallMethod[objc.Object](d_, objc.GetSelector("item"))
	return rv
}
