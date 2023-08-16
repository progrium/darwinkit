// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UbiquitousKeyValueStore] class.
var UbiquitousKeyValueStoreClass = _UbiquitousKeyValueStoreClass{objc.GetClass("NSUbiquitousKeyValueStore")}

type _UbiquitousKeyValueStoreClass struct {
	objc.Class
}

// An interface definition for the [UbiquitousKeyValueStore] class.
type IUbiquitousKeyValueStore interface {
	objc.IObject
	ObjectForKey(aKey string) objc.Object
	SetLongLongForKey(value int64, aKey string)
	SetBoolForKey(value bool, aKey string)
	ArrayForKey(aKey string) []objc.Object
	StringForKey(aKey string) string
	BoolForKey(aKey string) bool
	SetDictionaryForKey(aDictionary map[string]objc.IObject, aKey string)
	SetStringForKey(aString string, aKey string)
	SetObjectForKey(anObject objc.IObject, aKey string)
	SetArrayForKey(anArray []objc.IObject, aKey string)
	SetDoubleForKey(value float64, aKey string)
	DataForKey(aKey string) []byte
	DoubleForKey(aKey string) float64
	SetDataForKey(aData []byte, aKey string)
	DictionaryForKey(aKey string) map[string]objc.Object
	Synchronize() bool
	RemoveObjectForKey(aKey string)
	LongLongForKey(aKey string) int64
	DictionaryRepresentation() map[string]objc.Object
}

// An iCloud-based container of key-value pairs you use to share data among instances of your app running on a user's connected devices. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsubiquitouskeyvaluestore?language=objc
type UbiquitousKeyValueStore struct {
	objc.Object
}

func UbiquitousKeyValueStoreFrom(ptr unsafe.Pointer) UbiquitousKeyValueStore {
	return UbiquitousKeyValueStore{
		Object: objc.ObjectFrom(ptr),
	}
}

func (uc _UbiquitousKeyValueStoreClass) Alloc() UbiquitousKeyValueStore {
	rv := objc.Call[UbiquitousKeyValueStore](uc, objc.Sel("alloc"))
	return rv
}

func UbiquitousKeyValueStore_Alloc() UbiquitousKeyValueStore {
	return UbiquitousKeyValueStoreClass.Alloc()
}

