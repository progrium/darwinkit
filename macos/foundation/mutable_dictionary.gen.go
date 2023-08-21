// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MutableDictionary] class.
var MutableDictionaryClass = _MutableDictionaryClass{objc.GetClass("NSMutableDictionary")}

type _MutableDictionaryClass struct {
	objc.Class
}

// An interface definition for the [MutableDictionary] class.
type IMutableDictionary interface {
	IDictionary
	AddWhoHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objc.Object
	InitWithContentsOfFile(path string) MutableDictionary
	AddAuthorizationResponseHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objc.Object
	AddApplicationParameterHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objc.Object
	GetHeaderBytes() MutableData
	AddTargetHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objc.Object
	AddConnectionIDHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objc.Object
	AddImageHandleHeader(type_ string) objc.Object
	AddImageDescriptorHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objc.Object
	AddLengthHeader(length uint32) objc.Object
	AddEntriesFromDictionary(otherDictionary Dictionary)
	AddObjectClassHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objc.Object
	SetDictionary(otherDictionary Dictionary)
	SetObjectForKey(anObject objc.IObject, aKey PCopying)
	SetObjectForKeyObject(anObject objc.IObject, aKeyObject objc.IObject)
	AddTimeISOHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objc.Object
	AddTypeHeader(type_ string) objc.Object
	RemoveObjectsForKeys(keyArray []objc.IObject)
	AddAuthorizationChallengeHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objc.Object
	InitWithContentsOfURL(url IURL) MutableDictionary
	AddBodyHeaderLengthEndOfBody(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, isEndOfBody bool) objc.Object
	AddTime4ByteHeader(time4Byte uint32) objc.Object
	AddDescriptionHeader(inDescriptionString string) objc.Object
	RemoveObjectForKey(aKey objc.IObject)
	AddUserDefinedHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objc.Object
	AddNameHeader(inNameString string) objc.Object
	AddCountHeader(inCount uint32) objc.Object
	AddByteSequenceHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objc.Object
	SetValueForKey(value objc.IObject, key string)
	AddHTTPHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objc.Object
	RemoveAllObjects()
}

// A dynamic collection of objects associated with unique keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary?language=objc
type MutableDictionary struct {
	Dictionary
}

func MutableDictionaryFrom(ptr unsafe.Pointer) MutableDictionary {
	return MutableDictionary{
		Dictionary: DictionaryFrom(ptr),
	}
}

func (m_ MutableDictionary) InitWithCapacity(numItems uint) MutableDictionary {
	rv := objc.Call[MutableDictionary](m_, objc.Sel("initWithCapacity:"), numItems)
	return rv
}

// Initializes a newly allocated mutable dictionary, allocating enough memory to hold numItems entries. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1417898-initwithcapacity?language=objc
func NewMutableDictionaryWithCapacity(numItems uint) MutableDictionary {
	instance := MutableDictionaryClass.Alloc().InitWithCapacity(numItems)
	instance.Autorelease()
	return instance
}

