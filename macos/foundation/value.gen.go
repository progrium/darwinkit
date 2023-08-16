// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Value] class.
var ValueClass = _ValueClass{objc.GetClass("NSValue")}

type _ValueClass struct {
	objc.Class
}

// An interface definition for the [Value] class.
type IValue interface {
	objc.IObject
	GetValueSize(value unsafe.Pointer, size uint)
	IsEqualToValue(value IValue) bool
	ObjCType() *uint8
	CMTimeValue() coremedia.Time
	PointValue() Point
	PointerValue() unsafe.Pointer
	CMTimeRangeValue() coremedia.TimeRange
	SizeValue() Size
	EdgeInsetsValue() EdgeInsets
	RangeValue() Range
	CMTimeMappingValue() coremedia.TimeMapping
	RectValue() Rect
	NonretainedObjectValue() objc.Object
}

// A simple container for a single C or Objective-C data item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue?language=objc
type Value struct {
	objc.Object
}

func ValueFrom(ptr unsafe.Pointer) Value {
	return Value{
		Object: objc.ObjectFrom(ptr),
	}
}

func (v_ Value) InitWithBytesObjCType(value unsafe.Pointer, type_ *uint8) Value {
	rv := objc.Call[Value](v_, objc.Sel("initWithBytes:objCType:"), value, type_)
	return rv
}

// Initializes a value object to contain the specified value, interpreted with the specified Objective-C type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1411621-initwithbytes?language=objc
func Value_InitWithBytesObjCType(value unsafe.Pointer, type_ *uint8) Value {
	return ValueClass.Alloc().InitWithBytesObjCType(value, type_)
}

func (vc _ValueClass) Alloc() Value {
	rv := objc.Call[Value](vc, objc.Sel("alloc"))
	return rv
}

func Value_Alloc() Value {
	return ValueClass.Alloc()
}

