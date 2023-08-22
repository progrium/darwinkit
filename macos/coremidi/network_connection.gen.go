// AUTO-GENERATED CODE, DO NOT MODIFY

package coremidi

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NetworkConnection] class.
var NetworkConnectionClass = _NetworkConnectionClass{objc.GetClass("MIDINetworkConnection")}

type _NetworkConnectionClass struct {
	objc.Class
}

// An interface definition for the [NetworkConnection] class.
type INetworkConnection interface {
	objc.IObject
	Host() NetworkHost
}

// An object that connects a session to a host. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworkconnection?language=objc
type NetworkConnection struct {
	objc.Object
}

func NetworkConnectionFrom(ptr unsafe.Pointer) NetworkConnection {
	return NetworkConnection{
		Object: objc.ObjectFrom(ptr),
	}
}

func (nc _NetworkConnectionClass) ConnectionWithHost(host INetworkHost) NetworkConnection {
	rv := objc.Call[NetworkConnection](nc, objc.Sel("connectionWithHost:"), objc.Ptr(host))
	return rv
}

// Creates a connection to the specified host. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworkconnection/1619340-connectionwithhost?language=objc
func NetworkConnection_ConnectionWithHost(host INetworkHost) NetworkConnection {
	return NetworkConnectionClass.ConnectionWithHost(host)
}

func (nc _NetworkConnectionClass) Alloc() NetworkConnection {
	rv := objc.Call[NetworkConnection](nc, objc.Sel("alloc"))
	return rv
}

func NetworkConnection_Alloc() NetworkConnection {
	return NetworkConnectionClass.Alloc()
}

func (nc _NetworkConnectionClass) New() NetworkConnection {
	rv := objc.Call[NetworkConnection](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNetworkConnection() NetworkConnection {
	return NetworkConnectionClass.New()
}

func (n_ NetworkConnection) Init() NetworkConnection {
	rv := objc.Call[NetworkConnection](n_, objc.Sel("init"))
	return rv
}

// The host connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworkconnection/1619334-host?language=objc
func (n_ NetworkConnection) Host() NetworkHost {
	rv := objc.Call[NetworkHost](n_, objc.Sel("host"))
	return rv
}
