// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DecimalNumber] class.
var DecimalNumberClass = _DecimalNumberClass{objc.GetClass("NSDecimalNumber")}

type _DecimalNumberClass struct {
	objc.Class
}

// An interface definition for the [DecimalNumber] class.
type IDecimalNumber interface {
	INumber
	DecimalNumberByMultiplyingByWithBehavior(decimalNumber IDecimalNumber, behavior PDecimalNumberBehaviors) DecimalNumber
	DecimalNumberByMultiplyingByWithBehaviorObject(decimalNumber IDecimalNumber, behaviorObject objc.IObject) DecimalNumber
	DecimalNumberByRoundingAccordingToBehavior(behavior PDecimalNumberBehaviors) DecimalNumber
	DecimalNumberByRoundingAccordingToBehaviorObject(behaviorObject objc.IObject) DecimalNumber
	DecimalNumberBySubtracting(decimalNumber IDecimalNumber) DecimalNumber
	DecimalNumberByMultiplyingByPowerOf10WithBehavior(power int, behavior PDecimalNumberBehaviors) DecimalNumber
	DecimalNumberByMultiplyingByPowerOf10WithBehaviorObject(power int, behaviorObject objc.IObject) DecimalNumber
	DecimalNumberByAddingWithBehavior(decimalNumber IDecimalNumber, behavior PDecimalNumberBehaviors) DecimalNumber
	DecimalNumberByAddingWithBehaviorObject(decimalNumber IDecimalNumber, behaviorObject objc.IObject) DecimalNumber
	DecimalNumberByDividingByWithBehavior(decimalNumber IDecimalNumber, behavior PDecimalNumberBehaviors) DecimalNumber
	DecimalNumberByDividingByWithBehaviorObject(decimalNumber IDecimalNumber, behaviorObject objc.IObject) DecimalNumber
	DecimalNumberByRaisingToPower(power uint) DecimalNumber
}

// An object for representing and performing arithmetic on base-10 numbers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber?language=objc
type DecimalNumber struct {
	Number
}

func DecimalNumberFrom(ptr unsafe.Pointer) DecimalNumber {
	return DecimalNumber{
		Number: NumberFrom(ptr),
	}
}

func (d_ DecimalNumber) InitWithString(numberValue string) DecimalNumber {
	rv := objc.Call[DecimalNumber](d_, objc.Sel("initWithString:"), numberValue)
	return rv
}

// Initializes a decimal number so that its value is equivalent to that in a given numeric string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1409902-initwithstring?language=objc
func NewDecimalNumberWithString(numberValue string) DecimalNumber {
	instance := DecimalNumberClass.Alloc().InitWithString(numberValue)
	instance.Autorelease()
	return instance
}

func (d_ DecimalNumber) InitWithDecimal(dcm Decimal) DecimalNumber {
	rv := objc.Call[DecimalNumber](d_, objc.Sel("initWithDecimal:"), dcm)
	return rv
}

// Initializes a decimal number to represent a given decimal. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1412692-initwithdecimal?language=objc
func NewDecimalNumberWithDecimal(dcm Decimal) DecimalNumber {
	instance := DecimalNumberClass.Alloc().InitWithDecimal(dcm)
	instance.Autorelease()
	return instance
}

func (d_ DecimalNumber) InitWithMantissaExponentIsNegative(mantissa int64, exponent int, flag bool) DecimalNumber {
	rv := objc.Call[DecimalNumber](d_, objc.Sel("initWithMantissa:exponent:isNegative:"), mantissa, exponent, flag)
	return rv
}

// Initializes a decimal number using the given mantissa, exponent, and sign. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1416003-initwithmantissa?language=objc
func NewDecimalNumberWithMantissaExponentIsNegative(mantissa int64, exponent int, flag bool) DecimalNumber {
	instance := DecimalNumberClass.Alloc().InitWithMantissaExponentIsNegative(mantissa, exponent, flag)
	instance.Autorelease()
	return instance
}

func (dc _DecimalNumberClass) Alloc() DecimalNumber {
	rv := objc.Call[DecimalNumber](dc, objc.Sel("alloc"))
	return rv
}

func DecimalNumber_Alloc() DecimalNumber {
	return DecimalNumberClass.Alloc()
}

