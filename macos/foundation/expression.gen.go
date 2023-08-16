// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Expression] class.
var ExpressionClass = _ExpressionClass{objc.GetClass("NSExpression")}

type _ExpressionClass struct {
	objc.Class
}

// An interface definition for the [Expression] class.
type IExpression interface {
	objc.IObject
	AllowEvaluation()
	ExpressionValueWithObjectContext(object objc.IObject, context IMutableDictionary) objc.Object
	ExpressionType() ExpressionType
	Function() string
	Arguments() []Expression
	ExpressionBlock() func(arg0 objc.Object, arg1 []Expression, arg2 MutableDictionary) objc.Object
	ConstantValue() objc.Object
	Variable() string
	Predicate() Predicate
	FalseExpression() Expression
	KeyPath() string
	RightExpression() Expression
	LeftExpression() Expression
	Operand() Expression
	Collection() objc.Object
	TrueExpression() Expression
}

// An expression for use in a comparison predicate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression?language=objc
type Expression struct {
	objc.Object
}

func ExpressionFrom(ptr unsafe.Pointer) Expression {
	return Expression{
		Object: objc.ObjectFrom(ptr),
	}
}

func (e_ Expression) InitWithExpressionType(type_ ExpressionType) Expression {
	rv := objc.Call[Expression](e_, objc.Sel("initWithExpressionType:"), type_)
	return rv
}

// Creates the expression with the specified expression type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1418351-initwithexpressiontype?language=objc
func Expression_InitWithExpressionType(type_ ExpressionType) Expression {
	return ExpressionClass.Alloc().InitWithExpressionType(type_)
}

func (ec _ExpressionClass) Alloc() Expression {
	rv := objc.Call[Expression](ec, objc.Sel("alloc"))
	return rv
}

func Expression_Alloc() Expression {
	return ExpressionClass.Alloc()
}

