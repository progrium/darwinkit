// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MetadataCatBodyObject] class.
var MetadataCatBodyObjectClass = _MetadataCatBodyObjectClass{objc.GetClass("AVMetadataCatBodyObject")}

type _MetadataCatBodyObjectClass struct {
	objc.Class
}

// An interface definition for the [MetadataCatBodyObject] class.
type IMetadataCatBodyObject interface {
	IMetadataBodyObject
}

// An object representing a single detected cat body in a picture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatacatbodyobject?language=objc
type MetadataCatBodyObject struct {
	MetadataBodyObject
}

func MetadataCatBodyObjectFrom(ptr unsafe.Pointer) MetadataCatBodyObject {
	return MetadataCatBodyObject{
		MetadataBodyObject: MetadataBodyObjectFrom(ptr),
	}
}

func (mc _MetadataCatBodyObjectClass) Alloc() MetadataCatBodyObject {
	rv := objc.Call[MetadataCatBodyObject](mc, objc.Sel("alloc"))
	return rv
}

func MetadataCatBodyObject_Alloc() MetadataCatBodyObject {
	return MetadataCatBodyObjectClass.Alloc()
}

func (mc _MetadataCatBodyObjectClass) New() MetadataCatBodyObject {
	rv := objc.Call[MetadataCatBodyObject](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMetadataCatBodyObject() MetadataCatBodyObject {
	return MetadataCatBodyObjectClass.New()
}

func (m_ MetadataCatBodyObject) Init() MetadataCatBodyObject {
	rv := objc.Call[MetadataCatBodyObject](m_, objc.Sel("init"))
	return rv
}
