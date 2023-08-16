// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AttributeDescription] class.
var AttributeDescriptionClass = _AttributeDescriptionClass{objc.GetClass("NSAttributeDescription")}

type _AttributeDescriptionClass struct {
	objc.Class
}

// An interface definition for the [AttributeDescription] class.
type IAttributeDescription interface {
	IPropertyDescription
	AllowsExternalBinaryDataStorage() bool
	SetAllowsExternalBinaryDataStorage(value bool)
	DefaultValue() objc.Object
	SetDefaultValue(value objc.IObject)
	ValueTransformerName() string
	SetValueTransformerName(value string)
	AttributeType() AttributeType
	SetAttributeType(value AttributeType)
	AllowsCloudEncryption() bool
	SetAllowsCloudEncryption(value bool)
	PreservesValueInHistoryOnDeletion() bool
	SetPreservesValueInHistoryOnDeletion(value bool)
	AttributeValueClassName() string
	SetAttributeValueClassName(value string)
}

// A description of a single attribute belonging to an entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription?language=objc
type AttributeDescription struct {
	PropertyDescription
}

func AttributeDescriptionFrom(ptr unsafe.Pointer) AttributeDescription {
	return AttributeDescription{
		PropertyDescription: PropertyDescriptionFrom(ptr),
	}
}

func (ac _AttributeDescriptionClass) Alloc() AttributeDescription {
	rv := objc.Call[AttributeDescription](ac, objc.Sel("alloc"))
	return rv
}

func AttributeDescription_Alloc() AttributeDescription {
	return AttributeDescriptionClass.Alloc()
}

func (ac _AttributeDescriptionClass) New() AttributeDescription {
	rv := objc.Call[AttributeDescription](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAttributeDescription() AttributeDescription {
	return AttributeDescriptionClass.New()
}

func (a_ AttributeDescription) Init() AttributeDescription {
	rv := objc.Call[AttributeDescription](a_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether the attribute allows external binary storage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/1498295-allowsexternalbinarydatastorage?language=objc
func (a_ AttributeDescription) AllowsExternalBinaryDataStorage() bool {
	rv := objc.Call[bool](a_, objc.Sel("allowsExternalBinaryDataStorage"))
	return rv
}

// A Boolean value that indicates whether the attribute allows external binary storage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/1498295-allowsexternalbinarydatastorage?language=objc
func (a_ AttributeDescription) SetAllowsExternalBinaryDataStorage(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setAllowsExternalBinaryDataStorage:"), value)
}

// The default value of the attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/1498302-defaultvalue?language=objc
func (a_ AttributeDescription) DefaultValue() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("defaultValue"))
	return rv
}

// The default value of the attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/1498302-defaultvalue?language=objc
func (a_ AttributeDescription) SetDefaultValue(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setDefaultValue:"), value)
}

// The name of the transformer to use for the attribute value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/1498305-valuetransformername?language=objc
func (a_ AttributeDescription) ValueTransformerName() string {
	rv := objc.Call[string](a_, objc.Sel("valueTransformerName"))
	return rv
}

// The name of the transformer to use for the attribute value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/1498305-valuetransformername?language=objc
func (a_ AttributeDescription) SetValueTransformerName(value string) {
	objc.Call[objc.Void](a_, objc.Sel("setValueTransformerName:"), value)
}

// The attribute’s type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/1498291-attributetype?language=objc
func (a_ AttributeDescription) AttributeType() AttributeType {
	rv := objc.Call[AttributeType](a_, objc.Sel("attributeType"))
	return rv
}

// The attribute’s type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/1498291-attributetype?language=objc
func (a_ AttributeDescription) SetAttributeType(value AttributeType) {
	objc.Call[objc.Void](a_, objc.Sel("setAttributeType:"), value)
}

// A Boolean value that determines whether to encrypt the attribute’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/3746827-allowscloudencryption?language=objc
func (a_ AttributeDescription) AllowsCloudEncryption() bool {
	rv := objc.Call[bool](a_, objc.Sel("allowsCloudEncryption"))
	return rv
}

// A Boolean value that determines whether to encrypt the attribute’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/3746827-allowscloudencryption?language=objc
func (a_ AttributeDescription) SetAllowsCloudEncryption(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setAllowsCloudEncryption:"), value)
}

// A Boolean value that indicates whether the attribute records its value in the persistent history transaction for a managed object’s deletion. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/3180042-preservesvalueinhistoryondeletio?language=objc
func (a_ AttributeDescription) PreservesValueInHistoryOnDeletion() bool {
	rv := objc.Call[bool](a_, objc.Sel("preservesValueInHistoryOnDeletion"))
	return rv
}

// A Boolean value that indicates whether the attribute records its value in the persistent history transaction for a managed object’s deletion. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/3180042-preservesvalueinhistoryondeletio?language=objc
func (a_ AttributeDescription) SetPreservesValueInHistoryOnDeletion(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setPreservesValueInHistoryOnDeletion:"), value)
}

// The class name that represents the attribute’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/1498309-attributevalueclassname?language=objc
func (a_ AttributeDescription) AttributeValueClassName() string {
	rv := objc.Call[string](a_, objc.Sel("attributeValueClassName"))
	return rv
}

// The class name that represents the attribute’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsattributedescription/1498309-attributevalueclassname?language=objc
func (a_ AttributeDescription) SetAttributeValueClassName(value string) {
	objc.Call[objc.Void](a_, objc.Sel("setAttributeValueClassName:"), value)
}
