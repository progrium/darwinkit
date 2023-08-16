// AUTO-GENERATED CODE, DO NOT MODIFY

package corelocation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Heading] class.
var HeadingClass = _HeadingClass{objc.GetClass("CLHeading")}

type _HeadingClass struct {
	objc.Class
}

// An interface definition for the [Heading] class.
type IHeading interface {
	objc.IObject
	MagneticHeading() LocationDirection
	X() HeadingComponentValue
	SetX(value HeadingComponentValue)
	Y() HeadingComponentValue
	SetY(value HeadingComponentValue)
	TrueHeading() LocationDirection
	HeadingAccuracy() LocationDirection
	Timestamp() foundation.Date
	Z() HeadingComponentValue
	SetZ(value HeadingComponentValue)
}

// The azimuth (orientation) of the user’s device, relative to true or magnetic north. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clheading?language=objc
type Heading struct {
	objc.Object
}

func HeadingFrom(ptr unsafe.Pointer) Heading {
	return Heading{
		Object: objc.ObjectFrom(ptr),
	}
}

func (hc _HeadingClass) Alloc() Heading {
	rv := objc.Call[Heading](hc, objc.Sel("alloc"))
	return rv
}

func Heading_Alloc() Heading {
	return HeadingClass.Alloc()
}

func (hc _HeadingClass) New() Heading {
	rv := objc.Call[Heading](hc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewHeading() Heading {
	return HeadingClass.New()
}

func (h_ Heading) Init() Heading {
	rv := objc.Call[Heading](h_, objc.Sel("init"))
	return rv
}

// The heading (measured in degrees) relative to magnetic north. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clheading/1423763-magneticheading?language=objc
func (h_ Heading) MagneticHeading() LocationDirection {
	rv := objc.Call[LocationDirection](h_, objc.Sel("magneticHeading"))
	return rv
}

// The geomagnetic data (measured in microteslas) for the x-axis. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clheading/1423685-x?language=objc
func (h_ Heading) X() HeadingComponentValue {
	rv := objc.Call[HeadingComponentValue](h_, objc.Sel("x"))
	return rv
}

// The geomagnetic data (measured in microteslas) for the x-axis. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clheading/1423685-x?language=objc
func (h_ Heading) SetX(value HeadingComponentValue) {
	objc.Call[objc.Void](h_, objc.Sel("setX:"), value)
}

// The geomagnetic data (measured in microteslas) for the y-axis. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clheading/1423617-y?language=objc
func (h_ Heading) Y() HeadingComponentValue {
	rv := objc.Call[HeadingComponentValue](h_, objc.Sel("y"))
	return rv
}

// The geomagnetic data (measured in microteslas) for the y-axis. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clheading/1423617-y?language=objc
func (h_ Heading) SetY(value HeadingComponentValue) {
	objc.Call[objc.Void](h_, objc.Sel("setY:"), value)
}

// The heading (measured in degrees) relative to true north. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clheading/1423568-trueheading?language=objc
func (h_ Heading) TrueHeading() LocationDirection {
	rv := objc.Call[LocationDirection](h_, objc.Sel("trueHeading"))
	return rv
}

// The maximum deviation (measured in degrees) between the reported heading and the true geomagnetic heading. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clheading/1423705-headingaccuracy?language=objc
func (h_ Heading) HeadingAccuracy() LocationDirection {
	rv := objc.Call[LocationDirection](h_, objc.Sel("headingAccuracy"))
	return rv
}

// The time at which this heading was determined. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clheading/1423525-timestamp?language=objc
func (h_ Heading) Timestamp() foundation.Date {
	rv := objc.Call[foundation.Date](h_, objc.Sel("timestamp"))
	return rv
}

// The geomagnetic data (measured in microteslas) for the z-axis. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clheading/1423609-z?language=objc
func (h_ Heading) Z() HeadingComponentValue {
	rv := objc.Call[HeadingComponentValue](h_, objc.Sel("z"))
	return rv
}

// The geomagnetic data (measured in microteslas) for the z-axis. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clheading/1423609-z?language=objc
func (h_ Heading) SetZ(value HeadingComponentValue) {
	objc.Call[objc.Void](h_, objc.Sel("setZ:"), value)
}
