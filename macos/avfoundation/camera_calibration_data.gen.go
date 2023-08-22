// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/kernel"
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CameraCalibrationData] class.
var CameraCalibrationDataClass = _CameraCalibrationDataClass{objc.GetClass("AVCameraCalibrationData")}

type _CameraCalibrationDataClass struct {
	objc.Class
}

// An interface definition for the [CameraCalibrationData] class.
type ICameraCalibrationData interface {
	objc.IObject
	LensDistortionCenter() coregraphics.Point
	InverseLensDistortionLookupTable() []byte
	LensDistortionLookupTable() []byte
	IntrinsicMatrixReferenceDimensions() coregraphics.Size
	IntrinsicMatrix() kernel.Matrix_float3x3
	ExtrinsicMatrix() kernel.Matrix_float4x3
	PixelSize() float64
}

// Information about the camera characteristics used to capture images and depth data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcameracalibrationdata?language=objc
type CameraCalibrationData struct {
	objc.Object
}

func CameraCalibrationDataFrom(ptr unsafe.Pointer) CameraCalibrationData {
	return CameraCalibrationData{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CameraCalibrationDataClass) Alloc() CameraCalibrationData {
	rv := objc.Call[CameraCalibrationData](cc, objc.Sel("alloc"))
	return rv
}

func CameraCalibrationData_Alloc() CameraCalibrationData {
	return CameraCalibrationDataClass.Alloc()
}

func (cc _CameraCalibrationDataClass) New() CameraCalibrationData {
	rv := objc.Call[CameraCalibrationData](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCameraCalibrationData() CameraCalibrationData {
	return CameraCalibrationDataClass.New()
}

func (c_ CameraCalibrationData) Init() CameraCalibrationData {
	rv := objc.Call[CameraCalibrationData](c_, objc.Sel("init"))
	return rv
}

// The offset of the distortion center of the camera lens from the top-left corner of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcameracalibrationdata/2881131-lensdistortioncenter?language=objc
func (c_ CameraCalibrationData) LensDistortionCenter() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](c_, objc.Sel("lensDistortionCenter"))
	return rv
}

// A map of floating-point values describing radial distortions for use in reapplying camera geometry to a rectified image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcameracalibrationdata/2881132-inverselensdistortionlookuptable?language=objc
func (c_ CameraCalibrationData) InverseLensDistortionLookupTable() []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("inverseLensDistortionLookupTable"))
	return rv
}

// A map of floating-point values describing radial distortions imparted by the camera lens, for use in rectifying camera images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcameracalibrationdata/2881129-lensdistortionlookuptable?language=objc
func (c_ CameraCalibrationData) LensDistortionLookupTable() []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("lensDistortionLookupTable"))
	return rv
}

// The image dimensions to which the camera’s intrinsic matrix values are relative. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcameracalibrationdata/2881134-intrinsicmatrixreferencedimensio?language=objc
func (c_ CameraCalibrationData) IntrinsicMatrixReferenceDimensions() coregraphics.Size {
	rv := objc.Call[coregraphics.Size](c_, objc.Sel("intrinsicMatrixReferenceDimensions"))
	return rv
}

// A matrix that relates a camera’s internal properties to an ideal pinhole-camera model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcameracalibrationdata/2881135-intrinsicmatrix?language=objc
func (c_ CameraCalibrationData) IntrinsicMatrix() kernel.Matrix_float3x3 {
	rv := objc.Call[kernel.Matrix_float3x3](c_, objc.Sel("intrinsicMatrix"))
	return rv
}

// A matrix relating a camera’s position and orientation to a world or scene coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcameracalibrationdata/2881130-extrinsicmatrix?language=objc
func (c_ CameraCalibrationData) ExtrinsicMatrix() kernel.Matrix_float4x3 {
	rv := objc.Call[kernel.Matrix_float4x3](c_, objc.Sel("extrinsicMatrix"))
	return rv
}

// The size, in millimeters, of one image pixel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcameracalibrationdata/2881128-pixelsize?language=objc
func (c_ CameraCalibrationData) PixelSize() float64 {
	rv := objc.Call[float64](c_, objc.Sel("pixelSize"))
	return rv
}
