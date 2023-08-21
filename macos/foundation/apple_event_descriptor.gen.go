// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AppleEventDescriptor] class.
var AppleEventDescriptorClass = _AppleEventDescriptorClass{objc.GetClass("NSAppleEventDescriptor")}

type _AppleEventDescriptorClass struct {
	objc.Class
}

// An interface definition for the [AppleEventDescriptor] class.
type IAppleEventDescriptor interface {
	objc.IObject
	InsertDescriptorAtIndex(descriptor IAppleEventDescriptor, index int)
	RemoveDescriptorAtIndex(index int)
	DescriptorAtIndex(index int) AppleEventDescriptor
	SendEventWithOptionsTimeoutError(sendOptions AppleEventSendOptions, timeoutInSeconds TimeInterval, error IError) AppleEventDescriptor
	Int32Value() int32
	TypeCodeValue() uint
	StringValue() string
	EnumCodeValue() uint
	FileURLValue() URL
	BooleanValue() bool
	Data() []byte
	DoubleValue() float64
	DateValue() Date
	NumberOfItems() int
	IsRecordDescriptor() bool
}

// A wrapper for the Apple event descriptor data type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor?language=objc
type AppleEventDescriptor struct {
	objc.Object
}

func AppleEventDescriptorFrom(ptr unsafe.Pointer) AppleEventDescriptor {
	return AppleEventDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (a_ AppleEventDescriptor) InitRecordDescriptor() AppleEventDescriptor {
	rv := objc.Call[AppleEventDescriptor](a_, objc.Sel("initRecordDescriptor"))
	return rv
}

// Initializes a newly allocated instance as a descriptor that is an Apple event record. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1416093-initrecorddescriptor?language=objc
func NewAppleEventDescriptorRecordDescriptor() AppleEventDescriptor {
	instance := AppleEventDescriptorClass.Alloc().InitRecordDescriptor()
	instance.Autorelease()
	return instance
}

func (a_ AppleEventDescriptor) InitListDescriptor() AppleEventDescriptor {
	rv := objc.Call[AppleEventDescriptor](a_, objc.Sel("initListDescriptor"))
	return rv
}

// Initializes a newly allocated instance as an empty list descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1416351-initlistdescriptor?language=objc
func NewAppleEventDescriptorListDescriptor() AppleEventDescriptor {
	instance := AppleEventDescriptorClass.Alloc().InitListDescriptor()
	instance.Autorelease()
	return instance
}

func (ac _AppleEventDescriptorClass) Alloc() AppleEventDescriptor {
	rv := objc.Call[AppleEventDescriptor](ac, objc.Sel("alloc"))
	return rv
}

func AppleEventDescriptor_Alloc() AppleEventDescriptor {
	return AppleEventDescriptorClass.Alloc()
}

func (ac _AppleEventDescriptorClass) New() AppleEventDescriptor {
	rv := objc.Call[AppleEventDescriptor](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAppleEventDescriptor() AppleEventDescriptor {
	return AppleEventDescriptorClass.New()
}

func (a_ AppleEventDescriptor) Init() AppleEventDescriptor {
	rv := objc.Call[AppleEventDescriptor](a_, objc.Sel("init"))
	return rv
}

// Creates a descriptor initialized with type typeUnicodeText that stores the text from the specified string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1415227-descriptorwithstring?language=objc
func (ac _AppleEventDescriptorClass) DescriptorWithString(string_ string) AppleEventDescriptor {
	rv := objc.Call[AppleEventDescriptor](ac, objc.Sel("descriptorWithString:"), string_)
	return rv
}

// Creates a descriptor initialized with type typeUnicodeText that stores the text from the specified string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1415227-descriptorwithstring?language=objc
func AppleEventDescriptor_DescriptorWithString(string_ string) AppleEventDescriptor {
	return AppleEventDescriptorClass.DescriptorWithString(string_)
}

// Creates a descriptor initialized with Apple event type typeSInt32 that stores the specified integer value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1411435-descriptorwithint32?language=objc
func (ac _AppleEventDescriptorClass) DescriptorWithInt32(signedInt int32) AppleEventDescriptor {
	rv := objc.Call[AppleEventDescriptor](ac, objc.Sel("descriptorWithInt32:"), signedInt)
	return rv
}

// Creates a descriptor initialized with Apple event type typeSInt32 that stores the specified integer value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1411435-descriptorwithint32?language=objc
func AppleEventDescriptor_DescriptorWithInt32(signedInt int32) AppleEventDescriptor {
	return AppleEventDescriptorClass.DescriptorWithInt32(signedInt)
}

// Creates a descriptor initialized with type typeType that stores the specified type value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1407547-descriptorwithtypecode?language=objc
func (ac _AppleEventDescriptorClass) DescriptorWithTypeCode(typeCode uint) AppleEventDescriptor {
	rv := objc.Call[AppleEventDescriptor](ac, objc.Sel("descriptorWithTypeCode:"), typeCode)
	return rv
}

// Creates a descriptor initialized with type typeType that stores the specified type value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1407547-descriptorwithtypecode?language=objc
func AppleEventDescriptor_DescriptorWithTypeCode(typeCode uint) AppleEventDescriptor {
	return AppleEventDescriptorClass.DescriptorWithTypeCode(typeCode)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1411349-descriptorwithfileurl?language=objc
func (ac _AppleEventDescriptorClass) DescriptorWithFileURL(fileURL IURL) AppleEventDescriptor {
	rv := objc.Call[AppleEventDescriptor](ac, objc.Sel("descriptorWithFileURL:"), objc.Ptr(fileURL))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1411349-descriptorwithfileurl?language=objc
func AppleEventDescriptor_DescriptorWithFileURL(fileURL IURL) AppleEventDescriptor {
	return AppleEventDescriptorClass.DescriptorWithFileURL(fileURL)
}

// Creates and initializes an empty list descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1412854-listdescriptor?language=objc
func (ac _AppleEventDescriptorClass) ListDescriptor() AppleEventDescriptor {
	rv := objc.Call[AppleEventDescriptor](ac, objc.Sel("listDescriptor"))
	return rv
}

// Creates and initializes an empty list descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1412854-listdescriptor?language=objc
func AppleEventDescriptor_ListDescriptor() AppleEventDescriptor {
	return AppleEventDescriptorClass.ListDescriptor()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1415521-descriptorwithbundleidentifier?language=objc
func (ac _AppleEventDescriptorClass) DescriptorWithBundleIdentifier(bundleIdentifier string) AppleEventDescriptor {
	rv := objc.Call[AppleEventDescriptor](ac, objc.Sel("descriptorWithBundleIdentifier:"), bundleIdentifier)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1415521-descriptorwithbundleidentifier?language=objc
func AppleEventDescriptor_DescriptorWithBundleIdentifier(bundleIdentifier string) AppleEventDescriptor {
	return AppleEventDescriptorClass.DescriptorWithBundleIdentifier(bundleIdentifier)
}

// Creates a descriptor initialized with type typeEnumerated that stores the specified enumerator data type value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1413358-descriptorwithenumcode?language=objc
func (ac _AppleEventDescriptorClass) DescriptorWithEnumCode(enumerator uint) AppleEventDescriptor {
	rv := objc.Call[AppleEventDescriptor](ac, objc.Sel("descriptorWithEnumCode:"), enumerator)
	return rv
}

// Creates a descriptor initialized with type typeEnumerated that stores the specified enumerator data type value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1413358-descriptorwithenumcode?language=objc
func AppleEventDescriptor_DescriptorWithEnumCode(enumerator uint) AppleEventDescriptor {
	return AppleEventDescriptorClass.DescriptorWithEnumCode(enumerator)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1409710-descriptorwithdate?language=objc
func (ac _AppleEventDescriptorClass) DescriptorWithDate(date IDate) AppleEventDescriptor {
	rv := objc.Call[AppleEventDescriptor](ac, objc.Sel("descriptorWithDate:"), objc.Ptr(date))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1409710-descriptorwithdate?language=objc
func AppleEventDescriptor_DescriptorWithDate(date IDate) AppleEventDescriptor {
	return AppleEventDescriptorClass.DescriptorWithDate(date)
}

// Creates a descriptor initialized with type typeBoolean that stores the specified Boolean value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1410832-descriptorwithboolean?language=objc
func (ac _AppleEventDescriptorClass) DescriptorWithBoolean(boolean bool) AppleEventDescriptor {
	rv := objc.Call[AppleEventDescriptor](ac, objc.Sel("descriptorWithBoolean:"), boolean)
	return rv
}

// Creates a descriptor initialized with type typeBoolean that stores the specified Boolean value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1410832-descriptorwithboolean?language=objc
func AppleEventDescriptor_DescriptorWithBoolean(boolean bool) AppleEventDescriptor {
	return AppleEventDescriptorClass.DescriptorWithBoolean(boolean)
}

// Inserts a descriptor at the specified (one-based) position in the receiving descriptor list, replacing the existing descriptor, if any, at that position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1409452-insertdescriptor?language=objc
func (a_ AppleEventDescriptor) InsertDescriptorAtIndex(descriptor IAppleEventDescriptor, index int) {
	objc.Call[objc.Void](a_, objc.Sel("insertDescriptor:atIndex:"), objc.Ptr(descriptor), index)
}

// Creates and initializes a descriptor for an Apple event record whose data has yet to be set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1414380-recorddescriptor?language=objc
func (ac _AppleEventDescriptorClass) RecordDescriptor() AppleEventDescriptor {
	rv := objc.Call[AppleEventDescriptor](ac, objc.Sel("recordDescriptor"))
	return rv
}

// Creates and initializes a descriptor for an Apple event record whose data has yet to be set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1414380-recorddescriptor?language=objc
func AppleEventDescriptor_RecordDescriptor() AppleEventDescriptor {
	return AppleEventDescriptorClass.RecordDescriptor()
}

// Creates and initializes a descriptor with no parameter or attribute values set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1412488-nulldescriptor?language=objc
func (ac _AppleEventDescriptorClass) NullDescriptor() AppleEventDescriptor {
	rv := objc.Call[AppleEventDescriptor](ac, objc.Sel("nullDescriptor"))
	return rv
}

// Creates and initializes a descriptor with no parameter or attribute values set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1412488-nulldescriptor?language=objc
func AppleEventDescriptor_NullDescriptor() AppleEventDescriptor {
	return AppleEventDescriptorClass.NullDescriptor()
}

// Removes the descriptor at the specified (one-based) position in the receiving descriptor list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1413644-removedescriptoratindex?language=objc
func (a_ AppleEventDescriptor) RemoveDescriptorAtIndex(index int) {
	objc.Call[objc.Void](a_, objc.Sel("removeDescriptorAtIndex:"), index)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1416147-descriptorwithdouble?language=objc
func (ac _AppleEventDescriptorClass) DescriptorWithDouble(doubleValue float64) AppleEventDescriptor {
	rv := objc.Call[AppleEventDescriptor](ac, objc.Sel("descriptorWithDouble:"), doubleValue)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1416147-descriptorwithdouble?language=objc
func AppleEventDescriptor_DescriptorWithDouble(doubleValue float64) AppleEventDescriptor {
	return AppleEventDescriptorClass.DescriptorWithDouble(doubleValue)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1416614-currentprocessdescriptor?language=objc
func (ac _AppleEventDescriptorClass) CurrentProcessDescriptor() AppleEventDescriptor {
	rv := objc.Call[AppleEventDescriptor](ac, objc.Sel("currentProcessDescriptor"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1416614-currentprocessdescriptor?language=objc
func AppleEventDescriptor_CurrentProcessDescriptor() AppleEventDescriptor {
	return AppleEventDescriptorClass.CurrentProcessDescriptor()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1412834-descriptorwithapplicationurl?language=objc
func (ac _AppleEventDescriptorClass) DescriptorWithApplicationURL(applicationURL IURL) AppleEventDescriptor {
	rv := objc.Call[AppleEventDescriptor](ac, objc.Sel("descriptorWithApplicationURL:"), objc.Ptr(applicationURL))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1412834-descriptorwithapplicationurl?language=objc
func AppleEventDescriptor_DescriptorWithApplicationURL(applicationURL IURL) AppleEventDescriptor {
	return AppleEventDescriptorClass.DescriptorWithApplicationURL(applicationURL)
}

// Returns the descriptor at the specified (one-based) position in the receiving descriptor list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1408027-descriptoratindex?language=objc
func (a_ AppleEventDescriptor) DescriptorAtIndex(index int) AppleEventDescriptor {
	rv := objc.Call[AppleEventDescriptor](a_, objc.Sel("descriptorAtIndex:"), index)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1409337-sendeventwithoptions?language=objc
func (a_ AppleEventDescriptor) SendEventWithOptionsTimeoutError(sendOptions AppleEventSendOptions, timeoutInSeconds TimeInterval, error IError) AppleEventDescriptor {
	rv := objc.Call[AppleEventDescriptor](a_, objc.Sel("sendEventWithOptions:timeout:error:"), sendOptions, timeoutInSeconds, objc.Ptr(error))
	return rv
}

// The contents of the receiver as an integer, coercing (to typeSInt32) if necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1407270-int32value?language=objc
func (a_ AppleEventDescriptor) Int32Value() int32 {
	rv := objc.Call[int32](a_, objc.Sel("int32Value"))
	return rv
}

// The contents of the receiver as a type, coercing to typeType if necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1409662-typecodevalue?language=objc
func (a_ AppleEventDescriptor) TypeCodeValue() uint {
	rv := objc.Call[uint](a_, objc.Sel("typeCodeValue"))
	return rv
}

// The contents of the receiver as a Unicode text string, coercing to typeUnicodeText if necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1407584-stringvalue?language=objc
func (a_ AppleEventDescriptor) StringValue() string {
	rv := objc.Call[string](a_, objc.Sel("stringValue"))
	return rv
}

// The contents of the receiver as an enumeration type, coercing to typeEnumerated if necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1408039-enumcodevalue?language=objc
func (a_ AppleEventDescriptor) EnumCodeValue() uint {
	rv := objc.Call[uint](a_, objc.Sel("enumCodeValue"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1409448-fileurlvalue?language=objc
func (a_ AppleEventDescriptor) FileURLValue() URL {
	rv := objc.Call[URL](a_, objc.Sel("fileURLValue"))
	return rv
}

// The contents of the receiver as a Boolean value, coercing (to typeBoolean) if necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1412412-booleanvalue?language=objc
func (a_ AppleEventDescriptor) BooleanValue() bool {
	rv := objc.Call[bool](a_, objc.Sel("booleanValue"))
	return rv
}

// The receiver’s data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1413486-data?language=objc
func (a_ AppleEventDescriptor) Data() []byte {
	rv := objc.Call[[]byte](a_, objc.Sel("data"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1417847-doublevalue?language=objc
func (a_ AppleEventDescriptor) DoubleValue() float64 {
	rv := objc.Call[float64](a_, objc.Sel("doubleValue"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1416330-datevalue?language=objc
func (a_ AppleEventDescriptor) DateValue() Date {
	rv := objc.Call[Date](a_, objc.Sel("dateValue"))
	return rv
}

// The number of descriptors in the receiver’s descriptor list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1416786-numberofitems?language=objc
func (a_ AppleEventDescriptor) NumberOfItems() int {
	rv := objc.Call[int](a_, objc.Sel("numberOfItems"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/1410136-isrecorddescriptor?language=objc
func (a_ AppleEventDescriptor) IsRecordDescriptor() bool {
	rv := objc.Call[bool](a_, objc.Sel("isRecordDescriptor"))
	return rv
}
