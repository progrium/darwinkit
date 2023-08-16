// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [LayoutConstraint] class.
var LayoutConstraintClass = _LayoutConstraintClass{objc.GetClass("NSLayoutConstraint")}

type _LayoutConstraintClass struct {
	objc.Class
}

// An interface definition for the [LayoutConstraint] class.
type ILayoutConstraint interface {
	objc.IObject
	Relation() LayoutRelation
	Priority() LayoutPriority
	SetPriority(value LayoutPriority)
	Multiplier() float64
	IsActive() bool
	SetActive(value bool)
	SecondItem() objc.Object
	FirstItem() objc.Object
	FirstAttribute() LayoutAttribute
	FirstAnchor() LayoutAnchor
	SecondAttribute() LayoutAttribute
	Constant() float64
	SetConstant(value float64)
	ShouldBeArchived() bool
	SetShouldBeArchived(value bool)
	SecondAnchor() LayoutAnchor
	Identifier() string
	SetIdentifier(value string)
}

// The relationship between two user interface objects that must be satisfied by the constraint-based layout system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutconstraint?language=objc
type LayoutConstraint struct {
	objc.Object
}

func LayoutConstraintFrom(ptr unsafe.Pointer) LayoutConstraint {
	return LayoutConstraint{
		Object: objc.ObjectFrom(ptr),
	}
}

func (lc _LayoutConstraintClass) ConstraintWithItemAttributeRelatedByToItemAttributeMultiplierConstant(view1 objc.IObject, attr1 LayoutAttribute, relation LayoutRelation, view2 objc.IObject, attr2 LayoutAttribute, multiplier float64, c float64) LayoutConstraint {
	rv := objc.Call[LayoutConstraint](lc, objc.Sel("constraintWithItem:attribute:relatedBy:toItem:attribute:multiplier:constant:"), view1, attr1, relation, view2, attr2, multiplier, c)
	return rv
}

// Creates a constraint that defines the relationship between the specified attributes of the given views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutconstraint/1526954-constraintwithitem?language=objc
func LayoutConstraint_ConstraintWithItemAttributeRelatedByToItemAttributeMultiplierConstant(view1 objc.IObject, attr1 LayoutAttribute, relation LayoutRelation, view2 objc.IObject, attr2 LayoutAttribute, multiplier float64, c float64) LayoutConstraint {
	return LayoutConstraintClass.ConstraintWithItemAttributeRelatedByToItemAttributeMultiplierConstant(view1, attr1, relation, view2, attr2, multiplier, c)
}

func (lc _LayoutConstraintClass) Alloc() LayoutConstraint {
	rv := objc.Call[LayoutConstraint](lc, objc.Sel("alloc"))
	return rv
}

func LayoutConstraint_Alloc() LayoutConstraint {
	return LayoutConstraintClass.Alloc()
}

