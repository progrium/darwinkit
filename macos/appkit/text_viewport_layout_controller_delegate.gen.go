// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// Optional methods that delegates implement to reposond to viewport layout changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextviewportlayoutcontrollerdelegate?language=objc
type PTextViewportLayoutControllerDelegate interface {
	// optional
	ViewportBoundsForTextViewportLayoutController(textViewportLayoutController TextViewportLayoutController) coregraphics.Rect
	HasViewportBoundsForTextViewportLayoutController() bool

	// optional
	TextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment(textViewportLayoutController TextViewportLayoutController, textLayoutFragment TextLayoutFragment)
	HasTextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment() bool

	// optional
	TextViewportLayoutControllerDidLayout(textViewportLayoutController TextViewportLayoutController)
	HasTextViewportLayoutControllerDidLayout() bool

	// optional
	TextViewportLayoutControllerWillLayout(textViewportLayoutController TextViewportLayoutController)
	HasTextViewportLayoutControllerWillLayout() bool
}

// A delegate implementation builder for the [PTextViewportLayoutControllerDelegate] protocol.
type TextViewportLayoutControllerDelegate struct {
	_ViewportBoundsForTextViewportLayoutController                              func(textViewportLayoutController TextViewportLayoutController) coregraphics.Rect
	_TextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment func(textViewportLayoutController TextViewportLayoutController, textLayoutFragment TextLayoutFragment)
	_TextViewportLayoutControllerDidLayout                                      func(textViewportLayoutController TextViewportLayoutController)
	_TextViewportLayoutControllerWillLayout                                     func(textViewportLayoutController TextViewportLayoutController)
}

func (di *TextViewportLayoutControllerDelegate) HasViewportBoundsForTextViewportLayoutController() bool {
	return di._ViewportBoundsForTextViewportLayoutController != nil
}

// Returns the current viewport, which is the view visible bounds plus the overdraw area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextviewportlayoutcontrollerdelegate/3824779-viewportboundsfortextviewportlay?language=objc
func (di *TextViewportLayoutControllerDelegate) SetViewportBoundsForTextViewportLayoutController(f func(textViewportLayoutController TextViewportLayoutController) coregraphics.Rect) {
	di._ViewportBoundsForTextViewportLayoutController = f
}

// Returns the current viewport, which is the view visible bounds plus the overdraw area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextviewportlayoutcontrollerdelegate/3824779-viewportboundsfortextviewportlay?language=objc
func (di *TextViewportLayoutControllerDelegate) ViewportBoundsForTextViewportLayoutController(textViewportLayoutController TextViewportLayoutController) coregraphics.Rect {
	return di._ViewportBoundsForTextViewportLayoutController(textViewportLayoutController)
}
func (di *TextViewportLayoutControllerDelegate) HasTextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment() bool {
	return di._TextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment != nil
}

// The method the framework calls when the layout controller lays out a text layout fragment in the UI. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextviewportlayoutcontrollerdelegate/3824776-textviewportlayoutcontroller?language=objc
func (di *TextViewportLayoutControllerDelegate) SetTextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment(f func(textViewportLayoutController TextViewportLayoutController, textLayoutFragment TextLayoutFragment)) {
	di._TextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment = f
}

// The method the framework calls when the layout controller lays out a text layout fragment in the UI. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextviewportlayoutcontrollerdelegate/3824776-textviewportlayoutcontroller?language=objc
func (di *TextViewportLayoutControllerDelegate) TextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment(textViewportLayoutController TextViewportLayoutController, textLayoutFragment TextLayoutFragment) {
	di._TextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment(textViewportLayoutController, textLayoutFragment)
}
func (di *TextViewportLayoutControllerDelegate) HasTextViewportLayoutControllerDidLayout() bool {
	return di._TextViewportLayoutControllerDidLayout != nil
}

// The method the framework calls when the text viewport layout controller finishes its layout process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextviewportlayoutcontrollerdelegate/3824777-textviewportlayoutcontrollerdidl?language=objc
func (di *TextViewportLayoutControllerDelegate) SetTextViewportLayoutControllerDidLayout(f func(textViewportLayoutController TextViewportLayoutController)) {
	di._TextViewportLayoutControllerDidLayout = f
}

