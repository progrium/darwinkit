// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var LayoutConstraintClass = _LayoutConstraintClass{objc.GetClass("NSLayoutConstraint")}

type _LayoutConstraintClass struct {
	objc.Class
}

type ILayoutConstraint interface {
	objc.IObject
	IsActive() bool
	SetActive(value bool)
	FirstItem() objc.Object
	FirstAttribute() LayoutAttribute
	Relation() LayoutRelation
	SecondItem() objc.Object
	SecondAttribute() LayoutAttribute
	Multiplier() float64
	Constant() float64
	SetConstant(value float64)
	FirstAnchor() LayoutAnchor
	SecondAnchor() LayoutAnchor
	Priority() LayoutPriority
	SetPriority(value LayoutPriority)
	Identifier() string
	SetIdentifier(value string)
	ShouldBeArchived() bool
	SetShouldBeArchived(value bool)
}

type LayoutConstraint struct {
	objc.Object
}

func MakeLayoutConstraint(ptr unsafe.Pointer) LayoutConstraint {
	return LayoutConstraint{
		Object: objc.MakeObject(ptr),
	}
}

func (lc _LayoutConstraintClass) ConstraintWithItemAttributeRelatedByToItemAttributeMultiplierConstant(view1 objc.IObject, attr1 LayoutAttribute, relation LayoutRelation, view2 objc.IObject, attr2 LayoutAttribute, multiplier float64, c float64) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](lc, objc.GetSelector("constraintWithItem:attribute:relatedBy:toItem:attribute:multiplier:constant:"), objc.ExtractPtr(view1), attr1, relation, objc.ExtractPtr(view2), attr2, multiplier, c)
	return rv
}

func LayoutConstraint_ConstraintWithItemAttributeRelatedByToItemAttributeMultiplierConstant(view1 objc.IObject, attr1 LayoutAttribute, relation LayoutRelation, view2 objc.IObject, attr2 LayoutAttribute, multiplier float64, c float64) LayoutConstraint {
	return LayoutConstraintClass.ConstraintWithItemAttributeRelatedByToItemAttributeMultiplierConstant(view1, attr1, relation, view2, attr2, multiplier, c)
}

func (lc _LayoutConstraintClass) Alloc() LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](lc, objc.GetSelector("alloc"))
	return rv
}

func LayoutConstraint_Alloc() LayoutConstraint {
	return LayoutConstraintClass.Alloc()
}

func (lc _LayoutConstraintClass) New() LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](lc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewLayoutConstraint() LayoutConstraint {
	return LayoutConstraintClass.New()
}

func LayoutConstraint_New() LayoutConstraint {
	return LayoutConstraintClass.New()
}

func (l_ LayoutConstraint) Init() LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.GetSelector("init"))
	return rv
}

func LayoutConstraint_Init() LayoutConstraint {
	return LayoutConstraintClass.Alloc().Init()
}

func (lc _LayoutConstraintClass) ConstraintsWithVisualFormatOptionsMetricsViews(format string, opts LayoutFormatOptions, metrics map[string]objc.IObject, views map[string]objc.IObject) []LayoutConstraint {
	rv := objc.CallMethod[[]LayoutConstraint](lc, objc.GetSelector("constraintsWithVisualFormat:options:metrics:views:"), format, opts, metrics, views)
	// TODO: convert slice items...
	return rv
}

func LayoutConstraint_ConstraintsWithVisualFormatOptionsMetricsViews(format string, opts LayoutFormatOptions, metrics map[string]objc.IObject, views map[string]objc.IObject) []LayoutConstraint {
	return LayoutConstraintClass.ConstraintsWithVisualFormatOptionsMetricsViews(format, opts, metrics, views)
}

func (lc _LayoutConstraintClass) ActivateConstraints(constraints []ILayoutConstraint) {
	objc.CallMethod[objc.Void](lc, objc.GetSelector("activateConstraints:"), constraints)
}

func LayoutConstraint_ActivateConstraints(constraints []ILayoutConstraint) {
	LayoutConstraintClass.ActivateConstraints(constraints)
}

func (lc _LayoutConstraintClass) DeactivateConstraints(constraints []ILayoutConstraint) {
	objc.CallMethod[objc.Void](lc, objc.GetSelector("deactivateConstraints:"), constraints)
}

func LayoutConstraint_DeactivateConstraints(constraints []ILayoutConstraint) {
	LayoutConstraintClass.DeactivateConstraints(constraints)
}

func (l_ LayoutConstraint) IsActive() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("isActive"))
	return rv
}

func (l_ LayoutConstraint) SetActive(value bool) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setActive:"), value)
}

func (l_ LayoutConstraint) FirstItem() objc.Object {
	rv := objc.CallMethod[objc.Object](l_, objc.GetSelector("firstItem"))
	return rv
}

func (l_ LayoutConstraint) FirstAttribute() LayoutAttribute {
	rv := objc.CallMethod[LayoutAttribute](l_, objc.GetSelector("firstAttribute"))
	return rv
}

func (l_ LayoutConstraint) Relation() LayoutRelation {
	rv := objc.CallMethod[LayoutRelation](l_, objc.GetSelector("relation"))
	return rv
}

func (l_ LayoutConstraint) SecondItem() objc.Object {
	rv := objc.CallMethod[objc.Object](l_, objc.GetSelector("secondItem"))
	return rv
}

func (l_ LayoutConstraint) SecondAttribute() LayoutAttribute {
	rv := objc.CallMethod[LayoutAttribute](l_, objc.GetSelector("secondAttribute"))
	return rv
}

func (l_ LayoutConstraint) Multiplier() float64 {
	rv := objc.CallMethod[float64](l_, objc.GetSelector("multiplier"))
	return rv
}

func (l_ LayoutConstraint) Constant() float64 {
	rv := objc.CallMethod[float64](l_, objc.GetSelector("constant"))
	return rv
}

func (l_ LayoutConstraint) SetConstant(value float64) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setConstant:"), value)
}

func (l_ LayoutConstraint) FirstAnchor() LayoutAnchor {
	rv := objc.CallMethod[LayoutAnchor](l_, objc.GetSelector("firstAnchor"))
	return rv
}

func (l_ LayoutConstraint) SecondAnchor() LayoutAnchor {
	rv := objc.CallMethod[LayoutAnchor](l_, objc.GetSelector("secondAnchor"))
	return rv
}

func (l_ LayoutConstraint) Priority() LayoutPriority {
	rv := objc.CallMethod[LayoutPriority](l_, objc.GetSelector("priority"))
	return rv
}

func (l_ LayoutConstraint) SetPriority(value LayoutPriority) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setPriority:"), value)
}

func (l_ LayoutConstraint) Identifier() string {
	rv := objc.CallMethod[string](l_, objc.GetSelector("identifier"))
	return rv
}

func (l_ LayoutConstraint) SetIdentifier(value string) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setIdentifier:"), value)
}

func (l_ LayoutConstraint) ShouldBeArchived() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("shouldBeArchived"))
	return rv
}

func (l_ LayoutConstraint) SetShouldBeArchived(value bool) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setShouldBeArchived:"), value)
}
