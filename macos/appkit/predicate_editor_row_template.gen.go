// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coredata"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PredicateEditorRowTemplate] class.
var PredicateEditorRowTemplateClass = _PredicateEditorRowTemplateClass{objc.GetClass("NSPredicateEditorRowTemplate")}

type _PredicateEditorRowTemplateClass struct {
	objc.Class
}

// An interface definition for the [PredicateEditorRowTemplate] class.
type IPredicateEditorRowTemplate interface {
	objc.IObject
	SetPredicate(predicate foundation.IPredicate)
	DisplayableSubpredicatesOfPredicate(predicate foundation.IPredicate) []foundation.Predicate
	MatchForPredicate(predicate foundation.IPredicate) float64
	PredicateWithSubpredicates(subpredicates []foundation.IPredicate) foundation.Predicate
	TemplateViews() []View
	Options() uint
	Operators() []foundation.Number
	RightExpressions() []foundation.Expression
	RightExpressionAttributeType() coredata.AttributeType
	CompoundTypes() []foundation.Number
	Modifier() foundation.ComparisonPredicateModifier
	LeftExpressions() []foundation.Expression
}

// A template that describes available predicates and how to display them. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditorrowtemplate?language=objc
type PredicateEditorRowTemplate struct {
	objc.Object
}

func PredicateEditorRowTemplateFrom(ptr unsafe.Pointer) PredicateEditorRowTemplate {
	return PredicateEditorRowTemplate{
		Object: objc.ObjectFrom(ptr),
	}
}

func (p_ PredicateEditorRowTemplate) InitWithLeftExpressionsRightExpressionsModifierOperatorsOptions(leftExpressions []foundation.IExpression, rightExpressions []foundation.IExpression, modifier foundation.ComparisonPredicateModifier, operators []foundation.INumber, options uint) PredicateEditorRowTemplate {
	rv := objc.Call[PredicateEditorRowTemplate](p_, objc.Sel("initWithLeftExpressions:rightExpressions:modifier:operators:options:"), leftExpressions, rightExpressions, modifier, operators, options)
	return rv
}

// Initializes and returns a “pop-up-pop-up-pop-up”–style row template. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditorrowtemplate/1401175-initwithleftexpressions?language=objc
func PredicateEditorRowTemplate_InitWithLeftExpressionsRightExpressionsModifierOperatorsOptions(leftExpressions []foundation.IExpression, rightExpressions []foundation.IExpression, modifier foundation.ComparisonPredicateModifier, operators []foundation.INumber, options uint) PredicateEditorRowTemplate {
	return PredicateEditorRowTemplateClass.Alloc().InitWithLeftExpressionsRightExpressionsModifierOperatorsOptions(leftExpressions, rightExpressions, modifier, operators, options)
}

func (p_ PredicateEditorRowTemplate) InitWithCompoundTypes(compoundTypes []foundation.INumber) PredicateEditorRowTemplate {
	rv := objc.Call[PredicateEditorRowTemplate](p_, objc.Sel("initWithCompoundTypes:"), compoundTypes)
	return rv
}

// Initializes and returns a row template suitable for displaying compound predicates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditorrowtemplate/1401177-initwithcompoundtypes?language=objc
func PredicateEditorRowTemplate_InitWithCompoundTypes(compoundTypes []foundation.INumber) PredicateEditorRowTemplate {
	return PredicateEditorRowTemplateClass.Alloc().InitWithCompoundTypes(compoundTypes)
}

func (pc _PredicateEditorRowTemplateClass) Alloc() PredicateEditorRowTemplate {
	rv := objc.Call[PredicateEditorRowTemplate](pc, objc.Sel("alloc"))
	return rv
}

func PredicateEditorRowTemplate_Alloc() PredicateEditorRowTemplate {
	return PredicateEditorRowTemplateClass.Alloc()
}

