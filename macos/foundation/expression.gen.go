// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var ExpressionClass = _ExpressionClass{objc.GetClass("NSExpression")}

type _ExpressionClass struct {
	objc.Class
}

type IExpression interface {
	objc.IObject
	ExpressionValueWithObjectContext(object objc.IObject, context IMutableDictionary) objc.Object
	AllowEvaluation()
	Arguments() []Expression
	Collection() objc.Object
	ConstantValue() objc.Object
	ExpressionType() ExpressionType
	Function() string
	KeyPath() string
	Operand() Expression
	Predicate() Predicate
	LeftExpression() Expression
	RightExpression() Expression
	Variable() string
	FalseExpression() Expression
	TrueExpression() Expression
	ExpressionBlock() func(param1 objc.Object, param2 []Expression, param3 MutableDictionary) objc.Object
}

type Expression struct {
	objc.Object
}

func MakeExpression(ptr unsafe.Pointer) Expression {
	return Expression{
		Object: objc.MakeObject(ptr),
	}
}

func (e_ Expression) InitWithExpressionType(type_ ExpressionType) Expression {
	rv := objc.CallMethod[Expression](e_, objc.GetSelector("initWithExpressionType:"), type_)
	return rv
}

func Expression_InitWithExpressionType(type_ ExpressionType) Expression {
	return ExpressionClass.Alloc().InitWithExpressionType(type_)
}

func (ec _ExpressionClass) Alloc() Expression {
	rv := objc.CallMethod[Expression](ec, objc.GetSelector("alloc"))
	return rv
}

func Expression_Alloc() Expression {
	return ExpressionClass.Alloc()
}

