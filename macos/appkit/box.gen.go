// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Box] class.
var BoxClass = _BoxClass{objc.GetClass("NSBox")}

type _BoxClass struct {
	objc.Class
}

// An interface definition for the [Box] class.
type IBox interface {
	IView
	SizeToFit()
	SetFrameFromContentFrame(contentFrame foundation.Rect)
	IsTransparent() bool
	SetTransparent(value bool)
	TitlePosition() TitlePosition
	SetTitlePosition(value TitlePosition)
	ContentView() View
	SetContentView(value IView)
	TitleRect() foundation.Rect
	TitleCell() objc.Object
	CornerRadius() float64
	SetCornerRadius(value float64)
	FillColor() Color
	SetFillColor(value IColor)
	BoxType() BoxType
	SetBoxType(value BoxType)
	Title() string
	SetTitle(value string)
	BorderColor() Color
	SetBorderColor(value IColor)
	BorderWidth() float64
	SetBorderWidth(value float64)
	TitleFont() Font
	SetTitleFont(value IFont)
	ContentViewMargins() foundation.Size
	SetContentViewMargins(value foundation.Size)
	BorderRect() foundation.Rect
}

// A stylized rectangular box with an optional title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbox?language=objc
type Box struct {
	View
}

func BoxFrom(ptr unsafe.Pointer) Box {
	return Box{
		View: ViewFrom(ptr),
	}
}

func (bc _BoxClass) Alloc() Box {
	rv := objc.Call[Box](bc, objc.Sel("alloc"))
	return rv
}

func Box_Alloc() Box {
	return BoxClass.Alloc()
}

