// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DecimalNumberHandler] class.
var DecimalNumberHandlerClass = _DecimalNumberHandlerClass{objc.GetClass("NSDecimalNumberHandler")}

type _DecimalNumberHandlerClass struct {
	objc.Class
}

// An interface definition for the [DecimalNumberHandler] class.
type IDecimalNumberHandler interface {
	objc.IObject
}

// A class that adopts the decimal number behaviors protocol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumberhandler?language=objc
type DecimalNumberHandler struct {
	objc.Object
}

func DecimalNumberHandlerFrom(ptr unsafe.Pointer) DecimalNumberHandler {
	return DecimalNumberHandler{
		Object: objc.ObjectFrom(ptr),
	}
}

func (d_ DecimalNumberHandler) InitWithRoundingModeScaleRaiseOnExactnessRaiseOnOverflowRaiseOnUnderflowRaiseOnDivideByZero(roundingMode RoundingMode, scale int, exact bool, overflow bool, underflow bool, divideByZero bool) DecimalNumberHandler {
	rv := objc.Call[DecimalNumberHandler](d_, objc.Sel("initWithRoundingMode:scale:raiseOnExactness:raiseOnOverflow:raiseOnUnderflow:raiseOnDivideByZero:"), roundingMode, scale, exact, overflow, underflow, divideByZero)
	return rv
}

// Returns an NSDecimalNumberHandler object initialized so it behaves as specified by the method’s arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumberhandler/1416492-initwithroundingmode?language=objc
func NewDecimalNumberHandlerWithRoundingModeScaleRaiseOnExactnessRaiseOnOverflowRaiseOnUnderflowRaiseOnDivideByZero(roundingMode RoundingMode, scale int, exact bool, overflow bool, underflow bool, divideByZero bool) DecimalNumberHandler {
	instance := DecimalNumberHandlerClass.Alloc().InitWithRoundingModeScaleRaiseOnExactnessRaiseOnOverflowRaiseOnUnderflowRaiseOnDivideByZero(roundingMode, scale, exact, overflow, underflow, divideByZero)
	instance.Autorelease()
	return instance
}

func (dc _DecimalNumberHandlerClass) DecimalNumberHandlerWithRoundingModeScaleRaiseOnExactnessRaiseOnOverflowRaiseOnUnderflowRaiseOnDivideByZero(roundingMode RoundingMode, scale int, exact bool, overflow bool, underflow bool, divideByZero bool) DecimalNumberHandler {
	rv := objc.Call[DecimalNumberHandler](dc, objc.Sel("decimalNumberHandlerWithRoundingMode:scale:raiseOnExactness:raiseOnOverflow:raiseOnUnderflow:raiseOnDivideByZero:"), roundingMode, scale, exact, overflow, underflow, divideByZero)
	return rv
}

// Returns an NSDecimalNumberHandler object with customized behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumberhandler/1578295-decimalnumberhandlerwithrounding?language=objc
func DecimalNumberHandler_DecimalNumberHandlerWithRoundingModeScaleRaiseOnExactnessRaiseOnOverflowRaiseOnUnderflowRaiseOnDivideByZero(roundingMode RoundingMode, scale int, exact bool, overflow bool, underflow bool, divideByZero bool) DecimalNumberHandler {
	return DecimalNumberHandlerClass.DecimalNumberHandlerWithRoundingModeScaleRaiseOnExactnessRaiseOnOverflowRaiseOnUnderflowRaiseOnDivideByZero(roundingMode, scale, exact, overflow, underflow, divideByZero)
}

func (dc _DecimalNumberHandlerClass) Alloc() DecimalNumberHandler {
	rv := objc.Call[DecimalNumberHandler](dc, objc.Sel("alloc"))
	return rv
}

func DecimalNumberHandler_Alloc() DecimalNumberHandler {
	return DecimalNumberHandlerClass.Alloc()
}

func (dc _DecimalNumberHandlerClass) New() DecimalNumberHandler {
	rv := objc.Call[DecimalNumberHandler](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDecimalNumberHandler() DecimalNumberHandler {
	return DecimalNumberHandlerClass.New()
}

func (d_ DecimalNumberHandler) Init() DecimalNumberHandler {
	rv := objc.Call[DecimalNumberHandler](d_, objc.Sel("init"))
	return rv
}

// Returns the default instance of NSDecimalNumberHandler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumberhandler/1407825-defaultdecimalnumberhandler?language=objc
func (dc _DecimalNumberHandlerClass) DefaultDecimalNumberHandler() DecimalNumberHandler {
	rv := objc.Call[DecimalNumberHandler](dc, objc.Sel("defaultDecimalNumberHandler"))
	return rv
}

// Returns the default instance of NSDecimalNumberHandler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumberhandler/1407825-defaultdecimalnumberhandler?language=objc
func DecimalNumberHandler_DefaultDecimalNumberHandler() DecimalNumberHandler {
	return DecimalNumberHandlerClass.DefaultDecimalNumberHandler()
}
