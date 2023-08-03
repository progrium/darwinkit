package foundation

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

type IString interface {
	objc.IObject
	String() string
}

type String struct {
	objc.Object
}

func MakeString(ptr unsafe.Pointer) String {
	return String{
		Object: objc.MakeObject(ptr),
	}
}

func NewString(str string) String {
	return MakeString(objc.ToNSString(str))
}

func (N String) String() string {
	return objc.ToGoString(N.Ptr())
}

func (N String) ToString() string {
	return objc.ToGoString(N.Ptr())
}

type IMutableString interface {
	IString
	AppendString(str string)
}

type MutableString struct {
	String
}

func MakeMutableString(ptr unsafe.Pointer) MutableString {
	return MutableString{
		String: MakeString(ptr),
	}
}

var _MutableStringClass = objc.GetClass("NSMutableString")

func NewMutableString() MutableString {
	return objc.CallMethod[MutableString](_MutableStringClass, objc.GetSelector("string"))
}

func NewMutableStringWithString(str string) MutableString {
	return objc.CallMethod[MutableString](_MutableStringClass, objc.GetSelector("stringWithString:"), str)
}

func NewMutableStringWithCapacity(capacity uint) MutableString {
	return objc.CallMethod[MutableString](_MutableStringClass, objc.GetSelector("stringWithCapacity:"), capacity)
}

func (s MutableString) AppendString(str string) {
	objc.CallMethod[objc.Void](s, objc.GetSelector("appendString:"), str)
}

type IData interface {
	objc.IObject
	ToBytes() []byte
}

type Data struct {
	objc.Object
}

func MakeData(ptr unsafe.Pointer) Data {
	return Data{
		Object: objc.MakeObject(ptr),
	}
}

func NewData(bytes []byte) Data {
	return MakeData(objc.ToNSData(bytes))
}

func (d Data) ToBytes() []byte {
	return objc.ToGoBytes(d.Ptr())
}

type IMutableData interface {
	IData
	AppendData(data []byte)
}

type MutableData struct {
	Data
}

func MakeMutableData(ptr unsafe.Pointer) MutableData {
	return MutableData{
		Data: MakeData(ptr),
	}
}

var _MutableDataClass = objc.GetClass("NSMutableData")

func NewMutableData() MutableData {
	return objc.CallMethod[MutableData](_MutableDataClass, objc.GetSelector("data"))
}

func NewMutableDataWithData(data []byte) MutableData {
	return objc.CallMethod[MutableData](_MutableDataClass, objc.GetSelector("dataWithData:"), data)
}

func NewMutableDataWithCapacity(capacity uint) MutableData {
	return objc.CallMethod[MutableData](_MutableDataClass, objc.GetSelector("dataWithCapacity:"), capacity)
}

func NewMutableDataWithLength(length uint) MutableData {
	return objc.CallMethod[MutableData](_MutableDataClass, objc.GetSelector("dataWithLength:"), length)
}

func (d MutableData) AppendData(data []byte) {
	objc.CallMethod[objc.Void](d, objc.GetSelector("appendData:"), data)
}

type ISet interface {
	objc.IObject
	Count() uint
	AllObjects() Array
}

type Set struct {
	objc.Object
}

func MakeSet(ptr unsafe.Pointer) Set {
	return Set{
		Object: objc.MakeObject(ptr),
	}
}

var _SetClass = objc.GetClass("NSSet")

func NewSet() Set {
	return objc.CallMethod[Set](_SetClass, objc.GetSelector("set"))
}

func NewSetWithArray(array IArray) Set {
	return objc.CallMethod[Set](_SetClass, objc.GetSelector("setWithArray:"), array)
}

func (s Set) Count() uint {
	return objc.CallMethod[uint](s, objc.GetSelector("count"))
}

func (s Set) AllObjects() Array {
	return objc.CallMethod[Array](s, objc.GetSelector("allObjects"))
}

type IMutableSet interface {
	ISet
	AddObject(o objc.IObject)
}

type MutableSet struct {
	Set
}

func MakeMutableSet(ptr unsafe.Pointer) MutableSet {
	return MutableSet{
		Set: MakeSet(ptr),
	}
}