func (uc _UbiquitousKeyValueStoreClass) New() UbiquitousKeyValueStore {
	rv := objc.Call[UbiquitousKeyValueStore](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUbiquitousKeyValueStore() UbiquitousKeyValueStore {
	return UbiquitousKeyValueStoreClass.New()
}

func (u_ UbiquitousKeyValueStore) Init() UbiquitousKeyValueStore {
	rv := objc.Call[UbiquitousKeyValueStore](u_, objc.Sel("init"))
	return rv
}

// Returns the object associated with the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsubiquitouskeyvaluestore/1407436-objectforkey?language=objc
func (u_ UbiquitousKeyValueStore) ObjectForKey(aKey string) objc.Object {
	rv := objc.Call[objc.Object](u_, objc.Sel("objectForKey:"), aKey)
	return rv
}

// Sets a long long value for the specified key in the key-value store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsubiquitouskeyvaluestore/1407812-setlonglong?language=objc
func (u_ UbiquitousKeyValueStore) SetLongLongForKey(value int64, aKey string) {
	objc.Call[objc.Void](u_, objc.Sel("setLongLong:forKey:"), value, aKey)
}

// Sets a Boolean value for the specified key in the key-value store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsubiquitouskeyvaluestore/1415054-setbool?language=objc
func (u_ UbiquitousKeyValueStore) SetBoolForKey(value bool, aKey string) {
	objc.Call[objc.Void](u_, objc.Sel("setBool:forKey:"), value, aKey)
}

// Returns the array associated with the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsubiquitouskeyvaluestore/1412191-arrayforkey?language=objc
func (u_ UbiquitousKeyValueStore) ArrayForKey(aKey string) []objc.Object {
	rv := objc.Call[[]objc.Object](u_, objc.Sel("arrayForKey:"), aKey)
	return rv
}

// Returns the string associated with the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsubiquitouskeyvaluestore/1418249-stringforkey?language=objc
func (u_ UbiquitousKeyValueStore) StringForKey(aKey string) string {
	rv := objc.Call[string](u_, objc.Sel("stringForKey:"), aKey)
	return rv
}

// Returns the Boolean value associated with the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsubiquitouskeyvaluestore/1417350-boolforkey?language=objc
func (u_ UbiquitousKeyValueStore) BoolForKey(aKey string) bool {
	rv := objc.Call[bool](u_, objc.Sel("boolForKey:"), aKey)
	return rv
}

// Sets a dictionary object for the specified key in the key-value store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsubiquitouskeyvaluestore/1417155-setdictionary?language=objc
func (u_ UbiquitousKeyValueStore) SetDictionaryForKey(aDictionary map[string]objc.IObject, aKey string) {
	objc.Call[objc.Void](u_, objc.Sel("setDictionary:forKey:"), aDictionary, aKey)
}

// Sets a string object for the specified key in the key-value store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsubiquitouskeyvaluestore/1414610-setstring?language=objc
func (u_ UbiquitousKeyValueStore) SetStringForKey(aString string, aKey string) {
	objc.Call[objc.Void](u_, objc.Sel("setString:forKey:"), aString, aKey)
}

// Sets an object for the specified key in the key-value store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsubiquitouskeyvaluestore/1416434-setobject?language=objc
func (u_ UbiquitousKeyValueStore) SetObjectForKey(anObject objc.IObject, aKey string) {
	objc.Call[objc.Void](u_, objc.Sel("setObject:forKey:"), anObject, aKey)
}

// Sets an array object for the specified key in the key-value store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsubiquitouskeyvaluestore/1417721-setarray?language=objc
func (u_ UbiquitousKeyValueStore) SetArrayForKey(anArray []objc.IObject, aKey string) {
	objc.Call[objc.Void](u_, objc.Sel("setArray:forKey:"), anArray, aKey)
}

// Sets a double value for the specified key in the key-value store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsubiquitouskeyvaluestore/1412608-setdouble?language=objc
func (u_ UbiquitousKeyValueStore) SetDoubleForKey(value float64, aKey string) {
	objc.Call[objc.Void](u_, objc.Sel("setDouble:forKey:"), value, aKey)
}

// Returns the data object associated with the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsubiquitouskeyvaluestore/1417800-dataforkey?language=objc
func (u_ UbiquitousKeyValueStore) DataForKey(aKey string) []byte {
	rv := objc.Call[[]byte](u_, objc.Sel("dataForKey:"), aKey)
	return rv
}

// Returns the double value associated with the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsubiquitouskeyvaluestore/1409319-doubleforkey?language=objc
func (u_ UbiquitousKeyValueStore) DoubleForKey(aKey string) float64 {
	rv := objc.Call[float64](u_, objc.Sel("doubleForKey:"), aKey)
	return rv
}

// Sets a data object for the specified key in the key-value store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsubiquitouskeyvaluestore/1416218-setdata?language=objc
func (u_ UbiquitousKeyValueStore) SetDataForKey(aData []byte, aKey string) {
	objc.Call[objc.Void](u_, objc.Sel("setData:forKey:"), aData, aKey)
}

// Returns the dictionary object associated with the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsubiquitouskeyvaluestore/1416241-dictionaryforkey?language=objc
func (u_ UbiquitousKeyValueStore) DictionaryForKey(aKey string) map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](u_, objc.Sel("dictionaryForKey:"), aKey)
	return rv
}

// Explicitly synchronizes in-memory keys and values with those stored on disk. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsubiquitouskeyvaluestore/1415989-synchronize?language=objc
func (u_ UbiquitousKeyValueStore) Synchronize() bool {
	rv := objc.Call[bool](u_, objc.Sel("synchronize"))
	return rv
}

// Removes the value associated with the specified key from the key-value store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsubiquitouskeyvaluestore/1410916-removeobjectforkey?language=objc
func (u_ UbiquitousKeyValueStore) RemoveObjectForKey(aKey string) {
	objc.Call[objc.Void](u_, objc.Sel("removeObjectForKey:"), aKey)
}

// Returns the long long value associated with the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsubiquitouskeyvaluestore/1413240-longlongforkey?language=objc
func (u_ UbiquitousKeyValueStore) LongLongForKey(aKey string) int64 {
	rv := objc.Call[int64](u_, objc.Sel("longLongForKey:"), aKey)
	return rv
}

// Returns the shared iCloud key-value store object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsubiquitouskeyvaluestore/1413949-defaultstore?language=objc
func (uc _UbiquitousKeyValueStoreClass) DefaultStore() UbiquitousKeyValueStore {
	rv := objc.Call[UbiquitousKeyValueStore](uc, objc.Sel("defaultStore"))
	return rv
}

// Returns the shared iCloud key-value store object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsubiquitouskeyvaluestore/1413949-defaultstore?language=objc
func UbiquitousKeyValueStore_DefaultStore() UbiquitousKeyValueStore {
	return UbiquitousKeyValueStoreClass.DefaultStore()
}

// A dictionary containing all of the key-value pairs in the key-value store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsubiquitouskeyvaluestore/1411129-dictionaryrepresentation?language=objc
func (u_ UbiquitousKeyValueStore) DictionaryRepresentation() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](u_, objc.Sel("dictionaryRepresentation"))
	return rv
}
