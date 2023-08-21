// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [KeyedUnarchiver] class.
var KeyedUnarchiverClass = _KeyedUnarchiverClass{objc.GetClass("NSKeyedUnarchiver")}

type _KeyedUnarchiverClass struct {
	objc.Class
}

// An interface definition for the [KeyedUnarchiver] class.
type IKeyedUnarchiver interface {
	ICoder
	SetClassForClassName(cls objc.IClass, codedName string)
	FinishDecoding()
	SetDecodingFailurePolicy(value DecodingFailurePolicy)
	Delegate() KeyedUnarchiverDelegateWrapper
	SetDelegate(value PKeyedUnarchiverDelegate)
	SetDelegateObject(valueObject objc.IObject)
	SetRequiresSecureCoding(value bool)
}

// A decoder that restores data from an archive referenced by keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiver?language=objc
type KeyedUnarchiver struct {
	Coder
}

func KeyedUnarchiverFrom(ptr unsafe.Pointer) KeyedUnarchiver {
	return KeyedUnarchiver{
		Coder: CoderFrom(ptr),
	}
}

func (k_ KeyedUnarchiver) InitForReadingFromDataError(data []byte, error IError) KeyedUnarchiver {
	rv := objc.Call[KeyedUnarchiver](k_, objc.Sel("initForReadingFromData:error:"), data, objc.Ptr(error))
	return rv
}

// Initializes an archiver to decode data from the specified location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiver/2962883-initforreadingfromdata?language=objc
func NewKeyedUnarchiverForReadingFromDataError(data []byte, error IError) KeyedUnarchiver {
	instance := KeyedUnarchiverClass.Alloc().InitForReadingFromDataError(data, error)
	instance.Autorelease()
	return instance
}

func (kc _KeyedUnarchiverClass) Alloc() KeyedUnarchiver {
	rv := objc.Call[KeyedUnarchiver](kc, objc.Sel("alloc"))
	return rv
}

func KeyedUnarchiver_Alloc() KeyedUnarchiver {
	return KeyedUnarchiverClass.Alloc()
}