var _MutableSetClass = objc.GetClass("NSMutableSet")

func NewMutableSet() MutableSet {
	return objc.CallMethod[MutableSet](_MutableSetClass, objc.GetSelector("set"))
}

func NewMutableSetWithCapacity(size uint) MutableSet {
	return objc.CallMethod[MutableSet](_MutableSetClass, objc.GetSelector("setWithCapacity:"), size)
}

func NewMutableSetWithArray(array IArray) Set {
	return objc.CallMethod[Set](_MutableSetClass, objc.GetSelector("setWithArray:"), array)
}

func (s MutableSet) AddObject(o objc.IObject) {
	objc.CallMethod[objc.Void](s, objc.GetSelector("addObject:"), o)
}

type IArray interface {
	objc.IObject
	ObjectAtIndex(index uint) objc.Object
	Count() uint
}

type Array struct {
	objc.Object
}

func MakeArray(ptr unsafe.Pointer) Array {
	return Array{
		Object: objc.MakeObject(ptr),
	}
}

func (a Array) ObjectAtIndex(index uint) objc.Object {
	return objc.CallMethod[objc.Object](a, objc.GetSelector("objectAtIndex:"), index)
}

func (a Array) Count() uint {
	return objc.CallMethod[uint](a, objc.GetSelector("count"))
}

func ArrayOf[T any](slice []T) Array {
	return MakeArray(objc.ToNSArray(reflect.ValueOf(slice)))
}

func ArrayToSlice[T any](a Array) []T {
	var zero []T
	return objc.ToGoSlice(a.Ptr(), reflect.TypeOf(zero)).Interface().([]T)
}

type IMutableArray interface {
	IArray
	AddObject(o objc.IObject)
	InsertObjectAtIndex(o objc.IObject, index uint)
	RemoveObjectAtIndex(index uint)
	ReplaceObjectAtIndex_WithObject(index uint, o objc.IObject)
}

type MutableArray struct {
	Array
}

func MakeMutableArray(ptr unsafe.Pointer) MutableArray {
	return MutableArray{
		Array: MakeArray(ptr),
	}
}

var _MutableArrayClass = objc.GetClass("NSMutableArray")

func NewMutableArray() MutableArray {
	return objc.CallMethod[MutableArray](_MutableArrayClass, objc.GetSelector("array"))
}

func NewMutableArrayWithCapacity(size uint) MutableArray {
	return objc.CallMethod[MutableArray](_MutableArrayClass, objc.GetSelector("arrayWithCapacity:"), size)
}

func NewMutableSetWithSet(set ISet) MutableSet {
	return objc.CallMethod[MutableSet](_MutableSetClass, objc.GetSelector("setWithSet:"), set)
}

func (a MutableArray) AddObject(o objc.IObject) {
	objc.CallMethod[objc.Void](a, objc.GetSelector("addObject:"), o)
}

func (a MutableArray) InsertObjectAtIndex(o objc.IObject, index uint) {
	objc.CallMethod[objc.Void](a, objc.GetSelector("insertObject:atIndex:"), o, index)
}

func (a MutableArray) RemoveObjectAtIndex(index uint) {
	objc.CallMethod[objc.Void](a, objc.GetSelector("removeObjectAtIndex:"), index)
}

func (a MutableArray) ReplaceObjectAtIndex_WithObject(index uint, o objc.IObject) {
	objc.CallMethod[objc.Void](a, objc.GetSelector("replaceObjectAtIndex:withObject:"), index, o)
}

type IDictionary interface {
	objc.IObject
	ObjectForKey(key objc.IObject) objc.Object
	AllKeys() Array
	AllValues() Array
	Count() uint
}

type Dictionary struct {
	objc.Object
}

func MakeDictionary(ptr unsafe.Pointer) Dictionary {
	return Dictionary{
		Object: objc.MakeObject(ptr),
	}
}

func (d Dictionary) ObjectForKey(key objc.IObject) objc.Object {
	return objc.CallMethod[objc.Object](d, objc.GetSelector("objectForKey:"), key)
}

func (d Dictionary) AllKeys() Array {
	return objc.CallMethod[Array](d, objc.GetSelector("allKeys"))
}

