// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FetchRequestExpression] class.
var FetchRequestExpressionClass = _FetchRequestExpressionClass{objc.GetClass("NSFetchRequestExpression")}

type _FetchRequestExpressionClass struct {
	objc.Class
}

// An interface definition for the [FetchRequestExpression] class.
type IFetchRequestExpression interface {
	foundation.IExpression
	RequestExpression() foundation.Expression
	IsCountOnlyRequest() bool
	ContextExpression() foundation.Expression
}

// An expression that evaluates the result of a fetch request on a managed object context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequestexpression?language=objc
type FetchRequestExpression struct {
	foundation.Expression
}

func FetchRequestExpressionFrom(ptr unsafe.Pointer) FetchRequestExpression {
	return FetchRequestExpression{
		Expression: foundation.ExpressionFrom(ptr),
	}
}

func (fc _FetchRequestExpressionClass) Alloc() FetchRequestExpression {
	rv := objc.Call[FetchRequestExpression](fc, objc.Sel("alloc"))
	return rv
}

func FetchRequestExpression_Alloc() FetchRequestExpression {
	return FetchRequestExpressionClass.Alloc()
}

func (fc _FetchRequestExpressionClass) New() FetchRequestExpression {
	rv := objc.Call[FetchRequestExpression](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFetchRequestExpression() FetchRequestExpression {
	return FetchRequestExpressionClass.New()
}

func (f_ FetchRequestExpression) Init() FetchRequestExpression {
	rv := objc.Call[FetchRequestExpression](f_, objc.Sel("init"))
	return rv
}

func (f_ FetchRequestExpression) InitWithExpressionType(type_ foundation.ExpressionType) FetchRequestExpression {
	rv := objc.Call[FetchRequestExpression](f_, objc.Sel("initWithExpressionType:"), type_)
	return rv
}

// Creates the expression with the specified expression type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexpression/1418351-initwithexpressiontype?language=objc
func NewFetchRequestExpressionWithExpressionType(type_ foundation.ExpressionType) FetchRequestExpression {
	instance := FetchRequestExpressionClass.Alloc().InitWithExpressionType(type_)
	instance.Autorelease()
	return instance
}

// Returns an expression which will evaluate to the result of executing a fetch request on a context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequestexpression/1391661-expressionforfetch?language=objc
func (fc _FetchRequestExpressionClass) ExpressionForFetchContextCountOnly(fetch foundation.IExpression, context foundation.IExpression, countFlag bool) foundation.Expression {
	rv := objc.Call[foundation.Expression](fc, objc.Sel("expressionForFetch:context:countOnly:"), objc.Ptr(fetch), objc.Ptr(context), countFlag)
	return rv
}

// Returns an expression which will evaluate to the result of executing a fetch request on a context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequestexpression/1391661-expressionforfetch?language=objc
func FetchRequestExpression_ExpressionForFetchContextCountOnly(fetch foundation.IExpression, context foundation.IExpression, countFlag bool) foundation.Expression {
	return FetchRequestExpressionClass.ExpressionForFetchContextCountOnly(fetch, context, countFlag)
}

// The expression for the receiver’s fetch request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequestexpression/1391672-requestexpression?language=objc
func (f_ FetchRequestExpression) RequestExpression() foundation.Expression {
	rv := objc.Call[foundation.Expression](f_, objc.Sel("requestExpression"))
	return rv
}

// Returns a Boolean value that indicates whether the receiver represents a count-only fetch request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequestexpression/1391670-countonlyrequest?language=objc
func (f_ FetchRequestExpression) IsCountOnlyRequest() bool {
	rv := objc.Call[bool](f_, objc.Sel("isCountOnlyRequest"))
	return rv
}

// The expression for the receiver’s managed object context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequestexpression/1391665-contextexpression?language=objc
func (f_ FetchRequestExpression) ContextExpression() foundation.Expression {
	rv := objc.Call[foundation.Expression](f_, objc.Sel("contextExpression"))
	return rv
}
