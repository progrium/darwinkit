// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ManagedObject] class.
var ManagedObjectClass = _ManagedObjectClass{objc.GetClass("NSManagedObject")}

type _ManagedObjectClass struct {
	objc.Class
}

// An interface definition for the [ManagedObject] class.
type IManagedObject interface {
	objc.IObject
	PrimitiveValueForKey(key string) objc.Object
	DidChangeValueForKey(key string)
	ValidateValueForKeyError(value objc.IObject, key string, error foundation.IError) bool
	DidAccessValueForKey(key string)
	ObservationInfo() unsafe.Pointer
	ChangedValues() map[string]objc.Object
	AwakeFromInsert()
	ObjectIDsForRelationshipNamed(key string) []ManagedObjectID
	WillSave()
	DidSave()
	ValidateForDelete(error foundation.IError) bool
	SetPrimitiveValueForKey(value objc.IObject, key string)
	WillTurnIntoFault()
	AwakeFromFetch()
	ValidateForInsert(error foundation.IError) bool
	PrepareForDeletion()
	ValidateForUpdate(error foundation.IError) bool
	AwakeFromSnapshotEvents(flags SnapshotEventType)
	CommittedValuesForKeys(keys []string) map[string]objc.Object
	HasFaultForRelationshipNamed(key string) bool
	WillAccessValueForKey(key string)
	DidTurnIntoFault()
	SetValueForKey(value objc.IObject, key string)
	ChangedValuesForCurrentEvent() map[string]objc.Object
	WillChangeValueForKey(key string)
	SetObservationInfo(inObservationInfo unsafe.Pointer)
	ValueForKey(key string) objc.Object
	InitWithEntityInsertIntoManagedObjectContext(entity IEntityDescription, context IManagedObjectContext) ManagedObject
	Entity() EntityDescription
	IsInserted() bool
	HasPersistentChangedValues() bool
	FaultingState() uint
	IsUpdated() bool
	IsDeleted() bool
	ObjectID() ManagedObjectID
	ManagedObjectContext() ManagedObjectContext
	HasChanges() bool
	IsFault() bool
}

// The base class that all Core Data model objects inherit from. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject?language=objc
type ManagedObject struct {
	objc.Object
}

func ManagedObjectFrom(ptr unsafe.Pointer) ManagedObject {
	return ManagedObject{
		Object: objc.ObjectFrom(ptr),
	}
}

func (m_ ManagedObject) InitWithContext(moc IManagedObjectContext) ManagedObject {
	rv := objc.Call[ManagedObject](m_, objc.Sel("initWithContext:"), objc.Ptr(moc))
	return rv
}

// Initializes a managed object subclass and inserts it into the specified managed object context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1640602-initwithcontext?language=objc
func ManagedObject_InitWithContext(moc IManagedObjectContext) ManagedObject {
	return ManagedObjectClass.Alloc().InitWithContext(moc)
}

func (mc _ManagedObjectClass) Alloc() ManagedObject {
	rv := objc.Call[ManagedObject](mc, objc.Sel("alloc"))
	return rv
}

func ManagedObject_Alloc() ManagedObject {
	return ManagedObjectClass.Alloc()
}

func (mc _ManagedObjectClass) New() ManagedObject {
	rv := objc.Call[ManagedObject](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewManagedObject() ManagedObject {
	return ManagedObjectClass.New()
}

func (m_ ManagedObject) Init() ManagedObject {
	rv := objc.Call[ManagedObject](m_, objc.Sel("init"))
	return rv
}

// Returns the value for the specified property from the managed object’s private internal storage . [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506728-primitivevalueforkey?language=objc
func (m_ ManagedObject) PrimitiveValueForKey(key string) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("primitiveValueForKey:"), key)
	return rv
}

// Provides an opportunity to respond when a value of a given property has changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506976-didchangevalueforkey?language=objc
func (m_ ManagedObject) DidChangeValueForKey(key string) {
	objc.Call[objc.Void](m_, objc.Sel("didChangeValueForKey:"), key)
}

// Validates a property value for a given key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506776-validatevalue?language=objc
func (m_ ManagedObject) ValidateValueForKeyError(value objc.IObject, key string, error foundation.IError) bool {
	rv := objc.Call[bool](m_, objc.Sel("validateValue:forKey:error:"), value, key, objc.Ptr(error))
	return rv
}

// Provides support for key-value observing access notification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506865-didaccessvalueforkey?language=objc
func (m_ ManagedObject) DidAccessValueForKey(key string) {
	objc.Call[objc.Void](m_, objc.Sel("didAccessValueForKey:"), key)
}

// Returns the observation info of the managed object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506210-observationinfo?language=objc
func (m_ ManagedObject) ObservationInfo() unsafe.Pointer {
	rv := objc.Call[unsafe.Pointer](m_, objc.Sel("observationInfo"))
	return rv
}

// Returns an initialized fetch request with the entity this subclass represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1640605-fetchrequest?language=objc
func (mc _ManagedObjectClass) FetchRequest() FetchRequest {
	rv := objc.Call[FetchRequest](mc, objc.Sel("fetchRequest"))
	return rv
}

