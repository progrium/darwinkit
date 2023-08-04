// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var CollectionViewItemClass = _CollectionViewItemClass{objc.GetClass("NSCollectionViewItem")}

type _CollectionViewItemClass struct {
	objc.Class
}

type ICollectionViewItem interface {
	IViewController
	ImageView() ImageView
	SetImageView(value IImageView)
	TextField() TextField
	SetTextField(value ITextField)
	IsSelected() bool
	SetSelected(value bool)
	HighlightState() CollectionViewItemHighlightState
	SetHighlightState(value CollectionViewItemHighlightState)
	CollectionView() CollectionView
	DraggingImageComponents() []DraggingImageComponent
}

type CollectionViewItem struct {
	ViewController
}

func MakeCollectionViewItem(ptr unsafe.Pointer) CollectionViewItem {
	return CollectionViewItem{
		ViewController: MakeViewController(ptr),
	}
}

func (c_ CollectionViewItem) InitWithNibNameBundle(nibNameOrNil NibName, nibBundleOrNil foundation.IBundle) CollectionViewItem {
	rv := objc.CallMethod[CollectionViewItem](c_, objc.GetSelector("initWithNibName:bundle:"), nibNameOrNil, objc.ExtractPtr(nibBundleOrNil))
	return rv
}

func CollectionViewItem_InitWithNibNameBundle(nibNameOrNil NibName, nibBundleOrNil foundation.IBundle) CollectionViewItem {
	return CollectionViewItemClass.Alloc().InitWithNibNameBundle(nibNameOrNil, nibBundleOrNil)
}

func (c_ CollectionViewItem) Init() CollectionViewItem {
	rv := objc.CallMethod[CollectionViewItem](c_, objc.GetSelector("init"))
	return rv
}

func CollectionViewItem_Init() CollectionViewItem {
	return CollectionViewItemClass.Alloc().Init()
}

func (cc _CollectionViewItemClass) Alloc() CollectionViewItem {
	rv := objc.CallMethod[CollectionViewItem](cc, objc.GetSelector("alloc"))
	return rv
}

func CollectionViewItem_Alloc() CollectionViewItem {
	return CollectionViewItemClass.Alloc()
}

func (cc _CollectionViewItemClass) New() CollectionViewItem {
	rv := objc.CallMethod[CollectionViewItem](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionViewItem() CollectionViewItem {
	return CollectionViewItemClass.New()
}

func CollectionViewItem_New() CollectionViewItem {
	return CollectionViewItemClass.New()
}

func (c_ CollectionViewItem) ImageView() ImageView {
	rv := objc.CallMethod[ImageView](c_, objc.GetSelector("imageView"))
	return rv
}

func (c_ CollectionViewItem) SetImageView(value IImageView) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setImageView:"), objc.ExtractPtr(value))
}

func (c_ CollectionViewItem) TextField() TextField {
	rv := objc.CallMethod[TextField](c_, objc.GetSelector("textField"))
	return rv
}

func (c_ CollectionViewItem) SetTextField(value ITextField) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setTextField:"), objc.ExtractPtr(value))
}

func (c_ CollectionViewItem) IsSelected() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isSelected"))
	return rv
}

func (c_ CollectionViewItem) SetSelected(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setSelected:"), value)
}

func (c_ CollectionViewItem) HighlightState() CollectionViewItemHighlightState {
	rv := objc.CallMethod[CollectionViewItemHighlightState](c_, objc.GetSelector("highlightState"))
	return rv
}

func (c_ CollectionViewItem) SetHighlightState(value CollectionViewItemHighlightState) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setHighlightState:"), value)
}

func (c_ CollectionViewItem) CollectionView() CollectionView {
	rv := objc.CallMethod[CollectionView](c_, objc.GetSelector("collectionView"))
	return rv
}

func (c_ CollectionViewItem) DraggingImageComponents() []DraggingImageComponent {
	rv := objc.CallMethod[[]DraggingImageComponent](c_, objc.GetSelector("draggingImageComponents"))
	return rv
}
