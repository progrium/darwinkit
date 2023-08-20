// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Coder] class.
var CoderClass = _CoderClass{objc.GetClass("NSCoder")}

type _CoderClass struct {
	objc.Class
}

// An interface definition for the [Coder] class.
type ICoder interface {
	objc.IObject
	DecodeDataObject() []byte
	DecodeSizeForKey(key string) Size
	EncodeCMTimeMappingForKey(timeMapping coremedia.TimeMapping, key string)
	DecodeFloatForKey(key string) float64
	DecodeSize() Size
	EncodeRectForKey(rect Rect, key string)
	DecodeDictionaryWithKeysOfClassesObjectsOfClassesForKey(keyClasses ISet, objectClasses ISet, key string) Dictionary
	EncodePropertyList(aPropertyList objc.IObject)
	EncodeObject(object objc.IObject)
	DecodeDoubleForKey(key string) float64
	VersionForClassName(className string) int
	EncodeCMTimeRangeForKey(timeRange coremedia.TimeRange, key string)
	EncodeArrayOfObjCTypeCountAt(type_ *uint8, count uint, array unsafe.Pointer)
	EncodeValueOfObjCTypeAt(type_ *uint8, addr unsafe.Pointer)
	EncodeConditionalObject(object objc.IObject)
	DecodeValuesOfObjCTypes(types *uint8, args ...any)
	DecodeObject() objc.Object
	DecodeObjectOfClassForKey(aClass objc.IClass, key string) objc.Object
	FailWithError(error IError)
	DecodeInt64ForKey(key string) int64
	DecodeBytesForKeyReturnedLength(key string, lengthp *uint) *uint8
	DecodeCMTimeForKey(key string) coremedia.Time
	DecodePoint() Point
	EncodeBycopyObject(anObject objc.IObject)
	EncodeBoolForKey(value bool, key string)
	DecodeTopLevelObjectForKeyError(key string, error IError) objc.Object
	DecodeTopLevelObjectOfClassForKeyError(aClass objc.IClass, key string, error IError) objc.Object
	DecodeObjectOfClassesForKey(classes ISet, key string) objc.Object
	DecodeArrayOfObjectsOfClassForKey(cls objc.IClass, key string) []objc.Object
	DecodeTopLevelObjectAndReturnError(error IError) objc.Object
	DecodeArrayOfObjCTypeCountAt(itemType *uint8, count uint, array unsafe.Pointer)
	EncodeDataObject(data []byte)
	DecodeValueOfObjCTypeAtSize(type_ *uint8, data unsafe.Pointer, size uint)
	EncodeCMTimeForKey(time coremedia.Time, key string)
	ObjectZone() unsafe.Pointer
	DecodeInt32ForKey(key string) int32
	DecodeObjectForKey(key string) objc.Object
	EncodeRootObject(rootObject objc.IObject)
	DecodeRectForKey(key string) Rect
	DecodeRect() Rect
	SetObjectZone(zone unsafe.Pointer)
	DecodeIntForKey(key string) int
	EncodeSizeForKey(size Size, key string)
	DecodeIntegerForKey(key string) int
	EncodeIntegerForKey(value int, key string)
	EncodeByrefObject(anObject objc.IObject)
	EncodeInt32ForKey(value int32, key string)
	EncodePointForKey(point Point, key string)
	DecodePropertyList() objc.Object
	DecodePointForKey(key string) Point
	DecodeDictionaryWithKeysOfClassObjectsOfClassForKey(keyCls objc.IClass, objectCls objc.IClass, key string) Dictionary
	EncodeBytesLength(byteaddr unsafe.Pointer, length uint)
	DecodePropertyListForKey(key string) objc.Object
	DecodeTopLevelObjectOfClassesForKeyError(classes ISet, key string, error IError) objc.Object
	DecodeBoolForKey(key string) bool
	DecodeArrayOfObjectsOfClassesForKey(classes ISet, key string) []objc.Object
	DecodeCMTimeRangeForKey(key string) coremedia.TimeRange
	EncodeFloatForKey(value float64, key string)
	EncodeValuesOfObjCTypes(types *uint8, args ...any)
	ContainsValueForKey(key string) bool
	EncodeDoubleForKey(value float64, key string)
	EncodeIntForKey(value int, key string)
	DecodeCMTimeMappingForKey(key string) coremedia.TimeMapping
	DecodeBytesWithReturnedLength(lengthp *uint) unsafe.Pointer
	EncodeInt64ForKey(value int64, key string)
	Error() Error
	DecodingFailurePolicy() DecodingFailurePolicy
	AllowedClasses() Set
	AllowsKeyedCoding() bool
	SystemVersion() int
	RequiresSecureCoding() bool
}

