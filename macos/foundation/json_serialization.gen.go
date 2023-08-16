// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [JSONSerialization] class.
var JSONSerializationClass = _JSONSerializationClass{objc.GetClass("NSJSONSerialization")}

type _JSONSerializationClass struct {
	objc.Class
}

// An interface definition for the [JSONSerialization] class.
type IJSONSerialization interface {
	objc.IObject
}

// An object that converts between JSON and the equivalent Foundation objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsjsonserialization?language=objc
type JSONSerialization struct {
	objc.Object
}

func JSONSerializationFrom(ptr unsafe.Pointer) JSONSerialization {
	return JSONSerialization{
		Object: objc.ObjectFrom(ptr),
	}
}

func (jc _JSONSerializationClass) Alloc() JSONSerialization {
	rv := objc.Call[JSONSerialization](jc, objc.Sel("alloc"))
	return rv
}

func JSONSerialization_Alloc() JSONSerialization {
	return JSONSerializationClass.Alloc()
}

func (jc _JSONSerializationClass) New() JSONSerialization {
	rv := objc.Call[JSONSerialization](jc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewJSONSerialization() JSONSerialization {
	return JSONSerializationClass.New()
}

func (j_ JSONSerialization) Init() JSONSerialization {
	rv := objc.Call[JSONSerialization](j_, objc.Sel("init"))
	return rv
}

// Returns a Foundation object from JSON data in a given stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsjsonserialization/1418059-jsonobjectwithstream?language=objc
func (jc _JSONSerializationClass) JSONObjectWithStreamOptionsError(stream IInputStream, opt JSONReadingOptions, error IError) objc.Object {
	rv := objc.Call[objc.Object](jc, objc.Sel("JSONObjectWithStream:options:error:"), objc.Ptr(stream), opt, objc.Ptr(error))
	return rv
}

// Returns a Foundation object from JSON data in a given stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsjsonserialization/1418059-jsonobjectwithstream?language=objc
func JSONSerialization_JSONObjectWithStreamOptionsError(stream IInputStream, opt JSONReadingOptions, error IError) objc.Object {
	return JSONSerializationClass.JSONObjectWithStreamOptionsError(stream, opt, error)
}

// Returns a Foundation object from given JSON data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsjsonserialization/1415493-jsonobjectwithdata?language=objc
func (jc _JSONSerializationClass) JSONObjectWithDataOptionsError(data []byte, opt JSONReadingOptions, error IError) objc.Object {
	rv := objc.Call[objc.Object](jc, objc.Sel("JSONObjectWithData:options:error:"), data, opt, objc.Ptr(error))
	return rv
}

// Returns a Foundation object from given JSON data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsjsonserialization/1415493-jsonobjectwithdata?language=objc
func JSONSerialization_JSONObjectWithDataOptionsError(data []byte, opt JSONReadingOptions, error IError) objc.Object {
	return JSONSerializationClass.JSONObjectWithDataOptionsError(data, opt, error)
}

// Writes a given JSON object to a stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsjsonserialization/1417433-writejsonobject?language=objc
func (jc _JSONSerializationClass) WriteJSONObjectToStreamOptionsError(obj objc.IObject, stream IOutputStream, opt JSONWritingOptions, error IError) int {
	rv := objc.Call[int](jc, objc.Sel("writeJSONObject:toStream:options:error:"), obj, objc.Ptr(stream), opt, objc.Ptr(error))
	return rv
}

// Writes a given JSON object to a stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsjsonserialization/1417433-writejsonobject?language=objc
func JSONSerialization_WriteJSONObjectToStreamOptionsError(obj objc.IObject, stream IOutputStream, opt JSONWritingOptions, error IError) int {
	return JSONSerializationClass.WriteJSONObjectToStreamOptionsError(obj, stream, opt, error)
}

// Returns a Boolean value that indicates whether the serializer can convert a given object to JSON data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsjsonserialization/1418461-isvalidjsonobject?language=objc
func (jc _JSONSerializationClass) IsValidJSONObject(obj objc.IObject) bool {
	rv := objc.Call[bool](jc, objc.Sel("isValidJSONObject:"), obj)
	return rv
}

// Returns a Boolean value that indicates whether the serializer can convert a given object to JSON data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsjsonserialization/1418461-isvalidjsonobject?language=objc
func JSONSerialization_IsValidJSONObject(obj objc.IObject) bool {
	return JSONSerializationClass.IsValidJSONObject(obj)
}

// Returns JSON data from a Foundation object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsjsonserialization/1413636-datawithjsonobject?language=objc
func (jc _JSONSerializationClass) DataWithJSONObjectOptionsError(obj objc.IObject, opt JSONWritingOptions, error IError) []byte {
	rv := objc.Call[[]byte](jc, objc.Sel("dataWithJSONObject:options:error:"), obj, opt, objc.Ptr(error))
	return rv
}

// Returns JSON data from a Foundation object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsjsonserialization/1413636-datawithjsonobject?language=objc
func JSONSerialization_DataWithJSONObjectOptionsError(obj objc.IObject, opt JSONWritingOptions, error IError) []byte {
	return JSONSerializationClass.DataWithJSONObjectOptionsError(obj, opt, error)
}
