// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [GeometryUtils] class.
var GeometryUtilsClass = _GeometryUtilsClass{objc.GetClass("VNGeometryUtils")}

type _GeometryUtilsClass struct {
	objc.Class
}

// An interface definition for the [GeometryUtils] class.
type IGeometryUtils interface {
	objc.IObject
}

// Utility methods to determine the geometries of various Vision types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngeometryutils?language=objc
type GeometryUtils struct {
	objc.Object
}

func GeometryUtilsFrom(ptr unsafe.Pointer) GeometryUtils {
	return GeometryUtils{
		Object: objc.ObjectFrom(ptr),
	}
}

func (gc _GeometryUtilsClass) Alloc() GeometryUtils {
	rv := objc.Call[GeometryUtils](gc, objc.Sel("alloc"))
	return rv
}

func GeometryUtils_Alloc() GeometryUtils {
	return GeometryUtilsClass.Alloc()
}

func (gc _GeometryUtilsClass) New() GeometryUtils {
	rv := objc.Call[GeometryUtils](gc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewGeometryUtils() GeometryUtils {
	return GeometryUtilsClass.New()
}

func (g_ GeometryUtils) Init() GeometryUtils {
	rv := objc.Call[GeometryUtils](g_, objc.Sel("init"))
	return rv
}

// Calculates the area for the specified contour. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngeometryutils/3548357-calculatearea?language=objc
func (gc _GeometryUtilsClass) CalculateAreaForContourOrientedAreaError(area *float64, contour IContour, orientedArea bool, error foundation.IError) bool {
	rv := objc.Call[bool](gc, objc.Sel("calculateArea:forContour:orientedArea:error:"), area, objc.Ptr(contour), orientedArea, objc.Ptr(error))
	return rv
}

// Calculates the area for the specified contour. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngeometryutils/3548357-calculatearea?language=objc
func GeometryUtils_CalculateAreaForContourOrientedAreaError(area *float64, contour IContour, orientedArea bool, error foundation.IError) bool {
	return GeometryUtilsClass.CalculateAreaForContourOrientedAreaError(area, contour, orientedArea, error)
}

// Calculates a bounding circle for the specified array of points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngeometryutils/3548355-boundingcircleforpoints?language=objc
func (gc _GeometryUtilsClass) BoundingCircleForPointsError(points []IPoint, error foundation.IError) Circle {
	rv := objc.Call[Circle](gc, objc.Sel("boundingCircleForPoints:error:"), points, objc.Ptr(error))
	return rv
}

// Calculates a bounding circle for the specified array of points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngeometryutils/3548355-boundingcircleforpoints?language=objc
func GeometryUtils_BoundingCircleForPointsError(points []IPoint, error foundation.IError) Circle {
	return GeometryUtilsClass.BoundingCircleForPointsError(points, error)
}

// Calculates a bounding circle for the specified points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngeometryutils/3548356-boundingcircleforsimdpoints?language=objc
func (gc _GeometryUtilsClass) BoundingCircleForSIMDPointsPointCountError(points objc.IObject, pointCount int, error foundation.IError) Circle {
	rv := objc.Call[Circle](gc, objc.Sel("boundingCircleForSIMDPoints:pointCount:error:"), objc.Ptr(points), pointCount, objc.Ptr(error))
	return rv
}

// Calculates a bounding circle for the specified points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngeometryutils/3548356-boundingcircleforsimdpoints?language=objc
func GeometryUtils_BoundingCircleForSIMDPointsPointCountError(points objc.IObject, pointCount int, error foundation.IError) Circle {
	return GeometryUtilsClass.BoundingCircleForSIMDPointsPointCountError(points, pointCount, error)
}

// Calculates a bounding circle for the specified contour object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngeometryutils/3548354-boundingcircleforcontour?language=objc
func (gc _GeometryUtilsClass) BoundingCircleForContourError(contour IContour, error foundation.IError) Circle {
	rv := objc.Call[Circle](gc, objc.Sel("boundingCircleForContour:error:"), objc.Ptr(contour), objc.Ptr(error))
	return rv
}

// Calculates a bounding circle for the specified contour object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngeometryutils/3548354-boundingcircleforcontour?language=objc
func GeometryUtils_BoundingCircleForContourError(contour IContour, error foundation.IError) Circle {
	return GeometryUtilsClass.BoundingCircleForContourError(contour, error)
}

// Calculates the perimeter of a closed contour. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngeometryutils/3548358-calculateperimeter?language=objc
func (gc _GeometryUtilsClass) CalculatePerimeterForContourError(perimeter *float64, contour IContour, error foundation.IError) bool {
	rv := objc.Call[bool](gc, objc.Sel("calculatePerimeter:forContour:error:"), perimeter, objc.Ptr(contour), objc.Ptr(error))
	return rv
}

// Calculates the perimeter of a closed contour. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngeometryutils/3548358-calculateperimeter?language=objc
func GeometryUtils_CalculatePerimeterForContourError(perimeter *float64, contour IContour, error foundation.IError) bool {
	return GeometryUtilsClass.CalculatePerimeterForContourError(perimeter, contour, error)
}