// An abstract class that serves as the basis for objects that enable archiving and distribution of other objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder?language=objc
type Coder struct {
	objc.Object
}

func CoderFrom(ptr unsafe.Pointer) Coder {
	return Coder{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CoderClass) Alloc() Coder {
	rv := objc.Call[Coder](cc, objc.Sel("alloc"))
	return rv
}

func Coder_Alloc() Coder {
	return CoderClass.Alloc()
}

func (cc _CoderClass) New() Coder {
	rv := objc.Call[Coder](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCoder() Coder {
	return CoderClass.New()
}

func (c_ Coder) Init() Coder {
	rv := objc.Call[Coder](c_, objc.Sel("init"))
	return rv
}

// Decodes and returns an NSData object that was previously encoded with encodeDataObject:. Subclasses must override this method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1409876-decodedataobject?language=objc
func (c_ Coder) DecodeDataObject() []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("decodeDataObject"))
	return rv
}

// Decodes and returns an NSSize structure that was previously encoded with encodeSize:forKey:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1391253-decodesizeforkey?language=objc
func (c_ Coder) DecodeSizeForKey(key string) Size {
	rv := objc.Call[Size](c_, objc.Sel("decodeSizeForKey:"), key)
	return rv
}

// Encodes a given Core Media time mapping structure and associates it with a specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1389496-encodecmtimemapping?language=objc
func (c_ Coder) EncodeCMTimeMappingForKey(timeMapping coremedia.TimeMapping, key string) {
	objc.Call[objc.Void](c_, objc.Sel("encodeCMTimeMapping:forKey:"), timeMapping, key)
}

// Decodes and returns a float value that was previously encoded with encodeFloat:forKey: or encodeDouble:forKey: and associated with the string key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1408104-decodefloatforkey?language=objc
func (c_ Coder) DecodeFloatForKey(key string) float64 {
	rv := objc.Call[float64](c_, objc.Sel("decodeFloatForKey:"), key)
	return rv
}

// Decodes and returns an NSSize structure that was previously encoded with encodeSize:forKey:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1391144-decodesize?language=objc
func (c_ Coder) DecodeSize() Size {
	rv := objc.Call[Size](c_, objc.Sel("decodeSize"))
	return rv
}

// Encodes a rectangle structure and associates it with the string key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1391287-encoderect?language=objc
func (c_ Coder) EncodeRectForKey(rect Rect, key string) {
	objc.Call[objc.Void](c_, objc.Sel("encodeRect:forKey:"), rect, key)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/3563980-decodedictionarywithkeysofclasse?language=objc
func (c_ Coder) DecodeDictionaryWithKeysOfClassesObjectsOfClassesForKey(keyClasses ISet, objectClasses ISet, key string) Dictionary {
	rv := objc.Call[Dictionary](c_, objc.Sel("decodeDictionaryWithKeysOfClasses:objectsOfClasses:forKey:"), objc.Ptr(keyClasses), objc.Ptr(objectClasses), key)
	return rv
}

// Encodes a property list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1410643-encodepropertylist?language=objc
func (c_ Coder) EncodePropertyList(aPropertyList objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("encodePropertyList:"), aPropertyList)
}

// Encodes an object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1417647-encodeobject?language=objc
func (c_ Coder) EncodeObject(object objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("encodeObject:"), object)
}

// Decodes and returns a double value that was previously encoded with either encodeFloat:forKey: or encodeDouble:forKey: and associated with the string key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1409374-decodedoubleforkey?language=objc
func (c_ Coder) DecodeDoubleForKey(key string) float64 {
	rv := objc.Call[float64](c_, objc.Sel("decodeDoubleForKey:"), key)
	return rv
}

// This method is present for historical reasons and is not used with keyed archivers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1417703-versionforclassname?language=objc
func (c_ Coder) VersionForClassName(className string) int {
	rv := objc.Call[int](c_, objc.Sel("versionForClassName:"), className)
	return rv
}

// Encodes a given Core Media time range structure and associates it with a specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1386649-encodecmtimerange?language=objc
func (c_ Coder) EncodeCMTimeRangeForKey(timeRange coremedia.TimeRange, key string) {
	objc.Call[objc.Void](c_, objc.Sel("encodeCMTimeRange:forKey:"), timeRange, key)
}

