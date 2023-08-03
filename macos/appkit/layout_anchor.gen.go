// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var LayoutAnchorClass = _LayoutAnchorClass{objc.GetClass("NSLayoutAnchor")}

type _LayoutAnchorClass struct {
	objc.Class
}

type ILayoutAnchor interface {
	objc.IObject
	ConstraintEqualToAnchor(anchor ILayoutAnchor) LayoutConstraint
	ConstraintEqualToAnchorConstant(anchor ILayoutAnchor, c float64) LayoutConstraint
	ConstraintGreaterThanOrEqualToAnchor(anchor ILayoutAnchor) LayoutConstraint
	ConstraintGreaterThanOrEqualToAnchorConstant(anchor ILayoutAnchor, c float64) LayoutConstraint
	ConstraintLessThanOrEqualToAnchor(anchor ILayoutAnchor) LayoutConstraint
	ConstraintLessThanOrEqualToAnchorConstant(anchor ILayoutAnchor, c float64) LayoutConstraint
	ConstraintsAffectingLayout() []LayoutConstraint
	HasAmbiguousLayout() bool
	Name() string
	Item() objc.Object
}

type LayoutAnchor struct {
	objc.Object
}

func MakeLayoutAnchor(ptr unsafe.Pointer) LayoutAnchor {
	return LayoutAnchor{
		Object: objc.MakeObject(ptr),
	}
}

func (lc _LayoutAnchorClass) Alloc() LayoutAnchor {
	rv := objc.CallMethod[LayoutAnchor](lc, objc.GetSelector("alloc"))
	return rv
}

func LayoutAnchor_Alloc() LayoutAnchor {
	return LayoutAnchorClass.Alloc()
}

func (lc _LayoutAnchorClass) New() LayoutAnchor {
	rv := objc.CallMethod[LayoutAnchor](lc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewLayoutAnchor() LayoutAnchor {
	return LayoutAnchorClass.New()
}

func LayoutAnchor_New() LayoutAnchor {
	return LayoutAnchorClass.New()
}

func (l_ LayoutAnchor) Init() LayoutAnchor {
	rv := objc.CallMethod[LayoutAnchor](l_, objc.GetSelector("init"))
	return rv
}

func LayoutAnchor_Init() LayoutAnchor {
	return LayoutAnchorClass.Alloc().Init()
}

func (l_ LayoutAnchor) ConstraintEqualToAnchor(anchor ILayoutAnchor) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.GetSelector("constraintEqualToAnchor:"), objc.ExtractPtr(anchor))
	return rv
}

func (l_ LayoutAnchor) ConstraintEqualToAnchorConstant(anchor ILayoutAnchor, c float64) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.GetSelector("constraintEqualToAnchor:constant:"), objc.ExtractPtr(anchor), c)
	return rv
}

func (l_ LayoutAnchor) ConstraintGreaterThanOrEqualToAnchor(anchor ILayoutAnchor) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.GetSelector("constraintGreaterThanOrEqualToAnchor:"), objc.ExtractPtr(anchor))
	return rv
}

func (l_ LayoutAnchor) ConstraintGreaterThanOrEqualToAnchorConstant(anchor ILayoutAnchor, c float64) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.GetSelector("constraintGreaterThanOrEqualToAnchor:constant:"), objc.ExtractPtr(anchor), c)
	return rv
}

func (l_ LayoutAnchor) ConstraintLessThanOrEqualToAnchor(anchor ILayoutAnchor) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.GetSelector("constraintLessThanOrEqualToAnchor:"), objc.ExtractPtr(anchor))
	return rv
}

func (l_ LayoutAnchor) ConstraintLessThanOrEqualToAnchorConstant(anchor ILayoutAnchor, c float64) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.GetSelector("constraintLessThanOrEqualToAnchor:constant:"), objc.ExtractPtr(anchor), c)
	return rv
}

func (l_ LayoutAnchor) ConstraintsAffectingLayout() []LayoutConstraint {
	rv := objc.CallMethod[[]LayoutConstraint](l_, objc.GetSelector("constraintsAffectingLayout"))
	// TODO: convert slice items...
	return rv
}

func (l_ LayoutAnchor) HasAmbiguousLayout() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("hasAmbiguousLayout"))
	return rv
}

func (l_ LayoutAnchor) Name() string {
	rv := objc.CallMethod[string](l_, objc.GetSelector("name"))
	return rv
}

func (l_ LayoutAnchor) Item() objc.Object {
	rv := objc.CallMethod[objc.Object](l_, objc.GetSelector("item"))
	return rv
}
