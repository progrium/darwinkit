// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/corevideo"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/imageio"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PortraitEffectsMatte] class.
var PortraitEffectsMatteClass = _PortraitEffectsMatteClass{objc.GetClass("AVPortraitEffectsMatte")}

type _PortraitEffectsMatteClass struct {
	objc.Class
}

// An interface definition for the [PortraitEffectsMatte] class.
type IPortraitEffectsMatte interface {
	objc.IObject
	DictionaryRepresentationForAuxiliaryDataType(outAuxDataType string) foundation.Dictionary
	PixelFormatType() uint
	MattingImage() corevideo.PixelBufferRef
}

// An auxiliary image used to separate foreground from background with high resolution. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avportraiteffectsmatte?language=objc
type PortraitEffectsMatte struct {
	objc.Object
}

func PortraitEffectsMatteFrom(ptr unsafe.Pointer) PortraitEffectsMatte {
	return PortraitEffectsMatte{
		Object: objc.ObjectFrom(ptr),
	}
}

func (p_ PortraitEffectsMatte) PortraitEffectsMatteByReplacingPortraitEffectsMatteWithPixelBufferError(pixelBuffer corevideo.PixelBufferRef, outError foundation.IError) PortraitEffectsMatte {
	rv := objc.Call[PortraitEffectsMatte](p_, objc.Sel("portraitEffectsMatteByReplacingPortraitEffectsMatteWithPixelBuffer:error:"), pixelBuffer, objc.Ptr(outError))
	return rv
}

// Returns a portrait effects matte by wrapping the replacement pixel buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avportraiteffectsmatte/2976124-portraiteffectsmattebyreplacingp?language=objc
func PortraitEffectsMatte_PortraitEffectsMatteByReplacingPortraitEffectsMatteWithPixelBufferError(pixelBuffer corevideo.PixelBufferRef, outError foundation.IError) PortraitEffectsMatte {
	instance := PortraitEffectsMatteClass.Alloc().PortraitEffectsMatteByReplacingPortraitEffectsMatteWithPixelBufferError(pixelBuffer, outError)
	instance.Autorelease()
	return instance
}

func (p_ PortraitEffectsMatte) PortraitEffectsMatteByApplyingExifOrientation(exifOrientation imageio.ImagePropertyOrientation) PortraitEffectsMatte {
	rv := objc.Call[PortraitEffectsMatte](p_, objc.Sel("portraitEffectsMatteByApplyingExifOrientation:"), exifOrientation)
	return rv
}

// Returns a derivative portrait effects matte after applying the specified EXIF orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avportraiteffectsmatte/2976123-portraiteffectsmattebyapplyingex?language=objc
func PortraitEffectsMatte_PortraitEffectsMatteByApplyingExifOrientation(exifOrientation imageio.ImagePropertyOrientation) PortraitEffectsMatte {
	instance := PortraitEffectsMatteClass.Alloc().PortraitEffectsMatteByApplyingExifOrientation(exifOrientation)
	instance.Autorelease()
	return instance
}

func (pc _PortraitEffectsMatteClass) PortraitEffectsMatteFromDictionaryRepresentationError(imageSourceAuxDataInfoDictionary foundation.Dictionary, outError foundation.IError) PortraitEffectsMatte {
	rv := objc.Call[PortraitEffectsMatte](pc, objc.Sel("portraitEffectsMatteFromDictionaryRepresentation:error:"), imageSourceAuxDataInfoDictionary, objc.Ptr(outError))
	return rv
}

// Initializes a portrait effects matte instance from auxiliary image information in an image file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avportraiteffectsmatte/2976125-portraiteffectsmattefromdictiona?language=objc
func PortraitEffectsMatte_PortraitEffectsMatteFromDictionaryRepresentationError(imageSourceAuxDataInfoDictionary foundation.Dictionary, outError foundation.IError) PortraitEffectsMatte {
	return PortraitEffectsMatteClass.PortraitEffectsMatteFromDictionaryRepresentationError(imageSourceAuxDataInfoDictionary, outError)
}

func (pc _PortraitEffectsMatteClass) Alloc() PortraitEffectsMatte {
	rv := objc.Call[PortraitEffectsMatte](pc, objc.Sel("alloc"))
	return rv
}

func PortraitEffectsMatte_Alloc() PortraitEffectsMatte {
	return PortraitEffectsMatteClass.Alloc()
}

func (pc _PortraitEffectsMatteClass) New() PortraitEffectsMatte {
	rv := objc.Call[PortraitEffectsMatte](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPortraitEffectsMatte() PortraitEffectsMatte {
	return PortraitEffectsMatteClass.New()
}

func (p_ PortraitEffectsMatte) Init() PortraitEffectsMatte {
	rv := objc.Call[PortraitEffectsMatte](p_, objc.Sel("init"))
	return rv
}

// A dictionary of primitive map information used for writing an image file with a portrait effects matte. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avportraiteffectsmatte/2976120-dictionaryrepresentationforauxil?language=objc
func (p_ PortraitEffectsMatte) DictionaryRepresentationForAuxiliaryDataType(outAuxDataType string) foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](p_, objc.Sel("dictionaryRepresentationForAuxiliaryDataType:"), outAuxDataType)
	return rv
}

// The pixel format type of this portrait effects matte's internal image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avportraiteffectsmatte/2976122-pixelformattype?language=objc
func (p_ PortraitEffectsMatte) PixelFormatType() uint {
	rv := objc.Call[uint](p_, objc.Sel("pixelFormatType"))
	return rv
}

// The portrait effects matte's internal image, formatted as a pixel buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avportraiteffectsmatte/2976121-mattingimage?language=objc
func (p_ PortraitEffectsMatte) MattingImage() corevideo.PixelBufferRef {
	rv := objc.Call[corevideo.PixelBufferRef](p_, objc.Sel("mattingImage"))
	return rv
}
