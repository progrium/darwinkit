// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var DecimalNumberClass = _DecimalNumberClass{objc.GetClass("NSDecimalNumber")}

type _DecimalNumberClass struct {
	objc.Class
}

type IDecimalNumber interface {
	INumber
	DecimalNumberByAdding(decimalNumber IDecimalNumber) DecimalNumber
	DecimalNumberBySubtracting(decimalNumber IDecimalNumber) DecimalNumber
	DecimalNumberByMultiplyingBy(decimalNumber IDecimalNumber) DecimalNumber
	DecimalNumberByDividingBy(decimalNumber IDecimalNumber) DecimalNumber
	DecimalNumberByRaisingToPower(power uint) DecimalNumber
	DecimalNumberByMultiplyingByPowerOf10(power int16) DecimalNumber
}

type DecimalNumber struct {
	Number
}

func MakeDecimalNumber(ptr unsafe.Pointer) DecimalNumber {
	return DecimalNumber{
		Number: MakeNumber(ptr),
	}
}

func (d_ DecimalNumber) InitWithDecimal(dcm Decimal) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.GetSelector("initWithDecimal:"), dcm)
	return rv
}

func DecimalNumber_InitWithDecimal(dcm Decimal) DecimalNumber {
	return DecimalNumberClass.Alloc().InitWithDecimal(dcm)
}

func (d_ DecimalNumber) InitWithMantissaExponentIsNegative(mantissa uint64, exponent int16, flag bool) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.GetSelector("initWithMantissa:exponent:isNegative:"), mantissa, exponent, flag)
	return rv
}

func DecimalNumber_InitWithMantissaExponentIsNegative(mantissa uint64, exponent int16, flag bool) DecimalNumber {
	return DecimalNumberClass.Alloc().InitWithMantissaExponentIsNegative(mantissa, exponent, flag)
}

func (d_ DecimalNumber) InitWithString(numberValue string) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.GetSelector("initWithString:"), numberValue)
	return rv
}

func DecimalNumber_InitWithString(numberValue string) DecimalNumber {
	return DecimalNumberClass.Alloc().InitWithString(numberValue)
}

func (d_ DecimalNumber) InitWithStringLocale(numberValue string, locale objc.IObject) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.GetSelector("initWithString:locale:"), numberValue, objc.ExtractPtr(locale))
	return rv
}

func DecimalNumber_InitWithStringLocale(numberValue string, locale objc.IObject) DecimalNumber {
	return DecimalNumberClass.Alloc().InitWithStringLocale(numberValue, locale)
}

func (d_ DecimalNumber) InitWithBytesObjCType(value unsafe.Pointer, type_ *byte) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.GetSelector("initWithBytes:objCType:"), value, type_)
	return rv
}

func DecimalNumber_InitWithBytesObjCType(value unsafe.Pointer, type_ *byte) DecimalNumber {
	return DecimalNumberClass.Alloc().InitWithBytesObjCType(value, type_)
}

func (dc _DecimalNumberClass) Alloc() DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](dc, objc.GetSelector("alloc"))
	return rv
}

func DecimalNumber_Alloc() DecimalNumber {
	return DecimalNumberClass.Alloc()
}

func (dc _DecimalNumberClass) New() DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](dc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewDecimalNumber() DecimalNumber {
	return DecimalNumberClass.New()
}

func DecimalNumber_New() DecimalNumber {
	return DecimalNumberClass.New()
}

func (d_ DecimalNumber) Init() DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.GetSelector("init"))
	return rv
}

func DecimalNumber_Init() DecimalNumber {
	return DecimalNumberClass.Alloc().Init()
}

func (dc _DecimalNumberClass) DecimalNumberWithDecimal(dcm Decimal) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](dc, objc.GetSelector("decimalNumberWithDecimal:"), dcm)
	return rv
}

func DecimalNumber_DecimalNumberWithDecimal(dcm Decimal) DecimalNumber {
	return DecimalNumberClass.DecimalNumberWithDecimal(dcm)
}

