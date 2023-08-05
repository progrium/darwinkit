// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var TableCellViewClass = _TableCellViewClass{objc.GetClass("NSTableCellView")}

type _TableCellViewClass struct {
	objc.Class
}

type ITableCellView interface {
	IView
	ObjectValue() objc.Object
	SetObjectValue(value objc.IObject)
	ImageView() ImageView
	SetImageView(value IImageView)
	TextField() TextField
	SetTextField(value ITextField)
	BackgroundStyle() BackgroundStyle
	SetBackgroundStyle(value BackgroundStyle)
	RowSizeStyle() TableViewRowSizeStyle
	SetRowSizeStyle(value TableViewRowSizeStyle)
	DraggingImageComponents() []DraggingImageComponent
}

type TableCellView struct {
	View
}

func MakeTableCellView(ptr unsafe.Pointer) TableCellView {
	return TableCellView{
		View: MakeView(ptr),
	}
}

func (t_ TableCellView) InitWithFrame(frameRect foundation.Rect) TableCellView {
	rv := objc.CallMethod[TableCellView](t_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func TableCellView_InitWithFrame(frameRect foundation.Rect) TableCellView {
	return TableCellViewClass.Alloc().InitWithFrame(frameRect)
}

func (t_ TableCellView) Init() TableCellView {
	rv := objc.CallMethod[TableCellView](t_, objc.GetSelector("init"))
	return rv
}

func TableCellView_Init() TableCellView {
	return TableCellViewClass.Alloc().Init()
}

func (tc _TableCellViewClass) Alloc() TableCellView {
	rv := objc.CallMethod[TableCellView](tc, objc.GetSelector("alloc"))
	return rv
}

func TableCellView_Alloc() TableCellView {
	return TableCellViewClass.Alloc()
}

func (tc _TableCellViewClass) New() TableCellView {
	rv := objc.CallMethod[TableCellView](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTableCellView() TableCellView {
	return TableCellViewClass.New()
}

func TableCellView_New() TableCellView {
	return TableCellViewClass.New()
}

func (t_ TableCellView) ObjectValue() objc.Object {
	rv := objc.CallMethod[objc.Object](t_, objc.GetSelector("objectValue"))
	return rv
}

func (t_ TableCellView) SetObjectValue(value objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setObjectValue:"), objc.ExtractPtr(value))
}

func (t_ TableCellView) ImageView() ImageView {
	rv := objc.CallMethod[ImageView](t_, objc.GetSelector("imageView"))
	return rv
}

func (t_ TableCellView) SetImageView(value IImageView) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setImageView:"), objc.ExtractPtr(value))
}

func (t_ TableCellView) TextField() TextField {
	rv := objc.CallMethod[TextField](t_, objc.GetSelector("textField"))
	return rv
}

func (t_ TableCellView) SetTextField(value ITextField) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTextField:"), objc.ExtractPtr(value))
}

func (t_ TableCellView) BackgroundStyle() BackgroundStyle {
	rv := objc.CallMethod[BackgroundStyle](t_, objc.GetSelector("backgroundStyle"))
	return rv
}

func (t_ TableCellView) SetBackgroundStyle(value BackgroundStyle) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setBackgroundStyle:"), value)
}

func (t_ TableCellView) RowSizeStyle() TableViewRowSizeStyle {
	rv := objc.CallMethod[TableViewRowSizeStyle](t_, objc.GetSelector("rowSizeStyle"))
	return rv
}

func (t_ TableCellView) SetRowSizeStyle(value TableViewRowSizeStyle) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setRowSizeStyle:"), value)
}

func (t_ TableCellView) DraggingImageComponents() []DraggingImageComponent {
	rv := objc.CallMethod[[]DraggingImageComponent](t_, objc.GetSelector("draggingImageComponents"))
	return rv
}
