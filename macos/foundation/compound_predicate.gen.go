// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CompoundPredicate] class.
var CompoundPredicateClass = _CompoundPredicateClass{objc.GetClass("NSCompoundPredicate")}

type _CompoundPredicateClass struct {
	objc.Class
}

// An interface definition for the [CompoundPredicate] class.
type ICompoundPredicate interface {
	IPredicate
	Subpredicates() []objc.Object
	CompoundPredicateType() CompoundPredicateType
}

// A specialized predicate that evaluates logical combinations of other predicates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompoundpredicate?language=objc
type CompoundPredicate struct {
	Predicate
}

func CompoundPredicateFrom(ptr unsafe.Pointer) CompoundPredicate {
	return CompoundPredicate{
		Predicate: PredicateFrom(ptr),
	}
}

func (c_ CompoundPredicate) InitWithTypeSubpredicates(type_ CompoundPredicateType, subpredicates []IPredicate) CompoundPredicate {
	rv := objc.Call[CompoundPredicate](c_, objc.Sel("initWithType:subpredicates:"), type_, subpredicates)
	return rv
}

// Returns the receiver that a specified type initializes using predicates from a specified array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompoundpredicate/1407744-initwithtype?language=objc
func NewCompoundPredicateWithTypeSubpredicates(type_ CompoundPredicateType, subpredicates []IPredicate) CompoundPredicate {
	instance := CompoundPredicateClass.Alloc().InitWithTypeSubpredicates(type_, subpredicates)
	instance.Autorelease()
	return instance
}

func (cc _CompoundPredicateClass) Alloc() CompoundPredicate {
	rv := objc.Call[CompoundPredicate](cc, objc.Sel("alloc"))
	return rv
}

func CompoundPredicate_Alloc() CompoundPredicate {
	return CompoundPredicateClass.Alloc()
}

func (cc _CompoundPredicateClass) New() CompoundPredicate {
	rv := objc.Call[CompoundPredicate](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCompoundPredicate() CompoundPredicate {
	return CompoundPredicateClass.New()
}

func (c_ CompoundPredicate) Init() CompoundPredicate {
	rv := objc.Call[CompoundPredicate](c_, objc.Sel("init"))
	return rv
}

func (c_ CompoundPredicate) PredicateWithSubstitutionVariables(variables map[string]objc.IObject) CompoundPredicate {
	rv := objc.Call[CompoundPredicate](c_, objc.Sel("predicateWithSubstitutionVariables:"), variables)
	return rv
}

// Returns a copy of the predicate and substitutes the predicates variables with specified values from a specified substitution variables dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspredicate/1413227-predicatewithsubstitutionvariabl?language=objc
func CompoundPredicate_PredicateWithSubstitutionVariables(variables map[string]objc.IObject) CompoundPredicate {
	instance := CompoundPredicateClass.Alloc().PredicateWithSubstitutionVariables(variables)
	instance.Autorelease()
	return instance
}

// Returns a new predicate that you form using a NOT operation on a specified predicate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompoundpredicate/1409462-notpredicatewithsubpredicate?language=objc
func (cc _CompoundPredicateClass) NotPredicateWithSubpredicate(predicate IPredicate) CompoundPredicate {
	rv := objc.Call[CompoundPredicate](cc, objc.Sel("notPredicateWithSubpredicate:"), objc.Ptr(predicate))
	return rv
}

// Returns a new predicate that you form using a NOT operation on a specified predicate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompoundpredicate/1409462-notpredicatewithsubpredicate?language=objc
func CompoundPredicate_NotPredicateWithSubpredicate(predicate IPredicate) CompoundPredicate {
	return CompoundPredicateClass.NotPredicateWithSubpredicate(predicate)
}

// Returns a new predicate that you form using an AND operation on the predicates in a specified array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompoundpredicate/1407855-andpredicatewithsubpredicates?language=objc
func (cc _CompoundPredicateClass) AndPredicateWithSubpredicates(subpredicates []IPredicate) CompoundPredicate {
	rv := objc.Call[CompoundPredicate](cc, objc.Sel("andPredicateWithSubpredicates:"), subpredicates)
	return rv
}

// Returns a new predicate that you form using an AND operation on the predicates in a specified array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompoundpredicate/1407855-andpredicatewithsubpredicates?language=objc
func CompoundPredicate_AndPredicateWithSubpredicates(subpredicates []IPredicate) CompoundPredicate {
	return CompoundPredicateClass.AndPredicateWithSubpredicates(subpredicates)
}

// Returns a new predicate that you form using an OR operation on the predicates in a specified array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompoundpredicate/1417873-orpredicatewithsubpredicates?language=objc
func (cc _CompoundPredicateClass) OrPredicateWithSubpredicates(subpredicates []IPredicate) CompoundPredicate {
	rv := objc.Call[CompoundPredicate](cc, objc.Sel("orPredicateWithSubpredicates:"), subpredicates)
	return rv
}

// Returns a new predicate that you form using an OR operation on the predicates in a specified array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompoundpredicate/1417873-orpredicatewithsubpredicates?language=objc
func CompoundPredicate_OrPredicateWithSubpredicates(subpredicates []IPredicate) CompoundPredicate {
	return CompoundPredicateClass.OrPredicateWithSubpredicates(subpredicates)
}

// The receiver’s subpredicates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompoundpredicate/1410273-subpredicates?language=objc
func (c_ CompoundPredicate) Subpredicates() []objc.Object {
	rv := objc.Call[[]objc.Object](c_, objc.Sel("subpredicates"))
	return rv
}

// The predicate type for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscompoundpredicate/1412973-compoundpredicatetype?language=objc
func (c_ CompoundPredicate) CompoundPredicateType() CompoundPredicateType {
	rv := objc.Call[CompoundPredicateType](c_, objc.Sel("compoundPredicateType"))
	return rv
}
