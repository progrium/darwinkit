package foundation

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/progrium/darwinkit/macos/coregraphics"
	"github.com/progrium/darwinkit/objc"
)

type AppleEventManagerSuspensionID unsafe.Pointer
type RangePointer unsafe.Pointer
type RectPointer unsafe.Pointer
type PointPointer unsafe.Pointer
type RectArray unsafe.Pointer
type SizeArray unsafe.Pointer
type PointArray unsafe.Pointer

type PropertyListReadOptions PropertyListMutabilityOptions

type Point = coregraphics.Point
type Size = coregraphics.Size
type Rect = coregraphics.Rect

type FastEnumerationState struct {
	// TODO
}

func (N *String) String() string {
	return objc.ToGoString(N.Ptr())
}

func Data_ToBytes(data IData) []byte {
	return objc.ToGoBytes(data.Ptr())
}

func ArrayOf[T any](slice []T) Array {
	return ArrayFrom(objc.ToNSArray(reflect.ValueOf(slice)))
}

func ArrayToSlice[T any](a Array) []T {
	var zero []T
	return objc.ToGoSlice(a.Ptr(), reflect.TypeOf(zero)).Interface().([]T)
}

func DictOf[K comparable, V any](m map[K]V) Dictionary {
	return DictionaryFrom(objc.ToNSDict(reflect.ValueOf(m)))
}

func DictToMap[K comparable, V any](d Dictionary) map[K]V {
	var zero map[K]V
	return objc.ToGoMap(d.Ptr(), reflect.TypeOf(zero)).Interface().(map[K]V)
}

func ToNSError(code int, err error) Error {
	if err == nil {
		return Error{}
	}
	return ErrorClass.ErrorWithDomainCodeUserInfo("darwinkit.progrium.github.com", code, map[ErrorUserInfoKey]objc.IObject{
		"Error reason": NewStringWithString(err.Error()),
	})
}

// Error convert

type ocErr struct {
	code    int
	message string
}

// Error implements error
func (e ocErr) Error() string {
	return fmt.Sprintf("object error, code: %v, message: %v", e.code, e.message)
}

func ToGoError(err Error) error {
	if err.IsNil() {
		return nil
	}
	return ocErr{
		code:    err.Code(),
		message: err.Description(),
	}
}

var _NSJSONSerializationClass = objc.GetClass("NSJSONSerialization")

func JSONObjectWithData(data []byte, options JSONReadingOptions) (objc.Object, error) {
	var err Error
	r := objc.Call[objc.Object](_NSJSONSerializationClass, objc.Sel("JSONObjectWithData:options:error:"), data, options, unsafe.Pointer(&err))
	return r, ToGoError(err)
}

func DataWithJSONObject(o objc.IObject, options JSONWritingOptions) ([]byte, error) {
	var err Error
	r := objc.Call[[]byte](_NSJSONSerializationClass, objc.Sel("dataWithJSONObject:options:error:"), o, options, unsafe.Pointer(&err))
	return r, ToGoError(err)
}

func IsValidJSONObject(o objc.IObject) bool {
	return objc.Call[bool](_NSJSONSerializationClass, objc.Sel("isValidJSONObject:"), o)
}
