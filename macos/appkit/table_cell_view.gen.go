// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TableCellView] class.
var TableCellViewClass = _TableCellViewClass{objc.GetClass("NSTableCellView")}

type _TableCellViewClass struct {
	objc.Class
}

// An interface definition for the [TableCellView] class.
type ITableCellView interface {
	IView
	BackgroundStyle() BackgroundStyle
	SetBackgroundStyle(value BackgroundStyle)
	RowSizeStyle() TableViewRowSizeStyle
	SetRowSizeStyle(value TableViewRowSizeStyle)
	TextField() TextField
	SetTextField(value ITextField)
	ObjectValue() objc.Object
	SetObjectValue(value objc.IObject)
	DraggingImageComponents() []DraggingImageComponent
	ImageView() ImageView
	SetImageView(value IImageView)
}

// A reusable container view shown for a particular cell in a table view that uses rows for content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecellview?language=objc
type TableCellView struct {
	View
}

func TableCellViewFrom(ptr unsafe.Pointer) TableCellView {
	return TableCellView{
		View: ViewFrom(ptr),
	}
}

func (tc _TableCellViewClass) Alloc() TableCellView {
	rv := objc.Call[TableCellView](tc, objc.Sel("alloc"))
	return rv
}

func TableCellView_Alloc() TableCellView {
	return TableCellViewClass.Alloc()
}

func (tc _TableCellViewClass) New() TableCellView {
	rv := objc.Call[TableCellView](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTableCellView() TableCellView {
	return TableCellViewClass.New()
}

func (t_ TableCellView) Init() TableCellView {
	rv := objc.Call[TableCellView](t_, objc.Sel("init"))
	return rv
}

func (t_ TableCellView) InitWithFrame(frameRect foundation.Rect) TableCellView {
	rv := objc.Call[TableCellView](t_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func NewTableCellViewWithFrame(frameRect foundation.Rect) TableCellView {
	instance := TableCellViewClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

// This property is automatically set by the enclosing row view to let this view know what its background looks like. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecellview/1483206-backgroundstyle?language=objc
func (t_ TableCellView) BackgroundStyle() BackgroundStyle {
	rv := objc.Call[BackgroundStyle](t_, objc.Sel("backgroundStyle"))
	return rv
}

// This property is automatically set by the enclosing row view to let this view know what its background looks like. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecellview/1483206-backgroundstyle?language=objc
func (t_ TableCellView) SetBackgroundStyle(value BackgroundStyle) {
	objc.Call[objc.Void](t_, objc.Sel("setBackgroundStyle:"), value)
}

// Returns the row size style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecellview/1483211-rowsizestyle?language=objc
func (t_ TableCellView) RowSizeStyle() TableViewRowSizeStyle {
	rv := objc.Call[TableViewRowSizeStyle](t_, objc.Sel("rowSizeStyle"))
	return rv
}

// Returns the row size style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecellview/1483211-rowsizestyle?language=objc
func (t_ TableCellView) SetRowSizeStyle(value TableViewRowSizeStyle) {
	objc.Call[objc.Void](t_, objc.Sel("setRowSizeStyle:"), value)
}

// Text displayed by the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecellview/1483202-textfield?language=objc
func (t_ TableCellView) TextField() TextField {
	rv := objc.Call[TextField](t_, objc.Sel("textField"))
	return rv
}

// Text displayed by the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecellview/1483202-textfield?language=objc
func (t_ TableCellView) SetTextField(value ITextField) {
	objc.Call[objc.Void](t_, objc.Sel("setTextField:"), objc.Ptr(value))
}

// The object that represents the cell data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecellview/1483204-objectvalue?language=objc
func (t_ TableCellView) ObjectValue() objc.Object {
	rv := objc.Call[objc.Object](t_, objc.Sel("objectValue"))
	return rv
}

// The object that represents the cell data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecellview/1483204-objectvalue?language=objc
func (t_ TableCellView) SetObjectValue(value objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("setObjectValue:"), value)
}

// Returns dragging images for the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecellview/1483199-draggingimagecomponents?language=objc
func (t_ TableCellView) DraggingImageComponents() []DraggingImageComponent {
	rv := objc.Call[[]DraggingImageComponent](t_, objc.Sel("draggingImageComponents"))
	return rv
}

// Image displayed by the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecellview/1483213-imageview?language=objc
func (t_ TableCellView) ImageView() ImageView {
	rv := objc.Call[ImageView](t_, objc.Sel("imageView"))
	return rv
}

// Image displayed by the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecellview/1483213-imageview?language=objc
func (t_ TableCellView) SetImageView(value IImageView) {
	objc.Call[objc.Void](t_, objc.Sel("setImageView:"), objc.Ptr(value))
}
