// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DerivedAttributeDescription] class.
var DerivedAttributeDescriptionClass = _DerivedAttributeDescriptionClass{objc.GetClass("NSDerivedAttributeDescription")}

type _DerivedAttributeDescriptionClass struct {
	objc.Class
}

// An interface definition for the [DerivedAttributeDescription] class.
type IDerivedAttributeDescription interface {
	IAttributeDescription
	DerivationExpression() foundation.Expression
	SetDerivationExpression(value foundation.IExpression)
}

// A description of an attribute that derives its value by performing a calculation on a related attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsderivedattributedescription?language=objc
type DerivedAttributeDescription struct {
	AttributeDescription
}

func DerivedAttributeDescriptionFrom(ptr unsafe.Pointer) DerivedAttributeDescription {
	return DerivedAttributeDescription{
		AttributeDescription: AttributeDescriptionFrom(ptr),
	}
}

func (dc _DerivedAttributeDescriptionClass) Alloc() DerivedAttributeDescription {
	rv := objc.Call[DerivedAttributeDescription](dc, objc.Sel("alloc"))
	return rv
}

func DerivedAttributeDescription_Alloc() DerivedAttributeDescription {
	return DerivedAttributeDescriptionClass.Alloc()
}

func (dc _DerivedAttributeDescriptionClass) New() DerivedAttributeDescription {
	rv := objc.Call[DerivedAttributeDescription](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDerivedAttributeDescription() DerivedAttributeDescription {
	return DerivedAttributeDescriptionClass.New()
}

func (d_ DerivedAttributeDescription) Init() DerivedAttributeDescription {
	rv := objc.Call[DerivedAttributeDescription](d_, objc.Sel("init"))
	return rv
}

// An expression for generating derived data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsderivedattributedescription/3174854-derivationexpression?language=objc
func (d_ DerivedAttributeDescription) DerivationExpression() foundation.Expression {
	rv := objc.Call[foundation.Expression](d_, objc.Sel("derivationExpression"))
	return rv
}

// An expression for generating derived data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsderivedattributedescription/3174854-derivationexpression?language=objc
func (d_ DerivedAttributeDescription) SetDerivationExpression(value foundation.IExpression) {
	objc.Call[objc.Void](d_, objc.Sel("setDerivationExpression:"), objc.Ptr(value))
}
