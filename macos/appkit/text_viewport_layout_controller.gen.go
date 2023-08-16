// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TextViewportLayoutController] class.
var TextViewportLayoutControllerClass = _TextViewportLayoutControllerClass{objc.GetClass("NSTextViewportLayoutController")}

type _TextViewportLayoutControllerClass struct {
	objc.Class
}

// An interface definition for the [TextViewportLayoutController] class.
type ITextViewportLayoutController interface {
	objc.IObject
	AdjustViewportByVerticalOffset(verticalOffset float64)
	LayoutViewport()
	RelocateViewportToTextLocation(textLocation PTextLocation) float64
	RelocateViewportToTextLocationObject(textLocationObject objc.IObject) float64
	TextLayoutManager() TextLayoutManager
	ViewportRange() TextRange
	Delegate() TextViewportLayoutControllerDelegateWrapper
	SetDelegate(value PTextViewportLayoutControllerDelegate)
	SetDelegateObject(valueObject objc.IObject)
	ViewportBounds() coregraphics.Rect
}

// Manages the layout process inside the viewport interacting with its delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextviewportlayoutcontroller?language=objc
type TextViewportLayoutController struct {
	objc.Object
}

func TextViewportLayoutControllerFrom(ptr unsafe.Pointer) TextViewportLayoutController {
	return TextViewportLayoutController{
		Object: objc.ObjectFrom(ptr),
	}
}

func (t_ TextViewportLayoutController) InitWithTextLayoutManager(textLayoutManager ITextLayoutManager) TextViewportLayoutController {
	rv := objc.Call[TextViewportLayoutController](t_, objc.Sel("initWithTextLayoutManager:"), objc.Ptr(textLayoutManager))
	return rv
}

// Creates a new instance with the text layout manager you provide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextviewportlayoutcontroller/3824769-initwithtextlayoutmanager?language=objc
func TextViewportLayoutController_InitWithTextLayoutManager(textLayoutManager ITextLayoutManager) TextViewportLayoutController {
	return TextViewportLayoutControllerClass.Alloc().InitWithTextLayoutManager(textLayoutManager)
}

func (tc _TextViewportLayoutControllerClass) Alloc() TextViewportLayoutController {
	rv := objc.Call[TextViewportLayoutController](tc, objc.Sel("alloc"))
	return rv
}

func TextViewportLayoutController_Alloc() TextViewportLayoutController {
	return TextViewportLayoutControllerClass.Alloc()
}

func (tc _TextViewportLayoutControllerClass) New() TextViewportLayoutController {
	rv := objc.Call[TextViewportLayoutController](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTextViewportLayoutController() TextViewportLayoutController {
	return TextViewportLayoutControllerClass.New()
}

func (t_ TextViewportLayoutController) Init() TextViewportLayoutController {
	rv := objc.Call[TextViewportLayoutController](t_, objc.Sel("init"))
	return rv
}

// Adjusts the viewport rect by the specified offset if needed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextviewportlayoutcontroller/3852589-adjustviewportbyverticaloffset?language=objc
func (t_ TextViewportLayoutController) AdjustViewportByVerticalOffset(verticalOffset float64) {
	objc.Call[objc.Void](t_, objc.Sel("adjustViewportByVerticalOffset:"), verticalOffset)
}

// Performs layout in the viewport. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextviewportlayoutcontroller/3824770-layoutviewport?language=objc
func (t_ TextViewportLayoutController) LayoutViewport() {
	objc.Call[objc.Void](t_, objc.Sel("layoutViewport"))
}

// Relocates the viewport to the location you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextviewportlayoutcontroller/3852590-relocateviewporttotextlocation?language=objc
func (t_ TextViewportLayoutController) RelocateViewportToTextLocation(textLocation PTextLocation) float64 {
	po0 := objc.WrapAsProtocol("NSTextLocation", textLocation)
	rv := objc.Call[float64](t_, objc.Sel("relocateViewportToTextLocation:"), po0)
	return rv
}

// Relocates the viewport to the location you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextviewportlayoutcontroller/3852590-relocateviewporttotextlocation?language=objc
func (t_ TextViewportLayoutController) RelocateViewportToTextLocationObject(textLocationObject objc.IObject) float64 {
	rv := objc.Call[float64](t_, objc.Sel("relocateViewportToTextLocation:"), objc.Ptr(textLocationObject))
	return rv
}

// Returns the text layout manager for this viewport layout controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextviewportlayoutcontroller/3824772-textlayoutmanager?language=objc
func (t_ TextViewportLayoutController) TextLayoutManager() TextLayoutManager {
	rv := objc.Call[TextLayoutManager](t_, objc.Sel("textLayoutManager"))
	return rv
}

// Returns the text range of the current viewport layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextviewportlayoutcontroller/3824774-viewportrange?language=objc
func (t_ TextViewportLayoutController) ViewportRange() TextRange {
	rv := objc.Call[TextRange](t_, objc.Sel("viewportRange"))
	return rv
}

// The delegate for the text layout manager object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextviewportlayoutcontroller/3824768-delegate?language=objc
func (t_ TextViewportLayoutController) Delegate() TextViewportLayoutControllerDelegateWrapper {
	rv := objc.Call[TextViewportLayoutControllerDelegateWrapper](t_, objc.Sel("delegate"))
	return rv
}

// The delegate for the text layout manager object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextviewportlayoutcontroller/3824768-delegate?language=objc
func (t_ TextViewportLayoutController) SetDelegate(value PTextViewportLayoutControllerDelegate) {
	po0 := objc.WrapAsProtocol("NSTextViewportLayoutControllerDelegate", value)
	objc.SetAssociatedObject(t_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](t_, objc.Sel("setDelegate:"), po0)
}

// The delegate for the text layout manager object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextviewportlayoutcontroller/3824768-delegate?language=objc
func (t_ TextViewportLayoutController) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// Returns the visible bounds of the view, plus the overdraw area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextviewportlayoutcontroller/3824773-viewportbounds?language=objc
func (t_ TextViewportLayoutController) ViewportBounds() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](t_, objc.Sel("viewportBounds"))
	return rv
}
