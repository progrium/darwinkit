// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CaptionConversionWarning] class.
var CaptionConversionWarningClass = _CaptionConversionWarningClass{objc.GetClass("AVCaptionConversionWarning")}

type _CaptionConversionWarningClass struct {
	objc.Class
}

// An interface definition for the [CaptionConversionWarning] class.
type ICaptionConversionWarning interface {
	objc.IObject
	WarningType() CaptionConversionWarningType
	RangeOfCaptions() foundation.Range
	Adjustment() CaptionConversionAdjustment
}

// An object that represents a conversion warning produced by a validator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionwarning?language=objc
type CaptionConversionWarning struct {
	objc.Object
}

func CaptionConversionWarningFrom(ptr unsafe.Pointer) CaptionConversionWarning {
	return CaptionConversionWarning{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CaptionConversionWarningClass) Alloc() CaptionConversionWarning {
	rv := objc.Call[CaptionConversionWarning](cc, objc.Sel("alloc"))
	return rv
}

func CaptionConversionWarning_Alloc() CaptionConversionWarning {
	return CaptionConversionWarningClass.Alloc()
}

func (cc _CaptionConversionWarningClass) New() CaptionConversionWarning {
	rv := objc.Call[CaptionConversionWarning](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptionConversionWarning() CaptionConversionWarning {
	return CaptionConversionWarningClass.New()
}

func (c_ CaptionConversionWarning) Init() CaptionConversionWarning {
	rv := objc.Call[CaptionConversionWarning](c_, objc.Sel("init"))
	return rv
}

// A type that indicates the nature of the validation warning. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionwarning/3752951-warningtype?language=objc
func (c_ CaptionConversionWarning) WarningType() CaptionConversionWarningType {
	rv := objc.Call[CaptionConversionWarningType](c_, objc.Sel("warningType"))
	return rv
}

// The range of the captions for which the system issued a warning. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionwarning/3752950-rangeofcaptions?language=objc
func (c_ CaptionConversionWarning) RangeOfCaptions() foundation.Range {
	rv := objc.Call[foundation.Range](c_, objc.Sel("rangeOfCaptions"))
	return rv
}

// A correction the converter makes when it converts a caption to a specific format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionwarning/3752949-adjustment?language=objc
func (c_ CaptionConversionWarning) Adjustment() CaptionConversionAdjustment {
	rv := objc.Call[CaptionConversionAdjustment](c_, objc.Sel("adjustment"))
	return rv
}