// Encodes an array of the given Objective-C type, provided the number of items and a pointer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1417865-encodearrayofobjctype?language=objc
func (c_ Coder) EncodeArrayOfObjCTypeCountAt(type_ *uint8, count uint, array unsafe.Pointer) {
	objc.Call[objc.Void](c_, objc.Sel("encodeArrayOfObjCType:count:at:"), type_, count, array)
}

// Encodes a value of the given type at the given address. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1414648-encodevalueofobjctype?language=objc
func (c_ Coder) EncodeValueOfObjCTypeAt(type_ *uint8, addr unsafe.Pointer) {
	objc.Call[objc.Void](c_, objc.Sel("encodeValueOfObjCType:at:"), type_, addr)
}

// An encoding method for subclasses to override to conditionally encode an object, preserving common references to it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1415196-encodeconditionalobject?language=objc
func (c_ Coder) EncodeConditionalObject(object objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("encodeConditionalObject:"), object)
}

// Decodes a series of potentially different Objective-C types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1442577-decodevaluesofobjctypes?language=objc
func (c_ Coder) DecodeValuesOfObjCTypes(types *uint8, args ...any) {
	objc.Call[objc.Void](c_, objc.Sel("decodeValuesOfObjCTypes:"), append([]any{types}, args...)...)
}

// Decodes and returns an object that was previously encoded with any of the encode…Object methods. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1414478-decodeobject?language=objc
func (c_ Coder) DecodeObject() objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("decodeObject"))
	return rv
}

// Decodes an object for the key, restricted to the specified class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1442558-decodeobjectofclass?language=objc
func (c_ Coder) DecodeObjectOfClassForKey(aClass objc.IClass, key string) objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("decodeObjectOfClass:forKey:"), objc.Ptr(aClass), key)
	return rv
}

// Signals to this coder that the decode operation has failed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1411455-failwitherror?language=objc
func (c_ Coder) FailWithError(error IError) {
	objc.Call[objc.Void](c_, objc.Sel("failWithError:"), objc.Ptr(error))
}

// Decodes and returns a 64-bit integer value that was previously encoded with encodeInt:forKey:, encodeInteger:forKey:, encodeInt32:forKey:, or encodeInt64:forKey: and associated with the string key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1407878-decodeint64forkey?language=objc
func (c_ Coder) DecodeInt64ForKey(key string) int64 {
	rv := objc.Call[int64](c_, objc.Sel("decodeInt64ForKey:"), key)
	return rv
}

// Decodes a buffer of data that was previously encoded with encodeBytes:length: and associated with the string key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1411987-decodebytesforkey?language=objc
func (c_ Coder) DecodeBytesForKeyReturnedLength(key string, lengthp *uint) *uint8 {
	rv := objc.Call[*uint8](c_, objc.Sel("decodeBytesForKey:returnedLength:"), key, lengthp)
	return rv
}

// Returns the Core Media time structure associated with a given key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1389544-decodecmtimeforkey?language=objc
func (c_ Coder) DecodeCMTimeForKey(key string) coremedia.Time {
	rv := objc.Call[coremedia.Time](c_, objc.Sel("decodeCMTimeForKey:"), key)
	return rv
}

// Decodes and returns an NSPoint structure that was previously encoded with encodePoint:forKey:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1391189-decodepoint?language=objc
func (c_ Coder) DecodePoint() Point {
	rv := objc.Call[Point](c_, objc.Sel("decodePoint"))
	return rv
}

// An encoding method for subclasses to override such that it creates a copy, rather than a proxy, when decoded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1418225-encodebycopyobject?language=objc
func (c_ Coder) EncodeBycopyObject(anObject objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("encodeBycopyObject:"), anObject)
}

// Encodes a Boolean value and associates it with the string key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1409236-encodebool?language=objc
func (c_ Coder) EncodeBoolForKey(value bool, key string) {
	objc.Call[objc.Void](c_, objc.Sel("encodeBool:forKey:"), value, key)
}

// Decodes the previously-encoded object associated by a key, populating an error if decoding fails. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1442541-decodetoplevelobjectforkey?language=objc
func (c_ Coder) DecodeTopLevelObjectForKeyError(key string, error IError) objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("decodeTopLevelObjectForKey:error:"), key, objc.Ptr(error))
	return rv
}

