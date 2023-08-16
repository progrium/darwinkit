// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PropertyDescription] class.
var PropertyDescriptionClass = _PropertyDescriptionClass{objc.GetClass("NSPropertyDescription")}

type _PropertyDescriptionClass struct {
	objc.Class
}

// An interface definition for the [PropertyDescription] class.
type IPropertyDescription interface {
	objc.IObject
	SetValidationPredicatesWithValidationWarnings(validationPredicates []foundation.IPredicate, validationWarnings []string)
	Entity() EntityDescription
	Name() string
	SetName(value string)
	VersionHash() []byte
	IsTransient() bool
	SetTransient(value bool)
	ValidationPredicates() []foundation.Predicate
	UserInfo() foundation.Dictionary
	SetUserInfo(value foundation.Dictionary)
	ValidationWarnings() []objc.Object
	VersionHashModifier() string
	SetVersionHashModifier(value string)
	RenamingIdentifier() string
	SetRenamingIdentifier(value string)
	IsOptional() bool
	SetOptional(value bool)
	IsIndexedBySpotlight() bool
	SetIndexedBySpotlight(value bool)
}

// A description of a single property belonging to an entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription?language=objc
type PropertyDescription struct {
	objc.Object
}

func PropertyDescriptionFrom(ptr unsafe.Pointer) PropertyDescription {
	return PropertyDescription{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PropertyDescriptionClass) Alloc() PropertyDescription {
	rv := objc.Call[PropertyDescription](pc, objc.Sel("alloc"))
	return rv
}

func PropertyDescription_Alloc() PropertyDescription {
	return PropertyDescriptionClass.Alloc()
}

func (pc _PropertyDescriptionClass) New() PropertyDescription {
	rv := objc.Call[PropertyDescription](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPropertyDescription() PropertyDescription {
	return PropertyDescriptionClass.New()
}

func (p_ PropertyDescription) Init() PropertyDescription {
	rv := objc.Call[PropertyDescription](p_, objc.Sel("init"))
	return rv
}

// Sets the validation predicates and warnings of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/1506852-setvalidationpredicates?language=objc
func (p_ PropertyDescription) SetValidationPredicatesWithValidationWarnings(validationPredicates []foundation.IPredicate, validationWarnings []string) {
	objc.Call[objc.Void](p_, objc.Sel("setValidationPredicates:withValidationWarnings:"), validationPredicates, validationWarnings)
}

// The entity description of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/1506745-entity?language=objc
func (p_ PropertyDescription) Entity() EntityDescription {
	rv := objc.Call[EntityDescription](p_, objc.Sel("entity"))
	return rv
}

// The name of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/1506759-name?language=objc
func (p_ PropertyDescription) Name() string {
	rv := objc.Call[string](p_, objc.Sel("name"))
	return rv
}

// The name of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/1506759-name?language=objc
func (p_ PropertyDescription) SetName(value string) {
	objc.Call[objc.Void](p_, objc.Sel("setName:"), value)
}

// The version hash for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/1506198-versionhash?language=objc
func (p_ PropertyDescription) VersionHash() []byte {
	rv := objc.Call[[]byte](p_, objc.Sel("versionHash"))
	return rv
}

// A Boolean value that indicates whether the receiver is transient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/1506766-transient?language=objc
func (p_ PropertyDescription) IsTransient() bool {
	rv := objc.Call[bool](p_, objc.Sel("isTransient"))
	return rv
}

// A Boolean value that indicates whether the receiver is transient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/1506766-transient?language=objc
func (p_ PropertyDescription) SetTransient(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setTransient:"), value)
}

// The validation predicates of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/1506842-validationpredicates?language=objc
func (p_ PropertyDescription) ValidationPredicates() []foundation.Predicate {
	rv := objc.Call[[]foundation.Predicate](p_, objc.Sel("validationPredicates"))
	return rv
}

// The user info dictionary of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/1506833-userinfo?language=objc
func (p_ PropertyDescription) UserInfo() foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](p_, objc.Sel("userInfo"))
	return rv
}

// The user info dictionary of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/1506833-userinfo?language=objc
func (p_ PropertyDescription) SetUserInfo(value foundation.Dictionary) {
	objc.Call[objc.Void](p_, objc.Sel("setUserInfo:"), value)
}

// The error strings associated with the receiver’s validation predicates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/1506886-validationwarnings?language=objc
func (p_ PropertyDescription) ValidationWarnings() []objc.Object {
	rv := objc.Call[[]objc.Object](p_, objc.Sel("validationWarnings"))
	return rv
}

// The version hash modifier for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/1506214-versionhashmodifier?language=objc
func (p_ PropertyDescription) VersionHashModifier() string {
	rv := objc.Call[string](p_, objc.Sel("versionHashModifier"))
	return rv
}

// The version hash modifier for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/1506214-versionhashmodifier?language=objc
func (p_ PropertyDescription) SetVersionHashModifier(value string) {
	objc.Call[objc.Void](p_, objc.Sel("setVersionHashModifier:"), value)
}

// The renaming identifier for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/1506641-renamingidentifier?language=objc
func (p_ PropertyDescription) RenamingIdentifier() string {
	rv := objc.Call[string](p_, objc.Sel("renamingIdentifier"))
	return rv
}

// The renaming identifier for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/1506641-renamingidentifier?language=objc
func (p_ PropertyDescription) SetRenamingIdentifier(value string) {
	objc.Call[objc.Void](p_, objc.Sel("setRenamingIdentifier:"), value)
}

// A Boolean value that indicates whether the receiver is optional. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/1506735-optional?language=objc
func (p_ PropertyDescription) IsOptional() bool {
	rv := objc.Call[bool](p_, objc.Sel("isOptional"))
	return rv
}

// A Boolean value that indicates whether the receiver is optional. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/1506735-optional?language=objc
func (p_ PropertyDescription) SetOptional(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setOptional:"), value)
}

// A Boolean value that indicates whether Core Data adds the property’s value to the Core Spotlight index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/1506784-indexedbyspotlight?language=objc
func (p_ PropertyDescription) IsIndexedBySpotlight() bool {
	rv := objc.Call[bool](p_, objc.Sel("isIndexedBySpotlight"))
	return rv
}

// A Boolean value that indicates whether Core Data adds the property’s value to the Core Spotlight index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertydescription/1506784-indexedbyspotlight?language=objc
func (p_ PropertyDescription) SetIndexedBySpotlight(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setIndexedBySpotlight:"), value)
}
