// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FaceLandmarks] class.
var FaceLandmarksClass = _FaceLandmarksClass{objc.GetClass("VNFaceLandmarks")}

type _FaceLandmarksClass struct {
	objc.Class
}

// An interface definition for the [FaceLandmarks] class.
type IFaceLandmarks interface {
	objc.IObject
	Confidence() Confidence
}

// The abstract superclass for containers of face landmark information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfacelandmarks?language=objc
type FaceLandmarks struct {
	objc.Object
}

func FaceLandmarksFrom(ptr unsafe.Pointer) FaceLandmarks {
	return FaceLandmarks{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FaceLandmarksClass) Alloc() FaceLandmarks {
	rv := objc.Call[FaceLandmarks](fc, objc.Sel("alloc"))
	return rv
}

func FaceLandmarks_Alloc() FaceLandmarks {
	return FaceLandmarksClass.Alloc()
}

func (fc _FaceLandmarksClass) New() FaceLandmarks {
	rv := objc.Call[FaceLandmarks](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFaceLandmarks() FaceLandmarks {
	return FaceLandmarksClass.New()
}

func (f_ FaceLandmarks) Init() FaceLandmarks {
	rv := objc.Call[FaceLandmarks](f_, objc.Sel("init"))
	return rv
}

// A confidence estimate for the detected landmarks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfacelandmarks/2909054-confidence?language=objc
func (f_ FaceLandmarks) Confidence() Confidence {
	rv := objc.Call[Confidence](f_, objc.Sel("confidence"))
	return rv
}