func (kc _KeyedUnarchiverClass) New() KeyedUnarchiver {
	rv := objc.Call[KeyedUnarchiver](kc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewKeyedUnarchiver() KeyedUnarchiver {
	return KeyedUnarchiverClass.New()
}

func (k_ KeyedUnarchiver) Init() KeyedUnarchiver {
	rv := objc.Call[KeyedUnarchiver](k_, objc.Sel("init"))
	return rv
}

// Decodes a previously-archived object graph, that returns the root object as the specified type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiver/2962884-unarchivedobjectofclass?language=objc
func (kc _KeyedUnarchiverClass) UnarchivedObjectOfClassFromDataError(cls objc.IClass, data []byte, error IError) objc.Object {
	rv := objc.Call[objc.Object](kc, objc.Sel("unarchivedObjectOfClass:fromData:error:"), objc.Ptr(cls), data, objc.Ptr(error))
	return rv
}

// Decodes a previously-archived object graph, that returns the root object as the specified type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiver/2962884-unarchivedobjectofclass?language=objc
func KeyedUnarchiver_UnarchivedObjectOfClassFromDataError(cls objc.IClass, data []byte, error IError) objc.Object {
	return KeyedUnarchiverClass.UnarchivedObjectOfClassFromDataError(cls, data, error)
}

// Returns the class from which this unarchiver instantiates an encoded object with a given class name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiver/1416807-classforclassname?language=objc
func (kc _KeyedUnarchiverClass) ClassForClassName(codedName string) objc.Class {
	rv := objc.Call[objc.Class](kc, objc.Sel("classForClassName:"), codedName)
	return rv
}

// Returns the class from which this unarchiver instantiates an encoded object with a given class name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiver/1416807-classforclassname?language=objc
func KeyedUnarchiver_ClassForClassName(codedName string) objc.Class {
	return KeyedUnarchiverClass.ClassForClassName(codedName)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiver/3563983-unarchiveddictionarywithkeysofcl?language=objc
func (kc _KeyedUnarchiverClass) UnarchivedDictionaryWithKeysOfClassObjectsOfClassFromDataError(keyCls objc.IClass, valueCls objc.IClass, data []byte, error IError) Dictionary {
	rv := objc.Call[Dictionary](kc, objc.Sel("unarchivedDictionaryWithKeysOfClass:objectsOfClass:fromData:error:"), objc.Ptr(keyCls), objc.Ptr(valueCls), data, objc.Ptr(error))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiver/3563983-unarchiveddictionarywithkeysofcl?language=objc
func KeyedUnarchiver_UnarchivedDictionaryWithKeysOfClassObjectsOfClassFromDataError(keyCls objc.IClass, valueCls objc.IClass, data []byte, error IError) Dictionary {
	return KeyedUnarchiverClass.UnarchivedDictionaryWithKeysOfClassObjectsOfClassFromDataError(keyCls, valueCls, data, error)
}

// Sets a translation mapping on this unarchiver to decode objects encoded with a given class name as instances of a given class instead. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiver/1414659-setclass?language=objc
func (k_ KeyedUnarchiver) SetClassForClassName(cls objc.IClass, codedName string) {
	objc.Call[objc.Void](k_, objc.Sel("setClass:forClassName:"), objc.Ptr(cls), codedName)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiver/3563981-unarchivedarrayofobjectsofclass?language=objc
func (kc _KeyedUnarchiverClass) UnarchivedArrayOfObjectsOfClassFromDataError(cls objc.IClass, data []byte, error IError) []objc.Object {
	rv := objc.Call[[]objc.Object](kc, objc.Sel("unarchivedArrayOfObjectsOfClass:fromData:error:"), objc.Ptr(cls), data, objc.Ptr(error))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiver/3563981-unarchivedarrayofobjectsofclass?language=objc
func KeyedUnarchiver_UnarchivedArrayOfObjectsOfClassFromDataError(cls objc.IClass, data []byte, error IError) []objc.Object {
	return KeyedUnarchiverClass.UnarchivedArrayOfObjectsOfClassFromDataError(cls, data, error)
}

// Decodes a previously-archived object graph, returning the root object as one of the specified classes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiver/2962885-unarchivedobjectofclasses?language=objc
func (kc _KeyedUnarchiverClass) UnarchivedObjectOfClassesFromDataError(classes ISet, data []byte, error IError) objc.Object {
	rv := objc.Call[objc.Object](kc, objc.Sel("unarchivedObjectOfClasses:fromData:error:"), objc.Ptr(classes), data, objc.Ptr(error))
	return rv
}

// Decodes a previously-archived object graph, returning the root object as one of the specified classes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiver/2962885-unarchivedobjectofclasses?language=objc
func KeyedUnarchiver_UnarchivedObjectOfClassesFromDataError(classes ISet, data []byte, error IError) objc.Object {
	return KeyedUnarchiverClass.UnarchivedObjectOfClassesFromDataError(classes, data, error)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiver/3563982-unarchivedarrayofobjectsofclasse?language=objc
func (kc _KeyedUnarchiverClass) UnarchivedArrayOfObjectsOfClassesFromDataError(classes ISet, data []byte, error IError) []objc.Object {
	rv := objc.Call[[]objc.Object](kc, objc.Sel("unarchivedArrayOfObjectsOfClasses:fromData:error:"), objc.Ptr(classes), data, objc.Ptr(error))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiver/3563982-unarchivedarrayofobjectsofclasse?language=objc
func KeyedUnarchiver_UnarchivedArrayOfObjectsOfClassesFromDataError(classes ISet, data []byte, error IError) []objc.Object {
	return KeyedUnarchiverClass.UnarchivedArrayOfObjectsOfClassesFromDataError(classes, data, error)
}

// Tells the receiver that you are finished decoding objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiver/1418233-finishdecoding?language=objc
func (k_ KeyedUnarchiver) FinishDecoding() {
	objc.Call[objc.Void](k_, objc.Sel("finishDecoding"))
}

// The action to take when this unarchiver fails to decode an entry. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiver/1643164-decodingfailurepolicy?language=objc
func (k_ KeyedUnarchiver) SetDecodingFailurePolicy(value DecodingFailurePolicy) {
	objc.Call[objc.Void](k_, objc.Sel("setDecodingFailurePolicy:"), value)
}

// The receiver’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiver/1415688-delegate?language=objc
func (k_ KeyedUnarchiver) Delegate() KeyedUnarchiverDelegateWrapper {
	rv := objc.Call[KeyedUnarchiverDelegateWrapper](k_, objc.Sel("delegate"))
	return rv
}

// The receiver’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiver/1415688-delegate?language=objc
func (k_ KeyedUnarchiver) SetDelegate(value PKeyedUnarchiverDelegate) {
	po0 := objc.WrapAsProtocol("NSKeyedUnarchiverDelegate", value)
	objc.Call[objc.Void](k_, objc.Sel("setDelegate:"), po0)
}

// The receiver’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiver/1415688-delegate?language=objc
func (k_ KeyedUnarchiver) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](k_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// Indicates whether the receiver requires all unarchived classes to conform to NSSecureCoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nskeyedunarchiver/1410824-requiressecurecoding?language=objc
func (k_ KeyedUnarchiver) SetRequiresSecureCoding(value bool) {
	objc.Call[objc.Void](k_, objc.Sel("setRequiresSecureCoding:"), value)
}