// Decode an object as an expected type, failing if the archived type does not match. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1442575-decodetoplevelobjectofclass?language=objc
func (c_ Coder) DecodeTopLevelObjectOfClassForKeyError(aClass objc.IClass, key string, error IError) objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("decodeTopLevelObjectOfClass:forKey:error:"), objc.Ptr(aClass), key, objc.Ptr(error))
	return rv
}

// Decodes an object for the key, restricted to the specified classes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1442560-decodeobjectofclasses?language=objc
func (c_ Coder) DecodeObjectOfClassesForKey(classes ISet, key string) objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("decodeObjectOfClasses:forKey:"), objc.Ptr(classes), key)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/3563977-decodearrayofobjectsofclass?language=objc
func (c_ Coder) DecodeArrayOfObjectsOfClassForKey(cls objc.IClass, key string) []objc.Object {
	rv := objc.Call[[]objc.Object](c_, objc.Sel("decodeArrayOfObjectsOfClass:forKey:"), objc.Ptr(cls), key)
	return rv
}

// Decodes a previously-encoded object, populating an error if decoding fails. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1442553-decodetoplevelobjectandreturnerr?language=objc
func (c_ Coder) DecodeTopLevelObjectAndReturnError(error IError) objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("decodeTopLevelObjectAndReturnError:"), objc.Ptr(error))
	return rv
}

// Decodes an array of count items, whose Objective-C type is given by itemType. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1408354-decodearrayofobjctype?language=objc
func (c_ Coder) DecodeArrayOfObjCTypeCountAt(itemType *uint8, count uint, array unsafe.Pointer) {
	objc.Call[objc.Void](c_, objc.Sel("decodeArrayOfObjCType:count:at:"), itemType, count, array)
}

// Encodes a given data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1411473-encodedataobject?language=objc
func (c_ Coder) EncodeDataObject(data []byte) {
	objc.Call[objc.Void](c_, objc.Sel("encodeDataObject:"), data)
}

// Decodes a single value of a known type from the specified data buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/2919430-decodevalueofobjctype?language=objc
func (c_ Coder) DecodeValueOfObjCTypeAtSize(type_ *uint8, data unsafe.Pointer, size uint) {
	objc.Call[objc.Void](c_, objc.Sel("decodeValueOfObjCType:at:size:"), type_, data, size)
}

// Encodes a given Core Media time structure and associates it with a specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1388869-encodecmtime?language=objc
func (c_ Coder) EncodeCMTimeForKey(time coremedia.Time, key string) {
	objc.Call[objc.Void](c_, objc.Sel("encodeCMTime:forKey:"), time, key)
}

// This method is present for historical reasons and has no effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1442570-objectzone?language=objc
func (c_ Coder) ObjectZone() unsafe.Pointer {
	rv := objc.Call[unsafe.Pointer](c_, objc.Sel("objectZone"))
	return rv
}

// Decodes and returns a 32-bit integer value that was previously encoded with encodeInt:forKey:, encodeInteger:forKey:, encodeInt32:forKey:, or encodeInt64:forKey: and associated with the string key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1408918-decodeint32forkey?language=objc
func (c_ Coder) DecodeInt32ForKey(key string) int32 {
	rv := objc.Call[int32](c_, objc.Sel("decodeInt32ForKey:"), key)
	return rv
}

// Decodes and returns a previously-encoded object that was previously encoded with encodeObject: or encodeConditionalObject: and associated with the string key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1418185-decodeobjectforkey?language=objc
func (c_ Coder) DecodeObjectForKey(key string) objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("decodeObjectForKey:"), key)
	return rv
}

// An encoding method for subclasses to override to encode an interconnected group of objects, starting with the provided root object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1409439-encoderootobject?language=objc
func (c_ Coder) EncodeRootObject(rootObject objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("encodeRootObject:"), rootObject)
}

// Decodes and returns an NSRect structure that was previously encoded with encodeRect:forKey:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1391116-decoderectforkey?language=objc
func (c_ Coder) DecodeRectForKey(key string) Rect {
	rv := objc.Call[Rect](c_, objc.Sel("decodeRectForKey:"), key)
	return rv
}

// Decodes and returns an NSRect structure that was previously encoded with encodeRect:forKey:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1391269-decoderect?language=objc
func (c_ Coder) DecodeRect() Rect {
	rv := objc.Call[Rect](c_, objc.Sel("decodeRect"))
	return rv
}

// This method is present for historical reasons and has no effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1442547-setobjectzone?language=objc
func (c_ Coder) SetObjectZone(zone unsafe.Pointer) {
	objc.Call[objc.Void](c_, objc.Sel("setObjectZone:"), zone)
}

