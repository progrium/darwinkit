// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Cursor] class.
var CursorClass = _CursorClass{objc.GetClass("NSCursor")}

type _CursorClass struct {
	objc.Class
}

// An interface definition for the [Cursor] class.
type ICursor interface {
	objc.IObject
	Set()
	Push()
	HotSpot() foundation.Point
	Image() Image
}

// A pointer (also called a cursor). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor?language=objc
type Cursor struct {
	objc.Object
}

func CursorFrom(ptr unsafe.Pointer) Cursor {
	return Cursor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CursorClass) Alloc() Cursor {
	rv := objc.Call[Cursor](cc, objc.Sel("alloc"))
	return rv
}

func Cursor_Alloc() Cursor {
	return CursorClass.Alloc()
}

func (cc _CursorClass) New() Cursor {
	rv := objc.Call[Cursor](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCursor() Cursor {
	return CursorClass.New()
}

func (c_ Cursor) Init() Cursor {
	rv := objc.Call[Cursor](c_, objc.Sel("init"))
	return rv
}

// Makes the current cursor invisible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1527345-hide?language=objc
func (cc _CursorClass) Hide() {
	objc.Call[objc.Void](cc, objc.Sel("hide"))
}

// Makes the current cursor invisible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1527345-hide?language=objc
func Cursor_Hide() {
	CursorClass.Hide()
}

// Makes the receiver the current cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1526148-set?language=objc
func (c_ Cursor) Set() {
	objc.Call[objc.Void](c_, objc.Sel("set"))
}

// Puts the receiver on top of the cursor stack and makes it the current cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1532500-push?language=objc
func (c_ Cursor) Push() {
	objc.Call[objc.Void](c_, objc.Sel("push"))
}

// Negates an earlier call to hide by showing the current cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1532996-unhide?language=objc
func (cc _CursorClass) Unhide() {
	objc.Call[objc.Void](cc, objc.Sel("unhide"))
}

// Negates an earlier call to hide by showing the current cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1532996-unhide?language=objc
func Cursor_Unhide() {
	CursorClass.Unhide()
}

// Pops the current cursor off the top of the stack. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1532104-pop?language=objc
func (cc _CursorClass) Pop() {
	objc.Call[objc.Void](cc, objc.Sel("pop"))
}

// Pops the current cursor off the top of the stack. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1532104-pop?language=objc
func Cursor_Pop() {
	CursorClass.Pop()
}

// Sets whether the cursor is hidden until the mouse moves. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1534665-sethiddenuntilmousemoves?language=objc
func (cc _CursorClass) SetHiddenUntilMouseMoves(flag bool) {
	objc.Call[objc.Void](cc, objc.Sel("setHiddenUntilMouseMoves:"), flag)
}

// Sets whether the cursor is hidden until the mouse moves. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1534665-sethiddenuntilmousemoves?language=objc
func Cursor_SetHiddenUntilMouseMoves(flag bool) {
	CursorClass.SetHiddenUntilMouseMoves(flag)
}

// Returns the resize-right system cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1526314-resizerightcursor?language=objc
func (cc _CursorClass) ResizeRightCursor() Cursor {
	rv := objc.Call[Cursor](cc, objc.Sel("resizeRightCursor"))
	return rv
}

// Returns the resize-right system cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1526314-resizerightcursor?language=objc
func Cursor_ResizeRightCursor() Cursor {
	return CursorClass.ResizeRightCursor()
}

// Returns the resize-left system cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1535416-resizeleftcursor?language=objc
func (cc _CursorClass) ResizeLeftCursor() Cursor {
	rv := objc.Call[Cursor](cc, objc.Sel("resizeLeftCursor"))
	return rv
}

// Returns the resize-left system cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1535416-resizeleftcursor?language=objc
func Cursor_ResizeLeftCursor() Cursor {
	return CursorClass.ResizeLeftCursor()
}

// Returns a cursor indicating that the current operation will result in a copy action. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1529900-dragcopycursor?language=objc
func (cc _CursorClass) DragCopyCursor() Cursor {
	rv := objc.Call[Cursor](cc, objc.Sel("dragCopyCursor"))
	return rv
}

// Returns a cursor indicating that the current operation will result in a copy action. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1529900-dragcopycursor?language=objc
func Cursor_DragCopyCursor() Cursor {
	return CursorClass.DragCopyCursor()
}

// Returns the resize-up-and-down system cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1524641-resizeupdowncursor?language=objc
func (cc _CursorClass) ResizeUpDownCursor() Cursor {
	rv := objc.Call[Cursor](cc, objc.Sel("resizeUpDownCursor"))
	return rv
}

// Returns the resize-up-and-down system cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1524641-resizeupdowncursor?language=objc
func Cursor_ResizeUpDownCursor() Cursor {
	return CursorClass.ResizeUpDownCursor()
}

// Returns a cursor indicating that the current operation will result in a disappearing item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1534280-disappearingitemcursor?language=objc
func (cc _CursorClass) DisappearingItemCursor() Cursor {
	rv := objc.Call[Cursor](cc, objc.Sel("disappearingItemCursor"))
	return rv
}

// Returns a cursor indicating that the current operation will result in a disappearing item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1534280-disappearingitemcursor?language=objc
func Cursor_DisappearingItemCursor() Cursor {
	return CursorClass.DisappearingItemCursor()
}

// Returns the pointing-hand system cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1531896-pointinghandcursor?language=objc
func (cc _CursorClass) PointingHandCursor() Cursor {
	rv := objc.Call[Cursor](cc, objc.Sel("pointingHandCursor"))
	return rv
}

// Returns the pointing-hand system cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1531896-pointinghandcursor?language=objc
func Cursor_PointingHandCursor() Cursor {
	return CursorClass.PointingHandCursor()
}

// Returns the cursor for editing vertical layout text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1525182-ibeamcursorforverticallayout?language=objc
func (cc _CursorClass) IBeamCursorForVerticalLayout() Cursor {
	rv := objc.Call[Cursor](cc, objc.Sel("IBeamCursorForVerticalLayout"))
	return rv
}

// Returns the cursor for editing vertical layout text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1525182-ibeamcursorforverticallayout?language=objc
func Cursor_IBeamCursorForVerticalLayout() Cursor {
	return CursorClass.IBeamCursorForVerticalLayout()
}

// Returns the resize-up system cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1532226-resizeupcursor?language=objc
func (cc _CursorClass) ResizeUpCursor() Cursor {
	rv := objc.Call[Cursor](cc, objc.Sel("resizeUpCursor"))
	return rv
}

// Returns the resize-up system cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1532226-resizeupcursor?language=objc
func Cursor_ResizeUpCursor() Cursor {
	return CursorClass.ResizeUpCursor()
}

// The position of the cursor's hot spot. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1529096-hotspot?language=objc
func (c_ Cursor) HotSpot() foundation.Point {
	rv := objc.Call[foundation.Point](c_, objc.Sel("hotSpot"))
	return rv
}

// Returns the contextual menu system cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1529142-contextualmenucursor?language=objc
func (cc _CursorClass) ContextualMenuCursor() Cursor {
	rv := objc.Call[Cursor](cc, objc.Sel("contextualMenuCursor"))
	return rv
}

// Returns the contextual menu system cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1529142-contextualmenucursor?language=objc
func Cursor_ContextualMenuCursor() Cursor {
	return CursorClass.ContextualMenuCursor()
}

// Returns the current system cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1533611-currentsystemcursor?language=objc
func (cc _CursorClass) CurrentSystemCursor() Cursor {
	rv := objc.Call[Cursor](cc, objc.Sel("currentSystemCursor"))
	return rv
}

// Returns the current system cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1533611-currentsystemcursor?language=objc
func Cursor_CurrentSystemCursor() Cursor {
	return CursorClass.CurrentSystemCursor()
}

// Returns the cross-hair system cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1525359-crosshaircursor?language=objc
func (cc _CursorClass) CrosshairCursor() Cursor {
	rv := objc.Call[Cursor](cc, objc.Sel("crosshairCursor"))
	return rv
}

// Returns the cross-hair system cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1525359-crosshaircursor?language=objc
func Cursor_CrosshairCursor() Cursor {
	return CursorClass.CrosshairCursor()
}

// Returns the default cursor, the arrow cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1527160-arrowcursor?language=objc
func (cc _CursorClass) ArrowCursor() Cursor {
	rv := objc.Call[Cursor](cc, objc.Sel("arrowCursor"))
	return rv
}

// Returns the default cursor, the arrow cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1527160-arrowcursor?language=objc
func Cursor_ArrowCursor() Cursor {
	return CursorClass.ArrowCursor()
}

// Returns a cursor that looks like a capital I with a tiny crossbeam at its middle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1526109-ibeamcursor?language=objc
func (cc _CursorClass) IBeamCursor() Cursor {
	rv := objc.Call[Cursor](cc, objc.Sel("IBeamCursor"))
	return rv
}

// Returns a cursor that looks like a capital I with a tiny crossbeam at its middle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1526109-ibeamcursor?language=objc
func Cursor_IBeamCursor() Cursor {
	return CursorClass.IBeamCursor()
}

// Returns the resize-down system cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1531340-resizedowncursor?language=objc
func (cc _CursorClass) ResizeDownCursor() Cursor {
	rv := objc.Call[Cursor](cc, objc.Sel("resizeDownCursor"))
	return rv
}

// Returns the resize-down system cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1531340-resizedowncursor?language=objc
func Cursor_ResizeDownCursor() Cursor {
	return CursorClass.ResizeDownCursor()
}

// Returns a cursor indicating that the current operation will result in a link action. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1534435-draglinkcursor?language=objc
func (cc _CursorClass) DragLinkCursor() Cursor {
	rv := objc.Call[Cursor](cc, objc.Sel("dragLinkCursor"))
	return rv
}

// Returns a cursor indicating that the current operation will result in a link action. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1534435-draglinkcursor?language=objc
func Cursor_DragLinkCursor() Cursor {
	return CursorClass.DragLinkCursor()
}

// Returns the closed-hand system cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1524603-closedhandcursor?language=objc
func (cc _CursorClass) ClosedHandCursor() Cursor {
	rv := objc.Call[Cursor](cc, objc.Sel("closedHandCursor"))
	return rv
}

// Returns the closed-hand system cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1524603-closedhandcursor?language=objc
func Cursor_ClosedHandCursor() Cursor {
	return CursorClass.ClosedHandCursor()
}

// Returns the operation not allowed cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1525180-operationnotallowedcursor?language=objc
func (cc _CursorClass) OperationNotAllowedCursor() Cursor {
	rv := objc.Call[Cursor](cc, objc.Sel("operationNotAllowedCursor"))
	return rv
}

// Returns the operation not allowed cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1525180-operationnotallowedcursor?language=objc
func Cursor_OperationNotAllowedCursor() Cursor {
	return CursorClass.OperationNotAllowedCursor()
}

// The cursor’s image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1527062-image?language=objc
func (c_ Cursor) Image() Image {
	rv := objc.Call[Image](c_, objc.Sel("image"))
	return rv
}

// Returns the resize-left-and-right system cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1525074-resizeleftrightcursor?language=objc
func (cc _CursorClass) ResizeLeftRightCursor() Cursor {
	rv := objc.Call[Cursor](cc, objc.Sel("resizeLeftRightCursor"))
	return rv
}

// Returns the resize-left-and-right system cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1525074-resizeleftrightcursor?language=objc
func Cursor_ResizeLeftRightCursor() Cursor {
	return CursorClass.ResizeLeftRightCursor()
}

// Returns the open-hand system cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1528540-openhandcursor?language=objc
func (cc _CursorClass) OpenHandCursor() Cursor {
	rv := objc.Call[Cursor](cc, objc.Sel("openHandCursor"))
	return rv
}

// Returns the open-hand system cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1528540-openhandcursor?language=objc
func Cursor_OpenHandCursor() Cursor {
	return CursorClass.OpenHandCursor()
}

// Returns the application’s current cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1524595-currentcursor?language=objc
func (cc _CursorClass) CurrentCursor() Cursor {
	rv := objc.Call[Cursor](cc, objc.Sel("currentCursor"))
	return rv
}

// Returns the application’s current cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/1524595-currentcursor?language=objc
func Cursor_CurrentCursor() Cursor {
	return CursorClass.CurrentCursor()
}
