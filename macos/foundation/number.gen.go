// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var NumberClass = _NumberClass{objc.GetClass("NSNumber")}

type _NumberClass struct {
	objc.Class
}

type INumber interface {
	IValue
	InitWithBool(value bool) Number
	InitWithChar(value byte) Number
	InitWithDouble(value float64) Number
	InitWithFloat(value float32) Number
	InitWithInt(value int32) Number
	InitWithInteger(value int) Number
	InitWithLong(value int64) Number
	InitWithLongLong(value int64) Number
	InitWithShort(value int16) Number
	InitWithUnsignedChar(value byte) Number
	InitWithUnsignedInt(value uint32) Number
	InitWithUnsignedInteger(value uint) Number
	InitWithUnsignedLong(value uint64) Number
	InitWithUnsignedLongLong(value uint64) Number
	InitWithUnsignedShort(value uint16) Number
	DescriptionWithLocale(locale objc.IObject) string
	Compare(otherNumber INumber) ComparisonResult
	IsEqualToNumber(number INumber) bool
	BoolValue() bool
	CharValue() byte
	DecimalValue() Decimal
	DoubleValue() float64
	FloatValue() float32
	IntValue() int32
	IntegerValue() int
	LongLongValue() int64
	LongValue() int64
	ShortValue() int16
	UnsignedCharValue() byte
	UnsignedIntegerValue() uint
	UnsignedIntValue() uint32
	UnsignedLongLongValue() uint64
	UnsignedLongValue() uint64
	UnsignedShortValue() uint16
	StringValue() string
}

type Number struct {
	Value
}

func MakeNumber(ptr unsafe.Pointer) Number {
	return Number{
		Value: MakeValue(ptr),
	}
}

func (n_ Number) InitWithBytesObjCType(value unsafe.Pointer, type_ *byte) Number {
	rv := objc.CallMethod[Number](n_, objc.GetSelector("initWithBytes:objCType:"), value, type_)
	return rv
}

func Number_InitWithBytesObjCType(value unsafe.Pointer, type_ *byte) Number {
	return NumberClass.Alloc().InitWithBytesObjCType(value, type_)
}

func (nc _NumberClass) Alloc() Number {
	rv := objc.CallMethod[Number](nc, objc.GetSelector("alloc"))
	return rv
}

func Number_Alloc() Number {
	return NumberClass.Alloc()
}