// Decodes and returns an int value that was previously encoded with encodeInt:forKey:, encodeInteger:forKey:, encodeInt32:forKey:, or encodeInt64:forKey: and associated with the string key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1411168-decodeintforkey?language=objc
func (c_ Coder) DecodeIntForKey(key string) int {
	rv := objc.Call[int](c_, objc.Sel("decodeIntForKey:"), key)
	return rv
}

// Encodes a size structure and associates it with the given string key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1391176-encodesize?language=objc
func (c_ Coder) EncodeSizeForKey(size Size, key string) {
	objc.Call[objc.Void](c_, objc.Sel("encodeSize:forKey:"), size, key)
}

// Decodes and returns an NSInteger value that was previously encoded with encodeInt:forKey:, encodeInteger:forKey:, encodeInt32:forKey:, or encodeInt64:forKey: and associated with the string key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1409246-decodeintegerforkey?language=objc
func (c_ Coder) DecodeIntegerForKey(key string) int {
	rv := objc.Call[int](c_, objc.Sel("decodeIntegerForKey:"), key)
	return rv
}

// Encodes an integer value and associates it with the string key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1411551-encodeinteger?language=objc
func (c_ Coder) EncodeIntegerForKey(value int, key string) {
	objc.Call[objc.Void](c_, objc.Sel("encodeInteger:forKey:"), value, key)
}

// An encoding method for subclasses to override such that it creates a proxy, rather than a copy, when decoded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1416279-encodebyrefobject?language=objc
func (c_ Coder) EncodeByrefObject(anObject objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("encodeByrefObject:"), anObject)
}

// Encodes a 32-bit integer value and associates it with the string key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1412395-encodeint32?language=objc
func (c_ Coder) EncodeInt32ForKey(value int32, key string) {
	objc.Call[objc.Void](c_, objc.Sel("encodeInt32:forKey:"), value, key)
}

// Encodes a point and associates it with the string key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1391114-encodepoint?language=objc
func (c_ Coder) EncodePointForKey(point Point, key string) {
	objc.Call[objc.Void](c_, objc.Sel("encodePoint:forKey:"), point, key)
}

// Decodes a property list that was previously encoded with encodePropertyList:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1411916-decodepropertylist?language=objc
func (c_ Coder) DecodePropertyList() objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("decodePropertyList"))
	return rv
}

// Decodes and returns an NSPoint structure that was previously encoded with encodePoint:forKey:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1391214-decodepointforkey?language=objc
func (c_ Coder) DecodePointForKey(key string) Point {
	rv := objc.Call[Point](c_, objc.Sel("decodePointForKey:"), key)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/3563979-decodedictionarywithkeysofclass?language=objc
func (c_ Coder) DecodeDictionaryWithKeysOfClassObjectsOfClassForKey(keyCls objc.IClass, objectCls objc.IClass, key string) Dictionary {
	rv := objc.Call[Dictionary](c_, objc.Sel("decodeDictionaryWithKeysOfClass:objectsOfClass:forKey:"), objc.Ptr(keyCls), objc.Ptr(objectCls), key)
	return rv
}

// Encodes a buffer of data of an unspecified type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1411664-encodebytes?language=objc
func (c_ Coder) EncodeBytesLength(byteaddr unsafe.Pointer, length uint) {
	objc.Call[objc.Void](c_, objc.Sel("encodeBytes:length:"), byteaddr, length)
}

// Returns a decoded property list for the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1416284-decodepropertylistforkey?language=objc
func (c_ Coder) DecodePropertyListForKey(key string) objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("decodePropertyListForKey:"), key)
	return rv
}

// Decode an object as one of several expected types, failing if the archived type does not match. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1442539-decodetoplevelobjectofclasses?language=objc
func (c_ Coder) DecodeTopLevelObjectOfClassesForKeyError(classes ISet, key string, error IError) objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("decodeTopLevelObjectOfClasses:forKey:error:"), objc.Ptr(classes), key, objc.Ptr(error))
	return rv
}

