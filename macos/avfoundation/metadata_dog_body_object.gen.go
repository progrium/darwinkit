// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MetadataDogBodyObject] class.
var MetadataDogBodyObjectClass = _MetadataDogBodyObjectClass{objc.GetClass("AVMetadataDogBodyObject")}

type _MetadataDogBodyObjectClass struct {
	objc.Class
}

// An interface definition for the [MetadataDogBodyObject] class.
type IMetadataDogBodyObject interface {
	IMetadataBodyObject
}

// An object representing a single detected dog body in a picture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatadogbodyobject?language=objc
type MetadataDogBodyObject struct {
	MetadataBodyObject
}

func MetadataDogBodyObjectFrom(ptr unsafe.Pointer) MetadataDogBodyObject {
	return MetadataDogBodyObject{
		MetadataBodyObject: MetadataBodyObjectFrom(ptr),
	}
}

func (mc _MetadataDogBodyObjectClass) Alloc() MetadataDogBodyObject {
	rv := objc.Call[MetadataDogBodyObject](mc, objc.Sel("alloc"))
	return rv
}

func MetadataDogBodyObject_Alloc() MetadataDogBodyObject {
	return MetadataDogBodyObjectClass.Alloc()
}

func (mc _MetadataDogBodyObjectClass) New() MetadataDogBodyObject {
	rv := objc.Call[MetadataDogBodyObject](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMetadataDogBodyObject() MetadataDogBodyObject {
	return MetadataDogBodyObjectClass.New()
}

func (m_ MetadataDogBodyObject) Init() MetadataDogBodyObject {
	rv := objc.Call[MetadataDogBodyObject](m_, objc.Sel("init"))
	return rv
}
