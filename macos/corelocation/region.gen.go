// AUTO-GENERATED CODE, DO NOT MODIFY

package corelocation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Region] class.
var RegionClass = _RegionClass{objc.GetClass("CLRegion")}

type _RegionClass struct {
	objc.Class
}

// An interface definition for the [Region] class.
type IRegion interface {
	objc.IObject
	NotifyOnEntry() bool
	SetNotifyOnEntry(value bool)
	NotifyOnExit() bool
	SetNotifyOnExit(value bool)
	Identifier() string
}

// A base class representing an area that can be monitored. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clregion?language=objc
type Region struct {
	objc.Object
}

func RegionFrom(ptr unsafe.Pointer) Region {
	return Region{
		Object: objc.ObjectFrom(ptr),
	}
}

func (rc _RegionClass) Alloc() Region {
	rv := objc.Call[Region](rc, objc.Sel("alloc"))
	return rv
}

func Region_Alloc() Region {
	return RegionClass.Alloc()
}

func (rc _RegionClass) New() Region {
	rv := objc.Call[Region](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRegion() Region {
	return RegionClass.New()
}

func (r_ Region) Init() Region {
	rv := objc.Call[Region](r_, objc.Sel("init"))
	return rv
}

// A Boolean indicating that notifications are generated upon entry into the region. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clregion/1423566-notifyonentry?language=objc
func (r_ Region) NotifyOnEntry() bool {
	rv := objc.Call[bool](r_, objc.Sel("notifyOnEntry"))
	return rv
}

// A Boolean indicating that notifications are generated upon entry into the region. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clregion/1423566-notifyonentry?language=objc
func (r_ Region) SetNotifyOnEntry(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setNotifyOnEntry:"), value)
}

// A Boolean indicating that notifications are generated upon exit from the region. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clregion/1423595-notifyonexit?language=objc
func (r_ Region) NotifyOnExit() bool {
	rv := objc.Call[bool](r_, objc.Sel("notifyOnExit"))
	return rv
}

// A Boolean indicating that notifications are generated upon exit from the region. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clregion/1423595-notifyonexit?language=objc
func (r_ Region) SetNotifyOnExit(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setNotifyOnExit:"), value)
}

// The identifier for the region object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clregion/1423583-identifier?language=objc
func (r_ Region) Identifier() string {
	rv := objc.Call[string](r_, objc.Sel("identifier"))
	return rv
}
