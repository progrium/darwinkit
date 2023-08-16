// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TextureReferenceType] class.
var TextureReferenceTypeClass = _TextureReferenceTypeClass{objc.GetClass("MTLTextureReferenceType")}

type _TextureReferenceTypeClass struct {
	objc.Class
}

// An interface definition for the [TextureReferenceType] class.
type ITextureReferenceType interface {
	IType
	TextureType() TextureType
	Access() objc.Object
	IsDepthTexture() bool
	TextureDataType() DataType
}

// A description of a texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturereferencetype?language=objc
type TextureReferenceType struct {
	Type
}

func TextureReferenceTypeFrom(ptr unsafe.Pointer) TextureReferenceType {
	return TextureReferenceType{
		Type: TypeFrom(ptr),
	}
}

func (tc _TextureReferenceTypeClass) Alloc() TextureReferenceType {
	rv := objc.Call[TextureReferenceType](tc, objc.Sel("alloc"))
	return rv
}

func TextureReferenceType_Alloc() TextureReferenceType {
	return TextureReferenceTypeClass.Alloc()
}

func (tc _TextureReferenceTypeClass) New() TextureReferenceType {
	rv := objc.Call[TextureReferenceType](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTextureReferenceType() TextureReferenceType {
	return TextureReferenceTypeClass.New()
}

func (t_ TextureReferenceType) Init() TextureReferenceType {
	rv := objc.Call[TextureReferenceType](t_, objc.Sel("init"))
	return rv
}

// The texture type of the texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturereferencetype/2877442-texturetype?language=objc
func (t_ TextureReferenceType) TextureType() TextureType {
	rv := objc.Call[TextureType](t_, objc.Sel("textureType"))
	return rv
}

// The texture's read/write access to the argument. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturereferencetype/2877456-access?language=objc
func (t_ TextureReferenceType) Access() objc.Object {
	rv := objc.Call[objc.Object](t_, objc.Sel("access"))
	return rv
}

// A Boolean value that indicates whether the texture is a depth texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturereferencetype/2877449-isdepthtexture?language=objc
func (t_ TextureReferenceType) IsDepthTexture() bool {
	rv := objc.Call[bool](t_, objc.Sel("isDepthTexture"))
	return rv
}

// The data type of the texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturereferencetype/2877443-texturedatatype?language=objc
func (t_ TextureReferenceType) TextureDataType() DataType {
	rv := objc.Call[DataType](t_, objc.Sel("textureDataType"))
	return rv
}
