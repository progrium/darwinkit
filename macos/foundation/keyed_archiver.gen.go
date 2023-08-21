// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [KeyedArchiver] class.
var KeyedArchiverClass = _KeyedArchiverClass{objc.GetClass("NSKeyedArchiver")}

type _KeyedArchiverClass struct {
	objc.Class
}

// An interface definition for the [KeyedArchiver] class.
type IKeyedArchiver interface {
	ICoder
	EncodeObjectForKey(object objc.IObject, key string)
	EncodeConditionalObjectForKey(object objc.IObject, key string)
	FinishEncoding()
	ClassNameForClass(cls objc.IClass) string
	SetClassNameForClass(codedName string, cls objc.IClass)
	EncodeBytesLengthForKey(bytes *uint8, length uint, key string)
	OutputFormat() PropertyListFormat
	SetOutputFormat(value PropertyListFormat)
	Delegate() KeyedArchiverDelegateWrapper
	SetDelegate(value PKeyedArchiverDelegate)
	SetDelegateObject(valueObject objc.IObject)
	EncodedData() []byte
	SetRequiresSecureCoding(value bool)
}

// An encoder that stores an object’s data to an archive referenced by keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiver?language=objc
type KeyedArchiver struct {
	Coder
}

func KeyedArchiverFrom(ptr unsafe.Pointer) KeyedArchiver {
	return KeyedArchiver{
		Coder: CoderFrom(ptr),
	}
}

func (k_ KeyedArchiver) InitRequiringSecureCoding(requiresSecureCoding bool) KeyedArchiver {
	rv := objc.Call[KeyedArchiver](k_, objc.Sel("initRequiringSecureCoding:"), requiresSecureCoding)
	return rv
}

// Creates an archiver to encode data, and optionally disables secure coding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiver/2962881-initrequiringsecurecoding?language=objc
func NewKeyedArchiverRequiringSecureCoding(requiresSecureCoding bool) KeyedArchiver {
	instance := KeyedArchiverClass.Alloc().InitRequiringSecureCoding(requiresSecureCoding)
	instance.Autorelease()
	return instance
}

func (kc _KeyedArchiverClass) Alloc() KeyedArchiver {
	rv := objc.Call[KeyedArchiver](kc, objc.Sel("alloc"))
	return rv
}

func KeyedArchiver_Alloc() KeyedArchiver {
	return KeyedArchiverClass.Alloc()
}

func (kc _KeyedArchiverClass) New() KeyedArchiver {
	rv := objc.Call[KeyedArchiver](kc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewKeyedArchiver() KeyedArchiver {
	return KeyedArchiverClass.New()
}

func (k_ KeyedArchiver) Init() KeyedArchiver {
	rv := objc.Call[KeyedArchiver](k_, objc.Sel("init"))
	return rv
}

// Encodes a given object and associates it with a given key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiver/1412067-encodeobject?language=objc
func (k_ KeyedArchiver) EncodeObjectForKey(object objc.IObject, key string) {
	objc.Call[objc.Void](k_, objc.Sel("encodeObject:forKey:"), object, key)
}

// Encodes a reference to a given object and associates it with a key only if it has been unconditionally encoded elsewhere in the archive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiver/1413677-encodeconditionalobject?language=objc
func (k_ KeyedArchiver) EncodeConditionalObjectForKey(object objc.IObject, key string) {
	objc.Call[objc.Void](k_, objc.Sel("encodeConditionalObject:forKey:"), object, key)
}

// Instructs the receiver to construct the final data stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiver/1413904-finishencoding?language=objc
func (k_ KeyedArchiver) FinishEncoding() {
	objc.Call[objc.Void](k_, objc.Sel("finishEncoding"))
}

// Returns the class name with which this archiver encodes instances of a given class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiver/1407245-classnameforclass?language=objc
func (k_ KeyedArchiver) ClassNameForClass(cls objc.IClass) string {
	rv := objc.Call[string](k_, objc.Sel("classNameForClass:"), objc.Ptr(cls))
	return rv
}

// Sets a mapping for this archiver to encode instances of a given class with the provided name, rather than their real name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiver/1414746-setclassname?language=objc
func (k_ KeyedArchiver) SetClassNameForClass(codedName string, cls objc.IClass) {
	objc.Call[objc.Void](k_, objc.Sel("setClassName:forClass:"), codedName, objc.Ptr(cls))
}

// Encodes a given number of bytes from a given C array of bytes and associates them with a key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiver/1417696-encodebytes?language=objc
func (k_ KeyedArchiver) EncodeBytesLengthForKey(bytes *uint8, length uint, key string) {
	objc.Call[objc.Void](k_, objc.Sel("encodeBytes:length:forKey:"), bytes, length, key)
}

// Encodes an object graph with the given root object into a data representation, optionally requiring secure coding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiver/2962880-archiveddatawithrootobject?language=objc
func (kc _KeyedArchiverClass) ArchivedDataWithRootObjectRequiringSecureCodingError(object objc.IObject, requiresSecureCoding bool, error IError) []byte {
	rv := objc.Call[[]byte](kc, objc.Sel("archivedDataWithRootObject:requiringSecureCoding:error:"), object, requiresSecureCoding, objc.Ptr(error))
	return rv
}

// Encodes an object graph with the given root object into a data representation, optionally requiring secure coding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiver/2962880-archiveddatawithrootobject?language=objc
func KeyedArchiver_ArchivedDataWithRootObjectRequiringSecureCodingError(object objc.IObject, requiresSecureCoding bool, error IError) []byte {
	return KeyedArchiverClass.ArchivedDataWithRootObjectRequiringSecureCodingError(object, requiresSecureCoding, error)
}

// The format in which the receiver encodes its data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiver/1417520-outputformat?language=objc
func (k_ KeyedArchiver) OutputFormat() PropertyListFormat {
	rv := objc.Call[PropertyListFormat](k_, objc.Sel("outputFormat"))
	return rv
}

// The format in which the receiver encodes its data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiver/1417520-outputformat?language=objc
func (k_ KeyedArchiver) SetOutputFormat(value PropertyListFormat) {
	objc.Call[objc.Void](k_, objc.Sel("setOutputFormat:"), value)
}

// The archiver’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiver/1412809-delegate?language=objc
func (k_ KeyedArchiver) Delegate() KeyedArchiverDelegateWrapper {
	rv := objc.Call[KeyedArchiverDelegateWrapper](k_, objc.Sel("delegate"))
	return rv
}

// The archiver’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiver/1412809-delegate?language=objc
func (k_ KeyedArchiver) SetDelegate(value PKeyedArchiverDelegate) {
	po0 := objc.WrapAsProtocol("NSKeyedArchiverDelegate", value)
	objc.Call[objc.Void](k_, objc.Sel("setDelegate:"), po0)
}

// The archiver’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiver/1412809-delegate?language=objc
func (k_ KeyedArchiver) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](k_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The encoded data for the archiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiver/1643042-encodeddata?language=objc
func (k_ KeyedArchiver) EncodedData() []byte {
	rv := objc.Call[[]byte](k_, objc.Sel("encodedData"))
	return rv
}

// Indicates whether the archiver requires all archived classes to resist object substitution attacks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedarchiver/1417084-requiressecurecoding?language=objc
func (k_ KeyedArchiver) SetRequiresSecureCoding(value bool) {
	objc.Call[objc.Void](k_, objc.Sel("setRequiresSecureCoding:"), value)
}
