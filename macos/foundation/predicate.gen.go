// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Predicate] class.
var PredicateClass = _PredicateClass{objc.GetClass("NSPredicate")}

type _PredicateClass struct {
	objc.Class
}

// An interface definition for the [Predicate] class.
type IPredicate interface {
	objc.IObject
	AllowEvaluation()
	EvaluateWithObject(object objc.IObject) bool
	PredicateFormat() string
}

// A definition of logical conditions for constraining a search for a fetch or for in-memory filtering. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspredicate?language=objc
type Predicate struct {
	objc.Object
}

func PredicateFrom(ptr unsafe.Pointer) Predicate {
	return Predicate{
		Object: objc.ObjectFrom(ptr),
	}
}

func (p_ Predicate) PredicateWithSubstitutionVariables(variables map[string]objc.IObject) Predicate {
	rv := objc.Call[Predicate](p_, objc.Sel("predicateWithSubstitutionVariables:"), variables)
	return rv
}

// Returns a copy of the predicate and substitutes the predicates variables with specified values from a specified substitution variables dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspredicate/1413227-predicatewithsubstitutionvariabl?language=objc
func Predicate_PredicateWithSubstitutionVariables(variables map[string]objc.IObject) Predicate {
	instance := PredicateClass.Alloc().PredicateWithSubstitutionVariables(variables)
	instance.Autorelease()
	return instance
}

func (pc _PredicateClass) Alloc() Predicate {
	rv := objc.Call[Predicate](pc, objc.Sel("alloc"))
	return rv
}

func Predicate_Alloc() Predicate {
	return PredicateClass.Alloc()
}

func (pc _PredicateClass) New() Predicate {
	rv := objc.Call[Predicate](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPredicate() Predicate {
	return PredicateClass.New()
}

func (p_ Predicate) Init() Predicate {
	rv := objc.Call[Predicate](p_, objc.Sel("init"))
	return rv
}

// Creates a predicate that evaluates using a specified block object and bindings dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspredicate/1416182-predicatewithblock?language=objc
func (pc _PredicateClass) PredicateWithBlock(block func(evaluatedObject objc.Object, bindings map[string]objc.Object) bool) Predicate {
	rv := objc.Call[Predicate](pc, objc.Sel("predicateWithBlock:"), block)
	return rv
}

// Creates a predicate that evaluates using a specified block object and bindings dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspredicate/1416182-predicatewithblock?language=objc
func Predicate_PredicateWithBlock(block func(evaluatedObject objc.Object, bindings map[string]objc.Object) bool) Predicate {
	return PredicateClass.PredicateWithBlock(block)
}

// Forces a securely decoded predicate to allow evaluation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspredicate/1416310-allowevaluation?language=objc
func (p_ Predicate) AllowEvaluation() {
	objc.Call[objc.Void](p_, objc.Sel("allowEvaluation"))
}

// Creates and returns a predicate that always evaluates to a specified Boolean value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspredicate/1417329-predicatewithvalue?language=objc
func (pc _PredicateClass) PredicateWithValue(value bool) Predicate {
	rv := objc.Call[Predicate](pc, objc.Sel("predicateWithValue:"), value)
	return rv
}

// Creates and returns a predicate that always evaluates to a specified Boolean value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspredicate/1417329-predicatewithvalue?language=objc
func Predicate_PredicateWithValue(value bool) Predicate {
	return PredicateClass.PredicateWithValue(value)
}

// Returns a Boolean value that indicates whether the specified object matches the conditions that the predicate specifies. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspredicate/1417924-evaluatewithobject?language=objc
func (p_ Predicate) EvaluateWithObject(object objc.IObject) bool {
	rv := objc.Call[bool](p_, objc.Sel("evaluateWithObject:"), object)
	return rv
}

// Creates a predicate with a metadata query string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspredicate/1417831-predicatefrommetadataquerystring?language=objc
func (pc _PredicateClass) PredicateFromMetadataQueryString(queryString string) Predicate {
	rv := objc.Call[Predicate](pc, objc.Sel("predicateFromMetadataQueryString:"), queryString)
	return rv
}

// Creates a predicate with a metadata query string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspredicate/1417831-predicatefrommetadataquerystring?language=objc
func Predicate_PredicateFromMetadataQueryString(queryString string) Predicate {
	return PredicateClass.PredicateFromMetadataQueryString(queryString)
}

// The predicate's format string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspredicate/1411605-predicateformat?language=objc
func (p_ Predicate) PredicateFormat() string {
	rv := objc.Call[string](p_, objc.Sel("predicateFormat"))
	return rv
}
