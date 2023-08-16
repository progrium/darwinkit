// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AttributeDescriptorArray] class.
var AttributeDescriptorArrayClass = _AttributeDescriptorArrayClass{objc.GetClass("MTLAttributeDescriptorArray")}

type _AttributeDescriptorArrayClass struct {
	objc.Class
}

// An interface definition for the [AttributeDescriptorArray] class.
type IAttributeDescriptorArray interface {
	objc.IObject
	ObjectAtIndexedSubscript(index uint) AttributeDescriptor
	SetObjectAtIndexedSubscript(attributeDesc IAttributeDescriptor, index uint)
}

// An array of attribute descriptor objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlattributedescriptorarray?language=objc
type AttributeDescriptorArray struct {
	objc.Object
}

func AttributeDescriptorArrayFrom(ptr unsafe.Pointer) AttributeDescriptorArray {
	return AttributeDescriptorArray{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AttributeDescriptorArrayClass) Alloc() AttributeDescriptorArray {
	rv := objc.Call[AttributeDescriptorArray](ac, objc.Sel("alloc"))
	return rv
}

func AttributeDescriptorArray_Alloc() AttributeDescriptorArray {
	return AttributeDescriptorArrayClass.Alloc()
}

func (ac _AttributeDescriptorArrayClass) New() AttributeDescriptorArray {
	rv := objc.Call[AttributeDescriptorArray](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAttributeDescriptorArray() AttributeDescriptorArray {
	return AttributeDescriptorArrayClass.New()
}

func (a_ AttributeDescriptorArray) Init() AttributeDescriptorArray {
	rv := objc.Call[AttributeDescriptorArray](a_, objc.Sel("init"))
	return rv
}

// Returns the state of the specified attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlattributedescriptorarray/2097215-objectatindexedsubscript?language=objc
func (a_ AttributeDescriptorArray) ObjectAtIndexedSubscript(index uint) AttributeDescriptor {
	rv := objc.Call[AttributeDescriptor](a_, objc.Sel("objectAtIndexedSubscript:"), index)
	return rv
}

// Sets state for the specified attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlattributedescriptorarray/2097293-setobject?language=objc
func (a_ AttributeDescriptorArray) SetObjectAtIndexedSubscript(attributeDesc IAttributeDescriptor, index uint) {
	objc.Call[objc.Void](a_, objc.Sel("setObject:atIndexedSubscript:"), objc.Ptr(attributeDesc), index)
}
