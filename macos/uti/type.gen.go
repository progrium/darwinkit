// AUTO-GENERATED CODE, DO NOT MODIFY

package uti

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Type] class.
var TypeClass = _TypeClass{objc.GetClass("UTType")}

type _TypeClass struct {
	objc.Class
}

// An interface definition for the [Type] class.
type IType interface {
	objc.IObject
	IsSupertypeOfType(type_ IType) bool
	ConformsToType(type_ IType) bool
	IsSubtypeOfType(type_ IType) bool
	LocalizedDescription() string
	Version() foundation.Number
	PreferredFilenameExtension() string
	Supertypes() foundation.Set
	IsDynamic() bool
	IsPublicType() bool
	ReferenceURL() foundation.URL
	PreferredMIMEType() string
	IsDeclared() bool
	Tags() map[string][]string
	Identifier() string
}

// An object that represents a type of data to load, send, or receive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype?language=objc
type Type struct {
	objc.Object
}

func TypeFrom(ptr unsafe.Pointer) Type {
	return Type{
		Object: objc.ObjectFrom(ptr),
	}
}

func (tc _TypeClass) TypeWithMIMEType(mimeType string) Type {
	rv := objc.Call[Type](tc, objc.Sel("typeWithMIMEType:"), mimeType)
	return rv
}

// Creates a type based on a MIME type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype/3548219-typewithmimetype?language=objc
func Type_TypeWithMIMEType(mimeType string) Type {
	return TypeClass.TypeWithMIMEType(mimeType)
}

func (tc _TypeClass) TypeWithFilenameExtension(filenameExtension string) Type {
	rv := objc.Call[Type](tc, objc.Sel("typeWithFilenameExtension:"), filenameExtension)
	return rv
}

// Creates a type that represents the specified filename extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype/3548216-typewithfilenameextension?language=objc
func Type_TypeWithFilenameExtension(filenameExtension string) Type {
	return TypeClass.TypeWithFilenameExtension(filenameExtension)
}

func (tc _TypeClass) TypeWithTagTagClassConformingToType(tag string, tagClass string, supertype IType) Type {
	rv := objc.Call[Type](tc, objc.Sel("typeWithTag:tagClass:conformingToType:"), tag, tagClass, objc.Ptr(supertype))
	return rv
}

// Creates a type that represents the specified tag and tag class and which conforms to an existing type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype/3548221-typewithtag?language=objc
func Type_TypeWithTagTagClassConformingToType(tag string, tagClass string, supertype IType) Type {
	return TypeClass.TypeWithTagTagClassConformingToType(tag, tagClass, supertype)
}

func (tc _TypeClass) TypeWithIdentifier(identifier string) Type {
	rv := objc.Call[Type](tc, objc.Sel("typeWithIdentifier:"), identifier)
	return rv
}

// Creates a type based on an identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype/3548218-typewithidentifier?language=objc
func Type_TypeWithIdentifier(identifier string) Type {
	return TypeClass.TypeWithIdentifier(identifier)
}

func (tc _TypeClass) Alloc() Type {
	rv := objc.Call[Type](tc, objc.Sel("alloc"))
	return rv
}

func Type_Alloc() Type {
	return TypeClass.Alloc()
}

