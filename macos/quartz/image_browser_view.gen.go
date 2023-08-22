// AUTO-GENERATED CODE, DO NOT MODIFY

package quartz

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageBrowserView] class.
var ImageBrowserViewClass = _ImageBrowserViewClass{objc.GetClass("IKImageBrowserView")}

type _ImageBrowserViewClass struct {
	objc.Class
}

// An interface definition for the [ImageBrowserView] class.
type IImageBrowserView interface {
	appkit.IView
}

// The IKImageBrowserView class is a view for displaying and browsing a large amount of images and movies efficiently. This class will be deprecated in a future release. Please switch to NSCollectionView instead. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimagebrowserview?language=objc
type ImageBrowserView struct {
	appkit.View
}

func ImageBrowserViewFrom(ptr unsafe.Pointer) ImageBrowserView {
	return ImageBrowserView{
		View: appkit.ViewFrom(ptr),
	}
}

func (ic _ImageBrowserViewClass) Alloc() ImageBrowserView {
	rv := objc.Call[ImageBrowserView](ic, objc.Sel("alloc"))
	return rv
}

func ImageBrowserView_Alloc() ImageBrowserView {
	return ImageBrowserViewClass.Alloc()
}

func (ic _ImageBrowserViewClass) New() ImageBrowserView {
	rv := objc.Call[ImageBrowserView](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageBrowserView() ImageBrowserView {
	return ImageBrowserViewClass.New()
}

func (i_ ImageBrowserView) Init() ImageBrowserView {
	rv := objc.Call[ImageBrowserView](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageBrowserView) InitWithFrame(frameRect foundation.Rect) ImageBrowserView {
	rv := objc.Call[ImageBrowserView](i_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func NewImageBrowserViewWithFrame(frameRect foundation.Rect) ImageBrowserView {
	instance := ImageBrowserViewClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}
