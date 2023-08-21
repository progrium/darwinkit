// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/corelocation"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [LocationSortDescriptor] class.
var LocationSortDescriptorClass = _LocationSortDescriptorClass{objc.GetClass("CKLocationSortDescriptor")}

type _LocationSortDescriptorClass struct {
	objc.Class
}

// An interface definition for the [LocationSortDescriptor] class.
type ILocationSortDescriptor interface {
	foundation.ISortDescriptor
	RelativeLocation() corelocation.Location
}

// An object for sorting records that contain location data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cklocationsortdescriptor?language=objc
type LocationSortDescriptor struct {
	foundation.SortDescriptor
}

func LocationSortDescriptorFrom(ptr unsafe.Pointer) LocationSortDescriptor {
	return LocationSortDescriptor{
		SortDescriptor: foundation.SortDescriptorFrom(ptr),
	}
}

func (l_ LocationSortDescriptor) InitWithKeyRelativeLocation(key string, relativeLocation corelocation.ILocation) LocationSortDescriptor {
	rv := objc.Call[LocationSortDescriptor](l_, objc.Sel("initWithKey:relativeLocation:"), key, objc.Ptr(relativeLocation))
	return rv
}

// Creates a location sort descriptor using the specified key and relative location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cklocationsortdescriptor/1515071-initwithkey?language=objc
func NewLocationSortDescriptorWithKeyRelativeLocation(key string, relativeLocation corelocation.ILocation) LocationSortDescriptor {
	instance := LocationSortDescriptorClass.Alloc().InitWithKeyRelativeLocation(key, relativeLocation)
	instance.Autorelease()
	return instance
}

func (lc _LocationSortDescriptorClass) Alloc() LocationSortDescriptor {
	rv := objc.Call[LocationSortDescriptor](lc, objc.Sel("alloc"))
	return rv
}

func LocationSortDescriptor_Alloc() LocationSortDescriptor {
	return LocationSortDescriptorClass.Alloc()
}

func (lc _LocationSortDescriptorClass) New() LocationSortDescriptor {
	rv := objc.Call[LocationSortDescriptor](lc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewLocationSortDescriptor() LocationSortDescriptor {
	return LocationSortDescriptorClass.New()
}

func (l_ LocationSortDescriptor) Init() LocationSortDescriptor {
	rv := objc.Call[LocationSortDescriptor](l_, objc.Sel("init"))
	return rv
}

func (lc _LocationSortDescriptorClass) SortDescriptorWithKeyAscendingComparator(key string, ascending bool, cmptr foundation.Comparator) LocationSortDescriptor {
	rv := objc.Call[LocationSortDescriptor](lc, objc.Sel("sortDescriptorWithKey:ascending:comparator:"), key, ascending, cmptr)
	return rv
}

// Creates and returns a sort descriptor initialized with the specified key path and ordering, and a comparator block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssortdescriptor/1503734-sortdescriptorwithkey?language=objc
func LocationSortDescriptor_SortDescriptorWithKeyAscendingComparator(key string, ascending bool, cmptr foundation.Comparator) LocationSortDescriptor {
	return LocationSortDescriptorClass.SortDescriptorWithKeyAscendingComparator(key, ascending, cmptr)
}

func (l_ LocationSortDescriptor) InitWithKeyAscending(key string, ascending bool) LocationSortDescriptor {
	rv := objc.Call[LocationSortDescriptor](l_, objc.Sel("initWithKey:ascending:"), key, ascending)
	return rv
}

// Creates a sort descriptor with a specified string key path and sort order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssortdescriptor/1413572-initwithkey?language=objc
func NewLocationSortDescriptorWithKeyAscending(key string, ascending bool) LocationSortDescriptor {
	instance := LocationSortDescriptorClass.Alloc().InitWithKeyAscending(key, ascending)
	instance.Autorelease()
	return instance
}

// The reference location for sorting records. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cklocationsortdescriptor/1514915-relativelocation?language=objc
func (l_ LocationSortDescriptor) RelativeLocation() corelocation.Location {
	rv := objc.Call[corelocation.Location](l_, objc.Sel("relativeLocation"))
	return rv
}
