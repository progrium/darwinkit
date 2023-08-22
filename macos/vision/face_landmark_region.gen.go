// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FaceLandmarkRegion] class.
var FaceLandmarkRegionClass = _FaceLandmarkRegionClass{objc.GetClass("VNFaceLandmarkRegion")}

type _FaceLandmarkRegionClass struct {
	objc.Class
}

// An interface definition for the [FaceLandmarkRegion] class.
type IFaceLandmarkRegion interface {
	objc.IObject
	PointCount() uint
}

// The abstract superclass for information about a specific face landmark. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfacelandmarkregion?language=objc
type FaceLandmarkRegion struct {
	objc.Object
}

func FaceLandmarkRegionFrom(ptr unsafe.Pointer) FaceLandmarkRegion {
	return FaceLandmarkRegion{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FaceLandmarkRegionClass) Alloc() FaceLandmarkRegion {
	rv := objc.Call[FaceLandmarkRegion](fc, objc.Sel("alloc"))
	return rv
}

func FaceLandmarkRegion_Alloc() FaceLandmarkRegion {
	return FaceLandmarkRegionClass.Alloc()
}

func (fc _FaceLandmarkRegionClass) New() FaceLandmarkRegion {
	rv := objc.Call[FaceLandmarkRegion](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFaceLandmarkRegion() FaceLandmarkRegion {
	return FaceLandmarkRegionClass.New()
}

func (f_ FaceLandmarkRegion) Init() FaceLandmarkRegion {
	rv := objc.Call[FaceLandmarkRegion](f_, objc.Sel("init"))
	return rv
}

// The number of points in the face region. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfacelandmarkregion/2909053-pointcount?language=objc
func (f_ FaceLandmarkRegion) PointCount() uint {
	rv := objc.Call[uint](f_, objc.Sel("pointCount"))
	return rv
}
