// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [LayoutXAxisAnchor] class.
var LayoutXAxisAnchorClass = _LayoutXAxisAnchorClass{objc.GetClass("NSLayoutXAxisAnchor")}

type _LayoutXAxisAnchorClass struct {
	objc.Class
}

// An interface definition for the [LayoutXAxisAnchor] class.
type ILayoutXAxisAnchor interface {
	ILayoutAnchor
	ConstraintLessThanOrEqualToSystemSpacingAfterAnchorMultiplier(anchor ILayoutXAxisAnchor, multiplier float64) LayoutConstraint
	ConstraintGreaterThanOrEqualToSystemSpacingAfterAnchorMultiplier(anchor ILayoutXAxisAnchor, multiplier float64) LayoutConstraint
	ConstraintEqualToSystemSpacingAfterAnchorMultiplier(anchor ILayoutXAxisAnchor, multiplier float64) LayoutConstraint
	AnchorWithOffsetToAnchor(otherAnchor ILayoutXAxisAnchor) LayoutDimension
}

// A factory class for creating horizontal layout constraint objects using a fluent API. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutxaxisanchor?language=objc
type LayoutXAxisAnchor struct {
	LayoutAnchor
}

func LayoutXAxisAnchorFrom(ptr unsafe.Pointer) LayoutXAxisAnchor {
	return LayoutXAxisAnchor{
		LayoutAnchor: LayoutAnchorFrom(ptr),
	}
}

func (lc _LayoutXAxisAnchorClass) Alloc() LayoutXAxisAnchor {
	rv := objc.Call[LayoutXAxisAnchor](lc, objc.Sel("alloc"))
	return rv
}

func LayoutXAxisAnchor_Alloc() LayoutXAxisAnchor {
	return LayoutXAxisAnchorClass.Alloc()
}

func (lc _LayoutXAxisAnchorClass) New() LayoutXAxisAnchor {
	rv := objc.Call[LayoutXAxisAnchor](lc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewLayoutXAxisAnchor() LayoutXAxisAnchor {
	return LayoutXAxisAnchorClass.New()
}

func (l_ LayoutXAxisAnchor) Init() LayoutXAxisAnchor {
	rv := objc.Call[LayoutXAxisAnchor](l_, objc.Sel("init"))
	return rv
}

// Returns a constraint that defines the maximum amount by which the current anchor trails the specified anchor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutxaxisanchor/2866018-constraintlessthanorequaltosyste?language=objc
func (l_ LayoutXAxisAnchor) ConstraintLessThanOrEqualToSystemSpacingAfterAnchorMultiplier(anchor ILayoutXAxisAnchor, multiplier float64) LayoutConstraint {
	rv := objc.Call[LayoutConstraint](l_, objc.Sel("constraintLessThanOrEqualToSystemSpacingAfterAnchor:multiplier:"), objc.Ptr(anchor), multiplier)
	return rv
}

// Returns a constraint that defines the minimum amount by which the current anchor trails the specified anchor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutxaxisanchor/2865871-constraintgreaterthanorequaltosy?language=objc
func (l_ LayoutXAxisAnchor) ConstraintGreaterThanOrEqualToSystemSpacingAfterAnchorMultiplier(anchor ILayoutXAxisAnchor, multiplier float64) LayoutConstraint {
	rv := objc.Call[LayoutConstraint](l_, objc.Sel("constraintGreaterThanOrEqualToSystemSpacingAfterAnchor:multiplier:"), objc.Ptr(anchor), multiplier)
	return rv
}

// Returns a constraint that defines by how much the current anchor trails the specified anchor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutxaxisanchor/2866112-constraintequaltosystemspacingaf?language=objc
func (l_ LayoutXAxisAnchor) ConstraintEqualToSystemSpacingAfterAnchorMultiplier(anchor ILayoutXAxisAnchor, multiplier float64) LayoutConstraint {
	rv := objc.Call[LayoutConstraint](l_, objc.Sel("constraintEqualToSystemSpacingAfterAnchor:multiplier:"), objc.Ptr(anchor), multiplier)
	return rv
}

// Creates a layout dimension object from two anchors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutxaxisanchor/2866024-anchorwithoffsettoanchor?language=objc
func (l_ LayoutXAxisAnchor) AnchorWithOffsetToAnchor(otherAnchor ILayoutXAxisAnchor) LayoutDimension {
	rv := objc.Call[LayoutDimension](l_, objc.Sel("anchorWithOffsetToAnchor:"), objc.Ptr(otherAnchor))
	return rv
}
