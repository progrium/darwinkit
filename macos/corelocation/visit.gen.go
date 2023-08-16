// AUTO-GENERATED CODE, DO NOT MODIFY

package corelocation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Visit] class.
var VisitClass = _VisitClass{objc.GetClass("CLVisit")}

type _VisitClass struct {
	objc.Class
}

// An interface definition for the [Visit] class.
type IVisit interface {
	objc.IObject
	HorizontalAccuracy() LocationAccuracy
	Coordinate() LocationCoordinate2D
	ArrivalDate() foundation.Date
	DepartureDate() foundation.Date
}

// Information about the user's location during a specific period of time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clvisit?language=objc
type Visit struct {
	objc.Object
}

func VisitFrom(ptr unsafe.Pointer) Visit {
	return Visit{
		Object: objc.ObjectFrom(ptr),
	}
}

func (vc _VisitClass) Alloc() Visit {
	rv := objc.Call[Visit](vc, objc.Sel("alloc"))
	return rv
}

func Visit_Alloc() Visit {
	return VisitClass.Alloc()
}

func (vc _VisitClass) New() Visit {
	rv := objc.Call[Visit](vc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewVisit() Visit {
	return VisitClass.New()
}

func (v_ Visit) Init() Visit {
	rv := objc.Call[Visit](v_, objc.Sel("init"))
	return rv
}

// The horizontal accuracy (in meters) of the specified coordinate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clvisit/1614679-horizontalaccuracy?language=objc
func (v_ Visit) HorizontalAccuracy() LocationAccuracy {
	rv := objc.Call[LocationAccuracy](v_, objc.Sel("horizontalAccuracy"))
	return rv
}

// The geographical coordinate information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clvisit/1614677-coordinate?language=objc
func (v_ Visit) Coordinate() LocationCoordinate2D {
	rv := objc.Call[LocationCoordinate2D](v_, objc.Sel("coordinate"))
	return rv
}

// The approximate time at which the user arrived at the specified location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clvisit/1614681-arrivaldate?language=objc
func (v_ Visit) ArrivalDate() foundation.Date {
	rv := objc.Call[foundation.Date](v_, objc.Sel("arrivalDate"))
	return rv
}

// The approximate time at which the user left the specified location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clvisit/1614685-departuredate?language=objc
func (v_ Visit) DepartureDate() foundation.Date {
	rv := objc.Call[foundation.Date](v_, objc.Sel("departureDate"))
	return rv
}