func (vc _ValueClass) New() Value {
	rv := objc.Call[Value](vc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewValue() Value {
	return ValueClass.New()
}

func (v_ Value) Init() Value {
	rv := objc.Call[Value](v_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1391181-valuewithedgeinsets?language=objc
func (vc _ValueClass) ValueWithEdgeInsets(insets EdgeInsets) Value {
	rv := objc.Call[Value](vc, objc.Sel("valueWithEdgeInsets:"), insets)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1391181-valuewithedgeinsets?language=objc
func Value_ValueWithEdgeInsets(insets EdgeInsets) Value {
	return ValueClass.ValueWithEdgeInsets(insets)
}

// Creates a new value object containing the specified CoreMedia time structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1388561-valuewithcmtime?language=objc
func (vc _ValueClass) ValueWithCMTime(time coremedia.Time) Value {
	rv := objc.Call[Value](vc, objc.Sel("valueWithCMTime:"), time)
	return rv
}

// Creates a new value object containing the specified CoreMedia time structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1388561-valuewithcmtime?language=objc
func Value_ValueWithCMTime(time coremedia.Time) Value {
	return ValueClass.ValueWithCMTime(time)
}

// Creates a value object containing the specified value, interpreted with the specified Objective-C type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1417400-value?language=objc
func (vc _ValueClass) ValueWithObjCType(value unsafe.Pointer, type_ *uint8) Value {
	rv := objc.Call[Value](vc, objc.Sel("value:withObjCType:"), value, type_)
	return rv
}

// Creates a value object containing the specified value, interpreted with the specified Objective-C type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1417400-value?language=objc
func Value_ValueWithObjCType(value unsafe.Pointer, type_ *uint8) Value {
	return ValueClass.ValueWithObjCType(value, type_)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/2919632-getvalue?language=objc
func (v_ Value) GetValueSize(value unsafe.Pointer, size uint) {
	objc.Call[objc.Void](v_, objc.Sel("getValue:size:"), value, size)
}

// Creates a new value object containing the specified Foundation rectangle structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1391281-valuewithrect?language=objc
func (vc _ValueClass) ValueWithRect(rect Rect) Value {
	rv := objc.Call[Value](vc, objc.Sel("valueWithRect:"), rect)
	return rv
}

// Creates a new value object containing the specified Foundation rectangle structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1391281-valuewithrect?language=objc
func Value_ValueWithRect(rect Rect) Value {
	return ValueClass.ValueWithRect(rect)
}

// Returns a Boolean value that indicates whether the value object and another value object are equal. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1409038-isequaltovalue?language=objc
func (v_ Value) IsEqualToValue(value IValue) bool {
	rv := objc.Call[bool](v_, objc.Sel("isEqualToValue:"), objc.Ptr(value))
	return rv
}

// Creates a value object containing the specified object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1408098-valuewithnonretainedobject?language=objc
func (vc _ValueClass) ValueWithNonretainedObject(anObject objc.IObject) Value {
	rv := objc.Call[Value](vc, objc.Sel("valueWithNonretainedObject:"), anObject)
	return rv
}

// Creates a value object containing the specified object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1408098-valuewithnonretainedobject?language=objc
func Value_ValueWithNonretainedObject(anObject objc.IObject) Value {
	return ValueClass.ValueWithNonretainedObject(anObject)
}

// Creates a value object containing the specified value, interpreted with the specified Objective-C type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1551466-valuewithbytes?language=objc
func (vc _ValueClass) ValueWithBytesObjCType(value unsafe.Pointer, type_ *uint8) Value {
	rv := objc.Call[Value](vc, objc.Sel("valueWithBytes:objCType:"), value, type_)
	return rv
}

// Creates a value object containing the specified value, interpreted with the specified Objective-C type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1551466-valuewithbytes?language=objc
func Value_ValueWithBytesObjCType(value unsafe.Pointer, type_ *uint8) Value {
	return ValueClass.ValueWithBytesObjCType(value, type_)
}

// Creates a new value object containing the specified Foundation point structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1391106-valuewithpoint?language=objc
func (vc _ValueClass) ValueWithPoint(point Point) Value {
	rv := objc.Call[Value](vc, objc.Sel("valueWithPoint:"), point)
	return rv
}

// Creates a new value object containing the specified Foundation point structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1391106-valuewithpoint?language=objc
func Value_ValueWithPoint(point Point) Value {
	return ValueClass.ValueWithPoint(point)
}

// Creates a new value object containing the specified Foundation range structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1410315-valuewithrange?language=objc
func (vc _ValueClass) ValueWithRange(range_ Range) Value {
	rv := objc.Call[Value](vc, objc.Sel("valueWithRange:"), range_)
	return rv
}

// Creates a new value object containing the specified Foundation range structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1410315-valuewithrange?language=objc
func Value_ValueWithRange(range_ Range) Value {
	return ValueClass.ValueWithRange(range_)
}

// Creates a value object containing the specified pointer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1415975-valuewithpointer?language=objc
func (vc _ValueClass) ValueWithPointer(pointer unsafe.Pointer) Value {
	rv := objc.Call[Value](vc, objc.Sel("valueWithPointer:"), pointer)
	return rv
}

// Creates a value object containing the specified pointer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1415975-valuewithpointer?language=objc
func Value_ValueWithPointer(pointer unsafe.Pointer) Value {
	return ValueClass.ValueWithPointer(pointer)
}

// Creates a new value object containing the specified CoreMedia time range structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1386915-valuewithcmtimerange?language=objc
func (vc _ValueClass) ValueWithCMTimeRange(timeRange coremedia.TimeRange) Value {
	rv := objc.Call[Value](vc, objc.Sel("valueWithCMTimeRange:"), timeRange)
	return rv
}

// Creates a new value object containing the specified CoreMedia time range structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1386915-valuewithcmtimerange?language=objc
func Value_ValueWithCMTimeRange(timeRange coremedia.TimeRange) Value {
	return ValueClass.ValueWithCMTimeRange(timeRange)
}

// Creates a new value object containing the specified CoreMedia time mapping structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1387556-valuewithcmtimemapping?language=objc
func (vc _ValueClass) ValueWithCMTimeMapping(timeMapping coremedia.TimeMapping) Value {
	rv := objc.Call[Value](vc, objc.Sel("valueWithCMTimeMapping:"), timeMapping)
	return rv
}

// Creates a new value object containing the specified CoreMedia time mapping structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1387556-valuewithcmtimemapping?language=objc
func Value_ValueWithCMTimeMapping(timeMapping coremedia.TimeMapping) Value {
	return ValueClass.ValueWithCMTimeMapping(timeMapping)
}

// Creates a new value object containing the specified Foundation size structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1391199-valuewithsize?language=objc
func (vc _ValueClass) ValueWithSize(size Size) Value {
	rv := objc.Call[Value](vc, objc.Sel("valueWithSize:"), size)
	return rv
}

// Creates a new value object containing the specified Foundation size structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1391199-valuewithsize?language=objc
func Value_ValueWithSize(size Size) Value {
	return ValueClass.ValueWithSize(size)
}

// A C string containing the Objective-C type of the data contained in the value object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1412365-objctype?language=objc
func (v_ Value) ObjCType() *uint8 {
	rv := objc.Call[*uint8](v_, objc.Sel("objCType"))
	return rv
}

// The CoreMedia time structure representation of the value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1388151-cmtimevalue?language=objc
func (v_ Value) CMTimeValue() coremedia.Time {
	rv := objc.Call[coremedia.Time](v_, objc.Sel("CMTimeValue"))
	return rv
}

// The Foundation point structure representation of the value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1391255-pointvalue?language=objc
func (v_ Value) PointValue() Point {
	rv := objc.Call[Point](v_, objc.Sel("pointValue"))
	return rv
}

// Returns the value as an untyped pointer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1410668-pointervalue?language=objc
func (v_ Value) PointerValue() unsafe.Pointer {
	rv := objc.Call[unsafe.Pointer](v_, objc.Sel("pointerValue"))
	return rv
}

// The CoreMedia time range structure representation of the value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1385930-cmtimerangevalue?language=objc
func (v_ Value) CMTimeRangeValue() coremedia.TimeRange {
	rv := objc.Call[coremedia.TimeRange](v_, objc.Sel("CMTimeRangeValue"))
	return rv
}

// The Foundation size structure representation of the value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1391301-sizevalue?language=objc
func (v_ Value) SizeValue() Size {
	rv := objc.Call[Size](v_, objc.Sel("sizeValue"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1391123-edgeinsetsvalue?language=objc
func (v_ Value) EdgeInsetsValue() EdgeInsets {
	rv := objc.Call[EdgeInsets](v_, objc.Sel("edgeInsetsValue"))
	return rv
}

// The Foundation range structure representation of the value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1413902-rangevalue?language=objc
func (v_ Value) RangeValue() Range {
	rv := objc.Call[Range](v_, objc.Sel("rangeValue"))
	return rv
}

// The CoreMedia time mapping structure representation of the value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1387277-cmtimemappingvalue?language=objc
func (v_ Value) CMTimeMappingValue() coremedia.TimeMapping {
	rv := objc.Call[coremedia.TimeMapping](v_, objc.Sel("CMTimeMappingValue"))
	return rv
}

// The Foundation rectangle structure representation of the value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1391171-rectvalue?language=objc
func (v_ Value) RectValue() Rect {
	rv := objc.Call[Rect](v_, objc.Sel("rectValue"))
	return rv
}

// The value as a non-retained pointer to an object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/1412287-nonretainedobjectvalue?language=objc
func (v_ Value) NonretainedObjectValue() objc.Object {
	rv := objc.Call[objc.Object](v_, objc.Sel("nonretainedObjectValue"))
	return rv
}
