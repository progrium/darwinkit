// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ScrubberImageItemView] class.
var ScrubberImageItemViewClass = _ScrubberImageItemViewClass{objc.GetClass("NSScrubberImageItemView")}

type _ScrubberImageItemViewClass struct {
	objc.Class
}

// An interface definition for the [ScrubberImageItemView] class.
type IScrubberImageItemView interface {
	IScrubberItemView
	ImageAlignment() ImageAlignment
	SetImageAlignment(value ImageAlignment)
	ImageView() ImageView
	Image() Image
	SetImage(value IImage)
}

// A concrete view subclass for displaying images in a scrubber items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberimageitemview?language=objc
type ScrubberImageItemView struct {
	ScrubberItemView
}

func ScrubberImageItemViewFrom(ptr unsafe.Pointer) ScrubberImageItemView {
	return ScrubberImageItemView{
		ScrubberItemView: ScrubberItemViewFrom(ptr),
	}
}

func (sc _ScrubberImageItemViewClass) Alloc() ScrubberImageItemView {
	rv := objc.Call[ScrubberImageItemView](sc, objc.Sel("alloc"))
	return rv
}

func ScrubberImageItemView_Alloc() ScrubberImageItemView {
	return ScrubberImageItemViewClass.Alloc()
}

func (sc _ScrubberImageItemViewClass) New() ScrubberImageItemView {
	rv := objc.Call[ScrubberImageItemView](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewScrubberImageItemView() ScrubberImageItemView {
	return ScrubberImageItemViewClass.New()
}

func (s_ ScrubberImageItemView) Init() ScrubberImageItemView {
	rv := objc.Call[ScrubberImageItemView](s_, objc.Sel("init"))
	return rv
}

func (s_ ScrubberImageItemView) InitWithFrame(frameRect foundation.Rect) ScrubberImageItemView {
	rv := objc.Call[ScrubberImageItemView](s_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func NewScrubberImageItemViewWithFrame(frameRect foundation.Rect) ScrubberImageItemView {
	instance := ScrubberImageItemViewClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

// The alignment of the image within the scrubber item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberimageitemview/2544674-imagealignment?language=objc
func (s_ ScrubberImageItemView) ImageAlignment() ImageAlignment {
	rv := objc.Call[ImageAlignment](s_, objc.Sel("imageAlignment"))
	return rv
}

// The alignment of the image within the scrubber item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberimageitemview/2544674-imagealignment?language=objc
func (s_ ScrubberImageItemView) SetImageAlignment(value ImageAlignment) {
	objc.Call[objc.Void](s_, objc.Sel("setImageAlignment:"), value)
}

// The image view that the scrubber item uses to display its image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberimageitemview/2544876-imageview?language=objc
func (s_ ScrubberImageItemView) ImageView() ImageView {
	rv := objc.Call[ImageView](s_, objc.Sel("imageView"))
	return rv
}

// The image displayed by the scrubber item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberimageitemview/2544761-image?language=objc
func (s_ ScrubberImageItemView) Image() Image {
	rv := objc.Call[Image](s_, objc.Sel("image"))
	return rv
}

// The image displayed by the scrubber item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberimageitemview/2544761-image?language=objc
func (s_ ScrubberImageItemView) SetImage(value IImage) {
	objc.Call[objc.Void](s_, objc.Sel("setImage:"), objc.Ptr(value))
}
