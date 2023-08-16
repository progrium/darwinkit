// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [EntityDescription] class.
var EntityDescriptionClass = _EntityDescriptionClass{objc.GetClass("NSEntityDescription")}

type _EntityDescriptionClass struct {
	objc.Class
}

// An interface definition for the [EntityDescription] class.
type IEntityDescription interface {
	objc.IObject
	RelationshipsWithDestinationEntity(entity IEntityDescription) []RelationshipDescription
	IsKindOfEntity(entity IEntityDescription) bool
	ManagedObjectModel() ManagedObjectModel
	PropertiesByName() map[string]PropertyDescription
	Name() string
	SetName(value string)
	VersionHash() []byte
	Properties() []PropertyDescription
	SetProperties(value []IPropertyDescription)
	UniquenessConstraints() [][]objc.Object
	SetUniquenessConstraints(value [][]objc.IObject)
	UserInfo() foundation.Dictionary
	SetUserInfo(value foundation.Dictionary)
	ManagedObjectClassName() string
	SetManagedObjectClassName(value string)
	Subentities() []EntityDescription
	SetSubentities(value []IEntityDescription)
	CoreSpotlightDisplayNameExpression() foundation.Expression
	SetCoreSpotlightDisplayNameExpression(value foundation.IExpression)
	VersionHashModifier() string
	SetVersionHashModifier(value string)
	RenamingIdentifier() string
	SetRenamingIdentifier(value string)
	RelationshipsByName() map[string]RelationshipDescription
	IsAbstract() bool
	SetAbstract(value bool)
	AttributesByName() map[string]AttributeDescription
	Indexes() []FetchIndexDescription
	SetIndexes(value []IFetchIndexDescription)
	Superentity() EntityDescription
	SubentitiesByName() map[string]EntityDescription
}

// A description of a Core Data entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription?language=objc
type EntityDescription struct {
	objc.Object
}

