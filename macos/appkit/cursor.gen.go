// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var CursorClass = _CursorClass{objc.GetClass("NSCursor")}

type _CursorClass struct {
	objc.Class
}

type ICursor interface {
	objc.IObject
	Push()
	Set()
	Image() Image
	HotSpot() foundation.Point
}

type Cursor struct {
	objc.Object
}

func MakeCursor(ptr unsafe.Pointer) Cursor {
	return Cursor{
		Object: objc.MakeObject(ptr),
	}
}

func (c_ Cursor) InitWithImageHotSpot(newImage IImage, point foundation.Point) Cursor {
	rv := objc.CallMethod[Cursor](c_, objc.GetSelector("initWithImage:hotSpot:"), objc.ExtractPtr(newImage), point)
	return rv
}

func Cursor_InitWithImageHotSpot(newImage IImage, point foundation.Point) Cursor {
	return CursorClass.Alloc().InitWithImageHotSpot(newImage, point)
}

func (c_ Cursor) InitWithImageForegroundColorHintBackgroundColorHintHotSpot(newImage IImage, fg IColor, bg IColor, hotSpot foundation.Point) Cursor {
	rv := objc.CallMethod[Cursor](c_, objc.GetSelector("initWithImage:foregroundColorHint:backgroundColorHint:hotSpot:"), objc.ExtractPtr(newImage), objc.ExtractPtr(fg), objc.ExtractPtr(bg), hotSpot)
	return rv
}

func Cursor_InitWithImageForegroundColorHintBackgroundColorHintHotSpot(newImage IImage, fg IColor, bg IColor, hotSpot foundation.Point) Cursor {
	return CursorClass.Alloc().InitWithImageForegroundColorHintBackgroundColorHintHotSpot(newImage, fg, bg, hotSpot)
}

func (cc _CursorClass) Alloc() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("alloc"))
	return rv
}

func Cursor_Alloc() Cursor {
	return CursorClass.Alloc()
}

func (cc _CursorClass) New() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCursor() Cursor {
	return CursorClass.New()
}

func Cursor_New() Cursor {
	return CursorClass.New()
}

func (c_ Cursor) Init() Cursor {
	rv := objc.CallMethod[Cursor](c_, objc.GetSelector("init"))
	return rv
}

func Cursor_Init() Cursor {
	return CursorClass.Alloc().Init()
}

func (cc _CursorClass) Hide() {
	objc.CallMethod[objc.Void](cc, objc.GetSelector("hide"))
}

func Cursor_Hide() {
	CursorClass.Hide()
}

func (cc _CursorClass) Unhide() {
	objc.CallMethod[objc.Void](cc, objc.GetSelector("unhide"))
}

func Cursor_Unhide() {
	CursorClass.Unhide()
}

func (cc _CursorClass) SetHiddenUntilMouseMoves(flag bool) {
	objc.CallMethod[objc.Void](cc, objc.GetSelector("setHiddenUntilMouseMoves:"), flag)
}

func Cursor_SetHiddenUntilMouseMoves(flag bool) {
	CursorClass.SetHiddenUntilMouseMoves(flag)
}

func (cc _CursorClass) Pop_() {
	objc.CallMethod[objc.Void](cc, objc.GetSelector("pop"))
}

func Cursor_Pop_() {
	CursorClass.Pop_()
}

func (c_ Cursor) Push() {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("push"))
}

func (c_ Cursor) Set() {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("set"))
}

func (c_ Cursor) Image() Image {
	rv := objc.CallMethod[Image](c_, objc.GetSelector("image"))
	return rv
}

func (c_ Cursor) HotSpot() foundation.Point {
	rv := objc.CallMethod[foundation.Point](c_, objc.GetSelector("hotSpot"))
	return rv
}

func (cc _CursorClass) CurrentCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("currentCursor"))
	return rv
}

func Cursor_CurrentCursor() Cursor {
	return CursorClass.CurrentCursor()
}

