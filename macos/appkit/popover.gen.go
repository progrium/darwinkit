// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var PopoverClass = _PopoverClass{objc.GetClass("NSPopover")}

type _PopoverClass struct {
	objc.Class
}

type IPopover interface {
	IResponder
	ShowRelativeToRectOfViewPreferredEdge(positioningRect foundation.Rect, positioningView IView, preferredEdge foundation.RectEdge)
	PerformClose(sender objc.IObject)
	Close()
	ContentViewController() ViewController
	SetContentViewController(value IViewController)
	Behavior() PopoverBehavior
	SetBehavior(value PopoverBehavior)
	PositioningRect() foundation.Rect
	SetPositioningRect(value foundation.Rect)
	Appearance() Appearance
	SetAppearance(value IAppearance)
	EffectiveAppearance() Appearance
	Animates() bool
	SetAnimates(value bool)
	ContentSize() foundation.Size
	SetContentSize(value foundation.Size)
	IsShown() bool
	IsDetached() bool
	Delegate() PopoverDelegateWrapper
	SetDelegate(value IPopoverDelegate)
	SetDelegate0(value objc.IObject)
}

type Popover struct {
	Responder
}

func MakePopover(ptr unsafe.Pointer) Popover {
	return Popover{
		Responder: MakeResponder(ptr),
	}
}

func (p_ Popover) Init() Popover {
	rv := objc.CallMethod[Popover](p_, objc.GetSelector("init"))
	return rv
}

func Popover_Init() Popover {
	return PopoverClass.Alloc().Init()
}

func (pc _PopoverClass) Alloc() Popover {
	rv := objc.CallMethod[Popover](pc, objc.GetSelector("alloc"))
	return rv
}

func Popover_Alloc() Popover {
	return PopoverClass.Alloc()
}

func (pc _PopoverClass) New() Popover {
	rv := objc.CallMethod[Popover](pc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewPopover() Popover {
	return PopoverClass.New()
}

func Popover_New() Popover {
	return PopoverClass.New()
}

func (p_ Popover) ShowRelativeToRectOfViewPreferredEdge(positioningRect foundation.Rect, positioningView IView, preferredEdge foundation.RectEdge) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("showRelativeToRect:ofView:preferredEdge:"), positioningRect, objc.ExtractPtr(positioningView), preferredEdge)
}

func (p_ Popover) PerformClose(sender objc.IObject) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("performClose:"), objc.ExtractPtr(sender))
}

func (p_ Popover) Close() {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("close"))
}

func (p_ Popover) ContentViewController() ViewController {
	rv := objc.CallMethod[ViewController](p_, objc.GetSelector("contentViewController"))
	return rv
}

func (p_ Popover) SetContentViewController(value IViewController) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setContentViewController:"), objc.ExtractPtr(value))
}

func (p_ Popover) Behavior() PopoverBehavior {
	rv := objc.CallMethod[PopoverBehavior](p_, objc.GetSelector("behavior"))
	return rv
}

func (p_ Popover) SetBehavior(value PopoverBehavior) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setBehavior:"), value)
}

func (p_ Popover) PositioningRect() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](p_, objc.GetSelector("positioningRect"))
	return rv
}

func (p_ Popover) SetPositioningRect(value foundation.Rect) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setPositioningRect:"), value)
}

func (p_ Popover) Appearance() Appearance {
	rv := objc.CallMethod[Appearance](p_, objc.GetSelector("appearance"))
	return rv
}

func (p_ Popover) SetAppearance(value IAppearance) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setAppearance:"), objc.ExtractPtr(value))
}

func (p_ Popover) EffectiveAppearance() Appearance {
	rv := objc.CallMethod[Appearance](p_, objc.GetSelector("effectiveAppearance"))
	return rv
}

func (p_ Popover) Animates() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("animates"))
	return rv
}

func (p_ Popover) SetAnimates(value bool) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setAnimates:"), value)
}

func (p_ Popover) ContentSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](p_, objc.GetSelector("contentSize"))
	return rv
}

func (p_ Popover) SetContentSize(value foundation.Size) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setContentSize:"), value)
}

func (p_ Popover) IsShown() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("isShown"))
	return rv
}

func (p_ Popover) IsDetached() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("isDetached"))
	return rv
}

func (p_ Popover) Delegate() PopoverDelegateWrapper {
	rv := objc.CallMethod[PopoverDelegateWrapper](p_, objc.GetSelector("delegate"))
	return rv
}

func (p_ Popover) SetDelegate(value IPopoverDelegate) {
	po := objc.WrapAsProtocol("NSPopoverDelegate", value)
	objc.SetAssociatedObject(p_, objc.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setDelegate:"), po)
}

func (p_ Popover) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setDelegate:"), objc.ExtractPtr(value))
}
