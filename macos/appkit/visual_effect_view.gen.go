// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var VisualEffectViewClass = _VisualEffectViewClass{objc.GetClass("NSVisualEffectView")}

type _VisualEffectViewClass struct {
	objc.Class
}

type IVisualEffectView interface {
	IView
	Material() VisualEffectMaterial
	SetMaterial(value VisualEffectMaterial)
	BlendingMode() VisualEffectBlendingMode
	SetBlendingMode(value VisualEffectBlendingMode)
	IsEmphasized() bool
	SetEmphasized(value bool)
	InteriorBackgroundStyle() BackgroundStyle
	MaskImage() Image
	SetMaskImage(value IImage)
	State() VisualEffectState
	SetState(value VisualEffectState)
}

type VisualEffectView struct {
	View
}

func MakeVisualEffectView(ptr unsafe.Pointer) VisualEffectView {
	return VisualEffectView{
		View: MakeView(ptr),
	}
}

func (v_ VisualEffectView) InitWithFrame(frameRect foundation.Rect) VisualEffectView {
	rv := objc.CallMethod[VisualEffectView](v_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func VisualEffectView_InitWithFrame(frameRect foundation.Rect) VisualEffectView {
	return VisualEffectViewClass.Alloc().InitWithFrame(frameRect)
}

func (v_ VisualEffectView) Init() VisualEffectView {
	rv := objc.CallMethod[VisualEffectView](v_, objc.GetSelector("init"))
	return rv
}

func VisualEffectView_Init() VisualEffectView {
	return VisualEffectViewClass.Alloc().Init()
}

func (vc _VisualEffectViewClass) Alloc() VisualEffectView {
	rv := objc.CallMethod[VisualEffectView](vc, objc.GetSelector("alloc"))
	return rv
}

func VisualEffectView_Alloc() VisualEffectView {
	return VisualEffectViewClass.Alloc()
}

func (vc _VisualEffectViewClass) New() VisualEffectView {
	rv := objc.CallMethod[VisualEffectView](vc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewVisualEffectView() VisualEffectView {
	return VisualEffectViewClass.New()
}

func VisualEffectView_New() VisualEffectView {
	return VisualEffectViewClass.New()
}

func (v_ VisualEffectView) Material() VisualEffectMaterial {
	rv := objc.CallMethod[VisualEffectMaterial](v_, objc.GetSelector("material"))
	return rv
}

func (v_ VisualEffectView) SetMaterial(value VisualEffectMaterial) {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("setMaterial:"), value)
}

func (v_ VisualEffectView) BlendingMode() VisualEffectBlendingMode {
	rv := objc.CallMethod[VisualEffectBlendingMode](v_, objc.GetSelector("blendingMode"))
	return rv
}

func (v_ VisualEffectView) SetBlendingMode(value VisualEffectBlendingMode) {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("setBlendingMode:"), value)
}

func (v_ VisualEffectView) IsEmphasized() bool {
	rv := objc.CallMethod[bool](v_, objc.GetSelector("isEmphasized"))
	return rv
}

func (v_ VisualEffectView) SetEmphasized(value bool) {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("setEmphasized:"), value)
}

func (v_ VisualEffectView) InteriorBackgroundStyle() BackgroundStyle {
	rv := objc.CallMethod[BackgroundStyle](v_, objc.GetSelector("interiorBackgroundStyle"))
	return rv
}

func (v_ VisualEffectView) MaskImage() Image {
	rv := objc.CallMethod[Image](v_, objc.GetSelector("maskImage"))
	return rv
}

func (v_ VisualEffectView) SetMaskImage(value IImage) {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("setMaskImage:"), objc.ExtractPtr(value))
}

func (v_ VisualEffectView) State() VisualEffectState {
	rv := objc.CallMethod[VisualEffectState](v_, objc.GetSelector("state"))
	return rv
}

func (v_ VisualEffectView) SetState(value VisualEffectState) {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("setState:"), value)
}