// Decodes and returns a boolean value that was previously encoded with encodeBool:forKey: and associated with the string key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1409293-decodeboolforkey?language=objc
func (c_ Coder) DecodeBoolForKey(key string) bool {
	rv := objc.Call[bool](c_, objc.Sel("decodeBoolForKey:"), key)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/3563978-decodearrayofobjectsofclasses?language=objc
func (c_ Coder) DecodeArrayOfObjectsOfClassesForKey(classes ISet, key string) []objc.Object {
	rv := objc.Call[[]objc.Object](c_, objc.Sel("decodeArrayOfObjectsOfClasses:forKey:"), objc.Ptr(classes), key)
	return rv
}

// Returns the Core Media time range structure associated with a given key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1385718-decodecmtimerangeforkey?language=objc
func (c_ Coder) DecodeCMTimeRangeForKey(key string) coremedia.TimeRange {
	rv := objc.Call[coremedia.TimeRange](c_, objc.Sel("decodeCMTimeRangeForKey:"), key)
	return rv
}

// Encodes a floating point value and associates it with the string key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1414384-encodefloat?language=objc
func (c_ Coder) EncodeFloatForKey(value float64, key string) {
	objc.Call[objc.Void](c_, objc.Sel("encodeFloat:forKey:"), value, key)
}

// Encodes a series of values of potentially differing Objective-C types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1442581-encodevaluesofobjctypes?language=objc
func (c_ Coder) EncodeValuesOfObjCTypes(types *uint8, args ...any) {
	objc.Call[objc.Void](c_, objc.Sel("encodeValuesOfObjCTypes:"), append([]any{types}, args...)...)
}

// Returns a Boolean value that indicates whether an encoded value is available for a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1416125-containsvalueforkey?language=objc
func (c_ Coder) ContainsValueForKey(key string) bool {
	rv := objc.Call[bool](c_, objc.Sel("containsValueForKey:"), key)
	return rv
}

// Encodes a double-precision floating point value and associates it with the string key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1409008-encodedouble?language=objc
func (c_ Coder) EncodeDoubleForKey(value float64, key string) {
	objc.Call[objc.Void](c_, objc.Sel("encodeDouble:forKey:"), value, key)
}

// Encodes a C integer value and associates it with the string key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1413906-encodeint?language=objc
func (c_ Coder) EncodeIntForKey(value int, key string) {
	objc.Call[objc.Void](c_, objc.Sel("encodeInt:forKey:"), value, key)
}

// Returns the Core Media time mapping structure associated with a given key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1389860-decodecmtimemappingforkey?language=objc
func (c_ Coder) DecodeCMTimeMappingForKey(key string) coremedia.TimeMapping {
	rv := objc.Call[coremedia.TimeMapping](c_, objc.Sel("decodeCMTimeMappingForKey:"), key)
	return rv
}

// Decodes a buffer of data whose types are unspecified. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1415441-decodebyteswithreturnedlength?language=objc
func (c_ Coder) DecodeBytesWithReturnedLength(lengthp *uint) unsafe.Pointer {
	rv := objc.Call[unsafe.Pointer](c_, objc.Sel("decodeBytesWithReturnedLength:"), lengthp)
	return rv
}

// Encodes a 64-bit integer value and associates it with the string key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1410440-encodeint64?language=objc
func (c_ Coder) EncodeInt64ForKey(value int64, key string) {
	objc.Call[objc.Void](c_, objc.Sel("encodeInt64:forKey:"), value, key)
}

// An error in the top-level encode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1643263-error?language=objc
func (c_ Coder) Error() Error {
	rv := objc.Call[Error](c_, objc.Sel("error"))
	return rv
}

// The action the coder should take when decoding fails. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1642984-decodingfailurepolicy?language=objc
func (c_ Coder) DecodingFailurePolicy() DecodingFailurePolicy {
	rv := objc.Call[DecodingFailurePolicy](c_, objc.Sel("decodingFailurePolicy"))
	return rv
}

// The set of coded classes allowed for secure coding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1412486-allowedclasses?language=objc
func (c_ Coder) AllowedClasses() Set {
	rv := objc.Call[Set](c_, objc.Sel("allowedClasses"))
	return rv
}

// A Boolean value that indicates whether the receiver supports keyed coding of objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1417541-allowskeyedcoding?language=objc
func (c_ Coder) AllowsKeyedCoding() bool {
	rv := objc.Call[bool](c_, objc.Sel("allowsKeyedCoding"))
	return rv
}

// The system version in effect for the archive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1413205-systemversion?language=objc
func (c_ Coder) SystemVersion() int {
	rv := objc.Call[int](c_, objc.Sel("systemVersion"))
	return rv
}

// Indicates whether the archiver requires all archived classes to resist object substitution attacks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/1409845-requiressecurecoding?language=objc
func (c_ Coder) RequiresSecureCoding() bool {
	rv := objc.Call[bool](c_, objc.Sel("requiresSecureCoding"))
	return rv
}
