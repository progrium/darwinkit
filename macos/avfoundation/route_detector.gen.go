// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RouteDetector] class.
var RouteDetectorClass = _RouteDetectorClass{objc.GetClass("AVRouteDetector")}

type _RouteDetectorClass struct {
	objc.Class
}

// An interface definition for the [RouteDetector] class.
type IRouteDetector interface {
	objc.IObject
	MultipleRoutesDetected() bool
	IsRouteDetectionEnabled() bool
	SetRouteDetectionEnabled(value bool)
}

// An object that detects available media playback routes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avroutedetector?language=objc
type RouteDetector struct {
	objc.Object
}

func RouteDetectorFrom(ptr unsafe.Pointer) RouteDetector {
	return RouteDetector{
		Object: objc.ObjectFrom(ptr),
	}
}

func (rc _RouteDetectorClass) Alloc() RouteDetector {
	rv := objc.Call[RouteDetector](rc, objc.Sel("alloc"))
	return rv
}

func RouteDetector_Alloc() RouteDetector {
	return RouteDetectorClass.Alloc()
}

func (rc _RouteDetectorClass) New() RouteDetector {
	rv := objc.Call[RouteDetector](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRouteDetector() RouteDetector {
	return RouteDetectorClass.New()
}

func (r_ RouteDetector) Init() RouteDetector {
	rv := objc.Call[RouteDetector](r_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether the object detects more than one playback route. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avroutedetector/2915760-multipleroutesdetected?language=objc
func (r_ RouteDetector) MultipleRoutesDetected() bool {
	rv := objc.Call[bool](r_, objc.Sel("multipleRoutesDetected"))
	return rv
}

// A Boolean value that indicates whether route detection is in an enabled state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avroutedetector/2915762-routedetectionenabled?language=objc
func (r_ RouteDetector) IsRouteDetectionEnabled() bool {
	rv := objc.Call[bool](r_, objc.Sel("isRouteDetectionEnabled"))
	return rv
}

// A Boolean value that indicates whether route detection is in an enabled state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avroutedetector/2915762-routedetectionenabled?language=objc
func (r_ RouteDetector) SetRouteDetectionEnabled(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setRouteDetectionEnabled:"), value)
}
