// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DraggingItem] class.
var DraggingItemClass = _DraggingItemClass{objc.GetClass("NSDraggingItem")}

type _DraggingItemClass struct {
	objc.Class
}

// An interface definition for the [DraggingItem] class.
type IDraggingItem interface {
	objc.IObject
	SetDraggingFrameContents(frame foundation.Rect, contents objc.IObject)
	Item() objc.Object
	DraggingFrame() foundation.Rect
	SetDraggingFrame(value foundation.Rect)
	ImageComponents() []DraggingImageComponent
	ImageComponentsProvider() func() []DraggingImageComponent
	SetImageComponentsProvider(value func() []DraggingImageComponent)
}

// A single dragged item within a dragging session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingitem?language=objc
type DraggingItem struct {
	objc.Object
}

func DraggingItemFrom(ptr unsafe.Pointer) DraggingItem {
	return DraggingItem{
		Object: objc.ObjectFrom(ptr),
	}
}

func (d_ DraggingItem) InitWithPasteboardWriter(pasteboardWriter PPasteboardWriting) DraggingItem {
	po0 := objc.WrapAsProtocol("NSPasteboardWriting", pasteboardWriter)
	rv := objc.Call[DraggingItem](d_, objc.Sel("initWithPasteboardWriter:"), po0)
	return rv
}

// Creates and returns a dragging item using the specified content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingitem/1535417-initwithpasteboardwriter?language=objc
func DraggingItem_InitWithPasteboardWriter(pasteboardWriter PPasteboardWriting) DraggingItem {
	return DraggingItemClass.Alloc().InitWithPasteboardWriter(pasteboardWriter)
}

func (dc _DraggingItemClass) Alloc() DraggingItem {
	rv := objc.Call[DraggingItem](dc, objc.Sel("alloc"))
	return rv
}

func DraggingItem_Alloc() DraggingItem {
	return DraggingItemClass.Alloc()
}

func (dc _DraggingItemClass) New() DraggingItem {
	rv := objc.Call[DraggingItem](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDraggingItem() DraggingItem {
	return DraggingItemClass.New()
}

func (d_ DraggingItem) Init() DraggingItem {
	rv := objc.Call[DraggingItem](d_, objc.Sel("init"))
	return rv
}

// Sets the item’s dragging frame and contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingitem/1528746-setdraggingframe?language=objc
func (d_ DraggingItem) SetDraggingFrameContents(frame foundation.Rect, contents objc.IObject) {
	objc.Call[objc.Void](d_, objc.Sel("setDraggingFrame:contents:"), frame, contents)
}

// The pasteboard reader or writer object dependent on the context where you use the dragging item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingitem/1533258-item?language=objc
func (d_ DraggingItem) Item() objc.Object {
	rv := objc.Call[objc.Object](d_, objc.Sel("item"))
	return rv
}

// The frame of the dragging item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingitem/1528559-draggingframe?language=objc
func (d_ DraggingItem) DraggingFrame() foundation.Rect {
	rv := objc.Call[foundation.Rect](d_, objc.Sel("draggingFrame"))
	return rv
}

// The frame of the dragging item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingitem/1528559-draggingframe?language=objc
func (d_ DraggingItem) SetDraggingFrame(value foundation.Rect) {
	objc.Call[objc.Void](d_, objc.Sel("setDraggingFrame:"), value)
}

// An array of dragging image components to use to create the drag image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingitem/1524302-imagecomponents?language=objc
func (d_ DraggingItem) ImageComponents() []DraggingImageComponent {
	rv := objc.Call[[]DraggingImageComponent](d_, objc.Sel("imageComponents"))
	return rv
}

// An array of blocks that provide the dragging image components. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingitem/1535607-imagecomponentsprovider?language=objc
func (d_ DraggingItem) ImageComponentsProvider() func() []DraggingImageComponent {
	rv := objc.Call[func() []DraggingImageComponent](d_, objc.Sel("imageComponentsProvider"))
	return rv
}

// An array of blocks that provide the dragging image components. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingitem/1535607-imagecomponentsprovider?language=objc
func (d_ DraggingItem) SetImageComponentsProvider(value func() []DraggingImageComponent) {
	objc.Call[objc.Void](d_, objc.Sel("setImageComponentsProvider:"), value)
}
