// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RelationshipDescription] class.
var RelationshipDescriptionClass = _RelationshipDescriptionClass{objc.GetClass("NSRelationshipDescription")}

type _RelationshipDescriptionClass struct {
	objc.Class
}

// An interface definition for the [RelationshipDescription] class.
type IRelationshipDescription interface {
	IPropertyDescription
	IsOrdered() bool
	SetOrdered(value bool)
	DestinationEntity() EntityDescription
	SetDestinationEntity(value IEntityDescription)
	DeleteRule() DeleteRule
	SetDeleteRule(value DeleteRule)
	MaxCount() uint
	SetMaxCount(value uint)
	IsToMany() bool
	MinCount() uint
	SetMinCount(value uint)
	InverseRelationship() RelationshipDescription
	SetInverseRelationship(value IRelationshipDescription)
}

// A description of a relationship between two entities. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsrelationshipdescription?language=objc
type RelationshipDescription struct {
	PropertyDescription
}

func RelationshipDescriptionFrom(ptr unsafe.Pointer) RelationshipDescription {
	return RelationshipDescription{
		PropertyDescription: PropertyDescriptionFrom(ptr),
	}
}

func (rc _RelationshipDescriptionClass) Alloc() RelationshipDescription {
	rv := objc.Call[RelationshipDescription](rc, objc.Sel("alloc"))
	return rv
}

func RelationshipDescription_Alloc() RelationshipDescription {
	return RelationshipDescriptionClass.Alloc()
}

func (rc _RelationshipDescriptionClass) New() RelationshipDescription {
	rv := objc.Call[RelationshipDescription](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRelationshipDescription() RelationshipDescription {
	return RelationshipDescriptionClass.New()
}

func (r_ RelationshipDescription) Init() RelationshipDescription {
	rv := objc.Call[RelationshipDescription](r_, objc.Sel("init"))
	return rv
}

// A Boolean value that determines whether the relationship preserves the order of the referenced managed objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsrelationshipdescription/1506382-ordered?language=objc
func (r_ RelationshipDescription) IsOrdered() bool {
	rv := objc.Call[bool](r_, objc.Sel("isOrdered"))
	return rv
}

// A Boolean value that determines whether the relationship preserves the order of the referenced managed objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsrelationshipdescription/1506382-ordered?language=objc
func (r_ RelationshipDescription) SetOrdered(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setOrdered:"), value)
}

// The type of object the relationship contains. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsrelationshipdescription/1506652-destinationentity?language=objc
func (r_ RelationshipDescription) DestinationEntity() EntityDescription {
	rv := objc.Call[EntityDescription](r_, objc.Sel("destinationEntity"))
	return rv
}

// The type of object the relationship contains. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsrelationshipdescription/1506652-destinationentity?language=objc
func (r_ RelationshipDescription) SetDestinationEntity(value IEntityDescription) {
	objc.Call[objc.Void](r_, objc.Sel("setDestinationEntity:"), objc.Ptr(value))
}

// The rule to apply when you delete the relationship’s owning managed object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsrelationshipdescription/1506870-deleterule?language=objc
func (r_ RelationshipDescription) DeleteRule() DeleteRule {
	rv := objc.Call[DeleteRule](r_, objc.Sel("deleteRule"))
	return rv
}

// The rule to apply when you delete the relationship’s owning managed object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsrelationshipdescription/1506870-deleterule?language=objc
func (r_ RelationshipDescription) SetDeleteRule(value DeleteRule) {
	objc.Call[objc.Void](r_, objc.Sel("setDeleteRule:"), value)
}

// The maximum number of managed objects the relationship can reference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsrelationshipdescription/1506780-maxcount?language=objc
func (r_ RelationshipDescription) MaxCount() uint {
	rv := objc.Call[uint](r_, objc.Sel("maxCount"))
	return rv
}

// The maximum number of managed objects the relationship can reference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsrelationshipdescription/1506780-maxcount?language=objc
func (r_ RelationshipDescription) SetMaxCount(value uint) {
	objc.Call[objc.Void](r_, objc.Sel("setMaxCount:"), value)
}

// Returns a Boolean value that indicates whether the relationship can contain many managed objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsrelationshipdescription/1506560-tomany?language=objc
func (r_ RelationshipDescription) IsToMany() bool {
	rv := objc.Call[bool](r_, objc.Sel("isToMany"))
	return rv
}

// The minimum number of managed objects the relationship can reference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsrelationshipdescription/1506599-mincount?language=objc
func (r_ RelationshipDescription) MinCount() uint {
	rv := objc.Call[uint](r_, objc.Sel("minCount"))
	return rv
}

// The minimum number of managed objects the relationship can reference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsrelationshipdescription/1506599-mincount?language=objc
func (r_ RelationshipDescription) SetMinCount(value uint) {
	objc.Call[objc.Void](r_, objc.Sel("setMinCount:"), value)
}

// The relationship that represents the inverse of the current relationship. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsrelationshipdescription/1506596-inverserelationship?language=objc
func (r_ RelationshipDescription) InverseRelationship() RelationshipDescription {
	rv := objc.Call[RelationshipDescription](r_, objc.Sel("inverseRelationship"))
	return rv
}

// The relationship that represents the inverse of the current relationship. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsrelationshipdescription/1506596-inverserelationship?language=objc
func (r_ RelationshipDescription) SetInverseRelationship(value IRelationshipDescription) {
	objc.Call[objc.Void](r_, objc.Sel("setInverseRelationship:"), objc.Ptr(value))
}
