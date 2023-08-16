// AUTO-GENERATED CODE, DO NOT MODIFY

package coreml

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SequenceConstraint] class.
var SequenceConstraintClass = _SequenceConstraintClass{objc.GetClass("MLSequenceConstraint")}

type _SequenceConstraintClass struct {
	objc.Class
}

// An interface definition for the [SequenceConstraint] class.
type ISequenceConstraint interface {
	objc.IObject
	ValueDescription() FeatureDescription
	CountRange() foundation.Range
}

// The constraints for a sequence feature. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlsequenceconstraint?language=objc
type SequenceConstraint struct {
	objc.Object
}

func SequenceConstraintFrom(ptr unsafe.Pointer) SequenceConstraint {
	return SequenceConstraint{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _SequenceConstraintClass) Alloc() SequenceConstraint {
	rv := objc.Call[SequenceConstraint](sc, objc.Sel("alloc"))
	return rv
}

func SequenceConstraint_Alloc() SequenceConstraint {
	return SequenceConstraintClass.Alloc()
}

func (sc _SequenceConstraintClass) New() SequenceConstraint {
	rv := objc.Call[SequenceConstraint](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSequenceConstraint() SequenceConstraint {
	return SequenceConstraintClass.New()
}

func (s_ SequenceConstraint) Init() SequenceConstraint {
	rv := objc.Call[SequenceConstraint](s_, objc.Sel("init"))
	return rv
}

// The description that all sequence elements must match. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlsequenceconstraint/2962878-valuedescription?language=objc
func (s_ SequenceConstraint) ValueDescription() FeatureDescription {
	rv := objc.Call[FeatureDescription](s_, objc.Sel("valueDescription"))
	return rv
}

// The range of values allowed for the sequence's length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlsequenceconstraint/2994319-countrange?language=objc
func (s_ SequenceConstraint) CountRange() foundation.Range {
	rv := objc.Call[foundation.Range](s_, objc.Sel("countRange"))
	return rv
}
