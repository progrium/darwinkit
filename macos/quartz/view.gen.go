// AUTO-GENERATED CODE, DO NOT MODIFY

package quartz

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [View] class.
var ViewClass = _ViewClass{objc.GetClass("QCView")}

type _ViewClass struct {
	objc.Class
}

// An interface definition for the [View] class.
type IView interface {
	appkit.IView
}

// The QCView class is a custom NSView class that loads, plays, and controls Quartz Composer compositions. It is an autonomous view that is driven by an internal timer running on the main thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/qcview?language=objc
type View struct {
	appkit.View
}

func ViewFrom(ptr unsafe.Pointer) View {
	return View{
		View: appkit.ViewFrom(ptr),
	}
}

func (vc _ViewClass) Alloc() View {
	rv := objc.Call[View](vc, objc.Sel("alloc"))
	return rv
}

func View_Alloc() View {
	return ViewClass.Alloc()
}

func (vc _ViewClass) New() View {
	rv := objc.Call[View](vc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewView() View {
	return ViewClass.New()
}

func (v_ View) Init() View {
	rv := objc.Call[View](v_, objc.Sel("init"))
	return rv
}

func (v_ View) InitWithFrame(frameRect foundation.Rect) View {
	rv := objc.Call[View](v_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func NewViewWithFrame(frameRect foundation.Rect) View {
	instance := ViewClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}
