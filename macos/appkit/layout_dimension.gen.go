// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var LayoutDimensionClass = _LayoutDimensionClass{objc.GetClass("NSLayoutDimension")}

type _LayoutDimensionClass struct {
	objc.Class
}

type ILayoutDimension interface {
	ILayoutAnchor
	ConstraintEqualToAnchorMultiplier(anchor ILayoutDimension, m float64) LayoutConstraint
	ConstraintEqualToAnchorMultiplierConstant(anchor ILayoutDimension, m float64, c float64) LayoutConstraint
	ConstraintEqualToConstant(c float64) LayoutConstraint
	ConstraintGreaterThanOrEqualToAnchorMultiplier(anchor ILayoutDimension, m float64) LayoutConstraint
	ConstraintGreaterThanOrEqualToAnchorMultiplierConstant(anchor ILayoutDimension, m float64, c float64) LayoutConstraint
	ConstraintGreaterThanOrEqualToConstant(c float64) LayoutConstraint
	ConstraintLessThanOrEqualToAnchorMultiplier(anchor ILayoutDimension, m float64) LayoutConstraint
	ConstraintLessThanOrEqualToAnchorMultiplierConstant(anchor ILayoutDimension, m float64, c float64) LayoutConstraint
	ConstraintLessThanOrEqualToConstant(c float64) LayoutConstraint
}

type LayoutDimension struct {
	LayoutAnchor
}

func MakeLayoutDimension(ptr unsafe.Pointer) LayoutDimension {
	return LayoutDimension{
		LayoutAnchor: MakeLayoutAnchor(ptr),
	}
}

func (lc _LayoutDimensionClass) Alloc() LayoutDimension {
	rv := objc.CallMethod[LayoutDimension](lc, objc.GetSelector("alloc"))
	return rv
}

func LayoutDimension_Alloc() LayoutDimension {
	return LayoutDimensionClass.Alloc()
}

func (lc _LayoutDimensionClass) New() LayoutDimension {
	rv := objc.CallMethod[LayoutDimension](lc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewLayoutDimension() LayoutDimension {
	return LayoutDimensionClass.New()
}

func LayoutDimension_New() LayoutDimension {
	return LayoutDimensionClass.New()
}

func (l_ LayoutDimension) Init() LayoutDimension {
	rv := objc.CallMethod[LayoutDimension](l_, objc.GetSelector("init"))
	return rv
}

func LayoutDimension_Init() LayoutDimension {
	return LayoutDimensionClass.Alloc().Init()
}

func (l_ LayoutDimension) ConstraintEqualToAnchorMultiplier(anchor ILayoutDimension, m float64) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.GetSelector("constraintEqualToAnchor:multiplier:"), objc.ExtractPtr(anchor), m)
	return rv
}

func (l_ LayoutDimension) ConstraintEqualToAnchorMultiplierConstant(anchor ILayoutDimension, m float64, c float64) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.GetSelector("constraintEqualToAnchor:multiplier:constant:"), objc.ExtractPtr(anchor), m, c)
	return rv
}

func (l_ LayoutDimension) ConstraintEqualToConstant(c float64) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.GetSelector("constraintEqualToConstant:"), c)
	return rv
}

func (l_ LayoutDimension) ConstraintGreaterThanOrEqualToAnchorMultiplier(anchor ILayoutDimension, m float64) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.GetSelector("constraintGreaterThanOrEqualToAnchor:multiplier:"), objc.ExtractPtr(anchor), m)
	return rv
}

func (l_ LayoutDimension) ConstraintGreaterThanOrEqualToAnchorMultiplierConstant(anchor ILayoutDimension, m float64, c float64) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.GetSelector("constraintGreaterThanOrEqualToAnchor:multiplier:constant:"), objc.ExtractPtr(anchor), m, c)
	return rv
}

func (l_ LayoutDimension) ConstraintGreaterThanOrEqualToConstant(c float64) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.GetSelector("constraintGreaterThanOrEqualToConstant:"), c)
	return rv
}

func (l_ LayoutDimension) ConstraintLessThanOrEqualToAnchorMultiplier(anchor ILayoutDimension, m float64) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.GetSelector("constraintLessThanOrEqualToAnchor:multiplier:"), objc.ExtractPtr(anchor), m)
	return rv
}

func (l_ LayoutDimension) ConstraintLessThanOrEqualToAnchorMultiplierConstant(anchor ILayoutDimension, m float64, c float64) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.GetSelector("constraintLessThanOrEqualToAnchor:multiplier:constant:"), objc.ExtractPtr(anchor), m, c)
	return rv
}

func (l_ LayoutDimension) ConstraintLessThanOrEqualToConstant(c float64) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.GetSelector("constraintLessThanOrEqualToConstant:"), c)
	return rv
}
