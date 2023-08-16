// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [LayoutDimension] class.
var LayoutDimensionClass = _LayoutDimensionClass{objc.GetClass("NSLayoutDimension")}

type _LayoutDimensionClass struct {
	objc.Class
}

// An interface definition for the [LayoutDimension] class.
type ILayoutDimension interface {
	ILayoutAnchor
	ConstraintLessThanOrEqualToAnchorMultiplier(anchor ILayoutDimension, m float64) LayoutConstraint
	ConstraintEqualToAnchorMultiplierConstant(anchor ILayoutDimension, m float64, c float64) LayoutConstraint
	ConstraintEqualToConstant(c float64) LayoutConstraint
	ConstraintLessThanOrEqualToConstant(c float64) LayoutConstraint
	ConstraintGreaterThanOrEqualToConstant(c float64) LayoutConstraint
	ConstraintGreaterThanOrEqualToAnchorMultiplier(anchor ILayoutDimension, m float64) LayoutConstraint
}

// A factory class for creating size-based layout constraint objects using a fluent API. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutdimension?language=objc
type LayoutDimension struct {
	LayoutAnchor
}

func LayoutDimensionFrom(ptr unsafe.Pointer) LayoutDimension {
	return LayoutDimension{
		LayoutAnchor: LayoutAnchorFrom(ptr),
	}
}

func (lc _LayoutDimensionClass) Alloc() LayoutDimension {
	rv := objc.Call[LayoutDimension](lc, objc.Sel("alloc"))
	return rv
}

func LayoutDimension_Alloc() LayoutDimension {
	return LayoutDimensionClass.Alloc()
}

func (lc _LayoutDimensionClass) New() LayoutDimension {
	rv := objc.Call[LayoutDimension](lc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewLayoutDimension() LayoutDimension {
	return LayoutDimensionClass.New()
}

func (l_ LayoutDimension) Init() LayoutDimension {
	rv := objc.Call[LayoutDimension](l_, objc.Sel("init"))
	return rv
}

// Returns a constraint that defines the anchor’s size attribute as less than or equal to the specified anchor multiplied by the constant. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutdimension/1500943-constraintlessthanorequaltoancho?language=objc
func (l_ LayoutDimension) ConstraintLessThanOrEqualToAnchorMultiplier(anchor ILayoutDimension, m float64) LayoutConstraint {
	rv := objc.Call[LayoutConstraint](l_, objc.Sel("constraintLessThanOrEqualToAnchor:multiplier:"), objc.Ptr(anchor), m)
	return rv
}

// Returns a constraint that defines the anchor’s size attribute as equal to the specified size attribute multiplied by a constant plus an offset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutdimension/1500934-constraintequaltoanchor?language=objc
func (l_ LayoutDimension) ConstraintEqualToAnchorMultiplierConstant(anchor ILayoutDimension, m float64, c float64) LayoutConstraint {
	rv := objc.Call[LayoutConstraint](l_, objc.Sel("constraintEqualToAnchor:multiplier:constant:"), objc.Ptr(anchor), m, c)
	return rv
}

// Returns a constraint that defines a constant size for the anchor’s size attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutdimension/1500941-constraintequaltoconstant?language=objc
func (l_ LayoutDimension) ConstraintEqualToConstant(c float64) LayoutConstraint {
	rv := objc.Call[LayoutConstraint](l_, objc.Sel("constraintEqualToConstant:"), c)
	return rv
}

// Returns a constraint that defines the maximum size for the anchor’s size attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutdimension/1500963-constraintlessthanorequaltoconst?language=objc
func (l_ LayoutDimension) ConstraintLessThanOrEqualToConstant(c float64) LayoutConstraint {
	rv := objc.Call[LayoutConstraint](l_, objc.Sel("constraintLessThanOrEqualToConstant:"), c)
	return rv
}

// Returns a constraint that defines the minimum size for the anchor’s size attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutdimension/1500939-constraintgreaterthanorequaltoco?language=objc
func (l_ LayoutDimension) ConstraintGreaterThanOrEqualToConstant(c float64) LayoutConstraint {
	rv := objc.Call[LayoutConstraint](l_, objc.Sel("constraintGreaterThanOrEqualToConstant:"), c)
	return rv
}

// Returns a constraint that defines the anchor’s size attribute as greater than or equal to the specified anchor multiplied by the constant. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutdimension/1500961-constraintgreaterthanorequaltoan?language=objc
func (l_ LayoutDimension) ConstraintGreaterThanOrEqualToAnchorMultiplier(anchor ILayoutDimension, m float64) LayoutConstraint {
	rv := objc.Call[LayoutConstraint](l_, objc.Sel("constraintGreaterThanOrEqualToAnchor:multiplier:"), objc.Ptr(anchor), m)
	return rv
}
