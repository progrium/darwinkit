// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [OpenGLView] class.
var OpenGLViewClass = _OpenGLViewClass{objc.GetClass("NSOpenGLView")}

type _OpenGLViewClass struct {
	objc.Class
}

// An interface definition for the [OpenGLView] class.
type IOpenGLView interface {
	IView
}

// A view that displays OpenGL content in a view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenglview?language=objc
type OpenGLView struct {
	View
}

func OpenGLViewFrom(ptr unsafe.Pointer) OpenGLView {
	return OpenGLView{
		View: ViewFrom(ptr),
	}
}

func (oc _OpenGLViewClass) Alloc() OpenGLView {
	rv := objc.Call[OpenGLView](oc, objc.Sel("alloc"))
	return rv
}

func OpenGLView_Alloc() OpenGLView {
	return OpenGLViewClass.Alloc()
}

func (oc _OpenGLViewClass) New() OpenGLView {
	rv := objc.Call[OpenGLView](oc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewOpenGLView() OpenGLView {
	return OpenGLViewClass.New()
}

func (o_ OpenGLView) Init() OpenGLView {
	rv := objc.Call[OpenGLView](o_, objc.Sel("init"))
	return rv
}

func (o_ OpenGLView) InitWithFrame(frameRect foundation.Rect) OpenGLView {
	rv := objc.Call[OpenGLView](o_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func OpenGLView_InitWithFrame(frameRect foundation.Rect) OpenGLView {
	return OpenGLViewClass.Alloc().InitWithFrame(frameRect)
}
