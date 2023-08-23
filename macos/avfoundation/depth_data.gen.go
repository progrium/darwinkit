// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/corevideo"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/imageio"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DepthData] class.
var DepthDataClass = _DepthDataClass{objc.GetClass("AVDepthData")}

type _DepthDataClass struct {
	objc.Class
}

// An interface definition for the [DepthData] class.
type IDepthData interface {
	objc.IObject
	DictionaryRepresentationForAuxiliaryDataType(outAuxDataType string) foundation.Dictionary
	AvailableDepthDataTypes() []foundation.Number
	IsDepthDataFiltered() bool
	DepthDataAccuracy() DepthDataAccuracy
	DepthDataQuality() DepthDataQuality
	DepthDataType() uint
	DepthDataMap() corevideo.PixelBufferRef
	CameraCalibrationData() CameraCalibrationData
}

// A container for per-pixel distance or disparity information captured by compatible camera devices. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdepthdata?language=objc
type DepthData struct {
	objc.Object
}

func DepthDataFrom(ptr unsafe.Pointer) DepthData {
	return DepthData{
		Object: objc.ObjectFrom(ptr),
	}
}

func (d_ DepthData) DepthDataByApplyingExifOrientation(exifOrientation imageio.ImagePropertyOrientation) DepthData {
	rv := objc.Call[DepthData](d_, objc.Sel("depthDataByApplyingExifOrientation:"), exifOrientation)
	return rv
}

// Returns a derivative depth data object by mirroring or rotating it to the specified orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdepthdata/2881225-depthdatabyapplyingexiforientati?language=objc
func DepthData_DepthDataByApplyingExifOrientation(exifOrientation imageio.ImagePropertyOrientation) DepthData {
	instance := DepthDataClass.Alloc().DepthDataByApplyingExifOrientation(exifOrientation)
	instance.Autorelease()
	return instance
}

func (dc _DepthDataClass) DepthDataFromDictionaryRepresentationError(imageSourceAuxDataInfoDictionary foundation.Dictionary, outError foundation.IError) DepthData {
	rv := objc.Call[DepthData](dc, objc.Sel("depthDataFromDictionaryRepresentation:error:"), imageSourceAuxDataInfoDictionary, objc.Ptr(outError))
	return rv
}

// Creates a depth data object from depth information such as that found in an image file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdepthdata/2881221-depthdatafromdictionaryrepresent?language=objc
func DepthData_DepthDataFromDictionaryRepresentationError(imageSourceAuxDataInfoDictionary foundation.Dictionary, outError foundation.IError) DepthData {
	return DepthDataClass.DepthDataFromDictionaryRepresentationError(imageSourceAuxDataInfoDictionary, outError)
}

func (d_ DepthData) DepthDataByConvertingToDepthDataType(depthDataType uint) DepthData {
	rv := objc.Call[DepthData](d_, objc.Sel("depthDataByConvertingToDepthDataType:"), depthDataType)
	return rv
}

// Returns a derivative depth data object by converting the depth data map to the specified data type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdepthdata/2881222-depthdatabyconvertingtodepthdata?language=objc
func DepthData_DepthDataByConvertingToDepthDataType(depthDataType uint) DepthData {
	instance := DepthDataClass.Alloc().DepthDataByConvertingToDepthDataType(depthDataType)
	instance.Autorelease()
	return instance
}

func (d_ DepthData) DepthDataByReplacingDepthDataMapWithPixelBufferError(pixelBuffer corevideo.PixelBufferRef, outError foundation.IError) DepthData {
	rv := objc.Call[DepthData](d_, objc.Sel("depthDataByReplacingDepthDataMapWithPixelBuffer:error:"), pixelBuffer, objc.Ptr(outError))
	return rv
}

// Returns a derivative depth data object by replacing the depth data map. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdepthdata/2881231-depthdatabyreplacingdepthdatamap?language=objc
func DepthData_DepthDataByReplacingDepthDataMapWithPixelBufferError(pixelBuffer corevideo.PixelBufferRef, outError foundation.IError) DepthData {
	instance := DepthDataClass.Alloc().DepthDataByReplacingDepthDataMapWithPixelBufferError(pixelBuffer, outError)
	instance.Autorelease()
	return instance
}

func (dc _DepthDataClass) Alloc() DepthData {
	rv := objc.Call[DepthData](dc, objc.Sel("alloc"))
	return rv
}

func DepthData_Alloc() DepthData {
	return DepthDataClass.Alloc()
}

func (dc _DepthDataClass) New() DepthData {
	rv := objc.Call[DepthData](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDepthData() DepthData {
	return DepthDataClass.New()
}

func (d_ DepthData) Init() DepthData {
	rv := objc.Call[DepthData](d_, objc.Sel("init"))
	return rv
}

// Returns a dictionary representation of the depth data suitable for writing into an image file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdepthdata/2873883-dictionaryrepresentationforauxil?language=objc
func (d_ DepthData) DictionaryRepresentationForAuxiliaryDataType(outAuxDataType string) foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](d_, objc.Sel("dictionaryRepresentationForAuxiliaryDataType:"), outAuxDataType)
	return rv
}

// The list of depth data formats to which this depth data can be converted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdepthdata/2881233-availabledepthdatatypes?language=objc
func (d_ DepthData) AvailableDepthDataTypes() []foundation.Number {
	rv := objc.Call[[]foundation.Number](d_, objc.Sel("availableDepthDataTypes"))
	return rv
}

// A Boolean value indicating whether the depth map contains temporally smoothed data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdepthdata/2881224-depthdatafiltered?language=objc
func (d_ DepthData) IsDepthDataFiltered() bool {
	rv := objc.Call[bool](d_, objc.Sel("isDepthDataFiltered"))
	return rv
}

// The general accuracy of depth data map values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdepthdata/2881229-depthdataaccuracy?language=objc
func (d_ DepthData) DepthDataAccuracy() DepthDataAccuracy {
	rv := objc.Call[DepthDataAccuracy](d_, objc.Sel("depthDataAccuracy"))
	return rv
}

// The overall quality of the depth map. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdepthdata/2919804-depthdataquality?language=objc
func (d_ DepthData) DepthDataQuality() DepthDataQuality {
	rv := objc.Call[DepthDataQuality](d_, objc.Sel("depthDataQuality"))
	return rv
}

// The pixel format of the depth data map. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdepthdata/2881228-depthdatatype?language=objc
func (d_ DepthData) DepthDataType() uint {
	rv := objc.Call[uint](d_, objc.Sel("depthDataType"))
	return rv
}

// A pixel buffer containing the depth data's per-pixel depth or disparity data map. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdepthdata/2881227-depthdatamap?language=objc
func (d_ DepthData) DepthDataMap() corevideo.PixelBufferRef {
	rv := objc.Call[corevideo.PixelBufferRef](d_, objc.Sel("depthDataMap"))
	return rv
}

// The imaging parameters with which this depth data was captured. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdepthdata/2881230-cameracalibrationdata?language=objc
func (d_ DepthData) CameraCalibrationData() CameraCalibrationData {
	rv := objc.Call[CameraCalibrationData](d_, objc.Sel("cameraCalibrationData"))
	return rv
}