func (d Dictionary) AllValues() Array {
	return objc.CallMethod[Array](d, objc.GetSelector("allValues"))
}

func (d Dictionary) Count() uint {
	return objc.CallMethod[uint](d, objc.GetSelector("count"))
}

func DictOf[K comparable, V any](m map[K]V) Dictionary {
	return MakeDictionary(objc.ToNSDict(reflect.ValueOf(m)))
}

func DictToMap[K comparable, V any](d Dictionary) map[K]V {
	var zero map[K]V
	return objc.ToGoMap(d.Ptr(), reflect.TypeOf(zero)).Interface().(map[K]V)
}

type IMutableDictionary interface {
	IDictionary
	SetObjectForKey(value objc.IObject, key objc.IObject)
	RemoveObjectForKey(key objc.IObject)
}

type MutableDictionary struct {
	Dictionary
}

func MakeMutableDictionary(ptr unsafe.Pointer) MutableDictionary {
	return MutableDictionary{
		Dictionary: MakeDictionary(ptr),
	}
}

var _MutableDictionaryClass = objc.GetClass("NSMutableDictionary")

func NewMutableDictionary() MutableDictionary {
	return objc.CallMethod[MutableDictionary](_MutableDictionaryClass, objc.GetSelector("dictionary"))
}

func NewMutableDictionaryWithDictionary(dict IDictionary) MutableDictionary {
	return objc.CallMethod[MutableDictionary](_MutableDictionaryClass, objc.GetSelector("dictionaryWithDictionary:"), dict)
}

func (d MutableDictionary) SetObjectForKey(value objc.IObject, key objc.IObject) {
	objc.CallMethod[objc.Void](d, objc.GetSelector("setObject:forKey:"), value, key)
}

func (d MutableDictionary) RemoveObjectForKey(key objc.IObject) {
	objc.CallMethod[objc.Void](d, objc.GetSelector("removeObjectsForKeys:"), key)
}

func ToNSError(code int, err error) Error {
	if err == nil {
		return Error{}
	}
	return ErrorClass.ErrorWithDomainCodeUserInfo("cocoa.hsiafan.github.com", code, map[ErrorUserInfoKey]objc.IObject{
		"Error reason": NewString(err.Error()),
	})
}

// Error convert

var _ error = ocErr{}

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

// JSONSerialization begin
type JSONReadingOptions uint

const (
	JSONReadingMutableContainers         JSONReadingOptions = 1 << 0
	JSONReadingMutableLeaves             JSONReadingOptions = 1 << 1
	JSONReadingFragmentsAllowed          JSONReadingOptions = 1 << 2
	JSONReadingJSON5Allowed              JSONReadingOptions = 1 << 3
	JSONReadingTopLevelDictionaryAssumed JSONReadingOptions = 1 << 4
)

type JSONWritingOptions uint

const (
	JSONWritingPrettyPrinted          JSONWritingOptions = 1 << 0
	JSONWritingSortedKeys             JSONWritingOptions = 1 << 1
	JSONWritingFragmentsAllowed       JSONWritingOptions = 1 << 2
	JSONWritingWithoutEscapingSlashes JSONWritingOptions = 1 << 3
)

var _NSJSONSerializationClass = objc.GetClass("NSJSONSerialization")

func JSONObjectWithData(data []byte, options JSONReadingOptions) (objc.Object, error) {
	var err Error
	r := objc.CallMethod[objc.Object](_NSJSONSerializationClass, objc.GetSelector("JSONObjectWithData:options:error:"), data, options, unsafe.Pointer(&err))
	return r, ToGoError(err)
}

func DataWithJSONObject(o objc.IObject, options JSONWritingOptions) ([]byte, error) {
	var err Error
	r := objc.CallMethod[[]byte](_NSJSONSerializationClass, objc.GetSelector("dataWithJSONObject:options:error:"), o, options, unsafe.Pointer(&err))
	return r, ToGoError(err)
}

func IsValidJSONObject(o objc.IObject) bool {
	return objc.CallMethod[bool](_NSJSONSerializationClass, objc.GetSelector("isValidJSONObject:"), o)
}
