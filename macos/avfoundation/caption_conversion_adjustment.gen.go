// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CaptionConversionAdjustment] class.
var CaptionConversionAdjustmentClass = _CaptionConversionAdjustmentClass{objc.GetClass("AVCaptionConversionAdjustment")}

type _CaptionConversionAdjustmentClass struct {
	objc.Class
}

// An interface definition for the [CaptionConversionAdjustment] class.
type ICaptionConversionAdjustment interface {
	objc.IObject
	AdjustmentType() CaptionConversionAdjustmentType
}

// An object that describes an adjustment to correct a problem found during validation of a caption conversion. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionadjustment?language=objc
type CaptionConversionAdjustment struct {
	objc.Object
}

func CaptionConversionAdjustmentFrom(ptr unsafe.Pointer) CaptionConversionAdjustment {
	return CaptionConversionAdjustment{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CaptionConversionAdjustmentClass) Alloc() CaptionConversionAdjustment {
	rv := objc.Call[CaptionConversionAdjustment](cc, objc.Sel("alloc"))
	return rv
}

func CaptionConversionAdjustment_Alloc() CaptionConversionAdjustment {
	return CaptionConversionAdjustmentClass.Alloc()
}

func (cc _CaptionConversionAdjustmentClass) New() CaptionConversionAdjustment {
	rv := objc.Call[CaptionConversionAdjustment](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptionConversionAdjustment() CaptionConversionAdjustment {
	return CaptionConversionAdjustmentClass.New()
}

func (c_ CaptionConversionAdjustment) Init() CaptionConversionAdjustment {
	rv := objc.Call[CaptionConversionAdjustment](c_, objc.Sel("init"))
	return rv
}

// The type of caption conversion adjustment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionadjustment/3752928-adjustmenttype?language=objc
func (c_ CaptionConversionAdjustment) AdjustmentType() CaptionConversionAdjustmentType {
	rv := objc.Call[CaptionConversionAdjustmentType](c_, objc.Sel("adjustmentType"))
	return rv
}
