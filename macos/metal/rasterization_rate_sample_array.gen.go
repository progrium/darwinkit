// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RasterizationRateSampleArray] class.
var RasterizationRateSampleArrayClass = _RasterizationRateSampleArrayClass{objc.GetClass("MTLRasterizationRateSampleArray")}

type _RasterizationRateSampleArrayClass struct {
	objc.Class
}

// An interface definition for the [RasterizationRateSampleArray] class.
type IRasterizationRateSampleArray interface {
	objc.IObject
	ObjectAtIndexedSubscript(index uint) foundation.Number
	SetObjectAtIndexedSubscript(value foundation.INumber, index uint)
}

// An array object that contains rasterization rates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratesamplearray?language=objc
type RasterizationRateSampleArray struct {
	objc.Object
}

func RasterizationRateSampleArrayFrom(ptr unsafe.Pointer) RasterizationRateSampleArray {
	return RasterizationRateSampleArray{
		Object: objc.ObjectFrom(ptr),
	}
}

func (rc _RasterizationRateSampleArrayClass) Alloc() RasterizationRateSampleArray {
	rv := objc.Call[RasterizationRateSampleArray](rc, objc.Sel("alloc"))
	return rv
}

func RasterizationRateSampleArray_Alloc() RasterizationRateSampleArray {
	return RasterizationRateSampleArrayClass.Alloc()
}

func (rc _RasterizationRateSampleArrayClass) New() RasterizationRateSampleArray {
	rv := objc.Call[RasterizationRateSampleArray](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRasterizationRateSampleArray() RasterizationRateSampleArray {
	return RasterizationRateSampleArrayClass.New()
}

func (r_ RasterizationRateSampleArray) Init() RasterizationRateSampleArray {
	rv := objc.Call[RasterizationRateSampleArray](r_, objc.Sel("init"))
	return rv
}

// Retrieves the sample value at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratesamplearray/3088875-objectatindexedsubscript?language=objc
func (r_ RasterizationRateSampleArray) ObjectAtIndexedSubscript(index uint) foundation.Number {
	rv := objc.Call[foundation.Number](r_, objc.Sel("objectAtIndexedSubscript:"), index)
	return rv
}

// Stores a sample value at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratesamplearray/3088876-setobject?language=objc
func (r_ RasterizationRateSampleArray) SetObjectAtIndexedSubscript(value foundation.INumber, index uint) {
	objc.Call[objc.Void](r_, objc.Sel("setObject:atIndexedSubscript:"), objc.Ptr(value), index)
}
