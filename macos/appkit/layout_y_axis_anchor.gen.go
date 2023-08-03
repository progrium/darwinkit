// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var LayoutYAxisAnchorClass = _LayoutYAxisAnchorClass{objc.GetClass("NSLayoutYAxisAnchor")}

type _LayoutYAxisAnchorClass struct {
	objc.Class
}

type ILayoutYAxisAnchor interface {
	ILayoutAnchor
	ConstraintEqualToSystemSpacingBelowAnchorMultiplier(anchor ILayoutYAxisAnchor, multiplier float64) LayoutConstraint
	ConstraintGreaterThanOrEqualToSystemSpacingBelowAnchorMultiplier(anchor ILayoutYAxisAnchor, multiplier float64) LayoutConstraint
	ConstraintLessThanOrEqualToSystemSpacingBelowAnchorMultiplier(anchor ILayoutYAxisAnchor, multiplier float64) LayoutConstraint
	AnchorWithOffsetToAnchor(otherAnchor ILayoutYAxisAnchor) LayoutDimension
}

type LayoutYAxisAnchor struct {
	LayoutAnchor
}

func MakeLayoutYAxisAnchor(ptr unsafe.Pointer) LayoutYAxisAnchor {
	return LayoutYAxisAnchor{
		LayoutAnchor: MakeLayoutAnchor(ptr),
	}
}

func (lc _LayoutYAxisAnchorClass) Alloc() LayoutYAxisAnchor {
	rv := objc.CallMethod[LayoutYAxisAnchor](lc, objc.GetSelector("alloc"))
	return rv
}

func LayoutYAxisAnchor_Alloc() LayoutYAxisAnchor {
	return LayoutYAxisAnchorClass.Alloc()
}

func (lc _LayoutYAxisAnchorClass) New() LayoutYAxisAnchor {
	rv := objc.CallMethod[LayoutYAxisAnchor](lc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewLayoutYAxisAnchor() LayoutYAxisAnchor {
	return LayoutYAxisAnchorClass.New()
}

func LayoutYAxisAnchor_New() LayoutYAxisAnchor {
	return LayoutYAxisAnchorClass.New()
}

func (l_ LayoutYAxisAnchor) Init() LayoutYAxisAnchor {
	rv := objc.CallMethod[LayoutYAxisAnchor](l_, objc.GetSelector("init"))
	return rv
}

func LayoutYAxisAnchor_Init() LayoutYAxisAnchor {
	return LayoutYAxisAnchorClass.Alloc().Init()
}

func (l_ LayoutYAxisAnchor) ConstraintEqualToSystemSpacingBelowAnchorMultiplier(anchor ILayoutYAxisAnchor, multiplier float64) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.GetSelector("constraintEqualToSystemSpacingBelowAnchor:multiplier:"), objc.ExtractPtr(anchor), multiplier)
	return rv
}

func (l_ LayoutYAxisAnchor) ConstraintGreaterThanOrEqualToSystemSpacingBelowAnchorMultiplier(anchor ILayoutYAxisAnchor, multiplier float64) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.GetSelector("constraintGreaterThanOrEqualToSystemSpacingBelowAnchor:multiplier:"), objc.ExtractPtr(anchor), multiplier)
	return rv
}

func (l_ LayoutYAxisAnchor) ConstraintLessThanOrEqualToSystemSpacingBelowAnchorMultiplier(anchor ILayoutYAxisAnchor, multiplier float64) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.GetSelector("constraintLessThanOrEqualToSystemSpacingBelowAnchor:multiplier:"), objc.ExtractPtr(anchor), multiplier)
	return rv
}

func (l_ LayoutYAxisAnchor) AnchorWithOffsetToAnchor(otherAnchor ILayoutYAxisAnchor) LayoutDimension {
	rv := objc.CallMethod[LayoutDimension](l_, objc.GetSelector("anchorWithOffsetToAnchor:"), objc.ExtractPtr(otherAnchor))
	return rv
}