func (mc _MutableDictionaryClass) DictionaryWithOBEXHeadersData(inHeadersData []byte) MutableDictionary {
	rv := objc.Call[MutableDictionary](mc, objc.Sel("dictionaryWithOBEXHeadersData:"), inHeadersData)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1428433-dictionarywithobexheadersdata?language=objc
func MutableDictionary_DictionaryWithOBEXHeadersData(inHeadersData []byte) MutableDictionary {
	return MutableDictionaryClass.DictionaryWithOBEXHeadersData(inHeadersData)
}

func (m_ MutableDictionary) Init() MutableDictionary {
	rv := objc.Call[MutableDictionary](m_, objc.Sel("init"))
	return rv
}

func (mc _MutableDictionaryClass) DictionaryWithCapacity(numItems uint) MutableDictionary {
	rv := objc.Call[MutableDictionary](mc, objc.Sel("dictionaryWithCapacity:"), numItems)
	return rv
}

// Creates and returns a mutable dictionary, initially giving it enough allocated memory to hold a given number of entries. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1574186-dictionarywithcapacity?language=objc
func MutableDictionary_DictionaryWithCapacity(numItems uint) MutableDictionary {
	return MutableDictionaryClass.DictionaryWithCapacity(numItems)
}

func (mc _MutableDictionaryClass) Alloc() MutableDictionary {
	rv := objc.Call[MutableDictionary](mc, objc.Sel("alloc"))
	return rv
}

func MutableDictionary_Alloc() MutableDictionary {
	return MutableDictionaryClass.Alloc()
}

func (mc _MutableDictionaryClass) New() MutableDictionary {
	rv := objc.Call[MutableDictionary](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMutableDictionary() MutableDictionary {
	return MutableDictionaryClass.New()
}

func (mc _MutableDictionaryClass) DictionaryWithObjectsForKeys(objects []objc.IObject, keys []PCopying) MutableDictionary {
	rv := objc.Call[MutableDictionary](mc, objc.Sel("dictionaryWithObjects:forKeys:"), objects, keys)
	return rv
}

// Creates a dictionary containing entries constructed from the contents of an array of keys and an array of values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1574183-dictionarywithobjects?language=objc
func MutableDictionary_DictionaryWithObjectsForKeys(objects []objc.IObject, keys []PCopying) MutableDictionary {
	return MutableDictionaryClass.DictionaryWithObjectsForKeys(objects, keys)
}

func (m_ MutableDictionary) InitWithObjectsAndKeys(firstObject objc.IObject, args ...any) MutableDictionary {
	rv := objc.Call[MutableDictionary](m_, objc.Sel("initWithObjectsAndKeys:"), append([]any{firstObject}, args...)...)
	return rv
}

// Initializes a newly allocated dictionary with entries constructed from the specified set of values and keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1574190-initwithobjectsandkeys?language=objc
func NewMutableDictionaryWithObjectsAndKeys(firstObject objc.IObject, args ...any) MutableDictionary {
	instance := MutableDictionaryClass.Alloc().InitWithObjectsAndKeys(firstObject, args...)
	instance.Autorelease()
	return instance
}

func (m_ MutableDictionary) InitWithObjectsForKeys(objects []objc.IObject, keys []PCopying) MutableDictionary {
	rv := objc.Call[MutableDictionary](m_, objc.Sel("initWithObjects:forKeys:"), objects, keys)
	return rv
}

// Initializes a newly allocated dictionary with key-value pairs constructed from the provided arrays of keys and objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1410010-initwithobjects?language=objc
func NewMutableDictionaryWithObjectsForKeys(objects []objc.IObject, keys []PCopying) MutableDictionary {
	instance := MutableDictionaryClass.Alloc().InitWithObjectsForKeys(objects, keys)
	instance.Autorelease()
	return instance
}

func (mc _MutableDictionaryClass) DictionaryWithObjectForKey(object objc.IObject, key PCopying) MutableDictionary {
	po1 := objc.WrapAsProtocol("NSCopying", key)
	rv := objc.Call[MutableDictionary](mc, objc.Sel("dictionaryWithObject:forKey:"), objc.Ptr(object), po1)
	return rv
}

// Creates a dictionary containing a given key and value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1414965-dictionarywithobject?language=objc
func MutableDictionary_DictionaryWithObjectForKey(object objc.IObject, key PCopying) MutableDictionary {
	return MutableDictionaryClass.DictionaryWithObjectForKey(object, key)
}

func (mc _MutableDictionaryClass) Dictionary() MutableDictionary {
	rv := objc.Call[MutableDictionary](mc, objc.Sel("dictionary"))
	return rv
}

// Creates an empty dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1574180-dictionary?language=objc
func MutableDictionary_Dictionary() MutableDictionary {
	return MutableDictionaryClass.Dictionary()
}

func (mc _MutableDictionaryClass) DictionaryWithObjectsAndKeys(firstObject objc.IObject, args ...any) MutableDictionary {
	rv := objc.Call[MutableDictionary](mc, objc.Sel("dictionaryWithObjectsAndKeys:"), append([]any{firstObject}, args...)...)
	return rv
}

// Creates a dictionary containing entries constructed from the specified set of values and keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1574181-dictionarywithobjectsandkeys?language=objc
func MutableDictionary_DictionaryWithObjectsAndKeys(firstObject objc.IObject, args ...any) MutableDictionary {
	return MutableDictionaryClass.DictionaryWithObjectsAndKeys(firstObject, args...)
}

func (mc _MutableDictionaryClass) DictionaryWithDictionary(dict Dictionary) MutableDictionary {
	rv := objc.Call[MutableDictionary](mc, objc.Sel("dictionaryWithDictionary:"), dict)
	return rv
}

// Creates a dictionary containing the keys and values from another given dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1574191-dictionarywithdictionary?language=objc
func MutableDictionary_DictionaryWithDictionary(dict Dictionary) MutableDictionary {
	return MutableDictionaryClass.DictionaryWithDictionary(dict)
}

func (m_ MutableDictionary) InitWithDictionary(otherDictionary Dictionary) MutableDictionary {
	rv := objc.Call[MutableDictionary](m_, objc.Sel("initWithDictionary:"), otherDictionary)
	return rv
}

// Initializes a newly allocated dictionary by placing in it the keys and values contained in another given dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1418434-initwithdictionary?language=objc
func NewMutableDictionaryWithDictionary(otherDictionary Dictionary) MutableDictionary {
	instance := MutableDictionaryClass.Alloc().InitWithDictionary(otherDictionary)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1433510-addwhoheader?language=objc
func (m_ MutableDictionary) AddWhoHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("addWhoHeader:length:"), inHeaderData, inHeaderDataLength)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1407593-initwithcontentsoffile?language=objc
func (m_ MutableDictionary) InitWithContentsOfFile(path string) MutableDictionary {
	rv := objc.Call[MutableDictionary](m_, objc.Sel("initWithContentsOfFile:"), path)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1432339-addauthorizationresponseheader?language=objc
func (m_ MutableDictionary) AddAuthorizationResponseHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("addAuthorizationResponseHeader:length:"), inHeaderData, inHeaderDataLength)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1431956-addapplicationparameterheader?language=objc
func (m_ MutableDictionary) AddApplicationParameterHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("addApplicationParameterHeader:length:"), inHeaderData, inHeaderDataLength)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1428890-getheaderbytes?language=objc
func (m_ MutableDictionary) GetHeaderBytes() MutableData {
	rv := objc.Call[MutableData](m_, objc.Sel("getHeaderBytes"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1428678-addtargetheader?language=objc
func (m_ MutableDictionary) AddTargetHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("addTargetHeader:length:"), inHeaderData, inHeaderDataLength)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1430403-addconnectionidheader?language=objc
func (m_ MutableDictionary) AddConnectionIDHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("addConnectionIDHeader:length:"), inHeaderData, inHeaderDataLength)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1430878-addimagehandleheader?language=objc
func (m_ MutableDictionary) AddImageHandleHeader(type_ string) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("addImageHandleHeader:"), type_)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1429678-addimagedescriptorheader?language=objc
func (m_ MutableDictionary) AddImageDescriptorHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("addImageDescriptorHeader:length:"), inHeaderData, inHeaderDataLength)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1433652-addlengthheader?language=objc
func (m_ MutableDictionary) AddLengthHeader(length uint32) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("addLengthHeader:"), length)
	return rv
}