func (cc _CursorClass) CurrentSystemCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("currentSystemCursor"))
	return rv
}

func Cursor_CurrentSystemCursor() Cursor {
	return CursorClass.CurrentSystemCursor()
}

func (cc _CursorClass) ArrowCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("arrowCursor"))
	return rv
}

func Cursor_ArrowCursor() Cursor {
	return CursorClass.ArrowCursor()
}

func (cc _CursorClass) ContextualMenuCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("contextualMenuCursor"))
	return rv
}

func Cursor_ContextualMenuCursor() Cursor {
	return CursorClass.ContextualMenuCursor()
}

func (cc _CursorClass) ClosedHandCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("closedHandCursor"))
	return rv
}

func Cursor_ClosedHandCursor() Cursor {
	return CursorClass.ClosedHandCursor()
}

func (cc _CursorClass) CrosshairCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("crosshairCursor"))
	return rv
}

func Cursor_CrosshairCursor() Cursor {
	return CursorClass.CrosshairCursor()
}

func (cc _CursorClass) DisappearingItemCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("disappearingItemCursor"))
	return rv
}

func Cursor_DisappearingItemCursor() Cursor {
	return CursorClass.DisappearingItemCursor()
}

func (cc _CursorClass) DragCopyCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("dragCopyCursor"))
	return rv
}

func Cursor_DragCopyCursor() Cursor {
	return CursorClass.DragCopyCursor()
}

func (cc _CursorClass) DragLinkCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("dragLinkCursor"))
	return rv
}

func Cursor_DragLinkCursor() Cursor {
	return CursorClass.DragLinkCursor()
}

func (cc _CursorClass) IBeamCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("IBeamCursor"))
	return rv
}

func Cursor_IBeamCursor() Cursor {
	return CursorClass.IBeamCursor()
}

func (cc _CursorClass) OpenHandCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("openHandCursor"))
	return rv
}

func Cursor_OpenHandCursor() Cursor {
	return CursorClass.OpenHandCursor()
}

func (cc _CursorClass) OperationNotAllowedCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("operationNotAllowedCursor"))
	return rv
}

func Cursor_OperationNotAllowedCursor() Cursor {
	return CursorClass.OperationNotAllowedCursor()
}

func (cc _CursorClass) PointingHandCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("pointingHandCursor"))
	return rv
}

func Cursor_PointingHandCursor() Cursor {
	return CursorClass.PointingHandCursor()
}

func (cc _CursorClass) ResizeDownCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("resizeDownCursor"))
	return rv
}

func Cursor_ResizeDownCursor() Cursor {
	return CursorClass.ResizeDownCursor()
}

func (cc _CursorClass) ResizeLeftCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("resizeLeftCursor"))
	return rv
}

func Cursor_ResizeLeftCursor() Cursor {
	return CursorClass.ResizeLeftCursor()
}

func (cc _CursorClass) ResizeLeftRightCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("resizeLeftRightCursor"))
	return rv
}

func Cursor_ResizeLeftRightCursor() Cursor {
	return CursorClass.ResizeLeftRightCursor()
}

func (cc _CursorClass) ResizeRightCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("resizeRightCursor"))
	return rv
}

func Cursor_ResizeRightCursor() Cursor {
	return CursorClass.ResizeRightCursor()
}

func (cc _CursorClass) ResizeUpCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("resizeUpCursor"))
	return rv
}

func Cursor_ResizeUpCursor() Cursor {
	return CursorClass.ResizeUpCursor()
}

func (cc _CursorClass) ResizeUpDownCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("resizeUpDownCursor"))
	return rv
}

func Cursor_ResizeUpDownCursor() Cursor {
	return CursorClass.ResizeUpDownCursor()
}

func (cc _CursorClass) IBeamCursorForVerticalLayout() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("IBeamCursorForVerticalLayout"))
	return rv
}

func Cursor_IBeamCursorForVerticalLayout() Cursor {
	return CursorClass.IBeamCursorForVerticalLayout()
}