func (dc _DecimalNumberClass) DecimalNumberWithMantissaExponentIsNegative(mantissa uint64, exponent int16, flag bool) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](dc, objc.GetSelector("decimalNumberWithMantissa:exponent:isNegative:"), mantissa, exponent, flag)
	return rv
}

func DecimalNumber_DecimalNumberWithMantissaExponentIsNegative(mantissa uint64, exponent int16, flag bool) DecimalNumber {
	return DecimalNumberClass.DecimalNumberWithMantissaExponentIsNegative(mantissa, exponent, flag)
}

func (dc _DecimalNumberClass) DecimalNumberWithString(numberValue string) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](dc, objc.GetSelector("decimalNumberWithString:"), numberValue)
	return rv
}

func DecimalNumber_DecimalNumberWithString(numberValue string) DecimalNumber {
	return DecimalNumberClass.DecimalNumberWithString(numberValue)
}

func (dc _DecimalNumberClass) DecimalNumberWithStringLocale(numberValue string, locale objc.IObject) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](dc, objc.GetSelector("decimalNumberWithString:locale:"), numberValue, objc.ExtractPtr(locale))
	return rv
}

func DecimalNumber_DecimalNumberWithStringLocale(numberValue string, locale objc.IObject) DecimalNumber {
	return DecimalNumberClass.DecimalNumberWithStringLocale(numberValue, locale)
}

func (d_ DecimalNumber) DecimalNumberByAdding(decimalNumber IDecimalNumber) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.GetSelector("decimalNumberByAdding:"), objc.ExtractPtr(decimalNumber))
	return rv
}

func (d_ DecimalNumber) DecimalNumberBySubtracting(decimalNumber IDecimalNumber) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.GetSelector("decimalNumberBySubtracting:"), objc.ExtractPtr(decimalNumber))
	return rv
}

func (d_ DecimalNumber) DecimalNumberByMultiplyingBy(decimalNumber IDecimalNumber) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.GetSelector("decimalNumberByMultiplyingBy:"), objc.ExtractPtr(decimalNumber))
	return rv
}

func (d_ DecimalNumber) DecimalNumberByDividingBy(decimalNumber IDecimalNumber) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.GetSelector("decimalNumberByDividingBy:"), objc.ExtractPtr(decimalNumber))
	return rv
}

func (d_ DecimalNumber) DecimalNumberByRaisingToPower(power uint) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.GetSelector("decimalNumberByRaisingToPower:"), power)
	return rv
}

func (d_ DecimalNumber) DecimalNumberByMultiplyingByPowerOf10(power int16) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.GetSelector("decimalNumberByMultiplyingByPowerOf10:"), power)
	return rv
}

func (dc _DecimalNumberClass) One() DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](dc, objc.GetSelector("one"))
	return rv
}

func DecimalNumber_One() DecimalNumber {
	return DecimalNumberClass.One()
}

func (dc _DecimalNumberClass) Zero() DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](dc, objc.GetSelector("zero"))
	return rv
}

func DecimalNumber_Zero() DecimalNumber {
	return DecimalNumberClass.Zero()
}

func (dc _DecimalNumberClass) NotANumber() DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](dc, objc.GetSelector("notANumber"))
	return rv
}

func DecimalNumber_NotANumber() DecimalNumber {
	return DecimalNumberClass.NotANumber()
}

func (dc _DecimalNumberClass) MaximumDecimalNumber() DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](dc, objc.GetSelector("maximumDecimalNumber"))
	return rv
}

func DecimalNumber_MaximumDecimalNumber() DecimalNumber {
	return DecimalNumberClass.MaximumDecimalNumber()
}

func (dc _DecimalNumberClass) MinimumDecimalNumber() DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](dc, objc.GetSelector("minimumDecimalNumber"))
	return rv
}

func DecimalNumber_MinimumDecimalNumber() DecimalNumber {
	return DecimalNumberClass.MinimumDecimalNumber()
}
