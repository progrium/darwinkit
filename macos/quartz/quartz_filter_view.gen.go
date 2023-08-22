// AUTO-GENERATED CODE, DO NOT MODIFY

package quartz

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [QuartzFilterView] class.
var QuartzFilterViewClass = _QuartzFilterViewClass{objc.GetClass("QuartzFilterView")}

type _QuartzFilterViewClass struct {
	objc.Class
}

// An interface definition for the [QuartzFilterView] class.
type IQuartzFilterView interface {
	appkit.IView
	SizeToFit()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/quartzfilterview?language=objc
type QuartzFilterView struct {
	appkit.View
}

func QuartzFilterViewFrom(ptr unsafe.Pointer) QuartzFilterView {
	return QuartzFilterView{
		View: appkit.ViewFrom(ptr),
	}
}

func (qc _QuartzFilterViewClass) Alloc() QuartzFilterView {
	rv := objc.Call[QuartzFilterView](qc, objc.Sel("alloc"))
	return rv
}

func QuartzFilterView_Alloc() QuartzFilterView {
	return QuartzFilterViewClass.Alloc()
}

func (qc _QuartzFilterViewClass) New() QuartzFilterView {
	rv := objc.Call[QuartzFilterView](qc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewQuartzFilterView() QuartzFilterView {
	return QuartzFilterViewClass.New()
}

func (q_ QuartzFilterView) Init() QuartzFilterView {
	rv := objc.Call[QuartzFilterView](q_, objc.Sel("init"))
	return rv
}

func (q_ QuartzFilterView) InitWithFrame(frameRect foundation.Rect) QuartzFilterView {
	rv := objc.Call[QuartzFilterView](q_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func NewQuartzFilterViewWithFrame(frameRect foundation.Rect) QuartzFilterView {
	instance := QuartzFilterViewClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/quartzfilterview/1505010-sizetofit?language=objc
func (q_ QuartzFilterView) SizeToFit() {
	objc.Call[objc.Void](q_, objc.Sel("sizeToFit"))
}
