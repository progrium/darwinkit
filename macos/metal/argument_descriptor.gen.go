// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ArgumentDescriptor] class.
var ArgumentDescriptorClass = _ArgumentDescriptorClass{objc.GetClass("MTLArgumentDescriptor")}

type _ArgumentDescriptorClass struct {
	objc.Class
}

// An interface definition for the [ArgumentDescriptor] class.
type IArgumentDescriptor interface {
	objc.IObject
	ArrayLength() uint
	SetArrayLength(value uint)
	ConstantBlockAlignment() uint
	SetConstantBlockAlignment(value uint)
	TextureType() TextureType
	SetTextureType(value TextureType)
	Access() objc.Object
	SetAccess(value objc.IObject)
	DataType() DataType
	SetDataType(value DataType)
	Index() uint
	SetIndex(value uint)
}

// A representation of an argument within an argument buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentdescriptor?language=objc
type ArgumentDescriptor struct {
	objc.Object
}

func ArgumentDescriptorFrom(ptr unsafe.Pointer) ArgumentDescriptor {
	return ArgumentDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _ArgumentDescriptorClass) Alloc() ArgumentDescriptor {
	rv := objc.Call[ArgumentDescriptor](ac, objc.Sel("alloc"))
	return rv
}

func ArgumentDescriptor_Alloc() ArgumentDescriptor {
	return ArgumentDescriptorClass.Alloc()
}

func (ac _ArgumentDescriptorClass) New() ArgumentDescriptor {
	rv := objc.Call[ArgumentDescriptor](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewArgumentDescriptor() ArgumentDescriptor {
	return ArgumentDescriptorClass.New()
}

func (a_ ArgumentDescriptor) Init() ArgumentDescriptor {
	rv := objc.Call[ArgumentDescriptor](a_, objc.Sel("init"))
	return rv
}

// Creates an empty argument descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentdescriptor/2915746-argumentdescriptor?language=objc
func (ac _ArgumentDescriptorClass) ArgumentDescriptor() ArgumentDescriptor {
	rv := objc.Call[ArgumentDescriptor](ac, objc.Sel("argumentDescriptor"))
	return rv
}

// Creates an empty argument descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentdescriptor/2915746-argumentdescriptor?language=objc
func ArgumentDescriptor_ArgumentDescriptor() ArgumentDescriptor {
	return ArgumentDescriptorClass.ArgumentDescriptor()
}

// The length of an array argument. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentdescriptor/2915734-arraylength?language=objc
func (a_ ArgumentDescriptor) ArrayLength() uint {
	rv := objc.Call[uint](a_, objc.Sel("arrayLength"))
	return rv
}

// The length of an array argument. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentdescriptor/2915734-arraylength?language=objc
func (a_ ArgumentDescriptor) SetArrayLength(value uint) {
	objc.Call[objc.Void](a_, objc.Sel("setArrayLength:"), value)
}

// The alignment of the constant block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentdescriptor/2915739-constantblockalignment?language=objc
func (a_ ArgumentDescriptor) ConstantBlockAlignment() uint {
	rv := objc.Call[uint](a_, objc.Sel("constantBlockAlignment"))
	return rv
}

// The alignment of the constant block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentdescriptor/2915739-constantblockalignment?language=objc
func (a_ ArgumentDescriptor) SetConstantBlockAlignment(value uint) {
	objc.Call[objc.Void](a_, objc.Sel("setConstantBlockAlignment:"), value)
}

// The texture type of a texture argument. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentdescriptor/2915741-texturetype?language=objc
func (a_ ArgumentDescriptor) TextureType() TextureType {
	rv := objc.Call[TextureType](a_, objc.Sel("textureType"))
	return rv
}

// The texture type of a texture argument. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentdescriptor/2915741-texturetype?language=objc
func (a_ ArgumentDescriptor) SetTextureType(value TextureType) {
	objc.Call[objc.Void](a_, objc.Sel("setTextureType:"), value)
}

// The access permissions of the argument. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentdescriptor/2915735-access?language=objc
func (a_ ArgumentDescriptor) Access() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("access"))
	return rv
}

// The access permissions of the argument. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentdescriptor/2915735-access?language=objc
func (a_ ArgumentDescriptor) SetAccess(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setAccess:"), objc.Ptr(value))
}

// The data type of the argument. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentdescriptor/2915733-datatype?language=objc
func (a_ ArgumentDescriptor) DataType() DataType {
	rv := objc.Call[DataType](a_, objc.Sel("dataType"))
	return rv
}

// The data type of the argument. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentdescriptor/2915733-datatype?language=objc
func (a_ ArgumentDescriptor) SetDataType(value DataType) {
	objc.Call[objc.Void](a_, objc.Sel("setDataType:"), value)
}

// The index ID of the argument. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentdescriptor/2915732-index?language=objc
func (a_ ArgumentDescriptor) Index() uint {
	rv := objc.Call[uint](a_, objc.Sel("index"))
	return rv
}

// The index ID of the argument. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentdescriptor/2915732-index?language=objc
func (a_ ArgumentDescriptor) SetIndex(value uint) {
	objc.Call[objc.Void](a_, objc.Sel("setIndex:"), value)
}
