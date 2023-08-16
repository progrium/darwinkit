// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Number] class.
var NumberClass = _NumberClass{objc.GetClass("NSNumber")}

type _NumberClass struct {
	objc.Class
}

// An interface definition for the [Number] class.
type INumber interface {
	IValue
	InitWithUnsignedLong(value int32) Number
	InitWithInteger(value int) Number
	InitWithUnsignedShort(value int) Number
	InitWithLong(value int32) Number
	InitWithUnsignedLongLong(value int64) Number
	InitWithInt(value int) Number
	InitWithLongLong(value int64) Number
	InitWithShort(value int) Number
	InitWithChar(value uint8) Number
	InitWithFloat(value float64) Number
	InitWithDouble(value float64) Number
	DescriptionWithLocale(locale objc.IObject) string
	InitWithUnsignedChar(value uint8) Number
	Compare(otherNumber INumber) ComparisonResult
	InitWithBool(value bool) Number
	InitWithUnsignedInteger(value uint) Number
	InitWithUnsignedInt(value int) Number
	IsEqualToNumber(number INumber) bool
	IntegerValue() int
	UnsignedLongLongValue() int64
	IntValue() int
	ShortValue() int
	UnsignedIntValue() int
	StringValue() string
	LongLongValue() int64
	UnsignedShortValue() int
	DoubleValue() float64
	UnsignedLongValue() int32
	CharValue() uint8
	DecimalValue() Decimal
	BoolValue() bool
	FloatValue() float64
	UnsignedIntegerValue() uint
	UnsignedCharValue() uint8
	LongValue() int32
}

// An object wrapper for primitive scalar numeric values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber?language=objc
type Number struct {
	Value
}

func NumberFrom(ptr unsafe.Pointer) Number {
	return Number{
		Value: ValueFrom(ptr),
	}
}

func (nc _NumberClass) Alloc() Number {
	rv := objc.Call[Number](nc, objc.Sel("alloc"))
	return rv
}

func Number_Alloc() Number {
	return NumberClass.Alloc()
}