func (tc _TypeClass) New() Type {
	rv := objc.Call[Type](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewType() Type {
	return TypeClass.New()
}

func (t_ Type) Init() Type {
	rv := objc.Call[Type](t_, objc.Sel("init"))
	return rv
}

// Returns a Boolean value that indicates whether a type is lower in a hierarchy than the type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype/3548208-issupertypeoftype?language=objc
func (t_ Type) IsSupertypeOfType(type_ IType) bool {
	rv := objc.Call[bool](t_, objc.Sel("isSupertypeOfType:"), objc.Ptr(type_))
	return rv
}

// Creates a type your app uses, but doesn’t own, based on an identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype/3600610-importedtypewithidentifier?language=objc
func (tc _TypeClass) ImportedTypeWithIdentifier(identifier string) Type {
	rv := objc.Call[Type](tc, objc.Sel("importedTypeWithIdentifier:"), identifier)
	return rv
}

// Creates a type your app uses, but doesn’t own, based on an identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype/3600610-importedtypewithidentifier?language=objc
func Type_ImportedTypeWithIdentifier(identifier string) Type {
	return TypeClass.ImportedTypeWithIdentifier(identifier)
}

// Creates a type your app owns based on an identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype/3600608-exportedtypewithidentifier?language=objc
func (tc _TypeClass) ExportedTypeWithIdentifier(identifier string) Type {
	rv := objc.Call[Type](tc, objc.Sel("exportedTypeWithIdentifier:"), identifier)
	return rv
}

// Creates a type your app owns based on an identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype/3600608-exportedtypewithidentifier?language=objc
func Type_ExportedTypeWithIdentifier(identifier string) Type {
	return TypeClass.ExportedTypeWithIdentifier(identifier)
}

// Returns a Boolean value that indicates whether a type conforms to the type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype/3548203-conformstotype?language=objc
func (t_ Type) ConformsToType(type_ IType) bool {
	rv := objc.Call[bool](t_, objc.Sel("conformsToType:"), objc.Ptr(type_))
	return rv
}

// Returns a Boolean value that indicates whether a type is higher in a hierarchy than the type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype/3548207-issubtypeoftype?language=objc
func (t_ Type) IsSubtypeOfType(type_ IType) bool {
	rv := objc.Call[bool](t_, objc.Sel("isSubtypeOfType:"), objc.Ptr(type_))
	return rv
}

// Returns an array of types from the provided tag and tag class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype/3548222-typeswithtag?language=objc
func (tc _TypeClass) TypesWithTagTagClassConformingToType(tag string, tagClass string, supertype IType) []Type {
	rv := objc.Call[[]Type](tc, objc.Sel("typesWithTag:tagClass:conformingToType:"), tag, tagClass, objc.Ptr(supertype))
	return rv
}

// Returns an array of types from the provided tag and tag class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype/3548222-typeswithtag?language=objc
func Type_TypesWithTagTagClassConformingToType(tag string, tagClass string, supertype IType) []Type {
	return TypeClass.TypesWithTagTagClassConformingToType(tag, tagClass, supertype)
}

// A localized description of the type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype/3548209-localizeddescription?language=objc
func (t_ Type) LocalizedDescription() string {
	rv := objc.Call[string](t_, objc.Sel("localizedDescription"))
	return rv
}

// The type’s version, if available. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype/3548223-version?language=objc
func (t_ Type) Version() foundation.Number {
	rv := objc.Call[foundation.Number](t_, objc.Sel("version"))
	return rv
}

// The preferred filename extension for the type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype/3548210-preferredfilenameextension?language=objc
func (t_ Type) PreferredFilenameExtension() string {
	rv := objc.Call[string](t_, objc.Sel("preferredFilenameExtension"))
	return rv
}

// The set of types the type directly or indirectly conforms to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype/3548214-supertypes?language=objc
func (t_ Type) Supertypes() foundation.Set {
	rv := objc.Call[foundation.Set](t_, objc.Sel("supertypes"))
	return rv
}

// A Boolean value that indicates whether the system generates the type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype/3548205-dynamic?language=objc
func (t_ Type) IsDynamic() bool {
	rv := objc.Call[bool](t_, objc.Sel("isDynamic"))
	return rv
}

// A Boolean value that indicates whether the type is in the public domain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype/3548212-publictype?language=objc
func (t_ Type) IsPublicType() bool {
	rv := objc.Call[bool](t_, objc.Sel("isPublicType"))
	return rv
}

// The reference URL for the type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype/3548213-referenceurl?language=objc
func (t_ Type) ReferenceURL() foundation.URL {
	rv := objc.Call[foundation.URL](t_, objc.Sel("referenceURL"))
	return rv
}

// The preferred MIME type for the type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype/3548211-preferredmimetype?language=objc
func (t_ Type) PreferredMIMEType() string {
	rv := objc.Call[string](t_, objc.Sel("preferredMIMEType"))
	return rv
}

// A Boolean value that indicates whether the system declares the type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype/3548204-declared?language=objc
func (t_ Type) IsDeclared() bool {
	rv := objc.Call[bool](t_, objc.Sel("isDeclared"))
	return rv
}

// The tag specification dictionary of the type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype/3548215-tags?language=objc
func (t_ Type) Tags() map[string][]string {
	rv := objc.Call[map[string][]string](t_, objc.Sel("tags"))
	return rv
}

// The string that represents the type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype/3548206-identifier?language=objc
func (t_ Type) Identifier() string {
	rv := objc.Call[string](t_, objc.Sel("identifier"))
	return rv
}
