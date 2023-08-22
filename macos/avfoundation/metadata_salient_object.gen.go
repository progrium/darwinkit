// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MetadataSalientObject] class.
var MetadataSalientObjectClass = _MetadataSalientObjectClass{objc.GetClass("AVMetadataSalientObject")}

type _MetadataSalientObjectClass struct {
	objc.Class
}

// An interface definition for the [MetadataSalientObject] class.
type IMetadataSalientObject interface {
	IMetadataObject
	ObjectID() int
}

// An object representing a single salient area in a picture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatasalientobject?language=objc
type MetadataSalientObject struct {
	MetadataObject
}

func MetadataSalientObjectFrom(ptr unsafe.Pointer) MetadataSalientObject {
	return MetadataSalientObject{
		MetadataObject: MetadataObjectFrom(ptr),
	}
}

func (mc _MetadataSalientObjectClass) Alloc() MetadataSalientObject {
	rv := objc.Call[MetadataSalientObject](mc, objc.Sel("alloc"))
	return rv
}

func MetadataSalientObject_Alloc() MetadataSalientObject {
	return MetadataSalientObjectClass.Alloc()
}

func (mc _MetadataSalientObjectClass) New() MetadataSalientObject {
	rv := objc.Call[MetadataSalientObject](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMetadataSalientObject() MetadataSalientObject {
	return MetadataSalientObjectClass.New()
}

func (m_ MetadataSalientObject) Init() MetadataSalientObject {
	rv := objc.Call[MetadataSalientObject](m_, objc.Sel("init"))
	return rv
}

// An integer value that defines the unique identifier of an object in a picture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatasalientobject/3152378-objectid?language=objc
func (m_ MetadataSalientObject) ObjectID() int {
	rv := objc.Call[int](m_, objc.Sel("objectID"))
	return rv
}
