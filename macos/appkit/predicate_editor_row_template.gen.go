// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var PredicateEditorRowTemplateClass = _PredicateEditorRowTemplateClass{objc.GetClass("NSPredicateEditorRowTemplate")}

type _PredicateEditorRowTemplateClass struct {
	objc.Class
}

type IPredicateEditorRowTemplate interface {
	objc.IObject
	MatchForPredicate(predicate foundation.IPredicate) float64
	SetPredicate(predicate foundation.IPredicate)
	DisplayableSubpredicatesOfPredicate(predicate foundation.IPredicate) []foundation.Predicate
	PredicateWithSubpredicates(subpredicates []foundation.IPredicate) foundation.Predicate
	TemplateViews() []View
	LeftExpressions() []foundation.Expression
	RightExpressions() []foundation.Expression
	CompoundTypes() []foundation.Number
	Modifier() foundation.ComparisonPredicateModifier
	Operators() []foundation.Number
	Options() uint
}

type PredicateEditorRowTemplate struct {
	objc.Object
}

func MakePredicateEditorRowTemplate(ptr unsafe.Pointer) PredicateEditorRowTemplate {
	return PredicateEditorRowTemplate{
		Object: objc.MakeObject(ptr),
	}
}

func (p_ PredicateEditorRowTemplate) InitWithLeftExpressionsRightExpressionsModifierOperatorsOptions(leftExpressions []foundation.IExpression, rightExpressions []foundation.IExpression, modifier foundation.ComparisonPredicateModifier, operators []foundation.INumber, options uint) PredicateEditorRowTemplate {
	rv := objc.CallMethod[PredicateEditorRowTemplate](p_, objc.GetSelector("initWithLeftExpressions:rightExpressions:modifier:operators:options:"), leftExpressions, rightExpressions, modifier, operators, options)
	return rv
}

func PredicateEditorRowTemplate_InitWithLeftExpressionsRightExpressionsModifierOperatorsOptions(leftExpressions []foundation.IExpression, rightExpressions []foundation.IExpression, modifier foundation.ComparisonPredicateModifier, operators []foundation.INumber, options uint) PredicateEditorRowTemplate {
	return PredicateEditorRowTemplateClass.Alloc().InitWithLeftExpressionsRightExpressionsModifierOperatorsOptions(leftExpressions, rightExpressions, modifier, operators, options)
}

func (p_ PredicateEditorRowTemplate) InitWithCompoundTypes(compoundTypes []foundation.INumber) PredicateEditorRowTemplate {
	rv := objc.CallMethod[PredicateEditorRowTemplate](p_, objc.GetSelector("initWithCompoundTypes:"), compoundTypes)
	return rv
}

func PredicateEditorRowTemplate_InitWithCompoundTypes(compoundTypes []foundation.INumber) PredicateEditorRowTemplate {
	return PredicateEditorRowTemplateClass.Alloc().InitWithCompoundTypes(compoundTypes)
}

func (pc _PredicateEditorRowTemplateClass) Alloc() PredicateEditorRowTemplate {
	rv := objc.CallMethod[PredicateEditorRowTemplate](pc, objc.GetSelector("alloc"))
	return rv
}

func PredicateEditorRowTemplate_Alloc() PredicateEditorRowTemplate {
	return PredicateEditorRowTemplateClass.Alloc()
}

func (pc _PredicateEditorRowTemplateClass) New() PredicateEditorRowTemplate {
	rv := objc.CallMethod[PredicateEditorRowTemplate](pc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewPredicateEditorRowTemplate() PredicateEditorRowTemplate {
	return PredicateEditorRowTemplateClass.New()
}

func PredicateEditorRowTemplate_New() PredicateEditorRowTemplate {
	return PredicateEditorRowTemplateClass.New()
}

func (p_ PredicateEditorRowTemplate) Init() PredicateEditorRowTemplate {
	rv := objc.CallMethod[PredicateEditorRowTemplate](p_, objc.GetSelector("init"))
	return rv
}

func PredicateEditorRowTemplate_Init() PredicateEditorRowTemplate {
	return PredicateEditorRowTemplateClass.Alloc().Init()
}

func (p_ PredicateEditorRowTemplate) MatchForPredicate(predicate foundation.IPredicate) float64 {
	rv := objc.CallMethod[float64](p_, objc.GetSelector("matchForPredicate:"), objc.ExtractPtr(predicate))
	return rv
}

func (p_ PredicateEditorRowTemplate) SetPredicate(predicate foundation.IPredicate) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setPredicate:"), objc.ExtractPtr(predicate))
}

func (p_ PredicateEditorRowTemplate) DisplayableSubpredicatesOfPredicate(predicate foundation.IPredicate) []foundation.Predicate {
	rv := objc.CallMethod[[]foundation.Predicate](p_, objc.GetSelector("displayableSubpredicatesOfPredicate:"), objc.ExtractPtr(predicate))
	return rv
}

func (p_ PredicateEditorRowTemplate) PredicateWithSubpredicates(subpredicates []foundation.IPredicate) foundation.Predicate {
	rv := objc.CallMethod[foundation.Predicate](p_, objc.GetSelector("predicateWithSubpredicates:"), subpredicates)
	return rv
}

func (p_ PredicateEditorRowTemplate) TemplateViews() []View {
	rv := objc.CallMethod[[]View](p_, objc.GetSelector("templateViews"))
	return rv
}

func (p_ PredicateEditorRowTemplate) LeftExpressions() []foundation.Expression {
	rv := objc.CallMethod[[]foundation.Expression](p_, objc.GetSelector("leftExpressions"))
	return rv
}

func (p_ PredicateEditorRowTemplate) RightExpressions() []foundation.Expression {
	rv := objc.CallMethod[[]foundation.Expression](p_, objc.GetSelector("rightExpressions"))
	return rv
}

func (p_ PredicateEditorRowTemplate) CompoundTypes() []foundation.Number {
	rv := objc.CallMethod[[]foundation.Number](p_, objc.GetSelector("compoundTypes"))
	return rv
}

func (p_ PredicateEditorRowTemplate) Modifier() foundation.ComparisonPredicateModifier {
	rv := objc.CallMethod[foundation.ComparisonPredicateModifier](p_, objc.GetSelector("modifier"))
	return rv
}

func (p_ PredicateEditorRowTemplate) Operators() []foundation.Number {
	rv := objc.CallMethod[[]foundation.Number](p_, objc.GetSelector("operators"))
	return rv
}

func (p_ PredicateEditorRowTemplate) Options() uint {
	rv := objc.CallMethod[uint](p_, objc.GetSelector("options"))
	return rv
}
