// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that declares three methods that control the discretionary aspects of working with decimal numbers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumberbehaviors?language=objc
type PDecimalNumberBehaviors interface {
	// optional
	Scale() int
	HasScale() bool

	// optional
	ExceptionDuringOperationErrorLeftOperandRightOperand(operation objc.Selector, error CalculationError, leftOperand DecimalNumber, rightOperand DecimalNumber) IDecimalNumber
	HasExceptionDuringOperationErrorLeftOperandRightOperand() bool

	// optional
	RoundingMode() RoundingMode
	HasRoundingMode() bool
}

// A concrete type wrapper for the [PDecimalNumberBehaviors] protocol.
type DecimalNumberBehaviorsWrapper struct {
	objc.Object
}

func (d_ DecimalNumberBehaviorsWrapper) HasScale() bool {
	return d_.RespondsToSelector(objc.Sel("scale"))
}

// Returns the number of digits allowed after the decimal separator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumberbehaviors/1411642-scale?language=objc
func (d_ DecimalNumberBehaviorsWrapper) Scale() int {
	rv := objc.Call[int](d_, objc.Sel("scale"))
	return rv
}

func (d_ DecimalNumberBehaviorsWrapper) HasExceptionDuringOperationErrorLeftOperandRightOperand() bool {
	return d_.RespondsToSelector(objc.Sel("exceptionDuringOperation:error:leftOperand:rightOperand:"))
}

// Specifies what an NSDecimalNumber object will do when it encounters an error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumberbehaviors/1411766-exceptionduringoperation?language=objc
func (d_ DecimalNumberBehaviorsWrapper) ExceptionDuringOperationErrorLeftOperandRightOperand(operation objc.Selector, error CalculationError, leftOperand IDecimalNumber, rightOperand IDecimalNumber) DecimalNumber {
	rv := objc.Call[DecimalNumber](d_, objc.Sel("exceptionDuringOperation:error:leftOperand:rightOperand:"), operation, error, objc.Ptr(leftOperand), objc.Ptr(rightOperand))
	return rv
}

func (d_ DecimalNumberBehaviorsWrapper) HasRoundingMode() bool {
	return d_.RespondsToSelector(objc.Sel("roundingMode"))
}

// Returns the way that NSDecimalNumber's decimalNumberBy... methods round their return values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumberbehaviors/1409414-roundingmode?language=objc
func (d_ DecimalNumberBehaviorsWrapper) RoundingMode() RoundingMode {
	rv := objc.Call[RoundingMode](d_, objc.Sel("roundingMode"))
	return rv
}
