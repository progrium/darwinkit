// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CaptionFormatConformer] class.
var CaptionFormatConformerClass = _CaptionFormatConformerClass{objc.GetClass("AVCaptionFormatConformer")}

type _CaptionFormatConformerClass struct {
	objc.Class
}

// An interface definition for the [CaptionFormatConformer] class.
type ICaptionFormatConformer interface {
	objc.IObject
	ConformedCaptionForCaptionError(caption ICaption, outError foundation.IError) Caption
	ConformsCaptionsToTimeRange() bool
	SetConformsCaptionsToTimeRange(value bool)
}

// An object that converts a canonical caption to a specific format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionformatconformer?language=objc
type CaptionFormatConformer struct {
	objc.Object
}

func CaptionFormatConformerFrom(ptr unsafe.Pointer) CaptionFormatConformer {
	return CaptionFormatConformer{
		Object: objc.ObjectFrom(ptr),
	}
}

func (c_ CaptionFormatConformer) InitWithConversionSettings(conversionSettings map[CaptionSettingsKey]objc.IObject) CaptionFormatConformer {
	rv := objc.Call[CaptionFormatConformer](c_, objc.Sel("initWithConversionSettings:"), conversionSettings)
	return rv
}

// Creates a new object with format conversion settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionformatconformer/3752958-initwithconversionsettings?language=objc
func NewCaptionFormatConformerWithConversionSettings(conversionSettings map[CaptionSettingsKey]objc.IObject) CaptionFormatConformer {
	instance := CaptionFormatConformerClass.Alloc().InitWithConversionSettings(conversionSettings)
	instance.Autorelease()
	return instance
}

func (cc _CaptionFormatConformerClass) CaptionFormatConformerWithConversionSettings(conversionSettings map[CaptionSettingsKey]objc.IObject) CaptionFormatConformer {
	rv := objc.Call[CaptionFormatConformer](cc, objc.Sel("captionFormatConformerWithConversionSettings:"), conversionSettings)
	return rv
}

// A class method that creates a new object with format conversion settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionformatconformer/3752955-captionformatconformerwithconver?language=objc
func CaptionFormatConformer_CaptionFormatConformerWithConversionSettings(conversionSettings map[CaptionSettingsKey]objc.IObject) CaptionFormatConformer {
	return CaptionFormatConformerClass.CaptionFormatConformerWithConversionSettings(conversionSettings)
}

func (cc _CaptionFormatConformerClass) Alloc() CaptionFormatConformer {
	rv := objc.Call[CaptionFormatConformer](cc, objc.Sel("alloc"))
	return rv
}

func CaptionFormatConformer_Alloc() CaptionFormatConformer {
	return CaptionFormatConformerClass.Alloc()
}

func (cc _CaptionFormatConformerClass) New() CaptionFormatConformer {
	rv := objc.Call[CaptionFormatConformer](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptionFormatConformer() CaptionFormatConformer {
	return CaptionFormatConformerClass.New()
}

func (c_ CaptionFormatConformer) Init() CaptionFormatConformer {
	rv := objc.Call[CaptionFormatConformer](c_, objc.Sel("init"))
	return rv
}

// Creates a caption that conforms to a specific format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionformatconformer/3752956-conformedcaptionforcaption?language=objc
func (c_ CaptionFormatConformer) ConformedCaptionForCaptionError(caption ICaption, outError foundation.IError) Caption {
	rv := objc.Call[Caption](c_, objc.Sel("conformedCaptionForCaption:error:"), objc.Ptr(caption), objc.Ptr(outError))
	return rv
}

// A Boolean value that indicates whether to conform the time range of a canonical caption. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionformatconformer/3857642-conformscaptionstotimerange?language=objc
func (c_ CaptionFormatConformer) ConformsCaptionsToTimeRange() bool {
	rv := objc.Call[bool](c_, objc.Sel("conformsCaptionsToTimeRange"))
	return rv
}

// A Boolean value that indicates whether to conform the time range of a canonical caption. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionformatconformer/3857642-conformscaptionstotimerange?language=objc
func (c_ CaptionFormatConformer) SetConformsCaptionsToTimeRange(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setConformsCaptionsToTimeRange:"), value)
}