// Adds to the receiving dictionary the entries from another dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1411035-addentriesfromdictionary?language=objc
func (m_ MutableDictionary) AddEntriesFromDictionary(otherDictionary Dictionary) {
	objc.Call[objc.Void](m_, objc.Sel("addEntriesFromDictionary:"), otherDictionary)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1434696-addobjectclassheader?language=objc
func (m_ MutableDictionary) AddObjectClassHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("addObjectClassHeader:length:"), inHeaderData, inHeaderDataLength)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1574188-dictionarywithcontentsoffile?language=objc
func (mc _MutableDictionaryClass) DictionaryWithContentsOfFile(path string) MutableDictionary {
	rv := objc.Call[MutableDictionary](mc, objc.Sel("dictionaryWithContentsOfFile:"), path)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1574188-dictionarywithcontentsoffile?language=objc
func MutableDictionary_DictionaryWithContentsOfFile(path string) MutableDictionary {
	return MutableDictionaryClass.DictionaryWithContentsOfFile(path)
}

// Sets the contents of the receiving dictionary to entries in a given dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1409566-setdictionary?language=objc
func (m_ MutableDictionary) SetDictionary(otherDictionary Dictionary) {
	objc.Call[objc.Void](m_, objc.Sel("setDictionary:"), otherDictionary)
}

// Adds a given key-value pair to the dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1411616-setobject?language=objc
func (m_ MutableDictionary) SetObjectForKey(anObject objc.IObject, aKey PCopying) {
	po1 := objc.WrapAsProtocol("NSCopying", aKey)
	objc.Call[objc.Void](m_, objc.Sel("setObject:forKey:"), objc.Ptr(anObject), po1)
}

// Adds a given key-value pair to the dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1411616-setobject?language=objc
func (m_ MutableDictionary) SetObjectForKeyObject(anObject objc.IObject, aKeyObject objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("setObject:forKey:"), objc.Ptr(anObject), objc.Ptr(aKeyObject))
}

