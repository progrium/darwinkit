// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MetadataItemFilter] class.
var MetadataItemFilterClass = _MetadataItemFilterClass{objc.GetClass("AVMetadataItemFilter")}

type _MetadataItemFilterClass struct {
	objc.Class
}

// An interface definition for the [MetadataItemFilter] class.
type IMetadataItemFilter interface {
	objc.IObject
}

// An object that filters selected information from a metadata item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitemfilter?language=objc
type MetadataItemFilter struct {
	objc.Object
}

func MetadataItemFilterFrom(ptr unsafe.Pointer) MetadataItemFilter {
	return MetadataItemFilter{
		Object: objc.ObjectFrom(ptr),
	}
}

func (mc _MetadataItemFilterClass) Alloc() MetadataItemFilter {
	rv := objc.Call[MetadataItemFilter](mc, objc.Sel("alloc"))
	return rv
}

func MetadataItemFilter_Alloc() MetadataItemFilter {
	return MetadataItemFilterClass.Alloc()
}

func (mc _MetadataItemFilterClass) New() MetadataItemFilter {
	rv := objc.Call[MetadataItemFilter](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMetadataItemFilter() MetadataItemFilter {
	return MetadataItemFilterClass.New()
}

func (m_ MetadataItemFilter) Init() MetadataItemFilter {
	rv := objc.Call[MetadataItemFilter](m_, objc.Sel("init"))
	return rv
}

// Returns a metadata filter to use for sharing assets. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitemfilter/1387905-metadataitemfilterforsharing?language=objc
func (mc _MetadataItemFilterClass) MetadataItemFilterForSharing() MetadataItemFilter {
	rv := objc.Call[MetadataItemFilter](mc, objc.Sel("metadataItemFilterForSharing"))
	return rv
}

// Returns a metadata filter to use for sharing assets. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitemfilter/1387905-metadataitemfilterforsharing?language=objc
func MetadataItemFilter_MetadataItemFilterForSharing() MetadataItemFilter {
	return MetadataItemFilterClass.MetadataItemFilterForSharing()
}