func EntityDescriptionFrom(ptr unsafe.Pointer) EntityDescription {
	return EntityDescription{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ec _EntityDescriptionClass) Alloc() EntityDescription {
	rv := objc.Call[EntityDescription](ec, objc.Sel("alloc"))
	return rv
}

func EntityDescription_Alloc() EntityDescription {
	return EntityDescriptionClass.Alloc()
}

func (ec _EntityDescriptionClass) New() EntityDescription {
	rv := objc.Call[EntityDescription](ec, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewEntityDescription() EntityDescription {
	return EntityDescriptionClass.New()
}

func (e_ EntityDescription) Init() EntityDescription {
	rv := objc.Call[EntityDescription](e_, objc.Sel("init"))
	return rv
}

// Returns the entity with the specified name from the managed object model associated with the specified managed object context’s persistent store coordinator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/1425111-entityforname?language=objc
func (ec _EntityDescriptionClass) EntityForNameInManagedObjectContext(entityName string, context IManagedObjectContext) EntityDescription {
	rv := objc.Call[EntityDescription](ec, objc.Sel("entityForName:inManagedObjectContext:"), entityName, objc.Ptr(context))
	return rv
}

// Returns the entity with the specified name from the managed object model associated with the specified managed object context’s persistent store coordinator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/1425111-entityforname?language=objc
func EntityDescription_EntityForNameInManagedObjectContext(entityName string, context IManagedObjectContext) EntityDescription {
	return EntityDescriptionClass.EntityForNameInManagedObjectContext(entityName, context)
}

// Returns an array containing the relationships of the receiver where the entity description of the relationship is a given entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/1425127-relationshipswithdestinationenti?language=objc
func (e_ EntityDescription) RelationshipsWithDestinationEntity(entity IEntityDescription) []RelationshipDescription {
	rv := objc.Call[[]RelationshipDescription](e_, objc.Sel("relationshipsWithDestinationEntity:"), objc.Ptr(entity))
	return rv
}

// Creates, configures, and returns an instance of the class for the entity with a given name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/1425093-insertnewobjectforentityforname?language=objc
func (ec _EntityDescriptionClass) InsertNewObjectForEntityForNameInManagedObjectContext(entityName string, context IManagedObjectContext) ManagedObject {
	rv := objc.Call[ManagedObject](ec, objc.Sel("insertNewObjectForEntityForName:inManagedObjectContext:"), entityName, objc.Ptr(context))
	return rv
}

// Creates, configures, and returns an instance of the class for the entity with a given name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/1425093-insertnewobjectforentityforname?language=objc
func EntityDescription_InsertNewObjectForEntityForNameInManagedObjectContext(entityName string, context IManagedObjectContext) ManagedObject {
	return EntityDescriptionClass.InsertNewObjectForEntityForNameInManagedObjectContext(entityName, context)
}

// Returns a Boolean value that indicates whether the receiver is a sub-entity of another given entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/1425113-iskindofentity?language=objc
func (e_ EntityDescription) IsKindOfEntity(entity IEntityDescription) bool {
	rv := objc.Call[bool](e_, objc.Sel("isKindOfEntity:"), objc.Ptr(entity))
	return rv
}

// The managed object model with which the receiver is associated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/1425121-managedobjectmodel?language=objc
func (e_ EntityDescription) ManagedObjectModel() ManagedObjectModel {
	rv := objc.Call[ManagedObjectModel](e_, objc.Sel("managedObjectModel"))
	return rv
}

// A dictionary containing the properties of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/1425137-propertiesbyname?language=objc
func (e_ EntityDescription) PropertiesByName() map[string]PropertyDescription {
	rv := objc.Call[map[string]PropertyDescription](e_, objc.Sel("propertiesByName"))
	return rv
}

// The entity name of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/1425101-name?language=objc
func (e_ EntityDescription) Name() string {
	rv := objc.Call[string](e_, objc.Sel("name"))
	return rv
}

// The entity name of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/1425101-name?language=objc
func (e_ EntityDescription) SetName(value string) {
	objc.Call[objc.Void](e_, objc.Sel("setName:"), value)
}

// The version hash for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/1425133-versionhash?language=objc
func (e_ EntityDescription) VersionHash() []byte {
	rv := objc.Call[[]byte](e_, objc.Sel("versionHash"))
	return rv
}

// An array containing the properties of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/1425125-properties?language=objc
func (e_ EntityDescription) Properties() []PropertyDescription {
	rv := objc.Call[[]PropertyDescription](e_, objc.Sel("properties"))
	return rv
}

// An array containing the properties of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/1425125-properties?language=objc
func (e_ EntityDescription) SetProperties(value []IPropertyDescription) {
	objc.Call[objc.Void](e_, objc.Sel("setProperties:"), value)
}

// An array of arrays that contains one or more attributes with a value that must be unique over the instances of that entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/1425095-uniquenessconstraints?language=objc
func (e_ EntityDescription) UniquenessConstraints() [][]objc.Object {
	rv := objc.Call[[][]objc.Object](e_, objc.Sel("uniquenessConstraints"))
	return rv
}

// An array of arrays that contains one or more attributes with a value that must be unique over the instances of that entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/1425095-uniquenessconstraints?language=objc
func (e_ EntityDescription) SetUniquenessConstraints(value [][]objc.IObject) {
	objc.Call[objc.Void](e_, objc.Sel("setUniquenessConstraints:"), value)
}

// The user info dictionary of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/1425117-userinfo?language=objc
func (e_ EntityDescription) UserInfo() foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](e_, objc.Sel("userInfo"))
	return rv
}

// The user info dictionary of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/1425117-userinfo?language=objc
func (e_ EntityDescription) SetUserInfo(value foundation.Dictionary) {
	objc.Call[objc.Void](e_, objc.Sel("setUserInfo:"), value)
}

// The name of the class that represents the receiver’s entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/1425131-managedobjectclassname?language=objc
func (e_ EntityDescription) ManagedObjectClassName() string {
	rv := objc.Call[string](e_, objc.Sel("managedObjectClassName"))
	return rv
}

// The name of the class that represents the receiver’s entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/1425131-managedobjectclassname?language=objc
func (e_ EntityDescription) SetManagedObjectClassName(value string) {
	objc.Call[objc.Void](e_, objc.Sel("setManagedObjectClassName:"), value)
}

// An array containing the sub-entities of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/1425104-subentities?language=objc
func (e_ EntityDescription) Subentities() []EntityDescription {
	rv := objc.Call[[]EntityDescription](e_, objc.Sel("subentities"))
	return rv
}

// An array containing the sub-entities of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/1425104-subentities?language=objc
func (e_ EntityDescription) SetSubentities(value []IEntityDescription) {
	objc.Call[objc.Void](e_, objc.Sel("setSubentities:"), value)
}

// The expression that computes the CoreSpotlight display name for instances of the entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/2892364-corespotlightdisplaynameexpressi?language=objc
func (e_ EntityDescription) CoreSpotlightDisplayNameExpression() foundation.Expression {
	rv := objc.Call[foundation.Expression](e_, objc.Sel("coreSpotlightDisplayNameExpression"))
	return rv
}

// The expression that computes the CoreSpotlight display name for instances of the entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/2892364-corespotlightdisplaynameexpressi?language=objc
func (e_ EntityDescription) SetCoreSpotlightDisplayNameExpression(value foundation.IExpression) {
	objc.Call[objc.Void](e_, objc.Sel("setCoreSpotlightDisplayNameExpression:"), objc.Ptr(value))
}

// The version hash modifier for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/1425119-versionhashmodifier?language=objc
func (e_ EntityDescription) VersionHashModifier() string {
	rv := objc.Call[string](e_, objc.Sel("versionHashModifier"))
	return rv
}

// The version hash modifier for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/1425119-versionhashmodifier?language=objc
func (e_ EntityDescription) SetVersionHashModifier(value string) {
	objc.Call[objc.Void](e_, objc.Sel("setVersionHashModifier:"), value)
}

// The renaming identifier for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/1425135-renamingidentifier?language=objc
func (e_ EntityDescription) RenamingIdentifier() string {
	rv := objc.Call[string](e_, objc.Sel("renamingIdentifier"))
	return rv
}

// The renaming identifier for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/1425135-renamingidentifier?language=objc
func (e_ EntityDescription) SetRenamingIdentifier(value string) {
	objc.Call[objc.Void](e_, objc.Sel("setRenamingIdentifier:"), value)
}

// The relationships of the receiver in a dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/1425106-relationshipsbyname?language=objc
func (e_ EntityDescription) RelationshipsByName() map[string]RelationshipDescription {
	rv := objc.Call[map[string]RelationshipDescription](e_, objc.Sel("relationshipsByName"))
	return rv
}

// A Boolean value that indicates whether the receiver represents an abstract entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/1425097-abstract?language=objc
func (e_ EntityDescription) IsAbstract() bool {
	rv := objc.Call[bool](e_, objc.Sel("isAbstract"))
	return rv
}

// A Boolean value that indicates whether the receiver represents an abstract entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/1425097-abstract?language=objc
func (e_ EntityDescription) SetAbstract(value bool) {
	objc.Call[objc.Void](e_, objc.Sel("setAbstract:"), value)
}

// The attributes of the receiver in a dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/1425099-attributesbyname?language=objc
func (e_ EntityDescription) AttributesByName() map[string]AttributeDescription {
	rv := objc.Call[map[string]AttributeDescription](e_, objc.Sel("attributesByName"))
	return rv
}

// An array of fetch index descriptions for the entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/2887059-indexes?language=objc
func (e_ EntityDescription) Indexes() []FetchIndexDescription {
	rv := objc.Call[[]FetchIndexDescription](e_, objc.Sel("indexes"))
	return rv
}

// An array of fetch index descriptions for the entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/2887059-indexes?language=objc
func (e_ EntityDescription) SetIndexes(value []IFetchIndexDescription) {
	objc.Call[objc.Void](e_, objc.Sel("setIndexes:"), value)
}

// The super-entity of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/1425129-superentity?language=objc
func (e_ EntityDescription) Superentity() EntityDescription {
	rv := objc.Call[EntityDescription](e_, objc.Sel("superentity"))
	return rv
}

// A dictionary containing the receiver’s sub-entities. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitydescription/1425123-subentitiesbyname?language=objc
func (e_ EntityDescription) SubentitiesByName() map[string]EntityDescription {
	rv := objc.Call[map[string]EntityDescription](e_, objc.Sel("subentitiesByName"))
	return rv
}
