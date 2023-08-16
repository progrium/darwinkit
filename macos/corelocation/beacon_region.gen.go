// AUTO-GENERATED CODE, DO NOT MODIFY

package corelocation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [BeaconRegion] class.
var BeaconRegionClass = _BeaconRegionClass{objc.GetClass("CLBeaconRegion")}

type _BeaconRegionClass struct {
	objc.Class
}

// An interface definition for the [BeaconRegion] class.
type IBeaconRegion interface {
	IRegion
}

// A region for detecting the presence of iBeacon devices. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clbeaconregion?language=objc
type BeaconRegion struct {
	Region
}

func BeaconRegionFrom(ptr unsafe.Pointer) BeaconRegion {
	return BeaconRegion{
		Region: RegionFrom(ptr),
	}
}

func (bc _BeaconRegionClass) Alloc() BeaconRegion {
	rv := objc.Call[BeaconRegion](bc, objc.Sel("alloc"))
	return rv
}

func BeaconRegion_Alloc() BeaconRegion {
	return BeaconRegionClass.Alloc()
}

func (bc _BeaconRegionClass) New() BeaconRegion {
	rv := objc.Call[BeaconRegion](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewBeaconRegion() BeaconRegion {
	return BeaconRegionClass.New()
}

func (b_ BeaconRegion) Init() BeaconRegion {
	rv := objc.Call[BeaconRegion](b_, objc.Sel("init"))
	return rv
}
