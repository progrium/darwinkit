// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var ClipViewClass = _ClipViewClass{objc.GetClass("NSClipView")}

type _ClipViewClass struct {
	objc.Class
}

type IClipView interface {
	IView
	ScrollToPoint(newOrigin foundation.Point)
	ConstrainBoundsRect(proposedBounds foundation.Rect) foundation.Rect
	ViewBoundsChanged(notification foundation.INotification)
	ViewFrameChanged(notification foundation.INotification)
	DocumentView() View
	SetDocumentView(value IView)
	ContentInsets() foundation.EdgeInsets
	SetContentInsets(value foundation.EdgeInsets)
	AutomaticallyAdjustsContentInsets() bool
	SetAutomaticallyAdjustsContentInsets(value bool)
	DocumentRect() foundation.Rect
	DocumentVisibleRect() foundation.Rect
	DocumentCursor() Cursor
	SetDocumentCursor(value ICursor)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
}

type ClipView struct {
	View
}

func MakeClipView(ptr unsafe.Pointer) ClipView {
	return ClipView{
		View: MakeView(ptr),
	}
}

func (c_ ClipView) InitWithFrame(frameRect foundation.Rect) ClipView {
	rv := objc.CallMethod[ClipView](c_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func ClipView_InitWithFrame(frameRect foundation.Rect) ClipView {
	return ClipViewClass.Alloc().InitWithFrame(frameRect)
}

func (c_ ClipView) Init() ClipView {
	rv := objc.CallMethod[ClipView](c_, objc.GetSelector("init"))
	return rv
}

func ClipView_Init() ClipView {
	return ClipViewClass.Alloc().Init()
}

func (cc _ClipViewClass) Alloc() ClipView {
	rv := objc.CallMethod[ClipView](cc, objc.GetSelector("alloc"))
	return rv
}

func ClipView_Alloc() ClipView {
	return ClipViewClass.Alloc()
}

func (cc _ClipViewClass) New() ClipView {
	rv := objc.CallMethod[ClipView](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewClipView() ClipView {
	return ClipViewClass.New()
}

func ClipView_New() ClipView {
	return ClipViewClass.New()
}

func (c_ ClipView) ScrollToPoint(newOrigin foundation.Point) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("scrollToPoint:"), newOrigin)
}

func (c_ ClipView) ConstrainBoundsRect(proposedBounds foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](c_, objc.GetSelector("constrainBoundsRect:"), proposedBounds)
	return rv
}

func (c_ ClipView) ViewBoundsChanged(notification foundation.INotification) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("viewBoundsChanged:"), objc.ExtractPtr(notification))
}

func (c_ ClipView) ViewFrameChanged(notification foundation.INotification) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("viewFrameChanged:"), objc.ExtractPtr(notification))
}

func (c_ ClipView) DocumentView() View {
	rv := objc.CallMethod[View](c_, objc.GetSelector("documentView"))
	return rv
}

func (c_ ClipView) SetDocumentView(value IView) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setDocumentView:"), objc.ExtractPtr(value))
}

func (c_ ClipView) ContentInsets() foundation.EdgeInsets {
	rv := objc.CallMethod[foundation.EdgeInsets](c_, objc.GetSelector("contentInsets"))
	return rv
}

func (c_ ClipView) SetContentInsets(value foundation.EdgeInsets) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setContentInsets:"), value)
}

func (c_ ClipView) AutomaticallyAdjustsContentInsets() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("automaticallyAdjustsContentInsets"))
	return rv
}

func (c_ ClipView) SetAutomaticallyAdjustsContentInsets(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setAutomaticallyAdjustsContentInsets:"), value)
}

func (c_ ClipView) DocumentRect() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](c_, objc.GetSelector("documentRect"))
	return rv
}

func (c_ ClipView) DocumentVisibleRect() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](c_, objc.GetSelector("documentVisibleRect"))
	return rv
}

func (c_ ClipView) DocumentCursor() Cursor {
	rv := objc.CallMethod[Cursor](c_, objc.GetSelector("documentCursor"))
	return rv
}

func (c_ ClipView) SetDocumentCursor(value ICursor) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setDocumentCursor:"), objc.ExtractPtr(value))
}

func (c_ ClipView) DrawsBackground() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("drawsBackground"))
	return rv
}

func (c_ ClipView) SetDrawsBackground(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setDrawsBackground:"), value)
}

func (c_ ClipView) BackgroundColor() Color {
	rv := objc.CallMethod[Color](c_, objc.GetSelector("backgroundColor"))
	return rv
}

func (c_ ClipView) SetBackgroundColor(value IColor) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setBackgroundColor:"), objc.ExtractPtr(value))
}