// Returns an initialized fetch request with the entity this subclass represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1640605-fetchrequest?language=objc
func ManagedObject_FetchRequest() FetchRequest {
	return ManagedObjectClass.FetchRequest()
}

// Returns a dictionary containing the keys and new values of persistent properties with changes since the last fetching or saving of the managed object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506775-changedvalues?language=objc
func (m_ ManagedObject) ChangedValues() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](m_, objc.Sel("changedValues"))
	return rv
}

// Provides an opportunity to add code into the life cycle of the managed object when initially creating it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506548-awakefrominsert?language=objc
func (m_ ManagedObject) AwakeFromInsert() {
	objc.Call[objc.Void](m_, objc.Sel("awakeFromInsert"))
}

// Returns the object IDs for all of the managed objects that are in the named relationship. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506201-objectidsforrelationshipnamed?language=objc
func (m_ ManagedObject) ObjectIDsForRelationshipNamed(key string) []ManagedObjectID {
	rv := objc.Call[[]ManagedObjectID](m_, objc.Sel("objectIDsForRelationshipNamed:"), key)
	return rv
}

// Provides an opportunity to add code into the life cycle of the managed object before saving it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506209-willsave?language=objc
func (m_ ManagedObject) WillSave() {
	objc.Call[objc.Void](m_, objc.Sel("willSave"))
}

// Provides an opportunity to add code into the life cycle of the managed object after the managed object’s context completes a save operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506585-didsave?language=objc
func (m_ ManagedObject) DidSave() {
	objc.Call[objc.Void](m_, objc.Sel("didSave"))
}

// Determines whether the managed object can be deleted in its current state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506195-validatefordelete?language=objc
func (m_ ManagedObject) ValidateForDelete(error foundation.IError) bool {
	rv := objc.Call[bool](m_, objc.Sel("validateForDelete:"), objc.Ptr(error))
	return rv
}

// Sets the value of a given property in the managed object's private internal storage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506960-setprimitivevalue?language=objc
func (m_ ManagedObject) SetPrimitiveValueForKey(value objc.IObject, key string) {
	objc.Call[objc.Void](m_, objc.Sel("setPrimitiveValue:forKey:"), value, key)
}

// Provides an opportunity to add code into the life cycle of the managed object before converting it to a fault. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506537-willturnintofault?language=objc
func (m_ ManagedObject) WillTurnIntoFault() {
	objc.Call[objc.Void](m_, objc.Sel("willTurnIntoFault"))
}

// Provides an opportunity to add code into the life cycle of the managed object when fufilling it from a fault. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506424-awakefromfetch?language=objc
func (m_ ManagedObject) AwakeFromFetch() {
	objc.Call[objc.Void](m_, objc.Sel("awakeFromFetch"))
}

// Determines whether the managed object can be inserted in its current state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506683-validateforinsert?language=objc
func (m_ ManagedObject) ValidateForInsert(error foundation.IError) bool {
	rv := objc.Call[bool](m_, objc.Sel("validateForInsert:"), objc.Ptr(error))
	return rv
}

// Provides an opportunity to add code into the life cycle of the managed object before deleting it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506674-preparefordeletion?language=objc
func (m_ ManagedObject) PrepareForDeletion() {
	objc.Call[objc.Void](m_, objc.Sel("prepareForDeletion"))
}

// Determines whether the managed object's current state is valid. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506998-validateforupdate?language=objc
func (m_ ManagedObject) ValidateForUpdate(error foundation.IError) bool {
	rv := objc.Call[bool](m_, objc.Sel("validateForUpdate:"), objc.Ptr(error))
	return rv
}

// Provides an opportunity to add code into the life cycle of the managed object when fulfilling it from a snapshot. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506861-awakefromsnapshotevents?language=objc
func (m_ ManagedObject) AwakeFromSnapshotEvents(flags SnapshotEventType) {
	objc.Call[objc.Void](m_, objc.Sel("awakeFromSnapshotEvents:"), flags)
}

// Returns a dictionary of the most recent fetched or saved values of the managed object for the properties of the specified keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506771-committedvaluesforkeys?language=objc
func (m_ ManagedObject) CommittedValuesForKeys(keys []string) map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](m_, objc.Sel("committedValuesForKeys:"), keys)
	return rv
}

// Returns a Boolean value that indicates whether the relationship for a given key is a fault. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506973-hasfaultforrelationshipnamed?language=objc
func (m_ ManagedObject) HasFaultForRelationshipNamed(key string) bool {
	rv := objc.Call[bool](m_, objc.Sel("hasFaultForRelationshipNamed:"), key)
	return rv
}

// Provides support for key-value observing access notification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1507001-willaccessvalueforkey?language=objc
func (m_ ManagedObject) WillAccessValueForKey(key string) {
	objc.Call[objc.Void](m_, objc.Sel("willAccessValueForKey:"), key)
}

// Provides an opportunity to add code into the life cycle of the managed object after converting it to a fault. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506470-didturnintofault?language=objc
func (m_ ManagedObject) DidTurnIntoFault() {
	objc.Call[objc.Void](m_, objc.Sel("didTurnIntoFault"))
}

