// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ComparisonPredicate] class.
var ComparisonPredicateClass = _ComparisonPredicateClass{objc.GetClass("NSComparisonPredicate")}

type _ComparisonPredicateClass struct {
	objc.Class
}

// An interface definition for the [ComparisonPredicate] class.
type IComparisonPredicate interface {
	IPredicate
	Options() ComparisonPredicateOptions
	CustomSelector() objc.Selector
	ComparisonPredicateModifier() ComparisonPredicateModifier
	RightExpression() Expression
	LeftExpression() Expression
	PredicateOperatorType() PredicateOperatorType
}

// A specialized predicate for comparing expressions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscomparisonpredicate?language=objc
type ComparisonPredicate struct {
	Predicate
}

func ComparisonPredicateFrom(ptr unsafe.Pointer) ComparisonPredicate {
	return ComparisonPredicate{
		Predicate: PredicateFrom(ptr),
	}
}

func (c_ ComparisonPredicate) InitWithLeftExpressionRightExpressionCustomSelector(lhs IExpression, rhs IExpression, selector objc.Selector) ComparisonPredicate {
	rv := objc.Call[ComparisonPredicate](c_, objc.Sel("initWithLeftExpression:rightExpression:customSelector:"), objc.Ptr(lhs), objc.Ptr(rhs), selector)
	return rv
}

// Creates a predicate that you form by combining specified left and right expressions using a specified selector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscomparisonpredicate/1409054-initwithleftexpression?language=objc
func ComparisonPredicate_InitWithLeftExpressionRightExpressionCustomSelector(lhs IExpression, rhs IExpression, selector objc.Selector) ComparisonPredicate {
	return ComparisonPredicateClass.Alloc().InitWithLeftExpressionRightExpressionCustomSelector(lhs, rhs, selector)
}

func (cc _ComparisonPredicateClass) Alloc() ComparisonPredicate {
	rv := objc.Call[ComparisonPredicate](cc, objc.Sel("alloc"))
	return rv
}

func ComparisonPredicate_Alloc() ComparisonPredicate {
	return ComparisonPredicateClass.Alloc()
}

func (cc _ComparisonPredicateClass) New() ComparisonPredicate {
	rv := objc.Call[ComparisonPredicate](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewComparisonPredicate() ComparisonPredicate {
	return ComparisonPredicateClass.New()
}

func (c_ ComparisonPredicate) Init() ComparisonPredicate {
	rv := objc.Call[ComparisonPredicate](c_, objc.Sel("init"))
	return rv
}

func (c_ ComparisonPredicate) PredicateWithSubstitutionVariables(variables map[string]objc.IObject) ComparisonPredicate {
	rv := objc.Call[ComparisonPredicate](c_, objc.Sel("predicateWithSubstitutionVariables:"), variables)
	return rv
}

// Returns a copy of the predicate and substitutes the predicates variables with specified values from a specified substitution variables dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspredicate/1413227-predicatewithsubstitutionvariabl?language=objc
func ComparisonPredicate_PredicateWithSubstitutionVariables(variables map[string]objc.IObject) ComparisonPredicate {
	return ComparisonPredicateClass.Alloc().PredicateWithSubstitutionVariables(variables)
}

// Returns a new predicate formed by combining the left and right expressions using a given selector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscomparisonpredicate/1568141-predicatewithleftexpression?language=objc
func (cc _ComparisonPredicateClass) PredicateWithLeftExpressionRightExpressionCustomSelector(lhs IExpression, rhs IExpression, selector objc.Selector) ComparisonPredicate {
	rv := objc.Call[ComparisonPredicate](cc, objc.Sel("predicateWithLeftExpression:rightExpression:customSelector:"), objc.Ptr(lhs), objc.Ptr(rhs), selector)
	return rv
}

// Returns a new predicate formed by combining the left and right expressions using a given selector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscomparisonpredicate/1568141-predicatewithleftexpression?language=objc
func ComparisonPredicate_PredicateWithLeftExpressionRightExpressionCustomSelector(lhs IExpression, rhs IExpression, selector objc.Selector) ComparisonPredicate {
	return ComparisonPredicateClass.PredicateWithLeftExpressionRightExpressionCustomSelector(lhs, rhs, selector)
}

// The options to use for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscomparisonpredicate/1414069-options?language=objc
func (c_ ComparisonPredicate) Options() ComparisonPredicateOptions {
	rv := objc.Call[ComparisonPredicateOptions](c_, objc.Sel("options"))
	return rv
}

// The selector for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscomparisonpredicate/1413661-customselector?language=objc
func (c_ ComparisonPredicate) CustomSelector() objc.Selector {
	rv := objc.Call[objc.Selector](c_, objc.Sel("customSelector"))
	return rv
}

// The comparison predicate modifier for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscomparisonpredicate/1416376-comparisonpredicatemodifier?language=objc
func (c_ ComparisonPredicate) ComparisonPredicateModifier() ComparisonPredicateModifier {
	rv := objc.Call[ComparisonPredicateModifier](c_, objc.Sel("comparisonPredicateModifier"))
	return rv
}

// The right expression for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscomparisonpredicate/1409469-rightexpression?language=objc
func (c_ ComparisonPredicate) RightExpression() Expression {
	rv := objc.Call[Expression](c_, objc.Sel("rightExpression"))
	return rv
}

// The left expression for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscomparisonpredicate/1412552-leftexpression?language=objc
func (c_ ComparisonPredicate) LeftExpression() Expression {
	rv := objc.Call[Expression](c_, objc.Sel("leftExpression"))
	return rv
}

// The predicate type for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscomparisonpredicate/1418327-predicateoperatortype?language=objc
func (c_ ComparisonPredicate) PredicateOperatorType() PredicateOperatorType {
	rv := objc.Call[PredicateOperatorType](c_, objc.Sel("predicateOperatorType"))
	return rv
}
