// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ContoursObservation] class.
var ContoursObservationClass = _ContoursObservationClass{objc.GetClass("VNContoursObservation")}

type _ContoursObservationClass struct {
	objc.Class
}

// An interface definition for the [ContoursObservation] class.
type IContoursObservation interface {
	IObservation
	ContourAtIndexError(contourIndex int, error foundation.IError) Contour
	ContourAtIndexPathError(indexPath foundation.IIndexPath, error foundation.IError) Contour
	ContourCount() int
	TopLevelContours() []Contour
	NormalizedPath() unsafe.Pointer
	TopLevelContourCount() int
}

// An object that represents the detected contours in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontoursobservation?language=objc
type ContoursObservation struct {
	Observation
}

func ContoursObservationFrom(ptr unsafe.Pointer) ContoursObservation {
	return ContoursObservation{
		Observation: ObservationFrom(ptr),
	}
}

func (cc _ContoursObservationClass) Alloc() ContoursObservation {
	rv := objc.Call[ContoursObservation](cc, objc.Sel("alloc"))
	return rv
}

func ContoursObservation_Alloc() ContoursObservation {
	return ContoursObservationClass.Alloc()
}

func (cc _ContoursObservationClass) New() ContoursObservation {
	rv := objc.Call[ContoursObservation](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewContoursObservation() ContoursObservation {
	return ContoursObservationClass.New()
}

func (c_ ContoursObservation) Init() ContoursObservation {
	rv := objc.Call[ContoursObservation](c_, objc.Sel("init"))
	return rv
}

// Retrieves the contour object at the specified index, irrespective of hierarchy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontoursobservation/3548360-contouratindex?language=objc
func (c_ ContoursObservation) ContourAtIndexError(contourIndex int, error foundation.IError) Contour {
	rv := objc.Call[Contour](c_, objc.Sel("contourAtIndex:error:"), contourIndex, objc.Ptr(error))
	return rv
}

// Retrieves the contour object at the specified index path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontoursobservation/3548361-contouratindexpath?language=objc
func (c_ ContoursObservation) ContourAtIndexPathError(indexPath foundation.IIndexPath, error foundation.IError) Contour {
	rv := objc.Call[Contour](c_, objc.Sel("contourAtIndexPath:error:"), objc.Ptr(indexPath), objc.Ptr(error))
	return rv
}

// The total number of detected contours. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontoursobservation/3548362-contourcount?language=objc
func (c_ ContoursObservation) ContourCount() int {
	rv := objc.Call[int](c_, objc.Sel("contourCount"))
	return rv
}

// An array of contours that don’t have another contour enclosing them. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontoursobservation/3548364-toplevelcontours?language=objc
func (c_ ContoursObservation) TopLevelContours() []Contour {
	rv := objc.Call[[]Contour](c_, objc.Sel("topLevelContours"))
	return rv
}

// The detected contours as a path object in normalized coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontoursobservation/3548363-normalizedpath?language=objc
func (c_ ContoursObservation) NormalizedPath() unsafe.Pointer {
	rv := objc.Call[unsafe.Pointer](c_, objc.Sel("normalizedPath"))
	return rv
}

// The total number of detected top-level contours. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontoursobservation/3675675-toplevelcontourcount?language=objc
func (c_ ContoursObservation) TopLevelContourCount() int {
	rv := objc.Call[int](c_, objc.Sel("topLevelContourCount"))
	return rv
}