// Sets the specified property of the managed object to the specified value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506397-setvalue?language=objc
func (m_ ManagedObject) SetValueForKey(value objc.IObject, key string) {
	objc.Call[objc.Void](m_, objc.Sel("setValue:forKey:"), value, key)
}

// Returns a dictionary containing the keys and new values of persistent properties with changes since the last fetching or saving of the managed object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506472-changedvaluesforcurrentevent?language=objc
func (m_ ManagedObject) ChangedValuesForCurrentEvent() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](m_, objc.Sel("changedValuesForCurrentEvent"))
	return rv
}

// Provides an opportunity to respond when a value of a given property is about to change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506229-willchangevalueforkey?language=objc
func (m_ ManagedObject) WillChangeValueForKey(key string) {
	objc.Call[objc.Void](m_, objc.Sel("willChangeValueForKey:"), key)
}

// Sets the observation info of the managed object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506535-setobservationinfo?language=objc
func (m_ ManagedObject) SetObservationInfo(inObservationInfo unsafe.Pointer) {
	objc.Call[objc.Void](m_, objc.Sel("setObservationInfo:"), inObservationInfo)
}

// Returns the value for the property specified by key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506613-valueforkey?language=objc
func (m_ ManagedObject) ValueForKey(key string) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("valueForKey:"), key)
	return rv
}

// Initializes a managed object from an entity description and inserts it into the specified managed object context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506357-initwithentity?language=objc
func (m_ ManagedObject) InitWithEntityInsertIntoManagedObjectContext(entity IEntityDescription, context IManagedObjectContext) ManagedObject {
	rv := objc.Call[ManagedObject](m_, objc.Sel("initWithEntity:insertIntoManagedObjectContext:"), objc.Ptr(entity), objc.Ptr(context))
	return rv
}

// The entity description of the managed object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506562-entity?language=objc
func (m_ ManagedObject) Entity() EntityDescription {
	rv := objc.Call[EntityDescription](m_, objc.Sel("entity"))
	return rv
}

// A Boolean value that indicates whether the managed object has been inserted in a managed object context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506281-inserted?language=objc
func (m_ ManagedObject) IsInserted() bool {
	rv := objc.Call[bool](m_, objc.Sel("isInserted"))
	return rv
}

// A Boolean value that indicates whether the managed object has persistent changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506240-haspersistentchangedvalues?language=objc
func (m_ ManagedObject) HasPersistentChangedValues() bool {
	rv := objc.Call[bool](m_, objc.Sel("hasPersistentChangedValues"))
	return rv
}

// The faulting state of the managed object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506720-faultingstate?language=objc
func (m_ ManagedObject) FaultingState() uint {
	rv := objc.Call[uint](m_, objc.Sel("faultingState"))
	return rv
}

// A Boolean value that indicates whether to mark instances of the class as having changes when an unmodeled property changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506727-contextshouldignoreunmodeledprop?language=objc
func (mc _ManagedObjectClass) ContextShouldIgnoreUnmodeledPropertyChanges() bool {
	rv := objc.Call[bool](mc, objc.Sel("contextShouldIgnoreUnmodeledPropertyChanges"))
	return rv
}

// A Boolean value that indicates whether to mark instances of the class as having changes when an unmodeled property changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506727-contextshouldignoreunmodeledprop?language=objc
func ManagedObject_ContextShouldIgnoreUnmodeledPropertyChanges() bool {
	return ManagedObjectClass.ContextShouldIgnoreUnmodeledPropertyChanges()
}

// A Boolean value that indicates whether the managed object has unsaved changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506867-updated?language=objc
func (m_ ManagedObject) IsUpdated() bool {
	rv := objc.Call[bool](m_, objc.Sel("isUpdated"))
	return rv
}

// A Boolean value that indicates whether the managed object will be deleted during the next save. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506681-deleted?language=objc
func (m_ ManagedObject) IsDeleted() bool {
	rv := objc.Call[bool](m_, objc.Sel("isDeleted"))
	return rv
}

// The object ID of the managed object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506848-objectid?language=objc
func (m_ ManagedObject) ObjectID() ManagedObjectID {
	rv := objc.Call[ManagedObjectID](m_, objc.Sel("objectID"))
	return rv
}

// The managed object context with which the managed object is registered. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506677-managedobjectcontext?language=objc
func (m_ ManagedObject) ManagedObjectContext() ManagedObjectContext {
	rv := objc.Call[ManagedObjectContext](m_, objc.Sel("managedObjectContext"))
	return rv
}

// A Boolean value that indicates whether the managed object has been inserted, has been deleted, or has unsaved changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506654-haschanges?language=objc
func (m_ ManagedObject) HasChanges() bool {
	rv := objc.Call[bool](m_, objc.Sel("hasChanges"))
	return rv
}

// A Boolean value that indicates whether the managed object is a fault. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506837-fault?language=objc
func (m_ ManagedObject) IsFault() bool {
	rv := objc.Call[bool](m_, objc.Sel("isFault"))
	return rv
}
