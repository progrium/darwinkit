// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CapturePhotoOutput] class.
var CapturePhotoOutputClass = _CapturePhotoOutputClass{objc.GetClass("AVCapturePhotoOutput")}

type _CapturePhotoOutputClass struct {
	objc.Class
}

// An interface definition for the [CapturePhotoOutput] class.
type ICapturePhotoOutput interface {
	ICaptureOutput
	SupportedPhotoPixelFormatTypesForFileType(fileType FileType) []foundation.Number
	SupportedPhotoCodecTypesForFileType(fileType FileType) []VideoCodecType
	CapturePhotoWithSettingsDelegate(settings ICapturePhotoSettings, delegate PCapturePhotoCaptureDelegate)
	CapturePhotoWithSettingsDelegateObject(settings ICapturePhotoSettings, delegateObject objc.IObject)
	AvailablePhotoCodecTypes() []VideoCodecType
	AvailablePhotoPixelFormatTypes() []foundation.Number
	AvailablePhotoFileTypes() []FileType
}

// A capture output for still image, Live Photos, and other photography workflows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput?language=objc
type CapturePhotoOutput struct {
	CaptureOutput
}

func CapturePhotoOutputFrom(ptr unsafe.Pointer) CapturePhotoOutput {
	return CapturePhotoOutput{
		CaptureOutput: CaptureOutputFrom(ptr),
	}
}

func (cc _CapturePhotoOutputClass) New() CapturePhotoOutput {
	rv := objc.Call[CapturePhotoOutput](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCapturePhotoOutput() CapturePhotoOutput {
	return CapturePhotoOutputClass.New()
}

func (c_ CapturePhotoOutput) Init() CapturePhotoOutput {
	rv := objc.Call[CapturePhotoOutput](c_, objc.Sel("init"))
	return rv
}

func (cc _CapturePhotoOutputClass) Alloc() CapturePhotoOutput {
	rv := objc.Call[CapturePhotoOutput](cc, objc.Sel("alloc"))
	return rv
}

func CapturePhotoOutput_Alloc() CapturePhotoOutput {
	return CapturePhotoOutputClass.Alloc()
}

// Returns the list of uncompressed pixel formats supported for photo data in the specified file type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/2873950-supportedphotopixelformattypesfo?language=objc
func (c_ CapturePhotoOutput) SupportedPhotoPixelFormatTypesForFileType(fileType FileType) []foundation.Number {
	rv := objc.Call[[]foundation.Number](c_, objc.Sel("supportedPhotoPixelFormatTypesForFileType:"), fileType)
	return rv
}

// Returns the list of photo codecs (such as JPEG or HEVC) supported for photo data in the specified file type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/2873916-supportedphotocodectypesforfilet?language=objc
func (c_ CapturePhotoOutput) SupportedPhotoCodecTypesForFileType(fileType FileType) []VideoCodecType {
	rv := objc.Call[[]VideoCodecType](c_, objc.Sel("supportedPhotoCodecTypesForFileType:"), fileType)
	return rv
}

// Initiates a photo capture using the specified settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/1648765-capturephotowithsettings?language=objc
func (c_ CapturePhotoOutput) CapturePhotoWithSettingsDelegate(settings ICapturePhotoSettings, delegate PCapturePhotoCaptureDelegate) {
	po1 := objc.WrapAsProtocol("AVCapturePhotoCaptureDelegate", delegate)
	objc.Call[objc.Void](c_, objc.Sel("capturePhotoWithSettings:delegate:"), objc.Ptr(settings), po1)
}

// Initiates a photo capture using the specified settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/1648765-capturephotowithsettings?language=objc
func (c_ CapturePhotoOutput) CapturePhotoWithSettingsDelegateObject(settings ICapturePhotoSettings, delegateObject objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("capturePhotoWithSettings:delegate:"), objc.Ptr(settings), objc.Ptr(delegateObject))
}

// The compression codecs this capture output currently supports for photo capture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/1648654-availablephotocodectypes?language=objc
func (c_ CapturePhotoOutput) AvailablePhotoCodecTypes() []VideoCodecType {
	rv := objc.Call[[]VideoCodecType](c_, objc.Sel("availablePhotoCodecTypes"))
	return rv
}

// The pixel formats the capture output supports for photo capture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/1778630-availablephotopixelformattypes?language=objc
func (c_ CapturePhotoOutput) AvailablePhotoPixelFormatTypes() []foundation.Number {
	rv := objc.Call[[]foundation.Number](c_, objc.Sel("availablePhotoPixelFormatTypes"))
	return rv
}

// The list of file types currently supported for photo capture and output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/2873918-availablephotofiletypes?language=objc
func (c_ CapturePhotoOutput) AvailablePhotoFileTypes() []FileType {
	rv := objc.Call[[]FileType](c_, objc.Sel("availablePhotoFileTypes"))
	return rv
}