func (bc _BoxClass) New() Box {
	rv := objc.Call[Box](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewBox() Box {
	return BoxClass.New()
}

func (b_ Box) Init() Box {
	rv := objc.Call[Box](b_, objc.Sel("init"))
	return rv
}

func (b_ Box) InitWithFrame(frameRect foundation.Rect) Box {
	rv := objc.Call[Box](b_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func Box_InitWithFrame(frameRect foundation.Rect) Box {
	return BoxClass.Alloc().InitWithFrame(frameRect)
}

// Resizes and moves the receiver’s content view so it just encloses its subviews. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbox/1429826-sizetofit?language=objc
func (b_ Box) SizeToFit() {
	objc.Call[objc.Void](b_, objc.Sel("sizeToFit"))
}

// Places the receiver so its content view lies on the specified frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbox/1429816-setframefromcontentframe?language=objc
func (b_ Box) SetFrameFromContentFrame(contentFrame foundation.Rect) {
	objc.Call[objc.Void](b_, objc.Sel("setFrameFromContentFrame:"), contentFrame)
}

// A Boolean value that indicates whether the receiver is transparent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbox/1429821-transparent?language=objc
func (b_ Box) IsTransparent() bool {
	rv := objc.Call[bool](b_, objc.Sel("isTransparent"))
	return rv
}

// A Boolean value that indicates whether the receiver is transparent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbox/1429821-transparent?language=objc
func (b_ Box) SetTransparent(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setTransparent:"), value)
}

// A constant representing the title position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbox/1429844-titleposition?language=objc
func (b_ Box) TitlePosition() TitlePosition {
	rv := objc.Call[TitlePosition](b_, objc.Sel("titlePosition"))
	return rv
}

// A constant representing the title position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbox/1429844-titleposition?language=objc
func (b_ Box) SetTitlePosition(value TitlePosition) {
	objc.Call[objc.Void](b_, objc.Sel("setTitlePosition:"), value)
}

// The receiver’s content view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbox/1429818-contentview?language=objc
func (b_ Box) ContentView() View {
	rv := objc.Call[View](b_, objc.Sel("contentView"))
	return rv
}

// The receiver’s content view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbox/1429818-contentview?language=objc
func (b_ Box) SetContentView(value IView) {
	objc.Call[objc.Void](b_, objc.Sel("setContentView:"), objc.Ptr(value))
}

// The rectangle in which the receiver’s title is drawn. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbox/1429785-titlerect?language=objc
func (b_ Box) TitleRect() foundation.Rect {
	rv := objc.Call[foundation.Rect](b_, objc.Sel("titleRect"))
	return rv
}

// The cell used to display the receiver’s title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbox/1429789-titlecell?language=objc
func (b_ Box) TitleCell() objc.Object {
	rv := objc.Call[objc.Object](b_, objc.Sel("titleCell"))
	return rv
}

// The radius of the receiver’s corners when the receiver is a custom box with a simple line border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbox/1429812-cornerradius?language=objc
func (b_ Box) CornerRadius() float64 {
	rv := objc.Call[float64](b_, objc.Sel("cornerRadius"))
	return rv
}

// The radius of the receiver’s corners when the receiver is a custom box with a simple line border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbox/1429812-cornerradius?language=objc
func (b_ Box) SetCornerRadius(value float64) {
	objc.Call[objc.Void](b_, objc.Sel("setCornerRadius:"), value)
}

// The color of the receiver’s background when the receiver is a custom box with a simple line border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbox/1429797-fillcolor?language=objc
func (b_ Box) FillColor() Color {
	rv := objc.Call[Color](b_, objc.Sel("fillColor"))
	return rv
}

// The color of the receiver’s background when the receiver is a custom box with a simple line border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbox/1429797-fillcolor?language=objc
func (b_ Box) SetFillColor(value IColor) {
	objc.Call[objc.Void](b_, objc.Sel("setFillColor:"), objc.Ptr(value))
}

// The receiver’s box type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbox/1429822-boxtype?language=objc
func (b_ Box) BoxType() BoxType {
	rv := objc.Call[BoxType](b_, objc.Sel("boxType"))
	return rv
}

// The receiver’s box type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbox/1429822-boxtype?language=objc
func (b_ Box) SetBoxType(value BoxType) {
	objc.Call[objc.Void](b_, objc.Sel("setBoxType:"), value)
}

// The receiver’s title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbox/1429795-title?language=objc
func (b_ Box) Title() string {
	rv := objc.Call[string](b_, objc.Sel("title"))
	return rv
}

// The receiver’s title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbox/1429795-title?language=objc
func (b_ Box) SetTitle(value string) {
	objc.Call[objc.Void](b_, objc.Sel("setTitle:"), value)
}

// The color of the receiver’s border when the receiver is a custom box with a simple line border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbox/1429839-bordercolor?language=objc
func (b_ Box) BorderColor() Color {
	rv := objc.Call[Color](b_, objc.Sel("borderColor"))
	return rv
}

// The color of the receiver’s border when the receiver is a custom box with a simple line border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbox/1429839-bordercolor?language=objc
func (b_ Box) SetBorderColor(value IColor) {
	objc.Call[objc.Void](b_, objc.Sel("setBorderColor:"), objc.Ptr(value))
}

// The width of the receiver’s border when the receiver is a custom box with a simple line border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbox/1429831-borderwidth?language=objc
func (b_ Box) BorderWidth() float64 {
	rv := objc.Call[float64](b_, objc.Sel("borderWidth"))
	return rv
}

// The width of the receiver’s border when the receiver is a custom box with a simple line border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbox/1429831-borderwidth?language=objc
func (b_ Box) SetBorderWidth(value float64) {
	objc.Call[objc.Void](b_, objc.Sel("setBorderWidth:"), value)
}

// The font object used to draw the receiver’s title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbox/1429791-titlefont?language=objc
func (b_ Box) TitleFont() Font {
	rv := objc.Call[Font](b_, objc.Sel("titleFont"))
	return rv
}

// The font object used to draw the receiver’s title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbox/1429791-titlefont?language=objc
func (b_ Box) SetTitleFont(value IFont) {
	objc.Call[objc.Void](b_, objc.Sel("setTitleFont:"), objc.Ptr(value))
}

// The distances between the border and the content view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbox/1429837-contentviewmargins?language=objc
func (b_ Box) ContentViewMargins() foundation.Size {
	rv := objc.Call[foundation.Size](b_, objc.Sel("contentViewMargins"))
	return rv
}

// The distances between the border and the content view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbox/1429837-contentviewmargins?language=objc
func (b_ Box) SetContentViewMargins(value foundation.Size) {
	objc.Call[objc.Void](b_, objc.Sel("setContentViewMargins:"), value)
}

// The rectangle in which the receiver’s border is drawn. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbox/1429787-borderrect?language=objc
func (b_ Box) BorderRect() foundation.Rect {
	rv := objc.Call[foundation.Rect](b_, objc.Sel("borderRect"))
	return rv
}
