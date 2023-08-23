// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FaceLandmarkRegion2D] class.
var FaceLandmarkRegion2DClass = _FaceLandmarkRegion2DClass{objc.GetClass("VNFaceLandmarkRegion2D")}

type _FaceLandmarkRegion2DClass struct {
	objc.Class
}

// An interface definition for the [FaceLandmarkRegion2D] class.
type IFaceLandmarkRegion2D interface {
	IFaceLandmarkRegion
	PointsInImageOfSize(imageSize coregraphics.Size) *coregraphics.Point
	PrecisionEstimatesPerPoint() []foundation.Number
	NormalizedPoints() *coregraphics.Point
}

// 2D geometry information for a specific facial feature. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfacelandmarkregion2d?language=objc
type FaceLandmarkRegion2D struct {
	FaceLandmarkRegion
}

func FaceLandmarkRegion2DFrom(ptr unsafe.Pointer) FaceLandmarkRegion2D {
	return FaceLandmarkRegion2D{
		FaceLandmarkRegion: FaceLandmarkRegionFrom(ptr),
	}
}

func (fc _FaceLandmarkRegion2DClass) Alloc() FaceLandmarkRegion2D {
	rv := objc.Call[FaceLandmarkRegion2D](fc, objc.Sel("alloc"))
	return rv
}

func FaceLandmarkRegion2D_Alloc() FaceLandmarkRegion2D {
	return FaceLandmarkRegion2DClass.Alloc()
}

func (fc _FaceLandmarkRegion2DClass) New() FaceLandmarkRegion2D {
	rv := objc.Call[FaceLandmarkRegion2D](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFaceLandmarkRegion2D() FaceLandmarkRegion2D {
	return FaceLandmarkRegion2DClass.New()
}

func (f_ FaceLandmarkRegion2D) Init() FaceLandmarkRegion2D {
	rv := objc.Call[FaceLandmarkRegion2D](f_, objc.Sel("init"))
	return rv
}

// A buffer in memory containing landmark points in the coordinate space of the specified image size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfacelandmarkregion2d/2923491-pointsinimageofsize?language=objc
func (f_ FaceLandmarkRegion2D) PointsInImageOfSize(imageSize coregraphics.Size) *coregraphics.Point {
	rv := objc.Call[*coregraphics.Point](f_, objc.Sel("pointsInImageOfSize:"), imageSize)
	return rv
}

// An array of precision estimates for each landmark point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfacelandmarkregion2d/3143672-precisionestimatesperpoint?language=objc
func (f_ FaceLandmarkRegion2D) PrecisionEstimatesPerPoint() []foundation.Number {
	rv := objc.Call[[]foundation.Number](f_, objc.Sel("precisionEstimatesPerPoint"))
	return rv
}

// A buffer in memory containing normalized landmark points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfacelandmarkregion2d/2923490-normalizedpoints?language=objc
func (f_ FaceLandmarkRegion2D) NormalizedPoints() *coregraphics.Point {
	rv := objc.Call[*coregraphics.Point](f_, objc.Sel("normalizedPoints"))
	return rv
}
