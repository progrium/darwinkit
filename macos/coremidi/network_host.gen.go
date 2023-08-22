// AUTO-GENERATED CODE, DO NOT MODIFY

package coremidi

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NetworkHost] class.
var NetworkHostClass = _NetworkHostClass{objc.GetClass("MIDINetworkHost")}

type _NetworkHostClass struct {
	objc.Class
}

// An interface definition for the [NetworkHost] class.
type INetworkHost interface {
	objc.IObject
	HasSameAddressAs(other INetworkHost) bool
	Name() string
	NetServiceName() string
	Address() string
	NetServiceDomain() string
	Port() uint
}

// An object that represents the host’s network address. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworkhost?language=objc
type NetworkHost struct {
	objc.Object
}

func NetworkHostFrom(ptr unsafe.Pointer) NetworkHost {
	return NetworkHost{
		Object: objc.ObjectFrom(ptr),
	}
}

func (nc _NetworkHostClass) HostWithNameNetService(name string, netService foundation.INetService) NetworkHost {
	rv := objc.Call[NetworkHost](nc, objc.Sel("hostWithName:netService:"), name, objc.Ptr(netService))
	return rv
}

// Creates a host with the specified name and net service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworkhost/1619371-hostwithname?language=objc
func NetworkHost_HostWithNameNetService(name string, netService foundation.INetService) NetworkHost {
	return NetworkHostClass.HostWithNameNetService(name, netService)
}

func (nc _NetworkHostClass) Alloc() NetworkHost {
	rv := objc.Call[NetworkHost](nc, objc.Sel("alloc"))
	return rv
}

func NetworkHost_Alloc() NetworkHost {
	return NetworkHostClass.Alloc()
}

func (nc _NetworkHostClass) New() NetworkHost {
	rv := objc.Call[NetworkHost](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNetworkHost() NetworkHost {
	return NetworkHostClass.New()
}

func (n_ NetworkHost) Init() NetworkHost {
	rv := objc.Call[NetworkHost](n_, objc.Sel("init"))
	return rv
}

// Compares this host instance with another to see if they share the same address value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworkhost/1619372-hassameaddressas?language=objc
func (n_ NetworkHost) HasSameAddressAs(other INetworkHost) bool {
	rv := objc.Call[bool](n_, objc.Sel("hasSameAddressAs:"), objc.Ptr(other))
	return rv
}

// The host name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworkhost/1619339-name?language=objc
func (n_ NetworkHost) Name() string {
	rv := objc.Call[string](n_, objc.Sel("name"))
	return rv
}

// The net service name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworkhost/1619362-netservicename?language=objc
func (n_ NetworkHost) NetServiceName() string {
	rv := objc.Call[string](n_, objc.Sel("netServiceName"))
	return rv
}

// The host address. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworkhost/1619353-address?language=objc
func (n_ NetworkHost) Address() string {
	rv := objc.Call[string](n_, objc.Sel("address"))
	return rv
}

// The net service domain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworkhost/1619344-netservicedomain?language=objc
func (n_ NetworkHost) NetServiceDomain() string {
	rv := objc.Call[string](n_, objc.Sel("netServiceDomain"))
	return rv
}

// The host port. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworkhost/1619333-port?language=objc
func (n_ NetworkHost) Port() uint {
	rv := objc.Call[uint](n_, objc.Sel("port"))
	return rv
}
