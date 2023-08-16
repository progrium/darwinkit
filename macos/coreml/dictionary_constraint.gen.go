// AUTO-GENERATED CODE, DO NOT MODIFY

package coreml

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DictionaryConstraint] class.
var DictionaryConstraintClass = _DictionaryConstraintClass{objc.GetClass("MLDictionaryConstraint")}

type _DictionaryConstraintClass struct {
	objc.Class
}

// An interface definition for the [DictionaryConstraint] class.
type IDictionaryConstraint interface {
	objc.IObject
	KeyType() FeatureType
}

// The constraint on the keys for a dictionary feature. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mldictionaryconstraint?language=objc
type DictionaryConstraint struct {
	objc.Object
}

func DictionaryConstraintFrom(ptr unsafe.Pointer) DictionaryConstraint {
	return DictionaryConstraint{
		Object: objc.ObjectFrom(ptr),
	}
}

func (dc _DictionaryConstraintClass) Alloc() DictionaryConstraint {
	rv := objc.Call[DictionaryConstraint](dc, objc.Sel("alloc"))
	return rv
}

func DictionaryConstraint_Alloc() DictionaryConstraint {
	return DictionaryConstraintClass.Alloc()
}

func (dc _DictionaryConstraintClass) New() DictionaryConstraint {
	rv := objc.Call[DictionaryConstraint](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDictionaryConstraint() DictionaryConstraint {
	return DictionaryConstraintClass.New()
}

func (d_ DictionaryConstraint) Init() DictionaryConstraint {
	rv := objc.Call[DictionaryConstraint](d_, objc.Sel("init"))
	return rv
}

// The key type for the dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mldictionaryconstraint/2921280-keytype?language=objc
func (d_ DictionaryConstraint) KeyType() FeatureType {
	rv := objc.Call[FeatureType](d_, objc.Sel("keyType"))
	return rv
}
