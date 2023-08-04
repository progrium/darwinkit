// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var TextLayoutFragmentClass = _TextLayoutFragmentClass{objc.GetClass("NSTextLayoutFragment")}

type _TextLayoutFragmentClass struct {
	objc.Class
}

type ITextLayoutFragment interface {
	objc.IObject
	DrawAtPointInContext(point coregraphics.Point, context coregraphics.ContextRef)
	InvalidateLayout()
	BottomMargin() float64
	LeadingPadding() float64
	TopMargin() float64
	TrailingPadding() float64
	TextLayoutManager() TextLayoutManager
	LayoutQueue() foundation.OperationQueue
	SetLayoutQueue(value foundation.IOperationQueue)
	LayoutFragmentFrame() coregraphics.Rect
	RenderingSurfaceBounds() coregraphics.Rect
	TextAttachmentViewProviders() []TextAttachmentViewProvider
	State() TextLayoutFragmentState
	RangeInElement() TextRange
	TextElement() TextElement
	TextLineFragments() []TextLineFragment
}

type TextLayoutFragment struct {
	objc.Object
}

func MakeTextLayoutFragment(ptr unsafe.Pointer) TextLayoutFragment {
	return TextLayoutFragment{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TextLayoutFragment) InitWithTextElementRange(textElement ITextElement, rangeInElement ITextRange) TextLayoutFragment {
	rv := objc.CallMethod[TextLayoutFragment](t_, objc.GetSelector("initWithTextElement:range:"), objc.ExtractPtr(textElement), objc.ExtractPtr(rangeInElement))
	return rv
}

func TextLayoutFragment_InitWithTextElementRange(textElement ITextElement, rangeInElement ITextRange) TextLayoutFragment {
	return TextLayoutFragmentClass.Alloc().InitWithTextElementRange(textElement, rangeInElement)
}

func (tc _TextLayoutFragmentClass) Alloc() TextLayoutFragment {
	rv := objc.CallMethod[TextLayoutFragment](tc, objc.GetSelector("alloc"))
	return rv
}

func TextLayoutFragment_Alloc() TextLayoutFragment {
	return TextLayoutFragmentClass.Alloc()
}

func (tc _TextLayoutFragmentClass) New() TextLayoutFragment {
	rv := objc.CallMethod[TextLayoutFragment](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTextLayoutFragment() TextLayoutFragment {
	return TextLayoutFragmentClass.New()
}

func TextLayoutFragment_New() TextLayoutFragment {
	return TextLayoutFragmentClass.New()
}

func (t_ TextLayoutFragment) Init() TextLayoutFragment {
	rv := objc.CallMethod[TextLayoutFragment](t_, objc.GetSelector("init"))
	return rv
}

func TextLayoutFragment_Init() TextLayoutFragment {
	return TextLayoutFragmentClass.Alloc().Init()
}

func (t_ TextLayoutFragment) DrawAtPointInContext(point coregraphics.Point, context coregraphics.ContextRef) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("drawAtPoint:inContext:"), point, context)
}

func (t_ TextLayoutFragment) InvalidateLayout() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("invalidateLayout"))
}

func (t_ TextLayoutFragment) BottomMargin() float64 {
	rv := objc.CallMethod[float64](t_, objc.GetSelector("bottomMargin"))
	return rv
}

func (t_ TextLayoutFragment) LeadingPadding() float64 {
	rv := objc.CallMethod[float64](t_, objc.GetSelector("leadingPadding"))
	return rv
}

func (t_ TextLayoutFragment) TopMargin() float64 {
	rv := objc.CallMethod[float64](t_, objc.GetSelector("topMargin"))
	return rv
}

func (t_ TextLayoutFragment) TrailingPadding() float64 {
	rv := objc.CallMethod[float64](t_, objc.GetSelector("trailingPadding"))
	return rv
}

func (t_ TextLayoutFragment) TextLayoutManager() TextLayoutManager {
	rv := objc.CallMethod[TextLayoutManager](t_, objc.GetSelector("textLayoutManager"))
	return rv
}

func (t_ TextLayoutFragment) LayoutQueue() foundation.OperationQueue {
	rv := objc.CallMethod[foundation.OperationQueue](t_, objc.GetSelector("layoutQueue"))
	return rv
}

func (t_ TextLayoutFragment) SetLayoutQueue(value foundation.IOperationQueue) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setLayoutQueue:"), objc.ExtractPtr(value))
}

func (t_ TextLayoutFragment) LayoutFragmentFrame() coregraphics.Rect {
	rv := objc.CallMethod[coregraphics.Rect](t_, objc.GetSelector("layoutFragmentFrame"))
	return rv
}

func (t_ TextLayoutFragment) RenderingSurfaceBounds() coregraphics.Rect {
	rv := objc.CallMethod[coregraphics.Rect](t_, objc.GetSelector("renderingSurfaceBounds"))
	return rv
}

func (t_ TextLayoutFragment) TextAttachmentViewProviders() []TextAttachmentViewProvider {
	rv := objc.CallMethod[[]TextAttachmentViewProvider](t_, objc.GetSelector("textAttachmentViewProviders"))
	return rv
}

func (t_ TextLayoutFragment) State() TextLayoutFragmentState {
	rv := objc.CallMethod[TextLayoutFragmentState](t_, objc.GetSelector("state"))
	return rv
}

func (t_ TextLayoutFragment) RangeInElement() TextRange {
	rv := objc.CallMethod[TextRange](t_, objc.GetSelector("rangeInElement"))
	return rv
}

func (t_ TextLayoutFragment) TextElement() TextElement {
	rv := objc.CallMethod[TextElement](t_, objc.GetSelector("textElement"))
	return rv
}

func (t_ TextLayoutFragment) TextLineFragments() []TextLineFragment {
	rv := objc.CallMethod[[]TextLineFragment](t_, objc.GetSelector("textLineFragments"))
	return rv
}
