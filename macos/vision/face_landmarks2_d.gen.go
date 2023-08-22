// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FaceLandmarks2D] class.
var FaceLandmarks2DClass = _FaceLandmarks2DClass{objc.GetClass("VNFaceLandmarks2D")}

type _FaceLandmarks2DClass struct {
	objc.Class
}

// An interface definition for the [FaceLandmarks2D] class.
type IFaceLandmarks2D interface {
	IFaceLandmarks
	RightPupil() FaceLandmarkRegion2D
	Nose() FaceLandmarkRegion2D
	InnerLips() FaceLandmarkRegion2D
	LeftPupil() FaceLandmarkRegion2D
	OuterLips() FaceLandmarkRegion2D
	MedianLine() FaceLandmarkRegion2D
	NoseCrest() FaceLandmarkRegion2D
	RightEyebrow() FaceLandmarkRegion2D
	FaceContour() FaceLandmarkRegion2D
	AllPoints() FaceLandmarkRegion2D
	LeftEyebrow() FaceLandmarkRegion2D
	LeftEye() FaceLandmarkRegion2D
	RightEye() FaceLandmarkRegion2D
}

// A collection of facial features that a request detects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfacelandmarks2d?language=objc
type FaceLandmarks2D struct {
	FaceLandmarks
}

func FaceLandmarks2DFrom(ptr unsafe.Pointer) FaceLandmarks2D {
	return FaceLandmarks2D{
		FaceLandmarks: FaceLandmarksFrom(ptr),
	}
}

func (fc _FaceLandmarks2DClass) Alloc() FaceLandmarks2D {
	rv := objc.Call[FaceLandmarks2D](fc, objc.Sel("alloc"))
	return rv
}

func FaceLandmarks2D_Alloc() FaceLandmarks2D {
	return FaceLandmarks2DClass.Alloc()
}

func (fc _FaceLandmarks2DClass) New() FaceLandmarks2D {
	rv := objc.Call[FaceLandmarks2D](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFaceLandmarks2D() FaceLandmarks2D {
	return FaceLandmarks2DClass.New()
}

func (f_ FaceLandmarks2D) Init() FaceLandmarks2D {
	rv := objc.Call[FaceLandmarks2D](f_, objc.Sel("init"))
	return rv
}

// The region containing the point where the right pupil is located. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfacelandmarks2d/2879441-rightpupil?language=objc
func (f_ FaceLandmarks2D) RightPupil() FaceLandmarkRegion2D {
	rv := objc.Call[FaceLandmarkRegion2D](f_, objc.Sel("rightPupil"))
	return rv
}

// The region containing points that outline the nose. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfacelandmarks2d/2879443-nose?language=objc
func (f_ FaceLandmarks2D) Nose() FaceLandmarkRegion2D {
	rv := objc.Call[FaceLandmarkRegion2D](f_, objc.Sel("nose"))
	return rv
}

// The region containing points that outline the space between the lips. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfacelandmarks2d/2879434-innerlips?language=objc
func (f_ FaceLandmarks2D) InnerLips() FaceLandmarkRegion2D {
	rv := objc.Call[FaceLandmarkRegion2D](f_, objc.Sel("innerLips"))
	return rv
}

// The region containing the point where the left pupil is located. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfacelandmarks2d/2879436-leftpupil?language=objc
func (f_ FaceLandmarks2D) LeftPupil() FaceLandmarkRegion2D {
	rv := objc.Call[FaceLandmarkRegion2D](f_, objc.Sel("leftPupil"))
	return rv
}

// The region containing points that outline the outside of the lips. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfacelandmarks2d/2879440-outerlips?language=objc
func (f_ FaceLandmarks2D) OuterLips() FaceLandmarkRegion2D {
	rv := objc.Call[FaceLandmarkRegion2D](f_, objc.Sel("outerLips"))
	return rv
}

// The region containing points that trace a vertical line down the center of the face. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfacelandmarks2d/2879427-medianline?language=objc
func (f_ FaceLandmarks2D) MedianLine() FaceLandmarkRegion2D {
	rv := objc.Call[FaceLandmarkRegion2D](f_, objc.Sel("medianLine"))
	return rv
}

// The region containing points that trace the center crest of the nose. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfacelandmarks2d/2879431-nosecrest?language=objc
func (f_ FaceLandmarks2D) NoseCrest() FaceLandmarkRegion2D {
	rv := objc.Call[FaceLandmarkRegion2D](f_, objc.Sel("noseCrest"))
	return rv
}

// The region containing points that trace the right eyebrow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfacelandmarks2d/2879432-righteyebrow?language=objc
func (f_ FaceLandmarks2D) RightEyebrow() FaceLandmarkRegion2D {
	rv := objc.Call[FaceLandmarkRegion2D](f_, objc.Sel("rightEyebrow"))
	return rv
}

// The region containing points that trace the face contour from the left cheek, over the chin, to the right cheek. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfacelandmarks2d/2879437-facecontour?language=objc
func (f_ FaceLandmarks2D) FaceContour() FaceLandmarkRegion2D {
	rv := objc.Call[FaceLandmarkRegion2D](f_, objc.Sel("faceContour"))
	return rv
}

// The region containing all face landmark points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfacelandmarks2d/2879430-allpoints?language=objc
func (f_ FaceLandmarks2D) AllPoints() FaceLandmarkRegion2D {
	rv := objc.Call[FaceLandmarkRegion2D](f_, objc.Sel("allPoints"))
	return rv
}

// The region containing points that trace the left eyebrow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfacelandmarks2d/2879438-lefteyebrow?language=objc
func (f_ FaceLandmarks2D) LeftEyebrow() FaceLandmarkRegion2D {
	rv := objc.Call[FaceLandmarkRegion2D](f_, objc.Sel("leftEyebrow"))
	return rv
}

// The region containing points that outline the left eye. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfacelandmarks2d/2879426-lefteye?language=objc
func (f_ FaceLandmarks2D) LeftEye() FaceLandmarkRegion2D {
	rv := objc.Call[FaceLandmarkRegion2D](f_, objc.Sel("leftEye"))
	return rv
}

// The region containing points that outline the right eye. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfacelandmarks2d/2879442-righteye?language=objc
func (f_ FaceLandmarks2D) RightEye() FaceLandmarkRegion2D {
	rv := objc.Call[FaceLandmarkRegion2D](f_, objc.Sel("rightEye"))
	return rv
}
