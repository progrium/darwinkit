// AUTO-GENERATED CODE, DO NOT MODIFY

package coreml

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NumericConstraint] class.
var NumericConstraintClass = _NumericConstraintClass{objc.GetClass("MLNumericConstraint")}

type _NumericConstraintClass struct {
	objc.Class
}

// An interface definition for the [NumericConstraint] class.
type INumericConstraint interface {
	objc.IObject
	MinNumber() foundation.Number
	EnumeratedNumbers() foundation.Set
	MaxNumber() foundation.Number
}

// The value limitations of a number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlnumericconstraint?language=objc
type NumericConstraint struct {
	objc.Object
}

func NumericConstraintFrom(ptr unsafe.Pointer) NumericConstraint {
	return NumericConstraint{
		Object: objc.ObjectFrom(ptr),
	}
}

func (nc _NumericConstraintClass) Alloc() NumericConstraint {
	rv := objc.Call[NumericConstraint](nc, objc.Sel("alloc"))
	return rv
}

func NumericConstraint_Alloc() NumericConstraint {
	return NumericConstraintClass.Alloc()
}

func (nc _NumericConstraintClass) New() NumericConstraint {
	rv := objc.Call[NumericConstraint](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNumericConstraint() NumericConstraint {
	return NumericConstraintClass.New()
}

func (n_ NumericConstraint) Init() NumericConstraint {
	rv := objc.Call[NumericConstraint](n_, objc.Sel("init"))
	return rv
}

// The smallest numerical value allowed by this constraint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlnumericconstraint/3180066-minnumber?language=objc
func (n_ NumericConstraint) MinNumber() foundation.Number {
	rv := objc.Call[foundation.Number](n_, objc.Sel("minNumber"))
	return rv
}

// A set of the numbers allowed in this constraint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlnumericconstraint/3180064-enumeratednumbers?language=objc
func (n_ NumericConstraint) EnumeratedNumbers() foundation.Set {
	rv := objc.Call[foundation.Set](n_, objc.Sel("enumeratedNumbers"))
	return rv
}

// The largest numerical value allowed by this constraint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlnumericconstraint/3180065-maxnumber?language=objc
func (n_ NumericConstraint) MaxNumber() foundation.Number {
	rv := objc.Call[foundation.Number](n_, objc.Sel("maxNumber"))
	return rv
}
