// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var LayoutXAxisAnchorClass = _LayoutXAxisAnchorClass{objc.GetClass("NSLayoutXAxisAnchor")}

type _LayoutXAxisAnchorClass struct {
	objc.Class
}

type ILayoutXAxisAnchor interface {
	ILayoutAnchor
	ConstraintEqualToSystemSpacingAfterAnchorMultiplier(anchor ILayoutXAxisAnchor, multiplier float64) LayoutConstraint
	ConstraintGreaterThanOrEqualToSystemSpacingAfterAnchorMultiplier(anchor ILayoutXAxisAnchor, multiplier float64) LayoutConstraint
	ConstraintLessThanOrEqualToSystemSpacingAfterAnchorMultiplier(anchor ILayoutXAxisAnchor, multiplier float64) LayoutConstraint
	AnchorWithOffsetToAnchor(otherAnchor ILayoutXAxisAnchor) LayoutDimension
}

type LayoutXAxisAnchor struct {
	LayoutAnchor
}

func MakeLayoutXAxisAnchor(ptr unsafe.Pointer) LayoutXAxisAnchor {
	return LayoutXAxisAnchor{
		LayoutAnchor: MakeLayoutAnchor(ptr),
	}
}

func (lc _LayoutXAxisAnchorClass) Alloc() LayoutXAxisAnchor {
	rv := objc.CallMethod[LayoutXAxisAnchor](lc, objc.GetSelector("alloc"))
	return rv
}

func LayoutXAxisAnchor_Alloc() LayoutXAxisAnchor {
	return LayoutXAxisAnchorClass.Alloc()
}

func (lc _LayoutXAxisAnchorClass) New() LayoutXAxisAnchor {
	rv := objc.CallMethod[LayoutXAxisAnchor](lc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewLayoutXAxisAnchor() LayoutXAxisAnchor {
	return LayoutXAxisAnchorClass.New()
}

func LayoutXAxisAnchor_New() LayoutXAxisAnchor {
	return LayoutXAxisAnchorClass.New()
}

func (l_ LayoutXAxisAnchor) Init() LayoutXAxisAnchor {
	rv := objc.CallMethod[LayoutXAxisAnchor](l_, objc.GetSelector("init"))
	return rv
}

func LayoutXAxisAnchor_Init() LayoutXAxisAnchor {
	return LayoutXAxisAnchorClass.Alloc().Init()
}

func (l_ LayoutXAxisAnchor) ConstraintEqualToSystemSpacingAfterAnchorMultiplier(anchor ILayoutXAxisAnchor, multiplier float64) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.GetSelector("constraintEqualToSystemSpacingAfterAnchor:multiplier:"), objc.ExtractPtr(anchor), multiplier)
	return rv
}

func (l_ LayoutXAxisAnchor) ConstraintGreaterThanOrEqualToSystemSpacingAfterAnchorMultiplier(anchor ILayoutXAxisAnchor, multiplier float64) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.GetSelector("constraintGreaterThanOrEqualToSystemSpacingAfterAnchor:multiplier:"), objc.ExtractPtr(anchor), multiplier)
	return rv
}

func (l_ LayoutXAxisAnchor) ConstraintLessThanOrEqualToSystemSpacingAfterAnchorMultiplier(anchor ILayoutXAxisAnchor, multiplier float64) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.GetSelector("constraintLessThanOrEqualToSystemSpacingAfterAnchor:multiplier:"), objc.ExtractPtr(anchor), multiplier)
	return rv
}

func (l_ LayoutXAxisAnchor) AnchorWithOffsetToAnchor(otherAnchor ILayoutXAxisAnchor) LayoutDimension {
	rv := objc.CallMethod[LayoutDimension](l_, objc.GetSelector("anchorWithOffsetToAnchor:"), objc.ExtractPtr(otherAnchor))
	return rv
}
