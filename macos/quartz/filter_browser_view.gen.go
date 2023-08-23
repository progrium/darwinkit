// AUTO-GENERATED CODE, DO NOT MODIFY

package quartz

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FilterBrowserView] class.
var FilterBrowserViewClass = _FilterBrowserViewClass{objc.GetClass("IKFilterBrowserView")}

type _FilterBrowserViewClass struct {
	objc.Class
}

// An interface definition for the [FilterBrowserView] class.
type IFilterBrowserView interface {
	appkit.IView
	SetPreviewState(inState bool)
	FilterName() string
}

// The IKFilterBrowserView class is used as a container for the elements of an IKFilterBrowserPanel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikfilterbrowserview?language=objc
type FilterBrowserView struct {
	appkit.View
}

func FilterBrowserViewFrom(ptr unsafe.Pointer) FilterBrowserView {
	return FilterBrowserView{
		View: appkit.ViewFrom(ptr),
	}
}

func (fc _FilterBrowserViewClass) Alloc() FilterBrowserView {
	rv := objc.Call[FilterBrowserView](fc, objc.Sel("alloc"))
	return rv
}

func FilterBrowserView_Alloc() FilterBrowserView {
	return FilterBrowserViewClass.Alloc()
}

func (fc _FilterBrowserViewClass) New() FilterBrowserView {
	rv := objc.Call[FilterBrowserView](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFilterBrowserView() FilterBrowserView {
	return FilterBrowserViewClass.New()
}

func (f_ FilterBrowserView) Init() FilterBrowserView {
	rv := objc.Call[FilterBrowserView](f_, objc.Sel("init"))
	return rv
}

func (f_ FilterBrowserView) InitWithFrame(frameRect foundation.Rect) FilterBrowserView {
	rv := objc.Call[FilterBrowserView](f_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func NewFilterBrowserViewWithFrame(frameRect foundation.Rect) FilterBrowserView {
	instance := FilterBrowserViewClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

// Sets the preview state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikfilterbrowserview/1405296-setpreviewstate?language=objc
func (f_ FilterBrowserView) SetPreviewState(inState bool) {
	objc.Call[objc.Void](f_, objc.Sel("setPreviewState:"), inState)
}

// Returns the name of the  filter that is currently selected in the filter browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikfilterbrowserview/1405294-filtername?language=objc
func (f_ FilterBrowserView) FilterName() string {
	rv := objc.Call[string](f_, objc.Sel("filterName"))
	return rv
}
