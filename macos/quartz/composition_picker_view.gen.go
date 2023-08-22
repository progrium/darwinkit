// AUTO-GENERATED CODE, DO NOT MODIFY

package quartz

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CompositionPickerView] class.
var CompositionPickerViewClass = _CompositionPickerViewClass{objc.GetClass("QCCompositionPickerView")}

type _CompositionPickerViewClass struct {
	objc.Class
}

// An interface definition for the [CompositionPickerView] class.
type ICompositionPickerView interface {
	appkit.IView
}

// The QCCompositionPickerView class allows users to browse compositions that are in the Quartz Composer composition repository, and to preview them. You can set the default input parameters for a composition preview  by using the method setDefaultValue:forInputKey:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/qccompositionpickerview?language=objc
type CompositionPickerView struct {
	appkit.View
}

func CompositionPickerViewFrom(ptr unsafe.Pointer) CompositionPickerView {
	return CompositionPickerView{
		View: appkit.ViewFrom(ptr),
	}
}

func (cc _CompositionPickerViewClass) Alloc() CompositionPickerView {
	rv := objc.Call[CompositionPickerView](cc, objc.Sel("alloc"))
	return rv
}

func CompositionPickerView_Alloc() CompositionPickerView {
	return CompositionPickerViewClass.Alloc()
}

func (cc _CompositionPickerViewClass) New() CompositionPickerView {
	rv := objc.Call[CompositionPickerView](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCompositionPickerView() CompositionPickerView {
	return CompositionPickerViewClass.New()
}

func (c_ CompositionPickerView) Init() CompositionPickerView {
	rv := objc.Call[CompositionPickerView](c_, objc.Sel("init"))
	return rv
}

func (c_ CompositionPickerView) InitWithFrame(frameRect foundation.Rect) CompositionPickerView {
	rv := objc.Call[CompositionPickerView](c_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func NewCompositionPickerViewWithFrame(frameRect foundation.Rect) CompositionPickerView {
	instance := CompositionPickerViewClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}
