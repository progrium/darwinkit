// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AttributeDescriptor] class.
var AttributeDescriptorClass = _AttributeDescriptorClass{objc.GetClass("MTLAttributeDescriptor")}

type _AttributeDescriptorClass struct {
	objc.Class
}

// An interface definition for the [AttributeDescriptor] class.
type IAttributeDescriptor interface {
	objc.IObject
	BufferIndex() uint
	SetBufferIndex(value uint)
	Offset() uint
	SetOffset(value uint)
	Format() AttributeFormat
	SetFormat(value AttributeFormat)
}

// An object that describes an argument's format and where its data is in memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlattributedescriptor?language=objc
type AttributeDescriptor struct {
	objc.Object
}

func AttributeDescriptorFrom(ptr unsafe.Pointer) AttributeDescriptor {
	return AttributeDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AttributeDescriptorClass) Alloc() AttributeDescriptor {
	rv := objc.Call[AttributeDescriptor](ac, objc.Sel("alloc"))
	return rv
}

func AttributeDescriptor_Alloc() AttributeDescriptor {
	return AttributeDescriptorClass.Alloc()
}

func (ac _AttributeDescriptorClass) New() AttributeDescriptor {
	rv := objc.Call[AttributeDescriptor](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAttributeDescriptor() AttributeDescriptor {
	return AttributeDescriptorClass.New()
}

func (a_ AttributeDescriptor) Init() AttributeDescriptor {
	rv := objc.Call[AttributeDescriptor](a_, objc.Sel("init"))
	return rv
}

// The index in the argument table for the buffer that contains the data for the attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlattributedescriptor/2097218-bufferindex?language=objc
func (a_ AttributeDescriptor) BufferIndex() uint {
	rv := objc.Call[uint](a_, objc.Sel("bufferIndex"))
	return rv
}

// The index in the argument table for the buffer that contains the data for the attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlattributedescriptor/2097218-bufferindex?language=objc
func (a_ AttributeDescriptor) SetBufferIndex(value uint) {
	objc.Call[objc.Void](a_, objc.Sel("setBufferIndex:"), value)
}

// The offset of the data from the start of the buffer data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlattributedescriptor/2097220-offset?language=objc
func (a_ AttributeDescriptor) Offset() uint {
	rv := objc.Call[uint](a_, objc.Sel("offset"))
	return rv
}

// The offset of the data from the start of the buffer data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlattributedescriptor/2097220-offset?language=objc
func (a_ AttributeDescriptor) SetOffset(value uint) {
	objc.Call[objc.Void](a_, objc.Sel("setOffset:"), value)
}

// The format of the attribute's data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlattributedescriptor/2097194-format?language=objc
func (a_ AttributeDescriptor) Format() AttributeFormat {
	rv := objc.Call[AttributeFormat](a_, objc.Sel("format"))
	return rv
}

// The format of the attribute's data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlattributedescriptor/2097194-format?language=objc
func (a_ AttributeDescriptor) SetFormat(value AttributeFormat) {
	objc.Call[objc.Void](a_, objc.Sel("setFormat:"), value)
}
