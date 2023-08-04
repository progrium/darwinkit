// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var LayoutGuideClass = _LayoutGuideClass{objc.GetClass("NSLayoutGuide")}

type _LayoutGuideClass struct {
	objc.Class
}

type ILayoutGuide interface {
	objc.IObject
	ConstraintsAffectingLayoutForOrientation(orientation LayoutConstraintOrientation) []LayoutConstraint
	Identifier() UserInterfaceItemIdentifier
	SetIdentifier(value UserInterfaceItemIdentifier)
	Frame() foundation.Rect
	OwningView() View
	SetOwningView(value IView)
	BottomAnchor() LayoutYAxisAnchor
	CenterXAnchor() LayoutXAxisAnchor
	CenterYAnchor() LayoutYAxisAnchor
	HeightAnchor() LayoutDimension
	LeadingAnchor() LayoutXAxisAnchor
	LeftAnchor() LayoutXAxisAnchor
	RightAnchor() LayoutXAxisAnchor
	TopAnchor() LayoutYAxisAnchor
	TrailingAnchor() LayoutXAxisAnchor
	WidthAnchor() LayoutDimension
	HasAmbiguousLayout() bool
}

type LayoutGuide struct {
	objc.Object
}

func MakeLayoutGuide(ptr unsafe.Pointer) LayoutGuide {
	return LayoutGuide{
		Object: objc.MakeObject(ptr),
	}
}

func (lc _LayoutGuideClass) Alloc() LayoutGuide {
	rv := objc.CallMethod[LayoutGuide](lc, objc.GetSelector("alloc"))
	return rv
}

func LayoutGuide_Alloc() LayoutGuide {
	return LayoutGuideClass.Alloc()
}

func (lc _LayoutGuideClass) New() LayoutGuide {
	rv := objc.CallMethod[LayoutGuide](lc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewLayoutGuide() LayoutGuide {
	return LayoutGuideClass.New()
}

func LayoutGuide_New() LayoutGuide {
	return LayoutGuideClass.New()
}

func (l_ LayoutGuide) Init() LayoutGuide {
	rv := objc.CallMethod[LayoutGuide](l_, objc.GetSelector("init"))
	return rv
}

func LayoutGuide_Init() LayoutGuide {
	return LayoutGuideClass.Alloc().Init()
}

func (l_ LayoutGuide) ConstraintsAffectingLayoutForOrientation(orientation LayoutConstraintOrientation) []LayoutConstraint {
	rv := objc.CallMethod[[]LayoutConstraint](l_, objc.GetSelector("constraintsAffectingLayoutForOrientation:"), orientation)
	return rv
}

func (l_ LayoutGuide) Identifier() UserInterfaceItemIdentifier {
	rv := objc.CallMethod[UserInterfaceItemIdentifier](l_, objc.GetSelector("identifier"))
	return rv
}

func (l_ LayoutGuide) SetIdentifier(value UserInterfaceItemIdentifier) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setIdentifier:"), value)
}

func (l_ LayoutGuide) Frame() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](l_, objc.GetSelector("frame"))
	return rv
}

func (l_ LayoutGuide) OwningView() View {
	rv := objc.CallMethod[View](l_, objc.GetSelector("owningView"))
	return rv
}

func (l_ LayoutGuide) SetOwningView(value IView) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setOwningView:"), objc.ExtractPtr(value))
}

func (l_ LayoutGuide) BottomAnchor() LayoutYAxisAnchor {
	rv := objc.CallMethod[LayoutYAxisAnchor](l_, objc.GetSelector("bottomAnchor"))
	return rv
}

func (l_ LayoutGuide) CenterXAnchor() LayoutXAxisAnchor {
	rv := objc.CallMethod[LayoutXAxisAnchor](l_, objc.GetSelector("centerXAnchor"))
	return rv
}

func (l_ LayoutGuide) CenterYAnchor() LayoutYAxisAnchor {
	rv := objc.CallMethod[LayoutYAxisAnchor](l_, objc.GetSelector("centerYAnchor"))
	return rv
}

func (l_ LayoutGuide) HeightAnchor() LayoutDimension {
	rv := objc.CallMethod[LayoutDimension](l_, objc.GetSelector("heightAnchor"))
	return rv
}

func (l_ LayoutGuide) LeadingAnchor() LayoutXAxisAnchor {
	rv := objc.CallMethod[LayoutXAxisAnchor](l_, objc.GetSelector("leadingAnchor"))
	return rv
}

func (l_ LayoutGuide) LeftAnchor() LayoutXAxisAnchor {
	rv := objc.CallMethod[LayoutXAxisAnchor](l_, objc.GetSelector("leftAnchor"))
	return rv
}

func (l_ LayoutGuide) RightAnchor() LayoutXAxisAnchor {
	rv := objc.CallMethod[LayoutXAxisAnchor](l_, objc.GetSelector("rightAnchor"))
	return rv
}

func (l_ LayoutGuide) TopAnchor() LayoutYAxisAnchor {
	rv := objc.CallMethod[LayoutYAxisAnchor](l_, objc.GetSelector("topAnchor"))
	return rv
}

func (l_ LayoutGuide) TrailingAnchor() LayoutXAxisAnchor {
	rv := objc.CallMethod[LayoutXAxisAnchor](l_, objc.GetSelector("trailingAnchor"))
	return rv
}

func (l_ LayoutGuide) WidthAnchor() LayoutDimension {
	rv := objc.CallMethod[LayoutDimension](l_, objc.GetSelector("widthAnchor"))
	return rv
}

func (l_ LayoutGuide) HasAmbiguousLayout() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("hasAmbiguousLayout"))
	return rv
}
