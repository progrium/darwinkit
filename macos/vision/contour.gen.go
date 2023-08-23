// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Contour] class.
var ContourClass = _ContourClass{objc.GetClass("VNContour")}

type _ContourClass struct {
	objc.Class
}

// An interface definition for the [Contour] class.
type IContour interface {
	objc.IObject
	PolygonApproximationWithEpsilonError(epsilon float64, error foundation.IError) Contour
	ChildContourAtIndexError(childContourIndex uint, error foundation.IError) Contour
	ChildContours() []Contour
	NormalizedPoints() objc.Object
	PointCount() int
	NormalizedPath() unsafe.Pointer
	IndexPath() foundation.IndexPath
	AspectRatio() float64
	ChildContourCount() int
}

// A class that represents a detected contour in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontour?language=objc
type Contour struct {
	objc.Object
}

func ContourFrom(ptr unsafe.Pointer) Contour {
	return Contour{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _ContourClass) Alloc() Contour {
	rv := objc.Call[Contour](cc, objc.Sel("alloc"))
	return rv
}

func Contour_Alloc() Contour {
	return ContourClass.Alloc()
}

func (cc _ContourClass) New() Contour {
	rv := objc.Call[Contour](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewContour() Contour {
	return ContourClass.New()
}

func (c_ Contour) Init() Contour {
	rv := objc.Call[Contour](c_, objc.Sel("init"))
	return rv
}

// Simplifies the contour to a polygon using a Ramer-Douglas-Peucker algorithm. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontour/3618958-polygonapproximationwithepsilon?language=objc
func (c_ Contour) PolygonApproximationWithEpsilonError(epsilon float64, error foundation.IError) Contour {
	rv := objc.Call[Contour](c_, objc.Sel("polygonApproximationWithEpsilon:error:"), epsilon, objc.Ptr(error))
	return rv
}

// Retrieves the child contour object at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontour/3548321-childcontouratindex?language=objc
func (c_ Contour) ChildContourAtIndexError(childContourIndex uint, error foundation.IError) Contour {
	rv := objc.Call[Contour](c_, objc.Sel("childContourAtIndex:error:"), childContourIndex, objc.Ptr(error))
	return rv
}

// An array of contours that this contour encloses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontour/3548323-childcontours?language=objc
func (c_ Contour) ChildContours() []Contour {
	rv := objc.Call[[]Contour](c_, objc.Sel("childContours"))
	return rv
}

// The contour’s array of points in normalized coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontour/3548326-normalizedpoints?language=objc
func (c_ Contour) NormalizedPoints() objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("normalizedPoints"))
	return rv
}

// The contour’s number of points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontour/3548327-pointcount?language=objc
func (c_ Contour) PointCount() int {
	rv := objc.Call[int](c_, objc.Sel("pointCount"))
	return rv
}

// The contour object as a path in normalized coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontour/3548325-normalizedpath?language=objc
func (c_ Contour) NormalizedPath() unsafe.Pointer {
	rv := objc.Call[unsafe.Pointer](c_, objc.Sel("normalizedPath"))
	return rv
}

// The contour object’s index path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontour/3548324-indexpath?language=objc
func (c_ Contour) IndexPath() foundation.IndexPath {
	rv := objc.Call[foundation.IndexPath](c_, objc.Sel("indexPath"))
	return rv
}

// The aspect ratio of the contour. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontour/3600614-aspectratio?language=objc
func (c_ Contour) AspectRatio() float64 {
	rv := objc.Call[float64](c_, objc.Sel("aspectRatio"))
	return rv
}

// The total number of detected child contours. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontour/3548322-childcontourcount?language=objc
func (c_ Contour) ChildContourCount() int {
	rv := objc.Call[int](c_, objc.Sel("childContourCount"))
	return rv
}