func (ec _ExpressionClass) New() Expression {
	rv := objc.Call[Expression](ec, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewExpression() Expression {
	return ExpressionClass.New()
}

func (e_ Expression) Init() Expression {
	rv := objc.Call[Expression](e_, objc.Sel("init"))
	return rv
}

// Creates an expression that returns a result, depending on the value of predicate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1418004-expressionforconditional?language=objc
func (ec _ExpressionClass) ExpressionForConditionalTrueExpressionFalseExpression(predicate IPredicate, trueExpression IExpression, falseExpression IExpression) Expression {
	rv := objc.Call[Expression](ec, objc.Sel("expressionForConditional:trueExpression:falseExpression:"), objc.Ptr(predicate), objc.Ptr(trueExpression), objc.Ptr(falseExpression))
	return rv
}

// Creates an expression that returns a result, depending on the value of predicate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1418004-expressionforconditional?language=objc
func Expression_ExpressionForConditionalTrueExpressionFalseExpression(predicate IPredicate, trueExpression IExpression, falseExpression IExpression) Expression {
	return ExpressionClass.ExpressionForConditionalTrueExpressionFalseExpression(predicate, trueExpression, falseExpression)
}

// Creates an expression that represents the object you’re evaluating. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1411845-expressionforevaluatedobject?language=objc
func (ec _ExpressionClass) ExpressionForEvaluatedObject() Expression {
	rv := objc.Call[Expression](ec, objc.Sel("expressionForEvaluatedObject"))
	return rv
}

// Creates an expression that represents the object you’re evaluating. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1411845-expressionforevaluatedobject?language=objc
func Expression_ExpressionForEvaluatedObject() Expression {
	return ExpressionClass.ExpressionForEvaluatedObject()
}

// Creates an expression that invokes the value function with a specified key path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1408892-expressionforkeypath?language=objc
func (ec _ExpressionClass) ExpressionForKeyPath(keyPath string) Expression {
	rv := objc.Call[Expression](ec, objc.Sel("expressionForKeyPath:"), keyPath)
	return rv
}

// Creates an expression that invokes the value function with a specified key path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1408892-expressionforkeypath?language=objc
func Expression_ExpressionForKeyPath(keyPath string) Expression {
	return ExpressionClass.ExpressionForKeyPath(keyPath)
}

// Creates an expression object that uses the block for evaluating objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1407823-expressionforblock?language=objc
func (ec _ExpressionClass) ExpressionForBlockArguments(block func(evaluatedObject objc.Object, expressions []Expression, context MutableDictionary) objc.Object, arguments []IExpression) Expression {
	rv := objc.Call[Expression](ec, objc.Sel("expressionForBlock:arguments:"), block, arguments)
	return rv
}

// Creates an expression object that uses the block for evaluating objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1407823-expressionforblock?language=objc
func Expression_ExpressionForBlockArguments(block func(evaluatedObject objc.Object, expressions []Expression, context MutableDictionary) objc.Object, arguments []IExpression) Expression {
	return ExpressionClass.ExpressionForBlockArguments(block, arguments)
}

// Creates an expression that represents any key for a Spotlight query. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1410198-expressionforanykey?language=objc
func (ec _ExpressionClass) ExpressionForAnyKey() Expression {
	rv := objc.Call[Expression](ec, objc.Sel("expressionForAnyKey"))
	return rv
}

// Creates an expression that represents any key for a Spotlight query. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1410198-expressionforanykey?language=objc
func Expression_ExpressionForAnyKey() Expression {
	return ExpressionClass.ExpressionForAnyKey()
}

// Forces a securely decoded expression to allow evaluation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1409244-allowevaluation?language=objc
func (e_ Expression) AllowEvaluation() {
	objc.Call[objc.Void](e_, objc.Sel("allowEvaluation"))
}

// Creates an expression object that represents the subtraction of a specified collection from a specified set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1417659-expressionforminusset?language=objc
func (ec _ExpressionClass) ExpressionForMinusSetWith(left IExpression, right IExpression) Expression {
	rv := objc.Call[Expression](ec, objc.Sel("expressionForMinusSet:with:"), objc.Ptr(left), objc.Ptr(right))
	return rv
}

// Creates an expression object that represents the subtraction of a specified collection from a specified set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1417659-expressionforminusset?language=objc
func Expression_ExpressionForMinusSetWith(left IExpression, right IExpression) Expression {
	return ExpressionClass.ExpressionForMinusSetWith(left, right)
}

// Creates an aggregate expression for a specified collection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1418366-expressionforaggregate?language=objc
func (ec _ExpressionClass) ExpressionForAggregate(subexpressions []IExpression) Expression {
	rv := objc.Call[Expression](ec, objc.Sel("expressionForAggregate:"), subexpressions)
	return rv
}

// Creates an aggregate expression for a specified collection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1418366-expressionforaggregate?language=objc
func Expression_ExpressionForAggregate(subexpressions []IExpression) Expression {
	return ExpressionClass.ExpressionForAggregate(subexpressions)
}

// Creates an expression object that represents the union of a specified set and collection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1411585-expressionforunionset?language=objc
func (ec _ExpressionClass) ExpressionForUnionSetWith(left IExpression, right IExpression) Expression {
	rv := objc.Call[Expression](ec, objc.Sel("expressionForUnionSet:with:"), objc.Ptr(left), objc.Ptr(right))
	return rv
}

// Creates an expression object that represents the union of a specified set and collection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1411585-expressionforunionset?language=objc
func Expression_ExpressionForUnionSetWith(left IExpression, right IExpression) Expression {
	return ExpressionClass.ExpressionForUnionSetWith(left, right)
}

// Creates an expression that filters a collection by storing elements in the collection in a specified variable and keeping the elements that the qualifier returns as true. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1411651-expressionforsubquery?language=objc
func (ec _ExpressionClass) ExpressionForSubqueryUsingIteratorVariablePredicate(expression IExpression, variable string, predicate IPredicate) Expression {
	rv := objc.Call[Expression](ec, objc.Sel("expressionForSubquery:usingIteratorVariable:predicate:"), objc.Ptr(expression), variable, objc.Ptr(predicate))
	return rv
}

// Creates an expression that filters a collection by storing elements in the collection in a specified variable and keeping the elements that the qualifier returns as true. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1411651-expressionforsubquery?language=objc
func Expression_ExpressionForSubqueryUsingIteratorVariablePredicate(expression IExpression, variable string, predicate IPredicate) Expression {
	return ExpressionClass.ExpressionForSubqueryUsingIteratorVariablePredicate(expression, variable, predicate)
}

// Creates an expression that extracts a value from the variable bindings dictionary for a specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1417593-expressionforvariable?language=objc
func (ec _ExpressionClass) ExpressionForVariable(string_ string) Expression {
	rv := objc.Call[Expression](ec, objc.Sel("expressionForVariable:"), string_)
	return rv
}

// Creates an expression that extracts a value from the variable bindings dictionary for a specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1417593-expressionforvariable?language=objc
func Expression_ExpressionForVariable(string_ string) Expression {
	return ExpressionClass.ExpressionForVariable(string_)
}

// Creates an expression object that represents the intersection of a specified set and collection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1409706-expressionforintersectset?language=objc
func (ec _ExpressionClass) ExpressionForIntersectSetWith(left IExpression, right IExpression) Expression {
	rv := objc.Call[Expression](ec, objc.Sel("expressionForIntersectSet:with:"), objc.Ptr(left), objc.Ptr(right))
	return rv
}

// Creates an expression object that represents the intersection of a specified set and collection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1409706-expressionforintersectset?language=objc
func Expression_ExpressionForIntersectSetWith(left IExpression, right IExpression) Expression {
	return ExpressionClass.ExpressionForIntersectSetWith(left, right)
}

// Creates an expression that represents a specified constant value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1415818-expressionforconstantvalue?language=objc
func (ec _ExpressionClass) ExpressionForConstantValue(obj objc.IObject) Expression {
	rv := objc.Call[Expression](ec, objc.Sel("expressionForConstantValue:"), obj)
	return rv
}

// Creates an expression that represents a specified constant value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1415818-expressionforconstantvalue?language=objc
func Expression_ExpressionForConstantValue(obj objc.IObject) Expression {
	return ExpressionClass.ExpressionForConstantValue(obj)
}

// Evaluates an expression using a specified object and context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1410363-expressionvaluewithobject?language=objc
func (e_ Expression) ExpressionValueWithObjectContext(object objc.IObject, context IMutableDictionary) objc.Object {
	rv := objc.Call[objc.Object](e_, objc.Sel("expressionValueWithObject:context:"), object, objc.Ptr(context))
	return rv
}

// Creates an expression that returns the result of invoking a selector with a specified name using specified arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1412905-expressionforfunction?language=objc
func (ec _ExpressionClass) ExpressionForFunctionSelectorNameArguments(target IExpression, name string, parameters []objc.IObject) Expression {
	rv := objc.Call[Expression](ec, objc.Sel("expressionForFunction:selectorName:arguments:"), objc.Ptr(target), name, parameters)
	return rv
}

// Creates an expression that returns the result of invoking a selector with a specified name using specified arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1412905-expressionforfunction?language=objc
func Expression_ExpressionForFunctionSelectorNameArguments(target IExpression, name string, parameters []objc.IObject) Expression {
	return ExpressionClass.ExpressionForFunctionSelectorNameArguments(target, name, parameters)
}

// Creates the expression with the specified expression arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1587937-expressionwithformat?language=objc
func (ec _ExpressionClass) ExpressionWithFormat(expressionFormat string) Expression {
	rv := objc.Call[Expression](ec, objc.Sel("expressionWithFormat:"), expressionFormat)
	return rv
}

// Creates the expression with the specified expression arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1587937-expressionwithformat?language=objc
func Expression_ExpressionWithFormat(expressionFormat string) Expression {
	return ExpressionClass.ExpressionWithFormat(expressionFormat)
}

// The expression type for the expression. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1416975-expressiontype?language=objc
func (e_ Expression) ExpressionType() ExpressionType {
	rv := objc.Call[ExpressionType](e_, objc.Sel("expressionType"))
	return rv
}

// The function for the expression. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1416200-function?language=objc
func (e_ Expression) Function() string {
	rv := objc.Call[string](e_, objc.Sel("function"))
	return rv
}

// The arguments for the expression. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1411559-arguments?language=objc
func (e_ Expression) Arguments() []Expression {
	rv := objc.Call[[]Expression](e_, objc.Sel("arguments"))
	return rv
}

// The block that executes to evaluate the expression. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1409139-expressionblock?language=objc
func (e_ Expression) ExpressionBlock() func(arg0 objc.Object, arg1 []Expression, arg2 MutableDictionary) objc.Object {
	rv := objc.Call[func(arg0 objc.Object, arg1 []Expression, arg2 MutableDictionary) objc.Object](e_, objc.Sel("expressionBlock"))
	return rv
}

// The constant value of the expression. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1418093-constantvalue?language=objc
func (e_ Expression) ConstantValue() objc.Object {
	rv := objc.Call[objc.Object](e_, objc.Sel("constantValue"))
	return rv
}

// The variable for the expression. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1413759-variable?language=objc
func (e_ Expression) Variable() string {
	rv := objc.Call[string](e_, objc.Sel("variable"))
	return rv
}

// The predicate of a subquery expression. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1407531-predicate?language=objc
func (e_ Expression) Predicate() Predicate {
	rv := objc.Call[Predicate](e_, objc.Sel("predicate"))
	return rv
}

// An expression to evalutate if a conditional expression’s predicate evaluates to false. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1416488-falseexpression?language=objc
func (e_ Expression) FalseExpression() Expression {
	rv := objc.Call[Expression](e_, objc.Sel("falseExpression"))
	return rv
}

// The key path for the expression. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1416071-keypath?language=objc
func (e_ Expression) KeyPath() string {
	rv := objc.Call[string](e_, objc.Sel("keyPath"))
	return rv
}

// The right expression of an aggregate expression. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1416583-rightexpression?language=objc
func (e_ Expression) RightExpression() Expression {
	rv := objc.Call[Expression](e_, objc.Sel("rightExpression"))
	return rv
}

// The left expression of an aggregate expression. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1415792-leftexpression?language=objc
func (e_ Expression) LeftExpression() Expression {
	rv := objc.Call[Expression](e_, objc.Sel("leftExpression"))
	return rv
}

// The operand for the expression. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1413698-operand?language=objc
func (e_ Expression) Operand() Expression {
	rv := objc.Call[Expression](e_, objc.Sel("operand"))
	return rv
}

// The collection of expressions in an aggregate expression, or the collection element of a subquery expression. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1415684-collection?language=objc
func (e_ Expression) Collection() objc.Object {
	rv := objc.Call[objc.Object](e_, objc.Sel("collection"))
	return rv
}

// An expression to evalutate if a conditional expression’s predicate evaluates to true. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1411874-trueexpression?language=objc
func (e_ Expression) TrueExpression() Expression {
	rv := objc.Call[Expression](e_, objc.Sel("trueExpression"))
	return rv
}