// Creates a mutable dictionary which is optimized for dealing with a known set of keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1412658-dictionarywithsharedkeyset?language=objc
func (mc _MutableDictionaryClass) DictionaryWithSharedKeySet(keyset objc.IObject) MutableDictionary {
	rv := objc.Call[MutableDictionary](mc, objc.Sel("dictionaryWithSharedKeySet:"), keyset)
	return rv
}

// Creates a mutable dictionary which is optimized for dealing with a known set of keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1412658-dictionarywithsharedkeyset?language=objc
func MutableDictionary_DictionaryWithSharedKeySet(keyset objc.IObject) MutableDictionary {
	return MutableDictionaryClass.DictionaryWithSharedKeySet(keyset)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1428661-addtimeisoheader?language=objc
func (m_ MutableDictionary) AddTimeISOHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("addTimeISOHeader:length:"), inHeaderData, inHeaderDataLength)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1433211-addtypeheader?language=objc
func (m_ MutableDictionary) AddTypeHeader(type_ string) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("addTypeHeader:"), type_)
	return rv
}

// Removes from the dictionary entries specified by elements in a given array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1410430-removeobjectsforkeys?language=objc
func (m_ MutableDictionary) RemoveObjectsForKeys(keyArray []objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("removeObjectsForKeys:"), keyArray)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1433422-addauthorizationchallengeheader?language=objc
func (m_ MutableDictionary) AddAuthorizationChallengeHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("addAuthorizationChallengeHeader:length:"), inHeaderData, inHeaderDataLength)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1410409-initwithcontentsofurl?language=objc
func (m_ MutableDictionary) InitWithContentsOfURL(url IURL) MutableDictionary {
	rv := objc.Call[MutableDictionary](m_, objc.Sel("initWithContentsOfURL:"), objc.Ptr(url))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1430411-addbodyheader?language=objc
func (m_ MutableDictionary) AddBodyHeaderLengthEndOfBody(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, isEndOfBody bool) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("addBodyHeader:length:endOfBody:"), inHeaderData, inHeaderDataLength, isEndOfBody)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1574182-dictionarywithcontentsofurl?language=objc
func (mc _MutableDictionaryClass) DictionaryWithContentsOfURL(url IURL) MutableDictionary {
	rv := objc.Call[MutableDictionary](mc, objc.Sel("dictionaryWithContentsOfURL:"), objc.Ptr(url))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1574182-dictionarywithcontentsofurl?language=objc
func MutableDictionary_DictionaryWithContentsOfURL(url IURL) MutableDictionary {
	return MutableDictionaryClass.DictionaryWithContentsOfURL(url)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1433527-addtime4byteheader?language=objc
func (m_ MutableDictionary) AddTime4ByteHeader(time4Byte uint32) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("addTime4ByteHeader:"), time4Byte)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1434905-adddescriptionheader?language=objc
func (m_ MutableDictionary) AddDescriptionHeader(inDescriptionString string) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("addDescriptionHeader:"), inDescriptionString)
	return rv
}

// Removes a given key and its associated value from the dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1416518-removeobjectforkey?language=objc
func (m_ MutableDictionary) RemoveObjectForKey(aKey objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("removeObjectForKey:"), objc.Ptr(aKey))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1434461-adduserdefinedheader?language=objc
func (m_ MutableDictionary) AddUserDefinedHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("addUserDefinedHeader:length:"), inHeaderData, inHeaderDataLength)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1430284-addnameheader?language=objc
func (m_ MutableDictionary) AddNameHeader(inNameString string) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("addNameHeader:"), inNameString)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1433317-addcountheader?language=objc
func (m_ MutableDictionary) AddCountHeader(inCount uint32) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("addCountHeader:"), inCount)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1428676-addbytesequenceheader?language=objc
func (m_ MutableDictionary) AddByteSequenceHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("addByteSequenceHeader:length:"), inHeaderData, inHeaderDataLength)
	return rv
}

// Adds a given key-value pair to the dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1416335-setvalue?language=objc
func (m_ MutableDictionary) SetValueForKey(value objc.IObject, key string) {
	objc.Call[objc.Void](m_, objc.Sel("setValue:forKey:"), objc.Ptr(value), key)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1434484-addhttpheader?language=objc
func (m_ MutableDictionary) AddHTTPHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("addHTTPHeader:length:"), inHeaderData, inHeaderDataLength)
	return rv
}

// Empties the dictionary of its entries. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutabledictionary/1408955-removeallobjects?language=objc
func (m_ MutableDictionary) RemoveAllObjects() {
	objc.Call[objc.Void](m_, objc.Sel("removeAllObjects"))
}