func (nc _NumberClass) New() Number {
	rv := objc.CallMethod[Number](nc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewNumber() Number {
	return NumberClass.New()
}

func Number_New() Number {
	return NumberClass.New()
}

func (n_ Number) Init() Number {
	rv := objc.CallMethod[Number](n_, objc.GetSelector("init"))
	return rv
}

func Number_Init() Number {
	return NumberClass.Alloc().Init()
}

func (nc _NumberClass) NumberWithBool(value bool) Number {
	rv := objc.CallMethod[Number](nc, objc.GetSelector("numberWithBool:"), value)
	return rv
}

func Number_NumberWithBool(value bool) Number {
	return NumberClass.NumberWithBool(value)
}

func (nc _NumberClass) NumberWithChar(value byte) Number {
	rv := objc.CallMethod[Number](nc, objc.GetSelector("numberWithChar:"), value)
	return rv
}

func Number_NumberWithChar(value byte) Number {
	return NumberClass.NumberWithChar(value)
}

func (nc _NumberClass) NumberWithDouble(value float64) Number {
	rv := objc.CallMethod[Number](nc, objc.GetSelector("numberWithDouble:"), value)
	return rv
}

func Number_NumberWithDouble(value float64) Number {
	return NumberClass.NumberWithDouble(value)
}

func (nc _NumberClass) NumberWithFloat(value float32) Number {
	rv := objc.CallMethod[Number](nc, objc.GetSelector("numberWithFloat:"), value)
	return rv
}

func Number_NumberWithFloat(value float32) Number {
	return NumberClass.NumberWithFloat(value)
}

func (nc _NumberClass) NumberWithInt(value int32) Number {
	rv := objc.CallMethod[Number](nc, objc.GetSelector("numberWithInt:"), value)
	return rv
}

func Number_NumberWithInt(value int32) Number {
	return NumberClass.NumberWithInt(value)
}

func (nc _NumberClass) NumberWithInteger(value int) Number {
	rv := objc.CallMethod[Number](nc, objc.GetSelector("numberWithInteger:"), value)
	return rv
}

func Number_NumberWithInteger(value int) Number {
	return NumberClass.NumberWithInteger(value)
}

func (nc _NumberClass) NumberWithLong(value int64) Number {
	rv := objc.CallMethod[Number](nc, objc.GetSelector("numberWithLong:"), value)
	return rv
}

func Number_NumberWithLong(value int64) Number {
	return NumberClass.NumberWithLong(value)
}

func (nc _NumberClass) NumberWithLongLong(value int64) Number {
	rv := objc.CallMethod[Number](nc, objc.GetSelector("numberWithLongLong:"), value)
	return rv
}

func Number_NumberWithLongLong(value int64) Number {
	return NumberClass.NumberWithLongLong(value)
}

func (nc _NumberClass) NumberWithShort(value int16) Number {
	rv := objc.CallMethod[Number](nc, objc.GetSelector("numberWithShort:"), value)
	return rv
}

func Number_NumberWithShort(value int16) Number {
	return NumberClass.NumberWithShort(value)
}

func (nc _NumberClass) NumberWithUnsignedChar(value byte) Number {
	rv := objc.CallMethod[Number](nc, objc.GetSelector("numberWithUnsignedChar:"), value)
	return rv
}

func Number_NumberWithUnsignedChar(value byte) Number {
	return NumberClass.NumberWithUnsignedChar(value)
}

func (nc _NumberClass) NumberWithUnsignedInt(value uint32) Number {
	rv := objc.CallMethod[Number](nc, objc.GetSelector("numberWithUnsignedInt:"), value)
	return rv
}

func Number_NumberWithUnsignedInt(value uint32) Number {
	return NumberClass.NumberWithUnsignedInt(value)
}

func (nc _NumberClass) NumberWithUnsignedInteger(value uint) Number {
	rv := objc.CallMethod[Number](nc, objc.GetSelector("numberWithUnsignedInteger:"), value)
	return rv
}

func Number_NumberWithUnsignedInteger(value uint) Number {
	return NumberClass.NumberWithUnsignedInteger(value)
}

func (nc _NumberClass) NumberWithUnsignedLong(value uint64) Number {
	rv := objc.CallMethod[Number](nc, objc.GetSelector("numberWithUnsignedLong:"), value)
	return rv
}

func Number_NumberWithUnsignedLong(value uint64) Number {
	return NumberClass.NumberWithUnsignedLong(value)
}

func (nc _NumberClass) NumberWithUnsignedLongLong(value uint64) Number {
	rv := objc.CallMethod[Number](nc, objc.GetSelector("numberWithUnsignedLongLong:"), value)
	return rv
}

func Number_NumberWithUnsignedLongLong(value uint64) Number {
	return NumberClass.NumberWithUnsignedLongLong(value)
}

func (nc _NumberClass) NumberWithUnsignedShort(value uint16) Number {
	rv := objc.CallMethod[Number](nc, objc.GetSelector("numberWithUnsignedShort:"), value)
	return rv
}

func Number_NumberWithUnsignedShort(value uint16) Number {
	return NumberClass.NumberWithUnsignedShort(value)
}

func (n_ Number) InitWithBool(value bool) Number {
	rv := objc.CallMethod[Number](n_, objc.GetSelector("initWithBool:"), value)
	return rv
}

func (n_ Number) InitWithChar(value byte) Number {
	rv := objc.CallMethod[Number](n_, objc.GetSelector("initWithChar:"), value)
	return rv
}

func (n_ Number) InitWithDouble(value float64) Number {
	rv := objc.CallMethod[Number](n_, objc.GetSelector("initWithDouble:"), value)
	return rv
}

func (n_ Number) InitWithFloat(value float32) Number {
	rv := objc.CallMethod[Number](n_, objc.GetSelector("initWithFloat:"), value)
	return rv
}

func (n_ Number) InitWithInt(value int32) Number {
	rv := objc.CallMethod[Number](n_, objc.GetSelector("initWithInt:"), value)
	return rv
}

func (n_ Number) InitWithInteger(value int) Number {
	rv := objc.CallMethod[Number](n_, objc.GetSelector("initWithInteger:"), value)
	return rv
}

func (n_ Number) InitWithLong(value int64) Number {
	rv := objc.CallMethod[Number](n_, objc.GetSelector("initWithLong:"), value)
	return rv
}

func (n_ Number) InitWithLongLong(value int64) Number {
	rv := objc.CallMethod[Number](n_, objc.GetSelector("initWithLongLong:"), value)
	return rv
}

func (n_ Number) InitWithShort(value int16) Number {
	rv := objc.CallMethod[Number](n_, objc.GetSelector("initWithShort:"), value)
	return rv
}

func (n_ Number) InitWithUnsignedChar(value byte) Number {
	rv := objc.CallMethod[Number](n_, objc.GetSelector("initWithUnsignedChar:"), value)
	return rv
}

func (n_ Number) InitWithUnsignedInt(value uint32) Number {
	rv := objc.CallMethod[Number](n_, objc.GetSelector("initWithUnsignedInt:"), value)
	return rv
}

func (n_ Number) InitWithUnsignedInteger(value uint) Number {
	rv := objc.CallMethod[Number](n_, objc.GetSelector("initWithUnsignedInteger:"), value)
	return rv
}

func (n_ Number) InitWithUnsignedLong(value uint64) Number {
	rv := objc.CallMethod[Number](n_, objc.GetSelector("initWithUnsignedLong:"), value)
	return rv
}

func (n_ Number) InitWithUnsignedLongLong(value uint64) Number {
	rv := objc.CallMethod[Number](n_, objc.GetSelector("initWithUnsignedLongLong:"), value)
	return rv
}

func (n_ Number) InitWithUnsignedShort(value uint16) Number {
	rv := objc.CallMethod[Number](n_, objc.GetSelector("initWithUnsignedShort:"), value)
	return rv
}

func (n_ Number) DescriptionWithLocale(locale objc.IObject) string {
	rv := objc.CallMethod[string](n_, objc.GetSelector("descriptionWithLocale:"), objc.ExtractPtr(locale))
	return rv
}

func (n_ Number) Compare(otherNumber INumber) ComparisonResult {
	rv := objc.CallMethod[ComparisonResult](n_, objc.GetSelector("compare:"), objc.ExtractPtr(otherNumber))
	return rv
}

func (n_ Number) IsEqualToNumber(number INumber) bool {
	rv := objc.CallMethod[bool](n_, objc.GetSelector("isEqualToNumber:"), objc.ExtractPtr(number))
	return rv
}

func (n_ Number) BoolValue() bool {
	rv := objc.CallMethod[bool](n_, objc.GetSelector("boolValue"))
	return rv
}

func (n_ Number) CharValue() byte {
	rv := objc.CallMethod[byte](n_, objc.GetSelector("charValue"))
	return rv
}

func (n_ Number) DecimalValue() Decimal {
	rv := objc.CallMethod[Decimal](n_, objc.GetSelector("decimalValue"))
	return rv
}

func (n_ Number) DoubleValue() float64 {
	rv := objc.CallMethod[float64](n_, objc.GetSelector("doubleValue"))
	return rv
}

func (n_ Number) FloatValue() float32 {
	rv := objc.CallMethod[float32](n_, objc.GetSelector("floatValue"))
	return rv
}

func (n_ Number) IntValue() int32 {
	rv := objc.CallMethod[int32](n_, objc.GetSelector("intValue"))
	return rv
}

func (n_ Number) IntegerValue() int {
	rv := objc.CallMethod[int](n_, objc.GetSelector("integerValue"))
	return rv
}

func (n_ Number) LongLongValue() int64 {
	rv := objc.CallMethod[int64](n_, objc.GetSelector("longLongValue"))
	return rv
}

func (n_ Number) LongValue() int64 {
	rv := objc.CallMethod[int64](n_, objc.GetSelector("longValue"))
	return rv
}

func (n_ Number) ShortValue() int16 {
	rv := objc.CallMethod[int16](n_, objc.GetSelector("shortValue"))
	return rv
}

func (n_ Number) UnsignedCharValue() byte {
	rv := objc.CallMethod[byte](n_, objc.GetSelector("unsignedCharValue"))
	return rv
}

func (n_ Number) UnsignedIntegerValue() uint {
	rv := objc.CallMethod[uint](n_, objc.GetSelector("unsignedIntegerValue"))
	return rv
}

func (n_ Number) UnsignedIntValue() uint32 {
	rv := objc.CallMethod[uint32](n_, objc.GetSelector("unsignedIntValue"))
	return rv
}

func (n_ Number) UnsignedLongLongValue() uint64 {
	rv := objc.CallMethod[uint64](n_, objc.GetSelector("unsignedLongLongValue"))
	return rv
}

func (n_ Number) UnsignedLongValue() uint64 {
	rv := objc.CallMethod[uint64](n_, objc.GetSelector("unsignedLongValue"))
	return rv
}

func (n_ Number) UnsignedShortValue() uint16 {
	rv := objc.CallMethod[uint16](n_, objc.GetSelector("unsignedShortValue"))
	return rv
}

func (n_ Number) StringValue() string {
	rv := objc.CallMethod[string](n_, objc.GetSelector("stringValue"))
	return rv
}
