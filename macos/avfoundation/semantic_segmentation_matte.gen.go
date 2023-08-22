// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/corefoundation"
	"github.com/progrium/macdriver/macos/corevideo"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/imageio"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SemanticSegmentationMatte] class.
var SemanticSegmentationMatteClass = _SemanticSegmentationMatteClass{objc.GetClass("AVSemanticSegmentationMatte")}

type _SemanticSegmentationMatteClass struct {
	objc.Class
}

// An interface definition for the [SemanticSegmentationMatte] class.
type ISemanticSegmentationMatte interface {
	objc.IObject
	DictionaryRepresentationForAuxiliaryDataType(outAuxDataType string) foundation.Dictionary
	MatteType() SemanticSegmentationMatteType
	PixelFormatType() uint
	MattingImage() corevideo.PixelBufferRef
}

// An object that wraps a matting image for a particular semantic segmentation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsemanticsegmentationmatte?language=objc
type SemanticSegmentationMatte struct {
	objc.Object
}

func SemanticSegmentationMatteFrom(ptr unsafe.Pointer) SemanticSegmentationMatte {
	return SemanticSegmentationMatte{
		Object: objc.ObjectFrom(ptr),
	}
}

func (s_ SemanticSegmentationMatte) SemanticSegmentationMatteByApplyingExifOrientation(exifOrientation imageio.ImagePropertyOrientation) SemanticSegmentationMatte {
	rv := objc.Call[SemanticSegmentationMatte](s_, objc.Sel("semanticSegmentationMatteByApplyingExifOrientation:"), exifOrientation)
	return rv
}

// Returns a new semantic segmentation matte instance with the specified Exif orientation applied. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsemanticsegmentationmatte/3152122-semanticsegmentationmattebyapply?language=objc
func SemanticSegmentationMatte_SemanticSegmentationMatteByApplyingExifOrientation(exifOrientation imageio.ImagePropertyOrientation) SemanticSegmentationMatte {
	instance := SemanticSegmentationMatteClass.Alloc().SemanticSegmentationMatteByApplyingExifOrientation(exifOrientation)
	instance.Autorelease()
	return instance
}

func (s_ SemanticSegmentationMatte) SemanticSegmentationMatteByReplacingSemanticSegmentationMatteWithPixelBufferError(pixelBuffer corevideo.PixelBufferRef, outError foundation.IError) SemanticSegmentationMatte {
	rv := objc.Call[SemanticSegmentationMatte](s_, objc.Sel("semanticSegmentationMatteByReplacingSemanticSegmentationMatteWithPixelBuffer:error:"), pixelBuffer, objc.Ptr(outError))
	return rv
}

// Returns a semantic segmentation matte instance that wraps the replacement pixel buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsemanticsegmentationmatte/3152123-semanticsegmentationmattebyrepla?language=objc
func SemanticSegmentationMatte_SemanticSegmentationMatteByReplacingSemanticSegmentationMatteWithPixelBufferError(pixelBuffer corevideo.PixelBufferRef, outError foundation.IError) SemanticSegmentationMatte {
	instance := SemanticSegmentationMatteClass.Alloc().SemanticSegmentationMatteByReplacingSemanticSegmentationMatteWithPixelBufferError(pixelBuffer, outError)
	instance.Autorelease()
	return instance
}

func (sc _SemanticSegmentationMatteClass) SemanticSegmentationMatteFromImageSourceAuxiliaryDataTypeDictionaryRepresentationError(imageSourceAuxiliaryDataType corefoundation.StringRef, imageSourceAuxiliaryDataInfoDictionary foundation.Dictionary, outError foundation.IError) SemanticSegmentationMatte {
	rv := objc.Call[SemanticSegmentationMatte](sc, objc.Sel("semanticSegmentationMatteFromImageSourceAuxiliaryDataType:dictionaryRepresentation:error:"), imageSourceAuxiliaryDataType, imageSourceAuxiliaryDataInfoDictionary, objc.Ptr(outError))
	return rv
}

// Returns a new semantic segmentation matte instance from auxiliary image information in an image file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsemanticsegmentationmatte/3152124-semanticsegmentationmattefromima?language=objc
func SemanticSegmentationMatte_SemanticSegmentationMatteFromImageSourceAuxiliaryDataTypeDictionaryRepresentationError(imageSourceAuxiliaryDataType corefoundation.StringRef, imageSourceAuxiliaryDataInfoDictionary foundation.Dictionary, outError foundation.IError) SemanticSegmentationMatte {
	return SemanticSegmentationMatteClass.SemanticSegmentationMatteFromImageSourceAuxiliaryDataTypeDictionaryRepresentationError(imageSourceAuxiliaryDataType, imageSourceAuxiliaryDataInfoDictionary, outError)
}

func (sc _SemanticSegmentationMatteClass) Alloc() SemanticSegmentationMatte {
	rv := objc.Call[SemanticSegmentationMatte](sc, objc.Sel("alloc"))
	return rv
}

func SemanticSegmentationMatte_Alloc() SemanticSegmentationMatte {
	return SemanticSegmentationMatteClass.Alloc()
}

func (sc _SemanticSegmentationMatteClass) New() SemanticSegmentationMatte {
	rv := objc.Call[SemanticSegmentationMatte](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSemanticSegmentationMatte() SemanticSegmentationMatte {
	return SemanticSegmentationMatteClass.New()
}

func (s_ SemanticSegmentationMatte) Init() SemanticSegmentationMatte {
	rv := objc.Call[SemanticSegmentationMatte](s_, objc.Sel("init"))
	return rv
}

// Returns a dictionary of primitive map information to use when writing an image file with a semantic segmentation matte. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsemanticsegmentationmatte/3152118-dictionaryrepresentationforauxil?language=objc
func (s_ SemanticSegmentationMatte) DictionaryRepresentationForAuxiliaryDataType(outAuxDataType string) foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](s_, objc.Sel("dictionaryRepresentationForAuxiliaryDataType:"), outAuxDataType)
	return rv
}

// The semantic segmentation matte image type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsemanticsegmentationmatte/3152119-mattetype?language=objc
func (s_ SemanticSegmentationMatte) MatteType() SemanticSegmentationMatteType {
	rv := objc.Call[SemanticSegmentationMatteType](s_, objc.Sel("matteType"))
	return rv
}

// The pixel format type for this object’s internal matting image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsemanticsegmentationmatte/3152121-pixelformattype?language=objc
func (s_ SemanticSegmentationMatte) PixelFormatType() uint {
	rv := objc.Call[uint](s_, objc.Sel("pixelFormatType"))
	return rv
}

// The semantic segmentation matte’s internal image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsemanticsegmentationmatte/3152120-mattingimage?language=objc
func (s_ SemanticSegmentationMatte) MattingImage() corevideo.PixelBufferRef {
	rv := objc.Call[corevideo.PixelBufferRef](s_, objc.Sel("mattingImage"))
	return rv
}
