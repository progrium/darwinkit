// AUTO-GENERATED CODE, DO NOT MODIFY

package coremidi

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NetworkSession] class.
var NetworkSessionClass = _NetworkSessionClass{objc.GetClass("MIDINetworkSession")}

type _NetworkSessionClass struct {
	objc.Class
}

// An interface definition for the [NetworkSession] class.
type INetworkSession interface {
	objc.IObject
	Connections() foundation.Set
	AddConnection(connection INetworkConnection) bool
	RemoveConnection(connection INetworkConnection) bool
	RemoveContact(contact INetworkHost) bool
	DestinationEndpoint() EndpointRef
	AddContact(contact INetworkHost) bool
	SourceEndpoint() EndpointRef
	Contacts() foundation.Set
	LocalName() string
	NetworkName() string
	ConnectionPolicy() NetworkConnectionPolicy
	SetConnectionPolicy(value NetworkConnectionPolicy)
	NetworkPort() uint
	IsEnabled() bool
	SetEnabled(value bool)
}

// An object that represents a pairing of a source and destination. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworksession?language=objc
type NetworkSession struct {
	objc.Object
}

func NetworkSessionFrom(ptr unsafe.Pointer) NetworkSession {
	return NetworkSession{
		Object: objc.ObjectFrom(ptr),
	}
}

func (nc _NetworkSessionClass) Alloc() NetworkSession {
	rv := objc.Call[NetworkSession](nc, objc.Sel("alloc"))
	return rv
}

func NetworkSession_Alloc() NetworkSession {
	return NetworkSessionClass.Alloc()
}

func (nc _NetworkSessionClass) New() NetworkSession {
	rv := objc.Call[NetworkSession](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNetworkSession() NetworkSession {
	return NetworkSessionClass.New()
}

func (n_ NetworkSession) Init() NetworkSession {
	rv := objc.Call[NetworkSession](n_, objc.Sel("init"))
	return rv
}

// Returns the session’s set of MIDI network connections. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworksession/1619366-connections?language=objc
func (n_ NetworkSession) Connections() foundation.Set {
	rv := objc.Call[foundation.Set](n_, objc.Sel("connections"))
	return rv
}

// Adds a new connection to this session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworksession/1619369-addconnection?language=objc
func (n_ NetworkSession) AddConnection(connection INetworkConnection) bool {
	rv := objc.Call[bool](n_, objc.Sel("addConnection:"), objc.Ptr(connection))
	return rv
}

// Returns the default singleton session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworksession/1619363-defaultsession?language=objc
func (nc _NetworkSessionClass) DefaultSession() NetworkSession {
	rv := objc.Call[NetworkSession](nc, objc.Sel("defaultSession"))
	return rv
}

// Returns the default singleton session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworksession/1619363-defaultsession?language=objc
func NetworkSession_DefaultSession() NetworkSession {
	return NetworkSessionClass.DefaultSession()
}

// Removes a connection from this session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworksession/1619346-removeconnection?language=objc
func (n_ NetworkSession) RemoveConnection(connection INetworkConnection) bool {
	rv := objc.Call[bool](n_, objc.Sel("removeConnection:"), objc.Ptr(connection))
	return rv
}

// Removes a host as a contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworksession/1619332-removecontact?language=objc
func (n_ NetworkSession) RemoveContact(contact INetworkHost) bool {
	rv := objc.Call[bool](n_, objc.Sel("removeContact:"), objc.Ptr(contact))
	return rv
}

// Returns the session’s destination endpoint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworksession/1619367-destinationendpoint?language=objc
func (n_ NetworkSession) DestinationEndpoint() EndpointRef {
	rv := objc.Call[EndpointRef](n_, objc.Sel("destinationEndpoint"))
	return rv
}

// Adds a host as a contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworksession/1619361-addcontact?language=objc
func (n_ NetworkSession) AddContact(contact INetworkHost) bool {
	rv := objc.Call[bool](n_, objc.Sel("addContact:"), objc.Ptr(contact))
	return rv
}

// Returns the session’s source endpoint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworksession/1619359-sourceendpoint?language=objc
func (n_ NetworkSession) SourceEndpoint() EndpointRef {
	rv := objc.Call[EndpointRef](n_, objc.Sel("sourceEndpoint"))
	return rv
}

// Returns the array of network hosts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworksession/1619335-contacts?language=objc
func (n_ NetworkSession) Contacts() foundation.Set {
	rv := objc.Call[foundation.Set](n_, objc.Sel("contacts"))
	return rv
}

// The name of this session’s entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworksession/1619358-localname?language=objc
func (n_ NetworkSession) LocalName() string {
	rv := objc.Call[string](n_, objc.Sel("localName"))
	return rv
}

// The name with which this session advertises itself over Bonjour. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworksession/1619336-networkname?language=objc
func (n_ NetworkSession) NetworkName() string {
	rv := objc.Call[string](n_, objc.Sel("networkName"))
	return rv
}

// The policy that determines who can connect to this session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworksession/1619360-connectionpolicy?language=objc
func (n_ NetworkSession) ConnectionPolicy() NetworkConnectionPolicy {
	rv := objc.Call[NetworkConnectionPolicy](n_, objc.Sel("connectionPolicy"))
	return rv
}

// The policy that determines who can connect to this session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworksession/1619360-connectionpolicy?language=objc
func (n_ NetworkSession) SetConnectionPolicy(value NetworkConnectionPolicy) {
	objc.Call[objc.Void](n_, objc.Sel("setConnectionPolicy:"), value)
}

// The session’s UDP port. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworksession/1619373-networkport?language=objc
func (n_ NetworkSession) NetworkPort() uint {
	rv := objc.Call[uint](n_, objc.Sel("networkPort"))
	return rv
}

// A Boolean value that determines whether the session is enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworksession/1619368-enabled?language=objc
func (n_ NetworkSession) IsEnabled() bool {
	rv := objc.Call[bool](n_, objc.Sel("isEnabled"))
	return rv
}

// A Boolean value that determines whether the session is enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworksession/1619368-enabled?language=objc
func (n_ NetworkSession) SetEnabled(value bool) {
	objc.Call[objc.Void](n_, objc.Sel("setEnabled:"), value)
}
