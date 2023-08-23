// AUTO-GENERATED CODE, DO NOT MODIFY

package quartz

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CompositionParameterView] class.
var CompositionParameterViewClass = _CompositionParameterViewClass{objc.GetClass("QCCompositionParameterView")}

type _CompositionParameterViewClass struct {
	objc.Class
}

// An interface definition for the [CompositionParameterView] class.
type ICompositionParameterView interface {
	appkit.IView
}

// A class that allows users to edit the input parameters of a composition in real time. The composition can be rendering in any of the following objects: QCRenderer, QCView, or QCCompositionLayer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/qccompositionparameterview?language=objc
type CompositionParameterView struct {
	appkit.View
}

func CompositionParameterViewFrom(ptr unsafe.Pointer) CompositionParameterView {
	return CompositionParameterView{
		View: appkit.ViewFrom(ptr),
	}
}

func (cc _CompositionParameterViewClass) Alloc() CompositionParameterView {
	rv := objc.Call[CompositionParameterView](cc, objc.Sel("alloc"))
	return rv
}

func CompositionParameterView_Alloc() CompositionParameterView {
	return CompositionParameterViewClass.Alloc()
}

func (cc _CompositionParameterViewClass) New() CompositionParameterView {
	rv := objc.Call[CompositionParameterView](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCompositionParameterView() CompositionParameterView {
	return CompositionParameterViewClass.New()
}

func (c_ CompositionParameterView) Init() CompositionParameterView {
	rv := objc.Call[CompositionParameterView](c_, objc.Sel("init"))
	return rv
}

func (c_ CompositionParameterView) InitWithFrame(frameRect foundation.Rect) CompositionParameterView {
	rv := objc.Call[CompositionParameterView](c_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func NewCompositionParameterViewWithFrame(frameRect foundation.Rect) CompositionParameterView {
	instance := CompositionParameterViewClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}
