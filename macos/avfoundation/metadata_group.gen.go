// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MetadataGroup] class.
var MetadataGroupClass = _MetadataGroupClass{objc.GetClass("AVMetadataGroup")}

type _MetadataGroupClass struct {
	objc.Class
}

// An interface definition for the [MetadataGroup] class.
type IMetadataGroup interface {
	objc.IObject
	Items() []MetadataItem
	ClassifyingLabel() string
	UniqueID() string
}

// A collection of metadata items associated with a timeline segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatagroup?language=objc
type MetadataGroup struct {
	objc.Object
}

func MetadataGroupFrom(ptr unsafe.Pointer) MetadataGroup {
	return MetadataGroup{
		Object: objc.ObjectFrom(ptr),
	}
}

func (mc _MetadataGroupClass) Alloc() MetadataGroup {
	rv := objc.Call[MetadataGroup](mc, objc.Sel("alloc"))
	return rv
}

func MetadataGroup_Alloc() MetadataGroup {
	return MetadataGroupClass.Alloc()
}

func (mc _MetadataGroupClass) New() MetadataGroup {
	rv := objc.Call[MetadataGroup](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMetadataGroup() MetadataGroup {
	return MetadataGroupClass.New()
}

func (m_ MetadataGroup) Init() MetadataGroup {
	rv := objc.Call[MetadataGroup](m_, objc.Sel("init"))
	return rv
}

// The array of metadata items associated with the metadata group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatagroup/1389935-items?language=objc
func (m_ MetadataGroup) Items() []MetadataItem {
	rv := objc.Call[[]MetadataItem](m_, objc.Sel("items"))
	return rv
}

// The classifying label associated with the metadata group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatagroup/1620087-classifyinglabel?language=objc
func (m_ MetadataGroup) ClassifyingLabel() string {
	rv := objc.Call[string](m_, objc.Sel("classifyingLabel"))
	return rv
}

// The unique identifier for the metadata group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatagroup/1620088-uniqueid?language=objc
func (m_ MetadataGroup) UniqueID() string {
	rv := objc.Call[string](m_, objc.Sel("uniqueID"))
	return rv
}
