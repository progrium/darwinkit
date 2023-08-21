// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FetchIndexElementDescription] class.
var FetchIndexElementDescriptionClass = _FetchIndexElementDescriptionClass{objc.GetClass("NSFetchIndexElementDescription")}

type _FetchIndexElementDescriptionClass struct {
	objc.Class
}

// An interface definition for the [FetchIndexElementDescription] class.
type IFetchIndexElementDescription interface {
	objc.IObject
	CollationType() FetchIndexElementType
	SetCollationType(value FetchIndexElementType)
	IndexDescription() FetchIndexDescription
	PropertyName() string
	Property() PropertyDescription
	IsAscending() bool
	SetAscending(value bool)
}

// Description of an Index Element [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchindexelementdescription?language=objc
type FetchIndexElementDescription struct {
	objc.Object
}

func FetchIndexElementDescriptionFrom(ptr unsafe.Pointer) FetchIndexElementDescription {
	return FetchIndexElementDescription{
		Object: objc.ObjectFrom(ptr),
	}
}

func (f_ FetchIndexElementDescription) InitWithPropertyCollationType(property IPropertyDescription, collationType FetchIndexElementType) FetchIndexElementDescription {
	rv := objc.Call[FetchIndexElementDescription](f_, objc.Sel("initWithProperty:collationType:"), objc.Ptr(property), collationType)
	return rv
}

// Creates an index element description using the specified property description and collation type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchindexelementdescription/2887043-initwithproperty?language=objc
func NewFetchIndexElementDescriptionWithPropertyCollationType(property IPropertyDescription, collationType FetchIndexElementType) FetchIndexElementDescription {
	instance := FetchIndexElementDescriptionClass.Alloc().InitWithPropertyCollationType(property, collationType)
	instance.Autorelease()
	return instance
}

func (fc _FetchIndexElementDescriptionClass) Alloc() FetchIndexElementDescription {
	rv := objc.Call[FetchIndexElementDescription](fc, objc.Sel("alloc"))
	return rv
}

func FetchIndexElementDescription_Alloc() FetchIndexElementDescription {
	return FetchIndexElementDescriptionClass.Alloc()
}

func (fc _FetchIndexElementDescriptionClass) New() FetchIndexElementDescription {
	rv := objc.Call[FetchIndexElementDescription](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFetchIndexElementDescription() FetchIndexElementDescription {
	return FetchIndexElementDescriptionClass.New()
}

func (f_ FetchIndexElementDescription) Init() FetchIndexElementDescription {
	rv := objc.Call[FetchIndexElementDescription](f_, objc.Sel("init"))
	return rv
}

// The type of collation that the index element uses, either binary or R-tree. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchindexelementdescription/2887048-collationtype?language=objc
func (f_ FetchIndexElementDescription) CollationType() FetchIndexElementType {
	rv := objc.Call[FetchIndexElementType](f_, objc.Sel("collationType"))
	return rv
}

// The type of collation that the index element uses, either binary or R-tree. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchindexelementdescription/2887048-collationtype?language=objc
func (f_ FetchIndexElementDescription) SetCollationType(value FetchIndexElementType) {
	objc.Call[objc.Void](f_, objc.Sel("setCollationType:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchindexelementdescription/2887047-indexdescription?language=objc
func (f_ FetchIndexElementDescription) IndexDescription() FetchIndexDescription {
	rv := objc.Call[FetchIndexDescription](f_, objc.Sel("indexDescription"))
	return rv
}

// The specified name in the property description. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchindexelementdescription/2887045-propertyname?language=objc
func (f_ FetchIndexElementDescription) PropertyName() string {
	rv := objc.Call[string](f_, objc.Sel("propertyName"))
	return rv
}

// A property description. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchindexelementdescription/2887049-property?language=objc
func (f_ FetchIndexElementDescription) Property() PropertyDescription {
	rv := objc.Call[PropertyDescription](f_, objc.Sel("property"))
	return rv
}

// A Boolean value that controls whether an index that supports direction is an ascending or descending index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchindexelementdescription/2887051-ascending?language=objc
func (f_ FetchIndexElementDescription) IsAscending() bool {
	rv := objc.Call[bool](f_, objc.Sel("isAscending"))
	return rv
}

// A Boolean value that controls whether an index that supports direction is an ascending or descending index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchindexelementdescription/2887051-ascending?language=objc
func (f_ FetchIndexElementDescription) SetAscending(value bool) {
	objc.Call[objc.Void](f_, objc.Sel("setAscending:"), value)
}