func (ec _ExpressionClass) New() Expression {
	rv := objc.CallMethod[Expression](ec, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewExpression() Expression {
	return ExpressionClass.New()
}

func Expression_New() Expression {
	return ExpressionClass.New()
}

func (e_ Expression) Init() Expression {
	rv := objc.CallMethod[Expression](e_, objc.GetSelector("init"))
	return rv
}

func Expression_Init() Expression {
	return ExpressionClass.Alloc().Init()
}

func (ec _ExpressionClass) ExpressionWithFormatArgumentArray(expressionFormat string, arguments []objc.IObject) Expression {
	rv := objc.CallMethod[Expression](ec, objc.GetSelector("expressionWithFormat:argumentArray:"), expressionFormat, arguments)
	return rv
}

func Expression_ExpressionWithFormatArgumentArray(expressionFormat string, arguments []objc.IObject) Expression {
	return ExpressionClass.ExpressionWithFormatArgumentArray(expressionFormat, arguments)
}

func (ec _ExpressionClass) ExpressionForConstantValue(obj objc.IObject) Expression {
	rv := objc.CallMethod[Expression](ec, objc.GetSelector("expressionForConstantValue:"), objc.ExtractPtr(obj))
	return rv
}

func Expression_ExpressionForConstantValue(obj objc.IObject) Expression {
	return ExpressionClass.ExpressionForConstantValue(obj)
}

func (ec _ExpressionClass) ExpressionForEvaluatedObject() Expression {
	rv := objc.CallMethod[Expression](ec, objc.GetSelector("expressionForEvaluatedObject"))
	return rv
}

func Expression_ExpressionForEvaluatedObject() Expression {
	return ExpressionClass.ExpressionForEvaluatedObject()
}

func (ec _ExpressionClass) ExpressionForKeyPath(keyPath string) Expression {
	rv := objc.CallMethod[Expression](ec, objc.GetSelector("expressionForKeyPath:"), keyPath)
	return rv
}

func Expression_ExpressionForKeyPath(keyPath string) Expression {
	return ExpressionClass.ExpressionForKeyPath(keyPath)
}

func (ec _ExpressionClass) ExpressionForVariable(string_ string) Expression {
	rv := objc.CallMethod[Expression](ec, objc.GetSelector("expressionForVariable:"), string_)
	return rv
}

func Expression_ExpressionForVariable(string_ string) Expression {
	return ExpressionClass.ExpressionForVariable(string_)
}

func (ec _ExpressionClass) ExpressionForAnyKey() Expression {
	rv := objc.CallMethod[Expression](ec, objc.GetSelector("expressionForAnyKey"))
	return rv
}

func Expression_ExpressionForAnyKey() Expression {
	return ExpressionClass.ExpressionForAnyKey()
}

func (ec _ExpressionClass) ExpressionForAggregate(subexpressions []IExpression) Expression {
	rv := objc.CallMethod[Expression](ec, objc.GetSelector("expressionForAggregate:"), subexpressions)
	return rv
}

func Expression_ExpressionForAggregate(subexpressions []IExpression) Expression {
	return ExpressionClass.ExpressionForAggregate(subexpressions)
}

func (ec _ExpressionClass) ExpressionForUnionSetWith(left IExpression, right IExpression) Expression {
	rv := objc.CallMethod[Expression](ec, objc.GetSelector("expressionForUnionSet:with:"), objc.ExtractPtr(left), objc.ExtractPtr(right))
	return rv
}

func Expression_ExpressionForUnionSetWith(left IExpression, right IExpression) Expression {
	return ExpressionClass.ExpressionForUnionSetWith(left, right)
}

func (ec _ExpressionClass) ExpressionForIntersectSetWith(left IExpression, right IExpression) Expression {
	rv := objc.CallMethod[Expression](ec, objc.GetSelector("expressionForIntersectSet:with:"), objc.ExtractPtr(left), objc.ExtractPtr(right))
	return rv
}

func Expression_ExpressionForIntersectSetWith(left IExpression, right IExpression) Expression {
	return ExpressionClass.ExpressionForIntersectSetWith(left, right)
}

func (ec _ExpressionClass) ExpressionForMinusSetWith(left IExpression, right IExpression) Expression {
	rv := objc.CallMethod[Expression](ec, objc.GetSelector("expressionForMinusSet:with:"), objc.ExtractPtr(left), objc.ExtractPtr(right))
	return rv
}

func Expression_ExpressionForMinusSetWith(left IExpression, right IExpression) Expression {
	return ExpressionClass.ExpressionForMinusSetWith(left, right)
}

func (ec _ExpressionClass) ExpressionForSubqueryUsingIteratorVariablePredicate(expression IExpression, variable string, predicate IPredicate) Expression {
	rv := objc.CallMethod[Expression](ec, objc.GetSelector("expressionForSubquery:usingIteratorVariable:predicate:"), objc.ExtractPtr(expression), variable, objc.ExtractPtr(predicate))
	return rv
}

func Expression_ExpressionForSubqueryUsingIteratorVariablePredicate(expression IExpression, variable string, predicate IPredicate) Expression {
	return ExpressionClass.ExpressionForSubqueryUsingIteratorVariablePredicate(expression, variable, predicate)
}

func (ec _ExpressionClass) ExpressionForConditionalTrueExpressionFalseExpression(predicate IPredicate, trueExpression IExpression, falseExpression IExpression) Expression {
	rv := objc.CallMethod[Expression](ec, objc.GetSelector("expressionForConditional:trueExpression:falseExpression:"), objc.ExtractPtr(predicate), objc.ExtractPtr(trueExpression), objc.ExtractPtr(falseExpression))
	return rv
}

func Expression_ExpressionForConditionalTrueExpressionFalseExpression(predicate IPredicate, trueExpression IExpression, falseExpression IExpression) Expression {
	return ExpressionClass.ExpressionForConditionalTrueExpressionFalseExpression(predicate, trueExpression, falseExpression)
}

func (ec _ExpressionClass) ExpressionForBlockArguments(block func(evaluatedObject objc.Object, expressions []Expression, context MutableDictionary) objc.Object, arguments []IExpression) Expression {
	rv := objc.CallMethod[Expression](ec, objc.GetSelector("expressionForBlock:arguments:"), block, arguments)
	return rv
}

func Expression_ExpressionForBlockArguments(block func(evaluatedObject objc.Object, expressions []Expression, context MutableDictionary) objc.Object, arguments []IExpression) Expression {
	return ExpressionClass.ExpressionForBlockArguments(block, arguments)
}

func (ec _ExpressionClass) ExpressionForFunctionArguments(name string, parameters []objc.IObject) Expression {
	rv := objc.CallMethod[Expression](ec, objc.GetSelector("expressionForFunction:arguments:"), name, parameters)
	return rv
}

func Expression_ExpressionForFunctionArguments(name string, parameters []objc.IObject) Expression {
	return ExpressionClass.ExpressionForFunctionArguments(name, parameters)
}

func (ec _ExpressionClass) ExpressionForFunctionSelectorNameArguments(target IExpression, name string, parameters []objc.IObject) Expression {
	rv := objc.CallMethod[Expression](ec, objc.GetSelector("expressionForFunction:selectorName:arguments:"), objc.ExtractPtr(target), name, parameters)
	return rv
}

func Expression_ExpressionForFunctionSelectorNameArguments(target IExpression, name string, parameters []objc.IObject) Expression {
	return ExpressionClass.ExpressionForFunctionSelectorNameArguments(target, name, parameters)
}

func (e_ Expression) ExpressionValueWithObjectContext(object objc.IObject, context IMutableDictionary) objc.Object {
	rv := objc.CallMethod[objc.Object](e_, objc.GetSelector("expressionValueWithObject:context:"), objc.ExtractPtr(object), objc.ExtractPtr(context))
	return rv
}

func (e_ Expression) AllowEvaluation() {
	objc.CallMethod[objc.Void](e_, objc.GetSelector("allowEvaluation"))
}

func (e_ Expression) Arguments() []Expression {
	rv := objc.CallMethod[[]Expression](e_, objc.GetSelector("arguments"))
	// TODO: convert slice items...
	return rv
}

func (e_ Expression) Collection() objc.Object {
	rv := objc.CallMethod[objc.Object](e_, objc.GetSelector("collection"))
	return rv
}

func (e_ Expression) ConstantValue() objc.Object {
	rv := objc.CallMethod[objc.Object](e_, objc.GetSelector("constantValue"))
	return rv
}

func (e_ Expression) ExpressionType() ExpressionType {
	rv := objc.CallMethod[ExpressionType](e_, objc.GetSelector("expressionType"))
	return rv
}

func (e_ Expression) Function() string {
	rv := objc.CallMethod[string](e_, objc.GetSelector("function"))
	return rv
}

func (e_ Expression) KeyPath() string {
	rv := objc.CallMethod[string](e_, objc.GetSelector("keyPath"))
	return rv
}

func (e_ Expression) Operand() Expression {
	rv := objc.CallMethod[Expression](e_, objc.GetSelector("operand"))
	return rv
}

func (e_ Expression) Predicate() Predicate {
	rv := objc.CallMethod[Predicate](e_, objc.GetSelector("predicate"))
	return rv
}

func (e_ Expression) LeftExpression() Expression {
	rv := objc.CallMethod[Expression](e_, objc.GetSelector("leftExpression"))
	return rv
}

func (e_ Expression) RightExpression() Expression {
	rv := objc.CallMethod[Expression](e_, objc.GetSelector("rightExpression"))
	return rv
}

func (e_ Expression) Variable() string {
	rv := objc.CallMethod[string](e_, objc.GetSelector("variable"))
	return rv
}

func (e_ Expression) FalseExpression() Expression {
	rv := objc.CallMethod[Expression](e_, objc.GetSelector("falseExpression"))
	return rv
}

func (e_ Expression) TrueExpression() Expression {
	rv := objc.CallMethod[Expression](e_, objc.GetSelector("trueExpression"))
	return rv
}

func (e_ Expression) ExpressionBlock() func(param1 objc.Object, param2 []Expression, param3 MutableDictionary) objc.Object {
	rv := objc.CallMethod[func(param1 objc.Object, param2 []Expression, param3 MutableDictionary) objc.Object](e_, objc.GetSelector("expressionBlock"))
	return rv
}
