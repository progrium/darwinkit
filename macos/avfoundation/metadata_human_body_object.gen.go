// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MetadataHumanBodyObject] class.
var MetadataHumanBodyObjectClass = _MetadataHumanBodyObjectClass{objc.GetClass("AVMetadataHumanBodyObject")}

type _MetadataHumanBodyObjectClass struct {
	objc.Class
}

// An interface definition for the [MetadataHumanBodyObject] class.
type IMetadataHumanBodyObject interface {
	IMetadataBodyObject
}

// An object representing a single detected human body in a picture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatahumanbodyobject?language=objc
type MetadataHumanBodyObject struct {
	MetadataBodyObject
}

func MetadataHumanBodyObjectFrom(ptr unsafe.Pointer) MetadataHumanBodyObject {
	return MetadataHumanBodyObject{
		MetadataBodyObject: MetadataBodyObjectFrom(ptr),
	}
}

func (mc _MetadataHumanBodyObjectClass) Alloc() MetadataHumanBodyObject {
	rv := objc.Call[MetadataHumanBodyObject](mc, objc.Sel("alloc"))
	return rv
}

func MetadataHumanBodyObject_Alloc() MetadataHumanBodyObject {
	return MetadataHumanBodyObjectClass.Alloc()
}

func (mc _MetadataHumanBodyObjectClass) New() MetadataHumanBodyObject {
	rv := objc.Call[MetadataHumanBodyObject](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMetadataHumanBodyObject() MetadataHumanBodyObject {
	return MetadataHumanBodyObjectClass.New()
}

func (m_ MetadataHumanBodyObject) Init() MetadataHumanBodyObject {
	rv := objc.Call[MetadataHumanBodyObject](m_, objc.Sel("init"))
	return rv
}
