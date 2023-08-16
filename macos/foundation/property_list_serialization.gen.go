// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PropertyListSerialization] class.
var PropertyListSerializationClass = _PropertyListSerializationClass{objc.GetClass("NSPropertyListSerialization")}

type _PropertyListSerializationClass struct {
	objc.Class
}

// An interface definition for the [PropertyListSerialization] class.
type IPropertyListSerialization interface {
	objc.IObject
}

// An object that converts between a property list and one of several serialized representations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylistserialization?language=objc
type PropertyListSerialization struct {
	objc.Object
}

func PropertyListSerializationFrom(ptr unsafe.Pointer) PropertyListSerialization {
	return PropertyListSerialization{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PropertyListSerializationClass) Alloc() PropertyListSerialization {
	rv := objc.Call[PropertyListSerialization](pc, objc.Sel("alloc"))
	return rv
}

func PropertyListSerialization_Alloc() PropertyListSerialization {
	return PropertyListSerializationClass.Alloc()
}

func (pc _PropertyListSerializationClass) New() PropertyListSerialization {
	rv := objc.Call[PropertyListSerialization](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPropertyListSerialization() PropertyListSerialization {
	return PropertyListSerializationClass.New()
}

func (p_ PropertyListSerialization) Init() PropertyListSerialization {
	rv := objc.Call[PropertyListSerialization](p_, objc.Sel("init"))
	return rv
}

// Returns an NSData object containing a given property list in a specified format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylistserialization/1418309-datawithpropertylist?language=objc
func (pc _PropertyListSerializationClass) DataWithPropertyListFormatOptionsError(plist objc.IObject, format PropertyListFormat, opt PropertyListWriteOptions, error IError) []byte {
	rv := objc.Call[[]byte](pc, objc.Sel("dataWithPropertyList:format:options:error:"), plist, format, opt, objc.Ptr(error))
	return rv
}

// Returns an NSData object containing a given property list in a specified format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylistserialization/1418309-datawithpropertylist?language=objc
func PropertyListSerialization_DataWithPropertyListFormatOptionsError(plist objc.IObject, format PropertyListFormat, opt PropertyListWriteOptions, error IError) []byte {
	return PropertyListSerializationClass.DataWithPropertyListFormatOptionsError(plist, format, opt, error)
}

// Returns a Boolean value that indicates whether a given property list is valid for a given format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylistserialization/1411020-propertylist?language=objc
func (pc _PropertyListSerializationClass) PropertyListIsValidForFormat(plist objc.IObject, format PropertyListFormat) bool {
	rv := objc.Call[bool](pc, objc.Sel("propertyList:isValidForFormat:"), plist, format)
	return rv
}

// Returns a Boolean value that indicates whether a given property list is valid for a given format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylistserialization/1411020-propertylist?language=objc
func PropertyListSerialization_PropertyListIsValidForFormat(plist objc.IObject, format PropertyListFormat) bool {
	return PropertyListSerializationClass.PropertyListIsValidForFormat(plist, format)
}

// Creates and returns a property list by reading from the specified stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylistserialization/1415468-propertylistwithstream?language=objc
func (pc _PropertyListSerializationClass) PropertyListWithStreamOptionsFormatError(stream IInputStream, opt PropertyListReadOptions, format *PropertyListFormat, error IError) objc.Object {
	rv := objc.Call[objc.Object](pc, objc.Sel("propertyListWithStream:options:format:error:"), objc.Ptr(stream), opt, format, objc.Ptr(error))
	return rv
}

// Creates and returns a property list by reading from the specified stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylistserialization/1415468-propertylistwithstream?language=objc
func PropertyListSerialization_PropertyListWithStreamOptionsFormatError(stream IInputStream, opt PropertyListReadOptions, format *PropertyListFormat, error IError) objc.Object {
	return PropertyListSerializationClass.PropertyListWithStreamOptionsFormatError(stream, opt, format, error)
}

// Creates and returns a property list from the specified data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylistserialization/1409678-propertylistwithdata?language=objc
func (pc _PropertyListSerializationClass) PropertyListWithDataOptionsFormatError(data []byte, opt PropertyListReadOptions, format *PropertyListFormat, error IError) objc.Object {
	rv := objc.Call[objc.Object](pc, objc.Sel("propertyListWithData:options:format:error:"), data, opt, format, objc.Ptr(error))
	return rv
}

// Creates and returns a property list from the specified data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylistserialization/1409678-propertylistwithdata?language=objc
func PropertyListSerialization_PropertyListWithDataOptionsFormatError(data []byte, opt PropertyListReadOptions, format *PropertyListFormat, error IError) objc.Object {
	return PropertyListSerializationClass.PropertyListWithDataOptionsFormatError(data, opt, format, error)
}

// Writes a property list to the specified stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylistserialization/1407862-writepropertylist?language=objc
func (pc _PropertyListSerializationClass) WritePropertyListToStreamFormatOptionsError(plist objc.IObject, stream IOutputStream, format PropertyListFormat, opt PropertyListWriteOptions, error IError) int {
	rv := objc.Call[int](pc, objc.Sel("writePropertyList:toStream:format:options:error:"), plist, objc.Ptr(stream), format, opt, objc.Ptr(error))
	return rv
}

// Writes a property list to the specified stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylistserialization/1407862-writepropertylist?language=objc
func PropertyListSerialization_WritePropertyListToStreamFormatOptionsError(plist objc.IObject, stream IOutputStream, format PropertyListFormat, opt PropertyListWriteOptions, error IError) int {
	return PropertyListSerializationClass.WritePropertyListToStreamFormatOptionsError(plist, stream, format, opt, error)
}
