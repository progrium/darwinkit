// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ExpressionDescription] class.
var ExpressionDescriptionClass = _ExpressionDescriptionClass{objc.GetClass("NSExpressionDescription")}

type _ExpressionDescriptionClass struct {
	objc.Class
}

// An interface definition for the [ExpressionDescription] class.
type IExpressionDescription interface {
	IPropertyDescription
	ExpressionResultType() AttributeType
	SetExpressionResultType(value AttributeType)
	Expression() foundation.Expression
	SetExpression(value foundation.IExpression)
}

// An object that describes an expression to include with a fetch request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsexpressiondescription?language=objc
type ExpressionDescription struct {
	PropertyDescription
}

func ExpressionDescriptionFrom(ptr unsafe.Pointer) ExpressionDescription {
	return ExpressionDescription{
		PropertyDescription: PropertyDescriptionFrom(ptr),
	}
}

func (ec _ExpressionDescriptionClass) Alloc() ExpressionDescription {
	rv := objc.Call[ExpressionDescription](ec, objc.Sel("alloc"))
	return rv
}

func ExpressionDescription_Alloc() ExpressionDescription {
	return ExpressionDescriptionClass.Alloc()
}

func (ec _ExpressionDescriptionClass) New() ExpressionDescription {
	rv := objc.Call[ExpressionDescription](ec, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewExpressionDescription() ExpressionDescription {
	return ExpressionDescriptionClass.New()
}

func (e_ ExpressionDescription) Init() ExpressionDescription {
	rv := objc.Call[ExpressionDescription](e_, objc.Sel("init"))
	return rv
}

// The attribute type of the expression’s result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsexpressiondescription/1506706-expressionresulttype?language=objc
func (e_ ExpressionDescription) ExpressionResultType() AttributeType {
	rv := objc.Call[AttributeType](e_, objc.Sel("expressionResultType"))
	return rv
}

// The attribute type of the expression’s result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsexpressiondescription/1506706-expressionresulttype?language=objc
func (e_ ExpressionDescription) SetExpressionResultType(value AttributeType) {
	objc.Call[objc.Void](e_, objc.Sel("setExpressionResultType:"), value)
}

// The expression to evaluate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsexpressiondescription/1506817-expression?language=objc
func (e_ ExpressionDescription) Expression() foundation.Expression {
	rv := objc.Call[foundation.Expression](e_, objc.Sel("expression"))
	return rv
}

// The expression to evaluate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsexpressiondescription/1506817-expression?language=objc
func (e_ ExpressionDescription) SetExpression(value foundation.IExpression) {
	objc.Call[objc.Void](e_, objc.Sel("setExpression:"), objc.Ptr(value))
}