// The method the framework calls when the text viewport layout controller finishes its layout process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextviewportlayoutcontrollerdelegate/3824777-textviewportlayoutcontrollerdidl?language=objc
func (di *TextViewportLayoutControllerDelegate) TextViewportLayoutControllerDidLayout(textViewportLayoutController TextViewportLayoutController) {
	di._TextViewportLayoutControllerDidLayout(textViewportLayoutController)
}
func (di *TextViewportLayoutControllerDelegate) HasTextViewportLayoutControllerWillLayout() bool {
	return di._TextViewportLayoutControllerWillLayout != nil
}

// The method the framework calls before the text viewport layout controller starts its layout process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextviewportlayoutcontrollerdelegate/3824778-textviewportlayoutcontrollerwill?language=objc
func (di *TextViewportLayoutControllerDelegate) SetTextViewportLayoutControllerWillLayout(f func(textViewportLayoutController TextViewportLayoutController)) {
	di._TextViewportLayoutControllerWillLayout = f
}

// The method the framework calls before the text viewport layout controller starts its layout process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextviewportlayoutcontrollerdelegate/3824778-textviewportlayoutcontrollerwill?language=objc
func (di *TextViewportLayoutControllerDelegate) TextViewportLayoutControllerWillLayout(textViewportLayoutController TextViewportLayoutController) {
	di._TextViewportLayoutControllerWillLayout(textViewportLayoutController)
}

// A concrete type wrapper for the [PTextViewportLayoutControllerDelegate] protocol.
type TextViewportLayoutControllerDelegateWrapper struct {
	objc.Object
}

func (t_ TextViewportLayoutControllerDelegateWrapper) HasViewportBoundsForTextViewportLayoutController() bool {
	return t_.RespondsToSelector(objc.Sel("viewportBoundsForTextViewportLayoutController:"))
}

// Returns the current viewport, which is the view visible bounds plus the overdraw area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextviewportlayoutcontrollerdelegate/3824779-viewportboundsfortextviewportlay?language=objc
func (t_ TextViewportLayoutControllerDelegateWrapper) ViewportBoundsForTextViewportLayoutController(textViewportLayoutController ITextViewportLayoutController) coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](t_, objc.Sel("viewportBoundsForTextViewportLayoutController:"), objc.Ptr(textViewportLayoutController))
	return rv
}

func (t_ TextViewportLayoutControllerDelegateWrapper) HasTextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment() bool {
	return t_.RespondsToSelector(objc.Sel("textViewportLayoutController:configureRenderingSurfaceForTextLayoutFragment:"))
}

// The method the framework calls when the layout controller lays out a text layout fragment in the UI. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextviewportlayoutcontrollerdelegate/3824776-textviewportlayoutcontroller?language=objc
func (t_ TextViewportLayoutControllerDelegateWrapper) TextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment(textViewportLayoutController ITextViewportLayoutController, textLayoutFragment ITextLayoutFragment) {
	objc.Call[objc.Void](t_, objc.Sel("textViewportLayoutController:configureRenderingSurfaceForTextLayoutFragment:"), objc.Ptr(textViewportLayoutController), objc.Ptr(textLayoutFragment))
}

func (t_ TextViewportLayoutControllerDelegateWrapper) HasTextViewportLayoutControllerDidLayout() bool {
	return t_.RespondsToSelector(objc.Sel("textViewportLayoutControllerDidLayout:"))
}

// The method the framework calls when the text viewport layout controller finishes its layout process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextviewportlayoutcontrollerdelegate/3824777-textviewportlayoutcontrollerdidl?language=objc
func (t_ TextViewportLayoutControllerDelegateWrapper) TextViewportLayoutControllerDidLayout(textViewportLayoutController ITextViewportLayoutController) {
	objc.Call[objc.Void](t_, objc.Sel("textViewportLayoutControllerDidLayout:"), objc.Ptr(textViewportLayoutController))
}

func (t_ TextViewportLayoutControllerDelegateWrapper) HasTextViewportLayoutControllerWillLayout() bool {
	return t_.RespondsToSelector(objc.Sel("textViewportLayoutControllerWillLayout:"))
}

// The method the framework calls before the text viewport layout controller starts its layout process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextviewportlayoutcontrollerdelegate/3824778-textviewportlayoutcontrollerwill?language=objc
func (t_ TextViewportLayoutControllerDelegateWrapper) TextViewportLayoutControllerWillLayout(textViewportLayoutController ITextViewportLayoutController) {
	objc.Call[objc.Void](t_, objc.Sel("textViewportLayoutControllerWillLayout:"), objc.Ptr(textViewportLayoutController))
}