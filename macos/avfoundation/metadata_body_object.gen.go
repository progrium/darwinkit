// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MetadataBodyObject] class.
var MetadataBodyObjectClass = _MetadataBodyObjectClass{objc.GetClass("AVMetadataBodyObject")}

type _MetadataBodyObjectClass struct {
	objc.Class
}

// An interface definition for the [MetadataBodyObject] class.
type IMetadataBodyObject interface {
	IMetadataObject
	BodyID() int
}

// An abstract class that defines the interface for a metadata body object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatabodyobject?language=objc
type MetadataBodyObject struct {
	MetadataObject
}

func MetadataBodyObjectFrom(ptr unsafe.Pointer) MetadataBodyObject {
	return MetadataBodyObject{
		MetadataObject: MetadataObjectFrom(ptr),
	}
}

func (mc _MetadataBodyObjectClass) Alloc() MetadataBodyObject {
	rv := objc.Call[MetadataBodyObject](mc, objc.Sel("alloc"))
	return rv
}

func MetadataBodyObject_Alloc() MetadataBodyObject {
	return MetadataBodyObjectClass.Alloc()
}

func (mc _MetadataBodyObjectClass) New() MetadataBodyObject {
	rv := objc.Call[MetadataBodyObject](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMetadataBodyObject() MetadataBodyObject {
	return MetadataBodyObjectClass.New()
}

func (m_ MetadataBodyObject) Init() MetadataBodyObject {
	rv := objc.Call[MetadataBodyObject](m_, objc.Sel("init"))
	return rv
}

// An integer value that defines the unique identifier of an object in a picture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatabodyobject/3174758-bodyid?language=objc
func (m_ MetadataBodyObject) BodyID() int {
	rv := objc.Call[int](m_, objc.Sel("bodyID"))
	return rv
}
