// AUTO-GENERATED CODE, DO NOT MODIFY

package corelocation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CircularRegion] class.
var CircularRegionClass = _CircularRegionClass{objc.GetClass("CLCircularRegion")}

type _CircularRegionClass struct {
	objc.Class
}

// An interface definition for the [CircularRegion] class.
type ICircularRegion interface {
	IRegion
}

// A circular geographic region that a center point and radius deine. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clcircularregion?language=objc
type CircularRegion struct {
	Region
}

func CircularRegionFrom(ptr unsafe.Pointer) CircularRegion {
	return CircularRegion{
		Region: RegionFrom(ptr),
	}
}

func (cc _CircularRegionClass) Alloc() CircularRegion {
	rv := objc.Call[CircularRegion](cc, objc.Sel("alloc"))
	return rv
}

func CircularRegion_Alloc() CircularRegion {
	return CircularRegionClass.Alloc()
}

func (cc _CircularRegionClass) New() CircularRegion {
	rv := objc.Call[CircularRegion](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCircularRegion() CircularRegion {
	return CircularRegionClass.New()
}

func (c_ CircularRegion) Init() CircularRegion {
	rv := objc.Call[CircularRegion](c_, objc.Sel("init"))
	return rv
}
