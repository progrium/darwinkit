// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CapturePhotoSettings] class.
var CapturePhotoSettingsClass = _CapturePhotoSettingsClass{objc.GetClass("AVCapturePhotoSettings")}

type _CapturePhotoSettingsClass struct {
	objc.Class
}

// An interface definition for the [CapturePhotoSettings] class.
type ICapturePhotoSettings interface {
	objc.IObject
	UniqueID() int64
	Format() map[string]objc.Object
	ProcessedFileType() FileType
}

// A specification of the features and settings to use for a single photo capture request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings?language=objc
type CapturePhotoSettings struct {
	objc.Object
}

func CapturePhotoSettingsFrom(ptr unsafe.Pointer) CapturePhotoSettings {
	return CapturePhotoSettings{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CapturePhotoSettingsClass) PhotoSettingsFromPhotoSettings(photoSettings ICapturePhotoSettings) CapturePhotoSettings {
	rv := objc.Call[CapturePhotoSettings](cc, objc.Sel("photoSettingsFromPhotoSettings:"), objc.Ptr(photoSettings))
	return rv
}

// Creates a unique photo settings object, copying all settings values from the specified photo settings object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/1778655-photosettingsfromphotosettings?language=objc
func CapturePhotoSettings_PhotoSettingsFromPhotoSettings(photoSettings ICapturePhotoSettings) CapturePhotoSettings {
	return CapturePhotoSettingsClass.PhotoSettingsFromPhotoSettings(photoSettings)
}

func (cc _CapturePhotoSettingsClass) PhotoSettings() CapturePhotoSettings {
	rv := objc.Call[CapturePhotoSettings](cc, objc.Sel("photoSettings"))
	return rv
}

// Creates a photo settings object with default settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/1649254-photosettings?language=objc
func CapturePhotoSettings_PhotoSettings() CapturePhotoSettings {
	return CapturePhotoSettingsClass.PhotoSettings()
}

func (cc _CapturePhotoSettingsClass) PhotoSettingsWithFormat(format map[string]objc.IObject) CapturePhotoSettings {
	rv := objc.Call[CapturePhotoSettings](cc, objc.Sel("photoSettingsWithFormat:"), format)
	return rv
}

// Creates a photo settings object with the specified output format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/1648673-photosettingswithformat?language=objc
func CapturePhotoSettings_PhotoSettingsWithFormat(format map[string]objc.IObject) CapturePhotoSettings {
	return CapturePhotoSettingsClass.PhotoSettingsWithFormat(format)
}

func (cc _CapturePhotoSettingsClass) Alloc() CapturePhotoSettings {
	rv := objc.Call[CapturePhotoSettings](cc, objc.Sel("alloc"))
	return rv
}

func CapturePhotoSettings_Alloc() CapturePhotoSettings {
	return CapturePhotoSettingsClass.Alloc()
}

func (cc _CapturePhotoSettingsClass) New() CapturePhotoSettings {
	rv := objc.Call[CapturePhotoSettings](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCapturePhotoSettings() CapturePhotoSettings {
	return CapturePhotoSettingsClass.New()
}

func (c_ CapturePhotoSettings) Init() CapturePhotoSettings {
	rv := objc.Call[CapturePhotoSettings](c_, objc.Sel("init"))
	return rv
}

// A unique identifier for this photo settings instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/1648767-uniqueid?language=objc
func (c_ CapturePhotoSettings) UniqueID() int64 {
	rv := objc.Call[int64](c_, objc.Sel("uniqueID"))
	return rv
}

// A dictionary describing the processed format (for example, JPEG) to deliver captured photos in. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/1648783-format?language=objc
func (c_ CapturePhotoSettings) Format() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](c_, objc.Sel("format"))
	return rv
}

// The container file format for eventual output of the processed image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/2873970-processedfiletype?language=objc
func (c_ CapturePhotoSettings) ProcessedFileType() FileType {
	rv := objc.Call[FileType](c_, objc.Sel("processedFileType"))
	return rv
}
