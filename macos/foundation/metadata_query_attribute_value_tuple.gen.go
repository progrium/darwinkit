// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MetadataQueryAttributeValueTuple] class.
var MetadataQueryAttributeValueTupleClass = _MetadataQueryAttributeValueTupleClass{objc.GetClass("NSMetadataQueryAttributeValueTuple")}

type _MetadataQueryAttributeValueTupleClass struct {
	objc.Class
}

// An interface definition for the [MetadataQueryAttributeValueTuple] class.
type IMetadataQueryAttributeValueTuple interface {
	objc.IObject
	Value() objc.Object
	Attribute() string
	Count() uint
}

// The NSMetadataQueryAttributeValueTuple class represents attribute-value tuples, which are objects that contain the attribute name and value of a metadata attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataqueryattributevaluetuple?language=objc
type MetadataQueryAttributeValueTuple struct {
	objc.Object
}

func MetadataQueryAttributeValueTupleFrom(ptr unsafe.Pointer) MetadataQueryAttributeValueTuple {
	return MetadataQueryAttributeValueTuple{
		Object: objc.ObjectFrom(ptr),
	}
}

func (mc _MetadataQueryAttributeValueTupleClass) Alloc() MetadataQueryAttributeValueTuple {
	rv := objc.Call[MetadataQueryAttributeValueTuple](mc, objc.Sel("alloc"))
	return rv
}

func MetadataQueryAttributeValueTuple_Alloc() MetadataQueryAttributeValueTuple {
	return MetadataQueryAttributeValueTupleClass.Alloc()
}

func (mc _MetadataQueryAttributeValueTupleClass) New() MetadataQueryAttributeValueTuple {
	rv := objc.Call[MetadataQueryAttributeValueTuple](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMetadataQueryAttributeValueTuple() MetadataQueryAttributeValueTuple {
	return MetadataQueryAttributeValueTupleClass.New()
}

func (m_ MetadataQueryAttributeValueTuple) Init() MetadataQueryAttributeValueTuple {
	rv := objc.Call[MetadataQueryAttributeValueTuple](m_, objc.Sel("init"))
	return rv
}

// The value of the tuple’s attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataqueryattributevaluetuple/1417894-value?language=objc
func (m_ MetadataQueryAttributeValueTuple) Value() objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("value"))
	return rv
}

// The attribute name for the tuple’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataqueryattributevaluetuple/1415060-attribute?language=objc
func (m_ MetadataQueryAttributeValueTuple) Attribute() string {
	rv := objc.Call[string](m_, objc.Sel("attribute"))
	return rv
}

// The number of instances of the value for the tuple’s attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataqueryattributevaluetuple/1414426-count?language=objc
func (m_ MetadataQueryAttributeValueTuple) Count() uint {
	rv := objc.Call[uint](m_, objc.Sel("count"))
	return rv
}