func (dc _DecimalNumberClass) New() DecimalNumber {
	rv := objc.Call[DecimalNumber](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDecimalNumber() DecimalNumber {
	return DecimalNumberClass.New()
}

func (d_ DecimalNumber) Init() DecimalNumber {
	rv := objc.Call[DecimalNumber](d_, objc.Sel("init"))
	return rv
}

func (d_ DecimalNumber) InitWithBytesObjCType(value unsafe.Pointer, type_ *uint8) DecimalNumber {
	rv := objc.Call[DecimalNumber](d_, objc.Sel("initWithBytes:objCType:"), value, type_)
	return rv
}

// Initializes a value object to contain the specified value, interpreted with the specified Objective-C type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1411621-initwithbytes?language=objc
func NewDecimalNumberWithBytesObjCType(value unsafe.Pointer, type_ *uint8) DecimalNumber {
	instance := DecimalNumberClass.Alloc().InitWithBytesObjCType(value, type_)
	instance.Autorelease()
	return instance
}

// Multiplies this number by another given number using the specified behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1414874-decimalnumberbymultiplyingby?language=objc
func (d_ DecimalNumber) DecimalNumberByMultiplyingByWithBehavior(decimalNumber IDecimalNumber, behavior PDecimalNumberBehaviors) DecimalNumber {
	po1 := objc.WrapAsProtocol("NSDecimalNumberBehaviors", behavior)
	rv := objc.Call[DecimalNumber](d_, objc.Sel("decimalNumberByMultiplyingBy:withBehavior:"), objc.Ptr(decimalNumber), po1)
	return rv
}

// Multiplies this number by another given number using the specified behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1414874-decimalnumberbymultiplyingby?language=objc
func (d_ DecimalNumber) DecimalNumberByMultiplyingByWithBehaviorObject(decimalNumber IDecimalNumber, behaviorObject objc.IObject) DecimalNumber {
	rv := objc.Call[DecimalNumber](d_, objc.Sel("decimalNumberByMultiplyingBy:withBehavior:"), objc.Ptr(decimalNumber), objc.Ptr(behaviorObject))
	return rv
}

// Returns a rounded version of the decimal number using the specified rounding behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1415436-decimalnumberbyroundingaccording?language=objc
func (d_ DecimalNumber) DecimalNumberByRoundingAccordingToBehavior(behavior PDecimalNumberBehaviors) DecimalNumber {
	po0 := objc.WrapAsProtocol("NSDecimalNumberBehaviors", behavior)
	rv := objc.Call[DecimalNumber](d_, objc.Sel("decimalNumberByRoundingAccordingToBehavior:"), po0)
	return rv
}

// Returns a rounded version of the decimal number using the specified rounding behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1415436-decimalnumberbyroundingaccording?language=objc
func (d_ DecimalNumber) DecimalNumberByRoundingAccordingToBehaviorObject(behaviorObject objc.IObject) DecimalNumber {
	rv := objc.Call[DecimalNumber](d_, objc.Sel("decimalNumberByRoundingAccordingToBehavior:"), objc.Ptr(behaviorObject))
	return rv
}

// Subtracts another given number from this one. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1416873-decimalnumberbysubtracting?language=objc
func (d_ DecimalNumber) DecimalNumberBySubtracting(decimalNumber IDecimalNumber) DecimalNumber {
	rv := objc.Call[DecimalNumber](d_, objc.Sel("decimalNumberBySubtracting:"), objc.Ptr(decimalNumber))
	return rv
}

// Creates and returns a decimal number equivalent to a given decimal structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1578293-decimalnumberwithdecimal?language=objc
func (dc _DecimalNumberClass) DecimalNumberWithDecimal(dcm Decimal) DecimalNumber {
	rv := objc.Call[DecimalNumber](dc, objc.Sel("decimalNumberWithDecimal:"), dcm)
	return rv
}

// Creates and returns a decimal number equivalent to a given decimal structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1578293-decimalnumberwithdecimal?language=objc
func DecimalNumber_DecimalNumberWithDecimal(dcm Decimal) DecimalNumber {
	return DecimalNumberClass.DecimalNumberWithDecimal(dcm)
}

// Creates and returns a decimal number equivalent to the number specified by the arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1578294-decimalnumberwithmantissa?language=objc
func (dc _DecimalNumberClass) DecimalNumberWithMantissaExponentIsNegative(mantissa int64, exponent int, flag bool) DecimalNumber {
	rv := objc.Call[DecimalNumber](dc, objc.Sel("decimalNumberWithMantissa:exponent:isNegative:"), mantissa, exponent, flag)
	return rv
}

// Creates and returns a decimal number equivalent to the number specified by the arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1578294-decimalnumberwithmantissa?language=objc
func DecimalNumber_DecimalNumberWithMantissaExponentIsNegative(mantissa int64, exponent int, flag bool) DecimalNumber {
	return DecimalNumberClass.DecimalNumberWithMantissaExponentIsNegative(mantissa, exponent, flag)
}

// Creates a decimal number whose value is equivalent to that in a given numeric string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1578292-decimalnumberwithstring?language=objc
func (dc _DecimalNumberClass) DecimalNumberWithString(numberValue string) DecimalNumber {
	rv := objc.Call[DecimalNumber](dc, objc.Sel("decimalNumberWithString:"), numberValue)
	return rv
}

// Creates a decimal number whose value is equivalent to that in a given numeric string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1578292-decimalnumberwithstring?language=objc
func DecimalNumber_DecimalNumberWithString(numberValue string) DecimalNumber {
	return DecimalNumberClass.DecimalNumberWithString(numberValue)
}

// Multiplies the number by 10 raised to the given power using the specified behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1414279-decimalnumberbymultiplyingbypowe?language=objc
func (d_ DecimalNumber) DecimalNumberByMultiplyingByPowerOf10WithBehavior(power int, behavior PDecimalNumberBehaviors) DecimalNumber {
	po1 := objc.WrapAsProtocol("NSDecimalNumberBehaviors", behavior)
	rv := objc.Call[DecimalNumber](d_, objc.Sel("decimalNumberByMultiplyingByPowerOf10:withBehavior:"), power, po1)
	return rv
}

// Multiplies the number by 10 raised to the given power using the specified behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1414279-decimalnumberbymultiplyingbypowe?language=objc
func (d_ DecimalNumber) DecimalNumberByMultiplyingByPowerOf10WithBehaviorObject(power int, behaviorObject objc.IObject) DecimalNumber {
	rv := objc.Call[DecimalNumber](d_, objc.Sel("decimalNumberByMultiplyingByPowerOf10:withBehavior:"), power, objc.Ptr(behaviorObject))
	return rv
}

// Adds this number to another given number using the specified behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1407456-decimalnumberbyadding?language=objc
func (d_ DecimalNumber) DecimalNumberByAddingWithBehavior(decimalNumber IDecimalNumber, behavior PDecimalNumberBehaviors) DecimalNumber {
	po1 := objc.WrapAsProtocol("NSDecimalNumberBehaviors", behavior)
	rv := objc.Call[DecimalNumber](d_, objc.Sel("decimalNumberByAdding:withBehavior:"), objc.Ptr(decimalNumber), po1)
	return rv
}

// Adds this number to another given number using the specified behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1407456-decimalnumberbyadding?language=objc
func (d_ DecimalNumber) DecimalNumberByAddingWithBehaviorObject(decimalNumber IDecimalNumber, behaviorObject objc.IObject) DecimalNumber {
	rv := objc.Call[DecimalNumber](d_, objc.Sel("decimalNumberByAdding:withBehavior:"), objc.Ptr(decimalNumber), objc.Ptr(behaviorObject))
	return rv
}

// Divides this number by another given number using the specified behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1415619-decimalnumberbydividingby?language=objc
func (d_ DecimalNumber) DecimalNumberByDividingByWithBehavior(decimalNumber IDecimalNumber, behavior PDecimalNumberBehaviors) DecimalNumber {
	po1 := objc.WrapAsProtocol("NSDecimalNumberBehaviors", behavior)
	rv := objc.Call[DecimalNumber](d_, objc.Sel("decimalNumberByDividingBy:withBehavior:"), objc.Ptr(decimalNumber), po1)
	return rv
}

// Divides this number by another given number using the specified behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1415619-decimalnumberbydividingby?language=objc
func (d_ DecimalNumber) DecimalNumberByDividingByWithBehaviorObject(decimalNumber IDecimalNumber, behaviorObject objc.IObject) DecimalNumber {
	rv := objc.Call[DecimalNumber](d_, objc.Sel("decimalNumberByDividingBy:withBehavior:"), objc.Ptr(decimalNumber), objc.Ptr(behaviorObject))
	return rv
}

// Raises the number to a given power. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1413922-decimalnumberbyraisingtopower?language=objc
func (d_ DecimalNumber) DecimalNumberByRaisingToPower(power uint) DecimalNumber {
	rv := objc.Call[DecimalNumber](d_, objc.Sel("decimalNumberByRaisingToPower:"), power)
	return rv
}

// Returns the largest possible value of a decimal number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1415841-maximumdecimalnumber?language=objc
func (dc _DecimalNumberClass) MaximumDecimalNumber() DecimalNumber {
	rv := objc.Call[DecimalNumber](dc, objc.Sel("maximumDecimalNumber"))
	return rv
}

// Returns the largest possible value of a decimal number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1415841-maximumdecimalnumber?language=objc
func DecimalNumber_MaximumDecimalNumber() DecimalNumber {
	return DecimalNumberClass.MaximumDecimalNumber()
}

// A decimal number equivalent to the number 0.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1413127-zero?language=objc
func (dc _DecimalNumberClass) Zero() DecimalNumber {
	rv := objc.Call[DecimalNumber](dc, objc.Sel("zero"))
	return rv
}

// A decimal number equivalent to the number 0.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1413127-zero?language=objc
func DecimalNumber_Zero() DecimalNumber {
	return DecimalNumberClass.Zero()
}

// A decimal number equivalent to the number 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1415711-one?language=objc
func (dc _DecimalNumberClass) One() DecimalNumber {
	rv := objc.Call[DecimalNumber](dc, objc.Sel("one"))
	return rv
}

// A decimal number equivalent to the number 1.0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1415711-one?language=objc
func DecimalNumber_One() DecimalNumber {
	return DecimalNumberClass.One()
}

// A decimal number that specifies no number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1413389-notanumber?language=objc
func (dc _DecimalNumberClass) NotANumber() DecimalNumber {
	rv := objc.Call[DecimalNumber](dc, objc.Sel("notANumber"))
	return rv
}

// A decimal number that specifies no number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1413389-notanumber?language=objc
func DecimalNumber_NotANumber() DecimalNumber {
	return DecimalNumberClass.NotANumber()
}

// Returns the smallest possible value of a decimal number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1413371-minimumdecimalnumber?language=objc
func (dc _DecimalNumberClass) MinimumDecimalNumber() DecimalNumber {
	rv := objc.Call[DecimalNumber](dc, objc.Sel("minimumDecimalNumber"))
	return rv
}

// Returns the smallest possible value of a decimal number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1413371-minimumdecimalnumber?language=objc
func DecimalNumber_MinimumDecimalNumber() DecimalNumber {
	return DecimalNumberClass.MinimumDecimalNumber()
}

// The way arithmetic methods round off and handle error conditions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1418084-defaultbehavior?language=objc
func (dc _DecimalNumberClass) DefaultBehavior() DecimalNumberBehaviorsWrapper {
	rv := objc.Call[DecimalNumberBehaviorsWrapper](dc, objc.Sel("defaultBehavior"))
	return rv
}

// The way arithmetic methods round off and handle error conditions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1418084-defaultbehavior?language=objc
func DecimalNumber_DefaultBehavior() DecimalNumberBehaviorsWrapper {
	return DecimalNumberClass.DefaultBehavior()
}

// The way arithmetic methods round off and handle error conditions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1418084-defaultbehavior?language=objc
func (dc _DecimalNumberClass) SetDefaultBehavior(value PDecimalNumberBehaviors) {
	po0 := objc.WrapAsProtocol("NSDecimalNumberBehaviors", value)
	objc.Call[objc.Void](dc, objc.Sel("setDefaultBehavior:"), po0)
}

// The way arithmetic methods round off and handle error conditions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/1418084-defaultbehavior?language=objc
func DecimalNumber_SetDefaultBehavior(value PDecimalNumberBehaviors) {
	DecimalNumberClass.SetDefaultBehavior(value)
}
