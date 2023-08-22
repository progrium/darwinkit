// AUTO-GENERATED CODE, DO NOT MODIFY

package coremidi

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CIDiscoveryManager] class.
var CIDiscoveryManagerClass = _CIDiscoveryManagerClass{objc.GetClass("MIDICIDiscoveryManager")}

type _CIDiscoveryManagerClass struct {
	objc.Class
}

// An interface definition for the [CIDiscoveryManager] class.
type ICIDiscoveryManager interface {
	objc.IObject
	DiscoverWithHandler(completedHandler CIDiscoveryResponseBlock)
}

// A singleton object that performs systemwide MIDI-CI discovery. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicidiscoverymanager?language=objc
type CIDiscoveryManager struct {
	objc.Object
}

func CIDiscoveryManagerFrom(ptr unsafe.Pointer) CIDiscoveryManager {
	return CIDiscoveryManager{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CIDiscoveryManagerClass) Alloc() CIDiscoveryManager {
	rv := objc.Call[CIDiscoveryManager](cc, objc.Sel("alloc"))
	return rv
}

func CIDiscoveryManager_Alloc() CIDiscoveryManager {
	return CIDiscoveryManagerClass.Alloc()
}

func (cc _CIDiscoveryManagerClass) New() CIDiscoveryManager {
	rv := objc.Call[CIDiscoveryManager](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCIDiscoveryManager() CIDiscoveryManager {
	return CIDiscoveryManagerClass.New()
}

func (c_ CIDiscoveryManager) Init() CIDiscoveryManager {
	rv := objc.Call[CIDiscoveryManager](c_, objc.Sel("init"))
	return rv
}

// Returns the singleton discovery manager instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicidiscoverymanager/3553242-sharedinstance?language=objc
func (cc _CIDiscoveryManagerClass) SharedInstance() CIDiscoveryManager {
	rv := objc.Call[CIDiscoveryManager](cc, objc.Sel("sharedInstance"))
	return rv
}

// Returns the singleton discovery manager instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicidiscoverymanager/3553242-sharedinstance?language=objc
func CIDiscoveryManager_SharedInstance() CIDiscoveryManager {
	return CIDiscoveryManagerClass.SharedInstance()
}

// Discovers the available MIDI-CI nodes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicidiscoverymanager/3553241-discoverwithhandler?language=objc
func (c_ CIDiscoveryManager) DiscoverWithHandler(completedHandler CIDiscoveryResponseBlock) {
	objc.Call[objc.Void](c_, objc.Sel("discoverWithHandler:"), completedHandler)
}