func (lc _LayoutConstraintClass) New() LayoutConstraint {
	rv := objc.Call[LayoutConstraint](lc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewLayoutConstraint() LayoutConstraint {
	return LayoutConstraintClass.New()
}

func (l_ LayoutConstraint) Init() LayoutConstraint {
	rv := objc.Call[LayoutConstraint](l_, objc.Sel("init"))
	return rv
}

// Deactivates each constraint in the specified array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutconstraint/1526066-deactivateconstraints?language=objc
func (lc _LayoutConstraintClass) DeactivateConstraints(constraints []ILayoutConstraint) {
	objc.Call[objc.Void](lc, objc.Sel("deactivateConstraints:"), constraints)
}

// Deactivates each constraint in the specified array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutconstraint/1526066-deactivateconstraints?language=objc
func LayoutConstraint_DeactivateConstraints(constraints []ILayoutConstraint) {
	LayoutConstraintClass.DeactivateConstraints(constraints)
}

// Activates each constraint in the specified array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutconstraint/1526955-activateconstraints?language=objc
func (lc _LayoutConstraintClass) ActivateConstraints(constraints []ILayoutConstraint) {
	objc.Call[objc.Void](lc, objc.Sel("activateConstraints:"), constraints)
}

// Activates each constraint in the specified array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutconstraint/1526955-activateconstraints?language=objc
func LayoutConstraint_ActivateConstraints(constraints []ILayoutConstraint) {
	LayoutConstraintClass.ActivateConstraints(constraints)
}

// Creates constraints described by an ASCII art-like visual format string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutconstraint/1526944-constraintswithvisualformat?language=objc
func (lc _LayoutConstraintClass) ConstraintsWithVisualFormatOptionsMetricsViews(format string, opts LayoutFormatOptions, metrics map[string]objc.IObject, views map[string]objc.IObject) []LayoutConstraint {
	rv := objc.Call[[]LayoutConstraint](lc, objc.Sel("constraintsWithVisualFormat:options:metrics:views:"), format, opts, metrics, views)
	return rv
}

// Creates constraints described by an ASCII art-like visual format string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutconstraint/1526944-constraintswithvisualformat?language=objc
func LayoutConstraint_ConstraintsWithVisualFormatOptionsMetricsViews(format string, opts LayoutFormatOptions, metrics map[string]objc.IObject, views map[string]objc.IObject) []LayoutConstraint {
	return LayoutConstraintClass.ConstraintsWithVisualFormatOptionsMetricsViews(format, opts, metrics, views)
}

// The relation between the two attributes in the constraint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutconstraint/1526549-relation?language=objc
func (l_ LayoutConstraint) Relation() LayoutRelation {
	rv := objc.Call[LayoutRelation](l_, objc.Sel("relation"))
	return rv
}

// The priority of the constraint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutconstraint/1526946-priority?language=objc
func (l_ LayoutConstraint) Priority() LayoutPriority {
	rv := objc.Call[LayoutPriority](l_, objc.Sel("priority"))
	return rv
}

// The priority of the constraint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutconstraint/1526946-priority?language=objc
func (l_ LayoutConstraint) SetPriority(value LayoutPriority) {
	objc.Call[objc.Void](l_, objc.Sel("setPriority:"), value)
}

// The multiplier applied to the second attribute participating in the constraint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutconstraint/1526920-multiplier?language=objc
func (l_ LayoutConstraint) Multiplier() float64 {
	rv := objc.Call[float64](l_, objc.Sel("multiplier"))
	return rv
}

// The active state of the constraint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutconstraint/1527000-active?language=objc
func (l_ LayoutConstraint) IsActive() bool {
	rv := objc.Call[bool](l_, objc.Sel("isActive"))
	return rv
}

// The active state of the constraint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutconstraint/1527000-active?language=objc
func (l_ LayoutConstraint) SetActive(value bool) {
	objc.Call[objc.Void](l_, objc.Sel("setActive:"), value)
}

// The second object participating in the constraint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutconstraint/1526868-seconditem?language=objc
func (l_ LayoutConstraint) SecondItem() objc.Object {
	rv := objc.Call[objc.Object](l_, objc.Sel("secondItem"))
	return rv
}

// The first object participating in the constraint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutconstraint/1526860-firstitem?language=objc
func (l_ LayoutConstraint) FirstItem() objc.Object {
	rv := objc.Call[objc.Object](l_, objc.Sel("firstItem"))
	return rv
}

// The attribute of the first object participating in the constraint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutconstraint/1525204-firstattribute?language=objc
func (l_ LayoutConstraint) FirstAttribute() LayoutAttribute {
	rv := objc.Call[LayoutAttribute](l_, objc.Sel("firstAttribute"))
	return rv
}

// The first anchor that defines the constraint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutconstraint/1644261-firstanchor?language=objc
func (l_ LayoutConstraint) FirstAnchor() LayoutAnchor {
	rv := objc.Call[LayoutAnchor](l_, objc.Sel("firstAnchor"))
	return rv
}

// The attribute of the second object participating in the constraint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutconstraint/1526941-secondattribute?language=objc
func (l_ LayoutConstraint) SecondAttribute() LayoutAttribute {
	rv := objc.Call[LayoutAttribute](l_, objc.Sel("secondAttribute"))
	return rv
}

// The constant added to the multiplied second attribute participating in the constraint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutconstraint/1526928-constant?language=objc
func (l_ LayoutConstraint) Constant() float64 {
	rv := objc.Call[float64](l_, objc.Sel("constant"))
	return rv
}

// The constant added to the multiplied second attribute participating in the constraint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutconstraint/1526928-constant?language=objc
func (l_ LayoutConstraint) SetConstant(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setConstant:"), value)
}

// A Boolean value that determines whether the constraint should be archived by its owning view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutconstraint/1525647-shouldbearchived?language=objc
func (l_ LayoutConstraint) ShouldBeArchived() bool {
	rv := objc.Call[bool](l_, objc.Sel("shouldBeArchived"))
	return rv
}

// A Boolean value that determines whether the constraint should be archived by its owning view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutconstraint/1525647-shouldbearchived?language=objc
func (l_ LayoutConstraint) SetShouldBeArchived(value bool) {
	objc.Call[objc.Void](l_, objc.Sel("setShouldBeArchived:"), value)
}

// The second anchor that defines the constraint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutconstraint/1644260-secondanchor?language=objc
func (l_ LayoutConstraint) SecondAnchor() LayoutAnchor {
	rv := objc.Call[LayoutAnchor](l_, objc.Sel("secondAnchor"))
	return rv
}

// The name that identifies the constraint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutconstraint/1526879-identifier?language=objc
func (l_ LayoutConstraint) Identifier() string {
	rv := objc.Call[string](l_, objc.Sel("identifier"))
	return rv
}

// The name that identifies the constraint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutconstraint/1526879-identifier?language=objc
func (l_ LayoutConstraint) SetIdentifier(value string) {
	objc.Call[objc.Void](l_, objc.Sel("setIdentifier:"), value)
}