func (pc _PredicateEditorRowTemplateClass) New() PredicateEditorRowTemplate {
	rv := objc.Call[PredicateEditorRowTemplate](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPredicateEditorRowTemplate() PredicateEditorRowTemplate {
	return PredicateEditorRowTemplateClass.New()
}

func (p_ PredicateEditorRowTemplate) Init() PredicateEditorRowTemplate {
	rv := objc.Call[PredicateEditorRowTemplate](p_, objc.Sel("init"))
	return rv
}

// Returns an array of predicate templates for the given attribute key paths for a given entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditorrowtemplate/1401187-templateswithattributekeypaths?language=objc
func (pc _PredicateEditorRowTemplateClass) TemplatesWithAttributeKeyPathsInEntityDescription(keyPaths []string, entityDescription coredata.IEntityDescription) []PredicateEditorRowTemplate {
	rv := objc.Call[[]PredicateEditorRowTemplate](pc, objc.Sel("templatesWithAttributeKeyPaths:inEntityDescription:"), keyPaths, objc.Ptr(entityDescription))
	return rv
}

// Returns an array of predicate templates for the given attribute key paths for a given entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditorrowtemplate/1401187-templateswithattributekeypaths?language=objc
func PredicateEditorRowTemplate_TemplatesWithAttributeKeyPathsInEntityDescription(keyPaths []string, entityDescription coredata.IEntityDescription) []PredicateEditorRowTemplate {
	return PredicateEditorRowTemplateClass.TemplatesWithAttributeKeyPathsInEntityDescription(keyPaths, entityDescription)
}

// Sets the value of the views according to the given predicate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditorrowtemplate/1401189-setpredicate?language=objc
func (p_ PredicateEditorRowTemplate) SetPredicate(predicate foundation.IPredicate) {
	objc.Call[objc.Void](p_, objc.Sel("setPredicate:"), objc.Ptr(predicate))
}

// Returns the subpredicates that should be made sub-rows of a given predicate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditorrowtemplate/1401183-displayablesubpredicatesofpredic?language=objc
func (p_ PredicateEditorRowTemplate) DisplayableSubpredicatesOfPredicate(predicate foundation.IPredicate) []foundation.Predicate {
	rv := objc.Call[[]foundation.Predicate](p_, objc.Sel("displayableSubpredicatesOfPredicate:"), objc.Ptr(predicate))
	return rv
}

// Returns a positive number if the receiver can represent a given predicate, and 0 if it cannot. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditorrowtemplate/1401185-matchforpredicate?language=objc
func (p_ PredicateEditorRowTemplate) MatchForPredicate(predicate foundation.IPredicate) float64 {
	rv := objc.Call[float64](p_, objc.Sel("matchForPredicate:"), objc.Ptr(predicate))
	return rv
}

// Returns the predicate represented by the receiver’s views' values and the given sub-predicates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditorrowtemplate/1401179-predicatewithsubpredicates?language=objc
func (p_ PredicateEditorRowTemplate) PredicateWithSubpredicates(subpredicates []foundation.IPredicate) foundation.Predicate {
	rv := objc.Call[foundation.Predicate](p_, objc.Sel("predicateWithSubpredicates:"), subpredicates)
	return rv
}

// Returns the views that display this template’s predicate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditorrowtemplate/1401193-templateviews?language=objc
func (p_ PredicateEditorRowTemplate) TemplateViews() []View {
	rv := objc.Call[[]View](p_, objc.Sel("templateViews"))
	return rv
}

// Returns the comparison predicate options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditorrowtemplate/1401197-options?language=objc
func (p_ PredicateEditorRowTemplate) Options() uint {
	rv := objc.Call[uint](p_, objc.Sel("options"))
	return rv
}

// Returns the array of comparison predicate operators. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditorrowtemplate/1401195-operators?language=objc
func (p_ PredicateEditorRowTemplate) Operators() []foundation.Number {
	rv := objc.Call[[]foundation.Number](p_, objc.Sel("operators"))
	return rv
}

// Returns the right hand expressions for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditorrowtemplate/1401173-rightexpressions?language=objc
func (p_ PredicateEditorRowTemplate) RightExpressions() []foundation.Expression {
	rv := objc.Call[[]foundation.Expression](p_, objc.Sel("rightExpressions"))
	return rv
}

// Returns the attribute type of the receiver’s right expression. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditorrowtemplate/1401206-rightexpressionattributetype?language=objc
func (p_ PredicateEditorRowTemplate) RightExpressionAttributeType() coredata.AttributeType {
	rv := objc.Call[coredata.AttributeType](p_, objc.Sel("rightExpressionAttributeType"))
	return rv
}

// Returns the compound predicate types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditorrowtemplate/1401202-compoundtypes?language=objc
func (p_ PredicateEditorRowTemplate) CompoundTypes() []foundation.Number {
	rv := objc.Call[[]foundation.Number](p_, objc.Sel("compoundTypes"))
	return rv
}

// Returns the comparison predicate modifier for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditorrowtemplate/1401204-modifier?language=objc
func (p_ PredicateEditorRowTemplate) Modifier() foundation.ComparisonPredicateModifier {
	rv := objc.Call[foundation.ComparisonPredicateModifier](p_, objc.Sel("modifier"))
	return rv
}

// Returns the left hand expressions for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditorrowtemplate/1401191-leftexpressions?language=objc
func (p_ PredicateEditorRowTemplate) LeftExpressions() []foundation.Expression {
	rv := objc.Call[[]foundation.Expression](p_, objc.Sel("leftExpressions"))
	return rv
}
