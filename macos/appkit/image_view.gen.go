// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageView] class.
var ImageViewClass = _ImageViewClass{objc.GetClass("NSImageView")}

type _ImageViewClass struct {
	objc.Class
}

// An interface definition for the [ImageView] class.
type IImageView interface {
	IControl
	ContentTintColor() Color
	SetContentTintColor(value IColor)
	ImageAlignment() ImageAlignment
	SetImageAlignment(value ImageAlignment)
	AllowsCutCopyPaste() bool
	SetAllowsCutCopyPaste(value bool)
	IsEditable() bool
	SetEditable(value bool)
	ImageScaling() ImageScaling
	SetImageScaling(value ImageScaling)
	SymbolConfiguration() ImageSymbolConfiguration
	SetSymbolConfiguration(value IImageSymbolConfiguration)
	ImageFrameStyle() ImageFrameStyle
	SetImageFrameStyle(value ImageFrameStyle)
	Animates() bool
	SetAnimates(value bool)
	Image() Image
	SetImage(value IImage)
}

// A display of image data in a frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview?language=objc
type ImageView struct {
	Control
}

func ImageViewFrom(ptr unsafe.Pointer) ImageView {
	return ImageView{
		Control: ControlFrom(ptr),
	}
}

func (ic _ImageViewClass) ImageViewWithImage(image IImage) ImageView {
	rv := objc.Call[ImageView](ic, objc.Sel("imageViewWithImage:"), objc.Ptr(image))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/1644708-imageviewwithimage?language=objc
func ImageView_ImageViewWithImage(image IImage) ImageView {
	return ImageViewClass.ImageViewWithImage(image)
}

func (ic _ImageViewClass) Alloc() ImageView {
	rv := objc.Call[ImageView](ic, objc.Sel("alloc"))
	return rv
}

func ImageView_Alloc() ImageView {
	return ImageViewClass.Alloc()
}

func (ic _ImageViewClass) New() ImageView {
	rv := objc.Call[ImageView](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageView() ImageView {
	return ImageViewClass.New()
}

func (i_ ImageView) Init() ImageView {
	rv := objc.Call[ImageView](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageView) InitWithFrame(frameRect foundation.Rect) ImageView {
	rv := objc.Call[ImageView](i_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes a control with the specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428900-initwithframe?language=objc
func ImageView_InitWithFrame(frameRect foundation.Rect) ImageView {
	return ImageViewClass.Alloc().InitWithFrame(frameRect)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/3000783-contenttintcolor?language=objc
func (i_ ImageView) ContentTintColor() Color {
	rv := objc.Call[Color](i_, objc.Sel("contentTintColor"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/3000783-contenttintcolor?language=objc
func (i_ ImageView) SetContentTintColor(value IColor) {
	objc.Call[objc.Void](i_, objc.Sel("setContentTintColor:"), objc.Ptr(value))
}

// The alignment of the cell’s image inside the image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/1404963-imagealignment?language=objc
func (i_ ImageView) ImageAlignment() ImageAlignment {
	rv := objc.Call[ImageAlignment](i_, objc.Sel("imageAlignment"))
	return rv
}

// The alignment of the cell’s image inside the image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/1404963-imagealignment?language=objc
func (i_ ImageView) SetImageAlignment(value ImageAlignment) {
	objc.Call[objc.Void](i_, objc.Sel("setImageAlignment:"), value)
}

// A Boolean value indicating whether the image view lets the user cut, copy, and paste the image contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/1404961-allowscutcopypaste?language=objc
func (i_ ImageView) AllowsCutCopyPaste() bool {
	rv := objc.Call[bool](i_, objc.Sel("allowsCutCopyPaste"))
	return rv
}

// A Boolean value indicating whether the image view lets the user cut, copy, and paste the image contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/1404961-allowscutcopypaste?language=objc
func (i_ ImageView) SetAllowsCutCopyPaste(value bool) {
	objc.Call[objc.Void](i_, objc.Sel("setAllowsCutCopyPaste:"), value)
}

// A Boolean value indicating whether the user can drag a new image into the image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/1404954-editable?language=objc
func (i_ ImageView) IsEditable() bool {
	rv := objc.Call[bool](i_, objc.Sel("isEditable"))
	return rv
}

// A Boolean value indicating whether the user can drag a new image into the image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/1404954-editable?language=objc
func (i_ ImageView) SetEditable(value bool) {
	objc.Call[objc.Void](i_, objc.Sel("setEditable:"), value)
}

// The scaling mode applied to make the cell’s image fit the frame of the image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/1404956-imagescaling?language=objc
func (i_ ImageView) ImageScaling() ImageScaling {
	rv := objc.Call[ImageScaling](i_, objc.Sel("imageScaling"))
	return rv
}

// The scaling mode applied to make the cell’s image fit the frame of the image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/1404956-imagescaling?language=objc
func (i_ ImageView) SetImageScaling(value ImageScaling) {
	objc.Call[objc.Void](i_, objc.Sel("setImageScaling:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/3667456-symbolconfiguration?language=objc
func (i_ ImageView) SymbolConfiguration() ImageSymbolConfiguration {
	rv := objc.Call[ImageSymbolConfiguration](i_, objc.Sel("symbolConfiguration"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/3667456-symbolconfiguration?language=objc
func (i_ ImageView) SetSymbolConfiguration(value IImageSymbolConfiguration) {
	objc.Call[objc.Void](i_, objc.Sel("setSymbolConfiguration:"), objc.Ptr(value))
}

// The style of frame that appears around the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/1404948-imageframestyle?language=objc
func (i_ ImageView) ImageFrameStyle() ImageFrameStyle {
	rv := objc.Call[ImageFrameStyle](i_, objc.Sel("imageFrameStyle"))
	return rv
}

// The style of frame that appears around the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/1404948-imageframestyle?language=objc
func (i_ ImageView) SetImageFrameStyle(value ImageFrameStyle) {
	objc.Call[objc.Void](i_, objc.Sel("setImageFrameStyle:"), value)
}

// A Boolean value indicating whether the image view automatically plays animated images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/1404950-animates?language=objc
func (i_ ImageView) Animates() bool {
	rv := objc.Call[bool](i_, objc.Sel("animates"))
	return rv
}

// A Boolean value indicating whether the image view automatically plays animated images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/1404950-animates?language=objc
func (i_ ImageView) SetAnimates(value bool) {
	objc.Call[objc.Void](i_, objc.Sel("setAnimates:"), value)
}

// The image displayed by the image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/1404952-image?language=objc
func (i_ ImageView) Image() Image {
	rv := objc.Call[Image](i_, objc.Sel("image"))
	return rv
}

// The image displayed by the image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/1404952-image?language=objc
func (i_ ImageView) SetImage(value IImage) {
	objc.Call[objc.Void](i_, objc.Sel("setImage:"), objc.Ptr(value))
}
