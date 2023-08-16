// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [WebFrameView] class.
var WebFrameViewClass = _WebFrameViewClass{objc.GetClass("WebFrameView")}

type _WebFrameViewClass struct {
	objc.Class
}

// An interface definition for the [WebFrameView] class.
type IWebFrameView interface {
	appkit.IView
}

// WebFrameView objects and their subviews display the web content contained in a frame. You never create instances of WebFrameView directly—WebView objects create and manage a hierarchy of WebFrameView objects, one for each frame. WebFrameView objects use a scroll view whose document view conforms to the WebDocumentView protocol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webframeview?language=objc
type WebFrameView struct {
	appkit.View
}

func WebFrameViewFrom(ptr unsafe.Pointer) WebFrameView {
	return WebFrameView{
		View: appkit.ViewFrom(ptr),
	}
}

func (wc _WebFrameViewClass) Alloc() WebFrameView {
	rv := objc.Call[WebFrameView](wc, objc.Sel("alloc"))
	return rv
}

func WebFrameView_Alloc() WebFrameView {
	return WebFrameViewClass.Alloc()
}

func (wc _WebFrameViewClass) New() WebFrameView {
	rv := objc.Call[WebFrameView](wc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewWebFrameView() WebFrameView {
	return WebFrameViewClass.New()
}

func (w_ WebFrameView) Init() WebFrameView {
	rv := objc.Call[WebFrameView](w_, objc.Sel("init"))
	return rv
}

func (w_ WebFrameView) InitWithFrame(frameRect foundation.Rect) WebFrameView {
	rv := objc.Call[WebFrameView](w_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func WebFrameView_InitWithFrame(frameRect foundation.Rect) WebFrameView {
	return WebFrameViewClass.Alloc().InitWithFrame(frameRect)
}