func (nc _NumberClass) New() Number {
	rv := objc.Call[Number](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNumber() Number {
	return NumberClass.New()
}

func (n_ Number) Init() Number {
	rv := objc.Call[Number](n_, objc.Sel("init"))
	return rv
}

func (n_ Number) InitWithBytesObjCType(value unsafe.Pointer, type_ *uint8) Number {
	rv := objc.Call[Number](n_, objc.Sel("initWithBytes:objCType:"), value, type_)
	return rv
}

// Initializes a value object to contain the specified value, interpreted with the specified Objective-C type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1411621-initwithbytes?language=objc
func Number_InitWithBytesObjCType(value unsafe.Pointer, type_ *uint8) Number {
	return NumberClass.Alloc().InitWithBytesObjCType(value, type_)
}

// Returns an NSNumber object initialized to contain a given value, treated as an unsigned long. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1410728-initwithunsignedlong?language=objc
func (n_ Number) InitWithUnsignedLong(value int32) Number {
	rv := objc.Call[Number](n_, objc.Sel("initWithUnsignedLong:"), value)
	return rv
}

// Creates and returns an NSNumber object containing a given value, treating it as a signed char. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1551464-numberwithchar?language=objc
func (nc _NumberClass) NumberWithChar(value uint8) Number {
	rv := objc.Call[Number](nc, objc.Sel("numberWithChar:"), value)
	return rv
}

// Creates and returns an NSNumber object containing a given value, treating it as a signed char. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1551464-numberwithchar?language=objc
func Number_NumberWithChar(value uint8) Number {
	return NumberClass.NumberWithChar(value)
}

// Returns an NSNumber object initialized to contain a given value, treated as an NSInteger. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1409397-initwithinteger?language=objc
func (n_ Number) InitWithInteger(value int) Number {
	rv := objc.Call[Number](n_, objc.Sel("initWithInteger:"), value)
	return rv
}

// Creates and returns an NSNumber object containing a given value, treating it as a double. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1551463-numberwithdouble?language=objc
func (nc _NumberClass) NumberWithDouble(value float64) Number {
	rv := objc.Call[Number](nc, objc.Sel("numberWithDouble:"), value)
	return rv
}

// Creates and returns an NSNumber object containing a given value, treating it as a double. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1551463-numberwithdouble?language=objc
func Number_NumberWithDouble(value float64) Number {
	return NumberClass.NumberWithDouble(value)
}

// Creates and returns an NSNumber object containing a given value, treating it as an unsigned long long. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1551465-numberwithunsignedlonglong?language=objc
func (nc _NumberClass) NumberWithUnsignedLongLong(value int64) Number {
	rv := objc.Call[Number](nc, objc.Sel("numberWithUnsignedLongLong:"), value)
	return rv
}

// Creates and returns an NSNumber object containing a given value, treating it as an unsigned long long. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1551465-numberwithunsignedlonglong?language=objc
func Number_NumberWithUnsignedLongLong(value int64) Number {
	return NumberClass.NumberWithUnsignedLongLong(value)
}

// Creates and returns an NSNumber object containing a given value, treating it as an NSInteger. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1551473-numberwithinteger?language=objc
func (nc _NumberClass) NumberWithInteger(value int) Number {
	rv := objc.Call[Number](nc, objc.Sel("numberWithInteger:"), value)
	return rv
}

// Creates and returns an NSNumber object containing a given value, treating it as an NSInteger. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1551473-numberwithinteger?language=objc
func Number_NumberWithInteger(value int) Number {
	return NumberClass.NumberWithInteger(value)
}

// Returns an NSNumber object initialized to contain a given value, treated as an unsigned short. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1407718-initwithunsignedshort?language=objc
func (n_ Number) InitWithUnsignedShort(value int) Number {
	rv := objc.Call[Number](n_, objc.Sel("initWithUnsignedShort:"), value)
	return rv
}

// Creates and returns an NSNumber object containing value, treating it as a signed short. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1551476-numberwithshort?language=objc
func (nc _NumberClass) NumberWithShort(value int) Number {
	rv := objc.Call[Number](nc, objc.Sel("numberWithShort:"), value)
	return rv
}

// Creates and returns an NSNumber object containing value, treating it as a signed short. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1551476-numberwithshort?language=objc
func Number_NumberWithShort(value int) Number {
	return NumberClass.NumberWithShort(value)
}

// Creates and returns an NSNumber object containing a given value, treating it as an unsigned long. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1551477-numberwithunsignedlong?language=objc
func (nc _NumberClass) NumberWithUnsignedLong(value int32) Number {
	rv := objc.Call[Number](nc, objc.Sel("numberWithUnsignedLong:"), value)
	return rv
}

// Creates and returns an NSNumber object containing a given value, treating it as an unsigned long. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1551477-numberwithunsignedlong?language=objc
func Number_NumberWithUnsignedLong(value int32) Number {
	return NumberClass.NumberWithUnsignedLong(value)
}

// Returns an NSNumber object initialized to contain a given value, treated as a signed long. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1410739-initwithlong?language=objc
func (n_ Number) InitWithLong(value int32) Number {
	rv := objc.Call[Number](n_, objc.Sel("initWithLong:"), value)
	return rv
}

// Returns an NSNumber object initialized to contain a given value, treated as an unsigned long long. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1416550-initwithunsignedlonglong?language=objc
func (n_ Number) InitWithUnsignedLongLong(value int64) Number {
	rv := objc.Call[Number](n_, objc.Sel("initWithUnsignedLongLong:"), value)
	return rv
}

// Returns an NSNumber object initialized to contain a given value, treated as a signed int. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1407580-initwithint?language=objc
func (n_ Number) InitWithInt(value int) Number {
	rv := objc.Call[Number](n_, objc.Sel("initWithInt:"), value)
	return rv
}

// Creates and returns an NSNumber object containing a given value, treating it as a signed int. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1551470-numberwithint?language=objc
func (nc _NumberClass) NumberWithInt(value int) Number {
	rv := objc.Call[Number](nc, objc.Sel("numberWithInt:"), value)
	return rv
}

// Creates and returns an NSNumber object containing a given value, treating it as a signed int. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1551470-numberwithint?language=objc
func Number_NumberWithInt(value int) Number {
	return NumberClass.NumberWithInt(value)
}

// Returns an NSNumber object initialized to contain value, treated as a signed long long. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1408171-initwithlonglong?language=objc
func (n_ Number) InitWithLongLong(value int64) Number {
	rv := objc.Call[Number](n_, objc.Sel("initWithLongLong:"), value)
	return rv
}

// Creates and returns an NSNumber object containing a given value, treating it as a float. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1551471-numberwithfloat?language=objc
func (nc _NumberClass) NumberWithFloat(value float64) Number {
	rv := objc.Call[Number](nc, objc.Sel("numberWithFloat:"), value)
	return rv
}

// Creates and returns an NSNumber object containing a given value, treating it as a float. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1551471-numberwithfloat?language=objc
func Number_NumberWithFloat(value float64) Number {
	return NumberClass.NumberWithFloat(value)
}

// Creates and returns an NSNumber object containing a given value, treating it as an unsigned short. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1551467-numberwithunsignedshort?language=objc
func (nc _NumberClass) NumberWithUnsignedShort(value int) Number {
	rv := objc.Call[Number](nc, objc.Sel("numberWithUnsignedShort:"), value)
	return rv
}

// Creates and returns an NSNumber object containing a given value, treating it as an unsigned short. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1551467-numberwithunsignedshort?language=objc
func Number_NumberWithUnsignedShort(value int) Number {
	return NumberClass.NumberWithUnsignedShort(value)
}

// Creates and returns an NSNumber object containing a given value, treating it as a BOOL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1551475-numberwithbool?language=objc
func (nc _NumberClass) NumberWithBool(value bool) Number {
	rv := objc.Call[Number](nc, objc.Sel("numberWithBool:"), value)
	return rv
}

// Creates and returns an NSNumber object containing a given value, treating it as a BOOL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1551475-numberwithbool?language=objc
func Number_NumberWithBool(value bool) Number {
	return NumberClass.NumberWithBool(value)
}

// Returns an NSNumber object initialized to contain a given value, treated as a signed short. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1411886-initwithshort?language=objc
func (n_ Number) InitWithShort(value int) Number {
	rv := objc.Call[Number](n_, objc.Sel("initWithShort:"), value)
	return rv
}

// Creates and returns an NSNumber object containing a given value, treating it as an unsigned char. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1551468-numberwithunsignedchar?language=objc
func (nc _NumberClass) NumberWithUnsignedChar(value uint8) Number {
	rv := objc.Call[Number](nc, objc.Sel("numberWithUnsignedChar:"), value)
	return rv
}

// Creates and returns an NSNumber object containing a given value, treating it as an unsigned char. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1551468-numberwithunsignedchar?language=objc
func Number_NumberWithUnsignedChar(value uint8) Number {
	return NumberClass.NumberWithUnsignedChar(value)
}

// Returns an NSNumber object initialized to contain a given value, treated as a signed char. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1409777-initwithchar?language=objc
func (n_ Number) InitWithChar(value uint8) Number {
	rv := objc.Call[Number](n_, objc.Sel("initWithChar:"), value)
	return rv
}

// Creates and returns an NSNumber object containing a given value, treating it as an unsigned int. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1551472-numberwithunsignedint?language=objc
func (nc _NumberClass) NumberWithUnsignedInt(value int) Number {
	rv := objc.Call[Number](nc, objc.Sel("numberWithUnsignedInt:"), value)
	return rv
}

// Creates and returns an NSNumber object containing a given value, treating it as an unsigned int. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1551472-numberwithunsignedint?language=objc
func Number_NumberWithUnsignedInt(value int) Number {
	return NumberClass.NumberWithUnsignedInt(value)
}

// Returns an NSNumber object initialized to contain a given value, treated as a float. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1412999-initwithfloat?language=objc
func (n_ Number) InitWithFloat(value float64) Number {
	rv := objc.Call[Number](n_, objc.Sel("initWithFloat:"), value)
	return rv
}

// Returns an NSNumber object initialized to contain value, treated as a double. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1407545-initwithdouble?language=objc
func (n_ Number) InitWithDouble(value float64) Number {
	rv := objc.Call[Number](n_, objc.Sel("initWithDouble:"), value)
	return rv
}

// Creates and returns an NSNumber object containing a given value, treating it as a signed long long. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1551462-numberwithlonglong?language=objc
func (nc _NumberClass) NumberWithLongLong(value int64) Number {
	rv := objc.Call[Number](nc, objc.Sel("numberWithLongLong:"), value)
	return rv
}

// Creates and returns an NSNumber object containing a given value, treating it as a signed long long. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1551462-numberwithlonglong?language=objc
func Number_NumberWithLongLong(value int64) Number {
	return NumberClass.NumberWithLongLong(value)
}

// Creates and returns an NSNumber object containing a given value, treating it as an NSUInteger. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1551469-numberwithunsignedinteger?language=objc
func (nc _NumberClass) NumberWithUnsignedInteger(value uint) Number {
	rv := objc.Call[Number](nc, objc.Sel("numberWithUnsignedInteger:"), value)
	return rv
}

// Creates and returns an NSNumber object containing a given value, treating it as an NSUInteger. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1551469-numberwithunsignedinteger?language=objc
func Number_NumberWithUnsignedInteger(value uint) Number {
	return NumberClass.NumberWithUnsignedInteger(value)
}

// Returns a string that represents the contents of the number object for a given locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1409984-descriptionwithlocale?language=objc
func (n_ Number) DescriptionWithLocale(locale objc.IObject) string {
	rv := objc.Call[string](n_, objc.Sel("descriptionWithLocale:"), locale)
	return rv
}

// Creates and returns an NSNumber object containing a given value, treating it as a signed long. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1551474-numberwithlong?language=objc
func (nc _NumberClass) NumberWithLong(value int32) Number {
	rv := objc.Call[Number](nc, objc.Sel("numberWithLong:"), value)
	return rv
}

// Creates and returns an NSNumber object containing a given value, treating it as a signed long. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1551474-numberwithlong?language=objc
func Number_NumberWithLong(value int32) Number {
	return NumberClass.NumberWithLong(value)
}

// Returns an NSNumber object initialized to contain a given value, treated as an unsigned char. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1416533-initwithunsignedchar?language=objc
func (n_ Number) InitWithUnsignedChar(value uint8) Number {
	rv := objc.Call[Number](n_, objc.Sel("initWithUnsignedChar:"), value)
	return rv
}

// Returns an NSComparisonResult value that indicates whether the number object’s value is greater than, equal to, or less than a given number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1413562-compare?language=objc
func (n_ Number) Compare(otherNumber INumber) ComparisonResult {
	rv := objc.Call[ComparisonResult](n_, objc.Sel("compare:"), objc.Ptr(otherNumber))
	return rv
}

// Returns an NSNumber object initialized to contain a given value, treated as a BOOL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1415728-initwithbool?language=objc
func (n_ Number) InitWithBool(value bool) Number {
	rv := objc.Call[Number](n_, objc.Sel("initWithBool:"), value)
	return rv
}

// Returns an NSNumber object initialized to contain a given value, treated as an NSUInteger. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1412531-initwithunsignedinteger?language=objc
func (n_ Number) InitWithUnsignedInteger(value uint) Number {
	rv := objc.Call[Number](n_, objc.Sel("initWithUnsignedInteger:"), value)
	return rv
}

// Returns an NSNumber object initialized to contain a given value, treated as an unsigned int. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1414598-initwithunsignedint?language=objc
func (n_ Number) InitWithUnsignedInt(value int) Number {
	rv := objc.Call[Number](n_, objc.Sel("initWithUnsignedInt:"), value)
	return rv
}

// Returns a Boolean value that indicates whether the number object’s value and a given number are equal. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1411315-isequaltonumber?language=objc
func (n_ Number) IsEqualToNumber(number INumber) bool {
	rv := objc.Call[bool](n_, objc.Sel("isEqualToNumber:"), objc.Ptr(number))
	return rv
}

// The number object's value expressed as an NSInteger object, converted as necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1412554-integervalue?language=objc
func (n_ Number) IntegerValue() int {
	rv := objc.Call[int](n_, objc.Sel("integerValue"))
	return rv
}

// The number object’s value expressed as an unsigned long long, converted as necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1414524-unsignedlonglongvalue?language=objc
func (n_ Number) UnsignedLongLongValue() int64 {
	rv := objc.Call[int64](n_, objc.Sel("unsignedLongLongValue"))
	return rv
}

// The number object's value expressed as an int, converted as necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1407153-intvalue?language=objc
func (n_ Number) IntValue() int {
	rv := objc.Call[int](n_, objc.Sel("intValue"))
	return rv
}

// The number object's value expressed as a short, converted as necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1407601-shortvalue?language=objc
func (n_ Number) ShortValue() int {
	rv := objc.Call[int](n_, objc.Sel("shortValue"))
	return rv
}

// The number object's value expressed as an unsigned int, converted as necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1417875-unsignedintvalue?language=objc
func (n_ Number) UnsignedIntValue() int {
	rv := objc.Call[int](n_, objc.Sel("unsignedIntValue"))
	return rv
}

// The number object's value expressed as a human-readable string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1415802-stringvalue?language=objc
func (n_ Number) StringValue() string {
	rv := objc.Call[string](n_, objc.Sel("stringValue"))
	return rv
}

// The number object’s value expressed as a long long, converted as necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1416870-longlongvalue?language=objc
func (n_ Number) LongLongValue() int64 {
	rv := objc.Call[int64](n_, objc.Sel("longLongValue"))
	return rv
}

// The number object's value expressed as an unsigned short, converted as necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1410604-unsignedshortvalue?language=objc
func (n_ Number) UnsignedShortValue() int {
	rv := objc.Call[int](n_, objc.Sel("unsignedShortValue"))
	return rv
}

// The number object's value expressed as a double, converted as necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1414104-doublevalue?language=objc
func (n_ Number) DoubleValue() float64 {
	rv := objc.Call[float64](n_, objc.Sel("doubleValue"))
	return rv
}

// The number object's value expressed as an unsigned long, converted as necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1415252-unsignedlongvalue?language=objc
func (n_ Number) UnsignedLongValue() int32 {
	rv := objc.Call[int32](n_, objc.Sel("unsignedLongValue"))
	return rv
}

// The number object's value expressed as a char. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1407838-charvalue?language=objc
func (n_ Number) CharValue() uint8 {
	rv := objc.Call[uint8](n_, objc.Sel("charValue"))
	return rv
}

// The number object's value expressed as an NSDecimal structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1407409-decimalvalue?language=objc
func (n_ Number) DecimalValue() Decimal {
	rv := objc.Call[Decimal](n_, objc.Sel("decimalValue"))
	return rv
}

// The number object's value expressed as a Boolean value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1410865-boolvalue?language=objc
func (n_ Number) BoolValue() bool {
	rv := objc.Call[bool](n_, objc.Sel("boolValue"))
	return rv
}

// The number object's value expressed as a float, converted as necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1418317-floatvalue?language=objc
func (n_ Number) FloatValue() float64 {
	rv := objc.Call[float64](n_, objc.Sel("floatValue"))
	return rv
}

// The number object's value expressed as an NSUInteger object, converted as necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1413324-unsignedintegervalue?language=objc
func (n_ Number) UnsignedIntegerValue() uint {
	rv := objc.Call[uint](n_, objc.Sel("unsignedIntegerValue"))
	return rv
}

// The number object's value expressed as an unsigned char, converted as necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1409016-unsignedcharvalue?language=objc
func (n_ Number) UnsignedCharValue() uint8 {
	rv := objc.Call[uint8](n_, objc.Sel("unsignedCharValue"))
	return rv
}

// The number object's value expressed as a long, converted as necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/1412566-longvalue?language=objc
func (n_ Number) LongValue() int32 {
	rv := objc.Call[int32](n_, objc.Sel("longValue"))
	return rv
}
